# Look up the hipmatch match list by its ID.
data "scm_hipmatch_match_list" "hipmatch_match_list_ds" {
  id = "ee509dcd-207e-43ca-bd35-d2f459aa7fb0"
}

# Output the details of the hipmatch match list
output "hipmatch_match_list_data_source_results" {
  description = "The details of the hipmatch match list object read from the data source."
  value       = data.scm_hipmatch_match_list.hipmatch_match_list_ds
}