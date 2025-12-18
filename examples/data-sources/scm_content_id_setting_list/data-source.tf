data "scm_content_id_setting_list" "all_settings" {
  # The 'folder' is required to specify where to look for the setting.
  folder = "Prisma Access"
}

# -----------------------------------------------------------------------------
# OUTPUT: Display the fetched list data
# -----------------------------------------------------------------------------

output "fetched_setting_list_summary" {
  description = "Summary of settings retrieved by the SCM Content ID List data source."
  value = {
    count_of_settings_fetched = data.scm_content_id_setting_list.all_settings.total
    value                     = data.scm_content_id_setting_list.all_settings.data[0]
  }
}