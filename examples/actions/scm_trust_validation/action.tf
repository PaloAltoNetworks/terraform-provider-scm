# Validate a trust relationship between tenants.
#
# Usage:
#   terraform apply -invoke=action.scm_trust_validation.verify

action "scm_trust_validation" "verify" {
  config {
    tsg                   = "1234567890"
    donor_tenant_name     = "DONOR-LAB"
    recipient_tenant_name = "RECIPIENT-LAB"
    trust_id              = 1337
    psk                   = "6d870763-ce10-41b6-806e-c78e84c3167a"
  }
}
