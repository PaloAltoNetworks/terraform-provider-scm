package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
			if isSensitiveHTTPName(name) {
				value = "Bearer-TOKEN-REDACTED"
			}
			cmd = append(cmd, fmt.Sprintf("-H '%s: %s'", name, value))
		}
	}
	if req.Body != nil {
		bodyBytes, _ := io.ReadAll(req.Body)
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		if len(bodyBytes) > 0 {
			cmd = append(cmd, fmt.Sprintf("-d '%s'", redactHTTPBody(bodyBytes)))
		}
	}
	cmd = append(cmd, fmt.Sprintf("'%s'", redactHTTPURL(req.URL)))

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
			respBody = redactHTTPBody(bodyBytes)
		}
	}

	// Use tflog.Debug to log the response.
	tflog.Debug(lrt.Ctx, "Terraform Provider HTTP RESPONSE", map[string]interface{}{
		"status": resp.Status,
		"body":   respBody,
	})

	return resp, err
}

// sensitiveHTTPNameFragments are case-insensitive fragments of JSON field names,
// header names and query parameter names whose values must never be logged.
var sensitiveHTTPNameFragments = []string{
	"passphrase", "password", "passwd", "secret", "token", "credential",
	"private", "key_file", "keyfile", "psk", "pre_shared", "preshared",
	"authorization", "cookie", "api_key", "apikey", "auth_key", "jwt",
}

func isSensitiveHTTPName(name string) bool {
	n := strings.ToLower(name)
	if strings.HasSuffix(n, "key") {
		return true
	}
	for _, frag := range sensitiveHTTPNameFragments {
		if strings.Contains(n, frag) {
			return true
		}
	}
	return false
}

// redactHTTPValue walks decoded JSON and replaces values of sensitive fields.
func redactHTTPValue(v interface{}) interface{} {
	switch t := v.(type) {
	case map[string]interface{}:
		for k, val := range t {
			if isSensitiveHTTPName(k) {
				t[k] = "****"
			} else {
				t[k] = redactHTTPValue(val)
			}
		}
		return t
	case []interface{}:
		for i, val := range t {
			t[i] = redactHTTPValue(val)
		}
		return t
	}
	return v
}

// redactHTTPBody returns a loggable form of an HTTP body with secret-like JSON
// fields masked. Non-JSON bodies are not logged because they cannot be
// redacted reliably.
func redactHTTPBody(body []byte) string {
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.UseNumber()
	var v interface{}
	if err := dec.Decode(&v); err != nil {
		return fmt.Sprintf("[non-JSON body omitted, %d bytes]", len(body))
	}
	out, err := json.Marshal(redactHTTPValue(v))
	if err != nil {
		return fmt.Sprintf("[body omitted, %d bytes]", len(body))
	}
	return string(out)
}

// redactHTTPURL returns the URL with secret-like query parameter values masked.
func redactHTTPURL(u *url.URL) string {
	if u == nil {
		return ""
	}
	c := *u
	c.User = nil
	q := c.Query()
	changed := false
	for k := range q {
		if isSensitiveHTTPName(k) {
			q.Set(k, "****")
			changed = true
		}
	}
	if changed {
		c.RawQuery = q.Encode()
	}
	return c.String()
}
