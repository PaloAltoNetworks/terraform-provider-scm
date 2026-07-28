package models

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: deployment_services
// This file contains models for the deployment_services SDK package

// Locations represents the Terraform model for Locations
type Locations struct {
	Tfid            types.String           `tfsdk:"tfid"`
	AggregateRegion basetypes.StringValue  `tfsdk:"aggregate_region"`
	Continent       basetypes.StringValue  `tfsdk:"continent"`
	Display         basetypes.StringValue  `tfsdk:"display"`
	Latitude        basetypes.Float64Value `tfsdk:"latitude"`
	Longitude       basetypes.Float64Value `tfsdk:"longitude"`
	Region          basetypes.StringValue  `tfsdk:"region"`
	Value           basetypes.StringValue  `tfsdk:"value"`
}

// AttrTypes defines the attribute types for the Locations model.
func (o Locations) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":             basetypes.StringType{},
		"aggregate_region": basetypes.StringType{},
		"continent":        basetypes.StringType{},
		"display":          basetypes.StringType{},
		"latitude":         basetypes.Float64Type{},
		"longitude":        basetypes.Float64Type{},
		"region":           basetypes.StringType{},
		"value":            basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of Locations objects.
func (o Locations) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// LocationsResourceSchema defines the schema for Locations resource
var LocationsResourceSchema = schema.Schema{
	MarkdownDescription: "Location resource",
	Attributes: map[string]schema.Attribute{
		"aggregate_region": schema.StringAttribute{
			MarkdownDescription: "Aggregate region",
			Optional:            true,
		},
		"continent": schema.StringAttribute{
			MarkdownDescription: "The continent in which the location exists",
			Optional:            true,
		},
		"display": schema.StringAttribute{
			MarkdownDescription: "The location as displayed in the Strata Cloud Manager portal",
			Optional:            true,
		},
		"latitude": schema.Float64Attribute{
			Validators: []validator.Float64{
				float64validator.Between(-90.000000, 90.000000),
			},
			MarkdownDescription: "The latitudinal position of the location",
			Optional:            true,
		},
		"longitude": schema.Float64Attribute{
			Validators: []validator.Float64{
				float64validator.Between(-180.000000, 180.000000),
			},
			MarkdownDescription: "The longitudinal position of the location",
			Optional:            true,
		},
		"region": schema.StringAttribute{
			MarkdownDescription: "Region",
			Optional:            true,
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"value": schema.StringAttribute{
			MarkdownDescription: "Value",
			Optional:            true,
		},
	},
}

// LocationsDataSourceSchema defines the schema for Locations data source
var LocationsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Location data source",
	Attributes: map[string]dsschema.Attribute{
		"aggregate_region": dsschema.StringAttribute{
			MarkdownDescription: "Aggregate region",
			Computed:            true,
		},
		"continent": dsschema.StringAttribute{
			MarkdownDescription: "The continent in which the location exists",
			Computed:            true,
		},
		"display": dsschema.StringAttribute{
			MarkdownDescription: "The location as displayed in the Strata Cloud Manager portal",
			Computed:            true,
		},
		"latitude": dsschema.Float64Attribute{
			MarkdownDescription: "The latitudinal position of the location",
			Computed:            true,
		},
		"longitude": dsschema.Float64Attribute{
			MarkdownDescription: "The longitudinal position of the location",
			Computed:            true,
		},
		"region": dsschema.StringAttribute{
			MarkdownDescription: "Region",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"value": dsschema.StringAttribute{
			MarkdownDescription: "Value",
			Computed:            true,
		},
	},
}

// LocationsListModel represents the data model for a list data source.
type LocationsListModel struct {
	Tfid    types.String `tfsdk:"tfid"`
	Data    []Locations  `tfsdk:"data"`
	Limit   types.Int64  `tfsdk:"limit"`
	Offset  types.Int64  `tfsdk:"offset"`
	Name    types.String `tfsdk:"name"`
	Total   types.Int64  `tfsdk:"total"`
	Folder  types.String `tfsdk:"folder"`
	Snippet types.String `tfsdk:"snippet"`
	Device  types.String `tfsdk:"device"`
}

// LocationsListDataSourceSchema defines the schema for a list data source.
var LocationsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: LocationsDataSourceSchema.Attributes,
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
