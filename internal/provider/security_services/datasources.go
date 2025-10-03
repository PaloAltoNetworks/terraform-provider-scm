package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// GetDataSources returns the list of data sources for this package.
func GetDataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewVulnerabilityProtectionSignatureDataSource,
		NewDnsSecurityProfileDataSource,
		NewHttpHeaderProfileDataSource,
		NewAntiSpywareSignatureDataSource,
		NewSecurityRuleDataSource,
		NewDecryptionRuleDataSource,
		NewDosProtectionProfileDataSource,
		NewUrlAccessProfileDataSource,
		NewVulnerabilityProtectionProfileDataSource,
		NewProfileGroupDataSource,
		NewDecryptionExclusionDataSource,
		NewWildfireAntiVirusProfileDataSource,
		NewFileBlockingProfileDataSource,
		NewDecryptionProfileDataSource,
		NewUrlCategoryDataSource,
		NewAntiSpywareProfileDataSource,
		NewAppOverrideRuleDataSource,
		NewDosProtectionRuleDataSource,
		NewVulnerabilityProtectionSignatureListDataSource,
		NewDnsSecurityProfileListDataSource,
		NewHttpHeaderProfileListDataSource,
		NewAntiSpywareSignatureListDataSource,
		NewSecurityRuleListDataSource,
		NewDecryptionRuleListDataSource,
		NewDosProtectionProfileListDataSource,
		NewUrlAccessProfileListDataSource,
		NewVulnerabilityProtectionProfileListDataSource,
		NewProfileGroupListDataSource,
		NewDecryptionExclusionListDataSource,
		NewWildfireAntiVirusProfileListDataSource,
		NewFileBlockingProfileListDataSource,
		NewDecryptionProfileListDataSource,
		NewUrlCategoryListDataSource,
		NewAntiSpywareProfileListDataSource,
		NewAppOverrideRuleListDataSource,
		NewDosProtectionRuleListDataSource,
	}
}
