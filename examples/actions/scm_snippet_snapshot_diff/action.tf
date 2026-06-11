# Show the diff of a specific object between two snippet snapshot versions.
#
# Usage:
#   terraform apply -invoke=action.scm_snippet_snapshot_diff.inspect

action "scm_snippet_snapshot_diff" "inspect" {
  config {
    snippet_id        = "f23e1c22-de94-44cd-b67f-36f2516618a7"
    object_id         = "ddad1e64-0b64-41a4-b361-c6199761a8f2"
    version           = 100
    comparing_version = 101
  }
}
