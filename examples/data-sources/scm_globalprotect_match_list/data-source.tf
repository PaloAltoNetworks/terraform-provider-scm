# Look up the globalprotect match list by its ID.
data "scm_globalprotect_match_list" "globalprotect_match_list_ds" {
  id = "89c5fc2d-dba8-4560-b853-4eb1ecd36025"
}

# Output the details of the globalprotect match list
output "globalprotect_match_list_data_source_results" {
  description = "The details of the globalprotect match list object read from the data source."
  value       = data.scm_globalprotect_match_list.globalprotect_match_list_ds
}