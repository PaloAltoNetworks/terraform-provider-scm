variable "folder_scope" {
  description = "The folder scope for the SCM resource (e.g., 'Shared', 'Predefined', or a specific folder name)."
  type        = string
  default     = "Service Connections"
}

# ------------------------------------------------------------------
# Data Source List Lookup
# ------------------------------------------------------------------

data "scm_service_connection_list" "all_connections_in_folder" {
  folder = var.folder_scope
  limit  = 50
}

output "list_of_all_connection_names" {
  description = "A list of names for all Service Connections found in the target folder."
  value = [
    for conn in data.scm_service_connection_list.all_connections_in_folder.data : conn.name
  ]
}

output "total_connections_count" {
  description = "The total number of connections found by the list data source."
  value       = length(data.scm_service_connection_list.all_connections_in_folder.data)
}
