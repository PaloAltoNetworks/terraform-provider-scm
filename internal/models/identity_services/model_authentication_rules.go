package models

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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

// Package: identity_services
// This file contains models for the identity_services SDK package

// AuthenticationRules represents the Terraform model for AuthenticationRules
type AuthenticationRules struct {
	Tfid                      types.String          `tfsdk:"tfid"`
	AuthenticationEnforcement basetypes.StringValue `tfsdk:"authentication_enforcement"`
	Category                  basetypes.ListValue   `tfsdk:"category"`
	Description               basetypes.StringValue `tfsdk:"description"`
	Destination               basetypes.ListValue   `tfsdk:"destination"`
	DestinationHip            basetypes.ListValue   `tfsdk:"destination_hip"`
	Device                    basetypes.StringValue `tfsdk:"device"`
	Disabled                  basetypes.BoolValue   `tfsdk:"disabled"`
	Folder                    basetypes.StringValue `tfsdk:"folder"`
	From                      basetypes.ListValue   `tfsdk:"from"`
	GroupTag                  basetypes.StringValue `tfsdk:"group_tag"`
	HipProfiles               basetypes.ListValue   `tfsdk:"hip_profiles"`
	Id                        basetypes.StringValue `tfsdk:"id"`
	LogAuthenticationTimeout  basetypes.BoolValue   `tfsdk:"log_authentication_timeout"`
	LogSetting                basetypes.StringValue `tfsdk:"log_setting"`
	Name                      basetypes.StringValue `tfsdk:"name"`
	NegateDestination         basetypes.BoolValue   `tfsdk:"negate_destination"`
	NegateSource              basetypes.BoolValue   `tfsdk:"negate_source"`
	Service                   basetypes.ListValue   `tfsdk:"service"`
	Snippet                   basetypes.StringValue `tfsdk:"snippet"`
	Source                    basetypes.ListValue   `tfsdk:"source"`
	SourceHip                 basetypes.ListValue   `tfsdk:"source_hip"`
	SourceUser                basetypes.ListValue   `tfsdk:"source_user"`
	Tag                       basetypes.ListValue   `tfsdk:"tag"`
	Timeout                   basetypes.Int64Value  `tfsdk:"timeout"`
	To                        basetypes.ListValue   `tfsdk:"to"`
	Position                  basetypes.StringValue `tfsdk:"position"`
}

// AttrTypes defines the attribute types for the AuthenticationRules model.
func (o AuthenticationRules) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":                       basetypes.StringType{},
		"authentication_enforcement": basetypes.StringType{},
		"category":                   basetypes.ListType{ElemType: basetypes.StringType{}},
		"description":                basetypes.StringType{},
		"destination":                basetypes.ListType{ElemType: basetypes.StringType{}},
		"destination_hip":            basetypes.ListType{ElemType: basetypes.StringType{}},
		"device":                     basetypes.StringType{},
		"disabled":                   basetypes.BoolType{},
		"folder":                     basetypes.StringType{},
		"from":                       basetypes.ListType{ElemType: basetypes.StringType{}},
		"group_tag":                  basetypes.StringType{},
		"hip_profiles":               basetypes.ListType{ElemType: basetypes.StringType{}},
		"id":                         basetypes.StringType{},
		"log_authentication_timeout": basetypes.BoolType{},
		"log_setting":                basetypes.StringType{},
		"name":                       basetypes.StringType{},
		"negate_destination":         basetypes.BoolType{},
		"negate_source":              basetypes.BoolType{},
		"service":                    basetypes.ListType{ElemType: basetypes.StringType{}},
		"snippet":                    basetypes.StringType{},
		"source":                     basetypes.ListType{ElemType: basetypes.StringType{}},
		"source_hip":                 basetypes.ListType{ElemType: basetypes.StringType{}},
		"source_user":                basetypes.ListType{ElemType: basetypes.StringType{}},
		"tag":                        basetypes.ListType{ElemType: basetypes.StringType{}},
		"timeout":                    basetypes.Int64Type{},
		"to":                         basetypes.ListType{ElemType: basetypes.StringType{}},
		"position":                   basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AuthenticationRules objects.
