# --- 2. ANCHOR QOS POLICY RULE (Used for relative positioning) ---
resource "scm_qos_policy_rule" "anchor_qos_rule" {
  name        = "anchor-qos-rule"
  description = "Base rule for testing 'before' and 'after' positioning."
  folder      = "All"
  position    = "pre"

  # Action: Set the traffic to QoS Class B
  action = {
    class = "2"
  }

  # Optional fields
  schedule = "non-work-hours" # Assuming a schedule object exists

  # DSCP Marking: Set traffic to Expedited Forwarding (EF)
  dscp_tos = {
    codepoints = [
      {
        name = "Set-EF" # Friendly name for the codepoint
        type = {
          ef = {} # Using the 'ef' (Expedited Forwarding) type, which is an empty object
        }
      }
    ]
  }
}

# --- 3. ABSOLUTE POSITIONING Examples ("top" and "bottom") ---

resource "scm_qos_policy_rule" "rule_top_qos_rule" {
  name        = "top-absolute-qos-rule"
  description = "Placed at the very TOP of the QoS rulebase (Highest Priority)."
  folder      = "All"
  position    = "pre"

  # POSITIONING: Sets rule at the absolute beginning of the rulebase
  relative_position = "top"

  # Action: Set the traffic to QoS Class 2
  action = {
    class = "2"
  }
}


resource "scm_qos_policy_rule" "rule_bottom_qos_rule" {
  name        = "bottom-absolute-qos-rule"
  description = "Placed at the very BOTTOM of the QoS rulebase (Lowest Priority)"
  folder      = "All"
  position    = "pre"

  # POSITIONING: Sets rule at the absolute end of the rulebase
  relative_position = "bottom"

  # Action: Set the traffic to QoS Class 3
  action = {
    class = "3"
  }
}

# --- 4. RELATIVE POSITIONING Examples ("before" and "after") ---

resource "scm_qos_policy_rule" "rule_before_anchor_qos" {
  name        = "before-anchor-qos-rule"
  description = "Positioned immediately BEFORE the anchor-qos-rule."
  folder      = "All"
  position    = "pre"

  # POSITIONING: Requires both relative_position and target_rule
  relative_position = "before"
  target_rule       = scm_qos_policy_rule.anchor_qos_rule.id # Reference anchor by its ID

  # Action: Set the traffic to QoS Class 3
  action = {
    class = "5"
  }


}

resource "scm_qos_policy_rule" "rule_after_anchor_qos" {
  name        = "after-anchor-qos-rule"
  description = "Positioned immediately AFTER the anchor-qos-rule."
  folder      = "All"
  position    = "pre"

  # POSITIONING: Requires both relative_position and target_rule
  relative_position = "after"
  target_rule       = scm_qos_policy_rule.anchor_qos_rule.id # Reference anchor by its ID

  # Action: Set the traffic to QoS Class 4
  action = {
    class = "4"
  }
}
