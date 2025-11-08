resource "scm_radius_server_profile" "chap_radius_profile" {
  # Standard fields
  name    = "CHAP_only_rsp_1"
  folder  = "All"
  retries = 5
  timeout = 60

  protocol = {
    chap = {}
  }

  # Server list
  server = [
    {
      name       = "Chap_Server_Primary"
      ip_address = "10.1.1.10"
      port       = 1812
      secret     = "-AQ==lhyuV6U/j9Trb9JL9L0UoBecg9Y=kTOWntGhZ1KFyLD+etKQ3g=="
    },
  ]
}

resource "scm_radius_server_profile" "pap_radius_profile" {
  # Standard fields
  name    = "pap_only_rsp_1"
  folder  = "All"
  retries = 5
  timeout = 60

  protocol = {
    pap = {}
  }

  # Server list
  server = [
    {
      name       = "pap_Server_Primary"
      ip_address = "10.1.1.10"
      port       = 1812
      secret     = "-AQ==lhyuV6U/j9Trb9JL9L0UoBecg9Y=kTOWntGhZ1KFyLD+etKQ3g=="
    },
  ]
}