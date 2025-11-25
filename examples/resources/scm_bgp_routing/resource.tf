resource "scm_bgp_routing" "example" {
  # Enum: no-asymmetric-routing, asymmetric-routing-only, asymmetric-routing-with-load-share
  backbone_routing = "no-asymmetric-routing"

  accept_route_over_sc       = true
  add_host_route_to_ike_peer = false
  withdraw_static_route      = true

  outbound_routes_for_services = ["10.0.0.0/9"]

  # Nested Object (OneOf logic in spec, usually handled as optional blocks)
  routing_preference = {
    default = {}
    # OR
    # hot_potato_routing = {}
  }
}
