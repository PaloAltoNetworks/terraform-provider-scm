

resource "scm_security_rule" "standard_web_access" {
  folder      = "All"
  name        = "Allow Standard Web Access DS1"
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
}
# --- Data Source Calls to Fetch Existing Rules ---

# 1. Fetch by ID (Best for direct lookup)
data "scm_security_rule" "standard_web_access_by_id" {
  # We use the computed ID from the resource
  id = scm_security_rule.standard_web_access.id
}
# --- Outputs to Verify Fetched Data ---

output "fetched_standard_web_id" {
  description = "ID of the rule fetched by its ID."
  value       = data.scm_security_rule.standard_web_access_by_id.id
}

output "fetched_standard_web_name" {
  description = "Name of the rule fetched by its ID."
  value       = data.scm_security_rule.standard_web_access_by_id.name
}

output "fetched_standard_web_description" {
  description = "Description of the rule fetched by its ID."
  value       = data.scm_security_rule.standard_web_access_by_id.description
}