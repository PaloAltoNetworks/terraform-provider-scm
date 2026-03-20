#
# Data source to retrieve a list of SCM Data Filtering Porofiles.
#

# Fetch a list of all SCM Data Filtering Porofiles in the "ngfw-shared" folder.
data "scm_data_filtering_profile_list" "all_ngfw_shared" {
  folder = "ngfw-shared"
}

# Output the list of all SCM Decryption Profile objects from the "All" folder.
output "scm_data_filtering_profile_list_ngfw_shared" {
  description = "A list of all SCM Data Filtering Profiles from the ngfw-shared folder."
  value       = data.scm_data_filtering_profile_list.all_ngfw_shared.data
}