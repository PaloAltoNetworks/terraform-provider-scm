
data "scm_radius_server_profile" "single_profile_by_id" {
  # Requires the unique UUID of the rule
  id = "50e5f694-19a2-467b-90a8-9db168600327"
}

output "single_rsp__dump" {
  description = "Dump of all attributes of the single profile fetched by ID."
  # Dumping the entire fetched object, similar to your request for App Override
  value = data.scm_radius_server_profile.single_profile_by_id.name
}