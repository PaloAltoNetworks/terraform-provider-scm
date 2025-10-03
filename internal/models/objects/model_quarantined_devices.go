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

// Package: objects
// This file contains models for the objects SDK package

// QuarantinedDevices represents the Terraform model for QuarantinedDevices
type QuarantinedDevices struct {
	Tfid         types.String          `tfsdk:"tfid"`
	HostId       basetypes.StringValue `tfsdk:"host_id"`
	SerialNumber basetypes.StringValue `tfsdk:"serial_number"`
}

// AttrTypes defines the attribute types for the QuarantinedDevices model.
func (o QuarantinedDevices) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":          basetypes.StringType{},
		"host_id":       basetypes.StringType{},
		"serial_number": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of QuarantinedDevices objects.
func (o QuarantinedDevices) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// QuarantinedDevicesResourceSchema defines the schema for QuarantinedDevices resource
var QuarantinedDevicesResourceSchema = schema.Schema{
	MarkdownDescription: "Manages SCM QuarantinedDevices objects",
	Attributes: map[string]schema.Attribute{
		"host_id": schema.StringAttribute{
			MarkdownDescription: "Device host ID",
			Required:            true,
		},
		"serial_number": schema.StringAttribute{
			MarkdownDescription: "Device serial number",
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

// QuarantinedDevicesDataSourceSchema defines the schema for QuarantinedDevices data source
var QuarantinedDevicesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "QuarantinedDevices data source",
	Attributes: map[string]dsschema.Attribute{
		"host_id": dsschema.StringAttribute{
			MarkdownDescription: "Device host ID",
			Computed:            true,
		},
		"serial_number": dsschema.StringAttribute{
			MarkdownDescription: "Device serial number",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// QuarantinedDevicesListModel represents the data model for a list data source.
type QuarantinedDevicesListModel struct {
	Tfid    types.String         `tfsdk:"tfid"`
	Data    []QuarantinedDevices `tfsdk:"data"`
	Limit   types.Int64          `tfsdk:"limit"`
	Offset  types.Int64          `tfsdk:"offset"`
	Name    types.String         `tfsdk:"name"`
	Total   types.Int64          `tfsdk:"total"`
	Folder  types.String         `tfsdk:"folder"`
	Snippet types.String         `tfsdk:"snippet"`
	Device  types.String         `tfsdk:"device"`
}

// QuarantinedDevicesListDataSourceSchema defines the schema for a list data source.
var QuarantinedDevicesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: QuarantinedDevicesDataSourceSchema.Attributes,
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
