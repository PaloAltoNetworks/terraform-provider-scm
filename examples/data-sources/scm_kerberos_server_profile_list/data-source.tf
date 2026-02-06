#
# Data source to retrieve a list of SCM Kerberos Server Profile objects.
#

# Example 1: Fetch a list of all SCM Kerberos Server Profile in the "All" folder.
data "scm_kerberos_server_profile_list" "all_shared" {
  folder = "All"
}

# Output the list of all SCM Kerberos Server Profile objects from the "All" folder.
output "scm_kerberos_server_profile_all_shared" {
  description = "A list of all SCM Kerberos Server Profile from the Shared folder."
  value       = data.scm_kerberos_server_profile_list.all_shared.data
}
