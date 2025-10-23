resource "scm_label" "scm_label_1" {
  name = "scm_label"
}

//snippet with label
resource "scm_snippet" "scm_snippet_1" {
  name        = "scm_snippet"
  description = "Adding a Description from Terraform"
  labels      = [scm_label.scm_label_1.name]
}

resource "scm_snippet" "scm_snippet_2" {
  name        = "scm_snippet_2"
  description = "Adding a Description from Terraform"
}

resource "scm_snippet" "scm_snippet_3" {
  name        = "scm_snippet_3"
  description = "Adding a Description from Terraform"
}