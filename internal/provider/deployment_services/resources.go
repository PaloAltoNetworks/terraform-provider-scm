package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the deployment_services package
func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		// 		NewSiteResource,
		// 		NewBgpRoutingResource,
		NewServiceConnectionResource,
		NewServiceConnectionGroupResource,
		// 		NewBandwidthAllocationResource,
		NewInternalDnsServerResource,
		NewTrafficSteeringRuleResource,
		NewRemoteNetworkResource,
		// 		NewSharedInfrastructureSettingResource,
	}
}
