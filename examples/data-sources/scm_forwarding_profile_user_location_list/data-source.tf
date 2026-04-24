# This data source will call the "ListForwardingProfileUserLocations" API endpoint
# and return all forwarding profile user locations in the "Mobile Users" folder.

# 1. Use a single data block to fetch ALL user locations in the "Mobile Users" folder.
data "scm_forwarding_profile_user_location_list" "all_mobile_users" {
  folder = "Mobile Users"
}

# 2. Use a 'for' expression to transform the list into a map.
# This creates a map where the key is the user location id and the value is the user location object.
output "forwarding_profile_user_location_data_source_results_from_list" {
  description = "A map of all forwarding profile user location objects in the Mobile Users folder, keyed by id."
  value       = { for loc in data.scm_forwarding_profile_user_location_list.all_mobile_users.data : loc.id => loc }
}

data "scm_forwarding_profile_user_location_list" "paginated_user_locations_example" {
  folder = "Mobile Users"
  limit  = 5
  offset = 0
}

output "paginated_user_locations" {
  description = "A list of 5 forwarding profile user locations, skipping none - 0 offset"
  value       = { for loc in data.scm_forwarding_profile_user_location_list.paginated_user_locations_example.data : loc.id => loc }
}

output "pagination_user_locations_details" {
  description = "Details about the pagination."
  value = {
    total_objects_in_folder = data.scm_forwarding_profile_user_location_list.paginated_user_locations_example.total
    limit_used              = data.scm_forwarding_profile_user_location_list.paginated_user_locations_example.limit
  }
}