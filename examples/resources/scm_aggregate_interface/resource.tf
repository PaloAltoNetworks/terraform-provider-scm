#
# Creates a layer 2 aggregate interface without vlan configuration
#

resource "scm_aggregate_interface" "scm_aggregate_intf_l2" {
  name    = "$scm_aggregate_intf_l2"
  comment = "Managed by Terraform"
  folder  = "ngfw-shared"
  layer2  = {}
}

#
# Creates a layer 2 aggregate interface with vlan and lacp configuration
#

resource "scm_aggregate_interface" "scm_aggregate_intf_l2_lacp" {
  name    = "$scm_aggregate_intf_l2_lacp"
  comment = "Managed by Terraform"
  folder  = "ngfw-shared"
  layer2 = {
    vlan_tag = 1234
    lacp = {
      enable            = true
      fast_failover     = true
      systen_priority   = 32768
      transmission_rate = "fast"
    }
    lldp = {
      enable = false
    }
  }
}

#
# Creates a layer3 aggregate interface without ip configuration
#

resource "scm_aggregate_interface" "scm_aggregate_intf_l3" {
  name    = "$scm_aggregate_intf_l3"
  comment = "Managed by Terraform"
  folder  = "ngfw-shared"
  layer3  = {}
}

#
# Creates a layer3 aggregate interface with static ip address and lacp
#

resource "scm_aggregate_interface" "scm_aggregate_intf_l3_static" {
  name    = "$scm_aggregate_intf_l3_static"
  comment = "Managed by Terraform"
  folder  = "ngfw-shared"
  layer3 = {
    ip = [
      {
        name = "198.18.1.1/24"
      }
    ]
    lacp = {
      enable            = true
      fast_failover     = true
      systen_priority   = 32768
      transmission_rate = "fast"
    }
  }
}

#
# Creates a layer3 aggregate interface with dhcp-assigned ip address
#

resource "scm_aggregate_interface" "scm_aggregate_intf_l3_dhcp" {
  name    = "$scm_aggregate_intf_l3_dhcp"
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
# Creates a layer3 aggregate interface with multiple static ip addresses
#

resource "scm_aggregate_interface" "scm_aggregate_intf_l3_complex" {
  name    = "$scm_aggregate_intf_l3_complex"
  comment = "Managed by Terraform"
  folder  = "ngfw-shared"
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