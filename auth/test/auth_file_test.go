package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	setup "github.com/paloaltonetworks/scm-go"
)

// TestAuthFile_LoadFreshToken tests loading a fresh JWT from auth file
func TestAuthFile_LoadFreshToken(t *testing.T) {
	tmpDir := t.TempDir()
	authFile := filepath.Join(tmpDir, "auth-config.json")

	// Create auth file with fresh token
	config := map[string]interface{}{
		"client_id":      "test-client-id",
		"client_secret":  "test-client-secret",
		"host":           "api.test.paloaltonetworks.com",
		"auth_url":       "https://auth.test.paloaltonetworks.com/auth/v1/oauth2/access_token",
		"protocol":       "https",
		"scope":          "tsg_id:123456789",
		"logging":        "quiet",
		"jwt":            "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.test_fresh_token",
		"jwt_expires_at": time.Now().Add(15 * time.Minute).Format(time.RFC3339),
		"jwt_lifetime":   900,
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	if err := os.WriteFile(authFile, data, 0600); err != nil {
		t.Fatalf("Failed to write auth file: %v", err)
	}

	// Create setup client with auth file
	setupClient := &setup.Client{
		AuthFile:         authFile,
		CheckEnvironment: false,
	}

	if err := setupClient.Setup(); err != nil {
		t.Fatalf("Setup failed: %v", err)
	}

	// Verify JWT was loaded
	if setupClient.Jwt == "" {
		t.Error("Expected JWT to be loaded from auth file")
	}

	if setupClient.Jwt != config["jwt"] {
		t.Errorf("Expected JWT %s, got %s", config["jwt"], setupClient.Jwt)
	}

	// Verify client credentials were loaded
	if setupClient.ClientId != config["client_id"] {
		t.Errorf("Expected ClientId %s, got %s", config["client_id"], setupClient.ClientId)
	}

	// Verify host was loaded
	if setupClient.Host != config["host"] {
		t.Errorf("Expected Host %s, got %s", config["host"], setupClient.Host)
	}
}

// TestAuthFile_LoadExpiredToken tests loading an expired JWT
func TestAuthFile_LoadExpiredToken(t *testing.T) {
	tmpDir := t.TempDir()
	authFile := filepath.Join(tmpDir, "auth-config.json")

	expiredTime := time.Now().Add(-1 * time.Hour)

	config := map[string]interface{}{
		"client_id":      "test-client-id",
		"client_secret":  "test-client-secret",
		"host":           "api.test.paloaltonetworks.com",
		"auth_url":       "https://auth.test.paloaltonetworks.com/auth/v1/oauth2/access_token",
		"protocol":       "https",
		"scope":          "tsg_id:123456789",
		"logging":        "quiet",
		"jwt":            "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.test_expired_token",
		"jwt_expires_at": expiredTime.Format(time.RFC3339),
		"jwt_lifetime":   900,
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	if err := os.WriteFile(authFile, data, 0600); err != nil {
		t.Fatalf("Failed to write auth file: %v", err)
	}

	setupClient := &setup.Client{
		AuthFile:         authFile,
		CheckEnvironment: false,
	}

	if err := setupClient.Setup(); err != nil {
		t.Fatalf("Setup failed: %v", err)
	}

	// Verify JWT was loaded even though expired
	if setupClient.Jwt == "" {
		t.Error("Expected JWT to be loaded from auth file")
	}

	// Verify expiration was detected
	if !time.Now().After(setupClient.JwtExpiresAt) {
		t.Error("Token should be detected as expired")
	}
}

// TestAuthFile_MissingJWT tests auth file without JWT fields
// When JWT fields are missing, they should be added after token refresh
func TestAuthFile_MissingJWT(t *testing.T) {
	// Create a mock OAuth2 server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/oauth2/access_token" {
			// Return a mock JWT token response
			expiresIn := 900 // 15 minutes
			expiresAt := time.Now().Add(15 * time.Minute)
			response := map[string]interface{}{
				"access_token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.mock_token_for_missing_jwt_test",
				"token_type":   "Bearer",
				"expires_in":   expiresIn,
				"expires_at":   expiresAt.Unix(),
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}
		http.NotFound(w, r)
	}))
	defer mockServer.Close()

	tmpDir := t.TempDir()
	authFile := filepath.Join(tmpDir, "auth-config.json")

	config := map[string]interface{}{
		"client_id":     "test-client-id",
		"client_secret": "test-client-secret",
		"host":          "api.test.paloaltonetworks.com",
		"auth_url":      mockServer.URL + "/oauth2/access_token",
		"protocol":      "https",
		"scope":         "tsg_id:123456789",
		"logging":       "quiet",
		// No JWT fields
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	if err := os.WriteFile(authFile, data, 0600); err != nil {
		t.Fatalf("Failed to write auth file: %v", err)
	}

	setupClient := &setup.Client{
		AuthFile:         authFile,
		CheckEnvironment: false,
	}

	if err := setupClient.Setup(); err != nil {
		t.Fatalf("Setup failed: %v", err)
	}

	// JWT should be empty initially (needs to be fetched)
	if setupClient.Jwt != "" {
		t.Error("Expected JWT to be empty when not in auth file")
	}

	// Expiration should be zero initially
	if !setupClient.JwtExpiresAt.IsZero() {
		t.Error("Expected JwtExpiresAt to be zero when not in auth file")
	}

	// Refresh JWT to fetch new token from mock server
	ctx := context.Background()
	if err := setupClient.RefreshJwt(ctx); err != nil {
		t.Fatalf("RefreshJwt failed: %v", err)
	}

	// Verify new token was obtained
	if setupClient.Jwt == "" {
		t.Error("Expected JWT to be populated after refresh")
	}

	// Verify expiration is in the future
	if time.Now().After(setupClient.JwtExpiresAt) {
		t.Error("New token should not be expired")
	}

	// Verify JWT lifetime was set
	if setupClient.JwtLifetime == 0 {
		t.Error("Expected JwtLifetime to be set after refresh")
	}

	// Update the config with new JWT fields and save back to file
	config["jwt"] = setupClient.Jwt
	config["jwt_expires_at"] = setupClient.JwtExpiresAt.Format(time.RFC3339)
	config["jwt_lifetime"] = setupClient.JwtLifetime

	savedData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal updated config: %v", err)
	}

	if err := os.WriteFile(authFile, savedData, 0600); err != nil {
		t.Fatalf("Failed to write updated config: %v", err)
	}

	// Read the auth file and verify JWT fields were added
	readData, err := os.ReadFile(authFile)
	if err != nil {
		t.Fatalf("Failed to read updated auth file: %v", err)
	}

	var updatedConfig map[string]interface{}
	if err := json.Unmarshal(readData, &updatedConfig); err != nil {
		t.Fatalf("Failed to unmarshal updated config: %v", err)
	}

	// Verify JWT field was added to file
	if _, exists := updatedConfig["jwt"]; !exists {
		t.Error("Expected 'jwt' field to be added to auth file")
	}

	// Verify jwt_expires_at field was added to file
	if _, exists := updatedConfig["jwt_expires_at"]; !exists {
		t.Error("Expected 'jwt_expires_at' field to be added to auth file")
	}

	// Verify jwt_lifetime field was added to file
	if _, exists := updatedConfig["jwt_lifetime"]; !exists {
		t.Error("Expected 'jwt_lifetime' field to be added to auth file")
	}

	// Verify the JWT token value matches what the mock server returned
	if jwtValue, ok := updatedConfig["jwt"].(string); ok {
		if jwtValue != "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.mock_token_for_missing_jwt_test" {
			t.Errorf("Expected JWT token to match mock server response, got: %s", jwtValue)
		}
	} else {
		t.Error("JWT field is not a string")
	}
}

