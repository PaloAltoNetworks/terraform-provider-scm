# First, create the tag objects that you will reference.
resource "scm_tag" "outbound_tag" {
  folder = "All"
  name   = "outbound143"
  color  = "Red"
}

resource "scm_tag" "web_tag" {
  folder = "All"
  name   = "web143"
  color  = "Blue"
}

# --- Existing Rules (Backward Compatibility) ---

resource "scm_security_rule" "standard_web_access" {
  folder      = "All"
  name        = "Allow Standard Web Access143"
  description = "Allow outbound web traffic to any destination..."
  position    = "pre" #

  # Standard security rule fields
  action      = "allow"
  category    = ["any"]
  application = ["web-browsing", "ssl"]
  service     = ["service-http", "service-https"]

  # Zones (From/To)
  from = ["untrust", "trust"]
  to   = ["trust"]

  # Addresses
  source      = ["any"]
  destination = ["any"]

  negate_source      = false
  negate_destination = false # Security-only

  # Identity & Content
  source_user     = ["any"] #
  source_hip      = ["any"] # Security-only
  destination_hip = ["any"] # Security-only


  # Logging
  log_start = true # Security-only
  log_end   = true # Security-only

  # Optional fields
  disabled = false

  # Use the names of the tags you just created.
  tag = [scm_tag.outbound_tag.name, scm_tag.web_tag.name] #
}

resource "scm_security_rule" "block_risky_saas" {
  folder      = "All"
  name        = "Block Risky SaaS Applications143"
  description = "Prevent data exfiltration by blocking risky SaaS apps..."
  # Internet rule fields
  action      = "deny"
  policy_type = "Internet" #
  security_settings = {
    # These settings are only available for internet rules
    anti_spyware                = "yes"
    vulnerability               = "yes"
    virus_and_wildfire_analysis = "yes"
  }
  block_web_application = [
    "facebook-posting",
  ]
  log_settings = {
    log_sessions = true #
  }

  # Common fields
  from        = ["any"]
  to          = ["any"]
  source      = ["any"]
  destination = ["any"]
  source_user = ["any"]


  # Optional fields
  disabled = false
  tag      = [scm_tag.outbound_tag.name, scm_tag.web_tag.name]
}


# --- NEW Examples Demonstrating Rule Ordering ---

# Example 1: Place a critical block rule at the absolute top
resource "scm_security_rule" "critical_block_top" {
  folder            = "All"
  name              = "CRITICAL Block Malicious IPs Top143"
  description       = "Always block known malicious IPs first."
  relative_position = "top" # Place at the very top of the 'pre' rulebase

  action      = "deny"
  from        = ["any"]
  to          = ["any"]
  source      = ["any"]
  destination = ["any"]
  source_user = ["any"]
  category    = ["any"]
  application = ["any"]
  service     = ["any"]
  log_end     = true
  tag         = [scm_tag.outbound_tag.name]
}

# Example 2: Place a cleanup rule at the absolute bottom
resource "scm_security_rule" "cleanup_deny_bottom" {
  folder            = "All"
  name              = "Cleanup Deny All Bottom143"
  description       = "Deny any traffic not explicitly allowed."
  relative_position = "bottom" # Place at the very bottom of the 'pre' rulebase (default rulebase)

  action      = "deny"
  from        = ["any"]
  to          = ["any"]
  source      = ["any"]
  destination = ["any"]
  source_user = ["any"]
  category    = ["any"]
  application = ["any"]
  service     = ["any"]
  log_end     = true
  tag         = [scm_tag.outbound_tag.name]
}


# Example 3: Place a rule *before* the standard web access rule
resource "scm_security_rule" "allow_updates_before_web" {
  folder            = "All"
  name              = "Allow OS Updates Before Web143"
  description       = "Allow specific OS update traffic before general web access."
  relative_position = "before"
  target_rule       = scm_security_rule.standard_web_access.id # Reference by id

  action      = "allow"
  from        = ["trust"]
  to          = ["untrust"]
  source      = ["any"] # Assumes this group exists
  destination = ["any"] # Assumes this group exists
  source_user = ["any"]
  category    = ["any"]
  application = ["ms-update", "apple-update"] # Example apps
  service     = ["service-https"]
  log_end     = true
  tag         = [scm_tag.outbound_tag.name]
}

# Example 4: Place a rule *after* the standard web access rule
resource "scm_security_rule" "allow_corp_apps_after_web" {
  folder            = "All"
  name              = "Allow Corp Apps After Web143"
  description       = "Allow access to specific corporate apps after general web access."
  relative_position = "after"
  target_rule       = scm_security_rule.standard_web_access.id # Reference by id

  action      = "allow"
  from        = ["trust"]
  to          = ["untrust"]
  source      = ["any"]
  destination = ["any"] # Assumes this group exists
  source_user = ["any"]
  category    = ["any"]
  application = ["ms-update"]     # Example app
  service     = ["service-https"] # Example services
  log_end     = true
  tag         = [scm_tag.web_tag.name]
}