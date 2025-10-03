# Look up a single certificate profile object by its ID.
# The ID used here is from the API response log you provided.
data "scm_certificate_profile" "scm_certificate_profile_ds" {
  id = "8e64859b-eba9-4e25-9005-754c90c2b02d"
}

# Output the details of the single application object found.
output "example_cp_ds_result" {
  description = "The details of the specific certificate profile object read from the data source."
  value       = data.scm_certificate_profile.scm_certificate_profile_ds
}