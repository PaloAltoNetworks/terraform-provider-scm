//go:build !docsgen

package provider

import (
	"context"
	"sort"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// registeredType is a type name and whether its schema is marked deprecated.
type registeredType struct {
	name       string
	deprecated bool
}

func registeredResources(t *testing.T) []registeredType {
	t.Helper()
	ctx := context.Background()
	var out []registeredType
	for _, c := range New("test").(*ScmProvider).Resources(ctx) {
		r := c()
		var m resource.MetadataResponse
		r.Metadata(ctx, resource.MetadataRequest{ProviderTypeName: "scm"}, &m)
		var s resource.SchemaResponse
		r.Schema(ctx, resource.SchemaRequest{}, &s)
		out = append(out, registeredType{m.TypeName, s.Schema.DeprecationMessage != ""})
	}
	return out
}

func registeredDataSources(t *testing.T) []registeredType {
	t.Helper()
	ctx := context.Background()
	var out []registeredType
	for _, c := range New("test").(*ScmProvider).DataSources(ctx) {
		d := c()
		var m datasource.MetadataResponse
		d.Metadata(ctx, datasource.MetadataRequest{ProviderTypeName: "scm"}, &m)
		var s datasource.SchemaResponse
		d.Schema(ctx, datasource.SchemaRequest{}, &s)
		out = append(out, registeredType{m.TypeName, s.Schema.DeprecationMessage != ""})
	}
	return out
}

func registeredActions(t *testing.T) []registeredType {
	t.Helper()
	ctx := context.Background()
	var out []registeredType
	for _, c := range New("test").(*ScmProvider).Actions(ctx) {
		a := c()
		var m action.MetadataResponse
		a.Metadata(ctx, action.MetadataRequest{ProviderTypeName: "scm"}, &m)
		var s action.SchemaResponse
		a.Schema(ctx, action.SchemaRequest{}, &s)
		out = append(out, registeredType{m.TypeName, s.Schema.DeprecationMessage != ""})
	}
	return out
}

// checkNaming enforces the naming rule for one kind of type: names are
// unique, every name starts with scm_ or is a deprecated alias of an existing
// scm_ type, and every such alias is marked deprecated.
func checkNaming(t *testing.T, kind string, types []registeredType) {
	t.Helper()
	byName := map[string]registeredType{}
	for _, ty := range types {
		if _, dup := byName[ty.name]; dup {
			t.Errorf("%s %q is registered more than once", kind, ty.name)
		}
		byName[ty.name] = ty
	}

	var names []string
	for n := range byName {
		names = append(names, n)
	}
	sort.Strings(names)

	for _, n := range names {
		ty := byName[n]
		if strings.HasPrefix(n, scmTypePrefix) {
			if ty.deprecated {
				t.Errorf("%s %q starts with scm_ and must not be deprecated", kind, n)
			}
			continue
		}
		// Not scm_-prefixed: Terraform cannot infer this provider from the
		// name, so it is only allowed as a deprecated alias of a scm_ type.
		if _, ok := byName[scmTypePrefix+n]; !ok {
			t.Errorf("%s %q does not start with scm_ and has no scm_ counterpart %q", kind, n, scmTypePrefix+n)
		}
		if !ty.deprecated {
			t.Errorf("%s %q does not start with scm_, so it must be marked deprecated", kind, n)
		}
	}
}

func TestRegisteredTypeNames(t *testing.T) {
	checkNaming(t, "resource", registeredResources(t))
	checkNaming(t, "data source", registeredDataSources(t))
	checkNaming(t, "action", registeredActions(t))
}

// countUnprefixed returns how many registered types do not start with scm_.
func countUnprefixed(types []registeredType) int {
	n := 0
	for _, ty := range types {
		if !strings.HasPrefix(ty.name, scmTypePrefix) {
			n++
		}
	}
	return n
}

// Only the types that shipped without the scm_ prefix keep a deprecated alias.
// A type added later must be registered as scm_<name> only, so these counts
// must never grow.
func TestNoNewDeprecatedAliases(t *testing.T) {
	for _, c := range []struct {
		kind  string
		types []registeredType
		want  int
	}{
		{"resource", registeredResources(t), 8},
		{"data source", registeredDataSources(t), 29},
		{"action", registeredActions(t), 2},
	} {
		if got := countUnprefixed(c.types); got != c.want {
			t.Errorf("%s: %d types without the scm_ prefix, want %d (new types must not get an unprefixed alias)", c.kind, got, c.want)
		}
	}
}
