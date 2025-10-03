data "scm_application_filter" "scm_application_filter_ds" {
  id = "52ee6475-a99c-42d7-be0a-e251c05e805b"
}

output "application_filters_data_source_results" {
  description = "The details of the application filter read by the data source."
  value       = data.scm_application_filter.scm_application_filter_ds
}