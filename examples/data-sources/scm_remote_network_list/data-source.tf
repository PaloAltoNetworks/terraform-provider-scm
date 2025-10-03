#
# Data source to retrieve a list of remote_network objects.
#

# Look up a list of Remote Networks, filtering by folder.
data "scm_remote_network_list" "example" {
  folder = "Remote Networks"
}

# Output the list of Remote Network objects found.
# We access the ".data" attribute which contains the list.
output "remote_network_list" {
  description = "A list of Remote Network objects."
  value       = [for network in data.scm_remote_network_list.example.data : network.name]
}