#
# Data source to retrieve a single SCM Decryption Profile object by its ID.
#

# Replace the ID with the UUID of the SCM Decryption Profile you want to find.
data "scm_decryption_profile" "scm_dp_profile" {
  id = "c7629092-d286-400b-ba3f-1d57b8065645"
}

# Output the details of the single SCM Decryption Profile object found.
output "scm_decryption_profile_details" {
  value = {
    profile_id = data.scm_decryption_profile.scm_dp_profile.id
    folder     = data.scm_decryption_profile.scm_dp_profile.folder
    name       = data.scm_decryption_profile.scm_dp_profile.name
  }
}