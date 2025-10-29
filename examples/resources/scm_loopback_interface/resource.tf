#
# Creates a loopback interface with static ipv4 address
#

resource "scm_loopback_interface" "scm_loopback_intf" {
  name    = "$scm_loopback_intf"
  comment = "Managed by Terraform"
  folder  = "ngfw-shared"
  ip = [
    {
      "name" : "198.18.1.1/32"
    }
  ]
}

#
# Creates a loopback interface with static ipv4 address, with default value loopback.123
#

resource "scm_loopback_interface" "scm_loopback_intf_2" {
  name          = "$scm_loopback_intf_2"
  comment       = "Managed by Terraform"
  folder        = "ngfw-shared"
  default_value = "loopback.123"
  ip = [
    {
      "name" : "198.18.1.2/32"
    }
  ]
}
