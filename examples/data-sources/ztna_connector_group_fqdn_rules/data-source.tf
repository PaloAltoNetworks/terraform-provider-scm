# List FQDN application rules associated with a ZTNA Connector Group.
data "ztna_connector_group_fqdn_rules" "example" {
  oid = "11111111-1111-1111-1111-111111111111"
}

output "connector_group_fqdn_rules" {
  value = data.ztna_connector_group_fqdn_rules.example.applications
}
