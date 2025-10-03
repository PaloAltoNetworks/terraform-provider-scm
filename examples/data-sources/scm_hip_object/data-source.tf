# Look up a single HIP Profile by its ID.
data "scm_hip_object" "scm_hip_object_ds" {
  id = "aba16b3c-8d43-4bac-aa76-572f1d36dbc5"
}

# Output the details of the single HIP Profile object found.
output "hip_objects_ds_result" {
  description = "The details of the specific HIP Object."
  value       = data.scm_hip_object.scm_hip_object_ds
}
