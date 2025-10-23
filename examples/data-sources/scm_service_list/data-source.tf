# Data source to fetch all services in the "Shared" folder.
data "scm_service_list" "all_shared_services" {
  folder = "All"
}

# Creates a map where the key is the service ID and the value is the full object.
output "all_shared_services_map" {
  description = "A map of all service objects in the Shared folder, keyed by ID."
  value       = { for svc in data.scm_service_list.all_shared_services.data : svc.name => svc }
}

# Example of using pagination to get the first 5 services.
data "scm_service_list" "paginated_services" {
  folder = "All"
  limit  = 5
  offset = 0
}

# Outputs the list of paginated services.
output "paginated_services_list" {
  description = "A list of the first 5 services in the Shared folder."
  value       = data.scm_service_list.paginated_services.data
}

# Outputs pagination details, such as the total count and the limit used.
output "paginated_services_details" {
  description = "Details about the service list pagination."
  value = {
    total_in_folder = data.scm_service_list.paginated_services.total
    limit_used      = data.scm_service_list.paginated_services.limit
    offset_used     = data.scm_service_list.paginated_services.offset
  }
}
