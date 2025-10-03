package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// GetDataSources returns the list of data sources for this package.
func GetDataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewLocalUserDataSource,
		NewLocalUserGroupDataSource,
		NewAuthenticationRuleDataSource,
		NewCertificateProfileDataSource,
		NewSamlServerProfileDataSource,
		NewAuthenticationProfileDataSource,
		NewTlsServiceProfileDataSource,
		// 		NewAuthenticationPortalDataSource,
		NewKerberosServerProfileDataSource,
		// 		NewMfaServerDataSource,
		NewLdapServerProfileDataSource,
		NewScepProfileDataSource,
		// 		NewOcspResponderDataSource,
		NewTacacsServerProfileDataSource,
		NewRadiusServerProfileDataSource,
		NewAuthenticationSequenceDataSource,
		NewLocalUserListDataSource,
		NewLocalUserGroupListDataSource,
		NewAuthenticationRuleListDataSource,
		NewCertificateProfileListDataSource,
		NewSamlServerProfileListDataSource,
		NewAuthenticationProfileListDataSource,
		NewTlsServiceProfileListDataSource,
		// 		NewAuthenticationPortalListDataSource,
		NewKerberosServerProfileListDataSource,
		// 		NewMfaServerListDataSource,
		NewLdapServerProfileListDataSource,
		NewScepProfileListDataSource,
		// 		NewOcspResponderListDataSource,
		NewTacacsServerProfileListDataSource,
		NewRadiusServerProfileListDataSource,
		NewAuthenticationSequenceListDataSource,
	}
}
