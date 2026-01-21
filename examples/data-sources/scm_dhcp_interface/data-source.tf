# We use the ID from the resource created above.
data "scm_dhcp_interface" "single_inf_by_id" {
  # Requires the unique UUID of the rule
  id = "b3544acb-fc55-4c6f-921d-4128b5a1d135"
}

output "single_dhcp_inf_name" {
  description = "The name of the single DHCP Interface fetched by ID."
  value       = data.scm_dhcp_interface.single_inf_by_id
}
