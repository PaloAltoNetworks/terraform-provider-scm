package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: network_services
// This file contains models for the network_services SDK package

// ZoneProtectionProfiles represents the Terraform model for ZoneProtectionProfiles
type ZoneProtectionProfiles struct {
	Tfid                                   types.String          `tfsdk:"tfid"`
	AsymmetricPath                         basetypes.StringValue `tfsdk:"asymmetric_path"`
	Description                            basetypes.StringValue `tfsdk:"description"`
	Device                                 basetypes.StringValue `tfsdk:"device"`
	DiscardIcmpEmbeddedError               basetypes.BoolValue   `tfsdk:"discard_icmp_embedded_error"`
	Flood                                  basetypes.ObjectValue `tfsdk:"flood"`
	Folder                                 basetypes.StringValue `tfsdk:"folder"`
	FragmentedTrafficDiscard               basetypes.BoolValue   `tfsdk:"fragmented_traffic_discard"`
	IcmpFragDiscard                        basetypes.BoolValue   `tfsdk:"icmp_frag_discard"`
	IcmpLargePacketDiscard                 basetypes.BoolValue   `tfsdk:"icmp_large_packet_discard"`
	IcmpPingZeroIdDiscard                  basetypes.BoolValue   `tfsdk:"icmp_ping_zero_id_discard"`
	Id                                     basetypes.StringValue `tfsdk:"id"`
	Ipv6                                   basetypes.ObjectValue `tfsdk:"ipv6"`
	L2SecGroupTagProtection                basetypes.ObjectValue `tfsdk:"l2_sec_group_tag_protection"`
	LooseSourceRoutingDiscard              basetypes.BoolValue   `tfsdk:"loose_source_routing_discard"`
	MalformedOptionDiscard                 basetypes.BoolValue   `tfsdk:"malformed_option_discard"`
	MismatchedOverlappingTcpSegmentDiscard basetypes.BoolValue   `tfsdk:"mismatched_overlapping_tcp_segment_discard"`
	MptcpOptionStrip                       basetypes.StringValue `tfsdk:"mptcp_option_strip"`
	Name                                   basetypes.StringValue `tfsdk:"name"`
	NonIpProtocol                          basetypes.ObjectValue `tfsdk:"non_ip_protocol"`
	RecordRouteDiscard                     basetypes.BoolValue   `tfsdk:"record_route_discard"`
	RejectNonSynTcp                        basetypes.StringValue `tfsdk:"reject_non_syn_tcp"`
	Scan                                   basetypes.ListValue   `tfsdk:"scan"`
	ScanWhiteList                          basetypes.ListValue   `tfsdk:"scan_white_list"`
	SecurityDiscard                        basetypes.BoolValue   `tfsdk:"security_discard"`
	Snippet                                basetypes.StringValue `tfsdk:"snippet"`
	SpoofedIpDiscard                       basetypes.BoolValue   `tfsdk:"spoofed_ip_discard"`
	StreamIdDiscard                        basetypes.BoolValue   `tfsdk:"stream_id_discard"`
	StrictIpCheck                          basetypes.BoolValue   `tfsdk:"strict_ip_check"`
	StrictSourceRoutingDiscard             basetypes.BoolValue   `tfsdk:"strict_source_routing_discard"`
	SuppressIcmpNeedfrag                   basetypes.BoolValue   `tfsdk:"suppress_icmp_needfrag"`
	SuppressIcmpTimeexceeded               basetypes.BoolValue   `tfsdk:"suppress_icmp_timeexceeded"`
	TcpFastOpenAndDataStrip                basetypes.BoolValue   `tfsdk:"tcp_fast_open_and_data_strip"`
	TcpHandshakeDiscard                    basetypes.BoolValue   `tfsdk:"tcp_handshake_discard"`
	TcpSynWithDataDiscard                  basetypes.BoolValue   `tfsdk:"tcp_syn_with_data_discard"`
	TcpSynackWithDataDiscard               basetypes.BoolValue   `tfsdk:"tcp_synack_with_data_discard"`
	TcpTimestampStrip                      basetypes.BoolValue   `tfsdk:"tcp_timestamp_strip"`
	TimestampDiscard                       basetypes.BoolValue   `tfsdk:"timestamp_discard"`
	UnknownOptionDiscard                   basetypes.BoolValue   `tfsdk:"unknown_option_discard"`
}

// ZoneProtectionProfilesFlood represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesFlood struct {
	Icmp     basetypes.ObjectValue `tfsdk:"icmp"`
	Icmpv6   basetypes.ObjectValue `tfsdk:"icmpv6"`
	OtherIp  basetypes.ObjectValue `tfsdk:"other_ip"`
	SctpInit basetypes.ObjectValue `tfsdk:"sctp_init"`
	TcpSyn   basetypes.ObjectValue `tfsdk:"tcp_syn"`
	Udp      basetypes.ObjectValue `tfsdk:"udp"`
}

// ZoneProtectionProfilesFloodIcmp represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesFloodIcmp struct {
	Enable basetypes.BoolValue   `tfsdk:"enable"`
	Red    basetypes.ObjectValue `tfsdk:"red"`
}

// ZoneProtectionProfilesFloodIcmpRed represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesFloodIcmpRed struct {
	ActivateRate basetypes.Int64Value `tfsdk:"activate_rate"`
	AlarmRate    basetypes.Int64Value `tfsdk:"alarm_rate"`
	MaximalRate  basetypes.Int64Value `tfsdk:"maximal_rate"`
}

// ZoneProtectionProfilesFloodIcmpv6 represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesFloodIcmpv6 struct {
	Enable basetypes.BoolValue   `tfsdk:"enable"`
	Red    basetypes.ObjectValue `tfsdk:"red"`
}

// ZoneProtectionProfilesFloodIcmpv6Red represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesFloodIcmpv6Red struct {
	ActivateRate basetypes.Int64Value `tfsdk:"activate_rate"`
	AlarmRate    basetypes.Int64Value `tfsdk:"alarm_rate"`
	MaximalRate  basetypes.Int64Value `tfsdk:"maximal_rate"`
}

// ZoneProtectionProfilesFloodOtherIp represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesFloodOtherIp struct {
	Enable basetypes.BoolValue   `tfsdk:"enable"`
	Red    basetypes.ObjectValue `tfsdk:"red"`
}

// ZoneProtectionProfilesFloodOtherIpRed represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesFloodOtherIpRed struct {
	ActivateRate basetypes.Int64Value `tfsdk:"activate_rate"`
	AlarmRate    basetypes.Int64Value `tfsdk:"alarm_rate"`
	MaximalRate  basetypes.Int64Value `tfsdk:"maximal_rate"`
}

// ZoneProtectionProfilesFloodSctpInit represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesFloodSctpInit struct {
	Enable basetypes.BoolValue   `tfsdk:"enable"`
	Red    basetypes.ObjectValue `tfsdk:"red"`
}

// ZoneProtectionProfilesFloodSctpInitRed represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesFloodSctpInitRed struct {
	ActivateRate basetypes.Int64Value `tfsdk:"activate_rate"`
	AlarmRate    basetypes.Int64Value `tfsdk:"alarm_rate"`
	MaximalRate  basetypes.Int64Value `tfsdk:"maximal_rate"`
}

// ZoneProtectionProfilesFloodTcpSyn represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesFloodTcpSyn struct {
	Enable     basetypes.BoolValue   `tfsdk:"enable"`
	Red        basetypes.ObjectValue `tfsdk:"red"`
	SynCookies basetypes.ObjectValue `tfsdk:"syn_cookies"`
}

// ZoneProtectionProfilesFloodTcpSynRed represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesFloodTcpSynRed struct {
	ActivateRate basetypes.Int64Value `tfsdk:"activate_rate"`
	AlarmRate    basetypes.Int64Value `tfsdk:"alarm_rate"`
	MaximalRate  basetypes.Int64Value `tfsdk:"maximal_rate"`
}

// ZoneProtectionProfilesFloodTcpSynSynCookies represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesFloodTcpSynSynCookies struct {
	ActivateRate basetypes.Int64Value `tfsdk:"activate_rate"`
	AlarmRate    basetypes.Int64Value `tfsdk:"alarm_rate"`
	MaximalRate  basetypes.Int64Value `tfsdk:"maximal_rate"`
}

// ZoneProtectionProfilesFloodUdp represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesFloodUdp struct {
	Enable basetypes.BoolValue   `tfsdk:"enable"`
	Red    basetypes.ObjectValue `tfsdk:"red"`
}

// ZoneProtectionProfilesFloodUdpRed represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesFloodUdpRed struct {
	ActivateRate basetypes.Int64Value `tfsdk:"activate_rate"`
	AlarmRate    basetypes.Int64Value `tfsdk:"alarm_rate"`
	MaximalRate  basetypes.Int64Value `tfsdk:"maximal_rate"`
}

