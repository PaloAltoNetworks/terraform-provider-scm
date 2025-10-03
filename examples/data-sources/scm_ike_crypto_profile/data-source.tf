# Example of looking up an individual IPsec Crypto Profile by its ID.
data "scm_ike_crypto_profile" "scm_ike_crypto_profile_ds" {
  id = "f3a1251a-bb9e-437d-8048-f5d54617d0be"
}

# Outputs to display the results of the data source lookups.
output "ike_profile_by_id" {
  value = data.scm_ike_crypto_profile.scm_ike_crypto_profile_ds
}