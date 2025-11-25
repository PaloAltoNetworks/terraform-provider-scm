#
# Auto VPN Cluster Configuration
#
resource "scm_auto_vpn_cluster" "test_cluster" {
  name         = "test-will-be-deleted"
  type         = "hub-spoke"
  enable_sdwan = true

  #
  # Gateway Configuration (Hub)
  #
  gateways = [
    {
      name                       = "008199000033876" # Serial Number / Device ID
      bgp_redistribution_profile = "All-Connected-Routes"
      priority                   = 1
      logical_router             = "untrust-logical-router"

      # Note: 'interfaces' is also a list of objects inside the gateway object
      interfaces = [
        {
          name = "$mesh-interface" # Replace with actual interface name or reference

          sdwan_link_settings = {
            sdwan_interface_profile = "test-sdwan-profile"
            sdwan_gateway           = "$mesh-subnet-gateway-address"

            upstream_nat = {
              enable = false
            }
          }
        }
      ]
    }
  ]

  #
  # Branch Configuration
  #
  branches = [
    {
      name                       = "008199000033879" # Serial Number / Device ID
      bgp_redistribution_profile = "All-Connected-Routes"
      logical_router             = "trust-logical-router"

      # Note: 'interfaces' is also a list of objects inside the branch object
      interfaces = [
        {
          name = "$trust-interface" # Replace with actual interface name or reference

          sdwan_link_settings = {
            sdwan_interface_profile = "test-sdwan-profile"

            upstream_nat = {
              enable = true

              static_ip = {
                ip_address = "$trust-subnet-gateway-address"
              }
            }
          }
        }
      ]
    }
  ]
}