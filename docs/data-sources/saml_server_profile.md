---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_saml_server_profile Data Source - scm"
subcategory: ""
description: |-
  Retrieves a config item.
---

# scm_saml_server_profile (Data Source)

Retrieves a config item.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The Id param.

### Read-Only

- `certificate` (String) The Certificate param. String length must not exceed 63 characters.
- `entity_id` (String) The EntityId param. String length must be between 1 and 1024 characters.
- `max_clock_skew` (Number) The MaxClockSkew param. Value must be between 1 and 900.
- `slo_bindings` (String) The SloBindings param. String must be one of these: `"post"`, `"redirect"`.
- `sso_bindings` (String) The SsoBindings param. String must be one of these: `"post"`, `"redirect"`.
- `sso_url` (String) The SsoUrl param. String length must be between 1 and 255 characters.
- `tfid` (String) The Terraform ID.
- `validate_idp_certificate` (Boolean) The ValidateIdpCertificate param.
- `want_auth_requests_signed` (Boolean) The WantAuthRequestsSigned param.