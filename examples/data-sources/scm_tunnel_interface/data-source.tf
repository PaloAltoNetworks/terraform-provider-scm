# Look up tunnel interface by its ID.
data "scm_tunnel_interface" "scm_tunnel_intf_ds" {
  id = "ddad1e64-0b64-41a4-b361-c6191169a8f1"
}

# Output various attributes from the found tunnel interface to verify the lookups were successful.
output "scm_tunnel_interface_data_source_results" {
  description = "The details of the tunnel interface read from the data source"
  value = {
    id      = data.scm_tunnel_interface.scm_tunnel_intf_ds.id
    name    = data.scm_tunnel_interface.scm_tunnel_intf_ds.name
    comment = data.scm_tunnel_interface.scm_tunnel_intf_ds.comment
    ip      = data.scm_tunnel_interface.scm_tunnel_intf_ds.ip
    folder  = data.scm_tunnel_interface.scm_tunnel_intf_ds.folder
  }
}
