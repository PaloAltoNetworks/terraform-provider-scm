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

// Package: security_services
// This file contains models for the security_services SDK package

// DataFilteringProfiles represents the Terraform model for DataFilteringProfiles
type DataFilteringProfiles struct {
	Tfid            types.String          `tfsdk:"tfid"`
	DataCapture     basetypes.BoolValue   `tfsdk:"data_capture"`
	Description     basetypes.StringValue `tfsdk:"description"`
	Device          basetypes.StringValue `tfsdk:"device"`
	DisableOverride basetypes.StringValue `tfsdk:"disable_override"`
	Folder          basetypes.StringValue `tfsdk:"folder"`
	Id              basetypes.StringValue `tfsdk:"id"`
	Name            basetypes.StringValue `tfsdk:"name"`
	Rules           basetypes.ListValue   `tfsdk:"rules"`
	Snippet         basetypes.StringValue `tfsdk:"snippet"`
}

// DataFilteringProfilesRulesInner represents a nested structure within the DataFilteringProfiles model
type DataFilteringProfilesRulesInner struct {
	AlertThreshold basetypes.Int64Value  `tfsdk:"alert_threshold"`
	Application    basetypes.ListValue   `tfsdk:"application"`
	BlockThreshold basetypes.Int64Value  `tfsdk:"block_threshold"`
	DataObject     basetypes.StringValue `tfsdk:"data_object"`
	Direction      basetypes.StringValue `tfsdk:"direction"`
	FileType       basetypes.ListValue   `tfsdk:"file_type"`
	LogSeverity    basetypes.StringValue `tfsdk:"log_severity"`
	Name           basetypes.StringValue `tfsdk:"name"`
}

// AttrTypes defines the attribute types for the DataFilteringProfiles model.
func (o DataFilteringProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":             basetypes.StringType{},
		"data_capture":     basetypes.BoolType{},
		"description":      basetypes.StringType{},
		"device":           basetypes.StringType{},
		"disable_override": basetypes.StringType{},
		"folder":           basetypes.StringType{},
		"id":               basetypes.StringType{},
		"name":             basetypes.StringType{},
		"rules": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"alert_threshold": basetypes.Int64Type{},
				"application":     basetypes.ListType{ElemType: basetypes.StringType{}},
				"block_threshold": basetypes.Int64Type{},
				"data_object":     basetypes.StringType{},
				"direction":       basetypes.StringType{},
				"file_type":       basetypes.ListType{ElemType: basetypes.StringType{}},
				"log_severity":    basetypes.StringType{},
				"name":            basetypes.StringType{},
			},
		}},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DataFilteringProfiles objects.
func (o DataFilteringProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DataFilteringProfilesRulesInner model.
func (o DataFilteringProfilesRulesInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"alert_threshold": basetypes.Int64Type{},
		"application":     basetypes.ListType{ElemType: basetypes.StringType{}},
		"block_threshold": basetypes.Int64Type{},
		"data_object":     basetypes.StringType{},
		"direction":       basetypes.StringType{},
		"file_type":       basetypes.ListType{ElemType: basetypes.StringType{}},
		"log_severity":    basetypes.StringType{},
		"name":            basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DataFilteringProfilesRulesInner objects.
func (o DataFilteringProfilesRulesInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// DataFilteringProfilesResourceSchema defines the schema for DataFilteringProfiles resource
var DataFilteringProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "DataFilteringProfile resource",
	Attributes: map[string]schema.Attribute{
		"data_capture": schema.BoolAttribute{
			MarkdownDescription: "Data capture",
			Optional:            true,
		},
		"description": schema.StringAttribute{
			MarkdownDescription: "The description of the data filtering profile",
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
		"disable_override": schema.StringAttribute{
			MarkdownDescription: "Disable override",
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
			MarkdownDescription: "The UUID of the data filtering profile",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the data filtering profile",
			Optional:            true,
		},
		"rules": schema.ListNestedAttribute{
			MarkdownDescription: "Rules",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"alert_threshold": schema.Int64Attribute{
						MarkdownDescription: "Alert threshold",
						Optional:            true,
					},
					"application": schema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "Application",
						Optional:            true,
					},
					"block_threshold": schema.Int64Attribute{
						MarkdownDescription: "Block threshold",
						Optional:            true,
					},
					"data_object": schema.StringAttribute{
						MarkdownDescription: "Data object",
						Optional:            true,
					},
					"direction": schema.StringAttribute{
						MarkdownDescription: "Direction",
						Optional:            true,
					},
					"file_type": schema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "File type",
						Optional:            true,
					},
					"log_severity": schema.StringAttribute{
						MarkdownDescription: "Log severity",
						Optional:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "Name",
						Optional:            true,
					},
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

// DataFilteringProfilesDataSourceSchema defines the schema for DataFilteringProfiles data source
var DataFilteringProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "DataFilteringProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"data_capture": dsschema.BoolAttribute{
			MarkdownDescription: "Data capture",
			Computed:            true,
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "The description of the data filtering profile",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"disable_override": dsschema.StringAttribute{
			MarkdownDescription: "Disable override",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the data filtering profile",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the data filtering profile",
			Optional:            true,
			Computed:            true,
		},
		"rules": dsschema.ListNestedAttribute{
			MarkdownDescription: "Rules",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"alert_threshold": dsschema.Int64Attribute{
						MarkdownDescription: "Alert threshold",
						Computed:            true,
					},
					"application": dsschema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "Application",
						Computed:            true,
					},
					"block_threshold": dsschema.Int64Attribute{
						MarkdownDescription: "Block threshold",
						Computed:            true,
					},
					"data_object": dsschema.StringAttribute{
						MarkdownDescription: "Data object",
						Computed:            true,
					},
					"direction": dsschema.StringAttribute{
						MarkdownDescription: "Direction",
						Computed:            true,
					},
					"file_type": dsschema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "File type",
						Computed:            true,
					},
					"log_severity": dsschema.StringAttribute{
						MarkdownDescription: "Log severity",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "Name",
						Computed:            true,
					},
				},
			},
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

// DataFilteringProfilesListModel represents the data model for a list data source.
type DataFilteringProfilesListModel struct {
	Tfid    types.String            `tfsdk:"tfid"`
	Data    []DataFilteringProfiles `tfsdk:"data"`
	Limit   types.Int64             `tfsdk:"limit"`
	Offset  types.Int64             `tfsdk:"offset"`
	Name    types.String            `tfsdk:"name"`
	Total   types.Int64             `tfsdk:"total"`
	Folder  types.String            `tfsdk:"folder"`
	Snippet types.String            `tfsdk:"snippet"`
	Device  types.String            `tfsdk:"device"`
}

// DataFilteringProfilesListDataSourceSchema defines the schema for a list data source.
var DataFilteringProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: DataFilteringProfilesDataSourceSchema.Attributes,
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
