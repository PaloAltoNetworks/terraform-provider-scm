resource "scm_authentication_sequence" "test_sequence" {
  name                    = "test_auth_sequence_1"
  folder                  = "All"
  authentication_profiles = ["test_auth_profile"] // authentication_profiles added should exist 
  use_domain_find_profile = false
}

data "scm_authentication_sequence" "sequence_data" {
  id = scm_authentication_sequence.test_sequence.id

}

output "scm_authentication_sequence" {
  description = "The UUID of the fetched authentication prfile."
  value       = data.scm_authentication_sequence.sequence_data.id
}

output "fetched_sequence" {
  description = "Fetched Auth Profile"
  value       = data.scm_authentication_sequence.sequence_data
}