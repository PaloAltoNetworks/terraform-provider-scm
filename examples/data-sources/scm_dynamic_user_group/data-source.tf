# Look up the dynamic user group by its ID.
data "scm_dynamic_user_group" "scm_dynamic_user_group_ds" {
  id = "c8ac5c18-023b-4be5-bc39-65e585cff9c7"
}

# Output the details of the dynamic user group.
output "dynamic_user_group_details" {
  value = {
    id          = data.scm_dynamic_user_group.scm_dynamic_user_group_ds.id
    name        = data.scm_dynamic_user_group.scm_dynamic_user_group_ds.name
    folder      = data.scm_dynamic_user_group.scm_dynamic_user_group_ds.folder
    description = data.scm_dynamic_user_group.scm_dynamic_user_group_ds.description
    filter      = data.scm_dynamic_user_group.scm_dynamic_user_group_ds.filter
    tags        = data.scm_dynamic_user_group.scm_dynamic_user_group_ds.tag
  }
}