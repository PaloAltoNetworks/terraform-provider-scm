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

// CertificatesGet represents the Terraform model for CertificatesGet
type CertificatesGet struct {
	Tfid           types.String          `tfsdk:"tfid"`
	Algorithm      basetypes.StringValue `tfsdk:"algorithm"`
	Ca             basetypes.BoolValue   `tfsdk:"ca"`
	CommonName     basetypes.StringValue `tfsdk:"common_name"`
	CommonNameInt  basetypes.StringValue `tfsdk:"common_name_int"`
	Device         basetypes.StringValue `tfsdk:"device"`
	ExpiryEpoch    basetypes.StringValue `tfsdk:"expiry_epoch"`
	Folder         basetypes.StringValue `tfsdk:"folder"`
	Id             basetypes.StringValue `tfsdk:"id"`
	Issuer         basetypes.StringValue `tfsdk:"issuer"`
	IssuerHash     basetypes.StringValue `tfsdk:"issuer_hash"`
	Name           basetypes.StringValue `tfsdk:"name"`
	NotValidAfter  basetypes.StringValue `tfsdk:"not_valid_after"`
	NotValidBefore basetypes.StringValue `tfsdk:"not_valid_before"`
	PublicKey      basetypes.StringValue `tfsdk:"public_key"`
	Snippet        basetypes.StringValue `tfsdk:"snippet"`
	Subject        basetypes.StringValue `tfsdk:"subject"`
	SubjectHash    basetypes.StringValue `tfsdk:"subject_hash"`
	SubjectInt     basetypes.StringValue `tfsdk:"subject_int"`
}

// CertificatesPost represents a nested structure within the CertificatesGet model
type CertificatesPost struct {
	Tfid                   types.String          `tfsdk:"tfid"`
	Algorithm              basetypes.ObjectValue `tfsdk:"algorithm"`
	AlternateEmail         basetypes.ListValue   `tfsdk:"alternate_email"`
	CertificateName        basetypes.StringValue `tfsdk:"certificate_name"`
	CommonName             basetypes.StringValue `tfsdk:"common_name"`
	CountryCode            basetypes.StringValue `tfsdk:"country_code"`
	DayTillExpiration      basetypes.Int64Value  `tfsdk:"day_till_expiration"`
	Department             basetypes.ListValue   `tfsdk:"department"`
	Device                 basetypes.StringValue `tfsdk:"device"`
	Digest                 basetypes.StringValue `tfsdk:"digest"`
	Email                  basetypes.StringValue `tfsdk:"email"`
	Folder                 basetypes.StringValue `tfsdk:"folder"`
	Hostname               basetypes.ListValue   `tfsdk:"hostname"`
	Ip                     basetypes.ListValue   `tfsdk:"ip"`
	IsBlockPrivateKey      basetypes.BoolValue   `tfsdk:"is_block_private_key"`
	IsCertificateAuthority basetypes.BoolValue   `tfsdk:"is_certificate_authority"`
	Locality               basetypes.StringValue `tfsdk:"locality"`
	OcspResponderUrl       basetypes.StringValue `tfsdk:"ocsp_responder_url"`
	SignedBy               basetypes.StringValue `tfsdk:"signed_by"`
	Snippet                basetypes.StringValue `tfsdk:"snippet"`
	State                  basetypes.StringValue `tfsdk:"state"`
}

// CertificatesPostAlgorithm represents a nested structure within the CertificatesGet model
type CertificatesPostAlgorithm struct {
	EcdsaNumberOfBits basetypes.Float64Value `tfsdk:"ecdsa_number_of_bits"`
	RsaNumberOfBits   basetypes.Float64Value `tfsdk:"rsa_number_of_bits"`
}

// AttrTypes defines the attribute types for the CertificatesGet model.
func (o CertificatesGet) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":             basetypes.StringType{},
		"algorithm":        basetypes.StringType{},
		"ca":               basetypes.BoolType{},
		"common_name":      basetypes.StringType{},
		"common_name_int":  basetypes.StringType{},
		"device":           basetypes.StringType{},
		"expiry_epoch":     basetypes.StringType{},
		"folder":           basetypes.StringType{},
		"id":               basetypes.StringType{},
		"issuer":           basetypes.StringType{},
		"issuer_hash":      basetypes.StringType{},
		"name":             basetypes.StringType{},
		"not_valid_after":  basetypes.StringType{},
		"not_valid_before": basetypes.StringType{},
		"public_key":       basetypes.StringType{},
		"snippet":          basetypes.StringType{},
		"subject":          basetypes.StringType{},
		"subject_hash":     basetypes.StringType{},
		"subject_int":      basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of CertificatesGet objects.
func (o CertificatesGet) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the CertificatesPost model.
func (o CertificatesPost) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"algorithm": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ecdsa_number_of_bits": basetypes.Float64Type{},
				"rsa_number_of_bits":   basetypes.Float64Type{},
			},
		},
		"alternate_email":          basetypes.ListType{ElemType: basetypes.StringType{}},
		"certificate_name":         basetypes.StringType{},
		"common_name":              basetypes.StringType{},
		"country_code":             basetypes.StringType{},
		"day_till_expiration":      basetypes.Int64Type{},
		"department":               basetypes.ListType{ElemType: basetypes.StringType{}},
		"device":                   basetypes.StringType{},
		"digest":                   basetypes.StringType{},
		"email":                    basetypes.StringType{},
		"folder":                   basetypes.StringType{},
		"hostname":                 basetypes.ListType{ElemType: basetypes.StringType{}},
		"ip":                       basetypes.ListType{ElemType: basetypes.StringType{}},
		"is_block_private_key":     basetypes.BoolType{},
		"is_certificate_authority": basetypes.BoolType{},
		"locality":                 basetypes.StringType{},
		"ocsp_responder_url":       basetypes.StringType{},
		"signed_by":                basetypes.StringType{},
		"snippet":                  basetypes.StringType{},
		"state":                    basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of CertificatesPost objects.
