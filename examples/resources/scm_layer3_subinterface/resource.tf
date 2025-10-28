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