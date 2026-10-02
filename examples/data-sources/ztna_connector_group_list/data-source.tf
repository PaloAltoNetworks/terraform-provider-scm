# List all ZTNA Connector Groups in the tenant.
data "ztna_connector_group_list" "all" {}

output "connector_groups" {
  description = "A map of all connector groups keyed by oid."
  value       = { for cg in data.ztna_connector_group_list.all.data : cg.oid => cg }
}

output "connector_group_list_total" {
  description = "Total number of connector groups."
  value       = data.ztna_connector_group_list.all.total
}
