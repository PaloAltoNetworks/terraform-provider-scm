resource "scm_hipmatch_match_list" "hipmatch_match_list" {
  name             = "test-hipmatch-list"
  description      = "HIP match list for monitoring host information profile events and endpoint compliance status"
  folder           = "ngfw-shared"
  filter           = "All Logs"
  send_syslog      = ["test-syslog"]
  send_http        = ["some-http-profile"]
  send_snmptrap    = ["snmp_test"]
  send_email       = ["test-email"]
  quarantine       = false
  send_to_panorama = false
}
