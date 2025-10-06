package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
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

// Package: network_services
// This file contains models for the network_services SDK package

// IkeGateways represents the Terraform model for IkeGateways
type IkeGateways struct {
	Tfid            types.String          `tfsdk:"tfid"`
	EncryptedValues basetypes.MapValue    `tfsdk:"encrypted_values"`
	Authentication  basetypes.ObjectValue `tfsdk:"authentication"`
	Device          basetypes.StringValue `tfsdk:"device"`
	Folder          basetypes.StringValue `tfsdk:"folder"`
	Id              basetypes.StringValue `tfsdk:"id"`
	LocalAddress    basetypes.ObjectValue `tfsdk:"local_address"`
	LocalId         basetypes.ObjectValue `tfsdk:"local_id"`
	Name            basetypes.StringValue `tfsdk:"name"`
	PeerAddress     basetypes.ObjectValue `tfsdk:"peer_address"`
	PeerId          basetypes.ObjectValue `tfsdk:"peer_id"`
	Protocol        basetypes.ObjectValue `tfsdk:"protocol"`
	ProtocolCommon  basetypes.ObjectValue `tfsdk:"protocol_common"`
	Snippet         basetypes.StringValue `tfsdk:"snippet"`
}

// IkeGatewaysAuthentication represents a nested structure within the IkeGateways model
type IkeGatewaysAuthentication struct {
	Certificate  basetypes.ObjectValue `tfsdk:"certificate"`
	PreSharedKey basetypes.ObjectValue `tfsdk:"pre_shared_key"`
}

// IkeGatewaysAuthenticationCertificate represents a nested structure within the IkeGateways model
type IkeGatewaysAuthenticationCertificate struct {
	AllowIdPayloadMismatch     basetypes.BoolValue   `tfsdk:"allow_id_payload_mismatch"`
	CertificateProfile         basetypes.StringValue `tfsdk:"certificate_profile"`
	LocalCertificate           basetypes.ObjectValue `tfsdk:"local_certificate"`
	StrictValidationRevocation basetypes.BoolValue   `tfsdk:"strict_validation_revocation"`
	UseManagementAsSource      basetypes.BoolValue   `tfsdk:"use_management_as_source"`
}

// IkeGatewaysAuthenticationCertificateLocalCertificate represents a nested structure within the IkeGateways model
type IkeGatewaysAuthenticationCertificateLocalCertificate struct {
	LocalCertificateName basetypes.StringValue `tfsdk:"local_certificate_name"`
}

// IkeGatewaysAuthenticationPreSharedKey represents a nested structure within the IkeGateways model
type IkeGatewaysAuthenticationPreSharedKey struct {
	Key basetypes.StringValue `tfsdk:"key"`
}

// IkeGatewaysLocalAddress represents a nested structure within the IkeGateways model
type IkeGatewaysLocalAddress struct {
	Interface basetypes.StringValue `tfsdk:"interface"`
}

// IkeGatewaysLocalId represents a nested structure within the IkeGateways model
type IkeGatewaysLocalId struct {
	Id   basetypes.StringValue `tfsdk:"id"`
	Type basetypes.StringValue `tfsdk:"type"`
}

// IkeGatewaysPeerAddress represents a nested structure within the IkeGateways model
type IkeGatewaysPeerAddress struct {
	Dynamic basetypes.ObjectValue `tfsdk:"dynamic"`
	Fqdn    basetypes.StringValue `tfsdk:"fqdn"`
	Ip      basetypes.StringValue `tfsdk:"ip"`
}

// IkeGatewaysPeerId represents a nested structure within the IkeGateways model
type IkeGatewaysPeerId struct {
	Id   basetypes.StringValue `tfsdk:"id"`
	Type basetypes.StringValue `tfsdk:"type"`
}

// IkeGatewaysProtocol represents a nested structure within the IkeGateways model
type IkeGatewaysProtocol struct {
	Ikev1   basetypes.ObjectValue `tfsdk:"ikev1"`
	Ikev2   basetypes.ObjectValue `tfsdk:"ikev2"`
	Version basetypes.StringValue `tfsdk:"version"`
}

