# Look up a ZTNA FQDN application by its ID.
data "scm_ztna_fqdn_application" "example" {
  oid = "33333333-3333-3333-3333-333333333333"
}

output "fqdn_application_output" {
  value = {
    oid         = data.scm_ztna_fqdn_application.example.oid
    name        = data.scm_ztna_fqdn_application.example.name
    description = data.scm_ztna_fqdn_application.example.description
    fqdn        = data.scm_ztna_fqdn_application.example.fqdn
    group       = data.scm_ztna_fqdn_application.example.group
    app_enabled = data.scm_ztna_fqdn_application.example.app_enabled
  }
}
