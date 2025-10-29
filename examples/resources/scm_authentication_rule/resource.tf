resource "scm_tag" "app_access_tag" {
  folder = "All"
  name   = "app-access-test_25"
  color  = "Blue"
}

# -----------------------------------------------------------------------------
# 2. ANCHOR RULE (Used for relative positioning by other rules)
# -----------------------------------------------------------------------------

resource "scm_authentication_rule" "anchor_rule" {
  name        = "test_anchor_rule_251"
  description = "Base rule. Used to test 'before' and 'after' positioning"
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
  timeout                    = 1200
  negate_source              = false
  negate_destination         = false
  tag                        = [scm_tag.app_access_tag.name]
  category                   = ["any"]
  destination_hip            = ["any"]
  log_authentication_timeout = false
  disabled                   = false


}

# # -----------------------------------------------------------------------------
# # 3. ABSOLUTE POSITIONING Examples ("top" and "bottom")
# # -----------------------------------------------------------------------------

resource "scm_authentication_rule" "rule_top_of_list" {
  name        = "test_top_rule_25"
  description = "Placed at the very top of the 'pre' rulebase."
  folder      = "All"
  position    = "pre"

  # POSITIONING: Sets rule at the absolute beginning of the rulebase
  relative_position = "top"

  # Core fields
  destination = ["any"]
  from        = ["untrust"]
  to          = ["trust"]
  source      = ["any"]
  service     = ["any"]
  source_user = ["any"]
}


resource "scm_authentication_rule" "rule_bottom_of_list" {
  name        = "test_bottom_rule_25"
  description = "Placed at the very bottom of the 'pre' rulebase."
  folder      = "All"
  position    = "pre"

  # POSITIONING: Sets rule at the absolute end of the rulebase
  relative_position = "bottom"

  # Core fields
  destination = ["any"]
  from        = ["any"]
  to          = ["any"]
  source      = ["any"]
  service     = ["any"]
  source_user = ["any"]
}

# -----------------------------------------------------------------------------
# 4. RELATIVE POSITIONING Examples ("before" and "after")
# -----------------------------------------------------------------------------

resource "scm_authentication_rule" "rule_before_anchor" {
  name        = "test_before_rule_25_updating"
  description = "Positioned immediately BEFORE the anchor_rule."
  folder      = "All"
  position    = "pre"

  # POSITIONING: Requires both relative_position and target_rule
  relative_position = "before"
  target_rule       = scm_authentication_rule.anchor_rule.id # Reference anchor by its ID

  # Core fields
  destination = ["any"]
  from        = ["any"]
  to          = ["any"]
  source      = ["any"]
  service     = ["any"]
  source_user = ["any"]
}

resource "scm_authentication_rule" "rule_after_anchor" {
  name        = "test_after_rule_25"
  description = "Positioned immediately AFTER the anchor_rule."
  folder      = "All"
  position    = "pre"

  # POSITIONING: Requires both relative_position and target_rule
  relative_position = "after"
  target_rule       = scm_authentication_rule.anchor_rule.id # Reference anchor by its ID

  # Core fields
  destination = ["any"]
  from        = ["any"]
  to          = ["any"]
  source      = ["any"]
  service     = ["any"]
  source_user = ["any"] # Example specific user
}
