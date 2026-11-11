# Look up a ZTNA Connector by its ID.
data "scm_ztna_connector" "example" {
  oid = "22222222-2222-2222-2222-222222222222"
}

output "connector_output" {
  value = {
    oid         = data.scm_ztna_connector.example.oid
    name        = data.scm_ztna_connector.example.name
    description = data.scm_ztna_connector.example.description
    group       = data.scm_ztna_connector.example.group
  }
}
