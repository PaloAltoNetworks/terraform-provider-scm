---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_qos_policy_rule_list Data Source - scm"
subcategory: "NetSec"
description: |-
  Retrieves a listing of config items.
---

# scm_qos_policy_rule_list (Data Source)

Retrieves a listing of config items.

## Example Usage

```terraform
data "scm_qos_policy_rule_list" "example" {
  folder   = "Shared"
  position = "pre"
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
- `position` (String) The Position param. String must be one of these: `"pre"`, `"post"`. Default: `"pre"`.
- `snippet` (String) The Snippet param.

### Read-Only

- `data` (Attributes List) The Data param. (see [below for nested schema](#nestedatt--data))
- `tfid` (String) The Terraform ID.
- `total` (Number) The Total param.

<a id="nestedatt--data"></a>
### Nested Schema for `data`

Read-Only:

- `action` (Attributes) The Action param. (see [below for nested schema](#nestedatt--data--action))
- `description` (String) The Description param.
- `dscp_tos` (Attributes) The DscpTos param. (see [below for nested schema](#nestedatt--data--dscp_tos))
- `id` (String) UUID of the resource.
- `name` (String) The Name param.
- `schedule` (String) The Schedule param.

<a id="nestedatt--data--action"></a>
### Nested Schema for `data.action`

Read-Only:

- `class` (String) The Class param.


<a id="nestedatt--data--dscp_tos"></a>
### Nested Schema for `data.dscp_tos`

Read-Only:

- `codepoints` (Attributes List) The Codepoints param. (see [below for nested schema](#nestedatt--data--dscp_tos--codepoints))

<a id="nestedatt--data--dscp_tos--codepoints"></a>
### Nested Schema for `data.dscp_tos.codepoints`

Read-Only:

- `name` (String) The Name param.
- `type` (Attributes) The Type param. (see [below for nested schema](#nestedatt--data--dscp_tos--codepoints--type))

<a id="nestedatt--data--dscp_tos--codepoints--type"></a>
### Nested Schema for `data.dscp_tos.codepoints.type`

Read-Only:

- `af` (Attributes) The Af param. (see [below for nested schema](#nestedatt--data--dscp_tos--codepoints--type--af))
- `cs` (Attributes) The Cs param. (see [below for nested schema](#nestedatt--data--dscp_tos--codepoints--type--cs))
- `custom` (Attributes) The Custom param. (see [below for nested schema](#nestedatt--data--dscp_tos--codepoints--type--custom))
- `ef` (Boolean) The Ef param.
- `tos` (Attributes) The Tos param. (see [below for nested schema](#nestedatt--data--dscp_tos--codepoints--type--tos))

<a id="nestedatt--data--dscp_tos--codepoints--type--af"></a>
### Nested Schema for `data.dscp_tos.codepoints.type.af`

Read-Only:

- `codepoint` (String) The Codepoint param.


<a id="nestedatt--data--dscp_tos--codepoints--type--cs"></a>
### Nested Schema for `data.dscp_tos.codepoints.type.cs`

Read-Only:

- `codepoint` (String) The Codepoint param.


<a id="nestedatt--data--dscp_tos--codepoints--type--custom"></a>
### Nested Schema for `data.dscp_tos.codepoints.type.custom`

Read-Only:

- `codepoint` (Attributes) The Codepoint param. (see [below for nested schema](#nestedatt--data--dscp_tos--codepoints--type--custom--codepoint))

<a id="nestedatt--data--dscp_tos--codepoints--type--custom--codepoint"></a>
### Nested Schema for `data.dscp_tos.codepoints.type.custom.codepoint`

Read-Only:

- `binary_value` (String) The BinaryValue param.
- `codepoint_name` (String) The CodepointName param.



<a id="nestedatt--data--dscp_tos--codepoints--type--tos"></a>
### Nested Schema for `data.dscp_tos.codepoints.type.tos`

Read-Only:

- `codepoint` (String) The Codepoint param.
