#
# Creates a ethernet interface with static ip address
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
# Creates a ethernet interface with dhcp-assigned ip address
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
# Creates a ethernet interface with pppoe
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
# Creates a ethernet interface with multiple static ip addresses
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
