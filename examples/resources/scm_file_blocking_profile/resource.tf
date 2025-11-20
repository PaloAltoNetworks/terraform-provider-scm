resource "scm_file_blocking_profile" "scm_file_blocking_base" {
  folder = "ngfw-shared"
  name   = "base_file_blocking"
}

resource "scm_file_blocking_profile" "scm_file_blocking_profile" {
  folder      = "ngfw-shared"
  name        = "file_blocking_profile_complete"
  description = "alert, block, and continue"
  rules = [
    {
      name   = "block_rule"
      action = "block"
      application = [
        "any"
      ]
      direction = "upload"
      file_type = [
        "any"
      ]
    },
    {
      name   = "block_rule_two"
      action = "block"
      application = [
        "8x8"
      ]
      direction = "upload"
      file_type = [
        "7z",
        "bat",
        "chm",
        "class",
        "cpl",
        "dll",
        "hlp",
        "hta",
        "jar",
        "ocx",
        "pif",
        "scr",
        "torrent",
        "vbe",
        "wsf"
      ]
    },
    {
      name   = "alert_rule"
      action = "alert"
      application = [
        "access-grid",
        "adobe-update"
      ]
      direction = "both"
      file_type = [
        "ico"
      ]
    },
    {
      name   = "continue_rule"
      action = "continue"
      application = [
        "apple-appstore",
        "limelight"
      ]
      direction = "download"
      file_type = [
        "doc",
        "bmp",
        "dsn",
        "dwf"
      ]
    },
  ]
}