resource "scm_ike_gateway" "example" {
  folder = "Remote Networks"
  name   = "gw1"
  peer_address = {
    dynamic_address = true
  }
  protocol = {
    version = "ikev2"
    ikev2 = {
      dpd = {
        enable = false
      }
    }
  }
  authentication = {
    pre_shared_key = {
      key = "secret"
    }
  }
}
