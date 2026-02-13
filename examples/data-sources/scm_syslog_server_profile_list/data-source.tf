#
# Data source to retrieve a list of SCM Syslog Server Profile objects.
#

# Example 1: Fetch a list of all SCM Syslog Server Profile in the "All" folder.
data "scm_syslog_server_profile_list" "all_shared" {
  folder = "All"
}

# Output the list of all SCM Syslog Server Profile objects from the "All" folder.
output "scm_syslog_server_profile_list_all_shared" {
  description = "A list of all SCM Syslog Server Profile from the Shared folder."
  value       = data.scm_syslog_server_profile_list.all_shared.data
}