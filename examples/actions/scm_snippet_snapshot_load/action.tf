# Load a specific version of a snippet snapshot into the candidate configuration.
#
# Usage:
#   terraform apply -invoke=action.scm_snippet_snapshot_load.restore

action "scm_snippet_snapshot_load" "restore" {
  config {
    id      = "f23e1c22-de94-44cd-b67f-36f2516618a7"
    version = "100"
  }
}
