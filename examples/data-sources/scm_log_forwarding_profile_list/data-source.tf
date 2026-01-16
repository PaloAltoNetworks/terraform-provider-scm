#
# Data source to retrieve a list of SCM Log Forwarding Profile objects.
#

# Example 1: Fetch a list of all SCM Log Forwarding Profile in the "All" folder.
data "scm_log_forwarding_profile_list" "all_shared" {
  folder = "All"
}

# Output the list of all SCM Log Forwarding Profile objects from the "All" folder.
output "scm_log_forwarding_profile_list_all_shared" {
  description = "A list of all SCM Log Forwarding Profile from the All folder."
  value       = data.scm_log_forwarding_profile_list.all_shared.data
}