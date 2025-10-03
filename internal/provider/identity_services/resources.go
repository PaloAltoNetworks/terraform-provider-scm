package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the identity_services package
func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		NewLocalUserGroupResource,
		NewCertificateProfileResource,
		NewLocalUserResource,
		NewAuthenticationSequenceResource,
		NewMfaServerResource,
		// 		NewOcspResponderResource,
		NewLdapServerProfileResource,
		// 		NewCertificatesGetResource,
		NewRadiusServerProfileResource,
		NewTlsServiceProfileResource,
		NewAuthenticationProfileResource,
		NewTacacsServerProfileResource,
		// 		NewAuthenticationPortalResource,
		NewAuthenticationRuleResource,
		NewKerberosServerProfileResource,
		NewSamlServerProfileResource,
		NewScepProfileResource,
	}
}
