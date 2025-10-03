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

// Package: deployment_services
// This file contains models for the deployment_services SDK package

// ServiceConnectionGroups represents the Terraform model for ServiceConnectionGroups
type ServiceConnectionGroups struct {
	Tfid        types.String          `tfsdk:"tfid"`
	DisableSnat basetypes.BoolValue   `tfsdk:"disable_snat"`
	Folder      basetypes.StringValue `tfsdk:"folder"`
	Id          basetypes.StringValue `tfsdk:"id"`
	Name        basetypes.StringValue `tfsdk:"name"`
	PbfOnly     basetypes.BoolValue   `tfsdk:"pbf_only"`
	Target      basetypes.ListValue   `tfsdk:"target"`
}

// AttrTypes defines the attribute types for the ServiceConnectionGroups model.
func (o ServiceConnectionGroups) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":         basetypes.StringType{},
		"disable_snat": basetypes.BoolType{},
		"folder":       basetypes.StringType{},
		"id":           basetypes.StringType{},
		"name":         basetypes.StringType{},
		"pbf_only":     basetypes.BoolType{},
		"target":       basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of ServiceConnectionGroups objects.
func (o ServiceConnectionGroups) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// ServiceConnectionGroupsResourceSchema defines the schema for ServiceConnectionGroups resource
var ServiceConnectionGroupsResourceSchema = schema.Schema{
	MarkdownDescription: "Manages SCM ServiceConnectionGroups objects",
	Attributes: map[string]schema.Attribute{
		"disable_snat": schema.BoolAttribute{
			MarkdownDescription: "Disable snat",
			Optional:            true,
		},
		"folder": schema.StringAttribute{
			MarkdownDescription: "The folder containing the service connection group",
			Required:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the service connection group",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "Name",
			Required:            true,
		},
		"pbf_only": schema.BoolAttribute{
			MarkdownDescription: "Pbf only",
			Optional:            true,
		},
		"target": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Target",
			Required:            true,
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

// ServiceConnectionGroupsDataSourceSchema defines the schema for ServiceConnectionGroups data source
var ServiceConnectionGroupsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "ServiceConnectionGroups data source",
	Attributes: map[string]dsschema.Attribute{
		"disable_snat": dsschema.BoolAttribute{
			MarkdownDescription: "Disable snat",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder containing the service connection group",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the service connection group",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Name",
			Optional:            true,
			Computed:            true,
		},
		"pbf_only": dsschema.BoolAttribute{
			MarkdownDescription: "Pbf only",
			Computed:            true,
		},
		"target": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Target",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// ServiceConnectionGroupsListModel represents the data model for a list data source.
type ServiceConnectionGroupsListModel struct {
	Tfid    types.String              `tfsdk:"tfid"`
	Data    []ServiceConnectionGroups `tfsdk:"data"`
	Limit   types.Int64               `tfsdk:"limit"`
	Offset  types.Int64               `tfsdk:"offset"`
	Name    types.String              `tfsdk:"name"`
	Total   types.Int64               `tfsdk:"total"`
	Folder  types.String              `tfsdk:"folder"`
	Snippet types.String              `tfsdk:"snippet"`
	Device  types.String              `tfsdk:"device"`
}

// ServiceConnectionGroupsListDataSourceSchema defines the schema for a list data source.
var ServiceConnectionGroupsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: ServiceConnectionGroupsDataSourceSchema.Attributes,
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
