# 1. RESOURCE: Create a rule to ensure a predictable target for lookups
resource "scm_decryption_rule" "test_decryption_rule" {
  name        = "data-source-test-rule"
  description = "Rule created specifically for data source testing."
  folder      = "All"
  position    = "pre"
  action      = "decrypt"

  # Core Match Criteria (REQUIRED fields)
  from        = ["trust"]
  to          = ["untrust"]
  source      = ["any"]
  destination = ["any"]
  service     = ["service-https"]
  category    = ["high-risk"]
  source_user = ["any"]

  # Type definition (using assignment syntax as corrected previously)
  type = {
    ssl_forward_proxy = {}
  }
}

# We use the ID from the resource created above.
data "scm_decryption_rule" "single_rule_by_id" {
  # Requires the unique UUID of the rule
  id = scm_decryption_rule.test_decryption_rule.id
}

output "single_decryption_rule_name" {
  description = "The name of the single rule fetched by ID."
  value       = data.scm_decryption_rule.single_rule_by_id
}
