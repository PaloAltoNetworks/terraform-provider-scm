# List available filter fields for ZTNA Discovered Applications.
# Optionally specify a field name and search term to narrow results.
data "ztna_discovered_application_filters" "example" {}

output "discovered_application_filter_fields" {
  value = data.ztna_discovered_application_filters.example.filters
}
