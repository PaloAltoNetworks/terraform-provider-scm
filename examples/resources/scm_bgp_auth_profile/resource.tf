resource "scm_bgp_auth_profile" "scm_bgp_auth_profile" {
  folder = "ngfw-shared"
  name   = "scm_bgp_auth_profile"
  secret = "ExampleSecret123"
}
