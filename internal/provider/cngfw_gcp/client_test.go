package provider

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	setup "github.com/paloaltonetworks/scm-go"

	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

func newTestSetupClient(t *testing.T, srv *httptest.Server) *setup.Client {
	t.Helper()
	u, err := url.Parse(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	port, err := strconv.Atoi(u.Port())
	if err != nil {
		t.Fatal(err)
	}
	return &setup.Client{
		Protocol:     "http",
		Host:         u.Hostname(),
		Port:         port,
		Agent:        "test-agent",
		Headers:      map[string]string{"X-Custom": "custom"},
		Jwt:          "test-jwt",
		JwtExpiresAt: time.Now().Add(time.Hour),
		Transport:    &http.Transport{},
	}
}

func newTestClient(t *testing.T, h http.HandlerFunc) *Client {
	t.Helper()
	srv := httptest.NewServer(h)
	t.Cleanup(srv.Close)
	return NewClient(context.Background(), newTestSetupClient(t, srv))
}

func writeJSON(t *testing.T, w http.ResponseWriter, status int, v any) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("x-request-id", "req-123")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		t.Error(err)
	}
}

func TestGenerateLinkCode(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/cngfw-gcp/v2/link/generate" {
			t.Errorf("got %s %s, want POST /cngfw-gcp/v2/link/generate", r.Method, r.URL.Path)
		}
		for h, want := range map[string]string{
			"Authorization": "Bearer test-jwt",
			"Content-Type":  "application/json",
			"User-Agent":    "test-agent",
			"X-Custom":      "custom",
			"X-Panw-Region": "americas",
		} {
			if got := r.Header.Get(h); got != want {
				t.Errorf("header %s = %q, want %q", h, got, want)
			}
		}
		// The service takes the TSG from the gateway and looks up everything else itself.
		b, _ := io.ReadAll(r.Body)
		if strings.TrimSpace(string(b)) != `{"GCPTenantID":"projects/my-proj"}` {
			t.Errorf("body = %s", b)
		}
		writeJSON(t, w, http.StatusOK, map[string]any{
			"Response":       map[string]any{"Code": "957RPA7PHO", "TSGID": "1778860705", "CSPID": "123", "ExpiresAt": 1792171508},
			"ResponseStatus": map[string]any{"ErrorCode": 0, "Code": "OK", "Reason": ""},
		})
	})

	got, err := c.GenerateLinkCode(context.Background(), "americas", "projects/my-proj")
	if err != nil {
		t.Fatal(err)
	}
	if got.GetCode() != "957RPA7PHO" || got.GetTSGID() != "1778860705" || got.GetCSPID() != "123" || got.GetExpiresAt() != 1792171508 {
		t.Errorf("got %+v", got)
	}
}

func TestGenerateLinkCodeOmitsRegion(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if _, ok := r.Header["X-Panw-Region"]; ok {
			t.Errorf("x-panw-region sent with empty region")
		}
		writeJSON(t, w, http.StatusOK, map[string]any{
			"Response":       map[string]any{"Code": "ABC", "ExpiresAt": 1},
			"ResponseStatus": map[string]any{"ErrorCode": 0, "Code": "OK"},
		})
	})
	if _, err := c.GenerateLinkCode(context.Background(), "", "my-proj"); err != nil {
		t.Fatal(err)
	}
}

func TestGenerateLinkCodeErrors(t *testing.T) {
	cases := []struct {
		name       string
		status     int
		body       string
		wantCode   string
		wantSubstr []string
	}{
		{
			name:       "error envelope",
			status:     http.StatusBadRequest,
			body:       `{"Response":null,"ResponseStatus":{"ErrorCode":1,"Code":"SCM_NOT_PROVISIONED","Reason":""}}`,
			wantCode:   "SCM_NOT_PROVISIONED",
			wantSubstr: []string{"[HTTP 400]", "SCM_NOT_PROVISIONED", "provision it", "x-request-id: req-123"},
		},
		{
			name:       "error envelope with reason",
			status:     http.StatusInternalServerError,
			body:       `{"Response":null,"ResponseStatus":{"ErrorCode":1,"Code":"INTERNAL_ERROR","Reason":"tsg has 2 tenants"}}`,
			wantCode:   "INTERNAL_ERROR",
			wantSubstr: []string{"[HTTP 500]", "INTERNAL_ERROR", "(tsg has 2 tenants)"},
		},
		{
			name:       "error code on 200",
			status:     http.StatusOK,
			body:       `{"Response":null,"ResponseStatus":{"ErrorCode":1,"Code":"FIELD_INVALID","Reason":""}}`,
			wantCode:   "FIELD_INVALID",
			wantSubstr: []string{"[HTTP 200]", "FIELD_INVALID"},
		},
		{
			name:       "gateway error without envelope",
			status:     http.StatusUnauthorized,
			body:       `{"_errors":[{"message":"invalid token"}]}`,
			wantSubstr: []string{"[HTTP 401]", "invalid token"},
		},
		{
			name:       "plain text body",
			status:     http.StatusBadGateway,
			body:       "upstream down",
			wantSubstr: []string{"[HTTP 502]", "upstream down"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("x-request-id", "req-123")
				w.WriteHeader(tc.status)
				_, _ = io.WriteString(w, tc.body)
			})
			_, err := c.GenerateLinkCode(context.Background(), "", "my-proj")
			var apiErr *APIError
			if !errors.As(err, &apiErr) {
				t.Fatalf("got %v, want *APIError", err)
			}
			if apiErr.StatusCode != tc.status || apiErr.Code != tc.wantCode {
				t.Errorf("got status %d code %q, want %d %q", apiErr.StatusCode, apiErr.Code, tc.status, tc.wantCode)
			}
			for _, s := range tc.wantSubstr {
				if !strings.Contains(err.Error(), s) {
					t.Errorf("error %q does not contain %q", err, s)
				}
			}
		})
	}
}

