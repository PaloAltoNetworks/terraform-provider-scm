# This file is embedded using go:embed
# Tags for organizing and categorizing resources
resource "scm_tag" "scm_tag_1" {
  folder = "ngfw-shared"

  name  = "scm_tag_1"
  color = "Red"
}

resource "scm_tag" "scm_tag_2" {
  folder = "ngfw-shared"

  name  = "scm_tag_2"
  color = "Blue"
}

resource "scm_tag" "scm_tag_3" {
  folder = "ngfw-shared"

  name  = "scm_tag_3"
  color = "Orange"
}