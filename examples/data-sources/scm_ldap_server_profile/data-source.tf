#
# Data source to retrieve a single SCM LDAP Server Profile object by its ID.
#

# Replace the ID with the UUID of the LDAP Server Profile you want to find.
data "scm_ldap_server_profile" "ldap_server_prof" {
  id = "a5006a3e-52b3-435e-9e3b-39b50dc72401"
}

# Output the details of the single LDAP Server Profile object found.
output "scm_ldap_server_profile_details" {
  value = {
    id     = data.scm_ldap_server_profile.ldap_server_prof.id
    folder = data.scm_ldap_server_profile.ldap_server_prof.folder
    name   = data.scm_ldap_server_profile.ldap_server_prof.name
  }
}