# Define a data source for listing NAT rules
data "scm_nat_rule_list" "paged_nat_rules_list" {
  folder   = "All"
  limit    = 10
  offset   = 10
  position = "pre"
}

# --- Output Block to Print Retrieved Data ---

output "fetched_NAT_rule_list_summary" {
  description = "Summary of rules retrieved by the SCM NAT Rule List data source."
  value = {
    total_rules_in_list = data.scm_nat_rule_list.paged_nat_rules_list.total
    # Access the 'data' attribute (which contains the list of rules)
    all_rules = data.scm_nat_rule_list.paged_nat_rules_list.data
  }
}
