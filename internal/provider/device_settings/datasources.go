package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// GetDataSources returns the list of data sources for this package.
func GetDataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAuthenticationSettingDataSource,
		NewContentIdSettingDataSource,
		NewDeviceRedistributionCollectorDataSource,
		NewGeneralSettingDataSource,
		NewManagementInterfaceDataSource,
		NewMotdBannerSettingDataSource,
		NewServiceRouteDataSource,
		NewServiceSettingDataSource,
		NewSessionSettingDataSource,
		NewSessionTimeoutDataSource,
		NewTcpSettingDataSource,
		NewUpdateScheduleDataSource,
		NewVpnSettingDataSource,
		NewAuthenticationSettingListDataSource,
		NewContentIdSettingListDataSource,
		NewDeviceRedistributionCollectorListDataSource,
		NewGeneralSettingListDataSource,
		NewManagementInterfaceListDataSource,
		NewMotdBannerSettingListDataSource,
		NewServiceRouteListDataSource,
		NewServiceSettingListDataSource,
		NewSessionSettingListDataSource,
		NewSessionTimeoutListDataSource,
		NewTcpSettingListDataSource,
		NewUpdateScheduleListDataSource,
		NewVpnSettingListDataSource,
	}
}
