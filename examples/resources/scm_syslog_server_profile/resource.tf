resource "scm_syslog_server_profile" "scm_syslog_server_prof_1" {
  folder = "All"
  name   = "syslog-server-prof-base"
  server = [
    {
      name   = "Server-Primary"
      server = "192.168.1.10"
    }
  ]
}

resource "scm_syslog_server_profile" "scm_syslog_server_prof_2" {
  folder = "All"
  name   = "syslog-server-prof-mixed"
  server = [
    {
      name      = "Server-Mixed"
      server    = "10.0.0.50"
      transport = "TCP"
      port      = 601
      format    = "IETF"
      facility  = "LOG_LOCAL4"
    }
  ]
  format = {
    traffic       = "$bytes"
    threat        = "$app"
    globalprotect = "$cloud"
  }
}

resource "scm_syslog_server_profile" "scm_syslog_server_prof_3" {
  folder = "All"
  name   = "syslog-server-prof-complete"
  server = [
    {
      name      = "Server-A"
      server    = "172.16.10.1"
      transport = "UDP"
      port      = 514
      format    = "BSD"
      facility  = "LOG_LOCAL7"
    },
    {
      name      = "Server-B"
      server    = "172.16.10.2"
      transport = "TCP"
      port      = 6514
      format    = "IETF"
      facility  = "LOG_LOCAL3"
    },
    {
      name      = "Server-C"
      server    = "192.168.1.10"
      transport = "UDP"
      port      = 514
      format    = "BSD"
      facility  = "LOG_USER"
    },
  ]
  format = {
    escaping = {
      escape_character   = "*"
      escaped_characters = "&\\#"
    },
    traffic       = "$actionflags"
    threat        = "$error + $errorcode"
    wildfire      = "$client_os"
    url           = "$type"
    data          = "$srcregion"
    gtp           = "$time_generated"
    sctp          = "$device_name and $contenttype"
    tunnel        = "$tunnel_type"
    auth          = "$hostid + $portal"
    userid        = "$hostid and $location"
    iptag         = "dg_hier_level_1"
    decryption    = "dg_hier_level_2"
    config        = "dg_hier_level_3"
    system        = "$vsys_name + $status"
    globalprotect = "default"
    hip_match     = "custom"
    correlation   = "custom"
  }
}