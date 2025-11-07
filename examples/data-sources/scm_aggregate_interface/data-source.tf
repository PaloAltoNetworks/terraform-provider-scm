# Look up aggregate interface by its ID.
data "scm_aggregate_interface" "scm_aggregate_interface_ds" {
  id = "ddad1e64-0b64-41a4-b361-c6199761a8f1"
}

# Output various attributes from the found aggregate interface to verify the lookups were successful.
output "aggregate_interface_data_source_results" {
  description = "The details of the aggregate interface read from the data source."
  value = {
    id      = data.scm_aggregate_interface.scm_aggregate_interface_ds.id
    name    = data.scm_aggregate_interface.scm_aggregate_interface_ds.name
    comment = data.scm_aggregate_interface.scm_aggregate_interface_ds.comment
    layer3  = data.scm_aggregate_interface.scm_aggregate_interface_ds.layer3
    folder  = data.scm_aggregate_interface.scm_aggregate_interface_ds.folder
  }
}
