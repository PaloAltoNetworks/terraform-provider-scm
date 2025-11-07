#
# Creates various resources used for subsequent examples
#

resource "scm_variable" "scm_next_hop" {
  folder      = "All"
  name        = "$scm_next_hop"
  description = "Managed by Terraform"
  type        = "ip-netmask"
  value       = "198.18.1.1"
}

resource "scm_variable" "scm_next_hop_fqdn" {
  folder      = "All"
  name        = "$scm_next_hop_fqdn"
  description = "Managed by Terraform"
  type        = "fqdn"
  value       = "nexthop.example.com"
}

resource "scm_ethernet_interface" "scm_ethernet_interface" {
  name    = "$scm_ethernet_interface"
  comment = "Managed by Terraform"
  folder  = "ngfw-shared"
  layer3 = {
    ip = [
      {
        name = "198.18.11.1/24"
      }
    ]
  }
}

resource "scm_ethernet_interface" "scm_bgp_interface" {
  name    = "$scm_bgp_interface"
  comment = "Managed by Terraform"
  folder  = "ngfw-shared"
  layer3 = {
    ip = [
      {
        name = "198.18.12.1/24"
      }
    ]
  }
}

resource "scm_bgp_auth_profile" "bgp_auth_profile" {
  folder = "ngfw-shared"
  name   = "bgp_auth_profile"
  secret = "Example123"
}

#
# Creates a logical router with static routes
#

resource "scm_logical_router" "scm_logical_router" {
  folder        = "ngfw-shared"
  name          = "scm_logical_router"
  routing_stack = "advanced"
  vrf = [
    {
      name = "default"
      interface = [
        "$scm_ethernet_interface",
      ]
      routing_table = {
        ip = {
          static_route = [
            {
              name        = "default-route"
              destination = "0.0.0.0/0"
              preference  = 10
              nexthop = {
                ip_address = "198.18.1.1"
              }
            },
            {
              name        = "internal-route"
              interface   = "$scm_ethernet_interface"
              destination = "192.168.1.0/24"
              preference  = 1
              nexthop = {
                ip_address = "$scm_next_hop"
              }
            },
            {
              name        = "route-with-fqdn-nh"
              interface   = "$scm_ethernet_interface"
              destination = "192.168.2.0/24"
              preference  = 1
              nexthop = {
                fqdn = "$scm_next_hop"
              }
              bfd = {
                profile = "default"
              }
            }
          ]
        }
      }
    }
  ]
  depends_on = [
    scm_variable.scm_next_hop,
    scm_variable.scm_next_hop_fqdn,
    scm_ethernet_interface.scm_ethernet_interface
  ]
}

#
# Creates a logical router with bgp configuration
#

resource "scm_logical_router" "scm_bgp_router" {
  folder        = "ngfw-shared"
  name          = "scm_bgp_router"
  routing_stack = "advanced"
  vrf = [
    {
      name = "default"
      interface = [
        "$scm_bgp_interface",
      ]
      bgp = {
        enable               = true
        router_id            = "198.18.1.254"
        local_as             = "65535"
        install_route        = true
        reject_default_route = false
        peer_group = [
          {
            name = "prisma-access"
            address_family = {
              ipv4 = "default"
            }
            connection_options = {
              authentication = "bgp_auth_profile"
            }
            peer = [
              {
                name    = "primary-access-primary"
                enable  = true
                peer_as = 65515
                peer_address = {
                  ip = "198.18.1.100"
                }
                local_address = {
                  interface = "$scm_bgp_interface"
                }
                connection_options = {
                  multihop = "3"
                }
              }
            ]
          }
        ]
      }
    }
  ]
  depends_on = [
    scm_ethernet_interface.scm_bgp_interface,
    scm_bgp_auth_profile.bgp_auth_profile
  ]
}