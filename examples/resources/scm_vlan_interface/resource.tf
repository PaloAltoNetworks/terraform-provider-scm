#
# Creates a vlan interface used as parent-interface for subsequent examples
#

resource "scm_vlan_interface" "scm_vlan_interface_ipv4" {
  name     = "$scm_vlan_interface_ipv4"
  comment  = "Managed by Terraform"
  folder   = "ngfw-shared"
  vlan_tag = "1234"
  ip = [
    { name = "198.18.1.1/24" }
  ]
  adjust_tcp_mss = {
    enable              = true
    ipv4_mss_adjustment = 40
    ipv6_mss_adjustment = 60
  }
}
