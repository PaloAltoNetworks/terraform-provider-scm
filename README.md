Terraform Provider for Palo Alto Networks Strata Cloud Manager API
==================================================================

**NOTE**: This provider is auto-generated.

---
## Beta Release Disclaimer

**This software is a pre-release version and is not ready for production use.**

*   **No Warranty:** This software is provided "as is," without any warranty of any kind, either expressed or implied, including, but not limited to, the implied warranties of merchantability and fitness for a particular purpose.
*   **Instability:** The beta software may contain defects, may not operate correctly, and may be substantially modified or withdrawn at any time.
*   **Limitation of Liability:** In no event shall the authors or copyright holders be liable for any claim, damages, or other liability, whether in an action of contract, tort, or otherwise, arising from, out of, or in connection with the beta software or the use or other dealings in the beta software.
*   **Feedback:** We encourage and appreciate your feedback and bug reports. However, you acknowledge that any feedback you provide is non-confidential.

By using this software, you agree to these terms.
---

Warranty
--------

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

THIS SOFTWARE IS RELEASED AS A PROOF OF CONCEPT FOR EXPERIMENTAL PURPOSES ONLY. USE IT AT OWN RISK. THIS SOFTWARE IS NOT SUPPORTED.

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) v1+
- [Go](https://go.dev) v1.21+ (to build the provider from source)


Building the Provider
---------------------

1. Install [Go](https://go.dev/dl)

2. Clone the SDK repo:

```sh
git clone https://github.com/paloaltonetworks/scm-go
```

3. Clone this repo:

```sh
git clone https://github.com/paloaltonetworks/terraform-provider-scm
```

4. Build the provider:

```sh
cd terraform-provider-scm
go build
```

5. Specify the `dev_overrides` configuration per the next section below. This tells Terraform where to find the provider you just built. The directory to specify is the full path to the cloned provider repo.


Developing the Provider
-----------------------

With Terraform v1 and later, [development overrides for provider developers](https://www.terraform.io/docs/cli/config/config-file.html#development-overrides-for-provider-developers) can be leveraged in order to use the provider built from source.

To do this, populate a Terraform CLI configuration file (`~/.terraformrc` for all platforms other than Windows; `terraform.rc` in the `%APPDATA%` directory when using Windows) with at least the following options:

```hcl
provider_installation {
  dev_overrides {
    "registry.terraform.io/paloaltonetworks-local/scm" = "/directory/containing/the/provider/binary/here"
  }

  direct {}
}
```

Then when referencing the locally built provider, use the local name in the `terraform` configuration block like so:

```hcl
terraform {
    required_providers {
        scm = {
            source = "paloaltonetworks-local/scm"
            version = "1.0.0"
        }
    }
}
```

Running Concurrent Terraform Operations
----------------------------------------

### The Challenge

The Strata Cloud Manager authentication API has rate limits on token requests (approximately 10 concurrent requests per tenant). When running multiple concurrent Terraform operations (e.g., parallel CI/CD pipelines, multiple workspaces, or `terraform apply` with high parallelism), you may encounter authentication failures due to these rate limits.

Each provider instance requests a new JWT token during initialization, which can quickly exhaust the rate limit.

### The Solution: JWT Token Caching

Starting from version 1.0.8, the Terraform provider supports passing pre-existing JWT tokens via the `auth_file` parameter. This allows you to implement a token caching solution that shares tokens across multiple Terraform runs.

### Configuration

You can configure the provider to use a cached token by specifying an `auth_file` that includes JWT fields:

```hcl
provider "scm" {
  auth_file = "/var/cache/scm/auth-token.json"
}
```

The auth file should contain:

```json
{
  "client_id": "your-client-id",
  "client_secret": "your-client-secret",
  "host": "api.strata.paloaltonetworks.com",
  "protocol": "https",
  "scope": "tsg_id:1234567890",
  "jwt": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9...",
  "jwt_expires_at": "2026-01-21T10:30:00Z",
  "jwt_lifetime": 900
}
```

**Important Security Note**: Only share JWT tokens among Terraform runs that use the **same** `client_id` and `client_secret`. Different service principals with different RBAC permissions must never share tokens, as this would be a privilege escalation risk.

### Implementing Token Caching

The provider does not include a built-in token caching service. Instead, you should implement your own caching mechanism that fits your infrastructure and security requirements.

#### Architecture Overview

```
┌─────────────────────────────────────────────────────────────────────────┐
│                       Token Caching Architecture                         │
└─────────────────────────────────────────────────────────────────────────┘

                         ┌──────────────────────┐
                         │   SCM Auth API       │
                         │  (Rate Limited ~10   │
                         │  concurrent requests)│
                         └──────────┬───────────┘
                                    │
                                    │ 1. Fetch JWT Token
                                    │    (Once every 10-12 min)
                                    │
                         ┌──────────▼───────────┐
                         │  Token Cache Service │
                         │   (Cron Job/Timer)   │
                         │                      │
                         │  • Checks expiration │
                         │  • Fetches new token │
                         │  • Updates auth file │
                         └──────────┬───────────┘
                                    │
                                    │ 2. Write (Atomic)
                                    │    jwt + jwt_expires_at
                                    │
                    ┌───────────────▼────────────────┐
                    │    Shared Auth File            │
                    │   /var/cache/scm/auth.json     │
                    │                                 │
                    │  {                              │
                    │    "client_id": "...",          │
                    │    "client_secret": "...",      │
                    │    "jwt": "eyJ...",             │
                    │    "jwt_expires_at": "...",     │
                    │    "jwt_lifetime": 900          │
                    │  }                              │
                    └─┬────────┬────────┬────────┬───┘
                      │        │        │        │
          3. Read JWT │        │        │        │ 3. Read JWT
          (No API     │        │        │        │ (No API
           call)      │        │        │        │  call)
                      │        │        │        │
        ┌─────────────▼──┐  ┌──▼────────▼──┐  ┌─▼─────────────┐
        │  Terraform      │  │  Terraform   │  │  Terraform    │
        │  Process #1     │  │  Process #2  │  │  Process #N   │
        │                 │  │              │  │               │
        │  terraform      │  │  terraform   │  │  terraform    │
        │  apply          │  │  apply       │  │  apply        │
        │  (Workspace A)  │  │  (Workspace B)  │  (CI/CD Job) │
        └─────────────────┘  └──────────────┘  └───────────────┘

┌─────────────────────────────────────────────────────────────────────────┐
│  Benefits:                                                               │
│  • Only 1 auth request per tenant every 10-12 minutes                   │
│  • Supports unlimited concurrent Terraform processes                     │
│  • No rate limit errors during parallel operations                      │
│  • Cached token shared safely (read-only for Terraform processes)       │
└─────────────────────────────────────────────────────────────────────────┘
```

**How It Works:**

1. **Token Cache Service** (cron job/systemd timer) runs every 10-12 minutes
   - Checks if cached token is expired or expiring soon (60s buffer)
   - Fetches new JWT token from SCM Auth API if needed
   - Writes updated token to shared auth file (atomic write operation)

2. **Shared Auth File** (`/var/cache/scm/auth.json` or similar)
   - Contains `client_id`, `client_secret`, and cached `jwt` fields
   - Updated atomically by token cache service
   - Read by all Terraform provider instances

3. **Multiple Terraform Processes** (concurrent operations)
   - Each process reads the shared auth file on startup
   - Uses cached JWT token (no API call needed)
   - Can run unlimited concurrent operations without hitting rate limits
   - All processes must use the same `client_id`/`client_secret`

**Disclaimer**: The example code below is provided "as is" without warranty. It is intended as a reference implementation only. You are responsible for ensuring it meets your organization's security and operational requirements.

#### Example 1: Go Token Cache Service

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	scm "github.com/paloaltonetworks/scm-go"
)

// Config represents the SCM configuration with JWT caching
type Config struct {
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

func main() {
	// Use absolute path for cron compatibility
	configPath := os.Getenv("SCM_CONFIG_PATH")
	if configPath == "" {
		configPath := "/path/to/scm-config.json"
	}

	// Load existing config
	config, err := loadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Check if token needs refresh (5 minute buffer)
	needsRefresh := config.JWT == "" || time.Now().After(config.JWTExpiresAt.Add(-300*time.Second))

	if needsRefresh {
		log.Println("Token expired or missing, fetching new token...")

		// Create SCM client (will fetch new token)
		client := &scm.Client{
			ClientId:         config.ClientID,
			ClientSecret:     config.ClientSecret,
			Host:             config.Host,
			Protocol:         config.Protocol,
			Scope:            config.Scope,
			AuthUrl:          config.AuthUrl,
			CheckEnvironment: false, // Don't check env vars
		}

		if err := client.Setup(); err != nil {
			log.Fatalf("Failed to setup client: %v", err)
		}

		// Get JWT token
		if err := client.RefreshJwt(context.Background()); err != nil {
			log.Fatalf("Failed to refresh JWT: %v", err)
		}

		// Update config with new token
		config.JWT = client.Jwt
		config.JWTExpiresAt = client.JwtExpiresAt
		config.JWTLifetime = client.JwtLifetime

		// Save updated config atomically
		if err := saveConfigAtomic(configPath, config); err != nil {
			log.Fatalf("Failed to save config: %v", err)
		}

		log.Printf("Token refreshed, expires at: %s\n", config.JWTExpiresAt)
	} else {
		log.Printf("Using cached token, expires at: %s\n", config.JWTExpiresAt)
	}
}

func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func saveConfigAtomic(path string, config *Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	// Write to temp file first
	tmpPath := path + ".tmp"
	if err := os.WriteFile(tmpPath, data, 0600); err != nil {
		return err
	}

	// Atomic rename
	return os.Rename(tmpPath, path)
}
```

**Usage**: Run this as a cron job or systemd timer every 10-12 minutes to keep tokens fresh.

```bash
# Run every 10 minutes
*/10 * * * * /path/to/token-cache-service
```

#### Example 2: Python Token Cache Service

```python
#!/usr/bin/env python3
"""
SCM Token Cache Service
Fetches and caches JWT tokens for concurrent SCM operations
"""

import json
import os
import sys
import tempfile
from datetime import datetime, timedelta
from pathlib import Path

from scm import Scm


def load_config(config_path: str) -> dict:
    """Load configuration from JSON file"""
    with open(config_path, 'r') as f:
        return json.load(f)


def save_config_atomic(config_path: str, config: dict) -> None:
    """Save configuration atomically with proper permissions"""
    config_dir = os.path.dirname(config_path)

    # Write to temporary file first
    with tempfile.NamedTemporaryFile(
        mode='w',
        dir=config_dir,
        delete=False,
        prefix='.scm-config-',
        suffix='.tmp'
    ) as tmp_file:
        json.dump(config, tmp_file, indent=2)
        tmp_path = tmp_file.name

    # Set restrictive permissions
    os.chmod(tmp_path, 0o600)

    # Atomic rename
    os.rename(tmp_path, config_path)


def needs_refresh(config: dict, buffer_seconds: int = 60) -> bool:
    """Check if token needs refresh"""
    if not config.get('jwt') or not config.get('jwt_expires_at'):
        return True

    try:
        expires_at = datetime.fromisoformat(
            config['jwt_expires_at'].replace('Z', '+00:00')
        )
        # Refresh if expiring within buffer period
        return datetime.now(expires_at.tzinfo) >= (
            expires_at - timedelta(seconds=buffer_seconds)
        )
    except (ValueError, TypeError):
        return True


def main():
    config_path = os.path.expanduser("~/.scm/config.json")

    # Ensure config directory exists
    Path(config_path).parent.mkdir(parents=True, exist_ok=True, mode=0o700)

    # Load config
    try:
        config = load_config(config_path)
    except FileNotFoundError:
        print(f"Error: Config file not found: {config_path}", file=sys.stderr)
        sys.exit(1)

    # Check if refresh needed
    if needs_refresh(config):
        print("Token expired or missing, fetching new token...")

        # Create SCM client (will fetch new token)
        client = Scm(
            client_id=config.get('client_id'),
            client_secret=config.get('client_secret'),
            tsg_id=config.get('tsg_id') or config.get('scope', '').replace('tsg_id:', ''),
            host=config.get('host', 'api.sase.paloaltonetworks.com'),
        )

        # Get token details
        token = client.access_token
        expires_at = client._token_expires_at
        lifetime = client._jwt_lifetime

        # Update config
        config['jwt'] = token
        config['jwt_expires_at'] = expires_at.isoformat() if expires_at else None
        config['jwt_lifetime'] = lifetime

        # Save atomically
        save_config_atomic(config_path, config)

        print(f"Token refreshed, expires at: {expires_at}")
    else:
        print(f"Using cached token, expires at: {config['jwt_expires_at']}")


if __name__ == '__main__':
    main()
```

**Usage**: Run this as a cron job every 10-12 minutes to keep tokens fresh.

```bash
# Run every 10 minutes
*/10 * * * * /usr/bin/python3 /path/to/token_cache_service.py
```

#### Example 3: Simple Bash Token Cache

```bash
#!/bin/bash
# Simple token cache updater using scm-go binary

CONFIG_FILE="${HOME}/.scm/scm-config.json"
GO_CACHE_TOOL="/path/to/scm-token-refresh"  # Build from Example 1

# Check if token needs refresh
if [ -f "$CONFIG_FILE" ]; then
    EXPIRES_AT=$(jq -r '.jwt_expires_at // empty' "$CONFIG_FILE")
    if [ -n "$EXPIRES_AT" ]; then
        EXPIRES_EPOCH=$(date -d "$EXPIRES_AT" +%s 2>/dev/null || date -j -f "%Y-%m-%dT%H:%M:%S" "${EXPIRES_AT:0:19}" +%s 2>/dev/null)
        NOW_EPOCH=$(date +%s)
        BUFFER=120  # 2 minute buffer

        if [ $((EXPIRES_EPOCH - NOW_EPOCH)) -gt $BUFFER ]; then
            echo "Token still valid, expires at: $EXPIRES_AT"
            exit 0
        fi
    fi
fi

echo "Refreshing token..."
"$GO_CACHE_TOOL"
```

### Best Practices

1. **Token Lifecycle Management**
   - Refresh tokens before they expire (recommend 2-5 minute buffer)
   - Use atomic file operations (write to temp, then rename)
   - Handle token refresh failures gracefully

2. **Security Considerations**
   - Store cache files with restrictive permissions (`chmod 600`)
   - Use separate cache files for different credential sets
   - Never commit auth files to version control
   - Rotate credentials regularly

3. **Concurrent Access**
   - Multiple processes can read the same cache file safely
   - Only one process should write (use file locking if needed)
   - Consider using a shared filesystem for distributed runners

### Troubleshooting

**Issue**: Still getting authentication rate limit errors
- Verify all Terraform runs use the same cache file
- Check that token refresh logic is working
- Ensure cache file is readable by all processes

**Issue**: Different environments need different credentials
- Use separate cache files per credential set with provider aliases

### Additional Resources

- [scm-go JWT Token Caching Documentation](https://github.com/PaloAltoNetworks/scm-go#jwt-token-caching-for-concurrent-operations)
- [GitHub Issue #77: Limited concurrent IaC operations](https://github.com/PaloAltoNetworks/terraform-provider-scm/issues/77)
- [GitHub Issue #13: Allow passing JWTs to client](https://github.com/PaloAltoNetworks/scm-go/issues/13)
