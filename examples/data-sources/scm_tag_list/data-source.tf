# Fetch a list of all application objects.
# This data source will call the "List" API endpoint.
data "scm_tag_list" "all_tags" {
  folder = "All"
}

# Output the raw list of all application objects
output "tags_list_raw" {
  description = "A list of all tags objects from the Shared folder."

  # This directly outputs the list of data from the data source
  value = data.scm_tag_list.all_tags.data
}