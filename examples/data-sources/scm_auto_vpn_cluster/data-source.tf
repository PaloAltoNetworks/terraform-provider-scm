# Look up the application group by its ID from the terraform.tfstate file.
data "scm_auto_vpn_cluster" "scm_auto_vpn_cluster_ds" {
  id = "e3c732fb-85c7-4116-b5ed-ba3977209b29"
}

# Output the members of the application group.
output "auto_vpn_cluster_outputs" {
  value = {
    id           = data.scm_auto_vpn_cluster.scm_auto_vpn_cluster_ds.id
    name         = data.scm_auto_vpn_cluster.scm_auto_vpn_cluster_ds.name
    branches     = data.scm_auto_vpn_cluster.scm_auto_vpn_cluster_ds.branches
    gateways     = data.scm_auto_vpn_cluster.scm_auto_vpn_cluster_ds.gateways
    enable_sdwan = data.scm_auto_vpn_cluster.scm_auto_vpn_cluster_ds.enable_sdwan
  }
}