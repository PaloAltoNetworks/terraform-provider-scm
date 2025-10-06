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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: network_services
// This file contains models for the network_services SDK package

// IpsecTunnels represents the Terraform model for IpsecTunnels
type IpsecTunnels struct {
	Tfid                   types.String          `tfsdk:"tfid"`
	AntiReplay             basetypes.BoolValue   `tfsdk:"anti_replay"`
	AutoKey                basetypes.ObjectValue `tfsdk:"auto_key"`
	CopyTos                basetypes.BoolValue   `tfsdk:"copy_tos"`
	Device                 basetypes.StringValue `tfsdk:"device"`
	EnableGreEncapsulation basetypes.BoolValue   `tfsdk:"enable_gre_encapsulation"`
	Folder                 basetypes.StringValue `tfsdk:"folder"`
	Id                     basetypes.StringValue `tfsdk:"id"`
	Name                   basetypes.StringValue `tfsdk:"name"`
	Snippet                basetypes.StringValue `tfsdk:"snippet"`
	TunnelInterface        basetypes.StringValue `tfsdk:"tunnel_interface"`
	TunnelMonitor          basetypes.ObjectValue `tfsdk:"tunnel_monitor"`
}

// IpsecTunnelsAutoKey represents a nested structure within the IpsecTunnels model
type IpsecTunnelsAutoKey struct {
	IkeGateway         basetypes.ListValue   `tfsdk:"ike_gateway"`
	IpsecCryptoProfile basetypes.StringValue `tfsdk:"ipsec_crypto_profile"`
	ProxyId            basetypes.ListValue   `tfsdk:"proxy_id"`
	ProxyIdV6          basetypes.ListValue   `tfsdk:"proxy_id_v6"`
}

// IpsecTunnelsAutoKeyIkeGatewayInner represents a nested structure within the IpsecTunnels model
type IpsecTunnelsAutoKeyIkeGatewayInner struct {
	Name basetypes.StringValue `tfsdk:"name"`
}

// IpsecTunnelsAutoKeyProxyIdInner represents a nested structure within the IpsecTunnels model
type IpsecTunnelsAutoKeyProxyIdInner struct {
	Local    basetypes.StringValue `tfsdk:"local"`
	Name     basetypes.StringValue `tfsdk:"name"`
	Protocol basetypes.ObjectValue `tfsdk:"protocol"`
	Remote   basetypes.StringValue `tfsdk:"remote"`
}

// IpsecTunnelsAutoKeyProxyIdInnerProtocol represents a nested structure within the IpsecTunnels model
type IpsecTunnelsAutoKeyProxyIdInnerProtocol struct {
	Number basetypes.Int64Value  `tfsdk:"number"`
	Tcp    basetypes.ObjectValue `tfsdk:"tcp"`
	Udp    basetypes.ObjectValue `tfsdk:"udp"`
}

// IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp represents a nested structure within the IpsecTunnels model
type IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp struct {
	LocalPort  basetypes.Int64Value `tfsdk:"local_port"`
	RemotePort basetypes.Int64Value `tfsdk:"remote_port"`
}

// IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp represents a nested structure within the IpsecTunnels model
type IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp struct {
	LocalPort  basetypes.Int64Value `tfsdk:"local_port"`
	RemotePort basetypes.Int64Value `tfsdk:"remote_port"`
}

// IpsecTunnelsTunnelMonitor represents a nested structure within the IpsecTunnels model
type IpsecTunnelsTunnelMonitor struct {
	DestinationIp basetypes.StringValue `tfsdk:"destination_ip"`
	Enable        basetypes.BoolValue   `tfsdk:"enable"`
	ProxyId       basetypes.StringValue `tfsdk:"proxy_id"`
}

