data "scm_authentication_rule_list" "all_pre_rules" {
  # The 'folder' is required to specify where to look for the rules.
  limit    = 10
  offset   = 15
  position = "pre"
  folder   = "All"
}

# -----------------------------------------------------------------------------
# 5. OUTPUT: Display the fetched list data
# -----------------------------------------------------------------------------

output "fetched_rule_list_summary" {
  description = "Summary of rules retrieved by the SCM Authentication Rule List data source."
  value = {
    count_of_rules_fetched = data.scm_authentication_rule_list.all_pre_rules.total
    # We access the list of rules using the 'data' attribute, and then take the first one
    first_rule_name = data.scm_authentication_rule_list.all_pre_rules.data[0].name
  }
}