# Define a data source for listing pbf rules
data "scm_pbf_rule_list" "paged_pbf_rules_list" {
  folder = "All"
  limit  = 10


}

# --- Output Block to Print Retrieved Data ---

output "fetched_pbf_rule_list_summary" {
  description = "Summary of rules retrieved by the SCM PBF Rule List data source."
  value = {
    total_rules_in_list = data.scm_pbf_rule_list.paged_pbf_rules_list.total
    # Access the 'data' attribute (which contains the list of rules)
    all_rules = data.scm_pbf_rule_list.paged_pbf_rules_list.data
  }
}
