package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
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

// DnsProxies represents the Terraform model for DnsProxies
type DnsProxies struct {
	Tfid          types.String          `tfsdk:"tfid"`
	Cache         basetypes.ObjectValue `tfsdk:"cache"`
	Default       basetypes.ObjectValue `tfsdk:"default"`
	Device        basetypes.StringValue `tfsdk:"device"`
	DomainServers basetypes.ListValue   `tfsdk:"domain_servers"`
	Enabled       basetypes.BoolValue   `tfsdk:"enabled"`
	Folder        basetypes.StringValue `tfsdk:"folder"`
	Id            basetypes.StringValue `tfsdk:"id"`
	Interface     basetypes.ListValue   `tfsdk:"interface"`
	Name          basetypes.StringValue `tfsdk:"name"`
	Snippet       basetypes.StringValue `tfsdk:"snippet"`
	StaticEntries basetypes.ListValue   `tfsdk:"static_entries"`
	TcpQueries    basetypes.ObjectValue `tfsdk:"tcp_queries"`
	UdpQueries    basetypes.ObjectValue `tfsdk:"udp_queries"`
}

// DnsProxiesCache represents a nested structure within the DnsProxies model
type DnsProxiesCache struct {
	CacheEdns basetypes.BoolValue   `tfsdk:"cache_edns"`
	Enabled   basetypes.BoolValue   `tfsdk:"enabled"`
	MaxTtl    basetypes.ObjectValue `tfsdk:"max_ttl"`
}

// DnsProxiesCacheMaxTtl represents a nested structure within the DnsProxies model
type DnsProxiesCacheMaxTtl struct {
	Enabled    basetypes.BoolValue  `tfsdk:"enabled"`
	TimeToLive basetypes.Int64Value `tfsdk:"time_to_live"`
}

// DnsProxiesDefault represents a nested structure within the DnsProxies model
type DnsProxiesDefault struct {
	Inheritance basetypes.ObjectValue `tfsdk:"inheritance"`
	Primary     basetypes.StringValue `tfsdk:"primary"`
	Secondary   basetypes.StringValue `tfsdk:"secondary"`
}

// DnsProxiesDefaultInheritance represents a nested structure within the DnsProxies model
type DnsProxiesDefaultInheritance struct {
	Source basetypes.StringValue `tfsdk:"source"`
}

// DnsProxiesDomainServersInner represents a nested structure within the DnsProxies model
type DnsProxiesDomainServersInner struct {
	Cacheable  basetypes.BoolValue   `tfsdk:"cacheable"`
	DomainName basetypes.ListValue   `tfsdk:"domain_name"`
	Name       basetypes.StringValue `tfsdk:"name"`
	Primary    basetypes.StringValue `tfsdk:"primary"`
	Secondary  basetypes.StringValue `tfsdk:"secondary"`
}

// DnsProxiesStaticEntriesInner represents a nested structure within the DnsProxies model
type DnsProxiesStaticEntriesInner struct {
	Address basetypes.ListValue   `tfsdk:"address"`
	Domain  basetypes.StringValue `tfsdk:"domain"`
	Name    basetypes.StringValue `tfsdk:"name"`
}

// DnsProxiesTcpQueries represents a nested structure within the DnsProxies model
type DnsProxiesTcpQueries struct {
	Enabled            basetypes.BoolValue  `tfsdk:"enabled"`
	MaxPendingRequests basetypes.Int64Value `tfsdk:"max_pending_requests"`
}

// DnsProxiesUdpQueries represents a nested structure within the DnsProxies model
type DnsProxiesUdpQueries struct {
	Retries basetypes.ObjectValue `tfsdk:"retries"`
}

// DnsProxiesUdpQueriesRetries represents a nested structure within the DnsProxies model
type DnsProxiesUdpQueriesRetries struct {
	Attempts basetypes.Int64Value `tfsdk:"attempts"`
	Interval basetypes.Int64Value `tfsdk:"interval"`
}

