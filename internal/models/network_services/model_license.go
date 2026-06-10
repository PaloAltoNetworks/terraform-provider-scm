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

// LicenseInfo represents the Terraform model for LicenseInfo
type LicenseInfo struct {
	Tfid        types.String          `tfsdk:"tfid"`
	Count       basetypes.Int64Value  `tfsdk:"count"`
	LicenseType basetypes.StringValue `tfsdk:"license_type"`
}

// AttrTypes defines the attribute types for the LicenseInfo model.
func (o LicenseInfo) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":         basetypes.StringType{},
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

// LicenseInfoResourceSchema defines the schema for LicenseInfo resource
var LicenseInfoResourceSchema = schema.Schema{
	MarkdownDescription: "LicenseInfo resource",
	Attributes: map[string]schema.Attribute{
		"count": schema.Int64Attribute{
			MarkdownDescription: "Count",
			Optional:            true,
			Computed:            true,
		},
		"license_type": schema.StringAttribute{
			MarkdownDescription: "License type",
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

// LicenseInfoDataSourceSchema defines the schema for LicenseInfo data source
var LicenseInfoDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "LicenseInfo data source",
	Attributes: map[string]dsschema.Attribute{
		"count": dsschema.Int64Attribute{
			MarkdownDescription: "Count",
			Computed:            true,
		},
		"license_type": dsschema.StringAttribute{
			MarkdownDescription: "License type",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// LicenseInfoListModel represents the data model for a list data source.
type LicenseInfoListModel struct {
	Tfid    types.String  `tfsdk:"tfid"`
	Data    []LicenseInfo `tfsdk:"data"`
	Limit   types.Int64   `tfsdk:"limit"`
	Offset  types.Int64   `tfsdk:"offset"`
	Name    types.String  `tfsdk:"name"`
	Total   types.Int64   `tfsdk:"total"`
	Folder  types.String  `tfsdk:"folder"`
	Snippet types.String  `tfsdk:"snippet"`
	Device  types.String  `tfsdk:"device"`
}

// LicenseInfoListDataSourceSchema defines the schema for a list data source.
var LicenseInfoListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: LicenseInfoDataSourceSchema.Attributes,
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
