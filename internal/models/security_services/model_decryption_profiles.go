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

// DecryptionProfiles represents the Terraform model for DecryptionProfiles
type DecryptionProfiles struct {
	Tfid                types.String          `tfsdk:"tfid"`
	Device              basetypes.StringValue `tfsdk:"device"`
	Folder              basetypes.StringValue `tfsdk:"folder"`
	Id                  basetypes.StringValue `tfsdk:"id"`
	Name                basetypes.StringValue `tfsdk:"name"`
	Snippet             basetypes.StringValue `tfsdk:"snippet"`
	SslForwardProxy     basetypes.ObjectValue `tfsdk:"ssl_forward_proxy"`
	SslInboundProxy     basetypes.ObjectValue `tfsdk:"ssl_inbound_proxy"`
	SslNoProxy          basetypes.ObjectValue `tfsdk:"ssl_no_proxy"`
	SslProtocolSettings basetypes.ObjectValue `tfsdk:"ssl_protocol_settings"`
}

// DecryptionProfilesSslForwardProxy represents a nested structure within the DecryptionProfiles model
type DecryptionProfilesSslForwardProxy struct {
	AutoIncludeAltname            basetypes.BoolValue `tfsdk:"auto_include_altname"`
	BlockClientCert               basetypes.BoolValue `tfsdk:"block_client_cert"`
	BlockExpiredCertificate       basetypes.BoolValue `tfsdk:"block_expired_certificate"`
	BlockTimeoutCert              basetypes.BoolValue `tfsdk:"block_timeout_cert"`
	BlockTls13DowngradeNoResource basetypes.BoolValue `tfsdk:"block_tls13_downgrade_no_resource"`
	BlockUnknownCert              basetypes.BoolValue `tfsdk:"block_unknown_cert"`
	BlockUnsupportedCipher        basetypes.BoolValue `tfsdk:"block_unsupported_cipher"`
	BlockUnsupportedVersion       basetypes.BoolValue `tfsdk:"block_unsupported_version"`
	BlockUntrustedIssuer          basetypes.BoolValue `tfsdk:"block_untrusted_issuer"`
	RestrictCertExts              basetypes.BoolValue `tfsdk:"restrict_cert_exts"`
	StripAlpn                     basetypes.BoolValue `tfsdk:"strip_alpn"`
}

// DecryptionProfilesSslInboundProxy represents a nested structure within the DecryptionProfiles model
type DecryptionProfilesSslInboundProxy struct {
	BlockIfHsmUnavailable   basetypes.BoolValue `tfsdk:"block_if_hsm_unavailable"`
	BlockIfNoResource       basetypes.BoolValue `tfsdk:"block_if_no_resource"`
	BlockUnsupportedCipher  basetypes.BoolValue `tfsdk:"block_unsupported_cipher"`
	BlockUnsupportedVersion basetypes.BoolValue `tfsdk:"block_unsupported_version"`
}

// DecryptionProfilesSslNoProxy represents a nested structure within the DecryptionProfiles model
type DecryptionProfilesSslNoProxy struct {
	BlockExpiredCertificate basetypes.BoolValue `tfsdk:"block_expired_certificate"`
	BlockUntrustedIssuer    basetypes.BoolValue `tfsdk:"block_untrusted_issuer"`
}

// DecryptionProfilesSslProtocolSettings represents a nested structure within the DecryptionProfiles model
type DecryptionProfilesSslProtocolSettings struct {
	AuthAlgoMd5             basetypes.BoolValue   `tfsdk:"auth_algo_md5"`
	AuthAlgoSha1            basetypes.BoolValue   `tfsdk:"auth_algo_sha1"`
	AuthAlgoSha256          basetypes.BoolValue   `tfsdk:"auth_algo_sha256"`
	AuthAlgoSha384          basetypes.BoolValue   `tfsdk:"auth_algo_sha384"`
	EncAlgo3des             basetypes.BoolValue   `tfsdk:"enc_algo_3des"`
	EncAlgoAes128Cbc        basetypes.BoolValue   `tfsdk:"enc_algo_aes_128_cbc"`
	EncAlgoAes128Gcm        basetypes.BoolValue   `tfsdk:"enc_algo_aes_128_gcm"`
	EncAlgoAes256Cbc        basetypes.BoolValue   `tfsdk:"enc_algo_aes_256_cbc"`
	EncAlgoAes256Gcm        basetypes.BoolValue   `tfsdk:"enc_algo_aes_256_gcm"`
	EncAlgoChacha20Poly1305 basetypes.BoolValue   `tfsdk:"enc_algo_chacha20_poly1305"`
	EncAlgoRc4              basetypes.BoolValue   `tfsdk:"enc_algo_rc4"`
	KeyxchgAlgoDhe          basetypes.BoolValue   `tfsdk:"keyxchg_algo_dhe"`
	KeyxchgAlgoEcdhe        basetypes.BoolValue   `tfsdk:"keyxchg_algo_ecdhe"`
	KeyxchgAlgoRsa          basetypes.BoolValue   `tfsdk:"keyxchg_algo_rsa"`
	MaxVersion              basetypes.StringValue `tfsdk:"max_version"`
	MinVersion              basetypes.StringValue `tfsdk:"min_version"`
}

