---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_local_user Data Source - scm"
subcategory: "NetSec"
description: |-
  Retrieves a config item.
---

# scm_local_user (Data Source)

Retrieves a config item.

## Example Usage

```terraform
data "scm_local_user" "example" {
  id = "1234-56-789"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The Id param.

### Read-Only

- `disabled` (Boolean) The Disabled param.
- `name` (String) The Name param. String length must not exceed 31 characters.
- `password` (String, Sensitive) The Password param. String length must not exceed 63 characters.
- `tfid` (String) The Terraform ID.
