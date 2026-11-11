# Initiate ZTNA tenant offboarding. This action marks the tenant for deletion
# and triggers the cleanup process. This operation is irreversible.
#
# Usage:
#   terraform apply -invoke=action.scm_ztna_tenant_start_offboarding.run

action "scm_ztna_tenant_start_offboarding" "run" {
  config {}
}
