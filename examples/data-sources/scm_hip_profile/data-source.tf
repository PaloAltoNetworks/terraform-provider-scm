# Look up a single HIP Profile by its ID.
data "scm_hip_profile" "scm_hip_profile_ds" {
  id = "e0a970b8-98d2-42e9-a273-53fbf67607c2"
}

# Output the details of the single HIP Profile object found.
output "hip_profile_ds_result" {
  description = "The details of the specific HIP Profile object."
  value       = data.scm_hip_profile.scm_hip_profile_ds
}
