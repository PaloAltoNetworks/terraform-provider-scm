# Fetch a list of all layer2 sub-interfaces
data "scm_layer2_subinterface_list" "all_layer2_subinterfaces" {
  folder = "ngfw-shared"
}

# Output the raw list of all layer2 sub-interfaces
output "scm_layer2_subinterface_list" {
  description = "A list of all layer2 sub-interfaces from the 'All Firewalls' folder."
  value       = { for interface in data.scm_layer2_subinterface_list.all_layer2_subinterfaces.data : interface.name => interface }
}