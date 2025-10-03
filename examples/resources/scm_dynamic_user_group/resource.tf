# First, create the tags that will be used in the dynamic user group's filter.
resource "scm_tag" "scm_dug_tag_1" {
  folder = "Shared"
  name   = "scm_dug_tag_1"
  color  = "Red"
}

resource "scm_tag" "scm_dug_tag_2" {
  folder = "Shared"
  name   = "scm_dug_tag_2"
  color  = "Blue"
}

resource "scm_tag" "scm_dug_tag_3" {
  folder = "Shared"
  name   = "scm_dug_tag_3"
  color  = "Green"
}

# Create the dynamic user group that references the tags above.
resource "scm_dynamic_user_group" "scm_dug_1" {
  folder      = "Shared"
  name        = "scm_dug_1"
  description = "DUG created for Terraform"

  # This filter uses an OR condition for the first two tags and an AND for the third.
  filter = "'${scm_tag.scm_dug_tag_1.name}' or '${scm_tag.scm_dug_tag_2.name}' and '${scm_tag.scm_dug_tag_3.name}'"

  tag = [scm_tag.scm_dug_tag_1.name]

  # The depends_on block ensures the tags are created before the group.
  depends_on = [
    scm_tag.scm_dug_tag_1,
    scm_tag.scm_dug_tag_2,
    scm_tag.scm_dug_tag_3,
  ]
}
