#
# Data source to retrieve a list of SCM SAML Server Profile objects.
#

# Fetch a list of all SCM SAML Server Profile in the "All" folder.
data "scm_saml_server_profile_list" "all_shared" {
  folder = "All"
}

# Output the list of all SCM SAML Server Profile objects from the "All" folder.
output "scm_saml_server_profile_list_all_shared" {
  description = "A list of all SCM SAML Server Profile from the All folder."
  value       = data.scm_saml_server_profile_list.all_shared.data
}
