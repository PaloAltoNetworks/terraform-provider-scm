# Retrieve the scheduled upgrade configuration for a ZTNA Connector.
data "ztna_connector_scheduled_upgrade" "example" {
  oid = "22222222-2222-2222-2222-222222222222"
}

output "connector_scheduled_upgrade" {
  value = {
    image_id           = data.ztna_connector_scheduled_upgrade.example.image_id
    scheduled_download = data.ztna_connector_scheduled_upgrade.example.scheduled_download
    scheduled_upgrade  = data.ztna_connector_scheduled_upgrade.example.scheduled_upgrade
  }
}
