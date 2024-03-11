resource "scm_hip_object" "example" {
  folder      = "Shared"
  name        = "myExample"
  description = "Made by Terraform"
  disk_backup = {
    criteria = {
      is_installed = true
      last_backup_time = {
        within = {
          days = 1
        }
      }
    }
    exclude_vendor = false
  }
}
