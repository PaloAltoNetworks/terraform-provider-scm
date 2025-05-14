resource "scm_folder" "example" {
  name        = "my folder"
  parent      = "Shared"
  description = "Made by Terraform"
  snippets = [
    scm_snippet.snip1.name,
  ]
}

resource "scm_snippet" "snip1" {
  name = "foobar"
}
