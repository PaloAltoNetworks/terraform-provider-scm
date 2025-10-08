package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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

// Package: identity_services
// This file contains models for the identity_services SDK package

// SamlServerProfiles represents the Terraform model for SamlServerProfiles
type SamlServerProfiles struct {
	Tfid                   types.String          `tfsdk:"tfid"`
	Certificate            basetypes.StringValue `tfsdk:"certificate"`
	Device                 basetypes.StringValue `tfsdk:"device"`
	EntityId               basetypes.StringValue `tfsdk:"entity_id"`
	Folder                 basetypes.StringValue `tfsdk:"folder"`
	Id                     basetypes.StringValue `tfsdk:"id"`
	MaxClockSkew           basetypes.Int64Value  `tfsdk:"max_clock_skew"`
	Name                   basetypes.StringValue `tfsdk:"name"`
	SloBindings            basetypes.StringValue `tfsdk:"slo_bindings"`
	Snippet                basetypes.StringValue `tfsdk:"snippet"`
	SsoBindings            basetypes.StringValue `tfsdk:"sso_bindings"`
	SsoUrl                 basetypes.StringValue `tfsdk:"sso_url"`
	ValidateIdpCertificate basetypes.BoolValue   `tfsdk:"validate_idp_certificate"`
	WantAuthRequestsSigned basetypes.BoolValue   `tfsdk:"want_auth_requests_signed"`
}

// AttrTypes defines the attribute types for the SamlServerProfiles model.
func (o SamlServerProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":                      basetypes.StringType{},
		"certificate":               basetypes.StringType{},
		"device":                    basetypes.StringType{},
		"entity_id":                 basetypes.StringType{},
		"folder":                    basetypes.StringType{},
		"id":                        basetypes.StringType{},
		"max_clock_skew":            basetypes.Int64Type{},
		"name":                      basetypes.StringType{},
		"slo_bindings":              basetypes.StringType{},
		"snippet":                   basetypes.StringType{},
		"sso_bindings":              basetypes.StringType{},
		"sso_url":                   basetypes.StringType{},
		"validate_idp_certificate":  basetypes.BoolType{},
		"want_auth_requests_signed": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of SamlServerProfiles objects.
func (o SamlServerProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// SamlServerProfilesResourceSchema defines the schema for SamlServerProfiles resource
var SamlServerProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "SamlServerProfile resource",
	Attributes: map[string]schema.Attribute{
		"certificate": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(63),
			},
			MarkdownDescription: "The identity provider certificate",
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
			MarkdownDescription: "The device in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"entity_id": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(1024),
				stringvalidator.LengthAtLeast(1),
			},
			MarkdownDescription: "The identity provider ID",
			Required:            true,
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
			MarkdownDescription: "The folder in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the SAML server profile",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"max_clock_skew": schema.Int64Attribute{
			Validators: []validator.Int64{
				int64validator.Between(1, 900),
			},
			MarkdownDescription: "Maxiumum clock skew",
			Optional:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the SAML server profile",
			Required:            true,
		},
		"slo_bindings": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("post", "redirect"),
			},
			MarkdownDescription: "SAML HTTP binding for SLO requests to the identity provider",
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
			MarkdownDescription: "The snippet in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"sso_bindings": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("post", "redirect"),
			},
			MarkdownDescription: "SAML HTTP binding for SSO requests to the identity provider",
			Required:            true,
		},
		"sso_url": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(255),
				stringvalidator.LengthAtLeast(1),
			},
			MarkdownDescription: "Identity provider SSO URL",
			Required:            true,
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"validate_idp_certificate": schema.BoolAttribute{
			MarkdownDescription: "Validate the identity provider certificate?",
			Optional:            true,
		},
		"want_auth_requests_signed": schema.BoolAttribute{
			MarkdownDescription: "Sign SAML message to the identity provider?",
			Optional:            true,
		},
	},
}

// SamlServerProfilesDataSourceSchema defines the schema for SamlServerProfiles data source
var SamlServerProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "SamlServerProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"certificate": dsschema.StringAttribute{
			MarkdownDescription: "The identity provider certificate",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"entity_id": dsschema.StringAttribute{
			MarkdownDescription: "The identity provider ID",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the SAML server profile",
			Required:            true,
		},
		"max_clock_skew": dsschema.Int64Attribute{
			MarkdownDescription: "Maxiumum clock skew",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the SAML server profile",
			Optional:            true,
			Computed:            true,
		},
		"slo_bindings": dsschema.StringAttribute{
			MarkdownDescription: "SAML HTTP binding for SLO requests to the identity provider",
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"sso_bindings": dsschema.StringAttribute{
			MarkdownDescription: "SAML HTTP binding for SSO requests to the identity provider",
			Computed:            true,
		},
		"sso_url": dsschema.StringAttribute{
			MarkdownDescription: "Identity provider SSO URL",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"validate_idp_certificate": dsschema.BoolAttribute{
			MarkdownDescription: "Validate the identity provider certificate?",
			Computed:            true,
		},
		"want_auth_requests_signed": dsschema.BoolAttribute{
			MarkdownDescription: "Sign SAML message to the identity provider?",
			Computed:            true,
		},
	},
}

// SamlServerProfilesListModel represents the data model for a list data source.
type SamlServerProfilesListModel struct {
	Tfid    types.String         `tfsdk:"tfid"`
	Data    []SamlServerProfiles `tfsdk:"data"`
	Limit   types.Int64          `tfsdk:"limit"`
	Offset  types.Int64          `tfsdk:"offset"`
	Name    types.String         `tfsdk:"name"`
	Total   types.Int64          `tfsdk:"total"`
	Folder  types.String         `tfsdk:"folder"`
	Snippet types.String         `tfsdk:"snippet"`
	Device  types.String         `tfsdk:"device"`
}

// SamlServerProfilesListDataSourceSchema defines the schema for a list data source.
var SamlServerProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: SamlServerProfilesDataSourceSchema.Attributes,
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
