# List all discovered applications for the ZTNA tenant.
data "ztna_discovered_application_list" "example" {}

output "discovered_applications" {
  value = data.ztna_discovered_application_list.example.applications
}
