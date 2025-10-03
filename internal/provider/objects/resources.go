package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the objects package
func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		// 		NewQuarantinedDeviceResource,
		NewTagResource,
		NewSyslogServerProfileResource,
		NewHipObjectResource,
		NewApplicationResource,
		NewAddressGroupResource,
		NewApplicationFilterResource,
		NewServiceGroupResource,
		NewServiceResource,
		NewHttpServerProfileResource,
		NewExternalDynamicListResource,
		NewHipProfileResource,
		NewApplicationGroupResource,
		NewLogForwardingProfileResource,
		NewAddressResource,
		// 		NewAutoTagActionResource,
		NewDynamicUserGroupResource,
		NewScheduleResource,
		NewRegionResource,
	}
}