// AttrTypes defines the attribute types for the DecryptionProfiles model.
func (o DecryptionProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":    basetypes.StringType{},
		"device":  basetypes.StringType{},
		"folder":  basetypes.StringType{},
		"id":      basetypes.StringType{},
		"name":    basetypes.StringType{},
		"snippet": basetypes.StringType{},
		"ssl_forward_proxy": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"auto_include_altname":              basetypes.BoolType{},
				"block_client_cert":                 basetypes.BoolType{},
				"block_expired_certificate":         basetypes.BoolType{},
				"block_timeout_cert":                basetypes.BoolType{},
				"block_tls13_downgrade_no_resource": basetypes.BoolType{},
				"block_unknown_cert":                basetypes.BoolType{},
				"block_unsupported_cipher":          basetypes.BoolType{},
				"block_unsupported_version":         basetypes.BoolType{},
				"block_untrusted_issuer":            basetypes.BoolType{},
				"restrict_cert_exts":                basetypes.BoolType{},
				"strip_alpn":                        basetypes.BoolType{},
			},
		},
		"ssl_inbound_proxy": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"block_if_hsm_unavailable":  basetypes.BoolType{},
				"block_if_no_resource":      basetypes.BoolType{},
				"block_unsupported_cipher":  basetypes.BoolType{},
				"block_unsupported_version": basetypes.BoolType{},
			},
		},
		"ssl_no_proxy": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"block_expired_certificate": basetypes.BoolType{},
				"block_untrusted_issuer":    basetypes.BoolType{},
			},
		},
		"ssl_protocol_settings": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"auth_algo_md5":              basetypes.BoolType{},
				"auth_algo_sha1":             basetypes.BoolType{},
				"auth_algo_sha256":           basetypes.BoolType{},
				"auth_algo_sha384":           basetypes.BoolType{},
				"enc_algo_3des":              basetypes.BoolType{},
				"enc_algo_aes_128_cbc":       basetypes.BoolType{},
				"enc_algo_aes_128_gcm":       basetypes.BoolType{},
				"enc_algo_aes_256_cbc":       basetypes.BoolType{},
				"enc_algo_aes_256_gcm":       basetypes.BoolType{},
				"enc_algo_chacha20_poly1305": basetypes.BoolType{},
				"enc_algo_rc4":               basetypes.BoolType{},
				"keyxchg_algo_dhe":           basetypes.BoolType{},
				"keyxchg_algo_ecdhe":         basetypes.BoolType{},
				"keyxchg_algo_rsa":           basetypes.BoolType{},
				"max_version":                basetypes.StringType{},
				"min_version":                basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of DecryptionProfiles objects.
