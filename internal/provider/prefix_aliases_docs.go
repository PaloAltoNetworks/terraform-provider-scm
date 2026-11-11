//go:build docsgen

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Built with the "docsgen" tag: no deprecated unprefixed aliases, so the docs
// contain exactly one page per scm_ type.

func legacyResources(ctors []func() resource.Resource) []func() resource.Resource { return ctors }

func legacyDataSources(ctors []func() datasource.DataSource) []func() datasource.DataSource {
	return ctors
}

func legacyActions(ctors []func() action.Action) []func() action.Action { return ctors }
