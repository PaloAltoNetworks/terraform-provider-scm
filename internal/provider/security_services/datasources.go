package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// GetDataSources returns the list of data sources for this package.
func GetDataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAntiSpywareProfileDataSource,
		NewAntiSpywareSignatureDataSource,
		NewAppOverrideRuleDataSource,
		NewDataFilteringProfileDataSource,
		NewDataObjectDataSource,
		NewDecryptionExclusionDataSource,
		NewDecryptionProfileDataSource,
		NewDecryptionRuleDataSource,
		NewDnsSecurityProfileDataSource,
		NewDosProtectionProfileDataSource,
		NewDosProtectionRuleDataSource,
		NewFileBlockingProfileDataSource,
		NewHttpHeaderProfileDataSource,
		NewProfileGroupDataSource,
		NewSecurityRuleDataSource,
		// 		NewSslDecryptionSettingDataSource,
		NewUrlAccessProfileDataSource,
		NewUrlCategoryDataSource,
		NewVulnerabilityProtectionProfileDataSource,
		NewVulnerabilityProtectionSignatureDataSource,
		NewWildfireAntiVirusProfileDataSource,
		NewAntiSpywareProfileListDataSource,
		NewAntiSpywareSignatureListDataSource,
		NewAppOverrideRuleListDataSource,
		NewDataFilteringProfileListDataSource,
		NewDataObjectListDataSource,
		NewDecryptionExclusionListDataSource,
		NewDecryptionProfileListDataSource,
		NewDecryptionRuleListDataSource,
		NewDnsSecurityProfileListDataSource,
		NewDosProtectionProfileListDataSource,
		NewDosProtectionRuleListDataSource,
		NewFileBlockingProfileListDataSource,
		NewHttpHeaderProfileListDataSource,
		NewProfileGroupListDataSource,
		NewSecurityRuleListDataSource,
		NewUrlAccessProfileListDataSource,
		NewUrlCategoryListDataSource,
		NewVulnerabilityProtectionProfileListDataSource,
		NewVulnerabilityProtectionSignatureListDataSource,
		NewWildfireAntiVirusProfileListDataSource,
	}
}
