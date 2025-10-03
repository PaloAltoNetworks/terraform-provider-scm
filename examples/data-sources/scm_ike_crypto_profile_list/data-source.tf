# Example of listing all IPsec Crypto Profiles within a specific folder.
data "scm_ike_crypto_profile_list" "all_in_folder" {
  folder = "Prisma Access"
  limit  = 100
}

output "all_ike_profiles" {
  value = data.scm_ike_crypto_profile_list.all_in_folder.data
}