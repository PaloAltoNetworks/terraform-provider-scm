# The resource block defines a new IPsec Crypto Profile.
resource "scm_ipsec_crypto_profile" "scm_ipsec_crypto_profile_2" {
  name   = "scm_ipsec_crypto_profile_2"
  folder = "Prisma Access"

  # ESP (Encapsulating Security Payload) settings.
  esp = {
    authentication = ["sha256", "sha384"]
    encryption     = ["aes-256-gcm", "aes-128-cbc"]
  }

  # DH Group (Diffie-Hellman) for Perfect Forward Secrecy (PFS).
  dh_group = "group14"

  # Lifetime specifies how long the IPsec security association (SA) is valid.
  # Only one of days, hours, minutes, or seconds can be specified.
  lifetime = {
    hours = 1
  }

  # Lifesize specifies the amount of data that can be processed before the SA expires.
  # Only one of gb, kb, mb, or tb can be specified.
  lifesize = {
    gb = 10
  }
}