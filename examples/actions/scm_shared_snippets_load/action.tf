# Load a shared snippet into the current tenant's configuration.
#
# Usage:
#   terraform apply -invoke=action.scm_shared_snippets_load.import

action "scm_shared_snippets_load" "import" {
  config {
    id         = "f23e1c22-de94-44cd-b67f-36f2516618a7"
    validation = true
  }
}
