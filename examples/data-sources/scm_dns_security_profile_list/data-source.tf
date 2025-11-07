#
# Data source to retrieve a list of SCM DNS Security Profile objects.
#

# Fetch a list of all SCM DNS Security Profile in the "Shared" folder.
data "scm_dns_security_profile_list" "all_shared" {
  folder = "All"
}

# Output the list of all SCM DNS Security Profile objects from the "Shared" folder.
output "scm_dns_security_profile_list_all_shared" {
  description = "A list of all SCM DNS Security Profile from the Shared folder."
  value       = data.scm_dns_security_profile_list.all_shared.data
}