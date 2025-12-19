
data "scm_service_route" "settings_data" {
  id = "a232c048-98d6-4507-a299-bb10f66fca01"

}

output "fetched_settings" {
  description = "Fetched Service Route Settings"
  value       = data.scm_service_route.settings_data
}