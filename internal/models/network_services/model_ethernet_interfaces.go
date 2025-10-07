package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: network_services
// This file contains models for the network_services SDK package

// EthernetInterfaces represents the Terraform model for EthernetInterfaces
type EthernetInterfaces struct {
	Tfid            types.String          `tfsdk:"tfid"`
	EncryptedValues basetypes.MapValue    `tfsdk:"encrypted_values"`
	Comment         basetypes.StringValue `tfsdk:"comment"`
	DefaultValue    basetypes.StringValue `tfsdk:"default_value"`
	Device          basetypes.StringValue `tfsdk:"device"`
	Folder          basetypes.StringValue `tfsdk:"folder"`
	Id              basetypes.StringValue `tfsdk:"id"`
	Layer2          basetypes.ObjectValue `tfsdk:"layer2"`
	Layer3          basetypes.ObjectValue `tfsdk:"layer3"`
	LinkDuplex      basetypes.StringValue `tfsdk:"link_duplex"`
	LinkSpeed       basetypes.StringValue `tfsdk:"link_speed"`
	LinkState       basetypes.StringValue `tfsdk:"link_state"`
	Name            basetypes.StringValue `tfsdk:"name"`
	Poe             basetypes.ObjectValue `tfsdk:"poe"`
	Snippet         basetypes.StringValue `tfsdk:"snippet"`
	Tap             basetypes.ObjectValue `tfsdk:"tap"`
}

// EthernetInterfacesLayer2 represents a nested structure within the EthernetInterfaces model
type EthernetInterfacesLayer2 struct {
	VlanTag basetypes.Int64Value `tfsdk:"vlan_tag"`
}

// EthernetInterfacesLayer3 represents a nested structure within the EthernetInterfaces model
type EthernetInterfacesLayer3 struct {
	Arp                        basetypes.ListValue   `tfsdk:"arp"`
	DdnsConfig                 basetypes.ObjectValue `tfsdk:"ddns_config"`
	DhcpClient                 basetypes.ObjectValue `tfsdk:"dhcp_client"`
	InterfaceManagementProfile basetypes.StringValue `tfsdk:"interface_management_profile"`
	Ip                         basetypes.ListValue   `tfsdk:"ip"`
	Mtu                        basetypes.Int64Value  `tfsdk:"mtu"`
	Pppoe                      basetypes.ObjectValue `tfsdk:"pppoe"`
}

// ArpInner represents a nested structure within the EthernetInterfaces model
type ArpInner struct {
	HwAddress basetypes.StringValue `tfsdk:"hw_address"`
	Name      basetypes.StringValue `tfsdk:"name"`
}

// EthernetInterfacesLayer3DdnsConfig represents a nested structure within the EthernetInterfaces model
type EthernetInterfacesLayer3DdnsConfig struct {
	DdnsCertProfile    basetypes.StringValue `tfsdk:"ddns_cert_profile"`
	DdnsEnabled        basetypes.BoolValue   `tfsdk:"ddns_enabled"`
	DdnsHostname       basetypes.StringValue `tfsdk:"ddns_hostname"`
	DdnsIp             basetypes.StringValue `tfsdk:"ddns_ip"`
	DdnsUpdateInterval basetypes.Int64Value  `tfsdk:"ddns_update_interval"`
	DdnsVendor         basetypes.StringValue `tfsdk:"ddns_vendor"`
	DdnsVendorConfig   basetypes.StringValue `tfsdk:"ddns_vendor_config"`
}

// DhcpClient represents a nested structure within the EthernetInterfaces model
type DhcpClient struct {
	DhcpClient basetypes.ObjectValue `tfsdk:"dhcp_client"`
}

// AggregateEthernetInterfacesLayer3DhcpClient represents a nested structure within the EthernetInterfaces model
type AggregateEthernetInterfacesLayer3DhcpClient struct {
	CreateDefaultRoute basetypes.BoolValue   `tfsdk:"create_default_route"`
	DefaultRouteMetric basetypes.Int64Value  `tfsdk:"default_route_metric"`
	Enable             basetypes.BoolValue   `tfsdk:"enable"`
	SendHostname       basetypes.ObjectValue `tfsdk:"send_hostname"`
}

