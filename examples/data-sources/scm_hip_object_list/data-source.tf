# Fetch a list of all HIP Object objects in the Shared folder.
data "scm_hip_object_list" "all_shared" {
  folder = "All"
}

# Output the raw list of all HIP Object objects.
output "hip_objects_list_all_shared" {
  description = "A list of all HIP Object objects from the Shared folder."
  value       = data.scm_hip_object_list.all_shared.data
}

# Example of using pagination to get the first 10 HIP objects.
data "scm_hip_object_list" "paginated" {
  folder = "All"
  limit  = 10
  offset = 0
}

output "hip_objects_list_paginated" {
  description = "A paginated list of HIP Object objects."
  value       = data.scm_hip_object_list.paginated.data
}
