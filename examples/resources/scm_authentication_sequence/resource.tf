resource "scm_authentication_profile" "test_ui_example" {
  name              = "Test_UI"
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

resource "scm_authentication_sequence" "test_sequence" {
  name                    = "test_auth_sequence_1"
  folder                  = "All"
  authentication_profiles = ["test_auth_profile"] // authentication_profiles added should exist 
  use_domain_find_profile = false
}

resource "scm_authentication_sequence" "test_sequence_2" {
  name                    = "test_auth_sequence_2"
  folder                  = "All"
  authentication_profiles = [scm_authentication_profile.test_ui_example.name, "test_auth_profile"] // authentication_profiles added should exist 
  use_domain_find_profile = false
}