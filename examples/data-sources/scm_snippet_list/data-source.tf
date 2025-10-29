data "scm_snippet_list" "all_snippets" {
  limit = 10
}

# Output the raw list of all application objects
output "snippets_list_raw" {
  description = "A list of all tags objects from the Shared folder."

  # This directly outputs the list of data from the data source
  value = data.scm_snippet_list.all_snippets.data
}