# Retrieve the upgrade status of a ZTNA Connector.
data "ztna_connector_upgrade_status" "example" {
  oid = "22222222-2222-2222-2222-222222222222"
}

output "connector_upgrade_status" {
  value = {
    active_image_id  = data.ztna_connector_upgrade_status.example.active_image_id
    active_version   = data.ztna_connector_upgrade_status.example.active_version
    upgrade_state    = data.ztna_connector_upgrade_status.example.upgrade_state
    upgrade_image_id = data.ztna_connector_upgrade_status.example.upgrade_image_id
  }
}
