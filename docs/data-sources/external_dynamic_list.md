---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_external_dynamic_list Data Source - scm"
subcategory: ""
description: |-
  Retrieves a config item.
---

# scm_external_dynamic_list (Data Source)

Retrieves a config item.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The Id param.

### Read-Only

- `name` (String) Alphanumeric string [ 0-9a-zA-Z._-]. String length must not exceed 63 characters.
- `tfid` (String) The Terraform ID.
- `type` (Attributes) The Type param. (see [below for nested schema](#nestedatt--type))

<a id="nestedatt--type"></a>
### Nested Schema for `type`

Read-Only:

- `domain` (Attributes) The Domain param. (see [below for nested schema](#nestedatt--type--domain))
- `imei` (Attributes) The Imei param. (see [below for nested schema](#nestedatt--type--imei))
- `imsi` (Attributes) The Imsi param. (see [below for nested schema](#nestedatt--type--imsi))
- `ip` (Attributes) The Ip param. (see [below for nested schema](#nestedatt--type--ip))
- `predefined_ip` (Attributes) The PredefinedIp param. (see [below for nested schema](#nestedatt--type--predefined_ip))
- `predefined_url` (Attributes) The PredefinedUrl param. (see [below for nested schema](#nestedatt--type--predefined_url))
- `url` (Attributes) The Url param. (see [below for nested schema](#nestedatt--type--url))

<a id="nestedatt--type--domain"></a>
### Nested Schema for `type.domain`

Read-Only:

- `certificate_profile` (String) Profile for authenticating client certificates. Default: `"None"`.
- `description` (String) The Description param. String length must not exceed 255 characters.
- `domain_auth` (Attributes) The DomainAuth param. (see [below for nested schema](#nestedatt--type--domain--domain_auth))
- `exception_list` (List of String) The ExceptionList param.
- `expand_domain` (Boolean) Enable/Disable expand domain. Default: `false`.
- `recurring` (Attributes) The Recurring param. (see [below for nested schema](#nestedatt--type--domain--recurring))
- `url` (String) The Url param. String length must not exceed 255 characters. Default: `"http://"`.

<a id="nestedatt--type--domain--domain_auth"></a>
### Nested Schema for `type.domain.domain_auth`

Read-Only:

- `password` (String) The Password param. String length must not exceed 255 characters.
- `username` (String) The Username param. String length must be between 1 and 255 characters.


<a id="nestedatt--type--domain--recurring"></a>
### Nested Schema for `type.domain.recurring`

Read-Only:

- `daily` (Attributes) The Daily param. (see [below for nested schema](#nestedatt--type--domain--recurring--daily))
- `five_minute` (Boolean) The FiveMinute param.
- `hourly` (Boolean) The Hourly param.
- `monthly` (Attributes) The Monthly param. (see [below for nested schema](#nestedatt--type--domain--recurring--monthly))
- `weekly` (Attributes) The Weekly param. (see [below for nested schema](#nestedatt--type--domain--recurring--weekly))

<a id="nestedatt--type--domain--recurring--daily"></a>
### Nested Schema for `type.domain.recurring.weekly`

Read-Only:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.


<a id="nestedatt--type--domain--recurring--monthly"></a>
### Nested Schema for `type.domain.recurring.weekly`

Read-Only:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.
- `day_of_month` (Number) The DayOfMonth param. Value must be between 1 and 31.


<a id="nestedatt--type--domain--recurring--weekly"></a>
### Nested Schema for `type.domain.recurring.weekly`

Read-Only:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.
- `day_of_week` (String) The DayOfWeek param. String must be one of these: `"sunday"`, `"monday"`, `"tuesday"`, `"wednesday"`, `"thursday"`, `"friday"`, `"saturday"`.




<a id="nestedatt--type--imei"></a>
### Nested Schema for `type.imei`

Read-Only:

- `certificate_profile` (String) Profile for authenticating client certificates. Default: `"None"`.
- `description` (String) The Description param. String length must not exceed 255 characters.
- `exception_list` (List of String) The ExceptionList param.
- `imei_auth` (Attributes) The ImeiAuth param. (see [below for nested schema](#nestedatt--type--imei--imei_auth))
- `recurring` (Attributes) The Recurring param. (see [below for nested schema](#nestedatt--type--imei--recurring))
- `url` (String) The Url param. String length must not exceed 255 characters. Default: `"http://"`.

<a id="nestedatt--type--imei--imei_auth"></a>
### Nested Schema for `type.imei.imei_auth`

Read-Only:

- `password` (String) The Password param. String length must not exceed 255 characters.
- `username` (String) The Username param. String length must be between 1 and 255 characters.


<a id="nestedatt--type--imei--recurring"></a>
### Nested Schema for `type.imei.recurring`

Read-Only:

- `daily` (Attributes) The Daily param. (see [below for nested schema](#nestedatt--type--imei--recurring--daily))
- `five_minute` (Boolean) The FiveMinute param.
- `hourly` (Boolean) The Hourly param.
- `monthly` (Attributes) The Monthly param. (see [below for nested schema](#nestedatt--type--imei--recurring--monthly))
- `weekly` (Attributes) The Weekly param. (see [below for nested schema](#nestedatt--type--imei--recurring--weekly))

<a id="nestedatt--type--imei--recurring--daily"></a>
### Nested Schema for `type.imei.recurring.weekly`

Read-Only:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.


<a id="nestedatt--type--imei--recurring--monthly"></a>
### Nested Schema for `type.imei.recurring.weekly`

Read-Only:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.
- `day_of_month` (Number) The DayOfMonth param. Value must be between 1 and 31.


<a id="nestedatt--type--imei--recurring--weekly"></a>
### Nested Schema for `type.imei.recurring.weekly`

Read-Only:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.
- `day_of_week` (String) The DayOfWeek param. String must be one of these: `"sunday"`, `"monday"`, `"tuesday"`, `"wednesday"`, `"thursday"`, `"friday"`, `"saturday"`.




<a id="nestedatt--type--imsi"></a>
### Nested Schema for `type.imsi`

Read-Only:

- `certificate_profile` (String) Profile for authenticating client certificates. Default: `"None"`.
- `description` (String) The Description param. String length must not exceed 255 characters.
- `exception_list` (List of String) The ExceptionList param.
- `imsi_auth` (Attributes) The ImsiAuth param. (see [below for nested schema](#nestedatt--type--imsi--imsi_auth))
- `recurring` (Attributes) The Recurring param. (see [below for nested schema](#nestedatt--type--imsi--recurring))
- `url` (String) The Url param. String length must not exceed 255 characters. Default: `"http://"`.

<a id="nestedatt--type--imsi--imsi_auth"></a>
### Nested Schema for `type.imsi.imsi_auth`

Read-Only:

- `password` (String) The Password param. String length must not exceed 255 characters.
- `username` (String) The Username param. String length must be between 1 and 255 characters.


<a id="nestedatt--type--imsi--recurring"></a>
### Nested Schema for `type.imsi.recurring`

Read-Only:

- `daily` (Attributes) The Daily param. (see [below for nested schema](#nestedatt--type--imsi--recurring--daily))
- `five_minute` (Boolean) The FiveMinute param.
- `hourly` (Boolean) The Hourly param.
- `monthly` (Attributes) The Monthly param. (see [below for nested schema](#nestedatt--type--imsi--recurring--monthly))
- `weekly` (Attributes) The Weekly param. (see [below for nested schema](#nestedatt--type--imsi--recurring--weekly))

<a id="nestedatt--type--imsi--recurring--daily"></a>
### Nested Schema for `type.imsi.recurring.weekly`

Read-Only:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.


<a id="nestedatt--type--imsi--recurring--monthly"></a>
### Nested Schema for `type.imsi.recurring.weekly`

Read-Only:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.
- `day_of_month` (Number) The DayOfMonth param. Value must be between 1 and 31.


<a id="nestedatt--type--imsi--recurring--weekly"></a>
### Nested Schema for `type.imsi.recurring.weekly`

Read-Only:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.
- `day_of_week` (String) The DayOfWeek param. String must be one of these: `"sunday"`, `"monday"`, `"tuesday"`, `"wednesday"`, `"thursday"`, `"friday"`, `"saturday"`.




<a id="nestedatt--type--ip"></a>
### Nested Schema for `type.ip`

Read-Only:

- `certificate_profile` (String) Profile for authenticating client certificates. Default: `"None"`.
- `description` (String) The Description param. String length must not exceed 255 characters.
- `exception_list` (List of String) The ExceptionList param.
- `ip_auth` (Attributes) The IpAuth param. (see [below for nested schema](#nestedatt--type--ip--ip_auth))
- `recurring` (Attributes) The Recurring param. (see [below for nested schema](#nestedatt--type--ip--recurring))
- `url` (String) The Url param. String length must not exceed 255 characters. Default: `"http://"`.

<a id="nestedatt--type--ip--ip_auth"></a>
### Nested Schema for `type.ip.ip_auth`

Read-Only:

- `password` (String, Sensitive) The Password param. String length must not exceed 255 characters.
- `username` (String) The Username param. String length must be between 1 and 255 characters.


<a id="nestedatt--type--ip--recurring"></a>
### Nested Schema for `type.ip.recurring`

Read-Only:

- `daily` (Attributes) The Daily param. (see [below for nested schema](#nestedatt--type--ip--recurring--daily))
- `five_minute` (Boolean) The FiveMinute param.
- `hourly` (Boolean) The Hourly param.
- `monthly` (Attributes) The Monthly param. (see [below for nested schema](#nestedatt--type--ip--recurring--monthly))
- `weekly` (Attributes) The Weekly param. (see [below for nested schema](#nestedatt--type--ip--recurring--weekly))

<a id="nestedatt--type--ip--recurring--daily"></a>
### Nested Schema for `type.ip.recurring.weekly`

Read-Only:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.


<a id="nestedatt--type--ip--recurring--monthly"></a>
### Nested Schema for `type.ip.recurring.weekly`

Read-Only:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.
- `day_of_month` (Number) The DayOfMonth param. Value must be between 1 and 31.


<a id="nestedatt--type--ip--recurring--weekly"></a>
### Nested Schema for `type.ip.recurring.weekly`

Read-Only:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.
- `day_of_week` (String) The DayOfWeek param. String must be one of these: `"sunday"`, `"monday"`, `"tuesday"`, `"wednesday"`, `"thursday"`, `"friday"`, `"saturday"`.




<a id="nestedatt--type--predefined_ip"></a>
### Nested Schema for `type.predefined_ip`

Read-Only:

- `description` (String) The Description param. String length must not exceed 255 characters.
- `exception_list` (List of String) The ExceptionList param.
- `url` (String) The Url param.


<a id="nestedatt--type--predefined_url"></a>
### Nested Schema for `type.predefined_url`

Read-Only:

- `description` (String) The Description param. String length must not exceed 255 characters.
- `exception_list` (List of String) The ExceptionList param.
- `url` (String) The Url param.


<a id="nestedatt--type--url"></a>
### Nested Schema for `type.url`

Read-Only:

- `certificate_profile` (String) Profile for authenticating client certificates. Default: `"None"`.
- `description` (String) The Description param. String length must not exceed 255 characters.
- `exception_list` (List of String) The ExceptionList param.
- `recurring` (Attributes) The Recurring param. (see [below for nested schema](#nestedatt--type--url--recurring))
- `url` (String) The Url param. String length must not exceed 255 characters. Default: `"http://"`.
- `url_auth` (Attributes) The UrlAuth param. (see [below for nested schema](#nestedatt--type--url--url_auth))

<a id="nestedatt--type--url--recurring"></a>
### Nested Schema for `type.url.recurring`

Read-Only:

- `daily` (Attributes) The Daily param. (see [below for nested schema](#nestedatt--type--url--recurring--daily))
- `five_minute` (Boolean) The FiveMinute param.
- `hourly` (Boolean) The Hourly param.
- `monthly` (Attributes) The Monthly param. (see [below for nested schema](#nestedatt--type--url--recurring--monthly))
- `weekly` (Attributes) The Weekly param. (see [below for nested schema](#nestedatt--type--url--recurring--weekly))

<a id="nestedatt--type--url--recurring--daily"></a>
### Nested Schema for `type.url.recurring.weekly`

Read-Only:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.


<a id="nestedatt--type--url--recurring--monthly"></a>
### Nested Schema for `type.url.recurring.weekly`

Read-Only:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.
- `day_of_month` (Number) The DayOfMonth param. Value must be between 1 and 31.


<a id="nestedatt--type--url--recurring--weekly"></a>
### Nested Schema for `type.url.recurring.weekly`

Read-Only:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.
- `day_of_week` (String) The DayOfWeek param. String must be one of these: `"sunday"`, `"monday"`, `"tuesday"`, `"wednesday"`, `"thursday"`, `"friday"`, `"saturday"`.



<a id="nestedatt--type--url--url_auth"></a>
### Nested Schema for `type.url.url_auth`

Read-Only:

- `password` (String) The Password param. String length must not exceed 255 characters.
- `username` (String) The Username param. String length must be between 1 and 255 characters.