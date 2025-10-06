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

// UrlAccessProfiles represents the Terraform model for UrlAccessProfiles
type UrlAccessProfiles struct {
	Tfid                  types.String          `tfsdk:"tfid"`
	Alert                 basetypes.ListValue   `tfsdk:"alert"`
	Allow                 basetypes.ListValue   `tfsdk:"allow"`
	Block                 basetypes.ListValue   `tfsdk:"block"`
	CloudInlineCat        basetypes.BoolValue   `tfsdk:"cloud_inline_cat"`
	Continue              basetypes.ListValue   `tfsdk:"continue"`
	CredentialEnforcement basetypes.ObjectValue `tfsdk:"credential_enforcement"`
	Description           basetypes.StringValue `tfsdk:"description"`
	Device                basetypes.StringValue `tfsdk:"device"`
	Folder                basetypes.StringValue `tfsdk:"folder"`
	Id                    basetypes.StringValue `tfsdk:"id"`
	LocalInlineCat        basetypes.BoolValue   `tfsdk:"local_inline_cat"`
	LogContainerPageOnly  basetypes.BoolValue   `tfsdk:"log_container_page_only"`
	LogHttpHdrReferer     basetypes.BoolValue   `tfsdk:"log_http_hdr_referer"`
	LogHttpHdrUserAgent   basetypes.BoolValue   `tfsdk:"log_http_hdr_user_agent"`
	LogHttpHdrXff         basetypes.BoolValue   `tfsdk:"log_http_hdr_xff"`
	MlavCategoryException basetypes.ListValue   `tfsdk:"mlav_category_exception"`
	Name                  basetypes.StringValue `tfsdk:"name"`
	Redirect              basetypes.ListValue   `tfsdk:"redirect"`
	SafeSearchEnforcement basetypes.BoolValue   `tfsdk:"safe_search_enforcement"`
	Snippet               basetypes.StringValue `tfsdk:"snippet"`
}

// UrlAccessProfilesCredentialEnforcement represents a nested structure within the UrlAccessProfiles model
type UrlAccessProfilesCredentialEnforcement struct {
	Alert       basetypes.ListValue   `tfsdk:"alert"`
	Allow       basetypes.ListValue   `tfsdk:"allow"`
	Block       basetypes.ListValue   `tfsdk:"block"`
	Continue    basetypes.ListValue   `tfsdk:"continue"`
	LogSeverity basetypes.StringValue `tfsdk:"log_severity"`
	Mode        basetypes.ObjectValue `tfsdk:"mode"`
}

// UrlAccessProfilesCredentialEnforcementMode represents a nested structure within the UrlAccessProfiles model
type UrlAccessProfilesCredentialEnforcementMode struct {
	Disabled          basetypes.ObjectValue `tfsdk:"disabled"`
	DomainCredentials basetypes.ObjectValue `tfsdk:"domain_credentials"`
	GroupMapping      basetypes.StringValue `tfsdk:"group_mapping"`
	IpUser            basetypes.ObjectValue `tfsdk:"ip_user"`
}

// AttrTypes defines the attribute types for the UrlAccessProfiles model.
func (o UrlAccessProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":             basetypes.StringType{},
		"alert":            basetypes.ListType{ElemType: basetypes.StringType{}},
		"allow":            basetypes.ListType{ElemType: basetypes.StringType{}},
		"block":            basetypes.ListType{ElemType: basetypes.StringType{}},
		"cloud_inline_cat": basetypes.BoolType{},
		"continue":         basetypes.ListType{ElemType: basetypes.StringType{}},
		"credential_enforcement": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"alert":        basetypes.ListType{ElemType: basetypes.StringType{}},
				"allow":        basetypes.ListType{ElemType: basetypes.StringType{}},
				"block":        basetypes.ListType{ElemType: basetypes.StringType{}},
				"continue":     basetypes.ListType{ElemType: basetypes.StringType{}},
				"log_severity": basetypes.StringType{},
				"mode": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"disabled": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"domain_credentials": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"group_mapping": basetypes.StringType{},
						"ip_user": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
			},
		},
		"description":             basetypes.StringType{},
		"device":                  basetypes.StringType{},
		"folder":                  basetypes.StringType{},
		"id":                      basetypes.StringType{},
		"local_inline_cat":        basetypes.BoolType{},
		"log_container_page_only": basetypes.BoolType{},
		"log_http_hdr_referer":    basetypes.BoolType{},
		"log_http_hdr_user_agent": basetypes.BoolType{},
		"log_http_hdr_xff":        basetypes.BoolType{},
		"mlav_category_exception": basetypes.ListType{ElemType: basetypes.StringType{}},
		"name":                    basetypes.StringType{},
		"redirect":                basetypes.ListType{ElemType: basetypes.StringType{}},
		"safe_search_enforcement": basetypes.BoolType{},
		"snippet":                 basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of UrlAccessProfiles objects.
