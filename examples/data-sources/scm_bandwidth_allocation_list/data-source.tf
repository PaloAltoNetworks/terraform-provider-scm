data "scm_bandwidth_allocation_list" "example" {
  # According to your model schema, this parameter is Required for the list call
  name = "taiwan"

  # Optional pagination parameters
  limit  = 50
  offset = 0
}

output "allocation_list" {
  value = data.scm_bandwidth_allocation_list.example.data
}