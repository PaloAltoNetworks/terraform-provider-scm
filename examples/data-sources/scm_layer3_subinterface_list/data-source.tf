# Fetch a list of all layer3 sub-interfaces
data "scm_layer3_subinterface_list" "all_layer3_subinterfaces" {
  folder = "ngfw-shared"
}

# Output the raw list of all layer3 sub-interfaces
output "scm_layer3_subinterface_list" {
  description = "A list of all layer3 sub-interfaces from the 'All Firewalls' folder."
  value       = { for interface in data.scm_layer3_subinterface_list.all_layer3_subinterfaces.data : interface.name => interface }
}