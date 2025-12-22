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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: identity_services
// This file contains models for the identity_services SDK package

// ScepProfiles represents the Terraform model for ScepProfiles
type ScepProfiles struct {
	Tfid                  types.String          `tfsdk:"tfid"`
	EncryptedValues       basetypes.MapValue    `tfsdk:"encrypted_values"`
	Algorithm             basetypes.ObjectValue `tfsdk:"algorithm"`
	CaIdentityName        basetypes.StringValue `tfsdk:"ca_identity_name"`
	CertificateAttributes basetypes.ObjectValue `tfsdk:"certificate_attributes"`
	Device                basetypes.StringValue `tfsdk:"device"`
	Digest                basetypes.StringValue `tfsdk:"digest"`
	Fingerprint           basetypes.StringValue `tfsdk:"fingerprint"`
	Folder                basetypes.StringValue `tfsdk:"folder"`
	Id                    basetypes.StringValue `tfsdk:"id"`
	Name                  basetypes.StringValue `tfsdk:"name"`
	ScepCaCert            basetypes.StringValue `tfsdk:"scep_ca_cert"`
	ScepChallenge         basetypes.ObjectValue `tfsdk:"scep_challenge"`
	ScepClientCert        basetypes.StringValue `tfsdk:"scep_client_cert"`
	ScepUrl               basetypes.StringValue `tfsdk:"scep_url"`
	Snippet               basetypes.StringValue `tfsdk:"snippet"`
	Subject               basetypes.StringValue `tfsdk:"subject"`
	UseAsDigitalSignature basetypes.BoolValue   `tfsdk:"use_as_digital_signature"`
	UseForKeyEncipherment basetypes.BoolValue   `tfsdk:"use_for_key_encipherment"`
}

// ScepProfilesAlgorithm represents a nested structure within the ScepProfiles model
type ScepProfilesAlgorithm struct {
	Rsa basetypes.ObjectValue `tfsdk:"rsa"`
}

// ScepProfilesAlgorithmRsa represents a nested structure within the ScepProfiles model
type ScepProfilesAlgorithmRsa struct {
	RsaNbits basetypes.Int64Value `tfsdk:"rsa_nbits"`
}

// ScepProfilesCertificateAttributes represents a nested structure within the ScepProfiles model
type ScepProfilesCertificateAttributes struct {
	Dnsname                   basetypes.StringValue `tfsdk:"dnsname"`
	Rfc822name                basetypes.StringValue `tfsdk:"rfc822name"`
	UniformResourceIdentifier basetypes.StringValue `tfsdk:"uniform_resource_identifier"`
}

// ScepProfilesScepChallenge represents a nested structure within the ScepProfiles model
type ScepProfilesScepChallenge struct {
	Dynamic basetypes.ObjectValue `tfsdk:"dynamic"`
	Fixed   basetypes.StringValue `tfsdk:"fixed"`
	None    basetypes.StringValue `tfsdk:"none"`
}

// ScepProfilesScepChallengeDynamic represents a nested structure within the ScepProfiles model
type ScepProfilesScepChallengeDynamic struct {
	OtpServerUrl basetypes.StringValue `tfsdk:"otp_server_url"`
	Password     basetypes.StringValue `tfsdk:"password"`
	Username     basetypes.StringValue `tfsdk:"username"`
}