// AttrTypes defines the attribute types for the IpsecTunnels model.
func (o IpsecTunnels) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":        basetypes.StringType{},
		"anti_replay": basetypes.BoolType{},
		"auto_key": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ike_gateway": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name": basetypes.StringType{},
					},
				}},
				"ipsec_crypto_profile": basetypes.StringType{},
				"proxy_id": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"local": basetypes.StringType{},
						"name":  basetypes.StringType{},
						"protocol": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"number": basetypes.Int64Type{},
								"tcp": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"local_port":  basetypes.Int64Type{},
										"remote_port": basetypes.Int64Type{},
									},
								},
								"udp": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"local_port":  basetypes.Int64Type{},
										"remote_port": basetypes.Int64Type{},
									},
								},
							},
						},
						"remote": basetypes.StringType{},
					},
				}},
				"proxy_id_v6": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"local": basetypes.StringType{},
						"name":  basetypes.StringType{},
						"protocol": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"number": basetypes.Int64Type{},
								"tcp": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"local_port":  basetypes.Int64Type{},
										"remote_port": basetypes.Int64Type{},
									},
								},
								"udp": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"local_port":  basetypes.Int64Type{},
										"remote_port": basetypes.Int64Type{},
									},
								},
							},
						},
						"remote": basetypes.StringType{},
					},
				}},
			},
		},
		"copy_tos":                 basetypes.BoolType{},
		"device":                   basetypes.StringType{},
		"enable_gre_encapsulation": basetypes.BoolType{},
		"folder":                   basetypes.StringType{},
		"id":                       basetypes.StringType{},
		"name":                     basetypes.StringType{},
		"snippet":                  basetypes.StringType{},
		"tunnel_interface":         basetypes.StringType{},
		"tunnel_monitor": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"destination_ip": basetypes.StringType{},
				"enable":         basetypes.BoolType{},
				"proxy_id":       basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of IpsecTunnels objects.
func (o IpsecTunnels) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IpsecTunnelsAutoKey model.
func (o IpsecTunnelsAutoKey) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ike_gateway": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name": basetypes.StringType{},
			},
		}},
		"ipsec_crypto_profile": basetypes.StringType{},
		"proxy_id": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"local": basetypes.StringType{},
				"name":  basetypes.StringType{},
				"protocol": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"number": basetypes.Int64Type{},
						"tcp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"local_port":  basetypes.Int64Type{},
								"remote_port": basetypes.Int64Type{},
							},
						},
						"udp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"local_port":  basetypes.Int64Type{},
								"remote_port": basetypes.Int64Type{},
							},
						},
					},
				},
				"remote": basetypes.StringType{},
			},
		}},
		"proxy_id_v6": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"local": basetypes.StringType{},
				"name":  basetypes.StringType{},
				"protocol": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"number": basetypes.Int64Type{},
						"tcp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"local_port":  basetypes.Int64Type{},
								"remote_port": basetypes.Int64Type{},
							},
						},
						"udp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"local_port":  basetypes.Int64Type{},
								"remote_port": basetypes.Int64Type{},
							},
						},
					},
				},
				"remote": basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of IpsecTunnelsAutoKey objects.
func (o IpsecTunnelsAutoKey) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IpsecTunnelsAutoKeyIkeGatewayInner model.
func (o IpsecTunnelsAutoKeyIkeGatewayInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of IpsecTunnelsAutoKeyIkeGatewayInner objects.
func (o IpsecTunnelsAutoKeyIkeGatewayInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IpsecTunnelsAutoKeyProxyIdInner model.
func (o IpsecTunnelsAutoKeyProxyIdInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"local": basetypes.StringType{},
		"name":  basetypes.StringType{},
		"protocol": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"number": basetypes.Int64Type{},
				"tcp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"local_port":  basetypes.Int64Type{},
						"remote_port": basetypes.Int64Type{},
					},
				},
				"udp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"local_port":  basetypes.Int64Type{},
						"remote_port": basetypes.Int64Type{},
					},
				},
			},
		},
		"remote": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of IpsecTunnelsAutoKeyProxyIdInner objects.
func (o IpsecTunnelsAutoKeyProxyIdInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IpsecTunnelsAutoKeyProxyIdInnerProtocol model.
func (o IpsecTunnelsAutoKeyProxyIdInnerProtocol) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"number": basetypes.Int64Type{},
		"tcp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"local_port":  basetypes.Int64Type{},
				"remote_port": basetypes.Int64Type{},
			},
		},
		"udp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"local_port":  basetypes.Int64Type{},
				"remote_port": basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of IpsecTunnelsAutoKeyProxyIdInnerProtocol objects.
