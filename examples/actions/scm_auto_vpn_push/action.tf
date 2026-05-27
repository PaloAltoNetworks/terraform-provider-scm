# Push Auto VPN configuration to one or more VPN clusters.
#
# Usage:
#   terraform apply -invoke=action.scm_auto_vpn_push.deploy

action "scm_auto_vpn_push" "deploy" {
  config {
    auto_vpn_devices = [
      {
        name        = "Demo-Cluster"
        refresh_psk = false
      },
    ]
  }
}
