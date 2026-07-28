# Without field param: returns all available filter field names categorised by type.
# free_form_filters and static_filters are populated; values is empty.
data "ztna_subnet_filters" "all_fields" {}

output "subnet_free_form_filters" {
  value = data.ztna_subnet_filters.all_fields.free_form_filters
}

output "subnet_static_filters" {
  value = data.ztna_subnet_filters.all_fields.static_filters
}

# With field param: returns the possible values for the specified field.
# values is populated; free_form_filters and static_filters are empty.
data "ztna_subnet_filters" "by_field" {
  field = "name"
}

output "subnet_name_values" {
  value = data.ztna_subnet_filters.by_field.values
}
