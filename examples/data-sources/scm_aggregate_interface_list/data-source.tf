# Fetch a list of all aggregate interfaces
data "scm_aggregate_interface_list" "all_aggregate_interfaces" {
  folder = "ngfw-shared"
}

# Output the raw list of all aggregate interfaces
output "scm_aggregate_interface_list" {
  description = "A list of all aggregate interfaces from the 'All Firewalls' folder."
  value       = { for interface in data.scm_aggregate_interface_list.all_aggregate_interfaces.data : interface.name => interface }
}