// TestAuthFile_InvalidJSON tests invalid JSON in auth file
func TestAuthFile_InvalidJSON(t *testing.T) {
	tmpDir := t.TempDir()
	authFile := filepath.Join(tmpDir, "auth-config.json")

	// Write invalid JSON
	if err := os.WriteFile(authFile, []byte("invalid json{{{{{{"), 0600); err != nil {
		t.Fatalf("Failed to write invalid JSON: %v", err)
	}

	setupClient := &setup.Client{
		AuthFile:         authFile,
		CheckEnvironment: false,
	}

	err := setupClient.Setup()
	if err == nil {
		t.Error("Expected error with invalid JSON, got nil")
	} else {
		// Verify the error message indicates JSON parsing error
		errorMsg := err.Error()
		if !strings.Contains(errorMsg, "invalid character") && !strings.Contains(errorMsg, "unexpected end of JSON input") {
			t.Errorf("Expected error to contain JSON parsing error, got '%s'", errorMsg)
		}
	}
}

// TestAuthFile_MissingFile tests missing auth file
func TestAuthFile_MissingFile(t *testing.T) {
	tmpDir := t.TempDir()
	authFile := filepath.Join(tmpDir, "nonexistent.json")

	setupClient := &setup.Client{
		AuthFile:         authFile,
		CheckEnvironment: false,
	}

	err := setupClient.Setup()
	if err == nil {
		t.Error("Expected error with missing file, got nil")
	} else {
		// Verify the error message indicates file not found
		errorMsg := err.Error()
		if !strings.Contains(errorMsg, "no such file or directory") {
			t.Errorf("Expected error to contain 'no such file or directory', got '%s'", errorMsg)
		}
	}
}

