# This resource creates a comprehensive HIP Object with many criteria enabled.
# This showcases the use of nested objects and lists of objects that the
# corrected provider schema now supports.

resource "scm_hip_object" "scm_hip_object_1" {
  folder      = "Shared"
  name        = "scm_hip_object_1"
  description = "HIP object with multiple advanced criteria configured"

  # Anti-Malware check with a specific vendor and product
  anti_malware = {
    criteria = {
      is_installed         = true
      real_time_protection = "yes"
      virdef_version = {
        not_within = {
          days = 10
        }
      }
    }
    vendor = [{
      name    = "Microsoft"
      product = ["Microsoft Defender"]
    }]
  }

  # Data Loss Prevention check with a specific vendor
  data_loss_prevention = {
    criteria = {
      is_installed = true
      is_enabled   = "yes"
    }
    vendor = [{
      name = "Symantec"
    }]
  }

  # Disk Backup check
  disk_backup = {
    criteria = {
      is_installed = true
      last_backup_time = {
        within = {
          days = 7
        }
      }
    }
    vendor = [{
      name = "Veeam"
    }]
  }

  # Disk Encryption check for specific locations and their states
  disk_encryption = {
    criteria = {
      is_installed = true
      encrypted_locations = [{
        name             = "C:\\"
        encryption_state = { is = "encrypted" }
        }, {
        name             = "D:\\Users\\"
        encryption_state = { is_not = "unencrypted" }
      }]
    }
    vendor = [{
      name = "BitLocker"
    }]
  }

  # Firewall check for a specific vendor
  firewall = {
    criteria = {
      is_installed = true
      is_enabled   = "yes"
    }
    vendor = [{
      name = "Microsoft"
    }]
  }

  # Host Info check for OS and domain
  host_info = {
    criteria = {
      os = {
        contains = {
          microsoft = "Microsoft Windows 11"
        }
      }
      domain = {
        is = "corp.example.com"
      }
    }
  }

  # Mobile Device checks for unmanaged apps and malware
  mobile_device = {
    criteria = {
      jailbroken   = false
      passcode_set = true
      applications = {
        has_unmanaged_app = false
        has_malware = {
          no = {}
        }
      }
    }
  }

  # Network Info check
  network_info = {
    criteria = {
      network = {
        is = {
          wifi = {
            ssid = "Corporate-WLAN"
          }
        }
      }
    }
  }

  # Patch Management check with a list of specific patches
  patch_management = {
    criteria = {
      is_installed = true
      missing_patches = {
        check   = "has-none"
        patches = ["KB4012212", "KB4012213"]
        severity = {
          greater_than = 5
        }
      }
    }
    vendor = [{
      name = "Microsoft"
    }]
  }

  # Custom Checks for specific processes and registry keys with values
  custom_checks = {
    criteria = {
      process_list = [{
        name    = "evil_process.exe"
        running = false
        }, {
        name    = "corp_security_agent.exe"
        running = true
      }]
      registry_key = [{
        name = "HKEY_LOCAL_MACHINE\\Software\\PaloAltoNetworks"
        registry_value = [{
          name       = "AllowRemoteAccess"
          value_data = "false"
        }]
      }]
    }
  }
}

