resource "scm_saml_server_profile" "scm_saml_server_profile_1" {
  folder       = "All"
  name         = "scm-saml-server-prof-1"
  certificate  = "Global Authentication Cookie Cert"
  entity_id    = "123"
  sso_url      = "http://example.com"
  sso_bindings = "post"
}

resource "scm_saml_server_profile" "scm_saml_server_profile_2" {
  folder                   = "All"
  name                     = "scm-saml-server-prof-2"
  certificate              = "Global Authentication Cookie Cert"
  entity_id                = "test_id"
  max_clock_skew           = 100
  slo_bindings             = "redirect"
  sso_bindings             = "redirect"
  sso_url                  = "http://target.com"
  validate_idp_certificate = true
}

resource "scm_saml_server_profile" "scm_saml_server_profile_3" {
  folder                    = "All"
  name                      = "scm-saml-server-prof-3"
  certificate               = "Global Authentication Cookie Cert"
  entity_id                 = "test_123"
  max_clock_skew            = 900
  slo_bindings              = "post"
  sso_bindings              = "redirect"
  slo_url                   = "http://auth.com"
  sso_url                   = "http://okta.com"
  validate_idp_certificate  = false
  want_auth_requests_signed = false
}