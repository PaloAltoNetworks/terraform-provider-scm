package models

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: deployment_services
// This file contains models for the deployment_services SDK package

// BandwidthAllocations represents the Terraform model for BandwidthAllocations
type BandwidthAllocations struct {
	Tfid               types.String          `tfsdk:"tfid"`
	AllocatedBandwidth basetypes.Int64Value  `tfsdk:"allocated_bandwidth"`
	Name               basetypes.StringValue `tfsdk:"name"`
	Qos                basetypes.ObjectValue `tfsdk:"qos"`
	SpnNameList        basetypes.ListValue   `tfsdk:"spn_name_list"`
}

// BandwidthAllocationsQos represents a nested structure within the BandwidthAllocations model
type BandwidthAllocationsQos struct {
	Customized      basetypes.BoolValue    `tfsdk:"customized"`
	Enabled         basetypes.BoolValue    `tfsdk:"enabled"`
	GuaranteedRatio basetypes.Float64Value `tfsdk:"guaranteed_ratio"`
	Profile         basetypes.StringValue  `tfsdk:"profile"`
}

// AttrTypes defines the attribute types for the BandwidthAllocations model.
func (o BandwidthAllocations) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":                basetypes.StringType{},
		"allocated_bandwidth": basetypes.Int64Type{},
		"name":                basetypes.StringType{},
		"qos": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"customized":       basetypes.BoolType{},
				"enabled":          basetypes.BoolType{},
				"guaranteed_ratio": basetypes.Float64Type{},
				"profile":          basetypes.StringType{},
			},
		},
		"spn_name_list": basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of BandwidthAllocations objects.
func (o BandwidthAllocations) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BandwidthAllocationsQos model.
func (o BandwidthAllocationsQos) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"customized":       basetypes.BoolType{},
		"enabled":          basetypes.BoolType{},
		"guaranteed_ratio": basetypes.Float64Type{},
		"profile":          basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BandwidthAllocationsQos objects.
func (o BandwidthAllocationsQos) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// BandwidthAllocationsResourceSchema defines the schema for BandwidthAllocations resource
var BandwidthAllocationsResourceSchema = schema.Schema{
	MarkdownDescription: "BandwidthAllocation resource",
	Attributes: map[string]schema.Attribute{
		"allocated_bandwidth": schema.Int64Attribute{
			MarkdownDescription: "bandwidth to allocate in Mbps",
			Required:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "name of the aggregated bandwidth region",
			Required:            true,
		},
		"qos": schema.SingleNestedAttribute{
			MarkdownDescription: "Qos",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"customized": schema.BoolAttribute{
					MarkdownDescription: "Customized",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"enabled": schema.BoolAttribute{
					MarkdownDescription: "Enabled",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"guaranteed_ratio": schema.Float64Attribute{
					MarkdownDescription: "Guaranteed ratio",
					Optional:            true,
					Computed:            true,
					Default:             float64default.StaticFloat64(0.000000),
				},
				"profile": schema.StringAttribute{
					MarkdownDescription: "Profile",
					Optional:            true,
					Computed:            true,
					Default:             stringdefault.StaticString(""),
				},
			},
		},
		"spn_name_list": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Spn name list",
			Optional:            true,
			Computed:            true,
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
	},
}

// BandwidthAllocationsDataSourceSchema defines the schema for BandwidthAllocations data source
var BandwidthAllocationsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "BandwidthAllocation data source",
	Attributes: map[string]dsschema.Attribute{
		"allocated_bandwidth": dsschema.Int64Attribute{
			MarkdownDescription: "bandwidth to allocate in Mbps",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "name of the aggregated bandwidth region",
			Optional:            true,
			Computed:            true,
		},
		"qos": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Qos",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"customized": dsschema.BoolAttribute{
					MarkdownDescription: "Customized",
					Computed:            true,
				},
				"enabled": dsschema.BoolAttribute{
					MarkdownDescription: "Enabled",
					Computed:            true,
				},
				"guaranteed_ratio": dsschema.Float64Attribute{
					MarkdownDescription: "Guaranteed ratio",
					Computed:            true,
				},
				"profile": dsschema.StringAttribute{
					MarkdownDescription: "Profile",
					Computed:            true,
				},
			},
		},
		"spn_name_list": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Spn name list",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// BandwidthAllocationsListModel represents the data model for a list data source.
type BandwidthAllocationsListModel struct {
	Tfid    types.String           `tfsdk:"tfid"`
	Data    []BandwidthAllocations `tfsdk:"data"`
	Limit   types.Int64            `tfsdk:"limit"`
	Offset  types.Int64            `tfsdk:"offset"`
	Name    types.String           `tfsdk:"name"`
	Total   types.Int64            `tfsdk:"total"`
	Folder  types.String           `tfsdk:"folder"`
	Snippet types.String           `tfsdk:"snippet"`
	Device  types.String           `tfsdk:"device"`
}

// BandwidthAllocationsListDataSourceSchema defines the schema for a list data source.
var BandwidthAllocationsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: BandwidthAllocationsDataSourceSchema.Attributes,
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
