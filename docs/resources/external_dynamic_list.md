---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_external_dynamic_list Resource - scm"
subcategory: ""
description: |-
  Retrieves a config item.
---

# scm_external_dynamic_list (Resource)

Retrieves a config item.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Alphanumeric string [ 0-9a-zA-Z._-]. String length must not exceed 63 characters.
- `type` (Attributes) The Type param. (see [below for nested schema](#nestedatt--type))

### Optional

- `device` (String) The Device param.
- `folder` (String) The Folder param.
- `snippet` (String) The Snippet param.

### Read-Only

- `encrypted_values` (Map of String, Sensitive) (Internal use) Encrypted values returned from the API.
- `id` (String) UUID of the resource.
- `tfid` (String) The Terraform ID.

<a id="nestedatt--type"></a>
### Nested Schema for `type`

Optional:

- `domain` (Attributes) The Domain param. (see [below for nested schema](#nestedatt--type--domain))
- `imei` (Attributes) The Imei param. (see [below for nested schema](#nestedatt--type--imei))
- `imsi` (Attributes) The Imsi param. (see [below for nested schema](#nestedatt--type--imsi))
- `ip` (Attributes) The Ip param. (see [below for nested schema](#nestedatt--type--ip))
- `predefined_ip` (Attributes) The PredefinedIp param. (see [below for nested schema](#nestedatt--type--predefined_ip))
- `predefined_url` (Attributes) The PredefinedUrl param. (see [below for nested schema](#nestedatt--type--predefined_url))
- `url` (Attributes) The Url param. (see [below for nested schema](#nestedatt--type--url))

<a id="nestedatt--type--domain"></a>
### Nested Schema for `type.domain`

Required:

- `recurring` (Attributes) The Recurring param. (see [below for nested schema](#nestedatt--type--domain--recurring))

Optional:

- `certificate_profile` (String) Profile for authenticating client certificates. Default: `"None"`.
- `description` (String) The Description param. String length must not exceed 255 characters.
- `domain_auth` (Attributes) The DomainAuth param. (see [below for nested schema](#nestedatt--type--domain--domain_auth))
- `exception_list` (List of String) The ExceptionList param.
- `expand_domain` (Boolean) Enable/Disable expand domain. Default: `false`.
- `url` (String) The Url param. String length must not exceed 255 characters. Default: `"http://"`.

<a id="nestedatt--type--domain--recurring"></a>
### Nested Schema for `type.domain.recurring`

Optional:

- `daily` (Attributes) The Daily param. (see [below for nested schema](#nestedatt--type--domain--recurring--daily))
- `five_minute` (Boolean) The FiveMinute param.
- `hourly` (Boolean) The Hourly param.
- `monthly` (Attributes) The Monthly param. (see [below for nested schema](#nestedatt--type--domain--recurring--monthly))
- `weekly` (Attributes) The Weekly param. (see [below for nested schema](#nestedatt--type--domain--recurring--weekly))

<a id="nestedatt--type--domain--recurring--daily"></a>
### Nested Schema for `type.domain.recurring.weekly`

Optional:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.


<a id="nestedatt--type--domain--recurring--monthly"></a>
### Nested Schema for `type.domain.recurring.weekly`

Required:

- `day_of_month` (Number) The DayOfMonth param. Value must be between 1 and 31.

Optional:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.


<a id="nestedatt--type--domain--recurring--weekly"></a>
### Nested Schema for `type.domain.recurring.weekly`

Required:

- `day_of_week` (String) The DayOfWeek param. String must be one of these: `"sunday"`, `"monday"`, `"tuesday"`, `"wednesday"`, `"thursday"`, `"friday"`, `"saturday"`.

Optional:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.



<a id="nestedatt--type--domain--domain_auth"></a>
### Nested Schema for `type.domain.domain_auth`

Required:

- `password` (String) The Password param. String length must not exceed 255 characters.
- `username` (String) The Username param. String length must be between 1 and 255 characters.



<a id="nestedatt--type--imei"></a>
### Nested Schema for `type.imei`

Required:

- `recurring` (Attributes) The Recurring param. (see [below for nested schema](#nestedatt--type--imei--recurring))

Optional:

- `certificate_profile` (String) Profile for authenticating client certificates. Default: `"None"`.
- `description` (String) The Description param. String length must not exceed 255 characters.
- `exception_list` (List of String) The ExceptionList param.
- `imei_auth` (Attributes) The ImeiAuth param. (see [below for nested schema](#nestedatt--type--imei--imei_auth))
- `url` (String) The Url param. String length must not exceed 255 characters. Default: `"http://"`.

<a id="nestedatt--type--imei--recurring"></a>
### Nested Schema for `type.imei.recurring`

Optional:

- `daily` (Attributes) The Daily param. (see [below for nested schema](#nestedatt--type--imei--recurring--daily))
- `five_minute` (Boolean) The FiveMinute param.
- `hourly` (Boolean) The Hourly param.
- `monthly` (Attributes) The Monthly param. (see [below for nested schema](#nestedatt--type--imei--recurring--monthly))
- `weekly` (Attributes) The Weekly param. (see [below for nested schema](#nestedatt--type--imei--recurring--weekly))

<a id="nestedatt--type--imei--recurring--daily"></a>
### Nested Schema for `type.imei.recurring.weekly`

Optional:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.


<a id="nestedatt--type--imei--recurring--monthly"></a>
### Nested Schema for `type.imei.recurring.weekly`

Required:

- `day_of_month` (Number) The DayOfMonth param. Value must be between 1 and 31.

Optional:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.


<a id="nestedatt--type--imei--recurring--weekly"></a>
### Nested Schema for `type.imei.recurring.weekly`

Required:

- `day_of_week` (String) The DayOfWeek param. String must be one of these: `"sunday"`, `"monday"`, `"tuesday"`, `"wednesday"`, `"thursday"`, `"friday"`, `"saturday"`.

Optional:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.



<a id="nestedatt--type--imei--imei_auth"></a>
### Nested Schema for `type.imei.imei_auth`

Required:

- `password` (String) The Password param. String length must not exceed 255 characters.
- `username` (String) The Username param. String length must be between 1 and 255 characters.



<a id="nestedatt--type--imsi"></a>
### Nested Schema for `type.imsi`

Required:

- `recurring` (Attributes) The Recurring param. (see [below for nested schema](#nestedatt--type--imsi--recurring))

Optional:

- `certificate_profile` (String) Profile for authenticating client certificates. Default: `"None"`.
- `description` (String) The Description param. String length must not exceed 255 characters.
- `exception_list` (List of String) The ExceptionList param.
- `imsi_auth` (Attributes) The ImsiAuth param. (see [below for nested schema](#nestedatt--type--imsi--imsi_auth))
- `url` (String) The Url param. String length must not exceed 255 characters. Default: `"http://"`.

<a id="nestedatt--type--imsi--recurring"></a>
### Nested Schema for `type.imsi.recurring`

Optional:

- `daily` (Attributes) The Daily param. (see [below for nested schema](#nestedatt--type--imsi--recurring--daily))
- `five_minute` (Boolean) The FiveMinute param.
- `hourly` (Boolean) The Hourly param.
- `monthly` (Attributes) The Monthly param. (see [below for nested schema](#nestedatt--type--imsi--recurring--monthly))
- `weekly` (Attributes) The Weekly param. (see [below for nested schema](#nestedatt--type--imsi--recurring--weekly))

<a id="nestedatt--type--imsi--recurring--daily"></a>
### Nested Schema for `type.imsi.recurring.weekly`

Optional:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.


<a id="nestedatt--type--imsi--recurring--monthly"></a>
### Nested Schema for `type.imsi.recurring.weekly`

Required:

- `day_of_month` (Number) The DayOfMonth param. Value must be between 1 and 31.

Optional:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.


<a id="nestedatt--type--imsi--recurring--weekly"></a>
### Nested Schema for `type.imsi.recurring.weekly`

Required:

- `day_of_week` (String) The DayOfWeek param. String must be one of these: `"sunday"`, `"monday"`, `"tuesday"`, `"wednesday"`, `"thursday"`, `"friday"`, `"saturday"`.

Optional:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.



<a id="nestedatt--type--imsi--imsi_auth"></a>
### Nested Schema for `type.imsi.imsi_auth`

Required:

- `password` (String) The Password param. String length must not exceed 255 characters.
- `username` (String) The Username param. String length must be between 1 and 255 characters.



<a id="nestedatt--type--ip"></a>
### Nested Schema for `type.ip`

Required:

- `recurring` (Attributes) The Recurring param. (see [below for nested schema](#nestedatt--type--ip--recurring))

Optional:

- `certificate_profile` (String) Profile for authenticating client certificates. Default: `"None"`.
- `description` (String) The Description param. String length must not exceed 255 characters.
- `exception_list` (List of String) The ExceptionList param.
- `ip_auth` (Attributes) The IpAuth param. (see [below for nested schema](#nestedatt--type--ip--ip_auth))
- `url` (String) The Url param. String length must not exceed 255 characters. Default: `"http://"`.

<a id="nestedatt--type--ip--recurring"></a>
### Nested Schema for `type.ip.recurring`

Optional:

- `daily` (Attributes) The Daily param. (see [below for nested schema](#nestedatt--type--ip--recurring--daily))
- `five_minute` (Boolean) The FiveMinute param.
- `hourly` (Boolean) The Hourly param.
- `monthly` (Attributes) The Monthly param. (see [below for nested schema](#nestedatt--type--ip--recurring--monthly))
- `weekly` (Attributes) The Weekly param. (see [below for nested schema](#nestedatt--type--ip--recurring--weekly))

<a id="nestedatt--type--ip--recurring--daily"></a>
### Nested Schema for `type.ip.recurring.weekly`

Optional:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.


<a id="nestedatt--type--ip--recurring--monthly"></a>
### Nested Schema for `type.ip.recurring.weekly`

Required:

- `day_of_month` (Number) The DayOfMonth param. Value must be between 1 and 31.

Optional:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.


<a id="nestedatt--type--ip--recurring--weekly"></a>
### Nested Schema for `type.ip.recurring.weekly`

Required:

- `day_of_week` (String) The DayOfWeek param. String must be one of these: `"sunday"`, `"monday"`, `"tuesday"`, `"wednesday"`, `"thursday"`, `"friday"`, `"saturday"`.

Optional:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.



<a id="nestedatt--type--ip--ip_auth"></a>
### Nested Schema for `type.ip.ip_auth`

Required:

- `password` (String, Sensitive) The Password param. String length must not exceed 255 characters.
- `username` (String) The Username param. String length must be between 1 and 255 characters.



<a id="nestedatt--type--predefined_ip"></a>
### Nested Schema for `type.predefined_ip`

Required:

- `url` (String) The Url param.

Optional:

- `description` (String) The Description param. String length must not exceed 255 characters.
- `exception_list` (List of String) The ExceptionList param.


<a id="nestedatt--type--predefined_url"></a>
### Nested Schema for `type.predefined_url`

Required:

- `url` (String) The Url param.

Optional:

- `description` (String) The Description param. String length must not exceed 255 characters.
- `exception_list` (List of String) The ExceptionList param.


<a id="nestedatt--type--url"></a>
### Nested Schema for `type.url`

Required:

- `recurring` (Attributes) The Recurring param. (see [below for nested schema](#nestedatt--type--url--recurring))

Optional:

- `certificate_profile` (String) Profile for authenticating client certificates. Default: `"None"`.
- `description` (String) The Description param. String length must not exceed 255 characters.
- `exception_list` (List of String) The ExceptionList param.
- `url` (String) The Url param. String length must not exceed 255 characters. Default: `"http://"`.
- `url_auth` (Attributes) The UrlAuth param. (see [below for nested schema](#nestedatt--type--url--url_auth))

<a id="nestedatt--type--url--recurring"></a>
### Nested Schema for `type.url.recurring`

Optional:

- `daily` (Attributes) The Daily param. (see [below for nested schema](#nestedatt--type--url--recurring--daily))
- `five_minute` (Boolean) The FiveMinute param.
- `hourly` (Boolean) The Hourly param.
- `monthly` (Attributes) The Monthly param. (see [below for nested schema](#nestedatt--type--url--recurring--monthly))
- `weekly` (Attributes) The Weekly param. (see [below for nested schema](#nestedatt--type--url--recurring--weekly))

<a id="nestedatt--type--url--recurring--daily"></a>
### Nested Schema for `type.url.recurring.weekly`

Optional:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.


<a id="nestedatt--type--url--recurring--monthly"></a>
### Nested Schema for `type.url.recurring.weekly`

Required:

- `day_of_month` (Number) The DayOfMonth param. Value must be between 1 and 31.

Optional:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.


<a id="nestedatt--type--url--recurring--weekly"></a>
### Nested Schema for `type.url.recurring.weekly`

Required:

- `day_of_week` (String) The DayOfWeek param. String must be one of these: `"sunday"`, `"monday"`, `"tuesday"`, `"wednesday"`, `"thursday"`, `"friday"`, `"saturday"`.

Optional:

- `at` (String) Time specification hh (e.g. 20). String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.



<a id="nestedatt--type--url--url_auth"></a>
### Nested Schema for `type.url.url_auth`

Required:

- `password` (String) The Password param. String length must not exceed 255 characters.
- `username` (String) The Username param. String length must be between 1 and 255 characters.