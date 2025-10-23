#
# Data source to retrieve a list of URL Category objects.
#

# Example 1: Fetch a list of all URL Categories in the "Shared" folder.
data "scm_url_category_list" "all_shared" {
  folder = "All"
}

# Output the list of all URL Category objects from the "Shared" folder.
output "url_categories_list_all_shared" {
  description = "A list of all URL Category objects from the Shared folder."
  value       = data.scm_url_category_list.all_shared.data
}

# Example 2: Use pagination to get the first 5 URL Categories by name.
data "scm_url_category_list" "paginated" {
  folder = "All"
  limit  = 5
  offset = 0
}

output "url_categories_list_paginated" {
  description = "A paginated list of URL Category objects."
  value       = data.scm_url_category_list.paginated.data
}