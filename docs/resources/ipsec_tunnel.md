---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_ipsec_tunnel Resource - scm"
subcategory: "NetSec"
description: |-
  Retrieves a config item.
---

# scm_ipsec_tunnel (Resource)

Retrieves a config item.

## Example Usage

```terraform
resource "scm_ipsec_tunnel" "example" {
  # Resource params
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `auto_key` (Attributes) The AutoKey param. (see [below for nested schema](#nestedatt--auto_key))
- `name` (String) Alphanumeric string begin with letter: [0-9a-zA-Z._-]. String length must not exceed 63 characters.

### Optional

- `anti_replay` (Boolean) Enable Anti-Replay check on this tunnel.
- `copy_tos` (Boolean) Copy IP TOS bits from inner packet to IPSec packet (not recommended). Default: `false`.
- `device` (String) The Device param.
- `enable_gre_encapsulation` (Boolean) allow GRE over IPSec. Default: `false`.
- `folder` (String) The Folder param.
- `snippet` (String) The Snippet param.
- `tunnel_monitor` (Attributes) The TunnelMonitor param. (see [below for nested schema](#nestedatt--tunnel_monitor))

### Read-Only

- `id` (String) UUID of the resource.
- `tfid` (String) The Terraform ID.

<a id="nestedatt--auto_key"></a>
### Nested Schema for `auto_key`

Required:

- `ike_gateways` (Attributes List) The IkeGateways param. (see [below for nested schema](#nestedatt--auto_key--ike_gateways))
- `ipsec_crypto_profile` (String) The IpsecCryptoProfile param.

Optional:

- `proxy_id_v6s` (Attributes List) IPv6 type of proxy_id values. (see [below for nested schema](#nestedatt--auto_key--proxy_id_v6s))
- `proxy_ids` (Attributes List) IPv4 type of proxy_id values. (see [below for nested schema](#nestedatt--auto_key--proxy_ids))

<a id="nestedatt--auto_key--ike_gateways"></a>
### Nested Schema for `auto_key.ike_gateways`

Optional:

- `name` (String) The Name param.


<a id="nestedatt--auto_key--proxy_id_v6s"></a>
### Nested Schema for `auto_key.proxy_id_v6s`

Required:

- `name` (String) The Name param.

Optional:

- `local` (String) The Local param.
- `protocol` (Attributes) The Protocol param. (see [below for nested schema](#nestedatt--auto_key--proxy_id_v6s--protocol))
- `remote` (String) The Remote param.

<a id="nestedatt--auto_key--proxy_id_v6s--protocol"></a>
### Nested Schema for `auto_key.proxy_id_v6s.protocol`

Optional:

- `number` (Number) IP protocol number. Value must be between 1 and 254. Ensure that only one of the following is specified: `number`, `tcp`, `udp`
- `tcp` (Attributes) The Tcp param. Ensure that only one of the following is specified: `number`, `tcp`, `udp` (see [below for nested schema](#nestedatt--auto_key--proxy_id_v6s--protocol--tcp))
- `udp` (Attributes) The Udp param. Ensure that only one of the following is specified: `number`, `tcp`, `udp` (see [below for nested schema](#nestedatt--auto_key--proxy_id_v6s--protocol--udp))

<a id="nestedatt--auto_key--proxy_id_v6s--protocol--tcp"></a>
### Nested Schema for `auto_key.proxy_id_v6s.protocol.udp`

Optional:

- `local_port` (Number) The LocalPort param. Value must be between 0 and 65535. Default: `0`.
- `remote_port` (Number) The RemotePort param. Value must be between 0 and 65535. Default: `0`.


<a id="nestedatt--auto_key--proxy_id_v6s--protocol--udp"></a>
### Nested Schema for `auto_key.proxy_id_v6s.protocol.udp`

Optional:

- `local_port` (Number) The LocalPort param. Value must be between 0 and 65535. Default: `0`.
- `remote_port` (Number) The RemotePort param. Value must be between 0 and 65535. Default: `0`.




<a id="nestedatt--auto_key--proxy_ids"></a>
### Nested Schema for `auto_key.proxy_ids`

Required:

- `name` (String) The Name param.

Optional:

- `local` (String) The Local param.
- `protocol` (Attributes) The Protocol param. (see [below for nested schema](#nestedatt--auto_key--proxy_ids--protocol))
- `remote` (String) The Remote param.

<a id="nestedatt--auto_key--proxy_ids--protocol"></a>
### Nested Schema for `auto_key.proxy_ids.protocol`

Optional:

- `number` (Number) IP protocol number. Value must be between 1 and 254. Ensure that only one of the following is specified: `number`, `tcp`, `udp`
- `tcp` (Attributes) The Tcp param. Ensure that only one of the following is specified: `number`, `tcp`, `udp` (see [below for nested schema](#nestedatt--auto_key--proxy_ids--protocol--tcp))
- `udp` (Attributes) The Udp param. Ensure that only one of the following is specified: `number`, `tcp`, `udp` (see [below for nested schema](#nestedatt--auto_key--proxy_ids--protocol--udp))

<a id="nestedatt--auto_key--proxy_ids--protocol--tcp"></a>
### Nested Schema for `auto_key.proxy_ids.protocol.udp`

Optional:

- `local_port` (Number) The LocalPort param. Value must be between 0 and 65535. Default: `0`.
- `remote_port` (Number) The RemotePort param. Value must be between 0 and 65535. Default: `0`.


<a id="nestedatt--auto_key--proxy_ids--protocol--udp"></a>
### Nested Schema for `auto_key.proxy_ids.protocol.udp`

Optional:

- `local_port` (Number) The LocalPort param. Value must be between 0 and 65535. Default: `0`.
- `remote_port` (Number) The RemotePort param. Value must be between 0 and 65535. Default: `0`.





<a id="nestedatt--tunnel_monitor"></a>
### Nested Schema for `tunnel_monitor`

Required:

- `destination_ip` (String) Destination IP to send ICMP probe.

Optional:

- `enable` (Boolean) Enable tunnel monitoring on this tunnel. Default: `true`.
- `proxy_id` (String) Which proxy-id (or proxy-id-v6) the monitoring traffic will use.
