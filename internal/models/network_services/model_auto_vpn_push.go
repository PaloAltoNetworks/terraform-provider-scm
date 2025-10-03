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

// Package: network_services
// This file contains models for the network_services SDK package

// AutoVpnPushResponse represents the Terraform model for AutoVpnPushResponse
type AutoVpnPushResponse struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Job     basetypes.StringValue `tfsdk:"job"`
	Message basetypes.StringValue `tfsdk:"message"`
	Success basetypes.BoolValue   `tfsdk:"success"`
}

// AutoVpnPushConfig represents a nested structure within the AutoVpnPushResponse model
type AutoVpnPushConfig struct {
	Tfid           types.String        `tfsdk:"tfid"`
	AutoVpnDevices basetypes.ListValue `tfsdk:"auto_vpn_devices"`
}

// AutoVpnPushConfigAutoVpnDevicesInner represents a nested structure within the AutoVpnPushResponse model
type AutoVpnPushConfigAutoVpnDevicesInner struct {
	Name       basetypes.StringValue `tfsdk:"name"`
	RefreshPsk basetypes.BoolValue   `tfsdk:"refresh_psk"`
}

// AttrTypes defines the attribute types for the AutoVpnPushResponse model.
func (o AutoVpnPushResponse) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":    basetypes.StringType{},
		"job":     basetypes.StringType{},
		"message": basetypes.StringType{},
		"success": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of AutoVpnPushResponse objects.
func (o AutoVpnPushResponse) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AutoVpnPushConfig model.
func (o AutoVpnPushConfig) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"auto_vpn_devices": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":        basetypes.StringType{},
				"refresh_psk": basetypes.BoolType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of AutoVpnPushConfig objects.
func (o AutoVpnPushConfig) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AutoVpnPushConfigAutoVpnDevicesInner model.
func (o AutoVpnPushConfigAutoVpnDevicesInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":        basetypes.StringType{},
		"refresh_psk": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of AutoVpnPushConfigAutoVpnDevicesInner objects.
func (o AutoVpnPushConfigAutoVpnDevicesInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AutoVpnPushResponseResourceSchema defines the schema for AutoVpnPushResponse resource
var AutoVpnPushResponseResourceSchema = schema.Schema{
	MarkdownDescription: "Manages SCM AutoVpnPushResponse objects",
	Attributes: map[string]schema.Attribute{
		"job": schema.StringAttribute{
			MarkdownDescription: "Job ID",
			Optional:            true,
		},
		"message": schema.StringAttribute{
			MarkdownDescription: "Job message",
			Optional:            true,
		},
		"success": schema.BoolAttribute{
			MarkdownDescription: "Push successful?",
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

// AutoVpnPushResponseDataSourceSchema defines the schema for AutoVpnPushResponse data source
var AutoVpnPushResponseDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "AutoVpnPushResponse data source",
	Attributes: map[string]dsschema.Attribute{
		"job": dsschema.StringAttribute{
			MarkdownDescription: "Job ID",
			Computed:            true,
		},
		"message": dsschema.StringAttribute{
			MarkdownDescription: "Job message",
			Computed:            true,
		},
		"success": dsschema.BoolAttribute{
			MarkdownDescription: "Push successful?",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// AutoVpnPushResponseListModel represents the data model for a list data source.
type AutoVpnPushResponseListModel struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Data    []AutoVpnPushResponse `tfsdk:"data"`
	Limit   types.Int64           `tfsdk:"limit"`
	Offset  types.Int64           `tfsdk:"offset"`
	Name    types.String          `tfsdk:"name"`
	Total   types.Int64           `tfsdk:"total"`
	Folder  types.String          `tfsdk:"folder"`
	Snippet types.String          `tfsdk:"snippet"`
	Device  types.String          `tfsdk:"device"`
}

// AutoVpnPushResponseListDataSourceSchema defines the schema for a list data source.
var AutoVpnPushResponseListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: AutoVpnPushResponseDataSourceSchema.Attributes,
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
