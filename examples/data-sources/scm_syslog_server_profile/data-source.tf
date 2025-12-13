#
# Data source to retrieve a single SCM Syslog Server Profile object by its ID.
#

# Replace the ID with the UUID of the SCM Syslog Server Profile you want to find.
data "scm_syslog_server_profile" "scm_syslog_server_prof" {
  id = "69f7ee97-7c0a-416d-a28d-d45929851f6e"
}

# Output the details of the single SCM Syslog Server Profile object found.
output "scm_syslog_server_profile_details" {
  value = {
    profile_id = data.scm_syslog_server_profile.scm_syslog_server_prof.id
    folder     = data.scm_syslog_server_profile.scm_syslog_server_prof.folder
    name       = data.scm_syslog_server_profile.scm_syslog_server_prof.name
  }
}