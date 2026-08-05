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

// Package: ztna_connector_all
// This file contains models for the ztna_connector_all SDK package

// License represents the Terraform model for License
type License struct {
	Tfid            types.String           `tfsdk:"tfid"`
	Applications    basetypes.Float64Value `tfsdk:"applications"`
	Connectors      basetypes.Float64Value `tfsdk:"connectors"`
	Expiry          basetypes.StringValue  `tfsdk:"expiry"`
	LicenseName     basetypes.StringValue  `tfsdk:"license_name"`
	MaxApplications basetypes.Float64Value `tfsdk:"max_applications"`
	MaxConnectors   basetypes.Float64Value `tfsdk:"max_connectors"`
}

// AttrTypes defines the attribute types for the License model.
func (o License) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":             basetypes.StringType{},
		"applications":     basetypes.Float64Type{},
		"connectors":       basetypes.Float64Type{},
		"expiry":           basetypes.StringType{},
		"license_name":     basetypes.StringType{},
		"max_applications": basetypes.Float64Type{},
		"max_connectors":   basetypes.Float64Type{},
	}
}

// AttrType returns the attribute type for a list of License objects.
func (o License) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// LicenseResourceSchema defines the schema for License resource
var LicenseResourceSchema = schema.Schema{
	MarkdownDescription: "License resource",
	Attributes: map[string]schema.Attribute{
		"applications": schema.Float64Attribute{
			MarkdownDescription: "Applications",
			Optional:            true,
			Computed:            true,
		},
		"connectors": schema.Float64Attribute{
			MarkdownDescription: "Connectors",
			Optional:            true,
			Computed:            true,
		},
		"expiry": schema.StringAttribute{
			MarkdownDescription: "Expiry",
			Optional:            true,
			Computed:            true,
		},
		"license_name": schema.StringAttribute{
			MarkdownDescription: "License name",
			Optional:            true,
			Computed:            true,
		},
		"max_applications": schema.Float64Attribute{
			MarkdownDescription: "Max applications",
			Optional:            true,
			Computed:            true,
		},
		"max_connectors": schema.Float64Attribute{
			MarkdownDescription: "Max connectors",
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

// LicenseDataSourceSchema defines the schema for License data source
var LicenseDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "License data source",
	Attributes: map[string]dsschema.Attribute{
		"applications": dsschema.Float64Attribute{
			MarkdownDescription: "Applications",
			Computed:            true,
		},
		"connectors": dsschema.Float64Attribute{
			MarkdownDescription: "Connectors",
			Computed:            true,
		},
		"expiry": dsschema.StringAttribute{
			MarkdownDescription: "Expiry",
			Computed:            true,
		},
		"license_name": dsschema.StringAttribute{
			MarkdownDescription: "License name",
			Computed:            true,
		},
		"max_applications": dsschema.Float64Attribute{
			MarkdownDescription: "Max applications",
			Computed:            true,
		},
		"max_connectors": dsschema.Float64Attribute{
			MarkdownDescription: "Max connectors",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// LicenseListModel represents the data model for a list data source.
type LicenseListModel struct {
	Tfid    types.String `tfsdk:"tfid"`
	Data    []License    `tfsdk:"data"`
	Limit   types.Int64  `tfsdk:"limit"`
	Offset  types.Int64  `tfsdk:"offset"`
	Name    types.String `tfsdk:"name"`
	Total   types.Int64  `tfsdk:"total"`
	Folder  types.String `tfsdk:"folder"`
	Snippet types.String `tfsdk:"snippet"`
	Device  types.String `tfsdk:"device"`
}

// LicenseListDataSourceSchema defines the schema for a list data source.
var LicenseListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: LicenseDataSourceSchema.Attributes,
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
