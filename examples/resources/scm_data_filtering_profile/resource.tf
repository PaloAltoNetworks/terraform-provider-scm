resource "scm_data_object" "scm_data_filtering_profile_example" {
  name   = "tf-data-object-df"
  folder = "ngfw-shared"
  pattern_type = {
    predefined = {
      pattern = [
        {
          file_type = ["text/html"]
          name      = "ABA-Routing-Number"
        }
      ]
    }
  }
}

resource "scm_data_filtering_profile" "scm_data_filtering_profile_example" {
  name         = "tf-data-filtering"
  data_capture = false
  description  = "Terraform Description"
  folder       = "ngfw-shared"
  rules = [
    {
      name            = "rule0"
      data_object     = scm_data_object.scm_data_filtering_profile_example.name
      application     = ["any"]
      file_type       = ["any"]
      direction       = "both"
      alert_threshold = 1
      block_threshold = 1
      log_severity    = "informational"
    }
  ]
}


