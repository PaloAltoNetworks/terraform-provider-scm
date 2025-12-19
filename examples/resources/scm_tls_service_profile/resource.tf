resource "scm_tls_service_profile" "tls_service_prof_1_upper" {
  folder      = "All"
  name        = "TLS_Service_Profile_1"
  certificate = "Authentication Cookie CA"

  protocol_settings = {
    keyxchg_algo_rsa = true
  }
}

resource "scm_tls_service_profile" "tls_service_prof_2_upper" {
  folder      = "All"
  name        = "TLS_Service_Profile_2"
  certificate = "Forward-Trust-CA"

  protocol_settings = {
    min_version = "tls1-0"
    max_version = "tls1-1"

    # Encryption
    enc_algo_aes_128_cbc = true
    enc_algo_aes_256_cbc = true
  }
}

resource "scm_tls_service_profile" "tls_service_prof_3_upper" {
  folder      = "All"
  name        = "TLS_Service_Profile_3"
  certificate = "Root CA"

  protocol_settings = {
    min_version = "tls1-1"
    max_version = "tls1-3"

    # Key Exchange
    keyxchg_algo_rsa   = true
    keyxchg_algo_dhe   = true
    keyxchg_algo_ecdhe = true

    # Encryption
    enc_algo_aes_128_cbc = true
    enc_algo_aes_128_gcm = true
    enc_algo_aes_256_cbc = true
    enc_algo_aes_256_gcm = true

    # Authentication
    auth_algo_sha1   = true
    auth_algo_sha256 = true
    auth_algo_sha384 = true
  }
}
