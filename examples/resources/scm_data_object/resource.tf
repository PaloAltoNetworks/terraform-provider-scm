resource "scm_data_object" "scm_data_filtering_profile_predefined" {
  name   = "tf-data-object-pre"
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

resource "scm_data_object" "scm_data_filtering_profile_file_properties" {
  name   = "tf-data-object-fp"
  folder = "ngfw-shared"
  pattern_type = {
    file_properties = {
      pattern = [
        {
          file_type      = "pdf"
          name           = "test_pdf"
          property_value = "test_value"
          file_property  = "panav-rsp-pdf-dlp-author"
        }
      ]
    }
  }
}