#
# Data source to retrieve a single SCM Log Forwarding Profile object by its ID.
#

# Replace the ID with the UUID of the SCM Log Forwarding Profile you want to find.
data "scm_log_forwarding_profile" "scm_log_forwarding_prof" {
  id = "712dc61d-94ed-45e4-82b7-b2d86669a2bf"
}

# Output the details of the single SCM Log Forwarding Profile object found.
output "scm_log_forwarding_profile_details" {
  value = {
    profile_id  = data.scm_log_forwarding_profile.scm_log_forwarding_prof.id
    folder      = data.scm_log_forwarding_profile.scm_log_forwarding_prof.folder
    name        = data.scm_log_forwarding_profile.scm_log_forwarding_prof.name
    description = data.scm_log_forwarding_profile.scm_log_forwarding_prof.description
  }
}