# Bootstrap Prisma Access application defaults (certificates, config nodes) on a
# fresh tenant. One-time, tenant-wide, and effectively irreversible — a fresh
# tenant needs this before other automation. Requires Terraform >= 1.14 to invoke.
#
# Usage:
#   terraform apply -invoke=action.scm_application_defaults.bootstrap

action "scm_application_defaults" "bootstrap" {
  config {}
}
