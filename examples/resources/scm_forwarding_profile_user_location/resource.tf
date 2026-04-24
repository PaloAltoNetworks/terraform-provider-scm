# This file is embedded using go:embed

# IP addresses only
resource "scm_forwarding_profile_user_location" "scm_user_location_1" {
  folder      = "Mobile Users"
  name        = "tf_user_location_1"
  description = "Made by Terraform"
  ip_addresses = [
    "10.0.0.0/8",
    "192.168.1.0/24"
  ]
}

# Internal host detection only
resource "scm_forwarding_profile_user_location" "scm_user_location_2" {
  folder      = "Mobile Users"
  name        = "tf_user_location_2"
  description = "Internal host detection test"
  internal_host_detection = {
    ip_address = "10.10.0.1"
    fqdn       = "internal.corp.example.com"
  }
}

# Multiple IP addresses with wildcard
resource "scm_forwarding_profile_user_location" "scm_user_location_4" {
  folder      = "Mobile Users"
  name        = "tf_user_location_4"
  description = "Multiple IP ranges for office locations"
  ip_addresses = [
    "10.0.0.0/8",
    "172.16.0.0/12",
    "192.168.0.0/16",
    "203.0.113.0/24"
  ]
}