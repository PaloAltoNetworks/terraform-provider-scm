resource "scm_interface_management_profile" "dc_postman_profile" {
  # Required Field
  name = "cd"

  # Contextual Field
  folder = "All"

  # Permitted IP Addresses (Inner block structure)
  permitted_ip = [
    {
      name = "10.0.0.0/24"
    },
    {
      name = "10.0.0.0/32"
    }
  ]

  # Service Flags (Booleans)
  http                       = true
  https                      = false
  telnet                     = false
  ssh                        = true
  ping                       = false
  http_ocsp                  = true
  userid_service             = true
  userid_syslog_listener_ssl = true
  userid_syslog_listener_udp = true
  response_pages             = false
}