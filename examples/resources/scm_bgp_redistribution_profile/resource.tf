resource "scm_bgp_redistribution_profile" "scm_bgp_redistribution_profile" {
  folder = "ngfw-shared"
  name   = "scm_bgp_redistribution_profile"
  ipv4 = {
    unicast = {
      static = {
        enable = true
        metric = 10
      }
      connected = {
        enable = true
        metric = 10
      }
      ospf = {
        enable = false
      }
    }
  }
}