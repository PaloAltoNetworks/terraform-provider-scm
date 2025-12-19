data "scm_general_setting" "settings_data" {
  id = "f7e61db6-b03a-493b-a70d-da39ed2e21b0"

}

output "fetched_settings" {
  description = "Fetched General Settings"
  value       = data.scm_general_setting.settings_data
}