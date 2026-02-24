resource "scm_radius_server_profile" "chap_radius_profile" {
  # Standard fields
  name    = "CHAP_only_rsp"
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

resource "scm_authentication_profile" "global_radius_access" {
  name              = "test_auth_profile_radius"
  folder            = "All"
  user_domain       = "default"
  username_modifier = "%USERINPUT%"
  allow_list        = ["all"]

  # Lockout Configuration
  lockout = {
    failed_attempts = 1
    lockout_time    = 3
  }

  # Authentication Method: RADIUS
  method = {
    radius = {
      checkgroup     = true
      server_profile = scm_radius_server_profile.chap_radius_profile.name // server_profile added should exist 
    }
  }

  # Single Sign-On Configuration
  single_sign_on = {
    realm = "EXAMPLE.COM"
  }
}

resource "scm_authentication_profile" "global_db_access" {
  name              = "test_auth_global_db"
  folder            = "All"
  user_domain       = "default"
  username_modifier = "%USERINPUT%"
  allow_list        = ["all"]

  # Lockout Configuration
  lockout = {
    failed_attempts = 3
    lockout_time    = 1
  }

  # Authentication Method: RADIUS
  method = {
    local_database = {}
  }

  # Single Sign-On Configuration
  single_sign_on = {
    realm = "EXAMPLE.COM"
  }
}