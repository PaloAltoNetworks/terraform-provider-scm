# Look up ZTNA tenant status.
data "ztna_tenant_status" "current" {}

output "tenant_status" {
  value       = data.ztna_tenant_status.current.status
  description = "Current tenant status (ok, not_found, delete_in_progress)"
}
