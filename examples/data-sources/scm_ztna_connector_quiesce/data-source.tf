# Retrieve the quiesce state of a ZTNA Connector.
data "scm_ztna_connector_quiesce" "example" {
  oid = "22222222-2222-2222-2222-222222222222"
}

output "connector_quiesce" {
  value = {
    mode = data.scm_ztna_connector_quiesce.example.mode
  }
}
