# Look up the userid match list by its ID.
data "scm_userid_match_list" "userid_match_list_ds" {
  id = "dcf09e7b-3e46-45d8-953d-30d47a49137f"
}

# Output the details of the userid match list
output "userid_match_list_data_source_results" {
  description = "The details of the userid match list object read from the data source."
  value       = data.scm_userid_match_list.userid_match_list_ds
}