func (o DecryptionProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DecryptionProfilesSslForwardProxy model.
func (o DecryptionProfilesSslForwardProxy) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"auto_include_altname":              basetypes.BoolType{},
		"block_client_cert":                 basetypes.BoolType{},
		"block_expired_certificate":         basetypes.BoolType{},
		"block_timeout_cert":                basetypes.BoolType{},
		"block_tls13_downgrade_no_resource": basetypes.BoolType{},
		"block_unknown_cert":                basetypes.BoolType{},
		"block_unsupported_cipher":          basetypes.BoolType{},
		"block_unsupported_version":         basetypes.BoolType{},
		"block_untrusted_issuer":            basetypes.BoolType{},
		"restrict_cert_exts":                basetypes.BoolType{},
		"strip_alpn":                        basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of DecryptionProfilesSslForwardProxy objects.
func (o DecryptionProfilesSslForwardProxy) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DecryptionProfilesSslInboundProxy model.
func (o DecryptionProfilesSslInboundProxy) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"block_if_hsm_unavailable":  basetypes.BoolType{},
		"block_if_no_resource":      basetypes.BoolType{},
		"block_unsupported_cipher":  basetypes.BoolType{},
		"block_unsupported_version": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of DecryptionProfilesSslInboundProxy objects.
func (o DecryptionProfilesSslInboundProxy) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DecryptionProfilesSslNoProxy model.
func (o DecryptionProfilesSslNoProxy) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"block_expired_certificate": basetypes.BoolType{},
		"block_untrusted_issuer":    basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of DecryptionProfilesSslNoProxy objects.
func (o DecryptionProfilesSslNoProxy) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DecryptionProfilesSslProtocolSettings model.
func (o DecryptionProfilesSslProtocolSettings) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"auth_algo_md5":              basetypes.BoolType{},
		"auth_algo_sha1":             basetypes.BoolType{},
		"auth_algo_sha256":           basetypes.BoolType{},
		"auth_algo_sha384":           basetypes.BoolType{},
		"enc_algo_3des":              basetypes.BoolType{},
		"enc_algo_aes_128_cbc":       basetypes.BoolType{},
		"enc_algo_aes_128_gcm":       basetypes.BoolType{},
		"enc_algo_aes_256_cbc":       basetypes.BoolType{},
		"enc_algo_aes_256_gcm":       basetypes.BoolType{},
		"enc_algo_chacha20_poly1305": basetypes.BoolType{},
		"enc_algo_rc4":               basetypes.BoolType{},
		"keyxchg_algo_dhe":           basetypes.BoolType{},
		"keyxchg_algo_ecdhe":         basetypes.BoolType{},
		"keyxchg_algo_rsa":           basetypes.BoolType{},
		"max_version":                basetypes.StringType{},
		"min_version":                basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DecryptionProfilesSslProtocolSettings objects.
func (o DecryptionProfilesSslProtocolSettings) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// DecryptionProfilesResourceSchema defines the schema for DecryptionProfiles resource
var DecryptionProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "Manages SCM DecryptionProfiles objects",
	Attributes: map[string]schema.Attribute{
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
		"id": schema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.RegexMatches(regexp.MustCompile("^[A-Za-z0-9]{1}[A-Za-z0-9_\\-\\.\\s]{0,}$"), "pattern must match "+"^[A-Za-z0-9]{1}[A-Za-z0-9_\\-\\.\\s]{0,}$"),
			},
			MarkdownDescription: "Must start with alphanumeric char and should contain only alphanemeric, underscore, hyphen, dot or space",
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
		"ssl_forward_proxy": schema.SingleNestedAttribute{
			MarkdownDescription: "Ssl forward proxy",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"auto_include_altname": schema.BoolAttribute{
					MarkdownDescription: "Auto include altname",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"block_client_cert": schema.BoolAttribute{
					MarkdownDescription: "Block client cert",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"block_expired_certificate": schema.BoolAttribute{
					MarkdownDescription: "Block expired certificate",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"block_timeout_cert": schema.BoolAttribute{
					MarkdownDescription: "Block timeout cert",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"block_tls13_downgrade_no_resource": schema.BoolAttribute{
					MarkdownDescription: "Block tls13 downgrade no resource",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"block_unknown_cert": schema.BoolAttribute{
					MarkdownDescription: "Block unknown cert",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"block_unsupported_cipher": schema.BoolAttribute{
					MarkdownDescription: "Block unsupported cipher",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"block_unsupported_version": schema.BoolAttribute{
					MarkdownDescription: "Block unsupported version",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"block_untrusted_issuer": schema.BoolAttribute{
					MarkdownDescription: "Block untrusted issuer",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"restrict_cert_exts": schema.BoolAttribute{
					MarkdownDescription: "Restrict cert exts",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"strip_alpn": schema.BoolAttribute{
					MarkdownDescription: "Strip alpn",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
			},
		},
		"ssl_inbound_proxy": schema.SingleNestedAttribute{
			MarkdownDescription: "Ssl inbound proxy",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"block_if_hsm_unavailable": schema.BoolAttribute{
					MarkdownDescription: "Block if hsm unavailable",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"block_if_no_resource": schema.BoolAttribute{
					MarkdownDescription: "Block if no resource",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"block_unsupported_cipher": schema.BoolAttribute{
					MarkdownDescription: "Block unsupported cipher",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"block_unsupported_version": schema.BoolAttribute{
					MarkdownDescription: "Block unsupported version",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
			},
		},
		"ssl_no_proxy": schema.SingleNestedAttribute{
			MarkdownDescription: "Ssl no proxy",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"block_expired_certificate": schema.BoolAttribute{
					MarkdownDescription: "Block expired certificate",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"block_untrusted_issuer": schema.BoolAttribute{
					MarkdownDescription: "Block untrusted issuer",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
			},
		},
		"ssl_protocol_settings": schema.SingleNestedAttribute{
			MarkdownDescription: "Ssl protocol settings",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"auth_algo_md5": schema.BoolAttribute{
					MarkdownDescription: "Auth algo md5",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"auth_algo_sha1": schema.BoolAttribute{
					MarkdownDescription: "Auth algo sha1",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"auth_algo_sha256": schema.BoolAttribute{
					MarkdownDescription: "Auth algo sha256",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"auth_algo_sha384": schema.BoolAttribute{
					MarkdownDescription: "Auth algo sha384",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"enc_algo_3des": schema.BoolAttribute{
					MarkdownDescription: "Enc algo3des",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"enc_algo_aes_128_cbc": schema.BoolAttribute{
					MarkdownDescription: "Enc algo aes128 cbc",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"enc_algo_aes_128_gcm": schema.BoolAttribute{
					MarkdownDescription: "Enc algo aes128 gcm",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"enc_algo_aes_256_cbc": schema.BoolAttribute{
					MarkdownDescription: "Enc algo aes256 cbc",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"enc_algo_aes_256_gcm": schema.BoolAttribute{
					MarkdownDescription: "Enc algo aes256 gcm",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"enc_algo_chacha20_poly1305": schema.BoolAttribute{
					MarkdownDescription: "Enc algo chacha20 poly1305",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"enc_algo_rc4": schema.BoolAttribute{
					MarkdownDescription: "Enc algo rc4",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"keyxchg_algo_dhe": schema.BoolAttribute{
					MarkdownDescription: "Keyxchg algo dhe",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"keyxchg_algo_ecdhe": schema.BoolAttribute{
					MarkdownDescription: "Keyxchg algo ecdhe",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"keyxchg_algo_rsa": schema.BoolAttribute{
					MarkdownDescription: "Keyxchg algo rsa",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"max_version": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("sslv3", "tls1-0", "tls1-1", "tls1-2", "tls1-3", "max"),
					},
					MarkdownDescription: "Max version",
					Optional:            true,
					Computed:            true,
					Default:             stringdefault.StaticString("tls1-2"),
				},
				"min_version": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("sslv3", "tls1-0", "tls1-1", "tls1-2", "tls1-3"),
					},
					MarkdownDescription: "Min version",
					Optional:            true,
					Computed:            true,
					Default:             stringdefault.StaticString("tls1-0"),
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
	},
}

// DecryptionProfilesDataSourceSchema defines the schema for DecryptionProfiles data source
var DecryptionProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "DecryptionProfiles data source",
	Attributes: map[string]dsschema.Attribute{
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
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Must start with alphanumeric char and should contain only alphanemeric, underscore, hyphen, dot or space",
			Optional:            true,
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"ssl_forward_proxy": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Ssl forward proxy",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"auto_include_altname": dsschema.BoolAttribute{
					MarkdownDescription: "Auto include altname",
					Computed:            true,
				},
				"block_client_cert": dsschema.BoolAttribute{
					MarkdownDescription: "Block client cert",
					Computed:            true,
				},
				"block_expired_certificate": dsschema.BoolAttribute{
					MarkdownDescription: "Block expired certificate",
					Computed:            true,
				},
				"block_timeout_cert": dsschema.BoolAttribute{
					MarkdownDescription: "Block timeout cert",
					Computed:            true,
				},
				"block_tls13_downgrade_no_resource": dsschema.BoolAttribute{
					MarkdownDescription: "Block tls13 downgrade no resource",
					Computed:            true,
				},
				"block_unknown_cert": dsschema.BoolAttribute{
					MarkdownDescription: "Block unknown cert",
					Computed:            true,
				},
				"block_unsupported_cipher": dsschema.BoolAttribute{
					MarkdownDescription: "Block unsupported cipher",
					Computed:            true,
				},
				"block_unsupported_version": dsschema.BoolAttribute{
					MarkdownDescription: "Block unsupported version",
					Computed:            true,
				},
				"block_untrusted_issuer": dsschema.BoolAttribute{
					MarkdownDescription: "Block untrusted issuer",
					Computed:            true,
				},
				"restrict_cert_exts": dsschema.BoolAttribute{
					MarkdownDescription: "Restrict cert exts",
					Computed:            true,
				},
				"strip_alpn": dsschema.BoolAttribute{
					MarkdownDescription: "Strip alpn",
					Computed:            true,
				},
			},
		},
		"ssl_inbound_proxy": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Ssl inbound proxy",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"block_if_hsm_unavailable": dsschema.BoolAttribute{
					MarkdownDescription: "Block if hsm unavailable",
					Computed:            true,
				},
				"block_if_no_resource": dsschema.BoolAttribute{
					MarkdownDescription: "Block if no resource",
					Computed:            true,
				},
				"block_unsupported_cipher": dsschema.BoolAttribute{
					MarkdownDescription: "Block unsupported cipher",
					Computed:            true,
				},
				"block_unsupported_version": dsschema.BoolAttribute{
					MarkdownDescription: "Block unsupported version",
					Computed:            true,
				},
			},
		},
		"ssl_no_proxy": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Ssl no proxy",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"block_expired_certificate": dsschema.BoolAttribute{
					MarkdownDescription: "Block expired certificate",
					Computed:            true,
				},
				"block_untrusted_issuer": dsschema.BoolAttribute{
					MarkdownDescription: "Block untrusted issuer",
					Computed:            true,
				},
			},
		},
		"ssl_protocol_settings": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Ssl protocol settings",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"auth_algo_md5": dsschema.BoolAttribute{
					MarkdownDescription: "Auth algo md5",
					Computed:            true,
				},
				"auth_algo_sha1": dsschema.BoolAttribute{
					MarkdownDescription: "Auth algo sha1",
					Computed:            true,
				},
				"auth_algo_sha256": dsschema.BoolAttribute{
					MarkdownDescription: "Auth algo sha256",
					Computed:            true,
				},
				"auth_algo_sha384": dsschema.BoolAttribute{
					MarkdownDescription: "Auth algo sha384",
					Computed:            true,
				},
				"enc_algo_3des": dsschema.BoolAttribute{
					MarkdownDescription: "Enc algo3des",
					Computed:            true,
				},
				"enc_algo_aes_128_cbc": dsschema.BoolAttribute{
					MarkdownDescription: "Enc algo aes128 cbc",
					Computed:            true,
				},
				"enc_algo_aes_128_gcm": dsschema.BoolAttribute{
					MarkdownDescription: "Enc algo aes128 gcm",
					Computed:            true,
				},
				"enc_algo_aes_256_cbc": dsschema.BoolAttribute{
					MarkdownDescription: "Enc algo aes256 cbc",
					Computed:            true,
				},
				"enc_algo_aes_256_gcm": dsschema.BoolAttribute{
					MarkdownDescription: "Enc algo aes256 gcm",
					Computed:            true,
				},
				"enc_algo_chacha20_poly1305": dsschema.BoolAttribute{
					MarkdownDescription: "Enc algo chacha20 poly1305",
					Computed:            true,
				},
				"enc_algo_rc4": dsschema.BoolAttribute{
					MarkdownDescription: "Enc algo rc4",
					Computed:            true,
				},
				"keyxchg_algo_dhe": dsschema.BoolAttribute{
					MarkdownDescription: "Keyxchg algo dhe",
					Computed:            true,
				},
				"keyxchg_algo_ecdhe": dsschema.BoolAttribute{
					MarkdownDescription: "Keyxchg algo ecdhe",
					Computed:            true,
				},
				"keyxchg_algo_rsa": dsschema.BoolAttribute{
					MarkdownDescription: "Keyxchg algo rsa",
					Computed:            true,
				},
				"max_version": dsschema.StringAttribute{
					MarkdownDescription: "Max version",
					Computed:            true,
				},
				"min_version": dsschema.StringAttribute{
					MarkdownDescription: "Min version",
					Computed:            true,
				},
			},
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// DecryptionProfilesListModel represents the data model for a list data source.
type DecryptionProfilesListModel struct {
	Tfid    types.String         `tfsdk:"tfid"`
	Data    []DecryptionProfiles `tfsdk:"data"`
	Limit   types.Int64          `tfsdk:"limit"`
	Offset  types.Int64          `tfsdk:"offset"`
	Name    types.String         `tfsdk:"name"`
	Total   types.Int64          `tfsdk:"total"`
	Folder  types.String         `tfsdk:"folder"`
	Snippet types.String         `tfsdk:"snippet"`
	Device  types.String         `tfsdk:"device"`
}

// DecryptionProfilesListDataSourceSchema defines the schema for a list data source.
var DecryptionProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: DecryptionProfilesDataSourceSchema.Attributes,
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
