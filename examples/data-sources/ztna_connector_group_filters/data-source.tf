# Without field param: returns all available filter field names categorised by type.
# free_form_filters and static_filters are populated; values is empty.
data "ztna_connector_group_filters" "all_fields" {}

output "connector_group_free_form_filters" {
  value = data.ztna_connector_group_filters.all_fields.free_form_filters
}

output "connector_group_static_filters" {
  value = data.ztna_connector_group_filters.all_fields.static_filters
}

# With field param: returns the possible values for the specified field.
# values is populated; free_form_filters and static_filters are empty.
data "ztna_connector_group_filters" "by_field" {
  field = "name"
}

output "connector_group_name_values" {
  value = data.ztna_connector_group_filters.by_field.values
}
