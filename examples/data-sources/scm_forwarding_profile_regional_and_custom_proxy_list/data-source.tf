# This data source will call the "ListForwardingProfileRegionalAndCustomProxies" API endpoint
# and return all forwarding profile regional and custom proxies in the "Mobile Users" folder.

# 1. Use a single data block to fetch ALL forwarding profile regional and custom proxies in the "Mobile Users" folder.
data "scm_forwarding_profile_regional_and_custom_proxy_list" "all_mobile_users" {
  folder = "Mobile Users"
}

# 2. Use a 'for' expression to transform the list into a map.
# This creates a map where the key is the proxy id and the value is the proxy object.
output "forwarding_profile_regional_and_custom_proxy_data_source_results_from_list" {
  description = "A map of all forwarding profile regional and custom proxy objects in the Mobile Users folder, keyed by id."
  value       = { for proxy in data.scm_forwarding_profile_regional_and_custom_proxy_list.all_mobile_users.data : proxy.id => proxy }
}

data "scm_forwarding_profile_regional_and_custom_proxy_list" "paginated_proxies_example" {
  folder = "Mobile Users"
  limit  = 5
  offset = 0
}

output "paginated_proxies" {
  description = "A list of 5 forwarding profile regional and custom proxy objects, skipping none - 0 offset"
  value       = { for proxy in data.scm_forwarding_profile_regional_and_custom_proxy_list.paginated_proxies_example.data : proxy.id => proxy }
}

output "pagination_proxies_details" {
  description = "Details about the pagination."
  value = {
    total_objects_in_folder = data.scm_forwarding_profile_regional_and_custom_proxy_list.paginated_proxies_example.total
    limit_used              = data.scm_forwarding_profile_regional_and_custom_proxy_list.paginated_proxies_example.limit
  }
}
