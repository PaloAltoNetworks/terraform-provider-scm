# This provider serves several API families from a single binary:
#   scm_address, ...      — Strata Cloud Manager resources (NGFW & Prisma Access)
#   scm_ztna_*            — ZTNA Connector resources (Zero Trust Network Access)
#
# SCM and ZTNA use different API hosts:
#   scm_* (non-ZTNA) resources use the "host" field      (e.g. api.strata.paloaltonetworks.com)
#   scm_ztna_* resources use the "ztna_host" field (e.g. api.sase.paloaltonetworks.com)
# Both "host" and "ztna_host" can be specified together in the same provider block
# when configuring credentials inline. If using an auth file, both fields can be
# included in the same JSON file.
#
# ZTNA resources also require the "x_panw_region" field, which sets the
# x-panw-region HTTP header on all ZTNA API requests. Valid values:
#   ae, americas, au, br, ca, ch, cn, de, es, europe, fr, id, il, in, it,
#   jp, kr, pl, qa, sa, sg, tw, uk, za
# This can be set via the provider block, the X_PANW_REGION environment variable,
# or the "x_panw_region" key in the JSON auth file.
#
# Every resource, data source and action starts with scm_ (scm_address,
# scm_ztna_subnet, scm_cngfw_ngfirewall, ...), so one provider block configures
# all of them. ZTNA types need the ztna_host and x_panw_region settings:
#
# terraform {
#   required_providers {
#     scm = {
#       source = "PaloAltoNetworks/scm"
#     }
#   }
# }
#
# provider "scm" {
#   auth_file     = "/path/to/scm-config.json"
#   ztna_host     = "api.sase.paloaltonetworks.com"
#   x_panw_region = "americas"   # see valid values above
#   logging       = "debug"
# }



provider "scm" {
  host          = ""
  auth_url      = ""
  client_id     = ""
  client_secret = ""
  scope         = ""
  logging       = "debug"
}

# This file is embedded using go:embed
terraform {
  required_providers {
    scm = {
      source = "paloaltonetworks-local/scm"
    }
  }
}