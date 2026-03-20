#
# Data source to retrieve a list of SCM LDAP Server Profile objects.
#

# Fetch a list of all SCM LDAP Server Profiles in the "All" folder.
data "scm_ldap_server_profile_list" "all_shared" {
  folder = "All"
}

# Output the list of all SCM LDAP Server Profile objects from the "All" folder.
output "scm_ldap_server_profile_list_all_shared_filtered" {
  description = "A list of all SCM LDAP Server Profiles from the All folder."
  value = [
    for profile in data.scm_ldap_server_profile_list.all_shared.data : {
      id     = profile.id
      name   = profile.name
      folder = profile.folder
      server = profile.server
    }
  ]
}