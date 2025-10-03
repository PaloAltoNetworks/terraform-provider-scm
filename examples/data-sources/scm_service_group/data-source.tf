# Look up a single service group object by its ID.
# The ID used here is from the terraform.tfstate file.
data "scm_service_group" "scm_service_group_ds" {
  id = "dc430d61-52ca-44bc-a797-e65123a94134"
}

# Output the details of the single service group object found.
output "service_group_ds_result" {
  description = "The details of the specific service group object read from the data source."
  value       = data.scm_service_group.scm_service_group_ds
}