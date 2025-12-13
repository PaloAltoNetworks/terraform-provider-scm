// Set Sessions with Config Rematch value set to True
resource "scm_session_setting" "session_example" {
  # Corresponds to: "folder": "All"
  folder = "All"

  session_settings = {

    # Basic Session Settings
    dhcp_bcast_session_on              = false
    erspan                             = false
    ipv6_firewalling                   = true
    rematch                            = false
    accelerated_aging_enable           = true
    accelerated_aging_threshold        = 80
    accelerated_aging_scaling_factor   = 2
    icmp_unreachable_rate              = 200
    multicast_route_setup_buffering    = false
    max_pending_mcast_pkts_per_session = 1000

    # Block Configuration
    icmpv6_rate_limit = {
      bucket_size = 100
      packet_rate = 100
    }

    nat = {
      dipp_oversub = "1x"
    }

    nat64 = {
      ipv6_min_network_mtu = 1280
    }

    jumbo_frame = {
      mtu = 9191
    }

    config = {
      rematch = true
    }

    # Packet Buffer Protection Configuration
    packet_buffer_protection_enable                  = false
    packet_buffer_protection_monitor_only            = false
    packet_buffer_protection_alert                   = 50
    packet_buffer_protection_activate                = 80
    packet_buffer_protection_block_countdown         = 80
    packet_buffer_protection_block_hold_time         = 60
    packet_buffer_protection_block_duration_time     = 3600
    packet_buffer_protection_use_latency             = false
    packet_buffer_protection_latency_alert           = 50
    packet_buffer_protection_latency_activate        = 200
    packet_buffer_protection_latency_max_tolerate    = 500
    packet_buffer_protection_latency_block_countdown = 500

  }
}


# Set sessions with the rematch config value set to false; in this case, you don’t need to provide the {} config object.
# resource "scm_session_setting" "session_example_config_false" {
#   # Corresponds to: "folder": "All"
#   folder = "All"

#   session_settings = {

#     # Basic Session Settings
#     dhcp_bcast_session_on              = false
#     erspan                             = false
#     ipv6_firewalling                   = true
#     rematch                            = false
#     accelerated_aging_enable           = true
#     accelerated_aging_threshold        = 80
#     accelerated_aging_scaling_factor   = 2
#     icmp_unreachable_rate              = 200
#     multicast_route_setup_buffering    = false
#     max_pending_mcast_pkts_per_session = 1000

#     # Block Configuration
#     icmpv6_rate_limit = {
#       bucket_size = 100
#       packet_rate = 100
#     }

#     nat = {
#       dipp_oversub = "1x"
#     }

#     nat64 = {
#       ipv6_min_network_mtu = 1280
#     }

#     jumbo_frame = {
#       mtu = 9191
#     }

#     # Packet Buffer Protection Configuration
#     packet_buffer_protection_enable                  = false
#     packet_buffer_protection_monitor_only            = false
#     packet_buffer_protection_alert                   = 50
#     packet_buffer_protection_activate                = 80
#     packet_buffer_protection_block_countdown         = 80
#     packet_buffer_protection_block_hold_time         = 60
#     packet_buffer_protection_block_duration_time     = 3600
#     packet_buffer_protection_use_latency             = false
#     packet_buffer_protection_latency_alert           = 50
#     packet_buffer_protection_latency_activate        = 200
#     packet_buffer_protection_latency_max_tolerate    = 500
#     packet_buffer_protection_latency_block_countdown = 500

#   }
# }