// ZoneProtectionProfilesIpv6 represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesIpv6 struct {
	AnycastSource               basetypes.BoolValue   `tfsdk:"anycast_source"`
	FilterExtHdr                basetypes.ObjectValue `tfsdk:"filter_ext_hdr"`
	Icmpv6TooBigSmallMtuDiscard basetypes.BoolValue   `tfsdk:"icmpv6_too_big_small_mtu_discard"`
	IgnoreInvPkt                basetypes.ObjectValue `tfsdk:"ignore_inv_pkt"`
	Ipv4CompatibleAddress       basetypes.BoolValue   `tfsdk:"ipv4_compatible_address"`
	NeedlessFragmentHdr         basetypes.BoolValue   `tfsdk:"needless_fragment_hdr"`
	OptionsInvalidIpv6Discard   basetypes.BoolValue   `tfsdk:"options_invalid_ipv6_discard"`
	ReservedFieldSetDiscard     basetypes.BoolValue   `tfsdk:"reserved_field_set_discard"`
	RoutingHeader0              basetypes.BoolValue   `tfsdk:"routing_header_0"`
	RoutingHeader1              basetypes.BoolValue   `tfsdk:"routing_header_1"`
	RoutingHeader253            basetypes.BoolValue   `tfsdk:"routing_header_253"`
	RoutingHeader254            basetypes.BoolValue   `tfsdk:"routing_header_254"`
	RoutingHeader255            basetypes.BoolValue   `tfsdk:"routing_header_255"`
	RoutingHeader3              basetypes.BoolValue   `tfsdk:"routing_header_3"`
	RoutingHeader4252           basetypes.BoolValue   `tfsdk:"routing_header_4_252"`
}

// ZoneProtectionProfilesIpv6FilterExtHdr represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesIpv6FilterExtHdr struct {
	DestOptionHdr basetypes.BoolValue `tfsdk:"dest_option_hdr"`
	HopByHopHdr   basetypes.BoolValue `tfsdk:"hop_by_hop_hdr"`
	RoutingHdr    basetypes.BoolValue `tfsdk:"routing_hdr"`
}

// ZoneProtectionProfilesIpv6IgnoreInvPkt represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesIpv6IgnoreInvPkt struct {
	DestUnreach  basetypes.BoolValue `tfsdk:"dest_unreach"`
	ParamProblem basetypes.BoolValue `tfsdk:"param_problem"`
	PktTooBig    basetypes.BoolValue `tfsdk:"pkt_too_big"`
	Redirect     basetypes.BoolValue `tfsdk:"redirect"`
	TimeExceeded basetypes.BoolValue `tfsdk:"time_exceeded"`
}

// ZoneProtectionProfilesL2SecGroupTagProtection represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesL2SecGroupTagProtection struct {
	Tags basetypes.ListValue `tfsdk:"tags"`
}

// ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner struct {
	Enable basetypes.BoolValue   `tfsdk:"enable"`
	Name   basetypes.StringValue `tfsdk:"name"`
	Tag    basetypes.StringValue `tfsdk:"tag"`
}

// ZoneProtectionProfilesNonIpProtocol represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesNonIpProtocol struct {
	ListType basetypes.StringValue `tfsdk:"list_type"`
	Protocol basetypes.ListValue   `tfsdk:"protocol"`
}

// ZoneProtectionProfilesNonIpProtocolProtocolInner represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesNonIpProtocolProtocolInner struct {
	Enable    basetypes.BoolValue   `tfsdk:"enable"`
	EtherType basetypes.StringValue `tfsdk:"ether_type"`
	Name      basetypes.StringValue `tfsdk:"name"`
}

// ZoneProtectionProfilesScanInner represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesScanInner struct {
	Action    basetypes.ObjectValue `tfsdk:"action"`
	Interval  basetypes.Int64Value  `tfsdk:"interval"`
	Name      basetypes.StringValue `tfsdk:"name"`
	Threshold basetypes.Int64Value  `tfsdk:"threshold"`
}

// ZoneProtectionProfilesScanInnerAction represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesScanInnerAction struct {
	Alert   basetypes.ObjectValue `tfsdk:"alert"`
	Allow   basetypes.ObjectValue `tfsdk:"allow"`
	Block   basetypes.ObjectValue `tfsdk:"block"`
	BlockIp basetypes.ObjectValue `tfsdk:"block_ip"`
}

// ZoneProtectionProfilesScanInnerActionBlockIp represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesScanInnerActionBlockIp struct {
	Duration basetypes.Int64Value  `tfsdk:"duration"`
	TrackBy  basetypes.StringValue `tfsdk:"track_by"`
}

// ZoneProtectionProfilesScanWhiteListInner represents a nested structure within the ZoneProtectionProfiles model
type ZoneProtectionProfilesScanWhiteListInner struct {
	Ipv4 basetypes.StringValue `tfsdk:"ipv4"`
	Ipv6 basetypes.StringValue `tfsdk:"ipv6"`
	Name basetypes.StringValue `tfsdk:"name"`
}

// AttrTypes defines the attribute types for the ZoneProtectionProfiles model.
func (o ZoneProtectionProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":                        basetypes.StringType{},
		"asymmetric_path":             basetypes.StringType{},
		"description":                 basetypes.StringType{},
		"device":                      basetypes.StringType{},
		"discard_icmp_embedded_error": basetypes.BoolType{},
		"flood": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"icmp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.BoolType{},
						"red": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"activate_rate": basetypes.Int64Type{},
								"alarm_rate":    basetypes.Int64Type{},
								"maximal_rate":  basetypes.Int64Type{},
							},
						},
					},
				},
				"icmpv6": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.BoolType{},
						"red": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"activate_rate": basetypes.Int64Type{},
								"alarm_rate":    basetypes.Int64Type{},
								"maximal_rate":  basetypes.Int64Type{},
							},
						},
					},
				},
				"other_ip": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.BoolType{},
						"red": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"activate_rate": basetypes.Int64Type{},
								"alarm_rate":    basetypes.Int64Type{},
								"maximal_rate":  basetypes.Int64Type{},
							},
						},
					},
				},
				"sctp_init": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.BoolType{},
						"red": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"activate_rate": basetypes.Int64Type{},
								"alarm_rate":    basetypes.Int64Type{},
								"maximal_rate":  basetypes.Int64Type{},
							},
						},
					},
				},
				"tcp_syn": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.BoolType{},
						"red": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"activate_rate": basetypes.Int64Type{},
								"alarm_rate":    basetypes.Int64Type{},
								"maximal_rate":  basetypes.Int64Type{},
							},
						},
						"syn_cookies": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"activate_rate": basetypes.Int64Type{},
								"alarm_rate":    basetypes.Int64Type{},
								"maximal_rate":  basetypes.Int64Type{},
							},
						},
					},
				},
				"udp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.BoolType{},
						"red": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"activate_rate": basetypes.Int64Type{},
								"alarm_rate":    basetypes.Int64Type{},
								"maximal_rate":  basetypes.Int64Type{},
							},
						},
					},
				},
			},
		},
		"folder":                     basetypes.StringType{},
		"fragmented_traffic_discard": basetypes.BoolType{},
		"icmp_frag_discard":          basetypes.BoolType{},
		"icmp_large_packet_discard":  basetypes.BoolType{},
		"icmp_ping_zero_id_discard":  basetypes.BoolType{},
		"id":                         basetypes.StringType{},
		"ipv6": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"anycast_source": basetypes.BoolType{},
				"filter_ext_hdr": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"dest_option_hdr": basetypes.BoolType{},
						"hop_by_hop_hdr":  basetypes.BoolType{},
						"routing_hdr":     basetypes.BoolType{},
					},
				},
				"icmpv6_too_big_small_mtu_discard": basetypes.BoolType{},
				"ignore_inv_pkt": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"dest_unreach":  basetypes.BoolType{},
						"param_problem": basetypes.BoolType{},
						"pkt_too_big":   basetypes.BoolType{},
						"redirect":      basetypes.BoolType{},
						"time_exceeded": basetypes.BoolType{},
					},
				},
				"ipv4_compatible_address":      basetypes.BoolType{},
				"needless_fragment_hdr":        basetypes.BoolType{},
				"options_invalid_ipv6_discard": basetypes.BoolType{},
				"reserved_field_set_discard":   basetypes.BoolType{},
				"routing_header_0":             basetypes.BoolType{},
				"routing_header_1":             basetypes.BoolType{},
				"routing_header_253":           basetypes.BoolType{},
				"routing_header_254":           basetypes.BoolType{},
				"routing_header_255":           basetypes.BoolType{},
				"routing_header_3":             basetypes.BoolType{},
				"routing_header_4_252":         basetypes.BoolType{},
			},
		},
		"l2_sec_group_tag_protection": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"tags": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.BoolType{},
						"name":   basetypes.StringType{},
						"tag":    basetypes.StringType{},
					},
				}},
			},
		},
		"loose_source_routing_discard":               basetypes.BoolType{},
		"malformed_option_discard":                   basetypes.BoolType{},
		"mismatched_overlapping_tcp_segment_discard": basetypes.BoolType{},
		"mptcp_option_strip":                         basetypes.StringType{},
		"name":                                       basetypes.StringType{},
		"non_ip_protocol": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"list_type": basetypes.StringType{},
				"protocol": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable":     basetypes.BoolType{},
						"ether_type": basetypes.StringType{},
						"name":       basetypes.StringType{},
					},
				}},
			},
		},
		"record_route_discard": basetypes.BoolType{},
		"reject_non_syn_tcp":   basetypes.StringType{},
		"scan": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"alert": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"allow": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"block": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"block_ip": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"duration": basetypes.Int64Type{},
								"track_by": basetypes.StringType{},
							},
						},
					},
				},
				"interval":  basetypes.Int64Type{},
				"name":      basetypes.StringType{},
				"threshold": basetypes.Int64Type{},
			},
		}},
		"scan_white_list": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ipv4": basetypes.StringType{},
				"ipv6": basetypes.StringType{},
				"name": basetypes.StringType{},
			},
		}},
		"security_discard":              basetypes.BoolType{},
		"snippet":                       basetypes.StringType{},
		"spoofed_ip_discard":            basetypes.BoolType{},
		"stream_id_discard":             basetypes.BoolType{},
		"strict_ip_check":               basetypes.BoolType{},
		"strict_source_routing_discard": basetypes.BoolType{},
		"suppress_icmp_needfrag":        basetypes.BoolType{},
		"suppress_icmp_timeexceeded":    basetypes.BoolType{},
		"tcp_fast_open_and_data_strip":  basetypes.BoolType{},
		"tcp_handshake_discard":         basetypes.BoolType{},
		"tcp_syn_with_data_discard":     basetypes.BoolType{},
		"tcp_synack_with_data_discard":  basetypes.BoolType{},
		"tcp_timestamp_strip":           basetypes.BoolType{},
		"timestamp_discard":             basetypes.BoolType{},
		"unknown_option_discard":        basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfiles objects.
