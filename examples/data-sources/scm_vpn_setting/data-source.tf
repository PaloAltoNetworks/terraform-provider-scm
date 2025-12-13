data "scm_vpn_setting" "settings_data" {
  id = "7e4287ab-dfec-48bc-866d-8fb2ae3e1c5f"

}

output "fetched_settings" {
  description = "Fetched VPN Settings"
  value       = data.scm_vpn_setting.settings_data
}