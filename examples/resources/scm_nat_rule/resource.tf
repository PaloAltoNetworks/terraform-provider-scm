resource "scm_tag" "example_tag" {
  folder = "All"
  name   = "example-tag"
  color  = "Red"
}

#Source Translation (SNAT) - Dynamic IP and Port
resource "scm_nat_rule" "example_nat_rule" {
  # Required Attributes
  name        = "snat-to-internet-1"
  from        = ["any"]         # Source zone(s) of the original packet
  to          = ["untrust"]     # Destination zone of the original packet
  source      = ["any"]         # Source address(es) of the original packet
  destination = ["any"]         # Destination address(es) of the original packet
  service     = "service-https" # The service of the original packet

  # Optional Attributes
  description = "Dynamic SNAT for internal traffic accessing the internet. Updating"
  disabled    = false
  nat_type    = "ipv4" # Default value, explicitly setting for clarity

  # Scope (Use either folder, device, or snippet)
  folder = "All"

  # Apply the created tag
  tag = [
    scm_tag.example_tag.name,
  ]

  # Source Translation (SNAT) - Dynamic IP and Port
  source_translation = {
    dynamic_ip_and_port = {
      translated_address = ["1.1.1.1", "1.1.1.5"]
    }
    # NOTE: You must only specify one of: dynamic_ip, dynamic_ip_and_port, or static_ip
  }

  # Destination Translation (DNAT) is for Inbound NAT, usually used alone or with Static SNAT.
  destination_translation = {
    translated_address = "192.168.1.10"
    translated_port    = 8080
  }

  active_active_device_binding = "1"

}

#Source Translation (SNAT) - Static IP - Bidirectional - no
resource "scm_nat_rule" "example_nat_static_rule" {
  # Required Attributes
  name        = "snat-to-bid-1"
  from        = ["any"]         # Source zone(s) of the original packet
  to          = ["untrust"]     # Destination zone of the original packet
  source      = ["any"]         # Source address(es) of the original packet
  destination = ["any"]         # Destination address(es) of the original packet
  service     = "service-https" # The service of the original packet

  # Optional Attributes
  description = "Dynamic SNAT for internal traffic accessing the internet. Updating"
  disabled    = false
  nat_type    = "ipv4" # Default value, explicitly setting for clarity

  # Scope (Use either folder, device, or snippet)
  folder = "All"

  # Apply the created tag
  tag = [
    scm_tag.example_tag.name,
  ]

  # Source Translation (SNAT) - Dynamic IP and Port
  source_translation = {
    static_ip = {
      translated_address = "1.1.1.5",
      bi_directional     = "no"
    }
    # NOTE: You must only specify one of: dynamic_ip, dynamic_ip_and_port, or static_ip
  }

  # Destination Translation (DNAT) is for Inbound NAT, usually used alone or with Static SNAT.
  destination_translation = {
    translated_address = "192.168.1.10"
    translated_port    = 8080
  }

  active_active_device_binding = "1"

}

#Source Translation (SNAT) - Static IP - Bidirectional - yes
resource "scm_nat_rule" "example_nat_static_rule_2" {
  # Required Attributes
  name        = "snat-to-bid-yes-1"
  from        = ["any"]         # Source zone(s) of the original packet
  to          = ["untrust"]     # Destination zone of the original packet
  source      = ["any"]         # Source address(es) of the original packet
  destination = ["any"]         # Destination address(es) of the original packet
  service     = "service-https" # The service of the original packet

  # Optional Attributes
  description = "Dynamic SNAT for internal traffic accessing the internet. Updating"
  disabled    = false
  nat_type    = "ipv4" # Default value, explicitly setting for clarity

  # Scope (Use either folder, device, or snippet)
  folder = "All"

  # Apply the created tag
  tag = [
    scm_tag.example_tag.name,
  ]

  # Source Translation (SNAT) - Dynamic IP and Port
  source_translation = {
    static_ip = {
      translated_address = "1.1.1.5",
      bi_directional     = "yes"
    }
    # NOTE: You must only specify one of: dynamic_ip, dynamic_ip_and_port, or static_ip
  }

  active_active_device_binding = "1"

}


#Source Translation (SNAT) - Dynamic IP 
resource "scm_nat_rule" "example_nat_dynamic_rule" {
  # Required Attributes
  name        = "snat-to-dyanamic-1"
  from        = ["any"]         # Source zone(s) of the original packet
  to          = ["untrust"]     # Destination zone of the original packet
  source      = ["any"]         # Source address(es) of the original packet
  destination = ["any"]         # Destination address(es) of the original packet
  service     = "service-https" # The service of the original packet

  # Optional Attributes
  description = "Dynamic SNAT for internal traffic accessing the internet. Updating"
  disabled    = false
  nat_type    = "ipv4" # Default value, explicitly setting for clarity

  # Scope (Use either folder, device, or snippet)
  folder = "All"

  # Apply the created tag
  tag = [
    scm_tag.example_tag.name,
  ]

  # Source Translation (SNAT) - Dynamic IP and Port
  source_translation = {
    dynamic_ip = {
      translated_address = ["1.1.1.0/24"],
      fallback = {
        translated_address = ["1.1.1.0"],
        interface_address = {
          interface = "ethernet1/1",
          ip        = "1.1.1.5"
        }
      }
    }
    # NOTE: You must only specify one of: dynamic_ip, dynamic_ip_and_port, or static_ip
  }

  # Destination Translation (DNAT) is for Inbound NAT, usually used alone or with Static SNAT.
  destination_translation = {
    translated_address = "192.168.1.10"
    translated_port    = 8080
  }

  active_active_device_binding = "1"

}