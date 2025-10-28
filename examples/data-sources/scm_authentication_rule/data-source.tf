
resource "scm_authentication_rule" "rule_to_fetch" {
  name        = "rule-to-be-queried-scm-105"
  description = "This rule is created purely to test the data source functionality."
  position    = "pre" # Default rulebase is 'pre'
  folder      = "All"

  # Core fields (REQUIRED by the API)
  destination = ["any"]
  from        = ["any"]
  to          = ["any"]
  source      = ["any"]
  service     = ["service-http", "service-https"]

  # Identity and Enforcement
  source_user = ["any"]

  # Optional fields
  timeout            = 1200
  negate_source      = false
  negate_destination = false
}

data "scm_authentication_rule" "rule_data" {
  # To look up the rule by name, you must provide both the 'name' and the 'folder'.
  # The 'folder' is necessary as it's part of the API scope for listing/fetching rules.

  id = scm_authentication_rule.rule_to_fetch.id

}

output "fetched_rule_id" {
  description = "The UUID of the fetched authentication rule."
  value       = data.scm_authentication_rule.rule_data.id
}

output "fetched_rule_timeout" {
  description = "The timeout value for the fetched authentication rule."
  value       = data.scm_authentication_rule.rule_data.timeout
}