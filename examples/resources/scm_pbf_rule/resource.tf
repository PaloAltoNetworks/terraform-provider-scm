resource "scm_tag" "example_tag" {
  folder = "All"
  name   = "pbf-rule-tag-test-1"
  color  = "Red"
}

# --- PBF Rule Resource with discard action---
resource "scm_pbf_rule" "example_pbf_rule" {
  # Required fields
  name   = "pbf-test-rule-discard"
  folder = "All"

  # Optional: Description of the rule
  description = "PBF rule for forwarding specific traffic."

  # The source and destination criteria
  # Assuming 'any' for all except where a value is specified.
  from = {
    zone = ["zone-untrust"] # Traffic coming from the 'zone-untrust' zone
  }
  source      = ["any"]          # Source addresses
  destination = ["any"]          # Destination addresses
  application = ["any"]          # Applications
  service     = ["service-http"] # Services (ports/protocols)
  source_user = ["any"]


  # Define the action for the PBF rule
  action = {
    discard = {}
  }

  # Apply the created tag
  tag = [
    scm_tag.example_tag.name
  ]

  # Optional: Enforce Symmetric Return
  enforce_symmetric_return = {
    enabled = false
    # nexthop_address_list is optional here, leave it out for a simpler example
  }

  schedule = "non-work-hours"
}

# --- PBF Rule Resource with no-pbf action---
resource "scm_pbf_rule" "example_no_pbf_rule" {
  # Required fields
  name   = "pbf-test-rule-no-pbf"
  folder = "All"

  # Optional: Description of the rule
  description = "PBF rule for forwarding specific traffic"

  # The source and destination criteria
  # Assuming 'any' for all except where a value is specified.
  from = {
    zone = ["zone-untrust"] # Traffic coming from the 'zone-untrust' zone
  }
  source      = ["any"]           # Source addresses
  destination = ["any"]           # Destination addresses
  application = ["any"]           # Applications
  service     = ["service-https"] # Services (ports/protocols)
  source_user = ["any"]


  # Define the action for the PBF rule
  action = {
    no_pbf = {}
  }

  # Apply the created tag
  tag = [
    scm_tag.example_tag.name
  ]

  # Optional: Enforce Symmetric Return
  enforce_symmetric_return = {
    enabled = false
  }

  schedule = "non-work-hours"
}

# --- PBF Rule Resource with forward action---
resource "scm_pbf_rule" "example_forward_pbf_rule" {
  # Required fields
  name   = "pbf-test-rule-forward"
  folder = "All"

  # Optional: Description of the rule
  description = "PBF rule for forwarding specific traffic"

  # The source and destination criteria
  # Assuming 'any' for all except where a value is specified.
  from = {
    zone = ["zone-untrust"] # Traffic coming from the 'zone-untrust' zone
  }
  source      = ["any"]          # Source addresses
  destination = ["any"]          # Destination addresses
  application = ["any"]          # Applications
  service     = ["service-http"] # Services (ports/protocols)
  source_user = ["any"]


  # Define the action for the PBF rule
  action = {
    forward = {
      egress_interface = "ethernet1/1" # The egress interface for forwarding
      nexthop = {
        ip_address = "192.168.1.254" # The next hop IP address
      }
      #Optional: Monitor settings
      monitor = {
        ip_address             = "8.8.8.10"
        profile                = "test_tf_profile" #this should exists 
        disable_if_unreachable = true
      }
    }
  }

  # Apply the created tag
  tag = [
    scm_tag.example_tag.name
  ]

  # Optional: Enforce Symmetric Return
  enforce_symmetric_return = {
    enabled = true
  }

  schedule = "non-work-hours"
}