# TCP Service for a web server
resource "scm_service" "scm_service_1" {
  folder      = "Shared"
  name        = "scm_service_1"
  description = "Service for TCP traffic to web server"
  protocol = {
    tcp = {
      port = "80,443"
      override = {
        timeout = 3600 # 1 hour
      }
    }
  }
}

# UDP Service for DNS
resource "scm_service" "scm_service_2" {
  folder      = "Shared"
  name        = "scm_service_2"
  description = "Service for UDP DNS traffic"
  protocol = {
    udp = {
      port = "53"
    }
  }
}
