resource "scm_kerberos_server_profile" "scm_kerberos_server_profile_1" {
  folder = "All"
  name   = "kerberos-server-prof-1"
  server = [
    {
      name = "server_a"
      host = "$test_ip"
    }
  ]
}

resource "scm_kerberos_server_profile" "scm_kerberos_server_profile_2" {
  folder = "All"
  name   = "kerberos-server-prof-2"
  server = [
    {
      name = "server_a"
      host = "host_a"
      port = 120
    }
  ]
}

resource "scm_kerberos_server_profile" "scm_kerberos_server_profile_3" {
  folder = "All"
  name   = "kerberos-server-prof-3"
  server = [
    {
      name = "server_a"
      host = "host_a"
      port = 1
    },
    {
      name = "server_b"
      host = "host_b"
      port = 65535
    },
    {
      name = "server_c"
      host = "192.100.50.135"
      port = 45
    }
  ]
}