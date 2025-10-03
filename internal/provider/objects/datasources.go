package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// GetDataSources returns the list of data sources for this package.
func GetDataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		// 		NewAutoTagActionDataSource,
		NewAddressDataSource,
		NewHipObjectDataSource,
		NewAddressGroupDataSource,
		NewTagDataSource,
		NewApplicationFilterDataSource,
		NewServiceGroupDataSource,
		NewRegionDataSource,
		NewApplicationDataSource,
		NewApplicationGroupDataSource,
		NewHipProfileDataSource,
		NewHttpServerProfileDataSource,
		NewExternalDynamicListDataSource,
		// 		NewQuarantinedDeviceDataSource,
		NewSyslogServerProfileDataSource,
		NewDynamicUserGroupDataSource,
		NewLogForwardingProfileDataSource,
		NewServiceDataSource,
		NewScheduleDataSource,
		// 		NewAutoTagActionListDataSource,
		NewAddressListDataSource,
		NewHipObjectListDataSource,
		NewAddressGroupListDataSource,
		NewTagListDataSource,
		NewApplicationFilterListDataSource,
		NewServiceGroupListDataSource,
		NewRegionListDataSource,
		NewApplicationListDataSource,
		NewApplicationGroupListDataSource,
		NewHipProfileListDataSource,
		NewHttpServerProfileListDataSource,
		NewExternalDynamicListListDataSource,
		// 		NewQuarantinedDeviceListDataSource,
		NewSyslogServerProfileListDataSource,
		NewDynamicUserGroupListDataSource,
		NewLogForwardingProfileListDataSource,
		NewServiceListDataSource,
		NewScheduleListDataSource,
	}
}
