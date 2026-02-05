package utils

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	scm "github.com/paloaltonetworks/scm-go"
)

// TokenRefreshTransport intercepts HTTP requests and checks if the JWT token
// needs to be refreshed by re-reading the shared auth file before the token expires.
// This allows long-running Terraform operations to use tokens refreshed by external
// processes (e.g., cron jobs) instead of generating new tokens themselves.
type TokenRefreshTransport struct {
	Wrapped   http.RoundTripper
	Client    *scm.Client
	AuthFile  string
	Ctx       context.Context
	lastCheck time.Time
	mu        sync.Mutex
}

// authFileConfig represents the structure of the shared auth file
type authFileConfig struct {
	ClientID     string    `json:"client_id"`
	ClientSecret string    `json:"client_secret"`
	Host         string    `json:"host"`
	Protocol     string    `json:"protocol"`
	Scope        string    `json:"scope"`
	JWT          string    `json:"jwt,omitempty"`
	JWTExpiresAt time.Time `json:"jwt_expires_at,omitempty"`
	JWTLifetime  int64     `json:"jwt_lifetime,omitempty"`
	AuthUrl      string    `json:"auth_url"`
}

func (t *TokenRefreshTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Check if current token will expire within 3 minutes
	if t.AuthFile != "" && time.Now().After(t.Client.JwtExpiresAt.Add(-3*time.Minute)) {
		// Only read file every 60 seconds to avoid excessive file I/O
		// Once we're in the 3-minute danger zone, throttle file reads
		t.mu.Lock()
		shouldRead := time.Since(t.lastCheck) > 60*time.Second
		if shouldRead {
			t.lastCheck = time.Now()
		}
		t.mu.Unlock()

		if shouldRead {
			// Try to read fresh token from shared auth file
			if config, err := t.loadAuthFile(); err == nil {
				// Update client with fresh token from file
				t.mu.Lock()
				oldExpiry := t.Client.JwtExpiresAt
				t.Client.Jwt = config.JWT
				t.Client.JwtExpiresAt = config.JWTExpiresAt
				t.Client.JwtLifetime = config.JWTLifetime
				t.mu.Unlock()

				tflog.Info(t.Ctx, "Token refreshed from shared auth file", map[string]interface{}{
					"auth_file":      t.AuthFile,
					"old_expiry":     oldExpiry.Format(time.RFC3339),
					"new_expiry":     config.JWTExpiresAt.Format(time.RFC3339),
					"token_lifetime": config.JWTLifetime,
				})
			} else {
				tflog.Warn(t.Ctx, "Failed to read token from shared auth file, falling back to SDK auto-refresh", map[string]interface{}{
					"auth_file": t.AuthFile,
					"error":     err.Error(),
				})
			}
			// If file read fails, we'll fall back to SDK's auto-refresh logic
		}
	}

	// Execute the actual request
	return t.Wrapped.RoundTrip(req)
}

func (t *TokenRefreshTransport) loadAuthFile() (*authFileConfig, error) {
	data, err := os.ReadFile(t.AuthFile)
	if err != nil {
		return nil, err
	}

	var config authFileConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
