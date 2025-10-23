# Basic Anti-Spyware Profile
resource "scm_anti_spyware_profile" "scm_anti_spyware_profile_1" {
  folder                = "All"
  name                  = "scm_anti_spyware_profile_1"
  description           = "Managed by Terraform"
  cloud_inline_analysis = true
}

# Required object that will be referenced in examples
resource "scm_address" "scm_address_1" {
  folder      = "Shared"
  name        = "scm_address_1"
  description = "Made by Terraform"
  ip_netmask  = "10.2.3.4"
}


# Anti-Spyware Profile with exception EDL
resource "scm_anti_spyware_profile" "scm_anti_spyware_profile_2" {
  folder                = "All"
  name                  = "scm_anti_spyware_profile_2"
  description           = "Managed by Terraform"
  cloud_inline_analysis = true
  inline_exception_ip_address = [
    "scm_address_1"
  ]
  depends_on = [scm_address.scm_address_1]
}


# Anti-Spyware Profile with rules
resource "scm_anti_spyware_profile" "scm_anti_spyware_profile_3" {
  folder                = "All"
  name                  = "scm_anti_spyware_profile_3"
  description           = "Managed by Terraform"
  cloud_inline_analysis = true
  rules = [
    {
      name           = "Custom Rule",
      notes          = "Managed by Terraform",
      packet_capture = "single-packet",
      category       = "net-worm"
      severity       = ["critical"]
      threat_name    = "data-theft"
    }
  ]
}
