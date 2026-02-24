resource "scm_authentication_profile" "test_ui_example" {
  name              = "test_auth_profile_settings"
  folder            = "Prisma Access"
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

resource "scm_authentication_setting" "auth_settings" {
  folder = "Prisma Access"
  authentication = {
    authentication_profile = scm_authentication_profile.global_radius_access.name
    certificate_profile    = "EDL-Hosting-Service-Profile" //this cer profile should be created separately
  }
}