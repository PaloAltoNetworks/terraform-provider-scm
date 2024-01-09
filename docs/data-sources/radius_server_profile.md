---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_radius_server_profile Data Source - scm"
subcategory: ""
description: |-
  Retrieves a config item.
---

# scm_radius_server_profile (Data Source)

Retrieves a config item.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The Id param.

### Read-Only

- `protocol` (Attributes) The Protocol param. (see [below for nested schema](#nestedatt--protocol))
- `retries` (Number) The Retries param. Value must be between 1 and 5.
- `servers` (Attributes List) The Servers param. (see [below for nested schema](#nestedatt--servers))
- `tfid` (String) The Terraform ID.
- `timeout` (Number) The Timeout param. Value must be between 1 and 120.

<a id="nestedatt--protocol"></a>
### Nested Schema for `protocol`

Read-Only:

- `chap` (Boolean) The Chap param.
- `eap_ttls_with_pap` (Attributes) The EapTtlsWithPap param. (see [below for nested schema](#nestedatt--protocol--eap_ttls_with_pap))
- `pap` (Boolean) The Pap param.
- `peap_mschap_v2` (Attributes) The PeapMschapV2 param. (see [below for nested schema](#nestedatt--protocol--peap_mschap_v2))
- `peap_with_gtc` (Attributes) The PeapWithGtc param. (see [below for nested schema](#nestedatt--protocol--peap_with_gtc))

<a id="nestedatt--protocol--eap_ttls_with_pap"></a>
### Nested Schema for `protocol.eap_ttls_with_pap`

Read-Only:

- `anon_outer_id` (Boolean) The AnonOuterId param.
- `radius_cert_profile` (String) The RadiusCertProfile param.


<a id="nestedatt--protocol--peap_mschap_v2"></a>
### Nested Schema for `protocol.peap_mschap_v2`

Read-Only:

- `allow_pwd_change` (Boolean) The AllowPwdChange param.
- `anon_outer_id` (Boolean) The AnonOuterId param.
- `radius_cert_profile` (String) The RadiusCertProfile param.


<a id="nestedatt--protocol--peap_with_gtc"></a>
### Nested Schema for `protocol.peap_with_gtc`

Read-Only:

- `anon_outer_id` (Boolean) The AnonOuterId param.
- `radius_cert_profile` (String) The RadiusCertProfile param.



<a id="nestedatt--servers"></a>
### Nested Schema for `servers`

Read-Only:

- `ip_address` (String) The IpAddress param.
- `name` (String) The Name param.
- `port` (Number) The Port param. Value must be between 1 and 65535.
- `secret` (String, Sensitive) The Secret param. String length must not exceed 64 characters.