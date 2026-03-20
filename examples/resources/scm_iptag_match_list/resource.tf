resource "scm_iptag_match_list" "iptag_match_list" {
  name             = "test-iptag-list"
  description      = "IP tag match list for tracking dynamic IP address tagging events and policy enforcement"
  folder           = "ngfw-shared"
  filter           = "All Logs"
  send_syslog      = ["test-syslog"]
  send_http        = ["some-http-profile"]
  send_snmptrap    = ["snmp_test"]
  send_email       = ["test-email"]
  quarantine       = false
  send_to_panorama = false
}
