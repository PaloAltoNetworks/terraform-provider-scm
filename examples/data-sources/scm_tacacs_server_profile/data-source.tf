# Look up the multi-tag address object by its ID.
data "scm_tacacs_server_profile" "scm_tacacs_server_profile5_ds" {
  id = "1967a784-402b-4c20-aa48-aab64d73cc06"
}

# Output various attributes from the found objects to verify the lookups were successful.
output "profile_data_source_results" {
  description = "The details of the tacacs server profile objects read from the data sources."
  value = {
    name                  = data.scm_tacacs_server_profile.scm_tacacs_server_profile5_ds.name
    protocol              = data.scm_tacacs_server_profile.scm_tacacs_server_profile5_ds.protocol
    timeout               = data.scm_tacacs_server_profile.scm_tacacs_server_profile5_ds.timeout
    folder                = data.scm_tacacs_server_profile.scm_tacacs_server_profile5_ds.folder
    use_single_connection = data.scm_tacacs_server_profile.scm_tacacs_server_profile5_ds.use_single_connection
  }
}

