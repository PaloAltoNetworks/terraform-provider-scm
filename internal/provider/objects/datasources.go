package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// GetDataSources returns the list of data sources for this package.
func GetDataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		// 		NewQuarantinedDeviceDataSource,
		NewTagDataSource,
		NewSyslogServerProfileDataSource,
		NewHipObjectDataSource,
		NewApplicationDataSource,
		NewAddressGroupDataSource,
		NewApplicationFilterDataSource,
		NewServiceGroupDataSource,
		NewServiceDataSource,
		NewHttpServerProfileDataSource,
		NewExternalDynamicListDataSource,
		NewHipProfileDataSource,
		NewApplicationGroupDataSource,
		NewLogForwardingProfileDataSource,
		NewAddressDataSource,
		// 		NewAutoTagActionDataSource,
		NewDynamicUserGroupDataSource,
		NewScheduleDataSource,
		NewRegionDataSource,
		// 		NewQuarantinedDeviceListDataSource,
		NewTagListDataSource,
		NewSyslogServerProfileListDataSource,
		NewHipObjectListDataSource,
		NewApplicationListDataSource,
		NewAddressGroupListDataSource,
		NewApplicationFilterListDataSource,
		NewServiceGroupListDataSource,
		NewServiceListDataSource,
		NewHttpServerProfileListDataSource,
		NewExternalDynamicListListDataSource,
		NewHipProfileListDataSource,
		NewApplicationGroupListDataSource,
		NewLogForwardingProfileListDataSource,
		NewAddressListDataSource,
		// 		NewAutoTagActionListDataSource,
		NewDynamicUserGroupListDataSource,
		NewScheduleListDataSource,
		NewRegionListDataSource,
	}
}
