# Data source to fetch all external dynamic lists in the "Shared" folder.
data "scm_external_dynamic_list_list" "all_shared_edls" {
  folder = "Shared"
}

# Creates a map where the key is the EDL ID and the value is the full object.
output "all_shared_edls_map" {
  description = "A map of all external dynamic list objects in the Shared folder, keyed by ID."
  value       = data.scm_external_dynamic_list_list.all_shared_edls.data
  sensitive   = true
}

# Example of using pagination to get the first 5 EDLs.
data "scm_external_dynamic_list_list" "paginated_edls" {
  folder = "Shared"
  limit  = 5
  offset = 0
}

# Outputs pagination details, such as the total count and the limit used.
output "paginated_edls_details" {
  description = "Details about the EDL list pagination."
  value = {
    total_in_folder = data.scm_external_dynamic_list_list.paginated_edls.total
    limit_used      = data.scm_external_dynamic_list_list.paginated_edls.limit
    offset_used     = data.scm_external_dynamic_list_list.paginated_edls.offset
  }
  sensitive = true
}