# Look up the iptag match list by its ID.
data "scm_iptag_match_list" "iptag_match_list_ds" {
  id = "108dfdb1-0723-497d-9352-50642b231b4e"
}

# Output the details of the iptag match list
output "iptag_match_list_data_source_results" {
  description = "The details of the iptag match list object read from the data source."
  value       = data.scm_iptag_match_list.iptag_match_list_ds
}