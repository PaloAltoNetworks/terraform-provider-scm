---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_authentication_rule_list Data Source - scm"
subcategory: "NetSec"
description: |-
  Retrieves a listing of config items.
---

# scm_authentication_rule_list (Data Source)

Retrieves a listing of config items.

## Example Usage

```terraform
data "scm_authentication_rule_list" "example" {
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

- `authentication_enforcement` (String) the authentication profile name to apply to authentication rule.
- `categories` (List of String) The Categories param.
- `description` (String) The Description param.
- `destination_hips` (List of String) The DestinationHips param.
- `destinations` (List of String) The Destinations param.
- `disabled` (Boolean) The Disabled param. Default: `false`.
- `froms` (List of String) The Froms param.
- `group_tag` (String) The GroupTag param.
- `hip_profiles` (List of String) The HipProfiles param.
- `id` (String) UUID of the resource.
- `log_authentication_timeout` (Boolean) The LogAuthenticationTimeout param. Default: `false`.
- `log_setting` (String) The LogSetting param.
- `name` (String) The Name param.
- `negate_destination` (Boolean) The NegateDestination param. Default: `false`.
- `negate_source` (Boolean) The NegateSource param. Default: `false`.
- `services` (List of String) The Services param.
- `source_hips` (List of String) The SourceHips param.
- `source_users` (List of String) The SourceUsers param.
- `sources` (List of String) The Sources param.
- `tags` (List of String) The Tags param.
- `timeout` (Number) The Timeout param. Value must be between 1 and 1440.
- `tos` (List of String) The Tos param.
