# Look up ZTNA tenant license information.
data "ztna_license" "current" {}

output "license_output" {
  value = {
    license_name     = data.ztna_license.current.license_name
    expiry           = data.ztna_license.current.expiry
    connectors       = data.ztna_license.current.connectors
    max_connectors   = data.ztna_license.current.max_connectors
    applications     = data.ztna_license.current.applications
    max_applications = data.ztna_license.current.max_applications
  }
}