// AggregateEthernetInterfacesLayer3DhcpClientSendHostname represents a nested structure within the EthernetInterfaces model
type AggregateEthernetInterfacesLayer3DhcpClientSendHostname struct {
	Enable   basetypes.BoolValue   `tfsdk:"enable"`
	Hostname basetypes.StringValue `tfsdk:"hostname"`
}

// EthernetInterfacesLayer3Pppoe represents a nested structure within the EthernetInterfaces model
type EthernetInterfacesLayer3Pppoe struct {
	AccessConcentrator basetypes.StringValue `tfsdk:"access_concentrator"`
	Authentication     basetypes.StringValue `tfsdk:"authentication"`
	DefaultRouteMetric basetypes.Int64Value  `tfsdk:"default_route_metric"`
	Enable             basetypes.BoolValue   `tfsdk:"enable"`
	Passive            basetypes.BoolValue   `tfsdk:"passive"`
	Password           basetypes.StringValue `tfsdk:"password"`
	Service            basetypes.StringValue `tfsdk:"service"`
	StaticAddress      basetypes.ObjectValue `tfsdk:"static_address"`
	Username           basetypes.StringValue `tfsdk:"username"`
}

// EthernetInterfacesLayer3PppoeStaticAddress represents a nested structure within the EthernetInterfaces model
type EthernetInterfacesLayer3PppoeStaticAddress struct {
	Ip basetypes.StringValue `tfsdk:"ip"`
}

// Poe represents a nested structure within the EthernetInterfaces model
type Poe struct {
	PoeEnabled basetypes.BoolValue  `tfsdk:"poe_enabled"`
	PoeRsvdPwr basetypes.Int64Value `tfsdk:"poe_rsvd_pwr"`
}

// AttrTypes defines the attribute types for the EthernetInterfaces model.
func (o EthernetInterfaces) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":             basetypes.StringType{},
		"encrypted_values": basetypes.MapType{ElemType: basetypes.StringType{}},
		"comment":          basetypes.StringType{},
		"default_value":    basetypes.StringType{},
		"device":           basetypes.StringType{},
		"folder":           basetypes.StringType{},
		"id":               basetypes.StringType{},
		"layer2": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"vlan_tag": basetypes.Int64Type{},
			},
		},
		"layer3": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"arp": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"hw_address": basetypes.StringType{},
						"name":       basetypes.StringType{},
					},
				}},
				"ddns_config": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ddns_cert_profile":    basetypes.StringType{},
						"ddns_enabled":         basetypes.BoolType{},
						"ddns_hostname":        basetypes.StringType{},
						"ddns_ip":              basetypes.StringType{},
						"ddns_update_interval": basetypes.Int64Type{},
						"ddns_vendor":          basetypes.StringType{},
						"ddns_vendor_config":   basetypes.StringType{},
					},
				},
				"dhcp_client": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"dhcp_client": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"create_default_route": basetypes.BoolType{},
								"default_route_metric": basetypes.Int64Type{},
								"enable":               basetypes.BoolType{},
								"send_hostname": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"enable":   basetypes.BoolType{},
										"hostname": basetypes.StringType{},
									},
								},
							},
						},
					},
				},
				"interface_management_profile": basetypes.StringType{},
				"ip":                           basetypes.ListType{ElemType: basetypes.StringType{}},
				"mtu":                          basetypes.Int64Type{},
				"pppoe": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"access_concentrator":  basetypes.StringType{},
						"authentication":       basetypes.StringType{},
						"default_route_metric": basetypes.Int64Type{},
						"enable":               basetypes.BoolType{},
						"passive":              basetypes.BoolType{},
						"password":             basetypes.StringType{},
						"service":              basetypes.StringType{},
						"static_address": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"ip": basetypes.StringType{},
							},
						},
						"username": basetypes.StringType{},
					},
				},
			},
		},
		"link_duplex": basetypes.StringType{},
		"link_speed":  basetypes.StringType{},
		"link_state":  basetypes.StringType{},
		"name":        basetypes.StringType{},
		"poe": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"poe_enabled":  basetypes.BoolType{},
				"poe_rsvd_pwr": basetypes.Int64Type{},
			},
		},
		"snippet": basetypes.StringType{},
		"tap": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of EthernetInterfaces objects.
