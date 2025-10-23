# Fetch a list of all tunnel interfaces
data "scm_tunnel_interface_list" "all_tunnel_interfaces" {
  folder = "ngfw-shared"
}

# Output the raw list of all tunnel interfaces
output "scm_tunnel_interface_list" {
  description = "A list of all tunnel interfaces from the 'All Firewalls' folder."
  value       = { for interface in data.scm_tunnel_interface_list.all_tunnel_interfaces.data : interface.name => interface }
}