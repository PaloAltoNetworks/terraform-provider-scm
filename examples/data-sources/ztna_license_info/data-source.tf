# Look up ZTNA tenant license information.
data "ztna_license_info" "current" {}

output "license_info_output" {
  value = {
    license_name     = data.ztna_license_info.current.license_name
    expiry           = data.ztna_license_info.current.expiry
    connectors       = data.ztna_license_info.current.connectors
    max_connectors   = data.ztna_license_info.current.max_connectors
    applications     = data.ztna_license_info.current.applications
    max_applications = data.ztna_license_info.current.max_applications
  }
}