func (o UrlAccessProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the UrlAccessProfilesCredentialEnforcement model.
func (o UrlAccessProfilesCredentialEnforcement) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"alert":        basetypes.ListType{ElemType: basetypes.StringType{}},
		"allow":        basetypes.ListType{ElemType: basetypes.StringType{}},
		"block":        basetypes.ListType{ElemType: basetypes.StringType{}},
		"continue":     basetypes.ListType{ElemType: basetypes.StringType{}},
		"log_severity": basetypes.StringType{},
		"mode": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"disabled": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"domain_credentials": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"group_mapping": basetypes.StringType{},
				"ip_user": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of UrlAccessProfilesCredentialEnforcement objects.
func (o UrlAccessProfilesCredentialEnforcement) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the UrlAccessProfilesCredentialEnforcementMode model.
func (o UrlAccessProfilesCredentialEnforcementMode) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"disabled": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"domain_credentials": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"group_mapping": basetypes.StringType{},
		"ip_user": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of UrlAccessProfilesCredentialEnforcementMode objects.
func (o UrlAccessProfilesCredentialEnforcementMode) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// UrlAccessProfilesResourceSchema defines the schema for UrlAccessProfiles resource
var UrlAccessProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "UrlAccessProfile resource",
	Attributes: map[string]schema.Attribute{
		"alert": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Alert",
			Optional:            true,
		},
		"allow": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Allow",
			Optional:            true,
		},
		"block": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Block",
			Optional:            true,
		},
		"cloud_inline_cat": schema.BoolAttribute{
			MarkdownDescription: "Cloud inline cat",
			Optional:            true,
		},
		"continue": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Continue",
			Optional:            true,
		},
		"credential_enforcement": schema.SingleNestedAttribute{
			MarkdownDescription: "Credential enforcement",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"alert": schema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Alert",
					Optional:            true,
					Computed:            true,
				},
				"allow": schema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Allow",
					Optional:            true,
					Computed:            true,
				},
				"block": schema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Block",
					Optional:            true,
					Computed:            true,
				},
				"continue": schema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Continue",
					Optional:            true,
					Computed:            true,
				},
				"log_severity": schema.StringAttribute{
					MarkdownDescription: "Log severity",
					Optional:            true,
					Computed:            true,
					Default:             stringdefault.StaticString("medium"),
				},
				"mode": schema.SingleNestedAttribute{
					MarkdownDescription: "Mode",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"disabled": schema.SingleNestedAttribute{
							MarkdownDescription: "Disabled",
							Optional:            true,
							Computed:            true,
							Attributes:          map[string]schema.Attribute{},
						},
						"domain_credentials": schema.SingleNestedAttribute{
							MarkdownDescription: "Domain credentials",
							Optional:            true,
							Computed:            true,
							Attributes:          map[string]schema.Attribute{},
						},
						"group_mapping": schema.StringAttribute{
							MarkdownDescription: "Group mapping",
							Optional:            true,
							Computed:            true,
						},
						"ip_user": schema.SingleNestedAttribute{
							MarkdownDescription: "Ip user",
							Optional:            true,
							Computed:            true,
							Attributes:          map[string]schema.Attribute{},
						},
					},
				},
			},
		},
		"description": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(255),
			},
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
		"id": schema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"local_inline_cat": schema.BoolAttribute{
			MarkdownDescription: "Local inline cat",
			Optional:            true,
		},
		"log_container_page_only": schema.BoolAttribute{
			MarkdownDescription: "Log container page only",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(true),
		},
		"log_http_hdr_referer": schema.BoolAttribute{
			MarkdownDescription: "Log http hdr referer",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
		},
		"log_http_hdr_user_agent": schema.BoolAttribute{
			MarkdownDescription: "Log http hdr user agent",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
		},
		"log_http_hdr_xff": schema.BoolAttribute{
			MarkdownDescription: "Log http hdr xff",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
		},
		"mlav_category_exception": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Mlav category exception",
			Optional:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "Name",
			Required:            true,
		},
		"redirect": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Redirect",
			Optional:            true,
		},
		"safe_search_enforcement": schema.BoolAttribute{
			MarkdownDescription: "Safe search enforcement",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
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
	},
}

