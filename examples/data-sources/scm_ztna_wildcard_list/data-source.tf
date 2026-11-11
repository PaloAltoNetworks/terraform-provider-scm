# List all ZTNA wildcard rules in the tenant.
data "scm_ztna_wildcard_list" "all" {}

output "wildcards" {
  description = "A map of all wildcard rules keyed by oid."
  value       = { for w in data.scm_ztna_wildcard_list.all.data : w.oid => w }
}

output "wildcard_list_total" {
  description = "Total number of wildcard rules."
  value       = data.scm_ztna_wildcard_list.all.total
}
