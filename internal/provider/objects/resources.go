package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the objects package
func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		// 		NewAutoTagActionResource,
		NewAddressResource,
		NewHipObjectResource,
		NewAddressGroupResource,
		NewTagResource,
		NewApplicationFilterResource,
		NewServiceGroupResource,
		NewRegionResource,
		NewApplicationResource,
		NewApplicationGroupResource,
		NewHipProfileResource,
		NewHttpServerProfileResource,
		NewExternalDynamicListResource,
		// 		NewQuarantinedDeviceResource,
		NewSyslogServerProfileResource,
		NewDynamicUserGroupResource,
		NewLogForwardingProfileResource,
		NewServiceResource,
		NewScheduleResource,
	}
}
