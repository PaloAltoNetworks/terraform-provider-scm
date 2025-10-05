package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the identity_services package
func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		// 		NewAuthenticationPortalResource,
		NewAuthenticationProfileResource,
		NewAuthenticationRuleResource,
		NewAuthenticationSequenceResource,
		NewCertificateProfileResource,
		// 		NewCertificatesGetResource,
		NewKerberosServerProfileResource,
		NewLdapServerProfileResource,
		NewLocalUserResource,
		NewLocalUserGroupResource,
		NewMfaServerResource,
		// 		NewOcspResponderResource,
		NewRadiusServerProfileResource,
		NewSamlServerProfileResource,
		NewScepProfileResource,
		NewTacacsServerProfileResource,
		NewTlsServiceProfileResource,
	}
}
