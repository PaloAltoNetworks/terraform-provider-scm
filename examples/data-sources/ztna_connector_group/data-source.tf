# Look up a ZTNA Connector Group by its ID.
data "ztna_connector_group" "example" {
  oid = "11111111-1111-1111-1111-111111111111"
}

output "connector_group_output" {
  value = {
    oid         = data.ztna_connector_group.example.oid
    name        = data.ztna_connector_group.example.name
    description = data.ztna_connector_group.example.description
  }
}
