//go:build docsgen

package provider

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// With the docsgen tag there are no deprecated aliases, so every registered
// type must start with scm_ (otherwise tfplugindocs would emit clashing pages).
func TestDocsBuildHasOnlyScmTypes(t *testing.T) {
	ctx := context.Background()
	p := New("test").(*ScmProvider)
	check := func(kind, name string) {
		if !strings.HasPrefix(name, "scm_") {
			t.Errorf("%s %q does not start with scm_ in the docs build", kind, name)
		}
	}
	for _, c := range p.Resources(ctx) {
		var m resource.MetadataResponse
		c().Metadata(ctx, resource.MetadataRequest{ProviderTypeName: "scm"}, &m)
		check("resource", m.TypeName)
	}
	for _, c := range p.DataSources(ctx) {
		var m datasource.MetadataResponse
		c().Metadata(ctx, datasource.MetadataRequest{ProviderTypeName: "scm"}, &m)
		check("data source", m.TypeName)
	}
	for _, c := range p.Actions(ctx) {
		var m action.MetadataResponse
		c().Metadata(ctx, action.MetadataRequest{ProviderTypeName: "scm"}, &m)
		check("action", m.TypeName)
	}
}
