# Look up a ZTNA Connector by its ID.
data "ztna_connector" "example" {
  oid = "22222222-2222-2222-2222-222222222222"
}

output "connector_output" {
  value = {
    oid         = data.ztna_connector.example.oid
    name        = data.ztna_connector.example.name
    description = data.ztna_connector.example.description
    group       = data.ztna_connector.example.group
  }
}
