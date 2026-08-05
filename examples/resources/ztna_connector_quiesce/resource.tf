# Manage the quiesce state of a ZTNA Connector.
resource "ztna_connector_quiesce" "example" {
  oid  = "22222222-2222-2222-2222-222222222222"
  mode = "<mode>"
}
