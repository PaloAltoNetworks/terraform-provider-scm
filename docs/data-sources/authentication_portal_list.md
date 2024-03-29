---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_authentication_portal_list Data Source - scm"
subcategory: "NetSec"
description: |-
  Retrieves a listing of config items.
---

# scm_authentication_portal_list (Data Source)

Retrieves a listing of config items.

## Example Usage

```terraform
data "scm_authentication_portal_list" "example" {
  folder = "Shared"
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
- `snippet` (String) The Snippet param.

### Read-Only

- `data` (Attributes List) The Data param. (see [below for nested schema](#nestedatt--data))
- `tfid` (String) The Terraform ID.
- `total` (Number) The Total param.

<a id="nestedatt--data"></a>
### Nested Schema for `data`

Read-Only:

- `authentication_profile` (String) The AuthenticationProfile param.
- `certificate_profile` (String) The CertificateProfile param.
- `gp_udp_port` (Number) The GpUdpPort param. Value must be between 1 and 65535.
- `id` (String) UUID of the resource.
- `idle_timer` (Number) The IdleTimer param. Value must be between 1 and 1440.
- `redirect_host` (String) The RedirectHost param.
- `timer` (Number) The Timer param. Value must be between 1 and 1440.
- `tls_service_profile` (String) The TlsServiceProfile param.
