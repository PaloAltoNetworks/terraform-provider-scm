data "scm_motd_banner_setting" "settings_data" {
  id = "bc31452f-7f57-4c29-9dab-e1cb461799eb"

}

output "fetched_settings" {
  description = "Fetched MOTD and Banner Settings"
  value       = data.scm_motd_banner_setting.settings_data
}