func TestGenerateLinkCodeMissingCode(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"Response":       nil,
			"ResponseStatus": map[string]any{"ErrorCode": 0, "Code": "OK"},
		})
	})
	if _, err := c.GenerateLinkCode(context.Background(), "", "my-proj"); err == nil {
		t.Fatal("expected an error for a response without a code")
	}
}

func TestUnlink(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/cngfw-gcp/v2/link/unlink" {
			t.Errorf("got %s %s", r.Method, r.URL.Path)
		}
		if got := r.Header.Get("x-panw-region"); got != "europe" {
			t.Errorf("x-panw-region = %q, want europe", got)
		}
		// Unlink is keyed by the link code; the tenant must be the one the code was generated for.
		b, _ := io.ReadAll(r.Body)
		if strings.TrimSpace(string(b)) != `{"Code":"ABCDE12345","GCPTenantID":"organizations/123"}` {
			t.Errorf("body = %s", b)
		}
		writeJSON(t, w, http.StatusOK, map[string]any{
			"Response":       map[string]any{"GCPTenantID": "organizations/123", "Unlinked": true},
			"ResponseStatus": map[string]any{"ErrorCode": 0, "Code": "OK"},
		})
	})
	if err := c.Unlink(context.Background(), "europe", "organizations/123", "ABCDE12345"); err != nil {
		t.Fatal(err)
	}
}

func TestUnlinkError(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		writeJSON(t, w, http.StatusBadRequest, map[string]any{
			"Response":       nil,
			"ResponseStatus": map[string]any{"ErrorCode": 1, "Code": "FIELD_MISSING"},
		})
	})
	err := c.Unlink(context.Background(), "", "organizations/123", "")
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.Code != "FIELD_MISSING" {
		t.Fatalf("got %v, want FIELD_MISSING", err)
	}
}

// failingRoundTripper fails the test if a request reaches it.
type failingRoundTripper struct{ t *testing.T }

func (f failingRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
	f.t.Error("request went through the provider's body-logging transport")
	return nil, errors.New("unexpected")
}

func TestNewClientBypassesBodyLoggingTransport(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"Response":       map[string]any{"Code": "SECRET", "ExpiresAt": 1},
			"ResponseStatus": map[string]any{"ErrorCode": 0, "Code": "OK"},
		})
	}))
	defer srv.Close()

	sc := newTestSetupClient(t, srv)
	// This is what the provider installs when logging = "debug".
	sc.HttpClient = &http.Client{
		Transport: &utils.LoggingRoundTripper{Wrapped: failingRoundTripper{t}, Ctx: context.Background()},
		Timeout:   5 * time.Second,
	}
	if got := newConfiguration(context.Background(), sc).HTTPClient.Timeout; got != 5*time.Second {
		t.Errorf("timeout = %s, want the provider client's timeout", got)
	}
	if _, err := NewClient(context.Background(), sc).GenerateLinkCode(context.Background(), "", "my-proj"); err != nil {
		t.Fatal(err)
	}
}

func TestNewConfiguration(t *testing.T) {
	for _, tc := range []struct {
		sc         setup.Client
		wantScheme string
		wantHost   string
	}{
		{setup.Client{Protocol: "https", Host: "api.strata.paloaltonetworks.com"}, "https", "api.strata.paloaltonetworks.com"},
		{setup.Client{Protocol: "http", Host: "localhost", Port: 8080}, "http", "localhost:8080"},
	} {
		cfg := newConfiguration(context.Background(), &tc.sc)
		if cfg.Scheme != tc.wantScheme || cfg.Host != tc.wantHost {
			t.Errorf("scheme %q host %q, want %q %q", cfg.Scheme, cfg.Host, tc.wantScheme, tc.wantHost)
		}
	}
}
