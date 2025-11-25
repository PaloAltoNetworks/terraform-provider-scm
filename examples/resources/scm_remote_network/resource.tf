# --- DEPENDENCY 1: IKE Crypto Profile ---
# This profile defines the encryption and authentication algorithms for the IKE Gateway.
# The values are taken from the 'createTestIKECryptoProfile' helper function.
resource "scm_ike_crypto_profile" "example" {
  name       = "example-ike-crypto-11"
  folder     = "Remote Networks"
  hash       = ["sha256"]
  dh_group   = ["group14"]
  encryption = ["aes-256-cbc"]
}

# --- DEPENDENCY 2: IKE Gateway ---
# This defines the VPN peer. It depends on the IKE Crypto Profile created above.
# The values are taken from the 'createTestIKEGateway' helper function.
resource "scm_ike_gateway" "example" {
  name   = "example-ike-gateway-11"
  folder = "Remote Networks"
  authentication = {
    pre_shared_key = {
      key = "secret" # Please use a secure, generated key in production.
    }
  }
  peer_address = {
    ip = "1.1.1.1"
  }
  protocol = {
    ikev1 = {
      ike_crypto_profile = scm_ike_crypto_profile.example.name
    }
  }

  depends_on = [
    scm_ike_crypto_profile.example
  ]
}

# --- DEPENDENCY 3: IPsec Tunnel ---
# This defines the tunnel interface itself and uses the IKE Gateway.
# The values are taken from the 'createTestIPsecTunnel' helper function.
resource "scm_ipsec_tunnel" "example" {
  name                     = "example-ipsec-tunnel-11"
  folder                   = "Remote Networks"
  anti_replay              = true
  copy_tos                 = false
  enable_gre_encapsulation = false
  auto_key = {
    ike_gateway = [
      {
        name = scm_ike_gateway.example.name
      }
    ]
    ipsec_crypto_profile = "PaloAlto-Networks-IPSec-Crypto" # This is a default, built-in profile.
  }

  depends_on = [
    scm_ike_gateway.example
  ]
}

# --- MAIN RESOURCE: Remote Network ---
# This is the final resource, which uses the IPsec Tunnel created above.
# The values are taken directly from the 'Test_deployment_services_RemoteNetworksAPIService_Create' test.
resource "scm_remote_network" "example" {
  name         = "example-remote-network-11"
  folder       = "Remote Networks"
  license_type = "FWAAS-AGGREGATE"
  region       = "us-west-2"
  spn_name     = "us-west-dakota"
  subnets      = ["192.168.1.0/24"]
  ipsec_tunnel = scm_ipsec_tunnel.example.name

  protocol = {
    bgp = {
      enable                       = true
      peer_as                      = "65000"
      local_ip_address             = "169.254.1.1"
      peer_ip_address              = "169.254.1.2"
      do_not_export_routes         = false
      originate_default_route      = false
      summarize_mobile_user_routes = false
    }
  }

  depends_on = [
    scm_ipsec_tunnel.example
  ]
}