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
            network               = "198.18.1.0/24"
          }
        }
      ]
    }
  }
}