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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: network_services
// This file contains models for the network_services SDK package

// VlanInterfaces represents the Terraform model for VlanInterfaces
type VlanInterfaces struct {
	Tfid                       types.String          `tfsdk:"tfid"`
	Arp                        basetypes.ListValue   `tfsdk:"arp"`
	Comment                    basetypes.StringValue `tfsdk:"comment"`
	DdnsConfig                 basetypes.ObjectValue `tfsdk:"ddns_config"`
	DefaultValue               basetypes.StringValue `tfsdk:"default_value"`
	Device                     basetypes.StringValue `tfsdk:"device"`
	DhcpClient                 basetypes.ObjectValue `tfsdk:"dhcp_client"`
	Folder                     basetypes.StringValue `tfsdk:"folder"`
	Id                         basetypes.StringValue `tfsdk:"id"`
	InterfaceManagementProfile basetypes.StringValue `tfsdk:"interface_management_profile"`
	Ip                         basetypes.ListValue   `tfsdk:"ip"`
	Mtu                        basetypes.Int64Value  `tfsdk:"mtu"`
	Name                       basetypes.StringValue `tfsdk:"name"`
	Snippet                    basetypes.StringValue `tfsdk:"snippet"`
	VlanTag                    basetypes.StringValue `tfsdk:"vlan_tag"`
}

// VlanInterfacesArpInner represents a nested structure within the VlanInterfaces model
type VlanInterfacesArpInner struct {
	HwAddress basetypes.StringValue `tfsdk:"hw_address"`
	Interface basetypes.StringValue `tfsdk:"interface"`
	Name      basetypes.StringValue `tfsdk:"name"`
}

// VlanInterfacesDdnsConfig represents a nested structure within the VlanInterfaces model
type VlanInterfacesDdnsConfig struct {
	DdnsCertProfile    basetypes.StringValue `tfsdk:"ddns_cert_profile"`
	DdnsEnabled        basetypes.BoolValue   `tfsdk:"ddns_enabled"`
	DdnsHostname       basetypes.StringValue `tfsdk:"ddns_hostname"`
	DdnsIp             basetypes.StringValue `tfsdk:"ddns_ip"`
	DdnsUpdateInterval basetypes.Int64Value  `tfsdk:"ddns_update_interval"`
	DdnsVendor         basetypes.StringValue `tfsdk:"ddns_vendor"`
	DdnsVendorConfig   basetypes.StringValue `tfsdk:"ddns_vendor_config"`
}

// VlanInterfacesDhcpClient represents a nested structure within the VlanInterfaces model
type VlanInterfacesDhcpClient struct {
	CreateDefaultRoute basetypes.BoolValue   `tfsdk:"create_default_route"`
	DefaultRouteMetric basetypes.Int64Value  `tfsdk:"default_route_metric"`
	Enable             basetypes.BoolValue   `tfsdk:"enable"`
	SendHostname       basetypes.ObjectValue `tfsdk:"send_hostname"`
}

// VlanInterfacesDhcpClientSendHostname represents a nested structure within the VlanInterfaces model
type VlanInterfacesDhcpClientSendHostname struct {
	Enable   basetypes.BoolValue   `tfsdk:"enable"`
	Hostname basetypes.StringValue `tfsdk:"hostname"`
}

// VlanInterfacesIpInner represents a nested structure within the VlanInterfaces model
type VlanInterfacesIpInner struct {
	Name basetypes.StringValue `tfsdk:"name"`
}

// AttrTypes defines the attribute types for the VlanInterfaces model.
func (o VlanInterfaces) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"arp": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"hw_address": basetypes.StringType{},
				"interface":  basetypes.StringType{},
				"name":       basetypes.StringType{},
			},
		}},
		"comment": basetypes.StringType{},
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
		"default_value": basetypes.StringType{},
		"device":        basetypes.StringType{},
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
		"folder":                       basetypes.StringType{},
		"id":                           basetypes.StringType{},
		"interface_management_profile": basetypes.StringType{},
		"ip": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name": basetypes.StringType{},
			},
		}},
		"mtu":      basetypes.Int64Type{},
		"name":     basetypes.StringType{},
		"snippet":  basetypes.StringType{},
		"vlan_tag": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of VlanInterfaces objects.
