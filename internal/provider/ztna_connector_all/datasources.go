package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// GetDataSources returns the list of data sources for this package.
func GetDataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewConnectorDataSource,
		NewConnectorGroupDataSource,
		NewFqdnApplicationDataSource,
		NewLicenseDataSource,
		NewSubnetDataSource,
		NewTenantStatusDataSource,
		NewWildcardDataSource,
		NewConnectorListDataSource,
		NewConnectorGroupListDataSource,
		NewConnectorImageListDataSource,
		NewDiscoveredApplicationListDataSource,
		NewFqdnApplicationListDataSource,
		NewSubnetListDataSource,
		NewWildcardListDataSource,
		NewApplicationFiltersDataSource,
		NewConnectorFiltersDataSource,
		NewConnectorGroupFiltersDataSource,
		NewDiscoveredApplicationFiltersDataSource,
		NewSubnetFiltersDataSource,
		NewWildcardFiltersDataSource,
		NewConnectorGroupConnectorsDataSource,
		NewConnectorGroupFqdnRulesDataSource,
		NewConnectorGroupScheduledUpgradeDataSource,
		NewConnectorGroupSubnetRulesDataSource,
		NewConnectorGroupUpgradeStatusDataSource,
		NewConnectorGroupWildcardsDataSource,
		NewConnectorQuiesceDataSource,
		NewConnectorScheduledUpgradeDataSource,
		NewConnectorUpgradeStatusDataSource,
	}
}
