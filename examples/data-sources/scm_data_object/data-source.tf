# Replace the ID with the UUID of the SCM Data Object you want to find.
data "scm_data_object" "scm_data_object_get" {
  id = "b1398675-254e-4eff-8050-007ef2f9c0a1"
}

# Output the details of the single SCM Dataobject found.
output "scm_data_object_details" {
  description = "The pattern_type of the single data object fetched by ID."
  value       = data.scm_data_object.scm_data_object_get.pattern_type
}