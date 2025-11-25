# Look up bgp address family profile by its ID.
data "scm_bgp_address_family_profile" "scm_bgp_address_family_profile_ds" {
  id = "83ccef34-c29a-4e88-a99b-d0355440174e"
}

# Output various attributes from the found bgp address family profile to verify the lookups were successful.
output "scm_bgp_address_family_profile_ds_results" {
  description = "The details of the bgp address family profile read from the data source."
  value = {
    id     = data.scm_bgp_address_family_profile.scm_bgp_address_family_profile_ds.id
    name   = data.scm_bgp_address_family_profile.scm_bgp_address_family_profile_ds.name
    ipv4   = data.scm_bgp_address_family_profile.scm_bgp_address_family_profile_ds.ipv4
    folder = data.scm_bgp_address_family_profile.scm_bgp_address_family_profile_ds.folder
  }
}