data "scm_authentication_profile_list" "all_profiles" {
  # The 'folder' is required to specify where to look for the profiles.
  limit  = 10
  folder = "All"
}

# -----------------------------------------------------------------------------
# OUTPUT: Display the fetched list data
# -----------------------------------------------------------------------------

output "fetched_profile_list_summary" {
  description = "Summary of profiles retrieved by the SCM Authentication Profile List data source."
  value = {
    count_of_rules_fetched = data.scm_authentication_profile_list.all_profiles.total
    # We access the list of profiles using the 'data' attribute, and then take the first one
    first_rule_name = data.scm_authentication_profile_list.all_profiles.data[0].name
  }
}