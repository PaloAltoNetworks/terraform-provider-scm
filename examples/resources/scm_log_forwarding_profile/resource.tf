resource "scm_log_forwarding_profile" "scm_log_forwarding_profile_1" {
  folder = "All"
  name   = "scm-log-fowarding-profile-1"
  match_list = [
    {
      name     = "profile_match"
      log_type = "threat"
      filter   = "(addr in 192.50.10.10) and (addr.dst notin 192.40.50.10)"
    }
  ]
}

resource "scm_log_forwarding_profile" "scm_log_forwarding_profile_2" {
  folder      = "All"
  name        = "scm-log-fowarding-profile-2"
  description = "Log Forwarding w/ HTTP Server Profile and Syslog Server Profile"
  match_list = [
    {
      name        = "profile_match"
      log_type    = "traffic"
      filter      = "(device_name eq test_device)"
      send_http   = ["test_http"]
      send_syslog = ["syslog-server-prof-mixed", "syslog-server-prof-complete"]
    }
  ]
}

resource "scm_log_forwarding_profile" "scm_log_forwarding_profile_3" {
  folder      = "All"
  name        = "scm-log-fowarding-profile-3"
  description = "Log Forwarding w/ All Server Profiles"
  match_list = [
    {
      name          = "profile_match"
      action_desc   = "all server profiles"
      log_type      = "dns-security"
      filter        = "All Logs"
      send_http     = ["test_http", "t10"]
      send_syslog   = ["syslog-server-prof-base", "syslog-server-prof-mixed", "syslog-server-prof-complete"]
      send_snmptrap = ["snmp_test"]
      send_email    = ["email_test", "email_test_2"]
    }
  ]
}

resource "scm_log_forwarding_profile" "scm_log_forwarding_profile_4" {
  folder      = "All"
  name        = "scm-log-fowarding-profile-4"
  description = "Log Forwarding w/ Multiple Match Lists"
  match_list = [
    {
      name          = "profile_match_1"
      action_desc   = "match list for url"
      log_type      = "url"
      filter        = "(sdwan_cluster contains 123)"
      send_http     = ["t10"]
      send_syslog   = ["syslog-server-prof-base"]
      send_snmptrap = ["snmp_test"]
    },
    {
      name        = "profile_match_2"
      log_type    = "data"
      filter      = "(link_switch_2 neq lnk_2) or (pkts_received geq 100)"
      send_http   = ["t5", "t10", "t20"]
      send_syslog = ["syslog-server-prof-mixed"]
      send_email  = ["email_test", "email_test_2"]
    },
    {
      name        = "profile_match_3"
      action_desc = "match list for wildfire"
      log_type    = "wildfire"
      filter      = "(imei contains test_server)"
      send_http   = ["t5", "t10", "t20", "t22", "t24"]
      send_syslog = ["syslog-server-prof-complete"]
      send_email  = ["email_test", "email_test_2"]
    }
  ]
}