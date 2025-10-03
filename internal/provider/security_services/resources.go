package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the security_services package
func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		NewVulnerabilityProtectionSignatureResource,
		NewDnsSecurityProfileResource,
		NewHttpHeaderProfileResource,
		NewAntiSpywareSignatureResource,
		NewSecurityRuleResource,
		NewDecryptionRuleResource,
		NewDosProtectionProfileResource,
		NewUrlAccessProfileResource,
		NewVulnerabilityProtectionProfileResource,
		NewProfileGroupResource,
		NewDecryptionExclusionResource,
		NewWildfireAntiVirusProfileResource,
		NewFileBlockingProfileResource,
		NewDecryptionProfileResource,
		NewUrlCategoryResource,
		NewAntiSpywareProfileResource,
		NewAppOverrideRuleResource,
		NewDosProtectionRuleResource,
	}
}
