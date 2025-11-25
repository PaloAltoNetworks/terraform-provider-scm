# Look up bgp route map by its ID.
data "scm_bgp_route_map" "scm_bgp_route_map_ds" {
  id = "f2ffd626-e92d-4de6-8ac1-37742fe80fb9"
}

# Output various attributes from the found bgp route maps to verify the lookups were successful.
output "bgp_route_map_data_source_results" {
  description = "The details of the bgp route map read from the data source."
  value = {
    id        = data.scm_bgp_route_map.scm_bgp_route_map_ds.id
    name      = data.scm_bgp_route_map.scm_bgp_route_map_ds.name
    route_map = data.scm_bgp_route_map.scm_bgp_route_map_ds.route_map
    folder    = data.scm_bgp_route_map.scm_bgp_route_map_ds.folder
  }
}