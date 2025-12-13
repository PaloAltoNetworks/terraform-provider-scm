data "scm_session_timeout" "settings_data" {
  id = "c86b4a2c-1621-4c9c-9f9c-c9798bc9da87"

}

output "fetched_settings" {
  description = "Fetched Session Timeout Settings"
  value       = data.scm_session_timeout.settings_data
}