// AttrTypes defines the attribute types for the ScepProfiles model.
func (o ScepProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":             basetypes.StringType{},
		"encrypted_values": basetypes.MapType{ElemType: basetypes.StringType{}},
		"algorithm": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"rsa": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"rsa_nbits": basetypes.Int64Type{},
					},
				},
			},
		},
		"ca_identity_name": basetypes.StringType{},
		"certificate_attributes": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"dnsname":                     basetypes.StringType{},
				"rfc822name":                  basetypes.StringType{},
				"uniform_resource_identifier": basetypes.StringType{},
			},
		},
		"device":       basetypes.StringType{},
		"digest":       basetypes.StringType{},
		"fingerprint":  basetypes.StringType{},
		"folder":       basetypes.StringType{},
		"id":           basetypes.StringType{},
		"name":         basetypes.StringType{},
		"scep_ca_cert": basetypes.StringType{},
		"scep_challenge": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"dynamic": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"otp_server_url": basetypes.StringType{},
						"password":       basetypes.StringType{},
						"username":       basetypes.StringType{},
					},
				},
				"fixed": basetypes.StringType{},
				"none":  basetypes.StringType{},
			},
		},
		"scep_client_cert":         basetypes.StringType{},
		"scep_url":                 basetypes.StringType{},
		"snippet":                  basetypes.StringType{},
		"subject":                  basetypes.StringType{},
		"use_as_digital_signature": basetypes.BoolType{},
		"use_for_key_encipherment": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of ScepProfiles objects.
func (o ScepProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ScepProfilesAlgorithm model.
func (o ScepProfilesAlgorithm) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"rsa": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"rsa_nbits": basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ScepProfilesAlgorithm objects.
func (o ScepProfilesAlgorithm) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ScepProfilesAlgorithmRsa model.
func (o ScepProfilesAlgorithmRsa) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"rsa_nbits": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of ScepProfilesAlgorithmRsa objects.
func (o ScepProfilesAlgorithmRsa) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ScepProfilesCertificateAttributes model.
func (o ScepProfilesCertificateAttributes) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"dnsname":                     basetypes.StringType{},
		"rfc822name":                  basetypes.StringType{},
		"uniform_resource_identifier": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ScepProfilesCertificateAttributes objects.
func (o ScepProfilesCertificateAttributes) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ScepProfilesScepChallenge model.
func (o ScepProfilesScepChallenge) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"dynamic": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"otp_server_url": basetypes.StringType{},
				"password":       basetypes.StringType{},
				"username":       basetypes.StringType{},
			},
		},
		"fixed": basetypes.StringType{},
		"none":  basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ScepProfilesScepChallenge objects.
func (o ScepProfilesScepChallenge) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ScepProfilesScepChallengeDynamic model.
func (o ScepProfilesScepChallengeDynamic) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"otp_server_url": basetypes.StringType{},
		"password":       basetypes.StringType{},
		"username":       basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ScepProfilesScepChallengeDynamic objects.
