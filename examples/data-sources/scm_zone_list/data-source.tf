# Fetch a list of all zones
data "scm_zone_list" "all_zones" {
  folder = "ngfw-shared"
}

# Output the raw list of all zones
output "scm_zone_list" {
  description = "A list of all zones from the 'All Firewalls' folder."
  value       = { for zone in data.scm_zone_list.all_zones.data : zone.name => zone }
}