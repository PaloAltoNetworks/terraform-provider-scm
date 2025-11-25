data "scm_authentication_sequence_list" "all_sequences" {
  # The 'folder' is required to specify where to look for the sequences.
  limit  = 10
  folder = "All"
}

# -----------------------------------------------------------------------------
# OUTPUT: Display the fetched list data
# -----------------------------------------------------------------------------

output "fetched_sequence_list_summary" {
  description = "Summary of profiles retrieved by the SCM Authentication Profile List data source."
  value = {
    count_of_sequences_fetched = data.scm_authentication_sequence_list.all_sequences.total
    # We access the list of sequences using the 'data' attribute, and then take the first one
    first_rule_name = data.scm_authentication_sequence_list.all_sequences.data[0].name
  }
}