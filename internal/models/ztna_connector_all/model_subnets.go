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

// Subnets represents the Terraform model for Subnets
type Subnets struct {
	Tfid        types.String          `tfsdk:"tfid"`
	AppEnabled  basetypes.BoolValue   `tfsdk:"app_enabled"`
	CreatedTime basetypes.StringValue `tfsdk:"created_time"`
	Description basetypes.StringValue `tfsdk:"description"`
	Group       basetypes.StringValue `tfsdk:"group"`
	IcmpAllowed basetypes.BoolValue   `tfsdk:"icmp_allowed"`
	IpSubnets   basetypes.StringValue `tfsdk:"ip_subnets"`
	Name        basetypes.StringValue `tfsdk:"name"`
	Oid         basetypes.StringValue `tfsdk:"oid"`
	UpdatedTime basetypes.StringValue `tfsdk:"updated_time"`
}

// AttrTypes defines the attribute types for the Subnets model.
func (o Subnets) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":         basetypes.StringType{},
		"app_enabled":  basetypes.BoolType{},
		"created_time": basetypes.StringType{},
		"description":  basetypes.StringType{},
		"group":        basetypes.StringType{},
		"icmp_allowed": basetypes.BoolType{},
		"ip_subnets":   basetypes.StringType{},
		"name":         basetypes.StringType{},
		"oid":          basetypes.StringType{},
		"updated_time": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of Subnets objects.
func (o Subnets) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// SubnetsResourceSchema defines the schema for Subnets resource
var SubnetsResourceSchema = schema.Schema{
	MarkdownDescription: "Subnet resource",
	Attributes: map[string]schema.Attribute{
		"app_enabled": schema.BoolAttribute{
			MarkdownDescription: "Whether the IP subnet rule is enabled.\n\nIf omitted, defaults to false.",
			Optional:            true,
			Computed:            true,
		},
		"created_time": schema.StringAttribute{
			MarkdownDescription: "Created time",
			Computed:            true,
		},
		"description": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[\\p{L}\\p{N}\\p{P}\\s,.:_-]*$"), "pattern must match "+"^[\\p{L}\\p{N}\\p{P}\\s,.:_-]*$"),
			},
			MarkdownDescription: "Description",
			Optional:            true,
			Computed:            true,
		},
		"group": schema.StringAttribute{
			MarkdownDescription: "A comma separated list of connector group IDs",
			Required:            true,
		},
		"icmp_allowed": schema.BoolAttribute{
			MarkdownDescription: "Whether ICMP is allowed for this IP subnet rule.\n\nIf omitted, defaults to true.",
			Optional:            true,
			Computed:            true,
		},
		"ip_subnets": schema.StringAttribute{
			MarkdownDescription: "IPv4 subnet in CIDR notation (x.x.x.x/y)",
			Required:            true,
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(64),
				stringvalidator.LengthAtLeast(1),
				stringvalidator.RegexMatches(regexp.MustCompile("^[\\p{L}\\p{N}\\p{P}\\s,.:_-]*$"), "pattern must match "+"^[\\p{L}\\p{N}\\p{P}\\s,.:_-]*$"),
			},
			MarkdownDescription: "Name of the IP Subnet rule.\n\nIt can only be 64 characters long\nand contain unicode text, space, dash, or underscore, or period.",
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

// SubnetsDataSourceSchema defines the schema for Subnets data source
var SubnetsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Subnet data source",
	Attributes: map[string]dsschema.Attribute{
		"app_enabled": dsschema.BoolAttribute{
			MarkdownDescription: "Whether the IP subnet rule is enabled.\n\nIf omitted, defaults to false.",
			Computed:            true,
		},
		"created_time": dsschema.StringAttribute{
			MarkdownDescription: "Created time",
			Computed:            true,
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Description",
			Computed:            true,
		},
		"group": dsschema.StringAttribute{
			MarkdownDescription: "A comma separated list of connector group IDs",
			Computed:            true,
		},
		"icmp_allowed": dsschema.BoolAttribute{
			MarkdownDescription: "Whether ICMP is allowed for this IP subnet rule.\n\nIf omitted, defaults to true.",
			Computed:            true,
		},
		"ip_subnets": dsschema.StringAttribute{
			MarkdownDescription: "IPv4 subnet in CIDR notation (x.x.x.x/y)",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Name of the IP Subnet rule.\n\nIt can only be 64 characters long\nand contain unicode text, space, dash, or underscore, or period.",
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

// SubnetsListModel represents the data model for a list data source.
type SubnetsListModel struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Data    []Subnets             `tfsdk:"data"`
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

// SubnetsListDataSourceSchema defines the schema for a list data source.
var SubnetsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: SubnetsDataSourceSchema.Attributes,
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
