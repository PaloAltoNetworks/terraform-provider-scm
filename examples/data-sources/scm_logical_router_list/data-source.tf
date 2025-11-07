# Fetch a list of all logical routers
data "scm_logical_router_list" "all_logical_routers" {
  folder = "ngfw-shared"
  limit  = 100
}

# Output the raw list of all logical routers
output "scm_logical_router_list" {
  description = "A list of all logical routers from the 'All Firewalls' folder."
  value       = { for router in data.scm_logical_router_list.all_logical_routers.data : router.name => router }
}