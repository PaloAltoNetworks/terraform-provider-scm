# Fetch a list of all service group objects.
# This data source will call the "List" API endpoint.
data "scm_service_group_list" "all_service_groups" {
  folder = "All"
}

# Output the raw list of all service group objects
output "service_groups_list_raw" {
  description = "A list of all service group objects from the Shared folder."

  # This directly outputs the list of data from the data source
  value = data.scm_service_group_list.all_service_groups.data
}
