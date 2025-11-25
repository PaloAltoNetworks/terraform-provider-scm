data "scm_authentication_sequence" "sequence_data" {
  id = "1ee14ac7-760c-401f-8dbc-c887db16106a"

}

output "scm_authentication_sequence" {
  description = "The UUID of the fetched authentication prfile."
  value       = data.scm_authentication_sequence.sequence_data.id
}

output "fetched_sequence" {
  description = "Fetched Auth Profile"
  value       = data.scm_authentication_sequence.sequence_data
}