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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: network_services
// This file contains models for the network_services SDK package

// DhcpInterfaces represents the Terraform model for DhcpInterfaces
type DhcpInterfaces struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Device  basetypes.StringValue `tfsdk:"device"`
	Folder  basetypes.StringValue `tfsdk:"folder"`
	Id      basetypes.StringValue `tfsdk:"id"`
	Name    basetypes.StringValue `tfsdk:"name"`
	Relay   basetypes.ObjectValue `tfsdk:"relay"`
	Server  basetypes.ObjectValue `tfsdk:"server"`
	Snippet basetypes.StringValue `tfsdk:"snippet"`
}

// DhcpInterfacesRelay represents a nested structure within the DhcpInterfaces model
type DhcpInterfacesRelay struct {
	Ip basetypes.ObjectValue `tfsdk:"ip"`
}

// DhcpInterfacesRelayIp represents a nested structure within the DhcpInterfaces model
type DhcpInterfacesRelayIp struct {
	Enabled basetypes.BoolValue `tfsdk:"enabled"`
	Server  basetypes.ListValue `tfsdk:"server"`
}

// DhcpInterfacesServer represents a nested structure within the DhcpInterfaces model
type DhcpInterfacesServer struct {
	IpPool   basetypes.ListValue   `tfsdk:"ip_pool"`
	Mode     basetypes.StringValue `tfsdk:"mode"`
	Option   basetypes.ObjectValue `tfsdk:"option"`
	ProbeIp  basetypes.BoolValue   `tfsdk:"probe_ip"`
	Reserved basetypes.ListValue   `tfsdk:"reserved"`
}

// DhcpInterfacesServerOption represents a nested structure within the DhcpInterfaces model
type DhcpInterfacesServerOption struct {
	Dns         basetypes.ObjectValue `tfsdk:"dns"`
	DnsSuffix   basetypes.StringValue `tfsdk:"dns_suffix"`
	Gateway     basetypes.StringValue `tfsdk:"gateway"`
	Inheritance basetypes.ObjectValue `tfsdk:"inheritance"`
	Lease       basetypes.ObjectValue `tfsdk:"lease"`
	Nis         basetypes.ObjectValue `tfsdk:"nis"`
	Ntp         basetypes.ObjectValue `tfsdk:"ntp"`
	Pop3Server  basetypes.StringValue `tfsdk:"pop3_server"`
	SmtpServer  basetypes.StringValue `tfsdk:"smtp_server"`
	SubnetMask  basetypes.StringValue `tfsdk:"subnet_mask"`
	UserDefined basetypes.ListValue   `tfsdk:"user_defined"`
	Wins        basetypes.ObjectValue `tfsdk:"wins"`
}

// DhcpInterfacesServerOptionDns represents a nested structure within the DhcpInterfaces model
type DhcpInterfacesServerOptionDns struct {
	Primary   basetypes.StringValue `tfsdk:"primary"`
	Secondary basetypes.StringValue `tfsdk:"secondary"`
}

// DhcpInterfacesServerOptionInheritance represents a nested structure within the DhcpInterfaces model
type DhcpInterfacesServerOptionInheritance struct {
	Source basetypes.StringValue `tfsdk:"source"`
}

// DhcpInterfacesServerOptionLease represents a nested structure within the DhcpInterfaces model
type DhcpInterfacesServerOptionLease struct {
	Timeout   basetypes.Int64Value  `tfsdk:"timeout"`
	Unlimited basetypes.ObjectValue `tfsdk:"unlimited"`
}

// DhcpInterfacesServerOptionNis represents a nested structure within the DhcpInterfaces model
type DhcpInterfacesServerOptionNis struct {
	Primary   basetypes.StringValue `tfsdk:"primary"`
	Secondary basetypes.StringValue `tfsdk:"secondary"`
}

// DhcpInterfacesServerOptionNtp represents a nested structure within the DhcpInterfaces model
type DhcpInterfacesServerOptionNtp struct {
	Primary   basetypes.StringValue `tfsdk:"primary"`
	Secondary basetypes.StringValue `tfsdk:"secondary"`
}

