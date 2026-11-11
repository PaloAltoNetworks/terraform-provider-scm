resource "scm_ztna_connector_group" "test_group" {
  name        = "terraform-test-connector-group"
  description = "Temporary group created by terraform test"
}

resource "scm_ztna_wildcard" "example" {
  name        = "terraform-test-wildcard"
  fqdn        = "*.example.com"
  group       = scm_ztna_connector_group.test_group.oid
  description = "Made by Terraform"
}
