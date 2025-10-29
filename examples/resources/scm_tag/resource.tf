# This file is embedded using go:embed
# Tags for organizing and categorizing resources
resource "scm_tag" "scm_tag_1" {
  folder = "All"
  name   = "scm_tag_1"
  color  = "Red"
}

resource "scm_tag" "scm_tag_2" {
  folder = "All"
  name   = "scm_tag_2"
  color  = "Blue"
}

resource "scm_tag" "scm_tag_3" {
  folder = "All"
  name   = "scm_tag_3"
  color  = "Orange"
}