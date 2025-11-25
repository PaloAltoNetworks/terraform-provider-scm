data "scm_site_list" "all_sites" {
  limit  = 50
  folder = "Remote Networks"
}

output "list_of_all_connection_names" {
  description = "A list of names for all Sites found in the RN Folder."
  value = [
    for conn in data.scm_site_list.all_sites.data : conn.name
  ]
}

output "total_connections_count" {
  description = "The total number of sites found by the list data source."
  value       = length(data.scm_site_list.all_sites.data)
}
