package models

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: deployment_services
// This file contains models for the deployment_services SDK package

// RemoteNetworks represents the Terraform model for RemoteNetworks
type RemoteNetworks struct {
	Tfid                 types.String          `tfsdk:"tfid"`
	EncryptedValues      basetypes.MapValue    `tfsdk:"encrypted_values"`
	EcmpLoadBalancing    basetypes.StringValue `tfsdk:"ecmp_load_balancing"`
	EcmpTunnels          basetypes.ListValue   `tfsdk:"ecmp_tunnels"`
	Folder               basetypes.StringValue `tfsdk:"folder"`
	Id                   basetypes.StringValue `tfsdk:"id"`
	IpsecTunnel          basetypes.StringValue `tfsdk:"ipsec_tunnel"`
	LicenseType          basetypes.StringValue `tfsdk:"license_type"`
	Name                 basetypes.StringValue `tfsdk:"name"`
	Protocol             basetypes.ObjectValue `tfsdk:"protocol"`
	Region               basetypes.StringValue `tfsdk:"region"`
	SecondaryIpsecTunnel basetypes.StringValue `tfsdk:"secondary_ipsec_tunnel"`
	SpnName              basetypes.StringValue `tfsdk:"spn_name"`
	Subnets              basetypes.ListValue   `tfsdk:"subnets"`
}

// RemoteNetworksEcmpTunnelsInner represents a nested structure within the RemoteNetworks model
type RemoteNetworksEcmpTunnelsInner struct {
	IpsecTunnel basetypes.StringValue `tfsdk:"ipsec_tunnel"`
	Name        basetypes.StringValue `tfsdk:"name"`
	Protocol    basetypes.ObjectValue `tfsdk:"protocol"`
}

// RemoteNetworksEcmpTunnelsInnerProtocol represents a nested structure within the RemoteNetworks model
type RemoteNetworksEcmpTunnelsInnerProtocol struct {
	Bgp basetypes.ObjectValue `tfsdk:"bgp"`
}

// RemoteNetworksProtocolBgp represents a nested structure within the RemoteNetworks model
type RemoteNetworksProtocolBgp struct {
	DoNotExportRoutes         basetypes.BoolValue   `tfsdk:"do_not_export_routes"`
	Enable                    basetypes.BoolValue   `tfsdk:"enable"`
	LocalIpAddress            basetypes.StringValue `tfsdk:"local_ip_address"`
	OriginateDefaultRoute     basetypes.BoolValue   `tfsdk:"originate_default_route"`
	PeerAs                    basetypes.StringValue `tfsdk:"peer_as"`
	PeerIpAddress             basetypes.StringValue `tfsdk:"peer_ip_address"`
	PeeringType               basetypes.StringValue `tfsdk:"peering_type"`
	Secret                    basetypes.StringValue `tfsdk:"secret"`
	SummarizeMobileUserRoutes basetypes.BoolValue   `tfsdk:"summarize_mobile_user_routes"`
}

// RemoteNetworksProtocol represents a nested structure within the RemoteNetworks model
type RemoteNetworksProtocol struct {
	Bgp     basetypes.ObjectValue `tfsdk:"bgp"`
	BgpPeer basetypes.ObjectValue `tfsdk:"bgp_peer"`
}

// RemoteNetworksProtocolBgpPeer represents a nested structure within the RemoteNetworks model
type RemoteNetworksProtocolBgpPeer struct {
	LocalIpAddress basetypes.StringValue `tfsdk:"local_ip_address"`
	PeerIpAddress  basetypes.StringValue `tfsdk:"peer_ip_address"`
	Secret         basetypes.StringValue `tfsdk:"secret"`
}

