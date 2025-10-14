package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: identity_services
// This file contains models for the identity_services SDK package

// MfaServers represents the Terraform model for MfaServers
type MfaServers struct {
	Tfid            types.String          `tfsdk:"tfid"`
	EncryptedValues basetypes.MapValue    `tfsdk:"encrypted_values"`
	Device          basetypes.StringValue `tfsdk:"device"`
	Folder          basetypes.StringValue `tfsdk:"folder"`
	Id              basetypes.StringValue `tfsdk:"id"`
	MfaCertProfile  basetypes.StringValue `tfsdk:"mfa_cert_profile"`
	MfaVendorType   basetypes.ObjectValue `tfsdk:"mfa_vendor_type"`
	Name            basetypes.StringValue `tfsdk:"name"`
	Snippet         basetypes.StringValue `tfsdk:"snippet"`
}

// MfaServersMfaVendorType represents a nested structure within the MfaServers model
type MfaServersMfaVendorType struct {
	DuoSecurityV2      basetypes.ObjectValue `tfsdk:"duo_security_v2"`
	OktaAdaptiveV1     basetypes.ObjectValue `tfsdk:"okta_adaptive_v1"`
	PingIdentityV1     basetypes.ObjectValue `tfsdk:"ping_identity_v1"`
	RsaSecuridAccessV1 basetypes.ObjectValue `tfsdk:"rsa_securid_access_v1"`
}

// MfaServersMfaVendorTypeDuoSecurityV2 represents a nested structure within the MfaServers model
type MfaServersMfaVendorTypeDuoSecurityV2 struct {
	DuoApiHost        basetypes.StringValue `tfsdk:"duo_api_host"`
	DuoBaseuri        basetypes.StringValue `tfsdk:"duo_baseuri"`
	DuoIntegrationKey basetypes.StringValue `tfsdk:"duo_integration_key"`
	DuoSecretKey      basetypes.StringValue `tfsdk:"duo_secret_key"`
	DuoTimeout        basetypes.Int64Value  `tfsdk:"duo_timeout"`
}

// MfaServersMfaVendorTypeOktaAdaptiveV1 represents a nested structure within the MfaServers model
type MfaServersMfaVendorTypeOktaAdaptiveV1 struct {
	OktaApiHost basetypes.StringValue `tfsdk:"okta_api_host"`
	OktaBaseuri basetypes.StringValue `tfsdk:"okta_baseuri"`
	OktaOrg     basetypes.StringValue `tfsdk:"okta_org"`
	OktaTimeout basetypes.Int64Value  `tfsdk:"okta_timeout"`
	OktaToken   basetypes.StringValue `tfsdk:"okta_token"`
}

// MfaServersMfaVendorTypePingIdentityV1 represents a nested structure within the MfaServers model
type MfaServersMfaVendorTypePingIdentityV1 struct {
	PingApiHost      basetypes.StringValue `tfsdk:"ping_api_host"`
	PingBaseuri      basetypes.StringValue `tfsdk:"ping_baseuri"`
	PingOrgAlias     basetypes.StringValue `tfsdk:"ping_org_alias"`
	PingTimeout      basetypes.Int64Value  `tfsdk:"ping_timeout"`
	PingToken        basetypes.StringValue `tfsdk:"ping_token"`
	PingUseBase64Key basetypes.StringValue `tfsdk:"ping_use_base64_key"`
}

// MfaServersMfaVendorTypeRsaSecuridAccessV1 represents a nested structure within the MfaServers model
type MfaServersMfaVendorTypeRsaSecuridAccessV1 struct {
	RsaAccessid          basetypes.StringValue `tfsdk:"rsa_accessid"`
	RsaAccesskey         basetypes.StringValue `tfsdk:"rsa_accesskey"`
	RsaApiHost           basetypes.StringValue `tfsdk:"rsa_api_host"`
	RsaAssurancepolicyid basetypes.StringValue `tfsdk:"rsa_assurancepolicyid"`
	RsaBaseuri           basetypes.StringValue `tfsdk:"rsa_baseuri"`
	RsaTimeout           basetypes.Int64Value  `tfsdk:"rsa_timeout"`
}

