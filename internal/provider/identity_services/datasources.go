package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// GetDataSources returns the list of data sources for this package.
func GetDataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		// 		NewAuthenticationPortalDataSource,
		NewAuthenticationProfileDataSource,
		NewAuthenticationRuleDataSource,
		NewAuthenticationSequenceDataSource,
		NewCertificateProfileDataSource,
		NewKerberosServerProfileDataSource,
		NewLdapServerProfileDataSource,
		NewLocalUserDataSource,
		NewLocalUserGroupDataSource,
		// 		NewMfaServerDataSource,
		// 		NewOcspResponderDataSource,
		NewRadiusServerProfileDataSource,
		NewSamlServerProfileDataSource,
		NewScepProfileDataSource,
		NewTacacsServerProfileDataSource,
		NewTlsServiceProfileDataSource,
		// 		NewAuthenticationPortalListDataSource,
		NewAuthenticationProfileListDataSource,
		NewAuthenticationRuleListDataSource,
		NewAuthenticationSequenceListDataSource,
		NewCertificateProfileListDataSource,
		NewKerberosServerProfileListDataSource,
		NewLdapServerProfileListDataSource,
		NewLocalUserListDataSource,
		NewLocalUserGroupListDataSource,
		// 		NewMfaServerListDataSource,
		// 		NewOcspResponderListDataSource,
		NewRadiusServerProfileListDataSource,
		NewSamlServerProfileListDataSource,
		NewScepProfileListDataSource,
		NewTacacsServerProfileListDataSource,
		NewTlsServiceProfileListDataSource,
	}
}
