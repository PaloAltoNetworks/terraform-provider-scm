data "scm_dhcp_interface_list" "paged_interface_list" {
  folder = "All"
  limit  = 10
}

output "fetched_interface_list_summary" {
  description = "Summary of interfaces retrieved by the SCM DHCP Interface List data source."
  value = {
    count_of_rules_fetched = data.scm_dhcp_interface_list.paged_interface_list.total
    # We access the list of rules using the 'data' attribute, and then take the first one
    first_rule_name = data.scm_dhcp_interface_list.paged_interface_list.data
  }
}