func (o EthernetInterfaces) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the EthernetInterfacesLayer2 model.
func (o EthernetInterfacesLayer2) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"vlan_tag": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of EthernetInterfacesLayer2 objects.
func (o EthernetInterfacesLayer2) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the EthernetInterfacesLayer3 model.
func (o EthernetInterfacesLayer3) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"arp": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"hw_address": basetypes.StringType{},
				"name":       basetypes.StringType{},
			},
		}},
		"ddns_config": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ddns_cert_profile":    basetypes.StringType{},
				"ddns_enabled":         basetypes.BoolType{},
				"ddns_hostname":        basetypes.StringType{},
				"ddns_ip":              basetypes.StringType{},
				"ddns_update_interval": basetypes.Int64Type{},
				"ddns_vendor":          basetypes.StringType{},
				"ddns_vendor_config":   basetypes.StringType{},
			},
		},
		"dhcp_client": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"dhcp_client": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"create_default_route": basetypes.BoolType{},
						"default_route_metric": basetypes.Int64Type{},
						"enable":               basetypes.BoolType{},
						"send_hostname": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable":   basetypes.BoolType{},
								"hostname": basetypes.StringType{},
							},
						},
					},
				},
			},
		},
		"interface_management_profile": basetypes.StringType{},
		"ip":                           basetypes.ListType{ElemType: basetypes.StringType{}},
		"mtu":                          basetypes.Int64Type{},
		"pppoe": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"access_concentrator":  basetypes.StringType{},
				"authentication":       basetypes.StringType{},
				"default_route_metric": basetypes.Int64Type{},
				"enable":               basetypes.BoolType{},
				"passive":              basetypes.BoolType{},
				"password":             basetypes.StringType{},
				"service":              basetypes.StringType{},
				"static_address": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ip": basetypes.StringType{},
					},
				},
				"username": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of EthernetInterfacesLayer3 objects.
