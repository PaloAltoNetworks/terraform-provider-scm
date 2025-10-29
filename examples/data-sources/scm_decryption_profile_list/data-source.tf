#
# Data source to retrieve a list of SCM Decryption Profile objects.
#

# Fetch a list of all SCM Decryption Profile in the "Shared" folder.
data "scm_decryption_profile_list" "all_shared" {
  folder = "All"
}

# Output the list of all SCM Decryption Profile objects from the "Shared" folder.
output "scm_decryption_profile_list_all_shared" {
  description = "A list of all SCM Decryption Profile from the Shared folder."
  value       = data.scm_decryption_profile_list.all_shared.data
}