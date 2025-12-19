data "scm_management_interface" "settings_data" {
  id = "c94505fb-9554-4e19-b56d-e6104c908fa7"
}

output "fetched_settings" {
  description = "Fetched Management Interface Settings"
  value       = data.scm_management_interface.settings_data
}