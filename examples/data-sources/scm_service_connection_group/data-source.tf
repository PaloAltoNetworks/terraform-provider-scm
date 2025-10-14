variable "folder_scope" {
  description = "The folder scope for the SCM resource (e.g., 'Shared', 'Predefined', or a specific folder name)."
  type        = string
  default     = "Service Connections"
}

## 1. IKE Crypto Profile (IKE Phase 1)
resource "scm_ike_crypto_profile" "example" {
  name       = "example-ike-crypto_sc_grp"
  folder     = var.folder_scope
  hash       = ["sha256"]
  dh_group   = ["group14"]
  encryption = ["aes-256-cbc"]
}

## 2. IPsec Crypto Profile (IKE Phase 2)
resource "scm_ipsec_crypto_profile" "example" {
  name   = "panw-IPSec-Crypto_sc_grp"
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

## 3. IKE Gateway
resource "scm_ike_gateway" "example" {
  name   = "example-gateway_sc_grp"
  folder = var.folder_scope

  peer_address = {
    ip = "1.1.1.1"
  }

  authentication = {
    pre_shared_key = {
      key = "secret"
    }
  }

  protocol = {
    ikev1 = {
      ike_crypto_profile = scm_ike_crypto_profile.example.name
    }
  }
}

## 4. IPsec Tunnel
resource "scm_ipsec_tunnel" "example" {
  name                     = "example-tunnel_sc_grp"
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

## 5. Service Connection (The target for the group)
resource "scm_service_connection" "site_a_vpn_sc" {
  name         = "creating_a_service_connection_sc_grp"
  region       = "us-west-1a"
  ipsec_tunnel = scm_ipsec_tunnel.example.name
  subnets = [
    "10.1.0.0/16",
    "172.16.0.0/24"
  ]
  source_nat = false
}

## 5. Service Connection (The target for the group)
resource "scm_service_connection" "site_a_vpn_sc_2" {
  name         = "creating_a_service_connection_sc_grp_2"
  region       = "us-west-1a"
  ipsec_tunnel = scm_ipsec_tunnel.example.name
  subnets = [
    "10.1.0.0/16",
    "172.16.0.0/24"
  ]
  source_nat = false
}

## 6. Service Connection Group (Groups the Service Connection created above)
resource "scm_service_connection_group" "example_group" {
  name = "service-connection-group-app_sc_grp"
  # The key attribute: references the name of the Service Connection resource
  target = [
    scm_service_connection.site_a_vpn_sc.name, scm_service_connection.site_a_vpn_sc_2.name
  ]

  disable_snat = false
  pbf_only     = true
}

# ------------------------------------------------------------------
# Data Source: SCM Service Connection Group (Single Lookup)
# ------------------------------------------------------------------

data "scm_service_connection_group" "group_lookup" {
  # Name is fetched from the resource creation block
  id = scm_service_connection_group.example_group.id
}

output "looked_up_service_connection_group_details" {
  description = "All attributes retrieved for the Service Connection Group via the data source lookup."
  # Directly reference the data source object
  value = data.scm_service_connection_group.group_lookup
}

