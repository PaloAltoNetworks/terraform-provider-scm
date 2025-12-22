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

// Package: security_services
// This file contains models for the security_services SDK package

// WildfireAntiVirusProfiles represents the Terraform model for WildfireAntiVirusProfiles
type WildfireAntiVirusProfiles struct {
	Tfid            types.String          `tfsdk:"tfid"`
	Description     basetypes.StringValue `tfsdk:"description"`
	Device          basetypes.StringValue `tfsdk:"device"`
	Folder          basetypes.StringValue `tfsdk:"folder"`
	Id              basetypes.StringValue `tfsdk:"id"`
	MlavException   basetypes.ListValue   `tfsdk:"mlav_exception"`
	Name            basetypes.StringValue `tfsdk:"name"`
	PacketCapture   basetypes.BoolValue   `tfsdk:"packet_capture"`
	Rules           basetypes.ListValue   `tfsdk:"rules"`
	Snippet         basetypes.StringValue `tfsdk:"snippet"`
	ThreatException basetypes.ListValue   `tfsdk:"threat_exception"`
}

// WildfireAntiVirusProfilesMlavExceptionInner represents a nested structure within the WildfireAntiVirusProfiles model
type WildfireAntiVirusProfilesMlavExceptionInner struct {
	Description basetypes.StringValue `tfsdk:"description"`
	Filename    basetypes.StringValue `tfsdk:"filename"`
	Name        basetypes.StringValue `tfsdk:"name"`
}

// WildfireAntiVirusProfilesRulesInner represents a nested structure within the WildfireAntiVirusProfiles model
type WildfireAntiVirusProfilesRulesInner struct {
	Analysis    basetypes.StringValue `tfsdk:"analysis"`
	Application basetypes.ListValue   `tfsdk:"application"`
	Direction   basetypes.StringValue `tfsdk:"direction"`
	FileType    basetypes.ListValue   `tfsdk:"file_type"`
	Name        basetypes.StringValue `tfsdk:"name"`
}

// WildfireAntiVirusProfilesThreatExceptionInner represents a nested structure within the WildfireAntiVirusProfiles model
type WildfireAntiVirusProfilesThreatExceptionInner struct {
	Name  basetypes.StringValue `tfsdk:"name"`
	Notes basetypes.StringValue `tfsdk:"notes"`
}

// AttrTypes defines the attribute types for the WildfireAntiVirusProfiles model.
func (o WildfireAntiVirusProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":        basetypes.StringType{},
		"description": basetypes.StringType{},
		"device":      basetypes.StringType{},
		"folder":      basetypes.StringType{},
		"id":          basetypes.StringType{},
		"mlav_exception": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"description": basetypes.StringType{},
				"filename":    basetypes.StringType{},
				"name":        basetypes.StringType{},
			},
		}},
		"name":           basetypes.StringType{},
		"packet_capture": basetypes.BoolType{},
		"rules": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"analysis":    basetypes.StringType{},
				"application": basetypes.ListType{ElemType: basetypes.StringType{}},
				"direction":   basetypes.StringType{},
				"file_type":   basetypes.ListType{ElemType: basetypes.StringType{}},
				"name":        basetypes.StringType{},
			},
		}},
		"snippet": basetypes.StringType{},
		"threat_exception": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":  basetypes.StringType{},
				"notes": basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of WildfireAntiVirusProfiles objects.
