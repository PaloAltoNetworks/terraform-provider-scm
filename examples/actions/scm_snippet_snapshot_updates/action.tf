# Check for updates to a snippet snapshot from the publisher tenant.
#
# Usage:
#   terraform apply -invoke=action.scm_snippet_snapshot_updates.check

action "scm_snippet_snapshot_updates" "check" {
  config {
    id        = "f23e1c22-de94-44cd-b67f-36f2516618a7"
    tenant_id = "1234567890"
  }
}