// AttrTypes defines the attribute types for the MfaServers model.
func (o MfaServers) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":             basetypes.StringType{},
		"encrypted_values": basetypes.MapType{ElemType: basetypes.StringType{}},
		"device":           basetypes.StringType{},
		"folder":           basetypes.StringType{},
		"id":               basetypes.StringType{},
		"mfa_cert_profile": basetypes.StringType{},
		"mfa_vendor_type": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"duo_security_v2": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"duo_api_host":        basetypes.StringType{},
						"duo_baseuri":         basetypes.StringType{},
						"duo_integration_key": basetypes.StringType{},
						"duo_secret_key":      basetypes.StringType{},
						"duo_timeout":         basetypes.Int64Type{},
					},
				},
				"okta_adaptive_v1": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"okta_api_host": basetypes.StringType{},
						"okta_baseuri":  basetypes.StringType{},
						"okta_org":      basetypes.StringType{},
						"okta_timeout":  basetypes.Int64Type{},
						"okta_token":    basetypes.StringType{},
					},
				},
				"ping_identity_v1": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ping_api_host":       basetypes.StringType{},
						"ping_baseuri":        basetypes.StringType{},
						"ping_org_alias":      basetypes.StringType{},
						"ping_timeout":        basetypes.Int64Type{},
						"ping_token":          basetypes.StringType{},
						"ping_use_base64_key": basetypes.StringType{},
					},
				},
				"rsa_securid_access_v1": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"rsa_accessid":          basetypes.StringType{},
						"rsa_accesskey":         basetypes.StringType{},
						"rsa_api_host":          basetypes.StringType{},
						"rsa_assurancepolicyid": basetypes.StringType{},
						"rsa_baseuri":           basetypes.StringType{},
						"rsa_timeout":           basetypes.Int64Type{},
					},
				},
			},
		},
		"name":    basetypes.StringType{},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of MfaServers objects.
func (o MfaServers) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the MfaServersMfaVendorType model.
func (o MfaServersMfaVendorType) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"duo_security_v2": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"duo_api_host":        basetypes.StringType{},
				"duo_baseuri":         basetypes.StringType{},
				"duo_integration_key": basetypes.StringType{},
				"duo_secret_key":      basetypes.StringType{},
				"duo_timeout":         basetypes.Int64Type{},
			},
		},
		"okta_adaptive_v1": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"okta_api_host": basetypes.StringType{},
				"okta_baseuri":  basetypes.StringType{},
				"okta_org":      basetypes.StringType{},
				"okta_timeout":  basetypes.Int64Type{},
				"okta_token":    basetypes.StringType{},
			},
		},
		"ping_identity_v1": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ping_api_host":       basetypes.StringType{},
				"ping_baseuri":        basetypes.StringType{},
				"ping_org_alias":      basetypes.StringType{},
				"ping_timeout":        basetypes.Int64Type{},
				"ping_token":          basetypes.StringType{},
				"ping_use_base64_key": basetypes.StringType{},
			},
		},
		"rsa_securid_access_v1": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"rsa_accessid":          basetypes.StringType{},
				"rsa_accesskey":         basetypes.StringType{},
				"rsa_api_host":          basetypes.StringType{},
				"rsa_assurancepolicyid": basetypes.StringType{},
				"rsa_baseuri":           basetypes.StringType{},
				"rsa_timeout":           basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of MfaServersMfaVendorType objects.
func (o MfaServersMfaVendorType) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the MfaServersMfaVendorTypeDuoSecurityV2 model.
func (o MfaServersMfaVendorTypeDuoSecurityV2) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"duo_api_host":        basetypes.StringType{},
		"duo_baseuri":         basetypes.StringType{},
		"duo_integration_key": basetypes.StringType{},
		"duo_secret_key":      basetypes.StringType{},
		"duo_timeout":         basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of MfaServersMfaVendorTypeDuoSecurityV2 objects.
func (o MfaServersMfaVendorTypeDuoSecurityV2) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the MfaServersMfaVendorTypeOktaAdaptiveV1 model.
func (o MfaServersMfaVendorTypeOktaAdaptiveV1) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"okta_api_host": basetypes.StringType{},
		"okta_baseuri":  basetypes.StringType{},
		"okta_org":      basetypes.StringType{},
		"okta_timeout":  basetypes.Int64Type{},
		"okta_token":    basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of MfaServersMfaVendorTypeOktaAdaptiveV1 objects.
