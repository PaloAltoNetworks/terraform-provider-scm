resource "scm_dhcp_interface" "dhcp_server_example" {
  folder = "ngfw-shared"
  name   = "$test-interface-must-exist"
  server = {
    ip_pool  = ["10.10.10.10-10.10.10.200"]
    mode     = "auto"
    probe_ip = true
    option = {
      dns = {
        primary   = "8.8.8.8"
        secondary = "1.1.1.1"
      }
      dns_suffix  = "example.local"
      gateway     = "10.10.10.1"
      subnet_mask = "255.255.255.0"
      pop3_server = "192.0.2.25"
      smtp_server = "192.0.2.26"
      lease = {
        timeout = 3600
      }
      ntp = {
        primary   = "129.6.15.28"
        secondary = "132.163.96.1"
      }
      wins = {
        primary   = "192.168.1.1"
        secondary = "192.168.1.2"
      }
      nis = {
        primary   = "192.0.2.10"
        secondary = "192.0.2.11"
      }
      user_defined = [
        {
          inherited = true
          name      = "option-foo"
          code      = 224
          ip        = ["192.0.2.10"]
        }
      ]
    }
    reserved = [
      {
        description = "Reserved lease for printer"
        mac         = "aa:bb:cc:dd:ee:ff"
        name        = "10.10.10.50"
      }
    ]
  }
}