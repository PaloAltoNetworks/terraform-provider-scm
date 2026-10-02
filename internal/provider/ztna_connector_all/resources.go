package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the ztna_connector_all package
func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		NewConnectorResource,
		NewConnectorGroupResource,
		NewFqdnApplicationResource,
		NewSubnetResource,
		NewWildcardResource,
		NewConnectorGroupScheduledUpgradeResource,
		NewConnectorQuiesceResource,
		NewConnectorScheduledUpgradeResource,
	}
}
