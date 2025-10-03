package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// GetDataSources returns the list of data sources for this package.
func GetDataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewUrlCategoryDataSource,
		NewAppOverrideRuleDataSource,
		NewDecryptionExclusionDataSource,
		NewSecurityRuleDataSource,
		NewAntiSpywareProfileDataSource,
		NewDosProtectionProfileDataSource,
		NewHttpHeaderProfileDataSource,
		NewWildfireAntiVirusProfileDataSource,
		NewDecryptionRuleDataSource,
		NewDosProtectionRuleDataSource,
		NewDnsSecurityProfileDataSource,
		NewUrlAccessProfileDataSource,
		NewFileBlockingProfileDataSource,
		NewDecryptionProfileDataSource,
		NewVulnerabilityProtectionSignatureDataSource,
		NewAntiSpywareSignatureDataSource,
		NewProfileGroupDataSource,
		NewVulnerabilityProtectionProfileDataSource,
		NewUrlCategoryListDataSource,
		NewAppOverrideRuleListDataSource,
		NewDecryptionExclusionListDataSource,
		NewSecurityRuleListDataSource,
		NewAntiSpywareProfileListDataSource,
		NewDosProtectionProfileListDataSource,
		NewHttpHeaderProfileListDataSource,
		NewWildfireAntiVirusProfileListDataSource,
		NewDecryptionRuleListDataSource,
		NewDosProtectionRuleListDataSource,
		NewDnsSecurityProfileListDataSource,
		NewUrlAccessProfileListDataSource,
		NewFileBlockingProfileListDataSource,
		NewDecryptionProfileListDataSource,
		NewVulnerabilityProtectionSignatureListDataSource,
		NewAntiSpywareSignatureListDataSource,
		NewProfileGroupListDataSource,
		NewVulnerabilityProtectionProfileListDataSource,
	}
}
