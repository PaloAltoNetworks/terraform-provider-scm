data "scm_auto_vpn_setting" "auto_vpn_setting" {
}

output "scm_auto_vpn_setting_output" {
  description = "Auto vpn setting object"
  value       = data.scm_auto_vpn_setting.auto_vpn_setting
}