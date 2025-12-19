resource "scm_ethernet_interface" "layer3_setting" {
  name        = "$layer3_setting"
  comment     = "Managed by Terraform"
  folder      = "ngfw-shared"
  link_speed  = "auto"
  link_duplex = "full"
  link_state  = "auto"
  layer3 = {
    ip = [
      {
        name = "10.10.10.11"
      }
    ]
  }
}

resource "scm_device_redistribution_collector" "drc_settings" {
  folder = "ngfw-shared"
  redistribution_collector = {
    interface = scm_ethernet_interface.layer3_setting.name
  }
}