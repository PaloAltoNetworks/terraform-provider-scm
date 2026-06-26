# List available filter fields for ZTNA Subnets.
# Optionally specify a field name and search term to narrow results.
data "ztna_subnet_filters" "example" {}

output "subnet_filter_fields" {
  value = data.ztna_subnet_filters.example.filters
}
