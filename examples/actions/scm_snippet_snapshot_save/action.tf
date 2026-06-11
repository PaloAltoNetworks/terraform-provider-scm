# Save a snapshot of a snippet's current configuration.
#
# Usage:
#   terraform apply -invoke=action.scm_snippet_snapshot_save.backup

action "scm_snippet_snapshot_save" "backup" {
  config {
    id          = "f23e1c22-de94-44cd-b67f-36f2516618a7"
    description = "Snapshot before Terraform changes"
  }
}
