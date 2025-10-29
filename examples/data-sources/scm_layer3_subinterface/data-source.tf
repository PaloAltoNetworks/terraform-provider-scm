# Look up layer3 sub-interface by its ID.
data "scm_layer3_subinterface" "scm_l3_subinterface_ds" {
  id = "88f730d1-6577-492b-88a6-73d4a513dc76"
}

# Output various attributes from the found layer3 sub-interface to verify the lookups were successful.
output "layer3_subinterface_data_source_results" {
  description = "The details of the layer3 sub-interface read from the data source."
  value = {
    id               = data.scm_layer3_subinterface.scm_l3_subinterface_ds.id
    name             = data.scm_layer3_subinterface.scm_l3_subinterface_ds.name
    comment          = data.scm_layer3_subinterface.scm_l3_subinterface_ds.comment
    ip               = data.scm_layer3_subinterface.scm_l3_subinterface_ds.ip
    tag              = data.scm_layer3_subinterface.scm_l3_subinterface_ds.tag
    parent_interface = data.scm_layer3_subinterface.scm_l3_subinterface_ds.parent_interface
    folder           = data.scm_layer3_subinterface.scm_l3_subinterface_ds.folder
  }
}
