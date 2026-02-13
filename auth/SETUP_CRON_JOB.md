# SCM Token Cache Service

Automated JWT token refresh service for Palo Alto Networks Strata Cloud Manager (SCM) API authentication. This service runs as a cron job to keep your JWT tokens fresh, preventing authentication failures in long-running integrations.

## Overview

The SCM API requires JWT tokens for authentication, which expire after a set period. This service:
- Checks token expiration status every 5 minutes
- Automatically refreshes tokens when they will expire within 5 minutes
- Caches tokens in a JSON configuration file
- Uses atomic file writes to prevent corruption
- Logs all operations for monitoring and debugging

## Prerequisites

- Go 1.16 or higher (for building the binary)
- Access to SCM API credentials (Client ID, Client Secret)
- Unix-based system with cron (Linux, macOS)
- Read/write permissions for the config and log directories

## File Structure

```
auth/
├── token_cache.go              # Main Go source code
├── token-cache-service         # Compiled binary (generated)
├── run-token-cache.sh          # Wrapper script for cron execution
└── logs/                       # Log directory (auto-created)
    └── token-cache-YYYY-MM-DD.log

secrets/
└── scm-auth.json              # Configuration and cached token
```

## Setup Instructions

### Step 1: Prepare Configuration File

Create a JSON configuration file at `../secrets/scm-auth.json` with your SCM credentials:

```json
{
  "client_id": "your-client-id",
  "client_secret": "your-client-secret",
  "host": "api.sase.paloaltonetworks.com",
  "protocol": "https",
  "scope": "tsg_id:your-tsg-id",
  "auth_url": "https://auth.apps.paloaltonetworks.com/oauth2/access_token"
}
```

**Important:** Ensure this file has restricted permissions:
```bash
chmod 600 ../secrets/scm-auth.json
```

### Step 2: Create the Token Cache Service

Create a file named `token_cache.go` in your `auth` directory. This Go program handles the JWT token caching and refresh logic.

**Note:** The source code for `token_cache.go` will be provided separately. (Main README.md Section: Example 1: Go Token Cache Service)

### Step 3: Build the Binary

Navigate to the `auth` directory and compile the Go code using this command:

```bash
cd /path/to/your/project/auth
go build -o token-cache-service token_cache.go
```

Verify the binary was created:
```bash
ls -lh token-cache-service
```

You should see an executable file (approximately 8-10 MB).

### Step 4: Configure the Wrapper Script

The `run-token-cache.sh` script is already provided. Update the paths if your installation differs:

```bash
#!/bin/bash

# Set up logging
LOG_DIR="/path/to/your/project/auth/logs"
mkdir -p "$LOG_DIR"
LOG_FILE="$LOG_DIR/token-cache-$(date +\%Y-\%m-\%d).log"

# Run the service and append to log
echo "=== $(date '+%Y-%m-%d %H:%M:%S') ===" >> "$LOG_FILE"
/path/to/your/project/auth/token-cache-service >> "$LOG_FILE" 2>&1
echo "" >> "$LOG_FILE"
```

Make it executable:
```bash
chmod +x run-token-cache.sh
```

### Step 5: Set Up Cron Job

Add the cron job to run every 5 minutes:

**Option A - Automatic:**
```bash
(crontab -l 2>/dev/null; echo "*/5 * * * * /path/to/your/project/auth/run-token-cache.sh") | crontab -
```

**Option B - Manual:**
```bash
crontab -e
```

Add this line:
```
*/5 * * * * /path/to/your/project/auth/run-token-cache.sh
```

Save and exit.

### Step 6: Verify Installation

Check that the cron job was added:
```bash
crontab -l
```

You should see:
```
*/5 * * * * /path/to/your/project/auth/run-token-cache.sh
```

## Verification

### Test Manual Execution

Before relying on cron, test the service manually:

```bash
cd /path/to/your/project/auth
./token-cache-service
```

Expected output (first run):
```
Token expired or missing, fetching new token...
Token refreshed, expires at: 2026-01-23 12:34:56 +0000 UTC
```

Expected output (subsequent runs before expiration):
```
Using cached token, expires at: 2026-01-23 12:34:56 +0000 UTC
```

### Monitor Logs

After the cron job runs, check the logs:

```bash
tail -f /path/to/your/project/auth/logs/token-cache-$(date +\%Y-\%m-\%d).log
```

You should see entries like:
```
=== 2026-01-23 10:00:00 ===
Using cached token, expires at: 2026-01-23 12:34:56 +0000 UTC

=== 2026-01-23 10:02:00 ===
Using cached token, expires at: 2026-01-23 12:34:56 +0000 UTC
```

