resource "scm_http_header_profile" "scm_http_header_profile_1" {
  folder = "All"
  name   = "base_http_header_profile_1"
}

resource "scm_http_header_profile" "scm_http_header_profile_2" {
  folder = "All"
  name   = "simple_http_header_profile_2"

  http_header_insertion = [
    {
      name = "header_one"
      type = [
        {
          name    = "Google Apps Access Control"
          domains = ["*.google.com"]
          headers = [
            {
              name   = "X-GooGAppls-Allowed-Domains"
              header = "X-GooGAppls-Allowed-Domains"
              value  = "user-allowed"
            }
          ]
        }
      ]
    }
  ]
}

resource "scm_http_header_profile" "scm_http_header_profile_3" {
  folder = "All"
  name   = "complete_http_header_profile_3"

  http_header_insertion = [
    {
      name = "header_insertion_one"
      type = [
        {
          name    = "Dropbox Network Control"
          domains = ["*.db.tt", "*.dropbox.com", "dropboxformum.com"]
          headers = [
            {
              name   = "X-Dropbox-allowed-Team-Ids"
              header = "X-Dropbox-allowed-Team-Ids"
              value  = "dropbox-users"
            },
            {
              name   = "custom_header"
              header = "custom_header"
              value  = "10-header"
            },
          ]
        }
      ]
    },
    {
      name = "header_insertion_two"
      type = [
        {
          name    = "Microsoft Office365 Tenant Restrictions"
          domains = ["login.mircosoft.com", "login.mircosoftonline.com", "login.windows.net"]
          headers = [
            {
              name   = "Restrict-Access-Context"
              header = "Restrict-Access-Context"
              value  = "denied-context"
            },
            {
              name   = "Restrict-Access-To-Tenants"
              header = "Restrict-Access-To-Tenants"
              value  = "denied-tenants"
            },
          ]
        }
      ]
    },
    {
      name = "header_insertion_three"
      type = [
        {
          name    = "Dynamic Fields"
          domains = ["custom_domain"]
          headers = [
            {
              name   = "Authorization"
              header = "Authorization"
              value  = "auth"
            }
          ]
        }
      ]
    },
    {
      name = "header_insertion_four"
      type = [
        {
          name    = "Youtube Safe Search"
          domains = ["m.youtube.com", "www.youtube.com"]
          headers = [
            {
              name   = "Youtube-Restrict"
              header = "Youtube-Restrict"
              value  = "denied-youtube"
            }
          ]
        }
      ]
    },
    {
      name = "header_insertion_five"
      type = [
        {
          name    = "Custom"
          domains = ["custom_1", "custom_2"]
          headers = [
            {
              name   = "custom_header"
              header = "custom_header"
              value  = "custom"
            }
          ]
        }
      ]
    },
  ]
}