func (o MfaServersMfaVendorTypeOktaAdaptiveV1) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the MfaServersMfaVendorTypePingIdentityV1 model.
func (o MfaServersMfaVendorTypePingIdentityV1) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ping_api_host":       basetypes.StringType{},
		"ping_baseuri":        basetypes.StringType{},
		"ping_org_alias":      basetypes.StringType{},
		"ping_timeout":        basetypes.Int64Type{},
		"ping_token":          basetypes.StringType{},
		"ping_use_base64_key": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of MfaServersMfaVendorTypePingIdentityV1 objects.
func (o MfaServersMfaVendorTypePingIdentityV1) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the MfaServersMfaVendorTypeRsaSecuridAccessV1 model.
func (o MfaServersMfaVendorTypeRsaSecuridAccessV1) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"rsa_accessid":          basetypes.StringType{},
		"rsa_accesskey":         basetypes.StringType{},
		"rsa_api_host":          basetypes.StringType{},
		"rsa_assurancepolicyid": basetypes.StringType{},
		"rsa_baseuri":           basetypes.StringType{},
		"rsa_timeout":           basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of MfaServersMfaVendorTypeRsaSecuridAccessV1 objects.
func (o MfaServersMfaVendorTypeRsaSecuridAccessV1) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// MfaServersResourceSchema defines the schema for MfaServers resource
var MfaServersResourceSchema = schema.Schema{
	MarkdownDescription: "MfaServer resource",
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
			MarkdownDescription: "The device in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"encrypted_values": schema.MapAttribute{
			ElementType:         basetypes.StringType{},
			MarkdownDescription: "Map of sensitive values returned from the API.",
			Computed:            true,
			Sensitive:           true,
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
			MarkdownDescription: "The UUID of the MFA server",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"mfa_cert_profile": schema.StringAttribute{
			MarkdownDescription: "The MFA server certificate profile",
			Required:            true,
		},
		"mfa_vendor_type": schema.SingleNestedAttribute{
			MarkdownDescription: "The MFA vendor type",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"duo_security_v2": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("okta_adaptive_v1"),
							path.MatchRelative().AtParent().AtName("ping_identity_v1"),
							path.MatchRelative().AtParent().AtName("rsa_securid_access_v1"),
						),
					},
					MarkdownDescription: "Integration with [Duo Security](https://duo.com/product)\n",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"duo_api_host": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(16),
							},
							MarkdownDescription: "Duo Security API hostname",
							Required:            true,
						},
						"duo_baseuri": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(2),
							},
							MarkdownDescription: "Duo Security API base URI",
							Required:            true,
						},
						"duo_integration_key": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(16),
							},
							MarkdownDescription: "Duo Security integration key",
							Required:            true,
						},
						"duo_secret_key": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(16),
							},
							MarkdownDescription: "Duo Security secret key",
							Required:            true,
							Sensitive:           true,
						},
						"duo_timeout": schema.Int64Attribute{
							Validators: []validator.Int64{
								int64validator.Between(5, 600),
							},
							MarkdownDescription: "Duo Security timeout (seconds)",
							Required:            true,
						},
					},
				},
				"okta_adaptive_v1": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("duo_security_v2"),
							path.MatchRelative().AtParent().AtName("ping_identity_v1"),
							path.MatchRelative().AtParent().AtName("rsa_securid_access_v1"),
						),
					},
					MarkdownDescription: "Integration with [Okta Adaptive MFA](https://www.okta.com/products/adaptive-multi-factor-authentication)",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"okta_api_host": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(10),
							},
							MarkdownDescription: "Okta API hostname",
							Required:            true,
						},
						"okta_baseuri": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(2),
							},
							MarkdownDescription: "Okta baseuri",
							Required:            true,
						},
						"okta_org": schema.StringAttribute{
							MarkdownDescription: "Okta organization",
							Required:            true,
						},
						"okta_timeout": schema.Int64Attribute{
							Validators: []validator.Int64{
								int64validator.Between(5, 600),
							},
							MarkdownDescription: "Okta timeout (seconds)",
							Required:            true,
						},
						"okta_token": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(8),
							},
							MarkdownDescription: "Okta API token",
							Required:            true,
							Sensitive:           true,
						},
					},
				},
				"ping_identity_v1": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("duo_security_v2"),
							path.MatchRelative().AtParent().AtName("okta_adaptive_v1"),
							path.MatchRelative().AtParent().AtName("rsa_securid_access_v1"),
						),
					},
					MarkdownDescription: "Integation with [Ping Identity](https://www.pingidentity.com/en/platform.html)",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"ping_api_host": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(16),
							},
							MarkdownDescription: "Ping Identity API hostname",
							Required:            true,
						},
						"ping_baseuri": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(2),
							},
							MarkdownDescription: "Ping Identity API base URI",
							Required:            true,
						},
						"ping_org_alias": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(8),
							},
							MarkdownDescription: "Ping Identity client organization ID",
							Optional:            true,
							Computed:            true,
						},
						"ping_timeout": schema.Int64Attribute{
							Validators: []validator.Int64{
								int64validator.Between(5, 600),
							},
							MarkdownDescription: "Ping Identity timeout (seconds)",
							Required:            true,
						},
						"ping_token": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(8),
							},
							MarkdownDescription: "Ping Identity API token",
							Required:            true,
						},
						"ping_use_base64_key": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(8),
							},
							MarkdownDescription: "Ping Identity Base64 key",
							Required:            true,
							Sensitive:           true,
						},
					},
				},
				"rsa_securid_access_v1": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("duo_security_v2"),
							path.MatchRelative().AtParent().AtName("okta_adaptive_v1"),
							path.MatchRelative().AtParent().AtName("ping_identity_v1"),
						),
					},
					MarkdownDescription: "Integration with [RSA SecurID](https://www.rsa.com/products/securid/)",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"rsa_accessid": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(8),
							},
							MarkdownDescription: "RSA SecurID access ID",
							Optional:            true,
							Computed:            true,
						},
						"rsa_accesskey": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(8),
							},
							MarkdownDescription: "RSA SecurID access key",
							Optional:            true,
							Computed:            true,
							Sensitive:           true,
						},
						"rsa_api_host": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(10),
							},
							MarkdownDescription: "RSA SecurID hostname",
							Optional:            true,
							Computed:            true,
						},
						"rsa_assurancepolicyid": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(3),
							},
							MarkdownDescription: "RSA SecurID assurance level",
							Optional:            true,
							Computed:            true,
						},
						"rsa_baseuri": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(2),
							},
							MarkdownDescription: "RSA SecurID API base URI",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("/mfa/v1_1"),
						},
						"rsa_timeout": schema.Int64Attribute{
							Validators: []validator.Int64{
								int64validator.Between(5, 600),
							},
							MarkdownDescription: "RSA SecurID timeout (seconds)",
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(30),
						},
					},
				},
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the MFA server profile",
			Required:            true,
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
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
	},
}

