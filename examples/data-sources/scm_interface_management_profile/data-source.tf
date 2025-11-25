data "scm_interface_management_profile" "single_profile_by_id" {
  # The data source type needs to match the resource type, 
  # but the name is arbitrary (e.g., single_profile_by_id).
  id = "f4358615-daba-4b71-a0ea-bd3ebb412fe3"
}

# --------------------------------------------------------------------------------

output "fetched_profile_name" {
  description = "The name of the single Interface Management Profile fetched by ID."
  # Access the 'name' attribute from the data source
  value = data.scm_interface_management_profile.single_profile_by_id.name
}

output "fetched_profile" {
  description = "The PING status retrieved from the fetched data source."
  # Access the 'ping' attribute from the data source
  value = data.scm_interface_management_profile.single_profile_by_id
}