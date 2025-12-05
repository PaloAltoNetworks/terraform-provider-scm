data "scm_authentication_portal_list" "all_portals" {
  # The 'folder' is required to specify where to look for the profiles.
  limit  = 10
  folder = "All"
}

# -----------------------------------------------------------------------------
# OUTPUT: Display the fetched list data
# -----------------------------------------------------------------------------

output "fetched_portal_list_summary" {
  description = "Summary of profiles retrieved by the SCM Authentication Portals List data source."
  value = {
    count_of_rules_fetched = data.scm_authentication_portal_list.all_portals.total
    # We access the list of profiles using the 'data' attribute, and then take the first one
    value = data.scm_authentication_portal_list.all_portals.data[0]
  }
}