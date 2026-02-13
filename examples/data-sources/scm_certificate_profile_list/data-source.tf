# Fetch a list of all certificate profile objects from the "All" folder.
data "scm_certificate_profile_list" "all_profiles" {
  folder = "All"
}

# Output the raw list of all certificate profile objects found.
output "certificate_profiles_list" {
  description = "A list of all certificate profile objects from the All folder."
  # This directly outputs the list of data from the data source
  value = data.scm_certificate_profile_list.all_profiles.data
}
