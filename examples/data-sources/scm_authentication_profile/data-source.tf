resource "scm_authentication_profile" "global_db_access" {
  name              = "test_auth_profile_ds_db_1"
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

data "scm_authentication_profile" "profile_data" {
  id = scm_authentication_profile.global_db_access.id

}

output "scm_authentication_profile" {
  description = "The UUID of the fetched authentication prfile."
  value       = data.scm_authentication_profile.profile_data.id
}

output "fetched_profile" {
  description = "Fetched Auth Profile"
  value       = data.scm_authentication_profile.profile_data
}