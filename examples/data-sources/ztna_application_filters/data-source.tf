# List available filter fields for ZTNA FQDN Applications.
# Optionally specify a field name and search term to narrow results.
data "ztna_application_filters" "example" {}

output "application_filter_fields" {
  value = data.ztna_application_filters.example.filters
}
