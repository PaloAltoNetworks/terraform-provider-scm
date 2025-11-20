resource "scm_authentication_sequence" "test_sequence" {
  name                    = "test_auth_sequence_1"
  folder                  = "All"
  authentication_profiles = ["test_auth_profile"] // authentication_profiles added should exist 
  use_domain_find_profile = false
}

resource "scm_authentication_sequence" "test_sequence_2" {
  name                    = "test_auth_sequence_2"
  folder                  = "All"
  authentication_profiles = ["Test_UI", "test_auth_profile"] // authentication_profiles added should exist 
  use_domain_find_profile = false
}