
# 1. RESOURCE: Create an Application Override rule to ensure a predictable target for lookups
resource "scm_app_override_rule" "test_app_override_rule" {
  name        = "data-source-app-override-test"
  description = "Rule created specifically for data source testing."
  folder      = "All"
  position    = "pre"

  # Core Match Criteria
  application = "ssl"
  protocol    = "tcp"
  port        = "8443"
  from        = ["trust"]
  to          = ["untrust"]
  source      = ["any"]
  destination = ["any"]
}

data "scm_app_override_rule" "single_rule_by_id" {
  # Requires the unique UUID of the rule
  id = scm_app_override_rule.test_app_override_rule.id
}

output "single_app_override_rule_name" {
  description = "The name of the single rule fetched by ID."
  value       = data.scm_app_override_rule.single_rule_by_id
}
