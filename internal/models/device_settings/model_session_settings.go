package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
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

// SessionSettings represents the Terraform model for SessionSettings
type SessionSettings struct {
	Tfid            types.String          `tfsdk:"tfid"`
	Device          basetypes.StringValue `tfsdk:"device"`
	Folder          basetypes.StringValue `tfsdk:"folder"`
	Id              basetypes.StringValue `tfsdk:"id"`
	SessionSettings basetypes.ObjectValue `tfsdk:"session_settings"`
	Snippet         basetypes.StringValue `tfsdk:"snippet"`
}

// SessionSettingsSessionSettings represents a nested structure within the SessionSettings model
type SessionSettingsSessionSettings struct {
	AcceleratedAgingEnable                      basetypes.BoolValue    `tfsdk:"accelerated_aging_enable"`
	AcceleratedAgingScalingFactor               basetypes.Float64Value `tfsdk:"accelerated_aging_scaling_factor"`
	AcceleratedAgingThreshold                   basetypes.Float64Value `tfsdk:"accelerated_aging_threshold"`
	Config                                      basetypes.ObjectValue  `tfsdk:"config"`
	DhcpBcastSessionOn                          basetypes.BoolValue    `tfsdk:"dhcp_bcast_session_on"`
	Erspan                                      basetypes.BoolValue    `tfsdk:"erspan"`
	IcmpUnreachableRate                         basetypes.Float64Value `tfsdk:"icmp_unreachable_rate"`
	Icmpv6RateLimit                             basetypes.ObjectValue  `tfsdk:"icmpv6_rate_limit"`
	Ipv6Firewalling                             basetypes.BoolValue    `tfsdk:"ipv6_firewalling"`
	JumboFrame                                  basetypes.ObjectValue  `tfsdk:"jumbo_frame"`
	MaxPendingMcastPktsPerSession               basetypes.Float64Value `tfsdk:"max_pending_mcast_pkts_per_session"`
	MulticastRouteSetupBuffering                basetypes.BoolValue    `tfsdk:"multicast_route_setup_buffering"`
	Nat                                         basetypes.ObjectValue  `tfsdk:"nat"`
	Nat64                                       basetypes.ObjectValue  `tfsdk:"nat64"`
	PacketBufferProtectionActivate              basetypes.Float64Value `tfsdk:"packet_buffer_protection_activate"`
	PacketBufferProtectionAlert                 basetypes.Int64Value   `tfsdk:"packet_buffer_protection_alert"`
	PacketBufferProtectionBlockCountdown        basetypes.Float64Value `tfsdk:"packet_buffer_protection_block_countdown"`
	PacketBufferProtectionBlockDurationTime     basetypes.Float64Value `tfsdk:"packet_buffer_protection_block_duration_time"`
	PacketBufferProtectionBlockHoldTime         basetypes.Float64Value `tfsdk:"packet_buffer_protection_block_hold_time"`
	PacketBufferProtectionEnable                basetypes.BoolValue    `tfsdk:"packet_buffer_protection_enable"`
	PacketBufferProtectionLatencyActivate       basetypes.Float64Value `tfsdk:"packet_buffer_protection_latency_activate"`
	PacketBufferProtectionLatencyAlert          basetypes.Float64Value `tfsdk:"packet_buffer_protection_latency_alert"`
	PacketBufferProtectionLatencyBlockCountdown basetypes.Float64Value `tfsdk:"packet_buffer_protection_latency_block_countdown"`
	PacketBufferProtectionLatencyMaxTolerate    basetypes.Float64Value `tfsdk:"packet_buffer_protection_latency_max_tolerate"`
	PacketBufferProtectionMonitorOnly           basetypes.BoolValue    `tfsdk:"packet_buffer_protection_monitor_only"`
	PacketBufferProtectionUseLatency            basetypes.BoolValue    `tfsdk:"packet_buffer_protection_use_latency"`
}

// SessionSettingsSessionSettingsConfig represents a nested structure within the SessionSettings model
type SessionSettingsSessionSettingsConfig struct {
	Rematch basetypes.BoolValue `tfsdk:"rematch"`
}

