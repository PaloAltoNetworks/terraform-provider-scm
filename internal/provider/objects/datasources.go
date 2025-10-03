package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// GetDataSources returns the list of data sources for this package.
func GetDataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewApplicationDataSource,
		NewServiceGroupDataSource,
		NewHipProfileDataSource,
		NewExternalDynamicListDataSource,
		NewAddressDataSource,
		NewDynamicUserGroupDataSource,
		NewTagDataSource,
		// 		NewQuarantinedDeviceDataSource,
		NewHipObjectDataSource,
		NewRegionDataSource,
		NewApplicationGroupDataSource,
		NewAddressGroupDataSource,
		NewLogForwardingProfileDataSource,
		NewServiceDataSource,
		// 		NewAutoTagActionDataSource,
		NewSyslogServerProfileDataSource,
		NewScheduleDataSource,
		NewApplicationFilterDataSource,
		NewHttpServerProfileDataSource,
		NewApplicationListDataSource,
		NewServiceGroupListDataSource,
		NewHipProfileListDataSource,
		NewExternalDynamicListListDataSource,
		NewAddressListDataSource,
		NewDynamicUserGroupListDataSource,
		NewTagListDataSource,
		// 		NewQuarantinedDeviceListDataSource,
		NewHipObjectListDataSource,
		NewRegionListDataSource,
		NewApplicationGroupListDataSource,
		NewAddressGroupListDataSource,
		NewLogForwardingProfileListDataSource,
		NewServiceListDataSource,
		// 		NewAutoTagActionListDataSource,
		NewSyslogServerProfileListDataSource,
		NewScheduleListDataSource,
		NewApplicationFilterListDataSource,
		NewHttpServerProfileListDataSource,
	}
}
