#
# Creates a ethernet interface used as parent-interface for subsequent examples
#

resource "scm_ethernet_interface" "scm_parent_interface" {
  name    = "$scm_parent_interface"
  comment = "Managed by Terraform"
  folder  = "ngfw-shared"
  layer3  = {}
}

#
# Creates a layer3 sub-interface with static ip address
#

resource "scm_layer3_subinterface" "scm_l3_subinterface" {
  name             = "$scm_parent_interface.100"
  comment          = "Managed by Terraform"
  folder           = "ngfw-shared"
  tag              = 100
  parent_interface = "$scm_parent_interface"
  depends_on       = [scm_ethernet_interface.scm_parent_interface]
  ip               = [{ name = "198.18.1.1/32" }]
}

resource "scm_ethernet_interface" "scm_parent_dhcp_interface" {
  name    = "$scm_parent_dhcp_interface"
  comment = "Managed by Terraform"
  folder  = "All"
  layer3  = {}
}

#
# Creates a layer3 sub-interface with dhcp
#

resource "scm_layer3_subinterface" "scm_l3_dhcp_subinterface" {
  name             = "$scm_parent_dhcp_interface.100"
  comment          = "Managed by Terraform"
  folder           = "All"
  tag              = 100
  parent_interface = "$scm_parent_dhcp_interface"
  depends_on       = [scm_ethernet_interface.scm_parent_dhcp_interface]
  dhcp_client = {
    # Enable DHCP client on this subinterface (Default is true)
    enable = true

    # Automatically create default route pointing to default gateway provided by server
    create_default_route = true

    # Metric of the default route created (Default is 10, valid 1-65535)
    default_route_metric = 20

    # Configuration to send a hostname in the DHCP request
    send_hostname = {
      enable = true
      # Set interface hostname (Default is "system-hostname")
      hostname = "client-vlan50-host"
    }
  }
}