# This data source will call the "List" API endpoint
# and return all application filters in the "Shared" folder.

# 1. Use a single data block to fetch ALL application filters in the "Shared" folder.
data "scm_application_filter_list" "all_shared" {
  folder = "Shared"
}

# Alternative approach - just output the raw data to see what we get
output "application_filters_raw_data" {
  description = "Raw application filters data to debug structure"
  value       = try(data.scm_application_filter_list.all_shared, null)
}

# Simplified pagination example with error handling
data "scm_application_filter_list" "paginated_application_filters_example" {
  folder = "Shared"
  limit  = 5
  offset = 0
}

output "paginated_application_filters" {
  description = "A list of 5 application filters, skipping none - 0 offset"
  value = try({
    for application_filter in data.scm_application_filter_list.paginated_application_filters_example.data : application_filter.id => application_filter
  }, {})
}

output "pagination_application_filters_details" {
  description = "Details about the pagination."
  value = try({
    total_objects_in_folder = data.scm_application_filter_list.paginated_application_filters_example.total
    limit_used              = data.scm_application_filter_list.paginated_application_filters_example.limit
    offset_used             = data.scm_application_filter_list.paginated_application_filters_example.offset
  }, {})
}