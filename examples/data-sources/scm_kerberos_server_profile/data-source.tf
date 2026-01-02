
#
# Data source to retrieve a single SCM Kerberos Server Profile object by its ID.
#

# Replace the ID with the UUID of the SCM Kerberos Server Profile you want to find.
data "scm_kerberos_server_profile" "scm_kerberos_prof" {
  id = "6bd818f8-9679-4031-86df-17b8b40842a0"
}

# Output the details of the single SCM Kerberos Server Profile object found.
output "scm_kerberos_server_profile_details" {
  value = {
    profile_id = data.scm_kerberos_server_profile.scm_kerberos_prof.id
    folder     = data.scm_kerberos_server_profile.scm_kerberos_prof.folder
    name       = data.scm_kerberos_server_profile.scm_kerberos_prof.name
  }
}