# Look up a folder by id
data "scm_folder" "scm_folder_ds" {
  id = "0f11d0d9-df7c-45da-a60c-4d80f8422544"
}

# Output the details of the found folder to verify they were read correctly.
output "scm_folder_output" {
  value = {
    id          = data.scm_folder.scm_folder_ds.id
    name        = data.scm_folder.scm_folder_ds.name
    description = data.scm_folder.scm_folder_ds.description
    parent      = data.scm_folder.scm_folder_ds.parent
  }
}
