package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the objects package
func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		NewAddressResource,
		NewAddressGroupResource,
		NewApplicationResource,
		NewApplicationFilterResource,
		NewApplicationGroupResource,
		// 		NewAutoTagActionResource,
		NewDynamicUserGroupResource,
		NewExternalDynamicListResource,
		NewHipObjectResource,
		NewHipProfileResource,
		NewHttpServerProfileResource,
		NewLogForwardingProfileResource,
		// 		NewQuarantinedDeviceResource,
		NewRegionResource,
		NewScheduleResource,
		NewServiceResource,
		NewServiceGroupResource,
		NewSyslogServerProfileResource,
		NewTagResource,
	}
}
