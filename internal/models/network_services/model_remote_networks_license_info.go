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

// Package: network_services
// This file contains models for the network_services SDK package

// LicenseResult represents the Terraform model for LicenseResult
type LicenseResult struct {
	Tfid               types.String          `tfsdk:"tfid"`
	ConfiguredLicenses basetypes.ListValue   `tfsdk:"configured_licenses"`
	LicenseModel       basetypes.ListValue   `tfsdk:"license_model"`
	OperationalLicense basetypes.StringValue `tfsdk:"operational_license"`
	PurchasedLicenses  basetypes.ListValue   `tfsdk:"purchased_licenses"`
}

// LicenseInfo represents a nested structure within the LicenseResult model
type LicenseInfo struct {
	Count       basetypes.Int64Value  `tfsdk:"count"`
	LicenseType basetypes.StringValue `tfsdk:"license_type"`
}

// AttrTypes defines the attribute types for the LicenseResult model.
func (o LicenseResult) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"configured_licenses": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"count":        basetypes.Int64Type{},
				"license_type": basetypes.StringType{},
			},
		}},
		"license_model":       basetypes.ListType{ElemType: basetypes.StringType{}},
		"operational_license": basetypes.StringType{},
		"purchased_licenses": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"count":        basetypes.Int64Type{},
				"license_type": basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of LicenseResult objects.
func (o LicenseResult) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LicenseInfo model.
func (o LicenseInfo) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"count":        basetypes.Int64Type{},
		"license_type": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LicenseInfo objects.
func (o LicenseInfo) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// LicenseResultResourceSchema defines the schema for LicenseResult resource
var LicenseResultResourceSchema = schema.Schema{
	MarkdownDescription: "Manages SCM LicenseResult objects",
	Attributes: map[string]schema.Attribute{
		"configured_licenses": schema.ListNestedAttribute{
			MarkdownDescription: "Configured licenses",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"count": schema.Int64Attribute{
						MarkdownDescription: "Count",
						Optional:            true,
					},
					"license_type": schema.StringAttribute{
						MarkdownDescription: "License type",
						Optional:            true,
					},
				},
			},
		},
		"license_model": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "License model",
			Optional:            true,
		},
		"operational_license": schema.StringAttribute{
			MarkdownDescription: "Operational license",
			Optional:            true,
		},
		"purchased_licenses": schema.ListNestedAttribute{
			MarkdownDescription: "Purchased licenses",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"count": schema.Int64Attribute{
						MarkdownDescription: "Count",
						Optional:            true,
					},
					"license_type": schema.StringAttribute{
						MarkdownDescription: "License type",
						Optional:            true,
					},
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

// LicenseResultDataSourceSchema defines the schema for LicenseResult data source
var LicenseResultDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "LicenseResult data source",
	Attributes: map[string]dsschema.Attribute{
		"configured_licenses": dsschema.ListNestedAttribute{
			MarkdownDescription: "Configured licenses",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"count": dsschema.Int64Attribute{
						MarkdownDescription: "Count",
						Computed:            true,
					},
					"license_type": dsschema.StringAttribute{
						MarkdownDescription: "License type",
						Computed:            true,
					},
				},
			},
		},
		"license_model": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "License model",
			Computed:            true,
		},
		"operational_license": dsschema.StringAttribute{
			MarkdownDescription: "Operational license",
			Computed:            true,
		},
		"purchased_licenses": dsschema.ListNestedAttribute{
			MarkdownDescription: "Purchased licenses",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"count": dsschema.Int64Attribute{
						MarkdownDescription: "Count",
						Computed:            true,
					},
					"license_type": dsschema.StringAttribute{
						MarkdownDescription: "License type",
						Computed:            true,
					},
				},
			},
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// LicenseResultListModel represents the data model for a list data source.
type LicenseResultListModel struct {
	Tfid    types.String    `tfsdk:"tfid"`
	Data    []LicenseResult `tfsdk:"data"`
	Limit   types.Int64     `tfsdk:"limit"`
	Offset  types.Int64     `tfsdk:"offset"`
	Name    types.String    `tfsdk:"name"`
	Total   types.Int64     `tfsdk:"total"`
	Folder  types.String    `tfsdk:"folder"`
	Snippet types.String    `tfsdk:"snippet"`
	Device  types.String    `tfsdk:"device"`
}

// LicenseResultListDataSourceSchema defines the schema for a list data source.
var LicenseResultListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: LicenseResultDataSourceSchema.Attributes,
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
