# Look up a device by id
data "scm_device" "my_device" {
  id = "007099000023642"
}

# Output the details of the found device
output "scm_device_output" {
  value = {
    id               = data.scm_device.my_device.id
    name             = data.scm_device.my_device.name
    hostname         = data.scm_device.my_device.hostname
    model            = data.scm_device.my_device.model
    family           = data.scm_device.my_device.family
    folder           = data.scm_device.my_device.folder
    is_connected     = data.scm_device.my_device.is_connected
    software_version = data.scm_device.my_device.software_version
    ip_address       = data.scm_device.my_device.ip_address
    mac_address      = data.scm_device.my_device.mac_address
    uptime           = data.scm_device.my_device.uptime
  }
}
