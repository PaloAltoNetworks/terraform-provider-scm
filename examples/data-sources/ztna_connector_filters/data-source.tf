# List available filter fields for ZTNA Connectors.
# Optionally specify a field name and search term to narrow results.
data "ztna_connector_filters" "example" {}

output "connector_filter_fields" {
  value = data.ztna_connector_filters.example.filters
}