func (o EthernetInterfacesLayer3) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ArpInner model.
func (o ArpInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"hw_address": basetypes.StringType{},
		"name":       basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ArpInner objects.
func (o ArpInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the EthernetInterfacesLayer3DdnsConfig model.
func (o EthernetInterfacesLayer3DdnsConfig) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ddns_cert_profile":    basetypes.StringType{},
		"ddns_enabled":         basetypes.BoolType{},
		"ddns_hostname":        basetypes.StringType{},
		"ddns_ip":              basetypes.StringType{},
		"ddns_update_interval": basetypes.Int64Type{},
		"ddns_vendor":          basetypes.StringType{},
		"ddns_vendor_config":   basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of EthernetInterfacesLayer3DdnsConfig objects.
func (o EthernetInterfacesLayer3DdnsConfig) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DhcpClient model.
func (o DhcpClient) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"dhcp_client": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"create_default_route": basetypes.BoolType{},
				"default_route_metric": basetypes.Int64Type{},
				"enable":               basetypes.BoolType{},
				"send_hostname": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable":   basetypes.BoolType{},
						"hostname": basetypes.StringType{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of DhcpClient objects.
func (o DhcpClient) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AggregateEthernetInterfacesLayer3DhcpClient model.
func (o AggregateEthernetInterfacesLayer3DhcpClient) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"create_default_route": basetypes.BoolType{},
		"default_route_metric": basetypes.Int64Type{},
		"enable":               basetypes.BoolType{},
		"send_hostname": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable":   basetypes.BoolType{},
				"hostname": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of AggregateEthernetInterfacesLayer3DhcpClient objects.
func (o AggregateEthernetInterfacesLayer3DhcpClient) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AggregateEthernetInterfacesLayer3DhcpClientSendHostname model.
func (o AggregateEthernetInterfacesLayer3DhcpClientSendHostname) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable":   basetypes.BoolType{},
		"hostname": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AggregateEthernetInterfacesLayer3DhcpClientSendHostname objects.
func (o AggregateEthernetInterfacesLayer3DhcpClientSendHostname) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the EthernetInterfacesLayer3Pppoe model.
func (o EthernetInterfacesLayer3Pppoe) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"access_concentrator":  basetypes.StringType{},
		"authentication":       basetypes.StringType{},
		"default_route_metric": basetypes.Int64Type{},
		"enable":               basetypes.BoolType{},
		"passive":              basetypes.BoolType{},
		"password":             basetypes.StringType{},
		"service":              basetypes.StringType{},
		"static_address": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ip": basetypes.StringType{},
			},
		},
		"username": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of EthernetInterfacesLayer3Pppoe objects.
func (o EthernetInterfacesLayer3Pppoe) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the EthernetInterfacesLayer3PppoeStaticAddress model.
func (o EthernetInterfacesLayer3PppoeStaticAddress) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ip": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of EthernetInterfacesLayer3PppoeStaticAddress objects.
func (o EthernetInterfacesLayer3PppoeStaticAddress) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the Poe model.
func (o Poe) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"poe_enabled":  basetypes.BoolType{},
		"poe_rsvd_pwr": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of Poe objects.
func (o Poe) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// EthernetInterfacesResourceSchema defines the schema for EthernetInterfaces resource
var EthernetInterfacesResourceSchema = schema.Schema{
	MarkdownDescription: "EthernetInterface resource",
	Attributes: map[string]schema.Attribute{
		"comment": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(1023),
			},
			MarkdownDescription: "Interface description",
			Optional:            true,
		},
		"default_value": schema.StringAttribute{
			MarkdownDescription: "Default value",
			Optional:            true,
		},
		"device": schema.StringAttribute{
			Validators: []validator.String{
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
		"layer2": schema.SingleNestedAttribute{
			MarkdownDescription: "Layer2",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"vlan_tag": schema.Int64Attribute{
					MarkdownDescription: "Vlan tag",
					Optional:            true,
				},
			},
		},
		"layer3": schema.SingleNestedAttribute{
			MarkdownDescription: "Layer3",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"arp": schema.ListNestedAttribute{
					MarkdownDescription: "ARP configuration",
					Optional:            true,
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"hw_address": schema.StringAttribute{
								MarkdownDescription: "Hw address",
								Optional:            true,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: "IP address",
								Optional:            true,
							},
						},
					},
				},
				"ddns_config": schema.SingleNestedAttribute{
					MarkdownDescription: "Ddns config",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"ddns_cert_profile": schema.StringAttribute{
							MarkdownDescription: "Ddns cert profile",
							Optional:            true,
							Computed:            true,
						},
						"ddns_enabled": schema.BoolAttribute{
							MarkdownDescription: "Ddns enabled",
							Optional:            true,
							Computed:            true,
						},
						"ddns_hostname": schema.StringAttribute{
							MarkdownDescription: "Ddns hostname",
							Optional:            true,
							Computed:            true,
						},
						"ddns_ip": schema.StringAttribute{
							MarkdownDescription: "Ddns ip",
							Optional:            true,
							Computed:            true,
						},
						"ddns_update_interval": schema.Int64Attribute{
							MarkdownDescription: "Ddns update interval",
							Optional:            true,
							Computed:            true,
						},
						"ddns_vendor": schema.StringAttribute{
							MarkdownDescription: "Ddns vendor",
							Optional:            true,
							Computed:            true,
						},
						"ddns_vendor_config": schema.StringAttribute{
							MarkdownDescription: "Ddns vendor config",
							Optional:            true,
							Computed:            true,
						},
					},
				},
				"dhcp_client": schema.SingleNestedAttribute{
					MarkdownDescription: "Dhcp client",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"dhcp_client": schema.SingleNestedAttribute{
							MarkdownDescription: "Dhcp client",
							Optional:            true,
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"create_default_route": schema.BoolAttribute{
									MarkdownDescription: "Create default route",
									Optional:            true,
									Computed:            true,
								},
								"default_route_metric": schema.Int64Attribute{
									MarkdownDescription: "Default route metric",
									Optional:            true,
									Computed:            true,
								},
								"enable": schema.BoolAttribute{
									MarkdownDescription: "Enable",
									Optional:            true,
									Computed:            true,
								},
								"send_hostname": schema.SingleNestedAttribute{
									MarkdownDescription: "Send hostname",
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{
										"enable": schema.BoolAttribute{
											MarkdownDescription: "Enable",
											Optional:            true,
											Computed:            true,
										},
										"hostname": schema.StringAttribute{
											MarkdownDescription: "Hostname",
											Optional:            true,
											Computed:            true,
										},
									},
								},
							},
						},
					},
				},
				"interface_management_profile": schema.StringAttribute{
					MarkdownDescription: "Interface management profile",
					Optional:            true,
					Computed:            true,
				},
				"ip": schema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Interface IP addresses",
					Validators: []validator.List{
						listvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("dhcp_client"),
							path.MatchRelative().AtParent().AtName("pppoe"),
						),
					},
					Optional: true,
					Computed: true,
				},
				"mtu": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(576, 9216),
					},
					MarkdownDescription: "MTU",
					Optional:            true,
					Computed:            true,
					Default:             int64default.StaticInt64(1500),
				},
				"pppoe": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("dhcp_client"),
							path.MatchRelative().AtParent().AtName("ip"),
						),
					},
					MarkdownDescription: "Pppoe",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"access_concentrator": schema.StringAttribute{
							MarkdownDescription: "Access concentrator",
							Optional:            true,
							Computed:            true,
						},
						"authentication": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("CHAP", "PAP", "auto"),
							},
							MarkdownDescription: "Authentication protocol",
							Optional:            true,
							Computed:            true,
						},
						"default_route_metric": schema.Int64Attribute{
							MarkdownDescription: "Default route metric",
							Optional:            true,
							Computed:            true,
						},
						"enable": schema.BoolAttribute{
							MarkdownDescription: "Enable",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(true),
						},
						"passive": schema.BoolAttribute{
							MarkdownDescription: "Passive",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"password": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(255),
							},
							MarkdownDescription: "Password",
							Required:            true,
							Sensitive:           true,
						},
						"service": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(255),
								stringvalidator.LengthAtLeast(1),
							},
							MarkdownDescription: "Service",
							Optional:            true,
							Computed:            true,
						},
						"static_address": schema.SingleNestedAttribute{
							MarkdownDescription: "Static address",
							Optional:            true,
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"ip": schema.StringAttribute{
									MarkdownDescription: "Ip",
									Optional:            true,
									Computed:            true,
								},
							},
						},
						"username": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(255),
								stringvalidator.LengthAtLeast(1),
							},
							MarkdownDescription: "Username",
							Required:            true,
						},
					},
				},
			},
		},
		"link_duplex": schema.StringAttribute{
			MarkdownDescription: "Link duplex",
			Optional:            true,
		},
		"link_speed": schema.StringAttribute{
			MarkdownDescription: "Link speed",
			Optional:            true,
		},
		"link_state": schema.StringAttribute{
			MarkdownDescription: "Link state",
			Optional:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "Interface name",
			Required:            true,
		},
		"poe": schema.SingleNestedAttribute{
			MarkdownDescription: "Poe",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"poe_enabled": schema.BoolAttribute{
					MarkdownDescription: "Poe enabled",
					Optional:            true,
				},
				"poe_rsvd_pwr": schema.Int64Attribute{
					MarkdownDescription: "Poe rsvd pwr",
					Optional:            true,
				},
			},
		},
		"snippet": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
			},
			MarkdownDescription: "The snippet in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"tap": schema.SingleNestedAttribute{
			MarkdownDescription: "Tap",
			Optional:            true,
			Computed:            true,
			Attributes:          map[string]schema.Attribute{},
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

