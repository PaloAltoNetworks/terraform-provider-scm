# Fetch a list of all anti-spyware profiles
data "scm_anti_spyware_signature_list" "all_anti_spyware_signatures" {
  folder = "All"
}

# Output the raw list of all anti-spyware profiles
output "scm_anti_spyware_signature_list" {
  description = "A list of all anti-spyware signature from the Shared folder."
  value       = data.scm_anti_spyware_signature_list.all_anti_spyware_signatures.data
}