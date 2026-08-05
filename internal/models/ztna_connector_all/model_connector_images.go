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

// ConnectorImages represents the Terraform model for ConnectorImages
type ConnectorImages struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Id      basetypes.StringValue `tfsdk:"id"`
	Version basetypes.StringValue `tfsdk:"version"`
}

// AttrTypes defines the attribute types for the ConnectorImages model.
func (o ConnectorImages) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":    basetypes.StringType{},
		"id":      basetypes.StringType{},
		"version": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ConnectorImages objects.
func (o ConnectorImages) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// ConnectorImagesResourceSchema defines the schema for ConnectorImages resource
var ConnectorImagesResourceSchema = schema.Schema{
	MarkdownDescription: "ConnectorImage resource",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			MarkdownDescription: "Id",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"version": schema.StringAttribute{
			MarkdownDescription: "Version",
			Optional:            true,
			Computed:            true,
		},
	},
}

// ConnectorImagesDataSourceSchema defines the schema for ConnectorImages data source
var ConnectorImagesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "ConnectorImage data source",
	Attributes: map[string]dsschema.Attribute{
		"id": dsschema.StringAttribute{
			MarkdownDescription: "Id",
			Required:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"version": dsschema.StringAttribute{
			MarkdownDescription: "Version",
			Computed:            true,
		},
	},
}

// ConnectorImagesListModel represents the data model for a list data source.
type ConnectorImagesListModel struct {
	Tfid    types.String      `tfsdk:"tfid"`
	Data    []ConnectorImages `tfsdk:"data"`
	Limit   types.Int64       `tfsdk:"limit"`
	Offset  types.Int64       `tfsdk:"offset"`
	Name    types.String      `tfsdk:"name"`
	Total   types.Int64       `tfsdk:"total"`
	Folder  types.String      `tfsdk:"folder"`
	Snippet types.String      `tfsdk:"snippet"`
	Device  types.String      `tfsdk:"device"`
}

// ConnectorImagesListDataSourceSchema defines the schema for a list data source.
var ConnectorImagesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: ConnectorImagesDataSourceSchema.Attributes,
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
