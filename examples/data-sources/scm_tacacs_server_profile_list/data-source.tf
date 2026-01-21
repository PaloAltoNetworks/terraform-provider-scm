# Look up the multi-tag address object by its ID.
data "scm_tacacs_server_profile_list" "scm_tacacs_server_profiles" {
  folder = "All"
}

# Output various attributes from the found objects to verify the lookups were successful.
output "profile_data_source_results" {
  description = "The details of the tacacs server profile objects read from the data sources."
  value       = data.scm_tacacs_server_profile_list.scm_tacacs_server_profiles.data
  sensitive   = true
}

# Output just the names using a for loop
output "profile_names" {
  description = "Names of all TACACS server profiles"
  value       = [for profile in data.scm_tacacs_server_profile_list.scm_tacacs_server_profiles.data : profile.name]
}
