# List all available ZTNA Connector image versions.
# Use the returned image IDs when scheduling connector upgrades.
data "ztna_connector_images" "example" {}

output "connector_image_versions" {
  value = data.ztna_connector_images.example.values
}
