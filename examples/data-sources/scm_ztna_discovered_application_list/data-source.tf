# List all discovered applications for the ZTNA tenant.
data "scm_ztna_discovered_application_list" "example" {}

output "discovered_applications" {
  value = data.scm_ztna_discovered_application_list.example.applications
}