// AttrTypes defines the attribute types for the DnsProxies model.
func (o DnsProxies) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"cache": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"cache_edns": basetypes.BoolType{},
				"enabled":    basetypes.BoolType{},
				"max_ttl": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enabled":      basetypes.BoolType{},
						"time_to_live": basetypes.Int64Type{},
					},
				},
			},
		},
		"default": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"inheritance": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"source": basetypes.StringType{},
					},
				},
				"primary":   basetypes.StringType{},
				"secondary": basetypes.StringType{},
			},
		},
		"device": basetypes.StringType{},
		"domain_servers": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"cacheable":   basetypes.BoolType{},
				"domain_name": basetypes.ListType{ElemType: basetypes.StringType{}},
				"name":        basetypes.StringType{},
				"primary":     basetypes.StringType{},
				"secondary":   basetypes.StringType{},
			},
		}},
		"enabled":   basetypes.BoolType{},
		"folder":    basetypes.StringType{},
		"id":        basetypes.StringType{},
		"interface": basetypes.ListType{ElemType: basetypes.StringType{}},
		"name":      basetypes.StringType{},
		"snippet":   basetypes.StringType{},
		"static_entries": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address": basetypes.ListType{ElemType: basetypes.StringType{}},
				"domain":  basetypes.StringType{},
				"name":    basetypes.StringType{},
			},
		}},
		"tcp_queries": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enabled":              basetypes.BoolType{},
				"max_pending_requests": basetypes.Int64Type{},
			},
		},
		"udp_queries": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"retries": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"attempts": basetypes.Int64Type{},
						"interval": basetypes.Int64Type{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of DnsProxies objects.
func (o DnsProxies) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DnsProxiesCache model.
func (o DnsProxiesCache) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"cache_edns": basetypes.BoolType{},
		"enabled":    basetypes.BoolType{},
		"max_ttl": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enabled":      basetypes.BoolType{},
				"time_to_live": basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of DnsProxiesCache objects.
func (o DnsProxiesCache) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DnsProxiesCacheMaxTtl model.
func (o DnsProxiesCacheMaxTtl) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enabled":      basetypes.BoolType{},
		"time_to_live": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of DnsProxiesCacheMaxTtl objects.
func (o DnsProxiesCacheMaxTtl) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DnsProxiesDefault model.
func (o DnsProxiesDefault) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"inheritance": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"source": basetypes.StringType{},
			},
		},
		"primary":   basetypes.StringType{},
		"secondary": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DnsProxiesDefault objects.
