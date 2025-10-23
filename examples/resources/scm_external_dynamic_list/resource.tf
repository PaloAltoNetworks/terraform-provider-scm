# This resource creates a domain-based External Dynamic List (EDL).
# The EDL will fetch a list of domains from the specified URL daily.
resource "scm_external_dynamic_list" "scm_edl_1" {
  folder = "All"
  name   = "scm_edl_1"

  type = {
    domain = {
      description = "List of malicious domains to block, updated daily."
      url         = "http://some-threat-feed.com/domains.txt"
      recurring = {
        daily = {
          at = "03"
        }
      }
    }
  }
}

resource "scm_external_dynamic_list" "scm_edl_2" {
  folder = "All"
  name   = "scm_edl_2"

  type = {
    ip = {
      description = "IP threat feed with basic auth, updated hourly."
      url         = "https://threats.example.com/ips.txt"
      recurring = {
        hourly = {} # The empty object indicates an hourly schedule.
      }
    }
  }
}

resource "scm_external_dynamic_list" "scm_edl_3" {
  folder = "All"
  name   = "scm_edl_3"

  type = {
    url = {
      description = "List of phishing URLs, updated every Monday at 2 AM."
      url         = "https://phish-block.example.com/urls.txt"
      recurring = {
        weekly = {
          day_of_week = "monday"
          at          = "02" # Time is in 24-hour format (e.g., "02" for 2 AM).
        }
      }
    }
  }
}

resource "scm_external_dynamic_list" "scm_edl_4" {
  folder = "All"
  name   = "scm_edl_4"

  type = {
    predefined_ip = {
      description = "Palo Alto Networks-provided list of high-risk IP addresses."
      url         = "panw-highrisk-ip-list" # This is a predefined URL keyword.
    }
  }
}

resource "scm_external_dynamic_list" "scm_edl_5" {
  folder = "All"
  name   = "scm_edl_5"

  type = {
    ip = {
      description         = "IP threat feed that requires authentication."
      url                 = "https://secure-feed.example.com/ips.txt"
      certificate_profile = "test-cert-list-qekwys"

      # This is the block that will trigger the encrypted value handling
      auth = {
        username = "my-api-user"
        password = "a-very-secret-password-123!"
      }

      recurring = {
        five_minute = {} # Check for updates every five minutes
      }
    }
  }
}