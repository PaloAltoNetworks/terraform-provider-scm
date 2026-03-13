#
# Creates an empty bgp filtering profile
#

resource "scm_bgp_filtering_profile" "scm_bgp_filtering_profile" {
  folder = "ngfw-shared"
  name   = "scm_bgp_filtering_profile"
  ipv4   = {}
}

#
# Creates a complex filtering profile that utilises previously created FL, PL and RM
#

resource "scm_bgp_filtering_profile" "scm_bgp_filtering_profile_complex" {
  folder = "ngfw-shared"
  name   = "scm_bgp_filtering_profile_complex"
  ipv4 = {
    unicast = {
      inbound_network_filters = {
        prefix_list = "scm_pl_inbound"
      }
      route_maps = {
        inbound  = "scm_rm_inbound"
        outbound = "scm_rm_outbound"
      }
    }
  }
  depends_on = [
    scm_route_prefix_list.scm_pl_inbound,
    scm_bgp_route_map.scm_rm_inbound,
    scm_bgp_route_map.scm_rm_outbound,
  ]
}

#
# Creates various resources used for scm_bgp_filtering_profile_complex
#

resource "scm_route_prefix_list" "scm_pl_inbound" {
  folder      = "ngfw-shared"
  name        = "scm_pl_inbound"
  description = "Managed by Terraform"
  type = {
    ipv4 = {
      ipv4_entry = [
        {
          name   = 10
          action = "permit"
          prefix = {
            greater_than_or_equal = 24
            network               = "any"
          }
        }
      ]
    }
  }
}

resource "scm_route_prefix_list" "scm_pl_outbound" {
  folder      = "ngfw-shared"
  name        = "scm_pl_outbound"
  description = "Managed by Terraform"
  type = {
    ipv4 = {
      ipv4_entry = [
        {
          name   = 10
          action = "permit"
          prefix = {
            greater_than_or_equal = 24
            network               = "any"
          }
        }
      ]
    }
  }
}

resource "scm_bgp_route_map" "scm_rm_inbound" {
  folder      = "ngfw-shared"
  name        = "scm_rm_inbound"
  description = "Managed by Terraform"
  route_map = [
    {
      name        = 10
      description = "No Export"
      match = {
        ipv4 = {
          address = {
            prefix_list = "scm_pl_inbound"
          }
        }
      }
      set = { regular_community = ["no-export"] }
    }
  ]
  depends_on = [scm_route_prefix_list.scm_pl_inbound]
}

resource "scm_bgp_route_map" "scm_rm_outbound" {
  folder      = "ngfw-shared"
  name        = "scm_rm_outbound"
  description = "Managed by Terraform"
  route_map = [
    {
      name        = 10
      description = "No Export"
      match = {
        ipv4 = {
          address = {
            prefix_list = "scm_pl_outbound"
          }
        }
      }
      set = { regular_community = ["no-export"] }
    }
  ]
  depends_on = [scm_route_prefix_list.scm_pl_outbound]
}