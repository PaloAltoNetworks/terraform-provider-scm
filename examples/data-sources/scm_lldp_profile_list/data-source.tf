data "scm_lldp_profile_list" "all_profiles" {
  # The 'folder' is required to specify where to look for the setting.
  folder = "All"
}

# -----------------------------------------------------------------------------
# OUTPUT: Display the fetched list data
# -----------------------------------------------------------------------------

output "fetched_profiles_list_summary" {
  description = "Summary of profiles retrieved by the SCM LLDP Profile List data source."
  value = {
    count_of_settings_fetched = data.scm_lldp_profile_list.all_profiles.total
    first_profile             = data.scm_lldp_profile_list.all_profiles.data[0]
  }
}