func (o ZoneProtectionProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesFlood model.
func (o ZoneProtectionProfilesFlood) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"icmp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable": basetypes.BoolType{},
				"red": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"activate_rate": basetypes.Int64Type{},
						"alarm_rate":    basetypes.Int64Type{},
						"maximal_rate":  basetypes.Int64Type{},
					},
				},
			},
		},
		"icmpv6": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable": basetypes.BoolType{},
				"red": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"activate_rate": basetypes.Int64Type{},
						"alarm_rate":    basetypes.Int64Type{},
						"maximal_rate":  basetypes.Int64Type{},
					},
				},
			},
		},
		"other_ip": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable": basetypes.BoolType{},
				"red": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"activate_rate": basetypes.Int64Type{},
						"alarm_rate":    basetypes.Int64Type{},
						"maximal_rate":  basetypes.Int64Type{},
					},
				},
			},
		},
		"sctp_init": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable": basetypes.BoolType{},
				"red": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"activate_rate": basetypes.Int64Type{},
						"alarm_rate":    basetypes.Int64Type{},
						"maximal_rate":  basetypes.Int64Type{},
					},
				},
			},
		},
		"tcp_syn": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable": basetypes.BoolType{},
				"red": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"activate_rate": basetypes.Int64Type{},
						"alarm_rate":    basetypes.Int64Type{},
						"maximal_rate":  basetypes.Int64Type{},
					},
				},
				"syn_cookies": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"activate_rate": basetypes.Int64Type{},
						"alarm_rate":    basetypes.Int64Type{},
						"maximal_rate":  basetypes.Int64Type{},
					},
				},
			},
		},
		"udp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable": basetypes.BoolType{},
				"red": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"activate_rate": basetypes.Int64Type{},
						"alarm_rate":    basetypes.Int64Type{},
						"maximal_rate":  basetypes.Int64Type{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesFlood objects.
func (o ZoneProtectionProfilesFlood) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesFloodIcmp model.
func (o ZoneProtectionProfilesFloodIcmp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable": basetypes.BoolType{},
		"red": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"activate_rate": basetypes.Int64Type{},
				"alarm_rate":    basetypes.Int64Type{},
				"maximal_rate":  basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesFloodIcmp objects.
func (o ZoneProtectionProfilesFloodIcmp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesFloodIcmpRed model.
func (o ZoneProtectionProfilesFloodIcmpRed) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"activate_rate": basetypes.Int64Type{},
		"alarm_rate":    basetypes.Int64Type{},
		"maximal_rate":  basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesFloodIcmpRed objects.
func (o ZoneProtectionProfilesFloodIcmpRed) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesFloodIcmpv6 model.
func (o ZoneProtectionProfilesFloodIcmpv6) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable": basetypes.BoolType{},
		"red": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"activate_rate": basetypes.Int64Type{},
				"alarm_rate":    basetypes.Int64Type{},
				"maximal_rate":  basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesFloodIcmpv6 objects.
func (o ZoneProtectionProfilesFloodIcmpv6) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesFloodIcmpv6Red model.
func (o ZoneProtectionProfilesFloodIcmpv6Red) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"activate_rate": basetypes.Int64Type{},
		"alarm_rate":    basetypes.Int64Type{},
		"maximal_rate":  basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesFloodIcmpv6Red objects.
func (o ZoneProtectionProfilesFloodIcmpv6Red) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesFloodOtherIp model.
func (o ZoneProtectionProfilesFloodOtherIp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable": basetypes.BoolType{},
		"red": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"activate_rate": basetypes.Int64Type{},
				"alarm_rate":    basetypes.Int64Type{},
				"maximal_rate":  basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesFloodOtherIp objects.
func (o ZoneProtectionProfilesFloodOtherIp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesFloodOtherIpRed model.
func (o ZoneProtectionProfilesFloodOtherIpRed) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"activate_rate": basetypes.Int64Type{},
		"alarm_rate":    basetypes.Int64Type{},
		"maximal_rate":  basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesFloodOtherIpRed objects.
func (o ZoneProtectionProfilesFloodOtherIpRed) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesFloodSctpInit model.
func (o ZoneProtectionProfilesFloodSctpInit) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable": basetypes.BoolType{},
		"red": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"activate_rate": basetypes.Int64Type{},
				"alarm_rate":    basetypes.Int64Type{},
				"maximal_rate":  basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesFloodSctpInit objects.
func (o ZoneProtectionProfilesFloodSctpInit) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesFloodSctpInitRed model.
func (o ZoneProtectionProfilesFloodSctpInitRed) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"activate_rate": basetypes.Int64Type{},
		"alarm_rate":    basetypes.Int64Type{},
		"maximal_rate":  basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesFloodSctpInitRed objects.
func (o ZoneProtectionProfilesFloodSctpInitRed) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesFloodTcpSyn model.
func (o ZoneProtectionProfilesFloodTcpSyn) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable": basetypes.BoolType{},
		"red": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"activate_rate": basetypes.Int64Type{},
				"alarm_rate":    basetypes.Int64Type{},
				"maximal_rate":  basetypes.Int64Type{},
			},
		},
		"syn_cookies": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"activate_rate": basetypes.Int64Type{},
				"alarm_rate":    basetypes.Int64Type{},
				"maximal_rate":  basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesFloodTcpSyn objects.
func (o ZoneProtectionProfilesFloodTcpSyn) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesFloodTcpSynRed model.
func (o ZoneProtectionProfilesFloodTcpSynRed) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"activate_rate": basetypes.Int64Type{},
		"alarm_rate":    basetypes.Int64Type{},
		"maximal_rate":  basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesFloodTcpSynRed objects.
func (o ZoneProtectionProfilesFloodTcpSynRed) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesFloodTcpSynSynCookies model.
func (o ZoneProtectionProfilesFloodTcpSynSynCookies) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"activate_rate": basetypes.Int64Type{},
		"alarm_rate":    basetypes.Int64Type{},
		"maximal_rate":  basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesFloodTcpSynSynCookies objects.
func (o ZoneProtectionProfilesFloodTcpSynSynCookies) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesFloodUdp model.
func (o ZoneProtectionProfilesFloodUdp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable": basetypes.BoolType{},
		"red": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"activate_rate": basetypes.Int64Type{},
				"alarm_rate":    basetypes.Int64Type{},
				"maximal_rate":  basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesFloodUdp objects.
func (o ZoneProtectionProfilesFloodUdp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesFloodUdpRed model.
func (o ZoneProtectionProfilesFloodUdpRed) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"activate_rate": basetypes.Int64Type{},
		"alarm_rate":    basetypes.Int64Type{},
		"maximal_rate":  basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesFloodUdpRed objects.
func (o ZoneProtectionProfilesFloodUdpRed) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesIpv6 model.
func (o ZoneProtectionProfilesIpv6) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"anycast_source": basetypes.BoolType{},
		"filter_ext_hdr": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"dest_option_hdr": basetypes.BoolType{},
				"hop_by_hop_hdr":  basetypes.BoolType{},
				"routing_hdr":     basetypes.BoolType{},
			},
		},
		"icmpv6_too_big_small_mtu_discard": basetypes.BoolType{},
		"ignore_inv_pkt": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"dest_unreach":  basetypes.BoolType{},
				"param_problem": basetypes.BoolType{},
				"pkt_too_big":   basetypes.BoolType{},
				"redirect":      basetypes.BoolType{},
				"time_exceeded": basetypes.BoolType{},
			},
		},
		"ipv4_compatible_address":      basetypes.BoolType{},
		"needless_fragment_hdr":        basetypes.BoolType{},
		"options_invalid_ipv6_discard": basetypes.BoolType{},
		"reserved_field_set_discard":   basetypes.BoolType{},
		"routing_header_0":             basetypes.BoolType{},
		"routing_header_1":             basetypes.BoolType{},
		"routing_header_253":           basetypes.BoolType{},
		"routing_header_254":           basetypes.BoolType{},
		"routing_header_255":           basetypes.BoolType{},
		"routing_header_3":             basetypes.BoolType{},
		"routing_header_4_252":         basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesIpv6 objects.
func (o ZoneProtectionProfilesIpv6) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesIpv6FilterExtHdr model.
func (o ZoneProtectionProfilesIpv6FilterExtHdr) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"dest_option_hdr": basetypes.BoolType{},
		"hop_by_hop_hdr":  basetypes.BoolType{},
		"routing_hdr":     basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesIpv6FilterExtHdr objects.
func (o ZoneProtectionProfilesIpv6FilterExtHdr) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesIpv6IgnoreInvPkt model.
func (o ZoneProtectionProfilesIpv6IgnoreInvPkt) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"dest_unreach":  basetypes.BoolType{},
		"param_problem": basetypes.BoolType{},
		"pkt_too_big":   basetypes.BoolType{},
		"redirect":      basetypes.BoolType{},
		"time_exceeded": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesIpv6IgnoreInvPkt objects.
func (o ZoneProtectionProfilesIpv6IgnoreInvPkt) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesL2SecGroupTagProtection model.
func (o ZoneProtectionProfilesL2SecGroupTagProtection) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tags": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable": basetypes.BoolType{},
				"name":   basetypes.StringType{},
				"tag":    basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesL2SecGroupTagProtection objects.
func (o ZoneProtectionProfilesL2SecGroupTagProtection) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner model.
func (o ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable": basetypes.BoolType{},
		"name":   basetypes.StringType{},
		"tag":    basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner objects.
func (o ZoneProtectionProfilesL2SecGroupTagProtectionTagsInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesNonIpProtocol model.
func (o ZoneProtectionProfilesNonIpProtocol) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"list_type": basetypes.StringType{},
		"protocol": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable":     basetypes.BoolType{},
				"ether_type": basetypes.StringType{},
				"name":       basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesNonIpProtocol objects.
func (o ZoneProtectionProfilesNonIpProtocol) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesNonIpProtocolProtocolInner model.
func (o ZoneProtectionProfilesNonIpProtocolProtocolInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable":     basetypes.BoolType{},
		"ether_type": basetypes.StringType{},
		"name":       basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesNonIpProtocolProtocolInner objects.
func (o ZoneProtectionProfilesNonIpProtocolProtocolInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesScanInner model.
func (o ZoneProtectionProfilesScanInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"alert": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"allow": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"block": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"block_ip": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"duration": basetypes.Int64Type{},
						"track_by": basetypes.StringType{},
					},
				},
			},
		},
		"interval":  basetypes.Int64Type{},
		"name":      basetypes.StringType{},
		"threshold": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesScanInner objects.
func (o ZoneProtectionProfilesScanInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesScanInnerAction model.
func (o ZoneProtectionProfilesScanInnerAction) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"alert": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"allow": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"block": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"block_ip": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"duration": basetypes.Int64Type{},
				"track_by": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesScanInnerAction objects.
func (o ZoneProtectionProfilesScanInnerAction) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesScanInnerActionBlockIp model.
func (o ZoneProtectionProfilesScanInnerActionBlockIp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"duration": basetypes.Int64Type{},
		"track_by": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesScanInnerActionBlockIp objects.
func (o ZoneProtectionProfilesScanInnerActionBlockIp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ZoneProtectionProfilesScanWhiteListInner model.
func (o ZoneProtectionProfilesScanWhiteListInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ipv4": basetypes.StringType{},
		"ipv6": basetypes.StringType{},
		"name": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ZoneProtectionProfilesScanWhiteListInner objects.
func (o ZoneProtectionProfilesScanWhiteListInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// ZoneProtectionProfilesResourceSchema defines the schema for ZoneProtectionProfiles resource
var ZoneProtectionProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "ZoneProtectionProfile resource",
	Attributes: map[string]schema.Attribute{
		"asymmetric_path": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("global", "drop", "bypass"),
			},
			MarkdownDescription: "Determine whether to drop or bypass packets that contain out-of-sync ACKs or out-of-window sequence numbers:\n* `global` — Use system-wide setting that is assigned through TCP Settings or the CLI.\n* `drop` — Drop packets that contain an asymmetric path.\n* `bypass` — Bypass scanning on packets that contain an asymmetric path.\n",
			Optional:            true,
		},
		"description": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(255),
			},
			MarkdownDescription: "The description of the profile",
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
		"discard_icmp_embedded_error": schema.BoolAttribute{
			MarkdownDescription: "Discard ICMP packets that are embedded with an error message.",
			Optional:            true,
		},
		"flood": schema.SingleNestedAttribute{
			MarkdownDescription: "Flood",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"icmp": schema.SingleNestedAttribute{
					MarkdownDescription: "Icmp",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"enable": schema.BoolAttribute{
							MarkdownDescription: "Enable protection against ICMP floods?",
							Optional:            true,
						},
						"red": schema.SingleNestedAttribute{
							MarkdownDescription: "Red",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"activate_rate": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(0, 2000000),
									},
									MarkdownDescription: "The number of ICMP packets (not matching an existing session) that the zone receives per second before subsequent ICMP packets are dropped.",
									Required:            true,
								},
								"alarm_rate": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(0, 2000000),
									},
									MarkdownDescription: "The number of ICMP echo requests (pings not matching an existing session) that the zone receives per second that triggers an attack alarm.",
									Required:            true,
								},
								"maximal_rate": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(0, 2000000),
									},
									MarkdownDescription: "The maximum number of ICMP packets (not matching an existing session) that the zone receives per second before packets exceeding the maximum are dropped.",
									Required:            true,
								},
							},
						},
					},
				},
				"icmpv6": schema.SingleNestedAttribute{
					MarkdownDescription: "Icmpv6",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"enable": schema.BoolAttribute{
							MarkdownDescription: "Enable protection against ICMPv6 floods?",
							Optional:            true,
						},
						"red": schema.SingleNestedAttribute{
							MarkdownDescription: "Red",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"activate_rate": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(0, 2000000),
									},
									MarkdownDescription: "The number of ICMPv6 packets (not matching an existing session) that the zone receives per second before subsequent ICMPv6 packets are dropped.",
									Required:            true,
								},
								"alarm_rate": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(0, 2000000),
									},
									MarkdownDescription: "The number of ICMPv6 echo requests (pings not matching an existing session) that the zone receives per second that triggers an attack alarm.",
									Required:            true,
								},
								"maximal_rate": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(0, 2000000),
									},
									MarkdownDescription: "The maximum number of ICMPv6 packets (not matching an existing session) that the zone receives per second before packets exceeding the maximum are dropped.",
									Required:            true,
								},
							},
						},
					},
				},
				"other_ip": schema.SingleNestedAttribute{
					MarkdownDescription: "Other ip",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"enable": schema.BoolAttribute{
							MarkdownDescription: "Enable protection against other IP (non-TCP, non-ICMP, non-ICMPv6, non-SCTP, and non-UDP) floods?",
							Optional:            true,
						},
						"red": schema.SingleNestedAttribute{
							MarkdownDescription: "Red",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"activate_rate": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(0, 2000000),
									},
									MarkdownDescription: "Activate rate",
									Required:            true,
								},
								"alarm_rate": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(0, 2000000),
									},
									MarkdownDescription: "Alarm rate",
									Required:            true,
								},
								"maximal_rate": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(0, 2000000),
									},
									MarkdownDescription: "Maximal rate",
									Required:            true,
								},
							},
						},
					},
				},
				"sctp_init": schema.SingleNestedAttribute{
					MarkdownDescription: "Sctp init",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"enable": schema.BoolAttribute{
							MarkdownDescription: "Enable protection against floods of Stream Control Transmission Protocol (SCTP) packets that contain an Initiation (INIT) chunk?",
							Optional:            true,
						},
						"red": schema.SingleNestedAttribute{
							MarkdownDescription: "Red",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"activate_rate": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(0, 2000000),
									},
									MarkdownDescription: "The number of SCTP INIT packets (not matching an existing session) that the zone receives per second before subsequent SCTP INIT packets are dropped.",
									Required:            true,
								},
								"alarm_rate": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(0, 2000000),
									},
									MarkdownDescription: "The number of SCTP INIT packets (not matching an existing session) that the zone receives per second that triggers an attack alarm.",
									Required:            true,
								},
								"maximal_rate": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(0, 2000000),
									},
									MarkdownDescription: "The maximum number of SCTP INIT packets (not matching an existing session) that the zone receives per second before packets exceeding the maximum are dropped.",
									Required:            true,
								},
							},
						},
					},
				},
				"tcp_syn": schema.SingleNestedAttribute{
					MarkdownDescription: "Tcp syn",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"enable": schema.BoolAttribute{
							MarkdownDescription: "Enable protection against SYN floods?",
							Optional:            true,
						},
						"red": schema.SingleNestedAttribute{
							MarkdownDescription: "Red",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"activate_rate": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(0, 2000000),
									},
									MarkdownDescription: "When the flow exceeds the `activate_rate`` threshold, the firewall drops individual SYN packets randomly to restrict the flow.",
									Required:            true,
								},
								"alarm_rate": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(0, 2000000),
									},
									MarkdownDescription: "When the flow exceeds the `alert_rate`` threshold, an alarm is generated.",
									Required:            true,
								},
								"maximal_rate": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(0, 2000000),
									},
									MarkdownDescription: "When the flow exceeds the `maximal_rate` threshold, 100% of incoming SYN packets are dropped.",
									Required:            true,
								},
							},
						},
						"syn_cookies": schema.SingleNestedAttribute{
							MarkdownDescription: "Syn cookies",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"activate_rate": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(0, 2000000),
									},
									MarkdownDescription: "When the flow exceeds the `activate_rate`` threshold, the firewall drops individual SYN packets randomly to restrict the flow.",
									Required:            true,
								},
								"alarm_rate": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(0, 2000000),
									},
									MarkdownDescription: "When the flow exceeds the `alert_rate`` threshold, an alarm is generated.",
									Required:            true,
								},
								"maximal_rate": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(0, 2000000),
									},
									MarkdownDescription: "When the flow exceeds the `maximal_rate` threshold, 100% of incoming SYN packets are dropped.",
									Required:            true,
								},
							},
						},
					},
				},
				"udp": schema.SingleNestedAttribute{
					MarkdownDescription: "Udp",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"enable": schema.BoolAttribute{
							MarkdownDescription: "Enable protection against UDP floods?",
							Optional:            true,
						},
						"red": schema.SingleNestedAttribute{
							MarkdownDescription: "Red",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"activate_rate": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(0, 2000000),
									},
									MarkdownDescription: "The number of UDP packets (not matching an existing session) that the zone receives per second that triggers random dropping of UDP packets.",
									Required:            true,
								},
								"alarm_rate": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(0, 2000000),
									},
									MarkdownDescription: "The number of UDP packets (not matching an existing session) that the zone receives per second that triggers an attack alarm.",
									Required:            true,
								},
								"maximal_rate": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(0, 2000000),
									},
									MarkdownDescription: "The maximum number of UDP packets (not matching an existing session) the zone receives per second before packets exceeding the maximum are dropped.",
									Required:            true,
								},
							},
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
			MarkdownDescription: "The folder in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"fragmented_traffic_discard": schema.BoolAttribute{
			MarkdownDescription: "Discard fragmented IP packets.\n",
			Optional:            true,
		},
		"icmp_frag_discard": schema.BoolAttribute{
			MarkdownDescription: "Discard packets that consist of ICMP fragments.",
			Optional:            true,
		},
		"icmp_large_packet_discard": schema.BoolAttribute{
			MarkdownDescription: "Discard ICMP packets that are larger than 1024 bytes.",
			Optional:            true,
		},
		"icmp_ping_zero_id_discard": schema.BoolAttribute{
			MarkdownDescription: "Discard packets if the ICMP ping packet has an identifier value of 0.\n",
			Optional:            true,
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"ipv6": schema.SingleNestedAttribute{
			MarkdownDescription: "Ipv6",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"anycast_source": schema.BoolAttribute{
					MarkdownDescription: "Discard IPv6 packets that contain an anycast source address.",
					Optional:            true,
				},
				"filter_ext_hdr": schema.SingleNestedAttribute{
					MarkdownDescription: "Filter ext hdr",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"dest_option_hdr": schema.BoolAttribute{
							MarkdownDescription: "Discard IPv6 packets that contain the Destination Options extension, which contains options intended only for the destination of the packet.",
							Optional:            true,
						},
						"hop_by_hop_hdr": schema.BoolAttribute{
							MarkdownDescription: "Discard IPv6 packets that contain the Hop-by-Hop Options extension header.",
							Optional:            true,
						},
						"routing_hdr": schema.BoolAttribute{
							MarkdownDescription: "Discard IPv6 packets that contain the Routing extension header, which directs packets to one or more intermediate nodes on its way to its destination.",
							Optional:            true,
						},
					},
				},
				"icmpv6_too_big_small_mtu_discard": schema.BoolAttribute{
					MarkdownDescription: "Discard IPv6 packets that contain a Packet Too Big ICMPv6 message when the maximum transmission unit (MTU) is less than 1,280 bytes.",
					Optional:            true,
				},
				"ignore_inv_pkt": schema.SingleNestedAttribute{
					MarkdownDescription: "Ignore inv pkt",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"dest_unreach": schema.BoolAttribute{
							MarkdownDescription: "Require an explicit Security policy match for Destination Unreachable ICMPv6 messages, even when the message is associated with an existing session.",
							Optional:            true,
						},
						"param_problem": schema.BoolAttribute{
							MarkdownDescription: "Require an explicit Security policy match for Parameter Problem ICMPv6 messages, even when the message is associated with an existing session.",
							Optional:            true,
						},
						"pkt_too_big": schema.BoolAttribute{
							MarkdownDescription: "Require an explicit Security policy match for Packet Too Big ICMPv6 messages, even when the message is associated with an existing session.",
							Optional:            true,
						},
						"redirect": schema.BoolAttribute{
							MarkdownDescription: "Require an explicit Security policy match for Redirect Message ICMPv6 messages, even when the message is associated with an existing session.",
							Optional:            true,
						},
						"time_exceeded": schema.BoolAttribute{
							MarkdownDescription: "Require an explicit Security policy match for Time Exceeded ICMPv6 messages, even when the message is associated with an existing session.",
							Optional:            true,
						},
					},
				},
				"ipv4_compatible_address": schema.BoolAttribute{
					MarkdownDescription: "Discard IPv6 packets that are defined as an RFC 4291 IPv4-Compatible IPv6 address.",
					Optional:            true,
				},
				"needless_fragment_hdr": schema.BoolAttribute{
					MarkdownDescription: "Discard IPv6 packets with the last fragment flag (M=0) and offset of zero.",
					Optional:            true,
				},
				"options_invalid_ipv6_discard": schema.BoolAttribute{
					MarkdownDescription: "Discard IPv6 packets that contain invalid IPv6 options in an extension header.",
					Optional:            true,
				},
				"reserved_field_set_discard": schema.BoolAttribute{
					MarkdownDescription: "Discard IPv6 packets that have a header with a reserved field not set to zero.",
					Optional:            true,
				},
				"routing_header_0": schema.BoolAttribute{
					MarkdownDescription: "Drop packets with type 0 routing header.",
					Optional:            true,
				},
				"routing_header_1": schema.BoolAttribute{
					MarkdownDescription: "Drop packets with type 1 routing header.",
					Optional:            true,
				},
				"routing_header_253": schema.BoolAttribute{
					MarkdownDescription: "Drop packets with type 253 routing header.",
					Optional:            true,
				},
				"routing_header_254": schema.BoolAttribute{
					MarkdownDescription: "Drop packets with type 254 routing header.",
					Optional:            true,
				},
				"routing_header_255": schema.BoolAttribute{
					MarkdownDescription: "Drop packets with type 255 routing header.",
					Optional:            true,
				},
				"routing_header_3": schema.BoolAttribute{
					MarkdownDescription: "Drop packets with type 3 routing header.",
					Optional:            true,
				},
				"routing_header_4_252": schema.BoolAttribute{
					MarkdownDescription: "Drop packets with type 4 to type 252 routing header.",
					Optional:            true,
				},
			},
		},
		"l2_sec_group_tag_protection": schema.SingleNestedAttribute{
			MarkdownDescription: "L2 sec group tag protection",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"tags": schema.ListNestedAttribute{
					MarkdownDescription: "Tags",
					Optional:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"enable": schema.BoolAttribute{
								MarkdownDescription: "Enable this exclude list for Ethernet SGT protection.",
								Optional:            true,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: "Name for the list of Security Group Tags (SGTs).",
								Required:            true,
							},
							"tag": schema.StringAttribute{
								MarkdownDescription: "The Layer 2 SGTs in headers of packets that you want to exclude (drop) when the SGT matches this list in the Zone Protection profile applied to a zone (range is 0 to 65,535).",
								Required:            true,
							},
						},
					},
				},
			},
		},
		"loose_source_routing_discard": schema.BoolAttribute{
			MarkdownDescription: "Discard packets with the Loose Source Routing IP option set. Loose Source Routing is an option whereby a source of a datagram provides routing information and a gateway or host is allowed to choose any route of a number of intermediate gateways to get the datagram to the next address in the route.\n",
			Optional:            true,
		},
		"malformed_option_discard": schema.BoolAttribute{
			MarkdownDescription: "Discard packets if they have incorrect combinations of class, number, and length based on RFCs 791, 1108, 1393, and 2113.\n",
			Optional:            true,
		},
		"mismatched_overlapping_tcp_segment_discard": schema.BoolAttribute{
			MarkdownDescription: "Drop packets with mismatched overlapping TCP segments.\n",
			Optional:            true,
		},
		"mptcp_option_strip": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("no", "yes", "global"),
			},
			MarkdownDescription: "MPTCP is an extension of TCP that allows a client to maintain a connection by simultaneously using multiple paths to connect to the destination host. By default, MPTCP support is disabled, based on the global MPTCP setting.  Review or adjust the MPTCP settings for the security zones associated with this profile:\n* `no` — Enable MPTCP support (do not strip the MPTCP option).\n* `yes` — Disable MPTCP support (strip the MPTCP option). With this configured, MPTCP connections are converted to standard TCP connections, as MPTCP is backwards compatible with TCP.\n* `global` — Support MPTCP based on the global MPTCP setting. By default, the global MPTCP setting is set to yes so that MPTCP is disabled (the MPTCP option is stripped from the packet).\n",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("global"),
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(31),
			},
			MarkdownDescription: "The profile name",
			Required:            true,
		},
		"non_ip_protocol": schema.SingleNestedAttribute{
			MarkdownDescription: "Non ip protocol",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"list_type": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("exclude", "include"),
					},
					MarkdownDescription: "Specify the type of list you are creating for protocol protection:\n* Include List—Only the protocols on the list are allowed—in addition to IPv4 (0x0800), IPv6 (0x86DD), ARP (0x0806), and VLAN tagged frames (0x8100). All other protocols are implicitly denied (blocked).\n* Exclude List—Only the protocols on the list are denied; all other protocols are implicitly allowed. You cannot exclude IPv4 (0x0800), IPv6 (0x86DD), ARP (0x0806), or VLAN tagged frames (0x8100).\n",
					Optional:            true,
				},
				"protocol": schema.ListNestedAttribute{
					MarkdownDescription: "Protocol",
					Optional:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"enable": schema.BoolAttribute{
								MarkdownDescription: "Enable the Ethertype code on the list.",
								Optional:            true,
							},
							"ether_type": schema.StringAttribute{
								MarkdownDescription: "Enter an Ethertype code (protocol) preceded by 0x to indicate hexadecimal (range is 0x0000 to 0xFFFF). A list can have a maximum of 64 Ethertypes. Some sources of Ethertype codes are:\n* [IEEE hexadecimal Ethertype](https://www.iana.org/assignments/ieee-802-numbers/ieee-802-numbers.xhtml)\n* [standards.ieee.org/develop/regauth/ethertype/eth.txt](https://standards-oui.ieee.org/ethertype/eth.txt)\n* [www.cavebear.com/archive/cavebear/Ethernet/type.html](https://www.cavebear.com/archive/cavebear/Ethernet/type.html)\n",
								Required:            true,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: "Enter the protocol name that corresponds to the Ethertype code you are adding to the list. The firewall does not verify that the protocol name matches the Ethertype code but the Ethertype code does determine the protocol filter.\n",
								Required:            true,
							},
						},
					},
				},
			},
		},
		"record_route_discard": schema.BoolAttribute{
			MarkdownDescription: "Discard packets with the Record Route IP option set. When a datagram has this option, each router that routes the datagram adds its own IP address to the header, thus providing the path to the recipient.\n",
			Optional:            true,
		},
		"reject_non_syn_tcp": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("global", "yes", "no"),
			},
			MarkdownDescription: "Determine whether to reject the packet if the first packet for the TCP session setup is not a SYN packet:\n* `global` — Use system-wide setting that is assigned through the CLI.\n* `yes` — Reject non-SYN TCP.\n* `no` — Accept non-SYN TCP.\n",
			Optional:            true,
		},
		"scan": schema.ListNestedAttribute{
			MarkdownDescription: "Scan",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"action": schema.SingleNestedAttribute{
						MarkdownDescription: "Action",
						Optional:            true,
						Attributes: map[string]schema.Attribute{
							"alert": schema.SingleNestedAttribute{
								MarkdownDescription: "Alert",
								Optional:            true,
								Attributes:          map[string]schema.Attribute{},
							},
							"allow": schema.SingleNestedAttribute{
								MarkdownDescription: "Allow",
								Optional:            true,
								Attributes:          map[string]schema.Attribute{},
							},
							"block": schema.SingleNestedAttribute{
								MarkdownDescription: "Block",
								Optional:            true,
								Attributes:          map[string]schema.Attribute{},
							},
							"block_ip": schema.SingleNestedAttribute{
								MarkdownDescription: "Block ip",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"duration": schema.Int64Attribute{
										Validators: []validator.Int64{
											int64validator.Between(1, 3600),
										},
										MarkdownDescription: "Duration",
										Required:            true,
									},
									"track_by": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("source-and-destination", "source"),
										},
										MarkdownDescription: "Track by",
										Required:            true,
									},
								},
							},
						},
					},
					"interval": schema.Int64Attribute{
						Validators: []validator.Int64{
							int64validator.Between(2, 65535),
						},
						MarkdownDescription: "Interval",
						Optional:            true,
					},
					"name": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("8001", "8002", "8003", "8006"),
						},
						MarkdownDescription: "The threat ID number.  These can be found in [Palo Alto Networks ThreatVault](https://threatvault.paloaltonetworks.com).\n* \"8001\" - TCP Port Scan\n* \"8002\" - Host Sweep\n* \"8003\" - UDP Port Scan\n* \"8006\" - Port Scan\n",
						Required:            true,
					},
					"threshold": schema.Int64Attribute{
						Validators: []validator.Int64{
							int64validator.Between(2, 65535),
						},
						MarkdownDescription: "Threshold",
						Optional:            true,
					},
				},
			},
		},
		"scan_white_list": schema.ListNestedAttribute{
			MarkdownDescription: "Scan white list",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"ipv4": schema.StringAttribute{
						MarkdownDescription: "Ipv4",
						Optional:            true,
					},
					"ipv6": schema.StringAttribute{
						MarkdownDescription: "Ipv6",
						Optional:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "A descriptive name for the address to exclude.",
						Required:            true,
					},
				},
			},
		},
		"security_discard": schema.BoolAttribute{
			MarkdownDescription: "Discard packets if the security option is defined.\n",
			Optional:            true,
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
		"spoofed_ip_discard": schema.BoolAttribute{
			MarkdownDescription: "Check that the source IP address of the ingress packet is routable and the routing interface is in the same zone as the ingress interface. If either condition is not true, discard the packet.\n",
			Optional:            true,
		},
		"stream_id_discard": schema.BoolAttribute{
			MarkdownDescription: "Discard packets if the Stream ID option is defined.\n",
			Optional:            true,
		},
		"strict_ip_check": schema.BoolAttribute{
			MarkdownDescription: "Check that both conditions are true:\n* The source IP address is not the subnet broadcast IP address of the ingress interface.\n* The source IP address is routable over the exact ingress interface.\nIf either condition is not true, discard the packet.\n",
			Optional:            true,
		},
		"strict_source_routing_discard": schema.BoolAttribute{
			MarkdownDescription: "Discard packets with the Strict Source Routing IP option set. Strict Source Routing is an option whereby a source of a datagram provides routing information through which a gateway or host must send the datagram.\n",
			Optional:            true,
		},
		"suppress_icmp_needfrag": schema.BoolAttribute{
			MarkdownDescription: "Stop sending ICMP fragmentation needed messages in response to packets that exceed the interface MTU and have the do not fragment (DF) bit set. This setting will interfere with the PMTUD process performed by hosts behind the firewall.\n",
			Optional:            true,
		},
		"suppress_icmp_timeexceeded": schema.BoolAttribute{
			MarkdownDescription: "Stop sending ICMP TTL expired messages.",
			Optional:            true,
		},
		"tcp_fast_open_and_data_strip": schema.BoolAttribute{
			MarkdownDescription: "Strip the TCP Fast Open option (and data payload, if any) from the TCP SYN or SYN-ACK packet during a TCP three-way handshake.\n",
			Optional:            true,
		},
		"tcp_handshake_discard": schema.BoolAttribute{
			MarkdownDescription: "Drop packets with split handshakes.\n",
			Optional:            true,
		},
		"tcp_syn_with_data_discard": schema.BoolAttribute{
			MarkdownDescription: "Prevent a TCP session from being established if the TCP SYN packet contains data during a three-way handshake.\n",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(true),
		},
		"tcp_synack_with_data_discard": schema.BoolAttribute{
			MarkdownDescription: "Prevent a TCP session from being established if the TCP SYN-ACK packet contains data during a three-way handshake.\n",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(true),
		},
		"tcp_timestamp_strip": schema.BoolAttribute{
			MarkdownDescription: "Determine whether the packet has a TCP timestamp in the header and, if it does, strip the timestamp from the header.\n",
			Optional:            true,
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"timestamp_discard": schema.BoolAttribute{
			MarkdownDescription: "Discard packets with the Timestamp IP option set.\n",
			Optional:            true,
		},
		"unknown_option_discard": schema.BoolAttribute{
			MarkdownDescription: "Discard packets if the class and number are unknown.\n",
			Optional:            true,
		},
	},
}

