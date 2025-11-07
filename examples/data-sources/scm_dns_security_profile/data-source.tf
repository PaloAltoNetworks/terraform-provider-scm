#
# Data source to retrieve a single SCM DNS Security Profile object by its ID.
#

# Replace the ID with the UUID of the SCM DNS Profile you want to find.
data "scm_dns_security_profile" "scm_dns_profile" {
  id = "18607c90-22fa-4627-8741-f0584d1fa7d6"
}

# Output the details of the single SCM DNS Security Profile object found.
output "scm_dns_security_profile_details" {
  value = {
    profile_id = data.scm_dns_security_profile.scm_dns_profile.id
    folder     = data.scm_dns_security_profile.scm_dns_profile.folder
    name       = data.scm_dns_security_profile.scm_dns_profile.name
  }
}