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

# This file is embedded using go:embed
# provider "scm" {
#   auth_file = "../../../secrets/scm-auth.json"
#   logging   = "debug"
#   protocol = "https"
# }

# This file is embedded using go:embed
terraform {
  required_providers {
    scm = {
      source = "paloaltonetworks/terraform-provider-scm"
    }
  }
}