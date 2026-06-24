resource "ztna_connector_group" "test_group" {
  name        = "terraform-test-connector-group"
  description = "Temporary group created by terraform test"
}

resource "ztna_subnet" "example" {
  name        = "terraform-test-subnet"
  ip_subnets  = "10.0.0.0/24"
  group       = ztna_connector_group.test_group.oid
  description = "Made by Terraform"
}
