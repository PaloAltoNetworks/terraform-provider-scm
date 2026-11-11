
resource "scm_cngfw_ngfirewall" "example" {
  region                = "us-east-1"
  description           = "Sample CNGFW firewall managed by Terraform"
  customer_zone_id_list = ["use1-az1", "use1-az2"]
  tier                  = "standard"
  scm_labels            = ["sample-dev", "sample-prod"]
  mode                  = "CNGFW"
  tags = [
    {
      key   = "FirewallName"
      value = "sample-terraform-cngfw"
    },
    {
      key   = "CreatedBy"
      value = "Sample Terraform CNGFW"
    }
  ]
}
