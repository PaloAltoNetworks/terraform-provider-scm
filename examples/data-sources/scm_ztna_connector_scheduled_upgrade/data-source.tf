# Retrieve the scheduled upgrade configuration for a ZTNA Connector.
data "scm_ztna_connector_scheduled_upgrade" "example" {
  oid = "22222222-2222-2222-2222-222222222222"
}

output "connector_scheduled_upgrade" {
  value = {
    image_id           = data.scm_ztna_connector_scheduled_upgrade.example.image_id
    scheduled_download = data.scm_ztna_connector_scheduled_upgrade.example.scheduled_download
    scheduled_upgrade  = data.scm_ztna_connector_scheduled_upgrade.example.scheduled_upgrade
  }
}
