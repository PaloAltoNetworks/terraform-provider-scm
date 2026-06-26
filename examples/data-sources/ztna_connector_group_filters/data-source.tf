# List available filter fields for ZTNA Connector Groups.
# Optionally specify a field name and search term to narrow results.
data "ztna_connector_group_filters" "example" {}

output "connector_group_filter_fields" {
  value = data.ztna_connector_group_filters.example.filters
}
