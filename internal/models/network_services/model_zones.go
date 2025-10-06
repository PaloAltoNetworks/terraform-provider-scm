package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: network_services
// This file contains models for the network_services SDK package

// Zones represents the Terraform model for Zones
type Zones struct {
	Tfid                       types.String          `tfsdk:"tfid"`
	Device                     basetypes.StringValue `tfsdk:"device"`
	DeviceAcl                  basetypes.ObjectValue `tfsdk:"device_acl"`
	DosLogSetting              basetypes.StringValue `tfsdk:"dos_log_setting"`
	DosProfile                 basetypes.StringValue `tfsdk:"dos_profile"`
	EnableDeviceIdentification basetypes.BoolValue   `tfsdk:"enable_device_identification"`
	EnableUserIdentification   basetypes.BoolValue   `tfsdk:"enable_user_identification"`
	Folder                     basetypes.StringValue `tfsdk:"folder"`
	Id                         basetypes.StringValue `tfsdk:"id"`
	Name                       basetypes.StringValue `tfsdk:"name"`
	Network                    basetypes.ObjectValue `tfsdk:"network"`
	Snippet                    basetypes.StringValue `tfsdk:"snippet"`
	UserAcl                    basetypes.ObjectValue `tfsdk:"user_acl"`
}

// ZonesDeviceAcl represents a nested structure within the Zones model
type ZonesDeviceAcl struct {
	ExcludeList basetypes.ListValue `tfsdk:"exclude_list"`
	IncludeList basetypes.ListValue `tfsdk:"include_list"`
}

// ZonesNetwork represents a nested structure within the Zones model
type ZonesNetwork struct {
	EnablePacketBufferProtection basetypes.BoolValue   `tfsdk:"enable_packet_buffer_protection"`
	LogSetting                   basetypes.StringValue `tfsdk:"log_setting"`
	ZoneProtectionProfile        basetypes.StringValue `tfsdk:"zone_protection_profile"`
}

// AttrTypes defines the attribute types for the Zones model.
func (o Zones) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":   basetypes.StringType{},
		"device": basetypes.StringType{},
		"device_acl": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"exclude_list": basetypes.ListType{ElemType: basetypes.StringType{}},
				"include_list": basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		},
		"dos_log_setting":              basetypes.StringType{},
		"dos_profile":                  basetypes.StringType{},
		"enable_device_identification": basetypes.BoolType{},
		"enable_user_identification":   basetypes.BoolType{},
		"folder":                       basetypes.StringType{},
		"id":                           basetypes.StringType{},
		"name":                         basetypes.StringType{},
		"network": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable_packet_buffer_protection": basetypes.BoolType{},
				"log_setting":                     basetypes.StringType{},
				"zone_protection_profile":         basetypes.StringType{},
			},
		},
		"snippet": basetypes.StringType{},
		"user_acl": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"exclude_list": basetypes.ListType{ElemType: basetypes.StringType{}},
				"include_list": basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		},
	}
}

// AttrType returns the attribute type for a list of Zones objects.
func (o Zones) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZonesDeviceAcl model.
func (o ZonesDeviceAcl) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"exclude_list": basetypes.ListType{ElemType: basetypes.StringType{}},
		"include_list": basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of ZonesDeviceAcl objects.
func (o ZonesDeviceAcl) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZonesNetwork model.
func (o ZonesNetwork) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable_packet_buffer_protection": basetypes.BoolType{},
		"log_setting":                     basetypes.StringType{},
		"zone_protection_profile":         basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ZonesNetwork objects.
