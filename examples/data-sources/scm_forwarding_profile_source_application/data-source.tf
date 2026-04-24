# Look up a single forwarding profile source application by its ID.
data "scm_forwarding_profile_source_application" "example" {
  id = "12345678-1234-1234-1234-123456789012"
}

# Output the details of the forwarding profile source application.
output "forwarding_profile_source_application_result" {
  description = "The details of the specific forwarding profile source application read from the data source."
  value       = data.scm_forwarding_profile_source_application.example
}