#
# Data source to retrieve a single SCM SAML Server Profile object by its ID.
#

# Replace the ID with the UUID of the SCM SAML Server Profile you want to find.
data "scm_saml_server_profile" "scm_saml_server_prof" {
  id = "a17abcfc-d37d-4b8a-bb09-102ffdc3abef"
}

# Output the details of the single SCM SAML Server Profile object found.
output "scm_saml_server_profile_details" {
  value = {
    folder  = data.scm_saml_server_profile.scm_saml_server_prof.folder
    name    = data.scm_saml_server_profile.scm_saml_server_prof.name
    id      = data.scm_saml_server_profile.scm_saml_server_prof.id
    sso_url = data.scm_saml_server_profile.scm_saml_server_prof.sso_url
  }
}