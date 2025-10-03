# This file is embedded using go:embed
provider "scm" {
  host          = ""
  auth_url      = ""
  client_id     = ""
  client_secret = ""
  scope         = ""
  logging       = ""
}

# This file is embedded using go:embed
terraform {
  required_providers {
    scm = {
      source = "paloaltonetworks/terraform-provider-scm"
    }
  }
}