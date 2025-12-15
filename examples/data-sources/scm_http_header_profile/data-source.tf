#
# Data source to retrieve a single SCM HTTP Header Profile object by its ID.
#

# Replace the ID with the UUID of the SCM HTTP Header Profile you want to find.
data "scm_http_header_profile" "scm_http_header_prof" {
  id = "2733cba4-c79d-4c98-8e07-4d3cbdd0ba11"
}

# Output the details of the single SCM HTTP Header Profile object found.
output "scm_http_header_profile_details" {
  value = {
    profile_id = data.scm_http_header_profile.scm_http_header_prof.id
    folder     = data.scm_http_header_profile.scm_http_header_prof.folder
    name       = data.scm_http_header_profile.scm_http_header_prof.name
  }
}