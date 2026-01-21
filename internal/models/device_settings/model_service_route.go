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

// Package: device_settings
// This file contains models for the device_settings SDK package

// ServiceRoute represents the Terraform model for ServiceRoute
type ServiceRoute struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Device  basetypes.StringValue `tfsdk:"device"`
	Folder  basetypes.StringValue `tfsdk:"folder"`
	Id      basetypes.StringValue `tfsdk:"id"`
	Route   basetypes.ObjectValue `tfsdk:"route"`
	Snippet basetypes.StringValue `tfsdk:"snippet"`
}

// ServiceRouteRoute represents a nested structure within the ServiceRoute model
type ServiceRouteRoute struct {
	Destination basetypes.ListValue `tfsdk:"destination"`
	Service     basetypes.ListValue `tfsdk:"service"`
}

// ServiceRouteRouteDestinationInner represents a nested structure within the ServiceRoute model
type ServiceRouteRouteDestinationInner struct {
	Name   basetypes.StringValue `tfsdk:"name"`
	Source basetypes.ObjectValue `tfsdk:"source"`
}

// ServiceRouteRouteDestinationInnerSource represents a nested structure within the ServiceRoute model
type ServiceRouteRouteDestinationInnerSource struct {
	Address   basetypes.StringValue `tfsdk:"address"`
	Interface basetypes.StringValue `tfsdk:"interface"`
}

// ServiceRouteRouteServiceInner represents a nested structure within the ServiceRoute model
type ServiceRouteRouteServiceInner struct {
	Name     basetypes.StringValue `tfsdk:"name"`
	Source   basetypes.ObjectValue `tfsdk:"source"`
	SourceV6 basetypes.ObjectValue `tfsdk:"source_v6"`
}

// ServiceRouteRouteServiceInnerSource represents a nested structure within the ServiceRoute model
type ServiceRouteRouteServiceInnerSource struct {
	Address   basetypes.StringValue `tfsdk:"address"`
	Interface basetypes.StringValue `tfsdk:"interface"`
}

// ServiceRouteRouteServiceInnerSourceV6 represents a nested structure within the ServiceRoute model
type ServiceRouteRouteServiceInnerSourceV6 struct {
	Address   basetypes.StringValue `tfsdk:"address"`
	Interface basetypes.StringValue `tfsdk:"interface"`
}

// AttrTypes defines the attribute types for the ServiceRoute model.
func (o ServiceRoute) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":   basetypes.StringType{},
		"device": basetypes.StringType{},
		"folder": basetypes.StringType{},
		"id":     basetypes.StringType{},
		"route": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"destination": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name": basetypes.StringType{},
						"source": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address":   basetypes.StringType{},
								"interface": basetypes.StringType{},
							},
						},
					},
				}},
				"service": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name": basetypes.StringType{},
						"source": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address":   basetypes.StringType{},
								"interface": basetypes.StringType{},
							},
						},
						"source_v6": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address":   basetypes.StringType{},
								"interface": basetypes.StringType{},
							},
						},
					},
				}},
			},
		},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ServiceRoute objects.
func (o ServiceRoute) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ServiceRouteRoute model.
func (o ServiceRouteRoute) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"destination": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name": basetypes.StringType{},
				"source": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address":   basetypes.StringType{},
						"interface": basetypes.StringType{},
					},
				},
			},
		}},
		"service": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name": basetypes.StringType{},
				"source": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address":   basetypes.StringType{},
						"interface": basetypes.StringType{},
					},
				},
				"source_v6": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address":   basetypes.StringType{},
						"interface": basetypes.StringType{},
					},
				},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of ServiceRouteRoute objects.
