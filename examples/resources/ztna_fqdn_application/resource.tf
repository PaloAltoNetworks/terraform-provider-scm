resource "ztna_connector_group" "test_group" {
  name        = "terraform-test-connector-group"
  description = "Temporary group created by terraform test"
}

resource "ztna_fqdn_application" "example" {
  name        = "terraform-test-fqdn"
  group       = ztna_connector_group.test_group.oid
  description = "Made by Terraform"

  spec = [
    {
      fqdn       = "terraform-test.example.com"
      tcp_port   = "443"
      probe_type = "tcp_ping"
      probe_port = "443"
    }
  ]
}
