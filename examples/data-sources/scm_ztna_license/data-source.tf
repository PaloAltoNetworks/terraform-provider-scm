# Look up ZTNA tenant license information.
data "scm_ztna_license" "current" {}

output "license_output" {
  value = {
    license_name     = data.scm_ztna_license.current.license_name
    expiry           = data.scm_ztna_license.current.expiry
    connectors       = data.scm_ztna_license.current.connectors
    max_connectors   = data.scm_ztna_license.current.max_connectors
    applications     = data.scm_ztna_license.current.applications
    max_applications = data.scm_ztna_license.current.max_applications
  }
}
