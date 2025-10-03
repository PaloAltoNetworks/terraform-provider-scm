package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the objects package
func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		NewApplicationResource,
		NewServiceGroupResource,
		NewHipProfileResource,
		NewExternalDynamicListResource,
		NewAddressResource,
		NewDynamicUserGroupResource,
		NewTagResource,
		// 		NewQuarantinedDeviceResource,
		NewHipObjectResource,
		NewRegionResource,
		NewApplicationGroupResource,
		NewAddressGroupResource,
		NewLogForwardingProfileResource,
		NewServiceResource,
		// 		NewAutoTagActionResource,
		NewSyslogServerProfileResource,
		NewScheduleResource,
		NewApplicationFilterResource,
		NewHttpServerProfileResource,
	}
}
