# Basic forwarding profile destination with FQDN
resource "scm_forwarding_profile_destination" "example_basic_fqdn" {
  folder = "Mobile Users"

  name = "example-dest-fqdn"
  fqdn = [
    {
      name = "example.com"
      port = 443
    }
  ]
}

# Forwarding profile destination with IP addresses
resource "scm_forwarding_profile_destination" "example_ip_addresses" {
  folder = "Mobile Users"

  name        = "example-dest-ip"
  description = "Managed by Terraform - Destination profile with IP addresses"
  ip_addresses = [
    {
      name = "10.0.0.0/8"
      port = 443
    },
    {
      name = "192.168.1.100"
      port = 8080
    }
  ]
}

# Forwarding profile destination with both FQDN and IP addresses
resource "scm_forwarding_profile_destination" "example_mixed" {
  folder = "Mobile Users"

  name        = "corp-destinations"
  description = "Corporate destinations profile with mixed types"
  fqdn = [
    {
      name = "*.corporate.com"
      port = 443
    },
    {
      name = "api.example.com"
      port = 8443
    }
  ]
  ip_addresses = [
    {
      name = "10.10.0.0/16"
      port = 443
    },
    {
      name = "172.16.0.1"
      port = 80
    }
  ]
}
