resource "scm_auto_vpn_setting" "example" {
  # The folder/scope is usually handled by the provider configuration
  # or implicitly if this is a global singleton.

  enable_mesh_between_hubs = true

  # Required list of CIDRs
  vpn_address_pool = [
    "10.91.0.0/25"
  ]

  # Required nested object
  as_range = {
    start = 65001
    end   = 65200
  }
}