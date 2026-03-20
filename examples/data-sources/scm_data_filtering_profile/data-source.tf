# Replace the ID with the UUID of the SCM Data filtering profile you want to find.
data "scm_data_filtering_profile" "scm_data_filtering_profile_get" {
  id = "d50c57e4-3a3e-419e-bd41-33fdd1c56a32"
}

# Output the details of the single SCM Data filtering profile object found.
output "scm_data_filtering_profile_details" {
  description = "The rules of the single data filtering profile fetched by ID."
  value       = data.scm_data_filtering_profile.scm_data_filtering_profile_get.rules
}


