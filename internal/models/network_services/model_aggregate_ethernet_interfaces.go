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

// AggregateEthernetInterfaces represents the Terraform model for AggregateEthernetInterfaces
type AggregateEthernetInterfaces struct {
	Tfid         types.String          `tfsdk:"tfid"`
	Comment      basetypes.StringValue `tfsdk:"comment"`
	DefaultValue basetypes.StringValue `tfsdk:"default_value"`
	Device       basetypes.StringValue `tfsdk:"device"`
	Folder       basetypes.StringValue `tfsdk:"folder"`
	Id           basetypes.StringValue `tfsdk:"id"`
	Layer2       basetypes.ObjectValue `tfsdk:"layer2"`
	Layer3       basetypes.ObjectValue `tfsdk:"layer3"`
	Name         basetypes.StringValue `tfsdk:"name"`
	Snippet      basetypes.StringValue `tfsdk:"snippet"`
}

// AggregateEthernetInterfacesLayer2 represents a nested structure within the AggregateEthernetInterfaces model
type AggregateEthernetInterfacesLayer2 struct {
	Lacp    basetypes.ObjectValue `tfsdk:"lacp"`
	VlanTag basetypes.Int64Value  `tfsdk:"vlan_tag"`
}

// Lacp represents a nested structure within the AggregateEthernetInterfaces model
type Lacp struct {
	Enable           basetypes.BoolValue   `tfsdk:"enable"`
	FastFailover     basetypes.BoolValue   `tfsdk:"fast_failover"`
	MaxPorts         basetypes.Int64Value  `tfsdk:"max_ports"`
	Mode             basetypes.StringValue `tfsdk:"mode"`
	SystemPriority   basetypes.Int64Value  `tfsdk:"system_priority"`
	TransmissionRate basetypes.StringValue `tfsdk:"transmission_rate"`
}

// AggregateEthernetInterfacesLayer3 represents a nested structure within the AggregateEthernetInterfaces model
type AggregateEthernetInterfacesLayer3 struct {
	Arp                        basetypes.ListValue   `tfsdk:"arp"`
	DdnsConfig                 basetypes.ObjectValue `tfsdk:"ddns_config"`
	DhcpClient                 basetypes.ObjectValue `tfsdk:"dhcp_client"`
	InterfaceManagementProfile basetypes.StringValue `tfsdk:"interface_management_profile"`
	Ip                         basetypes.ListValue   `tfsdk:"ip"`
	Lacp                       basetypes.ObjectValue `tfsdk:"lacp"`
	Mtu                        basetypes.Int64Value  `tfsdk:"mtu"`
}

// AggEthernetArpInner represents a nested structure within the AggregateEthernetInterfaces model
type AggEthernetArpInner struct {
	HwAddress basetypes.StringValue `tfsdk:"hw_address"`
	Name      basetypes.StringValue `tfsdk:"name"`
}

// AggregateEthernetInterfacesLayer3DdnsConfig represents a nested structure within the AggregateEthernetInterfaces model
type AggregateEthernetInterfacesLayer3DdnsConfig struct {
	DdnsCertProfile    basetypes.StringValue `tfsdk:"ddns_cert_profile"`
	DdnsEnabled        basetypes.BoolValue   `tfsdk:"ddns_enabled"`
	DdnsHostname       basetypes.StringValue `tfsdk:"ddns_hostname"`
	DdnsIp             basetypes.StringValue `tfsdk:"ddns_ip"`
	DdnsUpdateInterval basetypes.Int64Value  `tfsdk:"ddns_update_interval"`
	DdnsVendor         basetypes.StringValue `tfsdk:"ddns_vendor"`
	DdnsVendorConfig   basetypes.StringValue `tfsdk:"ddns_vendor_config"`
}

// AggEthernetDhcpClientDhcpClient represents a nested structure within the AggregateEthernetInterfaces model
type AggEthernetDhcpClientDhcpClient struct {
	CreateDefaultRoute basetypes.BoolValue   `tfsdk:"create_default_route"`
	DefaultRouteMetric basetypes.Int64Value  `tfsdk:"default_route_metric"`
	Enable             basetypes.BoolValue   `tfsdk:"enable"`
	SendHostname       basetypes.ObjectValue `tfsdk:"send_hostname"`
}

