data "scm_content_id_setting" "settings_data" {
  id = "d154bf67-5e4b-4907-b892-1d93cd8cafbc"

}

output "fetched_settings" {
  description = "Fetched Content ID Settings"
  value       = data.scm_content_id_setting.settings_data
}