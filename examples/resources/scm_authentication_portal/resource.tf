resource "scm_authentication_portal" "example_configuration" {
  redirect_host = "192.168.255.254"
  folder        = "ngfw-shared"


  # Optional Fields
  authentication_profile = "test_auth_profile"
  certificate_profile    = "EDL-Hosting-Service-Profile"
  gp_udp_port            = 12
  idle_timer             = 12
  timer                  = 10
}