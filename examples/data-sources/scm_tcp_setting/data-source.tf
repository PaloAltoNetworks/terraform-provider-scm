data "scm_tcp_setting" "settings_data" {
  id = "e7e9b9e1-e8db-4eed-b355-099a36a380c9"

}

output "fetched_settings" {
  description = "Fetched TCP Settings"
  value       = data.scm_tcp_setting.settings_data
}