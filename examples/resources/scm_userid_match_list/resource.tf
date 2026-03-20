resource "scm_userid_match_list" "userid_match_list" {
  name             = "test-userid-match-list"
  description      = "User ID match list for monitoring and forwarding user authentication events to external systems"
  folder           = "ngfw-shared"
  filter           = "All Logs"
  send_syslog      = ["test-syslog"]
  send_http        = ["some-http-profile"]
  send_snmptrap    = ["snmp_test"]
  send_email       = ["test-email"]
  quarantine       = false
  send_to_panorama = false
}
