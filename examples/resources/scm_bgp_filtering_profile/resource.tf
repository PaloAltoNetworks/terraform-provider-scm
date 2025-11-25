resource "scm_bgp_filtering_profile" "scm_bgp_filtering_profile" {
  folder = "ngfw-shared"
  name   = "scm_bgp_filtering_profile"
  ipv4   = {}
}

resource "scm_bgp_filtering_profile" "scm_bgp_filtering_profile_complex" {
  folder = "ngfw-shared"
  name   = "scm_bgp_filtering_profile_complex"
  ipv4 = {
    unicast = {
      filter_list = {
        inbound = "scm_filter_list"
      }
      inbound_network_filters = {
        prefix_list = "scm_pl_inbound"
      }
      outbound_network_filters = {
        distribute_list = "scm_distribute_list"
      }
      route_maps = {
        inbound  = "scm_rm_inbound"
        outbound = "scm_rm_outbound"
      }
    }
  }
}