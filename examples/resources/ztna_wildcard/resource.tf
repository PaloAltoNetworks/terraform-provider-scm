resource "ztna_connector_group" "test_group" {
  name        = "terraform-test-connector-group"
  description = "Temporary group created by terraform test"
}

resource "ztna_wildcard" "example" {
  name        = "terraform-test-wildcard"
  fqdn        = "*.example.com"
  group       = ztna_connector_group.test_group.oid
  description = "Made by Terraform"
}
