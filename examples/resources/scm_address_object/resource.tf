resource "scm_address_object" "example" {
  folder      = "Shared"
  name        = "example"
  description = "Made by Terraform"
  ip_netmask  = "10.2.3.4"
}
