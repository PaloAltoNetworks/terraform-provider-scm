#
# Data source to retrieve a list of SCM TLS Service Profile objects.
#

# Example 1: Fetch a list of all SCM TLS Service Profile in the "Shared" folder.
data "scm_tls_service_profile_list" "all_shared" {
  folder = "All"
}

# Output the list of all SCM TLS Service Profile objects from the "Shared" folder.
output "scm_tls_service_profile_list_all_shared" {
  description = "A list of all SCM TLS Service Profile from the Shared folder."
  value       = data.scm_tls_service_profile_list.all_shared.data
}