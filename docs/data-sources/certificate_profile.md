---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_certificate_profile Data Source - scm"
subcategory: "NetSec"
description: |-
  Retrieves a config item.
---

# scm_certificate_profile (Data Source)

Retrieves a config item.

## Example Usage

```terraform
data "scm_certificate_profile" "example" {
  id = "1234-56-789"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The Id param.

### Read-Only

- `block_expired_cert` (Boolean) The BlockExpiredCert param.
- `block_timeout_cert` (Boolean) The BlockTimeoutCert param.
- `block_unauthenticated_cert` (Boolean) The BlockUnauthenticatedCert param.
- `block_unknown_cert` (Boolean) The BlockUnknownCert param.
- `ca_certificates` (Attributes List) The CaCertificates param. (see [below for nested schema](#nestedatt--ca_certificates))
- `cert_status_timeout` (String) The CertStatusTimeout param.
- `crl_receive_timeout` (String) The CrlReceiveTimeout param.
- `domain` (String) The Domain param.
- `name` (String) Alphanumeric string [ 0-9a-zA-Z._-]. String length must not exceed 63 characters.
- `ocsp_receive_timeout` (String) The OcspReceiveTimeout param.
- `tfid` (String) The Terraform ID.
- `use_crl` (Boolean) The UseCrl param.
- `use_ocsp` (Boolean) The UseOcsp param.
- `username_field` (Attributes) The UsernameField param. (see [below for nested schema](#nestedatt--username_field))

<a id="nestedatt--ca_certificates"></a>
### Nested Schema for `ca_certificates`

Read-Only:

- `default_ocsp_url` (String) The DefaultOcspUrl param.
- `name` (String) The Name param.
- `ocsp_verify_cert` (String) The OcspVerifyCert param.
- `template_name` (String) The TemplateName param.


<a id="nestedatt--username_field"></a>
### Nested Schema for `username_field`

Read-Only:

- `subject` (String) The Subject param. String must be one of these: `"common-name"`.
- `subject_alt` (String) The SubjectAlt param. String must be one of these: `"email"`.
