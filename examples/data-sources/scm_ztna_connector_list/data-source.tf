# List all ZTNA Connectors in the tenant.
data "scm_ztna_connector_list" "all" {}

output "connectors" {
  description = "A map of all connectors keyed by oid."
  value       = { for c in data.scm_ztna_connector_list.all.data : c.oid => c }
}

output "connector_list_total" {
  description = "Total number of connectors."
  value       = data.scm_ztna_connector_list.all.total
}
