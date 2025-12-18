data "scm_session_timeout_list" "all_settings" {
  # The 'folder' is required to specify where to look for the setting.
  folder = "All"
}

# -----------------------------------------------------------------------------
# OUTPUT: Display the fetched list data
# -----------------------------------------------------------------------------

output "fetched_setting_list_summary" {
  description = "Summary of settings retrieved by the SCM Session Timeout List data source."
  value = {
    count_of_settings_fetched = data.scm_session_timeout_list.all_settings.total
    value                     = data.scm_session_timeout_list.all_settings.data[0]
  }
}
