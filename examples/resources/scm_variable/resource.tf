#
# Creates a variable in as-number format
#
resource "scm_variable" "scm_variable_asn" {
  folder      = "All"
  name        = "$tf_variable_asn"
  description = "Managed by Terraform"
  type        = "as-number"
  value       = "65535"
}

#
# Creates a variable in count format
#
resource "scm_variable" "scm_variable_count" {
  folder      = "All"
  name        = "$tf_variable_count"
  description = "Managed by Terraform"
  type        = "count"
  value       = "15"
}

#
# Creates a variable in fqdn format
#
resource "scm_variable" "scm_variable_fqdn" {
  folder      = "All"
  name        = "$tf_variable_fqdn"
  description = "Managed by Terraform"
  type        = "fqdn"
  value       = "scm.example.com"
}

#
# Creates a variable in group-id format
#
resource "scm_variable" "scm_variable_group_id" {
  folder      = "All"
  name        = "$tf_variable_group_id"
  description = "Managed by Terraform"
  type        = "group-id"
  value       = "10"
}

#
# Creates a variable in ip-range format
#
resource "scm_variable" "scm_variable_iprange" {
  folder      = "All"
  name        = "$tf_variable_iprange"
  description = "Managed by Terraform"
  type        = "ip-range"
  value       = "198.18.1.1-198.18.1.100"
}

#
# Creates a variable in ip-netmask format
#
resource "scm_variable" "scm_variable_ipaddr" {
  folder      = "All"
  name        = "$tf_variable_ipaddr"
  description = "Managed by Terraform"
  type        = "ip-netmask"
  value       = "198.18.2.0/24"
}

#
# Creates a variable in ip-wildcard format
#
resource "scm_variable" "scm_variable_ipwildcard" {
  folder      = "All"
  name        = "$tf_variable_ipwildcard"
  description = "Managed by Terraform"
  type        = "ip-wildcard"
  value       = "198.18.1.0/0.255.255.255"
}

#
# Creates a variable in percent format
#
resource "scm_variable" "scm_variable_percent" {
  folder      = "All"
  name        = "$tf_variable_percent"
  description = "Managed by Terraform"
  type        = "percent"
  value       = "10"
}

#
# Creates a variable in router-id format
#
resource "scm_variable" "scm_variable_router_id" {
  folder      = "All"
  name        = "$tf_variable_router_id"
  description = "Managed by Terraform"
  type        = "router-id"
  value       = "198.18.1.1"
}

#
# Creates a variable in timer format
#
resource "scm_variable" "scm_variable_timer" {
  folder      = "All"
  name        = "$tf_variable_timer"
  description = "Managed by Terraform"
  type        = "timer"
  value       = "1440"
}

#
# Creates a variable in zone format
#
resource "scm_variable" "scm_variable_zone" {
  folder      = "All"
  name        = "$tf_variable_zone"
  description = "Managed by Terraform"
  type        = "zone"
  value       = "internet"
}