func (o ServiceRouteRoute) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ServiceRouteRouteDestinationInner model.
func (o ServiceRouteRouteDestinationInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": basetypes.StringType{},
		"source": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address":   basetypes.StringType{},
				"interface": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ServiceRouteRouteDestinationInner objects.
func (o ServiceRouteRouteDestinationInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ServiceRouteRouteDestinationInnerSource model.
func (o ServiceRouteRouteDestinationInnerSource) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address":   basetypes.StringType{},
		"interface": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ServiceRouteRouteDestinationInnerSource objects.
func (o ServiceRouteRouteDestinationInnerSource) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ServiceRouteRouteServiceInner model.
func (o ServiceRouteRouteServiceInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": basetypes.StringType{},
		"source": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address":   basetypes.StringType{},
				"interface": basetypes.StringType{},
			},
		},
		"source_v6": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address":   basetypes.StringType{},
				"interface": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ServiceRouteRouteServiceInner objects.
func (o ServiceRouteRouteServiceInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ServiceRouteRouteServiceInnerSource model.
func (o ServiceRouteRouteServiceInnerSource) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address":   basetypes.StringType{},
		"interface": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ServiceRouteRouteServiceInnerSource objects.
func (o ServiceRouteRouteServiceInnerSource) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ServiceRouteRouteServiceInnerSourceV6 model.
func (o ServiceRouteRouteServiceInnerSourceV6) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address":   basetypes.StringType{},
		"interface": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ServiceRouteRouteServiceInnerSourceV6 objects.
func (o ServiceRouteRouteServiceInnerSourceV6) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// ServiceRouteResourceSchema defines the schema for ServiceRoute resource
var ServiceRouteResourceSchema = schema.Schema{
	MarkdownDescription: "ServiceRoute resource",
	Attributes: map[string]schema.Attribute{
		"device": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("folder"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
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
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
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
		"route": schema.SingleNestedAttribute{
			MarkdownDescription: "Route",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"destination": schema.ListNestedAttribute{
					MarkdownDescription: "Destination",
					Optional:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								MarkdownDescription: "Name",
								Optional:            true,
							},
							"source": schema.SingleNestedAttribute{
								MarkdownDescription: "Source",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: "Address",
										Optional:            true,
									},
									"interface": schema.StringAttribute{
										MarkdownDescription: "Interface",
										Optional:            true,
									},
								},
							},
						},
					},
				},
				"service": schema.ListNestedAttribute{
					MarkdownDescription: "Service",
					Optional:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("autofocus", "crl-status", "data-services", "ddns", "deployments", "dns", "edl-updates", "email", "hsm", "http", "iot", "kerberos", "ldap", "mdm", "mfa", "netflow", "ntp", "paloalto-networks-services", "panorama", "panorama-log-forwarding", "proxy", "radius", "scep", "snmp", "syslog", "tacplus", "uid-agent", "url-updates", "vmmonitor", "wildfire-private", "ztp"),
								},
								MarkdownDescription: "The follow list details the accepted `name` values and their corresponding service description.\n- `autofocus` = AutoFocus Cloud\n- `crl-status` = CRL servers\n- `data-services` = Data Services\n- `ddns` = DDNS server(s)\n- `deployments` = Panorama pushed updates\n- `dns` = DNS server(s)\n- `edl-updates` = External Dynamic List update server\n- `email` = SMTP gateway(s)\n- `hsm` = Hardware Security Module server(s)\n- `http` = HTTP Forwarding server(s)\n- `iot` = IOT service-route\n- `kerberos` = Kerberos server\n- `ldap` = LDAP server\n- `mdm` = MDM servers\n- `mfa` = Multi-Factor Authentication\n- `netflow` = Netflow server(s)\n- `ntp` = NTP server(s)\n- `paloalto-networks-services` = Palo Alto Networks Services\n- `panorama` = Panorama server\n- `panorama-log-forwarding` = Panorama Log Forwarding\n- `proxy` = Proxy server\n- `radius` = RADIUS server\n- `scep` = SCEP\n- `snmp` = SNMP server(s)\n- `syslog` = Syslog server(s)\n- `tacplus` = TACACS+ server\n- `uid-`agent = UID agent(s)\n- `url-`updates = URL update server\n- `vmmonitor` = VM monitor\n- `wildfire-`private = WildFire Appliance\n- `ztp` = ZTP and Auto-VPN DDNS\n",
								Optional:            true,
							},
							"source": schema.SingleNestedAttribute{
								MarkdownDescription: "Source",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: "Address",
										Optional:            true,
									},
									"interface": schema.StringAttribute{
										MarkdownDescription: "Interface",
										Optional:            true,
									},
								},
							},
							"source_v6": schema.SingleNestedAttribute{
								MarkdownDescription: "Source v6",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: "Address",
										Optional:            true,
									},
									"interface": schema.StringAttribute{
										MarkdownDescription: "Interface",
										Optional:            true,
									},
								},
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
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
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

