//go:build !docsgen

package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// scmTypePrefix is the prefix every canonical type name starts with. The
// deprecated alias of a type is its name without this prefix.
const scmTypePrefix = "scm_"

func legacyName(canonical string) string { return strings.TrimPrefix(canonical, scmTypePrefix) }

// Types that shipped without the scm_ prefix, by canonical (scm_) name. Only
// these keep a deprecated unprefixed alias. This list is frozen: a type added
// later is registered under its scm_ name only.
var (
	legacyResourceNames = map[string]bool{
		"scm_ztna_connector":                         true,
		"scm_ztna_connector_group":                   true,
		"scm_ztna_connector_group_scheduled_upgrade": true,
		"scm_ztna_connector_quiesce":                 true,
		"scm_ztna_connector_scheduled_upgrade":       true,
		"scm_ztna_fqdn_application":                  true,
		"scm_ztna_subnet":                            true,
		"scm_ztna_wildcard":                          true,
	}
	legacyDataSourceNames = map[string]bool{
		"scm_ztna_application_filters":               true,
		"scm_ztna_connector":                         true,
		"scm_ztna_connector_filters":                 true,
		"scm_ztna_connector_group":                   true,
		"scm_ztna_connector_group_connectors":        true,
		"scm_ztna_connector_group_filters":           true,
		"scm_ztna_connector_group_fqdn_rules":        true,
		"scm_ztna_connector_group_list":              true,
		"scm_ztna_connector_group_scheduled_upgrade": true,
		"scm_ztna_connector_group_subnet_rules":      true,
		"scm_ztna_connector_group_upgrade_status":    true,
		"scm_ztna_connector_group_wildcards":         true,
		"scm_ztna_connector_image_list":              true,
		"scm_ztna_connector_list":                    true,
		"scm_ztna_connector_quiesce":                 true,
		"scm_ztna_connector_scheduled_upgrade":       true,
		"scm_ztna_connector_upgrade_status":          true,
		"scm_ztna_discovered_application_filters":    true,
		"scm_ztna_discovered_application_list":       true,
		"scm_ztna_fqdn_application":                  true,
		"scm_ztna_fqdn_application_list":             true,
		"scm_ztna_license":                           true,
		"scm_ztna_subnet":                            true,
		"scm_ztna_subnet_filters":                    true,
		"scm_ztna_subnet_list":                       true,
		"scm_ztna_tenant_status":                     true,
		"scm_ztna_wildcard":                          true,
		"scm_ztna_wildcard_filters":                  true,
		"scm_ztna_wildcard_list":                     true,
	}
	legacyActionNames = map[string]bool{
		"scm_ztna_tenant_start_offboarding": true,
		"scm_ztna_tenant_start_onboarding":  true,
	}
)

func deprecationMessage(canonical string) string {
	return fmt.Sprintf("The type %q is deprecated and will be removed in a future major version. Use %q instead. Existing resources can be moved to the new type with a moved block.", legacyName(canonical), canonical)
}

// ---------------------------------------------------------------------------
// Resources
// ---------------------------------------------------------------------------

type baseResource struct{ inner resource.Resource }

func (b *baseResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	b.inner.Metadata(ctx, req, resp)
}

func (b *baseResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	b.inner.Schema(ctx, req, resp)
}

func (b *baseResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if c, ok := b.inner.(resource.ResourceWithConfigure); ok {
		c.Configure(ctx, req, resp)
	}
}

func (b *baseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	b.inner.Create(ctx, req, resp)
}

func (b *baseResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	b.inner.Read(ctx, req, resp)
}

func (b *baseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	b.inner.Update(ctx, req, resp)
}

func (b *baseResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	b.inner.Delete(ctx, req, resp)
}

func (b *baseResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if i, ok := b.inner.(resource.ResourceWithImportState); ok {
		i.ImportState(ctx, req, resp)
		return
	}
	resp.Diagnostics.AddError("Resource Import Not Implemented", "This resource does not support import. Please contact the provider developer for additional information.")
}

func (b *baseResource) canonicalName(ctx context.Context) string {
	var resp resource.MetadataResponse
	b.inner.Metadata(ctx, resource.MetadataRequest{ProviderTypeName: "scm"}, &resp)
	return resp.TypeName
}

// movableResource is the scm_ resource, able to take over state from its old
// unprefixed name through a moved block. The schemas are identical, so the
// state is carried over as is.
type movableResource struct{ baseResource }

func (r *movableResource) MoveState(ctx context.Context) []resource.StateMover {
	oldName := legacyName(r.canonicalName(ctx))
	var sch resource.SchemaResponse
	r.inner.Schema(ctx, resource.SchemaRequest{}, &sch)
	return []resource.StateMover{{
		SourceSchema: &sch.Schema,
		StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
			if req.SourceTypeName != oldName || !strings.HasSuffix(req.SourceProviderAddress, "/scm") {
				return
			}
			if req.SourceState == nil {
				resp.Diagnostics.AddError("Unable to move state", fmt.Sprintf("The state of %q could not be read.", oldName))
				return
			}
			resp.TargetState = *req.SourceState
		},
	}}
}

