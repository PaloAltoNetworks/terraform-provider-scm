data "scm_device_redistribution_collector" "settings_data" {
  id = "8c1f7d95-83bd-4ae6-877c-89e1a212ef14"

}

output "fetched_settings" {
  description = "Fetched Device Redistribution Collector Settings"
  value       = data.scm_device_redistribution_collector.settings_data
}