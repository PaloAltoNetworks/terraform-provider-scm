#
# Data source to retrieve a single URL Access Profile object by its ID.
#

# Replace the ID with the UUID of the URL Access Profile you want to find.
data "scm_url_access_profile" "example" {
  id = "e97c7e7e-9906-42d6-90a8-606ed5527125"
}

# Output the details of the single URL Access Profile object found.
output "url_access_profile_details" {
  description = "The details of the specific URL Access Profile object."
  value       = data.scm_url_access_profile.example
}