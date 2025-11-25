# Fetch a list of all bgp auth profiles
data "scm_bgp_auth_profile_list" "all_bgp_auth_profiles" {
  folder = "ngfw-shared"
}

# Output the raw list of all bgp auth profiles
output "scm_bgp_auth_profile_list" {
  description = "A list of all bgp auth profiles from the 'Global' folder."
  value       = { for profile in data.scm_bgp_auth_profile_list.all_bgp_auth_profiles.data : profile.id => profile }
}