func (o ScepProfilesScepChallengeDynamic) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// ScepProfilesResourceSchema defines the schema for ScepProfiles resource
var ScepProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "ScepProfile resource",
	Attributes: map[string]schema.Attribute{
		"algorithm": schema.SingleNestedAttribute{
			MarkdownDescription: "Algorithm",
			Required:            true,
			Attributes: map[string]schema.Attribute{
				"rsa": schema.SingleNestedAttribute{
					MarkdownDescription: "Key length (bits)",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"rsa_nbits": schema.Int64Attribute{
							MarkdownDescription: "Rsa nbits",
							Optional:            true,
						},
					},
				},
			},
		},
		"ca_identity_name": schema.StringAttribute{
			MarkdownDescription: "Certificate Authority identity",
			Required:            true,
		},
		"certificate_attributes": schema.SingleNestedAttribute{
			MarkdownDescription: "Subject Alternative name type",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"dnsname": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("rfc822name"),
							path.MatchRelative().AtParent().AtName("uniform_resource_identifier"),
						),
					},
					MarkdownDescription: "Dnsname\n\n> ℹ️ **Note:** You must specify exactly one of `dnsname`, `rfc822name`, and `uniform_resource_identifier`.",
					Optional:            true,
				},
				"rfc822name": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("dnsname"),
							path.MatchRelative().AtParent().AtName("uniform_resource_identifier"),
						),
					},
					MarkdownDescription: "Rfc822name\n\n> ℹ️ **Note:** You must specify exactly one of `dnsname`, `rfc822name`, and `uniform_resource_identifier`.",
					Optional:            true,
				},
				"uniform_resource_identifier": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("dnsname"),
							path.MatchRelative().AtParent().AtName("rfc822name"),
						),
					},
					MarkdownDescription: "Uniform resource identifier\n\n> ℹ️ **Note:** You must specify exactly one of `dnsname`, `rfc822name`, and `uniform_resource_identifier`.",
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
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
			},
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"digest": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("sha1", "sha256", "sha348", "sha512"),
			},
			MarkdownDescription: "Digest for CSR",
			Required:            true,
		},
		"encrypted_values": schema.MapAttribute{
			ElementType:         basetypes.StringType{},
			MarkdownDescription: "Map of sensitive values returned from the API.",
			Computed:            true,
			Sensitive:           true,
		},
		"fingerprint": schema.StringAttribute{
			MarkdownDescription: "CA certificate fingerprint",
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
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the SCEP profile",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(31),
			},
			MarkdownDescription: "The name of the SCEP profile",
			Required:            true,
		},
		"scep_ca_cert": schema.StringAttribute{
			MarkdownDescription: "SCEP server CA certificate",
			Optional:            true,
		},
		"scep_challenge": schema.SingleNestedAttribute{
			MarkdownDescription: "One Time Password challenge",
			Required:            true,
			Attributes: map[string]schema.Attribute{
				"dynamic": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("fixed"),
							path.MatchRelative().AtParent().AtName("none"),
						),
					},
					MarkdownDescription: "Dynamic\n\n> ℹ️ **Note:** You must specify exactly one of `dynamic`, `fixed`, and `none`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"otp_server_url": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(255),
							},
							MarkdownDescription: "OTP server URL",
							Optional:            true,
						},
						"password": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(255),
							},
							MarkdownDescription: "OTP password",
							Optional:            true,
							Sensitive:           true,
						},
						"username": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(255),
							},
							MarkdownDescription: "OTP username",
							Optional:            true,
						},
					},
				},
				"fixed": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("dynamic"),
							path.MatchRelative().AtParent().AtName("none"),
						),
						stringvalidator.LengthAtMost(1024),
					},
					MarkdownDescription: "Challenge to use for SCEP server on mobile clients\n\n> ℹ️ **Note:** You must specify exactly one of `dynamic`, `fixed`, and `none`.",
					Optional:            true,
				},
				"none": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("dynamic"),
							path.MatchRelative().AtParent().AtName("fixed"),
						),
						stringvalidator.OneOf(""),
					},
					MarkdownDescription: "No OTP\n\n> ℹ️ **Note:** You must specify exactly one of `dynamic`, `fixed`, and `none`.",
					Optional:            true,
				},
			},
		},
		"scep_client_cert": schema.StringAttribute{
			MarkdownDescription: "SCEP client ceertificate",
			Optional:            true,
		},
		"scep_url": schema.StringAttribute{
			MarkdownDescription: "SCEP server URL",
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
			MarkdownDescription: "The snippet in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"subject": schema.StringAttribute{
			MarkdownDescription: "Subject",
			Required:            true,
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"use_as_digital_signature": schema.BoolAttribute{
			MarkdownDescription: "Use as digital signature?",
			Optional:            true,
		},
		"use_for_key_encipherment": schema.BoolAttribute{
			MarkdownDescription: "Use for key encipherment?",
			Optional:            true,
		},
	},
}

