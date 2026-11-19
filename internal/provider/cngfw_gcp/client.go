package provider

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	setup "github.com/paloaltonetworks/scm-go"
	"github.com/paloaltonetworks/scm-go/generated/cloud_ngfw_gcp"

	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// maxErrorBodyLen caps how much of a non-JSON error body ends up in a diagnostic.
const maxErrorBodyLen = 512

// Client calls the Cloud NGFW for GCP link APIs through the scm-go cloud_ngfw_gcp client.
type Client struct {
	link *cloud_ngfw_gcp.LinkAPIService
}

// NewClient builds a Client from the configured SCM setup client.
//
// It does not use setup.GetCloudNgfwGcpAPIClient. That client sends requests through the
// provider's debug logger and the SDK's "detailed" logger, which both dump response bodies,
// and a generate response carries a live link code that must never reach a log. The
// transport chain is rebuilt from the base transport instead.
func NewClient(ctx context.Context, sc *setup.Client) *Client {
	return &Client{link: cloud_ngfw_gcp.NewAPIClient(newConfiguration(ctx, sc)).LinkAPI}
}

func newConfiguration(ctx context.Context, sc *setup.Client) *cloud_ngfw_gcp.Configuration {
	var base http.RoundTripper = http.DefaultTransport
	if sc.Transport != nil {
		base = sc.Transport
	}
	if sc.AuthFile != "" {
		base = &utils.TokenRefreshTransport{
			Wrapped:  base,
			Client:   sc,
			AuthFile: sc.AuthFile,
			Ctx:      ctx,
		}
	}

	var timeout time.Duration
	if sc.HttpClient != nil {
		timeout = sc.HttpClient.Timeout
	}

	cfg := cloud_ngfw_gcp.NewConfiguration()
	cfg.Scheme = sc.Protocol
	cfg.Host = sc.Host
	if sc.Port != 0 {
		cfg.Host = fmt.Sprintf("%s:%d", sc.Host, sc.Port)
	}
	if sc.Agent != "" {
		cfg.UserAgent = sc.Agent
	}
	for k, v := range sc.Headers {
		cfg.AddDefaultHeader(k, v)
	}
	cfg.HTTPClient = &http.Client{
		Transport: &metadataLogTransport{wrapped: &setup.JWTRefreshTransport{Wrapped: base, SetupClient: sc}},
		Timeout:   timeout,
	}
	return cfg
}

// metadataLogTransport logs the method, URL, status and request ID of each call. It never
// logs a body.
type metadataLogTransport struct {
	wrapped http.RoundTripper
}

func (t *metadataLogTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := t.wrapped.RoundTrip(req)
	fields := map[string]interface{}{
		"method": req.Method,
		"url":    req.URL.String(),
	}
	if resp != nil {
		fields["status"] = resp.StatusCode
		fields["x-request-id"] = resp.Header.Get("x-request-id")
	}
	if err != nil {
		fields["error"] = err.Error()
	}
	tflog.Debug(req.Context(), "Cloud NGFW for GCP API call", fields)
	return resp, err
}

// APIError is a failed Cloud NGFW for GCP call. Code and Reason come from the ResponseStatus
// of the link envelope; Code is empty when the response had no envelope, e.g. a gateway 401.
type APIError struct {
	StatusCode int
	Code       string
	Reason     string
	RequestID  string
	Body       string
}

func (e *APIError) Error() string {
	var b strings.Builder
	fmt.Fprintf(&b, "[HTTP %d]", e.StatusCode)
	if e.Code != "" {
		fmt.Fprintf(&b, " %s", e.Code)
		if hint := codeHint(e.Code); hint != "" {
			fmt.Fprintf(&b, ": %s", hint)
		}
		if e.Reason != "" {
			fmt.Fprintf(&b, " (%s)", e.Reason)
		}
	} else if e.Body != "" {
		fmt.Fprintf(&b, " %s", e.Body)
	}
	if e.RequestID != "" {
		fmt.Fprintf(&b, " (x-request-id: %s)", e.RequestID)
	}
	return b.String()
}

