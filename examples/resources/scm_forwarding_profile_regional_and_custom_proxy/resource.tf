# Basic gp-and-pac proxy with a primary proxy server
# proxy_1 and proxy_2 are only available for gp-and-pac
resource "scm_forwarding_profile_regional_and_custom_proxy" "example_gp_and_pac_basic" {
  folder = "Mobile Users"

  name = "example-gp-pac-proxy-1"
  type = "gp-and-pac"

  proxy_1 = {
    fqdn     = "proxy1.example.com"
    port     = 8080
    location = "us-east"
  }
}

# gp-and-pac with primary and secondary proxy servers for high availability
resource "scm_forwarding_profile_regional_and_custom_proxy" "example_gp_and_pac_dual_proxy" {
  folder = "Mobile Users"

  name        = "example-gp-pac-dual-proxy"
  description = "Managed by Terraform - Dual proxy configuration for high availability"
  type        = "gp-and-pac"

  proxy_1 = {
    fqdn     = "proxy1.example.com"
    port     = 8080
    location = "us-east"
  }

  proxy_2 = {
    fqdn     = "proxy2.example.com"
    port     = 8080
    location = "us-west"
  }
}

# Basic ztna-agent proxy
# connectivity_preference and fallback_option are only available for ztna-agent
# At least one connectivity_preference entry must have enabled = true
resource "scm_forwarding_profile_regional_and_custom_proxy" "example_ztna_basic" {
  folder = "Mobile Users"

  name            = "example-ztna-proxy"
  description     = "Managed by Terraform - ZTNA agent proxy configuration"
  type            = "ztna-agent"
  fallback_option = "fail-open"

  connectivity_preference = [
    {
      name    = "tunnel"
      enabled = true
    },
    {
      name    = "proxy"
      enabled = false
    }
  ]
}

# Full ztna-agent with location_preference and Prisma Access locations
# location_preference is only available for ztna-agent
resource "scm_forwarding_profile_regional_and_custom_proxy" "example_ztna_full" {
  folder = "Mobile Users"

  name                = "example-ztna-full"
  description         = "Managed by Terraform - Full ZTNA agent with location and connectivity settings"
  type                = "ztna-agent"
  fallback_option     = "fail-safe"
  location_preference = "specific-pa-location"

  prisma_access_locations = [
    {
      name      = "americas"
      locations = ["us-east-1", "us-west-1"]
    },
    {
      name      = "europe"
      locations = ["europe-west"]
    },
    {
      name      = "apac"
      locations = ["asia-southeast1"]
    }
  ]

  connectivity_preference = [
    {
      name    = "tunnel"
      enabled = true
    },
    {
      name    = "adns"
      enabled = true
    },
    {
      name    = "masque"
      enabled = true
    }
  ]
}