// IkeGatewaysProtocolIkev1 represents a nested structure within the IkeGateways model
type IkeGatewaysProtocolIkev1 struct {
	Dpd              basetypes.ObjectValue `tfsdk:"dpd"`
	IkeCryptoProfile basetypes.StringValue `tfsdk:"ike_crypto_profile"`
}

// IkeGatewaysProtocolIkev1Dpd represents a nested structure within the IkeGateways model
type IkeGatewaysProtocolIkev1Dpd struct {
	Enable basetypes.BoolValue `tfsdk:"enable"`
}

// IkeGatewaysProtocolCommon represents a nested structure within the IkeGateways model
type IkeGatewaysProtocolCommon struct {
	Fragmentation basetypes.ObjectValue `tfsdk:"fragmentation"`
	NatTraversal  basetypes.ObjectValue `tfsdk:"nat_traversal"`
	PassiveMode   basetypes.BoolValue   `tfsdk:"passive_mode"`
}

// IkeGatewaysProtocolCommonFragmentation represents a nested structure within the IkeGateways model
type IkeGatewaysProtocolCommonFragmentation struct {
	Enable basetypes.BoolValue `tfsdk:"enable"`
}

// IkeGatewaysProtocolCommonNatTraversal represents a nested structure within the IkeGateways model
type IkeGatewaysProtocolCommonNatTraversal struct {
	Enable basetypes.BoolValue `tfsdk:"enable"`
}

// AttrTypes defines the attribute types for the IkeGateways model.
func (o IkeGateways) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":             basetypes.StringType{},
		"encrypted_values": basetypes.MapType{ElemType: basetypes.StringType{}},
		"authentication": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"certificate": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"allow_id_payload_mismatch": basetypes.BoolType{},
						"certificate_profile":       basetypes.StringType{},
						"local_certificate": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"local_certificate_name": basetypes.StringType{},
							},
						},
						"strict_validation_revocation": basetypes.BoolType{},
						"use_management_as_source":     basetypes.BoolType{},
					},
				},
				"pre_shared_key": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"key": basetypes.StringType{},
					},
				},
			},
		},
		"device": basetypes.StringType{},
		"folder": basetypes.StringType{},
		"id":     basetypes.StringType{},
		"local_address": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"interface": basetypes.StringType{},
			},
		},
		"local_id": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"id":   basetypes.StringType{},
				"type": basetypes.StringType{},
			},
		},
		"name": basetypes.StringType{},
		"peer_address": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"dynamic": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"fqdn": basetypes.StringType{},
				"ip":   basetypes.StringType{},
			},
		},
		"peer_id": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"id":   basetypes.StringType{},
				"type": basetypes.StringType{},
			},
		},
		"protocol": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ikev1": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"dpd": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable": basetypes.BoolType{},
							},
						},
						"ike_crypto_profile": basetypes.StringType{},
					},
				},
				"ikev2": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"dpd": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable": basetypes.BoolType{},
							},
						},
						"ike_crypto_profile": basetypes.StringType{},
					},
				},
				"version": basetypes.StringType{},
			},
		},
		"protocol_common": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"fragmentation": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.BoolType{},
					},
				},
				"nat_traversal": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.BoolType{},
					},
				},
				"passive_mode": basetypes.BoolType{},
			},
		},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of IkeGateways objects.
func (o IkeGateways) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IkeGatewaysAuthentication model.
func (o IkeGatewaysAuthentication) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"certificate": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"allow_id_payload_mismatch": basetypes.BoolType{},
				"certificate_profile":       basetypes.StringType{},
				"local_certificate": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"local_certificate_name": basetypes.StringType{},
					},
				},
				"strict_validation_revocation": basetypes.BoolType{},
				"use_management_as_source":     basetypes.BoolType{},
			},
		},
		"pre_shared_key": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"key": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of IkeGatewaysAuthentication objects.
