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
// This file contains models for the Devices resource (moving firewalls into folders)

// DevicesTf represents the Terraform model for Devices
type DevicesTf struct {
	Tfid            types.String          `tfsdk:"tfid"`
	Id              basetypes.StringValue `tfsdk:"id"`
	Name            basetypes.StringValue `tfsdk:"name"`
	Folder          basetypes.StringValue `tfsdk:"folder"`
	Description     basetypes.StringValue `tfsdk:"description"`
	DisplayName     basetypes.StringValue `tfsdk:"display_name"`
	Labels          basetypes.ListValue   `tfsdk:"labels"`
	Snippets        basetypes.ListValue   `tfsdk:"snippets"`
	Hostname        basetypes.StringValue `tfsdk:"hostname"`
	IpAddress       basetypes.StringValue `tfsdk:"ip_address"`
	IpV6Address     basetypes.StringValue `tfsdk:"ipv6_address"`
	MacAddress      basetypes.StringValue `tfsdk:"mac_address"`
	Family          basetypes.StringValue `tfsdk:"family"`
	Model           basetypes.StringValue `tfsdk:"model"`
	SoftwareVersion basetypes.StringValue `tfsdk:"software_version"`
	IsConnected     basetypes.BoolValue   `tfsdk:"is_connected"`
}

// AttrTypes defines the attribute types for the DevicesTf model.
func (o DevicesTf) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":             basetypes.StringType{},
		"id":               basetypes.StringType{},
		"name":             basetypes.StringType{},
		"folder":           basetypes.StringType{},
		"description":      basetypes.StringType{},
		"display_name":     basetypes.StringType{},
		"labels":           basetypes.ListType{ElemType: basetypes.StringType{}},
		"snippets":         basetypes.ListType{ElemType: basetypes.StringType{}},
		"hostname":         basetypes.StringType{},
		"ip_address":       basetypes.StringType{},
		"ipv6_address":     basetypes.StringType{},
		"mac_address":      basetypes.StringType{},
		"family":           basetypes.StringType{},
		"model":            basetypes.StringType{},
		"software_version": basetypes.StringType{},
		"is_connected":     basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of DevicesTf objects.
func (o DevicesTf) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// DevicesResourceSchema defines the schema for the Devices resource.
// This resource manages device-to-folder assignment, labels, and snippets.
// The device itself is not created or deleted by Terraform — it already exists in SCM.
var DevicesResourceSchema = schema.Schema{
	MarkdownDescription: "Manages a device in Strata Cloud Manager. Use this resource to move firewalls into folders, update labels, and manage snippets.",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the device",
			Required:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"folder": schema.StringAttribute{
			MarkdownDescription: "The folder containing the device. Set this to move a device into a folder.",
			Required:            true,
		},
		"description": schema.StringAttribute{
			MarkdownDescription: "The description of the device",
			Optional:            true,
			Computed:            true,
		},
		"display_name": schema.StringAttribute{
			MarkdownDescription: "The display name of the device",
			Optional:            true,
			Computed:            true,
		},
		"labels": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Labels assigned to the device",
			Optional:            true,
			Computed:            true,
		},
		"snippets": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Snippets associated with the device",
			Optional:            true,
			Computed:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The serial number / name of the device",
			Computed:            true,
		},
		"hostname": schema.StringAttribute{
			MarkdownDescription: "The hostname of the device",
			Computed:            true,
		},
		"ip_address": schema.StringAttribute{
			MarkdownDescription: "The IPv4 address of the device",
			Computed:            true,
		},
		"ipv6_address": schema.StringAttribute{
			MarkdownDescription: "The IPv6 address of the device",
			Computed:            true,
		},
		"mac_address": schema.StringAttribute{
			MarkdownDescription: "The MAC address of the device",
			Computed:            true,
		},
		"family": schema.StringAttribute{
			MarkdownDescription: "The product family of the device",
			Computed:            true,
		},
		"model": schema.StringAttribute{
			MarkdownDescription: "The model of the device",
			Computed:            true,
		},
		"software_version": schema.StringAttribute{
			MarkdownDescription: "The software version running on the device",
			Computed:            true,
		},
		"is_connected": schema.BoolAttribute{
			MarkdownDescription: "Whether the device is currently connected",
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

// DevicesDataSourceSchema defines the schema for the Devices data source.
var DevicesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves information about a device in Strata Cloud Manager.",
	Attributes: map[string]dsschema.Attribute{
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the device",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The serial number / name of the device",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder containing the device",
			Computed:            true,
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "The description of the device",
			Computed:            true,
		},
		"display_name": dsschema.StringAttribute{
			MarkdownDescription: "The display name of the device",
			Computed:            true,
		},
		"labels": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Labels assigned to the device",
			Computed:            true,
		},
		"snippets": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Snippets associated with the device",
			Computed:            true,
		},
		"hostname": dsschema.StringAttribute{
			MarkdownDescription: "The hostname of the device",
			Computed:            true,
		},
		"ip_address": dsschema.StringAttribute{
			MarkdownDescription: "The IPv4 address of the device",
			Computed:            true,
		},
		"ipv6_address": dsschema.StringAttribute{
			MarkdownDescription: "The IPv6 address of the device",
			Computed:            true,
		},
		"mac_address": dsschema.StringAttribute{
			MarkdownDescription: "The MAC address of the device",
			Computed:            true,
		},
		"family": dsschema.StringAttribute{
			MarkdownDescription: "The product family of the device",
			Computed:            true,
		},
		"model": dsschema.StringAttribute{
			MarkdownDescription: "The model of the device",
			Computed:            true,
		},
		"software_version": dsschema.StringAttribute{
			MarkdownDescription: "The software version running on the device",
			Computed:            true,
		},
		"is_connected": dsschema.BoolAttribute{
			MarkdownDescription: "Whether the device is currently connected",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// DevicesListModel represents the data model for a list data source.
type DevicesListModel struct {
	Tfid   types.String `tfsdk:"tfid"`
	Data   []DevicesTf  `tfsdk:"data"`
	Limit  types.Int64  `tfsdk:"limit"`
	Offset types.Int64  `tfsdk:"offset"`
	Total  types.Int64  `tfsdk:"total"`
	Folder types.String `tfsdk:"folder"`
}

// DevicesListDataSourceSchema defines the schema for a list data source.
var DevicesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of devices in Strata Cloud Manager.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The list of devices.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: DevicesDataSourceSchema.Attributes,
			},
		},
		"limit":  dsschema.Int64Attribute{Description: "The max number of items to return. Default: 200.", Optional: true},
		"offset": dsschema.Int64Attribute{Description: "The offset of the first item to return.", Optional: true},
		"total":  dsschema.Int64Attribute{Description: "The total number of items.", Computed: true},
		"folder": dsschema.StringAttribute{Description: "Filter devices by folder.", Optional: true},
	},
}
