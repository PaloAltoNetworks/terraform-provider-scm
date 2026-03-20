resource "scm_system_match_list" "system_match_list" {
  name             = "test-system-list"
  description      = "System match list for capturing system-level events and forwarding to monitoring platform"
  folder           = "ngfw-shared"
  filter           = "All Logs"
  send_syslog      = ["test-syslog"]
  send_http        = ["some-http-profile"]
  send_snmptrap    = ["snmp_test"]
  send_email       = ["test-email"]
  send_to_panorama = false
}