data "scm_authentication_portal" "portal_data" {
  id = "5fb1d1e9-cd39-4293-bd82-339ee91ce32f"

}

output "scm_authentication_portal" {
  description = "The UUID of the fetched authentication portal."
  value       = data.scm_authentication_portal.portal_data.id
}

output "fetched_portal" {
  description = "Fetched Auth Portal"
  value       = data.scm_authentication_portal.portal_data
}