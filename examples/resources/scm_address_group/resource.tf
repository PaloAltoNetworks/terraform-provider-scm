# This file is embedded using go:embed
# First, create some addresses that will be used in the address group
resource "scm_address" "scm_address_ag_1" {
  folder      = "Shared"
  name        = "scm_address_ag_1"
  description = "First test address"
  ip_netmask  = "192.168.1.1/32"
}

resource "scm_address" "scm_address_ag_2" {
  folder      = "Shared"
  name        = "scm_address_ag_2"
  description = "Second test address"
  ip_netmask  = "192.168.1.2/32"
}

# Create the address group that references the addresses above
resource "scm_address_group" "scm_address_group_1" {
  folder      = "Shared"
  name        = "scm_address_group_1"
  description = "Sample address group created with Terraform"
  static = [
    scm_address.scm_address_ag_1.name,
    scm_address.scm_address_ag_2.name
  ]
}

# Create tags to be used for dynamic address group
resource "scm_tag" "scm_addressgroup_tag_1" {
  folder   = "Shared"
  name     = "scm_addressgroup_tag_1"
  comments = "Managed by Terraform"
  color    = "Orange"
}

resource "scm_tag" "scm_addressgroup_tag_2" {
  folder   = "Shared"
  name     = "scm_addressgroup_tag_2"
  comments = "Managed by Terraform"
  color    = "Blue"
}

# Create a dynamic addressgroup that matches both tags
resource "scm_address_group" "scm_addressgroup_dynamic" {
  folder      = "Shared"
  name        = "scm_addressgroup_dynamic"
  description = "Managed by Terraform"
  dynamic = {
    filter = "scm_addressgroup_tag_1 and scm_addressgroup_tag_2"
  }
}