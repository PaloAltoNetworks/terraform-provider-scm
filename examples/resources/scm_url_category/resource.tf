#
# Creates a custom URL Category object.
#
resource "scm_url_category" "example" {
  folder      = "Prisma Access"
  name        = "example_url_category"
  description = "Test URL Category for create API"
  list        = ["example.com", "test-create.com"]
  type        = "URL List"
}