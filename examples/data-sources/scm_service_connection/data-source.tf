
#------------------------------------------------------
# Data Soruce
#------------------------------------------------------
data "scm_service_connection" "created_conn_lookup" {
  id = "3d07bda7-2cfa-4fdc-b504-cd82847b2ec3"
}

output "created_service_connection_id" {
  description = "The provider-assigned unique ID of the newly created service connection."
  value       = data.scm_service_connection.created_conn_lookup.id
}

output "created_service_connection_region" {
  description = "The region assigned to the Service Connection."
  value       = data.scm_service_connection.created_conn_lookup.region
}

output "created_service_connection_subnets" {
  description = "The subnets defined for the Service Connection."
  value       = data.scm_service_connection.created_conn_lookup.subnets
}
