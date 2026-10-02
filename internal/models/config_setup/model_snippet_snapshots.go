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

// SaveSnippetSnapshotConfigResponse represents the Terraform model for SaveSnippetSnapshotConfigResponse
type SaveSnippetSnapshotConfigResponse struct {
	Tfid   types.String          `tfsdk:"tfid"`
	Result basetypes.ObjectValue `tfsdk:"result"`
	Status basetypes.StringValue `tfsdk:"status"`
}

// SaveSnippetSnapshotConfigResponseResult represents a nested structure within the SaveSnippetSnapshotConfigResponse model
type SaveSnippetSnapshotConfigResponseResult struct {
	Version basetypes.StringValue `tfsdk:"version"`
}

// AttrTypes defines the attribute types for the SaveSnippetSnapshotConfigResponse model.
func (o SaveSnippetSnapshotConfigResponse) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"result": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"version": basetypes.StringType{},
			},
		},
		"status": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of SaveSnippetSnapshotConfigResponse objects.
func (o SaveSnippetSnapshotConfigResponse) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SaveSnippetSnapshotConfigResponseResult model.
func (o SaveSnippetSnapshotConfigResponseResult) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"version": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of SaveSnippetSnapshotConfigResponseResult objects.
func (o SaveSnippetSnapshotConfigResponseResult) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// SaveSnippetSnapshotConfigResponseResourceSchema defines the schema for SaveSnippetSnapshotConfigResponse resource
var SaveSnippetSnapshotConfigResponseResourceSchema = schema.Schema{
	MarkdownDescription: "SaveSnippetSnapshotConfigResponse resource",
	Attributes: map[string]schema.Attribute{
		"result": schema.SingleNestedAttribute{
			MarkdownDescription: "Result",
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"version": schema.StringAttribute{
					MarkdownDescription: "Version",
					Optional:            true,
					Computed:            true,
				},
			},
		},
		"status": schema.StringAttribute{
			MarkdownDescription: "Status",
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

// SaveSnippetSnapshotConfigResponseDataSourceSchema defines the schema for SaveSnippetSnapshotConfigResponse data source
var SaveSnippetSnapshotConfigResponseDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "SaveSnippetSnapshotConfigResponse data source",
	Attributes: map[string]dsschema.Attribute{
		"result": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Result",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"version": dsschema.StringAttribute{
					MarkdownDescription: "Version",
					Computed:            true,
				},
			},
		},
		"status": dsschema.StringAttribute{
			MarkdownDescription: "Status",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// SaveSnippetSnapshotConfigResponseListModel represents the data model for a list data source.
type SaveSnippetSnapshotConfigResponseListModel struct {
	Tfid    types.String                        `tfsdk:"tfid"`
	Data    []SaveSnippetSnapshotConfigResponse `tfsdk:"data"`
	Limit   types.Int64                         `tfsdk:"limit"`
	Offset  types.Int64                         `tfsdk:"offset"`
	Name    types.String                        `tfsdk:"name"`
	Total   types.Int64                         `tfsdk:"total"`
	Folder  types.String                        `tfsdk:"folder"`
	Snippet types.String                        `tfsdk:"snippet"`
	Device  types.String                        `tfsdk:"device"`
}

// SaveSnippetSnapshotConfigResponseListDataSourceSchema defines the schema for a list data source.
var SaveSnippetSnapshotConfigResponseListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: SaveSnippetSnapshotConfigResponseDataSourceSchema.Attributes,
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
