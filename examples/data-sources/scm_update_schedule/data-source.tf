data "scm_update_schedule" "settings_data" {
  id = "4bde6878-8709-4231-a7b4-c51d8fb4008a"

}

output "fetched_settings" {
  description = "Fetched Update Schedule Settings"
  value       = data.scm_update_schedule.settings_data
}