func (o VlanInterfaces) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the VlanInterfacesArpInner model.
func (o VlanInterfacesArpInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"hw_address": basetypes.StringType{},
		"interface":  basetypes.StringType{},
		"name":       basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of VlanInterfacesArpInner objects.
func (o VlanInterfacesArpInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the VlanInterfacesDdnsConfig model.
func (o VlanInterfacesDdnsConfig) AttrTypes() map[string]attr.Type {
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

// AttrType returns the attribute type for a list of VlanInterfacesDdnsConfig objects.
func (o VlanInterfacesDdnsConfig) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the VlanInterfacesDhcpClient model.
func (o VlanInterfacesDhcpClient) AttrTypes() map[string]attr.Type {
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

// AttrType returns the attribute type for a list of VlanInterfacesDhcpClient objects.
func (o VlanInterfacesDhcpClient) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the VlanInterfacesDhcpClientSendHostname model.
func (o VlanInterfacesDhcpClientSendHostname) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable":   basetypes.BoolType{},
		"hostname": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of VlanInterfacesDhcpClientSendHostname objects.
func (o VlanInterfacesDhcpClientSendHostname) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the VlanInterfacesIpInner model.
func (o VlanInterfacesIpInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of VlanInterfacesIpInner objects.
func (o VlanInterfacesIpInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// VlanInterfacesResourceSchema defines the schema for VlanInterfaces resource
var VlanInterfacesResourceSchema = schema.Schema{
	MarkdownDescription: "VlanInterface resource",
	Attributes: map[string]schema.Attribute{
		"arp": schema.ListNestedAttribute{
			MarkdownDescription: "ARP configuration",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"hw_address": schema.StringAttribute{
						MarkdownDescription: "MAC address",
						Optional:            true,
					},
					"interface": schema.StringAttribute{
						MarkdownDescription: "ARP interface",
						Optional:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "IP address",
						Optional:            true,
					},
				},
			},
		},
		"comment": schema.StringAttribute{
			MarkdownDescription: "Description",
			Optional:            true,
		},
		"ddns_config": schema.SingleNestedAttribute{
			MarkdownDescription: "Dynamic DNS configuration specific to the Vlan Interfaces.",
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
		"default_value": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.RegexMatches(regexp.MustCompile("^vlan\\.([1-9]\\d{0,2}|[1-3]\\d{3}|40[0-8]\\d|409[0-6])$"), "pattern must match "+"^vlan\\.([1-9]\\d{0,2}|[1-3]\\d{3}|40[0-8]\\d|409[0-6])$"),
			},
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
			MarkdownDescription: "The device in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"dhcp_client": schema.SingleNestedAttribute{
			Validators: []validator.Object{
				objectvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("ip"),
				),
			},
			MarkdownDescription: "Vlan interfaces DHCP Client Object\n> ℹ️ **Note:** You must specify exactly one of `dhcp_client` and `ip`.",
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
					MarkdownDescription: "Send hostname",
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
		"folder": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("device"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
			},
			MarkdownDescription: "The folder in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
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
		"interface_management_profile": schema.StringAttribute{
			MarkdownDescription: "Interface management profile",
			Optional:            true,
		},
		"ip": schema.ListNestedAttribute{
			Validators: []validator.List{
				listvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("dhcp_client"),
				),
			},
			MarkdownDescription: "VLAN Interface IP Parent\n> ℹ️ **Note:** You must specify exactly one of `dhcp_client` and `ip`.",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						MarkdownDescription: "VLAN Interface IP address(es)",
						Required:            true,
					},
				},
			},
		},
		"mtu": schema.Int64Attribute{
			Validators: []validator.Int64{
				int64validator.Between(576, 9216),
			},
			MarkdownDescription: "MTU",
			Optional:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "L3 sub-interface name",
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
			MarkdownDescription: "The snippet in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
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
		"vlan_tag": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.RegexMatches(regexp.MustCompile("^([1-9]\\d{0,2}|[1-3]\\d{3}|40[0-8]\\d|409[0-6])$"), "pattern must match "+"^([1-9]\\d{0,2}|[1-3]\\d{3}|40[0-8]\\d|409[0-6])$"),
			},
			MarkdownDescription: "VLAN tag",
			Optional:            true,
		},
	},
}

// VlanInterfacesDataSourceSchema defines the schema for VlanInterfaces data source
var VlanInterfacesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "VlanInterface data source",
	Attributes: map[string]dsschema.Attribute{
		"arp": dsschema.ListNestedAttribute{
			MarkdownDescription: "ARP configuration",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"hw_address": dsschema.StringAttribute{
						MarkdownDescription: "MAC address",
						Computed:            true,
					},
					"interface": dsschema.StringAttribute{
						MarkdownDescription: "ARP interface",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "IP address",
						Computed:            true,
					},
				},
			},
		},
		"comment": dsschema.StringAttribute{
			MarkdownDescription: "Description",
			Computed:            true,
		},
		"ddns_config": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Dynamic DNS configuration specific to the Vlan Interfaces.",
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
		"default_value": dsschema.StringAttribute{
			MarkdownDescription: "Default interface assignment",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"dhcp_client": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Vlan interfaces DHCP Client Object\n> ℹ️ **Note:** You must specify exactly one of `dhcp_client` and `ip`.",
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
					MarkdownDescription: "Send hostname",
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
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"interface_management_profile": dsschema.StringAttribute{
			MarkdownDescription: "Interface management profile",
			Computed:            true,
		},
		"ip": dsschema.ListNestedAttribute{
			MarkdownDescription: "VLAN Interface IP Parent\n> ℹ️ **Note:** You must specify exactly one of `dhcp_client` and `ip`.",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"name": dsschema.StringAttribute{
						MarkdownDescription: "VLAN Interface IP address(es)",
						Computed:            true,
					},
				},
			},
		},
		"mtu": dsschema.Int64Attribute{
			MarkdownDescription: "MTU",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "L3 sub-interface name",
			Optional:            true,
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"vlan_tag": dsschema.StringAttribute{
			MarkdownDescription: "VLAN tag",
			Computed:            true,
		},
	},
}

// VlanInterfacesListModel represents the data model for a list data source.
type VlanInterfacesListModel struct {
	Tfid    types.String     `tfsdk:"tfid"`
	Data    []VlanInterfaces `tfsdk:"data"`
	Limit   types.Int64      `tfsdk:"limit"`
	Offset  types.Int64      `tfsdk:"offset"`
	Name    types.String     `tfsdk:"name"`
	Total   types.Int64      `tfsdk:"total"`
	Folder  types.String     `tfsdk:"folder"`
	Snippet types.String     `tfsdk:"snippet"`
	Device  types.String     `tfsdk:"device"`
}

// VlanInterfacesListDataSourceSchema defines the schema for a list data source.
var VlanInterfacesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: VlanInterfacesDataSourceSchema.Attributes,
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
