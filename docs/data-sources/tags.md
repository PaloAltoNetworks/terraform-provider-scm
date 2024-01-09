---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_tags Data Source - scm"
subcategory: ""
description: |-
  Retrieves a config item.
---

# scm_tags (Data Source)

Retrieves a config item.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The Id param.

### Read-Only

- `color` (String) The Color param. String must be one of these: `"Red"`, `"Green"`, `"Blue"`, `"Yellow"`, `"Copper"`, `"Orange"`, `"Purple"`, `"Gray"`, `"Light Green"`, `"Cyan"`, `"Light Gray"`, `"Blue Gray"`, `"Lime"`, `"Black"`, `"Gold"`, `"Brown"`, `"Olive"`, `"Maroon"`, `"Red-Orange"`, `"Yellow-Orange"`, `"Forest Green"`, `"Turquoise Blue"`, `"Azure Blue"`, `"Cerulean Blue"`, `"Midnight Blue"`, `"Medium Blue"`, `"Cobalt Blue"`, `"Violet Blue"`, `"Blue Violet"`, `"Medium Violet"`, `"Medium Rose"`, `"Lavender"`, `"Orchid"`, `"Thistle"`, `"Peach"`, `"Salmon"`, `"Magenta"`, `"Red Violet"`, `"Mahogany"`, `"Burnt Sienna"`, `"Chestnut"`.
- `comments` (String) The Comments param. String length must not exceed 1023 characters.
- `name` (String) The Name param. String length must not exceed 127 characters.
- `tfid` (String) The Terraform ID.