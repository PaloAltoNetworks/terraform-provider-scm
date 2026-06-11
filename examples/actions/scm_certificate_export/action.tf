# Export a certificate from Strata Cloud Manager.
#
# Limitation:
#   Certificate Response payload is currently written to stdout - no file export supported for now
# Usage:
#   terraform apply -invoke=action.scm_certificate_export.backup

action "scm_certificate_export" "backup" {
  config {
    id     = "f23e1c22-de94-44cd-b67f-36f2516618a7"
    format = "pem"
  }
}
