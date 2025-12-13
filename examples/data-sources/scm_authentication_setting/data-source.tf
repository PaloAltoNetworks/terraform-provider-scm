data "scm_authentication_setting" "settings_data" {
  id = "f4e09263-f8bf-4a4d-a37d-b54d6530810a"

}

output "fetched_settings" {
  description = "Fetched Authentication Settings"
  value       = data.scm_authentication_setting.settings_data
}