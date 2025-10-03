package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// GetDataSources returns the list of data sources for this package.
func GetDataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewScepProfileDataSource,
		// 		NewAuthenticationPortalDataSource,
		NewTacacsServerProfileDataSource,
		NewLdapServerProfileDataSource,
		NewAuthenticationRuleDataSource,
		// 		NewMfaServerDataSource,
		NewRadiusServerProfileDataSource,
		NewAuthenticationSequenceDataSource,
		NewCertificateProfileDataSource,
		NewAuthenticationProfileDataSource,
		NewSamlServerProfileDataSource,
		NewKerberosServerProfileDataSource,
		NewLocalUserGroupDataSource,
		NewTlsServiceProfileDataSource,
		// 		NewOcspResponderDataSource,
		NewLocalUserDataSource,
		NewScepProfileListDataSource,
		// 		NewAuthenticationPortalListDataSource,
		NewTacacsServerProfileListDataSource,
		NewLdapServerProfileListDataSource,
		NewAuthenticationRuleListDataSource,
		// 		NewMfaServerListDataSource,
		NewRadiusServerProfileListDataSource,
		NewAuthenticationSequenceListDataSource,
		NewCertificateProfileListDataSource,
		NewAuthenticationProfileListDataSource,
		NewSamlServerProfileListDataSource,
		NewKerberosServerProfileListDataSource,
		NewLocalUserGroupListDataSource,
		NewTlsServiceProfileListDataSource,
		// 		NewOcspResponderListDataSource,
		NewLocalUserListDataSource,
	}
}
