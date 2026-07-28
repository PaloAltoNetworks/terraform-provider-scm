package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// GetDataSources returns the list of data sources for this package.
func GetDataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewBandwidthAllocationDataSource,
		NewBgpRoutingDataSource,
		// 		NewInternalDnsServerDataSource,
		NewRemoteNetworkDataSource,
		NewServiceConnectionDataSource,
		NewServiceConnectionGroupDataSource,
		// 		NewSharedInfrastructureSettingDataSource,
		NewSiteDataSource,
		NewTrafficSteeringRuleDataSource,
		NewBandwidthAllocationListDataSource,
		// 		NewInternalDnsServerListDataSource,
		NewLocationListDataSource,
		NewRemoteNetworkListDataSource,
		NewServiceConnectionListDataSource,
		NewServiceConnectionGroupListDataSource,
		NewSiteListDataSource,
		NewTrafficSteeringRuleListDataSource,
	}
}
