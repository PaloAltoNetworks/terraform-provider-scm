# First, create the tag objects that you will reference.
resource "scm_tag" "outbound_tag" {
  folder = "Shared"
  name   = "outbound77"
  color  = "Red"
}

resource "scm_tag" "web_tag" {
  folder = "Shared"
  name   = "web77"
  color  = "Blue"
}


resource "scm_security_rule" "standard_web_access" {
  folder      = "Shared"
  name        = "Allow Standard Web Access77"
  description = "Allow outbound web traffic to any destination."
  position    = "pre"

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
  source_user     = ["any"]
  source_hip      = ["any"] # Security-only
  destination_hip = ["any"] # Security-only


  # Logging
  log_start = true # Security-only
  log_end   = true # Security-only

  # Optional fields
  disabled = false

  # Use the names of the tags you just created.
  tag = [scm_tag.outbound_tag.name, scm_tag.web_tag.name]
}

resource "scm_security_rule" "block_risky_saas" {
  folder      = "Shared"
  name        = "Block Risky SaaS Applications"
  description = "Prevent data exfiltration by blocking risky SaaS apps."

  # Internet rule fields
  action      = "deny"
  policy_type = "Internet"
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
    log_sessions = true
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