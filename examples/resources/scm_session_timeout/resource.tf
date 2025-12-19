resource "scm_session_timeout" "st_example" {
  # Corresponds to: "folder": "All"
  folder = "All"

  session_timeouts = {
    timeout_default            = 60
    timeout_discard_default    = 60
    timeout_discard_tcp        = 90
    timeout_discard_udp        = 60
    timeout_icmp               = 6
    timeout_scan               = 10
    timeout_tcp                = 3600
    timeout_tcphandshake       = 10
    timeout_tcpinit            = 5
    timeout_tcp_half_closed    = 120
    timeout_tcp_time_wait      = 15
    timeout_tcp_unverified_rst = 30
    timeout_udp                = 30
    timeout_captive_portal     = 30
  }
}