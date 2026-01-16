# scep profile w/ no challenge
resource "scm_scep_profile" "scm_scep_profile_1" {
  folder           = "All"
  name             = "scep-prof-1"
  scep_url         = "https://scep.example.com/"
  ca_identity_name = "Default"
  digest           = "sha1"
  subject          = "CN=$USERNAME"
  scep_challenge = {
    # none = {}
    fixed = "123"
  }

  algorithm = {
    rsa = {
      rsa_nbits = "1024"
    }
  }
}

# scep profile w/ fixed challenge
resource "scm_scep_profile" "scm_scep_profile_2" {
  folder           = "All"
  name             = "scep-prof-2"
  scep_url         = "https://example.target.com/"
  ca_identity_name = "user-scep"
  digest           = "sha256"
  subject          = "CN=$USERNAME"
  scep_ca_cert     = "Forward-Trust-CA-ECDSA"

  scep_challenge = {
    fixed = "Password123!"
  }

  certificate_attributes = {
    rfc822name = "user@example.com"
  }

  algorithm = {
    rsa = {
      rsa_nbits = "2048"
    }
  }
}

# scep profile w/ dynamic challenge
resource "scm_scep_profile" "scm_scep_profile_3" {
  folder                   = "All"
  name                     = "scep-prof-3"
  scep_url                 = "https://example.gateway.com/"
  ca_identity_name         = "vpn-gateway"
  digest                   = "sha384"
  subject                  = "CN=$CORP_VPN"
  scep_ca_cert             = "GlobalSign-Root-CA"
  fingerprint              = "64EC88CA00B268E5BA1A35678A1B5316D212F4F366B24772322342423123455A"
  use_as_digital_signature = false

  scep_challenge = {
    dynamic = {
      username       = "user123"
      password       = "Password123!"
      otp_server_url = "http://auth.com"
    }
  }

  certificate_attributes = {
    dnsname = "vpn-gateway.example.com"
  }

  algorithm = {
    rsa = {
      rsa_nbits = "3072"
    }
  }
}

# scep profile w/ all fields
resource "scm_scep_profile" "scm_scep_profile_4" {
  folder                   = "All"
  name                     = "scep-prof-4"
  scep_url                 = "https://example.wifi.com/"
  ca_identity_name         = "wifi"
  digest                   = "sha512"
  subject                  = "CN=$WIFI-ACCESS"
  scep_ca_cert             = "Root CA"
  scep_client_cert         = "Forward-UnTrust-CA-ECDSA"
  fingerprint              = "4448CA00B268E5BU690378A1B5316D212F4F366B2477232234394I"
  use_as_digital_signature = true
  use_for_key_encipherment = true

  scep_challenge = {
    dynamic = {
      username       = "admin"
      password       = "Pwd@123"
      otp_server_url = "http://auth.com"
    }
  }

  certificate_attributes = {
    uniform_resource_identifier = "file:///C:/Users/Documents/report.txt"
  }

  algorithm = {
    rsa = {
      rsa_nbits = "3072"
    }
  }
}