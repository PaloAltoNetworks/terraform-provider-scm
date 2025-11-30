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

// Package: identity_services
// This file contains models for the identity_services SDK package

// CertificateProfiles represents the Terraform model for CertificateProfiles
type CertificateProfiles struct {
	Tfid                     types.String          `tfsdk:"tfid"`
	BlockExpiredCert         basetypes.BoolValue   `tfsdk:"block_expired_cert"`
	BlockTimeoutCert         basetypes.BoolValue   `tfsdk:"block_timeout_cert"`
	BlockUnauthenticatedCert basetypes.BoolValue   `tfsdk:"block_unauthenticated_cert"`
	BlockUnknownCert         basetypes.BoolValue   `tfsdk:"block_unknown_cert"`
	CaCertificates           basetypes.ListValue   `tfsdk:"ca_certificates"`
	CertStatusTimeout        basetypes.StringValue `tfsdk:"cert_status_timeout"`
	CrlReceiveTimeout        basetypes.StringValue `tfsdk:"crl_receive_timeout"`
	Device                   basetypes.StringValue `tfsdk:"device"`
	Domain                   basetypes.StringValue `tfsdk:"domain"`
	Folder                   basetypes.StringValue `tfsdk:"folder"`
	Id                       basetypes.StringValue `tfsdk:"id"`
	Name                     basetypes.StringValue `tfsdk:"name"`
	OcspReceiveTimeout       basetypes.StringValue `tfsdk:"ocsp_receive_timeout"`
	Snippet                  basetypes.StringValue `tfsdk:"snippet"`
	UseCrl                   basetypes.BoolValue   `tfsdk:"use_crl"`
	UseOcsp                  basetypes.BoolValue   `tfsdk:"use_ocsp"`
	UsernameField            basetypes.ObjectValue `tfsdk:"username_field"`
}

// CertificateProfilesCaCertificatesInner represents a nested structure within the CertificateProfiles model
type CertificateProfilesCaCertificatesInner struct {
	DefaultOcspUrl basetypes.StringValue `tfsdk:"default_ocsp_url"`
	Name           basetypes.StringValue `tfsdk:"name"`
	OcspVerifyCert basetypes.StringValue `tfsdk:"ocsp_verify_cert"`
	TemplateName   basetypes.StringValue `tfsdk:"template_name"`
}

// CertificateProfilesUsernameField represents a nested structure within the CertificateProfiles model
type CertificateProfilesUsernameField struct {
	Subject    basetypes.StringValue `tfsdk:"subject"`
	SubjectAlt basetypes.StringValue `tfsdk:"subject_alt"`
}

// AttrTypes defines the attribute types for the CertificateProfiles model.
func (o CertificateProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":                       basetypes.StringType{},
		"block_expired_cert":         basetypes.BoolType{},
		"block_timeout_cert":         basetypes.BoolType{},
		"block_unauthenticated_cert": basetypes.BoolType{},
		"block_unknown_cert":         basetypes.BoolType{},
		"ca_certificates": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"default_ocsp_url": basetypes.StringType{},
				"name":             basetypes.StringType{},
				"ocsp_verify_cert": basetypes.StringType{},
				"template_name":    basetypes.StringType{},
			},
		}},
		"cert_status_timeout":  basetypes.StringType{},
		"crl_receive_timeout":  basetypes.StringType{},
		"device":               basetypes.StringType{},
		"domain":               basetypes.StringType{},
		"folder":               basetypes.StringType{},
		"id":                   basetypes.StringType{},
		"name":                 basetypes.StringType{},
		"ocsp_receive_timeout": basetypes.StringType{},
		"snippet":              basetypes.StringType{},
		"use_crl":              basetypes.BoolType{},
		"use_ocsp":             basetypes.BoolType{},
		"username_field": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"subject":     basetypes.StringType{},
				"subject_alt": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of CertificateProfiles objects.
