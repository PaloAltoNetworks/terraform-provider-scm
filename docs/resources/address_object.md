---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_address_object Resource - scm"
subcategory: "NetSec"
description: |-
  Retrieves a config item.
---

# scm_address_object (Resource)

Retrieves a config item.

## Example Usage

```terraform
resource "scm_address_object" "example" {
  folder      = "Shared"
  name        = "example"
  description = "Made by Terraform"
  ip_netmask  = "10.2.3.4"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Alphanumeric string [ 0-9a-zA-Z._-]. String length must not exceed 63 characters.

### Optional

- `description` (String) The Description param. String length must not exceed 1023 characters.
- `device` (String) The Device param.
- `folder` (String) The Folder param.
- `fqdn` (String) The Fqdn param. String length must be between 1 and 255 characters. String validation regex: `^[a-zA-Z0-9_]([a-zA-Z0-9._-])+[a-zA-Z0-9]$`. Ensure that only one of the following is specified: `fqdn`, `ip_netmask`, `ip_range`, `ip_wildcard`
- `ip_netmask` (String) The IpNetmask param. Ensure that only one of the following is specified: `fqdn`, `ip_netmask`, `ip_range`, `ip_wildcard`
- `ip_range` (String) The IpRange param. Ensure that only one of the following is specified: `fqdn`, `ip_netmask`, `ip_range`, `ip_wildcard`
- `ip_wildcard` (String) The IpWildcard param. Ensure that only one of the following is specified: `fqdn`, `ip_netmask`, `ip_range`, `ip_wildcard`
- `snippet` (String) The Snippet param.
- `tags` (List of String) Tags for address object. List must contain at most 64 elements. Individual elements in this list are subject to additional validation. String length must not exceed 127 characters.

### Read-Only

- `id` (String) UUID of the resource.
- `tfid` (String) The Terraform ID.
- `type` (String) The Type param.
