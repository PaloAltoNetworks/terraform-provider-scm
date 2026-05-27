# Publish a snippet snapshot to subscriber tenants.
#
# Usage:
#   terraform apply -invoke=action.scm_snippet_snapshot_publish.release

action "scm_snippet_snapshot_publish" "release" {
  config {
    id         = "f23e1c22-de94-44cd-b67f-36f2516618a7"
    name       = "example-snippet"
    version    = 100
    validation = true
    tsgs       = ["1234567890", "1234567891"]
  }
}
