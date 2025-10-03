#
# Data source to retrieve a single remote_network object.
#

# Look up a single Remote Network by its ID.
# Replace the ID with the UUID of the remote network you want to find.
data "scm_remote_network" "example" {
  id = "7fc59ec2-46b3-4a0e-9c86-9b7416426a70"
}

# Output the details of the single Remote Network object found.
output "remote_network_details" {
  description = "The details of the specific Remote Network object."
  value       = data.scm_remote_network.example
  sensitive   = true
}