package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// GetDataSources returns the list of data sources for this package.
func GetDataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAddressDataSource,
		NewAddressGroupDataSource,
		NewApplicationDataSource,
		NewApplicationFilterDataSource,
		NewApplicationGroupDataSource,
		// 		NewAutoTagActionDataSource,
		NewDynamicUserGroupDataSource,
		NewExternalDynamicListDataSource,
		NewHipObjectDataSource,
		NewHipProfileDataSource,
		NewHttpServerProfileDataSource,
		NewLogForwardingProfileDataSource,
		// 		NewQuarantinedDeviceDataSource,
		NewRegionDataSource,
		NewScheduleDataSource,
		NewServiceDataSource,
		NewServiceGroupDataSource,
		NewSyslogServerProfileDataSource,
		NewTagDataSource,
		NewAddressListDataSource,
		NewAddressGroupListDataSource,
		NewApplicationListDataSource,
		NewApplicationFilterListDataSource,
		NewApplicationGroupListDataSource,
		NewDynamicUserGroupListDataSource,
		NewExternalDynamicListListDataSource,
		NewHipObjectListDataSource,
		NewHipProfileListDataSource,
		NewHttpServerProfileListDataSource,
		NewLogForwardingProfileListDataSource,
		NewRegionListDataSource,
		NewScheduleListDataSource,
		NewServiceListDataSource,
		NewServiceGroupListDataSource,
		NewSyslogServerProfileListDataSource,
		NewTagListDataSource,
	}
}
