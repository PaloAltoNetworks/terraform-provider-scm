# Terraform Provider Authentication Tests

This directory contains comprehensive tests for JWT token caching functionality in the terraform-provider-scm.

## Test Files

### auth_file_test.go

Unit tests for auth file loading and JWT token caching.

**Test Coverage:**

#### Positive Test Cases ✅

1. **TestAuthFile_LoadFreshToken**
   - Loads a fresh JWT token from auth file
   - Verifies all fields (token, expiration, credentials) are correctly loaded
   - Validates token is not marked as expired

2. **TestAuthFile_RefreshExpiredToken** (Integration)
   - Tests automatic refresh of expired token
   - Requires real SCM credentials
   - Validates new token is obtained

#### Negative Test Cases ❌

1. **TestAuthFile_LoadExpiredToken**
   - Loads an expired JWT token
   - Verifies token is loaded but marked as expired
   - Tests expiration detection

2. **TestAuthFile_MissingJWT**
   - Tests auth file without JWT fields
   - Verifies client handles missing JWT gracefully

3. **TestAuthFile_InvalidJSON**
   - Tests behavior with malformed JSON
   - Expects Setup() to return an error

4. **TestAuthFile_MissingFile**
   - Tests behavior when auth file doesn't exist
   - Expects Setup() to return an error

5. **TestAuthFile_PartialConfig**
   - Tests incomplete configuration
   - Expects Setup() to return an error

## Running the Tests

### Unit Tests (Fast, No External Dependencies)

```bash
# Run all unit tests
go test -v ./auth/test/

# Run specific test
go test -v ./auth/test/ -run TestAuthFile_LoadFreshToken

# Run with short mode (skip integration tests)
go test -v -short ./auth/test/

# Run with race detector
go test -v -race ./auth/test/
```

### Integration Tests (Require Real Credentials)

```bash
# Set credentials
export SCM_CLIENT_ID="your-client-id"
export SCM_CLIENT_SECRET="your-client-secret"
export SCM_SCOPE="tsg_id:1234567890"

# Run integration tests
go test -v ./auth/test/ -run TestAuthFile_RefreshExpiredToken
```

## Test Categories

### Unit Tests
- ✅ No network access required
- ✅ No real credentials required
- ✅ Fast execution (< 1 second)
- ✅ Can run in CI/CD
- Tests: `TestAuthFile_*` (except `RefreshExpiredToken`)

### Integration Tests
- ⚠️ Requires real SCM credentials
- ⚠️ Makes real API calls
- ⚠️ Slower execution (1-5 seconds)
- Tests: `TestAuthFile_RefreshExpiredToken`

## Auth File Format

Tests validate this JSON structure:

```json
{
  "client_id": "your-client-id",
  "client_secret": "your-client-secret",
  "host": "api.strata.paloaltonetworks.com",
  "auth_url": "https://auth.apps.paloaltonetworks.com/auth/v1/oauth2/access_token",
  "protocol": "https",
  "scope": "tsg_id:1234567890",
  "logging": "quiet",
  "jwt": "eyJ0eXAi...",
  "jwt_expires_at": "2026-01-21T10:30:00Z",
  "jwt_lifetime": 900
}
```

## Expected Test Results

### Unit Tests (Without Integration)

```
=== RUN   TestAuthFile_LoadFreshToken
--- PASS: TestAuthFile_LoadFreshToken (0.00s)
=== RUN   TestAuthFile_LoadExpiredToken
--- PASS: TestAuthFile_LoadExpiredToken (0.00s)
=== RUN   TestAuthFile_MissingJWT
--- PASS: TestAuthFile_MissingJWT (0.00s)
=== RUN   TestAuthFile_InvalidJSON
--- PASS: TestAuthFile_InvalidJSON (0.00s)
=== RUN   TestAuthFile_MissingFile
--- PASS: TestAuthFile_MissingFile (0.00s)
=== RUN   TestAuthFile_PartialConfig
--- PASS: TestAuthFile_PartialConfig (0.00s)
PASS
```

### With Integration Tests

```
=== RUN   TestAuthFile_RefreshExpiredToken
--- PASS: TestAuthFile_RefreshExpiredToken (2.15s)
PASS
```

## Continuous Integration

### GitHub Actions Example

```yaml
name: Auth Tests

on: [push, pull_request]

jobs:
  unit-tests:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'

      - name: Run Unit Tests
        run: go test -v -short -race ./auth/test/

  integration-tests:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'

      - name: Run Integration Tests
        env:
          SCM_CLIENT_ID: ${{ secrets.SCM_CLIENT_ID }}
          SCM_CLIENT_SECRET: ${{ secrets.SCM_CLIENT_SECRET }}
          SCM_SCOPE: ${{ secrets.SCM_SCOPE }}
        run: go test -v ./auth/test/ -run TestAuthFile_RefreshExpiredToken
```

## Common Issues

### Test: TestAuthFile_RefreshExpiredToken Skipped

```
--- SKIP: TestAuthFile_RefreshExpiredToken (0.00s)
    auth_file_test.go:XXX: SCM_CLIENT_ID not set, skipping integration test
```

**Solution:** Set environment variables:
```bash
export SCM_CLIENT_ID="your-client-id"
export SCM_CLIENT_SECRET="your-client-secret"
export SCM_SCOPE="tsg_id:1234567890"
```

### Test: Invalid Credentials

```
--- FAIL: TestAuthFile_RefreshExpiredToken (1.5s)
    auth_file_test.go:XXX: RefreshJwt failed: authentication error
```

**Solution:** Verify your credentials are correct and have proper permissions.

### Test: Rate Limit Errors

If running tests repeatedly in quick succession, you may hit API rate limits.

**Solution:**
- Add delays between test runs
- Use test tokens with longer lifetimes
- Implement retry logic in integration tests

## Test Development Guidelines

1. **Unit tests should be fast**
   - Use temporary directories (`t.TempDir()`)
   - No network calls
   - No real credentials

2. **Mark integration tests**
   - Use `testing.Short()` to skip in short mode
   - Check for environment variables before running
   - Provide clear skip messages

3. **Clean up resources**
   - Use `t.TempDir()` for automatic cleanup
   - Don't leave test artifacts

4. **Test error cases**
   - Invalid input
   - Missing files
   - Malformed data
   - Edge cases

5. **Document expected behavior**
   - Clear test names
   - Comments explaining what's being tested
   - Expected vs actual comparisons
