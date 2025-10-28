# Look up layer2 sub-interface by its ID.
data "scm_layer2_subinterface" "scm_l2_subinterface_ds" {
  id = "88f730d1-6577-492b-88a6-73d4a513dc76"
}

# Output various attributes from the found layer2 sub-interface to verify the lookups were successful.
output "layer2_subinterface_data_source_results" {
  description = "The details of the layer2 sub-interface read from the data source."
  value = {
    id               = data.scm_layer2_subinterface.scm_l2_subinterface_ds.id
    name             = data.scm_layer2_subinterface.scm_l2_subinterface_ds.name
    comment          = data.scm_layer2_subinterface.scm_l2_subinterface_ds.comment
    vlan_tag         = data.scm_layer2_subinterface.scm_l2_subinterface_ds.vlan_tag
    parent_interface = data.scm_layer2_subinterface.scm_l2_subinterface_ds.parent_interface
    folder           = data.scm_layer2_subinterface.scm_l2_subinterface_ds.folder
  }
}