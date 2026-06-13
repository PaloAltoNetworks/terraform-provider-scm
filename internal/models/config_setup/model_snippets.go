package models

import (
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

// Package: config_setup
// This file contains models for the config_setup SDK package

// Snippets represents the Terraform model for Snippets
type Snippets struct {
	Tfid        types.String          `tfsdk:"tfid"`
	Description basetypes.StringValue `tfsdk:"description"`
	Id          basetypes.StringValue `tfsdk:"id"`
	Labels      basetypes.ListValue   `tfsdk:"labels"`
	Name        basetypes.StringValue `tfsdk:"name"`
	Type        basetypes.StringValue `tfsdk:"type"`
}

// AttrTypes defines the attribute types for the Snippets model.
func (o Snippets) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":        basetypes.StringType{},
		"description": basetypes.StringType{},
		"id":          basetypes.StringType{},
		"labels":      basetypes.ListType{ElemType: basetypes.StringType{}},
		"name":        basetypes.StringType{},
		"type":        basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of Snippets objects.
func (o Snippets) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// SnippetsResourceSchema defines the schema for Snippets resource
var SnippetsResourceSchema = schema.Schema{
	MarkdownDescription: "Snippet resource",
	Attributes: map[string]schema.Attribute{
		"description": schema.StringAttribute{
			MarkdownDescription: "The description of the snippet",
			Optional:            true,
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the snippet",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"labels": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Labels applied to the snippet",
			Optional:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the snippet",
			Required:            true,
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"type": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("predefined", "custom", "readonly"),
			},
			MarkdownDescription: "The snippet type",
			Computed:            true,
		},
	},
}

// SnippetsDataSourceSchema defines the schema for Snippets data source
var SnippetsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Snippet data source",
	Attributes: map[string]dsschema.Attribute{
		"description": dsschema.StringAttribute{
			MarkdownDescription: "The description of the snippet",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the snippet",
			Required:            true,
		},
		"labels": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Labels applied to the snippet",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the snippet",
			Optional:            true,
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"type": dsschema.StringAttribute{
			MarkdownDescription: "The snippet type",
			Computed:            true,
		},
	},
}

// SnippetsListModel represents the data model for a list data source.
type SnippetsListModel struct {
	Tfid    types.String `tfsdk:"tfid"`
	Data    []Snippets   `tfsdk:"data"`
	Limit   types.Int64  `tfsdk:"limit"`
	Offset  types.Int64  `tfsdk:"offset"`
	Name    types.String `tfsdk:"name"`
	Total   types.Int64  `tfsdk:"total"`
	Folder  types.String `tfsdk:"folder"`
	Snippet types.String `tfsdk:"snippet"`
	Device  types.String `tfsdk:"device"`
}

// SnippetsListDataSourceSchema defines the schema for a list data source.
var SnippetsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: SnippetsDataSourceSchema.Attributes,
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
