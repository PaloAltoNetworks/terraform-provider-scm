# Look up BGP Auth Profile by its ID.
data "scm_bgp_auth_profile" "scm_bgp_auth_profile_ds" {
  id = "f2ffd626-e92d-4de6-8ac1-37742fe80fb9"
}

# Output various attributes from the found bgp auth profile to verify the lookups were successful.
output "bgp_auth_profile_data_source_results" {
  description = "The details of the bgp auth profile read from the data source."
  value = {
    id     = data.scm_bgp_auth_profile.scm_bgp_auth_profile_ds.id
    name   = data.scm_bgp_auth_profile.scm_bgp_auth_profile_ds.name
    secret = data.scm_bgp_auth_profile.scm_bgp_auth_profile_ds.secret
    folder = data.scm_bgp_auth_profile.scm_bgp_auth_profile_ds.folder
  }
  sensitive = true
}