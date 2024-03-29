---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_security_rule_list Data Source - scm"
subcategory: "NetSec"
description: |-
  Retrieves a listing of config items.
---

# scm_security_rule_list (Data Source)

Retrieves a listing of config items.

## Example Usage

```terraform
data "scm_security_rule_list" "example" {
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

- `action` (String) The action to be taken when the rule is matched. String must be one of these: `"allow"`, `"deny"`, `"drop"`, `"reset-client"`, `"reset-server"`, `"reset-both"`.
- `applications` (List of String) The application(s) being accessed.
- `categories` (List of String) The URL categories being accessed.
- `description` (String) The description of the security rule.
- `destination_hips` (List of String) The destination Host Integrity Profile(s).
- `destinations` (List of String) The destination address(es).
- `disabled` (Boolean) The state of the security rule. Default: `false`.
- `froms` (List of String) The source security zone(s).
- `id` (String) The UUID of the security rule.
- `log_setting` (String) The external log forwarding profile.
- `name` (String) The name of the security rule.
- `negate_destination` (Boolean) Negate the destination addresses(es). Default: `false`.
- `negate_source` (Boolean) Negate the source address(es). Default: `false`.
- `profile_setting` (Attributes) The security profile object. (see [below for nested schema](#nestedatt--data--profile_setting))
- `services` (List of String) The service(s) being accessed.
- `source_hips` (List of String) The source Host Integrity Profile(s).
- `source_users` (List of String) The source user(s) or group(s).
- `sources` (List of String) The source address(es).
- `tags` (List of String) The tags associated with the security rule.
- `tos` (List of String) The destination security zone(s).

<a id="nestedatt--data--profile_setting"></a>
### Nested Schema for `data.profile_setting`

Read-Only:

- `group` (List of String) The security profile group.
