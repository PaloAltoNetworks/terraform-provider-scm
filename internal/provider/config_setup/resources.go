package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the config_setup package
func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		NewFolderResource,
		NewLabelResource,
		NewSnippetResource,
		NewVariableResource,
	}
}
