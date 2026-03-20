# ldpap server profile w/ required fields
resource "scm_ldap_server_profile" "scm_ldap_server_profile_one" {
  folder = "All"
  name   = "simple-ldap-profile"

  server = [
    {
      name    = "primary-ldap"
      address = "$tst_68081_1"
      port    = 389
    },
    {
      name    = "secondary-ldap"
      address = "$tst_68081_2"
      port    = 1
    },
    {
      name    = "extra-ldap"
      address = "$test_ip"
      port    = 65535
    }
  ]
}

# ldpap server profile w/ some fields
resource "scm_ldap_server_profile" "scm_ldap_server_profile_two" {
  folder    = "All"
  name      = "intermediate-ldap-profile-one"
  ldap_type = "active-directory"
  base      = "dc=example,dc=com"
  ssl       = true

  server = [
    {
      name    = "extra-ldap"
      address = "$test_ip"
      port    = 25
    }
  ]
}

resource "scm_ldap_server_profile" "scm_ldap_server_profile_three" {
  folder        = "All"
  name          = "intermediate-ldap-profile-two"
  ldap_type     = "sun"
  base          = "DC=internal,DC=company,DC=com"
  bind_dn       = "CN=LDAP Bind,OU=Service Accounts,DC=internal,DC=company,DC=com"
  bind_password = "SecurePwd123!"
  timelimit     = 30

  server = [
    {
      name    = "primary-ldap"
      address = "$tst_68081_1"
      port    = 5000
    },
    {
      name    = "secondary-ldap"
      address = "150.25.25.60"
      port    = 25000
    }
  ]
}

# ldpap server profile w/ all fields
resource "scm_ldap_server_profile" "scm_ldap_server_profile_four" {
  folder                    = "All"
  name                      = "complex-ldap-profile-one"
  ldap_type                 = "e-directory"
  base                      = "ou=users,dc=corp,dc=local"
  bind_dn                   = "cn=admin,dc=corp,dc=local"
  bind_password             = "MyPwd123!"
  bind_timelimit            = "20"
  retry_interval            = 1000
  timelimit                 = 10
  ssl                       = true
  verify_server_certificate = false

  server = [
    {
      name    = "global-directory-main"
      address = "$scm_variable_ipaddr"
      port    = 636
    }
  ]
}

resource "scm_ldap_server_profile" "scm_ldap_server_profile_five" {
  folder                    = "All"
  name                      = "complex-ldap-profile-two"
  ldap_type                 = "other"
  base                      = "OU=Employees,DC=global,DC=enterprise,DC=net"
  bind_dn                   = "CN=SVC_LDAP_Search,OU=Service Accounts,DC=global,DC=enterprise,DC=net"
  bind_password             = "ExtremelyComplexP@ssw0rd!"
  bind_timelimit            = "15"
  retry_interval            = 300
  timelimit                 = 30
  ssl                       = true
  verify_server_certificate = true

  server = [
    {
      name    = "dc-us-east"
      address = "$scm_variable_ipaddr"
      port    = 720
    },
    {
      name    = "dc-us-west"
      address = "192.10.10.10"
      port    = 1400
    },
    {
      name    = "dc-eu-central"
      address = "3.3.3.3"
      port    = 20000
    }
  ]
}