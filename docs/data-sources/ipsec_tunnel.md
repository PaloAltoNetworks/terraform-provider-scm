---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_ipsec_tunnel Data Source - scm"
subcategory: ""
description: |-
  Retrieves a config item.
---

# scm_ipsec_tunnel (Data Source)

Retrieves a config item.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The Id param.

### Read-Only

- `anti_replay` (Boolean) Enable Anti-Replay check on this tunnel.
- `auto_key` (Attributes) The AutoKey param. (see [below for nested schema](#nestedatt--auto_key))
- `copy_tos` (Boolean) Copy IP TOS bits from inner packet to IPSec packet (not recommended). Default: `false`.
- `enable_gre_encapsulation` (Boolean) allow GRE over IPSec. Default: `false`.
- `name` (String) Alphanumeric string begin with letter: [0-9a-zA-Z._-]. String length must not exceed 63 characters.
- `tfid` (String) The Terraform ID.
- `tunnel_monitor` (Attributes) The TunnelMonitor param. (see [below for nested schema](#nestedatt--tunnel_monitor))

<a id="nestedatt--auto_key"></a>
### Nested Schema for `auto_key`

Read-Only:

- `ike_gateways` (Attributes List) The IkeGateways param. (see [below for nested schema](#nestedatt--auto_key--ike_gateways))
- `ipsec_crypto_profile` (String) The IpsecCryptoProfile param.
- `proxy_id_v6s` (Attributes List) IPv6 type of proxy_id values. (see [below for nested schema](#nestedatt--auto_key--proxy_id_v6s))
- `proxy_ids` (Attributes List) IPv4 type of proxy_id values. (see [below for nested schema](#nestedatt--auto_key--proxy_ids))

<a id="nestedatt--auto_key--ike_gateways"></a>
### Nested Schema for `auto_key.ike_gateways`

Read-Only:

- `name` (String) The Name param.


<a id="nestedatt--auto_key--proxy_id_v6s"></a>
### Nested Schema for `auto_key.proxy_id_v6s`

Read-Only:

- `local` (String) The Local param.
- `name` (String) The Name param.
- `protocol` (Attributes) The Protocol param. (see [below for nested schema](#nestedatt--auto_key--proxy_id_v6s--protocol))
- `remote` (String) The Remote param.

<a id="nestedatt--auto_key--proxy_id_v6s--protocol"></a>
### Nested Schema for `auto_key.proxy_id_v6s.protocol`

Read-Only:

- `number` (Number) IP protocol number. Value must be between 1 and 254.
- `tcp` (Attributes) The Tcp param. (see [below for nested schema](#nestedatt--auto_key--proxy_id_v6s--protocol--tcp))
- `udp` (Attributes) The Udp param. (see [below for nested schema](#nestedatt--auto_key--proxy_id_v6s--protocol--udp))

<a id="nestedatt--auto_key--proxy_id_v6s--protocol--tcp"></a>
### Nested Schema for `auto_key.proxy_id_v6s.protocol.udp`

Read-Only:

- `local_port` (Number) The LocalPort param. Value must be between 0 and 65535. Default: `0`.
- `remote_port` (Number) The RemotePort param. Value must be between 0 and 65535. Default: `0`.


<a id="nestedatt--auto_key--proxy_id_v6s--protocol--udp"></a>
### Nested Schema for `auto_key.proxy_id_v6s.protocol.udp`

Read-Only:

- `local_port` (Number) The LocalPort param. Value must be between 0 and 65535. Default: `0`.
- `remote_port` (Number) The RemotePort param. Value must be between 0 and 65535. Default: `0`.




<a id="nestedatt--auto_key--proxy_ids"></a>
### Nested Schema for `auto_key.proxy_ids`

Read-Only:

- `local` (String) The Local param.
- `name` (String) The Name param.
- `protocol` (Attributes) The Protocol param. (see [below for nested schema](#nestedatt--auto_key--proxy_ids--protocol))
- `remote` (String) The Remote param.

<a id="nestedatt--auto_key--proxy_ids--protocol"></a>
### Nested Schema for `auto_key.proxy_ids.protocol`

Read-Only:

- `number` (Number) IP protocol number. Value must be between 1 and 254.
- `tcp` (Attributes) The Tcp param. (see [below for nested schema](#nestedatt--auto_key--proxy_ids--protocol--tcp))
- `udp` (Attributes) The Udp param. (see [below for nested schema](#nestedatt--auto_key--proxy_ids--protocol--udp))

<a id="nestedatt--auto_key--proxy_ids--protocol--tcp"></a>
### Nested Schema for `auto_key.proxy_ids.protocol.udp`

Read-Only:

- `local_port` (Number) The LocalPort param. Value must be between 0 and 65535. Default: `0`.
- `remote_port` (Number) The RemotePort param. Value must be between 0 and 65535. Default: `0`.


<a id="nestedatt--auto_key--proxy_ids--protocol--udp"></a>
### Nested Schema for `auto_key.proxy_ids.protocol.udp`

Read-Only:

- `local_port` (Number) The LocalPort param. Value must be between 0 and 65535. Default: `0`.
- `remote_port` (Number) The RemotePort param. Value must be between 0 and 65535. Default: `0`.





<a id="nestedatt--tunnel_monitor"></a>
### Nested Schema for `tunnel_monitor`

Read-Only:

- `destination_ip` (String) Destination IP to send ICMP probe.
- `enable` (Boolean) Enable tunnel monitoring on this tunnel. Default: `true`.
- `proxy_id` (String) Which proxy-id (or proxy-id-v6) the monitoring traffic will use.