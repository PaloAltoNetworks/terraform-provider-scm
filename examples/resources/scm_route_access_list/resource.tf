resource "scm_route_access_list" "example" {
  folder = "ngfw-shared"
  name   = "EXAMPLE-ACL"
  type = {
    ipv4 = {
      ipv4_entry = [
        {
          name   = 10
          action = "permit"
          destination_address = {
            entry = {
              address  = "10.0.0.0"
              wildcard = "0.0.0.7"
            }
          }
        }
      ]
    }
  }
}