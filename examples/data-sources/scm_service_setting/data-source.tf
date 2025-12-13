data "scm_service_setting" "settings_data" {
  id = "d64ba07a-c8fa-4143-992e-0dcae477a263"

}

output "fetched_settings" {
  description = "Fetched Service Settings"
  value = {
    # General Setting
    folder = data.scm_service_setting.settings_data.folder
    id     = data.scm_service_setting.settings_data.id

    # DNS Settings
    dns_primary   = data.scm_service_setting.settings_data.services.dns_setting.servers.primary
    dns_secondary = data.scm_service_setting.settings_data.services.dns_setting.servers.secondary

    # FQDN Settings
    fqdn_refresh_time        = data.scm_service_setting.settings_data.services.fqdn_refresh_time
    fqdn_stale_entry_timeout = data.scm_service_setting.settings_data.services.fqdn_stale_entry_timeout

    # # NTP Settings (Only Address, not Auth Type for brevity)
    ntp_primary_address   = data.scm_service_setting.settings_data.services.ntp_servers.primary_ntp_server.ntp_server_address
    ntp_secondary_address = data.scm_service_setting.settings_data.services.ntp_servers.secondary_ntp_server.ntp_server_address

    # Update Server Settings
    update_server       = data.scm_service_setting.settings_data.services.update_server
    server_verification = data.scm_service_setting.settings_data.services.server_verification

    # Proxy Settings (EXCLUDING PASSWORD)
    secure_proxy_server = data.scm_service_setting.settings_data.services.secure_proxy_server
    secure_proxy_port   = data.scm_service_setting.settings_data.services.secure_proxy_port
    secure_proxy_user   = data.scm_service_setting.settings_data.services.secure_proxy_user
    lcaas_use_proxy     = data.scm_service_setting.settings_data.services.lcaas_use_proxy
    inline_cloud_proxy  = data.scm_service_setting.settings_data.services.inline_cloud_proxy
  }
}