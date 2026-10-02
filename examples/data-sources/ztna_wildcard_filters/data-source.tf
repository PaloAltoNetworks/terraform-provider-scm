# Without field param: returns all available filter field names categorised by type.
# free_form_filters and static_filters are populated; values is empty.
data "ztna_wildcard_filters" "all_fields" {}

output "wildcard_free_form_filters" {
  value = data.ztna_wildcard_filters.all_fields.free_form_filters
}

output "wildcard_static_filters" {
  value = data.ztna_wildcard_filters.all_fields.static_filters
}

# With field param: returns the possible values for the specified field.
# values is populated; free_form_filters and static_filters are empty.
data "ztna_wildcard_filters" "by_field" {
  field = "name"
}

output "wildcard_name_values" {
  value = data.ztna_wildcard_filters.by_field.values
}