// MfaServersDataSourceSchema defines the schema for MfaServers data source
var MfaServersDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "MfaServer data source",
	Attributes: map[string]dsschema.Attribute{
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"encrypted_values": dsschema.MapAttribute{
			ElementType:         basetypes.StringType{},
			MarkdownDescription: "Map of sensitive values returned from the API.",
			Computed:            true,
			Sensitive:           true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the MFA server",
			Required:            true,
		},
		"mfa_cert_profile": dsschema.StringAttribute{
			MarkdownDescription: "The MFA server certificate profile",
			Computed:            true,
		},
		"mfa_vendor_type": dsschema.SingleNestedAttribute{
			MarkdownDescription: "The MFA vendor type",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"duo_security_v2": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Integration with [Duo Security](https://duo.com/product)\n",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"duo_api_host": dsschema.StringAttribute{
							MarkdownDescription: "Duo Security API hostname",
							Computed:            true,
						},
						"duo_baseuri": dsschema.StringAttribute{
							MarkdownDescription: "Duo Security API base URI",
							Computed:            true,
						},
						"duo_integration_key": dsschema.StringAttribute{
							MarkdownDescription: "Duo Security integration key",
							Computed:            true,
						},
						"duo_secret_key": dsschema.StringAttribute{
							MarkdownDescription: "Duo Security secret key",
							Computed:            true,
							Sensitive:           true,
						},
						"duo_timeout": dsschema.Int64Attribute{
							MarkdownDescription: "Duo Security timeout (seconds)",
							Computed:            true,
						},
					},
				},
				"okta_adaptive_v1": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Integration with [Okta Adaptive MFA](https://www.okta.com/products/adaptive-multi-factor-authentication)",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"okta_api_host": dsschema.StringAttribute{
							MarkdownDescription: "Okta API hostname",
							Computed:            true,
						},
						"okta_baseuri": dsschema.StringAttribute{
							MarkdownDescription: "Okta baseuri",
							Computed:            true,
						},
						"okta_org": dsschema.StringAttribute{
							MarkdownDescription: "Okta organization",
							Computed:            true,
						},
						"okta_timeout": dsschema.Int64Attribute{
							MarkdownDescription: "Okta timeout (seconds)",
							Computed:            true,
						},
						"okta_token": dsschema.StringAttribute{
							MarkdownDescription: "Okta API token",
							Computed:            true,
							Sensitive:           true,
						},
					},
				},
				"ping_identity_v1": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Integation with [Ping Identity](https://www.pingidentity.com/en/platform.html)",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"ping_api_host": dsschema.StringAttribute{
							MarkdownDescription: "Ping Identity API hostname",
							Computed:            true,
						},
						"ping_baseuri": dsschema.StringAttribute{
							MarkdownDescription: "Ping Identity API base URI",
							Computed:            true,
						},
						"ping_org_alias": dsschema.StringAttribute{
							MarkdownDescription: "Ping Identity client organization ID",
							Computed:            true,
						},
						"ping_timeout": dsschema.Int64Attribute{
							MarkdownDescription: "Ping Identity timeout (seconds)",
							Computed:            true,
						},
						"ping_token": dsschema.StringAttribute{
							MarkdownDescription: "Ping Identity API token",
							Computed:            true,
						},
						"ping_use_base64_key": dsschema.StringAttribute{
							MarkdownDescription: "Ping Identity Base64 key",
							Computed:            true,
							Sensitive:           true,
						},
					},
				},
				"rsa_securid_access_v1": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Integration with [RSA SecurID](https://www.rsa.com/products/securid/)",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"rsa_accessid": dsschema.StringAttribute{
							MarkdownDescription: "RSA SecurID access ID",
							Computed:            true,
						},
						"rsa_accesskey": dsschema.StringAttribute{
							MarkdownDescription: "RSA SecurID access key",
							Computed:            true,
							Sensitive:           true,
						},
						"rsa_api_host": dsschema.StringAttribute{
							MarkdownDescription: "RSA SecurID hostname",
							Computed:            true,
						},
						"rsa_assurancepolicyid": dsschema.StringAttribute{
							MarkdownDescription: "RSA SecurID assurance level",
							Computed:            true,
						},
						"rsa_baseuri": dsschema.StringAttribute{
							MarkdownDescription: "RSA SecurID API base URI",
							Computed:            true,
						},
						"rsa_timeout": dsschema.Int64Attribute{
							MarkdownDescription: "RSA SecurID timeout (seconds)",
							Computed:            true,
						},
					},
				},
			},
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the MFA server profile",
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

// MfaServersListModel represents the data model for a list data source.
type MfaServersListModel struct {
	Tfid    types.String `tfsdk:"tfid"`
	Data    []MfaServers `tfsdk:"data"`
	Limit   types.Int64  `tfsdk:"limit"`
	Offset  types.Int64  `tfsdk:"offset"`
	Name    types.String `tfsdk:"name"`
	Total   types.Int64  `tfsdk:"total"`
	Folder  types.String `tfsdk:"folder"`
	Snippet types.String `tfsdk:"snippet"`
	Device  types.String `tfsdk:"device"`
}

// MfaServersListDataSourceSchema defines the schema for a list data source.
var MfaServersListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: MfaServersDataSourceSchema.Attributes,
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
