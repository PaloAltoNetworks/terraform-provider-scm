# List all connectors belonging to a specific ZTNA Connector Group.
data "ztna_connector_group_connectors" "example" {
  oid = "11111111-1111-1111-1111-111111111111"
}

output "connector_group_connectors" {
  value = data.ztna_connector_group_connectors.example.connectors
}
