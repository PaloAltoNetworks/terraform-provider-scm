# Convert a snippet snapshot (e.g., from shared to local).
#
# Usage:
#   terraform apply -invoke=action.scm_snippet_snapshot_convert.localize

action "scm_snippet_snapshot_convert" "localize" {
  config {
    id         = "f23e1c22-de94-44cd-b67f-36f2516618a7"
    keep_local = true
  }
}
