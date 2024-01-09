---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_ipsec_crypto_profile Data Source - scm"
subcategory: ""
description: |-
  Retrieves a config item.
---

# scm_ipsec_crypto_profile (Data Source)

Retrieves a config item.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The Id param.

### Read-Only

- `ah` (Attributes) The Ah param. (see [below for nested schema](#nestedatt--ah))
- `dh_group` (String) phase-2 DH group (PFS DH group). String must be one of these: `"no-pfs"`, `"group1"`, `"group2"`, `"group5"`, `"group14"`, `"group19"`, `"group20"`. Default: `"group2"`.
- `esp` (Attributes) The Esp param. (see [below for nested schema](#nestedatt--esp))
- `lifesize` (Attributes) The Lifesize param. (see [below for nested schema](#nestedatt--lifesize))
- `lifetime` (Attributes) The Lifetime param. (see [below for nested schema](#nestedatt--lifetime))
- `name` (String) Alphanumeric string begin with letter: [0-9a-zA-Z._-]. String length must not exceed 31 characters.
- `tfid` (String) The Terraform ID.

<a id="nestedatt--ah"></a>
### Nested Schema for `ah`

Read-Only:

- `authentications` (List of String) The Authentications param.


<a id="nestedatt--esp"></a>
### Nested Schema for `esp`

Read-Only:

- `authentications` (List of String) Authentication algorithm.
- `encryptions` (List of String) Encryption algorithm.


<a id="nestedatt--lifesize"></a>
### Nested Schema for `lifesize`

Read-Only:

- `gb` (Number) specify lifesize in gigabytes(GB). Value must be between 1 and 65535.
- `kb` (Number) specify lifesize in kilobytes(KB). Value must be between 1 and 65535.
- `mb` (Number) specify lifesize in megabytes(MB). Value must be between 1 and 65535.
- `tb` (Number) specify lifesize in terabytes(TB). Value must be between 1 and 65535.


<a id="nestedatt--lifetime"></a>
### Nested Schema for `lifetime`

Read-Only:

- `days` (Number) specify lifetime in days. Value must be between 1 and 365.
- `hours` (Number) specify lifetime in hours. Value must be between 1 and 65535.
- `minutes` (Number) specify lifetime in minutes. Value must be between 3 and 65535.
- `seconds` (Number) specify lifetime in seconds. Value must be between 180 and 65535.