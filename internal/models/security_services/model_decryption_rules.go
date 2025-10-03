package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: security_services
// This file contains models for the security_services SDK package

// DecryptionRules represents the Terraform model for DecryptionRules
type DecryptionRules struct {
	Tfid              types.String          `tfsdk:"tfid"`
	Action            basetypes.StringValue `tfsdk:"action"`
	Category          basetypes.ListValue   `tfsdk:"category"`
	Description       basetypes.StringValue `tfsdk:"description"`
	Destination       basetypes.ListValue   `tfsdk:"destination"`
	DestinationHip    basetypes.ListValue   `tfsdk:"destination_hip"`
	Device            basetypes.StringValue `tfsdk:"device"`
	Disabled          basetypes.BoolValue   `tfsdk:"disabled"`
	Folder            basetypes.StringValue `tfsdk:"folder"`
	From              basetypes.ListValue   `tfsdk:"from"`
	Id                basetypes.StringValue `tfsdk:"id"`
	LogFail           basetypes.BoolValue   `tfsdk:"log_fail"`
	LogSetting        basetypes.StringValue `tfsdk:"log_setting"`
	LogSuccess        basetypes.BoolValue   `tfsdk:"log_success"`
	Name              basetypes.StringValue `tfsdk:"name"`
	NegateDestination basetypes.BoolValue   `tfsdk:"negate_destination"`
	NegateSource      basetypes.BoolValue   `tfsdk:"negate_source"`
	Profile           basetypes.StringValue `tfsdk:"profile"`
	Service           basetypes.ListValue   `tfsdk:"service"`
	Snippet           basetypes.StringValue `tfsdk:"snippet"`
	Source            basetypes.ListValue   `tfsdk:"source"`
	SourceHip         basetypes.ListValue   `tfsdk:"source_hip"`
	SourceUser        basetypes.ListValue   `tfsdk:"source_user"`
	Tag               basetypes.ListValue   `tfsdk:"tag"`
	To                basetypes.ListValue   `tfsdk:"to"`
	Type              basetypes.ObjectValue `tfsdk:"type"`
	Position          basetypes.StringValue `tfsdk:"position"`
}

// DecryptionRulesType represents a nested structure within the DecryptionRules model
type DecryptionRulesType struct {
	SslForwardProxy      basetypes.ObjectValue `tfsdk:"ssl_forward_proxy"`
	SslInboundInspection basetypes.StringValue `tfsdk:"ssl_inbound_inspection"`
}

// AttrTypes defines the attribute types for the DecryptionRules model.
func (o DecryptionRules) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":               basetypes.StringType{},
		"action":             basetypes.StringType{},
		"category":           basetypes.ListType{ElemType: basetypes.StringType{}},
		"description":        basetypes.StringType{},
		"destination":        basetypes.ListType{ElemType: basetypes.StringType{}},
		"destination_hip":    basetypes.ListType{ElemType: basetypes.StringType{}},
		"device":             basetypes.StringType{},
		"disabled":           basetypes.BoolType{},
		"folder":             basetypes.StringType{},
		"from":               basetypes.ListType{ElemType: basetypes.StringType{}},
		"id":                 basetypes.StringType{},
		"log_fail":           basetypes.BoolType{},
		"log_setting":        basetypes.StringType{},
		"log_success":        basetypes.BoolType{},
		"name":               basetypes.StringType{},
		"negate_destination": basetypes.BoolType{},
		"negate_source":      basetypes.BoolType{},
		"profile":            basetypes.StringType{},
		"service":            basetypes.ListType{ElemType: basetypes.StringType{}},
		"snippet":            basetypes.StringType{},
		"source":             basetypes.ListType{ElemType: basetypes.StringType{}},
		"source_hip":         basetypes.ListType{ElemType: basetypes.StringType{}},
		"source_user":        basetypes.ListType{ElemType: basetypes.StringType{}},
		"tag":                basetypes.ListType{ElemType: basetypes.StringType{}},
		"to":                 basetypes.ListType{ElemType: basetypes.StringType{}},
		"type": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ssl_forward_proxy": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"ssl_inbound_inspection": basetypes.StringType{},
			},
		},
		"position": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DecryptionRules objects.