// ServiceRouteDataSourceSchema defines the schema for ServiceRoute data source
var ServiceRouteDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "ServiceRoute data source",
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
		"route": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Route",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"destination": dsschema.ListNestedAttribute{
					MarkdownDescription: "Destination",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"name": dsschema.StringAttribute{
								MarkdownDescription: "Name",
								Computed:            true,
							},
							"source": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Source",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"address": dsschema.StringAttribute{
										MarkdownDescription: "Address",
										Computed:            true,
									},
									"interface": dsschema.StringAttribute{
										MarkdownDescription: "Interface",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"service": dsschema.ListNestedAttribute{
					MarkdownDescription: "Service",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"name": dsschema.StringAttribute{
								MarkdownDescription: "The follow list details the accepted `name` values and their corresponding service description.\n- `autofocus` = AutoFocus Cloud\n- `crl-status` = CRL servers\n- `data-services` = Data Services\n- `ddns` = DDNS server(s)\n- `deployments` = Panorama pushed updates\n- `dns` = DNS server(s)\n- `edl-updates` = External Dynamic List update server\n- `email` = SMTP gateway(s)\n- `hsm` = Hardware Security Module server(s)\n- `http` = HTTP Forwarding server(s)\n- `iot` = IOT service-route\n- `kerberos` = Kerberos server\n- `ldap` = LDAP server\n- `mdm` = MDM servers\n- `mfa` = Multi-Factor Authentication\n- `netflow` = Netflow server(s)\n- `ntp` = NTP server(s)\n- `paloalto-networks-services` = Palo Alto Networks Services\n- `panorama` = Panorama server\n- `panorama-log-forwarding` = Panorama Log Forwarding\n- `proxy` = Proxy server\n- `radius` = RADIUS server\n- `scep` = SCEP\n- `snmp` = SNMP server(s)\n- `syslog` = Syslog server(s)\n- `tacplus` = TACACS+ server\n- `uid-`agent = UID agent(s)\n- `url-`updates = URL update server\n- `vmmonitor` = VM monitor\n- `wildfire-`private = WildFire Appliance\n- `ztp` = ZTP and Auto-VPN DDNS\n",
								Computed:            true,
							},
							"source": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Source",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"address": dsschema.StringAttribute{
										MarkdownDescription: "Address",
										Computed:            true,
									},
									"interface": dsschema.StringAttribute{
										MarkdownDescription: "Interface",
										Computed:            true,
									},
								},
							},
							"source_v6": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Source v6",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"address": dsschema.StringAttribute{
										MarkdownDescription: "Address",
										Computed:            true,
									},
									"interface": dsschema.StringAttribute{
										MarkdownDescription: "Interface",
										Computed:            true,
									},
								},
							},
						},
					},
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

// ServiceRouteListModel represents the data model for a list data source.
type ServiceRouteListModel struct {
	Tfid    types.String   `tfsdk:"tfid"`
	Data    []ServiceRoute `tfsdk:"data"`
	Limit   types.Int64    `tfsdk:"limit"`
	Offset  types.Int64    `tfsdk:"offset"`
	Name    types.String   `tfsdk:"name"`
	Total   types.Int64    `tfsdk:"total"`
	Folder  types.String   `tfsdk:"folder"`
	Snippet types.String   `tfsdk:"snippet"`
	Device  types.String   `tfsdk:"device"`
}

// ServiceRouteListDataSourceSchema defines the schema for a list data source.
var ServiceRouteListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: ServiceRouteDataSourceSchema.Attributes,
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
