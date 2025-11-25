# Fetch a list of all bgp route maps
data "scm_bgp_route_map_list" "all_bgp_route_maps" {
  folder = "ngfw-shared"
}

# Output the raw list of all bgp route maps
output "scm_bgp_route_map_list" {
  description = "A list of all bgp route maps from the 'All Firewalls' folder."
  value       = { for route_map in data.scm_bgp_route_map_list.all_bgp_route_maps.data : route_map.id => route_map }
}