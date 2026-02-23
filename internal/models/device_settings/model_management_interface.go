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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// Package: device_settings
// This file contains models for the device_settings SDK package

// ManagementInterface represents the Terraform model for ManagementInterface
type ManagementInterface struct {
	Tfid                types.String          `tfsdk:"tfid"`
	Device              basetypes.StringValue `tfsdk:"device"`
	Folder              basetypes.StringValue `tfsdk:"folder"`
	Id                  basetypes.StringValue `tfsdk:"id"`
	ManagementInterface basetypes.ObjectValue `tfsdk:"management_interface"`
	Snippet             basetypes.StringValue `tfsdk:"snippet"`
}

// ManagementInterfaceManagementInterface represents a nested structure within the ManagementInterface model
type ManagementInterfaceManagementInterface struct {
	MgmtType    basetypes.ObjectValue `tfsdk:"mgmt_type"`
	Mtu         basetypes.Int64Value  `tfsdk:"mtu"`
	PermittedIp basetypes.ListValue   `tfsdk:"permitted_ip"`
	Service     basetypes.ObjectValue `tfsdk:"service"`
	SpeedDuplex basetypes.StringValue `tfsdk:"speed_duplex"`
}

// ManagementInterfaceManagementInterfaceMgmtType represents a nested structure within the ManagementInterface model
type ManagementInterfaceManagementInterfaceMgmtType struct {
	DhcpClient basetypes.ObjectValue `tfsdk:"dhcp_client"`
	Static     basetypes.ObjectValue `tfsdk:"static"`
}

// ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient represents a nested structure within the ManagementInterface model
type ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient struct {
	AcceptDhcpDomain   basetypes.BoolValue `tfsdk:"accept_dhcp_domain"`
	AcceptDhcpHostname basetypes.BoolValue `tfsdk:"accept_dhcp_hostname"`
	SendClientId       basetypes.BoolValue `tfsdk:"send_client_id"`
	SendHostname       basetypes.BoolValue `tfsdk:"send_hostname"`
}

// ManagementInterfaceManagementInterfaceMgmtTypeStatic represents a nested structure within the ManagementInterface model
type ManagementInterfaceManagementInterfaceMgmtTypeStatic struct {
	DefaultGateway basetypes.StringValue `tfsdk:"default_gateway"`
	IpAddress      basetypes.StringValue `tfsdk:"ip_address"`
	Netmask        basetypes.StringValue `tfsdk:"netmask"`
}

// ManagementInterfaceManagementInterfacePermittedIpInner represents a nested structure within the ManagementInterface model
type ManagementInterfaceManagementInterfacePermittedIpInner struct {
	Description basetypes.StringValue `tfsdk:"description"`
	Name        basetypes.StringValue `tfsdk:"name"`
}

// ManagementInterfaceManagementInterfaceService represents a nested structure within the ManagementInterface model
type ManagementInterfaceManagementInterfaceService struct {
	DisableHttp                    basetypes.BoolValue `tfsdk:"disable_http"`
	DisableHttpOcsp                basetypes.BoolValue `tfsdk:"disable_http_ocsp"`
	DisableHttps                   basetypes.BoolValue `tfsdk:"disable_https"`
	DisableIcmp                    basetypes.BoolValue `tfsdk:"disable_icmp"`
	DisableSnmp                    basetypes.BoolValue `tfsdk:"disable_snmp"`
	DisableSsh                     basetypes.BoolValue `tfsdk:"disable_ssh"`
	DisableTelnet                  basetypes.BoolValue `tfsdk:"disable_telnet"`
	DisableUseridService           basetypes.BoolValue `tfsdk:"disable_userid_service"`
	DisableUseridSyslogListenerSsl basetypes.BoolValue `tfsdk:"disable_userid_syslog_listener_ssl"`
	DisableUseridSyslogListenerUdp basetypes.BoolValue `tfsdk:"disable_userid_syslog_listener_udp"`
}

