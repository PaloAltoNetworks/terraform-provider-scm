resource "scm_content_id_setting" "cid_example" {
  folder = "All"
  content_id = {
    allow_forward_decrypted_content = true
    extended_capture_segment        = 6
    tcp_bypass_exceed_queue         = true
    udp_bypass_exceed_queue         = true
    allow_http_range                = true
    x_forwarded_for                 = "0"
    strip_x_fwd_for                 = false
    application = {
      bypass_exceed_queue = false
    }
  }

}