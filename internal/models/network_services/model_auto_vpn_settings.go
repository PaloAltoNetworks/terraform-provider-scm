package models

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: network_services
// This file contains models for the network_services SDK package

// AutoVpnSettings represents the Terraform model for AutoVpnSettings
type AutoVpnSettings struct {
	Tfid                  types.String          `tfsdk:"tfid"`
	AsRange               basetypes.ObjectValue `tfsdk:"as_range"`
	EnableMeshBetweenHubs basetypes.BoolValue   `tfsdk:"enable_mesh_between_hubs"`
	VpnAddressPool        basetypes.ListValue   `tfsdk:"vpn_address_pool"`
}

// AutoVpnSettingsAsRange represents a nested structure within the AutoVpnSettings model
type AutoVpnSettingsAsRange struct {
	End   basetypes.Int64Value `tfsdk:"end"`
	Start basetypes.Int64Value `tfsdk:"start"`
}

// AttrTypes defines the attribute types for the AutoVpnSettings model.
func (o AutoVpnSettings) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"as_range": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"end":   basetypes.Int64Type{},
				"start": basetypes.Int64Type{},
			},
		},
		"enable_mesh_between_hubs": basetypes.BoolType{},
		"vpn_address_pool":         basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of AutoVpnSettings objects.
func (o AutoVpnSettings) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AutoVpnSettingsAsRange model.
func (o AutoVpnSettingsAsRange) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"end":   basetypes.Int64Type{},
		"start": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of AutoVpnSettingsAsRange objects.
func (o AutoVpnSettingsAsRange) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AutoVpnSettingsResourceSchema defines the schema for AutoVpnSettings resource
var AutoVpnSettingsResourceSchema = schema.Schema{
	MarkdownDescription: "AutoVpnSetting resource",
	Attributes: map[string]schema.Attribute{
		"as_range": schema.SingleNestedAttribute{
			MarkdownDescription: "As range",
			Required:            true,
			Attributes: map[string]schema.Attribute{
				"end": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(1, 65535),
					},
					MarkdownDescription: "End",
					Optional:            true,
				},
				"start": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(1, 65535),
					},
					MarkdownDescription: "Start",
					Optional:            true,
				},
			},
		},
		"enable_mesh_between_hubs": schema.BoolAttribute{
			MarkdownDescription: "Enable mesh connection between hubs?",
			Optional:            true,
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"vpn_address_pool": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "VPN address pool",
			Required:            true,
		},
	},
}

// AutoVpnSettingsDataSourceSchema defines the schema for AutoVpnSettings data source
var AutoVpnSettingsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "AutoVpnSetting data source",
	Attributes: map[string]dsschema.Attribute{
		"as_range": dsschema.SingleNestedAttribute{
			MarkdownDescription: "As range",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"end": dsschema.Int64Attribute{
					MarkdownDescription: "End",
					Computed:            true,
				},
				"start": dsschema.Int64Attribute{
					MarkdownDescription: "Start",
					Computed:            true,
				},
			},
		},
		"enable_mesh_between_hubs": dsschema.BoolAttribute{
			MarkdownDescription: "Enable mesh connection between hubs?",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"vpn_address_pool": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "VPN address pool",
			Computed:            true,
		},
	},
}

// AutoVpnSettingsListModel represents the data model for a list data source.
type AutoVpnSettingsListModel struct {
	Tfid    types.String      `tfsdk:"tfid"`
	Data    []AutoVpnSettings `tfsdk:"data"`
	Limit   types.Int64       `tfsdk:"limit"`
	Offset  types.Int64       `tfsdk:"offset"`
	Name    types.String      `tfsdk:"name"`
	Total   types.Int64       `tfsdk:"total"`
	Folder  types.String      `tfsdk:"folder"`
	Snippet types.String      `tfsdk:"snippet"`
	Device  types.String      `tfsdk:"device"`
}

// AutoVpnSettingsListDataSourceSchema defines the schema for a list data source.
var AutoVpnSettingsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: AutoVpnSettingsDataSourceSchema.Attributes,
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