// EthernetInterfacesDataSourceSchema defines the schema for EthernetInterfaces data source
var EthernetInterfacesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "EthernetInterface data source",
	Attributes: map[string]dsschema.Attribute{
		"comment": dsschema.StringAttribute{
			MarkdownDescription: "Interface description",
			Computed:            true,
		},
		"default_value": dsschema.StringAttribute{
			MarkdownDescription: "Default value",
			Computed:            true,
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
		"layer2": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Layer2",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"vlan_tag": dsschema.Int64Attribute{
					MarkdownDescription: "Vlan tag",
					Computed:            true,
				},
			},
		},
		"layer3": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Layer3",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"arp": dsschema.ListNestedAttribute{
					MarkdownDescription: "ARP configuration",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"hw_address": dsschema.StringAttribute{
								MarkdownDescription: "Hw address",
								Computed:            true,
							},
							"name": dsschema.StringAttribute{
								MarkdownDescription: "IP address",
								Computed:            true,
							},
						},
					},
				},
				"ddns_config": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Ddns config",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"ddns_cert_profile": dsschema.StringAttribute{
							MarkdownDescription: "Ddns cert profile",
							Computed:            true,
						},
						"ddns_enabled": dsschema.BoolAttribute{
							MarkdownDescription: "Ddns enabled",
							Computed:            true,
						},
						"ddns_hostname": dsschema.StringAttribute{
							MarkdownDescription: "Ddns hostname",
							Computed:            true,
						},
						"ddns_ip": dsschema.StringAttribute{
							MarkdownDescription: "Ddns ip",
							Computed:            true,
						},
						"ddns_update_interval": dsschema.Int64Attribute{
							MarkdownDescription: "Ddns update interval",
							Computed:            true,
						},
						"ddns_vendor": dsschema.StringAttribute{
							MarkdownDescription: "Ddns vendor",
							Computed:            true,
						},
						"ddns_vendor_config": dsschema.StringAttribute{
							MarkdownDescription: "Ddns vendor config",
							Computed:            true,
						},
					},
				},
				"dhcp_client": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Dhcp client",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"dhcp_client": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Dhcp client",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"create_default_route": dsschema.BoolAttribute{
									MarkdownDescription: "Create default route",
									Computed:            true,
								},
								"default_route_metric": dsschema.Int64Attribute{
									MarkdownDescription: "Default route metric",
									Computed:            true,
								},
								"enable": dsschema.BoolAttribute{
									MarkdownDescription: "Enable",
									Computed:            true,
								},
								"send_hostname": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Send hostname",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"enable": dsschema.BoolAttribute{
											MarkdownDescription: "Enable",
											Computed:            true,
										},
										"hostname": dsschema.StringAttribute{
											MarkdownDescription: "Hostname",
											Computed:            true,
										},
									},
								},
							},
						},
					},
				},
				"interface_management_profile": dsschema.StringAttribute{
					MarkdownDescription: "Interface management profile",
					Computed:            true,
				},
				"ip": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Interface IP addresses",
					Computed:            true,
				},
				"mtu": dsschema.Int64Attribute{
					MarkdownDescription: "MTU",
					Computed:            true,
				},
				"pppoe": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Pppoe",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"access_concentrator": dsschema.StringAttribute{
							MarkdownDescription: "Access concentrator",
							Computed:            true,
						},
						"authentication": dsschema.StringAttribute{
							MarkdownDescription: "Authentication protocol",
							Computed:            true,
						},
						"default_route_metric": dsschema.Int64Attribute{
							MarkdownDescription: "Default route metric",
							Computed:            true,
						},
						"enable": dsschema.BoolAttribute{
							MarkdownDescription: "Enable",
							Computed:            true,
						},
						"passive": dsschema.BoolAttribute{
							MarkdownDescription: "Passive",
							Computed:            true,
						},
						"password": dsschema.StringAttribute{
							MarkdownDescription: "Password",
							Computed:            true,
							Sensitive:           true,
						},
						"service": dsschema.StringAttribute{
							MarkdownDescription: "Service",
							Computed:            true,
						},
						"static_address": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Static address",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"ip": dsschema.StringAttribute{
									MarkdownDescription: "Ip",
									Computed:            true,
								},
							},
						},
						"username": dsschema.StringAttribute{
							MarkdownDescription: "Username",
							Computed:            true,
						},
					},
				},
			},
		},
		"link_duplex": dsschema.StringAttribute{
			MarkdownDescription: "Link duplex",
			Computed:            true,
		},
		"link_speed": dsschema.StringAttribute{
			MarkdownDescription: "Link speed",
			Computed:            true,
		},
		"link_state": dsschema.StringAttribute{
			MarkdownDescription: "Link state",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Interface name",
			Optional:            true,
			Computed:            true,
		},
		"poe": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Poe",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"poe_enabled": dsschema.BoolAttribute{
					MarkdownDescription: "Poe enabled",
					Computed:            true,
				},
				"poe_rsvd_pwr": dsschema.Int64Attribute{
					MarkdownDescription: "Poe rsvd pwr",
					Computed:            true,
				},
			},
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"tap": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Tap",
			Computed:            true,
			Attributes:          map[string]dsschema.Attribute{},
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// EthernetInterfacesListModel represents the data model for a list data source.
type EthernetInterfacesListModel struct {
	Tfid    types.String         `tfsdk:"tfid"`
	Data    []EthernetInterfaces `tfsdk:"data"`
	Limit   types.Int64          `tfsdk:"limit"`
	Offset  types.Int64          `tfsdk:"offset"`
	Name    types.String         `tfsdk:"name"`
	Total   types.Int64          `tfsdk:"total"`
	Folder  types.String         `tfsdk:"folder"`
	Snippet types.String         `tfsdk:"snippet"`
	Device  types.String         `tfsdk:"device"`
}

// EthernetInterfacesListDataSourceSchema defines the schema for a list data source.
var EthernetInterfacesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: EthernetInterfacesDataSourceSchema.Attributes,
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
