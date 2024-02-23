resource "scm_local_user" "example" {
  folder   = "Shared"
  name     = "user1"
  password = "secret"
}
