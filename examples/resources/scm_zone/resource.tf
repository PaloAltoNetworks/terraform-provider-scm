#
# Creates an empty layer3 zone
#

resource "scm_zone" "scm_layer3_zone" {
  name   = "scm_layer3_zone"
  folder = "ngfw-shared"
  network = {
    layer3 = []
  }
}

#
# Creates an empty layer2 zone
#

resource "scm_zone" "scm_layer2_zone" {
  name   = "scm_layer2_zone"
  folder = "ngfw-shared"
  network = {
    layer2 = []
  }
}

#
# Creates an empty tap zone
#

resource "scm_zone" "scm_tap_zone" {
  name   = "scm_tap_zone"
  folder = "ngfw-shared"
  network = {
    tap = []
  }
}

#
# Creates an empty vwire zone
#

resource "scm_zone" "scm_vwire_zone" {
  name   = "scm_vwire_zone"
  folder = "ngfw-shared"
  network = {
    virtual_wire = []
  }
}

#
# Creates a layer3 zone
# Requires Interface $scm_l3_interface to exist
#

resource "scm_zone" "scm_layer3_zone_complex" {
  name   = "scm_layer3_zone_complex"
  folder = "ngfw-shared"
  network = {
    layer3                          = ["$scm_l3_interface"]
    zone_protection_profile         = "best-practice"
    enable_packet_buffer_protection = true
  }
  enable_device_identification = true
  device_acl = {
    include_list = ["198.18.1.0/24"]
    exclude_list = ["198.18.2.0/24"]
  }
  enable_user_identification = true
  user_acl = {
    include_list = ["198.18.3.0/24"]
    exclude_list = ["198.18.4.0/24"]
  }
}
