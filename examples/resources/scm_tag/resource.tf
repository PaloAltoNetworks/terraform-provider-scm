resource "scm_tag" "example" {
  folder   = "Shared"
  name     = "myColor"
  color    = "Green"
  comments = "Made by Terraform"
}
