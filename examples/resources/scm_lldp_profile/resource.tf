resource "scm_lldp_profile" "example" {
  folder = "All"
  mode   = "transmit-receive"
  name   = "lldp-profile-tf-1"
  option_tlvs = {
    management_address = {
      enabled = true
    }
    port_description    = true
    system_capabilities = true
    system_description  = false
    system_name         = true
  }
  snmp_syslog_notification = true
}