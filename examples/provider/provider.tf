# This provider serves two resource prefixes from a single binary:
#   scm_*  — Strata Cloud Manager resources (NGFW & Prisma Access)
#   ztna_* — ZTNA Connector resources (Zero Trust Network Access)
#
# SCM and ZTNA use different API hosts:
#   scm_*  resources use the "host" field      (e.g. api.strata.paloaltonetworks.com)
#   ztna_* resources use the "ztna_host" field (e.g. api.sase.paloaltonetworks.com)
# Both "host" and "ztna_host" can be specified together in the same provider block
# when configuring credentials inline. If using an auth file, both fields can be
# included in the same JSON file.
#
# ZTNA resources also require the "x_panw_region" field, which sets the
# x-panw-region HTTP header on all ZTNA API requests. Valid values: americas, europe, apac.
# This can be set via the provider block, the X_PANW_REGION environment variable,
# or the "x_panw_region" key in the JSON auth file.
#
# To use both, declare two provider aliases in required_providers pointing to
# the same source, and configure each with its own provider block.
# The ztna_host and x_panw_region fields in the auth file or provider block are
# used for ZTNA resources.
#
# Example with both providers:
#
# terraform {
#   required_providers {
#     scm = {
#       source  = "PaloAltoNetworks/scm"
#     }
#     ztna = {
#       source  = "PaloAltoNetworks/scm"
#     }
#   }
# }
#
# provider "scm" {
#   auth_file = "/path/to/scm-config.json"
#   logging   = "debug"
# }
#
# provider "ztna" {
#   auth_file     = "/path/to/scm-config.json"
#   ztna_host     = "api.sase.paloaltonetworks.com"
#   x_panw_region = "americas"   # or "europe" / "apac"
#   logging       = "debug"
# }

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
