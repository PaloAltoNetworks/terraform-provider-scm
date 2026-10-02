# Without field param: returns all available filter field names categorised by type.
# free_form_filters and static_filters are populated; values is empty.
data "ztna_discovered_application_filters" "all_fields" {}

output "discovered_application_free_form_filters" {
  value = data.ztna_discovered_application_filters.all_fields.free_form_filters
}

output "discovered_application_static_filters" {
  value = data.ztna_discovered_application_filters.all_fields.static_filters
}

# With field param: returns the possible values for the specified field.
# values is populated; free_form_filters and static_filters are empty.
data "ztna_discovered_application_filters" "by_field" {
  field = "onboarded"
}

output "discovered_application_name_values" {
  value = data.ztna_discovered_application_filters.by_field.values
}
