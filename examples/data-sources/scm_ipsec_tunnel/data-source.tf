# Example of looking up an individual IPsec Tunnel by its ID.
data "scm_ipsec_tunnel" "scm_ipsec_tunnel_ds" {
  id = "7c237a82-8c11-4f09-bdbf-599e159019ce"
}

# Output to display the result of the data source lookup.
output "ipsec_tunnel_by_id" {
  value = data.scm_ipsec_tunnel.scm_ipsec_tunnel_ds
}