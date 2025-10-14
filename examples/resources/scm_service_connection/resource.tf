variable "folder_scope" {
  description = "The folder scope for the SCM resource (e.g., 'Shared', 'Predefined', or a specific folder name)."
  type        = string
  default     = "Service Connections"
}


## 1. Define the IKE Crypto Profile (IKE Phase 1)
# Note: The resource name is plural: "scm_ike_crypto_profile"
resource "scm_ike_crypto_profile" "example" {
  name       = "example-ike-crypto"
  folder     = var.folder_scope
  hash       = ["sha256"]
  dh_group   = ["group14"]
  encryption = ["aes-256-cbc"]
}

## 2. Define the IPsec Crypto Profile (IKE Phase 2)
# Note: The resource name is plural and nested blocks now use an equals sign (=).
resource "scm_ipsec_crypto_profile" "example" {
  name   = "panw-IPSec-Crypto"
  folder = var.folder_scope

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
  folder = var.folder_scope

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
  folder                   = var.folder_scope
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

resource "scm_service_connection" "site_a_vpn_sc" {
  name         = "creating_a_service_connection"
  region       = "us-west-1"
  ipsec_tunnel = scm_ipsec_tunnel.example.name
  subnets = [
    "10.1.0.0/16",
    "172.16.0.0/24"
  ]
  source_nat = true
}
