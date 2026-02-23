#
# Creates various resources used for subsequent examples
#

resource "scm_route_prefix_list" "scm_route_prefix_list" {
  folder      = "ngfw-shared"
  name        = "scm_bgp_prefix_list"
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

#
# Creates a bgp route map that sets no-export community for traffic matching prefix-list
# Requires: scm_bgp_prefix_list
#

resource "scm_bgp_route_map" "scm_bgp_route_map" {
  folder      = "ngfw-shared"
  name        = "scm_bgp_route_map"
  description = "Managed by Terraform"
  route_map = [
    {
      name        = 10
      description = "No Export"
      match = {
        ipv4 = {
          address = {
            prefix_list = "scm_bgp_prefix_list"
          }
        }
      }
      set = { regular_community = ["no-export"] }
    }
  ]
  depends_on = [scm_route_prefix_list.scm_route_prefix_list]
}