// DhcpInterfacesServerOptionUserDefinedInner represents a nested structure within the DhcpInterfaces model
type DhcpInterfacesServerOptionUserDefinedInner struct {
	Ascii     basetypes.ListValue   `tfsdk:"ascii"`
	Code      basetypes.Int64Value  `tfsdk:"code"`
	Hex       basetypes.ListValue   `tfsdk:"hex"`
	Inherited basetypes.BoolValue   `tfsdk:"inherited"`
	Ip        basetypes.ListValue   `tfsdk:"ip"`
	Name      basetypes.StringValue `tfsdk:"name"`
}

// DhcpInterfacesServerOptionWins represents a nested structure within the DhcpInterfaces model
type DhcpInterfacesServerOptionWins struct {
	Primary   basetypes.StringValue `tfsdk:"primary"`
	Secondary basetypes.StringValue `tfsdk:"secondary"`
}

// DhcpInterfacesServerReservedInner represents a nested structure within the DhcpInterfaces model
type DhcpInterfacesServerReservedInner struct {
	Description basetypes.StringValue `tfsdk:"description"`
	Mac         basetypes.StringValue `tfsdk:"mac"`
	Name        basetypes.StringValue `tfsdk:"name"`
}

// AttrTypes defines the attribute types for the DhcpInterfaces model.
func (o DhcpInterfaces) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":   basetypes.StringType{},
		"device": basetypes.StringType{},
		"folder": basetypes.StringType{},
		"id":     basetypes.StringType{},
		"name":   basetypes.StringType{},
		"relay": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ip": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enabled": basetypes.BoolType{},
						"server":  basetypes.ListType{ElemType: basetypes.StringType{}},
					},
				},
			},
		},
		"server": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ip_pool": basetypes.ListType{ElemType: basetypes.StringType{}},
				"mode":    basetypes.StringType{},
				"option": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"dns": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"primary":   basetypes.StringType{},
								"secondary": basetypes.StringType{},
							},
						},
						"dns_suffix": basetypes.StringType{},
						"gateway":    basetypes.StringType{},
						"inheritance": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"source": basetypes.StringType{},
							},
						},
						"lease": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"timeout": basetypes.Int64Type{},
								"unlimited": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
						"nis": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"primary":   basetypes.StringType{},
								"secondary": basetypes.StringType{},
							},
						},
						"ntp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"primary":   basetypes.StringType{},
								"secondary": basetypes.StringType{},
							},
						},
						"pop3_server": basetypes.StringType{},
						"smtp_server": basetypes.StringType{},
						"subnet_mask": basetypes.StringType{},
						"user_defined": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"ascii":     basetypes.ListType{ElemType: basetypes.StringType{}},
								"code":      basetypes.Int64Type{},
								"hex":       basetypes.ListType{ElemType: basetypes.StringType{}},
								"inherited": basetypes.BoolType{},
								"ip":        basetypes.ListType{ElemType: basetypes.StringType{}},
								"name":      basetypes.StringType{},
							},
						}},
						"wins": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"primary":   basetypes.StringType{},
								"secondary": basetypes.StringType{},
							},
						},
					},
				},
				"probe_ip": basetypes.BoolType{},
				"reserved": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"description": basetypes.StringType{},
						"mac":         basetypes.StringType{},
						"name":        basetypes.StringType{},
					},
				}},
			},
		},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DhcpInterfaces objects.
func (o DhcpInterfaces) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DhcpInterfacesRelay model.
func (o DhcpInterfacesRelay) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ip": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enabled": basetypes.BoolType{},
				"server":  basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		},
	}
}

