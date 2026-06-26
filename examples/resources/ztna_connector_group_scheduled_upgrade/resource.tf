# Schedule an upgrade for a ZTNA Connector Group.
# All connectors in the group will be upgraded to the specified image version.
resource "ztna_connector_group_scheduled_upgrade" "example" {
  oid                = "11111111-1111-1111-1111-111111111111"
  image_id           = "7.0.0-123"
  scheduled_download = "2025-01-15T02:00:00Z"
  scheduled_upgrade  = "2025-01-15T03:00:00Z"
  rolling_upgrade    = true
  drain_timeout      = 300
}
