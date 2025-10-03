# Look up the application group by its ID from the terraform.tfstate file.
data "scm_application_group" "scm_application_group_ds" {
  id = "91616221-ddeb-4b49-866d-48d64dedc056"
}

# Output the members of the application group.
output "application_group_outputs" {
  value = {
    group_id = data.scm_application_group.scm_application_group_ds.id
    folder   = data.scm_application_group.scm_application_group_ds.folder
    name     = data.scm_application_group.scm_application_group_ds.name
    members  = data.scm_application_group.scm_application_group_ds.members
  }
}