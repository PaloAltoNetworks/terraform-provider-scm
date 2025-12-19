
resource "scm_motd_banner_setting" "motd_example" {
  # The 'folder' field in your JSON maps to the 'folder' argument in the SCM provider.
  # "All" means it applies to all devices managed by the Panorama/Cloud Managed environment.
  folder = "All"
  motd_and_banner = {
    # MOTD Configuration (motd_and_banner fields)
    motd_enable               = true
    message                   = "Test Message"
    motd_do_not_display_again = false
    motd_title                = "Test title"
    motd_color                = "color1"
    severity                  = "warning"

    # Banner Header Configuration
    banner_header            = "Test banner Header"
    banner_header_color      = "color1"
    banner_header_text_color = "color1"

    # Banner Footer Configuration
    banner_header_footer_match = true
    banner_footer              = "Test banner Footer"
    banner_footer_color        = "color1"
    banner_footer_text_color   = "color1"
  }
}
