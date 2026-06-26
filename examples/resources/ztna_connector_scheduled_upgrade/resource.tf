# Schedule an upgrade for a specific ZTNA Connector.
# Use the ztna_connector_images data source to retrieve a valid image_id.
resource "ztna_connector_scheduled_upgrade" "example" {
  oid                = "22222222-2222-2222-2222-222222222222"
  image_id           = "7.0.0-123"
  scheduled_download = "2025-01-15T02:00:00Z"
  scheduled_upgrade  = "2025-01-15T03:00:00Z"
}
