---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_file_blocking_profile Data Source - scm"
subcategory: ""
description: |-
  Retrieves a config item.
---

# scm_file_blocking_profile (Data Source)

Retrieves a config item.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The Id param.

### Read-Only

- `description` (String) The Description param.
- `name` (String) The Name param.
- `rules` (Attributes List) The Rules param. (see [below for nested schema](#nestedatt--rules))
- `tfid` (String) The Terraform ID.

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Read-Only:

- `action` (String) The Action param. String must be one of these: `"alert"`, `"block"`, `"continue"`. Default: `"alert"`.
- `applications` (List of String) The Applications param. List must contain at least 1 elements.
- `direction` (String) The Direction param. String must be one of these: `"download"`, `"upload"`, `"both"`. Default: `"both"`.
- `file_types` (List of String) The FileTypes param. List must contain at least 1 elements.
- `name` (String) The Name param.