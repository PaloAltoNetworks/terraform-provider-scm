# Schedule an upgrade for a specific ZTNA Connector.
#
# IMPORTANT — image_id is a numeric inventory ID, NOT a version string.
# Use the scm_ztna_connector_image_list data source to look up valid IDs:
#
#   data "scm_ztna_connector_image_list" "images" {}
#   # data.scm_ztna_connector_image_list.images.data[*].image_id
#
# Both timestamps MUST be set to a future time (RFC3339, UTC).
# Omitting either field defaults to the current time, triggering an immediate upgrade.
resource "scm_ztna_connector_scheduled_upgrade" "example" {
  oid                = "22222222-2222-2222-2222-222222222222"
  image_id           = "1000000000000000000" # replace with ID from scm_ztna_connector_image_list
  scheduled_download = "2027-01-15T02:00:00Z"
  scheduled_upgrade  = "2027-01-15T03:00:00Z"
}
