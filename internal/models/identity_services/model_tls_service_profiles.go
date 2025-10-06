package models

import (
	"regexp"

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

// Package: identity_services
// This file contains models for the identity_services SDK package

// TlsServiceProfiles represents the Terraform model for TlsServiceProfiles
type TlsServiceProfiles struct {
	Tfid             types.String          `tfsdk:"tfid"`
	Certificate      basetypes.StringValue `tfsdk:"certificate"`
	Device           basetypes.StringValue `tfsdk:"device"`
	Folder           basetypes.StringValue `tfsdk:"folder"`
	Id               basetypes.StringValue `tfsdk:"id"`
	Name             basetypes.StringValue `tfsdk:"name"`
	ProtocolSettings basetypes.ObjectValue `tfsdk:"protocol_settings"`
	Snippet          basetypes.StringValue `tfsdk:"snippet"`
}

// TlsServiceProfilesProtocolSettings represents a nested structure within the TlsServiceProfiles model
type TlsServiceProfilesProtocolSettings struct {
	AuthAlgoSha1     basetypes.BoolValue   `tfsdk:"auth_algo_sha1"`
	AuthAlgoSha256   basetypes.BoolValue   `tfsdk:"auth_algo_sha256"`
	AuthAlgoSha384   basetypes.BoolValue   `tfsdk:"auth_algo_sha384"`
	EncAlgo3des      basetypes.BoolValue   `tfsdk:"enc_algo_3des"`
	EncAlgoAes128Cbc basetypes.BoolValue   `tfsdk:"enc_algo_aes_128_cbc"`
	EncAlgoAes128Gcm basetypes.BoolValue   `tfsdk:"enc_algo_aes_128_gcm"`
	EncAlgoAes256Cbc basetypes.BoolValue   `tfsdk:"enc_algo_aes_256_cbc"`
	EncAlgoAes256Gcm basetypes.BoolValue   `tfsdk:"enc_algo_aes_256_gcm"`
	EncAlgoRc4       basetypes.BoolValue   `tfsdk:"enc_algo_rc4"`
	KeyxchgAlgoDhe   basetypes.BoolValue   `tfsdk:"keyxchg_algo_dhe"`
	KeyxchgAlgoEcdhe basetypes.BoolValue   `tfsdk:"keyxchg_algo_ecdhe"`
	KeyxchgAlgoRsa   basetypes.BoolValue   `tfsdk:"keyxchg_algo_rsa"`
	MaxVersion       basetypes.StringValue `tfsdk:"max_version"`
	MinVersion       basetypes.StringValue `tfsdk:"min_version"`
}

// AttrTypes defines the attribute types for the TlsServiceProfiles model.
func (o TlsServiceProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":        basetypes.StringType{},
		"certificate": basetypes.StringType{},
		"device":      basetypes.StringType{},
		"folder":      basetypes.StringType{},
		"id":          basetypes.StringType{},
		"name":        basetypes.StringType{},
		"protocol_settings": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"auth_algo_sha1":       basetypes.BoolType{},
				"auth_algo_sha256":     basetypes.BoolType{},
				"auth_algo_sha384":     basetypes.BoolType{},
				"enc_algo_3des":        basetypes.BoolType{},
				"enc_algo_aes_128_cbc": basetypes.BoolType{},
				"enc_algo_aes_128_gcm": basetypes.BoolType{},
				"enc_algo_aes_256_cbc": basetypes.BoolType{},
				"enc_algo_aes_256_gcm": basetypes.BoolType{},
				"enc_algo_rc4":         basetypes.BoolType{},
				"keyxchg_algo_dhe":     basetypes.BoolType{},
				"keyxchg_algo_ecdhe":   basetypes.BoolType{},
				"keyxchg_algo_rsa":     basetypes.BoolType{},
				"max_version":          basetypes.StringType{},
				"min_version":          basetypes.StringType{},
			},
		},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of TlsServiceProfiles objects.
