resource "scm_config_match_list" "config_match_list" {
  name             = "test-config-list"
  description      = "Config match list for tracking configuration changes and audit log forwarding"
  folder           = "ngfw-shared"
  filter           = "All Logs"
  send_syslog      = ["test-syslog"]
  send_http        = ["some-http-profile"]
  send_snmptrap    = ["snmp_test"]
  send_email       = ["test-email"]
  send_to_panorama = false
}