func (o WildfireAntiVirusProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the WildfireAntiVirusProfilesMlavExceptionInner model.
func (o WildfireAntiVirusProfilesMlavExceptionInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"description": basetypes.StringType{},
		"filename":    basetypes.StringType{},
		"name":        basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of WildfireAntiVirusProfilesMlavExceptionInner objects.
func (o WildfireAntiVirusProfilesMlavExceptionInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the WildfireAntiVirusProfilesRulesInner model.
func (o WildfireAntiVirusProfilesRulesInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"analysis":    basetypes.StringType{},
		"application": basetypes.ListType{ElemType: basetypes.StringType{}},
		"direction":   basetypes.StringType{},
		"file_type":   basetypes.ListType{ElemType: basetypes.StringType{}},
		"name":        basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of WildfireAntiVirusProfilesRulesInner objects.
func (o WildfireAntiVirusProfilesRulesInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the WildfireAntiVirusProfilesThreatExceptionInner model.
func (o WildfireAntiVirusProfilesThreatExceptionInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":  basetypes.StringType{},
		"notes": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of WildfireAntiVirusProfilesThreatExceptionInner objects.
func (o WildfireAntiVirusProfilesThreatExceptionInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// WildfireAntiVirusProfilesResourceSchema defines the schema for WildfireAntiVirusProfiles resource
var WildfireAntiVirusProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "WildfireAntiVirusProfile resource",
	Attributes: map[string]schema.Attribute{
		"description": schema.StringAttribute{
			MarkdownDescription: "Description",
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
		"folder": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("device"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
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
		"mlav_exception": schema.ListNestedAttribute{
			MarkdownDescription: "Mlav exception",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						MarkdownDescription: "Description",
						Optional:            true,
					},
					"filename": schema.StringAttribute{
						MarkdownDescription: "Filename",
						Optional:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "Name",
						Optional:            true,
					},
				},
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9._-]+$"), "pattern must match "+"^[a-zA-Z0-9._-]+$"),
			},
			MarkdownDescription: "Name",
			Required:            true,
		},
		"packet_capture": schema.BoolAttribute{
			MarkdownDescription: "Packet capture",
			Optional:            true,
		},
		"rules": schema.ListNestedAttribute{
			MarkdownDescription: "Rules",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"analysis": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("public-cloud", "private-cloud"),
						},
						MarkdownDescription: "Analysis",
						Optional:            true,
					},
					"application": schema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "Application",
						Optional:            true,
					},
					"direction": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("download", "upload", "both"),
						},
						MarkdownDescription: "Direction",
						Optional:            true,
					},
					"file_type": schema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "File type",
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
		"threat_exception": schema.ListNestedAttribute{
			MarkdownDescription: "Threat exception",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						MarkdownDescription: "Name",
						Optional:            true,
					},
					"notes": schema.StringAttribute{
						MarkdownDescription: "Notes",
						Optional:            true,
					},
				},
			},
		},
	},
}

// WildfireAntiVirusProfilesDataSourceSchema defines the schema for WildfireAntiVirusProfiles data source
var WildfireAntiVirusProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "WildfireAntiVirusProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Description",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
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
		"mlav_exception": dsschema.ListNestedAttribute{
			MarkdownDescription: "Mlav exception",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"description": dsschema.StringAttribute{
						MarkdownDescription: "Description",
						Computed:            true,
					},
					"filename": dsschema.StringAttribute{
						MarkdownDescription: "Filename",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "Name",
						Computed:            true,
					},
				},
			},
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Name",
			Optional:            true,
			Computed:            true,
		},
		"packet_capture": dsschema.BoolAttribute{
			MarkdownDescription: "Packet capture",
			Computed:            true,
		},
		"rules": dsschema.ListNestedAttribute{
			MarkdownDescription: "Rules",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"analysis": dsschema.StringAttribute{
						MarkdownDescription: "Analysis",
						Computed:            true,
					},
					"application": dsschema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "Application",
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
		"threat_exception": dsschema.ListNestedAttribute{
			MarkdownDescription: "Threat exception",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"name": dsschema.StringAttribute{
						MarkdownDescription: "Name",
						Computed:            true,
					},
					"notes": dsschema.StringAttribute{
						MarkdownDescription: "Notes",
						Computed:            true,
					},
				},
			},
		},
	},
}

// WildfireAntiVirusProfilesListModel represents the data model for a list data source.
type WildfireAntiVirusProfilesListModel struct {
	Tfid    types.String                `tfsdk:"tfid"`
	Data    []WildfireAntiVirusProfiles `tfsdk:"data"`
	Limit   types.Int64                 `tfsdk:"limit"`
	Offset  types.Int64                 `tfsdk:"offset"`
	Name    types.String                `tfsdk:"name"`
	Total   types.Int64                 `tfsdk:"total"`
	Folder  types.String                `tfsdk:"folder"`
	Snippet types.String                `tfsdk:"snippet"`
	Device  types.String                `tfsdk:"device"`
}

// WildfireAntiVirusProfilesListDataSourceSchema defines the schema for a list data source.
var WildfireAntiVirusProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: WildfireAntiVirusProfilesDataSourceSchema.Attributes,
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