// AttrType returns the attribute type for a list of DhcpInterfacesRelay objects.
func (o DhcpInterfacesRelay) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DhcpInterfacesRelayIp model.
func (o DhcpInterfacesRelayIp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enabled": basetypes.BoolType{},
		"server":  basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of DhcpInterfacesRelayIp objects.
func (o DhcpInterfacesRelayIp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DhcpInterfacesServer model.
func (o DhcpInterfacesServer) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ip_pool": basetypes.ListType{ElemType: basetypes.StringType{}},
		"mode":    basetypes.StringType{},
		"option": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"dns": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"primary":   basetypes.StringType{},
						"secondary": basetypes.StringType{},
					},
				},
				"dns_suffix": basetypes.StringType{},
				"gateway":    basetypes.StringType{},
				"inheritance": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"source": basetypes.StringType{},
					},
				},
				"lease": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"timeout": basetypes.Int64Type{},
						"unlimited": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
				"nis": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"primary":   basetypes.StringType{},
						"secondary": basetypes.StringType{},
					},
				},
				"ntp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"primary":   basetypes.StringType{},
						"secondary": basetypes.StringType{},
					},
				},
				"pop3_server": basetypes.StringType{},
				"smtp_server": basetypes.StringType{},
				"subnet_mask": basetypes.StringType{},
				"user_defined": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ascii":     basetypes.ListType{ElemType: basetypes.StringType{}},
						"code":      basetypes.Int64Type{},
						"hex":       basetypes.ListType{ElemType: basetypes.StringType{}},
						"inherited": basetypes.BoolType{},
						"ip":        basetypes.ListType{ElemType: basetypes.StringType{}},
						"name":      basetypes.StringType{},
					},
				}},
				"wins": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"primary":   basetypes.StringType{},
						"secondary": basetypes.StringType{},
					},
				},
			},
		},
		"probe_ip": basetypes.BoolType{},
		"reserved": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"description": basetypes.StringType{},
				"mac":         basetypes.StringType{},
				"name":        basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of DhcpInterfacesServer objects.
func (o DhcpInterfacesServer) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DhcpInterfacesServerOption model.
func (o DhcpInterfacesServerOption) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"dns": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"primary":   basetypes.StringType{},
				"secondary": basetypes.StringType{},
			},
		},
		"dns_suffix": basetypes.StringType{},
		"gateway":    basetypes.StringType{},
		"inheritance": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"source": basetypes.StringType{},
			},
		},
		"lease": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"timeout": basetypes.Int64Type{},
				"unlimited": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
		"nis": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"primary":   basetypes.StringType{},
				"secondary": basetypes.StringType{},
			},
		},
		"ntp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"primary":   basetypes.StringType{},
				"secondary": basetypes.StringType{},
			},
		},
		"pop3_server": basetypes.StringType{},
		"smtp_server": basetypes.StringType{},
		"subnet_mask": basetypes.StringType{},
		"user_defined": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ascii":     basetypes.ListType{ElemType: basetypes.StringType{}},
				"code":      basetypes.Int64Type{},
				"hex":       basetypes.ListType{ElemType: basetypes.StringType{}},
				"inherited": basetypes.BoolType{},
				"ip":        basetypes.ListType{ElemType: basetypes.StringType{}},
				"name":      basetypes.StringType{},
			},
		}},
		"wins": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"primary":   basetypes.StringType{},
				"secondary": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of DhcpInterfacesServerOption objects.
func (o DhcpInterfacesServerOption) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DhcpInterfacesServerOptionDns model.
func (o DhcpInterfacesServerOptionDns) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"primary":   basetypes.StringType{},
		"secondary": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DhcpInterfacesServerOptionDns objects.
func (o DhcpInterfacesServerOptionDns) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DhcpInterfacesServerOptionInheritance model.
func (o DhcpInterfacesServerOptionInheritance) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"source": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DhcpInterfacesServerOptionInheritance objects.
func (o DhcpInterfacesServerOptionInheritance) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DhcpInterfacesServerOptionLease model.
func (o DhcpInterfacesServerOptionLease) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"timeout": basetypes.Int64Type{},
		"unlimited": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of DhcpInterfacesServerOptionLease objects.
func (o DhcpInterfacesServerOptionLease) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DhcpInterfacesServerOptionNis model.
func (o DhcpInterfacesServerOptionNis) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"primary":   basetypes.StringType{},
		"secondary": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DhcpInterfacesServerOptionNis objects.