// AttrTypes defines the attribute types for the ManagementInterface model.
func (o ManagementInterface) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":   basetypes.StringType{},
		"device": basetypes.StringType{},
		"folder": basetypes.StringType{},
		"id":     basetypes.StringType{},
		"management_interface": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"mgmt_type": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"dhcp_client": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"accept_dhcp_domain":   basetypes.BoolType{},
								"accept_dhcp_hostname": basetypes.BoolType{},
								"send_client_id":       basetypes.BoolType{},
								"send_hostname":        basetypes.BoolType{},
							},
						},
						"static": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"default_gateway": basetypes.StringType{},
								"ip_address":      basetypes.StringType{},
								"netmask":         basetypes.StringType{},
							},
						},
					},
				},
				"mtu": basetypes.Int64Type{},
				"permitted_ip": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"description": basetypes.StringType{},
						"name":        basetypes.StringType{},
					},
				}},
				"service": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"disable_http":                       basetypes.BoolType{},
						"disable_http_ocsp":                  basetypes.BoolType{},
						"disable_https":                      basetypes.BoolType{},
						"disable_icmp":                       basetypes.BoolType{},
						"disable_snmp":                       basetypes.BoolType{},
						"disable_ssh":                        basetypes.BoolType{},
						"disable_telnet":                     basetypes.BoolType{},
						"disable_userid_service":             basetypes.BoolType{},
						"disable_userid_syslog_listener_ssl": basetypes.BoolType{},
						"disable_userid_syslog_listener_udp": basetypes.BoolType{},
					},
				},
				"speed_duplex": basetypes.StringType{},
			},
		},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ManagementInterface objects.
func (o ManagementInterface) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ManagementInterfaceManagementInterface model.
func (o ManagementInterfaceManagementInterface) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"mgmt_type": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"dhcp_client": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"accept_dhcp_domain":   basetypes.BoolType{},
						"accept_dhcp_hostname": basetypes.BoolType{},
						"send_client_id":       basetypes.BoolType{},
						"send_hostname":        basetypes.BoolType{},
					},
				},
				"static": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"default_gateway": basetypes.StringType{},
						"ip_address":      basetypes.StringType{},
						"netmask":         basetypes.StringType{},
					},
				},
			},
		},
		"mtu": basetypes.Int64Type{},
		"permitted_ip": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"description": basetypes.StringType{},
				"name":        basetypes.StringType{},
			},
		}},
		"service": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"disable_http":                       basetypes.BoolType{},
				"disable_http_ocsp":                  basetypes.BoolType{},
				"disable_https":                      basetypes.BoolType{},
				"disable_icmp":                       basetypes.BoolType{},
				"disable_snmp":                       basetypes.BoolType{},
				"disable_ssh":                        basetypes.BoolType{},
				"disable_telnet":                     basetypes.BoolType{},
				"disable_userid_service":             basetypes.BoolType{},
				"disable_userid_syslog_listener_ssl": basetypes.BoolType{},
				"disable_userid_syslog_listener_udp": basetypes.BoolType{},
			},
		},
		"speed_duplex": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ManagementInterfaceManagementInterface objects.
func (o ManagementInterfaceManagementInterface) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ManagementInterfaceManagementInterfaceMgmtType model.
func (o ManagementInterfaceManagementInterfaceMgmtType) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"dhcp_client": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"accept_dhcp_domain":   basetypes.BoolType{},
				"accept_dhcp_hostname": basetypes.BoolType{},
				"send_client_id":       basetypes.BoolType{},
				"send_hostname":        basetypes.BoolType{},
			},
		},
		"static": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"default_gateway": basetypes.StringType{},
				"ip_address":      basetypes.StringType{},
				"netmask":         basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ManagementInterfaceManagementInterfaceMgmtType objects.
