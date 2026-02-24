resource "scm_authentication_profile" "global_radius_access" {
  name              = "test_auth_profile_radius_1"
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
      server_profile = "CHAP_only_rsp_1" // server_profile added should exist 
    }
  }

  # Single Sign-On Configuration
  single_sign_on = {
    realm = "EXAMPLE.COM"
  }
}

resource "scm_authentication_profile" "global_db_access" {
  name              = "test_auth_global_db_1"
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