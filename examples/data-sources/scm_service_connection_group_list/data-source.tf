variable "folder_scope" {
  description = "The folder scope for the SCM resource (e.g., 'Shared', 'Predefined', or a specific folder name)."
  type        = string
  default     = "Service Connections"
}

# ------------------------------------------------------------------
# Data Source List: SCM Service Connection Group (List Lookup)
# ------------------------------------------------------------------
data "scm_service_connection_group_list" "all_groups_in_folder" {
  # Filters the list to only include groups in the specified folder
  folder = var.folder_scope
  limit  = 50
}

output "list_of_all_group_names" {
  description = "A list of names for all Service Connection Groups found in the target folder."
  # Uses a 'for' expression to extract the name from each group object in the list
  value = [
    for group in data.scm_service_connection_group_list.all_groups_in_folder.data : group.name
  ]
}

output "total_groups_count" {
  description = "The total number of Service Connection Groups found by the list data source."
  # Gets the count of items in the returned list
  value = length(data.scm_service_connection_group_list.all_groups_in_folder.data)
}
