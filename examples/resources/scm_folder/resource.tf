#
# Creates a folder under the "All Firewalls" aka "ngfw-shared" folder
#
resource "scm_folder" "scm_folder_example" {
  name        = "scm_folder_example"
  parent      = "ngfw-shared"
  description = "Managed by Terraform"
}

#
# Creates a folder under the "scm_folder_example" folder created beforehand
#
resource "scm_folder" "scm_nested_folder_example" {
  name        = "scm_nested_folder_example"
  parent      = "scm_folder_example"
  description = "Managed by Terraform"
  depends_on  = [scm_folder.scm_folder_example]
}

#
# Creates a snippet that will be associated to a folder
#
resource "scm_snippet" "scm_snippet_example" {
  name        = "scm_snippet_example"
  description = "Managed by Terraform"
}

#
# Creates a folder with an attached snippet
#

resource "scm_folder" "scm_folder_with_snippets" {
  name        = "scm_folder_with_snippets"
  parent      = "ngfw-shared"
  description = "Managed by Terraform"
  snippets    = [scm_snippet.scm_snippet_example.name]
}

#
# Creates a label that will be associated to a folder
#
resource "scm_label" "scm_label_example" {
  name        = "scm_label_example"
  description = "Managed by Terraform"
}

#
# Creates a folder with an attached label
#
resource "scm_folder" "scm_folder_with_label" {
  name        = "scm_folder_with_label"
  parent      = "ngfw-shared"
  description = "Managed by Terraform"
  labels      = [scm_label.scm_label_example.name]
  depends_on  = [scm_label.scm_label_example]
}
