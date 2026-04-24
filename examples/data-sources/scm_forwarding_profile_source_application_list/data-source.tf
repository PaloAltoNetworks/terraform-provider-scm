# This data source will call the "List" API endpoint
# and return all forwarding profile source applications in the "Mobile Users" folder.

# 1. Use a single data block to fetch ALL forwarding profile source applications in the "Mobile Users" folder.
data "scm_forwarding_profile_source_application_list" "mobile_users" {
  folder = "Mobile Users"
}

# Alternative approach - just output the raw data to see what we get
output "forwarding_profile_source_applications_raw_data" {
  description = "Raw forwarding profile source applications data to debug structure"
  value       = try(data.scm_forwarding_profile_source_application_list.mobile_users, null)
}

# Simplified pagination example with error handling
data "scm_forwarding_profile_source_application_list" "paginated_example" {
  folder = "Mobile Users"
  limit  = 5
  offset = 0
}

output "paginated_forwarding_profile_source_applications" {
  description = "A list of 5 forwarding profile source applications, skipping none - 0 offset"
  value = try({
    for app in data.scm_forwarding_profile_source_application_list.paginated_example.data : app.id => app
  }, {})
}

output "pagination_details" {
  description = "Details about the pagination."
  value = try({
    total_objects_in_folder = data.scm_forwarding_profile_source_application_list.paginated_example.total
    limit_used              = data.scm_forwarding_profile_source_application_list.paginated_example.limit
    offset_used             = data.scm_forwarding_profile_source_application_list.paginated_example.offset
  }, {})
}
