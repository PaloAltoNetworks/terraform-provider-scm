# Fetch a list of all application objects.
# This data source will call the "List" API endpoint.
data "scm_application_list" "all_apps" {
  folder = "All"
}

# Output the raw list of all application objects
output "applications_list_raw" {
  description = "A list of all application objects from the Shared folder."

  # This directly outputs the list of data from the data source
  value = data.scm_application_list.all_apps.data
}