---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_authentication_sequence_list Data Source - scm"
subcategory: "NetSec"
description: |-
  Retrieves a listing of config items.
---

# scm_authentication_sequence_list (Data Source)

Retrieves a listing of config items.

## Example Usage

```terraform
data "scm_authentication_sequence_list" "example" {
  folder = "Shared"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `device` (String) The Device param.
- `folder` (String) The Folder param.
- `limit` (Number) The Limit param. A limit of -1 will return all configured items. Default: `200`.
- `name` (String) The Name param.
- `offset` (Number) The Offset param. Default: `0`.
- `snippet` (String) The Snippet param.

### Read-Only

- `data` (Attributes List) The Data param. (see [below for nested schema](#nestedatt--data))
- `tfid` (String) The Terraform ID.
- `total` (Number) The Total param.

<a id="nestedatt--data"></a>
### Nested Schema for `data`

Read-Only:

- `authentication_profiles` (List of String) The AuthenticationProfiles param.
- `id` (String) UUID of the resource.
- `name` (String) The Name param.
- `use_domain_find_profile` (Boolean) The UseDomainFindProfile param. Default: `true`.