// ZoneProtectionProfilesDataSourceSchema defines the schema for ZoneProtectionProfiles data source
var ZoneProtectionProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "ZoneProtectionProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"asymmetric_path": dsschema.StringAttribute{
			MarkdownDescription: "Determine whether to drop or bypass packets that contain out-of-sync ACKs or out-of-window sequence numbers:\n* `global` — Use system-wide setting that is assigned through TCP Settings or the CLI.\n* `drop` — Drop packets that contain an asymmetric path.\n* `bypass` — Bypass scanning on packets that contain an asymmetric path.\n",
			Computed:            true,
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "The description of the profile",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"discard_icmp_embedded_error": dsschema.BoolAttribute{
			MarkdownDescription: "Discard ICMP packets that are embedded with an error message.",
			Computed:            true,
		},
		"flood": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Flood",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"icmp": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Icmp",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"enable": dsschema.BoolAttribute{
							MarkdownDescription: "Enable protection against ICMP floods?",
							Computed:            true,
						},
						"red": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Red",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"activate_rate": dsschema.Int64Attribute{
									MarkdownDescription: "The number of ICMP packets (not matching an existing session) that the zone receives per second before subsequent ICMP packets are dropped.",
									Computed:            true,
								},
								"alarm_rate": dsschema.Int64Attribute{
									MarkdownDescription: "The number of ICMP echo requests (pings not matching an existing session) that the zone receives per second that triggers an attack alarm.",
									Computed:            true,
								},
								"maximal_rate": dsschema.Int64Attribute{
									MarkdownDescription: "The maximum number of ICMP packets (not matching an existing session) that the zone receives per second before packets exceeding the maximum are dropped.",
									Computed:            true,
								},
							},
						},
					},
				},
				"icmpv6": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Icmpv6",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"enable": dsschema.BoolAttribute{
							MarkdownDescription: "Enable protection against ICMPv6 floods?",
							Computed:            true,
						},
						"red": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Red",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"activate_rate": dsschema.Int64Attribute{
									MarkdownDescription: "The number of ICMPv6 packets (not matching an existing session) that the zone receives per second before subsequent ICMPv6 packets are dropped.",
									Computed:            true,
								},
								"alarm_rate": dsschema.Int64Attribute{
									MarkdownDescription: "The number of ICMPv6 echo requests (pings not matching an existing session) that the zone receives per second that triggers an attack alarm.",
									Computed:            true,
								},
								"maximal_rate": dsschema.Int64Attribute{
									MarkdownDescription: "The maximum number of ICMPv6 packets (not matching an existing session) that the zone receives per second before packets exceeding the maximum are dropped.",
									Computed:            true,
								},
							},
						},
					},
				},
				"other_ip": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Other ip",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"enable": dsschema.BoolAttribute{
							MarkdownDescription: "Enable protection against other IP (non-TCP, non-ICMP, non-ICMPv6, non-SCTP, and non-UDP) floods?",
							Computed:            true,
						},
						"red": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Red",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"activate_rate": dsschema.Int64Attribute{
									MarkdownDescription: "Activate rate",
									Computed:            true,
								},
								"alarm_rate": dsschema.Int64Attribute{
									MarkdownDescription: "Alarm rate",
									Computed:            true,
								},
								"maximal_rate": dsschema.Int64Attribute{
									MarkdownDescription: "Maximal rate",
									Computed:            true,
								},
							},
						},
					},
				},
				"sctp_init": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Sctp init",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"enable": dsschema.BoolAttribute{
							MarkdownDescription: "Enable protection against floods of Stream Control Transmission Protocol (SCTP) packets that contain an Initiation (INIT) chunk?",
							Computed:            true,
						},
						"red": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Red",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"activate_rate": dsschema.Int64Attribute{
									MarkdownDescription: "The number of SCTP INIT packets (not matching an existing session) that the zone receives per second before subsequent SCTP INIT packets are dropped.",
									Computed:            true,
								},
								"alarm_rate": dsschema.Int64Attribute{
									MarkdownDescription: "The number of SCTP INIT packets (not matching an existing session) that the zone receives per second that triggers an attack alarm.",
									Computed:            true,
								},
								"maximal_rate": dsschema.Int64Attribute{
									MarkdownDescription: "The maximum number of SCTP INIT packets (not matching an existing session) that the zone receives per second before packets exceeding the maximum are dropped.",
									Computed:            true,
								},
							},
						},
					},
				},
				"tcp_syn": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Tcp syn",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"enable": dsschema.BoolAttribute{
							MarkdownDescription: "Enable protection against SYN floods?",
							Computed:            true,
						},
						"red": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Red",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"activate_rate": dsschema.Int64Attribute{
									MarkdownDescription: "When the flow exceeds the `activate_rate`` threshold, the firewall drops individual SYN packets randomly to restrict the flow.",
									Computed:            true,
								},
								"alarm_rate": dsschema.Int64Attribute{
									MarkdownDescription: "When the flow exceeds the `alert_rate`` threshold, an alarm is generated.",
									Computed:            true,
								},
								"maximal_rate": dsschema.Int64Attribute{
									MarkdownDescription: "When the flow exceeds the `maximal_rate` threshold, 100% of incoming SYN packets are dropped.",
									Computed:            true,
								},
							},
						},
						"syn_cookies": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Syn cookies",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"activate_rate": dsschema.Int64Attribute{
									MarkdownDescription: "When the flow exceeds the `activate_rate`` threshold, the firewall drops individual SYN packets randomly to restrict the flow.",
									Computed:            true,
								},
								"alarm_rate": dsschema.Int64Attribute{
									MarkdownDescription: "When the flow exceeds the `alert_rate`` threshold, an alarm is generated.",
									Computed:            true,
								},
								"maximal_rate": dsschema.Int64Attribute{
									MarkdownDescription: "When the flow exceeds the `maximal_rate` threshold, 100% of incoming SYN packets are dropped.",
									Computed:            true,
								},
							},
						},
					},
				},
				"udp": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Udp",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"enable": dsschema.BoolAttribute{
							MarkdownDescription: "Enable protection against UDP floods?",
							Computed:            true,
						},
						"red": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Red",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"activate_rate": dsschema.Int64Attribute{
									MarkdownDescription: "The number of UDP packets (not matching an existing session) that the zone receives per second that triggers random dropping of UDP packets.",
									Computed:            true,
								},
								"alarm_rate": dsschema.Int64Attribute{
									MarkdownDescription: "The number of UDP packets (not matching an existing session) that the zone receives per second that triggers an attack alarm.",
									Computed:            true,
								},
								"maximal_rate": dsschema.Int64Attribute{
									MarkdownDescription: "The maximum number of UDP packets (not matching an existing session) the zone receives per second before packets exceeding the maximum are dropped.",
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined",
			Computed:            true,
		},
		"fragmented_traffic_discard": dsschema.BoolAttribute{
			MarkdownDescription: "Discard fragmented IP packets.\n",
			Computed:            true,
		},
		"icmp_frag_discard": dsschema.BoolAttribute{
			MarkdownDescription: "Discard packets that consist of ICMP fragments.",
			Computed:            true,
		},
		"icmp_large_packet_discard": dsschema.BoolAttribute{
			MarkdownDescription: "Discard ICMP packets that are larger than 1024 bytes.",
			Computed:            true,
		},
		"icmp_ping_zero_id_discard": dsschema.BoolAttribute{
			MarkdownDescription: "Discard packets if the ICMP ping packet has an identifier value of 0.\n",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"ipv6": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Ipv6",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"anycast_source": dsschema.BoolAttribute{
					MarkdownDescription: "Discard IPv6 packets that contain an anycast source address.",
					Computed:            true,
				},
				"filter_ext_hdr": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Filter ext hdr",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"dest_option_hdr": dsschema.BoolAttribute{
							MarkdownDescription: "Discard IPv6 packets that contain the Destination Options extension, which contains options intended only for the destination of the packet.",
							Computed:            true,
						},
						"hop_by_hop_hdr": dsschema.BoolAttribute{
							MarkdownDescription: "Discard IPv6 packets that contain the Hop-by-Hop Options extension header.",
							Computed:            true,
						},
						"routing_hdr": dsschema.BoolAttribute{
							MarkdownDescription: "Discard IPv6 packets that contain the Routing extension header, which directs packets to one or more intermediate nodes on its way to its destination.",
							Computed:            true,
						},
					},
				},
				"icmpv6_too_big_small_mtu_discard": dsschema.BoolAttribute{
					MarkdownDescription: "Discard IPv6 packets that contain a Packet Too Big ICMPv6 message when the maximum transmission unit (MTU) is less than 1,280 bytes.",
					Computed:            true,
				},
				"ignore_inv_pkt": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Ignore inv pkt",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"dest_unreach": dsschema.BoolAttribute{
							MarkdownDescription: "Require an explicit Security policy match for Destination Unreachable ICMPv6 messages, even when the message is associated with an existing session.",
							Computed:            true,
						},
						"param_problem": dsschema.BoolAttribute{
							MarkdownDescription: "Require an explicit Security policy match for Parameter Problem ICMPv6 messages, even when the message is associated with an existing session.",
							Computed:            true,
						},
						"pkt_too_big": dsschema.BoolAttribute{
							MarkdownDescription: "Require an explicit Security policy match for Packet Too Big ICMPv6 messages, even when the message is associated with an existing session.",
							Computed:            true,
						},
						"redirect": dsschema.BoolAttribute{
							MarkdownDescription: "Require an explicit Security policy match for Redirect Message ICMPv6 messages, even when the message is associated with an existing session.",
							Computed:            true,
						},
						"time_exceeded": dsschema.BoolAttribute{
							MarkdownDescription: "Require an explicit Security policy match for Time Exceeded ICMPv6 messages, even when the message is associated with an existing session.",
							Computed:            true,
						},
					},
				},
				"ipv4_compatible_address": dsschema.BoolAttribute{
					MarkdownDescription: "Discard IPv6 packets that are defined as an RFC 4291 IPv4-Compatible IPv6 address.",
					Computed:            true,
				},
				"needless_fragment_hdr": dsschema.BoolAttribute{
					MarkdownDescription: "Discard IPv6 packets with the last fragment flag (M=0) and offset of zero.",
					Computed:            true,
				},
				"options_invalid_ipv6_discard": dsschema.BoolAttribute{
					MarkdownDescription: "Discard IPv6 packets that contain invalid IPv6 options in an extension header.",
					Computed:            true,
				},
				"reserved_field_set_discard": dsschema.BoolAttribute{
					MarkdownDescription: "Discard IPv6 packets that have a header with a reserved field not set to zero.",
					Computed:            true,
				},
				"routing_header_0": dsschema.BoolAttribute{
					MarkdownDescription: "Drop packets with type 0 routing header.",
					Computed:            true,
				},
				"routing_header_1": dsschema.BoolAttribute{
					MarkdownDescription: "Drop packets with type 1 routing header.",
					Computed:            true,
				},
				"routing_header_253": dsschema.BoolAttribute{
					MarkdownDescription: "Drop packets with type 253 routing header.",
					Computed:            true,
				},
				"routing_header_254": dsschema.BoolAttribute{
					MarkdownDescription: "Drop packets with type 254 routing header.",
					Computed:            true,
				},
				"routing_header_255": dsschema.BoolAttribute{
					MarkdownDescription: "Drop packets with type 255 routing header.",
					Computed:            true,
				},
				"routing_header_3": dsschema.BoolAttribute{
					MarkdownDescription: "Drop packets with type 3 routing header.",
					Computed:            true,
				},
				"routing_header_4_252": dsschema.BoolAttribute{
					MarkdownDescription: "Drop packets with type 4 to type 252 routing header.",
					Computed:            true,
				},
			},
		},
		"l2_sec_group_tag_protection": dsschema.SingleNestedAttribute{
			MarkdownDescription: "L2 sec group tag protection",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"tags": dsschema.ListNestedAttribute{
					MarkdownDescription: "Tags",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"enable": dsschema.BoolAttribute{
								MarkdownDescription: "Enable this exclude list for Ethernet SGT protection.",
								Computed:            true,
							},
							"name": dsschema.StringAttribute{
								MarkdownDescription: "Name for the list of Security Group Tags (SGTs).",
								Computed:            true,
							},
							"tag": dsschema.StringAttribute{
								MarkdownDescription: "The Layer 2 SGTs in headers of packets that you want to exclude (drop) when the SGT matches this list in the Zone Protection profile applied to a zone (range is 0 to 65,535).",
								Computed:            true,
							},
						},
					},
				},
			},
		},
		"loose_source_routing_discard": dsschema.BoolAttribute{
			MarkdownDescription: "Discard packets with the Loose Source Routing IP option set. Loose Source Routing is an option whereby a source of a datagram provides routing information and a gateway or host is allowed to choose any route of a number of intermediate gateways to get the datagram to the next address in the route.\n",
			Computed:            true,
		},
		"malformed_option_discard": dsschema.BoolAttribute{
			MarkdownDescription: "Discard packets if they have incorrect combinations of class, number, and length based on RFCs 791, 1108, 1393, and 2113.\n",
			Computed:            true,
		},
		"mismatched_overlapping_tcp_segment_discard": dsschema.BoolAttribute{
			MarkdownDescription: "Drop packets with mismatched overlapping TCP segments.\n",
			Computed:            true,
		},
		"mptcp_option_strip": dsschema.StringAttribute{
			MarkdownDescription: "MPTCP is an extension of TCP that allows a client to maintain a connection by simultaneously using multiple paths to connect to the destination host. By default, MPTCP support is disabled, based on the global MPTCP setting.  Review or adjust the MPTCP settings for the security zones associated with this profile:\n* `no` — Enable MPTCP support (do not strip the MPTCP option).\n* `yes` — Disable MPTCP support (strip the MPTCP option). With this configured, MPTCP connections are converted to standard TCP connections, as MPTCP is backwards compatible with TCP.\n* `global` — Support MPTCP based on the global MPTCP setting. By default, the global MPTCP setting is set to yes so that MPTCP is disabled (the MPTCP option is stripped from the packet).\n",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The profile name",
			Optional:            true,
			Computed:            true,
		},
		"non_ip_protocol": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Non ip protocol",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"list_type": dsschema.StringAttribute{
					MarkdownDescription: "Specify the type of list you are creating for protocol protection:\n* Include List—Only the protocols on the list are allowed—in addition to IPv4 (0x0800), IPv6 (0x86DD), ARP (0x0806), and VLAN tagged frames (0x8100). All other protocols are implicitly denied (blocked).\n* Exclude List—Only the protocols on the list are denied; all other protocols are implicitly allowed. You cannot exclude IPv4 (0x0800), IPv6 (0x86DD), ARP (0x0806), or VLAN tagged frames (0x8100).\n",
					Computed:            true,
				},
				"protocol": dsschema.ListNestedAttribute{
					MarkdownDescription: "Protocol",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"enable": dsschema.BoolAttribute{
								MarkdownDescription: "Enable the Ethertype code on the list.",
								Computed:            true,
							},
							"ether_type": dsschema.StringAttribute{
								MarkdownDescription: "Enter an Ethertype code (protocol) preceded by 0x to indicate hexadecimal (range is 0x0000 to 0xFFFF). A list can have a maximum of 64 Ethertypes. Some sources of Ethertype codes are:\n* [IEEE hexadecimal Ethertype](https://www.iana.org/assignments/ieee-802-numbers/ieee-802-numbers.xhtml)\n* [standards.ieee.org/develop/regauth/ethertype/eth.txt](https://standards-oui.ieee.org/ethertype/eth.txt)\n* [www.cavebear.com/archive/cavebear/Ethernet/type.html](https://www.cavebear.com/archive/cavebear/Ethernet/type.html)\n",
								Computed:            true,
							},
							"name": dsschema.StringAttribute{
								MarkdownDescription: "Enter the protocol name that corresponds to the Ethertype code you are adding to the list. The firewall does not verify that the protocol name matches the Ethertype code but the Ethertype code does determine the protocol filter.\n",
								Computed:            true,
							},
						},
					},
				},
			},
		},
		"record_route_discard": dsschema.BoolAttribute{
			MarkdownDescription: "Discard packets with the Record Route IP option set. When a datagram has this option, each router that routes the datagram adds its own IP address to the header, thus providing the path to the recipient.\n",
			Computed:            true,
		},
		"reject_non_syn_tcp": dsschema.StringAttribute{
			MarkdownDescription: "Determine whether to reject the packet if the first packet for the TCP session setup is not a SYN packet:\n* `global` — Use system-wide setting that is assigned through the CLI.\n* `yes` — Reject non-SYN TCP.\n* `no` — Accept non-SYN TCP.\n",
			Computed:            true,
		},
		"scan": dsschema.ListNestedAttribute{
			MarkdownDescription: "Scan",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"action": dsschema.SingleNestedAttribute{
						MarkdownDescription: "Action",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"alert": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Alert",
								Computed:            true,
								Attributes:          map[string]dsschema.Attribute{},
							},
							"allow": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Allow",
								Computed:            true,
								Attributes:          map[string]dsschema.Attribute{},
							},
							"block": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Block",
								Computed:            true,
								Attributes:          map[string]dsschema.Attribute{},
							},
							"block_ip": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Block ip",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"duration": dsschema.Int64Attribute{
										MarkdownDescription: "Duration",
										Computed:            true,
									},
									"track_by": dsschema.StringAttribute{
										MarkdownDescription: "Track by",
										Computed:            true,
									},
								},
							},
						},
					},
					"interval": dsschema.Int64Attribute{
						MarkdownDescription: "Interval",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "The threat ID number.  These can be found in [Palo Alto Networks ThreatVault](https://threatvault.paloaltonetworks.com).\n* \"8001\" - TCP Port Scan\n* \"8002\" - Host Sweep\n* \"8003\" - UDP Port Scan\n* \"8006\" - Port Scan\n",
						Computed:            true,
					},
					"threshold": dsschema.Int64Attribute{
						MarkdownDescription: "Threshold",
						Computed:            true,
					},
				},
			},
		},
		"scan_white_list": dsschema.ListNestedAttribute{
			MarkdownDescription: "Scan white list",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"ipv4": dsschema.StringAttribute{
						MarkdownDescription: "Ipv4",
						Computed:            true,
					},
					"ipv6": dsschema.StringAttribute{
						MarkdownDescription: "Ipv6",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "A descriptive name for the address to exclude.",
						Computed:            true,
					},
				},
			},
		},
		"security_discard": dsschema.BoolAttribute{
			MarkdownDescription: "Discard packets if the security option is defined.\n",
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"spoofed_ip_discard": dsschema.BoolAttribute{
			MarkdownDescription: "Check that the source IP address of the ingress packet is routable and the routing interface is in the same zone as the ingress interface. If either condition is not true, discard the packet.\n",
			Computed:            true,
		},
		"stream_id_discard": dsschema.BoolAttribute{
			MarkdownDescription: "Discard packets if the Stream ID option is defined.\n",
			Computed:            true,
		},
		"strict_ip_check": dsschema.BoolAttribute{
			MarkdownDescription: "Check that both conditions are true:\n* The source IP address is not the subnet broadcast IP address of the ingress interface.\n* The source IP address is routable over the exact ingress interface.\nIf either condition is not true, discard the packet.\n",
			Computed:            true,
		},
		"strict_source_routing_discard": dsschema.BoolAttribute{
			MarkdownDescription: "Discard packets with the Strict Source Routing IP option set. Strict Source Routing is an option whereby a source of a datagram provides routing information through which a gateway or host must send the datagram.\n",
			Computed:            true,
		},
		"suppress_icmp_needfrag": dsschema.BoolAttribute{
			MarkdownDescription: "Stop sending ICMP fragmentation needed messages in response to packets that exceed the interface MTU and have the do not fragment (DF) bit set. This setting will interfere with the PMTUD process performed by hosts behind the firewall.\n",
			Computed:            true,
		},
		"suppress_icmp_timeexceeded": dsschema.BoolAttribute{
			MarkdownDescription: "Stop sending ICMP TTL expired messages.",
			Computed:            true,
		},
		"tcp_fast_open_and_data_strip": dsschema.BoolAttribute{
			MarkdownDescription: "Strip the TCP Fast Open option (and data payload, if any) from the TCP SYN or SYN-ACK packet during a TCP three-way handshake.\n",
			Computed:            true,
		},
		"tcp_handshake_discard": dsschema.BoolAttribute{
			MarkdownDescription: "Drop packets with split handshakes.\n",
			Computed:            true,
		},
		"tcp_syn_with_data_discard": dsschema.BoolAttribute{
			MarkdownDescription: "Prevent a TCP session from being established if the TCP SYN packet contains data during a three-way handshake.\n",
			Computed:            true,
		},
		"tcp_synack_with_data_discard": dsschema.BoolAttribute{
			MarkdownDescription: "Prevent a TCP session from being established if the TCP SYN-ACK packet contains data during a three-way handshake.\n",
			Computed:            true,
		},
		"tcp_timestamp_strip": dsschema.BoolAttribute{
			MarkdownDescription: "Determine whether the packet has a TCP timestamp in the header and, if it does, strip the timestamp from the header.\n",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"timestamp_discard": dsschema.BoolAttribute{
			MarkdownDescription: "Discard packets with the Timestamp IP option set.\n",
			Computed:            true,
		},
		"unknown_option_discard": dsschema.BoolAttribute{
			MarkdownDescription: "Discard packets if the class and number are unknown.\n",
			Computed:            true,
		},
	},
}

// ZoneProtectionProfilesListModel represents the data model for a list data source.
type ZoneProtectionProfilesListModel struct {
	Tfid    types.String             `tfsdk:"tfid"`
	Data    []ZoneProtectionProfiles `tfsdk:"data"`
	Limit   types.Int64              `tfsdk:"limit"`
	Offset  types.Int64              `tfsdk:"offset"`
	Name    types.String             `tfsdk:"name"`
	Total   types.Int64              `tfsdk:"total"`
	Folder  types.String             `tfsdk:"folder"`
	Snippet types.String             `tfsdk:"snippet"`
	Device  types.String             `tfsdk:"device"`
}

// ZoneProtectionProfilesListDataSourceSchema defines the schema for a list data source.
var ZoneProtectionProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: ZoneProtectionProfilesDataSourceSchema.Attributes,
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