func codeHint(code string) string {
	switch code {
	case "FIELD_MISSING":
		return "a required field is missing"
	case "FIELD_INVALID":
		return "a field has an invalid value"
	case "SCM_NOT_PROVISIONED":
		return "the TSG has no Strata Cloud Manager; provision it before linking"
	case "UNAUTHENTICATED":
		return "the request reached the service without a TSG; check the provider scope"
	case "INTERNAL_ERROR":
		return "the service failed; retry, and contact support with the request ID if it keeps failing"
	}
	return ""
}

// GenerateLinkCode mints a one-time link code that links a GCP tenant to the TSG the provider
// credentials are scoped to. region sets x-panw-region and may be empty.
func (c *Client) GenerateLinkCode(ctx context.Context, region, gcpTenantID string) (*cloud_ngfw_gcp.LinkGenerateResult, error) {
	req := c.link.GenerateLinkCode(ctx).LinkGenerateRequest(*cloud_ngfw_gcp.NewLinkGenerateRequest(gcpTenantID))
	if region != "" {
		req = req.XPanwRegion(region)
	}
	out, httpRes, err := req.Execute()
	if err != nil {
		return nil, toAPIError(httpRes, err)
	}
	if err := statusError(httpRes, out.GetResponseStatus()); err != nil {
		return nil, err
	}
	if out.GetResponse().Code == "" {
		return nil, errors.New("link generate returned no code")
	}
	return out.Response, nil
}

// Unlink retires the firewall binding made when Google redeemed code for gcpTenantID. It
// succeeds, retiring nothing, if the code was never redeemed.
func (c *Client) Unlink(ctx context.Context, region, gcpTenantID, code string) error {
	req := c.link.UnlinkTenant(ctx).LinkUnlinkRequest(*cloud_ngfw_gcp.NewLinkUnlinkRequest(code, gcpTenantID))
	if region != "" {
		req = req.XPanwRegion(region)
	}
	out, httpRes, err := req.Execute()
	if err != nil {
		return toAPIError(httpRes, err)
	}
	return statusError(httpRes, out.GetResponseStatus())
}

// toAPIError turns a failed call into an *APIError. Errors without a response, such as a
// connection failure, are returned unchanged.
func toAPIError(res *http.Response, err error) error {
	var openAPIErr *cloud_ngfw_gcp.GenericOpenAPIError
	if res == nil || !errors.As(err, &openAPIErr) {
		return err
	}
	apiErr := &APIError{StatusCode: res.StatusCode, RequestID: res.Header.Get("x-request-id")}
	if m, ok := openAPIErr.Model().(cloud_ngfw_gcp.LinkErrorResponse); ok && m.ResponseStatus != nil && m.ResponseStatus.GetCode() != "" {
		apiErr.Code = m.ResponseStatus.GetCode()
		apiErr.Reason = m.ResponseStatus.GetReason()
	} else {
		// Not the link envelope (e.g. a gateway error).
		apiErr.Body = truncate(openAPIErr.Body())
	}
	return apiErr
}

// statusError reports a 2xx response whose envelope still carries a non-zero ErrorCode.
func statusError(res *http.Response, rs cloud_ngfw_gcp.LinkResponseStatus) error {
	if rs.GetErrorCode() == 0 {
		return nil
	}
	return &APIError{
		StatusCode: res.StatusCode,
		Code:       rs.GetCode(),
		Reason:     rs.GetReason(),
		RequestID:  res.Header.Get("x-request-id"),
	}
}

func truncate(b []byte) string {
	s := strings.TrimSpace(string(b))
	if len(s) > maxErrorBodyLen {
		return s[:maxErrorBodyLen] + "..."
	}
	return s
}
