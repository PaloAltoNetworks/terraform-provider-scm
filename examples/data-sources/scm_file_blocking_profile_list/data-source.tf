#
# Data source to retrieve a list of SCM File Blocking Profile objects.
#

# Fetch a list of all SCM File Blocking Profile in the "Shared" folder.
data "scm_file_blocking_profile_list" "all_shared" {
  folder = "All"
}

# Output the list of all SCM File Blocking Profile objects from the "Shared" folder.
output "scm_file_blocking_profile_all_shared" {
  description = "A list of all SCM File Blocking Profile from the Shared folder."
  value       = data.scm_file_blocking_profile_list.all_shared.data
}