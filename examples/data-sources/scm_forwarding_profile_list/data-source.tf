# This data source will call the "ListForwardingProfiles" API endpoint
# and return all forwarding profiles in the "Mobile Users" folder.

# 1. Use a single data block to fetch ALL forwarding profiles in the "Mobile Users" folder.
data "scm_forwarding_profile_list" "all_mobile_users" {
  folder = "Mobile Users"
}

# 2. Use a 'for' expression to transform the list into a map.
# This creates a map where the key is the forwarding profile id and the value is the profile object.
output "forwarding_profile_data_source_results_from_list" {
  description = "A map of all forwarding profile objects in the Mobile Users folder, keyed by id."
  value       = { for fp in data.scm_forwarding_profile_list.all_mobile_users.data : fp.id => fp }
}

data "scm_forwarding_profile_list" "paginated_profiles_example" {
  folder = "Mobile Users"
  limit  = 5
  offset = 0
}

output "paginated_forwarding_profiles" {
  description = "A list of 5 forwarding profile names, skipping none - 0 offset"
  value       = { for fp in data.scm_forwarding_profile_list.paginated_profiles_example.data : fp.id => fp }
}

output "pagination_forwarding_profiles_details" {
  description = "Details about the pagination."
  value = {
    total_objects_in_folder = data.scm_forwarding_profile_list.paginated_profiles_example.total
    limit_used              = data.scm_forwarding_profile_list.paginated_profiles_example.limit
  }
}