func (o CertificateProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the CertificateProfilesCaCertificatesInner model.
func (o CertificateProfilesCaCertificatesInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"default_ocsp_url": basetypes.StringType{},
		"name":             basetypes.StringType{},
		"ocsp_verify_cert": basetypes.StringType{},
		"template_name":    basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of CertificateProfilesCaCertificatesInner objects.
func (o CertificateProfilesCaCertificatesInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the CertificateProfilesUsernameField model.
func (o CertificateProfilesUsernameField) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"subject":     basetypes.StringType{},
		"subject_alt": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of CertificateProfilesUsernameField objects.
func (o CertificateProfilesUsernameField) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// CertificateProfilesResourceSchema defines the schema for CertificateProfiles resource
var CertificateProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "CertificateProfile resource",
	Attributes: map[string]schema.Attribute{
		"block_expired_cert": schema.BoolAttribute{
			MarkdownDescription: "Block sessions with expired certificates?",
			Optional:            true,
		},
		"block_timeout_cert": schema.BoolAttribute{
			MarkdownDescription: "Block session if certificate status cannot be retrieved within timeout?",
			Optional:            true,
		},
		"block_unauthenticated_cert": schema.BoolAttribute{
			MarkdownDescription: "Block session if the certificate was not issued to the authenticating device?",
			Optional:            true,
		},
		"block_unknown_cert": schema.BoolAttribute{
			MarkdownDescription: "Block session if certificate status is unknown?",
			Optional:            true,
		},
		"ca_certificates": schema.ListNestedAttribute{
			MarkdownDescription: "An ordered list of CA certificates",
			Required:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"default_ocsp_url": schema.StringAttribute{
						MarkdownDescription: "Default OCSP URL",
						Optional:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "CA certificate name",
						Required:            true,
					},
					"ocsp_verify_cert": schema.StringAttribute{
						MarkdownDescription: "OCSP verify certificate",
						Optional:            true,
					},
					"template_name": schema.StringAttribute{
						MarkdownDescription: "Template name/OID",
						Optional:            true,
					},
				},
			},
		},
		"cert_status_timeout": schema.StringAttribute{
			MarkdownDescription: "Certificate status timeout",
			Optional:            true,
		},
		"crl_receive_timeout": schema.StringAttribute{
			MarkdownDescription: "CRL receive timeout (seconds)",
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
			MarkdownDescription: "The device in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"domain": schema.StringAttribute{
			MarkdownDescription: "User domain",
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
			},
			MarkdownDescription: "The folder in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the certificate profile",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(63),
			},
			MarkdownDescription: "The name of the certificate profile",
			Required:            true,
		},
		"ocsp_receive_timeout": schema.StringAttribute{
			MarkdownDescription: "OCSP receive timeout (seconds)",
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
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"use_crl": schema.BoolAttribute{
			MarkdownDescription: "Use CRL?",
			Optional:            true,
		},
		"use_ocsp": schema.BoolAttribute{
			MarkdownDescription: "Use OCSP?",
			Optional:            true,
		},
		"username_field": schema.SingleNestedAttribute{
			MarkdownDescription: "Certificate username field",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"subject": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("common-name"),
					},
					MarkdownDescription: "Common name",
					Optional:            true,
				},
				"subject_alt": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("email"),
					},
					MarkdownDescription: "Email address",
					Optional:            true,
				},
			},
		},
	},
}

// CertificateProfilesDataSourceSchema defines the schema for CertificateProfiles data source
var CertificateProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "CertificateProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"block_expired_cert": dsschema.BoolAttribute{
			MarkdownDescription: "Block sessions with expired certificates?",
			Computed:            true,
		},
		"block_timeout_cert": dsschema.BoolAttribute{
			MarkdownDescription: "Block session if certificate status cannot be retrieved within timeout?",
			Computed:            true,
		},
		"block_unauthenticated_cert": dsschema.BoolAttribute{
			MarkdownDescription: "Block session if the certificate was not issued to the authenticating device?",
			Computed:            true,
		},
		"block_unknown_cert": dsschema.BoolAttribute{
			MarkdownDescription: "Block session if certificate status is unknown?",
			Computed:            true,
		},
		"ca_certificates": dsschema.ListNestedAttribute{
			MarkdownDescription: "An ordered list of CA certificates",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"default_ocsp_url": dsschema.StringAttribute{
						MarkdownDescription: "Default OCSP URL",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "CA certificate name",
						Computed:            true,
					},
					"ocsp_verify_cert": dsschema.StringAttribute{
						MarkdownDescription: "OCSP verify certificate",
						Computed:            true,
					},
					"template_name": dsschema.StringAttribute{
						MarkdownDescription: "Template name/OID",
						Computed:            true,
					},
				},
			},
		},
		"cert_status_timeout": dsschema.StringAttribute{
			MarkdownDescription: "Certificate status timeout",
			Computed:            true,
		},
		"crl_receive_timeout": dsschema.StringAttribute{
			MarkdownDescription: "CRL receive timeout (seconds)",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"domain": dsschema.StringAttribute{
			MarkdownDescription: "User domain",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the certificate profile",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the certificate profile",
			Optional:            true,
			Computed:            true,
		},
		"ocsp_receive_timeout": dsschema.StringAttribute{
			MarkdownDescription: "OCSP receive timeout (seconds)",
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"use_crl": dsschema.BoolAttribute{
			MarkdownDescription: "Use CRL?",
			Computed:            true,
		},
		"use_ocsp": dsschema.BoolAttribute{
			MarkdownDescription: "Use OCSP?",
			Computed:            true,
		},
		"username_field": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Certificate username field",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"subject": dsschema.StringAttribute{
					MarkdownDescription: "Common name",
					Computed:            true,
				},
				"subject_alt": dsschema.StringAttribute{
					MarkdownDescription: "Email address",
					Computed:            true,
				},
			},
		},
	},
}

// CertificateProfilesListModel represents the data model for a list data source.
type CertificateProfilesListModel struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Data    []CertificateProfiles `tfsdk:"data"`
	Limit   types.Int64           `tfsdk:"limit"`
	Offset  types.Int64           `tfsdk:"offset"`
	Name    types.String          `tfsdk:"name"`
	Total   types.Int64           `tfsdk:"total"`
	Folder  types.String          `tfsdk:"folder"`
	Snippet types.String          `tfsdk:"snippet"`
	Device  types.String          `tfsdk:"device"`
}

// CertificateProfilesListDataSourceSchema defines the schema for a list data source.
var CertificateProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: CertificateProfilesDataSourceSchema.Attributes,
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
