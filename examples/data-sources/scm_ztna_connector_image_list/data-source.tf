# List all available ZTNA Connector image versions.
# Use the returned image IDs when scheduling connector upgrades.
data "scm_ztna_connector_image_list" "example" {}

output "connector_image_versions" {
  value = data.scm_ztna_connector_image_list.example.data
}
