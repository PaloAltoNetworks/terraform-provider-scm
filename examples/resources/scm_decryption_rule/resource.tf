# --- 1. TAG Resource ---
resource "scm_tag" "decryption_position_tag" {
  name   = "decryption-position-tag"
  folder = "All"
  color  = "Purple"
}

# --- 2. ANCHOR DECRYPTION RULE (Used for relative positioning) ---
resource "scm_decryption_rule" "anchor_decryption_rule" {
  name        = "anchor-decryption-rule"
  description = "Base rule for testing 'before' and 'after' positioning."
  folder      = "All"
  position    = "pre" # Decryption Rules typically exist in the 'pre' rulebase
  action      = "decrypt"

  # Core Match Criteria (REQUIRED fields)
  from        = ["trust"]         # Source security zone
  to          = ["untrust"]       # Destination security zone
  source      = ["any"]           # Source addresses
  destination = ["any"]           # Destination addresses
  service     = ["service-https"] # Services (e.g., standard HTTPS port)
  category    = ["high-risk"]     # Destination URL Category
  source_user = ["any"]           # Source user/group
  type = {
    ssl_forward_proxy = {}
  }
  destination_hip = ["any"]

  # Optional fields
  tag                = [scm_tag.decryption_position_tag.name]
  log_success        = true
  log_fail           = true
  disabled           = false
  negate_source      = false
  negate_destination = false
}

# --- 3. ABSOLUTE POSITIONING Examples ("top" and "bottom") ---

resource "scm_decryption_rule" "rule_top_decryption_rule" {
  name        = "top-absolute-decryption-rule"
  description = "Placed at the very TOP of the Decryption rulebase."
  folder      = "All"
  position    = "pre"
  action      = "no-decrypt" # Use no-decrypt for top rules (e.g., skip banking sites)

  # POSITIONING: Sets rule at the absolute beginning of the rulebase
  relative_position = "top"

  # Core Match Criteria
  from        = ["any"]
  to          = ["any"]
  source      = ["any"]
  destination = ["any"]
  service     = ["service-https"]
  category    = ["high-risk"]
  source_user = ["any"]
  type = {
    ssl_forward_proxy = {}
  }
}


resource "scm_decryption_rule" "rule_bottom_decryption_rule" {
  name        = "bottom-absolute-decryption-rule"
  description = "Placed at the very BOTTOM of the Decryption rulebase."
  folder      = "All"
  position    = "pre"
  action      = "decrypt"

  # POSITIONING: Sets rule at the absolute end of the rulebase
  relative_position = "bottom"

  # Core Match Criteria
  from        = ["any"]
  to          = ["any"]
  source      = ["any"]
  destination = ["any"]
  service     = ["service-https"]
  category    = ["high-risk"]
  source_user = ["any"]
  type = {
    ssl_forward_proxy = {}
  }
}

# --- 4. RELATIVE POSITIONING Examples ("before" and "after") ---

resource "scm_decryption_rule" "rule_before_anchor_decryption" {
  name        = "before-anchor-decryption-rule"
  description = "Positioned immediately BEFORE the anchor-decryption-rule. Updating"
  folder      = "All"
  position    = "pre"
  action      = "decrypt"

  # POSITIONING: Requires both relative_position and target_rule
  relative_position = "before"
  target_rule       = scm_decryption_rule.anchor_decryption_rule.id # Reference anchor by its ID

  # Core Match Criteria
  from        = ["trust"]
  to          = ["untrust"]
  source      = ["10.1.1.0/24"] # Specific source subnet
  destination = ["any"]
  service     = ["service-https"]
  category    = ["high-risk"]
  source_user = ["any"]
  type = {
    ssl_forward_proxy = {}
  }
}

resource "scm_decryption_rule" "rule_after_anchor_decryption" {
  name        = "after-anchor-decryption-rule_123"
  description = "Positioned immediately AFTER the anchor-decryption-rule."
  folder      = "All"
  position    = "pre"
  action      = "decrypt"

  # POSITIONING: Requires both relative_position and target_rule
  relative_position = "after"
  target_rule       = scm_decryption_rule.anchor_decryption_rule.id # Reference anchor by its ID

  # Core Match Criteria
  from        = ["any"]
  to          = ["untrust"]
  source      = ["any"]
  destination = ["192.168.1.10"] # Specific destination host
  service     = ["service-https"]
  category    = ["any"]
  source_user = ["any"]
  type = {
    ssl_forward_proxy = {}
  }
}

resource "scm_decryption_rule" "decryption_rule_ssl_inbound_inspection" {
  name        = "ssl_inbound_inspection_rule"
  description = "Decryption Rule with SSL Inbound Set"
  folder      = "All"
  position    = "pre" # Decryption Rules typically exist in the 'pre' rulebase
  action      = "decrypt"

  # Core Match Criteria (REQUIRED fields)
  from        = ["trust"]         # Source security zone
  to          = ["untrust"]       # Destination security zone
  source      = ["any"]           # Source addresses
  destination = ["any"]           # Destination addresses
  service     = ["service-https"] # Services (e.g., standard HTTPS port)
  category    = ["high-risk"]     # Destination URL Category
  source_user = ["any"]           # Source user/group
  type = {
    ssl_inbound_inspection = {
      certificates = ["Authentication Cookie CA"]
    }
  }
  destination_hip = ["any"]

  # Optional fields
  tag                = [scm_tag.decryption_position_tag.name]
  log_success        = true
  log_fail           = true
  disabled           = false
  negate_source      = false
  negate_destination = false
}
