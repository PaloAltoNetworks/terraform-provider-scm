resource "scm_management_interface" "mi_example" {
  # Corresponds to: "folder": "All"
  folder = "All"

  # The settings for the management interface itself are defined within the interface_setting block.
  management_interface = {
    # Corresponds to: "speed_duplex": "auto-negotiate"
    speed_duplex = "auto-negotiate"

    # Corresponds to: "mtu": 1500
    mtu = 1500

    # Corresponds to: "mgmt_type": { ... }
    mgmt_type = {
      # Corresponds to: "dhcp_client": { ... }
      dhcp_client = {
        send_hostname        = true  # "send_hostname" : false
        send_client_id       = false # "send_client_id" : false
        accept_dhcp_hostname = false # "accept_dhcp_hostname" : false
        accept_dhcp_domain   = false # "accept_dhcp_domain" : false
      }
      # The JSON did not specify "static_ip" or "pppoe" so we omit them.
    }

    # Corresponds to: "service": { ... }
    service = {
      disable_http                       = false # "disable_http": false
      disable_https                      = true  # "disable_https": true
      disable_telnet                     = false # "disable_telnet": false
      disable_ssh                        = true  # "disable_ssh": true
      disable_icmp                       = false # "disable_icmp": false
      disable_snmp                       = false # "disable_snmp": false
      disable_userid_service             = false # "disable_userid_service": false
      disable_userid_syslog_listener_ssl = false # "disable_userid_syslog_listener_ssl": false
      disable_userid_syslog_listener_udp = false # "disable_userid_syslog_listener_udp": false
      disable_http_ocsp                  = false # "disable_http_ocsp": false
    }

    # Corresponds to: "permitted_ip": [ { ... } ]
    # This is a list of allowed source IPs/networks for management access.
    permitted_ip = [
      {
        name        = "10.10.10.10" # "name": "10.10.10.10"
        description = "string"      # "description": "string"
      }
    ]
  }
}