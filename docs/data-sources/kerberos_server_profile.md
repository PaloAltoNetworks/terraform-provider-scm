---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_kerberos_server_profile Data Source - scm"
subcategory: ""
description: |-
  Retrieves a config item.
---

# scm_kerberos_server_profile (Data Source)

Retrieves a config item.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The Id param.

### Read-Only

- `servers` (Attributes List) The Servers param. (see [below for nested schema](#nestedatt--servers))
- `tfid` (String) The Terraform ID.

<a id="nestedatt--servers"></a>
### Nested Schema for `servers`

Read-Only:

- `host` (String) The Host param.
- `name` (String) The Name param.
- `port` (Number) The Port param. Value must be between 1 and 65535.