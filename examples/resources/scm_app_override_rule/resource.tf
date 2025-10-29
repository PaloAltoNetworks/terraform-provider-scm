# --- 1. TAG Resource ---
resource "scm_tag" "app_override_position_tag" {
  name   = "app-override-position-tag_1"
  folder = "All"
  color  = "Orange"
}

# --- 2. ANCHOR RULE (Used for relative positioning by other rules) ---
resource "scm_app_override_rule" "anchor_app_override" {
  name        = "anchor-app-override-rule"
  description = "Base rule for testing 'before' and 'after' positioning. Updating"
  folder      = "All"
  position    = "pre" # App Override Rules exist in the 'pre' rulebase

  # Core Match Criteria
  application = "ssl" # Reclassify the traffic as 'ssl'
  protocol    = "tcp"
  port        = "112"

  # Source and Destination Criteria
  from        = ["trust"]
  to          = ["untrust"]
  source      = ["any"]
  destination = ["any"]

  # Optional fields
  tag = [scm_tag.app_override_position_tag.name]
}

# --- 3. ABSOLUTE POSITIONING Examples ("top" and "bottom") ---

resource "scm_app_override_rule" "rule_top_app_override" {
  name        = "top-absolute-app-override"
  description = "Placed at the very TOP of the App Override rulebase."
  folder      = "All"
  position    = "pre"

  # POSITIONING: Sets rule at the absolute beginning of the rulebase
  relative_position = "bottom"

  # Core Match Criteria
  application = "ssl"
  protocol    = "tcp"
  port        = "443"
  from        = ["untrust"]
  to          = ["trust"]
  source      = ["any"]
  destination = ["any"]
}


resource "scm_app_override_rule" "rule_bottom_app_override" {
  name        = "bottom-absolute-app-override"
  description = "Placed at the very BOTTOM of the App Override rulebase."
  folder      = "All"
  position    = "pre"

  # POSITIONING: Sets rule at the absolute end of the rulebase
  relative_position = "bottom"

  # Core Match Criteria
  application = "ssl"
  protocol    = "tcp"
  port        = "443"
  from        = ["any"]
  to          = ["any"]
  source      = ["any"]
  destination = ["any"]
}

#--- 4. RELATIVE POSITIONING Examples ("before" and "after") ---

resource "scm_app_override_rule" "rule_before_anchor_override" {
  name        = "before-anchor-app-override"
  description = "Positioned immediately BEFORE the anchor-app-override-rule."
  folder      = "All"
  position    = "pre"

  # POSITIONING: Requires both relative_position and target_rule
  relative_position = "before"
  # Reference anchor by its ID (must use 'id' attribute)
  target_rule = scm_app_override_rule.anchor_app_override.id

  # Core Match Criteria
  application = "ssl"
  protocol    = "tcp"
  port        = "443"
  from        = ["trust"]
  to          = ["untrust"]
  source      = ["any"]
  destination = ["any"]
}

resource "scm_app_override_rule" "rule_after_anchor_override" {
  name        = "after-anchor-app-override"
  description = "Positioned immediately AFTER the anchor-app-override-rule."
  folder      = "All"
  position    = "pre"

  # POSITIONING: Requires both relative_position and target_rule
  relative_position = "before"
  # Reference anchor by its ID (must use 'id' attribute)
  target_rule = scm_app_override_rule.anchor_app_override.id

  # Core Match Criteria
  application = "ssl"
  protocol    = "tcp"
  port        = "443"
  from        = ["untrust"]
  to          = ["trust"]
  source      = ["any"]
  destination = ["any"]
}
