package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// GetDataSources returns the list of data sources for this package.
func GetDataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewSiteDataSource,
		// 		NewBgpRoutingDataSource,
		NewRemoteNetworkDataSource,
		NewTrafficSteeringRuleDataSource,
		// 		NewSharedInfrastructureSettingDataSource,
		NewServiceConnectionGroupDataSource,
		// 		NewBandwidthAllocationDataSource,
		NewServiceConnectionDataSource,
		// 		NewInternalDnsServerDataSource,
		NewSiteListDataSource,
		// 		NewBgpRoutingListDataSource,
		NewRemoteNetworkListDataSource,
		NewTrafficSteeringRuleListDataSource,
		// 		NewSharedInfrastructureSettingListDataSource,
		NewServiceConnectionGroupListDataSource,
		// 		NewBandwidthAllocationListDataSource,
		NewServiceConnectionListDataSource,
		// 		NewInternalDnsServerListDataSource,
	}
}
