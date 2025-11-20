resource "scm_authentication_portal" "example_configuration" {
  redirect_host = "192.168.255.254"
  folder        = "All"

  # Optional Fields
  authentication_profile = "test_auth_profile"
  certificate_profile    = "EDL-Hosting-Service-Profile"
  tls_service_profile    = "test_svc_profile"
  gp_udp_port            = 12
  idle_timer             = 12
  timer                  = 10
}