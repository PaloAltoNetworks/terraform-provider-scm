data "scm_bgp_routing" "all_routing" {
}

output "scm_bgp_routing_output" {
  description = "Bgp routing object"
  value       = data.scm_bgp_routing.all_routing
}