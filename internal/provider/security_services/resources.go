package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the security_services package
func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		NewUrlCategoryResource,
		NewAppOverrideRuleResource,
		NewDecryptionExclusionResource,
		NewSecurityRuleResource,
		NewAntiSpywareProfileResource,
		NewDosProtectionProfileResource,
		NewHttpHeaderProfileResource,
		NewWildfireAntiVirusProfileResource,
		NewDecryptionRuleResource,
		NewDosProtectionRuleResource,
		NewDnsSecurityProfileResource,
		NewUrlAccessProfileResource,
		NewFileBlockingProfileResource,
		NewDecryptionProfileResource,
		NewVulnerabilityProtectionSignatureResource,
		NewAntiSpywareSignatureResource,
		NewProfileGroupResource,
		NewVulnerabilityProtectionProfileResource,
	}
}