// AttrTypes defines the attribute types for the RemoteNetworks model.
func (o RemoteNetworks) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":                basetypes.StringType{},
		"encrypted_values":    basetypes.MapType{ElemType: basetypes.StringType{}},
		"ecmp_load_balancing": basetypes.StringType{},
		"ecmp_tunnels": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ipsec_tunnel": basetypes.StringType{},
				"name":         basetypes.StringType{},
				"protocol": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"bgp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"do_not_export_routes":         basetypes.BoolType{},
								"enable":                       basetypes.BoolType{},
								"local_ip_address":             basetypes.StringType{},
								"originate_default_route":      basetypes.BoolType{},
								"peer_as":                      basetypes.StringType{},
								"peer_ip_address":              basetypes.StringType{},
								"peering_type":                 basetypes.StringType{},
								"secret":                       basetypes.StringType{},
								"summarize_mobile_user_routes": basetypes.BoolType{},
							},
						},
					},
				},
			},
		}},
		"folder":       basetypes.StringType{},
		"id":           basetypes.StringType{},
		"ipsec_tunnel": basetypes.StringType{},
		"license_type": basetypes.StringType{},
		"name":         basetypes.StringType{},
		"protocol": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"bgp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"do_not_export_routes":         basetypes.BoolType{},
						"enable":                       basetypes.BoolType{},
						"local_ip_address":             basetypes.StringType{},
						"originate_default_route":      basetypes.BoolType{},
						"peer_as":                      basetypes.StringType{},
						"peer_ip_address":              basetypes.StringType{},
						"peering_type":                 basetypes.StringType{},
						"secret":                       basetypes.StringType{},
						"summarize_mobile_user_routes": basetypes.BoolType{},
					},
				},
				"bgp_peer": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"local_ip_address": basetypes.StringType{},
						"peer_ip_address":  basetypes.StringType{},
						"secret":           basetypes.StringType{},
					},
				},
			},
		},
		"region":                 basetypes.StringType{},
		"secondary_ipsec_tunnel": basetypes.StringType{},
		"spn_name":               basetypes.StringType{},
		"subnets":                basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of RemoteNetworks objects.
