package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for this package.
func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		NewLinkResource,
	}
}
