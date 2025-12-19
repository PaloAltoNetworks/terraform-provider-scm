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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: security_services
// This file contains models for the security_services SDK package

// AppOverrideRules represents the Terraform model for AppOverrideRules
type AppOverrideRules struct {
	Tfid              types.String          `tfsdk:"tfid"`
	RelativePosition  basetypes.StringValue `tfsdk:"relative_position"`
	TargetRule        basetypes.StringValue `tfsdk:"target_rule"`
	Application       basetypes.StringValue `tfsdk:"application"`
	Description       basetypes.StringValue `tfsdk:"description"`
	Destination       basetypes.ListValue   `tfsdk:"destination"`
	Device            basetypes.StringValue `tfsdk:"device"`
	Disabled          basetypes.BoolValue   `tfsdk:"disabled"`
	Folder            basetypes.StringValue `tfsdk:"folder"`
	From              basetypes.ListValue   `tfsdk:"from"`
	GroupTag          basetypes.StringValue `tfsdk:"group_tag"`
	Id                basetypes.StringValue `tfsdk:"id"`
	Name              basetypes.StringValue `tfsdk:"name"`
	NegateDestination basetypes.BoolValue   `tfsdk:"negate_destination"`
	NegateSource      basetypes.BoolValue   `tfsdk:"negate_source"`
	Port              basetypes.StringValue `tfsdk:"port"`
	Protocol          basetypes.StringValue `tfsdk:"protocol"`
	Snippet           basetypes.StringValue `tfsdk:"snippet"`
	Source            basetypes.ListValue   `tfsdk:"source"`
	Tag               basetypes.ListValue   `tfsdk:"tag"`
	To                basetypes.ListValue   `tfsdk:"to"`
	Position          basetypes.StringValue `tfsdk:"position"`
}

// AttrTypes defines the attribute types for the AppOverrideRules model.
func (o AppOverrideRules) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":               basetypes.StringType{},
		"relative_position":  basetypes.StringType{},
		"target_rule":        basetypes.StringType{},
		"application":        basetypes.StringType{},
		"description":        basetypes.StringType{},
		"destination":        basetypes.ListType{ElemType: basetypes.StringType{}},
		"device":             basetypes.StringType{},
		"disabled":           basetypes.BoolType{},
		"folder":             basetypes.StringType{},
		"from":               basetypes.ListType{ElemType: basetypes.StringType{}},
		"group_tag":          basetypes.StringType{},
		"id":                 basetypes.StringType{},
		"name":               basetypes.StringType{},
		"negate_destination": basetypes.BoolType{},
		"negate_source":      basetypes.BoolType{},
		"port":               basetypes.StringType{},
		"protocol":           basetypes.StringType{},
		"snippet":            basetypes.StringType{},
		"source":             basetypes.ListType{ElemType: basetypes.StringType{}},
		"tag":                basetypes.ListType{ElemType: basetypes.StringType{}},
		"to":                 basetypes.ListType{ElemType: basetypes.StringType{}},
		"position":           basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AppOverrideRules objects.
func (o AppOverrideRules) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AppOverrideRulesResourceSchema defines the schema for AppOverrideRules resource
var AppOverrideRulesResourceSchema = schema.Schema{
	MarkdownDescription: "AppOverrideRule resource",
	Attributes: map[string]schema.Attribute{
		"application": schema.StringAttribute{
			MarkdownDescription: "Application",
			Required:            true,
		},
		"description": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(1024),
			},
			MarkdownDescription: "Description",
			Optional:            true,
		},
		"destination": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Destination",
			Required:            true,
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
			MarkdownDescription: "The device in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"disabled": schema.BoolAttribute{
			MarkdownDescription: "Disabled",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
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
			MarkdownDescription: "The folder in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"from": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "From",
			Required:            true,
		},
		"group_tag": schema.StringAttribute{
			MarkdownDescription: "Group tag",
			Optional:            true,
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
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9._-]+$"), "pattern must match "+"^[a-zA-Z0-9._-]+$"),
			},
			MarkdownDescription: "Name",
			Required:            true,
		},
		"negate_destination": schema.BoolAttribute{
			MarkdownDescription: "Negate destination",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
		},
		"negate_source": schema.BoolAttribute{
			MarkdownDescription: "Negate source",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
		},
		"port": schema.StringAttribute{
			MarkdownDescription: "Port",
			Required:            true,
		},
		"position": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("pre", "post"),
			},
			MarkdownDescription: "The position of a security rule\n",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("pre"),
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"protocol": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("tcp", "udp"),
			},
			MarkdownDescription: "Protocol",
			Required:            true,
		},
		"relative_position": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("before", "after", "top", "bottom"),
			},
			MarkdownDescription: "Relative positioning rule. String must be one of these: `\"before\"`, `\"after\"`, `\"top\"`, `\"bottom\"`. If not specified, rule is created at the bottom of the ruleset.",
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
			MarkdownDescription: "The snippet in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"source": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Source",
			Required:            true,
		},
		"tag": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Tag",
			Optional:            true,
		},
		"target_rule": schema.StringAttribute{
			MarkdownDescription: "The name or UUID of the rule to position this rule relative to. Required when `relative_position` is `\"before\"` or `\"after\"`.",
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
			MarkdownDescription: "To",
			Required:            true,
		},
	},
}

