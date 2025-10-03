
## 1. Define the IKE Crypto Profile (IKE Phase 1)
# Note: The resource name is plural: "scm_ike_crypto_profile"
resource "scm_ike_crypto_profile" "example" {
  name       = "example-ike-crypto"
  folder     = "Remote Networks"
  hash       = ["sha256"]
  dh_group   = ["group14"]
  encryption = ["aes-256-cbc"]
}

## 2. Define the IPsec Crypto Profile (IKE Phase 2)
# Note: The resource name is plural and nested blocks now use an equals sign (=).
resource "scm_ipsec_crypto_profile" "example" {
  name   = "PaloAlto-Networks-IPSec-Crypto"
  folder = "Remote Networks"

  esp = {
    encryption     = ["aes-256-gcm"]
    authentication = ["sha256"]
  }

  dh_group = "group14"

  lifetime = {
    hours = 8
  }
}

## 3. Define the IKE Gateway
# Note: The resource name is plural and nested blocks now use an equals sign (=).
resource "scm_ike_gateway" "example" {
  name   = "example-gateway"
  folder = "Remote Networks"

  peer_address = {
    ip = "1.1.1.1"
  }

  authentication = {
    pre_shared_key = {
      # In a real-world scenario, use a variable marked as sensitive for the key.
      key = "secret"
    }
  }

  protocol = {
    ikev1 = {
      ike_crypto_profile = scm_ike_crypto_profile.example.name
    }
  }
}

## 4. Define the IPsec Tunnel
# Note: Nested 'auto_key' block uses an equals sign (=).
resource "scm_ipsec_tunnel" "example" {
  name                     = "example-tunnel"
  folder                   = "Remote Networks"
  tunnel_interface         = "tunnel"
  anti_replay              = true
  copy_tos                 = false
  enable_gre_encapsulation = false

  auto_key = {
    ike_gateway = [{
      name = scm_ike_gateway.example.name
    }]

    ipsec_crypto_profile = scm_ipsec_crypto_profile.example.name
  }

  depends_on = [
    scm_ike_gateway.example
  ]
}