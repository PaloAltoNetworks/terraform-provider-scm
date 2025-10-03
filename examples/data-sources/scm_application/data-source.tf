# Look up a single application object by its ID.
# The ID used here is from the API response log you provided.
data "scm_application" "scm_application_ds" {
  id = "bb16f631-4839-475e-8628-70585319ca75"
}

# Output the details of the single application object found.
output "application_ds_result" {
  description = "The details of the specific application object read from the data source."
  value       = data.scm_application.scm_application_ds
}
