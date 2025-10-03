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

// Package: objects
// This file contains models for the objects SDK package

// LogForwardingProfiles represents the Terraform model for LogForwardingProfiles
type LogForwardingProfiles struct {
	Tfid        types.String          `tfsdk:"tfid"`
	Description basetypes.StringValue `tfsdk:"description"`
	Device      basetypes.StringValue `tfsdk:"device"`
	Folder      basetypes.StringValue `tfsdk:"folder"`
	Id          basetypes.StringValue `tfsdk:"id"`
	MatchList   basetypes.ListValue   `tfsdk:"match_list"`
	Name        basetypes.StringValue `tfsdk:"name"`
	Snippet     basetypes.StringValue `tfsdk:"snippet"`
}

// LogForwardingProfilesMatchListInner represents a nested structure within the LogForwardingProfiles model
type LogForwardingProfilesMatchListInner struct {
	ActionDesc basetypes.StringValue `tfsdk:"action_desc"`
	Filter     basetypes.StringValue `tfsdk:"filter"`
	LogType    basetypes.StringValue `tfsdk:"log_type"`
	Name       basetypes.StringValue `tfsdk:"name"`
	SendHttp   basetypes.ListValue   `tfsdk:"send_http"`
	SendSyslog basetypes.ListValue   `tfsdk:"send_syslog"`
}

// AttrTypes defines the attribute types for the LogForwardingProfiles model.
func (o LogForwardingProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":        basetypes.StringType{},
		"description": basetypes.StringType{},
		"device":      basetypes.StringType{},
		"folder":      basetypes.StringType{},
		"id":          basetypes.StringType{},
		"match_list": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action_desc": basetypes.StringType{},
				"filter":      basetypes.StringType{},
				"log_type":    basetypes.StringType{},
				"name":        basetypes.StringType{},
				"send_http":   basetypes.ListType{ElemType: basetypes.StringType{}},
				"send_syslog": basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		}},
		"name":    basetypes.StringType{},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogForwardingProfiles objects.
func (o LogForwardingProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogForwardingProfilesMatchListInner model.
func (o LogForwardingProfilesMatchListInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action_desc": basetypes.StringType{},
		"filter":      basetypes.StringType{},
		"log_type":    basetypes.StringType{},
		"name":        basetypes.StringType{},
		"send_http":   basetypes.ListType{ElemType: basetypes.StringType{}},
		"send_syslog": basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of LogForwardingProfilesMatchListInner objects.
func (o LogForwardingProfilesMatchListInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// LogForwardingProfilesResourceSchema defines the schema for LogForwardingProfiles resource
var LogForwardingProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "Manages SCM LogForwardingProfiles objects",
	Attributes: map[string]schema.Attribute{
		"description": schema.StringAttribute{
			MarkdownDescription: "Log forwarding profile description",
			Optional:            true,
		},
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
		"folder": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("snippet"),
					path.MatchRelative().AtParent().AtName("device"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
			},
			MarkdownDescription: "The folder in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the log server profile",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"match_list": schema.ListNestedAttribute{
			MarkdownDescription: "Match list",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"action_desc": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.LengthAtMost(255),
						},
						MarkdownDescription: "Match profile description",
						Optional:            true,
					},
					"filter": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.LengthAtMost(65535),
						},
						MarkdownDescription: "Filter match criteria",
						Optional:            true,
					},
					"log_type": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("traffic", "threat", "wildfire", "url", "data", "tunnel", "auth", "decryption"),
						},
						MarkdownDescription: "Log type",
						Optional:            true,
					},
					"name": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.LengthAtMost(63),
						},
						MarkdownDescription: "Name of the match profile",
						Optional:            true,
					},
					"send_http": schema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "A list of HTTP server profiles",
						Optional:            true,
					},
					"send_syslog": schema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "A list of syslog server profiles",
						Optional:            true,
					},
				},
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(63),
			},
			MarkdownDescription: "The name of the log forwarding profile",
			Required:            true,
		},
		"snippet": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("folder"),
					path.MatchRelative().AtParent().AtName("device"),
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
	},
}

// LogForwardingProfilesDataSourceSchema defines the schema for LogForwardingProfiles data source
var LogForwardingProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "LogForwardingProfiles data source",
	Attributes: map[string]dsschema.Attribute{
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Log forwarding profile description",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the log server profile",
			Required:            true,
		},
		"match_list": dsschema.ListNestedAttribute{
			MarkdownDescription: "Match list",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"action_desc": dsschema.StringAttribute{
						MarkdownDescription: "Match profile description",
						Computed:            true,
					},
					"filter": dsschema.StringAttribute{
						MarkdownDescription: "Filter match criteria",
						Computed:            true,
					},
					"log_type": dsschema.StringAttribute{
						MarkdownDescription: "Log type",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "Name of the match profile",
						Computed:            true,
					},
					"send_http": dsschema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "A list of HTTP server profiles",
						Computed:            true,
					},
					"send_syslog": dsschema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "A list of syslog server profiles",
						Computed:            true,
					},
				},
			},
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the log forwarding profile",
			Optional:            true,
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// LogForwardingProfilesListModel represents the data model for a list data source.
type LogForwardingProfilesListModel struct {
	Tfid    types.String            `tfsdk:"tfid"`
	Data    []LogForwardingProfiles `tfsdk:"data"`
	Limit   types.Int64             `tfsdk:"limit"`
	Offset  types.Int64             `tfsdk:"offset"`
	Name    types.String            `tfsdk:"name"`
	Total   types.Int64             `tfsdk:"total"`
	Folder  types.String            `tfsdk:"folder"`
	Snippet types.String            `tfsdk:"snippet"`
	Device  types.String            `tfsdk:"device"`
}

// LogForwardingProfilesListDataSourceSchema defines the schema for a list data source.
var LogForwardingProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: LogForwardingProfilesDataSourceSchema.Attributes,
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
