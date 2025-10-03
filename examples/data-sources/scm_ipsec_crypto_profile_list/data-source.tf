# Example of listing all IPsec Crypto Profiles within a specific folder.
data "scm_ipsec_crypto_profile_list" "scm_ipsec_crypto_profile_list_ds" {
  folder = "Prisma Access"
  limit  = 100
}

output "all_ipsec_profiles" {
  value = data.scm_ipsec_crypto_profile_list.scm_ipsec_crypto_profile_list_ds.data
}