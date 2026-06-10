package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: ztna_connector_all
// This file contains models for the ztna_connector_all SDK package

// Connectors represents the Terraform model for Connectors
type Connectors struct {
	Tfid        types.String          `tfsdk:"tfid"`
	CreatedTime basetypes.StringValue `tfsdk:"created_time"`
	Description basetypes.StringValue `tfsdk:"description"`
	Group       basetypes.StringValue `tfsdk:"group"`
	Name        basetypes.StringValue `tfsdk:"name"`
	Oid         basetypes.StringValue `tfsdk:"oid"`
	UpdatedTime basetypes.StringValue `tfsdk:"updated_time"`
}

// AttrTypes defines the attribute types for the Connectors model.
func (o Connectors) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":         basetypes.StringType{},
		"created_time": basetypes.StringType{},
		"description":  basetypes.StringType{},
		"group":        basetypes.StringType{},
		"name":         basetypes.StringType{},
		"oid":          basetypes.StringType{},
		"updated_time": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of Connectors objects.
func (o Connectors) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// ConnectorsResourceSchema defines the schema for Connectors resource
var ConnectorsResourceSchema = schema.Schema{
	MarkdownDescription: "Connector resource",
	Attributes: map[string]schema.Attribute{
		"created_time": schema.StringAttribute{
			MarkdownDescription: "Created time",
			Computed:            true,
		},
		"description": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[\\p{L}\\p{N}\\p{P}\\s,.:_-]*$"), "pattern must match "+"^[\\p{L}\\p{N}\\p{P}\\s,.:_-]*$"),
			},
			MarkdownDescription: "Description of the connector.",
			Optional:            true,
			Computed:            true,
		},
		"group": schema.StringAttribute{
			MarkdownDescription: "The connector group id",
			Required:            true,
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(64),
				stringvalidator.LengthAtLeast(1),
				stringvalidator.RegexMatches(regexp.MustCompile("^[\\p{L}\\p{N}\\p{P}\\s,.:_-]*$"), "pattern must match "+"^[\\p{L}\\p{N}\\p{P}\\s,.:_-]*$"),
			},
			MarkdownDescription: "Name of the connector.\n\nIt can only be 64 characters long\nand contain unicode text, space, dash, or underscore, or period.",
			Required:            true,
		},
		"oid": schema.StringAttribute{
			MarkdownDescription: "Id of the entry.",
			Computed:            true,
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"updated_time": schema.StringAttribute{
			MarkdownDescription: "Updated time",
			Computed:            true,
		},
	},
}

// ConnectorsDataSourceSchema defines the schema for Connectors data source
var ConnectorsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Connector data source",
	Attributes: map[string]dsschema.Attribute{
		"created_time": dsschema.StringAttribute{
			MarkdownDescription: "Created time",
			Computed:            true,
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Description of the connector.",
			Computed:            true,
		},
		"group": dsschema.StringAttribute{
			MarkdownDescription: "The connector group id",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Name of the connector.\n\nIt can only be 64 characters long\nand contain unicode text, space, dash, or underscore, or period.",
			Optional:            true,
			Computed:            true,
		},
		"oid": dsschema.StringAttribute{
			MarkdownDescription: "Id of the entry.",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"updated_time": dsschema.StringAttribute{
			MarkdownDescription: "Updated time",
			Computed:            true,
		},
	},
}

// ConnectorsListModel represents the data model for a list data source.
type ConnectorsListModel struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Data    []Connectors          `tfsdk:"data"`
	Limit   types.Int64           `tfsdk:"limit"`
	Offset  types.Int64           `tfsdk:"offset"`
	Name    types.String          `tfsdk:"name"`
	Total   types.Int64           `tfsdk:"total"`
	Folder  types.String          `tfsdk:"folder"`
	Snippet types.String          `tfsdk:"snippet"`
	Device  types.String          `tfsdk:"device"`
	Sort    basetypes.StringValue `tfsdk:"sort"`
	Search  basetypes.StringValue `tfsdk:"search"`
	Filters basetypes.StringValue `tfsdk:"filters"`
}

// ConnectorsListDataSourceSchema defines the schema for a list data source.
var ConnectorsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: ConnectorsDataSourceSchema.Attributes,
			},
		},
		"limit":   dsschema.Int64Attribute{Description: "The max number of items to return. Default: 200.", Optional: true},
		"offset":  dsschema.Int64Attribute{Description: "The offset of the first item to return.", Optional: true},
		"name":    dsschema.StringAttribute{Description: "The name of the item.", Optional: true},
		"total":   dsschema.Int64Attribute{Description: "The total number of items.", Computed: true},
		"folder":  dsschema.StringAttribute{Description: "The folder of the item. Default: Shared.", Optional: true},
		"snippet": dsschema.StringAttribute{Description: "The snippet of the item.", Optional: true},
		"device":  dsschema.StringAttribute{Description: "The device of the item.", Optional: true},
		"sort": dsschema.StringAttribute{
			Description: "List of fields from item response to sort by.",
			Optional:    true,
		},
		"search": dsschema.StringAttribute{
			Description: "String to filter list results by. Is searched over multiple fields in each object. Multiple searches can be specified.",
			Optional:    true,
		},
		"filters": dsschema.StringAttribute{
			Description: "String to filter list results by searching one specified field of an object. Multiple filters can be specified.",
			Optional:    true,
		},
	},
}
