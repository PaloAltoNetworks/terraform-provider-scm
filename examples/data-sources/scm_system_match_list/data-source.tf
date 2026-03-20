# Look up the system match list by its ID.
data "scm_system_match_list" "system_match_list_ds" {
  id = "dc61ea2f-e7aa-4e86-a742-aa8b2aaf4bf2"
}

# Output the details of the system match list
output "system_match_list_data_source_results" {
  description = "The details of the system match list object read from the data source."
  value       = data.scm_system_match_list.system_match_list_ds
}
