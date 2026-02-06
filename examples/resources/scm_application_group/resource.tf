# First, create some applications that will be used in the application group.
resource "scm_application" "scm_ag_app_1" {
  folder      = "Prisma Access"
  name        = "scm_ag_app_1"
  description = "First test application"
  category    = "business-systems"
  subcategory = "database"
  technology  = "client-server"
  risk        = 3
}

resource "scm_application" "scm_ag_app_2" {
  folder      = "Prisma Access"
  name        = "scm_ag_app_2"
  description = "Second test application"
  category    = "business-systems"
  subcategory = "database"
  technology  = "client-server"
  risk        = 4
}

# Create the application group that references the applications above.
resource "scm_application_group" "scm_app_group_1" {
  folder = "Prisma Access"
  name   = "scm_app_group_1"
  members = [
    scm_application.scm_ag_app_1.name,
    scm_application.scm_ag_app_2.name,
  ]
}