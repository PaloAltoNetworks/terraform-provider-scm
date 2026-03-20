# Look up the config match list by its ID.
data "scm_config_match_list" "config_match_list_ds" {
  id = "e8592b6e-b135-4792-b722-fc583a8f7b8e"
}

# Output the details of the config match list
output "config_match_list_data_source_results" {
  description = "The details of the config match list object read from the data source."
  value       = data.scm_config_match_list.config_match_list_ds
}