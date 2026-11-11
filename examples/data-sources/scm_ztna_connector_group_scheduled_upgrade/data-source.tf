# Retrieve the scheduled upgrade configuration for a ZTNA Connector Group.
data "scm_ztna_connector_group_scheduled_upgrade" "example" {
  oid = "11111111-1111-1111-1111-111111111111"
}

output "connector_group_scheduled_upgrade" {
  value = {
    image_id           = data.scm_ztna_connector_group_scheduled_upgrade.example.image_id
    scheduled_download = data.scm_ztna_connector_group_scheduled_upgrade.example.scheduled_download
    scheduled_upgrade  = data.scm_ztna_connector_group_scheduled_upgrade.example.scheduled_upgrade
    rolling_upgrade    = data.scm_ztna_connector_group_scheduled_upgrade.example.rolling_upgrade
  }
}