func (o DecryptionRules) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DecryptionRulesType model.
func (o DecryptionRulesType) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ssl_forward_proxy": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"ssl_inbound_inspection": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DecryptionRulesType objects.
func (o DecryptionRulesType) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// DecryptionRulesResourceSchema defines the schema for DecryptionRules resource
var DecryptionRulesResourceSchema = schema.Schema{
	MarkdownDescription: "DecryptionRule resource",
	Attributes: map[string]schema.Attribute{
		"action": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("decrypt", "no-decrypt"),
			},
			MarkdownDescription: "The action to be taken",
			Required:            true,
		},
		"category": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The destination URL category",
			Required:            true,
		},
		"description": schema.StringAttribute{
			MarkdownDescription: "The description of the decryption rule",
			Optional:            true,
		},
		"destination": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The destination addresses",
			Required:            true,
		},
		"destination_hip": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The Host Integrity Profile of the destination host",
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
		"disabled": schema.BoolAttribute{
			MarkdownDescription: "Is the rule disabled?",
			Optional:            true,
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
		"from": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The source security zone",
			Required:            true,
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the decryption rule",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"log_fail": schema.BoolAttribute{
			MarkdownDescription: "Log failed decryption events?",
			Optional:            true,
		},
		"log_setting": schema.StringAttribute{
			MarkdownDescription: "The log settings of the decryption rule",
			Optional:            true,
		},
		"log_success": schema.BoolAttribute{
			MarkdownDescription: "Log successful decryption events?",
			Optional:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the decryption rule",
			Required:            true,
		},
		"negate_destination": schema.BoolAttribute{
			MarkdownDescription: "Negate the destination addresses?",
			Optional:            true,
		},
		"negate_source": schema.BoolAttribute{
			MarkdownDescription: "Negate the source addresses?",
			Optional:            true,
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
		"profile": schema.StringAttribute{
			MarkdownDescription: "The decryption profile associated with the decryption rule",
			Optional:            true,
		},
		"service": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The destination services and/or service groups",
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
		"source": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The source addresses",
			Required:            true,
		},
		"source_hip": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Source hip",
			Optional:            true,
		},
		"source_user": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of source users and/or groups.  Reserved words include `any`, `pre-login`, `known-user`, and `unknown`.",
			Required:            true,
		},
		"tag": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The tags associated with the decryption rule",
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
			MarkdownDescription: "The destination security zone",
			Required:            true,
		},
		"type": schema.SingleNestedAttribute{
			MarkdownDescription: "The type of decryption",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"ssl_forward_proxy": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("ssl_inbound_inspection"),
						),
					},
					MarkdownDescription: "Ssl forward proxy",
					Optional:            true,
					Attributes:          map[string]schema.Attribute{},
				},
				"ssl_inbound_inspection": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("ssl_forward_proxy"),
						),
					},
					MarkdownDescription: "add the certificate name for SSL inbound inspection",
					Optional:            true,
				},
			},
		},
	},
}

// DecryptionRulesDataSourceSchema defines the schema for DecryptionRules data source
var DecryptionRulesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "DecryptionRule data source",
	Attributes: map[string]dsschema.Attribute{
		"action": dsschema.StringAttribute{
			MarkdownDescription: "The action to be taken",
			Computed:            true,
		},
		"category": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The destination URL category",
			Computed:            true,
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "The description of the decryption rule",
			Computed:            true,
		},
		"destination": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The destination addresses",
			Computed:            true,
		},
		"destination_hip": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The Host Integrity Profile of the destination host",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"disabled": dsschema.BoolAttribute{
			MarkdownDescription: "Is the rule disabled?",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined",
			Computed:            true,
		},
		"from": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The source security zone",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the decryption rule",
			Required:            true,
		},
		"log_fail": dsschema.BoolAttribute{
			MarkdownDescription: "Log failed decryption events?",
			Computed:            true,
		},
		"log_setting": dsschema.StringAttribute{
			MarkdownDescription: "The log settings of the decryption rule",
			Computed:            true,
		},
		"log_success": dsschema.BoolAttribute{
			MarkdownDescription: "Log successful decryption events?",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the decryption rule",
			Optional:            true,
			Computed:            true,
		},
		"negate_destination": dsschema.BoolAttribute{
			MarkdownDescription: "Negate the destination addresses?",
			Computed:            true,
		},
		"negate_source": dsschema.BoolAttribute{
			MarkdownDescription: "Negate the source addresses?",
			Computed:            true,
		},
		"profile": dsschema.StringAttribute{
			MarkdownDescription: "The decryption profile associated with the decryption rule",
			Computed:            true,
		},
		"service": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The destination services and/or service groups",
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"source": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The source addresses",
			Computed:            true,
		},
		"source_hip": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Source hip",
			Computed:            true,
		},
		"source_user": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "List of source users and/or groups.  Reserved words include `any`, `pre-login`, `known-user`, and `unknown`.",
			Computed:            true,
		},
		"tag": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The tags associated with the decryption rule",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"to": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "The destination security zone",
			Computed:            true,
		},
		"type": dsschema.SingleNestedAttribute{
			MarkdownDescription: "The type of decryption",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"ssl_forward_proxy": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Ssl forward proxy",
					Computed:            true,
					Attributes:          map[string]dsschema.Attribute{},
				},
				"ssl_inbound_inspection": dsschema.StringAttribute{
					MarkdownDescription: "add the certificate name for SSL inbound inspection",
					Computed:            true,
				},
			},
		},
	},
}

// DecryptionRulesListModel represents the data model for a list data source.
type DecryptionRulesListModel struct {
	Tfid    types.String      `tfsdk:"tfid"`
	Data    []DecryptionRules `tfsdk:"data"`
	Limit   types.Int64       `tfsdk:"limit"`
	Offset  types.Int64       `tfsdk:"offset"`
	Name    types.String      `tfsdk:"name"`
	Total   types.Int64       `tfsdk:"total"`
	Folder  types.String      `tfsdk:"folder"`
	Snippet types.String      `tfsdk:"snippet"`
	Device  types.String      `tfsdk:"device"`
}

// DecryptionRulesListDataSourceSchema defines the schema for a list data source.
var DecryptionRulesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: DecryptionRulesDataSourceSchema.Attributes,
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
