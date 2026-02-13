resource "scm_application_filter" "scm_application_filter_1" {
  folder   = "Prisma Access"
  name     = "scm_application_filter_1"
  category = ["business-systems"]
  risk     = [1]
  evasive  = true
  tagging = {
    no_tag = true
  }
}