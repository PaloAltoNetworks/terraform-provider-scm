# Look up BGP Redistribution Profile by its ID.
data "scm_bgp_redistribution_profile" "scm_bgp_redistribution_profile_ds" {
  id = "491918e9-0205-4a08-955a-7e59a38d5dc4"
}

# Output various attributes from the found bgp redistribution profile to verify the lookups were successful.
output "bgp_redistribution_profile_data_source_results" {
  description = "The details of the bgp redistribution profile read from the data source."
  value = {
    id     = data.scm_bgp_redistribution_profile.scm_bgp_redistribution_profile_ds.id
    name   = data.scm_bgp_redistribution_profile.scm_bgp_redistribution_profile_ds.name
    ipv4   = data.scm_bgp_redistribution_profile.scm_bgp_redistribution_profile_ds.ipv4
    folder = data.scm_bgp_redistribution_profile.scm_bgp_redistribution_profile_ds.folder
  }
  sensitive = true
}