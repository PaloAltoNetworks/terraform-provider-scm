#
# Data source to retrieve a single SCM SCEP Profile object by its ID.
#

# Replace the ID with the UUID of the SCM SCEP Profile you want to find.
data "scm_scep_profile" "scm_scep_prof" {
  id = "06c1d4ea-e2b1-44c9-bf5a-3f66c7d180a1"
}

# Output the details of the single SCM SCEP Profile object found.
output "scm_scep_profile_details" {
  value = {
    profile_id = data.scm_scep_profile.scm_scep_prof.id
    folder     = data.scm_scep_profile.scm_scep_prof.folder
    name       = data.scm_scep_profile.scm_scep_prof.name
  }
}