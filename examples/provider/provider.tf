provider "scm" {
  host          = "api.strata.paloaltonetworks.com"
  client_id     = "your-id@12345"
  client_secret = "secret"
  scope         = "tsg_id:12345"
}

terraform {
  required_providers {
    scm = {
      source  = "paloaltonetworks/terraform-provider-scm"
      version = "0.1.0"
    }
  }
}
