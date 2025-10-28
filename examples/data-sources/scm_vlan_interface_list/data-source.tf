# Fetch a list of all vlan interfaces
data "scm_vlan_interface_list" "all_vlan_interfaces" {
  folder = "ngfw-shared"
}

# Output the raw list of all vlan interfaces
output "scm_vlan_interface_list" {
  description = "A list of all vlan interfaces from the 'All Firewalls' folder."
  value       = { for interface in data.scm_vlan_interface_list.all_vlan_interfaces.data : interface.name => interface }
}