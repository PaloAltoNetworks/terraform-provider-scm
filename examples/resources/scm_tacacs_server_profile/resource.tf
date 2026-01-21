resource "scm_tacacs_server_profile" "ise_tacacs_example" {
  name                  = "ISE-TACACS11"
  protocol              = "PAP"
  timeout               = 3
  folder                = "All"
  use_single_connection = true

  server = [
    {
      name    = "Server-1"
      address = "10.10.10.10"
      port    = 49
      secret  = "Test Secret1"
    },
    {
      name    = "Server-2"
      address = "10.10.10.10"
      port    = 49
      secret  = "Test Secret1"
    },
    {
      name    = "Server-3"
      address = "10.10.10.10"
      port    = 49
      secret  = "Test Secret1"
    }
  ]
}