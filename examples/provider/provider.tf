# This file is embedded using go:embed
provider "scm" {
  host          = ""
  auth_url      = ""
  client_id     = ""
  client_secret = ""
  scope         = ""
  logging       = ""
}

# OR with Auth File

# provider "scm" {
#   auth_file = "../../../secrets/scm-auth.json"
#   logging   = "debug"
#   protocol  = "https"
# }

terraform {
  required_providers {
    scm = {
      source = "paloaltonetworks/terraform-provider-scm"
    }
  }
}
