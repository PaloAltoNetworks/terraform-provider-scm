# Basic forwarding profile source application
resource "scm_forwarding_profile_source_application" "example_basic" {
  folder = "Mobile Users"

  name = "example-source-app"
  applications = [
    "ssl",
    "web-browsing"
  ]
}

# Forwarding profile source application with description
resource "scm_forwarding_profile_source_application" "example_with_desc" {
  folder = "Mobile Users"

  name        = "example-source-app-detailed"
  description = "Managed by Terraform - Source application profile for mobile users"
  applications = [
    "ssl",
    "web-browsing",
    "dns",
    "ntp"
  ]
}

# Forwarding profile source application with multiple applications
resource "scm_forwarding_profile_source_application" "example_multi_apps" {
  folder = "Mobile Users"

  name        = "corp-mobile-apps"
  description = "Corporate mobile applications profile"
  applications = [
    "ssl",
    "web-browsing",
    "dns",
    "ssh",
    "ping",
    "office365-consumer-access",
    "google-base"
  ]
}
