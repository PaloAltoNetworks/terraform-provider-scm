resource "scm_address_group" "example" {
  folder = "Shared"
  name   = "example"
  static_list = [
    scm_address_object.x.name,
  ]
}

resource "scm_address_object" "x" {
  folder      = "Shared"
  name        = "foo"
  description = "Made by Terraform"
  fqdn        = "www.example.com"
}
