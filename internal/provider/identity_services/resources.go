package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the identity_services package
func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		NewLocalUserResource,
		NewLocalUserGroupResource,
		NewAuthenticationRuleResource,
		NewCertificateProfileResource,
		NewSamlServerProfileResource,
		NewAuthenticationProfileResource,
		NewTlsServiceProfileResource,
		// 		NewAuthenticationPortalResource,
		NewKerberosServerProfileResource,
		NewMfaServerResource,
		NewLdapServerProfileResource,
		NewScepProfileResource,
		// 		NewOcspResponderResource,
		// 		NewCertificatesGetResource,
		NewTacacsServerProfileResource,
		NewRadiusServerProfileResource,
		NewAuthenticationSequenceResource,
	}
}
