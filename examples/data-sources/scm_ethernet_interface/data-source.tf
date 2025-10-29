# Look up ethernet interface by its ID.
data "scm_ethernet_interface" "scm_l3_intf_static_ds" {
  id = "ddad1e64-0b64-41a4-b361-c6199769a8f2"
}

# Output various attributes from the found ethernet interface to verify the lookups were successful.
output "ethernet_interface_data_source_results" {
  description = "The details of the ethernet interface read from the data source."
  value = {
    id      = data.scm_ethernet_interface.scm_l3_intf_static_ds.id
    name    = data.scm_ethernet_interface.scm_l3_intf_static_ds.name
    comment = data.scm_ethernet_interface.scm_l3_intf_static_ds.comment
    layer3  = data.scm_ethernet_interface.scm_l3_intf_static_ds.layer3
    folder  = data.scm_ethernet_interface.scm_l3_intf_static_ds.folder
  }
}
