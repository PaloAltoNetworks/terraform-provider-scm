package models

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: config_setup
// This file contains models for the config_setup SDK package

// TrustedTenantOverview represents the Terraform model for TrustedTenantOverview
type TrustedTenantOverview struct {
	Tfid       types.String          `tfsdk:"tfid"`
	Publisher  basetypes.ObjectValue `tfsdk:"publisher"`
	Subscriber basetypes.ObjectValue `tfsdk:"subscriber"`
}

// TrustedTenantOverviewPublisher represents a nested structure within the TrustedTenantOverview model
type TrustedTenantOverviewPublisher struct {
	Pending basetypes.Int64Value `tfsdk:"pending"`
	Total   basetypes.Int64Value `tfsdk:"total"`
}

// AttrTypes defines the attribute types for the TrustedTenantOverview model.
func (o TrustedTenantOverview) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"publisher": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"pending": basetypes.Int64Type{},
				"total":   basetypes.Int64Type{},
			},
		},
		"subscriber": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"pending": basetypes.Int64Type{},
				"total":   basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of TrustedTenantOverview objects.
func (o TrustedTenantOverview) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the TrustedTenantOverviewPublisher model.
func (o TrustedTenantOverviewPublisher) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"pending": basetypes.Int64Type{},
		"total":   basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of TrustedTenantOverviewPublisher objects.
func (o TrustedTenantOverviewPublisher) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// TrustedTenantOverviewResourceSchema defines the schema for TrustedTenantOverview resource
var TrustedTenantOverviewResourceSchema = schema.Schema{
	MarkdownDescription: "TrustedTenantOverview resource",
	Attributes: map[string]schema.Attribute{
		"publisher": schema.SingleNestedAttribute{
			MarkdownDescription: "Publisher",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"pending": schema.Int64Attribute{
					MarkdownDescription: "Pending",
					Computed:            true,
				},
				"total": schema.Int64Attribute{
					MarkdownDescription: "Total",
					Computed:            true,
				},
			},
		},
		"subscriber": schema.SingleNestedAttribute{
			MarkdownDescription: "Subscriber",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"pending": schema.Int64Attribute{
					MarkdownDescription: "Pending",
					Computed:            true,
				},
				"total": schema.Int64Attribute{
					MarkdownDescription: "Total",
					Computed:            true,
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
	},
}

// TrustedTenantOverviewDataSourceSchema defines the schema for TrustedTenantOverview data source
var TrustedTenantOverviewDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "TrustedTenantOverview data source",
	Attributes: map[string]dsschema.Attribute{
		"publisher": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Publisher",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"pending": dsschema.Int64Attribute{
					MarkdownDescription: "Pending",
					Computed:            true,
				},
				"total": dsschema.Int64Attribute{
					MarkdownDescription: "Total",
					Computed:            true,
				},
			},
		},
		"subscriber": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Subscriber",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"pending": dsschema.Int64Attribute{
					MarkdownDescription: "Pending",
					Computed:            true,
				},
				"total": dsschema.Int64Attribute{
					MarkdownDescription: "Total",
					Computed:            true,
				},
			},
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// TrustedTenantOverviewListModel represents the data model for a list data source.
type TrustedTenantOverviewListModel struct {
	Tfid    types.String            `tfsdk:"tfid"`
	Data    []TrustedTenantOverview `tfsdk:"data"`
	Limit   types.Int64             `tfsdk:"limit"`
	Offset  types.Int64             `tfsdk:"offset"`
	Name    types.String            `tfsdk:"name"`
	Total   types.Int64             `tfsdk:"total"`
	Folder  types.String            `tfsdk:"folder"`
	Snippet types.String            `tfsdk:"snippet"`
	Device  types.String            `tfsdk:"device"`
}

// TrustedTenantOverviewListDataSourceSchema defines the schema for a list data source.
var TrustedTenantOverviewListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: TrustedTenantOverviewDataSourceSchema.Attributes,
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
