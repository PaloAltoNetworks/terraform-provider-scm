resource "scm_label" "scm_label_1" {
  name = "scm_label"
}


resource "scm_snippet" "scm_snippet_1" {
  name        = "scm_snippet"
  description = "Adding a Description from Terraform"
  labels      = [scm_label.scm_label_1.name]
}

# Look up the "scm_snippet" tag by its id
data "scm_snippet" "scm_snippet_outputs_ds" {
  id = scm_snippet.scm_snippet_1.id
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