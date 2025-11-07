# Look up the logical router by its ID.
data "scm_logical_router" "scm_logical_router_ds" {
  id = "b7c6f00b-b20e-4073-af1c-1f42863a5983"
}

# Output various attributes from the found logical router to verify the lookups were successful.
output "scm_logical_router_data_source_output" {
  description = "The details of the logical router read from the data source."
  value = {
    id            = data.scm_logical_router.scm_logical_router_ds.id
    name          = data.scm_logical_router.scm_logical_router_ds.name
    routing_stack = data.scm_logical_router.scm_logical_router_ds.routing_stack
    vrf           = data.scm_logical_router.scm_logical_router_ds.vrf
    folder        = data.scm_logical_router.scm_logical_router_ds.folder
  }
}