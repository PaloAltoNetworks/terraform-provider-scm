#
# Creates a ethernet interface used as parent-interface for subsequent examples
#

resource "scm_ethernet_interface" "scm_parent_interface" {
  name    = "$scm_parent_interface"
  comment = "Managed by Terraform"
  folder  = "ngfw-shared"
  layer2  = {}
}

#
# Creates a layer2 sub-interface with vlan tag 100
#

resource "scm_layer2_subinterface" "scm_layer2_subinterface" {
  name             = "$scm_parent_interface.100"
  comment          = "Managed by Terraform"
  folder           = "ngfw-shared"
  vlan_tag         = 100
  parent_interface = "$scm_parent_interface"
  depends_on       = [scm_ethernet_interface.scm_parent_interface]
}