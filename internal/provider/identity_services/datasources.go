package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// GetDataSources returns the list of data sources for this package.
func GetDataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewLocalUserGroupDataSource,
		NewCertificateProfileDataSource,
		NewLocalUserDataSource,
		NewAuthenticationSequenceDataSource,
		// 		NewMfaServerDataSource,
		// 		NewOcspResponderDataSource,
		NewLdapServerProfileDataSource,
		NewRadiusServerProfileDataSource,
		NewTlsServiceProfileDataSource,
		NewAuthenticationProfileDataSource,
		NewTacacsServerProfileDataSource,
		// 		NewAuthenticationPortalDataSource,
		NewAuthenticationRuleDataSource,
		NewKerberosServerProfileDataSource,
		NewSamlServerProfileDataSource,
		NewScepProfileDataSource,
		NewLocalUserGroupListDataSource,
		NewCertificateProfileListDataSource,
		NewLocalUserListDataSource,
		NewAuthenticationSequenceListDataSource,
		// 		NewMfaServerListDataSource,
		// 		NewOcspResponderListDataSource,
		NewLdapServerProfileListDataSource,
		NewRadiusServerProfileListDataSource,
		NewTlsServiceProfileListDataSource,
		NewAuthenticationProfileListDataSource,
		NewTacacsServerProfileListDataSource,
		// 		NewAuthenticationPortalListDataSource,
		NewAuthenticationRuleListDataSource,
		NewKerberosServerProfileListDataSource,
		NewSamlServerProfileListDataSource,
		NewScepProfileListDataSource,
	}
}
