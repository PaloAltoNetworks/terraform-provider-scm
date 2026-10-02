# Retrieve the list of Prisma Access locations. This is read-only reference
# data (no create/update), so the data source takes no filter arguments — it
# returns every location the tenant can use for Service Connection / Remote
# Network region parameters.
#

data "scm_location_list" "all" {}

# Map of region value -> human-readable display name.
output "prisma_access_locations" {
  description = "All Prisma Access locations, keyed by region value."
  value       = { for loc in data.scm_location_list.all.data : loc.value => loc.display }
}

# Example: reference a specific location's region value elsewhere (e.g. as a
# Remote Network region), looked up by display name.
output "austria_rn_region" {
  description = "Region value for the location displayed as 'Austria'."
  value = one([
    for loc in data.scm_location_list.all.data : loc.value if loc.display == "Austria"
  ])
}

# Example: reference a specific location's region value elsewhere (e.g. as a
# Remote Network or Service Connection region), looked up by display name.
output "austria_sc_region" {
  description = "Region value for the location displayed as 'Austria'."
  value = one([
    for loc in data.scm_location_list.all.data : loc.value if loc.display == "Austria"
  ])
}

# Example: reference a specific location's aggregate_region value elsewhere (e.g. as a
# Remote Networks Bandwidth Region), looked up by display name.
output "austria_rn_bandwidth_region" {
  description = "Region value for the location displayed as 'Austria'."
  value = one([
    for loc in data.scm_location_list.all.data : loc.aggregate_region if loc.display == "Austria"
  ])
}

