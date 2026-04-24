# List all devices
data "scm_device_list" "all_devices" {
  limit  = 200
  offset = 0
}

# Output the devices
output "all_device_names" {
  value = [for d in data.scm_device_list.all_devices.data : d.name]
}

output "device_details" {
  value = [for d in data.scm_device_list.all_devices.data : {
    id               = d.id
    name             = d.name
    model            = d.model
    family           = d.family
    is_connected     = d.is_connected
    software_version = d.software_version
    ip_address       = d.ip_address
  }]
}
