data "scm_app_override_rule" "single_rule_by_id" {
  # Requires the unique UUID of the rule
  id = "8c285335-3c95-47c9-9bbd-829105b4a15c"
}

output "single_app_override_rule_name" {
  description = "The name of the single rule fetched by ID."
  value       = data.scm_app_override_rule.single_rule_by_id
}
