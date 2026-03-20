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

// SslDecryptionSettings represents the Terraform model for SslDecryptionSettings
type SslDecryptionSettings struct {
	Tfid                                 types.String          `tfsdk:"tfid"`
	Device                               basetypes.StringValue `tfsdk:"device"`
	DisabledSslExcludeCertFromPredefined basetypes.ListValue   `tfsdk:"disabled_ssl_exclude_cert_from_predefined"`
	Folder                               basetypes.StringValue `tfsdk:"folder"`
	ForwardTrustCertificate              basetypes.ObjectValue `tfsdk:"forward_trust_certificate"`
	ForwardUntrustCertificate            basetypes.ObjectValue `tfsdk:"forward_untrust_certificate"`
	RootCaExcludeList                    basetypes.ListValue   `tfsdk:"root_ca_exclude_list"`
	Snippet                              basetypes.StringValue `tfsdk:"snippet"`
	SslExcludeCert                       basetypes.ListValue   `tfsdk:"ssl_exclude_cert"`
	TrustedRootCA                        basetypes.ListValue   `tfsdk:"trusted_root_ca"`
}

// SslDecryptionSettingsForwardTrustCertificate represents a nested structure within the SslDecryptionSettings model
type SslDecryptionSettingsForwardTrustCertificate struct {
	Ecdsa basetypes.StringValue `tfsdk:"ecdsa"`
	Rsa   basetypes.StringValue `tfsdk:"rsa"`
}

// SslDecryptionSettingsSslExcludeCertInner represents a nested structure within the SslDecryptionSettings model
type SslDecryptionSettingsSslExcludeCertInner struct {
	Description basetypes.StringValue `tfsdk:"description"`
	Exclude     basetypes.BoolValue   `tfsdk:"exclude"`
	Name        basetypes.StringValue `tfsdk:"name"`
}

// AttrTypes defines the attribute types for the SslDecryptionSettings model.
func (o SslDecryptionSettings) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":   basetypes.StringType{},
		"device": basetypes.StringType{},
		"disabled_ssl_exclude_cert_from_predefined": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		}},
		"folder": basetypes.StringType{},
		"forward_trust_certificate": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ecdsa": basetypes.StringType{},
				"rsa":   basetypes.StringType{},
			},
		},
		"forward_untrust_certificate": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ecdsa": basetypes.StringType{},
				"rsa":   basetypes.StringType{},
			},
		},
		"root_ca_exclude_list": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		}},
		"snippet": basetypes.StringType{},
		"ssl_exclude_cert": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"description": basetypes.StringType{},
				"exclude":     basetypes.BoolType{},
				"name":        basetypes.StringType{},
			},
		}},
		"trusted_root_ca": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		}},
	}
}

// AttrType returns the attribute type for a list of SslDecryptionSettings objects.
func (o SslDecryptionSettings) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SslDecryptionSettingsForwardTrustCertificate model.
func (o SslDecryptionSettingsForwardTrustCertificate) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ecdsa": basetypes.StringType{},
		"rsa":   basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of SslDecryptionSettingsForwardTrustCertificate objects.
func (o SslDecryptionSettingsForwardTrustCertificate) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SslDecryptionSettingsSslExcludeCertInner model.
func (o SslDecryptionSettingsSslExcludeCertInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"description": basetypes.StringType{},
		"exclude":     basetypes.BoolType{},
		"name":        basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of SslDecryptionSettingsSslExcludeCertInner objects.
func (o SslDecryptionSettingsSslExcludeCertInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// SslDecryptionSettingsResourceSchema defines the schema for SslDecryptionSettings resource
var SslDecryptionSettingsResourceSchema = schema.Schema{
	MarkdownDescription: "SslDecryptionSetting resource",
	Attributes: map[string]schema.Attribute{
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
		"disabled_ssl_exclude_cert_from_predefined": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Disabled ssl exclude cert from predefined",
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
		"forward_trust_certificate": schema.SingleNestedAttribute{
			MarkdownDescription: "Forward trust certificate",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"ecdsa": schema.StringAttribute{
					MarkdownDescription: "Ecdsa",
					Optional:            true,
				},
				"rsa": schema.StringAttribute{
					MarkdownDescription: "Rsa",
					Optional:            true,
				},
			},
		},
		"forward_untrust_certificate": schema.SingleNestedAttribute{
			MarkdownDescription: "Forward untrust certificate",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"ecdsa": schema.StringAttribute{
					MarkdownDescription: "Ecdsa",
					Optional:            true,
				},
				"rsa": schema.StringAttribute{
					MarkdownDescription: "Rsa",
					Optional:            true,
				},
			},
		},
		"root_ca_exclude_list": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Root ca exclude list",
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
			MarkdownDescription: "The snippet in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"ssl_exclude_cert": schema.ListNestedAttribute{
			MarkdownDescription: "Ssl exclude cert",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						MarkdownDescription: "Description",
						Optional:            true,
					},
					"exclude": schema.BoolAttribute{
						MarkdownDescription: "Exclude",
						Optional:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "Name",
						Optional:            true,
					},
				},
			},
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"trusted_root_ca": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Trusted root c a",
			Optional:            true,
		},
	},
}

// SslDecryptionSettingsDataSourceSchema defines the schema for SslDecryptionSettings data source
var SslDecryptionSettingsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "SslDecryptionSetting data source",
	Attributes: map[string]dsschema.Attribute{
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"disabled_ssl_exclude_cert_from_predefined": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Disabled ssl exclude cert from predefined",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"forward_trust_certificate": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Forward trust certificate",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"ecdsa": dsschema.StringAttribute{
					MarkdownDescription: "Ecdsa",
					Computed:            true,
				},
				"rsa": dsschema.StringAttribute{
					MarkdownDescription: "Rsa",
					Computed:            true,
				},
			},
		},
		"forward_untrust_certificate": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Forward untrust certificate",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"ecdsa": dsschema.StringAttribute{
					MarkdownDescription: "Ecdsa",
					Computed:            true,
				},
				"rsa": dsschema.StringAttribute{
					MarkdownDescription: "Rsa",
					Computed:            true,
				},
			},
		},
		"root_ca_exclude_list": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Root ca exclude list",
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"ssl_exclude_cert": dsschema.ListNestedAttribute{
			MarkdownDescription: "Ssl exclude cert",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"description": dsschema.StringAttribute{
						MarkdownDescription: "Description",
						Computed:            true,
					},
					"exclude": dsschema.BoolAttribute{
						MarkdownDescription: "Exclude",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "Name",
						Computed:            true,
					},
				},
			},
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"trusted_root_ca": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Trusted root c a",
			Computed:            true,
		},
	},
}

// SslDecryptionSettingsListModel represents the data model for a list data source.
type SslDecryptionSettingsListModel struct {
	Tfid    types.String            `tfsdk:"tfid"`
	Data    []SslDecryptionSettings `tfsdk:"data"`
	Limit   types.Int64             `tfsdk:"limit"`
	Offset  types.Int64             `tfsdk:"offset"`
	Name    types.String            `tfsdk:"name"`
	Total   types.Int64             `tfsdk:"total"`
	Folder  types.String            `tfsdk:"folder"`
	Snippet types.String            `tfsdk:"snippet"`
	Device  types.String            `tfsdk:"device"`
}

// SslDecryptionSettingsListDataSourceSchema defines the schema for a list data source.
var SslDecryptionSettingsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: SslDecryptionSettingsDataSourceSchema.Attributes,
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
