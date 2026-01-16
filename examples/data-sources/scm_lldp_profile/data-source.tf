data "scm_lldp_profile" "profile_data" {
  id = "e46f6246-fd4a-4211-a18f-948b09f474bd"

}

output "fetched_profile" {
  description = "Fetched LLDP Profile"
  value       = data.scm_lldp_profile.profile_data
}