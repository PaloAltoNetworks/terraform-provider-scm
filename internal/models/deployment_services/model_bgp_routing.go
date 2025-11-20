package models

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: deployment_services
// This file contains models for the deployment_services SDK package

// BgpRouting represents the Terraform model for BgpRouting
type BgpRouting struct {
	Tfid                      types.String          `tfsdk:"tfid"`
	AcceptRouteOverSC         basetypes.BoolValue   `tfsdk:"accept_route_over_sc"`
	AddHostRouteToIkePeer     basetypes.BoolValue   `tfsdk:"add_host_route_to_ike_peer"`
	BackboneRouting           basetypes.StringValue `tfsdk:"backbone_routing"`
	OutboundRoutesForServices basetypes.ListValue   `tfsdk:"outbound_routes_for_services"`
	RoutingPreference         basetypes.ObjectValue `tfsdk:"routing_preference"`
	WithdrawStaticRoute       basetypes.BoolValue   `tfsdk:"withdraw_static_route"`
}

// BgpRoutingRoutingPreference represents a nested structure within the BgpRouting model
type BgpRoutingRoutingPreference struct {
	Default          basetypes.ObjectValue `tfsdk:"default"`
	HotPotatoRouting basetypes.ObjectValue `tfsdk:"hot_potato_routing"`
}

// AttrTypes defines the attribute types for the BgpRouting model.
func (o BgpRouting) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":                         basetypes.StringType{},
		"accept_route_over_sc":         basetypes.BoolType{},
		"add_host_route_to_ike_peer":   basetypes.BoolType{},
		"backbone_routing":             basetypes.StringType{},
		"outbound_routes_for_services": basetypes.ListType{ElemType: basetypes.StringType{}},
		"routing_preference": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"default": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"hot_potato_routing": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
		"withdraw_static_route": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouting objects.
func (o BgpRouting) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRoutingRoutingPreference model.
func (o BgpRoutingRoutingPreference) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"default": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"hot_potato_routing": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of BgpRoutingRoutingPreference objects.
func (o BgpRoutingRoutingPreference) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// BgpRoutingResourceSchema defines the schema for BgpRouting resource
var BgpRoutingResourceSchema = schema.Schema{
	MarkdownDescription: "BgpRouting resource",
	Attributes: map[string]schema.Attribute{
		"accept_route_over_sc": schema.BoolAttribute{
			MarkdownDescription: "Accept route over s c",
			Optional:            true,
		},
		"add_host_route_to_ike_peer": schema.BoolAttribute{
			MarkdownDescription: "Add host route to ike peer",
			Optional:            true,
		},
		"backbone_routing": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("no-asymmetric-routing", "asymmetric-routing-only", "asymmetric-routing-with-load-share"),
			},
			MarkdownDescription: "Backbone routing",
			Optional:            true,
		},
		"outbound_routes_for_services": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Outbound routes for services",
			Optional:            true,
		},
		"routing_preference": schema.SingleNestedAttribute{
			MarkdownDescription: "Routing preference",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"default": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("hot_potato_routing"),
						),
					},
					MarkdownDescription: "Default\n\n> ℹ️ **Note:** You must specify exactly one of `default` and `hot_potato_routing`.",
					Optional:            true,
					Attributes:          map[string]schema.Attribute{},
				},
				"hot_potato_routing": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("default"),
						),
					},
					MarkdownDescription: "Hot potato routing\n\n> ℹ️ **Note:** You must specify exactly one of `default` and `hot_potato_routing`.",
					Optional:            true,
					Attributes:          map[string]schema.Attribute{},
				},
			},
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"withdraw_static_route": schema.BoolAttribute{
			MarkdownDescription: "Withdraw static route",
			Optional:            true,
		},
	},
}

// BgpRoutingDataSourceSchema defines the schema for BgpRouting data source
var BgpRoutingDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "BgpRouting data source",
	Attributes: map[string]dsschema.Attribute{
		"accept_route_over_sc": dsschema.BoolAttribute{
			MarkdownDescription: "Accept route over s c",
			Computed:            true,
		},
		"add_host_route_to_ike_peer": dsschema.BoolAttribute{
			MarkdownDescription: "Add host route to ike peer",
			Computed:            true,
		},
		"backbone_routing": dsschema.StringAttribute{
			MarkdownDescription: "Backbone routing",
			Computed:            true,
		},
		"outbound_routes_for_services": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Outbound routes for services",
			Computed:            true,
		},
		"routing_preference": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Routing preference",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"default": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Default\n\n> ℹ️ **Note:** You must specify exactly one of `default` and `hot_potato_routing`.",
					Computed:            true,
					Attributes:          map[string]dsschema.Attribute{},
				},
				"hot_potato_routing": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Hot potato routing\n\n> ℹ️ **Note:** You must specify exactly one of `default` and `hot_potato_routing`.",
					Computed:            true,
					Attributes:          map[string]dsschema.Attribute{},
				},
			},
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"withdraw_static_route": dsschema.BoolAttribute{
			MarkdownDescription: "Withdraw static route",
			Computed:            true,
		},
	},
}

// BgpRoutingListModel represents the data model for a list data source.
type BgpRoutingListModel struct {
	Tfid    types.String `tfsdk:"tfid"`
	Data    []BgpRouting `tfsdk:"data"`
	Limit   types.Int64  `tfsdk:"limit"`
	Offset  types.Int64  `tfsdk:"offset"`
	Name    types.String `tfsdk:"name"`
	Total   types.Int64  `tfsdk:"total"`
	Folder  types.String `tfsdk:"folder"`
	Snippet types.String `tfsdk:"snippet"`
	Device  types.String `tfsdk:"device"`
}

// BgpRoutingListDataSourceSchema defines the schema for a list data source.
var BgpRoutingListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: BgpRoutingDataSourceSchema.Attributes,
			},
		},
		"limit":   dsschema.Int64Attribute{Description: "The max number of items to return. Default: 200.", Optional: true},
		"offset":  dsschema.Int64Attribute{Description: "The offset of the first item to return.", Optional: true},
		"name":    dsschema.StringAttribute{Description: "The name of the item.", Optional: true},
		"total":   dsschema.Int64Attribute{Description: "The total number of items.", Computed: true},
		"folder":  dsschema.StringAttribute{Description: "The folder of the item. Default: Shared.", Optional: true},
		"snippet": dsschema.StringAttribute{Description: "The snippet of the item.", Optional: true},
		"device":  dsschema.StringAttribute{Description: "The device of the item.", Optional: true},
	},
}
