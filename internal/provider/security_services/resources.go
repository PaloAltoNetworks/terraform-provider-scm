package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the security_services package
func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		NewWildfireAntiVirusProfileResource,
		NewDosProtectionRuleResource,
		NewAntiSpywareProfileResource,
		NewFileBlockingProfileResource,
		NewUrlAccessProfileResource,
		NewHttpHeaderProfileResource,
		NewSecurityRuleResource,
		NewDecryptionProfileResource,
		NewUrlCategoryResource,
		NewVulnerabilityProtectionSignatureResource,
		NewDnsSecurityProfileResource,
		NewAntiSpywareSignatureResource,
		NewDecryptionRuleResource,
		NewDecryptionExclusionResource,
		NewAppOverrideRuleResource,
		NewVulnerabilityProtectionProfileResource,
		NewDosProtectionProfileResource,
		NewProfileGroupResource,
	}
}
