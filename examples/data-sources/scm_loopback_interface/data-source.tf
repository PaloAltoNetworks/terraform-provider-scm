# Look up loopback interface by its ID.
data "scm_loopback_interface" "scm_loopback_intf_ds" {
  id = "ddad1e64-0b64-41a4-b361-c6199769a8f1"
}

# Output various attributes from the found loopback interface to verify the lookups were successful.
output "scm_loopback_interface_data_source_results" {
  description = "The details of the loopback interfaces read from the data source"
  value = {
    id      = data.scm_loopback_interface.scm_loopback_intf_ds.id
    name    = data.scm_loopback_interface.scm_loopback_intf_ds.name
    comment = data.scm_loopback_interface.scm_loopback_intf_ds.comment
    ip      = data.scm_loopback_interface.scm_loopback_intf_ds.ip
    folder  = data.scm_loopback_interface.scm_loopback_intf_ds.folder
  }
}
