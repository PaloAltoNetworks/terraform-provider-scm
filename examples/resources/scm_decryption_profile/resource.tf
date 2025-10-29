resource "scm_decryption_profile" "scm_decryption_profile_base" {
  folder = "ngfw-shared"
  name   = "dp_base"
}

resource "scm_decryption_profile" "scm_decryption_profile_forward_proxy" {
  folder = "ngfw-shared"
  name   = "dp_forward_proxy"
  ssl_forward_proxy = {
    auto_include_altname              = false
    block_client_cert                 = true
    block_expired_certificate         = true
    block_timeout_cert                = true
    block_tls13_downgrade_no_resource = false
    block_unknown_cert                = false
    block_unsupported_cipher          = true
    block_unsupported_version         = true
    block_untrusted_issuer            = true
    restrict_cert_exts                = false
    strip_alpn                        = true
  }
}

resource "scm_decryption_profile" "scm_decryption_profile_inbound_proxy" {
  folder = "ngfw-shared"
  name   = "dp_inbound_proxy"
  ssl_inbound_proxy = {
    block_if_hsm_unavailable  = true
    block_if_no_resource      = true
    block_unsupported_cipher  = false
    block_unsupported_version = true
  }
}

resource "scm_decryption_profile" "scm_decryption_profile_no_proxy" {
  folder = "ngfw-shared"
  name   = "dp_no_proxy"
  ssl_no_proxy = {
    block_expired_certificate = true
    block_untrusted_issuer    = false
  }
}

resource "scm_decryption_profile" "scm_decryption_profile_protocol_settings" {
  folder = "ngfw-shared"
  name   = "dp_protocol_settings"
  ssl_protocol_settings = {
    auth_algo_md5              = true
    auth_algo_sha1             = true
    auth_algo_sha256           = true
    auth_algo_sha384           = false
    enc_algo_3des              = false
    enc_algo_aes_128_cbc       = false
    enc_algo_aes_128_gcm       = true
    enc_algo_aes_256_cbc       = false
    enc_algo_aes_256_gcm       = true
    enc_algo_chacha20_poly1305 = false
    enc_algo_rc4               = false
    keyxchg_algo_dhe           = true
    keyxchg_algo_ecdhe         = true
    keyxchg_algo_rsa           = false
    max_version                = "max"
    min_version                = "tls1-2"
  }
}

resource "scm_decryption_profile" "mixed_decryption_profile" {
  folder = "ngfw-shared"
  name   = "mixed_dp"

  ssl_forward_proxy = {
    auto_include_altname      = true
    block_client_cert         = true
    block_expired_certificate = false
    restrict_cert_exts        = false
    strip_alpn                = true
  }

  ssl_inbound_proxy = {
    block_if_hsm_unavailable  = true
    block_if_no_resource      = true
    block_unsupported_cipher  = true
    block_unsupported_version = true
  }

  ssl_protocol_settings = {
    auth_algo_md5      = true
    auth_algo_sha1     = true
    auth_algo_sha256   = false
    auth_algo_sha384   = true
    enc_algo_3des      = true
    enc_algo_rc4       = true
    keyxchg_algo_dhe   = false
    keyxchg_algo_ecdhe = false
    max_version        = "tls1-3"
    min_version        = "tls1-1"
  }
}

resource "scm_decryption_profile" "full_mixed_decryption_profile" {
  folder = "ngfw-shared"
  name   = "full_mixed_dp"

  ssl_forward_proxy = {
    auto_include_altname      = true
    block_client_cert         = true
    block_expired_certificate = false
    block_timeout_cert        = true
    block_unknown_cert        = false
    block_unsupported_cipher  = true
    block_untrusted_issuer    = false
    restrict_cert_exts        = false
    strip_alpn                = true
  }

  ssl_inbound_proxy = {
    block_if_hsm_unavailable  = true
    block_if_no_resource      = false
    block_unsupported_cipher  = true
    block_unsupported_version = false
  }

  ssl_no_proxy = {
    block_expired_certificate = false
    block_untrusted_issuer    = true
  }

  ssl_protocol_settings = {
    auth_algo_md5        = false
    auth_algo_sha1       = true
    auth_algo_sha256     = false
    auth_algo_sha384     = true
    enc_algo_3des        = false
    enc_algo_aes_128_gcm = true
    enc_algo_aes_256_cbc = false
    enc_algo_aes_256_gcm = true
    enc_algo_rc4         = true
    keyxchg_algo_dhe     = false
    keyxchg_algo_ecdhe   = true
    keyxchg_algo_rsa     = false
    max_version          = "tls1-0"
    min_version          = "sslv3"
  }
}