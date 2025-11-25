# Look up the "scm_snippet" tag by its id
data "scm_snippet" "scm_snippet_outputs_ds" {
  id = "b4811c43-e5f9-4b28-8316-7f18f97ba244"
}

# Output the details of the found tags to verify they were read correctly.
output "snippet_outputs" {
  value = {
    production_id          = data.scm_snippet.scm_snippet_outputs_ds.id
    production_name        = data.scm_snippet.scm_snippet_outputs_ds.name
    production_description = data.scm_snippet.scm_snippet_outputs_ds.description
    production_labels      = data.scm_snippet.scm_snippet_outputs_ds.labels
  }
}