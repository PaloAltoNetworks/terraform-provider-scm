resource "scm_nat_rule" "example" {
  folder      = "folder1"
  name        = "example_nat1"
  description = "Example source nat"
  sources     = ["10.10.10.10/24"]
  froms       = ["any"]
  tos         = ["untrust"]
  source_translation = {
    translated_address_single = "192.168.10.100"
    bi_directional            = "no"
  }
}
