data "scm_site" "single_site_by_id" {
  # Requires the unique UUID of the rule
  id = "d037fe30-68ae-47ee-9a74-71bc63ac2c10"
}

output "single_site_details" {
  description = "Single Site fetched by ID."
  value       = data.scm_site.single_site_by_id
}