// AggEthernetDhcpClientDhcpClientSendHostname represents a nested structure within the AggregateEthernetInterfaces model
type AggEthernetDhcpClientDhcpClientSendHostname struct {
	Enable   basetypes.BoolValue   `tfsdk:"enable"`
	Hostname basetypes.StringValue `tfsdk:"hostname"`
}

// AttrTypes defines the attribute types for the AggregateEthernetInterfaces model.
func (o AggregateEthernetInterfaces) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":          basetypes.StringType{},
		"comment":       basetypes.StringType{},
		"default_value": basetypes.StringType{},
		"device":        basetypes.StringType{},
		"folder":        basetypes.StringType{},
		"id":            basetypes.StringType{},
		"layer2": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"lacp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable":            basetypes.BoolType{},
						"fast_failover":     basetypes.BoolType{},
						"max_ports":         basetypes.Int64Type{},
						"mode":              basetypes.StringType{},
						"system_priority":   basetypes.Int64Type{},
						"transmission_rate": basetypes.StringType{},
					},
				},
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
				"interface_management_profile": basetypes.StringType{},
				"ip":                           basetypes.ListType{ElemType: basetypes.StringType{}},
				"lacp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable":            basetypes.BoolType{},
						"fast_failover":     basetypes.BoolType{},
						"max_ports":         basetypes.Int64Type{},
						"mode":              basetypes.StringType{},
						"system_priority":   basetypes.Int64Type{},
						"transmission_rate": basetypes.StringType{},
					},
				},
				"mtu": basetypes.Int64Type{},
			},
		},
		"name":    basetypes.StringType{},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AggregateEthernetInterfaces objects.
func (o AggregateEthernetInterfaces) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AggregateEthernetInterfacesLayer2 model.
func (o AggregateEthernetInterfacesLayer2) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"lacp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable":            basetypes.BoolType{},
				"fast_failover":     basetypes.BoolType{},
				"max_ports":         basetypes.Int64Type{},
				"mode":              basetypes.StringType{},
				"system_priority":   basetypes.Int64Type{},
				"transmission_rate": basetypes.StringType{},
			},
		},
		"vlan_tag": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of AggregateEthernetInterfacesLayer2 objects.
