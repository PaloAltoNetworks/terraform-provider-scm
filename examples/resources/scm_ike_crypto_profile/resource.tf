# The resource block defines a new IKE Crypto Profile.
resource "scm_ike_crypto_profile" "scm_ike_crypto_profile_2" {
  name   = "scm_ike_crypto_profile_2"
  folder = "Prisma Access"

  # A list of hashing algorithms.
  hash = ["sha256", "sha384"]

  # A list of Diffie-Hellman (DH) groups.
  dh_group = ["group14", "group5"]

  # A list of encryption algorithms.
  encryption = ["aes-256-cbc", "aes-128-cbc"]

  # The lifetime specifies how long the IKE security association (SA) is valid.
  # Only one of days, hours, minutes, or seconds can be specified.
  lifetime = {
    hours = 8
  }

  # IKEv2 SA reauthentication interval. 0 means disabled.
  authentication_multiple = 10
}
