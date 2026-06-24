# List all ZTNA wildcard rules in the tenant.
data "ztna_wildcard_list" "all" {}

output "wildcards" {
  description = "A map of all wildcard rules keyed by oid."
  value       = { for w in data.ztna_wildcard_list.all.data : w.oid => w }
}

output "wildcard_list_total" {
  description = "Total number of wildcard rules."
  value       = data.ztna_wildcard_list.all.total
}