func (o AggregateEthernetInterfacesLayer2) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the Lacp model.
func (o Lacp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable":            basetypes.BoolType{},
		"fast_failover":     basetypes.BoolType{},
		"max_ports":         basetypes.Int64Type{},
		"mode":              basetypes.StringType{},
		"system_priority":   basetypes.Int64Type{},
		"transmission_rate": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of Lacp objects.
func (o Lacp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AggregateEthernetInterfacesLayer3 model.
func (o AggregateEthernetInterfacesLayer3) AttrTypes() map[string]attr.Type {
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
		"interface_management_profile": basetypes.StringType{},
		"ip":                           basetypes.ListType{ElemType: basetypes.StringType{}},
		"lacp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable":            basetypes.BoolType{},
				"fast_failover":     basetypes.BoolType{},
				"max_ports":         basetypes.Int64Type{},
				"mode":              basetypes.StringType{},
				"system_priority":   basetypes.Int64Type{},
				"transmission_rate": basetypes.StringType{},
			},
		},
		"mtu": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of AggregateEthernetInterfacesLayer3 objects.
func (o AggregateEthernetInterfacesLayer3) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AggEthernetArpInner model.
func (o AggEthernetArpInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"hw_address": basetypes.StringType{},
		"name":       basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AggEthernetArpInner objects.
func (o AggEthernetArpInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AggregateEthernetInterfacesLayer3DdnsConfig model.
func (o AggregateEthernetInterfacesLayer3DdnsConfig) AttrTypes() map[string]attr.Type {
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

// AttrType returns the attribute type for a list of AggregateEthernetInterfacesLayer3DdnsConfig objects.
func (o AggregateEthernetInterfacesLayer3DdnsConfig) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AggEthernetDhcpClientDhcpClient model.
func (o AggEthernetDhcpClientDhcpClient) AttrTypes() map[string]attr.Type {
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

// AttrType returns the attribute type for a list of AggEthernetDhcpClientDhcpClient objects.
func (o AggEthernetDhcpClientDhcpClient) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AggEthernetDhcpClientDhcpClientSendHostname model.
func (o AggEthernetDhcpClientDhcpClientSendHostname) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable":   basetypes.BoolType{},
		"hostname": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AggEthernetDhcpClientDhcpClientSendHostname objects.
func (o AggEthernetDhcpClientDhcpClientSendHostname) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AggregateEthernetInterfacesResourceSchema defines the schema for AggregateEthernetInterfaces resource
var AggregateEthernetInterfacesResourceSchema = schema.Schema{
	MarkdownDescription: "AggregateEthernetInterface resource",
	Attributes: map[string]schema.Attribute{
		"comment": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(1023),
			},
			MarkdownDescription: "Aggregate interface description",
			Optional:            true,
		},
		"default_value": schema.StringAttribute{
			MarkdownDescription: "Default interface assignment",
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
			MarkdownDescription: "UUID of the resource",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"layer2": schema.SingleNestedAttribute{
			Validators: []validator.Object{
				objectvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("layer3"),
				),
			},
			MarkdownDescription: "Layer2",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"lacp": schema.SingleNestedAttribute{
					MarkdownDescription: "Lacp",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"enable": schema.BoolAttribute{
							MarkdownDescription: "Enable LACP?",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"fast_failover": schema.BoolAttribute{
							MarkdownDescription: "Fast failover",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"max_ports": schema.Int64Attribute{
							Validators: []validator.Int64{
								int64validator.Between(1, 8),
							},
							MarkdownDescription: "Maximum number of physical ports bundled in the LAG",
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(8),
						},
						"mode": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("passive", "active"),
							},
							MarkdownDescription: "Mode",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("passive"),
						},
						"system_priority": schema.Int64Attribute{
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
							MarkdownDescription: "LACP system priority in system ID",
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(32768),
						},
						"transmission_rate": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("fast", "slow"),
							},
							MarkdownDescription: "Transmission mode",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("slow"),
						},
					},
				},
				"vlan_tag": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(1, 4096),
					},
					MarkdownDescription: "Assign interface to VLAN tag",
					Optional:            true,
					Computed:            true,
				},
			},
		},
		"layer3": schema.SingleNestedAttribute{
			Validators: []validator.Object{
				objectvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("layer2"),
				),
			},
			MarkdownDescription: "Layer3",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"arp": schema.ListNestedAttribute{
					MarkdownDescription: "Aggregate Ethernet ARP configuration",
					Optional:            true,
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"hw_address": schema.StringAttribute{
								MarkdownDescription: "MAC address",
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
					MarkdownDescription: "Dynamic DNS configuration specific to the Aggregate Ethernet Interface.",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"ddns_cert_profile": schema.StringAttribute{
							MarkdownDescription: "Certificate profile",
							Required:            true,
						},
						"ddns_enabled": schema.BoolAttribute{
							MarkdownDescription: "Enable DDNS?",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"ddns_hostname": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(255),
								stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_\\.\\-]+$"), "pattern must match "+"^[a-zA-Z0-9_\\.\\-]+$"),
							},
							MarkdownDescription: "Ddns hostname",
							Required:            true,
						},
						"ddns_ip": schema.StringAttribute{
							MarkdownDescription: "IP to register (static only)",
							Optional:            true,
							Computed:            true,
						},
						"ddns_update_interval": schema.Int64Attribute{
							Validators: []validator.Int64{
								int64validator.Between(1, 30),
							},
							MarkdownDescription: "Update interval (days)",
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(1),
						},
						"ddns_vendor": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(127),
							},
							MarkdownDescription: "DDNS vendor",
							Required:            true,
						},
						"ddns_vendor_config": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtMost(255),
							},
							MarkdownDescription: "DDNS vendor",
							Required:            true,
						},
					},
				},
				"dhcp_client": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("static"),
						),
					},
					MarkdownDescription: "Aggregate Ethernet DHCP Client Object",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"create_default_route": schema.BoolAttribute{
							MarkdownDescription: "Automatically create default route pointing to default gateway provided by server",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(true),
						},
						"default_route_metric": schema.Int64Attribute{
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
							MarkdownDescription: "Metric of the default route created",
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(10),
						},
						"enable": schema.BoolAttribute{
							MarkdownDescription: "Enable DHCP?",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(true),
						},
						"send_hostname": schema.SingleNestedAttribute{
							MarkdownDescription: "Aggregate Ethernet DHCP Client Send hostname",
							Optional:            true,
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"enable": schema.BoolAttribute{
									MarkdownDescription: "Enable",
									Optional:            true,
									Computed:            true,
									Default:             booldefault.StaticBool(true),
								},
								"hostname": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.LengthAtMost(64),
										stringvalidator.LengthAtLeast(1),
										stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9\\._-]+$"), "pattern must match "+"^[a-zA-Z0-9\\._-]+$"),
									},
									MarkdownDescription: "Set interface hostname",
									Optional:            true,
									Computed:            true,
									Default:             stringdefault.StaticString("system-hostname"),
								},
							},
						},
					},
				},
				"interface_management_profile": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.LengthAtMost(31),
					},
					MarkdownDescription: "Interface management profile",
					Optional:            true,
					Computed:            true,
				},
				"ip": schema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Interface IP addresses",
					Optional:            true,
					Computed:            true,
				},
				"lacp": schema.SingleNestedAttribute{
					MarkdownDescription: "Lacp",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"enable": schema.BoolAttribute{
							MarkdownDescription: "Enable LACP?",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"fast_failover": schema.BoolAttribute{
							MarkdownDescription: "Fast failover",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"max_ports": schema.Int64Attribute{
							Validators: []validator.Int64{
								int64validator.Between(1, 8),
							},
							MarkdownDescription: "Maximum number of physical ports bundled in the LAG",
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(8),
						},
						"mode": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("passive", "active"),
							},
							MarkdownDescription: "Mode",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("passive"),
						},
						"system_priority": schema.Int64Attribute{
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
							MarkdownDescription: "LACP system priority in system ID",
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(32768),
						},
						"transmission_rate": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("fast", "slow"),
							},
							MarkdownDescription: "Transmission mode",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("slow"),
						},
					},
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
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "Aggregate interface name",
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