// AppOverrideRulesDataSourceSchema defines the schema for AppOverrideRules data source
var AppOverrideRulesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "AppOverrideRule data source",
	Attributes: map[string]dsschema.Attribute{
		"application": dsschema.StringAttribute{
			MarkdownDescription: "Application",
			Computed:            true,
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Description",
			Computed:            true,
		},
		"destination": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Destination",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"disabled": dsschema.BoolAttribute{
			MarkdownDescription: "Disabled",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"from": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "From",
			Computed:            true,
		},
		"group_tag": dsschema.StringAttribute{
			MarkdownDescription: "Group tag",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Name",
			Optional:            true,
			Computed:            true,
		},
		"negate_destination": dsschema.BoolAttribute{
			MarkdownDescription: "Negate destination",
			Computed:            true,
		},
		"negate_source": dsschema.BoolAttribute{
			MarkdownDescription: "Negate source",
			Computed:            true,
		},
		"port": dsschema.StringAttribute{
			MarkdownDescription: "Port",
			Computed:            true,
		},
		"position": dsschema.StringAttribute{
			MarkdownDescription: "The position of a security rule\n",
			Computed:            true,
		},
		"protocol": dsschema.StringAttribute{
			MarkdownDescription: "Protocol",
			Computed:            true,
		},
		"relative_position": dsschema.StringAttribute{
			MarkdownDescription: "Relative positioning rule. String must be one of these: `\"before\"`, `\"after\"`, `\"top\"`, `\"bottom\"`. If not specified, rule is created at the bottom of the ruleset.",
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"source": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Source",
			Computed:            true,
		},
		"tag": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Tag",
			Computed:            true,
		},
		"target_rule": dsschema.StringAttribute{
			MarkdownDescription: "The name or UUID of the rule to position this rule relative to. Required when `relative_position` is `\"before\"` or `\"after\"`.",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"to": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "To",
			Computed:            true,
		},
	},
}

// AppOverrideRulesListModel represents the data model for a list data source.
type AppOverrideRulesListModel struct {
	Tfid     types.String          `tfsdk:"tfid"`
	Data     []AppOverrideRules    `tfsdk:"data"`
	Limit    types.Int64           `tfsdk:"limit"`
	Offset   types.Int64           `tfsdk:"offset"`
	Name     types.String          `tfsdk:"name"`
	Total    types.Int64           `tfsdk:"total"`
	Folder   types.String          `tfsdk:"folder"`
	Snippet  types.String          `tfsdk:"snippet"`
	Device   types.String          `tfsdk:"device"`
	Position basetypes.StringValue `tfsdk:"position"`
}

// AppOverrideRulesListDataSourceSchema defines the schema for a list data source.
var AppOverrideRulesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: AppOverrideRulesDataSourceSchema.Attributes,
			},
		},
		"limit":   dsschema.Int64Attribute{Description: "The max number of items to return. Default: 200.", Optional: true},
		"offset":  dsschema.Int64Attribute{Description: "The offset of the first item to return.", Optional: true},
		"name":    dsschema.StringAttribute{Description: "The name of the item.", Optional: true},
		"total":   dsschema.Int64Attribute{Description: "The total number of items.", Computed: true},
		"folder":  dsschema.StringAttribute{Description: "The folder of the item. Default: Shared.", Optional: true},
		"snippet": dsschema.StringAttribute{Description: "The snippet of the item.", Optional: true},
		"device":  dsschema.StringAttribute{Description: "The device of the item.", Optional: true},
		"position": dsschema.StringAttribute{
			Description: "The position of a security rule\n",
			Required:    true,
		},
	},
}
