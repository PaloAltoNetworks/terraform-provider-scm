package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// GetDataSources returns the list of data sources for this package.
func GetDataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewSiteDataSource,
		// 		NewBgpRoutingDataSource,
		NewServiceConnectionDataSource,
		NewServiceConnectionGroupDataSource,
		// 		NewBandwidthAllocationDataSource,
		// 		NewInternalDnsServerDataSource,
		NewTrafficSteeringRuleDataSource,
		NewRemoteNetworkDataSource,
		// 		NewSharedInfrastructureSettingDataSource,
		NewSiteListDataSource,
		// 		NewBgpRoutingListDataSource,
		NewServiceConnectionListDataSource,
		NewServiceConnectionGroupListDataSource,
		// 		NewBandwidthAllocationListDataSource,
		// 		NewInternalDnsServerListDataSource,
		NewTrafficSteeringRuleListDataSource,
		NewRemoteNetworkListDataSource,
		// 		NewSharedInfrastructureSettingListDataSource,
	}
}
