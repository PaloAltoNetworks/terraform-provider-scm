data "scm_service_setting_list" "all_settings" {
  # The 'folder' is required to specify where to look for the setting.
  folder = "All"
}

# -----------------------------------------------------------------------------
# OUTPUT: Display the fetched list data
# -----------------------------------------------------------------------------

output "fetched_setting_list_summary" {
  description = "Summary of settings retrieved by the SCM Service Settings List data source."
  value = {
    count_of_settings_fetched = data.scm_service_setting_list.all_settings.total
    dns_setting               = data.scm_service_setting_list.all_settings.data[0].services.dns_setting
  }
}
