# 1. Use a single data block to fetch ALL dynamic user groups in the "Shared" folder. [cite: 2]
data "scm_dynamic_user_group_list" "all_shared_dugs" {
  folder = "Shared"
}

# 2. Use a 'for' expression to transform the list into a map. [cite: 4]
# This creates a map where the key is the group's ID and the value is the full group object. [cite: 4]
output "dug_results_from_list" {
  description = "A map of all dynamic user group objects in the Shared folder, keyed by id."
  value       = { for group in data.scm_dynamic_user_group_list.all_shared_dugs.data : group.id => group }
}

# This data source block shows an example of pagination.
data "scm_dynamic_user_group_list" "paginated_dugs_example" {
  folder = "Shared"
  limit  = 5
  offset = 0
}

# This output will show the paginated list of dynamic user groups.
output "paginated_dugs" {
  description = "A list of 5 dynamic user groups, skipping none - 0 offset"
  value       = { for group in data.scm_dynamic_user_group_list.paginated_dugs_example.data : group.id => group }
}

# This output shows details about the pagination results. 
output "pagination_dugs_details" {
  description = "Details about the pagination."
  value = {
    total_objects_in_folder = data.scm_dynamic_user_group_list.paginated_dugs_example.total
    limit_used              = data.scm_dynamic_user_group_list.paginated_dugs_example.limit
  }
}