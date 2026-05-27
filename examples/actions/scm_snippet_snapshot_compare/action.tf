# Compare two versions of a snippet snapshot configuration.
#
# Usage:
#   terraform apply -invoke=action.scm_snippet_snapshot_compare.check

action "scm_snippet_snapshot_compare" "check" {
  config {
    id                = "f23d1c22-de94-44cd-b67f-36f2516618a7"
    version           = 100
    comparing_version = 101
  }
}
