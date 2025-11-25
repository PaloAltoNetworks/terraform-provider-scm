data "scm_authentication_profile" "profile_data" {
  id = "de491856-1987-4b53-a3f7-e4f4a52067e3"
}

output "scm_authentication_profile" {
  description = "The UUID of the fetched authentication prfile."
  value       = data.scm_authentication_profile.profile_data.id
}

output "fetched_profile" {
  description = "Fetched Auth Profile"
  value       = data.scm_authentication_profile.profile_data
}