# This resource creates the "all-slots-casino-3" application.
# The attributes are mapped from your JSON payload.
resource "scm_application" "scm_app_1" {
  folder                  = "Shared"
  name                    = "scm_app_1"
  description             = "online casino"
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


resource "scm_application" "scm_app_2" {
  folder                  = "Shared"
  name                    = "scm_app_2"
  description             = "online casino"
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