// SessionSettingsSessionSettingsIcmpv6RateLimit represents a nested structure within the SessionSettings model
type SessionSettingsSessionSettingsIcmpv6RateLimit struct {
	BucketSize basetypes.Int64Value `tfsdk:"bucket_size"`
	PacketRate basetypes.Int64Value `tfsdk:"packet_rate"`
}

// SessionSettingsSessionSettingsJumboFrame represents a nested structure within the SessionSettings model
type SessionSettingsSessionSettingsJumboFrame struct {
	Mtu basetypes.Int64Value `tfsdk:"mtu"`
}

// SessionSettingsSessionSettingsNat represents a nested structure within the SessionSettings model
type SessionSettingsSessionSettingsNat struct {
	DippOversub basetypes.StringValue `tfsdk:"dipp_oversub"`
}

// SessionSettingsSessionSettingsNat64 represents a nested structure within the SessionSettings model
type SessionSettingsSessionSettingsNat64 struct {
	Ipv6MinNetworkMtu basetypes.Int64Value `tfsdk:"ipv6_min_network_mtu"`
}

// AttrTypes defines the attribute types for the SessionSettings model.
func (o SessionSettings) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":   basetypes.StringType{},
		"device": basetypes.StringType{},
		"folder": basetypes.StringType{},
		"id":     basetypes.StringType{},
		"session_settings": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"accelerated_aging_enable":         basetypes.BoolType{},
				"accelerated_aging_scaling_factor": basetypes.Float64Type{},
				"accelerated_aging_threshold":      basetypes.Float64Type{},
				"config": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"rematch": basetypes.BoolType{},
					},
				},
				"dhcp_bcast_session_on": basetypes.BoolType{},
				"erspan":                basetypes.BoolType{},
				"icmp_unreachable_rate": basetypes.Float64Type{},
				"icmpv6_rate_limit": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"bucket_size": basetypes.Int64Type{},
						"packet_rate": basetypes.Int64Type{},
					},
				},
				"ipv6_firewalling": basetypes.BoolType{},
				"jumbo_frame": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"mtu": basetypes.Int64Type{},
					},
				},
				"max_pending_mcast_pkts_per_session": basetypes.Float64Type{},
				"multicast_route_setup_buffering":    basetypes.BoolType{},
				"nat": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"dipp_oversub": basetypes.StringType{},
					},
				},
				"nat64": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ipv6_min_network_mtu": basetypes.Int64Type{},
					},
				},
				"packet_buffer_protection_activate":                basetypes.Float64Type{},
				"packet_buffer_protection_alert":                   basetypes.Int64Type{},
				"packet_buffer_protection_block_countdown":         basetypes.Float64Type{},
				"packet_buffer_protection_block_duration_time":     basetypes.Float64Type{},
				"packet_buffer_protection_block_hold_time":         basetypes.Float64Type{},
				"packet_buffer_protection_enable":                  basetypes.BoolType{},
				"packet_buffer_protection_latency_activate":        basetypes.Float64Type{},
				"packet_buffer_protection_latency_alert":           basetypes.Float64Type{},
				"packet_buffer_protection_latency_block_countdown": basetypes.Float64Type{},
				"packet_buffer_protection_latency_max_tolerate":    basetypes.Float64Type{},
				"packet_buffer_protection_monitor_only":            basetypes.BoolType{},
				"packet_buffer_protection_use_latency":             basetypes.BoolType{},
			},
		},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of SessionSettings objects.
