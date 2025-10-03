# Data source to look up a single service by its ID.
data "scm_service" "scm_service_tcp_ds" {
  id = "ff135641-6735-4d7d-85c6-3401bba9dee8"
}

# Output the details of the looked-up service.
output "service_details_tcp" {
  description = "The details of the retrieved service object."
  value       = data.scm_service.scm_service_tcp_ds
}


# Data source to look up a single service by its ID.
data "scm_service" "scm_service_udp_ds" {
  id = "e087b703-aede-437e-853e-b11576f6dcbe"
}

# Output the details of the looked-up service.
output "service_details_udp" {
  description = "The details of the retrieved service object."
  value       = data.scm_service.scm_service_udp_ds
}