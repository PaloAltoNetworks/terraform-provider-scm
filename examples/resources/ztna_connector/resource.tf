resource "ztna_connector_group" "test_group" {
  name        = "terraform-test-connector-group"
  description = "Temporary group created by terraform test"
}

resource "ztna_connector" "example" {
  name        = "terraform-test-connector"
  group       = ztna_connector_group.test_group.oid
  description = "Made by Terraform"
}
