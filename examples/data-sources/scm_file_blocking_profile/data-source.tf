#
# Data source to retrieve a single SCM File Blocking Profile object by its ID.
#

# Replace the ID with the UUID of the SCM File Blocking Profile you want to find.
data "scm_file_blocking_profile" "scm_file_blocking_prof" {
  id = "f32697f8-a98b-4097-b249-22c89f7d8f7f"
}

# Output the details of the single SCM File Blocking Profile object found.
output "scm_file_blocking_profile_details" {
  value = {
    profile_id = data.scm_file_blocking_profile.scm_file_blocking_prof.id
    folder     = data.scm_file_blocking_profile.scm_file_blocking_prof.folder
    name       = data.scm_file_blocking_profile.scm_file_blocking_prof.name
  }
}