func (o AuthenticationRules) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AuthenticationRulesResourceSchema defines the schema for AuthenticationRules resource
var AuthenticationRulesResourceSchema = schema.Schema{
	MarkdownDescription: "Manages SCM AuthenticationRules objects",
	Attributes: map[string]schema.Attribute{
		"authentication_enforcement": schema.StringAttribute{
			MarkdownDescription: "The authentication profile name",
			Optional:            true,
		},
		"category": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The destination URL categories",
			Optional:            true,
		},
		"description": schema.StringAttribute{
			MarkdownDescription: "The description of the authentication rule",
			Optional:            true,
		},
		"destination": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The destination addresses",
			Required:            true,
		},
		"destination_hip": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The destination Host Integrity Profile (HIP)",
			Optional:            true,
		},
		"device": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("folder"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
			},
			MarkdownDescription: "Device",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"disabled": schema.BoolAttribute{
			MarkdownDescription: "Is the authentication rule disabled?",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
		},
		"folder": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("snippet"),
					path.MatchRelative().AtParent().AtName("device"),
				),
			},
			MarkdownDescription: "Folder",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"from": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The source security zones",
			Required:            true,
		},
		"group_tag": schema.StringAttribute{
			MarkdownDescription: "Group tag",
			Optional:            true,
		},
		"hip_profiles": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The source Host Integrity Profile (HIP)",
			Optional:            true,
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the authentication rule",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"log_authentication_timeout": schema.BoolAttribute{
			MarkdownDescription: "Log authentication timeouts?",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
		},
		"log_setting": schema.StringAttribute{
			MarkdownDescription: "The log forwarding profile name",
			Optional:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the authentication rule",
			Required:            true,
		},
		"negate_destination": schema.BoolAttribute{
			MarkdownDescription: "Are the destination addresses negated?",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
		},
		"negate_source": schema.BoolAttribute{
			MarkdownDescription: "Are the source addresses negated?",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
		},
		"position": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("pre", "post"),
			},
			MarkdownDescription: "The relative position of the rule\n",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("pre"),
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"service": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The destination ports",
			Required:            true,
		},
		"snippet": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("folder"),
					path.MatchRelative().AtParent().AtName("device"),
				),
			},
			MarkdownDescription: "Snippet",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"source": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The source addresses",
			Required:            true,
		},
		"source_hip": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The source Host Integrity Profile (HIP)",
			Optional:            true,
		},
		"source_user": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The source users",
			Optional:            true,
		},
		"tag": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The authentication rule tags",
			Optional:            true,
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"timeout": schema.Int64Attribute{
			Validators: []validator.Int64{
				int64validator.Between(1, 1440),
			},
			MarkdownDescription: "The authentication session timeout (seconds)",
			Optional:            true,
		},
		"to": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The destination security zones",
			Required:            true,
		},
	},
}

// AuthenticationRulesDataSourceSchema defines the schema for AuthenticationRules data source
var AuthenticationRulesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "AuthenticationRules data source",
	Attributes: map[string]dsschema.Attribute{
		"authentication_enforcement": dsschema.StringAttribute{
			MarkdownDescription: "The authentication profile name",
			Computed:            true,
		},
		"category": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The destination URL categories",
			Computed:            true,
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "The description of the authentication rule",
			Computed:            true,
		},
		"destination": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The destination addresses",
			Computed:            true,
		},
		"destination_hip": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The destination Host Integrity Profile (HIP)",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "Device",
			Computed:            true,
		},
		"disabled": dsschema.BoolAttribute{
			MarkdownDescription: "Is the authentication rule disabled?",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "Folder",
			Computed:            true,
		},
		"from": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The source security zones",
			Computed:            true,
		},
		"group_tag": dsschema.StringAttribute{
			MarkdownDescription: "Group tag",
			Computed:            true,
		},
		"hip_profiles": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The source Host Integrity Profile (HIP)",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the authentication rule",
			Required:            true,
		},
		"log_authentication_timeout": dsschema.BoolAttribute{
			MarkdownDescription: "Log authentication timeouts?",
			Computed:            true,
		},
		"log_setting": dsschema.StringAttribute{
			MarkdownDescription: "The log forwarding profile name",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the authentication rule",
			Optional:            true,
			Computed:            true,
		},
		"negate_destination": dsschema.BoolAttribute{
			MarkdownDescription: "Are the destination addresses negated?",
			Computed:            true,
		},
		"negate_source": dsschema.BoolAttribute{
			MarkdownDescription: "Are the source addresses negated?",
			Computed:            true,
		},
		"service": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The destination ports",
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "Snippet",
			Computed:            true,
		},
		"source": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The source addresses",
			Computed:            true,
		},
		"source_hip": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The source Host Integrity Profile (HIP)",
			Computed:            true,
		},
		"source_user": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The source users",
			Computed:            true,
		},
		"tag": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The authentication rule tags",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"timeout": dsschema.Int64Attribute{
			MarkdownDescription: "The authentication session timeout (seconds)",
			Computed:            true,
		},
		"to": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The destination security zones",
			Computed:            true,
		},
	},
}

// AuthenticationRulesListModel represents the data model for a list data source.
type AuthenticationRulesListModel struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Data    []AuthenticationRules `tfsdk:"data"`
	Limit   types.Int64           `tfsdk:"limit"`
	Offset  types.Int64           `tfsdk:"offset"`
	Name    types.String          `tfsdk:"name"`
	Total   types.Int64           `tfsdk:"total"`
	Folder  types.String          `tfsdk:"folder"`
	Snippet types.String          `tfsdk:"snippet"`
	Device  types.String          `tfsdk:"device"`
}

// AuthenticationRulesListDataSourceSchema defines the schema for a list data source.
var AuthenticationRulesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: AuthenticationRulesDataSourceSchema.Attributes,
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
