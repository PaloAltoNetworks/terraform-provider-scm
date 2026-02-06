# 1. Use a single data block to fetch ALL application groups in the "All" folder.
data "scm_application_group_list" "all_shared" {
  folder = "All"
}

# 2. Use a 'for' expression to transform the list into a map.
# This creates a map where the key is the group name and the value is the group object.
output "application_groups_data_source_results_from_list" {
  description = "A map of all application group objects in the All folder, keyed by id."
  value       = { for group in data.scm_application_group_list.all_shared.data : group.id => group }
}

# Example of using pagination to get a subset of application groups.
data "scm_application_group_list" "paginated_example" {
  folder = "All"
  limit  = 5
  offset = 0
}

output "paginated_application_groups" {
  description = "A list of 5 application groups, skipping none (0 offset)."
  value       = { for group in data.scm_application_group_list.paginated_example.data : group.id => group }
}

output "pagination_application_groups_details" {
  description = "Details about the pagination."
  value = {
    total_objects_in_folder = data.scm_application_group_list.paginated_example.total
    limit_used              = data.scm_application_group_list.paginated_example.limit
  }
}