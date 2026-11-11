# List all ZTNA FQDN applications in the tenant.
data "scm_ztna_fqdn_application_list" "all" {}

output "fqdn_applications" {
  description = "A map of all FQDN applications keyed by oid."
  value       = { for app in data.scm_ztna_fqdn_application_list.all.data : app.oid => app }
}

output "fqdn_application_list_total" {
  description = "Total number of FQDN applications."
  value       = data.scm_ztna_fqdn_application_list.all.total
}
