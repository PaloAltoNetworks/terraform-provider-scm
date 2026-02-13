# This resource creates a new HIP Profile.
resource "scm_hip_profile" "scm_hip_profile_1" {
  folder      = "Prisma Access"
  name        = "scm_hip_profile_1"
  description = "A HIP profile created by Terraform"
  match       = "\"is-win\" and \"is-anti-malware-and-rtp-enabled\""
}