// AggregateEthernetInterfacesDataSourceSchema defines the schema for AggregateEthernetInterfaces data source
var AggregateEthernetInterfacesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "AggregateEthernetInterface data source",
	Attributes: map[string]dsschema.Attribute{
		"comment": dsschema.StringAttribute{
			MarkdownDescription: "Aggregate interface description",
			Computed:            true,
		},
		"default_value": dsschema.StringAttribute{
			MarkdownDescription: "Default interface assignment",
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
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"layer2": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Layer2",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"lacp": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Lacp",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"enable": dsschema.BoolAttribute{
							MarkdownDescription: "Enable LACP?",
							Computed:            true,
						},
						"fast_failover": dsschema.BoolAttribute{
							MarkdownDescription: "Fast failover",
							Computed:            true,
						},
						"max_ports": dsschema.Int64Attribute{
							MarkdownDescription: "Maximum number of physical ports bundled in the LAG",
							Computed:            true,
						},
						"mode": dsschema.StringAttribute{
							MarkdownDescription: "Mode",
							Computed:            true,
						},
						"system_priority": dsschema.Int64Attribute{
							MarkdownDescription: "LACP system priority in system ID",
							Computed:            true,
						},
						"transmission_rate": dsschema.StringAttribute{
							MarkdownDescription: "Transmission mode",
							Computed:            true,
						},
					},
				},
				"vlan_tag": dsschema.Int64Attribute{
					MarkdownDescription: "Assign interface to VLAN tag",
					Computed:            true,
				},
			},
		},
		"layer3": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Layer3",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"arp": dsschema.ListNestedAttribute{
					MarkdownDescription: "Aggregate Ethernet ARP configuration",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"hw_address": dsschema.StringAttribute{
								MarkdownDescription: "MAC address",
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
					MarkdownDescription: "Dynamic DNS configuration specific to the Aggregate Ethernet Interface.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"ddns_cert_profile": dsschema.StringAttribute{
							MarkdownDescription: "Certificate profile",
							Computed:            true,
						},
						"ddns_enabled": dsschema.BoolAttribute{
							MarkdownDescription: "Enable DDNS?",
							Computed:            true,
						},
						"ddns_hostname": dsschema.StringAttribute{
							MarkdownDescription: "Ddns hostname",
							Computed:            true,
						},
						"ddns_ip": dsschema.StringAttribute{
							MarkdownDescription: "IP to register (static only)",
							Computed:            true,
						},
						"ddns_update_interval": dsschema.Int64Attribute{
							MarkdownDescription: "Update interval (days)",
							Computed:            true,
						},
						"ddns_vendor": dsschema.StringAttribute{
							MarkdownDescription: "DDNS vendor",
							Computed:            true,
						},
						"ddns_vendor_config": dsschema.StringAttribute{
							MarkdownDescription: "DDNS vendor",
							Computed:            true,
						},
					},
				},
				"dhcp_client": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Aggregate Ethernet DHCP Client Object",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"create_default_route": dsschema.BoolAttribute{
							MarkdownDescription: "Automatically create default route pointing to default gateway provided by server",
							Computed:            true,
						},
						"default_route_metric": dsschema.Int64Attribute{
							MarkdownDescription: "Metric of the default route created",
							Computed:            true,
						},
						"enable": dsschema.BoolAttribute{
							MarkdownDescription: "Enable DHCP?",
							Computed:            true,
						},
						"send_hostname": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Aggregate Ethernet DHCP Client Send hostname",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"enable": dsschema.BoolAttribute{
									MarkdownDescription: "Enable",
									Computed:            true,
								},
								"hostname": dsschema.StringAttribute{
									MarkdownDescription: "Set interface hostname",
									Computed:            true,
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
				"lacp": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Lacp",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"enable": dsschema.BoolAttribute{
							MarkdownDescription: "Enable LACP?",
							Computed:            true,
						},
						"fast_failover": dsschema.BoolAttribute{
							MarkdownDescription: "Fast failover",
							Computed:            true,
						},
						"max_ports": dsschema.Int64Attribute{
							MarkdownDescription: "Maximum number of physical ports bundled in the LAG",
							Computed:            true,
						},
						"mode": dsschema.StringAttribute{
							MarkdownDescription: "Mode",
							Computed:            true,
						},
						"system_priority": dsschema.Int64Attribute{
							MarkdownDescription: "LACP system priority in system ID",
							Computed:            true,
						},
						"transmission_rate": dsschema.StringAttribute{
							MarkdownDescription: "Transmission mode",
							Computed:            true,
						},
					},
				},
				"mtu": dsschema.Int64Attribute{
					MarkdownDescription: "MTU",
					Computed:            true,
				},
			},
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Aggregate interface name",
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

// AggregateEthernetInterfacesListModel represents the data model for a list data source.
type AggregateEthernetInterfacesListModel struct {
	Tfid    types.String                  `tfsdk:"tfid"`
	Data    []AggregateEthernetInterfaces `tfsdk:"data"`
	Limit   types.Int64                   `tfsdk:"limit"`
	Offset  types.Int64                   `tfsdk:"offset"`
	Name    types.String                  `tfsdk:"name"`
	Total   types.Int64                   `tfsdk:"total"`
	Folder  types.String                  `tfsdk:"folder"`
	Snippet types.String                  `tfsdk:"snippet"`
	Device  types.String                  `tfsdk:"device"`
}

// AggregateEthernetInterfacesListDataSourceSchema defines the schema for a list data source.
var AggregateEthernetInterfacesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: AggregateEthernetInterfacesDataSourceSchema.Attributes,
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
