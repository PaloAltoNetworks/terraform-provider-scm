resource "scm_vpn_setting" "tcp_example" {
  # Corresponds to = folder = All
  folder = "ngfw-shared"


  vpn = {
    ikev2 = {
      cookie_threshold       = 500
      max_half_opened_sa     = 65535
      certificate_cache_size = 122
    }
  }
}