#
# Creates a URL Access Profile object.
#
resource "scm_url_access_profile" "example" {
  folder      = "Shared"
  name        = "example_url_access_profile"
  description = "Test URL Access Profile for create API"
  block       = ["adult", "gambling"]
  alert       = ["high-risk", "phishing"]
}