# Look up a ZTNA wildcard rule by its ID.
data "scm_ztna_wildcard" "example" {
  oid = "55555555-5555-5555-5555-555555555555"
}

output "wildcard_output" {
  value = {
    oid         = data.scm_ztna_wildcard.example.oid
    name        = data.scm_ztna_wildcard.example.name
    description = data.scm_ztna_wildcard.example.description
    fqdn        = data.scm_ztna_wildcard.example.fqdn
    group       = data.scm_ztna_wildcard.example.group
    app_enabled = data.scm_ztna_wildcard.example.app_enabled
  }
}
