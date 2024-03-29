---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_decryption_rule_list Data Source - scm"
subcategory: "NetSec"
description: |-
  Retrieves a listing of config items.
---

# scm_decryption_rule_list (Data Source)

Retrieves a listing of config items.

## Example Usage

```terraform
data "scm_decryption_rule_list" "example" {
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

- `action` (String) The Action param. String must be one of these: `"decrypt"`, `"no-decrypt"`.
- `categories` (List of String) The Categories param.
- `description` (String) The Description param.
- `destination_hips` (List of String) The DestinationHips param.
- `destinations` (List of String) The Destinations param.
- `disabled` (Boolean) The Disabled param.
- `froms` (List of String) The Froms param.
- `id` (String) UUID of the resource.
- `log_fail` (Boolean) The LogFail param.
- `log_setting` (String) The LogSetting param.
- `log_success` (Boolean) The LogSuccess param.
- `name` (String) The Name param.
- `negate_destination` (Boolean) The NegateDestination param.
- `negate_source` (Boolean) The NegateSource param.
- `profile` (String) The Profile param.
- `services` (List of String) The Services param.
- `source_hips` (List of String) The SourceHips param.
- `source_users` (List of String) The SourceUsers param.
- `sources` (List of String) The Sources param.
- `tags` (List of String) The Tags param.
- `tos` (List of String) The Tos param.
- `type` (Attributes) The Type param. (see [below for nested schema](#nestedatt--data--type))

<a id="nestedatt--data--type"></a>
### Nested Schema for `data.type`

Read-Only:

- `ssl_forward_proxy` (Boolean) The SslForwardProxy param.
- `ssl_inbound_inspection` (String) add the certificate name for SSL inbound inspection.
