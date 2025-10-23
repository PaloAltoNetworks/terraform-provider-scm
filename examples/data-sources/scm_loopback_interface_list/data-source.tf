# Fetch a list of all loopback interfaces
data "scm_loopback_interface_list" "all_loopback_interfaces" {
  folder = "ngfw-shared"
}

# Output the raw list of all loopback interfaces
output "scm_loopback_interface_list" {
  description = "A list of all loopback interfaces from the 'All Firewalls' folder."
  value       = { for interface in data.scm_loopback_interface_list.all_loopback_interfaces.data : interface.name => interface }
}