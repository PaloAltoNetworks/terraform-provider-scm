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
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// Package: network_services
// This file contains models for the network_services SDK package

// ConfigMatchList represents the Terraform model for ConfigMatchList
type ConfigMatchList struct {
	Tfid           types.String          `tfsdk:"tfid"`
	Description    basetypes.StringValue `tfsdk:"description"`
	Device         basetypes.StringValue `tfsdk:"device"`
	Filter         basetypes.StringValue `tfsdk:"filter"`
	Folder         basetypes.StringValue `tfsdk:"folder"`
	Id             basetypes.StringValue `tfsdk:"id"`
	Name           basetypes.StringValue `tfsdk:"name"`
	SendEmail      basetypes.ListValue   `tfsdk:"send_email"`
	SendHttp       basetypes.ListValue   `tfsdk:"send_http"`
	SendSnmptrap   basetypes.ListValue   `tfsdk:"send_snmptrap"`
	SendSyslog     basetypes.ListValue   `tfsdk:"send_syslog"`
	SendToPanorama basetypes.BoolValue   `tfsdk:"send_to_panorama"`
	Snippet        basetypes.StringValue `tfsdk:"snippet"`
}

// AttrTypes defines the attribute types for the ConfigMatchList model.
func (o ConfigMatchList) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":             basetypes.StringType{},
		"description":      basetypes.StringType{},
		"device":           basetypes.StringType{},
		"filter":           basetypes.StringType{},
		"folder":           basetypes.StringType{},
		"id":               basetypes.StringType{},
		"name":             basetypes.StringType{},
		"send_email":       basetypes.ListType{ElemType: basetypes.StringType{}},
		"send_http":        basetypes.ListType{ElemType: basetypes.StringType{}},
		"send_snmptrap":    basetypes.ListType{ElemType: basetypes.StringType{}},
		"send_syslog":      basetypes.ListType{ElemType: basetypes.StringType{}},
		"send_to_panorama": basetypes.BoolType{},
		"snippet":          basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ConfigMatchList objects.
func (o ConfigMatchList) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// ConfigMatchListResourceSchema defines the schema for ConfigMatchList resource
var ConfigMatchListResourceSchema = schema.Schema{
	MarkdownDescription: "ConfigMatchList resource",
	Attributes: map[string]schema.Attribute{
		"description": schema.StringAttribute{
			MarkdownDescription: "Description of the config match list entry",
			Optional:            true,
		},
		"device": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("folder"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
			},
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"filter": schema.StringAttribute{
			MarkdownDescription: "Filter of the config match list entry",
			Optional:            true,
		},
		"folder": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("device"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
				utils.FolderValidator(),
			},
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
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
			MarkdownDescription: "Name of the config match list entry",
			Required:            true,
		},
		"send_email": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Send Email List of the config match list entry",
			Optional:            true,
		},
		"send_http": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Send HTTP List of the config match list entry",
			Optional:            true,
		},
		"send_snmptrap": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Send SNMP Trap List of the config match list entry",
			Optional:            true,
		},
		"send_syslog": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Send Sys Log List of the config match list entry",
			Optional:            true,
		},
		"send_to_panorama": schema.BoolAttribute{
			MarkdownDescription: "Send Panorama Flag of the config match list entry",
			Optional:            true,
		},
		"snippet": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("device"),
					path.MatchRelative().AtParent().AtName("folder"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
			},
			MarkdownDescription: "The snippet in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
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

// ConfigMatchListDataSourceSchema defines the schema for ConfigMatchList data source
var ConfigMatchListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "ConfigMatchList data source",
	Attributes: map[string]dsschema.Attribute{
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Description of the config match list entry",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"filter": dsschema.StringAttribute{
			MarkdownDescription: "Filter of the config match list entry",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Name of the config match list entry",
			Optional:            true,
			Computed:            true,
		},
		"send_email": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Send Email List of the config match list entry",
			Computed:            true,
		},
		"send_http": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Send HTTP List of the config match list entry",
			Computed:            true,
		},
		"send_snmptrap": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Send SNMP Trap List of the config match list entry",
			Computed:            true,
		},
		"send_syslog": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Send Sys Log List of the config match list entry",
			Computed:            true,
		},
		"send_to_panorama": dsschema.BoolAttribute{
			MarkdownDescription: "Send Panorama Flag of the config match list entry",
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// ConfigMatchListListModel represents the data model for a list data source.
type ConfigMatchListListModel struct {
	Tfid    types.String      `tfsdk:"tfid"`
	Data    []ConfigMatchList `tfsdk:"data"`
	Limit   types.Int64       `tfsdk:"limit"`
	Offset  types.Int64       `tfsdk:"offset"`
	Name    types.String      `tfsdk:"name"`
	Total   types.Int64       `tfsdk:"total"`
	Folder  types.String      `tfsdk:"folder"`
	Snippet types.String      `tfsdk:"snippet"`
	Device  types.String      `tfsdk:"device"`
}

// ConfigMatchListListDataSourceSchema defines the schema for a list data source.
var ConfigMatchListListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: ConfigMatchListDataSourceSchema.Attributes,
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
