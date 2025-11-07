resource "scm_radius_server_profile" "chap_radius_profile" {
  # Standard fields
  name    = "CHAP_only_rsp_ds_1"
  folder  = "All"
  retries = 5
  timeout = 60

  protocol = {
    c_h_a_p = {}
  }

  # Server list
  server = [
    {
      name       = "Chap_Server_Primary"
      ip_address = "10.1.1.10"
      port       = 1812
      secret     = "-AQ==lhyuV6U/j9Trb9JL9L0UoBecg9Y=kTOWntGhZ1KFyLD+etKQ3g=="
    },
  ]
}

data "scm_radius_server_profile" "single_profile_by_id" {
  # Requires the unique UUID of the rule
  id = scm_radius_server_profile.chap_radius_profile.id
}

output "single_rsp__dump" {
  description = "Dump of all attributes of the single profile fetched by ID."
  # Dumping the entire fetched object, similar to your request for App Override
  value = data.scm_radius_server_profile.single_profile_by_id.name
}