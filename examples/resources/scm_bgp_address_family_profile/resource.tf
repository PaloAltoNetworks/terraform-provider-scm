resource "scm_bgp_address_family_profile" "scm_bgp_address_family_profile" {
  folder = "ngfw-shared"
  name   = "scm_bgp_address_family_profile"
  ipv4 = {
    unicast = {
      enable = true
      allowas_in = {
        origin = {}
      }
      next_hop = {
        self = {}
      }
    }
  }
}