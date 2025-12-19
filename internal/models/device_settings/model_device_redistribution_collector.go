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

// Package: device_settings
// This file contains models for the device_settings SDK package

// DeviceRedistributionCollector represents the Terraform model for DeviceRedistributionCollector
type DeviceRedistributionCollector struct {
	Tfid                    types.String          `tfsdk:"tfid"`
	Device                  basetypes.StringValue `tfsdk:"device"`
	Folder                  basetypes.StringValue `tfsdk:"folder"`
	Id                      basetypes.StringValue `tfsdk:"id"`
	RedistributionCollector basetypes.ObjectValue `tfsdk:"redistribution_collector"`
	Snippet                 basetypes.StringValue `tfsdk:"snippet"`
}

// DeviceRedistributionCollectorRedistributionCollector represents a nested structure within the DeviceRedistributionCollector model
type DeviceRedistributionCollectorRedistributionCollector struct {
	Interface basetypes.StringValue `tfsdk:"interface"`
}

// AttrTypes defines the attribute types for the DeviceRedistributionCollector model.
func (o DeviceRedistributionCollector) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":   basetypes.StringType{},
		"device": basetypes.StringType{},
		"folder": basetypes.StringType{},
		"id":     basetypes.StringType{},
		"redistribution_collector": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"interface": basetypes.StringType{},
			},
		},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DeviceRedistributionCollector objects.
func (o DeviceRedistributionCollector) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DeviceRedistributionCollectorRedistributionCollector model.
func (o DeviceRedistributionCollectorRedistributionCollector) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"interface": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DeviceRedistributionCollectorRedistributionCollector objects.
func (o DeviceRedistributionCollectorRedistributionCollector) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// DeviceRedistributionCollectorResourceSchema defines the schema for DeviceRedistributionCollector resource
var DeviceRedistributionCollectorResourceSchema = schema.Schema{
	MarkdownDescription: "DeviceRedistributionCollector resource",
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
			MarkdownDescription: "The device in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"folder": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("device"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
			},
			MarkdownDescription: "The folder in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
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
		"redistribution_collector": schema.SingleNestedAttribute{
			MarkdownDescription: "Redistribution collector",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"interface": schema.StringAttribute{
					MarkdownDescription: "User-ID collector interface",
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
			MarkdownDescription: "The snippet in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
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
	},
}

// DeviceRedistributionCollectorDataSourceSchema defines the schema for DeviceRedistributionCollector data source
var DeviceRedistributionCollectorDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "DeviceRedistributionCollector data source",
	Attributes: map[string]dsschema.Attribute{
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"redistribution_collector": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Redistribution collector",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"interface": dsschema.StringAttribute{
					MarkdownDescription: "User-ID collector interface",
					Computed:            true,
				},
			},
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// DeviceRedistributionCollectorListModel represents the data model for a list data source.
type DeviceRedistributionCollectorListModel struct {
	Tfid    types.String                    `tfsdk:"tfid"`
	Data    []DeviceRedistributionCollector `tfsdk:"data"`
	Limit   types.Int64                     `tfsdk:"limit"`
	Offset  types.Int64                     `tfsdk:"offset"`
	Name    types.String                    `tfsdk:"name"`
	Total   types.Int64                     `tfsdk:"total"`
	Folder  types.String                    `tfsdk:"folder"`
	Snippet types.String                    `tfsdk:"snippet"`
	Device  types.String                    `tfsdk:"device"`
}

// DeviceRedistributionCollectorListDataSourceSchema defines the schema for a list data source.
var DeviceRedistributionCollectorListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: DeviceRedistributionCollectorDataSourceSchema.Attributes,
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
