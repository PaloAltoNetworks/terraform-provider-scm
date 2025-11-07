resource "scm_dns_security_profile" "scm_dns_security_profile_base" {
  folder = "All"
  name   = "dns_base"
}

resource "scm_dns_security_profile" "scm_dns_security_categories" {
  folder      = "All"
  name        = "test_dns_sec_categories"
  description = "dns security profile w/ dns security categories"

  botnet_domains = {
    dns_security_categories = [
      {
        name = "pan-dns-sec-recent"
      },
      {
        name           = "pan-dns-sec-grayware"
        action         = "allow"
        log_level      = "high"
        packet_capture = "disable"
      },
      {
        name           = "pan-dns-sec-proxy"
        action         = "block"
        log_level      = "default"
        packet_capture = "single-packet"
      },
      {
        name           = "pan-dns-sec-phishing"
        action         = "sinkhole"
        log_level      = "critical"
        packet_capture = "extended-capture"
      },
      {
        name           = "pan-dns-sec-malware"
        action         = "default"
        log_level      = "informational"
        packet_capture = "disable"
      },
    ]
  }
}

resource "scm_dns_security_profile" "scm_dns_lists" {
  folder      = "All"
  name        = "test_dns_lists"
  description = "dns security profile w/ dns lists"

  botnet_domains = {
    dns_lists = [
      {
        name           = "default-paloalto-dns"
        packet_capture = "disable"

        action = {
          alert = {}
        }
      },
      {
        name           = "update-edl"
        packet_capture = "extended-capture"

        action = {
          allow = {}
        }
      }
    ]
  }
}

resource "scm_dns_security_profile" "scm_dns_sinkhole" {
  folder      = "All"
  name        = "test_dns_sinkhole"
  description = "dns security profile w/ sinkhole"

  botnet_domains = {
    sinkhole = {
      ipv4_address = "127.0.0.1"
      ipv6_address = "::1"
    }
  }
}

resource "scm_dns_security_profile" "scm_dns_whitelist" {
  folder      = "All"
  name        = "test_dns_whitelist"
  description = "dns security profile w/ whitelist"

  botnet_domains = {
    whitelist = [
      {
        name = "example.com"
      },
      {
        name        = "example2.com"
        description = "creating whitelist"
      }
    ]
  }
}

resource "scm_dns_security_profile" "scm_dns_all" {
  folder      = "All"
  name        = "test_dns_all_test"
  description = "dns security profile w/ all"

  botnet_domains = {
    dns_security_categories = [
      {
        name           = "pan-dns-sec-ddns"
        action         = "block"
        log_level      = "low"
        packet_capture = "disable"
      },
    ]
    dns_lists = [
      {
        name           = "scm_edl_1"
        packet_capture = "single-packet"

        action = {
          block = {}
        }
      }
    ]
    sinkhole = {
      ipv4_address = "pan-sinkhole-default-ip"
      ipv6_address = "::1"
    }
    whitelist = [
      {
        name        = "ebay.com"
        description = "creating whitelist"
      },
    ]
  }
}