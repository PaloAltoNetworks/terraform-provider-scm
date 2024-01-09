---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_authentication_profile Data Source - scm"
subcategory: ""
description: |-
  Retrieves a config item.
---

# scm_authentication_profile (Data Source)

Retrieves a config item.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The Id param.

### Read-Only

- `allow_list` (List of String) The AllowList param.
- `lockout` (Attributes) The Lockout param. (see [below for nested schema](#nestedatt--lockout))
- `method` (Attributes) The Method param. (see [below for nested schema](#nestedatt--method))
- `multi_factor_auth` (Attributes) The MultiFactorAuth param. (see [below for nested schema](#nestedatt--multi_factor_auth))
- `name` (String) The Name param.
- `single_sign_on` (Attributes) The SingleSignOn param. (see [below for nested schema](#nestedatt--single_sign_on))
- `tfid` (String) The Terraform ID.
- `user_domain` (String) The UserDomain param. String length must not exceed 63 characters.
- `username_modifier` (String) The UsernameModifier param. String must be one of these: `"%USERINPUT%"`, `"%USERINPUT%@%USERDOMAIN%"`, `"%USERDOMAIN%\\%USERINPUT%"`.

<a id="nestedatt--lockout"></a>
### Nested Schema for `lockout`

Read-Only:

- `failed_attempts` (Number) The FailedAttempts param. Value must be between 0 and 10.
- `lockout_time` (Number) The LockoutTime param. Value must be between 0 and 60.


<a id="nestedatt--method"></a>
### Nested Schema for `method`

Read-Only:

- `cloud` (Attributes) The Cloud param. (see [below for nested schema](#nestedatt--method--cloud))
- `kerberos` (Attributes) The Kerberos param. (see [below for nested schema](#nestedatt--method--kerberos))
- `ldap` (Attributes) The Ldap param. (see [below for nested schema](#nestedatt--method--ldap))
- `local_database` (Boolean) The LocalDatabase param.
- `radius` (Attributes) The Radius param. (see [below for nested schema](#nestedatt--method--radius))
- `saml_idp` (Attributes) The SamlIdp param. (see [below for nested schema](#nestedatt--method--saml_idp))
- `tacplus` (Attributes) The Tacplus param. (see [below for nested schema](#nestedatt--method--tacplus))

<a id="nestedatt--method--cloud"></a>
### Nested Schema for `method.cloud`

Read-Only:

- `profile_name` (String) The tenant profile name.


<a id="nestedatt--method--kerberos"></a>
### Nested Schema for `method.kerberos`

Read-Only:

- `realm` (String) The Realm param.
- `server_profile` (String) The ServerProfile param.


<a id="nestedatt--method--ldap"></a>
### Nested Schema for `method.ldap`

Read-Only:

- `login_attribute` (String) The LoginAttribute param.
- `passwd_exp_days` (Number) The PasswdExpDays param.
- `server_profile` (String) The ServerProfile param.


<a id="nestedatt--method--radius"></a>
### Nested Schema for `method.radius`

Read-Only:

- `checkgroup` (Boolean) The Checkgroup param.
- `server_profile` (String) The ServerProfile param.


<a id="nestedatt--method--saml_idp"></a>
### Nested Schema for `method.saml_idp`

Read-Only:

- `attribute_name_usergroup` (String) The AttributeNameUsergroup param. String length must be between 1 and 63 characters.
- `attribute_name_username` (String) The AttributeNameUsername param. String length must be between 1 and 63 characters.
- `certificate_profile` (String) The CertificateProfile param. String length must not exceed 31 characters.
- `enable_single_logout` (Boolean) The EnableSingleLogout param.
- `request_signing_certificate` (String) The RequestSigningCertificate param. String length must not exceed 64 characters.
- `server_profile` (String) The ServerProfile param. String length must not exceed 63 characters.


<a id="nestedatt--method--tacplus"></a>
### Nested Schema for `method.tacplus`

Read-Only:

- `checkgroup` (Boolean) The Checkgroup param.
- `server_profile` (String) The ServerProfile param.



<a id="nestedatt--multi_factor_auth"></a>
### Nested Schema for `multi_factor_auth`

Read-Only:

- `factors` (List of String) The Factors param.
- `mfa_enable` (Boolean) The MfaEnable param.


<a id="nestedatt--single_sign_on"></a>
### Nested Schema for `single_sign_on`

Read-Only:

- `kerberos_keytab` (String) The KerberosKeytab param. String length must not exceed 8192 characters.
- `realm` (String) The Realm param. String length must not exceed 127 characters.