# Custom Application
resource "scm_application" "scm_media_app" {
  folder                  = "Shared"
  name                    = "scm_media_app"
  description             = "Managed by Terraform"
  category                = "media"
  subcategory             = "gaming"
  technology              = "client-server"
  risk                    = "4"
  evasive_behavior        = true
  pervasive_use           = true
  consume_big_bandwidth   = true
  has_known_vulnerability = true
  prone_to_misuse         = true
}

# Custom Application with tweaked timeouts and additional settings
resource "scm_application" "scm_risky_app" {
  folder                   = "Shared"
  name                     = "scm_risky_app"
  description              = "Managed by Terraform"
  category                 = "media"
  subcategory              = "gaming"
  technology               = "client-server"
  risk                     = "4"
  timeout                  = 3600
  tcp_timeout              = 3600
  tcp_half_closed_timeout  = 60
  tcp_time_wait_timeout    = 10
  evasive_behavior         = true
  pervasive_use            = true
  consume_big_bandwidth    = true
  has_known_vulnerability  = true
  prone_to_misuse          = true
  tunnel_other_application = true
  tunnel_applications      = true
  no_appid_caching         = true
  parent_app               = "bittorrent"
}

# Custom Application based on Custom Signature
resource "scm_application" "scm_custom_app" {
  folder      = "Shared"
  name        = "scm_custom_app"
  description = "Managed by Terraform"
  category    = "media"
  subcategory = "gaming"
  technology  = "client-server"
  risk        = "4"
  signature = [
    {
      name       = "Custom Signature"
      comment    = "Managed by Terraform"
      scope      = "session"
      order_free = false
      and_condition = [
        {
          name = "Example Condition"
          operator = {
            pattern_match = {
              context = "file-data"
              pattern = "^malware$"
            }
          }
        }
      ]
    }
  ]
}