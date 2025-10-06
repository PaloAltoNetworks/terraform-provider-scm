package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: network_services
// This file contains models for the network_services SDK package

// SdwanRules represents the Terraform model for SdwanRules
type SdwanRules struct {
	Tfid                   types.String          `tfsdk:"tfid"`
	Action                 basetypes.ObjectValue `tfsdk:"action"`
	Application            basetypes.ListValue   `tfsdk:"application"`
	Description            basetypes.StringValue `tfsdk:"description"`
	Destination            basetypes.ListValue   `tfsdk:"destination"`
	Device                 basetypes.StringValue `tfsdk:"device"`
	Disabled               basetypes.BoolValue   `tfsdk:"disabled"`
	ErrorCorrectionProfile basetypes.StringValue `tfsdk:"error_correction_profile"`
	Folder                 basetypes.StringValue `tfsdk:"folder"`
	From                   basetypes.ListValue   `tfsdk:"from"`
	Id                     basetypes.StringValue `tfsdk:"id"`
	Name                   basetypes.StringValue `tfsdk:"name"`
	NegateDestination      basetypes.BoolValue   `tfsdk:"negate_destination"`
	NegateSource           basetypes.BoolValue   `tfsdk:"negate_source"`
	PathQualityProfile     basetypes.StringValue `tfsdk:"path_quality_profile"`
	Position               basetypes.StringValue `tfsdk:"position"`
	SaasQualityProfile     basetypes.StringValue `tfsdk:"saas_quality_profile"`
	Service                basetypes.ListValue   `tfsdk:"service"`
	Snippet                basetypes.StringValue `tfsdk:"snippet"`
	Source                 basetypes.ListValue   `tfsdk:"source"`
	SourceUser             basetypes.ListValue   `tfsdk:"source_user"`
	Tag                    basetypes.ListValue   `tfsdk:"tag"`
	To                     basetypes.ListValue   `tfsdk:"to"`
}

// SdwanRulesAction represents a nested structure within the SdwanRules model
type SdwanRulesAction struct {
	TrafficDistributionProfile basetypes.StringValue `tfsdk:"traffic_distribution_profile"`
}

// AttrTypes defines the attribute types for the SdwanRules model.
func (o SdwanRules) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"action": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"traffic_distribution_profile": basetypes.StringType{},
			},
		},
		"application":              basetypes.ListType{ElemType: basetypes.StringType{}},
		"description":              basetypes.StringType{},
		"destination":              basetypes.ListType{ElemType: basetypes.StringType{}},
		"device":                   basetypes.StringType{},
		"disabled":                 basetypes.BoolType{},
		"error_correction_profile": basetypes.StringType{},
		"folder":                   basetypes.StringType{},
		"from":                     basetypes.ListType{ElemType: basetypes.StringType{}},
		"id":                       basetypes.StringType{},
		"name":                     basetypes.StringType{},
		"negate_destination":       basetypes.BoolType{},
		"negate_source":            basetypes.BoolType{},
		"path_quality_profile":     basetypes.StringType{},
		"position":                 basetypes.StringType{},
		"saas_quality_profile":     basetypes.StringType{},
		"service":                  basetypes.ListType{ElemType: basetypes.StringType{}},
		"snippet":                  basetypes.StringType{},
		"source":                   basetypes.ListType{ElemType: basetypes.StringType{}},
		"source_user":              basetypes.ListType{ElemType: basetypes.StringType{}},
		"tag":                      basetypes.ListType{ElemType: basetypes.StringType{}},
		"to":                       basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of SdwanRules objects.
