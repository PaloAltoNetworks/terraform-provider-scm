# Load a specific configuration version into the candidate configuration.
#
# Usage:
#   terraform apply -invoke=action.scm_config_load.rollback

action "scm_config_load" "rollback" {
  config {
    version = 100
  }
}
