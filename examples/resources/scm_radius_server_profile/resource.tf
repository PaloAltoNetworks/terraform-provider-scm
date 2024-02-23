resource "scm_radius_server_profile" "example" {
  folder  = "Shared"
  retries = 5
  timeout = 7
  protocol = {
    chap = true
  }
  servers = [
    {
      name       = "server1"
      ip_address = "11.2.3.5"
      secret     = "secret"
    },
  ]
}