func (o RemoteNetworks) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RemoteNetworksEcmpTunnelsInner model.
func (o RemoteNetworksEcmpTunnelsInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ipsec_tunnel": basetypes.StringType{},
		"name":         basetypes.StringType{},
		"protocol": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"bgp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"do_not_export_routes":         basetypes.BoolType{},
						"enable":                       basetypes.BoolType{},
						"local_ip_address":             basetypes.StringType{},
						"originate_default_route":      basetypes.BoolType{},
						"peer_as":                      basetypes.StringType{},
						"peer_ip_address":              basetypes.StringType{},
						"peering_type":                 basetypes.StringType{},
						"secret":                       basetypes.StringType{},
						"summarize_mobile_user_routes": basetypes.BoolType{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of RemoteNetworksEcmpTunnelsInner objects.
func (o RemoteNetworksEcmpTunnelsInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RemoteNetworksEcmpTunnelsInnerProtocol model.
func (o RemoteNetworksEcmpTunnelsInnerProtocol) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"bgp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"do_not_export_routes":         basetypes.BoolType{},
				"enable":                       basetypes.BoolType{},
				"local_ip_address":             basetypes.StringType{},
				"originate_default_route":      basetypes.BoolType{},
				"peer_as":                      basetypes.StringType{},
				"peer_ip_address":              basetypes.StringType{},
				"peering_type":                 basetypes.StringType{},
				"secret":                       basetypes.StringType{},
				"summarize_mobile_user_routes": basetypes.BoolType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of RemoteNetworksEcmpTunnelsInnerProtocol objects.
func (o RemoteNetworksEcmpTunnelsInnerProtocol) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RemoteNetworksProtocolBgp model.
func (o RemoteNetworksProtocolBgp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"do_not_export_routes":         basetypes.BoolType{},
		"enable":                       basetypes.BoolType{},
		"local_ip_address":             basetypes.StringType{},
		"originate_default_route":      basetypes.BoolType{},
		"peer_as":                      basetypes.StringType{},
		"peer_ip_address":              basetypes.StringType{},
		"peering_type":                 basetypes.StringType{},
		"secret":                       basetypes.StringType{},
		"summarize_mobile_user_routes": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of RemoteNetworksProtocolBgp objects.
func (o RemoteNetworksProtocolBgp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RemoteNetworksProtocol model.
func (o RemoteNetworksProtocol) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"bgp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"do_not_export_routes":         basetypes.BoolType{},
				"enable":                       basetypes.BoolType{},
				"local_ip_address":             basetypes.StringType{},
				"originate_default_route":      basetypes.BoolType{},
				"peer_as":                      basetypes.StringType{},
				"peer_ip_address":              basetypes.StringType{},
				"peering_type":                 basetypes.StringType{},
				"secret":                       basetypes.StringType{},
				"summarize_mobile_user_routes": basetypes.BoolType{},
			},
		},
		"bgp_peer": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"local_ip_address": basetypes.StringType{},
				"peer_ip_address":  basetypes.StringType{},
				"secret":           basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of RemoteNetworksProtocol objects.
func (o RemoteNetworksProtocol) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RemoteNetworksProtocolBgpPeer model.
func (o RemoteNetworksProtocolBgpPeer) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"local_ip_address": basetypes.StringType{},
		"peer_ip_address":  basetypes.StringType{},
		"secret":           basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of RemoteNetworksProtocolBgpPeer objects.
func (o RemoteNetworksProtocolBgpPeer) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// RemoteNetworksResourceSchema defines the schema for RemoteNetworks resource
var RemoteNetworksResourceSchema = schema.Schema{
	MarkdownDescription: "RemoteNetwork resource",
	Attributes: map[string]schema.Attribute{
		"ecmp_load_balancing": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("enable", "disable"),
			},
			MarkdownDescription: "Ecmp load balancing",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("disable"),
		},
		"ecmp_tunnels": schema.ListNestedAttribute{
			MarkdownDescription: "ecmp_tunnels is required when ecmp_load_balancing is enable",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"ipsec_tunnel": schema.StringAttribute{
						MarkdownDescription: "Ipsec tunnel",
						Required:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "Name",
						Required:            true,
					},
					"protocol": schema.SingleNestedAttribute{
						MarkdownDescription: "Protocol",
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"bgp": schema.SingleNestedAttribute{
								MarkdownDescription: "Bgp",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"do_not_export_routes": schema.BoolAttribute{
										MarkdownDescription: "Do not export routes?",
										Optional:            true,
									},
									"enable": schema.BoolAttribute{
										MarkdownDescription: "Enable BGP peering?",
										Optional:            true,
									},
									"local_ip_address": schema.StringAttribute{
										MarkdownDescription: "Local peer IP address",
										Optional:            true,
									},
									"originate_default_route": schema.BoolAttribute{
										MarkdownDescription: "Originate default route?",
										Optional:            true,
									},
									"peer_as": schema.StringAttribute{
										MarkdownDescription: "BGP peer ASN",
										Optional:            true,
									},
									"peer_ip_address": schema.StringAttribute{
										MarkdownDescription: "Remote peer IP address",
										Optional:            true,
									},
									"peering_type": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("exchange-v4-over-v4", "exchange-v4-v6-over-v4", "exchange-v4-over-v4-v6-over-v6", "exchange-v6-over-v6"),
										},
										MarkdownDescription: "Route exchange types",
										Optional:            true,
									},
									"secret": schema.StringAttribute{
										MarkdownDescription: "BGP peering secret",
										Optional:            true,
										Sensitive:           true,
									},
									"summarize_mobile_user_routes": schema.BoolAttribute{
										MarkdownDescription: "Summarize mobile user routes?",
										Optional:            true,
									},
								},
							},
						},
					},
				},
			},
		},
		"encrypted_values": schema.MapAttribute{
			ElementType:         basetypes.StringType{},
			MarkdownDescription: "Map of sensitive values returned from the API.",
			Computed:            true,
			Sensitive:           true,
		},
		"folder": schema.StringAttribute{
			MarkdownDescription: "The folder that contains the remote network",
			Required:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the remote network",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"ipsec_tunnel": schema.StringAttribute{
			MarkdownDescription: "ipsec_tunnel is required when ecmp_load_balancing is disable",
			Optional:            true,
		},
		"license_type": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtLeast(1),
			},
			MarkdownDescription: "New customer will only be on aggregate bandwidth licensing",
			Required:            true,
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(63),
			},
			MarkdownDescription: "The name of the remote network",
			Required:            true,
		},
		"protocol": schema.SingleNestedAttribute{
			MarkdownDescription: "setup the protocol when ecmp_load_balancing is disable",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"bgp": schema.SingleNestedAttribute{
					MarkdownDescription: "Bgp",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"do_not_export_routes": schema.BoolAttribute{
							MarkdownDescription: "Do not export routes?",
							Optional:            true,
						},
						"enable": schema.BoolAttribute{
							MarkdownDescription: "Enable BGP peering?",
							Optional:            true,
						},
						"local_ip_address": schema.StringAttribute{
							MarkdownDescription: "Local peer IP address",
							Optional:            true,
						},
						"originate_default_route": schema.BoolAttribute{
							MarkdownDescription: "Originate default route?",
							Optional:            true,
						},
						"peer_as": schema.StringAttribute{
							MarkdownDescription: "BGP peer ASN",
							Optional:            true,
						},
						"peer_ip_address": schema.StringAttribute{
							MarkdownDescription: "Remote peer IP address",
							Optional:            true,
						},
						"peering_type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("exchange-v4-over-v4", "exchange-v4-v6-over-v4", "exchange-v4-over-v4-v6-over-v6", "exchange-v6-over-v6"),
							},
							MarkdownDescription: "Route exchange types",
							Optional:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "BGP peering secret",
							Optional:            true,
							Sensitive:           true,
						},
						"summarize_mobile_user_routes": schema.BoolAttribute{
							MarkdownDescription: "Summarize mobile user routes?",
							Optional:            true,
						},
					},
				},
				"bgp_peer": schema.SingleNestedAttribute{
					MarkdownDescription: "secondary bgp routing as bgp_peer",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"local_ip_address": schema.StringAttribute{
							MarkdownDescription: "Local peer IP address (secondary WAN)",
							Optional:            true,
						},
						"peer_ip_address": schema.StringAttribute{
							MarkdownDescription: "Remote peer IP address (secondary WAN)",
							Optional:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "BGP peering secret (secondary WAN)",
							Optional:            true,
							Sensitive:           true,
						},
					},
				},
			},
		},
		"region": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtLeast(1),
			},
			MarkdownDescription: "Region",
			Required:            true,
		},
		"secondary_ipsec_tunnel": schema.StringAttribute{
			MarkdownDescription: "specify secondary ipsec_tunnel if needed",
			Optional:            true,
		},
		"spn_name": schema.StringAttribute{
			MarkdownDescription: "spn-name is needed when license_type is FWAAS-AGGREGATE",
			Optional:            true,
		},
		"subnets": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Subnets",
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

