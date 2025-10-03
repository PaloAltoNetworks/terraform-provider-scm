package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// GetDataSources returns the list of data sources for this package.
func GetDataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewTrafficSteeringRuleDataSource,
		NewServiceConnectionDataSource,
		NewServiceConnectionGroupDataSource,
		// 		NewInternalDnsServerDataSource,
		// 		NewBgpRoutingDataSource,
		// 		NewSharedInfrastructureSettingDataSource,
		NewSiteDataSource,
		// 		NewBandwidthAllocationDataSource,
		NewRemoteNetworkDataSource,
		NewTrafficSteeringRuleListDataSource,
		NewServiceConnectionListDataSource,
		NewServiceConnectionGroupListDataSource,
		// 		NewInternalDnsServerListDataSource,
		// 		NewBgpRoutingListDataSource,
		// 		NewSharedInfrastructureSettingListDataSource,
		NewSiteListDataSource,
		// 		NewBandwidthAllocationListDataSource,
		NewRemoteNetworkListDataSource,
	}
}
