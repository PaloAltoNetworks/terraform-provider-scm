#
# Data source to retrieve a list of SCM SCEP Profile objects.
#

# Example: Fetch a list of all SCM SCEP Profile in the "All" folder.
data "scm_scep_profile_list" "all_shared" {
  folder = "All"
}

# Output the list of all SCM SCEP Profile objects from the "All" folder.
output "scm_scep_profile_list_all_shared" {
  description = "A list of all SCM SCEP Profiles from the All folder."
  value       = data.scm_scep_profile_list.all_shared.data
}
