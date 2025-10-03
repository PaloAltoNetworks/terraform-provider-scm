# Example of looking up an individual IPsec Crypto Profile by its ID.
data "scm_ipsec_crypto_profile" "scm_ipsec_crypto_profile_ds" {
  id = "b89e8fe1-9e92-46fa-8a67-de84313128c9"
}

# Outputs to display the results of the data source lookups.
output "ipsec_profile_by_id" {
  value = data.scm_ipsec_crypto_profile.scm_ipsec_crypto_profile_ds
}