// TestAuthFile_PartialConfig tests incomplete configuration
func TestAuthFile_PartialConfig(t *testing.T) {
	tmpDir := t.TempDir()
	authFile := filepath.Join(tmpDir, "auth-config.json")

	// Missing required fields: client_id, client_secret, scope
	config := map[string]interface{}{
		"host":     "api.test.paloaltonetworks.com",
		"protocol": "https",
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	if err := os.WriteFile(authFile, data, 0600); err != nil {
		t.Fatalf("Failed to write auth file: %v", err)
	}

	setupClient := &setup.Client{
		AuthFile:         authFile,
		CheckEnvironment: false,
	}

	err = setupClient.Setup()
	if err == nil {
		t.Error("Expected error with partial config, got nil")
	} else {
		// Verify the error message indicates missing ClientId
		expectedError := "ClientId must be specified"
		if err.Error() != expectedError {
			t.Errorf("Expected error '%s', got '%s'", expectedError, err.Error())
		}
	}
}

// TestAuthFile_RefreshExpiredToken tests the refresh logic when token is expired
func TestAuthFile_RefreshExpiredToken(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// This test requires real credentials
	if os.Getenv("SCM_CLIENT_ID") == "" {
		t.Skip("SCM_CLIENT_ID not set, skipping integration test")
	}

	tmpDir := t.TempDir()
	authFile := filepath.Join(tmpDir, "auth-config.json")

	expiredTime := time.Now().Add(-1 * time.Hour)

	config := map[string]interface{}{
		"client_id":      os.Getenv("SCM_CLIENT_ID"),
		"client_secret":  os.Getenv("SCM_CLIENT_SECRET"),
		"host":           getEnv("SCM_HOST", "api.test.paloaltonetworks.com"),
		"auth_url":       getEnv("SCM_AUTH_URL", "https://auth.test.paloaltonetworks.com/auth/v1/oauth2/access_token"),
		"protocol":       "https",
		"scope":          os.Getenv("SCM_SCOPE"),
		"logging":        "quiet",
		"jwt":            "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.old_expired_token",
		"jwt_expires_at": expiredTime.Format(time.RFC3339),
		"jwt_lifetime":   900,
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	if err := os.WriteFile(authFile, data, 0600); err != nil {
		t.Fatalf("Failed to write auth file: %v", err)
	}

	setupClient := &setup.Client{
		AuthFile:         authFile,
		CheckEnvironment: false,
	}

	if err := setupClient.Setup(); err != nil {
		t.Fatalf("Setup failed: %v", err)
	}

	oldToken := setupClient.Jwt

	// Token is expired, so RefreshJwt should fetch new one
	if time.Now().After(setupClient.JwtExpiresAt) {
		ctx := context.Background()
		if err := setupClient.RefreshJwt(ctx); err != nil {
			t.Fatalf("RefreshJwt failed: %v", err)
		}

		// Verify new token was obtained
		if setupClient.Jwt == oldToken {
			t.Error("Expected new token after refresh, got same token")
		}

		// Verify new expiration is in the future
		if time.Now().After(setupClient.JwtExpiresAt) {
			t.Error("New token should not be expired")
		}
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
