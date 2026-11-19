# This file is embedded using go:embed
# Link a GCP project to the TSG in the provider scope. tsg_id and csp_id are
# read-only: the link service fills them in from the provider credentials.
resource "scm_cngfw_gcp_link" "project" {
  gcp_tenant_id = "projects/my-gcp-project"
}

# Link a GCP organization, storing the link in the europe region.
resource "scm_cngfw_gcp_link" "organization" {
  gcp_tenant_id = "organizations/123456789012"
  region        = "europe"
}

# Hand the link code to whoever creates the firewall endpoint in Google Cloud.
output "gcp_link_code" {
  value     = scm_cngfw_gcp_link.project.link_code
  sensitive = true
}

output "gcp_link_code_expires_at" {
  value = scm_cngfw_gcp_link.project.expires_at
}
