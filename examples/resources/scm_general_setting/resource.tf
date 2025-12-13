resource "scm_general_setting" "gs_example" {
  folder = "All"
  general = {
    domain           = "foo.com"
    login_banner     = "Test Banner"
    ack_login_banner = false
    locale           = "en"
    geo_location = {
      latitude  = "37.38314"
      longitude = "-121.98306"
    }
    timezone = "America/Los_Angeles"
    setting = {
      management = {
        auto_acquire_commit_lock            = true
        enable_certificate_expiration_check = false
      }
      auto_mac_detect     = false
      tunnel_acceleration = true
      fail_open           = false
    }
  }
}