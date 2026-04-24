# User location representing the US East office IP ranges
resource "scm_forwarding_profile_user_location" "example_us_east" {
  folder      = "Mobile Users"
  name        = "us-east-tf"
  description = "Managed by Terraform - US East office user location"
  ip_addresses = [
    "10.1.0.0/16",
    "10.2.0.0/16"
  ]
}

# Corporate destination profile with FQDN and IP ranges
resource "scm_forwarding_profile_destination" "corp_destinations" {
  folder      = "Mobile Users"
  name        = "corp-destinations-combined"
  description = "Managed by Terraform - Corporate FQDN and IP destinations (combined example)"
  fqdn = [
    {
      name = "*.corporate.com"
      port = 443
    },
    {
      name = "api.internal.com"
      port = 8443
    }
  ]
  ip_addresses = [
    {
      name = "10.0.0.0/8"
      port = 443
    },
    {
      name = "172.16.0.0/12"
      port = 443
    }
  ]
}

# Basic Global Protect proxy-based forwarding profile with a single forwarding rule
resource "scm_forwarding_profile" "example_gp_proxy_basic" {
  folder = "Mobile Users"

  name = "example-gp-proxy-basic"

  type = {
    global_protect_proxy = {
      forwarding_rules = [
        {
          name         = "direct-rule"
          connectivity = "direct"
          destinations = "Any"
          enabled      = true
        }
      ]
    }
  }

  # Ensure user locations and destinations are created first
  depends_on = [
    scm_forwarding_profile_user_location.example_us_east,
    scm_forwarding_profile_destination.corp_destinations
  ]
}

# Global Protect proxy-based forwarding profile with block rule and multiple forwarding rules
resource "scm_forwarding_profile" "example_gp_proxy_full" {
  folder = "Mobile Users"

  name        = "example-gp-proxy-full"
  description = "Managed by Terraform - GP proxy with block rule and multiple forwarding rules"

  type = {
    global_protect_proxy = {
      pac_upload = false

      block_rule = {
        enable = true
        allow_tcp = {
          enable_locations = true
          locations        = [scm_forwarding_profile_user_location.example_us_east.name]
        }
        allow_udp = {
          enable_locations    = true
          locations           = [scm_forwarding_profile_user_location.example_us_east.name]
          enable_destinations = true
          destinations        = "any"
        }
      }

      forwarding_rules = [
        {
          name           = "direct-rule"
          connectivity   = "direct"
          destinations   = scm_forwarding_profile_destination.corp_destinations.name
          user_locations = "Any"
          enabled        = true
        }
      ]
    }
  }

  # Dependencies are implicit via references, but being explicit helps
  depends_on = [
    scm_forwarding_profile_user_location.example_us_east,
    scm_forwarding_profile_destination.corp_destinations
  ]
}

# PAC file based forwarding profile
resource "scm_forwarding_profile" "example_pac_file" {
  folder = "Mobile Users"

  name        = "example-pac-file"
  description = "Managed by Terraform - PAC file based forwarding profile"

  type = {
    pac_file = {
      pac_upload = true

      forwarding_rules = [
        {
          name         = "pac-direct-rule"
          connectivity = "direct"
          destinations = "Any"
          enabled      = true
        }
      ]
    }
  }

  # Ensure user locations and destinations are created first
  depends_on = [
    scm_forwarding_profile_user_location.example_us_east,
    scm_forwarding_profile_destination.corp_destinations
  ]
}

# ZTNA agent-based forwarding profile with block rule and forwarding rules
resource "scm_forwarding_profile" "example_ztna_agent" {
  folder = "Mobile Users"

  name        = "example-ztna-agent"
  description = "Managed by Terraform - ZTNA agent forwarding profile with block rule"

  type = {
    ztna_agent = {
      block_rule = {
        block_all_other_unmatched_outbound_connections             = false
        block_inbound_access_when_connected_to_tunnel              = false
        block_non_tcp_non_udp_traffic_when_connected_to_tunnel     = false
        allow_icmp_for_troubleshooting                             = false
        block_outbound_lan_access_when_connected_to_tunnel         = false
        enforcer_fqdn_dns_resolution_via_dns_servers               = true
        resolve_all_fqdns_using_dns_servers_assigned_by_the_tunnel = true
      }

      forwarding_rules = [
        {
          name                = "ztna-dns-rule"
          connectivity        = "tunnel"
          destinations        = scm_forwarding_profile_destination.corp_destinations.name
          source_applications = "Any"
          traffic_type        = "dns-and-network-traffic"
          user_locations      = "Any"
          enabled             = true
        },
        {
          name                = "ztna-direct-rule"
          connectivity        = "direct"
          destinations        = "Any"
          source_applications = "Any"
          traffic_type        = "network-traffic"
          user_locations      = "Any"
          enabled             = true
        }
      ]
    }
  }

  # Dependencies are implicit via reference, but being explicit helps
  depends_on = [
    scm_forwarding_profile_destination.corp_destinations
  ]
}