func (o ManagementInterfaceManagementInterfaceMgmtType) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient model.
func (o ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"accept_dhcp_domain":   basetypes.BoolType{},
		"accept_dhcp_hostname": basetypes.BoolType{},
		"send_client_id":       basetypes.BoolType{},
		"send_hostname":        basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient objects.
func (o ManagementInterfaceManagementInterfaceMgmtTypeDhcpClient) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ManagementInterfaceManagementInterfaceMgmtTypeStatic model.
func (o ManagementInterfaceManagementInterfaceMgmtTypeStatic) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"default_gateway": basetypes.StringType{},
		"ip_address":      basetypes.StringType{},
		"netmask":         basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ManagementInterfaceManagementInterfaceMgmtTypeStatic objects.
func (o ManagementInterfaceManagementInterfaceMgmtTypeStatic) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ManagementInterfaceManagementInterfacePermittedIpInner model.
func (o ManagementInterfaceManagementInterfacePermittedIpInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"description": basetypes.StringType{},
		"name":        basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ManagementInterfaceManagementInterfacePermittedIpInner objects.
func (o ManagementInterfaceManagementInterfacePermittedIpInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ManagementInterfaceManagementInterfaceService model.
func (o ManagementInterfaceManagementInterfaceService) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"disable_http":                       basetypes.BoolType{},
		"disable_http_ocsp":                  basetypes.BoolType{},
		"disable_https":                      basetypes.BoolType{},
		"disable_icmp":                       basetypes.BoolType{},
		"disable_snmp":                       basetypes.BoolType{},
		"disable_ssh":                        basetypes.BoolType{},
		"disable_telnet":                     basetypes.BoolType{},
		"disable_userid_service":             basetypes.BoolType{},
		"disable_userid_syslog_listener_ssl": basetypes.BoolType{},
		"disable_userid_syslog_listener_udp": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of ManagementInterfaceManagementInterfaceService objects.
func (o ManagementInterfaceManagementInterfaceService) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// ManagementInterfaceResourceSchema defines the schema for ManagementInterface resource
var ManagementInterfaceResourceSchema = schema.Schema{
	MarkdownDescription: "ManagementInterface resource",
	Attributes: map[string]schema.Attribute{
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
		"folder": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("device"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
				utils.FolderValidator(),
			},
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
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
		"management_interface": schema.SingleNestedAttribute{
			MarkdownDescription: "Management interface",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"mgmt_type": schema.SingleNestedAttribute{
					MarkdownDescription: "IP type",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"dhcp_client": schema.SingleNestedAttribute{
							Validators: []validator.Object{
								objectvalidator.ConflictsWith(
									path.MatchRelative().AtParent().AtName("static"),
								),
							},
							MarkdownDescription: "Dhcp client\n\n> ℹ️ **Note:** You must specify exactly one of `dhcp_client` and `static`.",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"accept_dhcp_domain": schema.BoolAttribute{
									MarkdownDescription: "Accept DHCP server provided domain name",
									Optional:            true,
									Computed:            true,
									Default:             booldefault.StaticBool(false),
								},
								"accept_dhcp_hostname": schema.BoolAttribute{
									MarkdownDescription: "Accept DHCP server provided hostname",
									Optional:            true,
									Computed:            true,
									Default:             booldefault.StaticBool(false),
								},
								"send_client_id": schema.BoolAttribute{
									MarkdownDescription: "Send client ID",
									Optional:            true,
									Computed:            true,
									Default:             booldefault.StaticBool(false),
								},
								"send_hostname": schema.BoolAttribute{
									MarkdownDescription: "Send hostname",
									Optional:            true,
									Computed:            true,
									Default:             booldefault.StaticBool(false),
								},
							},
						},
						"static": schema.SingleNestedAttribute{
							Validators: []validator.Object{
								objectvalidator.ConflictsWith(
									path.MatchRelative().AtParent().AtName("dhcp_client"),
								),
							},
							MarkdownDescription: "Static\n\n> ℹ️ **Note:** You must specify exactly one of `dhcp_client` and `static`.",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"default_gateway": schema.StringAttribute{
									MarkdownDescription: "Default gateway",
									Required:            true,
								},
								"ip_address": schema.StringAttribute{
									MarkdownDescription: "IP address",
									Required:            true,
								},
								"netmask": schema.StringAttribute{
									MarkdownDescription: "Netmask",
									Required:            true,
								},
							},
						},
					},
				},
				"mtu": schema.Int64Attribute{
					MarkdownDescription: "MTU",
					Optional:            true,
					Computed:            true,
					Default:             int64default.StaticInt64(1500),
				},
				"permitted_ip": schema.ListNestedAttribute{
					MarkdownDescription: "Permitting IP addresses",
					Optional:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"description": schema.StringAttribute{
								MarkdownDescription: "Description",
								Optional:            true,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: "IP address",
								Optional:            true,
							},
						},
					},
				},
				"service": schema.SingleNestedAttribute{
					MarkdownDescription: "Network services",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"disable_http": schema.BoolAttribute{
							MarkdownDescription: "HTTP",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"disable_http_ocsp": schema.BoolAttribute{
							MarkdownDescription: "HTTP OCSP",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"disable_https": schema.BoolAttribute{
							MarkdownDescription: "HTTPS",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(true),
						},
						"disable_icmp": schema.BoolAttribute{
							MarkdownDescription: "Ping",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"disable_snmp": schema.BoolAttribute{
							MarkdownDescription: "SNMP",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"disable_ssh": schema.BoolAttribute{
							MarkdownDescription: "SSH",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(true),
						},
						"disable_telnet": schema.BoolAttribute{
							MarkdownDescription: "Telnet",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"disable_userid_service": schema.BoolAttribute{
							MarkdownDescription: "User-ID",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"disable_userid_syslog_listener_ssl": schema.BoolAttribute{
							MarkdownDescription: "User-ID syslog listener over SSL",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"disable_userid_syslog_listener_udp": schema.BoolAttribute{
							MarkdownDescription: "User-ID syslog listener over UDP",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
					},
				},
				"speed_duplex": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("auto-negotiate", "10Mbps-half-duplex", "10Mbps-full-duplex", "100Mbps-half-duplex", "100Mbps-full-duplex", "1Gbps-half-duplex", "1Gbps-full-duplex"),
					},
					MarkdownDescription: "Speed and duplex",
					Optional:            true,
					Computed:            true,
					Default:             stringdefault.StaticString("auto-negotiate"),
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
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
			},
			MarkdownDescription: "The snippet in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
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

// ManagementInterfaceDataSourceSchema defines the schema for ManagementInterface data source
var ManagementInterfaceDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "ManagementInterface data source",
	Attributes: map[string]dsschema.Attribute{
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"management_interface": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Management interface",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"mgmt_type": dsschema.SingleNestedAttribute{
					MarkdownDescription: "IP type",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"dhcp_client": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Dhcp client\n\n> ℹ️ **Note:** You must specify exactly one of `dhcp_client` and `static`.",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"accept_dhcp_domain": dsschema.BoolAttribute{
									MarkdownDescription: "Accept DHCP server provided domain name",
									Computed:            true,
								},
								"accept_dhcp_hostname": dsschema.BoolAttribute{
									MarkdownDescription: "Accept DHCP server provided hostname",
									Computed:            true,
								},
								"send_client_id": dsschema.BoolAttribute{
									MarkdownDescription: "Send client ID",
									Computed:            true,
								},
								"send_hostname": dsschema.BoolAttribute{
									MarkdownDescription: "Send hostname",
									Computed:            true,
								},
							},
						},
						"static": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Static\n\n> ℹ️ **Note:** You must specify exactly one of `dhcp_client` and `static`.",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"default_gateway": dsschema.StringAttribute{
									MarkdownDescription: "Default gateway",
									Computed:            true,
								},
								"ip_address": dsschema.StringAttribute{
									MarkdownDescription: "IP address",
									Computed:            true,
								},
								"netmask": dsschema.StringAttribute{
									MarkdownDescription: "Netmask",
									Computed:            true,
								},
							},
						},
					},
				},
				"mtu": dsschema.Int64Attribute{
					MarkdownDescription: "MTU",
					Computed:            true,
				},
				"permitted_ip": dsschema.ListNestedAttribute{
					MarkdownDescription: "Permitting IP addresses",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"description": dsschema.StringAttribute{
								MarkdownDescription: "Description",
								Computed:            true,
							},
							"name": dsschema.StringAttribute{
								MarkdownDescription: "IP address",
								Computed:            true,
							},
						},
					},
				},
				"service": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Network services",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"disable_http": dsschema.BoolAttribute{
							MarkdownDescription: "HTTP",
							Computed:            true,
						},
						"disable_http_ocsp": dsschema.BoolAttribute{
							MarkdownDescription: "HTTP OCSP",
							Computed:            true,
						},
						"disable_https": dsschema.BoolAttribute{
							MarkdownDescription: "HTTPS",
							Computed:            true,
						},
						"disable_icmp": dsschema.BoolAttribute{
							MarkdownDescription: "Ping",
							Computed:            true,
						},
						"disable_snmp": dsschema.BoolAttribute{
							MarkdownDescription: "SNMP",
							Computed:            true,
						},
						"disable_ssh": dsschema.BoolAttribute{
							MarkdownDescription: "SSH",
							Computed:            true,
						},
						"disable_telnet": dsschema.BoolAttribute{
							MarkdownDescription: "Telnet",
							Computed:            true,
						},
						"disable_userid_service": dsschema.BoolAttribute{
							MarkdownDescription: "User-ID",
							Computed:            true,
						},
						"disable_userid_syslog_listener_ssl": dsschema.BoolAttribute{
							MarkdownDescription: "User-ID syslog listener over SSL",
							Computed:            true,
						},
						"disable_userid_syslog_listener_udp": dsschema.BoolAttribute{
							MarkdownDescription: "User-ID syslog listener over UDP",
							Computed:            true,
						},
					},
				},
				"speed_duplex": dsschema.StringAttribute{
					MarkdownDescription: "Speed and duplex",
					Computed:            true,
				},
			},
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// ManagementInterfaceListModel represents the data model for a list data source.
type ManagementInterfaceListModel struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Data    []ManagementInterface `tfsdk:"data"`
	Limit   types.Int64           `tfsdk:"limit"`
	Offset  types.Int64           `tfsdk:"offset"`
	Name    types.String          `tfsdk:"name"`
	Total   types.Int64           `tfsdk:"total"`
	Folder  types.String          `tfsdk:"folder"`
	Snippet types.String          `tfsdk:"snippet"`
	Device  types.String          `tfsdk:"device"`
}

// ManagementInterfaceListDataSourceSchema defines the schema for a list data source.
var ManagementInterfaceListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: ManagementInterfaceDataSourceSchema.Attributes,
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
