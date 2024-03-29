---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_internal_dns_server_list Data Source - scm"
subcategory: "NetSec"
description: |-
  Retrieves a listing of config items.
---

# scm_internal_dns_server_list (Data Source)

Retrieves a listing of config items.

## Example Usage

```terraform
data "scm_internal_dns_server_list" "example" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `limit` (Number) The Limit param. A limit of -1 will return all configured items. Default: `200`.
- `name` (String) The Name param.
- `offset` (Number) The Offset param. Default: `0`.

### Read-Only

- `data` (Attributes List) The Data param. (see [below for nested schema](#nestedatt--data))
- `tfid` (String) The Terraform ID.
- `total` (Number) The Total param.

<a id="nestedatt--data"></a>
### Nested Schema for `data`

Read-Only:

- `domain_names` (List of String) The DomainNames param.
- `id` (String) UUID of the resource.
- `name` (String) The Name param.
- `primary` (String) The Primary param.
- `secondary` (String) The Secondary param.
