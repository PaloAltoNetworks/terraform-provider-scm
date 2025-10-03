#
# Data source to retrieve a single URL Category object by its ID.
#

# Replace the ID with the UUID of the URL Category you want to find.
data "scm_url_category" "example" {
  id = "5ae04e1a-bc7b-4ea3-99bb-86de23886b45"
}

# Output the details of the single URL Category object found.
output "url_category_details" {
  description = "The details of the specific URL Category object."
  value       = data.scm_url_category.example
}