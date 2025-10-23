# This data source will call the "ListAddresses" API endpoint
# and return all filters in the "Shared" folder.

# 1. Use a single data block to fetch ALL addresses in the "Shared" folder.
data "scm_address_list" "all_shared" {
  folder = "All"
}

# 2. Use a 'for' expression to transform the list into a map.
# This creates a map where the key is the address name and the value is the address object.
output "address_data_source_results_from_list" {
  description = "A map of all address objects in the Shared folder, keyed by id."
  value       = { for addr in data.scm_address_list.all_shared.data : addr.id => addr }
}

data "scm_address_list" "paginated_addresses_example" {
  folder = "All"
  limit  = 5
  offset = 0
}

output "paginated_addresses" {
  description = "A list of 5 address names, skipping none - 0 offset"
  value       = { for addr in data.scm_address_list.paginated_addresses_example.data : addr.id => addr }
}

output "pagination_addresses_details" {
  description = "Details about the pagination."
  value = {
    total_objects_in_folder = data.scm_address_list.paginated_addresses_example.total
    limit_used              = data.scm_address_list.paginated_addresses_example.limit
  }
}