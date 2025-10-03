package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the identity_services package
func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		NewScepProfileResource,
		// 		NewAuthenticationPortalResource,
		NewTacacsServerProfileResource,
		NewLdapServerProfileResource,
		NewAuthenticationRuleResource,
		NewMfaServerResource,
		// 		NewCertificatesGetResource,
		NewRadiusServerProfileResource,
		NewAuthenticationSequenceResource,
		NewCertificateProfileResource,
		NewAuthenticationProfileResource,
		NewSamlServerProfileResource,
		NewKerberosServerProfileResource,
		NewLocalUserGroupResource,
		NewTlsServiceProfileResource,
		// 		NewOcspResponderResource,
		NewLocalUserResource,
	}
}