func (o TlsServiceProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the TlsServiceProfilesProtocolSettings model.
func (o TlsServiceProfilesProtocolSettings) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"auth_algo_sha1":       basetypes.BoolType{},
		"auth_algo_sha256":     basetypes.BoolType{},
		"auth_algo_sha384":     basetypes.BoolType{},
		"enc_algo_3des":        basetypes.BoolType{},
		"enc_algo_aes_128_cbc": basetypes.BoolType{},
		"enc_algo_aes_128_gcm": basetypes.BoolType{},
		"enc_algo_aes_256_cbc": basetypes.BoolType{},
		"enc_algo_aes_256_gcm": basetypes.BoolType{},
		"enc_algo_rc4":         basetypes.BoolType{},
		"keyxchg_algo_dhe":     basetypes.BoolType{},
		"keyxchg_algo_ecdhe":   basetypes.BoolType{},
		"keyxchg_algo_rsa":     basetypes.BoolType{},
		"max_version":          basetypes.StringType{},
		"min_version":          basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of TlsServiceProfilesProtocolSettings objects.
func (o TlsServiceProfilesProtocolSettings) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// TlsServiceProfilesResourceSchema defines the schema for TlsServiceProfiles resource
var TlsServiceProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "TlsServiceProfile resource",
	Attributes: map[string]schema.Attribute{
		"certificate": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(255),
			},
			MarkdownDescription: "Certificate name",
			Required:            true,
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
			MarkdownDescription: "The UUID of the TLS service profile",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(127),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9._-]+$"), "pattern must match "+"^[a-zA-Z0-9._-]+$"),
			},
			MarkdownDescription: "TLS service profile name. The value is `muCustomDomainSSLProfile` when it is used on mobile-agent infra settings.",
			Required:            true,
		},
		"protocol_settings": schema.SingleNestedAttribute{
			MarkdownDescription: "Protocol settings",
			Required:            true,
			Attributes: map[string]schema.Attribute{
				"auth_algo_sha1": schema.BoolAttribute{
					MarkdownDescription: "Allow SHA1 authentication?",
					Optional:            true,
				},
				"auth_algo_sha256": schema.BoolAttribute{
					MarkdownDescription: "Allow SHA256 authentication?",
					Optional:            true,
				},
				"auth_algo_sha384": schema.BoolAttribute{
					MarkdownDescription: "Allow SHA384 authentication?",
					Optional:            true,
				},
				"enc_algo_3des": schema.BoolAttribute{
					MarkdownDescription: "Allow 3DES algorithm?",
					Optional:            true,
				},
				"enc_algo_aes_128_cbc": schema.BoolAttribute{
					MarkdownDescription: "Allow AES-128-CBC algorithm?",
					Optional:            true,
				},
				"enc_algo_aes_128_gcm": schema.BoolAttribute{
					MarkdownDescription: "Allow AES-128-GCM algorithm?",
					Optional:            true,
				},
				"enc_algo_aes_256_cbc": schema.BoolAttribute{
					MarkdownDescription: "Allow AES-256-CBC algorithm?",
					Optional:            true,
				},
				"enc_algo_aes_256_gcm": schema.BoolAttribute{
					MarkdownDescription: "Allow algorithm AES-256-GCM",
					Optional:            true,
				},
				"enc_algo_rc4": schema.BoolAttribute{
					MarkdownDescription: "Allow RC4 algorithm?",
					Optional:            true,
				},
				"keyxchg_algo_dhe": schema.BoolAttribute{
					MarkdownDescription: "Allow DHE algorithm?",
					Optional:            true,
				},
				"keyxchg_algo_ecdhe": schema.BoolAttribute{
					MarkdownDescription: "Allow ECDHE algorithm?",
					Optional:            true,
				},
				"keyxchg_algo_rsa": schema.BoolAttribute{
					MarkdownDescription: "Allow RSA algorithm?",
					Optional:            true,
				},
				"max_version": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("tls1-0", "tls1-1", "tls1-2", "tls1-3"),
					},
					MarkdownDescription: "Maximum TLS version",
					Optional:            true,
					Computed:            true,
					Default:             stringdefault.StaticString("tls1-3"),
				},
				"min_version": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("tls1-0", "tls1-1", "tls1-2"),
					},
					MarkdownDescription: "Minimum TLS version",
					Optional:            true,
					Computed:            true,
					Default:             stringdefault.StaticString("tls1-2"),
				},
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