func (o IpsecTunnelsAutoKeyProxyIdInnerProtocol) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp model.
func (o IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"local_port":  basetypes.Int64Type{},
		"remote_port": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp objects.
func (o IpsecTunnelsAutoKeyProxyIdInnerProtocolTcp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp model.
func (o IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"local_port":  basetypes.Int64Type{},
		"remote_port": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp objects.
func (o IpsecTunnelsAutoKeyProxyIdInnerProtocolUdp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the IpsecTunnelsTunnelMonitor model.
func (o IpsecTunnelsTunnelMonitor) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"destination_ip": basetypes.StringType{},
		"enable":         basetypes.BoolType{},
		"proxy_id":       basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of IpsecTunnelsTunnelMonitor objects.
func (o IpsecTunnelsTunnelMonitor) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// IpsecTunnelsResourceSchema defines the schema for IpsecTunnels resource
var IpsecTunnelsResourceSchema = schema.Schema{
	MarkdownDescription: "IpsecTunnel resource",
	Attributes: map[string]schema.Attribute{
		"anti_replay": schema.BoolAttribute{
			MarkdownDescription: "Enable Anti-Replay check on this tunnel",
			Optional:            true,
		},
		"auto_key": schema.SingleNestedAttribute{
			MarkdownDescription: "Auto key",
			Required:            true,
			Attributes: map[string]schema.Attribute{
				"ike_gateway": schema.ListNestedAttribute{
					MarkdownDescription: "Ike gateway",
					Required:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								MarkdownDescription: "Name",
								Optional:            true,
							},
						},
					},
				},
				"ipsec_crypto_profile": schema.StringAttribute{
					MarkdownDescription: "Ipsec crypto profile",
					Required:            true,
				},
				"proxy_id": schema.ListNestedAttribute{
					MarkdownDescription: "IPv4 type of proxy_id values",
					Optional:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"local": schema.StringAttribute{
								MarkdownDescription: "Local",
								Optional:            true,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: "Name",
								Required:            true,
							},
							"protocol": schema.SingleNestedAttribute{
								MarkdownDescription: "Protocol",
								Optional:            true,
								Computed:            true,
								Attributes: map[string]schema.Attribute{
									"number": schema.Int64Attribute{
										Validators: []validator.Int64{
											int64validator.ExactlyOneOf(
												path.MatchRelative().AtParent().AtName("tcp"),
												path.MatchRelative().AtParent().AtName("udp"),
											),
											int64validator.Between(1, 254),
										},
										MarkdownDescription: "IP protocol number",
										Optional:            true,
										Computed:            true,
									},
									"tcp": schema.SingleNestedAttribute{
										Validators: []validator.Object{
											objectvalidator.ExactlyOneOf(
												path.MatchRelative().AtParent().AtName("number"),
												path.MatchRelative().AtParent().AtName("udp"),
											),
										},
										MarkdownDescription: "TCP",
										Optional:            true,
										Computed:            true,
										Attributes: map[string]schema.Attribute{
											"local_port": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(0, 65535),
												},
												MarkdownDescription: "Local port",
												Optional:            true,
												Computed:            true,
												Default:             int64default.StaticInt64(0),
											},
											"remote_port": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(0, 65535),
												},
												MarkdownDescription: "Remote port",
												Optional:            true,
												Computed:            true,
												Default:             int64default.StaticInt64(0),
											},
										},
									},
									"udp": schema.SingleNestedAttribute{
										Validators: []validator.Object{
											objectvalidator.ExactlyOneOf(
												path.MatchRelative().AtParent().AtName("number"),
												path.MatchRelative().AtParent().AtName("tcp"),
											),
										},
										MarkdownDescription: "UDP",
										Optional:            true,
										Computed:            true,
										Attributes: map[string]schema.Attribute{
											"local_port": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(0, 65535),
												},
												MarkdownDescription: "Local port",
												Optional:            true,
												Computed:            true,
												Default:             int64default.StaticInt64(0),
											},
											"remote_port": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(0, 65535),
												},
												MarkdownDescription: "Remote port",
												Optional:            true,
												Computed:            true,
												Default:             int64default.StaticInt64(0),
											},
										},
									},
								},
							},
							"remote": schema.StringAttribute{
								MarkdownDescription: "Remote",
								Optional:            true,
							},
						},
					},
				},
				"proxy_id_v6": schema.ListNestedAttribute{
					MarkdownDescription: "IPv6 type of proxy_id values",
					Optional:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"local": schema.StringAttribute{
								MarkdownDescription: "Local",
								Optional:            true,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: "Name",
								Required:            true,
							},
							"protocol": schema.SingleNestedAttribute{
								MarkdownDescription: "Protocol",
								Optional:            true,
								Computed:            true,
								Attributes: map[string]schema.Attribute{
									"number": schema.Int64Attribute{
										Validators: []validator.Int64{
											int64validator.ExactlyOneOf(
												path.MatchRelative().AtParent().AtName("tcp"),
												path.MatchRelative().AtParent().AtName("udp"),
											),
											int64validator.Between(1, 254),
										},
										MarkdownDescription: "IP protocol number",
										Optional:            true,
										Computed:            true,
									},
									"tcp": schema.SingleNestedAttribute{
										Validators: []validator.Object{
											objectvalidator.ExactlyOneOf(
												path.MatchRelative().AtParent().AtName("number"),
												path.MatchRelative().AtParent().AtName("udp"),
											),
										},
										MarkdownDescription: "TCP",
										Optional:            true,
										Computed:            true,
										Attributes: map[string]schema.Attribute{
											"local_port": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(0, 65535),
												},
												MarkdownDescription: "Local port",
												Optional:            true,
												Computed:            true,
												Default:             int64default.StaticInt64(0),
											},
											"remote_port": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(0, 65535),
												},
												MarkdownDescription: "Remote port",
												Optional:            true,
												Computed:            true,
												Default:             int64default.StaticInt64(0),
											},
										},
									},
									"udp": schema.SingleNestedAttribute{
										Validators: []validator.Object{
											objectvalidator.ExactlyOneOf(
												path.MatchRelative().AtParent().AtName("number"),
												path.MatchRelative().AtParent().AtName("tcp"),
											),
										},
										MarkdownDescription: "UDP",
										Optional:            true,
										Computed:            true,
										Attributes: map[string]schema.Attribute{
											"local_port": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(0, 65535),
												},
												MarkdownDescription: "Local port",
												Optional:            true,
												Computed:            true,
												Default:             int64default.StaticInt64(0),
											},
											"remote_port": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(0, 65535),
												},
												MarkdownDescription: "Remote port",
												Optional:            true,
												Computed:            true,
												Default:             int64default.StaticInt64(0),
											},
										},
									},
								},
							},
							"remote": schema.StringAttribute{
								MarkdownDescription: "Remote",
								Optional:            true,
							},
						},
					},
				},
			},
		},
		"copy_tos": schema.BoolAttribute{
			MarkdownDescription: "Copy IP TOS bits from inner packet to IPSec packet (not recommended)",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
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
		"enable_gre_encapsulation": schema.BoolAttribute{
			MarkdownDescription: "allow GRE over IPSec",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
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
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(63),
			},
			MarkdownDescription: "Alphanumeric string begin with letter: [0-9a-zA-Z._-]",
			Required:            true,
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
		"tunnel_interface": schema.StringAttribute{
			MarkdownDescription: "Tunnel interface variable or hardcoded tunnel. Default will be tunnels.",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("tunnel"),
		},
		"tunnel_monitor": schema.SingleNestedAttribute{
			MarkdownDescription: "Tunnel monitor",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"destination_ip": schema.StringAttribute{
					MarkdownDescription: "Destination IP to send ICMP probe",
					Required:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Enable tunnel monitoring on this tunnel",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"proxy_id": schema.StringAttribute{
					MarkdownDescription: "Which proxy-id (or proxy-id-v6) the monitoring traffic will use",
					Optional:            true,
					Computed:            true,
				},
			},
		},
	},
}

