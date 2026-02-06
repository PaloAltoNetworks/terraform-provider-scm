resource "scm_certificate_profile" "scm_cp_1" {
  folder                     = "Prisma Access"
  name                       = "scm_cp_1"
  domain                     = "test"
  use_crl                    = true
  use_ocsp                   = true
  block_unknown_cert         = true
  block_timeout_cert         = true
  block_unauthenticated_cert = true
  block_expired_cert         = true
  crl_receive_timeout        = "5"
  ocsp_receive_timeout       = "5"
  cert_status_timeout        = "5"

  ca_certificates = [{
    name             = "Forward-Trust-CA"
    default_ocsp_url = "http://test.com"
    ocsp_verify_cert = "Forward-Trust-CA-ECDSA"
    template_name    = "something"
  }]

  username_field = {
    subject = "common-name"
  }
}