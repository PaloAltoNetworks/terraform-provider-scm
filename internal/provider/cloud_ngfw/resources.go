package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the cloud_ngfw package
func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		NewNgfirewallResource,
	}
}
