
# 1. Fetch by ID (Best for direct lookup)
data "scm_security_rule" "standard_web_access_by_id" {
  # We use the computed ID from the resource
  id = "2a550f26-3e98-47d0-984f-b51e4ff367de"
}
# --- Outputs to Verify Fetched Data ---

output "fetched_standard_web_id" {
  description = "ID of the rule fetched by its ID."
  value       = data.scm_security_rule.standard_web_access_by_id.id
}

output "fetched_standard_web_name" {
  description = "Name of the rule fetched by its ID."
  value       = data.scm_security_rule.standard_web_access_by_id.name
}

output "fetched_standard_web_description" {
  description = "Description of the rule fetched by its ID."
  value       = data.scm_security_rule.standard_web_access_by_id.description
}