### Verify Token File

Check that the token is being cached:

```bash
cat ../secrets/scm-auth.json
```

You should see the JWT fields populated:
```json
{
  "client_id": "your-client-id",
  "client_secret": "your-client-secret",
  "host": "api.sase.paloaltonetworks.com",
  "protocol": "https",
  "scope": "tsg_id:your-tsg-id",
  "auth_url": "https://auth.apps.paloaltonetworks.com/oauth2/access_token",
  "jwt": "eyJhbGc...",
  "jwt_expires_at": "2026-01-23T12:34:56Z",
  "jwt_lifetime": 3600
}
```

## Configuration Options

### Custom Config Path

Set a custom path using the `SCM_CONFIG_PATH` environment variable:

```bash
export SCM_CONFIG_PATH="/custom/path/to/scm-auth.json"
./token-cache-service
```

Or in cron:
```
*/5 * * * * SCM_CONFIG_PATH="/custom/path/to/config.json" /path/to/token-cache-service
```

### Adjust Refresh Interval

To change how often the cron job runs, modify the cron schedule:

- Every 1 minute: `*/1 * * * *`
- Every 2 minutes: `*/2 * * * *`
- Every 5 minutes: `*/5 * * * *` (recommended for 15-minute token lifetime)
- Every 10 minutes: `*/10 * * * *`

**Note:** The service has a built-in 5-minute expiration buffer. With tokens expiring in 15 minutes, the recommended 5-minute cron interval ensures tokens are always refreshed before expiration. Running every 5 minutes means the service checks at consistent intervals (0, 5, 10, 15, 20, etc.) and will always catch tokens within the 5-minute expiration window, preventing any request failures.

## Troubleshooting

### Cron Job Not Running

1. **Check cron service is running:**
   ```bash
   sudo systemctl status cron  # Linux
   # or
   sudo launchctl list | grep cron  # macOS
   ```

2. **Check cron logs:**
   ```bash
   grep CRON /var/log/syslog  # Linux
   # or
   tail -f /var/log/system.log | grep cron  # macOS
   ```

3. **Verify script permissions:**
   ```bash
   ls -l run-token-cache.sh token-cache-service
   ```
   Both should be executable (`-rwxr-xr-x`).

### Authentication Failures

If you see "Failed to refresh JWT" errors:

1. **Verify credentials in config file:**
   ```bash
   cat ../secrets/scm-auth.json
   ```

2. **Test SCM API connectivity:**
   ```bash
   curl -X POST https://auth.apps.paloaltonetworks.com/oauth2/access_token \
     -H "Content-Type: application/json" \
     -d '{
       "client_id": "your-client-id",
       "client_secret": "your-client-secret",
       "grant_type": "client_credentials",
       "scope": "tsg_id:your-tsg-id"
     }'
   ```

3. **Check network connectivity and firewall rules**

### Permission Errors

If you see "permission denied" errors:

```bash
chmod 600 ../secrets/scm-auth.json
chmod 755 auth/logs
chmod +x run-token-cache.sh token-cache-service
```

### Log File Growth

Logs are organized by date (`token-cache-YYYY-MM-DD.log`). To prevent excessive disk usage, set up log rotation:

Create `/etc/logrotate.d/scm-token-cache`:
```
/path/to/your/project/auth/logs/token-cache-*.log {
    daily
    rotate 7
    compress
    missingok
    notifempty
}
```

## Stopping the Service

To stop the cron job:

```bash
crontab -e
```

Comment out or delete the line:
```
# */5 * * * * /path/to/your/project/auth/run-token-cache.sh
```

Or remove it entirely:
```bash
crontab -l | grep -v "run-token-cache.sh" | crontab -
```

## Security Considerations

1. **Protect your credentials:**
   - Keep `scm-auth.json` with restrictive permissions (`chmod 600`)
   - Never commit credentials to version control
   - Add `secrets/` to `.gitignore`

2. **Secure the token cache:**
   - The JWT token has the same privileges as your client credentials
   - Ensure only authorized users can read the config file
   - Consider encrypting the config file at rest

3. **Monitor logs:**
   - Regularly check logs for unauthorized access attempts
   - Set up alerts for authentication failures

## Support

For issues or questions:
- Check the troubleshooting section above
- Review logs in `auth/logs/`
- Verify your SCM credentials and API access
- Ensure network connectivity to SCM API endpoints

## License

This service is provided as-is for use with Palo Alto Networks Strata Cloud Manager.
