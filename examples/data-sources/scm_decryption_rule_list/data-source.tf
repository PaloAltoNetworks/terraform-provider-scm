data "scm_decryption_rule_list" "paged_rules_list" {
  folder   = "All"
  offset   = 10
  position = "pre"
}

output "fetched_rule_list_summary" {
  description = "Summary of rules retrieved by the SCM Decryption Rule List data source."
  value = {
    count_of_rules_fetched = data.scm_decryption_rule_list.paged_rules_list.total
    # We access the list of rules using the 'data' attribute, and then take the first one
    first_rule_name = data.scm_decryption_rule_list.paged_rules_list.data[0].name
  }
}