func (o DhcpInterfacesServerOptionNis) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DhcpInterfacesServerOptionNtp model.
func (o DhcpInterfacesServerOptionNtp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"primary":   basetypes.StringType{},
		"secondary": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DhcpInterfacesServerOptionNtp objects.
func (o DhcpInterfacesServerOptionNtp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DhcpInterfacesServerOptionUserDefinedInner model.
func (o DhcpInterfacesServerOptionUserDefinedInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ascii":     basetypes.ListType{ElemType: basetypes.StringType{}},
		"code":      basetypes.Int64Type{},
		"hex":       basetypes.ListType{ElemType: basetypes.StringType{}},
		"inherited": basetypes.BoolType{},
		"ip":        basetypes.ListType{ElemType: basetypes.StringType{}},
		"name":      basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DhcpInterfacesServerOptionUserDefinedInner objects.
func (o DhcpInterfacesServerOptionUserDefinedInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DhcpInterfacesServerOptionWins model.
func (o DhcpInterfacesServerOptionWins) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"primary":   basetypes.StringType{},
		"secondary": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DhcpInterfacesServerOptionWins objects.
func (o DhcpInterfacesServerOptionWins) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DhcpInterfacesServerReservedInner model.
func (o DhcpInterfacesServerReservedInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"description": basetypes.StringType{},
		"mac":         basetypes.StringType{},
		"name":        basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DhcpInterfacesServerReservedInner objects.
func (o DhcpInterfacesServerReservedInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// DhcpInterfacesResourceSchema defines the schema for DhcpInterfaces resource
var DhcpInterfacesResourceSchema = schema.Schema{
	MarkdownDescription: "DhcpInterface resource",
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
			MarkdownDescription: "The device in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
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
		"name": schema.StringAttribute{
			MarkdownDescription: "Interface name",
			Required:            true,
		},
		"relay": schema.SingleNestedAttribute{
			Validators: []validator.Object{
				objectvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("server"),
				),
			},
			MarkdownDescription: "Relay\n> ℹ️ **Note:** You must specify exactly one of `relay` and `server`.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"ip": schema.SingleNestedAttribute{
					MarkdownDescription: "Ip",
					Required:            true,
					Attributes: map[string]schema.Attribute{
						"enabled": schema.BoolAttribute{
							MarkdownDescription: "Enabled?",
							Required:            true,
						},
						"server": schema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "Server",
							Required:            true,
						},
					},
				},
			},
		},
		"server": schema.SingleNestedAttribute{
			Validators: []validator.Object{
				objectvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("relay"),
				),
			},
			MarkdownDescription: "Server\n> ℹ️ **Note:** You must specify exactly one of `relay` and `server`.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"ip_pool": schema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "List of IP address pools",
					Optional:            true,
				},
				"mode": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("auto", "enabled", "disabled"),
					},
					MarkdownDescription: "DHCP server mode",
					Optional:            true,
				},
				"option": schema.SingleNestedAttribute{
					MarkdownDescription: "Option",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"dns": schema.SingleNestedAttribute{
							MarkdownDescription: "Dns",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"primary": schema.StringAttribute{
									MarkdownDescription: "Primary DNS server",
									Optional:            true,
								},
								"secondary": schema.StringAttribute{
									MarkdownDescription: "Secondary DNS server",
									Optional:            true,
								},
							},
						},
						"dns_suffix": schema.StringAttribute{
							MarkdownDescription: "DNS suffix",
							Optional:            true,
						},
						"gateway": schema.StringAttribute{
							MarkdownDescription: "Default gateway",
							Optional:            true,
						},
						"inheritance": schema.SingleNestedAttribute{
							MarkdownDescription: "Inheritance",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"source": schema.StringAttribute{
									MarkdownDescription: "Interface from which to inherit lease options",
									Optional:            true,
								},
							},
						},
						"lease": schema.SingleNestedAttribute{
							MarkdownDescription: "Lease",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"timeout": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("unlimited"),
										),
										int64validator.Between(0, 1000000),
									},
									MarkdownDescription: "DHCP lease timeout (minutes)\n> ℹ️ **Note:** You must specify exactly one of `timeout` and `unlimited`.",
									Optional:            true,
								},
								"unlimited": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ConflictsWith(
											path.MatchRelative().AtParent().AtName("timeout"),
										),
									},
									MarkdownDescription: "Unlimited\n> ℹ️ **Note:** You must specify exactly one of `timeout` and `unlimited`.",
									Optional:            true,
									Attributes:          map[string]schema.Attribute{},
								},
							},
						},
						"nis": schema.SingleNestedAttribute{
							MarkdownDescription: "Nis",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"primary": schema.StringAttribute{
									MarkdownDescription: "Primary NIS server",
									Optional:            true,
								},
								"secondary": schema.StringAttribute{
									MarkdownDescription: "Secondary NIS server",
									Optional:            true,
								},
							},
						},
						"ntp": schema.SingleNestedAttribute{
							MarkdownDescription: "Ntp",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"primary": schema.StringAttribute{
									MarkdownDescription: "Primary NTP server",
									Optional:            true,
								},
								"secondary": schema.StringAttribute{
									MarkdownDescription: "Secondary NTP server",
									Optional:            true,
								},
							},
						},
						"pop3_server": schema.StringAttribute{
							MarkdownDescription: "POP3 server",
							Optional:            true,
						},
						"smtp_server": schema.StringAttribute{
							MarkdownDescription: "SMTP server",
							Optional:            true,
						},
						"subnet_mask": schema.StringAttribute{
							MarkdownDescription: "Subnet mask",
							Optional:            true,
						},
						"user_defined": schema.ListNestedAttribute{
							MarkdownDescription: "Custom DHCP options",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ascii": schema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "Ascii",
										Optional:            true,
									},
									"code": schema.Int64Attribute{
										Validators: []validator.Int64{
											int64validator.Between(1, 254),
										},
										MarkdownDescription: "Option code",
										Optional:            true,
									},
									"hex": schema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "Hex",
										Optional:            true,
									},
									"inherited": schema.BoolAttribute{
										MarkdownDescription: "Inherited from DHCP server inheritance source?",
										Required:            true,
									},
									"ip": schema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "Ip",
										Optional:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: "Option name",
										Required:            true,
									},
								},
							},
						},
						"wins": schema.SingleNestedAttribute{
							MarkdownDescription: "Wins",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"primary": schema.StringAttribute{
									MarkdownDescription: "Primary WINS server",
									Optional:            true,
								},
								"secondary": schema.StringAttribute{
									MarkdownDescription: "Secondary WINS server",
									Optional:            true,
								},
							},
						},
					},
				},
				"probe_ip": schema.BoolAttribute{
					MarkdownDescription: "Ping IP before allocating?",
					Optional:            true,
				},
				"reserved": schema.ListNestedAttribute{
					MarkdownDescription: "List of IP reservations",
					Optional:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"description": schema.StringAttribute{
								MarkdownDescription: "Reservation description",
								Optional:            true,
							},
							"mac": schema.StringAttribute{
								MarkdownDescription: "Reserved MAC address",
								Optional:            true,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: "Reserved IP address",
								Optional:            true,
							},
						},
					},
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
	},
}

