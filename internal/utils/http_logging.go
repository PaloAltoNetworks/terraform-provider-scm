package utils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LoggingRoundTripper intercepts an HTTP request, logs it using tflog,
// executes it, and then logs the full HTTP response.
type LoggingRoundTripper struct {
	Wrapped http.RoundTripper
	Ctx     context.Context
}

func (lrt *LoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// Log the request as a curl command.
	var cmd []string
	cmd = append(cmd, "curl", "-X", req.Method)
	for name, values := range req.Header {
		for _, value := range values {
			if strings.EqualFold(name, "Authorization") {
				value = "Bearer-TOKEN-REDACTED"
			}
			if strings.EqualFold(name, "X-Auth-Jwt") {
				value = "Bearer-TOKEN-REDACTED"
			}
			cmd = append(cmd, fmt.Sprintf("-H '%s: %s'", name, value))
		}
	}
	if req.Body != nil {
		bodyBytes, _ := io.ReadAll(req.Body)
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		if len(bodyBytes) > 0 {
			cmd = append(cmd, fmt.Sprintf("-d '%s'", string(bodyBytes)))
		}
	}
	cmd = append(cmd, fmt.Sprintf("'%s'", req.URL.String()))

	// Use tflog.Debug to log the request.
	tflog.Debug(lrt.Ctx, "Terraform Provider HTTP REQUEST", map[string]interface{}{
		"curl_command": strings.Join(cmd, " "),
	})

	// Execute the actual request.
	resp, err := lrt.Wrapped.RoundTrip(req)
	if err != nil {
		return resp, err
	}

	// Log the response.
	respBody := "empty"
	if resp.Body != nil {
		bodyBytes, _ := io.ReadAll(resp.Body)
		resp.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		if len(bodyBytes) > 0 {
			respBody = string(bodyBytes)
		}
	}

	// Use tflog.Debug to log the response.
	tflog.Debug(lrt.Ctx, "Terraform Provider HTTP RESPONSE", map[string]interface{}{
		"status": resp.Status,
		"body":   respBody,
	})

	return resp, err
}
