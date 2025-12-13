resource "scm_service_setting" "service_settings" {
  folder = "All"

  services = {
    dns_setting = {
      servers = {
        primary   = "10.10.10.10"
        secondary = "10.10.10.11"
      }
    }

    fqdn_refresh_time        = 15
    fqdn_stale_entry_timeout = 1440

    ntp_servers = {
      primary_ntp_server = {
        ntp_server_address = "10.10.10.10"
        authentication_type = {
          # Empty block for 'autokey' authentication type
          autokey = {}
        }
      }

      secondary_ntp_server = {
        ntp_server_address = "11.11.11.11"
        authentication_type = {
          # Empty block for 'none' authentication type
          none = {}
        }
      }
    }

    update_server         = "updates.paloaltonetworks.com"
    server_verification   = true
    secure_proxy_server   = "test_proxy_server"
    secure_proxy_port     = 90
    secure_proxy_user     = "test_proxy_user"
    secure_proxy_password = "test_password"
    lcaas_use_proxy       = true
    inline_cloud_proxy    = false
  }
}