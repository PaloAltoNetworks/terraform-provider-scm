resource "scm_globalprotect_match_list" "globalprotect_match_list" {
  name             = "test-globalprotect-list"
  description      = "GlobalProtect match list for monitoring VPN connection events and remote access activities"
  folder           = "ngfw-shared"
  filter           = "All Logs"
  send_syslog      = ["test-syslog"]
  send_http        = ["some-http-profile"]
  send_snmptrap    = ["snmp_test"]
  send_email       = ["test-email"]
  quarantine       = false
  send_to_panorama = false
}
