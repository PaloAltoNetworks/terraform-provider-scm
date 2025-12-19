#
# Data source to retrieve a single SCM TLS Service Profile object by its ID.
#

# Replace the ID with the UUID of the SCM TLS Service Profile you want to find.
data "scm_tls_service_profile" "scm_tls_service_prof" {
  id = "b4d70015-5b0e-4491-a2a9-4305b01397d5"
}

# Output the details of the single SCM TLS Service Profile object found.
output "scm_tls_service_profile_details" {
  value = {
    profile_id = data.scm_tls_service_profile.scm_tls_service_prof.id
    folder     = data.scm_tls_service_profile.scm_tls_service_prof.folder
    name       = data.scm_tls_service_profile.scm_tls_service_prof.name
  }
}

