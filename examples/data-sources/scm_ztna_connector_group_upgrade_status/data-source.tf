# Retrieve the scheduled upgrade status for a ZTNA Connector Group,
# including the upgrade status of all connectors within the group.
data "scm_ztna_connector_group_upgrade_status" "example" {
  oid = "11111111-1111-1111-1111-111111111111"
}

output "connector_group_upgrade_status" {
  value = {
    name           = data.scm_ztna_connector_group_upgrade_status.example.name
    upgrade_status = data.scm_ztna_connector_group_upgrade_status.example.upgrade_status
    data           = data.scm_ztna_connector_group_upgrade_status.example.data
  }
}
