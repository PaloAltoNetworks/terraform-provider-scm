# List wildcard rules associated with a ZTNA Connector Group.
data "scm_ztna_connector_group_wildcards" "example" {
  oid = "11111111-1111-1111-1111-111111111111"
}

output "connector_group_wildcards" {
  value = data.scm_ztna_connector_group_wildcards.example.wildcards
}