func (o ZonesNetwork) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// ZonesResourceSchema defines the schema for Zones resource
var ZonesResourceSchema = schema.Schema{
	MarkdownDescription: "Zone resource",
	Attributes: map[string]schema.Attribute{
		"device": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("folder"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
			},
			MarkdownDescription: "The device in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"device_acl": schema.SingleNestedAttribute{
			MarkdownDescription: "Device acl",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"exclude_list": schema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Exclude list",
					Optional:            true,
				},
				"include_list": schema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Include list",
					Optional:            true,
				},
			},
		},
		"dos_log_setting": schema.StringAttribute{
			MarkdownDescription: "Dos log setting",
			Optional:            true,
		},
		"dos_profile": schema.StringAttribute{
			MarkdownDescription: "Dos profile",
			Optional:            true,
		},
		"enable_device_identification": schema.BoolAttribute{
			MarkdownDescription: "Enable device identification",
			Optional:            true,
		},
		"enable_user_identification": schema.BoolAttribute{
			MarkdownDescription: "Enable user identification",
			Optional:            true,
		},
		"folder": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("device"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
			},
			MarkdownDescription: "Folder",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(63),
			},
			MarkdownDescription: "Alphanumeric string begin with letter: [0-9a-zA-Z._-]",
			Required:            true,
		},
		"network": schema.SingleNestedAttribute{
			MarkdownDescription: "Network",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"enable_packet_buffer_protection": schema.BoolAttribute{
					MarkdownDescription: "Enable packet buffer protection",
					Optional:            true,
				},
				"log_setting": schema.StringAttribute{
					MarkdownDescription: "Log setting",
					Optional:            true,
				},
				"zone_protection_profile": schema.StringAttribute{
					MarkdownDescription: "Zone protection profile",
					Optional:            true,
				},
			},
		},
		"snippet": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("device"),
					path.MatchRelative().AtParent().AtName("folder"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
			},
			MarkdownDescription: "The snippet in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"user_acl": schema.SingleNestedAttribute{
			MarkdownDescription: "User acl",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"exclude_list": schema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Exclude list",
					Optional:            true,
				},
				"include_list": schema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Include list",
					Optional:            true,
				},
			},
		},
	},
}

// ZonesDataSourceSchema defines the schema for Zones data source
var ZonesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Zone data source",
	Attributes: map[string]dsschema.Attribute{
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"device_acl": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Device acl",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"exclude_list": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Exclude list",
					Computed:            true,
				},
				"include_list": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Include list",
					Computed:            true,
				},
			},
		},
		"dos_log_setting": dsschema.StringAttribute{
			MarkdownDescription: "Dos log setting",
			Computed:            true,
		},
		"dos_profile": dsschema.StringAttribute{
			MarkdownDescription: "Dos profile",
			Computed:            true,
		},
		"enable_device_identification": dsschema.BoolAttribute{
			MarkdownDescription: "Enable device identification",
			Computed:            true,
		},
		"enable_user_identification": dsschema.BoolAttribute{
			MarkdownDescription: "Enable user identification",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "Folder",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Alphanumeric string begin with letter: [0-9a-zA-Z._-]",
			Optional:            true,
			Computed:            true,
		},
		"network": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Network",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"enable_packet_buffer_protection": dsschema.BoolAttribute{
					MarkdownDescription: "Enable packet buffer protection",
					Computed:            true,
				},
				"log_setting": dsschema.StringAttribute{
					MarkdownDescription: "Log setting",
					Computed:            true,
				},
				"zone_protection_profile": dsschema.StringAttribute{
					MarkdownDescription: "Zone protection profile",
					Computed:            true,
				},
			},
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"user_acl": dsschema.SingleNestedAttribute{
			MarkdownDescription: "User acl",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"exclude_list": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Exclude list",
					Computed:            true,
				},
				"include_list": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Include list",
					Computed:            true,
				},
			},
		},
	},
}

// ZonesListModel represents the data model for a list data source.
type ZonesListModel struct {
	Tfid    types.String `tfsdk:"tfid"`
	Data    []Zones      `tfsdk:"data"`
	Limit   types.Int64  `tfsdk:"limit"`
	Offset  types.Int64  `tfsdk:"offset"`
	Name    types.String `tfsdk:"name"`
	Total   types.Int64  `tfsdk:"total"`
	Folder  types.String `tfsdk:"folder"`
	Snippet types.String `tfsdk:"snippet"`
	Device  types.String `tfsdk:"device"`
}

// ZonesListDataSourceSchema defines the schema for a list data source.
var ZonesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: ZonesDataSourceSchema.Attributes,
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
