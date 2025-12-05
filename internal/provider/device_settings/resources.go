package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the device_settings package
func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		NewAuthenticationSettingResource,
		NewContentIdSettingResource,
		NewDeviceRedistributionCollectorResource,
		NewGeneralSettingResource,
		NewManagementInterfaceResource,
		NewMotdBannerSettingResource,
		NewServiceRouteResource,
		NewServiceSettingResource,
		NewSessionSettingResource,
		NewSessionTimeoutResource,
		NewTcpSettingResource,
		NewUpdateScheduleResource,
		NewVpnSettingResource,
	}
}
