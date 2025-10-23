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