func (o SessionSettings) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SessionSettingsSessionSettings model.
func (o SessionSettingsSessionSettings) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"accelerated_aging_enable":         basetypes.BoolType{},
		"accelerated_aging_scaling_factor": basetypes.Float64Type{},
		"accelerated_aging_threshold":      basetypes.Float64Type{},
		"config": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"rematch": basetypes.BoolType{},
			},
		},
		"dhcp_bcast_session_on": basetypes.BoolType{},
		"erspan":                basetypes.BoolType{},
		"icmp_unreachable_rate": basetypes.Float64Type{},
		"icmpv6_rate_limit": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"bucket_size": basetypes.Int64Type{},
				"packet_rate": basetypes.Int64Type{},
			},
		},
		"ipv6_firewalling": basetypes.BoolType{},
		"jumbo_frame": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"mtu": basetypes.Int64Type{},
			},
		},
		"max_pending_mcast_pkts_per_session": basetypes.Float64Type{},
		"multicast_route_setup_buffering":    basetypes.BoolType{},
		"nat": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"dipp_oversub": basetypes.StringType{},
			},
		},
		"nat64": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ipv6_min_network_mtu": basetypes.Int64Type{},
			},
		},
		"packet_buffer_protection_activate":                basetypes.Float64Type{},
		"packet_buffer_protection_alert":                   basetypes.Int64Type{},
		"packet_buffer_protection_block_countdown":         basetypes.Float64Type{},
		"packet_buffer_protection_block_duration_time":     basetypes.Float64Type{},
		"packet_buffer_protection_block_hold_time":         basetypes.Float64Type{},
		"packet_buffer_protection_enable":                  basetypes.BoolType{},
		"packet_buffer_protection_latency_activate":        basetypes.Float64Type{},
		"packet_buffer_protection_latency_alert":           basetypes.Float64Type{},
		"packet_buffer_protection_latency_block_countdown": basetypes.Float64Type{},
		"packet_buffer_protection_latency_max_tolerate":    basetypes.Float64Type{},
		"packet_buffer_protection_monitor_only":            basetypes.BoolType{},
		"packet_buffer_protection_use_latency":             basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of SessionSettingsSessionSettings objects.
func (o SessionSettingsSessionSettings) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SessionSettingsSessionSettingsConfig model.
func (o SessionSettingsSessionSettingsConfig) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"rematch": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of SessionSettingsSessionSettingsConfig objects.
func (o SessionSettingsSessionSettingsConfig) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SessionSettingsSessionSettingsIcmpv6RateLimit model.
func (o SessionSettingsSessionSettingsIcmpv6RateLimit) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"bucket_size": basetypes.Int64Type{},
		"packet_rate": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of SessionSettingsSessionSettingsIcmpv6RateLimit objects.
