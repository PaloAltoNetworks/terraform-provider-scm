# TCP Service with multiple destination ports custom timeout
resource "scm_service" "scm_service_tcp_ports" {
  folder      = "Prisma Access"
  name        = "scm_service_tcp_ports"
  description = "Managed by Terraform"
  protocol = {
    tcp = {
      port = "80,443"
      override = {
        timeout = 3600
      }
    }
  }
}

# TCP Service with source port, destination port and custom timeout values
resource "scm_service" "scm_service_tcp_port_src_dst" {
  folder      = "Prisma Access"
  name        = "scm_service_tcp_port_src_dst"
  description = "Managed by Terraform"
  protocol = {
    tcp = {
      port        = "80"
      source_port = "49152-65535"
      override = {
        timeout           = 3600
        halfclose_timeout = 240
        timewait_timeout  = 30
      }
    }
  }
}

# UDP Service with single destination port
resource "scm_service" "scm_service_udp_port" {
  folder      = "Prisma Access"
  name        = "scm_service_udp_port"
  description = "Managed by Terraform"
  protocol = {
    udp = {
      port = "53"
    }
  }
}
