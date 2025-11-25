# Fetch a list of all bgp redistribution profiles
data "scm_bgp_redistribution_profile_list" "all_bgp_redistribution_profiles" {
  folder = "ngfw-shared"
}

# Output the raw list of all bgp redistribution profiles
output "scm_bgp_redistribution_profile_list" {
  description = "A list of all bgp redistribution profiles from the 'All Firewalls' folder."
  value       = { for profile in data.scm_bgp_redistribution_profile_list.all_bgp_redistribution_profiles.data : profile.id => profile }
}