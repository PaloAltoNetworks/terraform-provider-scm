# List all ZTNA subnet rules in the tenant.
data "scm_ztna_subnet_list" "all" {}

output "subnets" {
  description = "A map of all subnet rules keyed by oid."
  value       = { for s in data.scm_ztna_subnet_list.all.data : s.oid => s }
}

output "subnet_list_total" {
  description = "Total number of subnet rules."
  value       = data.scm_ztna_subnet_list.all.total
}