func (o SdwanRules) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SdwanRulesAction model.
func (o SdwanRulesAction) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"traffic_distribution_profile": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of SdwanRulesAction objects.
func (o SdwanRulesAction) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// SdwanRulesResourceSchema defines the schema for SdwanRules resource
var SdwanRulesResourceSchema = schema.Schema{
	MarkdownDescription: "SdwanRule resource",
	Attributes: map[string]schema.Attribute{
		"action": schema.SingleNestedAttribute{
			MarkdownDescription: "Action",
			Required:            true,
			Attributes: map[string]schema.Attribute{
				"traffic_distribution_profile": schema.StringAttribute{
					MarkdownDescription: "Traffic dstribution profile",
					Required:            true,
				},
			},
		},
		"application": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of applications",
			Required:            true,
		},
		"description": schema.StringAttribute{
			MarkdownDescription: "Rule description",
			Optional:            true,
		},
		"destination": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of destination addresses",
			Required:            true,
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
		"disabled": schema.BoolAttribute{
			MarkdownDescription: "Disable rule?",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
		},
		"error_correction_profile": schema.StringAttribute{
			MarkdownDescription: "Error correction profile",
			Optional:            true,
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
			MarkdownDescription: "The folder in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"from": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of source zones",
			Required:            true,
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "Rule name",
			Required:            true,
		},
		"negate_destination": schema.BoolAttribute{
			MarkdownDescription: "Negate destination address(es)?",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
		},
		"negate_source": schema.BoolAttribute{
			MarkdownDescription: "Negate source address(es)?",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
		},
		"path_quality_profile": schema.StringAttribute{
			MarkdownDescription: "Path quality profile",
			Required:            true,
		},
		"position": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("pre", "post"),
			},
			MarkdownDescription: "Rule postion relative to device rules",
			Required:            true,
		},
		"saas_quality_profile": schema.StringAttribute{
			MarkdownDescription: "SaaS quality profile",
			Optional:            true,
		},
		"service": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of services",
			Required:            true,
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
		"source": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of source addresses",
			Required:            true,
		},
		"source_user": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of source users",
			Required:            true,
		},
		"tag": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of tags",
			Optional:            true,
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"to": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of destination zones",
			Required:            true,
		},
	},
}

// SdwanRulesDataSourceSchema defines the schema for SdwanRules data source
var SdwanRulesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "SdwanRule data source",
	Attributes: map[string]dsschema.Attribute{
		"action": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Action",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"traffic_distribution_profile": dsschema.StringAttribute{
					MarkdownDescription: "Traffic dstribution profile",
					Computed:            true,
				},
			},
		},
		"application": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of applications",
			Computed:            true,
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Rule description",
			Computed:            true,
		},
		"destination": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of destination addresses",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"disabled": dsschema.BoolAttribute{
			MarkdownDescription: "Disable rule?",
			Computed:            true,
		},
		"error_correction_profile": dsschema.StringAttribute{
			MarkdownDescription: "Error correction profile",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined",
			Computed:            true,
		},
		"from": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of source zones",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Rule name",
			Optional:            true,
			Computed:            true,
		},
		"negate_destination": dsschema.BoolAttribute{
			MarkdownDescription: "Negate destination address(es)?",
			Computed:            true,
		},
		"negate_source": dsschema.BoolAttribute{
			MarkdownDescription: "Negate source address(es)?",
			Computed:            true,
		},
		"path_quality_profile": dsschema.StringAttribute{
			MarkdownDescription: "Path quality profile",
			Computed:            true,
		},
		"position": dsschema.StringAttribute{
			MarkdownDescription: "Rule postion relative to device rules",
			Computed:            true,
		},
		"saas_quality_profile": dsschema.StringAttribute{
			MarkdownDescription: "SaaS quality profile",
			Computed:            true,
		},
		"service": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of services",
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"source": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of source addresses",
			Computed:            true,
		},
		"source_user": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of source users",
			Computed:            true,
		},
		"tag": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of tags",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"to": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of destination zones",
			Computed:            true,
		},
	},
}

// SdwanRulesListModel represents the data model for a list data source.
type SdwanRulesListModel struct {
	Tfid    types.String `tfsdk:"tfid"`
	Data    []SdwanRules `tfsdk:"data"`
	Limit   types.Int64  `tfsdk:"limit"`
	Offset  types.Int64  `tfsdk:"offset"`
	Name    types.String `tfsdk:"name"`
	Total   types.Int64  `tfsdk:"total"`
	Folder  types.String `tfsdk:"folder"`
	Snippet types.String `tfsdk:"snippet"`
	Device  types.String `tfsdk:"device"`
}

// SdwanRulesListDataSourceSchema defines the schema for a list data source.
var SdwanRulesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: SdwanRulesDataSourceSchema.Attributes,
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
