---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_address_group Data Source - scm"
subcategory: ""
description: |-
  Retrieves a config item.
---

# scm_address_group (Data Source)

Retrieves a config item.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The Id param.

### Read-Only

- `description` (String) The Description param. String length must not exceed 1023 characters.
- `dynamic_value` (Attributes) The DynamicValue param. (see [below for nested schema](#nestedatt--dynamic_value))
- `name` (String) Alphanumeric string [ 0-9a-zA-Z._-]. String length must not exceed 63 characters.
- `static_list` (List of String) The StaticList param.
- `tags` (List of String) Tags for address group object. List must contain at most 64 elements.
- `tfid` (String) The Terraform ID.

<a id="nestedatt--dynamic_value"></a>
### Nested Schema for `dynamic_value`

Read-Only:

- `filter` (String) Tag based filter defining group membership e.g. `tag1 AND tag2 OR tag3`. String length must not exceed 2047 characters.