// IpsecTunnelsDataSourceSchema defines the schema for IpsecTunnels data source
var IpsecTunnelsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "IpsecTunnel data source",
	Attributes: map[string]dsschema.Attribute{
		"anti_replay": dsschema.BoolAttribute{
			MarkdownDescription: "Enable Anti-Replay check on this tunnel",
			Computed:            true,
		},
		"auto_key": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Auto key",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"ike_gateway": dsschema.ListNestedAttribute{
					MarkdownDescription: "Ike gateway",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"name": dsschema.StringAttribute{
								MarkdownDescription: "Name",
								Computed:            true,
							},
						},
					},
				},
				"ipsec_crypto_profile": dsschema.StringAttribute{
					MarkdownDescription: "Ipsec crypto profile",
					Computed:            true,
				},
				"proxy_id": dsschema.ListNestedAttribute{
					MarkdownDescription: "IPv4 type of proxy_id values",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"local": dsschema.StringAttribute{
								MarkdownDescription: "Local",
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
									"number": dsschema.Int64Attribute{
										MarkdownDescription: "IP protocol number",
										Computed:            true,
									},
									"tcp": dsschema.SingleNestedAttribute{
										MarkdownDescription: "TCP",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"local_port": dsschema.Int64Attribute{
												MarkdownDescription: "Local port",
												Computed:            true,
											},
											"remote_port": dsschema.Int64Attribute{
												MarkdownDescription: "Remote port",
												Computed:            true,
											},
										},
									},
									"udp": dsschema.SingleNestedAttribute{
										MarkdownDescription: "UDP",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"local_port": dsschema.Int64Attribute{
												MarkdownDescription: "Local port",
												Computed:            true,
											},
											"remote_port": dsschema.Int64Attribute{
												MarkdownDescription: "Remote port",
												Computed:            true,
											},
										},
									},
								},
							},
							"remote": dsschema.StringAttribute{
								MarkdownDescription: "Remote",
								Computed:            true,
							},
						},
					},
				},
				"proxy_id_v6": dsschema.ListNestedAttribute{
					MarkdownDescription: "IPv6 type of proxy_id values",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"local": dsschema.StringAttribute{
								MarkdownDescription: "Local",
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
									"number": dsschema.Int64Attribute{
										MarkdownDescription: "IP protocol number",
										Computed:            true,
									},
									"tcp": dsschema.SingleNestedAttribute{
										MarkdownDescription: "TCP",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"local_port": dsschema.Int64Attribute{
												MarkdownDescription: "Local port",
												Computed:            true,
											},
											"remote_port": dsschema.Int64Attribute{
												MarkdownDescription: "Remote port",
												Computed:            true,
											},
										},
									},
									"udp": dsschema.SingleNestedAttribute{
										MarkdownDescription: "UDP",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"local_port": dsschema.Int64Attribute{
												MarkdownDescription: "Local port",
												Computed:            true,
											},
											"remote_port": dsschema.Int64Attribute{
												MarkdownDescription: "Remote port",
												Computed:            true,
											},
										},
									},
								},
							},
							"remote": dsschema.StringAttribute{
								MarkdownDescription: "Remote",
								Computed:            true,
							},
						},
					},
				},
			},
		},
		"copy_tos": dsschema.BoolAttribute{
			MarkdownDescription: "Copy IP TOS bits from inner packet to IPSec packet (not recommended)",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"enable_gre_encapsulation": dsschema.BoolAttribute{
			MarkdownDescription: "allow GRE over IPSec",
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
			MarkdownDescription: "Alphanumeric string begin with letter: [0-9a-zA-Z._-]",
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
		"tunnel_interface": dsschema.StringAttribute{
			MarkdownDescription: "Tunnel interface variable or hardcoded tunnel. Default will be tunnels.",
			Computed:            true,
		},
		"tunnel_monitor": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Tunnel monitor",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"destination_ip": dsschema.StringAttribute{
					MarkdownDescription: "Destination IP to send ICMP probe",
					Computed:            true,
				},
				"enable": dsschema.BoolAttribute{
					MarkdownDescription: "Enable tunnel monitoring on this tunnel",
					Computed:            true,
				},
				"proxy_id": dsschema.StringAttribute{
					MarkdownDescription: "Which proxy-id (or proxy-id-v6) the monitoring traffic will use",
					Computed:            true,
				},
			},
		},
	},
}

// IpsecTunnelsListModel represents the data model for a list data source.
type IpsecTunnelsListModel struct {
	Tfid    types.String   `tfsdk:"tfid"`
	Data    []IpsecTunnels `tfsdk:"data"`
	Limit   types.Int64    `tfsdk:"limit"`
	Offset  types.Int64    `tfsdk:"offset"`
	Name    types.String   `tfsdk:"name"`
	Total   types.Int64    `tfsdk:"total"`
	Folder  types.String   `tfsdk:"folder"`
	Snippet types.String   `tfsdk:"snippet"`
	Device  types.String   `tfsdk:"device"`
}

// IpsecTunnelsListDataSourceSchema defines the schema for a list data source.
var IpsecTunnelsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: IpsecTunnelsDataSourceSchema.Attributes,
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
