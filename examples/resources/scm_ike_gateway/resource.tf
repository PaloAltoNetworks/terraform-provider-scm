// The scm_ike_crypto_profile resource is a prerequisite for the IKE gateway.
resource "scm_ike_crypto_profile" "scm_ike_gw_crypto_profile_1" {
  folder     = "Remote Networks"
  name       = "scm_ike_gw_crypto_profile_1"
  hash       = ["sha256"]
  dh_group   = ["group14"]
  encryption = ["aes-256-cbc"]
}

// This is the main scm_ike_gateway resource.
resource "scm_ike_gateway" "scm_ike_gateway_1" {
  folder = "Remote Networks"
  name   = "scm_ike_gateway_1"
  authentication = {
    pre_shared_key = {
      key = "123456"
    }
  }
  peer_address = {
    ip = "2.2.2.4"
  }
  peer_id = {
    type = "ipaddr"
    id   = "10.3.3.4"
  }
  local_id = {
    type = "ipaddr"
    id   = "10.3.4.4"
  }
  protocol = {
    ikev1 = {
      ike_crypto_profile = scm_ike_crypto_profile.scm_ike_gw_crypto_profile_1.name
      dpd = {
        enable = true
      }
    }
    ikev2 = {
      ike_crypto_profile = scm_ike_crypto_profile.scm_ike_gw_crypto_profile_1.name
      dpd = {
        enable = true
      }
    }
  }
}
