#
# Creates a tunnel interface with static ipv4 address
#

resource "scm_tunnel_interface" "scm_tunnel_intf" {
  name    = "$scm_tunnel_intf"
  comment = "Managed by Terraform"
  folder  = "ngfw-shared"
  ip = [
    {
      name = "198.18.1.1/32"
    }
  ]
}

#
# Creates a tunnel interface with static ipv4 address, with default value tunnel.123
#

resource "scm_tunnel_interface" "scm_tunnel_intf_2" {
  name    = "$scm_tunnel_intf_2"
  comment = "Managed by Terraform"
  folder  = "ngfw-shared"
  ip = [
    {
      name = "198.18.1.2/32"
    }
  ]
}

#
# Creates an ip subnet variable used in the subsequent example
#

resource "scm_variable" "scm_ipv6_prefix" {
  folder      = "ngfw-shared"
  name        = "$scm_ipv6_prefix"
  description = "Managed by Terraform"
  type        = "ip-netmask"
  value       = "2001:0db8:abcd:0001::/64"
}

#
# Creates a tunnel interface with ipv6 address, with default value tunnel.321
#

resource "scm_tunnel_interface" "scm_tunnel_intf_3" {
  name          = "$scm_tunnel_intf_3"
  comment       = "Managed by Terraform"
  folder        = "ngfw-shared"
  default_value = "tunnel.321"
  ipv6 = {
    enabled      = true
    interface_id = "EUI-64"
    address = [
      {
        name   = "$scm_ipv6_prefix"
        prefix = {}
      }
    ]
  }
  depends_on = [scm_variable.scm_ipv6_prefix]
}
