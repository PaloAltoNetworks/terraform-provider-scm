# This data source will call the "ListForwardingProfileDestinations" API endpoint
# and return all forwarding profile destinations in the "Mobile Users" folder.

# 1. Use a single data block to fetch ALL forwarding profile destinations in the "Mobile Users" folder.
data "scm_forwarding_profile_destination_list" "all_mobile_users" {
  folder = "Mobile Users"
}

# 2. Use a 'for' expression to transform the list into a map.
# This creates a map where the key is the destination id and the value is the destination object.
output "forwarding_profile_destination_data_source_results_from_list" {
  description = "A map of all forwarding profile destination objects in the Mobile Users folder, keyed by id."
  value       = { for dest in data.scm_forwarding_profile_destination_list.all_mobile_users.data : dest.id => dest }
}

data "scm_forwarding_profile_destination_list" "paginated_destinations_example" {
  folder = "Mobile Users"
  limit  = 5
  offset = 0
}

output "paginated_destinations" {
  description = "A list of 5 forwarding profile destinations, skipping none - 0 offset"
  value       = { for dest in data.scm_forwarding_profile_destination_list.paginated_destinations_example.data : dest.id => dest }
}

output "pagination_destinations_details" {
  description = "Details about the pagination."
  value = {
    total_objects_in_folder = data.scm_forwarding_profile_destination_list.paginated_destinations_example.total
    limit_used              = data.scm_forwarding_profile_destination_list.paginated_destinations_example.limit
  }
}