// UrlAccessProfilesDataSourceSchema defines the schema for UrlAccessProfiles data source
var UrlAccessProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "UrlAccessProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"alert": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Alert",
			Computed:            true,
		},
		"allow": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Allow",
			Computed:            true,
		},
		"block": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Block",
			Computed:            true,
		},
		"cloud_inline_cat": dsschema.BoolAttribute{
			MarkdownDescription: "Cloud inline cat",
			Computed:            true,
		},
		"continue": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Continue",
			Computed:            true,
		},
		"credential_enforcement": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Credential enforcement",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"alert": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Alert",
					Computed:            true,
				},
				"allow": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Allow",
					Computed:            true,
				},
				"block": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Block",
					Computed:            true,
				},
				"continue": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Continue",
					Computed:            true,
				},
				"log_severity": dsschema.StringAttribute{
					MarkdownDescription: "Log severity",
					Computed:            true,
				},
				"mode": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Mode",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"disabled": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Disabled",
							Computed:            true,
							Attributes:          map[string]dsschema.Attribute{},
						},
						"domain_credentials": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Domain credentials",
							Computed:            true,
							Attributes:          map[string]dsschema.Attribute{},
						},
						"group_mapping": dsschema.StringAttribute{
							MarkdownDescription: "Group mapping",
							Computed:            true,
						},
						"ip_user": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Ip user",
							Computed:            true,
							Attributes:          map[string]dsschema.Attribute{},
						},
					},
				},
			},
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Description",
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
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"local_inline_cat": dsschema.BoolAttribute{
			MarkdownDescription: "Local inline cat",
			Computed:            true,
		},
		"log_container_page_only": dsschema.BoolAttribute{
			MarkdownDescription: "Log container page only",
			Computed:            true,
		},
		"log_http_hdr_referer": dsschema.BoolAttribute{
			MarkdownDescription: "Log http hdr referer",
			Computed:            true,
		},
		"log_http_hdr_user_agent": dsschema.BoolAttribute{
			MarkdownDescription: "Log http hdr user agent",
			Computed:            true,
		},
		"log_http_hdr_xff": dsschema.BoolAttribute{
			MarkdownDescription: "Log http hdr xff",
			Computed:            true,
		},
		"mlav_category_exception": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Mlav category exception",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Name",
			Optional:            true,
			Computed:            true,
		},
		"redirect": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Redirect",
			Computed:            true,
		},
		"safe_search_enforcement": dsschema.BoolAttribute{
			MarkdownDescription: "Safe search enforcement",
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

// UrlAccessProfilesListModel represents the data model for a list data source.
type UrlAccessProfilesListModel struct {
	Tfid    types.String        `tfsdk:"tfid"`
	Data    []UrlAccessProfiles `tfsdk:"data"`
	Limit   types.Int64         `tfsdk:"limit"`
	Offset  types.Int64         `tfsdk:"offset"`
	Name    types.String        `tfsdk:"name"`
	Total   types.Int64         `tfsdk:"total"`
	Folder  types.String        `tfsdk:"folder"`
	Snippet types.String        `tfsdk:"snippet"`
	Device  types.String        `tfsdk:"device"`
}

// UrlAccessProfilesListDataSourceSchema defines the schema for a list data source.
var UrlAccessProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: UrlAccessProfilesDataSourceSchema.Attributes,
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
