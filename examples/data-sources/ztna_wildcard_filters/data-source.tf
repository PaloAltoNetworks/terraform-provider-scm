# List available filter fields for ZTNA Wildcards.
# Optionally specify a field name and search term to narrow results.
data "ztna_wildcard_filters" "example" {}

output "wildcard_filter_fields" {
  value = data.ztna_wildcard_filters.example.filters
}