func (o CertificatesPost) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the CertificatesPostAlgorithm model.
func (o CertificatesPostAlgorithm) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ecdsa_number_of_bits": basetypes.Float64Type{},
		"rsa_number_of_bits":   basetypes.Float64Type{},
	}
}

// AttrType returns the attribute type for a list of CertificatesPostAlgorithm objects.
func (o CertificatesPostAlgorithm) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// CertificatesGetResourceSchema defines the schema for CertificatesGet resource
var CertificatesGetResourceSchema = schema.Schema{
	MarkdownDescription: "CertificatesGet resource",
	Attributes: map[string]schema.Attribute{
		"algorithm": schema.StringAttribute{
			MarkdownDescription: "Algorithm",
			Optional:            true,
		},
		"ca": schema.BoolAttribute{
			MarkdownDescription: "CA certificate?",
			Optional:            true,
		},
		"common_name": schema.StringAttribute{
			MarkdownDescription: "Common name",
			Optional:            true,
		},
		"common_name_int": schema.StringAttribute{
			MarkdownDescription: "Common name int",
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
			MarkdownDescription: "The device in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"expiry_epoch": schema.StringAttribute{
			MarkdownDescription: "Expiry epoch",
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
			MarkdownDescription: "The folder in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the certificate",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"issuer": schema.StringAttribute{
			MarkdownDescription: "Issuer",
			Optional:            true,
		},
		"issuer_hash": schema.StringAttribute{
			MarkdownDescription: "Issue hash",
			Optional:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the certificate",
			Optional:            true,
		},
		"not_valid_after": schema.StringAttribute{
			MarkdownDescription: "Not valid after this date",
			Optional:            true,
		},
		"not_valid_before": schema.StringAttribute{
			MarkdownDescription: "Not valid before this date",
			Optional:            true,
		},
		"public_key": schema.StringAttribute{
			MarkdownDescription: "Public key",
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
		"subject": schema.StringAttribute{
			MarkdownDescription: "Subject",
			Optional:            true,
		},
		"subject_hash": schema.StringAttribute{
			MarkdownDescription: "Subject hash",
			Optional:            true,
		},
		"subject_int": schema.StringAttribute{
			MarkdownDescription: "Subject int",
			Optional:            true,
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

// CertificatesGetDataSourceSchema defines the schema for CertificatesGet data source
var CertificatesGetDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "CertificatesGet data source",
	Attributes: map[string]dsschema.Attribute{
		"algorithm": dsschema.StringAttribute{
			MarkdownDescription: "Algorithm",
			Computed:            true,
		},
		"ca": dsschema.BoolAttribute{
			MarkdownDescription: "CA certificate?",
			Computed:            true,
		},
		"common_name": dsschema.StringAttribute{
			MarkdownDescription: "Common name",
			Computed:            true,
		},
		"common_name_int": dsschema.StringAttribute{
			MarkdownDescription: "Common name int",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"expiry_epoch": dsschema.StringAttribute{
			MarkdownDescription: "Expiry epoch",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the certificate",
			Required:            true,
		},
		"issuer": dsschema.StringAttribute{
			MarkdownDescription: "Issuer",
			Computed:            true,
		},
		"issuer_hash": dsschema.StringAttribute{
			MarkdownDescription: "Issue hash",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the certificate",
			Optional:            true,
			Computed:            true,
		},
		"not_valid_after": dsschema.StringAttribute{
			MarkdownDescription: "Not valid after this date",
			Computed:            true,
		},
		"not_valid_before": dsschema.StringAttribute{
			MarkdownDescription: "Not valid before this date",
			Computed:            true,
		},
		"public_key": dsschema.StringAttribute{
			MarkdownDescription: "Public key",
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"subject": dsschema.StringAttribute{
			MarkdownDescription: "Subject",
			Computed:            true,
		},
		"subject_hash": dsschema.StringAttribute{
			MarkdownDescription: "Subject hash",
			Computed:            true,
		},
		"subject_int": dsschema.StringAttribute{
			MarkdownDescription: "Subject int",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// CertificatesGetListModel represents the data model for a list data source.
type CertificatesGetListModel struct {
	Tfid    types.String      `tfsdk:"tfid"`
	Data    []CertificatesGet `tfsdk:"data"`
	Limit   types.Int64       `tfsdk:"limit"`
	Offset  types.Int64       `tfsdk:"offset"`
	Name    types.String      `tfsdk:"name"`
	Total   types.Int64       `tfsdk:"total"`
	Folder  types.String      `tfsdk:"folder"`
	Snippet types.String      `tfsdk:"snippet"`
	Device  types.String      `tfsdk:"device"`
}

// CertificatesGetListDataSourceSchema defines the schema for a list data source.
var CertificatesGetListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: CertificatesGetDataSourceSchema.Attributes,
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
