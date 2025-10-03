# Fetch a list of all HIP Profile objects in the Shared folder.
data "scm_hip_profile_list" "all_shared" {
  folder = "Shared"
}

# Output the raw list of all HIP Profile objects.
output "hip_profiles_list_all_shared" {
  description = "A list of all HIP Profile objects from the Shared folder."
  value       = data.scm_hip_profile_list.all_shared.data
}

# Example of using pagination to get the first 10 HIP Profiles.
data "scm_hip_profile_list" "paginated" {
  folder = "Shared"
  limit  = 10
  offset = 0
}

output "hip_profiles_list_paginated" {
  description = "A paginated list of HIP Profile objects."
  value       = data.scm_hip_profile_list.paginated.data
}
