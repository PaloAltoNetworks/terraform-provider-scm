resource "scm_update_schedule" "us_settings" {
  folder = "All"

  update_schedule = {
    threats = {
      recurring = {
        threshold         = 300
        new_app_threshold = 300
        sync_to_peer      = false

        hourly = {
          at                  = 30
          action              = "download-only"
          disable_new_content = false
        }
      }
    }

    anti_virus = {
      recurring = {
        threshold    = 300
        sync_to_peer = false

        hourly = {
          at     = 30
          action = "download-only"
        }
      }
    }

    wildfire = {
      recurring = {
        every_hour = {
          at           = 30
          action       = "download-only"
          sync_to_peer = true
        }
      }
    }
  }
}


# -- Example 2 : With Daily recurring updates for recurring and anti-virus
resource "scm_update_schedule" "us_settings_daily" {
  folder = "All"

  update_schedule = {
    threats = {
      recurring = {
        threshold         = 300
        new_app_threshold = 300
        sync_to_peer      = false

        daily = {
          at                  = "02:13"
          action              = "download-only"
          disable_new_content = false
        }
      }
    }

    anti_virus = {
      recurring = {
        threshold    = 300
        sync_to_peer = true

        daily = {
          at     = "02:13"
          action = "download-only"
        }
      }
    }

    wildfire = {
      recurring = {
        every_30_mins = {
          at           = 20
          action       = "download-only"
          sync_to_peer = false
        }
      }
    }
  }
}