# This file is embedded using go:embed
resource "scm_tag" "scm_addr_tag_1" {
  folder = "All"
  name   = "tf_addr_tag_1"
  color  = "Red"
}

resource "scm_tag" "scm_addr_tag_2" {
  folder = "All"
  name   = "tf_addr_tag_2"
  color  = "Blue"
}

resource "scm_tag" "scm_addr_tag_3" {
  folder = "All"
  name   = "tf_addr_tag_3"
  color  = "Orange"
}

# IP Netmask
resource "scm_address" "scm_address_1" {
  folder      = "Prisma Access"
  name        = "tf_address_1"
  description = "Made by Terraform"
  ip_netmask  = "10.2.3.4"
}

# IP Range
resource "scm_address" "scm_address_2" {
  folder      = "Prisma Access"
  name        = "tf_address_2"
  description = "Small IP range test"
  ip_range    = "192.168.1.10-192.168.1.20"
}

# FQDN
resource "scm_address" "scm_address_3" {
  folder      = "Prisma Access"
  name        = "tf_address_3"
  description = "Simple FQDN test"
  fqdn        = "example.com"
}

# Class C wildcard
resource "scm_address" "scm_address_4" {
  folder      = "Prisma Access"
  name        = "tf_address_4"
  description = "Class C wildcard test"
  ip_wildcard = "192.168.1.0/0.0.0.255"
}

# Multiple tags
resource "scm_address" "scm_address_5" {
  folder      = "Prisma Access"
  name        = "tf_address_5"
  description = "Multiple tags test"
  ip_netmask  = "10.10.10.2/32"
  tag = [
    scm_tag.scm_addr_tag_1.name,
    scm_tag.scm_addr_tag_2.name,
    scm_tag.scm_addr_tag_3.name
  ]

  depends_on = [
    scm_tag.scm_addr_tag_1,
    scm_tag.scm_addr_tag_2,
    scm_tag.scm_addr_tag_3
  ]
}
