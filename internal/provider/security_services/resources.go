package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the security_services package
func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		NewAntiSpywareProfileResource,
		NewAntiSpywareSignatureResource,
		NewAppOverrideRuleResource,
		NewDataFilteringProfileResource,
		NewDataObjectResource,
		NewDecryptionExclusionResource,
		NewDecryptionProfileResource,
		NewDecryptionRuleResource,
		NewDnsSecurityProfileResource,
		NewDosProtectionProfileResource,
		NewDosProtectionRuleResource,
		NewFileBlockingProfileResource,
		NewHttpHeaderProfileResource,
		NewProfileGroupResource,
		NewSecurityRuleResource,
		// 		NewSslDecryptionSettingResource,
		NewUrlAccessProfileResource,
		NewUrlCategoryResource,
		NewVulnerabilityProtectionProfileResource,
		NewVulnerabilityProtectionSignatureResource,
		NewWildfireAntiVirusProfileResource,
	}
}
