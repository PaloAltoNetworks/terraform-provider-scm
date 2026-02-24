resource "scm_authentication_profile" "global_radius_access" {
  name              = "test_auth_profile_settings"
  folder            = "Prisma Access"
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
      server_profile = "CHAP_only_rsp_11" // server_profile added should exist 
    }
  }
}

resource "scm_authentication_setting" "auth_settings" {
  folder = "Prisma Access"
  authentication = {
    authentication_profile = scm_authentication_profile.global_radius_access.name
    certificate_profile    = "EDL-Hosting-Service-Profile" //this cer profile should be created separately
  }
}