resource "scm_tcp_setting" "tcp_example" {
  # Corresponds to = folder = All
  folder = "All"

  tcp = {
    bypass_exceed_oo_queue = false
    allow_challenge_ack    = true,
    check_timestamp_option = true,
    asymmetric_path        = "drop",
    urgent_data            = "clear",
    drop_zero_flag         = true,
    strip_mptcp_option     = true,
    siptcp_cleartext_proxy = "0",
    tcp_retransmit_scan    = true
  }
}