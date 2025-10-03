# Look up the address group by its ID.
data "scm_address_group" "scm_address_group_ds" {
  id = "99802bce-76c6-42c9-801e-e2e4529bb335"
}

# Output the static members of the address group.
output "address_group_outputs" {
  value = {
    group_id    = data.scm_address_group.scm_address_group_ds.id
    folder      = data.scm_address_group.scm_address_group_ds.folder
    name        = data.scm_address_group.scm_address_group_ds.name
    description = data.scm_address_group.scm_address_group_ds.description
    static      = data.scm_address_group.scm_address_group_ds.static
  }
}