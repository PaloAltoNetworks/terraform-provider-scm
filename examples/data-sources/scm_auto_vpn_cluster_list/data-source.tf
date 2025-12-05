# Fetch a list of all clusters from the "All" folder.
data "scm_auto_vpn_cluster_list" "all_clusters" {
  folder = "All"
  limit  = 10
}

# Output the raw list of all certificate profile objects found.
output "auto_vpn_clusters_list" {
  description = "A list of all auto vpn cluster objects from the All folder."
  # This directly outputs the list of data from the data source
  value = data.scm_auto_vpn_cluster_list.all_clusters
}