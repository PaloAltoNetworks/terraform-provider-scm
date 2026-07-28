# Initiate ZTNA tenant onboarding. This action starts the onboarding process
# for the tenant associated with the authenticated credentials.
#
# Usage:
#   terraform apply -invoke=action.ztna_tenant_start_onboarding.run

action "ztna_tenant_start_onboarding" "run" {
  config {}
}