// DhcpInterfacesDataSourceSchema defines the schema for DhcpInterfaces data source
var DhcpInterfacesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "DhcpInterface data source",
	Attributes: map[string]dsschema.Attribute{
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
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
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Interface name",
			Optional:            true,
			Computed:            true,
		},
		"relay": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Relay\n> ℹ️ **Note:** You must specify exactly one of `relay` and `server`.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"ip": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Ip",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"enabled": dsschema.BoolAttribute{
							MarkdownDescription: "Enabled?",
							Computed:            true,
						},
						"server": dsschema.ListAttribute{
							ElementType:         types.StringType,
							MarkdownDescription: "Server",
							Computed:            true,
						},
					},
				},
			},
		},
		"server": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Server\n> ℹ️ **Note:** You must specify exactly one of `relay` and `server`.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"ip_pool": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "List of IP address pools",
					Computed:            true,
				},
				"mode": dsschema.StringAttribute{
					MarkdownDescription: "DHCP server mode",
					Computed:            true,
				},
				"option": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Option",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"dns": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Dns",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"primary": dsschema.StringAttribute{
									MarkdownDescription: "Primary DNS server",
									Computed:            true,
								},
								"secondary": dsschema.StringAttribute{
									MarkdownDescription: "Secondary DNS server",
									Computed:            true,
								},
							},
						},
						"dns_suffix": dsschema.StringAttribute{
							MarkdownDescription: "DNS suffix",
							Computed:            true,
						},
						"gateway": dsschema.StringAttribute{
							MarkdownDescription: "Default gateway",
							Computed:            true,
						},
						"inheritance": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Inheritance",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"source": dsschema.StringAttribute{
									MarkdownDescription: "Interface from which to inherit lease options",
									Computed:            true,
								},
							},
						},
						"lease": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Lease",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"timeout": dsschema.Int64Attribute{
									MarkdownDescription: "DHCP lease timeout (minutes)\n> ℹ️ **Note:** You must specify exactly one of `timeout` and `unlimited`.",
									Computed:            true,
								},
								"unlimited": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Unlimited\n> ℹ️ **Note:** You must specify exactly one of `timeout` and `unlimited`.",
									Computed:            true,
									Attributes:          map[string]dsschema.Attribute{},
								},
							},
						},
						"nis": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Nis",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"primary": dsschema.StringAttribute{
									MarkdownDescription: "Primary NIS server",
									Computed:            true,
								},
								"secondary": dsschema.StringAttribute{
									MarkdownDescription: "Secondary NIS server",
									Computed:            true,
								},
							},
						},
						"ntp": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Ntp",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"primary": dsschema.StringAttribute{
									MarkdownDescription: "Primary NTP server",
									Computed:            true,
								},
								"secondary": dsschema.StringAttribute{
									MarkdownDescription: "Secondary NTP server",
									Computed:            true,
								},
							},
						},
						"pop3_server": dsschema.StringAttribute{
							MarkdownDescription: "POP3 server",
							Computed:            true,
						},
						"smtp_server": dsschema.StringAttribute{
							MarkdownDescription: "SMTP server",
							Computed:            true,
						},
						"subnet_mask": dsschema.StringAttribute{
							MarkdownDescription: "Subnet mask",
							Computed:            true,
						},
						"user_defined": dsschema.ListNestedAttribute{
							MarkdownDescription: "Custom DHCP options",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"ascii": dsschema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "Ascii",
										Computed:            true,
									},
									"code": dsschema.Int64Attribute{
										MarkdownDescription: "Option code",
										Computed:            true,
									},
									"hex": dsschema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "Hex",
										Computed:            true,
									},
									"inherited": dsschema.BoolAttribute{
										MarkdownDescription: "Inherited from DHCP server inheritance source?",
										Computed:            true,
									},
									"ip": dsschema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "Ip",
										Computed:            true,
									},
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Option name",
										Computed:            true,
									},
								},
							},
						},
						"wins": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Wins",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"primary": dsschema.StringAttribute{
									MarkdownDescription: "Primary WINS server",
									Computed:            true,
								},
								"secondary": dsschema.StringAttribute{
									MarkdownDescription: "Secondary WINS server",
									Computed:            true,
								},
							},
						},
					},
				},
				"probe_ip": dsschema.BoolAttribute{
					MarkdownDescription: "Ping IP before allocating?",
					Computed:            true,
				},
				"reserved": dsschema.ListNestedAttribute{
					MarkdownDescription: "List of IP reservations",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"description": dsschema.StringAttribute{
								MarkdownDescription: "Reservation description",
								Computed:            true,
							},
							"mac": dsschema.StringAttribute{
								MarkdownDescription: "Reserved MAC address",
								Computed:            true,
							},
							"name": dsschema.StringAttribute{
								MarkdownDescription: "Reserved IP address",
								Computed:            true,
							},
						},
					},
				},
			},
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
	},
}

// DhcpInterfacesListModel represents the data model for a list data source.
type DhcpInterfacesListModel struct {
	Tfid    types.String     `tfsdk:"tfid"`
	Data    []DhcpInterfaces `tfsdk:"data"`
	Limit   types.Int64      `tfsdk:"limit"`
	Offset  types.Int64      `tfsdk:"offset"`
	Name    types.String     `tfsdk:"name"`
	Total   types.Int64      `tfsdk:"total"`
	Folder  types.String     `tfsdk:"folder"`
	Snippet types.String     `tfsdk:"snippet"`
	Device  types.String     `tfsdk:"device"`
}

// DhcpInterfacesListDataSourceSchema defines the schema for a list data source.
var DhcpInterfacesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: DhcpInterfacesDataSourceSchema.Attributes,
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
