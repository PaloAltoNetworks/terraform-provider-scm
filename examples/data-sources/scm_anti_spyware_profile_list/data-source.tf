# Fetch a list of all anti-spyware profiles
data "scm_anti_spyware_profile_list" "all_anti_spyware_profiles" {
  folder = "Shared"
}

# Output the raw list of all anti-spyware profiles
output "scm_anti_spyware_profile_list" {
  description = "A list of all anti-spyware profiles from the Shared folder."
  value       = data.scm_anti_spyware_profile_list.all_anti_spyware_profiles.data
}