// RemoteNetworksDataSourceSchema defines the schema for RemoteNetworks data source
var RemoteNetworksDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "RemoteNetwork data source",
	Attributes: map[string]dsschema.Attribute{
		"ecmp_load_balancing": dsschema.StringAttribute{
			MarkdownDescription: "Ecmp load balancing",
			Computed:            true,
		},
		"ecmp_tunnels": dsschema.ListNestedAttribute{
			MarkdownDescription: "ecmp_tunnels is required when ecmp_load_balancing is enable",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"ipsec_tunnel": dsschema.StringAttribute{
						MarkdownDescription: "Ipsec tunnel",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "Name",
						Computed:            true,
					},
					"protocol": dsschema.SingleNestedAttribute{
						MarkdownDescription: "Protocol",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"bgp": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Bgp",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"do_not_export_routes": dsschema.BoolAttribute{
										MarkdownDescription: "Do not export routes?",
										Computed:            true,
									},
									"enable": dsschema.BoolAttribute{
										MarkdownDescription: "Enable BGP peering?",
										Computed:            true,
									},
									"local_ip_address": dsschema.StringAttribute{
										MarkdownDescription: "Local peer IP address",
										Computed:            true,
									},
									"originate_default_route": dsschema.BoolAttribute{
										MarkdownDescription: "Originate default route?",
										Computed:            true,
									},
									"peer_as": dsschema.StringAttribute{
										MarkdownDescription: "BGP peer ASN",
										Computed:            true,
									},
									"peer_ip_address": dsschema.StringAttribute{
										MarkdownDescription: "Remote peer IP address",
										Computed:            true,
									},
									"peering_type": dsschema.StringAttribute{
										MarkdownDescription: "Route exchange types",
										Computed:            true,
									},
									"secret": dsschema.StringAttribute{
										MarkdownDescription: "BGP peering secret",
										Computed:            true,
										Sensitive:           true,
									},
									"summarize_mobile_user_routes": dsschema.BoolAttribute{
										MarkdownDescription: "Summarize mobile user routes?",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
		"encrypted_values": dsschema.MapAttribute{
			ElementType:         basetypes.StringType{},
			MarkdownDescription: "Map of sensitive values returned from the API.",
			Computed:            true,
			Sensitive:           true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder that contains the remote network",
			Optional:            true,
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the remote network",
			Required:            true,
		},
		"ipsec_tunnel": dsschema.StringAttribute{
			MarkdownDescription: "ipsec_tunnel is required when ecmp_load_balancing is disable",
			Computed:            true,
		},
		"license_type": dsschema.StringAttribute{
			MarkdownDescription: "New customer will only be on aggregate bandwidth licensing",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the remote network",
			Optional:            true,
			Computed:            true,
		},
		"protocol": dsschema.SingleNestedAttribute{
			MarkdownDescription: "setup the protocol when ecmp_load_balancing is disable",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"bgp": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Bgp",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"do_not_export_routes": dsschema.BoolAttribute{
							MarkdownDescription: "Do not export routes?",
							Computed:            true,
						},
						"enable": dsschema.BoolAttribute{
							MarkdownDescription: "Enable BGP peering?",
							Computed:            true,
						},
						"local_ip_address": dsschema.StringAttribute{
							MarkdownDescription: "Local peer IP address",
							Computed:            true,
						},
						"originate_default_route": dsschema.BoolAttribute{
							MarkdownDescription: "Originate default route?",
							Computed:            true,
						},
						"peer_as": dsschema.StringAttribute{
							MarkdownDescription: "BGP peer ASN",
							Computed:            true,
						},
						"peer_ip_address": dsschema.StringAttribute{
							MarkdownDescription: "Remote peer IP address",
							Computed:            true,
						},
						"peering_type": dsschema.StringAttribute{
							MarkdownDescription: "Route exchange types",
							Computed:            true,
						},
						"secret": dsschema.StringAttribute{
							MarkdownDescription: "BGP peering secret",
							Computed:            true,
							Sensitive:           true,
						},
						"summarize_mobile_user_routes": dsschema.BoolAttribute{
							MarkdownDescription: "Summarize mobile user routes?",
							Computed:            true,
						},
					},
				},
				"bgp_peer": dsschema.SingleNestedAttribute{
					MarkdownDescription: "secondary bgp routing as bgp_peer",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"local_ip_address": dsschema.StringAttribute{
							MarkdownDescription: "Local peer IP address (secondary WAN)",
							Computed:            true,
						},
						"peer_ip_address": dsschema.StringAttribute{
							MarkdownDescription: "Remote peer IP address (secondary WAN)",
							Computed:            true,
						},
						"secret": dsschema.StringAttribute{
							MarkdownDescription: "BGP peering secret (secondary WAN)",
							Computed:            true,
							Sensitive:           true,
						},
					},
				},
			},
		},
		"region": dsschema.StringAttribute{
			MarkdownDescription: "Region",
			Computed:            true,
		},
		"secondary_ipsec_tunnel": dsschema.StringAttribute{
			MarkdownDescription: "specify secondary ipsec_tunnel if needed",
			Computed:            true,
		},
		"spn_name": dsschema.StringAttribute{
			MarkdownDescription: "spn-name is needed when license_type is FWAAS-AGGREGATE",
			Computed:            true,
		},
		"subnets": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Subnets",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// RemoteNetworksListModel represents the data model for a list data source.
type RemoteNetworksListModel struct {
	Tfid    types.String     `tfsdk:"tfid"`
	Data    []RemoteNetworks `tfsdk:"data"`
	Limit   types.Int64      `tfsdk:"limit"`
	Offset  types.Int64      `tfsdk:"offset"`
	Name    types.String     `tfsdk:"name"`
	Total   types.Int64      `tfsdk:"total"`
	Folder  types.String     `tfsdk:"folder"`
	Snippet types.String     `tfsdk:"snippet"`
	Device  types.String     `tfsdk:"device"`
}

// RemoteNetworksListDataSourceSchema defines the schema for a list data source.
var RemoteNetworksListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: RemoteNetworksDataSourceSchema.Attributes,
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
