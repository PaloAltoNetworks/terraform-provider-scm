#
# Data source to retrieve a list of SCM Data objects.
#

# Fetch a list of all SCM Data Objects in the "ngfw-shared" folder.
data "scm_data_object_list" "all_ngfw_shared" {
  folder = "ngfw-shared"
}

# Output the list of all SCM Data objects from the "ngfw-shared" folder.
output "scm_data_object_list_ngfw_shared" {
  description = "A list of all SCM Data Objects from the ngfw-shared folder."
  value       = data.scm_data_object_list.all_ngfw_shared.data
}