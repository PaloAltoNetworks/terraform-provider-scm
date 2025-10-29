# Look up vlan interface by its ID.
data "scm_vlan_interface" "scm_vlan_interface_ds" {
  id = "3f9382a3-5c93-46d9-ae06-a632c2d9ce0c"
}

# Output various attributes from the found vlan interface to verify the lookups were successful.
output "vlan_interface_data_source_results" {
  description = "The details of the vlan interface read from the data source."
  value = {
    id       = data.scm_vlan_interface.scm_vlan_interface_ds.id
    name     = data.scm_vlan_interface.scm_vlan_interface_ds.name
    comment  = data.scm_vlan_interface.scm_vlan_interface_ds.comment
    vlan_tag = data.scm_vlan_interface.scm_vlan_interface_ds.vlan_tag
    ip       = data.scm_vlan_interface.scm_vlan_interface_ds.ip
    folder   = data.scm_vlan_interface.scm_vlan_interface_ds.folder
  }
}
