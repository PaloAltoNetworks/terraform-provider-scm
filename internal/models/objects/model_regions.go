package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
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

// Package: objects
// This file contains models for the objects SDK package

// Regions represents the Terraform model for Regions
type Regions struct {
	Tfid        types.String          `tfsdk:"tfid"`
	Address     basetypes.ListValue   `tfsdk:"address"`
	Device      basetypes.StringValue `tfsdk:"device"`
	Folder      basetypes.StringValue `tfsdk:"folder"`
	GeoLocation basetypes.ObjectValue `tfsdk:"geo_location"`
	Id          basetypes.StringValue `tfsdk:"id"`
	Name        basetypes.StringValue `tfsdk:"name"`
	Snippet     basetypes.StringValue `tfsdk:"snippet"`
}

// RegionsGeoLocation represents a nested structure within the Regions model
type RegionsGeoLocation struct {
	Latitude  basetypes.Float64Value `tfsdk:"latitude"`
	Longitude basetypes.Float64Value `tfsdk:"longitude"`
}

// AttrTypes defines the attribute types for the Regions model.
func (o Regions) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":    basetypes.StringType{},
		"address": basetypes.ListType{ElemType: basetypes.StringType{}},
		"device":  basetypes.StringType{},
		"folder":  basetypes.StringType{},
		"geo_location": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"latitude":  basetypes.Float64Type{},
				"longitude": basetypes.Float64Type{},
			},
		},
		"id":      basetypes.StringType{},
		"name":    basetypes.StringType{},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of Regions objects.
func (o Regions) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RegionsGeoLocation model.
func (o RegionsGeoLocation) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"latitude":  basetypes.Float64Type{},
		"longitude": basetypes.Float64Type{},
	}
}

// AttrType returns the attribute type for a list of RegionsGeoLocation objects.
func (o RegionsGeoLocation) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// RegionsResourceSchema defines the schema for Regions resource
var RegionsResourceSchema = schema.Schema{
	MarkdownDescription: "Region resource",
	Attributes: map[string]schema.Attribute{
		"address": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Address",
			Optional:            true,
		},
		"device": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("folder"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
			},
			MarkdownDescription: "The device in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"folder": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("device"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
			},
			MarkdownDescription: "The folder in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"geo_location": schema.SingleNestedAttribute{
			MarkdownDescription: "Geo location",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"latitude": schema.Float64Attribute{
					Validators: []validator.Float64{
						float64validator.Between(-90.000000, 90.000000),
					},
					MarkdownDescription: "The latitudinal position of the region",
					Required:            true,
				},
				"longitude": schema.Float64Attribute{
					Validators: []validator.Float64{
						float64validator.Between(-180.000000, 180.000000),
					},
					MarkdownDescription: "The longitudinal postition of the region",
					Required:            true,
				},
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the region",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(31),
				stringvalidator.RegexMatches(regexp.MustCompile("^[ a-zA-Z\\d._-]+$"), "pattern must match "+"^[ a-zA-Z\\d._-]+$"),
			},
			MarkdownDescription: "The name of the region",
			Required:            true,
		},
		"snippet": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("device"),
					path.MatchRelative().AtParent().AtName("folder"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
			},
			MarkdownDescription: "The snippet in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
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

// RegionsDataSourceSchema defines the schema for Regions data source
var RegionsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Region data source",
	Attributes: map[string]dsschema.Attribute{
		"address": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Address",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"geo_location": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Geo location",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"latitude": dsschema.Float64Attribute{
					MarkdownDescription: "The latitudinal position of the region",
					Computed:            true,
				},
				"longitude": dsschema.Float64Attribute{
					MarkdownDescription: "The longitudinal postition of the region",
					Computed:            true,
				},
			},
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the region",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the region",
			Optional:            true,
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// RegionsListModel represents the data model for a list data source.
type RegionsListModel struct {
	Tfid    types.String `tfsdk:"tfid"`
	Data    []Regions    `tfsdk:"data"`
	Limit   types.Int64  `tfsdk:"limit"`
	Offset  types.Int64  `tfsdk:"offset"`
	Name    types.String `tfsdk:"name"`
	Total   types.Int64  `tfsdk:"total"`
	Folder  types.String `tfsdk:"folder"`
	Snippet types.String `tfsdk:"snippet"`
	Device  types.String `tfsdk:"device"`
}

// RegionsListDataSourceSchema defines the schema for a list data source.
var RegionsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: RegionsDataSourceSchema.Attributes,
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
