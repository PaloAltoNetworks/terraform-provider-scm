---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_hip_object_list Data Source - scm"
subcategory: "NetSec"
description: |-
  Retrieves a listing of config items.
---

# scm_hip_object_list (Data Source)

Retrieves a listing of config items.

## Example Usage

```terraform
data "scm_hip_object_list" "example" {
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

- `anti_malware` (Attributes) The AntiMalware param. (see [below for nested schema](#nestedatt--data--anti_malware))
- `certificate` (Attributes) The Certificate param. (see [below for nested schema](#nestedatt--data--certificate))
- `custom_checks` (Attributes) The CustomChecks param. (see [below for nested schema](#nestedatt--data--custom_checks))
- `data_loss_prevention` (Attributes) The DataLossPrevention param. (see [below for nested schema](#nestedatt--data--data_loss_prevention))
- `description` (String) The Description param. String length must not exceed 255 characters.
- `disk_backup` (Attributes) The DiskBackup param. (see [below for nested schema](#nestedatt--data--disk_backup))
- `disk_encryption` (Attributes) The DiskEncryption param. (see [below for nested schema](#nestedatt--data--disk_encryption))
- `firewall` (Attributes) The Firewall param. (see [below for nested schema](#nestedatt--data--firewall))
- `host_info` (Attributes) The HostInfo param. (see [below for nested schema](#nestedatt--data--host_info))
- `id` (String) UUID of the resource.
- `mobile_device` (Attributes) The MobileDevice param. (see [below for nested schema](#nestedatt--data--mobile_device))
- `name` (String) Alphanumeric string [ 0-9a-zA-Z._-]. String length must not exceed 31 characters.
- `network_info` (Attributes) The NetworkInfo param. (see [below for nested schema](#nestedatt--data--network_info))
- `patch_management` (Attributes) The PatchManagement param. (see [below for nested schema](#nestedatt--data--patch_management))

<a id="nestedatt--data--anti_malware"></a>
### Nested Schema for `data.anti_malware`

Read-Only:

- `criteria` (Attributes) The Criteria param. (see [below for nested schema](#nestedatt--data--anti_malware--criteria))
- `exclude_vendor` (Boolean) The ExcludeVendor param. Default: `false`.
- `vendors` (Attributes List) Vendor name. (see [below for nested schema](#nestedatt--data--anti_malware--vendors))

<a id="nestedatt--data--anti_malware--criteria"></a>
### Nested Schema for `data.anti_malware.criteria`

Read-Only:

- `is_installed` (Boolean) Is Installed. Default: `true`.
- `last_scan_time` (Attributes) The LastScanTime param. (see [below for nested schema](#nestedatt--data--anti_malware--criteria--last_scan_time))
- `product_version` (Attributes) The ProductVersion param. (see [below for nested schema](#nestedatt--data--anti_malware--criteria--product_version))
- `real_time_protection` (String) real time protection. String must be one of these: `"no"`, `"yes"`, `"not-available"`.
- `virdef_version` (Attributes) The VirdefVersion param. (see [below for nested schema](#nestedatt--data--anti_malware--criteria--virdef_version))

<a id="nestedatt--data--anti_malware--criteria--last_scan_time"></a>
### Nested Schema for `data.anti_malware.criteria.virdef_version`

Read-Only:

- `not_available` (Boolean) The NotAvailable param.
- `not_within` (Attributes) The NotWithin param. (see [below for nested schema](#nestedatt--data--anti_malware--criteria--virdef_version--not_within))
- `within` (Attributes) The Within param. (see [below for nested schema](#nestedatt--data--anti_malware--criteria--virdef_version--within))

<a id="nestedatt--data--anti_malware--criteria--virdef_version--not_within"></a>
### Nested Schema for `data.anti_malware.criteria.virdef_version.not_within`

Read-Only:

- `days` (Number) specify time in days. Value must be between 1 and 65535.
- `hours` (Number) specify time in hours. Value must be between 1 and 65535.


<a id="nestedatt--data--anti_malware--criteria--virdef_version--within"></a>
### Nested Schema for `data.anti_malware.criteria.virdef_version.within`

Read-Only:

- `days` (Number) specify time in days. Value must be between 1 and 65535.
- `hours` (Number) specify time in hours. Value must be between 1 and 65535.



<a id="nestedatt--data--anti_malware--criteria--product_version"></a>
### Nested Schema for `data.anti_malware.criteria.virdef_version`

Read-Only:

- `contains` (String) The Contains param. String length must not exceed 255 characters.
- `greater_equal` (String) The GreaterEqual param. String length must not exceed 255 characters.
- `greater_than` (String) The GreaterThan param. String length must not exceed 255 characters.
- `is` (String) The Is param. String length must not exceed 255 characters.
- `is_not` (String) The IsNot param. String length must not exceed 255 characters.
- `less_equal` (String) The LessEqual param. String length must not exceed 255 characters.
- `less_than` (String) The LessThan param. String length must not exceed 255 characters.
- `not_within` (Attributes) The NotWithin param. (see [below for nested schema](#nestedatt--data--anti_malware--criteria--virdef_version--not_within))
- `within` (Attributes) The Within param. (see [below for nested schema](#nestedatt--data--anti_malware--criteria--virdef_version--within))

<a id="nestedatt--data--anti_malware--criteria--virdef_version--not_within"></a>
### Nested Schema for `data.anti_malware.criteria.virdef_version.not_within`

Read-Only:

- `versions` (Number) versions range. Value must be between 1 and 65535. Default: `1`.


<a id="nestedatt--data--anti_malware--criteria--virdef_version--within"></a>
### Nested Schema for `data.anti_malware.criteria.virdef_version.within`

Read-Only:

- `versions` (Number) versions range. Value must be between 1 and 65535. Default: `1`.



<a id="nestedatt--data--anti_malware--criteria--virdef_version"></a>
### Nested Schema for `data.anti_malware.criteria.virdef_version`

Read-Only:

- `not_within` (Attributes) The NotWithin param. (see [below for nested schema](#nestedatt--data--anti_malware--criteria--virdef_version--not_within))
- `within` (Attributes) The Within param. (see [below for nested schema](#nestedatt--data--anti_malware--criteria--virdef_version--within))

<a id="nestedatt--data--anti_malware--criteria--virdef_version--not_within"></a>
### Nested Schema for `data.anti_malware.criteria.virdef_version.not_within`

Read-Only:

- `days` (Number) specify time in days. Value must be between 1 and 65535.
- `versions` (Number) specify versions range. Value must be between 1 and 65535.


<a id="nestedatt--data--anti_malware--criteria--virdef_version--within"></a>
### Nested Schema for `data.anti_malware.criteria.virdef_version.within`

Read-Only:

- `days` (Number) specify time in days. Value must be between 1 and 65535.
- `versions` (Number) specify versions range. Value must be between 1 and 65535.




<a id="nestedatt--data--anti_malware--vendors"></a>
### Nested Schema for `data.anti_malware.vendors`

Read-Only:

- `name` (String) The Name param. String length must not exceed 103 characters.
- `products` (List of String) The Products param. Individual elements in this list are subject to additional validation. String length must not exceed 1023 characters. String validation regex: `.*`.



<a id="nestedatt--data--certificate"></a>
### Nested Schema for `data.certificate`

Read-Only:

- `criteria` (Attributes) The Criteria param. (see [below for nested schema](#nestedatt--data--certificate--criteria))

<a id="nestedatt--data--certificate--criteria"></a>
### Nested Schema for `data.certificate.criteria`

Read-Only:

- `certificate_attributes` (Attributes List) The CertificateAttributes param. (see [below for nested schema](#nestedatt--data--certificate--criteria--certificate_attributes))
- `certificate_profile` (String) Profile for authenticating client certificates.

<a id="nestedatt--data--certificate--criteria--certificate_attributes"></a>
### Nested Schema for `data.certificate.criteria.certificate_profile`

Read-Only:

- `name` (String) Attribute Name.
- `value` (String) Key value. String length must not exceed 1024 characters. String validation regex: `.*`.




<a id="nestedatt--data--custom_checks"></a>
### Nested Schema for `data.custom_checks`

Read-Only:

- `criteria` (Attributes) The Criteria param. (see [below for nested schema](#nestedatt--data--custom_checks--criteria))

<a id="nestedatt--data--custom_checks--criteria"></a>
### Nested Schema for `data.custom_checks.criteria`

Read-Only:

- `plist` (Attributes List) The Plist param. (see [below for nested schema](#nestedatt--data--custom_checks--criteria--plist))
- `process_list` (Attributes List) The ProcessList param. (see [below for nested schema](#nestedatt--data--custom_checks--criteria--process_list))
- `registry_keys` (Attributes List) The RegistryKeys param. (see [below for nested schema](#nestedatt--data--custom_checks--criteria--registry_keys))

<a id="nestedatt--data--custom_checks--criteria--plist"></a>
### Nested Schema for `data.custom_checks.criteria.registry_keys`

Read-Only:

- `keys` (Attributes List) The Keys param. (see [below for nested schema](#nestedatt--data--custom_checks--criteria--registry_keys--keys))
- `name` (String) Preference list. String length must not exceed 1023 characters.
- `negate` (Boolean) Plist does not exist. Default: `false`.

<a id="nestedatt--data--custom_checks--criteria--registry_keys--keys"></a>
### Nested Schema for `data.custom_checks.criteria.registry_keys.keys`

Read-Only:

- `name` (String) Key name. String length must not exceed 1023 characters.
- `negate` (Boolean) Value does not exist or match specified value data. Default: `false`.
- `value` (String) Key value. String length must not exceed 1024 characters. String validation regex: `.*`.



<a id="nestedatt--data--custom_checks--criteria--process_list"></a>
### Nested Schema for `data.custom_checks.criteria.registry_keys`

Read-Only:

- `name` (String) Process Name. String length must not exceed 1023 characters.
- `running` (Boolean) The Running param. Default: `true`.


<a id="nestedatt--data--custom_checks--criteria--registry_keys"></a>
### Nested Schema for `data.custom_checks.criteria.registry_keys`

Read-Only:

- `default_value_data` (String) Registry key default value data. String length must not exceed 1024 characters. String validation regex: `.*`.
- `name` (String) Registry key. String length must not exceed 1023 characters.
- `negate` (Boolean) Key does not exist or match specified value data. Default: `false`.
- `registry_values` (Attributes List) The RegistryValues param. (see [below for nested schema](#nestedatt--data--custom_checks--criteria--registry_keys--registry_values))

<a id="nestedatt--data--custom_checks--criteria--registry_keys--registry_values"></a>
### Nested Schema for `data.custom_checks.criteria.registry_keys.registry_values`

Read-Only:

- `name` (String) Registry value name. String length must not exceed 1023 characters.
- `negate` (Boolean) Value does not exist or match specified value data. Default: `false`.
- `value_data` (String) Registry value data. String length must not exceed 1024 characters. String validation regex: `.*`.





<a id="nestedatt--data--data_loss_prevention"></a>
### Nested Schema for `data.data_loss_prevention`

Read-Only:

- `criteria` (Attributes) The Criteria param. (see [below for nested schema](#nestedatt--data--data_loss_prevention--criteria))
- `exclude_vendor` (Boolean) The ExcludeVendor param. Default: `false`.
- `vendors` (Attributes List) Vendor name. (see [below for nested schema](#nestedatt--data--data_loss_prevention--vendors))

<a id="nestedatt--data--data_loss_prevention--criteria"></a>
### Nested Schema for `data.data_loss_prevention.criteria`

Read-Only:

- `is_enabled` (String) is enabled. String must be one of these: `"no"`, `"yes"`, `"not-available"`.
- `is_installed` (Boolean) Is Installed. Default: `true`.


<a id="nestedatt--data--data_loss_prevention--vendors"></a>
### Nested Schema for `data.data_loss_prevention.vendors`

Read-Only:

- `name` (String) The Name param. String length must not exceed 103 characters.
- `products` (List of String) Product name. Individual elements in this list are subject to additional validation. String length must not exceed 1023 characters. String validation regex: `.*`.



<a id="nestedatt--data--disk_backup"></a>
### Nested Schema for `data.disk_backup`

Read-Only:

- `criteria` (Attributes) The Criteria param. (see [below for nested schema](#nestedatt--data--disk_backup--criteria))
- `exclude_vendor` (Boolean) The ExcludeVendor param. Default: `false`.
- `vendors` (Attributes List) Vendor name. (see [below for nested schema](#nestedatt--data--disk_backup--vendors))

<a id="nestedatt--data--disk_backup--criteria"></a>
### Nested Schema for `data.disk_backup.criteria`

Read-Only:

- `is_installed` (Boolean) Is Installed. Default: `true`.
- `last_backup_time` (Attributes) The LastBackupTime param. (see [below for nested schema](#nestedatt--data--disk_backup--criteria--last_backup_time))

<a id="nestedatt--data--disk_backup--criteria--last_backup_time"></a>
### Nested Schema for `data.disk_backup.criteria.last_backup_time`

Read-Only:

- `not_available` (Boolean) The NotAvailable param.
- `not_within` (Attributes) The NotWithin param. (see [below for nested schema](#nestedatt--data--disk_backup--criteria--last_backup_time--not_within))
- `within` (Attributes) The Within param. (see [below for nested schema](#nestedatt--data--disk_backup--criteria--last_backup_time--within))

<a id="nestedatt--data--disk_backup--criteria--last_backup_time--not_within"></a>
### Nested Schema for `data.disk_backup.criteria.last_backup_time.not_within`

Read-Only:

- `days` (Number) specify time in days. Value must be between 1 and 65535.
- `hours` (Number) specify time in hours. Value must be between 1 and 65535.


<a id="nestedatt--data--disk_backup--criteria--last_backup_time--within"></a>
### Nested Schema for `data.disk_backup.criteria.last_backup_time.within`

Read-Only:

- `days` (Number) specify time in days. Value must be between 1 and 65535.
- `hours` (Number) specify time in hours. Value must be between 1 and 65535.




<a id="nestedatt--data--disk_backup--vendors"></a>
### Nested Schema for `data.disk_backup.vendors`

Read-Only:

- `name` (String) The Name param. String length must not exceed 103 characters.
- `products` (List of String) The Products param. Individual elements in this list are subject to additional validation. String length must not exceed 1023 characters. String validation regex: `.*`.



<a id="nestedatt--data--disk_encryption"></a>
### Nested Schema for `data.disk_encryption`

Read-Only:

- `criteria` (Attributes) Encryption locations. (see [below for nested schema](#nestedatt--data--disk_encryption--criteria))
- `exclude_vendor` (Boolean) The ExcludeVendor param. Default: `false`.
- `vendors` (Attributes List) Vendor name. (see [below for nested schema](#nestedatt--data--disk_encryption--vendors))

<a id="nestedatt--data--disk_encryption--criteria"></a>
### Nested Schema for `data.disk_encryption.criteria`

Read-Only:

- `encrypted_locations` (Attributes List) The EncryptedLocations param. (see [below for nested schema](#nestedatt--data--disk_encryption--criteria--encrypted_locations))
- `is_installed` (Boolean) Is Installed. Default: `true`.

<a id="nestedatt--data--disk_encryption--criteria--encrypted_locations"></a>
### Nested Schema for `data.disk_encryption.criteria.is_installed`

Read-Only:

- `encryption_state` (Attributes) The EncryptionState param. (see [below for nested schema](#nestedatt--data--disk_encryption--criteria--is_installed--encryption_state))
- `name` (String) Encryption location. String length must not exceed 1023 characters.

<a id="nestedatt--data--disk_encryption--criteria--is_installed--encryption_state"></a>
### Nested Schema for `data.disk_encryption.criteria.is_installed.encryption_state`

Read-Only:

- `is` (String) The Is param. String must be one of these: `"encrypted"`, `"unencrypted"`, `"partial"`, `"unknown"`.
- `is_not` (String) The IsNot param. String must be one of these: `"encrypted"`, `"unencrypted"`, `"partial"`, `"unknown"`.




<a id="nestedatt--data--disk_encryption--vendors"></a>
### Nested Schema for `data.disk_encryption.vendors`

Read-Only:

- `name` (String) The Name param. String length must not exceed 103 characters.
- `products` (List of String) The Products param. Individual elements in this list are subject to additional validation. String length must not exceed 1023 characters. String validation regex: `.*`.



<a id="nestedatt--data--firewall"></a>
### Nested Schema for `data.firewall`

Read-Only:

- `criteria` (Attributes) The Criteria param. (see [below for nested schema](#nestedatt--data--firewall--criteria))
- `exclude_vendor` (Boolean) The ExcludeVendor param. Default: `false`.
- `vendors` (Attributes List) Vendor name. (see [below for nested schema](#nestedatt--data--firewall--vendors))

<a id="nestedatt--data--firewall--criteria"></a>
### Nested Schema for `data.firewall.criteria`

Read-Only:

- `is_enabled` (String) is enabled. String must be one of these: `"no"`, `"yes"`, `"not-available"`.
- `is_installed` (Boolean) Is Installed. Default: `true`.


<a id="nestedatt--data--firewall--vendors"></a>
### Nested Schema for `data.firewall.vendors`

Read-Only:

- `name` (String) The Name param. String length must not exceed 103 characters.
- `products` (List of String) The Products param. Individual elements in this list are subject to additional validation. String length must not exceed 1023 characters. String validation regex: `.*`.



<a id="nestedatt--data--host_info"></a>
### Nested Schema for `data.host_info`

Read-Only:

- `criteria` (Attributes) The Criteria param. (see [below for nested schema](#nestedatt--data--host_info--criteria))

<a id="nestedatt--data--host_info--criteria"></a>
### Nested Schema for `data.host_info.criteria`

Read-Only:

- `client_version` (Attributes) The ClientVersion param. (see [below for nested schema](#nestedatt--data--host_info--criteria--client_version))
- `domain` (Attributes) The Domain param. (see [below for nested schema](#nestedatt--data--host_info--criteria--domain))
- `host_id` (Attributes) The HostId param. (see [below for nested schema](#nestedatt--data--host_info--criteria--host_id))
- `host_name` (Attributes) The HostName param. (see [below for nested schema](#nestedatt--data--host_info--criteria--host_name))
- `managed` (Boolean) If device is managed.
- `os` (Attributes) The Os param. (see [below for nested schema](#nestedatt--data--host_info--criteria--os))
- `serial_number` (Attributes) The SerialNumber param. (see [below for nested schema](#nestedatt--data--host_info--criteria--serial_number))

<a id="nestedatt--data--host_info--criteria--client_version"></a>
### Nested Schema for `data.host_info.criteria.serial_number`

Read-Only:

- `contains` (String) The Contains param. String length must not exceed 255 characters.
- `is` (String) The Is param. String length must not exceed 255 characters.
- `is_not` (String) The IsNot param. String length must not exceed 255 characters.


<a id="nestedatt--data--host_info--criteria--domain"></a>
### Nested Schema for `data.host_info.criteria.serial_number`

Read-Only:

- `contains` (String) The Contains param. String length must not exceed 255 characters.
- `is` (String) The Is param. String length must not exceed 255 characters.
- `is_not` (String) The IsNot param. String length must not exceed 255 characters.


<a id="nestedatt--data--host_info--criteria--host_id"></a>
### Nested Schema for `data.host_info.criteria.serial_number`

Read-Only:

- `contains` (String) The Contains param. String length must not exceed 255 characters.
- `is` (String) The Is param. String length must not exceed 255 characters.
- `is_not` (String) The IsNot param. String length must not exceed 255 characters.


<a id="nestedatt--data--host_info--criteria--host_name"></a>
### Nested Schema for `data.host_info.criteria.serial_number`

Read-Only:

- `contains` (String) The Contains param. String length must not exceed 255 characters.
- `is` (String) The Is param. String length must not exceed 255 characters.
- `is_not` (String) The IsNot param. String length must not exceed 255 characters.


<a id="nestedatt--data--host_info--criteria--os"></a>
### Nested Schema for `data.host_info.criteria.serial_number`

Read-Only:

- `contains` (Attributes) The Contains param. (see [below for nested schema](#nestedatt--data--host_info--criteria--serial_number--contains))

<a id="nestedatt--data--host_info--criteria--serial_number--contains"></a>
### Nested Schema for `data.host_info.criteria.serial_number.contains`

Read-Only:

- `apple` (String) Apple vendor. String length must not exceed 255 characters.
- `google` (String) Google vendor. String length must not exceed 255 characters.
- `linux` (String) Linux vendor. String length must not exceed 255 characters.
- `microsoft` (String) Microsoft vendor. String length must not exceed 255 characters.
- `other` (String) Other vendor. String length must not exceed 255 characters.



<a id="nestedatt--data--host_info--criteria--serial_number"></a>
### Nested Schema for `data.host_info.criteria.serial_number`

Read-Only:

- `contains` (String) The Contains param. String length must not exceed 255 characters.
- `is` (String) The Is param. String length must not exceed 255 characters.
- `is_not` (String) The IsNot param. String length must not exceed 255 characters.




<a id="nestedatt--data--mobile_device"></a>
### Nested Schema for `data.mobile_device`

Read-Only:

- `criteria` (Attributes) The Criteria param. (see [below for nested schema](#nestedatt--data--mobile_device--criteria))

<a id="nestedatt--data--mobile_device--criteria"></a>
### Nested Schema for `data.mobile_device.criteria`

Read-Only:

- `applications` (Attributes) The Applications param. (see [below for nested schema](#nestedatt--data--mobile_device--criteria--applications))
- `disk_encrypted` (Boolean) If device's disk is encrypted.
- `imei` (Attributes) The Imei param. (see [below for nested schema](#nestedatt--data--mobile_device--criteria--imei))
- `jailbroken` (Boolean) If device is by rooted/jailbroken.
- `last_checkin_time` (Attributes) The LastCheckinTime param. (see [below for nested schema](#nestedatt--data--mobile_device--criteria--last_checkin_time))
- `model` (Attributes) The Model param. (see [below for nested schema](#nestedatt--data--mobile_device--criteria--model))
- `passcode_set` (Boolean) If device's passcode is present.
- `phone_number` (Attributes) The PhoneNumber param. (see [below for nested schema](#nestedatt--data--mobile_device--criteria--phone_number))
- `tag` (Attributes) The Tag param. (see [below for nested schema](#nestedatt--data--mobile_device--criteria--tag))

<a id="nestedatt--data--mobile_device--criteria--applications"></a>
### Nested Schema for `data.mobile_device.criteria.tag`

Read-Only:

- `has_malware` (Attributes) The HasMalware param. (see [below for nested schema](#nestedatt--data--mobile_device--criteria--tag--has_malware))
- `has_unmanaged_app` (Boolean) Has apps that are not managed.
- `includes` (Attributes List) The Includes param. (see [below for nested schema](#nestedatt--data--mobile_device--criteria--tag--includes))

<a id="nestedatt--data--mobile_device--criteria--tag--has_malware"></a>
### Nested Schema for `data.mobile_device.criteria.tag.has_malware`

Read-Only:

- `no` (Boolean) The No param.
- `yes` (Attributes) The Yes param. (see [below for nested schema](#nestedatt--data--mobile_device--criteria--tag--has_malware--yes))

<a id="nestedatt--data--mobile_device--criteria--tag--has_malware--yes"></a>
### Nested Schema for `data.mobile_device.criteria.tag.has_malware.yes`

Read-Only:

- `excludes` (Attributes List) The Excludes param. (see [below for nested schema](#nestedatt--data--mobile_device--criteria--tag--has_malware--yes--excludes))

<a id="nestedatt--data--mobile_device--criteria--tag--has_malware--yes--excludes"></a>
### Nested Schema for `data.mobile_device.criteria.tag.has_malware.yes.excludes`

Read-Only:

- `hash` (String) application hash. String length must not exceed 1024 characters. String validation regex: `.*`.
- `name` (String) The Name param. String length must not exceed 31 characters.
- `package` (String) application package name. String length must not exceed 1024 characters. String validation regex: `.*`.




<a id="nestedatt--data--mobile_device--criteria--tag--includes"></a>
### Nested Schema for `data.mobile_device.criteria.tag.includes`

Read-Only:

- `hash` (String) application hash. String length must not exceed 1024 characters. String validation regex: `.*`.
- `name` (String) The Name param. String length must not exceed 31 characters.
- `package` (String) application package name. String length must not exceed 1024 characters. String validation regex: `.*`.



<a id="nestedatt--data--mobile_device--criteria--imei"></a>
### Nested Schema for `data.mobile_device.criteria.tag`

Read-Only:

- `contains` (String) The Contains param. String length must not exceed 255 characters.
- `is` (String) The Is param. String length must not exceed 255 characters.
- `is_not` (String) The IsNot param. String length must not exceed 255 characters.


<a id="nestedatt--data--mobile_device--criteria--last_checkin_time"></a>
### Nested Schema for `data.mobile_device.criteria.tag`

Read-Only:

- `not_within` (Attributes) The NotWithin param. (see [below for nested schema](#nestedatt--data--mobile_device--criteria--tag--not_within))
- `within` (Attributes) The Within param. (see [below for nested schema](#nestedatt--data--mobile_device--criteria--tag--within))

<a id="nestedatt--data--mobile_device--criteria--tag--not_within"></a>
### Nested Schema for `data.mobile_device.criteria.tag.not_within`

Read-Only:

- `days` (Number) specify time in days. Value must be between 1 and 365. Default: `30`.


<a id="nestedatt--data--mobile_device--criteria--tag--within"></a>
### Nested Schema for `data.mobile_device.criteria.tag.within`

Read-Only:

- `days` (Number) specify time in days. Value must be between 1 and 365. Default: `30`.



<a id="nestedatt--data--mobile_device--criteria--model"></a>
### Nested Schema for `data.mobile_device.criteria.tag`

Read-Only:

- `contains` (String) The Contains param. String length must not exceed 255 characters.
- `is` (String) The Is param. String length must not exceed 255 characters.
- `is_not` (String) The IsNot param. String length must not exceed 255 characters.


<a id="nestedatt--data--mobile_device--criteria--phone_number"></a>
### Nested Schema for `data.mobile_device.criteria.tag`

Read-Only:

- `contains` (String) The Contains param. String length must not exceed 255 characters.
- `is` (String) The Is param. String length must not exceed 255 characters.
- `is_not` (String) The IsNot param. String length must not exceed 255 characters.


<a id="nestedatt--data--mobile_device--criteria--tag"></a>
### Nested Schema for `data.mobile_device.criteria.tag`

Read-Only:

- `contains` (String) The Contains param. String length must not exceed 255 characters.
- `is` (String) The Is param. String length must not exceed 255 characters.
- `is_not` (String) The IsNot param. String length must not exceed 255 characters.




<a id="nestedatt--data--network_info"></a>
### Nested Schema for `data.network_info`

Read-Only:

- `criteria` (Attributes) The Criteria param. (see [below for nested schema](#nestedatt--data--network_info--criteria))

<a id="nestedatt--data--network_info--criteria"></a>
### Nested Schema for `data.network_info.criteria`

Read-Only:

- `network` (Attributes) The Network param. (see [below for nested schema](#nestedatt--data--network_info--criteria--network))

<a id="nestedatt--data--network_info--criteria--network"></a>
### Nested Schema for `data.network_info.criteria.network`

Read-Only:

- `is` (Attributes) The Is param. (see [below for nested schema](#nestedatt--data--network_info--criteria--network--is))
- `is_not` (Attributes) The IsNot param. (see [below for nested schema](#nestedatt--data--network_info--criteria--network--is_not))

<a id="nestedatt--data--network_info--criteria--network--is"></a>
### Nested Schema for `data.network_info.criteria.network.is`

Read-Only:

- `mobile` (Attributes) The Mobile param. (see [below for nested schema](#nestedatt--data--network_info--criteria--network--is--mobile))
- `unknown` (Boolean) The Unknown param.
- `wifi` (Attributes) The Wifi param. (see [below for nested schema](#nestedatt--data--network_info--criteria--network--is--wifi))

<a id="nestedatt--data--network_info--criteria--network--is--mobile"></a>
### Nested Schema for `data.network_info.criteria.network.is.wifi`

Read-Only:

- `carrier` (String) The Carrier param. String length must not exceed 1023 characters. String validation regex: `.*`.


<a id="nestedatt--data--network_info--criteria--network--is--wifi"></a>
### Nested Schema for `data.network_info.criteria.network.is.wifi`

Read-Only:

- `ssid` (String) SSID. String length must not exceed 1023 characters. String validation regex: `.*`.



<a id="nestedatt--data--network_info--criteria--network--is_not"></a>
### Nested Schema for `data.network_info.criteria.network.is_not`

Read-Only:

- `ethernet` (Boolean) The Ethernet param.
- `mobile` (Attributes) The Mobile param. (see [below for nested schema](#nestedatt--data--network_info--criteria--network--is_not--mobile))
- `unknown` (Boolean) The Unknown param.
- `wifi` (Attributes) The Wifi param. (see [below for nested schema](#nestedatt--data--network_info--criteria--network--is_not--wifi))

<a id="nestedatt--data--network_info--criteria--network--is_not--mobile"></a>
### Nested Schema for `data.network_info.criteria.network.is_not.wifi`

Read-Only:

- `carrier` (String) The Carrier param. String length must not exceed 1023 characters. String validation regex: `.*`.


<a id="nestedatt--data--network_info--criteria--network--is_not--wifi"></a>
### Nested Schema for `data.network_info.criteria.network.is_not.wifi`

Read-Only:

- `ssid` (String) SSID. String length must not exceed 1023 characters. String validation regex: `.*`.






<a id="nestedatt--data--patch_management"></a>
### Nested Schema for `data.patch_management`

Read-Only:

- `criteria` (Attributes) The Criteria param. (see [below for nested schema](#nestedatt--data--patch_management--criteria))
- `exclude_vendor` (Boolean) The ExcludeVendor param. Default: `false`.
- `vendors` (Attributes List) Vendor name. (see [below for nested schema](#nestedatt--data--patch_management--vendors))

<a id="nestedatt--data--patch_management--criteria"></a>
### Nested Schema for `data.patch_management.criteria`

Read-Only:

- `is_enabled` (String) is enabled. String must be one of these: `"no"`, `"yes"`, `"not-available"`.
- `is_installed` (Boolean) Is Installed. Default: `true`.
- `missing_patches` (Attributes) The MissingPatches param. (see [below for nested schema](#nestedatt--data--patch_management--criteria--missing_patches))

<a id="nestedatt--data--patch_management--criteria--missing_patches"></a>
### Nested Schema for `data.patch_management.criteria.missing_patches`

Read-Only:

- `check` (String) The Check param. String must be one of these: `"has-any"`, `"has-none"`, `"has-all"`. Default: `"has-any"`.
- `patches` (List of String) The Patches param. Individual elements in this list are subject to additional validation. String length must not exceed 1023 characters. String validation regex: `.*`.
- `severity` (Attributes) The Severity param. (see [below for nested schema](#nestedatt--data--patch_management--criteria--missing_patches--severity))

<a id="nestedatt--data--patch_management--criteria--missing_patches--severity"></a>
### Nested Schema for `data.patch_management.criteria.missing_patches.severity`

Read-Only:

- `greater_equal` (Number) The GreaterEqual param. Value must be between 0 and 100000.
- `greater_than` (Number) The GreaterThan param. Value must be between 0 and 100000.
- `is` (Number) The Is param. Value must be between 0 and 100000.
- `is_not` (Number) The IsNot param. Value must be between 0 and 100000.
- `less_equal` (Number) The LessEqual param. Value must be between 0 and 100000.
- `less_than` (Number) The LessThan param. Value must be between 0 and 100000.




<a id="nestedatt--data--patch_management--vendors"></a>
### Nested Schema for `data.patch_management.vendors`

Read-Only:

- `name` (String) The Name param. String length must not exceed 103 characters.
- `products` (List of String) Product name. Individual elements in this list are subject to additional validation. String length must not exceed 1023 characters. String validation regex: `.*`.
