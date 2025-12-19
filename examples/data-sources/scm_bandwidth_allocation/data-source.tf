# Look up an existing Bandwidth Allocation by name
data "scm_bandwidth_allocation" "example" {
  name = "taiwan"
}

# Output the details retrieved from the API
output "allocated_bandwidth" {
  value = data.scm_bandwidth_allocation.example.allocated_bandwidth
}

output "spn_name_list" {
  value = data.scm_bandwidth_allocation.example.spn_name_list
}

output "qos_profile" {
  value = data.scm_bandwidth_allocation.example.qos.profile
}