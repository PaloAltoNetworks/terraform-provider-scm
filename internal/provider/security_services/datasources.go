package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// GetDataSources returns the list of data sources for this package.
func GetDataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewWildfireAntiVirusProfileDataSource,
		NewDosProtectionRuleDataSource,
		NewAntiSpywareProfileDataSource,
		NewFileBlockingProfileDataSource,
		NewUrlAccessProfileDataSource,
		NewHttpHeaderProfileDataSource,
		NewSecurityRuleDataSource,
		NewDecryptionProfileDataSource,
		NewUrlCategoryDataSource,
		NewVulnerabilityProtectionSignatureDataSource,
		NewDnsSecurityProfileDataSource,
		NewAntiSpywareSignatureDataSource,
		NewDecryptionRuleDataSource,
		NewDecryptionExclusionDataSource,
		NewAppOverrideRuleDataSource,
		NewVulnerabilityProtectionProfileDataSource,
		NewDosProtectionProfileDataSource,
		NewProfileGroupDataSource,
		NewWildfireAntiVirusProfileListDataSource,
		NewDosProtectionRuleListDataSource,
		NewAntiSpywareProfileListDataSource,
		NewFileBlockingProfileListDataSource,
		NewUrlAccessProfileListDataSource,
		NewHttpHeaderProfileListDataSource,
		NewSecurityRuleListDataSource,
		NewDecryptionProfileListDataSource,
		NewUrlCategoryListDataSource,
		NewVulnerabilityProtectionSignatureListDataSource,
		NewDnsSecurityProfileListDataSource,
		NewAntiSpywareSignatureListDataSource,
		NewDecryptionRuleListDataSource,
		NewDecryptionExclusionListDataSource,
		NewAppOverrideRuleListDataSource,
		NewVulnerabilityProtectionProfileListDataSource,
		NewDosProtectionProfileListDataSource,
		NewProfileGroupListDataSource,
	}
}
