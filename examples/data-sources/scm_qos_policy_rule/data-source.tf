resource "scm_qos_policy_rule" "test_qos_policy_rule" {
  name        = "data-source-qos-test"
  description = "Rule created specifically for data source testing with DSCP/TOS."
  folder      = "All"
  position    = "pre"
  schedule    = "non-work-hours" # Assuming this schedule exists

  # Action: Set traffic to Class 1
  action = {
    class = "1"
  }

  # DSCP/TOS: Set to Expedited Forwarding (EF)
  dscp_tos = {
    codepoints = [
      {
        name = "Expedited Forwarding"
        type = {
          ef = {}
        }
      }
    ]
  }
}

data "scm_qos_policy_rule" "single_rule_by_id" {
  # Requires the unique UUID of the rule
  id = scm_qos_policy_rule.test_qos_policy_rule.id
}

output "single_qos_policy_rule_dump" {
  description = "Dump of all attributes of the single rule fetched by ID."
  # Dumping the entire fetched object, similar to your request for App Override
  value = data.scm_qos_policy_rule.single_rule_by_id
}