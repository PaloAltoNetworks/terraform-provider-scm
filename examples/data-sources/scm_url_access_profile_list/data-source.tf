#
# Data source to retrieve a list of URL Access Profile objects.
#

# Example 1: Fetch a list of all URL Access Profiles in the "Shared" folder.
data "scm_url_access_profile_list" "all_shared" {
  folder = "Shared"
}

# Output the list of all URL Access Profile objects from the "Shared" folder.
output "url_access_profiles_list_all_shared" {
  description = "A list of all URL Access Profile objects from the Shared folder."
  value       = data.scm_url_access_profile_list.all_shared.data
}
