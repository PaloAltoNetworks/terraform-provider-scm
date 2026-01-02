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

// Package: device_settings
// This file contains models for the device_settings SDK package

// AuthenticationSettings represents the Terraform model for AuthenticationSettings
type AuthenticationSettings struct {
	Tfid           types.String          `tfsdk:"tfid"`
	Authentication basetypes.ObjectValue `tfsdk:"authentication"`
	Device         basetypes.StringValue `tfsdk:"device"`
	Folder         basetypes.StringValue `tfsdk:"folder"`
	Id             basetypes.StringValue `tfsdk:"id"`
	Snippet        basetypes.StringValue `tfsdk:"snippet"`
}

// AuthenticationSettingsAuthentication represents a nested structure within the AuthenticationSettings model
type AuthenticationSettingsAuthentication struct {
	AccountingServerProfile basetypes.StringValue `tfsdk:"accounting_server_profile"`
	AuthenticationProfile   basetypes.StringValue `tfsdk:"authentication_profile"`
	CertificateProfile      basetypes.StringValue `tfsdk:"certificate_profile"`
}

// AttrTypes defines the attribute types for the AuthenticationSettings model.
func (o AuthenticationSettings) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"authentication": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"accounting_server_profile": basetypes.StringType{},
				"authentication_profile":    basetypes.StringType{},
				"certificate_profile":       basetypes.StringType{},
			},
		},
		"device":  basetypes.StringType{},
		"folder":  basetypes.StringType{},
		"id":      basetypes.StringType{},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AuthenticationSettings objects.
func (o AuthenticationSettings) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AuthenticationSettingsAuthentication model.
func (o AuthenticationSettingsAuthentication) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"accounting_server_profile": basetypes.StringType{},
		"authentication_profile":    basetypes.StringType{},
		"certificate_profile":       basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AuthenticationSettingsAuthentication objects.
func (o AuthenticationSettingsAuthentication) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AuthenticationSettingsResourceSchema defines the schema for AuthenticationSettings resource
var AuthenticationSettingsResourceSchema = schema.Schema{
	MarkdownDescription: "AuthenticationSetting resource",
	Attributes: map[string]schema.Attribute{
		"authentication": schema.SingleNestedAttribute{
			MarkdownDescription: "Authentication",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"accounting_server_profile": schema.StringAttribute{
					MarkdownDescription: "Accounting server profile",
					Optional:            true,
				},
				"authentication_profile": schema.StringAttribute{
					MarkdownDescription: "Authentication profile",
					Optional:            true,
				},
				"certificate_profile": schema.StringAttribute{
					MarkdownDescription: "Certificate profile",
					Optional:            true,
				},
			},
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
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
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
		"snippet": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("device"),
					path.MatchRelative().AtParent().AtName("folder"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
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

// AuthenticationSettingsDataSourceSchema defines the schema for AuthenticationSettings data source
var AuthenticationSettingsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "AuthenticationSetting data source",
	Attributes: map[string]dsschema.Attribute{
		"authentication": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Authentication",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"accounting_server_profile": dsschema.StringAttribute{
					MarkdownDescription: "Accounting server profile",
					Computed:            true,
				},
				"authentication_profile": dsschema.StringAttribute{
					MarkdownDescription: "Authentication profile",
					Computed:            true,
				},
				"certificate_profile": dsschema.StringAttribute{
					MarkdownDescription: "Certificate profile",
					Computed:            true,
				},
			},
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

// AuthenticationSettingsListModel represents the data model for a list data source.
type AuthenticationSettingsListModel struct {
	Tfid    types.String             `tfsdk:"tfid"`
	Data    []AuthenticationSettings `tfsdk:"data"`
	Limit   types.Int64              `tfsdk:"limit"`
	Offset  types.Int64              `tfsdk:"offset"`
	Name    types.String             `tfsdk:"name"`
	Total   types.Int64              `tfsdk:"total"`
	Folder  types.String             `tfsdk:"folder"`
	Snippet types.String             `tfsdk:"snippet"`
	Device  types.String             `tfsdk:"device"`
}

// AuthenticationSettingsListDataSourceSchema defines the schema for a list data source.
var AuthenticationSettingsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: AuthenticationSettingsDataSourceSchema.Attributes,
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
