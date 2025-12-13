resource "scm_ethernet_interface" "sr_ethernet_interface_1" {
  name        = "$layer3_sr_1"
  comment     = "Managed by Terraform"
  folder      = "ngfw-shared"
  link_speed  = "auto"
  link_duplex = "full"
  link_state  = "auto"
  layer3 = {
    ip = [
      {
        name = "10.10.10.10"
      }
    ]
  }
}

resource "scm_ethernet_interface" "sr_ethernet_interface_2" {
  name        = "$layer3_sr_2"
  comment     = "Managed by Terraform"
  folder      = "ngfw-shared"
  link_speed  = "auto"
  link_duplex = "full"
  link_state  = "auto"
  layer3 = {
    ip = [
      {
        name = "10.10.10.20"
      }
    ]
  }
}

resource "scm_service_route" "sr_settings" {
  folder = "ngfw-shared"
  route = {
    service = [
      {
        name = "ddns"
        source = {
          interface = scm_ethernet_interface.sr_ethernet_interface_1.name
          address   = "10.10.10.10"
        }
      }
    ]
    destination = [
      {
        name = scm_ethernet_interface.sr_ethernet_interface_2.name,
        source = {
          address = "10.10.10.20"
        }
      }
    ]
  }
}