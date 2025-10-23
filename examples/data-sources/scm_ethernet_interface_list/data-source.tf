# Fetch a list of all ethernet interfaces
data "scm_ethernet_interface_list" "all_ethernet_interfaces" {
  folder = "ngfw-shared"
}

# Output the raw list of all anti-spyware profiles
output "scm_ethernet_interface_list" {
  description = "A list of all ethernet interfaces from the 'All Firewalls' folder."
  value       = { for interface in data.scm_ethernet_interface_list.all_ethernet_interfaces.data : interface.name => interface }
}