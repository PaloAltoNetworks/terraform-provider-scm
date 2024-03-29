---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_dns_security_profile Data Source - scm"
subcategory: "NetSec"
description: |-
  Retrieves a config item.
---

# scm_dns_security_profile (Data Source)

Retrieves a config item.

## Example Usage

```terraform
data "scm_dns_security_profile" "example" {
  id = "1234-56-789"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The Id param.

### Read-Only

- `botnet_domains` (Attributes) The BotnetDomains param. (see [below for nested schema](#nestedatt--botnet_domains))
- `description` (String) The Description param.
- `name` (String) The Name param.
- `tfid` (String) The Terraform ID.

<a id="nestedatt--botnet_domains"></a>
### Nested Schema for `botnet_domains`

Read-Only:

- `dns_security_categories` (Attributes List) The DnsSecurityCategories param. (see [below for nested schema](#nestedatt--botnet_domains--dns_security_categories))
- `lists` (Attributes List) The Lists param. (see [below for nested schema](#nestedatt--botnet_domains--lists))
- `sinkhole` (Attributes) The Sinkhole param. (see [below for nested schema](#nestedatt--botnet_domains--sinkhole))
- `whitelists` (Attributes List) The Whitelists param. (see [below for nested schema](#nestedatt--botnet_domains--whitelists))

<a id="nestedatt--botnet_domains--dns_security_categories"></a>
### Nested Schema for `botnet_domains.dns_security_categories`

Read-Only:

- `action` (String) The Action param. String must be one of these: `"default"`, `"allow"`, `"block"`, `"sinkhole"`. Default: `"default"`.
- `log_level` (String) The LogLevel param. String must be one of these: `"default"`, `"none"`, `"low"`, `"informational"`, `"medium"`, `"high"`, `"critical"`. Default: `"default"`.
- `name` (String) The Name param.
- `packet_capture` (String) The PacketCapture param. String must be one of these: `"disable"`, `"single-packet"`, `"extended-capture"`.


<a id="nestedatt--botnet_domains--lists"></a>
### Nested Schema for `botnet_domains.lists`

Read-Only:

- `action` (Attributes) The Action param. (see [below for nested schema](#nestedatt--botnet_domains--lists--action))
- `name` (String) The Name param.
- `packet_capture` (String) The PacketCapture param. String must be one of these: `"disable"`, `"single-packet"`, `"extended-capture"`.

<a id="nestedatt--botnet_domains--lists--action"></a>
### Nested Schema for `botnet_domains.lists.action`

Read-Only:

- `alert` (Boolean) The Alert param.
- `allow` (Boolean) The Allow param.
- `block` (Boolean) The Block param.
- `sinkhole` (Boolean) The Sinkhole param.



<a id="nestedatt--botnet_domains--sinkhole"></a>
### Nested Schema for `botnet_domains.sinkhole`

Read-Only:

- `ipv4_address` (String) The Ipv4Address param. String must be one of these: `"127.0.0.1"`, `"pan-sinkhole-default-ip"`.
- `ipv6_address` (String) The Ipv6Address param. String must be one of these: `"::1"`.


<a id="nestedatt--botnet_domains--whitelists"></a>
### Nested Schema for `botnet_domains.whitelists`

Read-Only:

- `description` (String) The Description param.
- `name` (String) The Name param.
