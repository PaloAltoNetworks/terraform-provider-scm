
data "scm_session_setting" "settings_data" {
  id = "e869d51d-0d34-4c89-a674-1824cec0eeae"

}

output "fetched_settings" {
  description = "Fetched Session Settings"
  value       = data.scm_session_setting.settings_data
}