func (o SessionSettingsSessionSettingsIcmpv6RateLimit) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SessionSettingsSessionSettingsJumboFrame model.
func (o SessionSettingsSessionSettingsJumboFrame) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"mtu": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of SessionSettingsSessionSettingsJumboFrame objects.
func (o SessionSettingsSessionSettingsJumboFrame) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SessionSettingsSessionSettingsNat model.
func (o SessionSettingsSessionSettingsNat) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"dipp_oversub": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of SessionSettingsSessionSettingsNat objects.
func (o SessionSettingsSessionSettingsNat) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SessionSettingsSessionSettingsNat64 model.
func (o SessionSettingsSessionSettingsNat64) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ipv6_min_network_mtu": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of SessionSettingsSessionSettingsNat64 objects.
func (o SessionSettingsSessionSettingsNat64) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// SessionSettingsResourceSchema defines the schema for SessionSettings resource
var SessionSettingsResourceSchema = schema.Schema{
	MarkdownDescription: "SessionSetting resource",
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
		"session_settings": schema.SingleNestedAttribute{
			MarkdownDescription: "Session settings",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"accelerated_aging_enable": schema.BoolAttribute{
					MarkdownDescription: "Enable accelerated aging",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"accelerated_aging_scaling_factor": schema.Float64Attribute{
					Validators: []validator.Float64{
						float64validator.Between(2.000000, 16.000000),
					},
					MarkdownDescription: "Accelerated aging scaling factor",
					Optional:            true,
					Computed:            true,
					Default:             float64default.StaticFloat64(2.000000),
				},
				"accelerated_aging_threshold": schema.Float64Attribute{
					Validators: []validator.Float64{
						float64validator.Between(50.000000, 99.000000),
					},
					MarkdownDescription: "Accelerated aging threshold",
					Optional:            true,
					Computed:            true,
					Default:             float64default.StaticFloat64(80.000000),
				},
				"config": schema.SingleNestedAttribute{
					MarkdownDescription: "Config",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"rematch": schema.BoolAttribute{
							MarkdownDescription: "Rematch all sessions on config policy change",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
					},
				},
				"dhcp_bcast_session_on": schema.BoolAttribute{
					MarkdownDescription: "Enable DHCP broadcast session",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"erspan": schema.BoolAttribute{
					MarkdownDescription: "Enable ERSPAN support",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"icmp_unreachable_rate": schema.Float64Attribute{
					Validators: []validator.Float64{
						float64validator.Between(1.000000, 65535.000000),
					},
					MarkdownDescription: "ICMP unreachable packet rate (per second)",
					Optional:            true,
					Computed:            true,
					Default:             float64default.StaticFloat64(200.000000),
				},
				"icmpv6_rate_limit": schema.SingleNestedAttribute{
					MarkdownDescription: "ICMPv6 rate limiting",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"bucket_size": schema.Int64Attribute{
							Validators: []validator.Int64{
								int64validator.Between(10, 65535),
							},
							MarkdownDescription: "ICMPv6 token bucket size",
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(100),
						},
						"packet_rate": schema.Int64Attribute{
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
							MarkdownDescription: "ICMPv6 error packet pate (per second)",
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(100),
						},
					},
				},
				"ipv6_firewalling": schema.BoolAttribute{
					MarkdownDescription: "Enable IPv6 firewalling",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"jumbo_frame": schema.SingleNestedAttribute{
					MarkdownDescription: "Enable jumbo frame support",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"mtu": schema.Int64Attribute{
							Validators: []validator.Int64{
								int64validator.Between(512, 9216),
							},
							MarkdownDescription: "Global MTU",
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(9192),
						},
					},
				},
				"max_pending_mcast_pkts_per_session": schema.Float64Attribute{
					Validators: []validator.Float64{
						float64validator.Between(1.000000, 2000.000000),
					},
					MarkdownDescription: "Multicast route setup buffer size",
					Optional:            true,
					Computed:            true,
					Default:             float64default.StaticFloat64(1000.000000),
				},
				"multicast_route_setup_buffering": schema.BoolAttribute{
					MarkdownDescription: "Multicast route setup buffering",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"nat": schema.SingleNestedAttribute{
					MarkdownDescription: "Nat",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"dipp_oversub": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("1x", "2x", "4x", "8x"),
							},
							MarkdownDescription: "NAT oversubscription rate",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("1x"),
						},
					},
				},
				"nat64": schema.SingleNestedAttribute{
					MarkdownDescription: "Nat64",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"ipv6_min_network_mtu": schema.Int64Attribute{
							Validators: []validator.Int64{
								int64validator.Between(1280, 9216),
							},
							MarkdownDescription: "NAT64 IPv6 minimum network MTU",
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(1280),
						},
					},
				},
				"packet_buffer_protection_activate": schema.Float64Attribute{
					Validators: []validator.Float64{
						float64validator.Between(0.000000, 99.000000),
					},
					MarkdownDescription: "Activate (%)",
					Optional:            true,
					Computed:            true,
					Default:             float64default.StaticFloat64(80.000000),
				},
				"packet_buffer_protection_alert": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(0, 99),
					},
					MarkdownDescription: "Alert (%)",
					Optional:            true,
					Computed:            true,
					Default:             int64default.StaticInt64(50),
				},
				"packet_buffer_protection_block_countdown": schema.Float64Attribute{
					Validators: []validator.Float64{
						float64validator.Between(0.000000, 99.000000),
					},
					MarkdownDescription: "Block countdown threshold (%)",
					Optional:            true,
					Computed:            true,
					Default:             float64default.StaticFloat64(80.000000),
				},
				"packet_buffer_protection_block_duration_time": schema.Float64Attribute{
					Validators: []validator.Float64{
						float64validator.Between(1.000000, 15999999.000000),
					},
					MarkdownDescription: "Block duration (seconds)",
					Optional:            true,
					Computed:            true,
					Default:             float64default.StaticFloat64(3600.000000),
				},
				"packet_buffer_protection_block_hold_time": schema.Float64Attribute{
					Validators: []validator.Float64{
						float64validator.Between(0.000000, 65535.000000),
					},
					MarkdownDescription: "Block hold time (seconds)",
					Optional:            true,
					Computed:            true,
					Default:             float64default.StaticFloat64(60.000000),
				},
				"packet_buffer_protection_enable": schema.BoolAttribute{
					MarkdownDescription: "Enable packet buffer protection",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"packet_buffer_protection_latency_activate": schema.Float64Attribute{
					Validators: []validator.Float64{
						float64validator.Between(1.000000, 20000.000000),
					},
					MarkdownDescription: "Latency activate (milliseconds)",
					Optional:            true,
					Computed:            true,
					Default:             float64default.StaticFloat64(200.000000),
				},
				"packet_buffer_protection_latency_alert": schema.Float64Attribute{
					Validators: []validator.Float64{
						float64validator.Between(1.000000, 20000.000000),
					},
					MarkdownDescription: "Latency alert (milliseconds)",
					Optional:            true,
					Computed:            true,
					Default:             float64default.StaticFloat64(50.000000),
				},
				"packet_buffer_protection_latency_block_countdown": schema.Float64Attribute{
					Validators: []validator.Float64{
						float64validator.Between(1.000000, 20000.000000),
					},
					MarkdownDescription: "Block countdown threshold (milliseconds)",
					Optional:            true,
					Computed:            true,
					Default:             float64default.StaticFloat64(500.000000),
				},
				"packet_buffer_protection_latency_max_tolerate": schema.Float64Attribute{
					Validators: []validator.Float64{
						float64validator.Between(1.000000, 20000.000000),
					},
					MarkdownDescription: "Latency max tolerate (milliseconds)",
					Optional:            true,
					Computed:            true,
					Default:             float64default.StaticFloat64(500.000000),
				},
				"packet_buffer_protection_monitor_only": schema.BoolAttribute{
					MarkdownDescription: "Packet buffer protection monitor only",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"packet_buffer_protection_use_latency": schema.BoolAttribute{
					MarkdownDescription: "Enabled latency-based activation",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
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

// SessionSettingsDataSourceSchema defines the schema for SessionSettings data source
var SessionSettingsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "SessionSetting data source",
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
		"session_settings": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Session settings",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"accelerated_aging_enable": dsschema.BoolAttribute{
					MarkdownDescription: "Enable accelerated aging",
					Computed:            true,
				},
				"accelerated_aging_scaling_factor": dsschema.Float64Attribute{
					MarkdownDescription: "Accelerated aging scaling factor",
					Computed:            true,
				},
				"accelerated_aging_threshold": dsschema.Float64Attribute{
					MarkdownDescription: "Accelerated aging threshold",
					Computed:            true,
				},
				"config": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Config",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"rematch": dsschema.BoolAttribute{
							MarkdownDescription: "Rematch all sessions on config policy change",
							Computed:            true,
						},
					},
				},
				"dhcp_bcast_session_on": dsschema.BoolAttribute{
					MarkdownDescription: "Enable DHCP broadcast session",
					Computed:            true,
				},
				"erspan": dsschema.BoolAttribute{
					MarkdownDescription: "Enable ERSPAN support",
					Computed:            true,
				},
				"icmp_unreachable_rate": dsschema.Float64Attribute{
					MarkdownDescription: "ICMP unreachable packet rate (per second)",
					Computed:            true,
				},
				"icmpv6_rate_limit": dsschema.SingleNestedAttribute{
					MarkdownDescription: "ICMPv6 rate limiting",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"bucket_size": dsschema.Int64Attribute{
							MarkdownDescription: "ICMPv6 token bucket size",
							Computed:            true,
						},
						"packet_rate": dsschema.Int64Attribute{
							MarkdownDescription: "ICMPv6 error packet pate (per second)",
							Computed:            true,
						},
					},
				},
				"ipv6_firewalling": dsschema.BoolAttribute{
					MarkdownDescription: "Enable IPv6 firewalling",
					Computed:            true,
				},
				"jumbo_frame": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Enable jumbo frame support",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"mtu": dsschema.Int64Attribute{
							MarkdownDescription: "Global MTU",
							Computed:            true,
						},
					},
				},
				"max_pending_mcast_pkts_per_session": dsschema.Float64Attribute{
					MarkdownDescription: "Multicast route setup buffer size",
					Computed:            true,
				},
				"multicast_route_setup_buffering": dsschema.BoolAttribute{
					MarkdownDescription: "Multicast route setup buffering",
					Computed:            true,
				},
				"nat": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Nat",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"dipp_oversub": dsschema.StringAttribute{
							MarkdownDescription: "NAT oversubscription rate",
							Computed:            true,
						},
					},
				},
				"nat64": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Nat64",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"ipv6_min_network_mtu": dsschema.Int64Attribute{
							MarkdownDescription: "NAT64 IPv6 minimum network MTU",
							Computed:            true,
						},
					},
				},
				"packet_buffer_protection_activate": dsschema.Float64Attribute{
					MarkdownDescription: "Activate (%)",
					Computed:            true,
				},
				"packet_buffer_protection_alert": dsschema.Int64Attribute{
					MarkdownDescription: "Alert (%)",
					Computed:            true,
				},
				"packet_buffer_protection_block_countdown": dsschema.Float64Attribute{
					MarkdownDescription: "Block countdown threshold (%)",
					Computed:            true,
				},
				"packet_buffer_protection_block_duration_time": dsschema.Float64Attribute{
					MarkdownDescription: "Block duration (seconds)",
					Computed:            true,
				},
				"packet_buffer_protection_block_hold_time": dsschema.Float64Attribute{
					MarkdownDescription: "Block hold time (seconds)",
					Computed:            true,
				},
				"packet_buffer_protection_enable": dsschema.BoolAttribute{
					MarkdownDescription: "Enable packet buffer protection",
					Computed:            true,
				},
				"packet_buffer_protection_latency_activate": dsschema.Float64Attribute{
					MarkdownDescription: "Latency activate (milliseconds)",
					Computed:            true,
				},
				"packet_buffer_protection_latency_alert": dsschema.Float64Attribute{
					MarkdownDescription: "Latency alert (milliseconds)",
					Computed:            true,
				},
				"packet_buffer_protection_latency_block_countdown": dsschema.Float64Attribute{
					MarkdownDescription: "Block countdown threshold (milliseconds)",
					Computed:            true,
				},
				"packet_buffer_protection_latency_max_tolerate": dsschema.Float64Attribute{
					MarkdownDescription: "Latency max tolerate (milliseconds)",
					Computed:            true,
				},
				"packet_buffer_protection_monitor_only": dsschema.BoolAttribute{
					MarkdownDescription: "Packet buffer protection monitor only",
					Computed:            true,
				},
				"packet_buffer_protection_use_latency": dsschema.BoolAttribute{
					MarkdownDescription: "Enabled latency-based activation",
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

// SessionSettingsListModel represents the data model for a list data source.
type SessionSettingsListModel struct {
	Tfid    types.String      `tfsdk:"tfid"`
	Data    []SessionSettings `tfsdk:"data"`
	Limit   types.Int64       `tfsdk:"limit"`
	Offset  types.Int64       `tfsdk:"offset"`
	Name    types.String      `tfsdk:"name"`
	Total   types.Int64       `tfsdk:"total"`
	Folder  types.String      `tfsdk:"folder"`
	Snippet types.String      `tfsdk:"snippet"`
	Device  types.String      `tfsdk:"device"`
}

// SessionSettingsListDataSourceSchema defines the schema for a list data source.
var SessionSettingsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: SessionSettingsDataSourceSchema.Attributes,
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
