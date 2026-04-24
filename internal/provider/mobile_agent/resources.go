package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the mobile_agent package
func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		NewForwardingProfileResource,
		NewForwardingProfileDestinationResource,
		NewForwardingProfileRegionalAndCustomProxyResource,
		NewForwardingProfileSourceApplicationResource,
		NewForwardingProfileUserLocationResource,
	}
}
