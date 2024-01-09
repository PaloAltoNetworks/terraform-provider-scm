---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_app_override_rule Data Source - scm"
subcategory: ""
description: |-
  Retrieves a config item.
---

# scm_app_override_rule (Data Source)

Retrieves a config item.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The Id param.

### Read-Only

- `application` (String) The Application param.
- `description` (String) The Description param. String length must not exceed 1024 characters.
- `destinations` (List of String) The Destinations param.
- `disabled` (Boolean) The Disabled param. Default: `false`.
- `froms` (List of String) The Froms param.
- `group_tag` (String) The GroupTag param.
- `name` (String) The Name param. String length must not exceed 63 characters. String validation regex: `^[a-zA-Z0-9._-]+$`.
- `negate_destination` (Boolean) The NegateDestination param. Default: `false`.
- `negate_source` (Boolean) The NegateSource param. Default: `false`.
- `port` (Number) The Port param. Value must be between 0 and 65535.
- `protocol` (String) The Protocol param. String must be one of these: `"tcp"`, `"udp"`.
- `sources` (List of String) The Sources param.
- `tags` (List of String) The Tags param.
- `tfid` (String) The Terraform ID.
- `tos` (List of String) The Tos param.