func (o IkeGatewaysAuthentication) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IkeGatewaysAuthenticationCertificate model.
func (o IkeGatewaysAuthenticationCertificate) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"allow_id_payload_mismatch": basetypes.BoolType{},
		"certificate_profile":       basetypes.StringType{},
		"local_certificate": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"local_certificate_name": basetypes.StringType{},
			},
		},
		"strict_validation_revocation": basetypes.BoolType{},
		"use_management_as_source":     basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of IkeGatewaysAuthenticationCertificate objects.
func (o IkeGatewaysAuthenticationCertificate) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IkeGatewaysAuthenticationCertificateLocalCertificate model.
func (o IkeGatewaysAuthenticationCertificateLocalCertificate) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"local_certificate_name": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of IkeGatewaysAuthenticationCertificateLocalCertificate objects.
func (o IkeGatewaysAuthenticationCertificateLocalCertificate) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IkeGatewaysAuthenticationPreSharedKey model.
func (o IkeGatewaysAuthenticationPreSharedKey) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"key": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of IkeGatewaysAuthenticationPreSharedKey objects.
func (o IkeGatewaysAuthenticationPreSharedKey) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IkeGatewaysLocalAddress model.
func (o IkeGatewaysLocalAddress) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"interface": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of IkeGatewaysLocalAddress objects.
func (o IkeGatewaysLocalAddress) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IkeGatewaysLocalId model.
func (o IkeGatewaysLocalId) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":   basetypes.StringType{},
		"type": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of IkeGatewaysLocalId objects.
func (o IkeGatewaysLocalId) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IkeGatewaysPeerAddress model.
func (o IkeGatewaysPeerAddress) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"dynamic": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"fqdn": basetypes.StringType{},
		"ip":   basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of IkeGatewaysPeerAddress objects.
func (o IkeGatewaysPeerAddress) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IkeGatewaysPeerId model.
func (o IkeGatewaysPeerId) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":   basetypes.StringType{},
		"type": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of IkeGatewaysPeerId objects.
func (o IkeGatewaysPeerId) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IkeGatewaysProtocol model.
func (o IkeGatewaysProtocol) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ikev1": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"dpd": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.BoolType{},
					},
				},
				"ike_crypto_profile": basetypes.StringType{},
			},
		},
		"ikev2": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"dpd": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.BoolType{},
					},
				},
				"ike_crypto_profile": basetypes.StringType{},
			},
		},
		"version": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of IkeGatewaysProtocol objects.
func (o IkeGatewaysProtocol) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IkeGatewaysProtocolIkev1 model.
func (o IkeGatewaysProtocolIkev1) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"dpd": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable": basetypes.BoolType{},
			},
		},
		"ike_crypto_profile": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of IkeGatewaysProtocolIkev1 objects.
func (o IkeGatewaysProtocolIkev1) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IkeGatewaysProtocolIkev1Dpd model.
func (o IkeGatewaysProtocolIkev1Dpd) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of IkeGatewaysProtocolIkev1Dpd objects.
func (o IkeGatewaysProtocolIkev1Dpd) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IkeGatewaysProtocolCommon model.
func (o IkeGatewaysProtocolCommon) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"fragmentation": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable": basetypes.BoolType{},
			},
		},
		"nat_traversal": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable": basetypes.BoolType{},
			},
		},
		"passive_mode": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of IkeGatewaysProtocolCommon objects.
