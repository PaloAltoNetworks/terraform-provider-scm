
data "scm_authentication_rule" "rule_data" {
  id = "1f1e08af-fe7b-4c36-882a-411101ad36d7"
}

output "fetched_rule_id" {
  description = "The UUID of the fetched authentication rule."
  value       = data.scm_authentication_rule.rule_data.id
}

output "fetched_rule_data" {
  description = "The timeout value for the fetched authentication rule."
  value       = data.scm_authentication_rule.rule_data
}