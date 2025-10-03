# TCP Service for a web server
resource "scm_service" "scm_sg_service_1" {
  folder      = "Shared"
  name        = "scm_sg_service_1"
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
resource "scm_service" "scm_sg_service_2" {
  folder      = "Shared"
  name        = "scm_sg_service_2"
  description = "Service for UDP DNS traffic"
  protocol = {
    udp = {
      port = "53"
    }
  }
}

# Service Group containing both services
resource "scm_service_group" "scm_service_group_1" {
  folder = "Shared"
  name   = "scm_service_group_1"
  members = [
    scm_service.scm_sg_service_1.name,
    scm_service.scm_sg_service_2.name
  ]
}
