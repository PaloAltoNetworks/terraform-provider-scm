# ------------------------------------------------------------------
# Data Source: SCM Service Connection Group (Single Lookup)
# ------------------------------------------------------------------

data "scm_service_connection_group" "group_lookup" {
  # Name is fetched from the resource creation block
  id = "1480fd9d-dae7-4bf3-94f6-4945e89b59b6"
}

output "looked_up_service_connection_group_details" {
  description = "All attributes retrieved for the Service Connection Group via the data source lookup."
  # Directly reference the data source object
  value = data.scm_service_connection_group.group_lookup
}