func (o IkeGatewaysProtocolCommon) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IkeGatewaysProtocolCommonFragmentation model.
func (o IkeGatewaysProtocolCommonFragmentation) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of IkeGatewaysProtocolCommonFragmentation objects.
func (o IkeGatewaysProtocolCommonFragmentation) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IkeGatewaysProtocolCommonNatTraversal model.
func (o IkeGatewaysProtocolCommonNatTraversal) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of IkeGatewaysProtocolCommonNatTraversal objects.
func (o IkeGatewaysProtocolCommonNatTraversal) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// IkeGatewaysResourceSchema defines the schema for IkeGateways resource
var IkeGatewaysResourceSchema = schema.Schema{
	MarkdownDescription: "IkeGateway resource",
	Attributes: map[string]schema.Attribute{
		"authentication": schema.SingleNestedAttribute{
			MarkdownDescription: "Authentication",
			Required:            true,
			Attributes: map[string]schema.Attribute{
				"certificate": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("pre_shared_key"),
						),
					},
					MarkdownDescription: "Certificate",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"allow_id_payload_mismatch": schema.BoolAttribute{
							MarkdownDescription: "Allow id payload mismatch",
							Optional:            true,
						},
						"certificate_profile": schema.StringAttribute{
							MarkdownDescription: "Certificate profile",
							Optional:            true,
						},
						"local_certificate": schema.SingleNestedAttribute{
							MarkdownDescription: "Local certificate",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"local_certificate_name": schema.StringAttribute{
									MarkdownDescription: "Local certificate name",
									Optional:            true,
								},
							},
						},
						"strict_validation_revocation": schema.BoolAttribute{
							MarkdownDescription: "Strict validation revocation",
							Optional:            true,
						},
						"use_management_as_source": schema.BoolAttribute{
							MarkdownDescription: "Use management as source",
							Optional:            true,
						},
					},
				},
				"pre_shared_key": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("certificate"),
						),
					},
					MarkdownDescription: "Pre shared key",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							MarkdownDescription: "Key",
							Optional:            true,
							Sensitive:           true,
						},
					},
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
		"local_address": schema.SingleNestedAttribute{
			MarkdownDescription: "Local address",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"interface": schema.StringAttribute{
					MarkdownDescription: "Interface variable or hardcoded vlan/loopback. vlan will be passed as default value",
					Optional:            true,
					Computed:            true,
					Default:             stringdefault.StaticString("vlan"),
				},
			},
		},
		"local_id": schema.SingleNestedAttribute{
			MarkdownDescription: "Local id",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.LengthAtMost(1024),
						stringvalidator.LengthAtLeast(1),
						stringvalidator.RegexMatches(regexp.MustCompile("^(.+\\@[a-zA-Z0-9.-]+)$|^([$a-zA-Z0-9_:.-]+)$|^(([[:xdigit:]][[:xdigit:]])+)$|^([a-zA-Z0-9.]+=(\\\\,|[^,])+[, ]+)*([a-zA-Z0-9.]+=(\\\\,|[^,])+)$"), "pattern must match "+"^(.+\\@[a-zA-Z0-9.-]+)$|^([$a-zA-Z0-9_:.-]+)$|^(([[:xdigit:]][[:xdigit:]])+)$|^([a-zA-Z0-9.]+=(\\\\,|[^,])+[, ]+)*([a-zA-Z0-9.]+=(\\\\,|[^,])+)$"),
					},
					MarkdownDescription: "Local ID string",
					Optional:            true,
				},
				"type": schema.StringAttribute{
					MarkdownDescription: "Type",
					Optional:            true,
				},
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(63),
			},
			MarkdownDescription: "Alphanumeric string begin with letter: [0-9a-zA-Z._-]",
			Required:            true,
		},
		"peer_address": schema.SingleNestedAttribute{
			MarkdownDescription: "Peer address",
			Required:            true,
			Attributes: map[string]schema.Attribute{
				"dynamic": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("fqdn"),
							path.MatchRelative().AtParent().AtName("ip"),
						),
					},
					MarkdownDescription: "Dynamic",
					Optional:            true,
					Attributes:          map[string]schema.Attribute{},
				},
				"fqdn": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("dynamic"),
							path.MatchRelative().AtParent().AtName("ip"),
						),
						stringvalidator.LengthAtMost(255),
					},
					MarkdownDescription: "peer gateway FQDN name",
					Optional:            true,
				},
				"ip": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("dynamic"),
							path.MatchRelative().AtParent().AtName("fqdn"),
						),
					},
					MarkdownDescription: "peer gateway has static IP address",
					Optional:            true,
				},
			},
		},
		"peer_id": schema.SingleNestedAttribute{
			MarkdownDescription: "Peer id",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.LengthAtMost(1024),
						stringvalidator.LengthAtLeast(1),
						stringvalidator.RegexMatches(regexp.MustCompile("^(.+\\@[\\*a-zA-Z0-9.-]+)$|^([\\*$a-zA-Z0-9_:.-]+)$|^(([[:xdigit:]][[:xdigit:]])+)$|^([a-zA-Z0-9.]+=(\\\\,|[^,])+[, ]+)*([a-zA-Z0-9.]+=(\\\\,|[^,])+)$"), "pattern must match "+"^(.+\\@[\\*a-zA-Z0-9.-]+)$|^([\\*$a-zA-Z0-9_:.-]+)$|^(([[:xdigit:]][[:xdigit:]])+)$|^([a-zA-Z0-9.]+=(\\\\,|[^,])+[, ]+)*([a-zA-Z0-9.]+=(\\\\,|[^,])+)$"),
					},
					MarkdownDescription: "Peer ID string",
					Optional:            true,
				},
				"type": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("ipaddr", "keyid", "fqdn", "ufqdn"),
					},
					MarkdownDescription: "Type",
					Optional:            true,
				},
			},
		},
		"protocol": schema.SingleNestedAttribute{
			MarkdownDescription: "Protocol",
			Required:            true,
			Attributes: map[string]schema.Attribute{
				"ikev1": schema.SingleNestedAttribute{
					MarkdownDescription: "Ikev1",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"dpd": schema.SingleNestedAttribute{
							MarkdownDescription: "Dpd",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"enable": schema.BoolAttribute{
									MarkdownDescription: "Enable",
									Optional:            true,
								},
							},
						},
						"ike_crypto_profile": schema.StringAttribute{
							MarkdownDescription: "Ike crypto profile",
							Optional:            true,
						},
					},
				},
				"ikev2": schema.SingleNestedAttribute{
					MarkdownDescription: "Ikev2",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"dpd": schema.SingleNestedAttribute{
							MarkdownDescription: "Dpd",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"enable": schema.BoolAttribute{
									MarkdownDescription: "Enable",
									Optional:            true,
								},
							},
						},
						"ike_crypto_profile": schema.StringAttribute{
							MarkdownDescription: "Ike crypto profile",
							Optional:            true,
						},
					},
				},
				"version": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("ikev2-preferred", "ikev1", "ikev2"),
					},
					MarkdownDescription: "Version",
					Optional:            true,
					Computed:            true,
					Default:             stringdefault.StaticString("ikev2-preferred"),
				},
			},
		},
		"protocol_common": schema.SingleNestedAttribute{
			MarkdownDescription: "Protocol common",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"fragmentation": schema.SingleNestedAttribute{
					MarkdownDescription: "Fragmentation",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"enable": schema.BoolAttribute{
							MarkdownDescription: "Enable",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
					},
				},
				"nat_traversal": schema.SingleNestedAttribute{
					MarkdownDescription: "Enables NAT traversal for the IKE gateway.",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"enable": schema.BoolAttribute{
							MarkdownDescription: "Enable",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(true),
						},
					},
				},
				"passive_mode": schema.BoolAttribute{
					MarkdownDescription: "Passive mode",
					Optional:            true,
					Computed:            true,
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

// IkeGatewaysDataSourceSchema defines the schema for IkeGateways data source
var IkeGatewaysDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "IkeGateway data source",
	Attributes: map[string]dsschema.Attribute{
		"authentication": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Authentication",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"certificate": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Certificate",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"allow_id_payload_mismatch": dsschema.BoolAttribute{
							MarkdownDescription: "Allow id payload mismatch",
							Computed:            true,
						},
						"certificate_profile": dsschema.StringAttribute{
							MarkdownDescription: "Certificate profile",
							Computed:            true,
						},
						"local_certificate": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Local certificate",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"local_certificate_name": dsschema.StringAttribute{
									MarkdownDescription: "Local certificate name",
									Computed:            true,
								},
							},
						},
						"strict_validation_revocation": dsschema.BoolAttribute{
							MarkdownDescription: "Strict validation revocation",
							Computed:            true,
						},
						"use_management_as_source": dsschema.BoolAttribute{
							MarkdownDescription: "Use management as source",
							Computed:            true,
						},
					},
				},
				"pre_shared_key": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Pre shared key",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"key": dsschema.StringAttribute{
							MarkdownDescription: "Key",
							Computed:            true,
							Sensitive:           true,
						},
					},
				},
			},
		},
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
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"local_address": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Local address",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"interface": dsschema.StringAttribute{
					MarkdownDescription: "Interface variable or hardcoded vlan/loopback. vlan will be passed as default value",
					Computed:            true,
				},
			},
		},
		"local_id": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Local id",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"id": dsschema.StringAttribute{
					MarkdownDescription: "Local ID string",
					Computed:            true,
				},
				"type": dsschema.StringAttribute{
					MarkdownDescription: "Type",
					Computed:            true,
				},
			},
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Alphanumeric string begin with letter: [0-9a-zA-Z._-]",
			Optional:            true,
			Computed:            true,
		},
		"peer_address": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Peer address",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"dynamic": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Dynamic",
					Computed:            true,
					Attributes:          map[string]dsschema.Attribute{},
				},
				"fqdn": dsschema.StringAttribute{
					MarkdownDescription: "peer gateway FQDN name",
					Computed:            true,
				},
				"ip": dsschema.StringAttribute{
					MarkdownDescription: "peer gateway has static IP address",
					Computed:            true,
				},
			},
		},
		"peer_id": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Peer id",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"id": dsschema.StringAttribute{
					MarkdownDescription: "Peer ID string",
					Computed:            true,
				},
				"type": dsschema.StringAttribute{
					MarkdownDescription: "Type",
					Computed:            true,
				},
			},
		},
		"protocol": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Protocol",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"ikev1": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Ikev1",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"dpd": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Dpd",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"enable": dsschema.BoolAttribute{
									MarkdownDescription: "Enable",
									Computed:            true,
								},
							},
						},
						"ike_crypto_profile": dsschema.StringAttribute{
							MarkdownDescription: "Ike crypto profile",
							Computed:            true,
						},
					},
				},
				"ikev2": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Ikev2",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"dpd": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Dpd",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"enable": dsschema.BoolAttribute{
									MarkdownDescription: "Enable",
									Computed:            true,
								},
							},
						},
						"ike_crypto_profile": dsschema.StringAttribute{
							MarkdownDescription: "Ike crypto profile",
							Computed:            true,
						},
					},
				},
				"version": dsschema.StringAttribute{
					MarkdownDescription: "Version",
					Computed:            true,
				},
			},
		},
		"protocol_common": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Protocol common",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"fragmentation": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Fragmentation",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"enable": dsschema.BoolAttribute{
							MarkdownDescription: "Enable",
							Computed:            true,
						},
					},
				},
				"nat_traversal": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Enables NAT traversal for the IKE gateway.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"enable": dsschema.BoolAttribute{
							MarkdownDescription: "Enable",
							Computed:            true,
						},
					},
				},
				"passive_mode": dsschema.BoolAttribute{
					MarkdownDescription: "Passive mode",
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

// IkeGatewaysListModel represents the data model for a list data source.
type IkeGatewaysListModel struct {
	Tfid    types.String  `tfsdk:"tfid"`
	Data    []IkeGateways `tfsdk:"data"`
	Limit   types.Int64   `tfsdk:"limit"`
	Offset  types.Int64   `tfsdk:"offset"`
	Name    types.String  `tfsdk:"name"`
	Total   types.Int64   `tfsdk:"total"`
	Folder  types.String  `tfsdk:"folder"`
	Snippet types.String  `tfsdk:"snippet"`
	Device  types.String  `tfsdk:"device"`
}

// IkeGatewaysListDataSourceSchema defines the schema for a list data source.
var IkeGatewaysListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: IkeGatewaysDataSourceSchema.Attributes,
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
