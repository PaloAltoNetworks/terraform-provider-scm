data "scm_qos_policy_rule" "single_rule_by_id" {
  # Requires the unique UUID of the rule
  id = "fbc3fdb2-8513-4949-be8a-f1be03b492da"
}

output "single_qos_policy_rule_dump" {
  description = "Dump of all attributes of the single rule fetched by ID."
  # Dumping the entire fetched object, similar to your request for App Override
  value = data.scm_qos_policy_rule.single_rule_by_id
}