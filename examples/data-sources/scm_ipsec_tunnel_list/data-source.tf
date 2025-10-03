# Example of listing all IPsec Tunnels within a specific folder.
data "scm_ipsec_tunnel_list" "all_in_folder" {
  folder = "Remote Networks"
  limit  = 100
}

output "all_ipsec_tunnels_in_folder" {
  value = data.scm_ipsec_tunnel_list.all_in_folder.data
}