// TlsServiceProfilesDataSourceSchema defines the schema for TlsServiceProfiles data source
var TlsServiceProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "TlsServiceProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"certificate": dsschema.StringAttribute{
			MarkdownDescription: "Certificate name",
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
			MarkdownDescription: "The UUID of the TLS service profile",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "TLS service profile name. The value is `muCustomDomainSSLProfile` when it is used on mobile-agent infra settings.",
			Optional:            true,
			Computed:            true,
		},
		"protocol_settings": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Protocol settings",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"auth_algo_sha1": dsschema.BoolAttribute{
					MarkdownDescription: "Allow SHA1 authentication?",
					Computed:            true,
				},
				"auth_algo_sha256": dsschema.BoolAttribute{
					MarkdownDescription: "Allow SHA256 authentication?",
					Computed:            true,
				},
				"auth_algo_sha384": dsschema.BoolAttribute{
					MarkdownDescription: "Allow SHA384 authentication?",
					Computed:            true,
				},
				"enc_algo_3des": dsschema.BoolAttribute{
					MarkdownDescription: "Allow 3DES algorithm?",
					Computed:            true,
				},
				"enc_algo_aes_128_cbc": dsschema.BoolAttribute{
					MarkdownDescription: "Allow AES-128-CBC algorithm?",
					Computed:            true,
				},
				"enc_algo_aes_128_gcm": dsschema.BoolAttribute{
					MarkdownDescription: "Allow AES-128-GCM algorithm?",
					Computed:            true,
				},
				"enc_algo_aes_256_cbc": dsschema.BoolAttribute{
					MarkdownDescription: "Allow AES-256-CBC algorithm?",
					Computed:            true,
				},
				"enc_algo_aes_256_gcm": dsschema.BoolAttribute{
					MarkdownDescription: "Allow algorithm AES-256-GCM",
					Computed:            true,
				},
				"enc_algo_rc4": dsschema.BoolAttribute{
					MarkdownDescription: "Allow RC4 algorithm?",
					Computed:            true,
				},
				"keyxchg_algo_dhe": dsschema.BoolAttribute{
					MarkdownDescription: "Allow DHE algorithm?",
					Computed:            true,
				},
				"keyxchg_algo_ecdhe": dsschema.BoolAttribute{
					MarkdownDescription: "Allow ECDHE algorithm?",
					Computed:            true,
				},
				"keyxchg_algo_rsa": dsschema.BoolAttribute{
					MarkdownDescription: "Allow RSA algorithm?",
					Computed:            true,
				},
				"max_version": dsschema.StringAttribute{
					MarkdownDescription: "Maximum TLS version",
					Computed:            true,
				},
				"min_version": dsschema.StringAttribute{
					MarkdownDescription: "Minimum TLS version",
					Computed:            true,
				},
			},
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

// TlsServiceProfilesListModel represents the data model for a list data source.
type TlsServiceProfilesListModel struct {
	Tfid    types.String         `tfsdk:"tfid"`
	Data    []TlsServiceProfiles `tfsdk:"data"`
	Limit   types.Int64          `tfsdk:"limit"`
	Offset  types.Int64          `tfsdk:"offset"`
	Name    types.String         `tfsdk:"name"`
	Total   types.Int64          `tfsdk:"total"`
	Folder  types.String         `tfsdk:"folder"`
	Snippet types.String         `tfsdk:"snippet"`
	Device  types.String         `tfsdk:"device"`
}

// TlsServiceProfilesListDataSourceSchema defines the schema for a list data source.
var TlsServiceProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: TlsServiceProfilesDataSourceSchema.Attributes,
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
