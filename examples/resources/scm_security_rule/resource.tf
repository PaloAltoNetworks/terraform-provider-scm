resource "scm_security_rule" "allow_http" {
  name         = "Allow-HTTP"
  action       = "allow"
  description  = "Allow HTTP traffic"
  
  # Matching criteria
  sources          = ["any"]
  destinations     = ["any"]
  source_users     = ["any"]
  categories       = ["any"]
  applications     = ["web-browsing"]
  services         = ["application-default"]
  froms            = ["trust"]
  tos              = ["untrust"]
  
  # Positioning parameters - add this rule at the top of pre-rules
  position         = "pre"
  relative_position = "top"
}

resource "scm_security_rule" "allow_https" {
  name         = "Allow-HTTPS"
  action       = "allow"
  description  = "Allow HTTPS traffic"
  
  # Matching criteria
  sources          = ["any"]
  destinations     = ["any"]
  source_users     = ["any"]
  categories       = ["any"]
  applications     = ["ssl"]
  services         = ["application-default"]
  froms            = ["trust"]
  tos              = ["untrust"]
  
  # Position this rule after the HTTP rule
  position         = "pre"
  relative_position = "after"
  target_rule      = scm_security_rule.allow_http.name
}

resource "scm_security_rule" "block_harmful" {
  name         = "Block-Harmful-Traffic" 
  action       = "deny"
  description  = "Block harmful traffic"
  
  # Matching criteria
  sources          = ["any"]
  destinations     = ["any"]
  source_users     = ["any"]
  categories       = ["malware"]
  applications     = ["any"]
  services         = ["any"]
  froms            = ["any"]
  tos              = ["any"]
  
  # Position this rule before the HTTP rule
  position         = "pre"
  relative_position = "before"
  target_rule      = scm_security_rule.allow_http.name
}