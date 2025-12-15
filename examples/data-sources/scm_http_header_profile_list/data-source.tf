#
# Data source to retrieve a list of SCM HTTP Header Profile objects.
#

# Example 1: Fetch a list of all SCM HTTP Header Profile in the "Shared" folder.
data "scm_http_header_profile_list" "all_shared" {
  folder = "All"
}

# Output the list of all SCM HTTP Header Profile objects from the "Shared" folder.
output "scm_http_header_profile_list_all_shared" {
  description = "A list of all SCM HTTP Header Profile from the Shared folder."
  value       = data.scm_http_header_profile_list.all_shared.data
}
