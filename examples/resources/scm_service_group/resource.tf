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

# Service Group containing multiple services
resource "scm_service_group" "scm_servicegroup" {
  folder = "All"
  name   = "scm_servicegroup"
  members = [
    scm_service.scm_service_tcp_ports.name,
    scm_service.scm_service_udp_port.name
  ]
}

# Service Group containing multiple services and another servicegroup
resource "scm_service_group" "scm_servicegroup_nested" {
  folder = "All"
  name   = "scm_servicegroup_nested"
  members = [
    scm_service.scm_service_tcp_ports.name,
    scm_service.scm_service_udp_port.name,
    scm_service_group.scm_servicegroup.name
  ]
}
