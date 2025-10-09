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

// Folders represents the Terraform model for Folders
type Folders struct {
	Tfid        types.String          `tfsdk:"tfid"`
	Description basetypes.StringValue `tfsdk:"description"`
	Id          basetypes.StringValue `tfsdk:"id"`
	Labels      basetypes.ListValue   `tfsdk:"labels"`
	Name        basetypes.StringValue `tfsdk:"name"`
	Parent      basetypes.StringValue `tfsdk:"parent"`
	Snippets    basetypes.ListValue   `tfsdk:"snippets"`
}

// AttrTypes defines the attribute types for the Folders model.
func (o Folders) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":        basetypes.StringType{},
		"description": basetypes.StringType{},
		"id":          basetypes.StringType{},
		"labels":      basetypes.ListType{ElemType: basetypes.StringType{}},
		"name":        basetypes.StringType{},
		"parent":      basetypes.StringType{},
		"snippets":    basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of Folders objects.
func (o Folders) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// FoldersResourceSchema defines the schema for Folders resource
var FoldersResourceSchema = schema.Schema{
	MarkdownDescription: "Folder resource",
	Attributes: map[string]schema.Attribute{
		"description": schema.StringAttribute{
			MarkdownDescription: "The description of the folder",
			Optional:            true,
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the folder",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"labels": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Labels assigned to the folder",
			Optional:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the folder",
			Required:            true,
		},
		"parent": schema.StringAttribute{
			MarkdownDescription: "The parent folder",
			Required:            true,
		},
		"snippets": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Snippets associated with the folder",
			Optional:            true,
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

// FoldersDataSourceSchema defines the schema for Folders data source
var FoldersDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Folder data source",
	Attributes: map[string]dsschema.Attribute{
		"description": dsschema.StringAttribute{
			MarkdownDescription: "The description of the folder",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the folder",
			Required:            true,
		},
		"labels": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Labels assigned to the folder",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the folder",
			Optional:            true,
			Computed:            true,
		},
		"parent": dsschema.StringAttribute{
			MarkdownDescription: "The parent folder",
			Computed:            true,
		},
		"snippets": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Snippets associated with the folder",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// FoldersListModel represents the data model for a list data source.
type FoldersListModel struct {
	Tfid    types.String `tfsdk:"tfid"`
	Data    []Folders    `tfsdk:"data"`
	Limit   types.Int64  `tfsdk:"limit"`
	Offset  types.Int64  `tfsdk:"offset"`
	Name    types.String `tfsdk:"name"`
	Total   types.Int64  `tfsdk:"total"`
	Folder  types.String `tfsdk:"folder"`
	Snippet types.String `tfsdk:"snippet"`
	Device  types.String `tfsdk:"device"`
}

// FoldersListDataSourceSchema defines the schema for a list data source.
var FoldersListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: FoldersDataSourceSchema.Attributes,
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