func (o DnsProxiesDefault) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DnsProxiesDefaultInheritance model.
func (o DnsProxiesDefaultInheritance) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"source": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DnsProxiesDefaultInheritance objects.
func (o DnsProxiesDefaultInheritance) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DnsProxiesDomainServersInner model.
func (o DnsProxiesDomainServersInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"cacheable":   basetypes.BoolType{},
		"domain_name": basetypes.ListType{ElemType: basetypes.StringType{}},
		"name":        basetypes.StringType{},
		"primary":     basetypes.StringType{},
		"secondary":   basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DnsProxiesDomainServersInner objects.
func (o DnsProxiesDomainServersInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DnsProxiesStaticEntriesInner model.
func (o DnsProxiesStaticEntriesInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address": basetypes.ListType{ElemType: basetypes.StringType{}},
		"domain":  basetypes.StringType{},
		"name":    basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DnsProxiesStaticEntriesInner objects.
func (o DnsProxiesStaticEntriesInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DnsProxiesTcpQueries model.
func (o DnsProxiesTcpQueries) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enabled":              basetypes.BoolType{},
		"max_pending_requests": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of DnsProxiesTcpQueries objects.
func (o DnsProxiesTcpQueries) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DnsProxiesUdpQueries model.
func (o DnsProxiesUdpQueries) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"retries": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"attempts": basetypes.Int64Type{},
				"interval": basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of DnsProxiesUdpQueries objects.
func (o DnsProxiesUdpQueries) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DnsProxiesUdpQueriesRetries model.
func (o DnsProxiesUdpQueriesRetries) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"attempts": basetypes.Int64Type{},
		"interval": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of DnsProxiesUdpQueriesRetries objects.
func (o DnsProxiesUdpQueriesRetries) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// DnsProxiesResourceSchema defines the schema for DnsProxies resource
var DnsProxiesResourceSchema = schema.Schema{
	MarkdownDescription: "DnsProxy resource",
	Attributes: map[string]schema.Attribute{
		"cache": schema.SingleNestedAttribute{
			MarkdownDescription: "Cache",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"cache_edns": schema.BoolAttribute{
					MarkdownDescription: "Cache EDNS UDP response",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"enabled": schema.BoolAttribute{
					MarkdownDescription: "Turn on caching for this DNS object",
					Required:            true,
				},
				"max_ttl": schema.SingleNestedAttribute{
					MarkdownDescription: "Max ttl",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"enabled": schema.BoolAttribute{
							MarkdownDescription: "Enable max ttl for this DNS object",
							Required:            true,
						},
						"time_to_live": schema.Int64Attribute{
							Validators: []validator.Int64{
								int64validator.Between(60, 86400),
							},
							MarkdownDescription: "Time in seconds after which entry is cleared",
							Optional:            true,
							Computed:            true,
						},
					},
				},
			},
		},
		"default": schema.SingleNestedAttribute{
			MarkdownDescription: "Default",
			Required:            true,
			Attributes: map[string]schema.Attribute{
				"inheritance": schema.SingleNestedAttribute{
					MarkdownDescription: "Inheritance",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"source": schema.StringAttribute{
							MarkdownDescription: "Dynamic interface",
							Optional:            true,
						},
					},
				},
				"primary": schema.StringAttribute{
					MarkdownDescription: "Primary DNS Name server IP address",
					Required:            true,
				},
				"secondary": schema.StringAttribute{
					MarkdownDescription: "Secondary DNS Name server IP address",
					Optional:            true,
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
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
			},
			MarkdownDescription: "The device in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"domain_servers": schema.ListNestedAttribute{
			MarkdownDescription: "DNS proxy rules",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"cacheable": schema.BoolAttribute{
						MarkdownDescription: "Enable caching for this DNS proxy rule?",
						Optional:            true,
					},
					"domain_name": schema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "Domain names(s) that will be matched",
						Validators: []validator.List{
							listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(128)),
						},
						Optional: true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "Proxy rule name",
						Required:            true,
					},
					"primary": schema.StringAttribute{
						MarkdownDescription: "Primary DNS server IP address",
						Required:            true,
					},
					"secondary": schema.StringAttribute{
						MarkdownDescription: "Secondary DNS server IP address",
						Optional:            true,
					},
				},
			},
		},
		"enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable DNS proxy?",
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
			MarkdownDescription: "UUID of the resource",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"interface": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Interfaces on which to enable DNS proxy service",
			Optional:            true,
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(31),
			},
			MarkdownDescription: "DNS proxy name",
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
		"static_entries": schema.ListNestedAttribute{
			MarkdownDescription: "Static entries",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"address": schema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "Address",
						Validators: []validator.List{
							listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(63)),
						},
						Required: true,
					},
					"domain": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.LengthAtMost(255),
						},
						MarkdownDescription: "Fully qualified domain name",
						Required:            true,
					},
					"name": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.LengthAtMost(31),
						},
						MarkdownDescription: "Static entry name",
						Required:            true,
					},
				},
			},
		},
		"tcp_queries": schema.SingleNestedAttribute{
			MarkdownDescription: "Tcp queries",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"enabled": schema.BoolAttribute{
					MarkdownDescription: "Turn on forwarding of TCP DNS queries?",
					Required:            true,
				},
				"max_pending_requests": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(64, 256),
					},
					MarkdownDescription: "Upper limit on number of concurrent TCP DNS requests",
					Optional:            true,
					Computed:            true,
					Default:             int64default.StaticInt64(64),
				},
			},
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"udp_queries": schema.SingleNestedAttribute{
			MarkdownDescription: "Udp queries",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"retries": schema.SingleNestedAttribute{
					MarkdownDescription: "Retries",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"attempts": schema.Int64Attribute{
							Validators: []validator.Int64{
								int64validator.Between(1, 30),
							},
							MarkdownDescription: "Maximum number of retries before trying next name server",
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(5),
						},
						"interval": schema.Int64Attribute{
							Validators: []validator.Int64{
								int64validator.Between(1, 30),
							},
							MarkdownDescription: "Time in seconds for another request to be sent",
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(2),
						},
					},
				},
			},
		},
	},
}

