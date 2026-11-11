# Initiate ZTNA tenant onboarding. This action starts the onboarding process
# for the tenant associated with the authenticated credentials.
#
# Usage:
#   terraform apply -invoke=action.scm_ztna_tenant_start_onboarding.run

action "scm_ztna_tenant_start_onboarding" "run" {
  config {}
}