// ScepProfilesDataSourceSchema defines the schema for ScepProfiles data source
var ScepProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "ScepProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"algorithm": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Algorithm",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"rsa": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Key length (bits)",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"rsa_nbits": dsschema.Int64Attribute{
							MarkdownDescription: "Rsa nbits",
							Computed:            true,
						},
					},
				},
			},
		},
		"ca_identity_name": dsschema.StringAttribute{
			MarkdownDescription: "Certificate Authority identity",
			Computed:            true,
		},
		"certificate_attributes": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Subject Alternative name type",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"dnsname": dsschema.StringAttribute{
					MarkdownDescription: "Dnsname\n\n> ℹ️ **Note:** You must specify exactly one of `dnsname`, `rfc822name`, and `uniform_resource_identifier`.",
					Computed:            true,
				},
				"rfc822name": dsschema.StringAttribute{
					MarkdownDescription: "Rfc822name\n\n> ℹ️ **Note:** You must specify exactly one of `dnsname`, `rfc822name`, and `uniform_resource_identifier`.",
					Computed:            true,
				},
				"uniform_resource_identifier": dsschema.StringAttribute{
					MarkdownDescription: "Uniform resource identifier\n\n> ℹ️ **Note:** You must specify exactly one of `dnsname`, `rfc822name`, and `uniform_resource_identifier`.",
					Computed:            true,
				},
			},
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"digest": dsschema.StringAttribute{
			MarkdownDescription: "Digest for CSR",
			Computed:            true,
		},
		"encrypted_values": dsschema.MapAttribute{
			ElementType:         basetypes.StringType{},
			MarkdownDescription: "Map of sensitive values returned from the API.",
			Computed:            true,
			Sensitive:           true,
		},
		"fingerprint": dsschema.StringAttribute{
			MarkdownDescription: "CA certificate fingerprint",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the SCEP profile",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the SCEP profile",
			Optional:            true,
			Computed:            true,
		},
		"scep_ca_cert": dsschema.StringAttribute{
			MarkdownDescription: "SCEP server CA certificate",
			Computed:            true,
		},
		"scep_challenge": dsschema.SingleNestedAttribute{
			MarkdownDescription: "One Time Password challenge",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"dynamic": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Dynamic\n\n> ℹ️ **Note:** You must specify exactly one of `dynamic`, `fixed`, and `none`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"otp_server_url": dsschema.StringAttribute{
							MarkdownDescription: "OTP server URL",
							Computed:            true,
						},
						"password": dsschema.StringAttribute{
							MarkdownDescription: "OTP password",
							Computed:            true,
							Sensitive:           true,
						},
						"username": dsschema.StringAttribute{
							MarkdownDescription: "OTP username",
							Computed:            true,
						},
					},
				},
				"fixed": dsschema.StringAttribute{
					MarkdownDescription: "Challenge to use for SCEP server on mobile clients\n\n> ℹ️ **Note:** You must specify exactly one of `dynamic`, `fixed`, and `none`.",
					Computed:            true,
				},
				"none": dsschema.StringAttribute{
					MarkdownDescription: "No OTP\n\n> ℹ️ **Note:** You must specify exactly one of `dynamic`, `fixed`, and `none`.",
					Computed:            true,
				},
			},
		},
		"scep_client_cert": dsschema.StringAttribute{
			MarkdownDescription: "SCEP client ceertificate",
			Computed:            true,
		},
		"scep_url": dsschema.StringAttribute{
			MarkdownDescription: "SCEP server URL",
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"subject": dsschema.StringAttribute{
			MarkdownDescription: "Subject",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"use_as_digital_signature": dsschema.BoolAttribute{
			MarkdownDescription: "Use as digital signature?",
			Computed:            true,
		},
		"use_for_key_encipherment": dsschema.BoolAttribute{
			MarkdownDescription: "Use for key encipherment?",
			Computed:            true,
		},
	},
}

// ScepProfilesListModel represents the data model for a list data source.
type ScepProfilesListModel struct {
	Tfid    types.String   `tfsdk:"tfid"`
	Data    []ScepProfiles `tfsdk:"data"`
	Limit   types.Int64    `tfsdk:"limit"`
	Offset  types.Int64    `tfsdk:"offset"`
	Name    types.String   `tfsdk:"name"`
	Total   types.Int64    `tfsdk:"total"`
	Folder  types.String   `tfsdk:"folder"`
	Snippet types.String   `tfsdk:"snippet"`
	Device  types.String   `tfsdk:"device"`
}

// ScepProfilesListDataSourceSchema defines the schema for a list data source.
var ScepProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: ScepProfilesDataSourceSchema.Attributes,
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