// DnsProxiesDataSourceSchema defines the schema for DnsProxies data source
var DnsProxiesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "DnsProxy data source",
	Attributes: map[string]dsschema.Attribute{
		"cache": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Cache",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"cache_edns": dsschema.BoolAttribute{
					MarkdownDescription: "Cache EDNS UDP response",
					Computed:            true,
				},
				"enabled": dsschema.BoolAttribute{
					MarkdownDescription: "Turn on caching for this DNS object",
					Computed:            true,
				},
				"max_ttl": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Max ttl",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"enabled": dsschema.BoolAttribute{
							MarkdownDescription: "Enable max ttl for this DNS object",
							Computed:            true,
						},
						"time_to_live": dsschema.Int64Attribute{
							MarkdownDescription: "Time in seconds after which entry is cleared",
							Computed:            true,
						},
					},
				},
			},
		},
		"default": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Default",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"inheritance": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Inheritance",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"source": dsschema.StringAttribute{
							MarkdownDescription: "Dynamic interface",
							Computed:            true,
						},
					},
				},
				"primary": dsschema.StringAttribute{
					MarkdownDescription: "Primary DNS Name server IP address",
					Computed:            true,
				},
				"secondary": dsschema.StringAttribute{
					MarkdownDescription: "Secondary DNS Name server IP address",
					Computed:            true,
				},
			},
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"domain_servers": dsschema.ListNestedAttribute{
			MarkdownDescription: "DNS proxy rules",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"cacheable": dsschema.BoolAttribute{
						MarkdownDescription: "Enable caching for this DNS proxy rule?",
						Computed:            true,
					},
					"domain_name": dsschema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "Domain names(s) that will be matched",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "Proxy rule name",
						Computed:            true,
					},
					"primary": dsschema.StringAttribute{
						MarkdownDescription: "Primary DNS server IP address",
						Computed:            true,
					},
					"secondary": dsschema.StringAttribute{
						MarkdownDescription: "Secondary DNS server IP address",
						Computed:            true,
					},
				},
			},
		},
		"enabled": dsschema.BoolAttribute{
			MarkdownDescription: "Enable DNS proxy?",
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
		"interface": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Interfaces on which to enable DNS proxy service",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "DNS proxy name",
			Optional:            true,
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"static_entries": dsschema.ListNestedAttribute{
			MarkdownDescription: "Static entries",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"address": dsschema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "Address",
						Computed:            true,
					},
					"domain": dsschema.StringAttribute{
						MarkdownDescription: "Fully qualified domain name",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "Static entry name",
						Computed:            true,
					},
				},
			},
		},
		"tcp_queries": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Tcp queries",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"enabled": dsschema.BoolAttribute{
					MarkdownDescription: "Turn on forwarding of TCP DNS queries?",
					Computed:            true,
				},
				"max_pending_requests": dsschema.Int64Attribute{
					MarkdownDescription: "Upper limit on number of concurrent TCP DNS requests",
					Computed:            true,
				},
			},
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"udp_queries": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Udp queries",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"retries": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Retries",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"attempts": dsschema.Int64Attribute{
							MarkdownDescription: "Maximum number of retries before trying next name server",
							Computed:            true,
						},
						"interval": dsschema.Int64Attribute{
							MarkdownDescription: "Time in seconds for another request to be sent",
							Computed:            true,
						},
					},
				},
			},
		},
	},
}

// DnsProxiesListModel represents the data model for a list data source.
type DnsProxiesListModel struct {
	Tfid    types.String `tfsdk:"tfid"`
	Data    []DnsProxies `tfsdk:"data"`
	Limit   types.Int64  `tfsdk:"limit"`
	Offset  types.Int64  `tfsdk:"offset"`
	Name    types.String `tfsdk:"name"`
	Total   types.Int64  `tfsdk:"total"`
	Folder  types.String `tfsdk:"folder"`
	Snippet types.String `tfsdk:"snippet"`
	Device  types.String `tfsdk:"device"`
}

// DnsProxiesListDataSourceSchema defines the schema for a list data source.
var DnsProxiesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: DnsProxiesDataSourceSchema.Attributes,
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
