#
# Creates various resources used for subsequent examples
#

resource "scm_aggregate_interface" "scm_ae_intf" {
  name    = "$scm_ae_intf"
  comment = "Managed by Terraform"
  folder  = "ngfw-shared"
  layer2  = {}
}

#
# Creates a layer 2 ethernet interface without vlan configuration
#

resource "scm_ethernet_interface" "scm_l2_intf" {
  name        = "$scm_l2_intf"
  comment     = "Managed by Terraform"
  folder      = "ngfw-shared"
  link_speed  = "auto"
  link_duplex = "full"
  link_state  = "auto"
  layer2      = {}
}

#
# Creates a tap ethernet interface without vlan configuration
#

resource "scm_ethernet_interface" "scm_tap_intf" {
  name        = "$scm_tap_intf"
  comment     = "Managed by Terraform"
  folder      = "ngfw-shared"
  link_speed  = "auto"
  link_duplex = "full"
  link_state  = "auto"
  tap         = {}
}

#
# Creates a layer3 ethernet interface without ip configuration
#

resource "scm_ethernet_interface" "scm_l3_intf" {
  name        = "$scm_l3_intf"
  comment     = "Managed by Terraform"
  folder      = "ngfw-shared"
  link_speed  = "auto"
  link_duplex = "full"
  link_state  = "auto"
  layer3      = {}
}

#
# Creates a layer3 ethernet interface with static ip address
#

resource "scm_ethernet_interface" "scm_l3_intf_static" {
  name    = "$scm_l3_intf_static"
  comment = "Managed by Terraform"
  folder  = "ngfw-shared"
  layer3 = {
    ip = [
      {
        name = "198.18.1.1/24"
      }
    ]
  }
}

#
# Creates a layer3 ethernet interface with dhcp-assigned ip address
#

resource "scm_ethernet_interface" "scm_l3_intf_dhcp" {
  name    = "$scm_l3_intf_dhcp"
  comment = "Managed by Terraform"
  folder  = "ngfw-shared"
  layer3 = {
    dhcp_client = {
      enable               = true
      create_default_route = true
      default_route_metric = 10
    }
  }
}

#
# Creates a layer3 ethernet interface with pppoe
#

resource "scm_ethernet_interface" "scm_l3_intf_pppoe" {
  name    = "$scm_l3_intf_pppoe"
  comment = "Managed by Terraform"
  folder  = "ngfw-shared"

  layer3 = {
    pppoe = {
      authentication       = "auto"
      enable               = true
      username             = "testname"
      password             = "testpass"
      create_default_route = true
      default_route_metric = 10
    }
  }
}

#
# Creates a layer3 ethernet interface with multiple static ip addresses
#

resource "scm_ethernet_interface" "scm_l3_intf_complex" {
  name        = "$scm_l3_intf_complex"
  comment     = "Managed by Terraform"
  folder      = "ngfw-shared"
  link_speed  = "auto"
  link_duplex = "full"
  link_state  = "auto"
  layer3 = {
    ip = [
      {
        name = "198.18.1.1/24"
        name = "198.18.1.2/32"
      }
    ]
    mtu = 1500
  }
}

#
# Creates an ethernet interface assigned to an AggregateEthernet Interface
#

resource "scm_ethernet_interface" "scm_ae_member_1" {
  name            = "$scm_ae_member_1"
  comment         = "Managed by Terraform"
  folder          = "ngfw-shared"
  aggregate_group = "$scm_ae_intf"
  link_speed      = "auto"
  link_duplex     = "full"
  link_state      = "auto"
  depends_on      = [scm_aggregate_interface.scm_ae_intf]
}

#
# Creates an ethernet interface assigned to an AggregateEthernet Interface
#

resource "scm_ethernet_interface" "scm_ae_member_2" {
  name            = "$scm_ae_member_2"
  comment         = "Managed by Terraform"
  folder          = "ngfw-shared"
  aggregate_group = "$scm_ae_intf"
  link_speed      = "auto"
  link_duplex     = "full"
  link_state      = "auto"
  depends_on      = [scm_aggregate_interface.scm_ae_intf]
}