// deprecatedResource is the resource under its old unprefixed name.
type deprecatedResource struct{ baseResource }

func (d *deprecatedResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	d.inner.Metadata(ctx, req, resp)
	resp.TypeName = legacyName(resp.TypeName)
}

func (d *deprecatedResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	d.inner.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = deprecationMessage(d.canonicalName(ctx))
}

var (
	_ resource.ResourceWithConfigure   = &movableResource{}
	_ resource.ResourceWithImportState = &movableResource{}
	_ resource.ResourceWithMoveState   = &movableResource{}
	_ resource.ResourceWithConfigure   = &deprecatedResource{}
	_ resource.ResourceWithImportState = &deprecatedResource{}
)

// legacyResources returns, for every constructor, the scm_ resource and the
// same resource under its old unprefixed (deprecated) name.
func legacyResources(ctors []func() resource.Resource) []func() resource.Resource {
	out := make([]func() resource.Resource, 0, 2*len(ctors))
	for _, c := range ctors {
		c := c
		if !legacyResourceNames[(&baseResource{inner: c()}).canonicalName(context.Background())] {
			out = append(out, c)
			continue
		}
		out = append(out,
			func() resource.Resource { return &movableResource{baseResource{inner: c()}} },
			func() resource.Resource { return &deprecatedResource{baseResource{inner: c()}} },
		)
	}
	return out
}

// ---------------------------------------------------------------------------
// Data sources
// ---------------------------------------------------------------------------

type deprecatedDataSource struct{ inner datasource.DataSource }

func (d *deprecatedDataSource) canonicalName(ctx context.Context) string {
	var resp datasource.MetadataResponse
	d.inner.Metadata(ctx, datasource.MetadataRequest{ProviderTypeName: "scm"}, &resp)
	return resp.TypeName
}

func (d *deprecatedDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	d.inner.Metadata(ctx, req, resp)
	resp.TypeName = legacyName(resp.TypeName)
}

func (d *deprecatedDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	d.inner.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = deprecationMessage(d.canonicalName(ctx))
}

func (d *deprecatedDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if c, ok := d.inner.(datasource.DataSourceWithConfigure); ok {
		c.Configure(ctx, req, resp)
	}
}

func (d *deprecatedDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.inner.Read(ctx, req, resp)
}

var _ datasource.DataSourceWithConfigure = &deprecatedDataSource{}

// legacyDataSources returns, for every constructor, the scm_ data source and
// the same data source under its old unprefixed (deprecated) name.
func legacyDataSources(ctors []func() datasource.DataSource) []func() datasource.DataSource {
	out := make([]func() datasource.DataSource, 0, 2*len(ctors))
	for _, c := range ctors {
		c := c
		if !legacyDataSourceNames[(&deprecatedDataSource{inner: c()}).canonicalName(context.Background())] {
			out = append(out, c)
			continue
		}
		out = append(out, c, func() datasource.DataSource { return &deprecatedDataSource{inner: c()} })
	}
	return out
}

// ---------------------------------------------------------------------------
// Actions
// ---------------------------------------------------------------------------

type deprecatedAction struct{ inner action.Action }

func (d *deprecatedAction) canonicalName(ctx context.Context) string {
	var resp action.MetadataResponse
	d.inner.Metadata(ctx, action.MetadataRequest{ProviderTypeName: "scm"}, &resp)
	return resp.TypeName
}

func (d *deprecatedAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	d.inner.Metadata(ctx, req, resp)
	resp.TypeName = legacyName(resp.TypeName)
}

func (d *deprecatedAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	d.inner.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = deprecationMessage(d.canonicalName(ctx))
}

func (d *deprecatedAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if c, ok := d.inner.(action.ActionWithConfigure); ok {
		c.Configure(ctx, req, resp)
	}
}

func (d *deprecatedAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	d.inner.Invoke(ctx, req, resp)
}

var _ action.ActionWithConfigure = &deprecatedAction{}

// legacyActions returns, for every constructor, the scm_ action and the same
// action under its old unprefixed (deprecated) name.
func legacyActions(ctors []func() action.Action) []func() action.Action {
	out := make([]func() action.Action, 0, 2*len(ctors))
	for _, c := range ctors {
		c := c
		if !legacyActionNames[(&deprecatedAction{inner: c()}).canonicalName(context.Background())] {
			out = append(out, c)
			continue
		}
		out = append(out, c, func() action.Action { return &deprecatedAction{inner: c()} })
	}
	return out
}
