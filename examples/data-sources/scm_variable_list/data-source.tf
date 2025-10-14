#
# Data source to retrieve a list of all variables in a folder.
#

# Example 1: Fetch a list of all variables in the "Global" aka "All" folder.
data "scm_variable_list" "all_global" {
  folder = "All"
}

# Output the list of all variables from the "Global" aka "All" folder.
output "scm_variable_list_global" {
  description = "A list of all variables in the 'Global' folder"
  value       = data.scm_variable_list.all_global.data
}

# Example 2: Use pagination to get the first variables by name.
data "scm_variable_list" "paginated" {
  folder = "All"
  limit  = 5
  offset = 0
}

output "scm_variable_list_paginated" {
  description = "A paginated list of variables."
  value       = data.scm_variable_list.paginated.data
}