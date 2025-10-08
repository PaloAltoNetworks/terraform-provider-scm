package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
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

// LogicalRouters represents the Terraform model for LogicalRouters
type LogicalRouters struct {
	Tfid         types.String          `tfsdk:"tfid"`
	Device       basetypes.StringValue `tfsdk:"device"`
	Folder       basetypes.StringValue `tfsdk:"folder"`
	Id           basetypes.StringValue `tfsdk:"id"`
	Name         basetypes.StringValue `tfsdk:"name"`
	RoutingStack basetypes.StringValue `tfsdk:"routing_stack"`
	Snippet      basetypes.StringValue `tfsdk:"snippet"`
	Vrf          basetypes.ListValue   `tfsdk:"vrf"`
}

// LogicalRoutersVrfInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInner struct {
	AdminDists   basetypes.ObjectValue  `tfsdk:"admin_dists"`
	Bgp          basetypes.ObjectValue  `tfsdk:"bgp"`
	Ecmp         basetypes.ObjectValue  `tfsdk:"ecmp"`
	GlobalVrid   basetypes.Float64Value `tfsdk:"global_vrid"`
	Interface    basetypes.ListValue    `tfsdk:"interface"`
	Multicast    basetypes.ObjectValue  `tfsdk:"multicast"`
	Name         basetypes.StringValue  `tfsdk:"name"`
	Ospf         basetypes.ObjectValue  `tfsdk:"ospf"`
	Ospfv3       basetypes.ObjectValue  `tfsdk:"ospfv3"`
	RibFilter    basetypes.ObjectValue  `tfsdk:"rib_filter"`
	Rip          basetypes.ObjectValue  `tfsdk:"rip"`
	RoutingTable basetypes.ObjectValue  `tfsdk:"routing_table"`
	SdwanType    basetypes.StringValue  `tfsdk:"sdwan_type"`
	VrAdminDists basetypes.ObjectValue  `tfsdk:"vr_admin_dists"`
	ZoneName     basetypes.StringValue  `tfsdk:"zone_name"`
}

// LogicalRoutersVrfInnerAdminDists represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerAdminDists struct {
	BgpExternal basetypes.Float64Value `tfsdk:"bgp_external"`
	BgpInternal basetypes.Float64Value `tfsdk:"bgp_internal"`
	BgpLocal    basetypes.Float64Value `tfsdk:"bgp_local"`
	OspfExt     basetypes.Float64Value `tfsdk:"ospf_ext"`
	OspfInter   basetypes.Float64Value `tfsdk:"ospf_inter"`
	OspfIntra   basetypes.Float64Value `tfsdk:"ospf_intra"`
	Ospfv3Ext   basetypes.Float64Value `tfsdk:"ospfv3_ext"`
	Ospfv3Inter basetypes.Float64Value `tfsdk:"ospfv3_inter"`
	Ospfv3Intra basetypes.Float64Value `tfsdk:"ospfv3_intra"`
	Rip         basetypes.Float64Value `tfsdk:"rip"`
	Static      basetypes.Float64Value `tfsdk:"static"`
	StaticIpv6  basetypes.Float64Value `tfsdk:"static_ipv6"`
}

// LogicalRoutersVrfInnerBgp represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgp struct {
	AdvertiseNetwork            basetypes.ObjectValue  `tfsdk:"advertise_network"`
	Aggregate                   basetypes.ObjectValue  `tfsdk:"aggregate"`
	AggregateRoutes             basetypes.ListValue    `tfsdk:"aggregate_routes"`
	AllowRedistDefaultRoute     basetypes.BoolValue    `tfsdk:"allow_redist_default_route"`
	AlwaysAdvertiseNetworkRoute basetypes.BoolValue    `tfsdk:"always_advertise_network_route"`
	AsFormat                    basetypes.StringValue  `tfsdk:"as_format"`
	ConfederationMemberAs       basetypes.StringValue  `tfsdk:"confederation_member_as"`
	DefaultLocalPreference      basetypes.Float64Value `tfsdk:"default_local_preference"`
	EcmpMultiAs                 basetypes.BoolValue    `tfsdk:"ecmp_multi_as"`
	Enable                      basetypes.BoolValue    `tfsdk:"enable"`
	EnforceFirstAs              basetypes.BoolValue    `tfsdk:"enforce_first_as"`
	FastExternalFailover        basetypes.BoolValue    `tfsdk:"fast_external_failover"`
	GlobalBfd                   basetypes.ObjectValue  `tfsdk:"global_bfd"`
	GracefulRestart             basetypes.ObjectValue  `tfsdk:"graceful_restart"`
	GracefulShutdown            basetypes.BoolValue    `tfsdk:"graceful_shutdown"`
	InstallRoute                basetypes.BoolValue    `tfsdk:"install_route"`
	LocalAs                     basetypes.StringValue  `tfsdk:"local_as"`
	Med                         basetypes.ObjectValue  `tfsdk:"med"`
	PeerGroup                   basetypes.ListValue    `tfsdk:"peer_group"`
	Policy                      basetypes.ObjectValue  `tfsdk:"policy"`
	RedistRules                 basetypes.ListValue    `tfsdk:"redist_rules"`
	RedistributionProfile       basetypes.ObjectValue  `tfsdk:"redistribution_profile"`
	RejectDefaultRoute          basetypes.BoolValue    `tfsdk:"reject_default_route"`
	RouterId                    basetypes.StringValue  `tfsdk:"router_id"`
}

// LogicalRoutersVrfInnerBgpAdvertiseNetwork represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpAdvertiseNetwork struct {
	Ipv4 basetypes.ObjectValue `tfsdk:"ipv4"`
	Ipv6 basetypes.ObjectValue `tfsdk:"ipv6"`
}

// LogicalRoutersVrfInnerBgpAdvertiseNetworkIpv4 represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpAdvertiseNetworkIpv4 struct {
	Network basetypes.ListValue `tfsdk:"network"`
}

// LogicalRoutersVrfInnerBgpAdvertiseNetworkIpv4NetworkInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpAdvertiseNetworkIpv4NetworkInner struct {
	Backdoor  basetypes.BoolValue   `tfsdk:"backdoor"`
	Multicast basetypes.BoolValue   `tfsdk:"multicast"`
	Name      basetypes.StringValue `tfsdk:"name"`
	Unicast   basetypes.BoolValue   `tfsdk:"unicast"`
}

// LogicalRoutersVrfInnerBgpAdvertiseNetworkIpv6 represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpAdvertiseNetworkIpv6 struct {
	Network basetypes.ListValue `tfsdk:"network"`
}

// LogicalRoutersVrfInnerBgpAdvertiseNetworkIpv6NetworkInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpAdvertiseNetworkIpv6NetworkInner struct {
	Name    basetypes.StringValue `tfsdk:"name"`
	Unicast basetypes.BoolValue   `tfsdk:"unicast"`
}

// LogicalRoutersVrfInnerBgpAggregate represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpAggregate struct {
	AggregateMed basetypes.BoolValue `tfsdk:"aggregate_med"`
}

// LogicalRoutersVrfInnerBgpAggregateRoutesInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpAggregateRoutesInner struct {
	AsSet       basetypes.BoolValue   `tfsdk:"as_set"`
	Description basetypes.StringValue `tfsdk:"description"`
	Enable      basetypes.BoolValue   `tfsdk:"enable"`
	Name        basetypes.StringValue `tfsdk:"name"`
	SameMed     basetypes.BoolValue   `tfsdk:"same_med"`
	SummaryOnly basetypes.BoolValue   `tfsdk:"summary_only"`
	Type        basetypes.ObjectValue `tfsdk:"type"`
}

// LogicalRoutersVrfInnerBgpAggregateRoutesInnerType represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpAggregateRoutesInnerType struct {
	Ipv4 basetypes.ObjectValue `tfsdk:"ipv4"`
	Ipv6 basetypes.ObjectValue `tfsdk:"ipv6"`
}

// LogicalRoutersVrfInnerBgpAggregateRoutesInnerTypeIpv4 represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpAggregateRoutesInnerTypeIpv4 struct {
	AttributeMap  basetypes.StringValue `tfsdk:"attribute_map"`
	SummaryPrefix basetypes.StringValue `tfsdk:"summary_prefix"`
	SuppressMap   basetypes.StringValue `tfsdk:"suppress_map"`
}

// LogicalRoutersVrfInnerBgpGlobalBfd represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpGlobalBfd struct {
	Profile basetypes.StringValue `tfsdk:"profile"`
}

// LogicalRoutersVrfInnerBgpGracefulRestart represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpGracefulRestart struct {
	Enable             basetypes.BoolValue    `tfsdk:"enable"`
	LocalRestartTime   basetypes.Float64Value `tfsdk:"local_restart_time"`
	MaxPeerRestartTime basetypes.Float64Value `tfsdk:"max_peer_restart_time"`
	StaleRouteTime     basetypes.Float64Value `tfsdk:"stale_route_time"`
}

// LogicalRoutersVrfInnerBgpMed represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpMed struct {
	AlwaysCompareMed           basetypes.BoolValue `tfsdk:"always_compare_med"`
	DeterministicMedComparison basetypes.BoolValue `tfsdk:"deterministic_med_comparison"`
}

// LogicalRoutersVrfInnerBgpPeerGroupInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPeerGroupInner struct {
	AddressFamily           basetypes.ObjectValue `tfsdk:"address_family"`
	AggregatedConfedAsPath  basetypes.BoolValue   `tfsdk:"aggregated_confed_as_path"`
	ConnectionOptions       basetypes.ObjectValue `tfsdk:"connection_options"`
	Enable                  basetypes.BoolValue   `tfsdk:"enable"`
	FilteringProfile        basetypes.ObjectValue `tfsdk:"filtering_profile"`
	Name                    basetypes.StringValue `tfsdk:"name"`
	Peer                    basetypes.ListValue   `tfsdk:"peer"`
	SoftResetWithStoredInfo basetypes.BoolValue   `tfsdk:"soft_reset_with_stored_info"`
	Type                    basetypes.ObjectValue `tfsdk:"type"`
}

// LogicalRoutersVrfInnerBgpPeerGroupInnerAddressFamily represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPeerGroupInnerAddressFamily struct {
	Ipv4 basetypes.StringValue `tfsdk:"ipv4"`
	Ipv6 basetypes.StringValue `tfsdk:"ipv6"`
}

// LogicalRoutersVrfInnerBgpPeerGroupInnerConnectionOptions represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPeerGroupInnerConnectionOptions struct {
	Authentication basetypes.StringValue  `tfsdk:"authentication"`
	Dampening      basetypes.StringValue  `tfsdk:"dampening"`
	Multihop       basetypes.Float64Value `tfsdk:"multihop"`
	Timers         basetypes.StringValue  `tfsdk:"timers"`
}

// LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInner struct {
	Bfd                               basetypes.ObjectValue `tfsdk:"bfd"`
	ConnectionOptions                 basetypes.ObjectValue `tfsdk:"connection_options"`
	Enable                            basetypes.BoolValue   `tfsdk:"enable"`
	EnableMpBgp                       basetypes.BoolValue   `tfsdk:"enable_mp_bgp"`
	EnableSenderSideLoopDetection     basetypes.BoolValue   `tfsdk:"enable_sender_side_loop_detection"`
	Inherit                           basetypes.ObjectValue `tfsdk:"inherit"`
	LocalAddress                      basetypes.ObjectValue `tfsdk:"local_address"`
	Name                              basetypes.StringValue `tfsdk:"name"`
	Passive                           basetypes.BoolValue   `tfsdk:"passive"`
	PeerAddress                       basetypes.ObjectValue `tfsdk:"peer_address"`
	PeerAs                            basetypes.StringValue `tfsdk:"peer_as"`
	PeeringType                       basetypes.StringValue `tfsdk:"peering_type"`
	ReflectorClient                   basetypes.StringValue `tfsdk:"reflector_client"`
	SubsequentAddressFamilyIdentifier basetypes.ObjectValue `tfsdk:"subsequent_address_family_identifier"`
}

// LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerBfd represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerBfd struct {
	Multihop basetypes.ObjectValue `tfsdk:"multihop"`
	Profile  basetypes.StringValue `tfsdk:"profile"`
}

// LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerBfdMultihop represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerBfdMultihop struct {
	MinReceivedTtl basetypes.Float64Value `tfsdk:"min_received_ttl"`
}

// LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerConnectionOptions represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerConnectionOptions struct {
	Authentication        basetypes.StringValue  `tfsdk:"authentication"`
	Dampening             basetypes.StringValue  `tfsdk:"dampening"`
	HoldTime              basetypes.StringValue  `tfsdk:"hold_time"`
	IdleHoldTime          basetypes.Float64Value `tfsdk:"idle_hold_time"`
	IncomingBgpConnection basetypes.ObjectValue  `tfsdk:"incoming_bgp_connection"`
	KeepAliveInterval     basetypes.StringValue  `tfsdk:"keep_alive_interval"`
	MaxPrefixes           basetypes.StringValue  `tfsdk:"max_prefixes"`
	MinRouteAdvInterval   basetypes.Float64Value `tfsdk:"min_route_adv_interval"`
	Multihop              basetypes.StringValue  `tfsdk:"multihop"`
	OpenDelayTime         basetypes.Float64Value `tfsdk:"open_delay_time"`
	OutgoingBgpConnection basetypes.ObjectValue  `tfsdk:"outgoing_bgp_connection"`
	Timers                basetypes.StringValue  `tfsdk:"timers"`
}

// LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerConnectionOptionsIncomingBgpConnection represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerConnectionOptionsIncomingBgpConnection struct {
	Allow      basetypes.BoolValue    `tfsdk:"allow"`
	RemotePort basetypes.Float64Value `tfsdk:"remote_port"`
}

// LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerConnectionOptionsOutgoingBgpConnection represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerConnectionOptionsOutgoingBgpConnection struct {
	Allow     basetypes.BoolValue    `tfsdk:"allow"`
	LocalPort basetypes.Float64Value `tfsdk:"local_port"`
}

// LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerInherit represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerInherit struct {
	No  basetypes.ObjectValue `tfsdk:"no"`
	Yes basetypes.ObjectValue `tfsdk:"yes"`
}

// LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerInheritNo represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerInheritNo struct {
	AddressFamily    basetypes.ObjectValue `tfsdk:"address_family"`
	FilteringProfile basetypes.ObjectValue `tfsdk:"filtering_profile"`
}

// LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerLocalAddress represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerLocalAddress struct {
	Interface basetypes.StringValue `tfsdk:"interface"`
	Ip        basetypes.StringValue `tfsdk:"ip"`
}

// LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerPeerAddress represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerPeerAddress struct {
	Fqdn basetypes.StringValue `tfsdk:"fqdn"`
	Ip   basetypes.StringValue `tfsdk:"ip"`
}

// LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerSubsequentAddressFamilyIdentifier represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerSubsequentAddressFamilyIdentifier struct {
	Multicast basetypes.BoolValue `tfsdk:"multicast"`
	Unicast   basetypes.BoolValue `tfsdk:"unicast"`
}

// LogicalRoutersVrfInnerBgpPeerGroupInnerType represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPeerGroupInnerType struct {
	Ebgp       basetypes.ObjectValue `tfsdk:"ebgp"`
	EbgpConfed basetypes.ObjectValue `tfsdk:"ebgp_confed"`
	Ibgp       basetypes.ObjectValue `tfsdk:"ibgp"`
	IbgpConfed basetypes.ObjectValue `tfsdk:"ibgp_confed"`
}

// LogicalRoutersVrfInnerBgpPeerGroupInnerTypeEbgp represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPeerGroupInnerTypeEbgp struct {
	ExportNexthop   basetypes.StringValue `tfsdk:"export_nexthop"`
	ImportNexthop   basetypes.StringValue `tfsdk:"import_nexthop"`
	RemovePrivateAs basetypes.BoolValue   `tfsdk:"remove_private_as"`
}

// LogicalRoutersVrfInnerBgpPeerGroupInnerTypeEbgpConfed represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPeerGroupInnerTypeEbgpConfed struct {
	ExportNexthop basetypes.StringValue `tfsdk:"export_nexthop"`
}

// LogicalRoutersVrfInnerBgpPolicy represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPolicy struct {
	Aggregation              basetypes.ObjectValue `tfsdk:"aggregation"`
	ConditionalAdvertisement basetypes.ObjectValue `tfsdk:"conditional_advertisement"`
	Export                   basetypes.ObjectValue `tfsdk:"export"`
	Import                   basetypes.ObjectValue `tfsdk:"import"`
}

// LogicalRoutersVrfInnerBgpPolicyAggregation represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPolicyAggregation struct {
	Address basetypes.ListValue `tfsdk:"address"`
}

// LogicalRoutersVrfInnerBgpPolicyAggregationAddressInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPolicyAggregationAddressInner struct {
	AdvertiseFilters         basetypes.ListValue   `tfsdk:"advertise_filters"`
	AggregateRouteAttributes basetypes.ObjectValue `tfsdk:"aggregate_route_attributes"`
	AsSet                    basetypes.BoolValue   `tfsdk:"as_set"`
	Enable                   basetypes.BoolValue   `tfsdk:"enable"`
	Name                     basetypes.StringValue `tfsdk:"name"`
	Prefix                   basetypes.StringValue `tfsdk:"prefix"`
	Summary                  basetypes.BoolValue   `tfsdk:"summary"`
	SuppressFilters          basetypes.ListValue   `tfsdk:"suppress_filters"`
}

// LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAdvertiseFiltersInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAdvertiseFiltersInner struct {
	Enable basetypes.BoolValue   `tfsdk:"enable"`
	Match  basetypes.ObjectValue `tfsdk:"match"`
	Name   basetypes.StringValue `tfsdk:"name"`
}

// LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAdvertiseFiltersInnerMatch represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAdvertiseFiltersInnerMatch struct {
	AddressPrefix     basetypes.ListValue    `tfsdk:"address_prefix"`
	Afi               basetypes.StringValue  `tfsdk:"afi"`
	AsPath            basetypes.ObjectValue  `tfsdk:"as_path"`
	Community         basetypes.ObjectValue  `tfsdk:"community"`
	ExtendedCommunity basetypes.ObjectValue  `tfsdk:"extended_community"`
	FromPeer          basetypes.ListValue    `tfsdk:"from_peer"`
	Med               basetypes.Float64Value `tfsdk:"med"`
	Nexthop           basetypes.ListValue    `tfsdk:"nexthop"`
	RouteTable        basetypes.StringValue  `tfsdk:"route_table"`
	Safi              basetypes.StringValue  `tfsdk:"safi"`
}

// LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAdvertiseFiltersInnerMatchAddressPrefixInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAdvertiseFiltersInnerMatchAddressPrefixInner struct {
	Exact basetypes.BoolValue   `tfsdk:"exact"`
	Name  basetypes.StringValue `tfsdk:"name"`
}

// LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAdvertiseFiltersInnerMatchAsPath represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAdvertiseFiltersInnerMatchAsPath struct {
	Regex basetypes.StringValue `tfsdk:"regex"`
}

// LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAggregateRouteAttributes represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAggregateRouteAttributes struct {
	AsPath            basetypes.ObjectValue  `tfsdk:"as_path"`
	AsPathLimit       basetypes.Float64Value `tfsdk:"as_path_limit"`
	Community         basetypes.ObjectValue  `tfsdk:"community"`
	ExtendedCommunity basetypes.ObjectValue  `tfsdk:"extended_community"`
	LocalPreference   basetypes.Float64Value `tfsdk:"local_preference"`
	Med               basetypes.Float64Value `tfsdk:"med"`
	Nexthop           basetypes.StringValue  `tfsdk:"nexthop"`
	Origin            basetypes.StringValue  `tfsdk:"origin"`
	Weight            basetypes.Float64Value `tfsdk:"weight"`
}

// LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAggregateRouteAttributesAsPath represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAggregateRouteAttributesAsPath struct {
	None             basetypes.ObjectValue  `tfsdk:"none"`
	Prepend          basetypes.Float64Value `tfsdk:"prepend"`
	Remove           basetypes.ObjectValue  `tfsdk:"remove"`
	RemoveAndPrepend basetypes.Float64Value `tfsdk:"remove_and_prepend"`
}

// LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAggregateRouteAttributesCommunity represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAggregateRouteAttributesCommunity struct {
	Append      basetypes.ListValue   `tfsdk:"append"`
	None        basetypes.ObjectValue `tfsdk:"none"`
	Overwrite   basetypes.ListValue   `tfsdk:"overwrite"`
	RemoveAll   basetypes.ObjectValue `tfsdk:"remove_all"`
	RemoveRegex basetypes.StringValue `tfsdk:"remove_regex"`
}

// LogicalRoutersVrfInnerBgpPolicyConditionalAdvertisement represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPolicyConditionalAdvertisement struct {
	Policy basetypes.ListValue `tfsdk:"policy"`
}

// LogicalRoutersVrfInnerBgpPolicyConditionalAdvertisementPolicyInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPolicyConditionalAdvertisementPolicyInner struct {
	AdvertiseFilters basetypes.ListValue   `tfsdk:"advertise_filters"`
	Enable           basetypes.BoolValue   `tfsdk:"enable"`
	Name             basetypes.StringValue `tfsdk:"name"`
	NonExistFilters  basetypes.ListValue   `tfsdk:"non_exist_filters"`
	UsedBy           basetypes.ListValue   `tfsdk:"used_by"`
}

// LogicalRoutersVrfInnerBgpPolicyExport represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPolicyExport struct {
	Rules basetypes.ListValue `tfsdk:"rules"`
}

// LogicalRoutersVrfInnerBgpPolicyExportRulesInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPolicyExportRulesInner struct {
	Action basetypes.ObjectValue `tfsdk:"action"`
	Enable basetypes.BoolValue   `tfsdk:"enable"`
	Match  basetypes.ObjectValue `tfsdk:"match"`
	Name   basetypes.StringValue `tfsdk:"name"`
	UsedBy basetypes.ListValue   `tfsdk:"used_by"`
}

// LogicalRoutersVrfInnerBgpPolicyExportRulesInnerAction represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPolicyExportRulesInnerAction struct {
	Allow basetypes.ObjectValue `tfsdk:"allow"`
	Deny  basetypes.ObjectValue `tfsdk:"deny"`
}

// LogicalRoutersVrfInnerBgpPolicyExportRulesInnerActionAllow represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPolicyExportRulesInnerActionAllow struct {
	Update basetypes.ObjectValue `tfsdk:"update"`
}

// LogicalRoutersVrfInnerBgpPolicyExportRulesInnerActionAllowUpdate represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPolicyExportRulesInnerActionAllowUpdate struct {
	AsPath            basetypes.ObjectValue  `tfsdk:"as_path"`
	AsPathLimit       basetypes.Float64Value `tfsdk:"as_path_limit"`
	Community         basetypes.ObjectValue  `tfsdk:"community"`
	ExtendedCommunity basetypes.ObjectValue  `tfsdk:"extended_community"`
	LocalPreference   basetypes.Float64Value `tfsdk:"local_preference"`
	Med               basetypes.Float64Value `tfsdk:"med"`
	Nexthop           basetypes.StringValue  `tfsdk:"nexthop"`
	Origin            basetypes.StringValue  `tfsdk:"origin"`
}

// LogicalRoutersVrfInnerBgpPolicyExportRulesInnerMatch represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPolicyExportRulesInnerMatch struct {
	AddressPrefix     basetypes.ListValue    `tfsdk:"address_prefix"`
	Afi               basetypes.StringValue  `tfsdk:"afi"`
	AsPath            basetypes.ObjectValue  `tfsdk:"as_path"`
	Community         basetypes.ObjectValue  `tfsdk:"community"`
	ExtendedCommunity basetypes.ObjectValue  `tfsdk:"extended_community"`
	FromPeer          basetypes.ListValue    `tfsdk:"from_peer"`
	Med               basetypes.Float64Value `tfsdk:"med"`
	Nexthop           basetypes.ListValue    `tfsdk:"nexthop"`
	RouteTable        basetypes.StringValue  `tfsdk:"route_table"`
	Safi              basetypes.StringValue  `tfsdk:"safi"`
}

// LogicalRoutersVrfInnerBgpPolicyExportRulesInnerMatchAddressPrefixInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPolicyExportRulesInnerMatchAddressPrefixInner struct {
	Exact basetypes.BoolValue   `tfsdk:"exact"`
	Name  basetypes.StringValue `tfsdk:"name"`
}

// LogicalRoutersVrfInnerBgpPolicyImport represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPolicyImport struct {
	Rules basetypes.ListValue `tfsdk:"rules"`
}

// LogicalRoutersVrfInnerBgpPolicyImportRulesInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPolicyImportRulesInner struct {
	Action basetypes.ObjectValue `tfsdk:"action"`
	Enable basetypes.BoolValue   `tfsdk:"enable"`
	Match  basetypes.ObjectValue `tfsdk:"match"`
	Name   basetypes.StringValue `tfsdk:"name"`
	UsedBy basetypes.ListValue   `tfsdk:"used_by"`
}

// LogicalRoutersVrfInnerBgpPolicyImportRulesInnerAction represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPolicyImportRulesInnerAction struct {
	Allow basetypes.ObjectValue `tfsdk:"allow"`
	Deny  basetypes.ObjectValue `tfsdk:"deny"`
}

// LogicalRoutersVrfInnerBgpPolicyImportRulesInnerActionAllow represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpPolicyImportRulesInnerActionAllow struct {
	Dampening basetypes.StringValue `tfsdk:"dampening"`
	Update    basetypes.ObjectValue `tfsdk:"update"`
}

// LogicalRoutersVrfInnerBgpRedistRulesInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpRedistRulesInner struct {
	AddressFamilyIdentifier basetypes.StringValue  `tfsdk:"address_family_identifier"`
	Enable                  basetypes.BoolValue    `tfsdk:"enable"`
	Metric                  basetypes.Float64Value `tfsdk:"metric"`
	Name                    basetypes.StringValue  `tfsdk:"name"`
	RouteTable              basetypes.StringValue  `tfsdk:"route_table"`
	SetAsPathLimit          basetypes.Float64Value `tfsdk:"set_as_path_limit"`
	SetCommunity            basetypes.ListValue    `tfsdk:"set_community"`
	SetExtendedCommunity    basetypes.ListValue    `tfsdk:"set_extended_community"`
	SetLocalPreference      basetypes.Float64Value `tfsdk:"set_local_preference"`
	SetMed                  basetypes.Float64Value `tfsdk:"set_med"`
	SetOrigin               basetypes.StringValue  `tfsdk:"set_origin"`
}

// LogicalRoutersVrfInnerBgpRedistributionProfile represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpRedistributionProfile struct {
	Ipv4 basetypes.ObjectValue `tfsdk:"ipv4"`
	Ipv6 basetypes.ObjectValue `tfsdk:"ipv6"`
}

// LogicalRoutersVrfInnerBgpRedistributionProfileIpv4 represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerBgpRedistributionProfileIpv4 struct {
	Unicast basetypes.StringValue `tfsdk:"unicast"`
}

// LogicalRoutersVrfInnerEcmp represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerEcmp struct {
	Algorithm        basetypes.ObjectValue  `tfsdk:"algorithm"`
	Enable           basetypes.BoolValue    `tfsdk:"enable"`
	MaxPath          basetypes.Float64Value `tfsdk:"max_path"`
	StrictSourcePath basetypes.BoolValue    `tfsdk:"strict_source_path"`
	SymmetricReturn  basetypes.BoolValue    `tfsdk:"symmetric_return"`
}

// LogicalRoutersVrfInnerEcmpAlgorithm represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerEcmpAlgorithm struct {
	BalancedRoundRobin basetypes.ObjectValue `tfsdk:"balanced_round_robin"`
	IpHash             basetypes.ObjectValue `tfsdk:"ip_hash"`
	IpModulo           basetypes.ObjectValue `tfsdk:"ip_modulo"`
	WeightedRoundRobin basetypes.ObjectValue `tfsdk:"weighted_round_robin"`
}

// LogicalRoutersVrfInnerEcmpAlgorithmIpHash represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerEcmpAlgorithmIpHash struct {
	HashSeed basetypes.Float64Value `tfsdk:"hash_seed"`
	SrcOnly  basetypes.BoolValue    `tfsdk:"src_only"`
	UsePort  basetypes.BoolValue    `tfsdk:"use_port"`
}

// LogicalRoutersVrfInnerEcmpAlgorithmWeightedRoundRobin represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerEcmpAlgorithmWeightedRoundRobin struct {
	Interface basetypes.ListValue `tfsdk:"interface"`
}

// LogicalRoutersVrfInnerEcmpAlgorithmWeightedRoundRobinInterfaceInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerEcmpAlgorithmWeightedRoundRobinInterfaceInner struct {
	Name   basetypes.StringValue  `tfsdk:"name"`
	Weight basetypes.Float64Value `tfsdk:"weight"`
}

// LogicalRoutersVrfInnerMulticast represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticast struct {
	Enable          basetypes.BoolValue    `tfsdk:"enable"`
	EnableV6        basetypes.BoolValue    `tfsdk:"enable_v6"`
	Igmp            basetypes.ObjectValue  `tfsdk:"igmp"`
	InterfaceGroup  basetypes.ListValue    `tfsdk:"interface_group"`
	Mode            basetypes.StringValue  `tfsdk:"mode"`
	Msdp            basetypes.ObjectValue  `tfsdk:"msdp"`
	Pim             basetypes.ObjectValue  `tfsdk:"pim"`
	RouteAgeoutTime basetypes.Float64Value `tfsdk:"route_ageout_time"`
	Rp              basetypes.ObjectValue  `tfsdk:"rp"`
	SptThreshold    basetypes.ListValue    `tfsdk:"spt_threshold"`
	SsmAddressSpace basetypes.ListValue    `tfsdk:"ssm_address_space"`
	StaticRoute     basetypes.ListValue    `tfsdk:"static_route"`
}

// LogicalRoutersVrfInnerMulticastIgmp represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastIgmp struct {
	Dynamic basetypes.ObjectValue `tfsdk:"dynamic"`
	Enable  basetypes.BoolValue   `tfsdk:"enable"`
	Static  basetypes.ListValue   `tfsdk:"static"`
}

// LogicalRoutersVrfInnerMulticastIgmpDynamic represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastIgmpDynamic struct {
	Interface basetypes.ListValue `tfsdk:"interface"`
}

// LogicalRoutersVrfInnerMulticastIgmpDynamicInterfaceInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastIgmpDynamicInterfaceInner struct {
	GroupFilter         basetypes.StringValue `tfsdk:"group_filter"`
	MaxGroups           basetypes.StringValue `tfsdk:"max_groups"`
	MaxSources          basetypes.StringValue `tfsdk:"max_sources"`
	Name                basetypes.StringValue `tfsdk:"name"`
	QueryProfile        basetypes.StringValue `tfsdk:"query_profile"`
	Robustness          basetypes.StringValue `tfsdk:"robustness"`
	RouterAlertPolicing basetypes.BoolValue   `tfsdk:"router_alert_policing"`
	Version             basetypes.StringValue `tfsdk:"version"`
}

// LogicalRoutersVrfInnerMulticastIgmpStaticInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastIgmpStaticInner struct {
	GroupAddress  basetypes.StringValue `tfsdk:"group_address"`
	Interface     basetypes.StringValue `tfsdk:"interface"`
	Name          basetypes.StringValue `tfsdk:"name"`
	SourceAddress basetypes.StringValue `tfsdk:"source_address"`
}

// LogicalRoutersVrfInnerMulticastInterfaceGroupInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastInterfaceGroupInner struct {
	Description     basetypes.StringValue `tfsdk:"description"`
	GroupPermission basetypes.ObjectValue `tfsdk:"group_permission"`
	Igmp            basetypes.ObjectValue `tfsdk:"igmp"`
	Interface       basetypes.ListValue   `tfsdk:"interface"`
	Name            basetypes.StringValue `tfsdk:"name"`
	Pim             basetypes.ObjectValue `tfsdk:"pim"`
}

// LogicalRoutersVrfInnerMulticastInterfaceGroupInnerGroupPermission represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastInterfaceGroupInnerGroupPermission struct {
	AnySourceMulticast      basetypes.ListValue `tfsdk:"any_source_multicast"`
	SourceSpecificMulticast basetypes.ListValue `tfsdk:"source_specific_multicast"`
}

// LogicalRoutersVrfInnerMulticastInterfaceGroupInnerGroupPermissionAnySourceMulticastInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastInterfaceGroupInnerGroupPermissionAnySourceMulticastInner struct {
	GroupAddress basetypes.StringValue `tfsdk:"group_address"`
	Included     basetypes.BoolValue   `tfsdk:"included"`
	Name         basetypes.StringValue `tfsdk:"name"`
}

// LogicalRoutersVrfInnerMulticastInterfaceGroupInnerGroupPermissionSourceSpecificMulticastInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastInterfaceGroupInnerGroupPermissionSourceSpecificMulticastInner struct {
	GroupAddress  basetypes.StringValue `tfsdk:"group_address"`
	Included      basetypes.BoolValue   `tfsdk:"included"`
	Name          basetypes.StringValue `tfsdk:"name"`
	SourceAddress basetypes.StringValue `tfsdk:"source_address"`
}

// LogicalRoutersVrfInnerMulticastInterfaceGroupInnerIgmp represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastInterfaceGroupInnerIgmp struct {
	Enable                  basetypes.BoolValue    `tfsdk:"enable"`
	ImmediateLeave          basetypes.BoolValue    `tfsdk:"immediate_leave"`
	LastMemberQueryInterval basetypes.Float64Value `tfsdk:"last_member_query_interval"`
	MaxGroups               basetypes.StringValue  `tfsdk:"max_groups"`
	MaxQueryResponseTime    basetypes.Float64Value `tfsdk:"max_query_response_time"`
	MaxSources              basetypes.StringValue  `tfsdk:"max_sources"`
	Mode                    basetypes.StringValue  `tfsdk:"mode"`
	QueryInterval           basetypes.Float64Value `tfsdk:"query_interval"`
	Robustness              basetypes.StringValue  `tfsdk:"robustness"`
	RouterAlertPolicing     basetypes.BoolValue    `tfsdk:"router_alert_policing"`
	Version                 basetypes.StringValue  `tfsdk:"version"`
}

// LogicalRoutersVrfInnerMulticastInterfaceGroupInnerPim represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastInterfaceGroupInnerPim struct {
	AllowedNeighbors  basetypes.ListValue    `tfsdk:"allowed_neighbors"`
	AssertInterval    basetypes.Float64Value `tfsdk:"assert_interval"`
	BsrBorder         basetypes.BoolValue    `tfsdk:"bsr_border"`
	DrPriority        basetypes.Float64Value `tfsdk:"dr_priority"`
	Enable            basetypes.BoolValue    `tfsdk:"enable"`
	HelloInterval     basetypes.Float64Value `tfsdk:"hello_interval"`
	JoinPruneInterval basetypes.Float64Value `tfsdk:"join_prune_interval"`
}

// LogicalRoutersVrfInnerMulticastInterfaceGroupInnerPimAllowedNeighborsInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastInterfaceGroupInnerPimAllowedNeighborsInner struct {
	Name basetypes.StringValue `tfsdk:"name"`
}

// LogicalRoutersVrfInnerMulticastMsdp represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastMsdp struct {
	Enable               basetypes.BoolValue   `tfsdk:"enable"`
	GlobalAuthentication basetypes.StringValue `tfsdk:"global_authentication"`
	GlobalTimer          basetypes.StringValue `tfsdk:"global_timer"`
	OriginatorId         basetypes.ObjectValue `tfsdk:"originator_id"`
	Peer                 basetypes.ListValue   `tfsdk:"peer"`
}

// LogicalRoutersVrfInnerMulticastMsdpPeerInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastMsdpPeerInner struct {
	Authentication   basetypes.StringValue  `tfsdk:"authentication"`
	Enable           basetypes.BoolValue    `tfsdk:"enable"`
	InboundSaFilter  basetypes.StringValue  `tfsdk:"inbound_sa_filter"`
	LocalAddress     basetypes.ObjectValue  `tfsdk:"local_address"`
	MaxSa            basetypes.Float64Value `tfsdk:"max_sa"`
	Name             basetypes.StringValue  `tfsdk:"name"`
	OutboundSaFilter basetypes.StringValue  `tfsdk:"outbound_sa_filter"`
	PeerAddress      basetypes.ObjectValue  `tfsdk:"peer_address"`
	PeerAs           basetypes.StringValue  `tfsdk:"peer_as"`
}

// LogicalRoutersVrfInnerMulticastPim represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastPim struct {
	Enable          basetypes.BoolValue    `tfsdk:"enable"`
	GroupPermission basetypes.StringValue  `tfsdk:"group_permission"`
	IfTimerGlobal   basetypes.StringValue  `tfsdk:"if_timer_global"`
	Interface       basetypes.ListValue    `tfsdk:"interface"`
	RouteAgeoutTime basetypes.Float64Value `tfsdk:"route_ageout_time"`
	Rp              basetypes.ObjectValue  `tfsdk:"rp"`
	RpfLookupMode   basetypes.StringValue  `tfsdk:"rpf_lookup_mode"`
	SptThreshold    basetypes.ListValue    `tfsdk:"spt_threshold"`
	SsmAddressSpace basetypes.ObjectValue  `tfsdk:"ssm_address_space"`
}

// LogicalRoutersVrfInnerMulticastPimInterfaceInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastPimInterfaceInner struct {
	Description    basetypes.StringValue  `tfsdk:"description"`
	DrPriority     basetypes.Float64Value `tfsdk:"dr_priority"`
	IfTimer        basetypes.StringValue  `tfsdk:"if_timer"`
	Name           basetypes.StringValue  `tfsdk:"name"`
	NeighborFilter basetypes.StringValue  `tfsdk:"neighbor_filter"`
	SendBsm        basetypes.BoolValue    `tfsdk:"send_bsm"`
}

// LogicalRoutersVrfInnerMulticastPimRp represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastPimRp struct {
	ExternalRp basetypes.ListValue   `tfsdk:"external_rp"`
	LocalRp    basetypes.ObjectValue `tfsdk:"local_rp"`
}

// LogicalRoutersVrfInnerMulticastPimRpExternalRpInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastPimRpExternalRpInner struct {
	GroupList basetypes.StringValue `tfsdk:"group_list"`
	Name      basetypes.StringValue `tfsdk:"name"`
	Override  basetypes.BoolValue   `tfsdk:"override"`
}

// LogicalRoutersVrfInnerMulticastPimRpLocalRp represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastPimRpLocalRp struct {
	CandidateRp basetypes.ObjectValue `tfsdk:"candidate_rp"`
	StaticRp    basetypes.ObjectValue `tfsdk:"static_rp"`
}

// LogicalRoutersVrfInnerMulticastPimRpLocalRpCandidateRp represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastPimRpLocalRpCandidateRp struct {
	Address               basetypes.StringValue  `tfsdk:"address"`
	AdvertisementInterval basetypes.Float64Value `tfsdk:"advertisement_interval"`
	GroupList             basetypes.StringValue  `tfsdk:"group_list"`
	Interface             basetypes.StringValue  `tfsdk:"interface"`
	Priority              basetypes.Float64Value `tfsdk:"priority"`
}

// LogicalRoutersVrfInnerMulticastPimRpLocalRpStaticRp represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastPimRpLocalRpStaticRp struct {
	Address   basetypes.StringValue `tfsdk:"address"`
	GroupList basetypes.StringValue `tfsdk:"group_list"`
	Interface basetypes.StringValue `tfsdk:"interface"`
	Override  basetypes.BoolValue   `tfsdk:"override"`
}

// LogicalRoutersVrfInnerMulticastPimSptThresholdInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastPimSptThresholdInner struct {
	Name      basetypes.StringValue `tfsdk:"name"`
	Threshold basetypes.StringValue `tfsdk:"threshold"`
}

// LogicalRoutersVrfInnerMulticastPimSsmAddressSpace represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastPimSsmAddressSpace struct {
	GroupList basetypes.StringValue `tfsdk:"group_list"`
}

// LogicalRoutersVrfInnerMulticastRp represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastRp struct {
	ExternalRp basetypes.ListValue   `tfsdk:"external_rp"`
	LocalRp    basetypes.ObjectValue `tfsdk:"local_rp"`
}

// LogicalRoutersVrfInnerMulticastRpExternalRpInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastRpExternalRpInner struct {
	GroupAddresses basetypes.ListValue   `tfsdk:"group_addresses"`
	Name           basetypes.StringValue `tfsdk:"name"`
	Override       basetypes.BoolValue   `tfsdk:"override"`
}

// LogicalRoutersVrfInnerMulticastRpLocalRp represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastRpLocalRp struct {
	CandidateRp basetypes.ObjectValue `tfsdk:"candidate_rp"`
	StaticRp    basetypes.ObjectValue `tfsdk:"static_rp"`
}

// LogicalRoutersVrfInnerMulticastRpLocalRpCandidateRp represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastRpLocalRpCandidateRp struct {
	Address               basetypes.StringValue  `tfsdk:"address"`
	AdvertisementInterval basetypes.Float64Value `tfsdk:"advertisement_interval"`
	GroupAddresses        basetypes.ListValue    `tfsdk:"group_addresses"`
	Interface             basetypes.StringValue  `tfsdk:"interface"`
	Priority              basetypes.Float64Value `tfsdk:"priority"`
}

// LogicalRoutersVrfInnerMulticastRpLocalRpStaticRp represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastRpLocalRpStaticRp struct {
	Address        basetypes.StringValue `tfsdk:"address"`
	GroupAddresses basetypes.ListValue   `tfsdk:"group_addresses"`
	Interface      basetypes.StringValue `tfsdk:"interface"`
	Override       basetypes.BoolValue   `tfsdk:"override"`
}

// LogicalRoutersVrfInnerMulticastStaticRouteInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastStaticRouteInner struct {
	Destination basetypes.StringValue  `tfsdk:"destination"`
	Interface   basetypes.StringValue  `tfsdk:"interface"`
	Name        basetypes.StringValue  `tfsdk:"name"`
	Nexthop     basetypes.ObjectValue  `tfsdk:"nexthop"`
	Preference  basetypes.Float64Value `tfsdk:"preference"`
}

// LogicalRoutersVrfInnerMulticastStaticRouteInnerNexthop represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerMulticastStaticRouteInnerNexthop struct {
	IpAddress basetypes.StringValue `tfsdk:"ip_address"`
}

// LogicalRoutersVrfInnerOspf represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspf struct {
	AllowRedistDefaultRoute basetypes.BoolValue   `tfsdk:"allow_redist_default_route"`
	Area                    basetypes.ListValue   `tfsdk:"area"`
	AuthProfile             basetypes.ListValue   `tfsdk:"auth_profile"`
	Enable                  basetypes.BoolValue   `tfsdk:"enable"`
	ExportRules             basetypes.ListValue   `tfsdk:"export_rules"`
	FloodPrevention         basetypes.ObjectValue `tfsdk:"flood_prevention"`
	GlobalBfd               basetypes.ObjectValue `tfsdk:"global_bfd"`
	GlobalIfTimer           basetypes.StringValue `tfsdk:"global_if_timer"`
	GracefulRestart         basetypes.ObjectValue `tfsdk:"graceful_restart"`
	RedistributionProfile   basetypes.StringValue `tfsdk:"redistribution_profile"`
	RejectDefaultRoute      basetypes.BoolValue   `tfsdk:"reject_default_route"`
	Rfc1583                 basetypes.BoolValue   `tfsdk:"rfc1583"`
	RouterId                basetypes.StringValue `tfsdk:"router_id"`
	SpfTimer                basetypes.StringValue `tfsdk:"spf_timer"`
	VrTimers                basetypes.ObjectValue `tfsdk:"vr_timers"`
}

// LogicalRoutersVrfInnerOspfAreaInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfAreaInner struct {
	Authentication basetypes.StringValue `tfsdk:"authentication"`
	Interface      basetypes.ListValue   `tfsdk:"interface"`
	Name           basetypes.StringValue `tfsdk:"name"`
	Range          basetypes.ListValue   `tfsdk:"range"`
	Type           basetypes.ObjectValue `tfsdk:"type"`
	VirtualLink    basetypes.ListValue   `tfsdk:"virtual_link"`
	VrRange        basetypes.ListValue   `tfsdk:"vr_range"`
}

// LogicalRoutersVrfInnerOspfAreaInnerInterfaceInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfAreaInnerInterfaceInner struct {
	Authentication basetypes.StringValue  `tfsdk:"authentication"`
	Bfd            basetypes.ObjectValue  `tfsdk:"bfd"`
	Enable         basetypes.BoolValue    `tfsdk:"enable"`
	LinkType       basetypes.ObjectValue  `tfsdk:"link_type"`
	Metric         basetypes.Float64Value `tfsdk:"metric"`
	MtuIgnore      basetypes.BoolValue    `tfsdk:"mtu_ignore"`
	Name           basetypes.StringValue  `tfsdk:"name"`
	Passive        basetypes.BoolValue    `tfsdk:"passive"`
	Priority       basetypes.Float64Value `tfsdk:"priority"`
	Timing         basetypes.StringValue  `tfsdk:"timing"`
	VrTiming       basetypes.ObjectValue  `tfsdk:"vr_timing"`
}

// LogicalRoutersVrfInnerOspfAreaInnerInterfaceInnerLinkType represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfAreaInnerInterfaceInnerLinkType struct {
	Broadcast basetypes.ObjectValue `tfsdk:"broadcast"`
	P2mp      basetypes.ObjectValue `tfsdk:"p2mp"`
	P2p       basetypes.ObjectValue `tfsdk:"p2p"`
}

// LogicalRoutersVrfInnerOspfAreaInnerInterfaceInnerLinkTypeP2mp represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfAreaInnerInterfaceInnerLinkTypeP2mp struct {
	Neighbor basetypes.ListValue `tfsdk:"neighbor"`
}

// LogicalRoutersVrfInnerOspfAreaInnerInterfaceInnerLinkTypeP2mpNeighborInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfAreaInnerInterfaceInnerLinkTypeP2mpNeighborInner struct {
	Name     basetypes.StringValue  `tfsdk:"name"`
	Priority basetypes.Float64Value `tfsdk:"priority"`
}

// LogicalRoutersVrfInnerOspfAreaInnerInterfaceInnerVrTiming represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfAreaInnerInterfaceInnerVrTiming struct {
	DeadCounts         basetypes.Float64Value `tfsdk:"dead_counts"`
	GrDelay            basetypes.Float64Value `tfsdk:"gr_delay"`
	HelloInterval      basetypes.Float64Value `tfsdk:"hello_interval"`
	RetransmitInterval basetypes.Float64Value `tfsdk:"retransmit_interval"`
	TransitDelay       basetypes.Float64Value `tfsdk:"transit_delay"`
}

// LogicalRoutersVrfInnerOspfAreaInnerRangeInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfAreaInnerRangeInner struct {
	Advertise  basetypes.BoolValue   `tfsdk:"advertise"`
	Name       basetypes.StringValue `tfsdk:"name"`
	Substitute basetypes.StringValue `tfsdk:"substitute"`
}

// LogicalRoutersVrfInnerOspfAreaInnerType represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfAreaInnerType struct {
	Normal basetypes.ObjectValue `tfsdk:"normal"`
	Nssa   basetypes.ObjectValue `tfsdk:"nssa"`
	Stub   basetypes.ObjectValue `tfsdk:"stub"`
}

// LogicalRoutersVrfInnerOspfAreaInnerTypeNormal represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfAreaInnerTypeNormal struct {
	Abr basetypes.ObjectValue `tfsdk:"abr"`
}

// LogicalRoutersVrfInnerOspfAreaInnerTypeNormalAbr represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfAreaInnerTypeNormalAbr struct {
	ExportList         basetypes.StringValue `tfsdk:"export_list"`
	ImportList         basetypes.StringValue `tfsdk:"import_list"`
	InboundFilterList  basetypes.StringValue `tfsdk:"inbound_filter_list"`
	OutboundFilterList basetypes.StringValue `tfsdk:"outbound_filter_list"`
}

// LogicalRoutersVrfInnerOspfAreaInnerTypeNssa represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfAreaInnerTypeNssa struct {
	Abr                         basetypes.ObjectValue `tfsdk:"abr"`
	AcceptSummary               basetypes.BoolValue   `tfsdk:"accept_summary"`
	DefaultInformationOriginate basetypes.ObjectValue `tfsdk:"default_information_originate"`
	DefaultRoute                basetypes.ObjectValue `tfsdk:"default_route"`
	NoSummary                   basetypes.BoolValue   `tfsdk:"no_summary"`
	NssaExtRange                basetypes.ListValue   `tfsdk:"nssa_ext_range"`
}

// LogicalRoutersVrfInnerOspfAreaInnerTypeNssaAbr represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfAreaInnerTypeNssaAbr struct {
	ExportList         basetypes.StringValue `tfsdk:"export_list"`
	ImportList         basetypes.StringValue `tfsdk:"import_list"`
	InboundFilterList  basetypes.StringValue `tfsdk:"inbound_filter_list"`
	NssaExtRange       basetypes.ListValue   `tfsdk:"nssa_ext_range"`
	OutboundFilterList basetypes.StringValue `tfsdk:"outbound_filter_list"`
}

// LogicalRoutersVrfInnerOspfAreaInnerTypeNssaAbrNssaExtRangeInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfAreaInnerTypeNssaAbrNssaExtRangeInner struct {
	Advertise basetypes.BoolValue    `tfsdk:"advertise"`
	Name      basetypes.StringValue  `tfsdk:"name"`
	RouteTag  basetypes.Float64Value `tfsdk:"route_tag"`
}

// LogicalRoutersVrfInnerOspfAreaInnerTypeNssaDefaultInformationOriginate represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfAreaInnerTypeNssaDefaultInformationOriginate struct {
	Metric     basetypes.Float64Value `tfsdk:"metric"`
	MetricType basetypes.StringValue  `tfsdk:"metric_type"`
}

// LogicalRoutersVrfInnerOspfAreaInnerTypeNssaDefaultRoute represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfAreaInnerTypeNssaDefaultRoute struct {
	Advertise basetypes.ObjectValue `tfsdk:"advertise"`
	Disable   basetypes.ObjectValue `tfsdk:"disable"`
}

// LogicalRoutersVrfInnerOspfAreaInnerTypeNssaDefaultRouteAdvertise represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfAreaInnerTypeNssaDefaultRouteAdvertise struct {
	Metric basetypes.Float64Value `tfsdk:"metric"`
	Type   basetypes.StringValue  `tfsdk:"type"`
}

// LogicalRoutersVrfInnerOspfAreaInnerTypeNssaNssaExtRangeInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfAreaInnerTypeNssaNssaExtRangeInner struct {
	Advertise basetypes.ObjectValue `tfsdk:"advertise"`
	Name      basetypes.StringValue `tfsdk:"name"`
	Suppress  basetypes.ObjectValue `tfsdk:"suppress"`
}

// LogicalRoutersVrfInnerOspfAreaInnerTypeStub represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfAreaInnerTypeStub struct {
	Abr                basetypes.ObjectValue  `tfsdk:"abr"`
	AcceptSummary      basetypes.BoolValue    `tfsdk:"accept_summary"`
	DefaultRoute       basetypes.ObjectValue  `tfsdk:"default_route"`
	DefaultRouteMetric basetypes.Float64Value `tfsdk:"default_route_metric"`
	NoSummary          basetypes.BoolValue    `tfsdk:"no_summary"`
}

// LogicalRoutersVrfInnerOspfAreaInnerTypeStubDefaultRoute represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfAreaInnerTypeStubDefaultRoute struct {
	Advertise basetypes.ObjectValue `tfsdk:"advertise"`
	Disable   basetypes.ObjectValue `tfsdk:"disable"`
}

// LogicalRoutersVrfInnerOspfAreaInnerTypeStubDefaultRouteAdvertise represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfAreaInnerTypeStubDefaultRouteAdvertise struct {
	Metric basetypes.Float64Value `tfsdk:"metric"`
}

// LogicalRoutersVrfInnerOspfAreaInnerVirtualLinkInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfAreaInnerVirtualLinkInner struct {
	Authentication basetypes.StringValue  `tfsdk:"authentication"`
	Bfd            basetypes.ObjectValue  `tfsdk:"bfd"`
	Enable         basetypes.BoolValue    `tfsdk:"enable"`
	InstanceId     basetypes.Float64Value `tfsdk:"instance_id"`
	InterfaceId    basetypes.Float64Value `tfsdk:"interface_id"`
	Name           basetypes.StringValue  `tfsdk:"name"`
	NeighborId     basetypes.StringValue  `tfsdk:"neighbor_id"`
	Passive        basetypes.BoolValue    `tfsdk:"passive"`
	Timing         basetypes.StringValue  `tfsdk:"timing"`
	TransitAreaId  basetypes.StringValue  `tfsdk:"transit_area_id"`
	VrTiming       basetypes.ObjectValue  `tfsdk:"vr_timing"`
}

// LogicalRoutersVrfInnerOspfAreaInnerVirtualLinkInnerVrTiming represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfAreaInnerVirtualLinkInnerVrTiming struct {
	DeadCounts         basetypes.Float64Value `tfsdk:"dead_counts"`
	HelloInterval      basetypes.Float64Value `tfsdk:"hello_interval"`
	RetransmitInterval basetypes.Float64Value `tfsdk:"retransmit_interval"`
	TransitDelay       basetypes.Float64Value `tfsdk:"transit_delay"`
}

// LogicalRoutersVrfInnerOspfAuthProfileInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfAuthProfileInner struct {
	Md5      basetypes.ListValue   `tfsdk:"md5"`
	Name     basetypes.StringValue `tfsdk:"name"`
	Password basetypes.StringValue `tfsdk:"password"`
}

// LogicalRoutersVrfInnerOspfAuthProfileInnerMd5Inner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfAuthProfileInnerMd5Inner struct {
	Key       basetypes.StringValue  `tfsdk:"key"`
	Name      basetypes.Float64Value `tfsdk:"name"`
	Preferred basetypes.BoolValue    `tfsdk:"preferred"`
}

// LogicalRoutersVrfInnerOspfExportRulesInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfExportRulesInner struct {
	Metric      basetypes.Float64Value `tfsdk:"metric"`
	Name        basetypes.StringValue  `tfsdk:"name"`
	NewPathType basetypes.StringValue  `tfsdk:"new_path_type"`
	NewTag      basetypes.StringValue  `tfsdk:"new_tag"`
}

// LogicalRoutersVrfInnerOspfFloodPrevention represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfFloodPrevention struct {
	Hello basetypes.ObjectValue `tfsdk:"hello"`
	Lsa   basetypes.ObjectValue `tfsdk:"lsa"`
}

// LogicalRoutersVrfInnerOspfFloodPreventionHello represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfFloodPreventionHello struct {
	Enable    basetypes.BoolValue    `tfsdk:"enable"`
	MaxPacket basetypes.Float64Value `tfsdk:"max_packet"`
}

// LogicalRoutersVrfInnerOspfGracefulRestart represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfGracefulRestart struct {
	Enable                 basetypes.BoolValue    `tfsdk:"enable"`
	GracePeriod            basetypes.Float64Value `tfsdk:"grace_period"`
	HelperEnable           basetypes.BoolValue    `tfsdk:"helper_enable"`
	MaxNeighborRestartTime basetypes.Float64Value `tfsdk:"max_neighbor_restart_time"`
	StrictLSAChecking      basetypes.BoolValue    `tfsdk:"strict__l_s_a_checking"`
}

// LogicalRoutersVrfInnerOspfVrTimers represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfVrTimers struct {
	LsaInterval         basetypes.Float64Value `tfsdk:"lsa_interval"`
	SpfCalculationDelay basetypes.Float64Value `tfsdk:"spf_calculation_delay"`
}

// LogicalRoutersVrfInnerOspfv3 represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfv3 struct {
	AllowRedistDefaultRoute basetypes.BoolValue   `tfsdk:"allow_redist_default_route"`
	Area                    basetypes.ListValue   `tfsdk:"area"`
	AuthProfile             basetypes.ListValue   `tfsdk:"auth_profile"`
	DisableTransitTraffic   basetypes.BoolValue   `tfsdk:"disable_transit_traffic"`
	Enable                  basetypes.BoolValue   `tfsdk:"enable"`
	ExportRules             basetypes.ListValue   `tfsdk:"export_rules"`
	GlobalBfd               basetypes.ObjectValue `tfsdk:"global_bfd"`
	GlobalIfTimer           basetypes.StringValue `tfsdk:"global_if_timer"`
	GracefulRestart         basetypes.ObjectValue `tfsdk:"graceful_restart"`
	RedistributionProfile   basetypes.StringValue `tfsdk:"redistribution_profile"`
	RejectDefaultRoute      basetypes.BoolValue   `tfsdk:"reject_default_route"`
	RouterId                basetypes.StringValue `tfsdk:"router_id"`
	SpfTimer                basetypes.StringValue `tfsdk:"spf_timer"`
	VrTimers                basetypes.ObjectValue `tfsdk:"vr_timers"`
}

// LogicalRoutersVrfInnerOspfv3AreaInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfv3AreaInner struct {
	Authentication basetypes.StringValue `tfsdk:"authentication"`
	Interface      basetypes.ListValue   `tfsdk:"interface"`
	Name           basetypes.StringValue `tfsdk:"name"`
	Range          basetypes.ListValue   `tfsdk:"range"`
	Type           basetypes.ObjectValue `tfsdk:"type"`
	VirtualLink    basetypes.ListValue   `tfsdk:"virtual_link"`
	VrRange        basetypes.ListValue   `tfsdk:"vr_range"`
}

// LogicalRoutersVrfInnerOspfv3AreaInnerInterfaceInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfv3AreaInnerInterfaceInner struct {
	Authentication basetypes.StringValue  `tfsdk:"authentication"`
	Bfd            basetypes.ObjectValue  `tfsdk:"bfd"`
	Enable         basetypes.BoolValue    `tfsdk:"enable"`
	InstanceId     basetypes.Float64Value `tfsdk:"instance_id"`
	LinkType       basetypes.ObjectValue  `tfsdk:"link_type"`
	Metric         basetypes.Float64Value `tfsdk:"metric"`
	MtuIgnore      basetypes.BoolValue    `tfsdk:"mtu_ignore"`
	Name           basetypes.StringValue  `tfsdk:"name"`
	Neighbor       basetypes.ListValue    `tfsdk:"neighbor"`
	Passive        basetypes.BoolValue    `tfsdk:"passive"`
	Priority       basetypes.Float64Value `tfsdk:"priority"`
	Timing         basetypes.StringValue  `tfsdk:"timing"`
	VrTiming       basetypes.ObjectValue  `tfsdk:"vr_timing"`
}

// LogicalRoutersVrfInnerOspfv3AreaInnerRangeInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfv3AreaInnerRangeInner struct {
	Advertise basetypes.BoolValue   `tfsdk:"advertise"`
	Name      basetypes.StringValue `tfsdk:"name"`
}

// LogicalRoutersVrfInnerOspfv3AreaInnerType represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfv3AreaInnerType struct {
	Normal basetypes.ObjectValue `tfsdk:"normal"`
	Nssa   basetypes.ObjectValue `tfsdk:"nssa"`
	Stub   basetypes.ObjectValue `tfsdk:"stub"`
}

// LogicalRoutersVrfInnerOspfv3AreaInnerTypeNssa represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfv3AreaInnerTypeNssa struct {
	Abr                         basetypes.ObjectValue `tfsdk:"abr"`
	AcceptSummary               basetypes.BoolValue   `tfsdk:"accept_summary"`
	DefaultInformationOriginate basetypes.ObjectValue `tfsdk:"default_information_originate"`
	DefaultRoute                basetypes.ObjectValue `tfsdk:"default_route"`
	NoSummary                   basetypes.BoolValue   `tfsdk:"no_summary"`
	NssaExtRange                basetypes.ListValue   `tfsdk:"nssa_ext_range"`
}

// LogicalRoutersVrfInnerOspfv3AreaInnerTypeNssaAbr represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfv3AreaInnerTypeNssaAbr struct {
	ExportList         basetypes.StringValue `tfsdk:"export_list"`
	ImportList         basetypes.StringValue `tfsdk:"import_list"`
	InboundFilterList  basetypes.StringValue `tfsdk:"inbound_filter_list"`
	NssaExtRange       basetypes.ListValue   `tfsdk:"nssa_ext_range"`
	OutboundFilterList basetypes.StringValue `tfsdk:"outbound_filter_list"`
}

// LogicalRoutersVrfInnerOspfv3AreaInnerTypeNssaAbrNssaExtRangeInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfv3AreaInnerTypeNssaAbrNssaExtRangeInner struct {
	Advertise basetypes.ObjectValue  `tfsdk:"advertise"`
	Name      basetypes.StringValue  `tfsdk:"name"`
	RouteTag  basetypes.Float64Value `tfsdk:"route_tag"`
	Suppress  basetypes.ObjectValue  `tfsdk:"suppress"`
}

// LogicalRoutersVrfInnerOspfv3AuthProfileInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfv3AuthProfileInner struct {
	Ah   basetypes.ObjectValue `tfsdk:"ah"`
	Esp  basetypes.ObjectValue `tfsdk:"esp"`
	Name basetypes.StringValue `tfsdk:"name"`
	Spi  basetypes.StringValue `tfsdk:"spi"`
}

// LogicalRoutersVrfInnerOspfv3AuthProfileInnerAh represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfv3AuthProfileInnerAh struct {
	Md5    basetypes.ObjectValue `tfsdk:"md5"`
	Sha1   basetypes.ObjectValue `tfsdk:"sha1"`
	Sha256 basetypes.ObjectValue `tfsdk:"sha256"`
	Sha384 basetypes.ObjectValue `tfsdk:"sha384"`
	Sha512 basetypes.ObjectValue `tfsdk:"sha512"`
}

// LogicalRoutersVrfInnerOspfv3AuthProfileInnerAhMd5 represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfv3AuthProfileInnerAhMd5 struct {
	Key basetypes.StringValue `tfsdk:"key"`
}

// LogicalRoutersVrfInnerOspfv3AuthProfileInnerEsp represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfv3AuthProfileInnerEsp struct {
	Authentication basetypes.ObjectValue `tfsdk:"authentication"`
	Encryption     basetypes.ObjectValue `tfsdk:"encryption"`
}

// LogicalRoutersVrfInnerOspfv3AuthProfileInnerEspAuthentication represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfv3AuthProfileInnerEspAuthentication struct {
	Md5    basetypes.ObjectValue `tfsdk:"md5"`
	None   basetypes.ObjectValue `tfsdk:"none"`
	Sha1   basetypes.ObjectValue `tfsdk:"sha1"`
	Sha256 basetypes.ObjectValue `tfsdk:"sha256"`
	Sha384 basetypes.ObjectValue `tfsdk:"sha384"`
	Sha512 basetypes.ObjectValue `tfsdk:"sha512"`
}

// LogicalRoutersVrfInnerOspfv3AuthProfileInnerEspEncryption represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerOspfv3AuthProfileInnerEspEncryption struct {
	Algorithm basetypes.StringValue `tfsdk:"algorithm"`
	Key       basetypes.StringValue `tfsdk:"key"`
}

// LogicalRoutersVrfInnerRibFilter represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerRibFilter struct {
	Ipv4 basetypes.ObjectValue `tfsdk:"ipv4"`
	Ipv6 basetypes.ObjectValue `tfsdk:"ipv6"`
}

// LogicalRoutersVrfInnerRibFilterIpv4 represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerRibFilterIpv4 struct {
	Bgp    basetypes.ObjectValue `tfsdk:"bgp"`
	Ospf   basetypes.ObjectValue `tfsdk:"ospf"`
	Rip    basetypes.ObjectValue `tfsdk:"rip"`
	Static basetypes.ObjectValue `tfsdk:"static"`
}

// LogicalRoutersVrfInnerRibFilterIpv4Bgp represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerRibFilterIpv4Bgp struct {
	RouteMap basetypes.StringValue `tfsdk:"route_map"`
}

// LogicalRoutersVrfInnerRibFilterIpv6 represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerRibFilterIpv6 struct {
	Bgp    basetypes.ObjectValue `tfsdk:"bgp"`
	Ospfv3 basetypes.ObjectValue `tfsdk:"ospfv3"`
	Static basetypes.ObjectValue `tfsdk:"static"`
}

// LogicalRoutersVrfInnerRip represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerRip struct {
	AuthProfile                  basetypes.StringValue `tfsdk:"auth_profile"`
	DefaultInformationOriginate  basetypes.BoolValue   `tfsdk:"default_information_originate"`
	Enable                       basetypes.BoolValue   `tfsdk:"enable"`
	GlobalBfd                    basetypes.ObjectValue `tfsdk:"global_bfd"`
	GlobalInboundDistributeList  basetypes.ObjectValue `tfsdk:"global_inbound_distribute_list"`
	GlobalOutboundDistributeList basetypes.ObjectValue `tfsdk:"global_outbound_distribute_list"`
	GlobalTimer                  basetypes.StringValue `tfsdk:"global_timer"`
	Interface                    basetypes.ListValue   `tfsdk:"interface"`
	RedistributionProfile        basetypes.StringValue `tfsdk:"redistribution_profile"`
}

// LogicalRoutersVrfInnerRipGlobalInboundDistributeList represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerRipGlobalInboundDistributeList struct {
	AccessList basetypes.StringValue `tfsdk:"access_list"`
}

// LogicalRoutersVrfInnerRipInterfaceInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerRipInterfaceInner struct {
	Authentication                  basetypes.StringValue `tfsdk:"authentication"`
	Bfd                             basetypes.ObjectValue `tfsdk:"bfd"`
	Enable                          basetypes.BoolValue   `tfsdk:"enable"`
	InterfaceInboundDistributeList  basetypes.ObjectValue `tfsdk:"interface_inbound_distribute_list"`
	InterfaceOutboundDistributeList basetypes.ObjectValue `tfsdk:"interface_outbound_distribute_list"`
	Mode                            basetypes.StringValue `tfsdk:"mode"`
	Name                            basetypes.StringValue `tfsdk:"name"`
	SplitHorizon                    basetypes.StringValue `tfsdk:"split_horizon"`
}

// LogicalRoutersVrfInnerRipInterfaceInnerInterfaceInboundDistributeList represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerRipInterfaceInnerInterfaceInboundDistributeList struct {
	AccessList basetypes.StringValue  `tfsdk:"access_list"`
	Metric     basetypes.Float64Value `tfsdk:"metric"`
}

// LogicalRoutersVrfInnerRoutingTable represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerRoutingTable struct {
	Ip   basetypes.ObjectValue `tfsdk:"ip"`
	Ipv6 basetypes.ObjectValue `tfsdk:"ipv6"`
}

// LogicalRoutersVrfInnerRoutingTableIp represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerRoutingTableIp struct {
	StaticRoute basetypes.ListValue `tfsdk:"static_route"`
}

// LogicalRoutersVrfInnerRoutingTableIpStaticRouteInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerRoutingTableIpStaticRouteInner struct {
	AdminDist   basetypes.Float64Value `tfsdk:"admin_dist"`
	Bfd         basetypes.ObjectValue  `tfsdk:"bfd"`
	Destination basetypes.StringValue  `tfsdk:"destination"`
	Interface   basetypes.StringValue  `tfsdk:"interface"`
	Metric      basetypes.Float64Value `tfsdk:"metric"`
	Name        basetypes.StringValue  `tfsdk:"name"`
	Nexthop     basetypes.ObjectValue  `tfsdk:"nexthop"`
	PathMonitor basetypes.ObjectValue  `tfsdk:"path_monitor"`
	RouteTable  basetypes.ObjectValue  `tfsdk:"route_table"`
}

// LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerNexthop represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerNexthop struct {
	Discard     basetypes.ObjectValue `tfsdk:"discard"`
	Fqdn        basetypes.StringValue `tfsdk:"fqdn"`
	Ipv6Address basetypes.StringValue `tfsdk:"ipv6_address"`
	NextLr      basetypes.StringValue `tfsdk:"next_lr"`
	NextVr      basetypes.StringValue `tfsdk:"next_vr"`
	Receive     basetypes.ObjectValue `tfsdk:"receive"`
	Tunnel      basetypes.StringValue `tfsdk:"tunnel"`
}

// LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerPathMonitor represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerPathMonitor struct {
	Enable              basetypes.BoolValue    `tfsdk:"enable"`
	FailureCondition    basetypes.StringValue  `tfsdk:"failure_condition"`
	HoldTime            basetypes.Float64Value `tfsdk:"hold_time"`
	MonitorDestinations basetypes.ListValue    `tfsdk:"monitor_destinations"`
}

// LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerPathMonitorMonitorDestinationsInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerPathMonitorMonitorDestinationsInner struct {
	Count           basetypes.Float64Value `tfsdk:"count"`
	Destination     basetypes.StringValue  `tfsdk:"destination"`
	DestinationFqdn basetypes.StringValue  `tfsdk:"destination_fqdn"`
	Enable          basetypes.BoolValue    `tfsdk:"enable"`
	Interval        basetypes.Float64Value `tfsdk:"interval"`
	Name            basetypes.StringValue  `tfsdk:"name"`
	Source          basetypes.StringValue  `tfsdk:"source"`
}

// LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerRouteTable represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerRouteTable struct {
	Both      basetypes.ObjectValue `tfsdk:"both"`
	Multicast basetypes.ObjectValue `tfsdk:"multicast"`
	NoInstall basetypes.ObjectValue `tfsdk:"no_install"`
	Unicast   basetypes.ObjectValue `tfsdk:"unicast"`
}

// LogicalRoutersVrfInnerRoutingTableIpv6 represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerRoutingTableIpv6 struct {
	StaticRoute basetypes.ListValue `tfsdk:"static_route"`
}

// LogicalRoutersVrfInnerRoutingTableIpv6StaticRouteInner represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerRoutingTableIpv6StaticRouteInner struct {
	AdminDist   basetypes.Float64Value `tfsdk:"admin_dist"`
	Bfd         basetypes.ObjectValue  `tfsdk:"bfd"`
	Destination basetypes.StringValue  `tfsdk:"destination"`
	Interface   basetypes.StringValue  `tfsdk:"interface"`
	Metric      basetypes.Float64Value `tfsdk:"metric"`
	Name        basetypes.StringValue  `tfsdk:"name"`
	Nexthop     basetypes.ObjectValue  `tfsdk:"nexthop"`
	Option      basetypes.ObjectValue  `tfsdk:"option"`
	PathMonitor basetypes.ObjectValue  `tfsdk:"path_monitor"`
	RouteTable  basetypes.ObjectValue  `tfsdk:"route_table"`
}

// LogicalRoutersVrfInnerRoutingTableIpv6StaticRouteInnerOption represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerRoutingTableIpv6StaticRouteInnerOption struct {
	Passive basetypes.ObjectValue `tfsdk:"passive"`
}

// LogicalRoutersVrfInnerVrAdminDists represents a nested structure within the LogicalRouters model
type LogicalRoutersVrfInnerVrAdminDists struct {
	Ebgp       basetypes.Float64Value `tfsdk:"ebgp"`
	Ibgp       basetypes.Float64Value `tfsdk:"ibgp"`
	OspfExt    basetypes.Float64Value `tfsdk:"ospf_ext"`
	OspfInt    basetypes.Float64Value `tfsdk:"ospf_int"`
	Ospfv3Ext  basetypes.Float64Value `tfsdk:"ospfv3_ext"`
	Ospfv3Int  basetypes.Float64Value `tfsdk:"ospfv3_int"`
	Rip        basetypes.Float64Value `tfsdk:"rip"`
	Static     basetypes.Float64Value `tfsdk:"static"`
	StaticIpv6 basetypes.Float64Value `tfsdk:"static_ipv6"`
}

// AttrTypes defines the attribute types for the LogicalRouters model.
func (o LogicalRouters) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":          basetypes.StringType{},
		"device":        basetypes.StringType{},
		"folder":        basetypes.StringType{},
		"id":            basetypes.StringType{},
		"name":          basetypes.StringType{},
		"routing_stack": basetypes.StringType{},
		"snippet":       basetypes.StringType{},
		"vrf": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"admin_dists": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"bgp_external": basetypes.NumberType{},
						"bgp_internal": basetypes.NumberType{},
						"bgp_local":    basetypes.NumberType{},
						"ospf_ext":     basetypes.NumberType{},
						"ospf_inter":   basetypes.NumberType{},
						"ospf_intra":   basetypes.NumberType{},
						"ospfv3_ext":   basetypes.NumberType{},
						"ospfv3_inter": basetypes.NumberType{},
						"ospfv3_intra": basetypes.NumberType{},
						"rip":          basetypes.NumberType{},
						"static":       basetypes.NumberType{},
						"static_ipv6":  basetypes.NumberType{},
					},
				},
				"bgp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"advertise_network": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"ipv4": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"network": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"backdoor":  basetypes.BoolType{},
												"multicast": basetypes.BoolType{},
												"name":      basetypes.StringType{},
												"unicast":   basetypes.BoolType{},
											},
										}},
									},
								},
								"ipv6": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"network": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"name":    basetypes.StringType{},
												"unicast": basetypes.BoolType{},
											},
										}},
									},
								},
							},
						},
						"aggregate": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"aggregate_med": basetypes.BoolType{},
							},
						},
						"aggregate_routes": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"as_set":       basetypes.BoolType{},
								"description":  basetypes.StringType{},
								"enable":       basetypes.BoolType{},
								"name":         basetypes.StringType{},
								"same_med":     basetypes.BoolType{},
								"summary_only": basetypes.BoolType{},
								"type": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"ipv4": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"attribute_map":  basetypes.StringType{},
												"summary_prefix": basetypes.StringType{},
												"suppress_map":   basetypes.StringType{},
											},
										},
										"ipv6": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"attribute_map":  basetypes.StringType{},
												"summary_prefix": basetypes.StringType{},
												"suppress_map":   basetypes.StringType{},
											},
										},
									},
								},
							},
						}},
						"allow_redist_default_route":     basetypes.BoolType{},
						"always_advertise_network_route": basetypes.BoolType{},
						"as_format":                      basetypes.StringType{},
						"confederation_member_as":        basetypes.StringType{},
						"default_local_preference":       basetypes.NumberType{},
						"ecmp_multi_as":                  basetypes.BoolType{},
						"enable":                         basetypes.BoolType{},
						"enforce_first_as":               basetypes.BoolType{},
						"fast_external_failover":         basetypes.BoolType{},
						"global_bfd": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"profile": basetypes.StringType{},
							},
						},
						"graceful_restart": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable":                basetypes.BoolType{},
								"local_restart_time":    basetypes.NumberType{},
								"max_peer_restart_time": basetypes.NumberType{},
								"stale_route_time":      basetypes.NumberType{},
							},
						},
						"graceful_shutdown": basetypes.BoolType{},
						"install_route":     basetypes.BoolType{},
						"local_as":          basetypes.StringType{},
						"med": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"always_compare_med":           basetypes.BoolType{},
								"deterministic_med_comparison": basetypes.BoolType{},
							},
						},
						"peer_group": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address_family": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"ipv4": basetypes.StringType{},
										"ipv6": basetypes.StringType{},
									},
								},
								"aggregated_confed_as_path": basetypes.BoolType{},
								"connection_options": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"authentication": basetypes.StringType{},
										"dampening":      basetypes.StringType{},
										"multihop":       basetypes.NumberType{},
										"timers":         basetypes.StringType{},
									},
								},
								"enable": basetypes.BoolType{},
								"filtering_profile": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"ipv4": basetypes.StringType{},
										"ipv6": basetypes.StringType{},
									},
								},
								"name": basetypes.StringType{},
								"peer": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"bfd": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"multihop": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"min_received_ttl": basetypes.NumberType{},
													},
												},
												"profile": basetypes.StringType{},
											},
										},
										"connection_options": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"authentication": basetypes.StringType{},
												"dampening":      basetypes.StringType{},
												"hold_time":      basetypes.StringType{},
												"idle_hold_time": basetypes.NumberType{},
												"incoming_bgp_connection": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"allow":       basetypes.BoolType{},
														"remote_port": basetypes.NumberType{},
													},
												},
												"keep_alive_interval":    basetypes.StringType{},
												"max_prefixes":           basetypes.StringType{},
												"min_route_adv_interval": basetypes.NumberType{},
												"multihop":               basetypes.StringType{},
												"open_delay_time":        basetypes.NumberType{},
												"outgoing_bgp_connection": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"allow":      basetypes.BoolType{},
														"local_port": basetypes.NumberType{},
													},
												},
												"timers": basetypes.StringType{},
											},
										},
										"enable":                            basetypes.BoolType{},
										"enable_mp_bgp":                     basetypes.BoolType{},
										"enable_sender_side_loop_detection": basetypes.BoolType{},
										"inherit": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"no": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"address_family": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"ipv4": basetypes.StringType{},
																"ipv6": basetypes.StringType{},
															},
														},
														"filtering_profile": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"ipv4": basetypes.StringType{},
																"ipv6": basetypes.StringType{},
															},
														},
													},
												},
												"yes": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
											},
										},
										"local_address": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"interface": basetypes.StringType{},
												"ip":        basetypes.StringType{},
											},
										},
										"name":    basetypes.StringType{},
										"passive": basetypes.BoolType{},
										"peer_address": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"fqdn": basetypes.StringType{},
												"ip":   basetypes.StringType{},
											},
										},
										"peer_as":          basetypes.StringType{},
										"peering_type":     basetypes.StringType{},
										"reflector_client": basetypes.StringType{},
										"subsequent_address_family_identifier": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"multicast": basetypes.BoolType{},
												"unicast":   basetypes.BoolType{},
											},
										},
									},
								}},
								"soft_reset_with_stored_info": basetypes.BoolType{},
								"type": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"ebgp": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"export_nexthop":    basetypes.StringType{},
												"import_nexthop":    basetypes.StringType{},
												"remove_private_as": basetypes.BoolType{},
											},
										},
										"ebgp_confed": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"export_nexthop": basetypes.StringType{},
											},
										},
										"ibgp": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"export_nexthop": basetypes.StringType{},
											},
										},
										"ibgp_confed": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"export_nexthop": basetypes.StringType{},
											},
										},
									},
								},
							},
						}},
						"policy": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"aggregation": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"address": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"advertise_filters": basetypes.ListType{ElemType: basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"enable": basetypes.BoolType{},
														"match": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{
																		"exact": basetypes.BoolType{},
																		"name":  basetypes.StringType{},
																	},
																}},
																"afi": basetypes.StringType{},
																"as_path": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{
																		"regex": basetypes.StringType{},
																	},
																},
																"community": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{
																		"regex": basetypes.StringType{},
																	},
																},
																"extended_community": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{
																		"regex": basetypes.StringType{},
																	},
																},
																"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
																"med":         basetypes.NumberType{},
																"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
																"route_table": basetypes.StringType{},
																"safi":        basetypes.StringType{},
															},
														},
														"name": basetypes.StringType{},
													},
												}},
												"aggregate_route_attributes": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"as_path": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"none": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{},
																},
																"prepend": basetypes.NumberType{},
																"remove": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{},
																},
																"remove_and_prepend": basetypes.NumberType{},
															},
														},
														"as_path_limit": basetypes.NumberType{},
														"community": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"append": basetypes.ListType{ElemType: basetypes.StringType{}},
																"none": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{},
																},
																"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
																"remove_all": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{},
																},
																"remove_regex": basetypes.StringType{},
															},
														},
														"extended_community": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"append": basetypes.ListType{ElemType: basetypes.StringType{}},
																"none": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{},
																},
																"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
																"remove_all": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{},
																},
																"remove_regex": basetypes.StringType{},
															},
														},
														"local_preference": basetypes.NumberType{},
														"med":              basetypes.NumberType{},
														"nexthop":          basetypes.StringType{},
														"origin":           basetypes.StringType{},
														"weight":           basetypes.NumberType{},
													},
												},
												"as_set":  basetypes.BoolType{},
												"enable":  basetypes.BoolType{},
												"name":    basetypes.StringType{},
												"prefix":  basetypes.StringType{},
												"summary": basetypes.BoolType{},
												"suppress_filters": basetypes.ListType{ElemType: basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"enable": basetypes.BoolType{},
														"match": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{
																		"exact": basetypes.BoolType{},
																		"name":  basetypes.StringType{},
																	},
																}},
																"afi": basetypes.StringType{},
																"as_path": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{
																		"regex": basetypes.StringType{},
																	},
																},
																"community": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{
																		"regex": basetypes.StringType{},
																	},
																},
																"extended_community": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{
																		"regex": basetypes.StringType{},
																	},
																},
																"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
																"med":         basetypes.NumberType{},
																"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
																"route_table": basetypes.StringType{},
																"safi":        basetypes.StringType{},
															},
														},
														"name": basetypes.StringType{},
													},
												}},
											},
										}},
									},
								},
								"conditional_advertisement": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"policy": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"advertise_filters": basetypes.ListType{ElemType: basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"enable": basetypes.BoolType{},
														"match": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{
																		"exact": basetypes.BoolType{},
																		"name":  basetypes.StringType{},
																	},
																}},
																"afi": basetypes.StringType{},
																"as_path": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{
																		"regex": basetypes.StringType{},
																	},
																},
																"community": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{
																		"regex": basetypes.StringType{},
																	},
																},
																"extended_community": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{
																		"regex": basetypes.StringType{},
																	},
																},
																"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
																"med":         basetypes.NumberType{},
																"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
																"route_table": basetypes.StringType{},
																"safi":        basetypes.StringType{},
															},
														},
														"name": basetypes.StringType{},
													},
												}},
												"enable": basetypes.BoolType{},
												"name":   basetypes.StringType{},
												"non_exist_filters": basetypes.ListType{ElemType: basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"enable": basetypes.BoolType{},
														"match": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{
																		"exact": basetypes.BoolType{},
																		"name":  basetypes.StringType{},
																	},
																}},
																"afi": basetypes.StringType{},
																"as_path": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{
																		"regex": basetypes.StringType{},
																	},
																},
																"community": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{
																		"regex": basetypes.StringType{},
																	},
																},
																"extended_community": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{
																		"regex": basetypes.StringType{},
																	},
																},
																"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
																"med":         basetypes.NumberType{},
																"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
																"route_table": basetypes.StringType{},
																"safi":        basetypes.StringType{},
															},
														},
														"name": basetypes.StringType{},
													},
												}},
												"used_by": basetypes.ListType{ElemType: basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												}},
											},
										}},
									},
								},
								"export": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"rules": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"action": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"allow": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"update": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{
																		"as_path": basetypes.ObjectType{
																			AttrTypes: map[string]attr.Type{
																				"none": basetypes.ObjectType{
																					AttrTypes: map[string]attr.Type{},
																				},
																				"prepend": basetypes.NumberType{},
																				"remove": basetypes.ObjectType{
																					AttrTypes: map[string]attr.Type{},
																				},
																				"remove_and_prepend": basetypes.NumberType{},
																			},
																		},
																		"as_path_limit": basetypes.NumberType{},
																		"community": basetypes.ObjectType{
																			AttrTypes: map[string]attr.Type{
																				"append": basetypes.ListType{ElemType: basetypes.StringType{}},
																				"none": basetypes.ObjectType{
																					AttrTypes: map[string]attr.Type{},
																				},
																				"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
																				"remove_all": basetypes.ObjectType{
																					AttrTypes: map[string]attr.Type{},
																				},
																				"remove_regex": basetypes.StringType{},
																			},
																		},
																		"extended_community": basetypes.ObjectType{
																			AttrTypes: map[string]attr.Type{
																				"append": basetypes.ListType{ElemType: basetypes.StringType{}},
																				"none": basetypes.ObjectType{
																					AttrTypes: map[string]attr.Type{},
																				},
																				"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
																				"remove_all": basetypes.ObjectType{
																					AttrTypes: map[string]attr.Type{},
																				},
																				"remove_regex": basetypes.StringType{},
																			},
																		},
																		"local_preference": basetypes.NumberType{},
																		"med":              basetypes.NumberType{},
																		"nexthop":          basetypes.StringType{},
																		"origin":           basetypes.StringType{},
																	},
																},
															},
														},
														"deny": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
													},
												},
												"enable": basetypes.BoolType{},
												"match": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"exact": basetypes.BoolType{},
																"name":  basetypes.StringType{},
															},
														}},
														"afi": basetypes.StringType{},
														"as_path": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"regex": basetypes.StringType{},
															},
														},
														"community": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"regex": basetypes.StringType{},
															},
														},
														"extended_community": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"regex": basetypes.StringType{},
															},
														},
														"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
														"med":         basetypes.NumberType{},
														"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
														"route_table": basetypes.StringType{},
														"safi":        basetypes.StringType{},
													},
												},
												"name":    basetypes.StringType{},
												"used_by": basetypes.ListType{ElemType: basetypes.StringType{}},
											},
										}},
									},
								},
								"import": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"rules": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"action": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"allow": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"dampening": basetypes.StringType{},
																"update": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{
																		"as_path": basetypes.ObjectType{
																			AttrTypes: map[string]attr.Type{
																				"none": basetypes.ObjectType{
																					AttrTypes: map[string]attr.Type{},
																				},
																				"prepend": basetypes.NumberType{},
																				"remove": basetypes.ObjectType{
																					AttrTypes: map[string]attr.Type{},
																				},
																				"remove_and_prepend": basetypes.NumberType{},
																			},
																		},
																		"as_path_limit": basetypes.NumberType{},
																		"community": basetypes.ObjectType{
																			AttrTypes: map[string]attr.Type{
																				"append": basetypes.ListType{ElemType: basetypes.StringType{}},
																				"none": basetypes.ObjectType{
																					AttrTypes: map[string]attr.Type{},
																				},
																				"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
																				"remove_all": basetypes.ObjectType{
																					AttrTypes: map[string]attr.Type{},
																				},
																				"remove_regex": basetypes.StringType{},
																			},
																		},
																		"extended_community": basetypes.ObjectType{
																			AttrTypes: map[string]attr.Type{
																				"append": basetypes.ListType{ElemType: basetypes.StringType{}},
																				"none": basetypes.ObjectType{
																					AttrTypes: map[string]attr.Type{},
																				},
																				"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
																				"remove_all": basetypes.ObjectType{
																					AttrTypes: map[string]attr.Type{},
																				},
																				"remove_regex": basetypes.StringType{},
																			},
																		},
																		"local_preference": basetypes.NumberType{},
																		"med":              basetypes.NumberType{},
																		"nexthop":          basetypes.StringType{},
																		"origin":           basetypes.StringType{},
																		"weight":           basetypes.NumberType{},
																	},
																},
															},
														},
														"deny": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
													},
												},
												"enable": basetypes.BoolType{},
												"match": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"exact": basetypes.BoolType{},
																"name":  basetypes.StringType{},
															},
														}},
														"afi": basetypes.StringType{},
														"as_path": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"regex": basetypes.StringType{},
															},
														},
														"community": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"regex": basetypes.StringType{},
															},
														},
														"extended_community": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"regex": basetypes.StringType{},
															},
														},
														"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
														"med":         basetypes.NumberType{},
														"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
														"route_table": basetypes.StringType{},
														"safi":        basetypes.StringType{},
													},
												},
												"name":    basetypes.StringType{},
												"used_by": basetypes.ListType{ElemType: basetypes.StringType{}},
											},
										}},
									},
								},
							},
						},
						"redist_rules": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address_family_identifier": basetypes.StringType{},
								"enable":                    basetypes.BoolType{},
								"metric":                    basetypes.NumberType{},
								"name":                      basetypes.StringType{},
								"route_table":               basetypes.StringType{},
								"set_as_path_limit":         basetypes.NumberType{},
								"set_community":             basetypes.ListType{ElemType: basetypes.StringType{}},
								"set_extended_community":    basetypes.ListType{ElemType: basetypes.StringType{}},
								"set_local_preference":      basetypes.NumberType{},
								"set_med":                   basetypes.NumberType{},
								"set_origin":                basetypes.StringType{},
							},
						}},
						"redistribution_profile": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"ipv4": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"unicast": basetypes.StringType{},
									},
								},
								"ipv6": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"unicast": basetypes.StringType{},
									},
								},
							},
						},
						"reject_default_route": basetypes.BoolType{},
						"router_id":            basetypes.StringType{},
					},
				},
				"ecmp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"algorithm": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"balanced_round_robin": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"ip_hash": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"hash_seed": basetypes.NumberType{},
										"src_only":  basetypes.BoolType{},
										"use_port":  basetypes.BoolType{},
									},
								},
								"ip_modulo": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"weighted_round_robin": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"name":   basetypes.StringType{},
												"weight": basetypes.NumberType{},
											},
										}},
									},
								},
							},
						},
						"enable":             basetypes.BoolType{},
						"max_path":           basetypes.NumberType{},
						"strict_source_path": basetypes.BoolType{},
						"symmetric_return":   basetypes.BoolType{},
					},
				},
				"global_vrid": basetypes.NumberType{},
				"interface":   basetypes.ListType{ElemType: basetypes.StringType{}},
				"multicast": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable":    basetypes.BoolType{},
						"enable_v6": basetypes.BoolType{},
						"igmp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"dynamic": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"group_filter":          basetypes.StringType{},
												"max_groups":            basetypes.StringType{},
												"max_sources":           basetypes.StringType{},
												"name":                  basetypes.StringType{},
												"query_profile":         basetypes.StringType{},
												"robustness":            basetypes.StringType{},
												"router_alert_policing": basetypes.BoolType{},
												"version":               basetypes.StringType{},
											},
										}},
									},
								},
								"enable": basetypes.BoolType{},
								"static": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"group_address":  basetypes.StringType{},
										"interface":      basetypes.StringType{},
										"name":           basetypes.StringType{},
										"source_address": basetypes.StringType{},
									},
								}},
							},
						},
						"interface_group": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"description": basetypes.StringType{},
								"group_permission": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"any_source_multicast": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"group_address": basetypes.StringType{},
												"included":      basetypes.BoolType{},
												"name":          basetypes.StringType{},
											},
										}},
										"source_specific_multicast": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"group_address":  basetypes.StringType{},
												"included":       basetypes.BoolType{},
												"name":           basetypes.StringType{},
												"source_address": basetypes.StringType{},
											},
										}},
									},
								},
								"igmp": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"enable":                     basetypes.BoolType{},
										"immediate_leave":            basetypes.BoolType{},
										"last_member_query_interval": basetypes.NumberType{},
										"max_groups":                 basetypes.StringType{},
										"max_query_response_time":    basetypes.NumberType{},
										"max_sources":                basetypes.StringType{},
										"mode":                       basetypes.StringType{},
										"query_interval":             basetypes.NumberType{},
										"robustness":                 basetypes.StringType{},
										"router_alert_policing":      basetypes.BoolType{},
										"version":                    basetypes.StringType{},
									},
								},
								"interface": basetypes.ListType{ElemType: basetypes.StringType{}},
								"name":      basetypes.StringType{},
								"pim": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"allowed_neighbors": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"name": basetypes.StringType{},
											},
										}},
										"assert_interval":     basetypes.NumberType{},
										"bsr_border":          basetypes.BoolType{},
										"dr_priority":         basetypes.NumberType{},
										"enable":              basetypes.BoolType{},
										"hello_interval":      basetypes.NumberType{},
										"join_prune_interval": basetypes.NumberType{},
									},
								},
							},
						}},
						"mode": basetypes.StringType{},
						"msdp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable":                basetypes.BoolType{},
								"global_authentication": basetypes.StringType{},
								"global_timer":          basetypes.StringType{},
								"originator_id": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"interface": basetypes.StringType{},
										"ip":        basetypes.StringType{},
									},
								},
								"peer": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"authentication":    basetypes.StringType{},
										"enable":            basetypes.BoolType{},
										"inbound_sa_filter": basetypes.StringType{},
										"local_address": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"interface": basetypes.StringType{},
												"ip":        basetypes.StringType{},
											},
										},
										"max_sa":             basetypes.NumberType{},
										"name":               basetypes.StringType{},
										"outbound_sa_filter": basetypes.StringType{},
										"peer_address": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"fqdn": basetypes.StringType{},
												"ip":   basetypes.StringType{},
											},
										},
										"peer_as": basetypes.StringType{},
									},
								}},
							},
						},
						"pim": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable":           basetypes.BoolType{},
								"group_permission": basetypes.StringType{},
								"if_timer_global":  basetypes.StringType{},
								"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"description":     basetypes.StringType{},
										"dr_priority":     basetypes.NumberType{},
										"if_timer":        basetypes.StringType{},
										"name":            basetypes.StringType{},
										"neighbor_filter": basetypes.StringType{},
										"send_bsm":        basetypes.BoolType{},
									},
								}},
								"route_ageout_time": basetypes.NumberType{},
								"rp": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"external_rp": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"group_list": basetypes.StringType{},
												"name":       basetypes.StringType{},
												"override":   basetypes.BoolType{},
											},
										}},
										"local_rp": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"candidate_rp": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"address":                basetypes.StringType{},
														"advertisement_interval": basetypes.NumberType{},
														"group_list":             basetypes.StringType{},
														"interface":              basetypes.StringType{},
														"priority":               basetypes.NumberType{},
													},
												},
												"static_rp": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"address":    basetypes.StringType{},
														"group_list": basetypes.StringType{},
														"interface":  basetypes.StringType{},
														"override":   basetypes.BoolType{},
													},
												},
											},
										},
									},
								},
								"rpf_lookup_mode": basetypes.StringType{},
								"spt_threshold": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"name":      basetypes.StringType{},
										"threshold": basetypes.StringType{},
									},
								}},
								"ssm_address_space": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"group_list": basetypes.StringType{},
									},
								},
							},
						},
						"route_ageout_time": basetypes.NumberType{},
						"rp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"external_rp": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"group_addresses": basetypes.ListType{ElemType: basetypes.StringType{}},
										"name":            basetypes.StringType{},
										"override":        basetypes.BoolType{},
									},
								}},
								"local_rp": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"candidate_rp": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"address":                basetypes.StringType{},
												"advertisement_interval": basetypes.NumberType{},
												"group_addresses":        basetypes.ListType{ElemType: basetypes.StringType{}},
												"interface":              basetypes.StringType{},
												"priority":               basetypes.NumberType{},
											},
										},
										"static_rp": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"address":         basetypes.StringType{},
												"group_addresses": basetypes.ListType{ElemType: basetypes.StringType{}},
												"interface":       basetypes.StringType{},
												"override":        basetypes.BoolType{},
											},
										},
									},
								},
							},
						},
						"spt_threshold": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":      basetypes.StringType{},
								"threshold": basetypes.StringType{},
							},
						}},
						"ssm_address_space": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"group_address": basetypes.StringType{},
								"included":      basetypes.BoolType{},
								"name":          basetypes.StringType{},
							},
						}},
						"static_route": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"destination": basetypes.StringType{},
								"interface":   basetypes.StringType{},
								"name":        basetypes.StringType{},
								"nexthop": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"ip_address": basetypes.StringType{},
									},
								},
								"preference": basetypes.NumberType{},
							},
						}},
					},
				},
				"name": basetypes.StringType{},
				"ospf": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"allow_redist_default_route": basetypes.BoolType{},
						"area": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"authentication": basetypes.StringType{},
								"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"authentication": basetypes.StringType{},
										"bfd": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"profile": basetypes.StringType{},
											},
										},
										"enable": basetypes.BoolType{},
										"link_type": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"broadcast": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"p2mp": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"neighbor": basetypes.ListType{ElemType: basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"name":     basetypes.StringType{},
																"priority": basetypes.NumberType{},
															},
														}},
													},
												},
												"p2p": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
											},
										},
										"metric":     basetypes.NumberType{},
										"mtu_ignore": basetypes.BoolType{},
										"name":       basetypes.StringType{},
										"passive":    basetypes.BoolType{},
										"priority":   basetypes.NumberType{},
										"timing":     basetypes.StringType{},
										"vr_timing": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"dead_counts":         basetypes.NumberType{},
												"gr_delay":            basetypes.NumberType{},
												"hello_interval":      basetypes.NumberType{},
												"retransmit_interval": basetypes.NumberType{},
												"transit_delay":       basetypes.NumberType{},
											},
										},
									},
								}},
								"name": basetypes.StringType{},
								"range": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"advertise":  basetypes.BoolType{},
										"name":       basetypes.StringType{},
										"substitute": basetypes.StringType{},
									},
								}},
								"type": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"normal": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"abr": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"export_list":          basetypes.StringType{},
														"import_list":          basetypes.StringType{},
														"inbound_filter_list":  basetypes.StringType{},
														"outbound_filter_list": basetypes.StringType{},
													},
												},
											},
										},
										"nssa": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"abr": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"export_list":         basetypes.StringType{},
														"import_list":         basetypes.StringType{},
														"inbound_filter_list": basetypes.StringType{},
														"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"advertise": basetypes.BoolType{},
																"name":      basetypes.StringType{},
																"route_tag": basetypes.NumberType{},
															},
														}},
														"outbound_filter_list": basetypes.StringType{},
													},
												},
												"accept_summary": basetypes.BoolType{},
												"default_information_originate": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"metric":      basetypes.NumberType{},
														"metric_type": basetypes.StringType{},
													},
												},
												"default_route": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"advertise": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"metric": basetypes.NumberType{},
																"type":   basetypes.StringType{},
															},
														},
														"disable": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
													},
												},
												"no_summary": basetypes.BoolType{},
												"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"advertise": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
														"name": basetypes.StringType{},
														"suppress": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
													},
												}},
											},
										},
										"stub": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"abr": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"export_list":          basetypes.StringType{},
														"import_list":          basetypes.StringType{},
														"inbound_filter_list":  basetypes.StringType{},
														"outbound_filter_list": basetypes.StringType{},
													},
												},
												"accept_summary": basetypes.BoolType{},
												"default_route": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"advertise": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"metric": basetypes.NumberType{},
															},
														},
														"disable": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
													},
												},
												"default_route_metric": basetypes.NumberType{},
												"no_summary":           basetypes.BoolType{},
											},
										},
									},
								},
								"virtual_link": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"authentication": basetypes.StringType{},
										"bfd": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"profile": basetypes.StringType{},
											},
										},
										"enable":          basetypes.BoolType{},
										"instance_id":     basetypes.NumberType{},
										"interface_id":    basetypes.NumberType{},
										"name":            basetypes.StringType{},
										"neighbor_id":     basetypes.StringType{},
										"passive":         basetypes.BoolType{},
										"timing":          basetypes.StringType{},
										"transit_area_id": basetypes.StringType{},
										"vr_timing": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"dead_counts":         basetypes.NumberType{},
												"hello_interval":      basetypes.NumberType{},
												"retransmit_interval": basetypes.NumberType{},
												"transit_delay":       basetypes.NumberType{},
											},
										},
									},
								}},
								"vr_range": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"advertise": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"name": basetypes.StringType{},
										"suppress": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								}},
							},
						}},
						"auth_profile": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"md5": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"key":       basetypes.StringType{},
										"name":      basetypes.NumberType{},
										"preferred": basetypes.BoolType{},
									},
								}},
								"name":     basetypes.StringType{},
								"password": basetypes.StringType{},
							},
						}},
						"enable": basetypes.BoolType{},
						"export_rules": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"metric":        basetypes.NumberType{},
								"name":          basetypes.StringType{},
								"new_path_type": basetypes.StringType{},
								"new_tag":       basetypes.StringType{},
							},
						}},
						"flood_prevention": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"hello": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"enable":     basetypes.BoolType{},
										"max_packet": basetypes.NumberType{},
									},
								},
								"lsa": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"enable":     basetypes.BoolType{},
										"max_packet": basetypes.NumberType{},
									},
								},
							},
						},
						"global_bfd": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"profile": basetypes.StringType{},
							},
						},
						"global_if_timer": basetypes.StringType{},
						"graceful_restart": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable":                    basetypes.BoolType{},
								"grace_period":              basetypes.NumberType{},
								"helper_enable":             basetypes.BoolType{},
								"max_neighbor_restart_time": basetypes.NumberType{},
								"strict__l_s_a_checking":    basetypes.BoolType{},
							},
						},
						"redistribution_profile": basetypes.StringType{},
						"reject_default_route":   basetypes.BoolType{},
						"rfc1583":                basetypes.BoolType{},
						"router_id":              basetypes.StringType{},
						"spf_timer":              basetypes.StringType{},
						"vr_timers": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"lsa_interval":          basetypes.NumberType{},
								"spf_calculation_delay": basetypes.NumberType{},
							},
						},
					},
				},
				"ospfv3": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"allow_redist_default_route": basetypes.BoolType{},
						"area": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"authentication": basetypes.StringType{},
								"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"authentication": basetypes.StringType{},
										"bfd": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"profile": basetypes.StringType{},
											},
										},
										"enable":      basetypes.BoolType{},
										"instance_id": basetypes.NumberType{},
										"link_type": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"broadcast": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"p2mp": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"neighbor": basetypes.ListType{ElemType: basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"name":     basetypes.StringType{},
																"priority": basetypes.NumberType{},
															},
														}},
													},
												},
												"p2p": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
											},
										},
										"metric":     basetypes.NumberType{},
										"mtu_ignore": basetypes.BoolType{},
										"name":       basetypes.StringType{},
										"neighbor": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"name": basetypes.StringType{},
											},
										}},
										"passive":  basetypes.BoolType{},
										"priority": basetypes.NumberType{},
										"timing":   basetypes.StringType{},
										"vr_timing": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"dead_counts":         basetypes.NumberType{},
												"gr_delay":            basetypes.NumberType{},
												"hello_interval":      basetypes.NumberType{},
												"retransmit_interval": basetypes.NumberType{},
												"transit_delay":       basetypes.NumberType{},
											},
										},
									},
								}},
								"name": basetypes.StringType{},
								"range": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"advertise": basetypes.BoolType{},
										"name":      basetypes.StringType{},
									},
								}},
								"type": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"normal": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"abr": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"export_list":          basetypes.StringType{},
														"import_list":          basetypes.StringType{},
														"inbound_filter_list":  basetypes.StringType{},
														"outbound_filter_list": basetypes.StringType{},
													},
												},
											},
										},
										"nssa": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"abr": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"export_list":         basetypes.StringType{},
														"import_list":         basetypes.StringType{},
														"inbound_filter_list": basetypes.StringType{},
														"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"advertise": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{},
																},
																"name":      basetypes.StringType{},
																"route_tag": basetypes.NumberType{},
																"suppress": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{},
																},
															},
														}},
														"outbound_filter_list": basetypes.StringType{},
													},
												},
												"accept_summary": basetypes.BoolType{},
												"default_information_originate": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"metric":      basetypes.NumberType{},
														"metric_type": basetypes.StringType{},
													},
												},
												"default_route": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"advertise": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"metric": basetypes.NumberType{},
																"type":   basetypes.StringType{},
															},
														},
														"disable": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
													},
												},
												"no_summary": basetypes.BoolType{},
												"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"advertise": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
														"name":      basetypes.StringType{},
														"route_tag": basetypes.NumberType{},
														"suppress": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
													},
												}},
											},
										},
										"stub": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"abr": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"export_list":          basetypes.StringType{},
														"import_list":          basetypes.StringType{},
														"inbound_filter_list":  basetypes.StringType{},
														"outbound_filter_list": basetypes.StringType{},
													},
												},
												"accept_summary": basetypes.BoolType{},
												"default_route": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"advertise": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"metric": basetypes.NumberType{},
															},
														},
														"disable": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
													},
												},
												"default_route_metric": basetypes.NumberType{},
												"no_summary":           basetypes.BoolType{},
											},
										},
									},
								},
								"virtual_link": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"authentication": basetypes.StringType{},
										"bfd": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"profile": basetypes.StringType{},
											},
										},
										"enable":          basetypes.BoolType{},
										"instance_id":     basetypes.NumberType{},
										"interface_id":    basetypes.NumberType{},
										"name":            basetypes.StringType{},
										"neighbor_id":     basetypes.StringType{},
										"passive":         basetypes.BoolType{},
										"timing":          basetypes.StringType{},
										"transit_area_id": basetypes.StringType{},
										"vr_timing": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"dead_counts":         basetypes.NumberType{},
												"hello_interval":      basetypes.NumberType{},
												"retransmit_interval": basetypes.NumberType{},
												"transit_delay":       basetypes.NumberType{},
											},
										},
									},
								}},
								"vr_range": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"advertise": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"name": basetypes.StringType{},
										"suppress": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								}},
							},
						}},
						"auth_profile": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"ah": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"md5": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"key": basetypes.StringType{},
											},
										},
										"sha1": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"key": basetypes.StringType{},
											},
										},
										"sha256": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"key": basetypes.StringType{},
											},
										},
										"sha384": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"key": basetypes.StringType{},
											},
										},
										"sha512": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"key": basetypes.StringType{},
											},
										},
									},
								},
								"esp": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"authentication": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"md5": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"key": basetypes.StringType{},
													},
												},
												"none": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"sha1": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"key": basetypes.StringType{},
													},
												},
												"sha256": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"key": basetypes.StringType{},
													},
												},
												"sha384": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"key": basetypes.StringType{},
													},
												},
												"sha512": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"key": basetypes.StringType{},
													},
												},
											},
										},
										"encryption": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"algorithm": basetypes.StringType{},
												"key":       basetypes.StringType{},
											},
										},
									},
								},
								"name": basetypes.StringType{},
								"spi":  basetypes.StringType{},
							},
						}},
						"disable_transit_traffic": basetypes.BoolType{},
						"enable":                  basetypes.BoolType{},
						"export_rules": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"metric":        basetypes.NumberType{},
								"name":          basetypes.StringType{},
								"new_path_type": basetypes.StringType{},
								"new_tag":       basetypes.StringType{},
							},
						}},
						"global_bfd": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"profile": basetypes.StringType{},
							},
						},
						"global_if_timer": basetypes.StringType{},
						"graceful_restart": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable":                    basetypes.BoolType{},
								"grace_period":              basetypes.NumberType{},
								"helper_enable":             basetypes.BoolType{},
								"max_neighbor_restart_time": basetypes.NumberType{},
								"strict__l_s_a_checking":    basetypes.BoolType{},
							},
						},
						"redistribution_profile": basetypes.StringType{},
						"reject_default_route":   basetypes.BoolType{},
						"router_id":              basetypes.StringType{},
						"spf_timer":              basetypes.StringType{},
						"vr_timers": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"lsa_interval":          basetypes.NumberType{},
								"spf_calculation_delay": basetypes.NumberType{},
							},
						},
					},
				},
				"rib_filter": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ipv4": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"bgp": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"route_map": basetypes.StringType{},
									},
								},
								"ospf": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"route_map": basetypes.StringType{},
									},
								},
								"rip": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"route_map": basetypes.StringType{},
									},
								},
								"static": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"route_map": basetypes.StringType{},
									},
								},
							},
						},
						"ipv6": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"bgp": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"route_map": basetypes.StringType{},
									},
								},
								"ospfv3": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"route_map": basetypes.StringType{},
									},
								},
								"static": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"route_map": basetypes.StringType{},
									},
								},
							},
						},
					},
				},
				"rip": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"auth_profile":                  basetypes.StringType{},
						"default_information_originate": basetypes.BoolType{},
						"enable":                        basetypes.BoolType{},
						"global_bfd": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"profile": basetypes.StringType{},
							},
						},
						"global_inbound_distribute_list": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"access_list": basetypes.StringType{},
							},
						},
						"global_outbound_distribute_list": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"access_list": basetypes.StringType{},
							},
						},
						"global_timer": basetypes.StringType{},
						"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"authentication": basetypes.StringType{},
								"bfd": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"profile": basetypes.StringType{},
									},
								},
								"enable": basetypes.BoolType{},
								"interface_inbound_distribute_list": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"access_list": basetypes.StringType{},
										"metric":      basetypes.NumberType{},
									},
								},
								"interface_outbound_distribute_list": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"access_list": basetypes.StringType{},
										"metric":      basetypes.NumberType{},
									},
								},
								"mode":          basetypes.StringType{},
								"name":          basetypes.StringType{},
								"split_horizon": basetypes.StringType{},
							},
						}},
						"redistribution_profile": basetypes.StringType{},
					},
				},
				"routing_table": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ip": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"static_route": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"admin_dist": basetypes.NumberType{},
										"bfd": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"profile": basetypes.StringType{},
											},
										},
										"destination": basetypes.StringType{},
										"interface":   basetypes.StringType{},
										"metric":      basetypes.NumberType{},
										"name":        basetypes.StringType{},
										"nexthop": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"discard": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"fqdn":         basetypes.StringType{},
												"ipv6_address": basetypes.StringType{},
												"next_lr":      basetypes.StringType{},
												"next_vr":      basetypes.StringType{},
												"receive": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"tunnel": basetypes.StringType{},
											},
										},
										"path_monitor": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"enable":            basetypes.BoolType{},
												"failure_condition": basetypes.StringType{},
												"hold_time":         basetypes.NumberType{},
												"monitor_destinations": basetypes.ListType{ElemType: basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"count":            basetypes.NumberType{},
														"destination":      basetypes.StringType{},
														"destination_fqdn": basetypes.StringType{},
														"enable":           basetypes.BoolType{},
														"interval":         basetypes.NumberType{},
														"name":             basetypes.StringType{},
														"source":           basetypes.StringType{},
													},
												}},
											},
										},
										"route_table": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"both": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"multicast": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"no_install": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"unicast": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
											},
										},
									},
								}},
							},
						},
						"ipv6": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"static_route": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"admin_dist": basetypes.NumberType{},
										"bfd": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"profile": basetypes.StringType{},
											},
										},
										"destination": basetypes.StringType{},
										"interface":   basetypes.StringType{},
										"metric":      basetypes.NumberType{},
										"name":        basetypes.StringType{},
										"nexthop": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"discard": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"fqdn":         basetypes.StringType{},
												"ipv6_address": basetypes.StringType{},
												"next_lr":      basetypes.StringType{},
												"next_vr":      basetypes.StringType{},
												"receive": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"tunnel": basetypes.StringType{},
											},
										},
										"option": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"passive": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
											},
										},
										"path_monitor": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"enable":            basetypes.BoolType{},
												"failure_condition": basetypes.StringType{},
												"hold_time":         basetypes.NumberType{},
												"monitor_destinations": basetypes.ListType{ElemType: basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"count":            basetypes.NumberType{},
														"destination":      basetypes.StringType{},
														"destination_fqdn": basetypes.StringType{},
														"enable":           basetypes.BoolType{},
														"interval":         basetypes.NumberType{},
														"name":             basetypes.StringType{},
														"source":           basetypes.StringType{},
													},
												}},
											},
										},
										"route_table": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"both": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"multicast": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"no_install": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"unicast": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
											},
										},
									},
								}},
							},
						},
					},
				},
				"sdwan_type": basetypes.StringType{},
				"vr_admin_dists": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ebgp":        basetypes.NumberType{},
						"ibgp":        basetypes.NumberType{},
						"ospf_ext":    basetypes.NumberType{},
						"ospf_int":    basetypes.NumberType{},
						"ospfv3_ext":  basetypes.NumberType{},
						"ospfv3_int":  basetypes.NumberType{},
						"rip":         basetypes.NumberType{},
						"static":      basetypes.NumberType{},
						"static_ipv6": basetypes.NumberType{},
					},
				},
				"zone_name": basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of LogicalRouters objects.
func (o LogicalRouters) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInner model.
func (o LogicalRoutersVrfInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"admin_dists": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"bgp_external": basetypes.NumberType{},
				"bgp_internal": basetypes.NumberType{},
				"bgp_local":    basetypes.NumberType{},
				"ospf_ext":     basetypes.NumberType{},
				"ospf_inter":   basetypes.NumberType{},
				"ospf_intra":   basetypes.NumberType{},
				"ospfv3_ext":   basetypes.NumberType{},
				"ospfv3_inter": basetypes.NumberType{},
				"ospfv3_intra": basetypes.NumberType{},
				"rip":          basetypes.NumberType{},
				"static":       basetypes.NumberType{},
				"static_ipv6":  basetypes.NumberType{},
			},
		},
		"bgp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"advertise_network": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ipv4": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"network": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"backdoor":  basetypes.BoolType{},
										"multicast": basetypes.BoolType{},
										"name":      basetypes.StringType{},
										"unicast":   basetypes.BoolType{},
									},
								}},
							},
						},
						"ipv6": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"network": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"name":    basetypes.StringType{},
										"unicast": basetypes.BoolType{},
									},
								}},
							},
						},
					},
				},
				"aggregate": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"aggregate_med": basetypes.BoolType{},
					},
				},
				"aggregate_routes": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"as_set":       basetypes.BoolType{},
						"description":  basetypes.StringType{},
						"enable":       basetypes.BoolType{},
						"name":         basetypes.StringType{},
						"same_med":     basetypes.BoolType{},
						"summary_only": basetypes.BoolType{},
						"type": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"ipv4": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"attribute_map":  basetypes.StringType{},
										"summary_prefix": basetypes.StringType{},
										"suppress_map":   basetypes.StringType{},
									},
								},
								"ipv6": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"attribute_map":  basetypes.StringType{},
										"summary_prefix": basetypes.StringType{},
										"suppress_map":   basetypes.StringType{},
									},
								},
							},
						},
					},
				}},
				"allow_redist_default_route":     basetypes.BoolType{},
				"always_advertise_network_route": basetypes.BoolType{},
				"as_format":                      basetypes.StringType{},
				"confederation_member_as":        basetypes.StringType{},
				"default_local_preference":       basetypes.NumberType{},
				"ecmp_multi_as":                  basetypes.BoolType{},
				"enable":                         basetypes.BoolType{},
				"enforce_first_as":               basetypes.BoolType{},
				"fast_external_failover":         basetypes.BoolType{},
				"global_bfd": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"profile": basetypes.StringType{},
					},
				},
				"graceful_restart": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable":                basetypes.BoolType{},
						"local_restart_time":    basetypes.NumberType{},
						"max_peer_restart_time": basetypes.NumberType{},
						"stale_route_time":      basetypes.NumberType{},
					},
				},
				"graceful_shutdown": basetypes.BoolType{},
				"install_route":     basetypes.BoolType{},
				"local_as":          basetypes.StringType{},
				"med": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"always_compare_med":           basetypes.BoolType{},
						"deterministic_med_comparison": basetypes.BoolType{},
					},
				},
				"peer_group": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address_family": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"ipv4": basetypes.StringType{},
								"ipv6": basetypes.StringType{},
							},
						},
						"aggregated_confed_as_path": basetypes.BoolType{},
						"connection_options": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"authentication": basetypes.StringType{},
								"dampening":      basetypes.StringType{},
								"multihop":       basetypes.NumberType{},
								"timers":         basetypes.StringType{},
							},
						},
						"enable": basetypes.BoolType{},
						"filtering_profile": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"ipv4": basetypes.StringType{},
								"ipv6": basetypes.StringType{},
							},
						},
						"name": basetypes.StringType{},
						"peer": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"bfd": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"multihop": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"min_received_ttl": basetypes.NumberType{},
											},
										},
										"profile": basetypes.StringType{},
									},
								},
								"connection_options": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"authentication": basetypes.StringType{},
										"dampening":      basetypes.StringType{},
										"hold_time":      basetypes.StringType{},
										"idle_hold_time": basetypes.NumberType{},
										"incoming_bgp_connection": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"allow":       basetypes.BoolType{},
												"remote_port": basetypes.NumberType{},
											},
										},
										"keep_alive_interval":    basetypes.StringType{},
										"max_prefixes":           basetypes.StringType{},
										"min_route_adv_interval": basetypes.NumberType{},
										"multihop":               basetypes.StringType{},
										"open_delay_time":        basetypes.NumberType{},
										"outgoing_bgp_connection": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"allow":      basetypes.BoolType{},
												"local_port": basetypes.NumberType{},
											},
										},
										"timers": basetypes.StringType{},
									},
								},
								"enable":                            basetypes.BoolType{},
								"enable_mp_bgp":                     basetypes.BoolType{},
								"enable_sender_side_loop_detection": basetypes.BoolType{},
								"inherit": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"no": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"address_family": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"ipv4": basetypes.StringType{},
														"ipv6": basetypes.StringType{},
													},
												},
												"filtering_profile": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"ipv4": basetypes.StringType{},
														"ipv6": basetypes.StringType{},
													},
												},
											},
										},
										"yes": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								},
								"local_address": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"interface": basetypes.StringType{},
										"ip":        basetypes.StringType{},
									},
								},
								"name":    basetypes.StringType{},
								"passive": basetypes.BoolType{},
								"peer_address": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"fqdn": basetypes.StringType{},
										"ip":   basetypes.StringType{},
									},
								},
								"peer_as":          basetypes.StringType{},
								"peering_type":     basetypes.StringType{},
								"reflector_client": basetypes.StringType{},
								"subsequent_address_family_identifier": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"multicast": basetypes.BoolType{},
										"unicast":   basetypes.BoolType{},
									},
								},
							},
						}},
						"soft_reset_with_stored_info": basetypes.BoolType{},
						"type": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"ebgp": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"export_nexthop":    basetypes.StringType{},
										"import_nexthop":    basetypes.StringType{},
										"remove_private_as": basetypes.BoolType{},
									},
								},
								"ebgp_confed": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"export_nexthop": basetypes.StringType{},
									},
								},
								"ibgp": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"export_nexthop": basetypes.StringType{},
									},
								},
								"ibgp_confed": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"export_nexthop": basetypes.StringType{},
									},
								},
							},
						},
					},
				}},
				"policy": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"aggregation": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"advertise_filters": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"enable": basetypes.BoolType{},
												"match": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"exact": basetypes.BoolType{},
																"name":  basetypes.StringType{},
															},
														}},
														"afi": basetypes.StringType{},
														"as_path": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"regex": basetypes.StringType{},
															},
														},
														"community": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"regex": basetypes.StringType{},
															},
														},
														"extended_community": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"regex": basetypes.StringType{},
															},
														},
														"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
														"med":         basetypes.NumberType{},
														"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
														"route_table": basetypes.StringType{},
														"safi":        basetypes.StringType{},
													},
												},
												"name": basetypes.StringType{},
											},
										}},
										"aggregate_route_attributes": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"as_path": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"none": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
														"prepend": basetypes.NumberType{},
														"remove": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
														"remove_and_prepend": basetypes.NumberType{},
													},
												},
												"as_path_limit": basetypes.NumberType{},
												"community": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"append": basetypes.ListType{ElemType: basetypes.StringType{}},
														"none": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
														"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
														"remove_all": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
														"remove_regex": basetypes.StringType{},
													},
												},
												"extended_community": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"append": basetypes.ListType{ElemType: basetypes.StringType{}},
														"none": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
														"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
														"remove_all": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
														"remove_regex": basetypes.StringType{},
													},
												},
												"local_preference": basetypes.NumberType{},
												"med":              basetypes.NumberType{},
												"nexthop":          basetypes.StringType{},
												"origin":           basetypes.StringType{},
												"weight":           basetypes.NumberType{},
											},
										},
										"as_set":  basetypes.BoolType{},
										"enable":  basetypes.BoolType{},
										"name":    basetypes.StringType{},
										"prefix":  basetypes.StringType{},
										"summary": basetypes.BoolType{},
										"suppress_filters": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"enable": basetypes.BoolType{},
												"match": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"exact": basetypes.BoolType{},
																"name":  basetypes.StringType{},
															},
														}},
														"afi": basetypes.StringType{},
														"as_path": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"regex": basetypes.StringType{},
															},
														},
														"community": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"regex": basetypes.StringType{},
															},
														},
														"extended_community": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"regex": basetypes.StringType{},
															},
														},
														"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
														"med":         basetypes.NumberType{},
														"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
														"route_table": basetypes.StringType{},
														"safi":        basetypes.StringType{},
													},
												},
												"name": basetypes.StringType{},
											},
										}},
									},
								}},
							},
						},
						"conditional_advertisement": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"policy": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"advertise_filters": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"enable": basetypes.BoolType{},
												"match": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"exact": basetypes.BoolType{},
																"name":  basetypes.StringType{},
															},
														}},
														"afi": basetypes.StringType{},
														"as_path": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"regex": basetypes.StringType{},
															},
														},
														"community": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"regex": basetypes.StringType{},
															},
														},
														"extended_community": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"regex": basetypes.StringType{},
															},
														},
														"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
														"med":         basetypes.NumberType{},
														"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
														"route_table": basetypes.StringType{},
														"safi":        basetypes.StringType{},
													},
												},
												"name": basetypes.StringType{},
											},
										}},
										"enable": basetypes.BoolType{},
										"name":   basetypes.StringType{},
										"non_exist_filters": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"enable": basetypes.BoolType{},
												"match": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"exact": basetypes.BoolType{},
																"name":  basetypes.StringType{},
															},
														}},
														"afi": basetypes.StringType{},
														"as_path": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"regex": basetypes.StringType{},
															},
														},
														"community": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"regex": basetypes.StringType{},
															},
														},
														"extended_community": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"regex": basetypes.StringType{},
															},
														},
														"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
														"med":         basetypes.NumberType{},
														"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
														"route_table": basetypes.StringType{},
														"safi":        basetypes.StringType{},
													},
												},
												"name": basetypes.StringType{},
											},
										}},
										"used_by": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										}},
									},
								}},
							},
						},
						"export": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"rules": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"action": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"allow": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"update": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"as_path": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{
																		"none": basetypes.ObjectType{
																			AttrTypes: map[string]attr.Type{},
																		},
																		"prepend": basetypes.NumberType{},
																		"remove": basetypes.ObjectType{
																			AttrTypes: map[string]attr.Type{},
																		},
																		"remove_and_prepend": basetypes.NumberType{},
																	},
																},
																"as_path_limit": basetypes.NumberType{},
																"community": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{
																		"append": basetypes.ListType{ElemType: basetypes.StringType{}},
																		"none": basetypes.ObjectType{
																			AttrTypes: map[string]attr.Type{},
																		},
																		"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
																		"remove_all": basetypes.ObjectType{
																			AttrTypes: map[string]attr.Type{},
																		},
																		"remove_regex": basetypes.StringType{},
																	},
																},
																"extended_community": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{
																		"append": basetypes.ListType{ElemType: basetypes.StringType{}},
																		"none": basetypes.ObjectType{
																			AttrTypes: map[string]attr.Type{},
																		},
																		"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
																		"remove_all": basetypes.ObjectType{
																			AttrTypes: map[string]attr.Type{},
																		},
																		"remove_regex": basetypes.StringType{},
																	},
																},
																"local_preference": basetypes.NumberType{},
																"med":              basetypes.NumberType{},
																"nexthop":          basetypes.StringType{},
																"origin":           basetypes.StringType{},
															},
														},
													},
												},
												"deny": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
											},
										},
										"enable": basetypes.BoolType{},
										"match": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"exact": basetypes.BoolType{},
														"name":  basetypes.StringType{},
													},
												}},
												"afi": basetypes.StringType{},
												"as_path": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"regex": basetypes.StringType{},
													},
												},
												"community": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"regex": basetypes.StringType{},
													},
												},
												"extended_community": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"regex": basetypes.StringType{},
													},
												},
												"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
												"med":         basetypes.NumberType{},
												"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
												"route_table": basetypes.StringType{},
												"safi":        basetypes.StringType{},
											},
										},
										"name":    basetypes.StringType{},
										"used_by": basetypes.ListType{ElemType: basetypes.StringType{}},
									},
								}},
							},
						},
						"import": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"rules": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"action": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"allow": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"dampening": basetypes.StringType{},
														"update": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"as_path": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{
																		"none": basetypes.ObjectType{
																			AttrTypes: map[string]attr.Type{},
																		},
																		"prepend": basetypes.NumberType{},
																		"remove": basetypes.ObjectType{
																			AttrTypes: map[string]attr.Type{},
																		},
																		"remove_and_prepend": basetypes.NumberType{},
																	},
																},
																"as_path_limit": basetypes.NumberType{},
																"community": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{
																		"append": basetypes.ListType{ElemType: basetypes.StringType{}},
																		"none": basetypes.ObjectType{
																			AttrTypes: map[string]attr.Type{},
																		},
																		"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
																		"remove_all": basetypes.ObjectType{
																			AttrTypes: map[string]attr.Type{},
																		},
																		"remove_regex": basetypes.StringType{},
																	},
																},
																"extended_community": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{
																		"append": basetypes.ListType{ElemType: basetypes.StringType{}},
																		"none": basetypes.ObjectType{
																			AttrTypes: map[string]attr.Type{},
																		},
																		"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
																		"remove_all": basetypes.ObjectType{
																			AttrTypes: map[string]attr.Type{},
																		},
																		"remove_regex": basetypes.StringType{},
																	},
																},
																"local_preference": basetypes.NumberType{},
																"med":              basetypes.NumberType{},
																"nexthop":          basetypes.StringType{},
																"origin":           basetypes.StringType{},
																"weight":           basetypes.NumberType{},
															},
														},
													},
												},
												"deny": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
											},
										},
										"enable": basetypes.BoolType{},
										"match": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"exact": basetypes.BoolType{},
														"name":  basetypes.StringType{},
													},
												}},
												"afi": basetypes.StringType{},
												"as_path": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"regex": basetypes.StringType{},
													},
												},
												"community": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"regex": basetypes.StringType{},
													},
												},
												"extended_community": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"regex": basetypes.StringType{},
													},
												},
												"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
												"med":         basetypes.NumberType{},
												"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
												"route_table": basetypes.StringType{},
												"safi":        basetypes.StringType{},
											},
										},
										"name":    basetypes.StringType{},
										"used_by": basetypes.ListType{ElemType: basetypes.StringType{}},
									},
								}},
							},
						},
					},
				},
				"redist_rules": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address_family_identifier": basetypes.StringType{},
						"enable":                    basetypes.BoolType{},
						"metric":                    basetypes.NumberType{},
						"name":                      basetypes.StringType{},
						"route_table":               basetypes.StringType{},
						"set_as_path_limit":         basetypes.NumberType{},
						"set_community":             basetypes.ListType{ElemType: basetypes.StringType{}},
						"set_extended_community":    basetypes.ListType{ElemType: basetypes.StringType{}},
						"set_local_preference":      basetypes.NumberType{},
						"set_med":                   basetypes.NumberType{},
						"set_origin":                basetypes.StringType{},
					},
				}},
				"redistribution_profile": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ipv4": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"unicast": basetypes.StringType{},
							},
						},
						"ipv6": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"unicast": basetypes.StringType{},
							},
						},
					},
				},
				"reject_default_route": basetypes.BoolType{},
				"router_id":            basetypes.StringType{},
			},
		},
		"ecmp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"algorithm": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"balanced_round_robin": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"ip_hash": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"hash_seed": basetypes.NumberType{},
								"src_only":  basetypes.BoolType{},
								"use_port":  basetypes.BoolType{},
							},
						},
						"ip_modulo": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"weighted_round_robin": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"name":   basetypes.StringType{},
										"weight": basetypes.NumberType{},
									},
								}},
							},
						},
					},
				},
				"enable":             basetypes.BoolType{},
				"max_path":           basetypes.NumberType{},
				"strict_source_path": basetypes.BoolType{},
				"symmetric_return":   basetypes.BoolType{},
			},
		},
		"global_vrid": basetypes.NumberType{},
		"interface":   basetypes.ListType{ElemType: basetypes.StringType{}},
		"multicast": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable":    basetypes.BoolType{},
				"enable_v6": basetypes.BoolType{},
				"igmp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"dynamic": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"group_filter":          basetypes.StringType{},
										"max_groups":            basetypes.StringType{},
										"max_sources":           basetypes.StringType{},
										"name":                  basetypes.StringType{},
										"query_profile":         basetypes.StringType{},
										"robustness":            basetypes.StringType{},
										"router_alert_policing": basetypes.BoolType{},
										"version":               basetypes.StringType{},
									},
								}},
							},
						},
						"enable": basetypes.BoolType{},
						"static": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"group_address":  basetypes.StringType{},
								"interface":      basetypes.StringType{},
								"name":           basetypes.StringType{},
								"source_address": basetypes.StringType{},
							},
						}},
					},
				},
				"interface_group": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"description": basetypes.StringType{},
						"group_permission": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"any_source_multicast": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"group_address": basetypes.StringType{},
										"included":      basetypes.BoolType{},
										"name":          basetypes.StringType{},
									},
								}},
								"source_specific_multicast": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"group_address":  basetypes.StringType{},
										"included":       basetypes.BoolType{},
										"name":           basetypes.StringType{},
										"source_address": basetypes.StringType{},
									},
								}},
							},
						},
						"igmp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable":                     basetypes.BoolType{},
								"immediate_leave":            basetypes.BoolType{},
								"last_member_query_interval": basetypes.NumberType{},
								"max_groups":                 basetypes.StringType{},
								"max_query_response_time":    basetypes.NumberType{},
								"max_sources":                basetypes.StringType{},
								"mode":                       basetypes.StringType{},
								"query_interval":             basetypes.NumberType{},
								"robustness":                 basetypes.StringType{},
								"router_alert_policing":      basetypes.BoolType{},
								"version":                    basetypes.StringType{},
							},
						},
						"interface": basetypes.ListType{ElemType: basetypes.StringType{}},
						"name":      basetypes.StringType{},
						"pim": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"allowed_neighbors": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"name": basetypes.StringType{},
									},
								}},
								"assert_interval":     basetypes.NumberType{},
								"bsr_border":          basetypes.BoolType{},
								"dr_priority":         basetypes.NumberType{},
								"enable":              basetypes.BoolType{},
								"hello_interval":      basetypes.NumberType{},
								"join_prune_interval": basetypes.NumberType{},
							},
						},
					},
				}},
				"mode": basetypes.StringType{},
				"msdp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable":                basetypes.BoolType{},
						"global_authentication": basetypes.StringType{},
						"global_timer":          basetypes.StringType{},
						"originator_id": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"interface": basetypes.StringType{},
								"ip":        basetypes.StringType{},
							},
						},
						"peer": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"authentication":    basetypes.StringType{},
								"enable":            basetypes.BoolType{},
								"inbound_sa_filter": basetypes.StringType{},
								"local_address": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"interface": basetypes.StringType{},
										"ip":        basetypes.StringType{},
									},
								},
								"max_sa":             basetypes.NumberType{},
								"name":               basetypes.StringType{},
								"outbound_sa_filter": basetypes.StringType{},
								"peer_address": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"fqdn": basetypes.StringType{},
										"ip":   basetypes.StringType{},
									},
								},
								"peer_as": basetypes.StringType{},
							},
						}},
					},
				},
				"pim": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable":           basetypes.BoolType{},
						"group_permission": basetypes.StringType{},
						"if_timer_global":  basetypes.StringType{},
						"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"description":     basetypes.StringType{},
								"dr_priority":     basetypes.NumberType{},
								"if_timer":        basetypes.StringType{},
								"name":            basetypes.StringType{},
								"neighbor_filter": basetypes.StringType{},
								"send_bsm":        basetypes.BoolType{},
							},
						}},
						"route_ageout_time": basetypes.NumberType{},
						"rp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"external_rp": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"group_list": basetypes.StringType{},
										"name":       basetypes.StringType{},
										"override":   basetypes.BoolType{},
									},
								}},
								"local_rp": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"candidate_rp": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"address":                basetypes.StringType{},
												"advertisement_interval": basetypes.NumberType{},
												"group_list":             basetypes.StringType{},
												"interface":              basetypes.StringType{},
												"priority":               basetypes.NumberType{},
											},
										},
										"static_rp": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"address":    basetypes.StringType{},
												"group_list": basetypes.StringType{},
												"interface":  basetypes.StringType{},
												"override":   basetypes.BoolType{},
											},
										},
									},
								},
							},
						},
						"rpf_lookup_mode": basetypes.StringType{},
						"spt_threshold": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":      basetypes.StringType{},
								"threshold": basetypes.StringType{},
							},
						}},
						"ssm_address_space": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"group_list": basetypes.StringType{},
							},
						},
					},
				},
				"route_ageout_time": basetypes.NumberType{},
				"rp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"external_rp": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"group_addresses": basetypes.ListType{ElemType: basetypes.StringType{}},
								"name":            basetypes.StringType{},
								"override":        basetypes.BoolType{},
							},
						}},
						"local_rp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"candidate_rp": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"address":                basetypes.StringType{},
										"advertisement_interval": basetypes.NumberType{},
										"group_addresses":        basetypes.ListType{ElemType: basetypes.StringType{}},
										"interface":              basetypes.StringType{},
										"priority":               basetypes.NumberType{},
									},
								},
								"static_rp": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"address":         basetypes.StringType{},
										"group_addresses": basetypes.ListType{ElemType: basetypes.StringType{}},
										"interface":       basetypes.StringType{},
										"override":        basetypes.BoolType{},
									},
								},
							},
						},
					},
				},
				"spt_threshold": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":      basetypes.StringType{},
						"threshold": basetypes.StringType{},
					},
				}},
				"ssm_address_space": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"group_address": basetypes.StringType{},
						"included":      basetypes.BoolType{},
						"name":          basetypes.StringType{},
					},
				}},
				"static_route": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"destination": basetypes.StringType{},
						"interface":   basetypes.StringType{},
						"name":        basetypes.StringType{},
						"nexthop": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"ip_address": basetypes.StringType{},
							},
						},
						"preference": basetypes.NumberType{},
					},
				}},
			},
		},
		"name": basetypes.StringType{},
		"ospf": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"allow_redist_default_route": basetypes.BoolType{},
				"area": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"authentication": basetypes.StringType{},
						"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"authentication": basetypes.StringType{},
								"bfd": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"profile": basetypes.StringType{},
									},
								},
								"enable": basetypes.BoolType{},
								"link_type": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"broadcast": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"p2mp": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"neighbor": basetypes.ListType{ElemType: basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"name":     basetypes.StringType{},
														"priority": basetypes.NumberType{},
													},
												}},
											},
										},
										"p2p": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								},
								"metric":     basetypes.NumberType{},
								"mtu_ignore": basetypes.BoolType{},
								"name":       basetypes.StringType{},
								"passive":    basetypes.BoolType{},
								"priority":   basetypes.NumberType{},
								"timing":     basetypes.StringType{},
								"vr_timing": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"dead_counts":         basetypes.NumberType{},
										"gr_delay":            basetypes.NumberType{},
										"hello_interval":      basetypes.NumberType{},
										"retransmit_interval": basetypes.NumberType{},
										"transit_delay":       basetypes.NumberType{},
									},
								},
							},
						}},
						"name": basetypes.StringType{},
						"range": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"advertise":  basetypes.BoolType{},
								"name":       basetypes.StringType{},
								"substitute": basetypes.StringType{},
							},
						}},
						"type": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"normal": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"abr": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"export_list":          basetypes.StringType{},
												"import_list":          basetypes.StringType{},
												"inbound_filter_list":  basetypes.StringType{},
												"outbound_filter_list": basetypes.StringType{},
											},
										},
									},
								},
								"nssa": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"abr": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"export_list":         basetypes.StringType{},
												"import_list":         basetypes.StringType{},
												"inbound_filter_list": basetypes.StringType{},
												"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"advertise": basetypes.BoolType{},
														"name":      basetypes.StringType{},
														"route_tag": basetypes.NumberType{},
													},
												}},
												"outbound_filter_list": basetypes.StringType{},
											},
										},
										"accept_summary": basetypes.BoolType{},
										"default_information_originate": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"metric":      basetypes.NumberType{},
												"metric_type": basetypes.StringType{},
											},
										},
										"default_route": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"advertise": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"metric": basetypes.NumberType{},
														"type":   basetypes.StringType{},
													},
												},
												"disable": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
											},
										},
										"no_summary": basetypes.BoolType{},
										"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"advertise": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"name": basetypes.StringType{},
												"suppress": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
											},
										}},
									},
								},
								"stub": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"abr": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"export_list":          basetypes.StringType{},
												"import_list":          basetypes.StringType{},
												"inbound_filter_list":  basetypes.StringType{},
												"outbound_filter_list": basetypes.StringType{},
											},
										},
										"accept_summary": basetypes.BoolType{},
										"default_route": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"advertise": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"metric": basetypes.NumberType{},
													},
												},
												"disable": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
											},
										},
										"default_route_metric": basetypes.NumberType{},
										"no_summary":           basetypes.BoolType{},
									},
								},
							},
						},
						"virtual_link": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"authentication": basetypes.StringType{},
								"bfd": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"profile": basetypes.StringType{},
									},
								},
								"enable":          basetypes.BoolType{},
								"instance_id":     basetypes.NumberType{},
								"interface_id":    basetypes.NumberType{},
								"name":            basetypes.StringType{},
								"neighbor_id":     basetypes.StringType{},
								"passive":         basetypes.BoolType{},
								"timing":          basetypes.StringType{},
								"transit_area_id": basetypes.StringType{},
								"vr_timing": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"dead_counts":         basetypes.NumberType{},
										"hello_interval":      basetypes.NumberType{},
										"retransmit_interval": basetypes.NumberType{},
										"transit_delay":       basetypes.NumberType{},
									},
								},
							},
						}},
						"vr_range": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"advertise": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"name": basetypes.StringType{},
								"suppress": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						}},
					},
				}},
				"auth_profile": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"md5": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"key":       basetypes.StringType{},
								"name":      basetypes.NumberType{},
								"preferred": basetypes.BoolType{},
							},
						}},
						"name":     basetypes.StringType{},
						"password": basetypes.StringType{},
					},
				}},
				"enable": basetypes.BoolType{},
				"export_rules": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"metric":        basetypes.NumberType{},
						"name":          basetypes.StringType{},
						"new_path_type": basetypes.StringType{},
						"new_tag":       basetypes.StringType{},
					},
				}},
				"flood_prevention": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"hello": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable":     basetypes.BoolType{},
								"max_packet": basetypes.NumberType{},
							},
						},
						"lsa": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable":     basetypes.BoolType{},
								"max_packet": basetypes.NumberType{},
							},
						},
					},
				},
				"global_bfd": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"profile": basetypes.StringType{},
					},
				},
				"global_if_timer": basetypes.StringType{},
				"graceful_restart": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable":                    basetypes.BoolType{},
						"grace_period":              basetypes.NumberType{},
						"helper_enable":             basetypes.BoolType{},
						"max_neighbor_restart_time": basetypes.NumberType{},
						"strict__l_s_a_checking":    basetypes.BoolType{},
					},
				},
				"redistribution_profile": basetypes.StringType{},
				"reject_default_route":   basetypes.BoolType{},
				"rfc1583":                basetypes.BoolType{},
				"router_id":              basetypes.StringType{},
				"spf_timer":              basetypes.StringType{},
				"vr_timers": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"lsa_interval":          basetypes.NumberType{},
						"spf_calculation_delay": basetypes.NumberType{},
					},
				},
			},
		},
		"ospfv3": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"allow_redist_default_route": basetypes.BoolType{},
				"area": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"authentication": basetypes.StringType{},
						"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"authentication": basetypes.StringType{},
								"bfd": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"profile": basetypes.StringType{},
									},
								},
								"enable":      basetypes.BoolType{},
								"instance_id": basetypes.NumberType{},
								"link_type": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"broadcast": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"p2mp": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"neighbor": basetypes.ListType{ElemType: basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"name":     basetypes.StringType{},
														"priority": basetypes.NumberType{},
													},
												}},
											},
										},
										"p2p": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								},
								"metric":     basetypes.NumberType{},
								"mtu_ignore": basetypes.BoolType{},
								"name":       basetypes.StringType{},
								"neighbor": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"name": basetypes.StringType{},
									},
								}},
								"passive":  basetypes.BoolType{},
								"priority": basetypes.NumberType{},
								"timing":   basetypes.StringType{},
								"vr_timing": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"dead_counts":         basetypes.NumberType{},
										"gr_delay":            basetypes.NumberType{},
										"hello_interval":      basetypes.NumberType{},
										"retransmit_interval": basetypes.NumberType{},
										"transit_delay":       basetypes.NumberType{},
									},
								},
							},
						}},
						"name": basetypes.StringType{},
						"range": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"advertise": basetypes.BoolType{},
								"name":      basetypes.StringType{},
							},
						}},
						"type": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"normal": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"abr": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"export_list":          basetypes.StringType{},
												"import_list":          basetypes.StringType{},
												"inbound_filter_list":  basetypes.StringType{},
												"outbound_filter_list": basetypes.StringType{},
											},
										},
									},
								},
								"nssa": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"abr": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"export_list":         basetypes.StringType{},
												"import_list":         basetypes.StringType{},
												"inbound_filter_list": basetypes.StringType{},
												"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"advertise": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
														"name":      basetypes.StringType{},
														"route_tag": basetypes.NumberType{},
														"suppress": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
													},
												}},
												"outbound_filter_list": basetypes.StringType{},
											},
										},
										"accept_summary": basetypes.BoolType{},
										"default_information_originate": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"metric":      basetypes.NumberType{},
												"metric_type": basetypes.StringType{},
											},
										},
										"default_route": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"advertise": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"metric": basetypes.NumberType{},
														"type":   basetypes.StringType{},
													},
												},
												"disable": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
											},
										},
										"no_summary": basetypes.BoolType{},
										"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"advertise": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"name":      basetypes.StringType{},
												"route_tag": basetypes.NumberType{},
												"suppress": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
											},
										}},
									},
								},
								"stub": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"abr": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"export_list":          basetypes.StringType{},
												"import_list":          basetypes.StringType{},
												"inbound_filter_list":  basetypes.StringType{},
												"outbound_filter_list": basetypes.StringType{},
											},
										},
										"accept_summary": basetypes.BoolType{},
										"default_route": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"advertise": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"metric": basetypes.NumberType{},
													},
												},
												"disable": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
											},
										},
										"default_route_metric": basetypes.NumberType{},
										"no_summary":           basetypes.BoolType{},
									},
								},
							},
						},
						"virtual_link": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"authentication": basetypes.StringType{},
								"bfd": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"profile": basetypes.StringType{},
									},
								},
								"enable":          basetypes.BoolType{},
								"instance_id":     basetypes.NumberType{},
								"interface_id":    basetypes.NumberType{},
								"name":            basetypes.StringType{},
								"neighbor_id":     basetypes.StringType{},
								"passive":         basetypes.BoolType{},
								"timing":          basetypes.StringType{},
								"transit_area_id": basetypes.StringType{},
								"vr_timing": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"dead_counts":         basetypes.NumberType{},
										"hello_interval":      basetypes.NumberType{},
										"retransmit_interval": basetypes.NumberType{},
										"transit_delay":       basetypes.NumberType{},
									},
								},
							},
						}},
						"vr_range": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"advertise": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"name": basetypes.StringType{},
								"suppress": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						}},
					},
				}},
				"auth_profile": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ah": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"md5": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"key": basetypes.StringType{},
									},
								},
								"sha1": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"key": basetypes.StringType{},
									},
								},
								"sha256": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"key": basetypes.StringType{},
									},
								},
								"sha384": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"key": basetypes.StringType{},
									},
								},
								"sha512": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"key": basetypes.StringType{},
									},
								},
							},
						},
						"esp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"authentication": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"md5": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"key": basetypes.StringType{},
											},
										},
										"none": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"sha1": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"key": basetypes.StringType{},
											},
										},
										"sha256": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"key": basetypes.StringType{},
											},
										},
										"sha384": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"key": basetypes.StringType{},
											},
										},
										"sha512": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"key": basetypes.StringType{},
											},
										},
									},
								},
								"encryption": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"algorithm": basetypes.StringType{},
										"key":       basetypes.StringType{},
									},
								},
							},
						},
						"name": basetypes.StringType{},
						"spi":  basetypes.StringType{},
					},
				}},
				"disable_transit_traffic": basetypes.BoolType{},
				"enable":                  basetypes.BoolType{},
				"export_rules": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"metric":        basetypes.NumberType{},
						"name":          basetypes.StringType{},
						"new_path_type": basetypes.StringType{},
						"new_tag":       basetypes.StringType{},
					},
				}},
				"global_bfd": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"profile": basetypes.StringType{},
					},
				},
				"global_if_timer": basetypes.StringType{},
				"graceful_restart": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable":                    basetypes.BoolType{},
						"grace_period":              basetypes.NumberType{},
						"helper_enable":             basetypes.BoolType{},
						"max_neighbor_restart_time": basetypes.NumberType{},
						"strict__l_s_a_checking":    basetypes.BoolType{},
					},
				},
				"redistribution_profile": basetypes.StringType{},
				"reject_default_route":   basetypes.BoolType{},
				"router_id":              basetypes.StringType{},
				"spf_timer":              basetypes.StringType{},
				"vr_timers": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"lsa_interval":          basetypes.NumberType{},
						"spf_calculation_delay": basetypes.NumberType{},
					},
				},
			},
		},
		"rib_filter": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ipv4": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"bgp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"route_map": basetypes.StringType{},
							},
						},
						"ospf": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"route_map": basetypes.StringType{},
							},
						},
						"rip": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"route_map": basetypes.StringType{},
							},
						},
						"static": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"route_map": basetypes.StringType{},
							},
						},
					},
				},
				"ipv6": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"bgp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"route_map": basetypes.StringType{},
							},
						},
						"ospfv3": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"route_map": basetypes.StringType{},
							},
						},
						"static": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"route_map": basetypes.StringType{},
							},
						},
					},
				},
			},
		},
		"rip": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"auth_profile":                  basetypes.StringType{},
				"default_information_originate": basetypes.BoolType{},
				"enable":                        basetypes.BoolType{},
				"global_bfd": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"profile": basetypes.StringType{},
					},
				},
				"global_inbound_distribute_list": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"access_list": basetypes.StringType{},
					},
				},
				"global_outbound_distribute_list": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"access_list": basetypes.StringType{},
					},
				},
				"global_timer": basetypes.StringType{},
				"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"authentication": basetypes.StringType{},
						"bfd": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"profile": basetypes.StringType{},
							},
						},
						"enable": basetypes.BoolType{},
						"interface_inbound_distribute_list": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"access_list": basetypes.StringType{},
								"metric":      basetypes.NumberType{},
							},
						},
						"interface_outbound_distribute_list": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"access_list": basetypes.StringType{},
								"metric":      basetypes.NumberType{},
							},
						},
						"mode":          basetypes.StringType{},
						"name":          basetypes.StringType{},
						"split_horizon": basetypes.StringType{},
					},
				}},
				"redistribution_profile": basetypes.StringType{},
			},
		},
		"routing_table": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ip": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"static_route": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"admin_dist": basetypes.NumberType{},
								"bfd": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"profile": basetypes.StringType{},
									},
								},
								"destination": basetypes.StringType{},
								"interface":   basetypes.StringType{},
								"metric":      basetypes.NumberType{},
								"name":        basetypes.StringType{},
								"nexthop": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"discard": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"fqdn":         basetypes.StringType{},
										"ipv6_address": basetypes.StringType{},
										"next_lr":      basetypes.StringType{},
										"next_vr":      basetypes.StringType{},
										"receive": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"tunnel": basetypes.StringType{},
									},
								},
								"path_monitor": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"enable":            basetypes.BoolType{},
										"failure_condition": basetypes.StringType{},
										"hold_time":         basetypes.NumberType{},
										"monitor_destinations": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"count":            basetypes.NumberType{},
												"destination":      basetypes.StringType{},
												"destination_fqdn": basetypes.StringType{},
												"enable":           basetypes.BoolType{},
												"interval":         basetypes.NumberType{},
												"name":             basetypes.StringType{},
												"source":           basetypes.StringType{},
											},
										}},
									},
								},
								"route_table": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"both": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"multicast": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"no_install": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"unicast": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								},
							},
						}},
					},
				},
				"ipv6": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"static_route": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"admin_dist": basetypes.NumberType{},
								"bfd": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"profile": basetypes.StringType{},
									},
								},
								"destination": basetypes.StringType{},
								"interface":   basetypes.StringType{},
								"metric":      basetypes.NumberType{},
								"name":        basetypes.StringType{},
								"nexthop": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"discard": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"fqdn":         basetypes.StringType{},
										"ipv6_address": basetypes.StringType{},
										"next_lr":      basetypes.StringType{},
										"next_vr":      basetypes.StringType{},
										"receive": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"tunnel": basetypes.StringType{},
									},
								},
								"option": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"passive": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								},
								"path_monitor": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"enable":            basetypes.BoolType{},
										"failure_condition": basetypes.StringType{},
										"hold_time":         basetypes.NumberType{},
										"monitor_destinations": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"count":            basetypes.NumberType{},
												"destination":      basetypes.StringType{},
												"destination_fqdn": basetypes.StringType{},
												"enable":           basetypes.BoolType{},
												"interval":         basetypes.NumberType{},
												"name":             basetypes.StringType{},
												"source":           basetypes.StringType{},
											},
										}},
									},
								},
								"route_table": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"both": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"multicast": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"no_install": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"unicast": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								},
							},
						}},
					},
				},
			},
		},
		"sdwan_type": basetypes.StringType{},
		"vr_admin_dists": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ebgp":        basetypes.NumberType{},
				"ibgp":        basetypes.NumberType{},
				"ospf_ext":    basetypes.NumberType{},
				"ospf_int":    basetypes.NumberType{},
				"ospfv3_ext":  basetypes.NumberType{},
				"ospfv3_int":  basetypes.NumberType{},
				"rip":         basetypes.NumberType{},
				"static":      basetypes.NumberType{},
				"static_ipv6": basetypes.NumberType{},
			},
		},
		"zone_name": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInner objects.
func (o LogicalRoutersVrfInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerAdminDists model.
func (o LogicalRoutersVrfInnerAdminDists) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"bgp_external": basetypes.NumberType{},
		"bgp_internal": basetypes.NumberType{},
		"bgp_local":    basetypes.NumberType{},
		"ospf_ext":     basetypes.NumberType{},
		"ospf_inter":   basetypes.NumberType{},
		"ospf_intra":   basetypes.NumberType{},
		"ospfv3_ext":   basetypes.NumberType{},
		"ospfv3_inter": basetypes.NumberType{},
		"ospfv3_intra": basetypes.NumberType{},
		"rip":          basetypes.NumberType{},
		"static":       basetypes.NumberType{},
		"static_ipv6":  basetypes.NumberType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerAdminDists objects.
func (o LogicalRoutersVrfInnerAdminDists) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgp model.
func (o LogicalRoutersVrfInnerBgp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"advertise_network": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ipv4": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"network": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"backdoor":  basetypes.BoolType{},
								"multicast": basetypes.BoolType{},
								"name":      basetypes.StringType{},
								"unicast":   basetypes.BoolType{},
							},
						}},
					},
				},
				"ipv6": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"network": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":    basetypes.StringType{},
								"unicast": basetypes.BoolType{},
							},
						}},
					},
				},
			},
		},
		"aggregate": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"aggregate_med": basetypes.BoolType{},
			},
		},
		"aggregate_routes": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"as_set":       basetypes.BoolType{},
				"description":  basetypes.StringType{},
				"enable":       basetypes.BoolType{},
				"name":         basetypes.StringType{},
				"same_med":     basetypes.BoolType{},
				"summary_only": basetypes.BoolType{},
				"type": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ipv4": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"attribute_map":  basetypes.StringType{},
								"summary_prefix": basetypes.StringType{},
								"suppress_map":   basetypes.StringType{},
							},
						},
						"ipv6": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"attribute_map":  basetypes.StringType{},
								"summary_prefix": basetypes.StringType{},
								"suppress_map":   basetypes.StringType{},
							},
						},
					},
				},
			},
		}},
		"allow_redist_default_route":     basetypes.BoolType{},
		"always_advertise_network_route": basetypes.BoolType{},
		"as_format":                      basetypes.StringType{},
		"confederation_member_as":        basetypes.StringType{},
		"default_local_preference":       basetypes.NumberType{},
		"ecmp_multi_as":                  basetypes.BoolType{},
		"enable":                         basetypes.BoolType{},
		"enforce_first_as":               basetypes.BoolType{},
		"fast_external_failover":         basetypes.BoolType{},
		"global_bfd": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"profile": basetypes.StringType{},
			},
		},
		"graceful_restart": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable":                basetypes.BoolType{},
				"local_restart_time":    basetypes.NumberType{},
				"max_peer_restart_time": basetypes.NumberType{},
				"stale_route_time":      basetypes.NumberType{},
			},
		},
		"graceful_shutdown": basetypes.BoolType{},
		"install_route":     basetypes.BoolType{},
		"local_as":          basetypes.StringType{},
		"med": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"always_compare_med":           basetypes.BoolType{},
				"deterministic_med_comparison": basetypes.BoolType{},
			},
		},
		"peer_group": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address_family": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ipv4": basetypes.StringType{},
						"ipv6": basetypes.StringType{},
					},
				},
				"aggregated_confed_as_path": basetypes.BoolType{},
				"connection_options": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"authentication": basetypes.StringType{},
						"dampening":      basetypes.StringType{},
						"multihop":       basetypes.NumberType{},
						"timers":         basetypes.StringType{},
					},
				},
				"enable": basetypes.BoolType{},
				"filtering_profile": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ipv4": basetypes.StringType{},
						"ipv6": basetypes.StringType{},
					},
				},
				"name": basetypes.StringType{},
				"peer": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"bfd": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"multihop": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"min_received_ttl": basetypes.NumberType{},
									},
								},
								"profile": basetypes.StringType{},
							},
						},
						"connection_options": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"authentication": basetypes.StringType{},
								"dampening":      basetypes.StringType{},
								"hold_time":      basetypes.StringType{},
								"idle_hold_time": basetypes.NumberType{},
								"incoming_bgp_connection": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"allow":       basetypes.BoolType{},
										"remote_port": basetypes.NumberType{},
									},
								},
								"keep_alive_interval":    basetypes.StringType{},
								"max_prefixes":           basetypes.StringType{},
								"min_route_adv_interval": basetypes.NumberType{},
								"multihop":               basetypes.StringType{},
								"open_delay_time":        basetypes.NumberType{},
								"outgoing_bgp_connection": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"allow":      basetypes.BoolType{},
										"local_port": basetypes.NumberType{},
									},
								},
								"timers": basetypes.StringType{},
							},
						},
						"enable":                            basetypes.BoolType{},
						"enable_mp_bgp":                     basetypes.BoolType{},
						"enable_sender_side_loop_detection": basetypes.BoolType{},
						"inherit": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"no": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"address_family": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"ipv4": basetypes.StringType{},
												"ipv6": basetypes.StringType{},
											},
										},
										"filtering_profile": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"ipv4": basetypes.StringType{},
												"ipv6": basetypes.StringType{},
											},
										},
									},
								},
								"yes": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
						"local_address": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"interface": basetypes.StringType{},
								"ip":        basetypes.StringType{},
							},
						},
						"name":    basetypes.StringType{},
						"passive": basetypes.BoolType{},
						"peer_address": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"fqdn": basetypes.StringType{},
								"ip":   basetypes.StringType{},
							},
						},
						"peer_as":          basetypes.StringType{},
						"peering_type":     basetypes.StringType{},
						"reflector_client": basetypes.StringType{},
						"subsequent_address_family_identifier": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"multicast": basetypes.BoolType{},
								"unicast":   basetypes.BoolType{},
							},
						},
					},
				}},
				"soft_reset_with_stored_info": basetypes.BoolType{},
				"type": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ebgp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"export_nexthop":    basetypes.StringType{},
								"import_nexthop":    basetypes.StringType{},
								"remove_private_as": basetypes.BoolType{},
							},
						},
						"ebgp_confed": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"export_nexthop": basetypes.StringType{},
							},
						},
						"ibgp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"export_nexthop": basetypes.StringType{},
							},
						},
						"ibgp_confed": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"export_nexthop": basetypes.StringType{},
							},
						},
					},
				},
			},
		}},
		"policy": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"aggregation": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"advertise_filters": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"enable": basetypes.BoolType{},
										"match": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"exact": basetypes.BoolType{},
														"name":  basetypes.StringType{},
													},
												}},
												"afi": basetypes.StringType{},
												"as_path": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"regex": basetypes.StringType{},
													},
												},
												"community": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"regex": basetypes.StringType{},
													},
												},
												"extended_community": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"regex": basetypes.StringType{},
													},
												},
												"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
												"med":         basetypes.NumberType{},
												"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
												"route_table": basetypes.StringType{},
												"safi":        basetypes.StringType{},
											},
										},
										"name": basetypes.StringType{},
									},
								}},
								"aggregate_route_attributes": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"as_path": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"none": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"prepend": basetypes.NumberType{},
												"remove": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"remove_and_prepend": basetypes.NumberType{},
											},
										},
										"as_path_limit": basetypes.NumberType{},
										"community": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"append": basetypes.ListType{ElemType: basetypes.StringType{}},
												"none": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
												"remove_all": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"remove_regex": basetypes.StringType{},
											},
										},
										"extended_community": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"append": basetypes.ListType{ElemType: basetypes.StringType{}},
												"none": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
												"remove_all": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"remove_regex": basetypes.StringType{},
											},
										},
										"local_preference": basetypes.NumberType{},
										"med":              basetypes.NumberType{},
										"nexthop":          basetypes.StringType{},
										"origin":           basetypes.StringType{},
										"weight":           basetypes.NumberType{},
									},
								},
								"as_set":  basetypes.BoolType{},
								"enable":  basetypes.BoolType{},
								"name":    basetypes.StringType{},
								"prefix":  basetypes.StringType{},
								"summary": basetypes.BoolType{},
								"suppress_filters": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"enable": basetypes.BoolType{},
										"match": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"exact": basetypes.BoolType{},
														"name":  basetypes.StringType{},
													},
												}},
												"afi": basetypes.StringType{},
												"as_path": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"regex": basetypes.StringType{},
													},
												},
												"community": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"regex": basetypes.StringType{},
													},
												},
												"extended_community": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"regex": basetypes.StringType{},
													},
												},
												"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
												"med":         basetypes.NumberType{},
												"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
												"route_table": basetypes.StringType{},
												"safi":        basetypes.StringType{},
											},
										},
										"name": basetypes.StringType{},
									},
								}},
							},
						}},
					},
				},
				"conditional_advertisement": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"policy": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"advertise_filters": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"enable": basetypes.BoolType{},
										"match": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"exact": basetypes.BoolType{},
														"name":  basetypes.StringType{},
													},
												}},
												"afi": basetypes.StringType{},
												"as_path": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"regex": basetypes.StringType{},
													},
												},
												"community": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"regex": basetypes.StringType{},
													},
												},
												"extended_community": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"regex": basetypes.StringType{},
													},
												},
												"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
												"med":         basetypes.NumberType{},
												"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
												"route_table": basetypes.StringType{},
												"safi":        basetypes.StringType{},
											},
										},
										"name": basetypes.StringType{},
									},
								}},
								"enable": basetypes.BoolType{},
								"name":   basetypes.StringType{},
								"non_exist_filters": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"enable": basetypes.BoolType{},
										"match": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"exact": basetypes.BoolType{},
														"name":  basetypes.StringType{},
													},
												}},
												"afi": basetypes.StringType{},
												"as_path": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"regex": basetypes.StringType{},
													},
												},
												"community": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"regex": basetypes.StringType{},
													},
												},
												"extended_community": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"regex": basetypes.StringType{},
													},
												},
												"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
												"med":         basetypes.NumberType{},
												"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
												"route_table": basetypes.StringType{},
												"safi":        basetypes.StringType{},
											},
										},
										"name": basetypes.StringType{},
									},
								}},
								"used_by": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								}},
							},
						}},
					},
				},
				"export": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"rules": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"allow": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"update": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"as_path": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"none": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{},
																},
																"prepend": basetypes.NumberType{},
																"remove": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{},
																},
																"remove_and_prepend": basetypes.NumberType{},
															},
														},
														"as_path_limit": basetypes.NumberType{},
														"community": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"append": basetypes.ListType{ElemType: basetypes.StringType{}},
																"none": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{},
																},
																"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
																"remove_all": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{},
																},
																"remove_regex": basetypes.StringType{},
															},
														},
														"extended_community": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"append": basetypes.ListType{ElemType: basetypes.StringType{}},
																"none": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{},
																},
																"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
																"remove_all": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{},
																},
																"remove_regex": basetypes.StringType{},
															},
														},
														"local_preference": basetypes.NumberType{},
														"med":              basetypes.NumberType{},
														"nexthop":          basetypes.StringType{},
														"origin":           basetypes.StringType{},
													},
												},
											},
										},
										"deny": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								},
								"enable": basetypes.BoolType{},
								"match": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"exact": basetypes.BoolType{},
												"name":  basetypes.StringType{},
											},
										}},
										"afi": basetypes.StringType{},
										"as_path": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"regex": basetypes.StringType{},
											},
										},
										"community": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"regex": basetypes.StringType{},
											},
										},
										"extended_community": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"regex": basetypes.StringType{},
											},
										},
										"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
										"med":         basetypes.NumberType{},
										"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
										"route_table": basetypes.StringType{},
										"safi":        basetypes.StringType{},
									},
								},
								"name":    basetypes.StringType{},
								"used_by": basetypes.ListType{ElemType: basetypes.StringType{}},
							},
						}},
					},
				},
				"import": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"rules": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"allow": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"dampening": basetypes.StringType{},
												"update": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"as_path": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"none": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{},
																},
																"prepend": basetypes.NumberType{},
																"remove": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{},
																},
																"remove_and_prepend": basetypes.NumberType{},
															},
														},
														"as_path_limit": basetypes.NumberType{},
														"community": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"append": basetypes.ListType{ElemType: basetypes.StringType{}},
																"none": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{},
																},
																"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
																"remove_all": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{},
																},
																"remove_regex": basetypes.StringType{},
															},
														},
														"extended_community": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"append": basetypes.ListType{ElemType: basetypes.StringType{}},
																"none": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{},
																},
																"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
																"remove_all": basetypes.ObjectType{
																	AttrTypes: map[string]attr.Type{},
																},
																"remove_regex": basetypes.StringType{},
															},
														},
														"local_preference": basetypes.NumberType{},
														"med":              basetypes.NumberType{},
														"nexthop":          basetypes.StringType{},
														"origin":           basetypes.StringType{},
														"weight":           basetypes.NumberType{},
													},
												},
											},
										},
										"deny": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								},
								"enable": basetypes.BoolType{},
								"match": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"exact": basetypes.BoolType{},
												"name":  basetypes.StringType{},
											},
										}},
										"afi": basetypes.StringType{},
										"as_path": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"regex": basetypes.StringType{},
											},
										},
										"community": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"regex": basetypes.StringType{},
											},
										},
										"extended_community": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"regex": basetypes.StringType{},
											},
										},
										"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
										"med":         basetypes.NumberType{},
										"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
										"route_table": basetypes.StringType{},
										"safi":        basetypes.StringType{},
									},
								},
								"name":    basetypes.StringType{},
								"used_by": basetypes.ListType{ElemType: basetypes.StringType{}},
							},
						}},
					},
				},
			},
		},
		"redist_rules": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address_family_identifier": basetypes.StringType{},
				"enable":                    basetypes.BoolType{},
				"metric":                    basetypes.NumberType{},
				"name":                      basetypes.StringType{},
				"route_table":               basetypes.StringType{},
				"set_as_path_limit":         basetypes.NumberType{},
				"set_community":             basetypes.ListType{ElemType: basetypes.StringType{}},
				"set_extended_community":    basetypes.ListType{ElemType: basetypes.StringType{}},
				"set_local_preference":      basetypes.NumberType{},
				"set_med":                   basetypes.NumberType{},
				"set_origin":                basetypes.StringType{},
			},
		}},
		"redistribution_profile": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ipv4": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"unicast": basetypes.StringType{},
					},
				},
				"ipv6": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"unicast": basetypes.StringType{},
					},
				},
			},
		},
		"reject_default_route": basetypes.BoolType{},
		"router_id":            basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgp objects.
func (o LogicalRoutersVrfInnerBgp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpAdvertiseNetwork model.
func (o LogicalRoutersVrfInnerBgpAdvertiseNetwork) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ipv4": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"network": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"backdoor":  basetypes.BoolType{},
						"multicast": basetypes.BoolType{},
						"name":      basetypes.StringType{},
						"unicast":   basetypes.BoolType{},
					},
				}},
			},
		},
		"ipv6": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"network": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":    basetypes.StringType{},
						"unicast": basetypes.BoolType{},
					},
				}},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpAdvertiseNetwork objects.
func (o LogicalRoutersVrfInnerBgpAdvertiseNetwork) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpAdvertiseNetworkIpv4 model.
func (o LogicalRoutersVrfInnerBgpAdvertiseNetworkIpv4) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"network": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"backdoor":  basetypes.BoolType{},
				"multicast": basetypes.BoolType{},
				"name":      basetypes.StringType{},
				"unicast":   basetypes.BoolType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpAdvertiseNetworkIpv4 objects.
func (o LogicalRoutersVrfInnerBgpAdvertiseNetworkIpv4) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpAdvertiseNetworkIpv4NetworkInner model.
func (o LogicalRoutersVrfInnerBgpAdvertiseNetworkIpv4NetworkInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"backdoor":  basetypes.BoolType{},
		"multicast": basetypes.BoolType{},
		"name":      basetypes.StringType{},
		"unicast":   basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpAdvertiseNetworkIpv4NetworkInner objects.
func (o LogicalRoutersVrfInnerBgpAdvertiseNetworkIpv4NetworkInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpAdvertiseNetworkIpv6 model.
func (o LogicalRoutersVrfInnerBgpAdvertiseNetworkIpv6) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"network": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":    basetypes.StringType{},
				"unicast": basetypes.BoolType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpAdvertiseNetworkIpv6 objects.
func (o LogicalRoutersVrfInnerBgpAdvertiseNetworkIpv6) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpAdvertiseNetworkIpv6NetworkInner model.
func (o LogicalRoutersVrfInnerBgpAdvertiseNetworkIpv6NetworkInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":    basetypes.StringType{},
		"unicast": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpAdvertiseNetworkIpv6NetworkInner objects.
func (o LogicalRoutersVrfInnerBgpAdvertiseNetworkIpv6NetworkInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpAggregate model.
func (o LogicalRoutersVrfInnerBgpAggregate) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"aggregate_med": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpAggregate objects.
func (o LogicalRoutersVrfInnerBgpAggregate) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpAggregateRoutesInner model.
func (o LogicalRoutersVrfInnerBgpAggregateRoutesInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"as_set":       basetypes.BoolType{},
		"description":  basetypes.StringType{},
		"enable":       basetypes.BoolType{},
		"name":         basetypes.StringType{},
		"same_med":     basetypes.BoolType{},
		"summary_only": basetypes.BoolType{},
		"type": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ipv4": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"attribute_map":  basetypes.StringType{},
						"summary_prefix": basetypes.StringType{},
						"suppress_map":   basetypes.StringType{},
					},
				},
				"ipv6": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"attribute_map":  basetypes.StringType{},
						"summary_prefix": basetypes.StringType{},
						"suppress_map":   basetypes.StringType{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpAggregateRoutesInner objects.
func (o LogicalRoutersVrfInnerBgpAggregateRoutesInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpAggregateRoutesInnerType model.
func (o LogicalRoutersVrfInnerBgpAggregateRoutesInnerType) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ipv4": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"attribute_map":  basetypes.StringType{},
				"summary_prefix": basetypes.StringType{},
				"suppress_map":   basetypes.StringType{},
			},
		},
		"ipv6": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"attribute_map":  basetypes.StringType{},
				"summary_prefix": basetypes.StringType{},
				"suppress_map":   basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpAggregateRoutesInnerType objects.
func (o LogicalRoutersVrfInnerBgpAggregateRoutesInnerType) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpAggregateRoutesInnerTypeIpv4 model.
func (o LogicalRoutersVrfInnerBgpAggregateRoutesInnerTypeIpv4) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"attribute_map":  basetypes.StringType{},
		"summary_prefix": basetypes.StringType{},
		"suppress_map":   basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpAggregateRoutesInnerTypeIpv4 objects.
func (o LogicalRoutersVrfInnerBgpAggregateRoutesInnerTypeIpv4) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpGlobalBfd model.
func (o LogicalRoutersVrfInnerBgpGlobalBfd) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"profile": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpGlobalBfd objects.
func (o LogicalRoutersVrfInnerBgpGlobalBfd) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpGracefulRestart model.
func (o LogicalRoutersVrfInnerBgpGracefulRestart) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable":                basetypes.BoolType{},
		"local_restart_time":    basetypes.NumberType{},
		"max_peer_restart_time": basetypes.NumberType{},
		"stale_route_time":      basetypes.NumberType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpGracefulRestart objects.
func (o LogicalRoutersVrfInnerBgpGracefulRestart) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpMed model.
func (o LogicalRoutersVrfInnerBgpMed) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"always_compare_med":           basetypes.BoolType{},
		"deterministic_med_comparison": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpMed objects.
func (o LogicalRoutersVrfInnerBgpMed) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPeerGroupInner model.
func (o LogicalRoutersVrfInnerBgpPeerGroupInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address_family": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ipv4": basetypes.StringType{},
				"ipv6": basetypes.StringType{},
			},
		},
		"aggregated_confed_as_path": basetypes.BoolType{},
		"connection_options": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"authentication": basetypes.StringType{},
				"dampening":      basetypes.StringType{},
				"multihop":       basetypes.NumberType{},
				"timers":         basetypes.StringType{},
			},
		},
		"enable": basetypes.BoolType{},
		"filtering_profile": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ipv4": basetypes.StringType{},
				"ipv6": basetypes.StringType{},
			},
		},
		"name": basetypes.StringType{},
		"peer": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"bfd": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"multihop": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"min_received_ttl": basetypes.NumberType{},
							},
						},
						"profile": basetypes.StringType{},
					},
				},
				"connection_options": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"authentication": basetypes.StringType{},
						"dampening":      basetypes.StringType{},
						"hold_time":      basetypes.StringType{},
						"idle_hold_time": basetypes.NumberType{},
						"incoming_bgp_connection": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"allow":       basetypes.BoolType{},
								"remote_port": basetypes.NumberType{},
							},
						},
						"keep_alive_interval":    basetypes.StringType{},
						"max_prefixes":           basetypes.StringType{},
						"min_route_adv_interval": basetypes.NumberType{},
						"multihop":               basetypes.StringType{},
						"open_delay_time":        basetypes.NumberType{},
						"outgoing_bgp_connection": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"allow":      basetypes.BoolType{},
								"local_port": basetypes.NumberType{},
							},
						},
						"timers": basetypes.StringType{},
					},
				},
				"enable":                            basetypes.BoolType{},
				"enable_mp_bgp":                     basetypes.BoolType{},
				"enable_sender_side_loop_detection": basetypes.BoolType{},
				"inherit": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"no": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address_family": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"ipv4": basetypes.StringType{},
										"ipv6": basetypes.StringType{},
									},
								},
								"filtering_profile": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"ipv4": basetypes.StringType{},
										"ipv6": basetypes.StringType{},
									},
								},
							},
						},
						"yes": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
				"local_address": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"interface": basetypes.StringType{},
						"ip":        basetypes.StringType{},
					},
				},
				"name":    basetypes.StringType{},
				"passive": basetypes.BoolType{},
				"peer_address": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"fqdn": basetypes.StringType{},
						"ip":   basetypes.StringType{},
					},
				},
				"peer_as":          basetypes.StringType{},
				"peering_type":     basetypes.StringType{},
				"reflector_client": basetypes.StringType{},
				"subsequent_address_family_identifier": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"multicast": basetypes.BoolType{},
						"unicast":   basetypes.BoolType{},
					},
				},
			},
		}},
		"soft_reset_with_stored_info": basetypes.BoolType{},
		"type": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ebgp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"export_nexthop":    basetypes.StringType{},
						"import_nexthop":    basetypes.StringType{},
						"remove_private_as": basetypes.BoolType{},
					},
				},
				"ebgp_confed": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"export_nexthop": basetypes.StringType{},
					},
				},
				"ibgp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"export_nexthop": basetypes.StringType{},
					},
				},
				"ibgp_confed": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"export_nexthop": basetypes.StringType{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPeerGroupInner objects.
func (o LogicalRoutersVrfInnerBgpPeerGroupInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPeerGroupInnerAddressFamily model.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerAddressFamily) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ipv4": basetypes.StringType{},
		"ipv6": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPeerGroupInnerAddressFamily objects.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerAddressFamily) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPeerGroupInnerConnectionOptions model.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerConnectionOptions) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"authentication": basetypes.StringType{},
		"dampening":      basetypes.StringType{},
		"multihop":       basetypes.NumberType{},
		"timers":         basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPeerGroupInnerConnectionOptions objects.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerConnectionOptions) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInner model.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"bfd": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"multihop": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"min_received_ttl": basetypes.NumberType{},
					},
				},
				"profile": basetypes.StringType{},
			},
		},
		"connection_options": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"authentication": basetypes.StringType{},
				"dampening":      basetypes.StringType{},
				"hold_time":      basetypes.StringType{},
				"idle_hold_time": basetypes.NumberType{},
				"incoming_bgp_connection": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"allow":       basetypes.BoolType{},
						"remote_port": basetypes.NumberType{},
					},
				},
				"keep_alive_interval":    basetypes.StringType{},
				"max_prefixes":           basetypes.StringType{},
				"min_route_adv_interval": basetypes.NumberType{},
				"multihop":               basetypes.StringType{},
				"open_delay_time":        basetypes.NumberType{},
				"outgoing_bgp_connection": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"allow":      basetypes.BoolType{},
						"local_port": basetypes.NumberType{},
					},
				},
				"timers": basetypes.StringType{},
			},
		},
		"enable":                            basetypes.BoolType{},
		"enable_mp_bgp":                     basetypes.BoolType{},
		"enable_sender_side_loop_detection": basetypes.BoolType{},
		"inherit": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"no": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address_family": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"ipv4": basetypes.StringType{},
								"ipv6": basetypes.StringType{},
							},
						},
						"filtering_profile": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"ipv4": basetypes.StringType{},
								"ipv6": basetypes.StringType{},
							},
						},
					},
				},
				"yes": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
		"local_address": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"interface": basetypes.StringType{},
				"ip":        basetypes.StringType{},
			},
		},
		"name":    basetypes.StringType{},
		"passive": basetypes.BoolType{},
		"peer_address": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"fqdn": basetypes.StringType{},
				"ip":   basetypes.StringType{},
			},
		},
		"peer_as":          basetypes.StringType{},
		"peering_type":     basetypes.StringType{},
		"reflector_client": basetypes.StringType{},
		"subsequent_address_family_identifier": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"multicast": basetypes.BoolType{},
				"unicast":   basetypes.BoolType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInner objects.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerBfd model.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerBfd) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"multihop": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"min_received_ttl": basetypes.NumberType{},
			},
		},
		"profile": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerBfd objects.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerBfd) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerBfdMultihop model.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerBfdMultihop) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"min_received_ttl": basetypes.NumberType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerBfdMultihop objects.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerBfdMultihop) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerConnectionOptions model.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerConnectionOptions) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"authentication": basetypes.StringType{},
		"dampening":      basetypes.StringType{},
		"hold_time":      basetypes.StringType{},
		"idle_hold_time": basetypes.NumberType{},
		"incoming_bgp_connection": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"allow":       basetypes.BoolType{},
				"remote_port": basetypes.NumberType{},
			},
		},
		"keep_alive_interval":    basetypes.StringType{},
		"max_prefixes":           basetypes.StringType{},
		"min_route_adv_interval": basetypes.NumberType{},
		"multihop":               basetypes.StringType{},
		"open_delay_time":        basetypes.NumberType{},
		"outgoing_bgp_connection": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"allow":      basetypes.BoolType{},
				"local_port": basetypes.NumberType{},
			},
		},
		"timers": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerConnectionOptions objects.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerConnectionOptions) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerConnectionOptionsIncomingBgpConnection model.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerConnectionOptionsIncomingBgpConnection) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"allow":       basetypes.BoolType{},
		"remote_port": basetypes.NumberType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerConnectionOptionsIncomingBgpConnection objects.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerConnectionOptionsIncomingBgpConnection) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerConnectionOptionsOutgoingBgpConnection model.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerConnectionOptionsOutgoingBgpConnection) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"allow":      basetypes.BoolType{},
		"local_port": basetypes.NumberType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerConnectionOptionsOutgoingBgpConnection objects.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerConnectionOptionsOutgoingBgpConnection) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerInherit model.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerInherit) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"no": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address_family": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ipv4": basetypes.StringType{},
						"ipv6": basetypes.StringType{},
					},
				},
				"filtering_profile": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ipv4": basetypes.StringType{},
						"ipv6": basetypes.StringType{},
					},
				},
			},
		},
		"yes": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerInherit objects.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerInherit) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerInheritNo model.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerInheritNo) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address_family": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ipv4": basetypes.StringType{},
				"ipv6": basetypes.StringType{},
			},
		},
		"filtering_profile": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ipv4": basetypes.StringType{},
				"ipv6": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerInheritNo objects.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerInheritNo) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerLocalAddress model.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerLocalAddress) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"interface": basetypes.StringType{},
		"ip":        basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerLocalAddress objects.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerLocalAddress) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerPeerAddress model.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerPeerAddress) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"fqdn": basetypes.StringType{},
		"ip":   basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerPeerAddress objects.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerPeerAddress) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerSubsequentAddressFamilyIdentifier model.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerSubsequentAddressFamilyIdentifier) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"multicast": basetypes.BoolType{},
		"unicast":   basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerSubsequentAddressFamilyIdentifier objects.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerPeerInnerSubsequentAddressFamilyIdentifier) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPeerGroupInnerType model.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerType) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ebgp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"export_nexthop":    basetypes.StringType{},
				"import_nexthop":    basetypes.StringType{},
				"remove_private_as": basetypes.BoolType{},
			},
		},
		"ebgp_confed": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"export_nexthop": basetypes.StringType{},
			},
		},
		"ibgp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"export_nexthop": basetypes.StringType{},
			},
		},
		"ibgp_confed": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"export_nexthop": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPeerGroupInnerType objects.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerType) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPeerGroupInnerTypeEbgp model.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerTypeEbgp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"export_nexthop":    basetypes.StringType{},
		"import_nexthop":    basetypes.StringType{},
		"remove_private_as": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPeerGroupInnerTypeEbgp objects.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerTypeEbgp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPeerGroupInnerTypeEbgpConfed model.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerTypeEbgpConfed) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"export_nexthop": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPeerGroupInnerTypeEbgpConfed objects.
func (o LogicalRoutersVrfInnerBgpPeerGroupInnerTypeEbgpConfed) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPolicy model.
func (o LogicalRoutersVrfInnerBgpPolicy) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"aggregation": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"advertise_filters": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable": basetypes.BoolType{},
								"match": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"exact": basetypes.BoolType{},
												"name":  basetypes.StringType{},
											},
										}},
										"afi": basetypes.StringType{},
										"as_path": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"regex": basetypes.StringType{},
											},
										},
										"community": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"regex": basetypes.StringType{},
											},
										},
										"extended_community": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"regex": basetypes.StringType{},
											},
										},
										"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
										"med":         basetypes.NumberType{},
										"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
										"route_table": basetypes.StringType{},
										"safi":        basetypes.StringType{},
									},
								},
								"name": basetypes.StringType{},
							},
						}},
						"aggregate_route_attributes": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"as_path": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"none": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"prepend": basetypes.NumberType{},
										"remove": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"remove_and_prepend": basetypes.NumberType{},
									},
								},
								"as_path_limit": basetypes.NumberType{},
								"community": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"append": basetypes.ListType{ElemType: basetypes.StringType{}},
										"none": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
										"remove_all": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"remove_regex": basetypes.StringType{},
									},
								},
								"extended_community": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"append": basetypes.ListType{ElemType: basetypes.StringType{}},
										"none": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
										"remove_all": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"remove_regex": basetypes.StringType{},
									},
								},
								"local_preference": basetypes.NumberType{},
								"med":              basetypes.NumberType{},
								"nexthop":          basetypes.StringType{},
								"origin":           basetypes.StringType{},
								"weight":           basetypes.NumberType{},
							},
						},
						"as_set":  basetypes.BoolType{},
						"enable":  basetypes.BoolType{},
						"name":    basetypes.StringType{},
						"prefix":  basetypes.StringType{},
						"summary": basetypes.BoolType{},
						"suppress_filters": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable": basetypes.BoolType{},
								"match": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"exact": basetypes.BoolType{},
												"name":  basetypes.StringType{},
											},
										}},
										"afi": basetypes.StringType{},
										"as_path": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"regex": basetypes.StringType{},
											},
										},
										"community": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"regex": basetypes.StringType{},
											},
										},
										"extended_community": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"regex": basetypes.StringType{},
											},
										},
										"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
										"med":         basetypes.NumberType{},
										"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
										"route_table": basetypes.StringType{},
										"safi":        basetypes.StringType{},
									},
								},
								"name": basetypes.StringType{},
							},
						}},
					},
				}},
			},
		},
		"conditional_advertisement": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"policy": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"advertise_filters": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable": basetypes.BoolType{},
								"match": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"exact": basetypes.BoolType{},
												"name":  basetypes.StringType{},
											},
										}},
										"afi": basetypes.StringType{},
										"as_path": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"regex": basetypes.StringType{},
											},
										},
										"community": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"regex": basetypes.StringType{},
											},
										},
										"extended_community": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"regex": basetypes.StringType{},
											},
										},
										"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
										"med":         basetypes.NumberType{},
										"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
										"route_table": basetypes.StringType{},
										"safi":        basetypes.StringType{},
									},
								},
								"name": basetypes.StringType{},
							},
						}},
						"enable": basetypes.BoolType{},
						"name":   basetypes.StringType{},
						"non_exist_filters": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable": basetypes.BoolType{},
								"match": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"exact": basetypes.BoolType{},
												"name":  basetypes.StringType{},
											},
										}},
										"afi": basetypes.StringType{},
										"as_path": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"regex": basetypes.StringType{},
											},
										},
										"community": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"regex": basetypes.StringType{},
											},
										},
										"extended_community": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"regex": basetypes.StringType{},
											},
										},
										"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
										"med":         basetypes.NumberType{},
										"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
										"route_table": basetypes.StringType{},
										"safi":        basetypes.StringType{},
									},
								},
								"name": basetypes.StringType{},
							},
						}},
						"used_by": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						}},
					},
				}},
			},
		},
		"export": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"rules": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"allow": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"update": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"as_path": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"none": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
														"prepend": basetypes.NumberType{},
														"remove": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
														"remove_and_prepend": basetypes.NumberType{},
													},
												},
												"as_path_limit": basetypes.NumberType{},
												"community": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"append": basetypes.ListType{ElemType: basetypes.StringType{}},
														"none": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
														"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
														"remove_all": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
														"remove_regex": basetypes.StringType{},
													},
												},
												"extended_community": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"append": basetypes.ListType{ElemType: basetypes.StringType{}},
														"none": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
														"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
														"remove_all": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
														"remove_regex": basetypes.StringType{},
													},
												},
												"local_preference": basetypes.NumberType{},
												"med":              basetypes.NumberType{},
												"nexthop":          basetypes.StringType{},
												"origin":           basetypes.StringType{},
											},
										},
									},
								},
								"deny": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
						"enable": basetypes.BoolType{},
						"match": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"exact": basetypes.BoolType{},
										"name":  basetypes.StringType{},
									},
								}},
								"afi": basetypes.StringType{},
								"as_path": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"regex": basetypes.StringType{},
									},
								},
								"community": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"regex": basetypes.StringType{},
									},
								},
								"extended_community": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"regex": basetypes.StringType{},
									},
								},
								"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
								"med":         basetypes.NumberType{},
								"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
								"route_table": basetypes.StringType{},
								"safi":        basetypes.StringType{},
							},
						},
						"name":    basetypes.StringType{},
						"used_by": basetypes.ListType{ElemType: basetypes.StringType{}},
					},
				}},
			},
		},
		"import": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"rules": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"allow": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"dampening": basetypes.StringType{},
										"update": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"as_path": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"none": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
														"prepend": basetypes.NumberType{},
														"remove": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
														"remove_and_prepend": basetypes.NumberType{},
													},
												},
												"as_path_limit": basetypes.NumberType{},
												"community": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"append": basetypes.ListType{ElemType: basetypes.StringType{}},
														"none": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
														"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
														"remove_all": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
														"remove_regex": basetypes.StringType{},
													},
												},
												"extended_community": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"append": basetypes.ListType{ElemType: basetypes.StringType{}},
														"none": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
														"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
														"remove_all": basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{},
														},
														"remove_regex": basetypes.StringType{},
													},
												},
												"local_preference": basetypes.NumberType{},
												"med":              basetypes.NumberType{},
												"nexthop":          basetypes.StringType{},
												"origin":           basetypes.StringType{},
												"weight":           basetypes.NumberType{},
											},
										},
									},
								},
								"deny": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
						"enable": basetypes.BoolType{},
						"match": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"exact": basetypes.BoolType{},
										"name":  basetypes.StringType{},
									},
								}},
								"afi": basetypes.StringType{},
								"as_path": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"regex": basetypes.StringType{},
									},
								},
								"community": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"regex": basetypes.StringType{},
									},
								},
								"extended_community": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"regex": basetypes.StringType{},
									},
								},
								"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
								"med":         basetypes.NumberType{},
								"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
								"route_table": basetypes.StringType{},
								"safi":        basetypes.StringType{},
							},
						},
						"name":    basetypes.StringType{},
						"used_by": basetypes.ListType{ElemType: basetypes.StringType{}},
					},
				}},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPolicy objects.
func (o LogicalRoutersVrfInnerBgpPolicy) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPolicyAggregation model.
func (o LogicalRoutersVrfInnerBgpPolicyAggregation) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"advertise_filters": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.BoolType{},
						"match": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"exact": basetypes.BoolType{},
										"name":  basetypes.StringType{},
									},
								}},
								"afi": basetypes.StringType{},
								"as_path": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"regex": basetypes.StringType{},
									},
								},
								"community": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"regex": basetypes.StringType{},
									},
								},
								"extended_community": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"regex": basetypes.StringType{},
									},
								},
								"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
								"med":         basetypes.NumberType{},
								"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
								"route_table": basetypes.StringType{},
								"safi":        basetypes.StringType{},
							},
						},
						"name": basetypes.StringType{},
					},
				}},
				"aggregate_route_attributes": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"as_path": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"none": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"prepend": basetypes.NumberType{},
								"remove": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"remove_and_prepend": basetypes.NumberType{},
							},
						},
						"as_path_limit": basetypes.NumberType{},
						"community": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"append": basetypes.ListType{ElemType: basetypes.StringType{}},
								"none": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
								"remove_all": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"remove_regex": basetypes.StringType{},
							},
						},
						"extended_community": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"append": basetypes.ListType{ElemType: basetypes.StringType{}},
								"none": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
								"remove_all": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"remove_regex": basetypes.StringType{},
							},
						},
						"local_preference": basetypes.NumberType{},
						"med":              basetypes.NumberType{},
						"nexthop":          basetypes.StringType{},
						"origin":           basetypes.StringType{},
						"weight":           basetypes.NumberType{},
					},
				},
				"as_set":  basetypes.BoolType{},
				"enable":  basetypes.BoolType{},
				"name":    basetypes.StringType{},
				"prefix":  basetypes.StringType{},
				"summary": basetypes.BoolType{},
				"suppress_filters": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.BoolType{},
						"match": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"exact": basetypes.BoolType{},
										"name":  basetypes.StringType{},
									},
								}},
								"afi": basetypes.StringType{},
								"as_path": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"regex": basetypes.StringType{},
									},
								},
								"community": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"regex": basetypes.StringType{},
									},
								},
								"extended_community": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"regex": basetypes.StringType{},
									},
								},
								"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
								"med":         basetypes.NumberType{},
								"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
								"route_table": basetypes.StringType{},
								"safi":        basetypes.StringType{},
							},
						},
						"name": basetypes.StringType{},
					},
				}},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPolicyAggregation objects.
func (o LogicalRoutersVrfInnerBgpPolicyAggregation) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPolicyAggregationAddressInner model.
func (o LogicalRoutersVrfInnerBgpPolicyAggregationAddressInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"advertise_filters": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable": basetypes.BoolType{},
				"match": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"exact": basetypes.BoolType{},
								"name":  basetypes.StringType{},
							},
						}},
						"afi": basetypes.StringType{},
						"as_path": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"regex": basetypes.StringType{},
							},
						},
						"community": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"regex": basetypes.StringType{},
							},
						},
						"extended_community": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"regex": basetypes.StringType{},
							},
						},
						"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
						"med":         basetypes.NumberType{},
						"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
						"route_table": basetypes.StringType{},
						"safi":        basetypes.StringType{},
					},
				},
				"name": basetypes.StringType{},
			},
		}},
		"aggregate_route_attributes": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"as_path": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"none": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"prepend": basetypes.NumberType{},
						"remove": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"remove_and_prepend": basetypes.NumberType{},
					},
				},
				"as_path_limit": basetypes.NumberType{},
				"community": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"append": basetypes.ListType{ElemType: basetypes.StringType{}},
						"none": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
						"remove_all": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"remove_regex": basetypes.StringType{},
					},
				},
				"extended_community": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"append": basetypes.ListType{ElemType: basetypes.StringType{}},
						"none": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
						"remove_all": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"remove_regex": basetypes.StringType{},
					},
				},
				"local_preference": basetypes.NumberType{},
				"med":              basetypes.NumberType{},
				"nexthop":          basetypes.StringType{},
				"origin":           basetypes.StringType{},
				"weight":           basetypes.NumberType{},
			},
		},
		"as_set":  basetypes.BoolType{},
		"enable":  basetypes.BoolType{},
		"name":    basetypes.StringType{},
		"prefix":  basetypes.StringType{},
		"summary": basetypes.BoolType{},
		"suppress_filters": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable": basetypes.BoolType{},
				"match": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"exact": basetypes.BoolType{},
								"name":  basetypes.StringType{},
							},
						}},
						"afi": basetypes.StringType{},
						"as_path": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"regex": basetypes.StringType{},
							},
						},
						"community": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"regex": basetypes.StringType{},
							},
						},
						"extended_community": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"regex": basetypes.StringType{},
							},
						},
						"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
						"med":         basetypes.NumberType{},
						"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
						"route_table": basetypes.StringType{},
						"safi":        basetypes.StringType{},
					},
				},
				"name": basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPolicyAggregationAddressInner objects.
func (o LogicalRoutersVrfInnerBgpPolicyAggregationAddressInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAdvertiseFiltersInner model.
func (o LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAdvertiseFiltersInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable": basetypes.BoolType{},
		"match": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"exact": basetypes.BoolType{},
						"name":  basetypes.StringType{},
					},
				}},
				"afi": basetypes.StringType{},
				"as_path": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"regex": basetypes.StringType{},
					},
				},
				"community": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"regex": basetypes.StringType{},
					},
				},
				"extended_community": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"regex": basetypes.StringType{},
					},
				},
				"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
				"med":         basetypes.NumberType{},
				"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
				"route_table": basetypes.StringType{},
				"safi":        basetypes.StringType{},
			},
		},
		"name": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAdvertiseFiltersInner objects.
func (o LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAdvertiseFiltersInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAdvertiseFiltersInnerMatch model.
func (o LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAdvertiseFiltersInnerMatch) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"exact": basetypes.BoolType{},
				"name":  basetypes.StringType{},
			},
		}},
		"afi": basetypes.StringType{},
		"as_path": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"regex": basetypes.StringType{},
			},
		},
		"community": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"regex": basetypes.StringType{},
			},
		},
		"extended_community": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"regex": basetypes.StringType{},
			},
		},
		"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
		"med":         basetypes.NumberType{},
		"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
		"route_table": basetypes.StringType{},
		"safi":        basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAdvertiseFiltersInnerMatch objects.
func (o LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAdvertiseFiltersInnerMatch) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAdvertiseFiltersInnerMatchAddressPrefixInner model.
func (o LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAdvertiseFiltersInnerMatchAddressPrefixInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"exact": basetypes.BoolType{},
		"name":  basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAdvertiseFiltersInnerMatchAddressPrefixInner objects.
func (o LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAdvertiseFiltersInnerMatchAddressPrefixInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAdvertiseFiltersInnerMatchAsPath model.
func (o LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAdvertiseFiltersInnerMatchAsPath) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"regex": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAdvertiseFiltersInnerMatchAsPath objects.
func (o LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAdvertiseFiltersInnerMatchAsPath) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAggregateRouteAttributes model.
func (o LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAggregateRouteAttributes) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"as_path": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"none": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"prepend": basetypes.NumberType{},
				"remove": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"remove_and_prepend": basetypes.NumberType{},
			},
		},
		"as_path_limit": basetypes.NumberType{},
		"community": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"append": basetypes.ListType{ElemType: basetypes.StringType{}},
				"none": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
				"remove_all": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"remove_regex": basetypes.StringType{},
			},
		},
		"extended_community": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"append": basetypes.ListType{ElemType: basetypes.StringType{}},
				"none": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
				"remove_all": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"remove_regex": basetypes.StringType{},
			},
		},
		"local_preference": basetypes.NumberType{},
		"med":              basetypes.NumberType{},
		"nexthop":          basetypes.StringType{},
		"origin":           basetypes.StringType{},
		"weight":           basetypes.NumberType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAggregateRouteAttributes objects.
func (o LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAggregateRouteAttributes) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAggregateRouteAttributesAsPath model.
func (o LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAggregateRouteAttributesAsPath) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"none": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"prepend": basetypes.NumberType{},
		"remove": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"remove_and_prepend": basetypes.NumberType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAggregateRouteAttributesAsPath objects.
func (o LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAggregateRouteAttributesAsPath) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAggregateRouteAttributesCommunity model.
func (o LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAggregateRouteAttributesCommunity) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"append": basetypes.ListType{ElemType: basetypes.StringType{}},
		"none": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
		"remove_all": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"remove_regex": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAggregateRouteAttributesCommunity objects.
func (o LogicalRoutersVrfInnerBgpPolicyAggregationAddressInnerAggregateRouteAttributesCommunity) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPolicyConditionalAdvertisement model.
func (o LogicalRoutersVrfInnerBgpPolicyConditionalAdvertisement) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"policy": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"advertise_filters": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.BoolType{},
						"match": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"exact": basetypes.BoolType{},
										"name":  basetypes.StringType{},
									},
								}},
								"afi": basetypes.StringType{},
								"as_path": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"regex": basetypes.StringType{},
									},
								},
								"community": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"regex": basetypes.StringType{},
									},
								},
								"extended_community": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"regex": basetypes.StringType{},
									},
								},
								"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
								"med":         basetypes.NumberType{},
								"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
								"route_table": basetypes.StringType{},
								"safi":        basetypes.StringType{},
							},
						},
						"name": basetypes.StringType{},
					},
				}},
				"enable": basetypes.BoolType{},
				"name":   basetypes.StringType{},
				"non_exist_filters": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.BoolType{},
						"match": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"exact": basetypes.BoolType{},
										"name":  basetypes.StringType{},
									},
								}},
								"afi": basetypes.StringType{},
								"as_path": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"regex": basetypes.StringType{},
									},
								},
								"community": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"regex": basetypes.StringType{},
									},
								},
								"extended_community": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"regex": basetypes.StringType{},
									},
								},
								"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
								"med":         basetypes.NumberType{},
								"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
								"route_table": basetypes.StringType{},
								"safi":        basetypes.StringType{},
							},
						},
						"name": basetypes.StringType{},
					},
				}},
				"used_by": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				}},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPolicyConditionalAdvertisement objects.
func (o LogicalRoutersVrfInnerBgpPolicyConditionalAdvertisement) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPolicyConditionalAdvertisementPolicyInner model.
func (o LogicalRoutersVrfInnerBgpPolicyConditionalAdvertisementPolicyInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"advertise_filters": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable": basetypes.BoolType{},
				"match": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"exact": basetypes.BoolType{},
								"name":  basetypes.StringType{},
							},
						}},
						"afi": basetypes.StringType{},
						"as_path": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"regex": basetypes.StringType{},
							},
						},
						"community": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"regex": basetypes.StringType{},
							},
						},
						"extended_community": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"regex": basetypes.StringType{},
							},
						},
						"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
						"med":         basetypes.NumberType{},
						"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
						"route_table": basetypes.StringType{},
						"safi":        basetypes.StringType{},
					},
				},
				"name": basetypes.StringType{},
			},
		}},
		"enable": basetypes.BoolType{},
		"name":   basetypes.StringType{},
		"non_exist_filters": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable": basetypes.BoolType{},
				"match": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"exact": basetypes.BoolType{},
								"name":  basetypes.StringType{},
							},
						}},
						"afi": basetypes.StringType{},
						"as_path": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"regex": basetypes.StringType{},
							},
						},
						"community": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"regex": basetypes.StringType{},
							},
						},
						"extended_community": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"regex": basetypes.StringType{},
							},
						},
						"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
						"med":         basetypes.NumberType{},
						"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
						"route_table": basetypes.StringType{},
						"safi":        basetypes.StringType{},
					},
				},
				"name": basetypes.StringType{},
			},
		}},
		"used_by": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		}},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPolicyConditionalAdvertisementPolicyInner objects.
func (o LogicalRoutersVrfInnerBgpPolicyConditionalAdvertisementPolicyInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPolicyExport model.
func (o LogicalRoutersVrfInnerBgpPolicyExport) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"rules": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"allow": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"update": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"as_path": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"none": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"prepend": basetypes.NumberType{},
												"remove": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"remove_and_prepend": basetypes.NumberType{},
											},
										},
										"as_path_limit": basetypes.NumberType{},
										"community": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"append": basetypes.ListType{ElemType: basetypes.StringType{}},
												"none": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
												"remove_all": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"remove_regex": basetypes.StringType{},
											},
										},
										"extended_community": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"append": basetypes.ListType{ElemType: basetypes.StringType{}},
												"none": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
												"remove_all": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"remove_regex": basetypes.StringType{},
											},
										},
										"local_preference": basetypes.NumberType{},
										"med":              basetypes.NumberType{},
										"nexthop":          basetypes.StringType{},
										"origin":           basetypes.StringType{},
									},
								},
							},
						},
						"deny": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
				"enable": basetypes.BoolType{},
				"match": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"exact": basetypes.BoolType{},
								"name":  basetypes.StringType{},
							},
						}},
						"afi": basetypes.StringType{},
						"as_path": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"regex": basetypes.StringType{},
							},
						},
						"community": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"regex": basetypes.StringType{},
							},
						},
						"extended_community": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"regex": basetypes.StringType{},
							},
						},
						"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
						"med":         basetypes.NumberType{},
						"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
						"route_table": basetypes.StringType{},
						"safi":        basetypes.StringType{},
					},
				},
				"name":    basetypes.StringType{},
				"used_by": basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPolicyExport objects.
func (o LogicalRoutersVrfInnerBgpPolicyExport) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPolicyExportRulesInner model.
func (o LogicalRoutersVrfInnerBgpPolicyExportRulesInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"allow": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"update": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"as_path": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"none": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"prepend": basetypes.NumberType{},
										"remove": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"remove_and_prepend": basetypes.NumberType{},
									},
								},
								"as_path_limit": basetypes.NumberType{},
								"community": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"append": basetypes.ListType{ElemType: basetypes.StringType{}},
										"none": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
										"remove_all": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"remove_regex": basetypes.StringType{},
									},
								},
								"extended_community": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"append": basetypes.ListType{ElemType: basetypes.StringType{}},
										"none": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
										"remove_all": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"remove_regex": basetypes.StringType{},
									},
								},
								"local_preference": basetypes.NumberType{},
								"med":              basetypes.NumberType{},
								"nexthop":          basetypes.StringType{},
								"origin":           basetypes.StringType{},
							},
						},
					},
				},
				"deny": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
		"enable": basetypes.BoolType{},
		"match": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"exact": basetypes.BoolType{},
						"name":  basetypes.StringType{},
					},
				}},
				"afi": basetypes.StringType{},
				"as_path": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"regex": basetypes.StringType{},
					},
				},
				"community": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"regex": basetypes.StringType{},
					},
				},
				"extended_community": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"regex": basetypes.StringType{},
					},
				},
				"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
				"med":         basetypes.NumberType{},
				"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
				"route_table": basetypes.StringType{},
				"safi":        basetypes.StringType{},
			},
		},
		"name":    basetypes.StringType{},
		"used_by": basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPolicyExportRulesInner objects.
func (o LogicalRoutersVrfInnerBgpPolicyExportRulesInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPolicyExportRulesInnerAction model.
func (o LogicalRoutersVrfInnerBgpPolicyExportRulesInnerAction) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"allow": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"update": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"as_path": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"none": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"prepend": basetypes.NumberType{},
								"remove": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"remove_and_prepend": basetypes.NumberType{},
							},
						},
						"as_path_limit": basetypes.NumberType{},
						"community": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"append": basetypes.ListType{ElemType: basetypes.StringType{}},
								"none": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
								"remove_all": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"remove_regex": basetypes.StringType{},
							},
						},
						"extended_community": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"append": basetypes.ListType{ElemType: basetypes.StringType{}},
								"none": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
								"remove_all": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"remove_regex": basetypes.StringType{},
							},
						},
						"local_preference": basetypes.NumberType{},
						"med":              basetypes.NumberType{},
						"nexthop":          basetypes.StringType{},
						"origin":           basetypes.StringType{},
					},
				},
			},
		},
		"deny": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPolicyExportRulesInnerAction objects.
func (o LogicalRoutersVrfInnerBgpPolicyExportRulesInnerAction) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPolicyExportRulesInnerActionAllow model.
func (o LogicalRoutersVrfInnerBgpPolicyExportRulesInnerActionAllow) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"update": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"as_path": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"none": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"prepend": basetypes.NumberType{},
						"remove": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"remove_and_prepend": basetypes.NumberType{},
					},
				},
				"as_path_limit": basetypes.NumberType{},
				"community": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"append": basetypes.ListType{ElemType: basetypes.StringType{}},
						"none": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
						"remove_all": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"remove_regex": basetypes.StringType{},
					},
				},
				"extended_community": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"append": basetypes.ListType{ElemType: basetypes.StringType{}},
						"none": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
						"remove_all": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"remove_regex": basetypes.StringType{},
					},
				},
				"local_preference": basetypes.NumberType{},
				"med":              basetypes.NumberType{},
				"nexthop":          basetypes.StringType{},
				"origin":           basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPolicyExportRulesInnerActionAllow objects.
func (o LogicalRoutersVrfInnerBgpPolicyExportRulesInnerActionAllow) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPolicyExportRulesInnerActionAllowUpdate model.
func (o LogicalRoutersVrfInnerBgpPolicyExportRulesInnerActionAllowUpdate) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"as_path": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"none": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"prepend": basetypes.NumberType{},
				"remove": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"remove_and_prepend": basetypes.NumberType{},
			},
		},
		"as_path_limit": basetypes.NumberType{},
		"community": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"append": basetypes.ListType{ElemType: basetypes.StringType{}},
				"none": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
				"remove_all": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"remove_regex": basetypes.StringType{},
			},
		},
		"extended_community": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"append": basetypes.ListType{ElemType: basetypes.StringType{}},
				"none": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
				"remove_all": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"remove_regex": basetypes.StringType{},
			},
		},
		"local_preference": basetypes.NumberType{},
		"med":              basetypes.NumberType{},
		"nexthop":          basetypes.StringType{},
		"origin":           basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPolicyExportRulesInnerActionAllowUpdate objects.
func (o LogicalRoutersVrfInnerBgpPolicyExportRulesInnerActionAllowUpdate) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPolicyExportRulesInnerMatch model.
func (o LogicalRoutersVrfInnerBgpPolicyExportRulesInnerMatch) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"exact": basetypes.BoolType{},
				"name":  basetypes.StringType{},
			},
		}},
		"afi": basetypes.StringType{},
		"as_path": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"regex": basetypes.StringType{},
			},
		},
		"community": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"regex": basetypes.StringType{},
			},
		},
		"extended_community": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"regex": basetypes.StringType{},
			},
		},
		"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
		"med":         basetypes.NumberType{},
		"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
		"route_table": basetypes.StringType{},
		"safi":        basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPolicyExportRulesInnerMatch objects.
func (o LogicalRoutersVrfInnerBgpPolicyExportRulesInnerMatch) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPolicyExportRulesInnerMatchAddressPrefixInner model.
func (o LogicalRoutersVrfInnerBgpPolicyExportRulesInnerMatchAddressPrefixInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"exact": basetypes.BoolType{},
		"name":  basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPolicyExportRulesInnerMatchAddressPrefixInner objects.
func (o LogicalRoutersVrfInnerBgpPolicyExportRulesInnerMatchAddressPrefixInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPolicyImport model.
func (o LogicalRoutersVrfInnerBgpPolicyImport) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"rules": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"allow": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"dampening": basetypes.StringType{},
								"update": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"as_path": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"none": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"prepend": basetypes.NumberType{},
												"remove": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"remove_and_prepend": basetypes.NumberType{},
											},
										},
										"as_path_limit": basetypes.NumberType{},
										"community": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"append": basetypes.ListType{ElemType: basetypes.StringType{}},
												"none": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
												"remove_all": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"remove_regex": basetypes.StringType{},
											},
										},
										"extended_community": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"append": basetypes.ListType{ElemType: basetypes.StringType{}},
												"none": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
												"remove_all": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"remove_regex": basetypes.StringType{},
											},
										},
										"local_preference": basetypes.NumberType{},
										"med":              basetypes.NumberType{},
										"nexthop":          basetypes.StringType{},
										"origin":           basetypes.StringType{},
										"weight":           basetypes.NumberType{},
									},
								},
							},
						},
						"deny": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
				"enable": basetypes.BoolType{},
				"match": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"exact": basetypes.BoolType{},
								"name":  basetypes.StringType{},
							},
						}},
						"afi": basetypes.StringType{},
						"as_path": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"regex": basetypes.StringType{},
							},
						},
						"community": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"regex": basetypes.StringType{},
							},
						},
						"extended_community": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"regex": basetypes.StringType{},
							},
						},
						"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
						"med":         basetypes.NumberType{},
						"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
						"route_table": basetypes.StringType{},
						"safi":        basetypes.StringType{},
					},
				},
				"name":    basetypes.StringType{},
				"used_by": basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPolicyImport objects.
func (o LogicalRoutersVrfInnerBgpPolicyImport) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPolicyImportRulesInner model.
func (o LogicalRoutersVrfInnerBgpPolicyImportRulesInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"allow": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"dampening": basetypes.StringType{},
						"update": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"as_path": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"none": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"prepend": basetypes.NumberType{},
										"remove": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"remove_and_prepend": basetypes.NumberType{},
									},
								},
								"as_path_limit": basetypes.NumberType{},
								"community": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"append": basetypes.ListType{ElemType: basetypes.StringType{}},
										"none": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
										"remove_all": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"remove_regex": basetypes.StringType{},
									},
								},
								"extended_community": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"append": basetypes.ListType{ElemType: basetypes.StringType{}},
										"none": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
										"remove_all": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"remove_regex": basetypes.StringType{},
									},
								},
								"local_preference": basetypes.NumberType{},
								"med":              basetypes.NumberType{},
								"nexthop":          basetypes.StringType{},
								"origin":           basetypes.StringType{},
								"weight":           basetypes.NumberType{},
							},
						},
					},
				},
				"deny": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
		"enable": basetypes.BoolType{},
		"match": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address_prefix": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"exact": basetypes.BoolType{},
						"name":  basetypes.StringType{},
					},
				}},
				"afi": basetypes.StringType{},
				"as_path": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"regex": basetypes.StringType{},
					},
				},
				"community": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"regex": basetypes.StringType{},
					},
				},
				"extended_community": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"regex": basetypes.StringType{},
					},
				},
				"from_peer":   basetypes.ListType{ElemType: basetypes.StringType{}},
				"med":         basetypes.NumberType{},
				"nexthop":     basetypes.ListType{ElemType: basetypes.StringType{}},
				"route_table": basetypes.StringType{},
				"safi":        basetypes.StringType{},
			},
		},
		"name":    basetypes.StringType{},
		"used_by": basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPolicyImportRulesInner objects.
func (o LogicalRoutersVrfInnerBgpPolicyImportRulesInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPolicyImportRulesInnerAction model.
func (o LogicalRoutersVrfInnerBgpPolicyImportRulesInnerAction) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"allow": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"dampening": basetypes.StringType{},
				"update": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"as_path": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"none": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"prepend": basetypes.NumberType{},
								"remove": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"remove_and_prepend": basetypes.NumberType{},
							},
						},
						"as_path_limit": basetypes.NumberType{},
						"community": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"append": basetypes.ListType{ElemType: basetypes.StringType{}},
								"none": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
								"remove_all": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"remove_regex": basetypes.StringType{},
							},
						},
						"extended_community": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"append": basetypes.ListType{ElemType: basetypes.StringType{}},
								"none": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
								"remove_all": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"remove_regex": basetypes.StringType{},
							},
						},
						"local_preference": basetypes.NumberType{},
						"med":              basetypes.NumberType{},
						"nexthop":          basetypes.StringType{},
						"origin":           basetypes.StringType{},
						"weight":           basetypes.NumberType{},
					},
				},
			},
		},
		"deny": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPolicyImportRulesInnerAction objects.
func (o LogicalRoutersVrfInnerBgpPolicyImportRulesInnerAction) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpPolicyImportRulesInnerActionAllow model.
func (o LogicalRoutersVrfInnerBgpPolicyImportRulesInnerActionAllow) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"dampening": basetypes.StringType{},
		"update": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"as_path": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"none": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"prepend": basetypes.NumberType{},
						"remove": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"remove_and_prepend": basetypes.NumberType{},
					},
				},
				"as_path_limit": basetypes.NumberType{},
				"community": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"append": basetypes.ListType{ElemType: basetypes.StringType{}},
						"none": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
						"remove_all": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"remove_regex": basetypes.StringType{},
					},
				},
				"extended_community": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"append": basetypes.ListType{ElemType: basetypes.StringType{}},
						"none": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"overwrite": basetypes.ListType{ElemType: basetypes.StringType{}},
						"remove_all": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"remove_regex": basetypes.StringType{},
					},
				},
				"local_preference": basetypes.NumberType{},
				"med":              basetypes.NumberType{},
				"nexthop":          basetypes.StringType{},
				"origin":           basetypes.StringType{},
				"weight":           basetypes.NumberType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpPolicyImportRulesInnerActionAllow objects.
func (o LogicalRoutersVrfInnerBgpPolicyImportRulesInnerActionAllow) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpRedistRulesInner model.
func (o LogicalRoutersVrfInnerBgpRedistRulesInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address_family_identifier": basetypes.StringType{},
		"enable":                    basetypes.BoolType{},
		"metric":                    basetypes.NumberType{},
		"name":                      basetypes.StringType{},
		"route_table":               basetypes.StringType{},
		"set_as_path_limit":         basetypes.NumberType{},
		"set_community":             basetypes.ListType{ElemType: basetypes.StringType{}},
		"set_extended_community":    basetypes.ListType{ElemType: basetypes.StringType{}},
		"set_local_preference":      basetypes.NumberType{},
		"set_med":                   basetypes.NumberType{},
		"set_origin":                basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpRedistRulesInner objects.
func (o LogicalRoutersVrfInnerBgpRedistRulesInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpRedistributionProfile model.
func (o LogicalRoutersVrfInnerBgpRedistributionProfile) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ipv4": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"unicast": basetypes.StringType{},
			},
		},
		"ipv6": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"unicast": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpRedistributionProfile objects.
func (o LogicalRoutersVrfInnerBgpRedistributionProfile) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerBgpRedistributionProfileIpv4 model.
func (o LogicalRoutersVrfInnerBgpRedistributionProfileIpv4) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"unicast": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerBgpRedistributionProfileIpv4 objects.
func (o LogicalRoutersVrfInnerBgpRedistributionProfileIpv4) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerEcmp model.
func (o LogicalRoutersVrfInnerEcmp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"algorithm": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"balanced_round_robin": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"ip_hash": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"hash_seed": basetypes.NumberType{},
						"src_only":  basetypes.BoolType{},
						"use_port":  basetypes.BoolType{},
					},
				},
				"ip_modulo": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"weighted_round_robin": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":   basetypes.StringType{},
								"weight": basetypes.NumberType{},
							},
						}},
					},
				},
			},
		},
		"enable":             basetypes.BoolType{},
		"max_path":           basetypes.NumberType{},
		"strict_source_path": basetypes.BoolType{},
		"symmetric_return":   basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerEcmp objects.
func (o LogicalRoutersVrfInnerEcmp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerEcmpAlgorithm model.
func (o LogicalRoutersVrfInnerEcmpAlgorithm) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"balanced_round_robin": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"ip_hash": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"hash_seed": basetypes.NumberType{},
				"src_only":  basetypes.BoolType{},
				"use_port":  basetypes.BoolType{},
			},
		},
		"ip_modulo": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"weighted_round_robin": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":   basetypes.StringType{},
						"weight": basetypes.NumberType{},
					},
				}},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerEcmpAlgorithm objects.
func (o LogicalRoutersVrfInnerEcmpAlgorithm) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerEcmpAlgorithmIpHash model.
func (o LogicalRoutersVrfInnerEcmpAlgorithmIpHash) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"hash_seed": basetypes.NumberType{},
		"src_only":  basetypes.BoolType{},
		"use_port":  basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerEcmpAlgorithmIpHash objects.
func (o LogicalRoutersVrfInnerEcmpAlgorithmIpHash) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerEcmpAlgorithmWeightedRoundRobin model.
func (o LogicalRoutersVrfInnerEcmpAlgorithmWeightedRoundRobin) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":   basetypes.StringType{},
				"weight": basetypes.NumberType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerEcmpAlgorithmWeightedRoundRobin objects.
func (o LogicalRoutersVrfInnerEcmpAlgorithmWeightedRoundRobin) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerEcmpAlgorithmWeightedRoundRobinInterfaceInner model.
func (o LogicalRoutersVrfInnerEcmpAlgorithmWeightedRoundRobinInterfaceInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":   basetypes.StringType{},
		"weight": basetypes.NumberType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerEcmpAlgorithmWeightedRoundRobinInterfaceInner objects.
func (o LogicalRoutersVrfInnerEcmpAlgorithmWeightedRoundRobinInterfaceInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticast model.
func (o LogicalRoutersVrfInnerMulticast) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable":    basetypes.BoolType{},
		"enable_v6": basetypes.BoolType{},
		"igmp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"dynamic": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"group_filter":          basetypes.StringType{},
								"max_groups":            basetypes.StringType{},
								"max_sources":           basetypes.StringType{},
								"name":                  basetypes.StringType{},
								"query_profile":         basetypes.StringType{},
								"robustness":            basetypes.StringType{},
								"router_alert_policing": basetypes.BoolType{},
								"version":               basetypes.StringType{},
							},
						}},
					},
				},
				"enable": basetypes.BoolType{},
				"static": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"group_address":  basetypes.StringType{},
						"interface":      basetypes.StringType{},
						"name":           basetypes.StringType{},
						"source_address": basetypes.StringType{},
					},
				}},
			},
		},
		"interface_group": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"description": basetypes.StringType{},
				"group_permission": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"any_source_multicast": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"group_address": basetypes.StringType{},
								"included":      basetypes.BoolType{},
								"name":          basetypes.StringType{},
							},
						}},
						"source_specific_multicast": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"group_address":  basetypes.StringType{},
								"included":       basetypes.BoolType{},
								"name":           basetypes.StringType{},
								"source_address": basetypes.StringType{},
							},
						}},
					},
				},
				"igmp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable":                     basetypes.BoolType{},
						"immediate_leave":            basetypes.BoolType{},
						"last_member_query_interval": basetypes.NumberType{},
						"max_groups":                 basetypes.StringType{},
						"max_query_response_time":    basetypes.NumberType{},
						"max_sources":                basetypes.StringType{},
						"mode":                       basetypes.StringType{},
						"query_interval":             basetypes.NumberType{},
						"robustness":                 basetypes.StringType{},
						"router_alert_policing":      basetypes.BoolType{},
						"version":                    basetypes.StringType{},
					},
				},
				"interface": basetypes.ListType{ElemType: basetypes.StringType{}},
				"name":      basetypes.StringType{},
				"pim": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"allowed_neighbors": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name": basetypes.StringType{},
							},
						}},
						"assert_interval":     basetypes.NumberType{},
						"bsr_border":          basetypes.BoolType{},
						"dr_priority":         basetypes.NumberType{},
						"enable":              basetypes.BoolType{},
						"hello_interval":      basetypes.NumberType{},
						"join_prune_interval": basetypes.NumberType{},
					},
				},
			},
		}},
		"mode": basetypes.StringType{},
		"msdp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable":                basetypes.BoolType{},
				"global_authentication": basetypes.StringType{},
				"global_timer":          basetypes.StringType{},
				"originator_id": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"interface": basetypes.StringType{},
						"ip":        basetypes.StringType{},
					},
				},
				"peer": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"authentication":    basetypes.StringType{},
						"enable":            basetypes.BoolType{},
						"inbound_sa_filter": basetypes.StringType{},
						"local_address": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"interface": basetypes.StringType{},
								"ip":        basetypes.StringType{},
							},
						},
						"max_sa":             basetypes.NumberType{},
						"name":               basetypes.StringType{},
						"outbound_sa_filter": basetypes.StringType{},
						"peer_address": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"fqdn": basetypes.StringType{},
								"ip":   basetypes.StringType{},
							},
						},
						"peer_as": basetypes.StringType{},
					},
				}},
			},
		},
		"pim": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable":           basetypes.BoolType{},
				"group_permission": basetypes.StringType{},
				"if_timer_global":  basetypes.StringType{},
				"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"description":     basetypes.StringType{},
						"dr_priority":     basetypes.NumberType{},
						"if_timer":        basetypes.StringType{},
						"name":            basetypes.StringType{},
						"neighbor_filter": basetypes.StringType{},
						"send_bsm":        basetypes.BoolType{},
					},
				}},
				"route_ageout_time": basetypes.NumberType{},
				"rp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"external_rp": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"group_list": basetypes.StringType{},
								"name":       basetypes.StringType{},
								"override":   basetypes.BoolType{},
							},
						}},
						"local_rp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"candidate_rp": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"address":                basetypes.StringType{},
										"advertisement_interval": basetypes.NumberType{},
										"group_list":             basetypes.StringType{},
										"interface":              basetypes.StringType{},
										"priority":               basetypes.NumberType{},
									},
								},
								"static_rp": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"address":    basetypes.StringType{},
										"group_list": basetypes.StringType{},
										"interface":  basetypes.StringType{},
										"override":   basetypes.BoolType{},
									},
								},
							},
						},
					},
				},
				"rpf_lookup_mode": basetypes.StringType{},
				"spt_threshold": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":      basetypes.StringType{},
						"threshold": basetypes.StringType{},
					},
				}},
				"ssm_address_space": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"group_list": basetypes.StringType{},
					},
				},
			},
		},
		"route_ageout_time": basetypes.NumberType{},
		"rp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"external_rp": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"group_addresses": basetypes.ListType{ElemType: basetypes.StringType{}},
						"name":            basetypes.StringType{},
						"override":        basetypes.BoolType{},
					},
				}},
				"local_rp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"candidate_rp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address":                basetypes.StringType{},
								"advertisement_interval": basetypes.NumberType{},
								"group_addresses":        basetypes.ListType{ElemType: basetypes.StringType{}},
								"interface":              basetypes.StringType{},
								"priority":               basetypes.NumberType{},
							},
						},
						"static_rp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address":         basetypes.StringType{},
								"group_addresses": basetypes.ListType{ElemType: basetypes.StringType{}},
								"interface":       basetypes.StringType{},
								"override":        basetypes.BoolType{},
							},
						},
					},
				},
			},
		},
		"spt_threshold": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":      basetypes.StringType{},
				"threshold": basetypes.StringType{},
			},
		}},
		"ssm_address_space": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"group_address": basetypes.StringType{},
				"included":      basetypes.BoolType{},
				"name":          basetypes.StringType{},
			},
		}},
		"static_route": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"destination": basetypes.StringType{},
				"interface":   basetypes.StringType{},
				"name":        basetypes.StringType{},
				"nexthop": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ip_address": basetypes.StringType{},
					},
				},
				"preference": basetypes.NumberType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticast objects.
func (o LogicalRoutersVrfInnerMulticast) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastIgmp model.
func (o LogicalRoutersVrfInnerMulticastIgmp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"dynamic": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"group_filter":          basetypes.StringType{},
						"max_groups":            basetypes.StringType{},
						"max_sources":           basetypes.StringType{},
						"name":                  basetypes.StringType{},
						"query_profile":         basetypes.StringType{},
						"robustness":            basetypes.StringType{},
						"router_alert_policing": basetypes.BoolType{},
						"version":               basetypes.StringType{},
					},
				}},
			},
		},
		"enable": basetypes.BoolType{},
		"static": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"group_address":  basetypes.StringType{},
				"interface":      basetypes.StringType{},
				"name":           basetypes.StringType{},
				"source_address": basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastIgmp objects.
func (o LogicalRoutersVrfInnerMulticastIgmp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastIgmpDynamic model.
func (o LogicalRoutersVrfInnerMulticastIgmpDynamic) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"group_filter":          basetypes.StringType{},
				"max_groups":            basetypes.StringType{},
				"max_sources":           basetypes.StringType{},
				"name":                  basetypes.StringType{},
				"query_profile":         basetypes.StringType{},
				"robustness":            basetypes.StringType{},
				"router_alert_policing": basetypes.BoolType{},
				"version":               basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastIgmpDynamic objects.
func (o LogicalRoutersVrfInnerMulticastIgmpDynamic) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastIgmpDynamicInterfaceInner model.
func (o LogicalRoutersVrfInnerMulticastIgmpDynamicInterfaceInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"group_filter":          basetypes.StringType{},
		"max_groups":            basetypes.StringType{},
		"max_sources":           basetypes.StringType{},
		"name":                  basetypes.StringType{},
		"query_profile":         basetypes.StringType{},
		"robustness":            basetypes.StringType{},
		"router_alert_policing": basetypes.BoolType{},
		"version":               basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastIgmpDynamicInterfaceInner objects.
func (o LogicalRoutersVrfInnerMulticastIgmpDynamicInterfaceInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastIgmpStaticInner model.
func (o LogicalRoutersVrfInnerMulticastIgmpStaticInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"group_address":  basetypes.StringType{},
		"interface":      basetypes.StringType{},
		"name":           basetypes.StringType{},
		"source_address": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastIgmpStaticInner objects.
func (o LogicalRoutersVrfInnerMulticastIgmpStaticInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastInterfaceGroupInner model.
func (o LogicalRoutersVrfInnerMulticastInterfaceGroupInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"description": basetypes.StringType{},
		"group_permission": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"any_source_multicast": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"group_address": basetypes.StringType{},
						"included":      basetypes.BoolType{},
						"name":          basetypes.StringType{},
					},
				}},
				"source_specific_multicast": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"group_address":  basetypes.StringType{},
						"included":       basetypes.BoolType{},
						"name":           basetypes.StringType{},
						"source_address": basetypes.StringType{},
					},
				}},
			},
		},
		"igmp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable":                     basetypes.BoolType{},
				"immediate_leave":            basetypes.BoolType{},
				"last_member_query_interval": basetypes.NumberType{},
				"max_groups":                 basetypes.StringType{},
				"max_query_response_time":    basetypes.NumberType{},
				"max_sources":                basetypes.StringType{},
				"mode":                       basetypes.StringType{},
				"query_interval":             basetypes.NumberType{},
				"robustness":                 basetypes.StringType{},
				"router_alert_policing":      basetypes.BoolType{},
				"version":                    basetypes.StringType{},
			},
		},
		"interface": basetypes.ListType{ElemType: basetypes.StringType{}},
		"name":      basetypes.StringType{},
		"pim": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"allowed_neighbors": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name": basetypes.StringType{},
					},
				}},
				"assert_interval":     basetypes.NumberType{},
				"bsr_border":          basetypes.BoolType{},
				"dr_priority":         basetypes.NumberType{},
				"enable":              basetypes.BoolType{},
				"hello_interval":      basetypes.NumberType{},
				"join_prune_interval": basetypes.NumberType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastInterfaceGroupInner objects.
func (o LogicalRoutersVrfInnerMulticastInterfaceGroupInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastInterfaceGroupInnerGroupPermission model.
func (o LogicalRoutersVrfInnerMulticastInterfaceGroupInnerGroupPermission) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"any_source_multicast": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"group_address": basetypes.StringType{},
				"included":      basetypes.BoolType{},
				"name":          basetypes.StringType{},
			},
		}},
		"source_specific_multicast": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"group_address":  basetypes.StringType{},
				"included":       basetypes.BoolType{},
				"name":           basetypes.StringType{},
				"source_address": basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastInterfaceGroupInnerGroupPermission objects.
func (o LogicalRoutersVrfInnerMulticastInterfaceGroupInnerGroupPermission) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastInterfaceGroupInnerGroupPermissionAnySourceMulticastInner model.
func (o LogicalRoutersVrfInnerMulticastInterfaceGroupInnerGroupPermissionAnySourceMulticastInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"group_address": basetypes.StringType{},
		"included":      basetypes.BoolType{},
		"name":          basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastInterfaceGroupInnerGroupPermissionAnySourceMulticastInner objects.
func (o LogicalRoutersVrfInnerMulticastInterfaceGroupInnerGroupPermissionAnySourceMulticastInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastInterfaceGroupInnerGroupPermissionSourceSpecificMulticastInner model.
func (o LogicalRoutersVrfInnerMulticastInterfaceGroupInnerGroupPermissionSourceSpecificMulticastInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"group_address":  basetypes.StringType{},
		"included":       basetypes.BoolType{},
		"name":           basetypes.StringType{},
		"source_address": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastInterfaceGroupInnerGroupPermissionSourceSpecificMulticastInner objects.
func (o LogicalRoutersVrfInnerMulticastInterfaceGroupInnerGroupPermissionSourceSpecificMulticastInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastInterfaceGroupInnerIgmp model.
func (o LogicalRoutersVrfInnerMulticastInterfaceGroupInnerIgmp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable":                     basetypes.BoolType{},
		"immediate_leave":            basetypes.BoolType{},
		"last_member_query_interval": basetypes.NumberType{},
		"max_groups":                 basetypes.StringType{},
		"max_query_response_time":    basetypes.NumberType{},
		"max_sources":                basetypes.StringType{},
		"mode":                       basetypes.StringType{},
		"query_interval":             basetypes.NumberType{},
		"robustness":                 basetypes.StringType{},
		"router_alert_policing":      basetypes.BoolType{},
		"version":                    basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastInterfaceGroupInnerIgmp objects.
func (o LogicalRoutersVrfInnerMulticastInterfaceGroupInnerIgmp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastInterfaceGroupInnerPim model.
func (o LogicalRoutersVrfInnerMulticastInterfaceGroupInnerPim) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"allowed_neighbors": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name": basetypes.StringType{},
			},
		}},
		"assert_interval":     basetypes.NumberType{},
		"bsr_border":          basetypes.BoolType{},
		"dr_priority":         basetypes.NumberType{},
		"enable":              basetypes.BoolType{},
		"hello_interval":      basetypes.NumberType{},
		"join_prune_interval": basetypes.NumberType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastInterfaceGroupInnerPim objects.
func (o LogicalRoutersVrfInnerMulticastInterfaceGroupInnerPim) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastInterfaceGroupInnerPimAllowedNeighborsInner model.
func (o LogicalRoutersVrfInnerMulticastInterfaceGroupInnerPimAllowedNeighborsInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastInterfaceGroupInnerPimAllowedNeighborsInner objects.
func (o LogicalRoutersVrfInnerMulticastInterfaceGroupInnerPimAllowedNeighborsInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastMsdp model.
func (o LogicalRoutersVrfInnerMulticastMsdp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable":                basetypes.BoolType{},
		"global_authentication": basetypes.StringType{},
		"global_timer":          basetypes.StringType{},
		"originator_id": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"interface": basetypes.StringType{},
				"ip":        basetypes.StringType{},
			},
		},
		"peer": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"authentication":    basetypes.StringType{},
				"enable":            basetypes.BoolType{},
				"inbound_sa_filter": basetypes.StringType{},
				"local_address": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"interface": basetypes.StringType{},
						"ip":        basetypes.StringType{},
					},
				},
				"max_sa":             basetypes.NumberType{},
				"name":               basetypes.StringType{},
				"outbound_sa_filter": basetypes.StringType{},
				"peer_address": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"fqdn": basetypes.StringType{},
						"ip":   basetypes.StringType{},
					},
				},
				"peer_as": basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastMsdp objects.
func (o LogicalRoutersVrfInnerMulticastMsdp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastMsdpPeerInner model.
func (o LogicalRoutersVrfInnerMulticastMsdpPeerInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"authentication":    basetypes.StringType{},
		"enable":            basetypes.BoolType{},
		"inbound_sa_filter": basetypes.StringType{},
		"local_address": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"interface": basetypes.StringType{},
				"ip":        basetypes.StringType{},
			},
		},
		"max_sa":             basetypes.NumberType{},
		"name":               basetypes.StringType{},
		"outbound_sa_filter": basetypes.StringType{},
		"peer_address": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"fqdn": basetypes.StringType{},
				"ip":   basetypes.StringType{},
			},
		},
		"peer_as": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastMsdpPeerInner objects.
func (o LogicalRoutersVrfInnerMulticastMsdpPeerInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastPim model.
func (o LogicalRoutersVrfInnerMulticastPim) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable":           basetypes.BoolType{},
		"group_permission": basetypes.StringType{},
		"if_timer_global":  basetypes.StringType{},
		"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"description":     basetypes.StringType{},
				"dr_priority":     basetypes.NumberType{},
				"if_timer":        basetypes.StringType{},
				"name":            basetypes.StringType{},
				"neighbor_filter": basetypes.StringType{},
				"send_bsm":        basetypes.BoolType{},
			},
		}},
		"route_ageout_time": basetypes.NumberType{},
		"rp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"external_rp": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"group_list": basetypes.StringType{},
						"name":       basetypes.StringType{},
						"override":   basetypes.BoolType{},
					},
				}},
				"local_rp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"candidate_rp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address":                basetypes.StringType{},
								"advertisement_interval": basetypes.NumberType{},
								"group_list":             basetypes.StringType{},
								"interface":              basetypes.StringType{},
								"priority":               basetypes.NumberType{},
							},
						},
						"static_rp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address":    basetypes.StringType{},
								"group_list": basetypes.StringType{},
								"interface":  basetypes.StringType{},
								"override":   basetypes.BoolType{},
							},
						},
					},
				},
			},
		},
		"rpf_lookup_mode": basetypes.StringType{},
		"spt_threshold": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":      basetypes.StringType{},
				"threshold": basetypes.StringType{},
			},
		}},
		"ssm_address_space": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"group_list": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastPim objects.
func (o LogicalRoutersVrfInnerMulticastPim) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastPimInterfaceInner model.
func (o LogicalRoutersVrfInnerMulticastPimInterfaceInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"description":     basetypes.StringType{},
		"dr_priority":     basetypes.NumberType{},
		"if_timer":        basetypes.StringType{},
		"name":            basetypes.StringType{},
		"neighbor_filter": basetypes.StringType{},
		"send_bsm":        basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastPimInterfaceInner objects.
func (o LogicalRoutersVrfInnerMulticastPimInterfaceInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastPimRp model.
func (o LogicalRoutersVrfInnerMulticastPimRp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"external_rp": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"group_list": basetypes.StringType{},
				"name":       basetypes.StringType{},
				"override":   basetypes.BoolType{},
			},
		}},
		"local_rp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"candidate_rp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address":                basetypes.StringType{},
						"advertisement_interval": basetypes.NumberType{},
						"group_list":             basetypes.StringType{},
						"interface":              basetypes.StringType{},
						"priority":               basetypes.NumberType{},
					},
				},
				"static_rp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address":    basetypes.StringType{},
						"group_list": basetypes.StringType{},
						"interface":  basetypes.StringType{},
						"override":   basetypes.BoolType{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastPimRp objects.
func (o LogicalRoutersVrfInnerMulticastPimRp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastPimRpExternalRpInner model.
func (o LogicalRoutersVrfInnerMulticastPimRpExternalRpInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"group_list": basetypes.StringType{},
		"name":       basetypes.StringType{},
		"override":   basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastPimRpExternalRpInner objects.
func (o LogicalRoutersVrfInnerMulticastPimRpExternalRpInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastPimRpLocalRp model.
func (o LogicalRoutersVrfInnerMulticastPimRpLocalRp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"candidate_rp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address":                basetypes.StringType{},
				"advertisement_interval": basetypes.NumberType{},
				"group_list":             basetypes.StringType{},
				"interface":              basetypes.StringType{},
				"priority":               basetypes.NumberType{},
			},
		},
		"static_rp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address":    basetypes.StringType{},
				"group_list": basetypes.StringType{},
				"interface":  basetypes.StringType{},
				"override":   basetypes.BoolType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastPimRpLocalRp objects.
func (o LogicalRoutersVrfInnerMulticastPimRpLocalRp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastPimRpLocalRpCandidateRp model.
func (o LogicalRoutersVrfInnerMulticastPimRpLocalRpCandidateRp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address":                basetypes.StringType{},
		"advertisement_interval": basetypes.NumberType{},
		"group_list":             basetypes.StringType{},
		"interface":              basetypes.StringType{},
		"priority":               basetypes.NumberType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastPimRpLocalRpCandidateRp objects.
func (o LogicalRoutersVrfInnerMulticastPimRpLocalRpCandidateRp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastPimRpLocalRpStaticRp model.
func (o LogicalRoutersVrfInnerMulticastPimRpLocalRpStaticRp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address":    basetypes.StringType{},
		"group_list": basetypes.StringType{},
		"interface":  basetypes.StringType{},
		"override":   basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastPimRpLocalRpStaticRp objects.
func (o LogicalRoutersVrfInnerMulticastPimRpLocalRpStaticRp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastPimSptThresholdInner model.
func (o LogicalRoutersVrfInnerMulticastPimSptThresholdInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":      basetypes.StringType{},
		"threshold": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastPimSptThresholdInner objects.
func (o LogicalRoutersVrfInnerMulticastPimSptThresholdInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastPimSsmAddressSpace model.
func (o LogicalRoutersVrfInnerMulticastPimSsmAddressSpace) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"group_list": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastPimSsmAddressSpace objects.
func (o LogicalRoutersVrfInnerMulticastPimSsmAddressSpace) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastRp model.
func (o LogicalRoutersVrfInnerMulticastRp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"external_rp": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"group_addresses": basetypes.ListType{ElemType: basetypes.StringType{}},
				"name":            basetypes.StringType{},
				"override":        basetypes.BoolType{},
			},
		}},
		"local_rp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"candidate_rp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address":                basetypes.StringType{},
						"advertisement_interval": basetypes.NumberType{},
						"group_addresses":        basetypes.ListType{ElemType: basetypes.StringType{}},
						"interface":              basetypes.StringType{},
						"priority":               basetypes.NumberType{},
					},
				},
				"static_rp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address":         basetypes.StringType{},
						"group_addresses": basetypes.ListType{ElemType: basetypes.StringType{}},
						"interface":       basetypes.StringType{},
						"override":        basetypes.BoolType{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastRp objects.
func (o LogicalRoutersVrfInnerMulticastRp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastRpExternalRpInner model.
func (o LogicalRoutersVrfInnerMulticastRpExternalRpInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"group_addresses": basetypes.ListType{ElemType: basetypes.StringType{}},
		"name":            basetypes.StringType{},
		"override":        basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastRpExternalRpInner objects.
func (o LogicalRoutersVrfInnerMulticastRpExternalRpInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastRpLocalRp model.
func (o LogicalRoutersVrfInnerMulticastRpLocalRp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"candidate_rp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address":                basetypes.StringType{},
				"advertisement_interval": basetypes.NumberType{},
				"group_addresses":        basetypes.ListType{ElemType: basetypes.StringType{}},
				"interface":              basetypes.StringType{},
				"priority":               basetypes.NumberType{},
			},
		},
		"static_rp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address":         basetypes.StringType{},
				"group_addresses": basetypes.ListType{ElemType: basetypes.StringType{}},
				"interface":       basetypes.StringType{},
				"override":        basetypes.BoolType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastRpLocalRp objects.
func (o LogicalRoutersVrfInnerMulticastRpLocalRp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastRpLocalRpCandidateRp model.
func (o LogicalRoutersVrfInnerMulticastRpLocalRpCandidateRp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address":                basetypes.StringType{},
		"advertisement_interval": basetypes.NumberType{},
		"group_addresses":        basetypes.ListType{ElemType: basetypes.StringType{}},
		"interface":              basetypes.StringType{},
		"priority":               basetypes.NumberType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastRpLocalRpCandidateRp objects.
func (o LogicalRoutersVrfInnerMulticastRpLocalRpCandidateRp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastRpLocalRpStaticRp model.
func (o LogicalRoutersVrfInnerMulticastRpLocalRpStaticRp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address":         basetypes.StringType{},
		"group_addresses": basetypes.ListType{ElemType: basetypes.StringType{}},
		"interface":       basetypes.StringType{},
		"override":        basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastRpLocalRpStaticRp objects.
func (o LogicalRoutersVrfInnerMulticastRpLocalRpStaticRp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastStaticRouteInner model.
func (o LogicalRoutersVrfInnerMulticastStaticRouteInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"destination": basetypes.StringType{},
		"interface":   basetypes.StringType{},
		"name":        basetypes.StringType{},
		"nexthop": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ip_address": basetypes.StringType{},
			},
		},
		"preference": basetypes.NumberType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastStaticRouteInner objects.
func (o LogicalRoutersVrfInnerMulticastStaticRouteInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerMulticastStaticRouteInnerNexthop model.
func (o LogicalRoutersVrfInnerMulticastStaticRouteInnerNexthop) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ip_address": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerMulticastStaticRouteInnerNexthop objects.
func (o LogicalRoutersVrfInnerMulticastStaticRouteInnerNexthop) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspf model.
func (o LogicalRoutersVrfInnerOspf) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"allow_redist_default_route": basetypes.BoolType{},
		"area": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"authentication": basetypes.StringType{},
				"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"authentication": basetypes.StringType{},
						"bfd": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"profile": basetypes.StringType{},
							},
						},
						"enable": basetypes.BoolType{},
						"link_type": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"broadcast": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"p2mp": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"neighbor": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"name":     basetypes.StringType{},
												"priority": basetypes.NumberType{},
											},
										}},
									},
								},
								"p2p": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
						"metric":     basetypes.NumberType{},
						"mtu_ignore": basetypes.BoolType{},
						"name":       basetypes.StringType{},
						"passive":    basetypes.BoolType{},
						"priority":   basetypes.NumberType{},
						"timing":     basetypes.StringType{},
						"vr_timing": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"dead_counts":         basetypes.NumberType{},
								"gr_delay":            basetypes.NumberType{},
								"hello_interval":      basetypes.NumberType{},
								"retransmit_interval": basetypes.NumberType{},
								"transit_delay":       basetypes.NumberType{},
							},
						},
					},
				}},
				"name": basetypes.StringType{},
				"range": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"advertise":  basetypes.BoolType{},
						"name":       basetypes.StringType{},
						"substitute": basetypes.StringType{},
					},
				}},
				"type": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"normal": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"abr": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"export_list":          basetypes.StringType{},
										"import_list":          basetypes.StringType{},
										"inbound_filter_list":  basetypes.StringType{},
										"outbound_filter_list": basetypes.StringType{},
									},
								},
							},
						},
						"nssa": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"abr": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"export_list":         basetypes.StringType{},
										"import_list":         basetypes.StringType{},
										"inbound_filter_list": basetypes.StringType{},
										"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"advertise": basetypes.BoolType{},
												"name":      basetypes.StringType{},
												"route_tag": basetypes.NumberType{},
											},
										}},
										"outbound_filter_list": basetypes.StringType{},
									},
								},
								"accept_summary": basetypes.BoolType{},
								"default_information_originate": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"metric":      basetypes.NumberType{},
										"metric_type": basetypes.StringType{},
									},
								},
								"default_route": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"advertise": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"metric": basetypes.NumberType{},
												"type":   basetypes.StringType{},
											},
										},
										"disable": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								},
								"no_summary": basetypes.BoolType{},
								"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"advertise": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"name": basetypes.StringType{},
										"suppress": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								}},
							},
						},
						"stub": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"abr": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"export_list":          basetypes.StringType{},
										"import_list":          basetypes.StringType{},
										"inbound_filter_list":  basetypes.StringType{},
										"outbound_filter_list": basetypes.StringType{},
									},
								},
								"accept_summary": basetypes.BoolType{},
								"default_route": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"advertise": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"metric": basetypes.NumberType{},
											},
										},
										"disable": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								},
								"default_route_metric": basetypes.NumberType{},
								"no_summary":           basetypes.BoolType{},
							},
						},
					},
				},
				"virtual_link": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"authentication": basetypes.StringType{},
						"bfd": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"profile": basetypes.StringType{},
							},
						},
						"enable":          basetypes.BoolType{},
						"instance_id":     basetypes.NumberType{},
						"interface_id":    basetypes.NumberType{},
						"name":            basetypes.StringType{},
						"neighbor_id":     basetypes.StringType{},
						"passive":         basetypes.BoolType{},
						"timing":          basetypes.StringType{},
						"transit_area_id": basetypes.StringType{},
						"vr_timing": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"dead_counts":         basetypes.NumberType{},
								"hello_interval":      basetypes.NumberType{},
								"retransmit_interval": basetypes.NumberType{},
								"transit_delay":       basetypes.NumberType{},
							},
						},
					},
				}},
				"vr_range": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"advertise": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"name": basetypes.StringType{},
						"suppress": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				}},
			},
		}},
		"auth_profile": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"md5": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"key":       basetypes.StringType{},
						"name":      basetypes.NumberType{},
						"preferred": basetypes.BoolType{},
					},
				}},
				"name":     basetypes.StringType{},
				"password": basetypes.StringType{},
			},
		}},
		"enable": basetypes.BoolType{},
		"export_rules": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"metric":        basetypes.NumberType{},
				"name":          basetypes.StringType{},
				"new_path_type": basetypes.StringType{},
				"new_tag":       basetypes.StringType{},
			},
		}},
		"flood_prevention": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"hello": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable":     basetypes.BoolType{},
						"max_packet": basetypes.NumberType{},
					},
				},
				"lsa": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable":     basetypes.BoolType{},
						"max_packet": basetypes.NumberType{},
					},
				},
			},
		},
		"global_bfd": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"profile": basetypes.StringType{},
			},
		},
		"global_if_timer": basetypes.StringType{},
		"graceful_restart": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable":                    basetypes.BoolType{},
				"grace_period":              basetypes.NumberType{},
				"helper_enable":             basetypes.BoolType{},
				"max_neighbor_restart_time": basetypes.NumberType{},
				"strict__l_s_a_checking":    basetypes.BoolType{},
			},
		},
		"redistribution_profile": basetypes.StringType{},
		"reject_default_route":   basetypes.BoolType{},
		"rfc1583":                basetypes.BoolType{},
		"router_id":              basetypes.StringType{},
		"spf_timer":              basetypes.StringType{},
		"vr_timers": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"lsa_interval":          basetypes.NumberType{},
				"spf_calculation_delay": basetypes.NumberType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspf objects.
func (o LogicalRoutersVrfInnerOspf) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfAreaInner model.
func (o LogicalRoutersVrfInnerOspfAreaInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"authentication": basetypes.StringType{},
		"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"authentication": basetypes.StringType{},
				"bfd": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"profile": basetypes.StringType{},
					},
				},
				"enable": basetypes.BoolType{},
				"link_type": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"broadcast": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"p2mp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"neighbor": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"name":     basetypes.StringType{},
										"priority": basetypes.NumberType{},
									},
								}},
							},
						},
						"p2p": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
				"metric":     basetypes.NumberType{},
				"mtu_ignore": basetypes.BoolType{},
				"name":       basetypes.StringType{},
				"passive":    basetypes.BoolType{},
				"priority":   basetypes.NumberType{},
				"timing":     basetypes.StringType{},
				"vr_timing": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"dead_counts":         basetypes.NumberType{},
						"gr_delay":            basetypes.NumberType{},
						"hello_interval":      basetypes.NumberType{},
						"retransmit_interval": basetypes.NumberType{},
						"transit_delay":       basetypes.NumberType{},
					},
				},
			},
		}},
		"name": basetypes.StringType{},
		"range": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"advertise":  basetypes.BoolType{},
				"name":       basetypes.StringType{},
				"substitute": basetypes.StringType{},
			},
		}},
		"type": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"normal": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"abr": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"export_list":          basetypes.StringType{},
								"import_list":          basetypes.StringType{},
								"inbound_filter_list":  basetypes.StringType{},
								"outbound_filter_list": basetypes.StringType{},
							},
						},
					},
				},
				"nssa": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"abr": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"export_list":         basetypes.StringType{},
								"import_list":         basetypes.StringType{},
								"inbound_filter_list": basetypes.StringType{},
								"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"advertise": basetypes.BoolType{},
										"name":      basetypes.StringType{},
										"route_tag": basetypes.NumberType{},
									},
								}},
								"outbound_filter_list": basetypes.StringType{},
							},
						},
						"accept_summary": basetypes.BoolType{},
						"default_information_originate": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"metric":      basetypes.NumberType{},
								"metric_type": basetypes.StringType{},
							},
						},
						"default_route": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"advertise": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"metric": basetypes.NumberType{},
										"type":   basetypes.StringType{},
									},
								},
								"disable": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
						"no_summary": basetypes.BoolType{},
						"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"advertise": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"name": basetypes.StringType{},
								"suppress": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						}},
					},
				},
				"stub": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"abr": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"export_list":          basetypes.StringType{},
								"import_list":          basetypes.StringType{},
								"inbound_filter_list":  basetypes.StringType{},
								"outbound_filter_list": basetypes.StringType{},
							},
						},
						"accept_summary": basetypes.BoolType{},
						"default_route": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"advertise": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"metric": basetypes.NumberType{},
									},
								},
								"disable": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
						"default_route_metric": basetypes.NumberType{},
						"no_summary":           basetypes.BoolType{},
					},
				},
			},
		},
		"virtual_link": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"authentication": basetypes.StringType{},
				"bfd": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"profile": basetypes.StringType{},
					},
				},
				"enable":          basetypes.BoolType{},
				"instance_id":     basetypes.NumberType{},
				"interface_id":    basetypes.NumberType{},
				"name":            basetypes.StringType{},
				"neighbor_id":     basetypes.StringType{},
				"passive":         basetypes.BoolType{},
				"timing":          basetypes.StringType{},
				"transit_area_id": basetypes.StringType{},
				"vr_timing": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"dead_counts":         basetypes.NumberType{},
						"hello_interval":      basetypes.NumberType{},
						"retransmit_interval": basetypes.NumberType{},
						"transit_delay":       basetypes.NumberType{},
					},
				},
			},
		}},
		"vr_range": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"advertise": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"name": basetypes.StringType{},
				"suppress": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfAreaInner objects.
func (o LogicalRoutersVrfInnerOspfAreaInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfAreaInnerInterfaceInner model.
func (o LogicalRoutersVrfInnerOspfAreaInnerInterfaceInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"authentication": basetypes.StringType{},
		"bfd": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"profile": basetypes.StringType{},
			},
		},
		"enable": basetypes.BoolType{},
		"link_type": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"broadcast": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"p2mp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"neighbor": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":     basetypes.StringType{},
								"priority": basetypes.NumberType{},
							},
						}},
					},
				},
				"p2p": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
		"metric":     basetypes.NumberType{},
		"mtu_ignore": basetypes.BoolType{},
		"name":       basetypes.StringType{},
		"passive":    basetypes.BoolType{},
		"priority":   basetypes.NumberType{},
		"timing":     basetypes.StringType{},
		"vr_timing": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"dead_counts":         basetypes.NumberType{},
				"gr_delay":            basetypes.NumberType{},
				"hello_interval":      basetypes.NumberType{},
				"retransmit_interval": basetypes.NumberType{},
				"transit_delay":       basetypes.NumberType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfAreaInnerInterfaceInner objects.
func (o LogicalRoutersVrfInnerOspfAreaInnerInterfaceInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfAreaInnerInterfaceInnerLinkType model.
func (o LogicalRoutersVrfInnerOspfAreaInnerInterfaceInnerLinkType) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"broadcast": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"p2mp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"neighbor": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":     basetypes.StringType{},
						"priority": basetypes.NumberType{},
					},
				}},
			},
		},
		"p2p": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfAreaInnerInterfaceInnerLinkType objects.
func (o LogicalRoutersVrfInnerOspfAreaInnerInterfaceInnerLinkType) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfAreaInnerInterfaceInnerLinkTypeP2mp model.
func (o LogicalRoutersVrfInnerOspfAreaInnerInterfaceInnerLinkTypeP2mp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"neighbor": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":     basetypes.StringType{},
				"priority": basetypes.NumberType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfAreaInnerInterfaceInnerLinkTypeP2mp objects.
func (o LogicalRoutersVrfInnerOspfAreaInnerInterfaceInnerLinkTypeP2mp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfAreaInnerInterfaceInnerLinkTypeP2mpNeighborInner model.
func (o LogicalRoutersVrfInnerOspfAreaInnerInterfaceInnerLinkTypeP2mpNeighborInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":     basetypes.StringType{},
		"priority": basetypes.NumberType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfAreaInnerInterfaceInnerLinkTypeP2mpNeighborInner objects.
func (o LogicalRoutersVrfInnerOspfAreaInnerInterfaceInnerLinkTypeP2mpNeighborInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfAreaInnerInterfaceInnerVrTiming model.
func (o LogicalRoutersVrfInnerOspfAreaInnerInterfaceInnerVrTiming) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"dead_counts":         basetypes.NumberType{},
		"gr_delay":            basetypes.NumberType{},
		"hello_interval":      basetypes.NumberType{},
		"retransmit_interval": basetypes.NumberType{},
		"transit_delay":       basetypes.NumberType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfAreaInnerInterfaceInnerVrTiming objects.
func (o LogicalRoutersVrfInnerOspfAreaInnerInterfaceInnerVrTiming) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfAreaInnerRangeInner model.
func (o LogicalRoutersVrfInnerOspfAreaInnerRangeInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"advertise":  basetypes.BoolType{},
		"name":       basetypes.StringType{},
		"substitute": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfAreaInnerRangeInner objects.
func (o LogicalRoutersVrfInnerOspfAreaInnerRangeInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfAreaInnerType model.
func (o LogicalRoutersVrfInnerOspfAreaInnerType) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"normal": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"abr": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"export_list":          basetypes.StringType{},
						"import_list":          basetypes.StringType{},
						"inbound_filter_list":  basetypes.StringType{},
						"outbound_filter_list": basetypes.StringType{},
					},
				},
			},
		},
		"nssa": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"abr": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"export_list":         basetypes.StringType{},
						"import_list":         basetypes.StringType{},
						"inbound_filter_list": basetypes.StringType{},
						"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"advertise": basetypes.BoolType{},
								"name":      basetypes.StringType{},
								"route_tag": basetypes.NumberType{},
							},
						}},
						"outbound_filter_list": basetypes.StringType{},
					},
				},
				"accept_summary": basetypes.BoolType{},
				"default_information_originate": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"metric":      basetypes.NumberType{},
						"metric_type": basetypes.StringType{},
					},
				},
				"default_route": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"advertise": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"metric": basetypes.NumberType{},
								"type":   basetypes.StringType{},
							},
						},
						"disable": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
				"no_summary": basetypes.BoolType{},
				"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"advertise": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"name": basetypes.StringType{},
						"suppress": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				}},
			},
		},
		"stub": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"abr": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"export_list":          basetypes.StringType{},
						"import_list":          basetypes.StringType{},
						"inbound_filter_list":  basetypes.StringType{},
						"outbound_filter_list": basetypes.StringType{},
					},
				},
				"accept_summary": basetypes.BoolType{},
				"default_route": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"advertise": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"metric": basetypes.NumberType{},
							},
						},
						"disable": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
				"default_route_metric": basetypes.NumberType{},
				"no_summary":           basetypes.BoolType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfAreaInnerType objects.
func (o LogicalRoutersVrfInnerOspfAreaInnerType) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfAreaInnerTypeNormal model.
func (o LogicalRoutersVrfInnerOspfAreaInnerTypeNormal) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"abr": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"export_list":          basetypes.StringType{},
				"import_list":          basetypes.StringType{},
				"inbound_filter_list":  basetypes.StringType{},
				"outbound_filter_list": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfAreaInnerTypeNormal objects.
func (o LogicalRoutersVrfInnerOspfAreaInnerTypeNormal) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfAreaInnerTypeNormalAbr model.
func (o LogicalRoutersVrfInnerOspfAreaInnerTypeNormalAbr) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"export_list":          basetypes.StringType{},
		"import_list":          basetypes.StringType{},
		"inbound_filter_list":  basetypes.StringType{},
		"outbound_filter_list": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfAreaInnerTypeNormalAbr objects.
func (o LogicalRoutersVrfInnerOspfAreaInnerTypeNormalAbr) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfAreaInnerTypeNssa model.
func (o LogicalRoutersVrfInnerOspfAreaInnerTypeNssa) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"abr": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"export_list":         basetypes.StringType{},
				"import_list":         basetypes.StringType{},
				"inbound_filter_list": basetypes.StringType{},
				"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"advertise": basetypes.BoolType{},
						"name":      basetypes.StringType{},
						"route_tag": basetypes.NumberType{},
					},
				}},
				"outbound_filter_list": basetypes.StringType{},
			},
		},
		"accept_summary": basetypes.BoolType{},
		"default_information_originate": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"metric":      basetypes.NumberType{},
				"metric_type": basetypes.StringType{},
			},
		},
		"default_route": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"advertise": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"metric": basetypes.NumberType{},
						"type":   basetypes.StringType{},
					},
				},
				"disable": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
		"no_summary": basetypes.BoolType{},
		"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"advertise": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"name": basetypes.StringType{},
				"suppress": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfAreaInnerTypeNssa objects.
func (o LogicalRoutersVrfInnerOspfAreaInnerTypeNssa) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfAreaInnerTypeNssaAbr model.
func (o LogicalRoutersVrfInnerOspfAreaInnerTypeNssaAbr) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"export_list":         basetypes.StringType{},
		"import_list":         basetypes.StringType{},
		"inbound_filter_list": basetypes.StringType{},
		"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"advertise": basetypes.BoolType{},
				"name":      basetypes.StringType{},
				"route_tag": basetypes.NumberType{},
			},
		}},
		"outbound_filter_list": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfAreaInnerTypeNssaAbr objects.
func (o LogicalRoutersVrfInnerOspfAreaInnerTypeNssaAbr) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfAreaInnerTypeNssaAbrNssaExtRangeInner model.
func (o LogicalRoutersVrfInnerOspfAreaInnerTypeNssaAbrNssaExtRangeInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"advertise": basetypes.BoolType{},
		"name":      basetypes.StringType{},
		"route_tag": basetypes.NumberType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfAreaInnerTypeNssaAbrNssaExtRangeInner objects.
func (o LogicalRoutersVrfInnerOspfAreaInnerTypeNssaAbrNssaExtRangeInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfAreaInnerTypeNssaDefaultInformationOriginate model.
func (o LogicalRoutersVrfInnerOspfAreaInnerTypeNssaDefaultInformationOriginate) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"metric":      basetypes.NumberType{},
		"metric_type": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfAreaInnerTypeNssaDefaultInformationOriginate objects.
func (o LogicalRoutersVrfInnerOspfAreaInnerTypeNssaDefaultInformationOriginate) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfAreaInnerTypeNssaDefaultRoute model.
func (o LogicalRoutersVrfInnerOspfAreaInnerTypeNssaDefaultRoute) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"advertise": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"metric": basetypes.NumberType{},
				"type":   basetypes.StringType{},
			},
		},
		"disable": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfAreaInnerTypeNssaDefaultRoute objects.
func (o LogicalRoutersVrfInnerOspfAreaInnerTypeNssaDefaultRoute) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfAreaInnerTypeNssaDefaultRouteAdvertise model.
func (o LogicalRoutersVrfInnerOspfAreaInnerTypeNssaDefaultRouteAdvertise) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"metric": basetypes.NumberType{},
		"type":   basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfAreaInnerTypeNssaDefaultRouteAdvertise objects.
func (o LogicalRoutersVrfInnerOspfAreaInnerTypeNssaDefaultRouteAdvertise) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfAreaInnerTypeNssaNssaExtRangeInner model.
func (o LogicalRoutersVrfInnerOspfAreaInnerTypeNssaNssaExtRangeInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"advertise": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"name": basetypes.StringType{},
		"suppress": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfAreaInnerTypeNssaNssaExtRangeInner objects.
func (o LogicalRoutersVrfInnerOspfAreaInnerTypeNssaNssaExtRangeInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfAreaInnerTypeStub model.
func (o LogicalRoutersVrfInnerOspfAreaInnerTypeStub) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"abr": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"export_list":          basetypes.StringType{},
				"import_list":          basetypes.StringType{},
				"inbound_filter_list":  basetypes.StringType{},
				"outbound_filter_list": basetypes.StringType{},
			},
		},
		"accept_summary": basetypes.BoolType{},
		"default_route": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"advertise": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"metric": basetypes.NumberType{},
					},
				},
				"disable": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
		"default_route_metric": basetypes.NumberType{},
		"no_summary":           basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfAreaInnerTypeStub objects.
func (o LogicalRoutersVrfInnerOspfAreaInnerTypeStub) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfAreaInnerTypeStubDefaultRoute model.
func (o LogicalRoutersVrfInnerOspfAreaInnerTypeStubDefaultRoute) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"advertise": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"metric": basetypes.NumberType{},
			},
		},
		"disable": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfAreaInnerTypeStubDefaultRoute objects.
func (o LogicalRoutersVrfInnerOspfAreaInnerTypeStubDefaultRoute) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfAreaInnerTypeStubDefaultRouteAdvertise model.
func (o LogicalRoutersVrfInnerOspfAreaInnerTypeStubDefaultRouteAdvertise) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"metric": basetypes.NumberType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfAreaInnerTypeStubDefaultRouteAdvertise objects.
func (o LogicalRoutersVrfInnerOspfAreaInnerTypeStubDefaultRouteAdvertise) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfAreaInnerVirtualLinkInner model.
func (o LogicalRoutersVrfInnerOspfAreaInnerVirtualLinkInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"authentication": basetypes.StringType{},
		"bfd": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"profile": basetypes.StringType{},
			},
		},
		"enable":          basetypes.BoolType{},
		"instance_id":     basetypes.NumberType{},
		"interface_id":    basetypes.NumberType{},
		"name":            basetypes.StringType{},
		"neighbor_id":     basetypes.StringType{},
		"passive":         basetypes.BoolType{},
		"timing":          basetypes.StringType{},
		"transit_area_id": basetypes.StringType{},
		"vr_timing": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"dead_counts":         basetypes.NumberType{},
				"hello_interval":      basetypes.NumberType{},
				"retransmit_interval": basetypes.NumberType{},
				"transit_delay":       basetypes.NumberType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfAreaInnerVirtualLinkInner objects.
func (o LogicalRoutersVrfInnerOspfAreaInnerVirtualLinkInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfAreaInnerVirtualLinkInnerVrTiming model.
func (o LogicalRoutersVrfInnerOspfAreaInnerVirtualLinkInnerVrTiming) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"dead_counts":         basetypes.NumberType{},
		"hello_interval":      basetypes.NumberType{},
		"retransmit_interval": basetypes.NumberType{},
		"transit_delay":       basetypes.NumberType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfAreaInnerVirtualLinkInnerVrTiming objects.
func (o LogicalRoutersVrfInnerOspfAreaInnerVirtualLinkInnerVrTiming) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfAuthProfileInner model.
func (o LogicalRoutersVrfInnerOspfAuthProfileInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"md5": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"key":       basetypes.StringType{},
				"name":      basetypes.NumberType{},
				"preferred": basetypes.BoolType{},
			},
		}},
		"name":     basetypes.StringType{},
		"password": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfAuthProfileInner objects.
func (o LogicalRoutersVrfInnerOspfAuthProfileInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfAuthProfileInnerMd5Inner model.
func (o LogicalRoutersVrfInnerOspfAuthProfileInnerMd5Inner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"key":       basetypes.StringType{},
		"name":      basetypes.NumberType{},
		"preferred": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfAuthProfileInnerMd5Inner objects.
func (o LogicalRoutersVrfInnerOspfAuthProfileInnerMd5Inner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfExportRulesInner model.
func (o LogicalRoutersVrfInnerOspfExportRulesInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"metric":        basetypes.NumberType{},
		"name":          basetypes.StringType{},
		"new_path_type": basetypes.StringType{},
		"new_tag":       basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfExportRulesInner objects.
func (o LogicalRoutersVrfInnerOspfExportRulesInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfFloodPrevention model.
func (o LogicalRoutersVrfInnerOspfFloodPrevention) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"hello": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable":     basetypes.BoolType{},
				"max_packet": basetypes.NumberType{},
			},
		},
		"lsa": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable":     basetypes.BoolType{},
				"max_packet": basetypes.NumberType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfFloodPrevention objects.
func (o LogicalRoutersVrfInnerOspfFloodPrevention) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfFloodPreventionHello model.
func (o LogicalRoutersVrfInnerOspfFloodPreventionHello) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable":     basetypes.BoolType{},
		"max_packet": basetypes.NumberType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfFloodPreventionHello objects.
func (o LogicalRoutersVrfInnerOspfFloodPreventionHello) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfGracefulRestart model.
func (o LogicalRoutersVrfInnerOspfGracefulRestart) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable":                    basetypes.BoolType{},
		"grace_period":              basetypes.NumberType{},
		"helper_enable":             basetypes.BoolType{},
		"max_neighbor_restart_time": basetypes.NumberType{},
		"strict__l_s_a_checking":    basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfGracefulRestart objects.
func (o LogicalRoutersVrfInnerOspfGracefulRestart) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfVrTimers model.
func (o LogicalRoutersVrfInnerOspfVrTimers) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"lsa_interval":          basetypes.NumberType{},
		"spf_calculation_delay": basetypes.NumberType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfVrTimers objects.
func (o LogicalRoutersVrfInnerOspfVrTimers) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfv3 model.
func (o LogicalRoutersVrfInnerOspfv3) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"allow_redist_default_route": basetypes.BoolType{},
		"area": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"authentication": basetypes.StringType{},
				"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"authentication": basetypes.StringType{},
						"bfd": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"profile": basetypes.StringType{},
							},
						},
						"enable":      basetypes.BoolType{},
						"instance_id": basetypes.NumberType{},
						"link_type": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"broadcast": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"p2mp": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"neighbor": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"name":     basetypes.StringType{},
												"priority": basetypes.NumberType{},
											},
										}},
									},
								},
								"p2p": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
						"metric":     basetypes.NumberType{},
						"mtu_ignore": basetypes.BoolType{},
						"name":       basetypes.StringType{},
						"neighbor": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name": basetypes.StringType{},
							},
						}},
						"passive":  basetypes.BoolType{},
						"priority": basetypes.NumberType{},
						"timing":   basetypes.StringType{},
						"vr_timing": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"dead_counts":         basetypes.NumberType{},
								"gr_delay":            basetypes.NumberType{},
								"hello_interval":      basetypes.NumberType{},
								"retransmit_interval": basetypes.NumberType{},
								"transit_delay":       basetypes.NumberType{},
							},
						},
					},
				}},
				"name": basetypes.StringType{},
				"range": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"advertise": basetypes.BoolType{},
						"name":      basetypes.StringType{},
					},
				}},
				"type": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"normal": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"abr": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"export_list":          basetypes.StringType{},
										"import_list":          basetypes.StringType{},
										"inbound_filter_list":  basetypes.StringType{},
										"outbound_filter_list": basetypes.StringType{},
									},
								},
							},
						},
						"nssa": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"abr": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"export_list":         basetypes.StringType{},
										"import_list":         basetypes.StringType{},
										"inbound_filter_list": basetypes.StringType{},
										"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"advertise": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
												"name":      basetypes.StringType{},
												"route_tag": basetypes.NumberType{},
												"suppress": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
											},
										}},
										"outbound_filter_list": basetypes.StringType{},
									},
								},
								"accept_summary": basetypes.BoolType{},
								"default_information_originate": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"metric":      basetypes.NumberType{},
										"metric_type": basetypes.StringType{},
									},
								},
								"default_route": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"advertise": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"metric": basetypes.NumberType{},
												"type":   basetypes.StringType{},
											},
										},
										"disable": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								},
								"no_summary": basetypes.BoolType{},
								"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"advertise": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"name":      basetypes.StringType{},
										"route_tag": basetypes.NumberType{},
										"suppress": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								}},
							},
						},
						"stub": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"abr": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"export_list":          basetypes.StringType{},
										"import_list":          basetypes.StringType{},
										"inbound_filter_list":  basetypes.StringType{},
										"outbound_filter_list": basetypes.StringType{},
									},
								},
								"accept_summary": basetypes.BoolType{},
								"default_route": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"advertise": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"metric": basetypes.NumberType{},
											},
										},
										"disable": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								},
								"default_route_metric": basetypes.NumberType{},
								"no_summary":           basetypes.BoolType{},
							},
						},
					},
				},
				"virtual_link": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"authentication": basetypes.StringType{},
						"bfd": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"profile": basetypes.StringType{},
							},
						},
						"enable":          basetypes.BoolType{},
						"instance_id":     basetypes.NumberType{},
						"interface_id":    basetypes.NumberType{},
						"name":            basetypes.StringType{},
						"neighbor_id":     basetypes.StringType{},
						"passive":         basetypes.BoolType{},
						"timing":          basetypes.StringType{},
						"transit_area_id": basetypes.StringType{},
						"vr_timing": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"dead_counts":         basetypes.NumberType{},
								"hello_interval":      basetypes.NumberType{},
								"retransmit_interval": basetypes.NumberType{},
								"transit_delay":       basetypes.NumberType{},
							},
						},
					},
				}},
				"vr_range": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"advertise": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"name": basetypes.StringType{},
						"suppress": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				}},
			},
		}},
		"auth_profile": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ah": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"md5": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"key": basetypes.StringType{},
							},
						},
						"sha1": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"key": basetypes.StringType{},
							},
						},
						"sha256": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"key": basetypes.StringType{},
							},
						},
						"sha384": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"key": basetypes.StringType{},
							},
						},
						"sha512": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"key": basetypes.StringType{},
							},
						},
					},
				},
				"esp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"authentication": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"md5": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"key": basetypes.StringType{},
									},
								},
								"none": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"sha1": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"key": basetypes.StringType{},
									},
								},
								"sha256": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"key": basetypes.StringType{},
									},
								},
								"sha384": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"key": basetypes.StringType{},
									},
								},
								"sha512": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"key": basetypes.StringType{},
									},
								},
							},
						},
						"encryption": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"algorithm": basetypes.StringType{},
								"key":       basetypes.StringType{},
							},
						},
					},
				},
				"name": basetypes.StringType{},
				"spi":  basetypes.StringType{},
			},
		}},
		"disable_transit_traffic": basetypes.BoolType{},
		"enable":                  basetypes.BoolType{},
		"export_rules": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"metric":        basetypes.NumberType{},
				"name":          basetypes.StringType{},
				"new_path_type": basetypes.StringType{},
				"new_tag":       basetypes.StringType{},
			},
		}},
		"global_bfd": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"profile": basetypes.StringType{},
			},
		},
		"global_if_timer": basetypes.StringType{},
		"graceful_restart": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable":                    basetypes.BoolType{},
				"grace_period":              basetypes.NumberType{},
				"helper_enable":             basetypes.BoolType{},
				"max_neighbor_restart_time": basetypes.NumberType{},
				"strict__l_s_a_checking":    basetypes.BoolType{},
			},
		},
		"redistribution_profile": basetypes.StringType{},
		"reject_default_route":   basetypes.BoolType{},
		"router_id":              basetypes.StringType{},
		"spf_timer":              basetypes.StringType{},
		"vr_timers": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"lsa_interval":          basetypes.NumberType{},
				"spf_calculation_delay": basetypes.NumberType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfv3 objects.
func (o LogicalRoutersVrfInnerOspfv3) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfv3AreaInner model.
func (o LogicalRoutersVrfInnerOspfv3AreaInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"authentication": basetypes.StringType{},
		"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"authentication": basetypes.StringType{},
				"bfd": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"profile": basetypes.StringType{},
					},
				},
				"enable":      basetypes.BoolType{},
				"instance_id": basetypes.NumberType{},
				"link_type": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"broadcast": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"p2mp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"neighbor": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"name":     basetypes.StringType{},
										"priority": basetypes.NumberType{},
									},
								}},
							},
						},
						"p2p": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
				"metric":     basetypes.NumberType{},
				"mtu_ignore": basetypes.BoolType{},
				"name":       basetypes.StringType{},
				"neighbor": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name": basetypes.StringType{},
					},
				}},
				"passive":  basetypes.BoolType{},
				"priority": basetypes.NumberType{},
				"timing":   basetypes.StringType{},
				"vr_timing": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"dead_counts":         basetypes.NumberType{},
						"gr_delay":            basetypes.NumberType{},
						"hello_interval":      basetypes.NumberType{},
						"retransmit_interval": basetypes.NumberType{},
						"transit_delay":       basetypes.NumberType{},
					},
				},
			},
		}},
		"name": basetypes.StringType{},
		"range": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"advertise": basetypes.BoolType{},
				"name":      basetypes.StringType{},
			},
		}},
		"type": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"normal": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"abr": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"export_list":          basetypes.StringType{},
								"import_list":          basetypes.StringType{},
								"inbound_filter_list":  basetypes.StringType{},
								"outbound_filter_list": basetypes.StringType{},
							},
						},
					},
				},
				"nssa": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"abr": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"export_list":         basetypes.StringType{},
								"import_list":         basetypes.StringType{},
								"inbound_filter_list": basetypes.StringType{},
								"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"advertise": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"name":      basetypes.StringType{},
										"route_tag": basetypes.NumberType{},
										"suppress": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								}},
								"outbound_filter_list": basetypes.StringType{},
							},
						},
						"accept_summary": basetypes.BoolType{},
						"default_information_originate": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"metric":      basetypes.NumberType{},
								"metric_type": basetypes.StringType{},
							},
						},
						"default_route": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"advertise": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"metric": basetypes.NumberType{},
										"type":   basetypes.StringType{},
									},
								},
								"disable": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
						"no_summary": basetypes.BoolType{},
						"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"advertise": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"name":      basetypes.StringType{},
								"route_tag": basetypes.NumberType{},
								"suppress": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						}},
					},
				},
				"stub": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"abr": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"export_list":          basetypes.StringType{},
								"import_list":          basetypes.StringType{},
								"inbound_filter_list":  basetypes.StringType{},
								"outbound_filter_list": basetypes.StringType{},
							},
						},
						"accept_summary": basetypes.BoolType{},
						"default_route": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"advertise": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"metric": basetypes.NumberType{},
									},
								},
								"disable": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
						"default_route_metric": basetypes.NumberType{},
						"no_summary":           basetypes.BoolType{},
					},
				},
			},
		},
		"virtual_link": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"authentication": basetypes.StringType{},
				"bfd": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"profile": basetypes.StringType{},
					},
				},
				"enable":          basetypes.BoolType{},
				"instance_id":     basetypes.NumberType{},
				"interface_id":    basetypes.NumberType{},
				"name":            basetypes.StringType{},
				"neighbor_id":     basetypes.StringType{},
				"passive":         basetypes.BoolType{},
				"timing":          basetypes.StringType{},
				"transit_area_id": basetypes.StringType{},
				"vr_timing": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"dead_counts":         basetypes.NumberType{},
						"hello_interval":      basetypes.NumberType{},
						"retransmit_interval": basetypes.NumberType{},
						"transit_delay":       basetypes.NumberType{},
					},
				},
			},
		}},
		"vr_range": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"advertise": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"name": basetypes.StringType{},
				"suppress": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfv3AreaInner objects.
func (o LogicalRoutersVrfInnerOspfv3AreaInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfv3AreaInnerInterfaceInner model.
func (o LogicalRoutersVrfInnerOspfv3AreaInnerInterfaceInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"authentication": basetypes.StringType{},
		"bfd": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"profile": basetypes.StringType{},
			},
		},
		"enable":      basetypes.BoolType{},
		"instance_id": basetypes.NumberType{},
		"link_type": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"broadcast": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"p2mp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"neighbor": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":     basetypes.StringType{},
								"priority": basetypes.NumberType{},
							},
						}},
					},
				},
				"p2p": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
		"metric":     basetypes.NumberType{},
		"mtu_ignore": basetypes.BoolType{},
		"name":       basetypes.StringType{},
		"neighbor": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name": basetypes.StringType{},
			},
		}},
		"passive":  basetypes.BoolType{},
		"priority": basetypes.NumberType{},
		"timing":   basetypes.StringType{},
		"vr_timing": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"dead_counts":         basetypes.NumberType{},
				"gr_delay":            basetypes.NumberType{},
				"hello_interval":      basetypes.NumberType{},
				"retransmit_interval": basetypes.NumberType{},
				"transit_delay":       basetypes.NumberType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfv3AreaInnerInterfaceInner objects.
func (o LogicalRoutersVrfInnerOspfv3AreaInnerInterfaceInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfv3AreaInnerRangeInner model.
func (o LogicalRoutersVrfInnerOspfv3AreaInnerRangeInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"advertise": basetypes.BoolType{},
		"name":      basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfv3AreaInnerRangeInner objects.
func (o LogicalRoutersVrfInnerOspfv3AreaInnerRangeInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfv3AreaInnerType model.
func (o LogicalRoutersVrfInnerOspfv3AreaInnerType) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"normal": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"abr": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"export_list":          basetypes.StringType{},
						"import_list":          basetypes.StringType{},
						"inbound_filter_list":  basetypes.StringType{},
						"outbound_filter_list": basetypes.StringType{},
					},
				},
			},
		},
		"nssa": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"abr": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"export_list":         basetypes.StringType{},
						"import_list":         basetypes.StringType{},
						"inbound_filter_list": basetypes.StringType{},
						"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"advertise": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"name":      basetypes.StringType{},
								"route_tag": basetypes.NumberType{},
								"suppress": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						}},
						"outbound_filter_list": basetypes.StringType{},
					},
				},
				"accept_summary": basetypes.BoolType{},
				"default_information_originate": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"metric":      basetypes.NumberType{},
						"metric_type": basetypes.StringType{},
					},
				},
				"default_route": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"advertise": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"metric": basetypes.NumberType{},
								"type":   basetypes.StringType{},
							},
						},
						"disable": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
				"no_summary": basetypes.BoolType{},
				"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"advertise": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"name":      basetypes.StringType{},
						"route_tag": basetypes.NumberType{},
						"suppress": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				}},
			},
		},
		"stub": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"abr": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"export_list":          basetypes.StringType{},
						"import_list":          basetypes.StringType{},
						"inbound_filter_list":  basetypes.StringType{},
						"outbound_filter_list": basetypes.StringType{},
					},
				},
				"accept_summary": basetypes.BoolType{},
				"default_route": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"advertise": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"metric": basetypes.NumberType{},
							},
						},
						"disable": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
				"default_route_metric": basetypes.NumberType{},
				"no_summary":           basetypes.BoolType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfv3AreaInnerType objects.
func (o LogicalRoutersVrfInnerOspfv3AreaInnerType) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfv3AreaInnerTypeNssa model.
func (o LogicalRoutersVrfInnerOspfv3AreaInnerTypeNssa) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"abr": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"export_list":         basetypes.StringType{},
				"import_list":         basetypes.StringType{},
				"inbound_filter_list": basetypes.StringType{},
				"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"advertise": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"name":      basetypes.StringType{},
						"route_tag": basetypes.NumberType{},
						"suppress": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				}},
				"outbound_filter_list": basetypes.StringType{},
			},
		},
		"accept_summary": basetypes.BoolType{},
		"default_information_originate": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"metric":      basetypes.NumberType{},
				"metric_type": basetypes.StringType{},
			},
		},
		"default_route": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"advertise": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"metric": basetypes.NumberType{},
						"type":   basetypes.StringType{},
					},
				},
				"disable": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
		"no_summary": basetypes.BoolType{},
		"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"advertise": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"name":      basetypes.StringType{},
				"route_tag": basetypes.NumberType{},
				"suppress": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfv3AreaInnerTypeNssa objects.
func (o LogicalRoutersVrfInnerOspfv3AreaInnerTypeNssa) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfv3AreaInnerTypeNssaAbr model.
func (o LogicalRoutersVrfInnerOspfv3AreaInnerTypeNssaAbr) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"export_list":         basetypes.StringType{},
		"import_list":         basetypes.StringType{},
		"inbound_filter_list": basetypes.StringType{},
		"nssa_ext_range": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"advertise": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"name":      basetypes.StringType{},
				"route_tag": basetypes.NumberType{},
				"suppress": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		}},
		"outbound_filter_list": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfv3AreaInnerTypeNssaAbr objects.
func (o LogicalRoutersVrfInnerOspfv3AreaInnerTypeNssaAbr) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfv3AreaInnerTypeNssaAbrNssaExtRangeInner model.
func (o LogicalRoutersVrfInnerOspfv3AreaInnerTypeNssaAbrNssaExtRangeInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"advertise": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"name":      basetypes.StringType{},
		"route_tag": basetypes.NumberType{},
		"suppress": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfv3AreaInnerTypeNssaAbrNssaExtRangeInner objects.
func (o LogicalRoutersVrfInnerOspfv3AreaInnerTypeNssaAbrNssaExtRangeInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfv3AuthProfileInner model.
func (o LogicalRoutersVrfInnerOspfv3AuthProfileInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ah": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"md5": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"key": basetypes.StringType{},
					},
				},
				"sha1": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"key": basetypes.StringType{},
					},
				},
				"sha256": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"key": basetypes.StringType{},
					},
				},
				"sha384": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"key": basetypes.StringType{},
					},
				},
				"sha512": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"key": basetypes.StringType{},
					},
				},
			},
		},
		"esp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"authentication": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"md5": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"key": basetypes.StringType{},
							},
						},
						"none": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"sha1": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"key": basetypes.StringType{},
							},
						},
						"sha256": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"key": basetypes.StringType{},
							},
						},
						"sha384": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"key": basetypes.StringType{},
							},
						},
						"sha512": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"key": basetypes.StringType{},
							},
						},
					},
				},
				"encryption": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"algorithm": basetypes.StringType{},
						"key":       basetypes.StringType{},
					},
				},
			},
		},
		"name": basetypes.StringType{},
		"spi":  basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfv3AuthProfileInner objects.
func (o LogicalRoutersVrfInnerOspfv3AuthProfileInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfv3AuthProfileInnerAh model.
func (o LogicalRoutersVrfInnerOspfv3AuthProfileInnerAh) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"md5": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"key": basetypes.StringType{},
			},
		},
		"sha1": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"key": basetypes.StringType{},
			},
		},
		"sha256": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"key": basetypes.StringType{},
			},
		},
		"sha384": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"key": basetypes.StringType{},
			},
		},
		"sha512": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"key": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfv3AuthProfileInnerAh objects.
func (o LogicalRoutersVrfInnerOspfv3AuthProfileInnerAh) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfv3AuthProfileInnerAhMd5 model.
func (o LogicalRoutersVrfInnerOspfv3AuthProfileInnerAhMd5) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"key": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfv3AuthProfileInnerAhMd5 objects.
func (o LogicalRoutersVrfInnerOspfv3AuthProfileInnerAhMd5) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfv3AuthProfileInnerEsp model.
func (o LogicalRoutersVrfInnerOspfv3AuthProfileInnerEsp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"authentication": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"md5": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"key": basetypes.StringType{},
					},
				},
				"none": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"sha1": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"key": basetypes.StringType{},
					},
				},
				"sha256": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"key": basetypes.StringType{},
					},
				},
				"sha384": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"key": basetypes.StringType{},
					},
				},
				"sha512": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"key": basetypes.StringType{},
					},
				},
			},
		},
		"encryption": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"algorithm": basetypes.StringType{},
				"key":       basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfv3AuthProfileInnerEsp objects.
func (o LogicalRoutersVrfInnerOspfv3AuthProfileInnerEsp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfv3AuthProfileInnerEspAuthentication model.
func (o LogicalRoutersVrfInnerOspfv3AuthProfileInnerEspAuthentication) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"md5": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"key": basetypes.StringType{},
			},
		},
		"none": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"sha1": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"key": basetypes.StringType{},
			},
		},
		"sha256": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"key": basetypes.StringType{},
			},
		},
		"sha384": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"key": basetypes.StringType{},
			},
		},
		"sha512": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"key": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfv3AuthProfileInnerEspAuthentication objects.
func (o LogicalRoutersVrfInnerOspfv3AuthProfileInnerEspAuthentication) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerOspfv3AuthProfileInnerEspEncryption model.
func (o LogicalRoutersVrfInnerOspfv3AuthProfileInnerEspEncryption) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"algorithm": basetypes.StringType{},
		"key":       basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerOspfv3AuthProfileInnerEspEncryption objects.
func (o LogicalRoutersVrfInnerOspfv3AuthProfileInnerEspEncryption) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerRibFilter model.
func (o LogicalRoutersVrfInnerRibFilter) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ipv4": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"bgp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"route_map": basetypes.StringType{},
					},
				},
				"ospf": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"route_map": basetypes.StringType{},
					},
				},
				"rip": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"route_map": basetypes.StringType{},
					},
				},
				"static": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"route_map": basetypes.StringType{},
					},
				},
			},
		},
		"ipv6": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"bgp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"route_map": basetypes.StringType{},
					},
				},
				"ospfv3": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"route_map": basetypes.StringType{},
					},
				},
				"static": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"route_map": basetypes.StringType{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerRibFilter objects.
func (o LogicalRoutersVrfInnerRibFilter) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerRibFilterIpv4 model.
func (o LogicalRoutersVrfInnerRibFilterIpv4) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"bgp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"route_map": basetypes.StringType{},
			},
		},
		"ospf": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"route_map": basetypes.StringType{},
			},
		},
		"rip": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"route_map": basetypes.StringType{},
			},
		},
		"static": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"route_map": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerRibFilterIpv4 objects.
func (o LogicalRoutersVrfInnerRibFilterIpv4) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerRibFilterIpv4Bgp model.
func (o LogicalRoutersVrfInnerRibFilterIpv4Bgp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"route_map": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerRibFilterIpv4Bgp objects.
func (o LogicalRoutersVrfInnerRibFilterIpv4Bgp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerRibFilterIpv6 model.
func (o LogicalRoutersVrfInnerRibFilterIpv6) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"bgp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"route_map": basetypes.StringType{},
			},
		},
		"ospfv3": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"route_map": basetypes.StringType{},
			},
		},
		"static": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"route_map": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerRibFilterIpv6 objects.
func (o LogicalRoutersVrfInnerRibFilterIpv6) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerRip model.
func (o LogicalRoutersVrfInnerRip) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"auth_profile":                  basetypes.StringType{},
		"default_information_originate": basetypes.BoolType{},
		"enable":                        basetypes.BoolType{},
		"global_bfd": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"profile": basetypes.StringType{},
			},
		},
		"global_inbound_distribute_list": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"access_list": basetypes.StringType{},
			},
		},
		"global_outbound_distribute_list": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"access_list": basetypes.StringType{},
			},
		},
		"global_timer": basetypes.StringType{},
		"interface": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"authentication": basetypes.StringType{},
				"bfd": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"profile": basetypes.StringType{},
					},
				},
				"enable": basetypes.BoolType{},
				"interface_inbound_distribute_list": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"access_list": basetypes.StringType{},
						"metric":      basetypes.NumberType{},
					},
				},
				"interface_outbound_distribute_list": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"access_list": basetypes.StringType{},
						"metric":      basetypes.NumberType{},
					},
				},
				"mode":          basetypes.StringType{},
				"name":          basetypes.StringType{},
				"split_horizon": basetypes.StringType{},
			},
		}},
		"redistribution_profile": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerRip objects.
func (o LogicalRoutersVrfInnerRip) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerRipGlobalInboundDistributeList model.
func (o LogicalRoutersVrfInnerRipGlobalInboundDistributeList) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"access_list": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerRipGlobalInboundDistributeList objects.
func (o LogicalRoutersVrfInnerRipGlobalInboundDistributeList) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerRipInterfaceInner model.
func (o LogicalRoutersVrfInnerRipInterfaceInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"authentication": basetypes.StringType{},
		"bfd": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"profile": basetypes.StringType{},
			},
		},
		"enable": basetypes.BoolType{},
		"interface_inbound_distribute_list": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"access_list": basetypes.StringType{},
				"metric":      basetypes.NumberType{},
			},
		},
		"interface_outbound_distribute_list": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"access_list": basetypes.StringType{},
				"metric":      basetypes.NumberType{},
			},
		},
		"mode":          basetypes.StringType{},
		"name":          basetypes.StringType{},
		"split_horizon": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerRipInterfaceInner objects.
func (o LogicalRoutersVrfInnerRipInterfaceInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerRipInterfaceInnerInterfaceInboundDistributeList model.
func (o LogicalRoutersVrfInnerRipInterfaceInnerInterfaceInboundDistributeList) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"access_list": basetypes.StringType{},
		"metric":      basetypes.NumberType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerRipInterfaceInnerInterfaceInboundDistributeList objects.
func (o LogicalRoutersVrfInnerRipInterfaceInnerInterfaceInboundDistributeList) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerRoutingTable model.
func (o LogicalRoutersVrfInnerRoutingTable) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ip": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"static_route": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"admin_dist": basetypes.NumberType{},
						"bfd": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"profile": basetypes.StringType{},
							},
						},
						"destination": basetypes.StringType{},
						"interface":   basetypes.StringType{},
						"metric":      basetypes.NumberType{},
						"name":        basetypes.StringType{},
						"nexthop": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"discard": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"fqdn":         basetypes.StringType{},
								"ipv6_address": basetypes.StringType{},
								"next_lr":      basetypes.StringType{},
								"next_vr":      basetypes.StringType{},
								"receive": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"tunnel": basetypes.StringType{},
							},
						},
						"path_monitor": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable":            basetypes.BoolType{},
								"failure_condition": basetypes.StringType{},
								"hold_time":         basetypes.NumberType{},
								"monitor_destinations": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"count":            basetypes.NumberType{},
										"destination":      basetypes.StringType{},
										"destination_fqdn": basetypes.StringType{},
										"enable":           basetypes.BoolType{},
										"interval":         basetypes.NumberType{},
										"name":             basetypes.StringType{},
										"source":           basetypes.StringType{},
									},
								}},
							},
						},
						"route_table": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"both": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"multicast": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"no_install": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"unicast": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
					},
				}},
			},
		},
		"ipv6": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"static_route": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"admin_dist": basetypes.NumberType{},
						"bfd": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"profile": basetypes.StringType{},
							},
						},
						"destination": basetypes.StringType{},
						"interface":   basetypes.StringType{},
						"metric":      basetypes.NumberType{},
						"name":        basetypes.StringType{},
						"nexthop": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"discard": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"fqdn":         basetypes.StringType{},
								"ipv6_address": basetypes.StringType{},
								"next_lr":      basetypes.StringType{},
								"next_vr":      basetypes.StringType{},
								"receive": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"tunnel": basetypes.StringType{},
							},
						},
						"option": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"passive": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
						"path_monitor": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable":            basetypes.BoolType{},
								"failure_condition": basetypes.StringType{},
								"hold_time":         basetypes.NumberType{},
								"monitor_destinations": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"count":            basetypes.NumberType{},
										"destination":      basetypes.StringType{},
										"destination_fqdn": basetypes.StringType{},
										"enable":           basetypes.BoolType{},
										"interval":         basetypes.NumberType{},
										"name":             basetypes.StringType{},
										"source":           basetypes.StringType{},
									},
								}},
							},
						},
						"route_table": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"both": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"multicast": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"no_install": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"unicast": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
					},
				}},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerRoutingTable objects.
func (o LogicalRoutersVrfInnerRoutingTable) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerRoutingTableIp model.
func (o LogicalRoutersVrfInnerRoutingTableIp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"static_route": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"admin_dist": basetypes.NumberType{},
				"bfd": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"profile": basetypes.StringType{},
					},
				},
				"destination": basetypes.StringType{},
				"interface":   basetypes.StringType{},
				"metric":      basetypes.NumberType{},
				"name":        basetypes.StringType{},
				"nexthop": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"discard": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"fqdn":         basetypes.StringType{},
						"ipv6_address": basetypes.StringType{},
						"next_lr":      basetypes.StringType{},
						"next_vr":      basetypes.StringType{},
						"receive": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"tunnel": basetypes.StringType{},
					},
				},
				"path_monitor": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable":            basetypes.BoolType{},
						"failure_condition": basetypes.StringType{},
						"hold_time":         basetypes.NumberType{},
						"monitor_destinations": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"count":            basetypes.NumberType{},
								"destination":      basetypes.StringType{},
								"destination_fqdn": basetypes.StringType{},
								"enable":           basetypes.BoolType{},
								"interval":         basetypes.NumberType{},
								"name":             basetypes.StringType{},
								"source":           basetypes.StringType{},
							},
						}},
					},
				},
				"route_table": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"both": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"multicast": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"no_install": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"unicast": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerRoutingTableIp objects.
func (o LogicalRoutersVrfInnerRoutingTableIp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerRoutingTableIpStaticRouteInner model.
func (o LogicalRoutersVrfInnerRoutingTableIpStaticRouteInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"admin_dist": basetypes.NumberType{},
		"bfd": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"profile": basetypes.StringType{},
			},
		},
		"destination": basetypes.StringType{},
		"interface":   basetypes.StringType{},
		"metric":      basetypes.NumberType{},
		"name":        basetypes.StringType{},
		"nexthop": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"discard": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"fqdn":         basetypes.StringType{},
				"ipv6_address": basetypes.StringType{},
				"next_lr":      basetypes.StringType{},
				"next_vr":      basetypes.StringType{},
				"receive": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"tunnel": basetypes.StringType{},
			},
		},
		"path_monitor": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable":            basetypes.BoolType{},
				"failure_condition": basetypes.StringType{},
				"hold_time":         basetypes.NumberType{},
				"monitor_destinations": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"count":            basetypes.NumberType{},
						"destination":      basetypes.StringType{},
						"destination_fqdn": basetypes.StringType{},
						"enable":           basetypes.BoolType{},
						"interval":         basetypes.NumberType{},
						"name":             basetypes.StringType{},
						"source":           basetypes.StringType{},
					},
				}},
			},
		},
		"route_table": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"both": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"multicast": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"no_install": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"unicast": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerRoutingTableIpStaticRouteInner objects.
func (o LogicalRoutersVrfInnerRoutingTableIpStaticRouteInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerNexthop model.
func (o LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerNexthop) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"discard": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"fqdn":         basetypes.StringType{},
		"ipv6_address": basetypes.StringType{},
		"next_lr":      basetypes.StringType{},
		"next_vr":      basetypes.StringType{},
		"receive": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"tunnel": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerNexthop objects.
func (o LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerNexthop) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerPathMonitor model.
func (o LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerPathMonitor) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable":            basetypes.BoolType{},
		"failure_condition": basetypes.StringType{},
		"hold_time":         basetypes.NumberType{},
		"monitor_destinations": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"count":            basetypes.NumberType{},
				"destination":      basetypes.StringType{},
				"destination_fqdn": basetypes.StringType{},
				"enable":           basetypes.BoolType{},
				"interval":         basetypes.NumberType{},
				"name":             basetypes.StringType{},
				"source":           basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerPathMonitor objects.
func (o LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerPathMonitor) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerPathMonitorMonitorDestinationsInner model.
func (o LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerPathMonitorMonitorDestinationsInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"count":            basetypes.NumberType{},
		"destination":      basetypes.StringType{},
		"destination_fqdn": basetypes.StringType{},
		"enable":           basetypes.BoolType{},
		"interval":         basetypes.NumberType{},
		"name":             basetypes.StringType{},
		"source":           basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerPathMonitorMonitorDestinationsInner objects.
func (o LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerPathMonitorMonitorDestinationsInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerRouteTable model.
func (o LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerRouteTable) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"both": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"multicast": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"no_install": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"unicast": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerRouteTable objects.
func (o LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerRouteTable) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerRoutingTableIpv6 model.
func (o LogicalRoutersVrfInnerRoutingTableIpv6) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"static_route": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"admin_dist": basetypes.NumberType{},
				"bfd": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"profile": basetypes.StringType{},
					},
				},
				"destination": basetypes.StringType{},
				"interface":   basetypes.StringType{},
				"metric":      basetypes.NumberType{},
				"name":        basetypes.StringType{},
				"nexthop": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"discard": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"fqdn":         basetypes.StringType{},
						"ipv6_address": basetypes.StringType{},
						"next_lr":      basetypes.StringType{},
						"next_vr":      basetypes.StringType{},
						"receive": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"tunnel": basetypes.StringType{},
					},
				},
				"option": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"passive": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
				"path_monitor": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable":            basetypes.BoolType{},
						"failure_condition": basetypes.StringType{},
						"hold_time":         basetypes.NumberType{},
						"monitor_destinations": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"count":            basetypes.NumberType{},
								"destination":      basetypes.StringType{},
								"destination_fqdn": basetypes.StringType{},
								"enable":           basetypes.BoolType{},
								"interval":         basetypes.NumberType{},
								"name":             basetypes.StringType{},
								"source":           basetypes.StringType{},
							},
						}},
					},
				},
				"route_table": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"both": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"multicast": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"no_install": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"unicast": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerRoutingTableIpv6 objects.
func (o LogicalRoutersVrfInnerRoutingTableIpv6) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerRoutingTableIpv6StaticRouteInner model.
func (o LogicalRoutersVrfInnerRoutingTableIpv6StaticRouteInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"admin_dist": basetypes.NumberType{},
		"bfd": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"profile": basetypes.StringType{},
			},
		},
		"destination": basetypes.StringType{},
		"interface":   basetypes.StringType{},
		"metric":      basetypes.NumberType{},
		"name":        basetypes.StringType{},
		"nexthop": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"discard": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"fqdn":         basetypes.StringType{},
				"ipv6_address": basetypes.StringType{},
				"next_lr":      basetypes.StringType{},
				"next_vr":      basetypes.StringType{},
				"receive": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"tunnel": basetypes.StringType{},
			},
		},
		"option": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"passive": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
		"path_monitor": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable":            basetypes.BoolType{},
				"failure_condition": basetypes.StringType{},
				"hold_time":         basetypes.NumberType{},
				"monitor_destinations": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"count":            basetypes.NumberType{},
						"destination":      basetypes.StringType{},
						"destination_fqdn": basetypes.StringType{},
						"enable":           basetypes.BoolType{},
						"interval":         basetypes.NumberType{},
						"name":             basetypes.StringType{},
						"source":           basetypes.StringType{},
					},
				}},
			},
		},
		"route_table": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"both": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"multicast": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"no_install": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"unicast": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerRoutingTableIpv6StaticRouteInner objects.
func (o LogicalRoutersVrfInnerRoutingTableIpv6StaticRouteInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerRoutingTableIpv6StaticRouteInnerOption model.
func (o LogicalRoutersVrfInnerRoutingTableIpv6StaticRouteInnerOption) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"passive": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerRoutingTableIpv6StaticRouteInnerOption objects.
func (o LogicalRoutersVrfInnerRoutingTableIpv6StaticRouteInnerOption) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LogicalRoutersVrfInnerVrAdminDists model.
func (o LogicalRoutersVrfInnerVrAdminDists) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ebgp":        basetypes.NumberType{},
		"ibgp":        basetypes.NumberType{},
		"ospf_ext":    basetypes.NumberType{},
		"ospf_int":    basetypes.NumberType{},
		"ospfv3_ext":  basetypes.NumberType{},
		"ospfv3_int":  basetypes.NumberType{},
		"rip":         basetypes.NumberType{},
		"static":      basetypes.NumberType{},
		"static_ipv6": basetypes.NumberType{},
	}
}

// AttrType returns the attribute type for a list of LogicalRoutersVrfInnerVrAdminDists objects.
func (o LogicalRoutersVrfInnerVrAdminDists) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// LogicalRoutersResourceSchema defines the schema for LogicalRouters resource
var LogicalRoutersResourceSchema = schema.Schema{
	MarkdownDescription: "LogicalRouter resource",
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
		"name": schema.StringAttribute{
			MarkdownDescription: "Name",
			Required:            true,
		},
		"routing_stack": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("legacy", "advanced"),
			},
			MarkdownDescription: "Routing stack",
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
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"vrf": schema.ListNestedAttribute{
			MarkdownDescription: "Vrf",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"admin_dists": schema.SingleNestedAttribute{
						MarkdownDescription: "Admin dists",
						Optional:            true,
						Attributes: map[string]schema.Attribute{
							"bgp_external": schema.Float64Attribute{
								MarkdownDescription: "Bgp external",
								Optional:            true,
							},
							"bgp_internal": schema.Float64Attribute{
								MarkdownDescription: "Bgp internal",
								Optional:            true,
							},
							"bgp_local": schema.Float64Attribute{
								MarkdownDescription: "Bgp local",
								Optional:            true,
							},
							"ospf_ext": schema.Float64Attribute{
								MarkdownDescription: "Ospf ext",
								Optional:            true,
							},
							"ospf_inter": schema.Float64Attribute{
								MarkdownDescription: "Ospf inter",
								Optional:            true,
							},
							"ospf_intra": schema.Float64Attribute{
								MarkdownDescription: "Ospf intra",
								Optional:            true,
							},
							"ospfv3_ext": schema.Float64Attribute{
								MarkdownDescription: "Ospfv3 ext",
								Optional:            true,
							},
							"ospfv3_inter": schema.Float64Attribute{
								MarkdownDescription: "Ospfv3 inter",
								Optional:            true,
							},
							"ospfv3_intra": schema.Float64Attribute{
								MarkdownDescription: "Ospfv3 intra",
								Optional:            true,
							},
							"rip": schema.Float64Attribute{
								MarkdownDescription: "Rip",
								Optional:            true,
							},
							"static": schema.Float64Attribute{
								MarkdownDescription: "Static",
								Optional:            true,
							},
							"static_ipv6": schema.Float64Attribute{
								MarkdownDescription: "Static ipv6",
								Optional:            true,
							},
						},
					},
					"bgp": schema.SingleNestedAttribute{
						MarkdownDescription: "Bgp",
						Optional:            true,
						Attributes: map[string]schema.Attribute{
							"advertise_network": schema.SingleNestedAttribute{
								MarkdownDescription: "Advertise network",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"ipv4": schema.SingleNestedAttribute{
										MarkdownDescription: "Ipv4",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"network": schema.ListNestedAttribute{
												MarkdownDescription: "Network",
												Optional:            true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"backdoor": schema.BoolAttribute{
															MarkdownDescription: "Backdoor",
															Optional:            true,
														},
														"multicast": schema.BoolAttribute{
															MarkdownDescription: "Multicast",
															Optional:            true,
														},
														"name": schema.StringAttribute{
															MarkdownDescription: "Name",
															Required:            true,
														},
														"unicast": schema.BoolAttribute{
															MarkdownDescription: "Unicast",
															Optional:            true,
														},
													},
												},
											},
										},
									},
									"ipv6": schema.SingleNestedAttribute{
										MarkdownDescription: "Ipv6",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"network": schema.ListNestedAttribute{
												MarkdownDescription: "Network",
												Optional:            true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															MarkdownDescription: "Name",
															Required:            true,
														},
														"unicast": schema.BoolAttribute{
															MarkdownDescription: "Unicast",
															Optional:            true,
														},
													},
												},
											},
										},
									},
								},
							},
							"aggregate": schema.SingleNestedAttribute{
								MarkdownDescription: "Aggregate",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"aggregate_med": schema.BoolAttribute{
										MarkdownDescription: "Aggregate med",
										Optional:            true,
									},
								},
							},
							"aggregate_routes": schema.ListNestedAttribute{
								MarkdownDescription: "Aggregate routes",
								Optional:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"as_set": schema.BoolAttribute{
											MarkdownDescription: "As set",
											Optional:            true,
										},
										"description": schema.StringAttribute{
											MarkdownDescription: "Description",
											Optional:            true,
										},
										"enable": schema.BoolAttribute{
											MarkdownDescription: "Enable",
											Optional:            true,
										},
										"name": schema.StringAttribute{
											MarkdownDescription: "Name",
											Required:            true,
										},
										"same_med": schema.BoolAttribute{
											MarkdownDescription: "Same med",
											Optional:            true,
										},
										"summary_only": schema.BoolAttribute{
											MarkdownDescription: "Summary only",
											Optional:            true,
										},
										"type": schema.SingleNestedAttribute{
											MarkdownDescription: "Type",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"ipv4": schema.SingleNestedAttribute{
													Validators: []validator.Object{
														objectvalidator.ExactlyOneOf(
															path.MatchRelative().AtParent().AtName("ipv6"),
														),
													},
													MarkdownDescription: "Ipv4",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"attribute_map": schema.StringAttribute{
															MarkdownDescription: "Attribute map",
															Optional:            true,
														},
														"summary_prefix": schema.StringAttribute{
															MarkdownDescription: "Summary prefix",
															Optional:            true,
														},
														"suppress_map": schema.StringAttribute{
															MarkdownDescription: "Suppress map",
															Optional:            true,
														},
													},
												},
												"ipv6": schema.SingleNestedAttribute{
													Validators: []validator.Object{
														objectvalidator.ExactlyOneOf(
															path.MatchRelative().AtParent().AtName("ipv4"),
														),
													},
													MarkdownDescription: "Ipv6",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"attribute_map": schema.StringAttribute{
															MarkdownDescription: "Attribute map",
															Optional:            true,
														},
														"summary_prefix": schema.StringAttribute{
															MarkdownDescription: "Summary prefix",
															Optional:            true,
														},
														"suppress_map": schema.StringAttribute{
															MarkdownDescription: "Suppress map",
															Optional:            true,
														},
													},
												},
											},
										},
									},
								},
							},
							"allow_redist_default_route": schema.BoolAttribute{
								MarkdownDescription: "Allow redist default route",
								Optional:            true,
							},
							"always_advertise_network_route": schema.BoolAttribute{
								MarkdownDescription: "Always advertise network route",
								Optional:            true,
							},
							"as_format": schema.StringAttribute{
								MarkdownDescription: "As format",
								Optional:            true,
							},
							"confederation_member_as": schema.StringAttribute{
								MarkdownDescription: "Confederation member as",
								Optional:            true,
							},
							"default_local_preference": schema.Float64Attribute{
								MarkdownDescription: "Default local preference",
								Optional:            true,
							},
							"ecmp_multi_as": schema.BoolAttribute{
								MarkdownDescription: "Ecmp multi as",
								Optional:            true,
							},
							"enable": schema.BoolAttribute{
								MarkdownDescription: "Enable",
								Optional:            true,
							},
							"enforce_first_as": schema.BoolAttribute{
								MarkdownDescription: "Enforce first as",
								Optional:            true,
							},
							"fast_external_failover": schema.BoolAttribute{
								MarkdownDescription: "Fast external failover",
								Optional:            true,
							},
							"global_bfd": schema.SingleNestedAttribute{
								MarkdownDescription: "Global bfd",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"profile": schema.StringAttribute{
										MarkdownDescription: "Profile",
										Optional:            true,
									},
								},
							},
							"graceful_restart": schema.SingleNestedAttribute{
								MarkdownDescription: "Graceful restart",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"enable": schema.BoolAttribute{
										MarkdownDescription: "Enable",
										Optional:            true,
									},
									"local_restart_time": schema.Float64Attribute{
										MarkdownDescription: "Local restart time",
										Optional:            true,
									},
									"max_peer_restart_time": schema.Float64Attribute{
										MarkdownDescription: "Max peer restart time",
										Optional:            true,
									},
									"stale_route_time": schema.Float64Attribute{
										MarkdownDescription: "Stale route time",
										Optional:            true,
									},
								},
							},
							"graceful_shutdown": schema.BoolAttribute{
								MarkdownDescription: "Graceful shutdown",
								Optional:            true,
							},
							"install_route": schema.BoolAttribute{
								MarkdownDescription: "Install route",
								Optional:            true,
							},
							"local_as": schema.StringAttribute{
								MarkdownDescription: "Local as",
								Optional:            true,
							},
							"med": schema.SingleNestedAttribute{
								MarkdownDescription: "Med",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"always_compare_med": schema.BoolAttribute{
										MarkdownDescription: "Always compare med",
										Optional:            true,
									},
									"deterministic_med_comparison": schema.BoolAttribute{
										MarkdownDescription: "Deterministic med comparison",
										Optional:            true,
									},
								},
							},
							"peer_group": schema.ListNestedAttribute{
								MarkdownDescription: "Peer group",
								Optional:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"address_family": schema.SingleNestedAttribute{
											MarkdownDescription: "Address family",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"ipv4": schema.StringAttribute{
													MarkdownDescription: "Ipv4",
													Optional:            true,
												},
												"ipv6": schema.StringAttribute{
													MarkdownDescription: "Ipv6",
													Optional:            true,
												},
											},
										},
										"aggregated_confed_as_path": schema.BoolAttribute{
											MarkdownDescription: "Aggregated confed as path",
											Optional:            true,
										},
										"connection_options": schema.SingleNestedAttribute{
											MarkdownDescription: "Connection options",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"authentication": schema.StringAttribute{
													MarkdownDescription: "Authentication",
													Optional:            true,
												},
												"dampening": schema.StringAttribute{
													MarkdownDescription: "Dampening",
													Optional:            true,
												},
												"multihop": schema.Float64Attribute{
													MarkdownDescription: "Multihop",
													Optional:            true,
												},
												"timers": schema.StringAttribute{
													MarkdownDescription: "Timers",
													Optional:            true,
												},
											},
										},
										"enable": schema.BoolAttribute{
											MarkdownDescription: "Enable",
											Optional:            true,
										},
										"filtering_profile": schema.SingleNestedAttribute{
											MarkdownDescription: "Filtering profile",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"ipv4": schema.StringAttribute{
													MarkdownDescription: "Ipv4",
													Optional:            true,
												},
												"ipv6": schema.StringAttribute{
													MarkdownDescription: "Ipv6",
													Optional:            true,
												},
											},
										},
										"name": schema.StringAttribute{
											MarkdownDescription: "Name",
											Required:            true,
										},
										"peer": schema.ListNestedAttribute{
											MarkdownDescription: "Peer",
											Optional:            true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"bfd": schema.SingleNestedAttribute{
														MarkdownDescription: "Bfd",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"multihop": schema.SingleNestedAttribute{
																MarkdownDescription: "Multihop",
																Optional:            true,
																Attributes: map[string]schema.Attribute{
																	"min_received_ttl": schema.Float64Attribute{
																		MarkdownDescription: "Min received ttl",
																		Optional:            true,
																	},
																},
															},
															"profile": schema.StringAttribute{
																MarkdownDescription: "Profile",
																Optional:            true,
															},
														},
													},
													"connection_options": schema.SingleNestedAttribute{
														MarkdownDescription: "Connection options",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"authentication": schema.StringAttribute{
																MarkdownDescription: "Authentication",
																Optional:            true,
															},
															"dampening": schema.StringAttribute{
																MarkdownDescription: "Dampening",
																Optional:            true,
															},
															"hold_time": schema.StringAttribute{
																MarkdownDescription: "Hold time",
																Optional:            true,
															},
															"idle_hold_time": schema.Float64Attribute{
																MarkdownDescription: "Idle hold time",
																Optional:            true,
															},
															"incoming_bgp_connection": schema.SingleNestedAttribute{
																MarkdownDescription: "Incoming bgp connection",
																Optional:            true,
																Attributes: map[string]schema.Attribute{
																	"allow": schema.BoolAttribute{
																		MarkdownDescription: "Allow",
																		Optional:            true,
																	},
																	"remote_port": schema.Float64Attribute{
																		MarkdownDescription: "Remote port",
																		Optional:            true,
																	},
																},
															},
															"keep_alive_interval": schema.StringAttribute{
																MarkdownDescription: "Keep alive interval",
																Optional:            true,
															},
															"max_prefixes": schema.StringAttribute{
																MarkdownDescription: "Max prefixes",
																Optional:            true,
															},
															"min_route_adv_interval": schema.Float64Attribute{
																MarkdownDescription: "Min route adv interval",
																Optional:            true,
															},
															"multihop": schema.StringAttribute{
																MarkdownDescription: "Multihop",
																Optional:            true,
															},
															"open_delay_time": schema.Float64Attribute{
																MarkdownDescription: "Open delay time",
																Optional:            true,
															},
															"outgoing_bgp_connection": schema.SingleNestedAttribute{
																MarkdownDescription: "Outgoing bgp connection",
																Optional:            true,
																Attributes: map[string]schema.Attribute{
																	"allow": schema.BoolAttribute{
																		MarkdownDescription: "Allow",
																		Optional:            true,
																	},
																	"local_port": schema.Float64Attribute{
																		MarkdownDescription: "Local port",
																		Optional:            true,
																	},
																},
															},
															"timers": schema.StringAttribute{
																MarkdownDescription: "Timers",
																Optional:            true,
															},
														},
													},
													"enable": schema.BoolAttribute{
														MarkdownDescription: "Enable",
														Optional:            true,
													},
													"enable_mp_bgp": schema.BoolAttribute{
														MarkdownDescription: "Enable mp bgp",
														Optional:            true,
													},
													"enable_sender_side_loop_detection": schema.BoolAttribute{
														MarkdownDescription: "Enable sender side loop detection",
														Optional:            true,
													},
													"inherit": schema.SingleNestedAttribute{
														MarkdownDescription: "Inherit",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"no": schema.SingleNestedAttribute{
																Validators: []validator.Object{
																	objectvalidator.ExactlyOneOf(
																		path.MatchRelative().AtParent().AtName("ipv4"),
																		path.MatchRelative().AtParent().AtName("yes"),
																	),
																},
																MarkdownDescription: "No",
																Optional:            true,
																Attributes: map[string]schema.Attribute{
																	"address_family": schema.SingleNestedAttribute{
																		MarkdownDescription: "Address family",
																		Optional:            true,
																		Attributes: map[string]schema.Attribute{
																			"ipv4": schema.StringAttribute{
																				MarkdownDescription: "Ipv4",
																				Optional:            true,
																			},
																			"ipv6": schema.StringAttribute{
																				MarkdownDescription: "Ipv6",
																				Optional:            true,
																			},
																		},
																	},
																	"filtering_profile": schema.SingleNestedAttribute{
																		MarkdownDescription: "Filtering profile",
																		Optional:            true,
																		Attributes: map[string]schema.Attribute{
																			"ipv4": schema.StringAttribute{
																				MarkdownDescription: "Ipv4",
																				Optional:            true,
																			},
																			"ipv6": schema.StringAttribute{
																				MarkdownDescription: "Ipv6",
																				Optional:            true,
																			},
																		},
																	},
																},
															},
															"yes": schema.SingleNestedAttribute{
																Validators: []validator.Object{
																	objectvalidator.ExactlyOneOf(
																		path.MatchRelative().AtParent().AtName("ipv4"),
																		path.MatchRelative().AtParent().AtName("no"),
																	),
																},
																MarkdownDescription: "Yes",
																Optional:            true,
																Attributes:          map[string]schema.Attribute{},
															},
														},
													},
													"local_address": schema.SingleNestedAttribute{
														MarkdownDescription: "Local address",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"interface": schema.StringAttribute{
																MarkdownDescription: "Interface",
																Optional:            true,
															},
															"ip": schema.StringAttribute{
																MarkdownDescription: "Ip",
																Optional:            true,
															},
														},
													},
													"name": schema.StringAttribute{
														MarkdownDescription: "Name",
														Required:            true,
													},
													"passive": schema.BoolAttribute{
														MarkdownDescription: "Passive",
														Optional:            true,
													},
													"peer_address": schema.SingleNestedAttribute{
														MarkdownDescription: "Peer address",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"fqdn": schema.StringAttribute{
																Validators: []validator.String{
																	stringvalidator.ExactlyOneOf(
																		path.MatchRelative().AtParent().AtName("ip"),
																	),
																},
																MarkdownDescription: "Fqdn",
																Optional:            true,
															},
															"ip": schema.StringAttribute{
																Validators: []validator.String{
																	stringvalidator.ExactlyOneOf(
																		path.MatchRelative().AtParent().AtName("fqdn"),
																	),
																},
																MarkdownDescription: "Ip",
																Optional:            true,
															},
														},
													},
													"peer_as": schema.StringAttribute{
														MarkdownDescription: "Peer as",
														Optional:            true,
													},
													"peering_type": schema.StringAttribute{
														MarkdownDescription: "Peering type",
														Optional:            true,
													},
													"reflector_client": schema.StringAttribute{
														MarkdownDescription: "Reflector client",
														Optional:            true,
													},
													"subsequent_address_family_identifier": schema.SingleNestedAttribute{
														MarkdownDescription: "Subsequent address family identifier",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"multicast": schema.BoolAttribute{
																MarkdownDescription: "Multicast",
																Optional:            true,
															},
															"unicast": schema.BoolAttribute{
																MarkdownDescription: "Unicast",
																Optional:            true,
															},
														},
													},
												},
											},
										},
										"soft_reset_with_stored_info": schema.BoolAttribute{
											MarkdownDescription: "Soft reset with stored info",
											Optional:            true,
										},
										"type": schema.SingleNestedAttribute{
											MarkdownDescription: "Type",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"ebgp": schema.SingleNestedAttribute{
													Validators: []validator.Object{
														objectvalidator.ExactlyOneOf(
															path.MatchRelative().AtParent().AtName("ebgp_confed"),
															path.MatchRelative().AtParent().AtName("ibgp"),
															path.MatchRelative().AtParent().AtName("ibgp_confed"),
														),
													},
													MarkdownDescription: "Ebgp",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"export_nexthop": schema.StringAttribute{
															MarkdownDescription: "Export nexthop",
															Optional:            true,
														},
														"import_nexthop": schema.StringAttribute{
															MarkdownDescription: "Import nexthop",
															Optional:            true,
														},
														"remove_private_as": schema.BoolAttribute{
															MarkdownDescription: "Remove private as",
															Optional:            true,
														},
													},
												},
												"ebgp_confed": schema.SingleNestedAttribute{
													Validators: []validator.Object{
														objectvalidator.ExactlyOneOf(
															path.MatchRelative().AtParent().AtName("ebgp"),
															path.MatchRelative().AtParent().AtName("ibgp"),
															path.MatchRelative().AtParent().AtName("ibgp_confed"),
														),
													},
													MarkdownDescription: "Ebgp confed",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"export_nexthop": schema.StringAttribute{
															MarkdownDescription: "Export nexthop",
															Optional:            true,
														},
													},
												},
												"ibgp": schema.SingleNestedAttribute{
													Validators: []validator.Object{
														objectvalidator.ExactlyOneOf(
															path.MatchRelative().AtParent().AtName("ebgp"),
															path.MatchRelative().AtParent().AtName("ebgp_confed"),
															path.MatchRelative().AtParent().AtName("ibgp_confed"),
														),
													},
													MarkdownDescription: "Ibgp",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"export_nexthop": schema.StringAttribute{
															MarkdownDescription: "Export nexthop",
															Optional:            true,
														},
													},
												},
												"ibgp_confed": schema.SingleNestedAttribute{
													Validators: []validator.Object{
														objectvalidator.ExactlyOneOf(
															path.MatchRelative().AtParent().AtName("ebgp"),
															path.MatchRelative().AtParent().AtName("ebgp_confed"),
															path.MatchRelative().AtParent().AtName("ibgp"),
														),
													},
													MarkdownDescription: "Ibgp confed",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"export_nexthop": schema.StringAttribute{
															MarkdownDescription: "Export nexthop",
															Optional:            true,
														},
													},
												},
											},
										},
									},
								},
							},
							"policy": schema.SingleNestedAttribute{
								MarkdownDescription: "Policy",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"aggregation": schema.SingleNestedAttribute{
										MarkdownDescription: "Aggregation",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"address": schema.ListNestedAttribute{
												MarkdownDescription: "Address",
												Optional:            true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"advertise_filters": schema.ListNestedAttribute{
															MarkdownDescription: "Advertise filters",
															Optional:            true,
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"enable": schema.BoolAttribute{
																		MarkdownDescription: "Enable",
																		Optional:            true,
																	},
																	"match": schema.SingleNestedAttribute{
																		MarkdownDescription: "Match",
																		Optional:            true,
																		Attributes: map[string]schema.Attribute{
																			"address_prefix": schema.ListNestedAttribute{
																				MarkdownDescription: "Address prefix",
																				Optional:            true,
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"exact": schema.BoolAttribute{
																							MarkdownDescription: "Exact",
																							Optional:            true,
																						},
																						"name": schema.StringAttribute{
																							MarkdownDescription: "Name",
																							Required:            true,
																						},
																					},
																				},
																			},
																			"afi": schema.StringAttribute{
																				Validators: []validator.String{
																					stringvalidator.OneOf("ip", "ipv6"),
																				},
																				MarkdownDescription: "Afi",
																				Optional:            true,
																			},
																			"as_path": schema.SingleNestedAttribute{
																				MarkdownDescription: "As path",
																				Optional:            true,
																				Attributes: map[string]schema.Attribute{
																					"regex": schema.StringAttribute{
																						MarkdownDescription: "Regex",
																						Optional:            true,
																					},
																				},
																			},
																			"community": schema.SingleNestedAttribute{
																				MarkdownDescription: "Community",
																				Optional:            true,
																				Attributes: map[string]schema.Attribute{
																					"regex": schema.StringAttribute{
																						MarkdownDescription: "Regex",
																						Optional:            true,
																					},
																				},
																			},
																			"extended_community": schema.SingleNestedAttribute{
																				MarkdownDescription: "Extended community",
																				Optional:            true,
																				Attributes: map[string]schema.Attribute{
																					"regex": schema.StringAttribute{
																						MarkdownDescription: "Regex",
																						Optional:            true,
																					},
																				},
																			},
																			"from_peer": schema.ListAttribute{
																				ElementType:         types.StringType,
																				MarkdownDescription: "From peer",
																				Optional:            true,
																			},
																			"med": schema.Float64Attribute{
																				MarkdownDescription: "Med",
																				Optional:            true,
																			},
																			"nexthop": schema.ListAttribute{
																				ElementType:         types.StringType,
																				MarkdownDescription: "Nexthop",
																				Optional:            true,
																			},
																			"route_table": schema.StringAttribute{
																				Validators: []validator.String{
																					stringvalidator.OneOf("unicast", "multicast", "both"),
																				},
																				MarkdownDescription: "Route table",
																				Optional:            true,
																			},
																			"safi": schema.StringAttribute{
																				Validators: []validator.String{
																					stringvalidator.OneOf("ip", "ipv6"),
																				},
																				MarkdownDescription: "Safi",
																				Optional:            true,
																			},
																		},
																	},
																	"name": schema.StringAttribute{
																		MarkdownDescription: "Name",
																		Required:            true,
																	},
																},
															},
														},
														"aggregate_route_attributes": schema.SingleNestedAttribute{
															MarkdownDescription: "Aggregate route attributes",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"as_path": schema.SingleNestedAttribute{
																	MarkdownDescription: "As path",
																	Optional:            true,
																	Attributes: map[string]schema.Attribute{
																		"none": schema.SingleNestedAttribute{
																			Validators: []validator.Object{
																				objectvalidator.ExactlyOneOf(
																					path.MatchRelative().AtParent().AtName("prepend"),
																					path.MatchRelative().AtParent().AtName("remove"),
																					path.MatchRelative().AtParent().AtName("remove_and_prepend"),
																				),
																			},
																			MarkdownDescription: "None",
																			Optional:            true,
																			Attributes:          map[string]schema.Attribute{},
																		},
																		"prepend": schema.Float64Attribute{
																			MarkdownDescription: "Prepend",
																			Optional:            true,
																		},
																		"remove": schema.SingleNestedAttribute{
																			Validators: []validator.Object{
																				objectvalidator.ExactlyOneOf(
																					path.MatchRelative().AtParent().AtName("none"),
																					path.MatchRelative().AtParent().AtName("prepend"),
																					path.MatchRelative().AtParent().AtName("remove_and_prepend"),
																				),
																			},
																			MarkdownDescription: "Remove",
																			Optional:            true,
																			Attributes:          map[string]schema.Attribute{},
																		},
																		"remove_and_prepend": schema.Float64Attribute{
																			MarkdownDescription: "Remove and prepend",
																			Optional:            true,
																		},
																	},
																},
																"as_path_limit": schema.Float64Attribute{
																	MarkdownDescription: "As path limit",
																	Optional:            true,
																},
																"community": schema.SingleNestedAttribute{
																	MarkdownDescription: "Community",
																	Optional:            true,
																	Attributes: map[string]schema.Attribute{
																		"append": schema.ListAttribute{
																			ElementType:         types.StringType,
																			MarkdownDescription: "Append",
																			Validators: []validator.List{
																				listvalidator.ExactlyOneOf(
																					path.MatchRelative().AtParent().AtName("none"),
																					path.MatchRelative().AtParent().AtName("overwrite"),
																					path.MatchRelative().AtParent().AtName("remove_all"),
																					path.MatchRelative().AtParent().AtName("remove_regex"),
																				),
																			},
																			Optional: true,
																		},
																		"none": schema.SingleNestedAttribute{
																			Validators: []validator.Object{
																				objectvalidator.ExactlyOneOf(
																					path.MatchRelative().AtParent().AtName("append"),
																					path.MatchRelative().AtParent().AtName("overwrite"),
																					path.MatchRelative().AtParent().AtName("remove_all"),
																					path.MatchRelative().AtParent().AtName("remove_regex"),
																				),
																			},
																			MarkdownDescription: "None",
																			Optional:            true,
																			Attributes:          map[string]schema.Attribute{},
																		},
																		"overwrite": schema.ListAttribute{
																			ElementType:         types.StringType,
																			MarkdownDescription: "Overwrite",
																			Validators: []validator.List{
																				listvalidator.ExactlyOneOf(
																					path.MatchRelative().AtParent().AtName("append"),
																					path.MatchRelative().AtParent().AtName("none"),
																					path.MatchRelative().AtParent().AtName("remove_all"),
																					path.MatchRelative().AtParent().AtName("remove_regex"),
																				),
																			},
																			Optional: true,
																		},
																		"remove_all": schema.SingleNestedAttribute{
																			Validators: []validator.Object{
																				objectvalidator.ExactlyOneOf(
																					path.MatchRelative().AtParent().AtName("append"),
																					path.MatchRelative().AtParent().AtName("none"),
																					path.MatchRelative().AtParent().AtName("overwrite"),
																					path.MatchRelative().AtParent().AtName("remove_regex"),
																				),
																			},
																			MarkdownDescription: "Remove all",
																			Optional:            true,
																			Attributes:          map[string]schema.Attribute{},
																		},
																		"remove_regex": schema.StringAttribute{
																			Validators: []validator.String{
																				stringvalidator.ExactlyOneOf(
																					path.MatchRelative().AtParent().AtName("append"),
																					path.MatchRelative().AtParent().AtName("none"),
																					path.MatchRelative().AtParent().AtName("overwrite"),
																					path.MatchRelative().AtParent().AtName("remove_all"),
																				),
																			},
																			MarkdownDescription: "Remove regex",
																			Optional:            true,
																		},
																	},
																},
																"extended_community": schema.SingleNestedAttribute{
																	MarkdownDescription: "Extended community",
																	Optional:            true,
																	Attributes: map[string]schema.Attribute{
																		"append": schema.ListAttribute{
																			ElementType:         types.StringType,
																			MarkdownDescription: "Append",
																			Validators: []validator.List{
																				listvalidator.ExactlyOneOf(
																					path.MatchRelative().AtParent().AtName("none"),
																					path.MatchRelative().AtParent().AtName("overwrite"),
																					path.MatchRelative().AtParent().AtName("remove_all"),
																					path.MatchRelative().AtParent().AtName("remove_regex"),
																				),
																			},
																			Optional: true,
																		},
																		"none": schema.SingleNestedAttribute{
																			Validators: []validator.Object{
																				objectvalidator.ExactlyOneOf(
																					path.MatchRelative().AtParent().AtName("append"),
																					path.MatchRelative().AtParent().AtName("overwrite"),
																					path.MatchRelative().AtParent().AtName("remove_all"),
																					path.MatchRelative().AtParent().AtName("remove_regex"),
																				),
																			},
																			MarkdownDescription: "None",
																			Optional:            true,
																			Attributes:          map[string]schema.Attribute{},
																		},
																		"overwrite": schema.ListAttribute{
																			ElementType:         types.StringType,
																			MarkdownDescription: "Overwrite",
																			Validators: []validator.List{
																				listvalidator.ExactlyOneOf(
																					path.MatchRelative().AtParent().AtName("append"),
																					path.MatchRelative().AtParent().AtName("none"),
																					path.MatchRelative().AtParent().AtName("remove_all"),
																					path.MatchRelative().AtParent().AtName("remove_regex"),
																				),
																			},
																			Optional: true,
																		},
																		"remove_all": schema.SingleNestedAttribute{
																			Validators: []validator.Object{
																				objectvalidator.ExactlyOneOf(
																					path.MatchRelative().AtParent().AtName("append"),
																					path.MatchRelative().AtParent().AtName("none"),
																					path.MatchRelative().AtParent().AtName("overwrite"),
																					path.MatchRelative().AtParent().AtName("remove_regex"),
																				),
																			},
																			MarkdownDescription: "Remove all",
																			Optional:            true,
																			Attributes:          map[string]schema.Attribute{},
																		},
																		"remove_regex": schema.StringAttribute{
																			Validators: []validator.String{
																				stringvalidator.ExactlyOneOf(
																					path.MatchRelative().AtParent().AtName("append"),
																					path.MatchRelative().AtParent().AtName("none"),
																					path.MatchRelative().AtParent().AtName("overwrite"),
																					path.MatchRelative().AtParent().AtName("remove_all"),
																				),
																			},
																			MarkdownDescription: "Remove regex",
																			Optional:            true,
																		},
																	},
																},
																"local_preference": schema.Float64Attribute{
																	MarkdownDescription: "Local preference",
																	Optional:            true,
																},
																"med": schema.Float64Attribute{
																	MarkdownDescription: "Med",
																	Optional:            true,
																},
																"nexthop": schema.StringAttribute{
																	MarkdownDescription: "Nexthop",
																	Optional:            true,
																},
																"origin": schema.StringAttribute{
																	Validators: []validator.String{
																		stringvalidator.OneOf("igp", "egp", "incomplete"),
																	},
																	MarkdownDescription: "Origin",
																	Optional:            true,
																},
																"weight": schema.Float64Attribute{
																	MarkdownDescription: "Weight",
																	Optional:            true,
																},
															},
														},
														"as_set": schema.BoolAttribute{
															MarkdownDescription: "As set",
															Optional:            true,
														},
														"enable": schema.BoolAttribute{
															MarkdownDescription: "Enable",
															Optional:            true,
														},
														"name": schema.StringAttribute{
															MarkdownDescription: "Name",
															Required:            true,
														},
														"prefix": schema.StringAttribute{
															MarkdownDescription: "Prefix",
															Optional:            true,
														},
														"summary": schema.BoolAttribute{
															MarkdownDescription: "Summary",
															Optional:            true,
														},
														"suppress_filters": schema.ListNestedAttribute{
															MarkdownDescription: "Suppress filters",
															Optional:            true,
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"enable": schema.BoolAttribute{
																		MarkdownDescription: "Enable",
																		Optional:            true,
																	},
																	"match": schema.SingleNestedAttribute{
																		MarkdownDescription: "Match",
																		Optional:            true,
																		Attributes: map[string]schema.Attribute{
																			"address_prefix": schema.ListNestedAttribute{
																				MarkdownDescription: "Address prefix",
																				Optional:            true,
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"exact": schema.BoolAttribute{
																							MarkdownDescription: "Exact",
																							Optional:            true,
																						},
																						"name": schema.StringAttribute{
																							MarkdownDescription: "Name",
																							Required:            true,
																						},
																					},
																				},
																			},
																			"afi": schema.StringAttribute{
																				Validators: []validator.String{
																					stringvalidator.OneOf("ip", "ipv6"),
																				},
																				MarkdownDescription: "Afi",
																				Optional:            true,
																			},
																			"as_path": schema.SingleNestedAttribute{
																				MarkdownDescription: "As path",
																				Optional:            true,
																				Attributes: map[string]schema.Attribute{
																					"regex": schema.StringAttribute{
																						MarkdownDescription: "Regex",
																						Optional:            true,
																					},
																				},
																			},
																			"community": schema.SingleNestedAttribute{
																				MarkdownDescription: "Community",
																				Optional:            true,
																				Attributes: map[string]schema.Attribute{
																					"regex": schema.StringAttribute{
																						MarkdownDescription: "Regex",
																						Optional:            true,
																					},
																				},
																			},
																			"extended_community": schema.SingleNestedAttribute{
																				MarkdownDescription: "Extended community",
																				Optional:            true,
																				Attributes: map[string]schema.Attribute{
																					"regex": schema.StringAttribute{
																						MarkdownDescription: "Regex",
																						Optional:            true,
																					},
																				},
																			},
																			"from_peer": schema.ListAttribute{
																				ElementType:         types.StringType,
																				MarkdownDescription: "From peer",
																				Optional:            true,
																			},
																			"med": schema.Float64Attribute{
																				MarkdownDescription: "Med",
																				Optional:            true,
																			},
																			"nexthop": schema.ListAttribute{
																				ElementType:         types.StringType,
																				MarkdownDescription: "Nexthop",
																				Optional:            true,
																			},
																			"route_table": schema.StringAttribute{
																				Validators: []validator.String{
																					stringvalidator.OneOf("unicast", "multicast", "both"),
																				},
																				MarkdownDescription: "Route table",
																				Optional:            true,
																			},
																			"safi": schema.StringAttribute{
																				Validators: []validator.String{
																					stringvalidator.OneOf("ip", "ipv6"),
																				},
																				MarkdownDescription: "Safi",
																				Optional:            true,
																			},
																		},
																	},
																	"name": schema.StringAttribute{
																		MarkdownDescription: "Name",
																		Required:            true,
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									"conditional_advertisement": schema.SingleNestedAttribute{
										MarkdownDescription: "Conditional advertisement",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"policy": schema.ListNestedAttribute{
												MarkdownDescription: "Policy",
												Optional:            true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"advertise_filters": schema.ListNestedAttribute{
															MarkdownDescription: "Advertise filters",
															Optional:            true,
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"enable": schema.BoolAttribute{
																		MarkdownDescription: "Enable",
																		Optional:            true,
																	},
																	"match": schema.SingleNestedAttribute{
																		MarkdownDescription: "Match",
																		Optional:            true,
																		Attributes: map[string]schema.Attribute{
																			"address_prefix": schema.ListNestedAttribute{
																				MarkdownDescription: "Address prefix",
																				Optional:            true,
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"exact": schema.BoolAttribute{
																							MarkdownDescription: "Exact",
																							Optional:            true,
																						},
																						"name": schema.StringAttribute{
																							MarkdownDescription: "Name",
																							Required:            true,
																						},
																					},
																				},
																			},
																			"afi": schema.StringAttribute{
																				Validators: []validator.String{
																					stringvalidator.OneOf("ip", "ipv6"),
																				},
																				MarkdownDescription: "Afi",
																				Optional:            true,
																			},
																			"as_path": schema.SingleNestedAttribute{
																				MarkdownDescription: "As path",
																				Optional:            true,
																				Attributes: map[string]schema.Attribute{
																					"regex": schema.StringAttribute{
																						MarkdownDescription: "Regex",
																						Optional:            true,
																					},
																				},
																			},
																			"community": schema.SingleNestedAttribute{
																				MarkdownDescription: "Community",
																				Optional:            true,
																				Attributes: map[string]schema.Attribute{
																					"regex": schema.StringAttribute{
																						MarkdownDescription: "Regex",
																						Optional:            true,
																					},
																				},
																			},
																			"extended_community": schema.SingleNestedAttribute{
																				MarkdownDescription: "Extended community",
																				Optional:            true,
																				Attributes: map[string]schema.Attribute{
																					"regex": schema.StringAttribute{
																						MarkdownDescription: "Regex",
																						Optional:            true,
																					},
																				},
																			},
																			"from_peer": schema.ListAttribute{
																				ElementType:         types.StringType,
																				MarkdownDescription: "From peer",
																				Optional:            true,
																			},
																			"med": schema.Float64Attribute{
																				MarkdownDescription: "Med",
																				Optional:            true,
																			},
																			"nexthop": schema.ListAttribute{
																				ElementType:         types.StringType,
																				MarkdownDescription: "Nexthop",
																				Optional:            true,
																			},
																			"route_table": schema.StringAttribute{
																				Validators: []validator.String{
																					stringvalidator.OneOf("unicast", "multicast", "both"),
																				},
																				MarkdownDescription: "Route table",
																				Optional:            true,
																			},
																			"safi": schema.StringAttribute{
																				Validators: []validator.String{
																					stringvalidator.OneOf("ip", "ipv6"),
																				},
																				MarkdownDescription: "Safi",
																				Optional:            true,
																			},
																		},
																	},
																	"name": schema.StringAttribute{
																		MarkdownDescription: "Name",
																		Required:            true,
																	},
																},
															},
														},
														"enable": schema.BoolAttribute{
															MarkdownDescription: "Enable",
															Optional:            true,
														},
														"name": schema.StringAttribute{
															MarkdownDescription: "Name",
															Required:            true,
														},
														"non_exist_filters": schema.ListNestedAttribute{
															MarkdownDescription: "Non exist filters",
															Optional:            true,
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"enable": schema.BoolAttribute{
																		MarkdownDescription: "Enable",
																		Optional:            true,
																	},
																	"match": schema.SingleNestedAttribute{
																		MarkdownDescription: "Match",
																		Optional:            true,
																		Attributes: map[string]schema.Attribute{
																			"address_prefix": schema.ListNestedAttribute{
																				MarkdownDescription: "Address prefix",
																				Optional:            true,
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"exact": schema.BoolAttribute{
																							MarkdownDescription: "Exact",
																							Optional:            true,
																						},
																						"name": schema.StringAttribute{
																							MarkdownDescription: "Name",
																							Required:            true,
																						},
																					},
																				},
																			},
																			"afi": schema.StringAttribute{
																				Validators: []validator.String{
																					stringvalidator.OneOf("ip", "ipv6"),
																				},
																				MarkdownDescription: "Afi",
																				Optional:            true,
																			},
																			"as_path": schema.SingleNestedAttribute{
																				MarkdownDescription: "As path",
																				Optional:            true,
																				Attributes: map[string]schema.Attribute{
																					"regex": schema.StringAttribute{
																						MarkdownDescription: "Regex",
																						Optional:            true,
																					},
																				},
																			},
																			"community": schema.SingleNestedAttribute{
																				MarkdownDescription: "Community",
																				Optional:            true,
																				Attributes: map[string]schema.Attribute{
																					"regex": schema.StringAttribute{
																						MarkdownDescription: "Regex",
																						Optional:            true,
																					},
																				},
																			},
																			"extended_community": schema.SingleNestedAttribute{
																				MarkdownDescription: "Extended community",
																				Optional:            true,
																				Attributes: map[string]schema.Attribute{
																					"regex": schema.StringAttribute{
																						MarkdownDescription: "Regex",
																						Optional:            true,
																					},
																				},
																			},
																			"from_peer": schema.ListAttribute{
																				ElementType:         types.StringType,
																				MarkdownDescription: "From peer",
																				Optional:            true,
																			},
																			"med": schema.Float64Attribute{
																				MarkdownDescription: "Med",
																				Optional:            true,
																			},
																			"nexthop": schema.ListAttribute{
																				ElementType:         types.StringType,
																				MarkdownDescription: "Nexthop",
																				Optional:            true,
																			},
																			"route_table": schema.StringAttribute{
																				Validators: []validator.String{
																					stringvalidator.OneOf("unicast", "multicast", "both"),
																				},
																				MarkdownDescription: "Route table",
																				Optional:            true,
																			},
																			"safi": schema.StringAttribute{
																				Validators: []validator.String{
																					stringvalidator.OneOf("ip", "ipv6"),
																				},
																				MarkdownDescription: "Safi",
																				Optional:            true,
																			},
																		},
																	},
																	"name": schema.StringAttribute{
																		MarkdownDescription: "Name",
																		Required:            true,
																	},
																},
															},
														},
														"used_by": schema.ListAttribute{
															ElementType:         types.StringType,
															MarkdownDescription: "Used by",
															Optional:            true,
														},
													},
												},
											},
										},
									},
									"export": schema.SingleNestedAttribute{
										MarkdownDescription: "Export",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"rules": schema.ListNestedAttribute{
												MarkdownDescription: "Rules",
												Optional:            true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"action": schema.SingleNestedAttribute{
															MarkdownDescription: "Action",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"allow": schema.SingleNestedAttribute{
																	Validators: []validator.Object{
																		objectvalidator.ExactlyOneOf(
																			path.MatchRelative().AtParent().AtName("deny"),
																		),
																	},
																	MarkdownDescription: "Allow",
																	Optional:            true,
																	Attributes: map[string]schema.Attribute{
																		"update": schema.SingleNestedAttribute{
																			MarkdownDescription: "Update",
																			Optional:            true,
																			Attributes: map[string]schema.Attribute{
																				"as_path": schema.SingleNestedAttribute{
																					MarkdownDescription: "As path",
																					Optional:            true,
																					Attributes: map[string]schema.Attribute{
																						"none": schema.SingleNestedAttribute{
																							Validators: []validator.Object{
																								objectvalidator.ExactlyOneOf(
																									path.MatchRelative().AtParent().AtName("prepend"),
																									path.MatchRelative().AtParent().AtName("remove"),
																									path.MatchRelative().AtParent().AtName("remove_and_prepend"),
																								),
																							},
																							MarkdownDescription: "None",
																							Optional:            true,
																							Attributes:          map[string]schema.Attribute{},
																						},
																						"prepend": schema.Float64Attribute{
																							MarkdownDescription: "Prepend",
																							Optional:            true,
																						},
																						"remove": schema.SingleNestedAttribute{
																							Validators: []validator.Object{
																								objectvalidator.ExactlyOneOf(
																									path.MatchRelative().AtParent().AtName("none"),
																									path.MatchRelative().AtParent().AtName("prepend"),
																									path.MatchRelative().AtParent().AtName("remove_and_prepend"),
																								),
																							},
																							MarkdownDescription: "Remove",
																							Optional:            true,
																							Attributes:          map[string]schema.Attribute{},
																						},
																						"remove_and_prepend": schema.Float64Attribute{
																							MarkdownDescription: "Remove and prepend",
																							Optional:            true,
																						},
																					},
																				},
																				"as_path_limit": schema.Float64Attribute{
																					MarkdownDescription: "As path limit",
																					Optional:            true,
																				},
																				"community": schema.SingleNestedAttribute{
																					MarkdownDescription: "Community",
																					Optional:            true,
																					Attributes: map[string]schema.Attribute{
																						"append": schema.ListAttribute{
																							ElementType:         types.StringType,
																							MarkdownDescription: "Append",
																							Validators: []validator.List{
																								listvalidator.ExactlyOneOf(
																									path.MatchRelative().AtParent().AtName("none"),
																									path.MatchRelative().AtParent().AtName("overwrite"),
																									path.MatchRelative().AtParent().AtName("remove_all"),
																									path.MatchRelative().AtParent().AtName("remove_regex"),
																								),
																							},
																							Optional: true,
																						},
																						"none": schema.SingleNestedAttribute{
																							Validators: []validator.Object{
																								objectvalidator.ExactlyOneOf(
																									path.MatchRelative().AtParent().AtName("append"),
																									path.MatchRelative().AtParent().AtName("overwrite"),
																									path.MatchRelative().AtParent().AtName("remove_all"),
																									path.MatchRelative().AtParent().AtName("remove_regex"),
																								),
																							},
																							MarkdownDescription: "None",
																							Optional:            true,
																							Attributes:          map[string]schema.Attribute{},
																						},
																						"overwrite": schema.ListAttribute{
																							ElementType:         types.StringType,
																							MarkdownDescription: "Overwrite",
																							Validators: []validator.List{
																								listvalidator.ExactlyOneOf(
																									path.MatchRelative().AtParent().AtName("append"),
																									path.MatchRelative().AtParent().AtName("none"),
																									path.MatchRelative().AtParent().AtName("remove_all"),
																									path.MatchRelative().AtParent().AtName("remove_regex"),
																								),
																							},
																							Optional: true,
																						},
																						"remove_all": schema.SingleNestedAttribute{
																							Validators: []validator.Object{
																								objectvalidator.ExactlyOneOf(
																									path.MatchRelative().AtParent().AtName("append"),
																									path.MatchRelative().AtParent().AtName("none"),
																									path.MatchRelative().AtParent().AtName("overwrite"),
																									path.MatchRelative().AtParent().AtName("remove_regex"),
																								),
																							},
																							MarkdownDescription: "Remove all",
																							Optional:            true,
																							Attributes:          map[string]schema.Attribute{},
																						},
																						"remove_regex": schema.StringAttribute{
																							Validators: []validator.String{
																								stringvalidator.ExactlyOneOf(
																									path.MatchRelative().AtParent().AtName("append"),
																									path.MatchRelative().AtParent().AtName("none"),
																									path.MatchRelative().AtParent().AtName("overwrite"),
																									path.MatchRelative().AtParent().AtName("remove_all"),
																								),
																							},
																							MarkdownDescription: "Remove regex",
																							Optional:            true,
																						},
																					},
																				},
																				"extended_community": schema.SingleNestedAttribute{
																					MarkdownDescription: "Extended community",
																					Optional:            true,
																					Attributes: map[string]schema.Attribute{
																						"append": schema.ListAttribute{
																							ElementType:         types.StringType,
																							MarkdownDescription: "Append",
																							Validators: []validator.List{
																								listvalidator.ExactlyOneOf(
																									path.MatchRelative().AtParent().AtName("none"),
																									path.MatchRelative().AtParent().AtName("overwrite"),
																									path.MatchRelative().AtParent().AtName("remove_all"),
																									path.MatchRelative().AtParent().AtName("remove_regex"),
																								),
																							},
																							Optional: true,
																						},
																						"none": schema.SingleNestedAttribute{
																							Validators: []validator.Object{
																								objectvalidator.ExactlyOneOf(
																									path.MatchRelative().AtParent().AtName("append"),
																									path.MatchRelative().AtParent().AtName("overwrite"),
																									path.MatchRelative().AtParent().AtName("remove_all"),
																									path.MatchRelative().AtParent().AtName("remove_regex"),
																								),
																							},
																							MarkdownDescription: "None",
																							Optional:            true,
																							Attributes:          map[string]schema.Attribute{},
																						},
																						"overwrite": schema.ListAttribute{
																							ElementType:         types.StringType,
																							MarkdownDescription: "Overwrite",
																							Validators: []validator.List{
																								listvalidator.ExactlyOneOf(
																									path.MatchRelative().AtParent().AtName("append"),
																									path.MatchRelative().AtParent().AtName("none"),
																									path.MatchRelative().AtParent().AtName("remove_all"),
																									path.MatchRelative().AtParent().AtName("remove_regex"),
																								),
																							},
																							Optional: true,
																						},
																						"remove_all": schema.SingleNestedAttribute{
																							Validators: []validator.Object{
																								objectvalidator.ExactlyOneOf(
																									path.MatchRelative().AtParent().AtName("append"),
																									path.MatchRelative().AtParent().AtName("none"),
																									path.MatchRelative().AtParent().AtName("overwrite"),
																									path.MatchRelative().AtParent().AtName("remove_regex"),
																								),
																							},
																							MarkdownDescription: "Remove all",
																							Optional:            true,
																							Attributes:          map[string]schema.Attribute{},
																						},
																						"remove_regex": schema.StringAttribute{
																							Validators: []validator.String{
																								stringvalidator.ExactlyOneOf(
																									path.MatchRelative().AtParent().AtName("append"),
																									path.MatchRelative().AtParent().AtName("none"),
																									path.MatchRelative().AtParent().AtName("overwrite"),
																									path.MatchRelative().AtParent().AtName("remove_all"),
																								),
																							},
																							MarkdownDescription: "Remove regex",
																							Optional:            true,
																						},
																					},
																				},
																				"local_preference": schema.Float64Attribute{
																					MarkdownDescription: "Local preference",
																					Optional:            true,
																				},
																				"med": schema.Float64Attribute{
																					MarkdownDescription: "Med",
																					Optional:            true,
																				},
																				"nexthop": schema.StringAttribute{
																					MarkdownDescription: "Nexthop",
																					Optional:            true,
																				},
																				"origin": schema.StringAttribute{
																					Validators: []validator.String{
																						stringvalidator.OneOf("igp", "egp", "multicast"),
																					},
																					MarkdownDescription: "Origin",
																					Optional:            true,
																				},
																			},
																		},
																	},
																},
																"deny": schema.SingleNestedAttribute{
																	Validators: []validator.Object{
																		objectvalidator.ExactlyOneOf(
																			path.MatchRelative().AtParent().AtName("allow"),
																		),
																	},
																	MarkdownDescription: "Deny",
																	Optional:            true,
																	Attributes:          map[string]schema.Attribute{},
																},
															},
														},
														"enable": schema.BoolAttribute{
															MarkdownDescription: "Enable",
															Optional:            true,
														},
														"match": schema.SingleNestedAttribute{
															MarkdownDescription: "Match",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"address_prefix": schema.ListNestedAttribute{
																	MarkdownDescription: "Address prefix",
																	Optional:            true,
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"exact": schema.BoolAttribute{
																				MarkdownDescription: "Exact",
																				Optional:            true,
																			},
																			"name": schema.StringAttribute{
																				MarkdownDescription: "Name",
																				Optional:            true,
																			},
																		},
																	},
																},
																"afi": schema.StringAttribute{
																	Validators: []validator.String{
																		stringvalidator.OneOf("ip", "ipv6"),
																	},
																	MarkdownDescription: "Afi",
																	Optional:            true,
																},
																"as_path": schema.SingleNestedAttribute{
																	MarkdownDescription: "As path",
																	Optional:            true,
																	Attributes: map[string]schema.Attribute{
																		"regex": schema.StringAttribute{
																			MarkdownDescription: "Regex",
																			Optional:            true,
																		},
																	},
																},
																"community": schema.SingleNestedAttribute{
																	MarkdownDescription: "Community",
																	Optional:            true,
																	Attributes: map[string]schema.Attribute{
																		"regex": schema.StringAttribute{
																			MarkdownDescription: "Regex",
																			Optional:            true,
																		},
																	},
																},
																"extended_community": schema.SingleNestedAttribute{
																	MarkdownDescription: "Extended community",
																	Optional:            true,
																	Attributes: map[string]schema.Attribute{
																		"regex": schema.StringAttribute{
																			MarkdownDescription: "Regex",
																			Optional:            true,
																		},
																	},
																},
																"from_peer": schema.ListAttribute{
																	ElementType:         types.StringType,
																	MarkdownDescription: "From peer",
																	Optional:            true,
																},
																"med": schema.Float64Attribute{
																	MarkdownDescription: "Med",
																	Optional:            true,
																},
																"nexthop": schema.ListAttribute{
																	ElementType:         types.StringType,
																	MarkdownDescription: "Nexthop",
																	Optional:            true,
																},
																"route_table": schema.StringAttribute{
																	Validators: []validator.String{
																		stringvalidator.OneOf("unicast", "multicast", "both"),
																	},
																	MarkdownDescription: "Route table",
																	Optional:            true,
																},
																"safi": schema.StringAttribute{
																	Validators: []validator.String{
																		stringvalidator.OneOf("ip", "ipv6"),
																	},
																	MarkdownDescription: "Safi",
																	Optional:            true,
																},
															},
														},
														"name": schema.StringAttribute{
															MarkdownDescription: "Name",
															Required:            true,
														},
														"used_by": schema.ListAttribute{
															ElementType:         types.StringType,
															MarkdownDescription: "Used by",
															Optional:            true,
														},
													},
												},
											},
										},
									},
									"import": schema.SingleNestedAttribute{
										MarkdownDescription: "Import",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"rules": schema.ListNestedAttribute{
												MarkdownDescription: "Rules",
												Optional:            true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"action": schema.SingleNestedAttribute{
															MarkdownDescription: "Action",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"allow": schema.SingleNestedAttribute{
																	Validators: []validator.Object{
																		objectvalidator.ExactlyOneOf(
																			path.MatchRelative().AtParent().AtName("deny"),
																		),
																	},
																	MarkdownDescription: "Allow",
																	Optional:            true,
																	Attributes: map[string]schema.Attribute{
																		"dampening": schema.StringAttribute{
																			MarkdownDescription: "Dampening",
																			Optional:            true,
																		},
																		"update": schema.SingleNestedAttribute{
																			MarkdownDescription: "Update",
																			Optional:            true,
																			Attributes: map[string]schema.Attribute{
																				"as_path": schema.SingleNestedAttribute{
																					MarkdownDescription: "As path",
																					Optional:            true,
																					Attributes: map[string]schema.Attribute{
																						"none": schema.SingleNestedAttribute{
																							Validators: []validator.Object{
																								objectvalidator.ExactlyOneOf(
																									path.MatchRelative().AtParent().AtName("prepend"),
																									path.MatchRelative().AtParent().AtName("remove"),
																									path.MatchRelative().AtParent().AtName("remove_and_prepend"),
																								),
																							},
																							MarkdownDescription: "None",
																							Optional:            true,
																							Attributes:          map[string]schema.Attribute{},
																						},
																						"prepend": schema.Float64Attribute{
																							MarkdownDescription: "Prepend",
																							Optional:            true,
																						},
																						"remove": schema.SingleNestedAttribute{
																							Validators: []validator.Object{
																								objectvalidator.ExactlyOneOf(
																									path.MatchRelative().AtParent().AtName("none"),
																									path.MatchRelative().AtParent().AtName("prepend"),
																									path.MatchRelative().AtParent().AtName("remove_and_prepend"),
																								),
																							},
																							MarkdownDescription: "Remove",
																							Optional:            true,
																							Attributes:          map[string]schema.Attribute{},
																						},
																						"remove_and_prepend": schema.Float64Attribute{
																							MarkdownDescription: "Remove and prepend",
																							Optional:            true,
																						},
																					},
																				},
																				"as_path_limit": schema.Float64Attribute{
																					MarkdownDescription: "As path limit",
																					Optional:            true,
																				},
																				"community": schema.SingleNestedAttribute{
																					MarkdownDescription: "Community",
																					Optional:            true,
																					Attributes: map[string]schema.Attribute{
																						"append": schema.ListAttribute{
																							ElementType:         types.StringType,
																							MarkdownDescription: "Append",
																							Validators: []validator.List{
																								listvalidator.ExactlyOneOf(
																									path.MatchRelative().AtParent().AtName("none"),
																									path.MatchRelative().AtParent().AtName("overwrite"),
																									path.MatchRelative().AtParent().AtName("remove_all"),
																									path.MatchRelative().AtParent().AtName("remove_regex"),
																								),
																							},
																							Optional: true,
																						},
																						"none": schema.SingleNestedAttribute{
																							Validators: []validator.Object{
																								objectvalidator.ExactlyOneOf(
																									path.MatchRelative().AtParent().AtName("append"),
																									path.MatchRelative().AtParent().AtName("overwrite"),
																									path.MatchRelative().AtParent().AtName("remove_all"),
																									path.MatchRelative().AtParent().AtName("remove_regex"),
																								),
																							},
																							MarkdownDescription: "None",
																							Optional:            true,
																							Attributes:          map[string]schema.Attribute{},
																						},
																						"overwrite": schema.ListAttribute{
																							ElementType:         types.StringType,
																							MarkdownDescription: "Overwrite",
																							Validators: []validator.List{
																								listvalidator.ExactlyOneOf(
																									path.MatchRelative().AtParent().AtName("append"),
																									path.MatchRelative().AtParent().AtName("none"),
																									path.MatchRelative().AtParent().AtName("remove_all"),
																									path.MatchRelative().AtParent().AtName("remove_regex"),
																								),
																							},
																							Optional: true,
																						},
																						"remove_all": schema.SingleNestedAttribute{
																							Validators: []validator.Object{
																								objectvalidator.ExactlyOneOf(
																									path.MatchRelative().AtParent().AtName("append"),
																									path.MatchRelative().AtParent().AtName("none"),
																									path.MatchRelative().AtParent().AtName("overwrite"),
																									path.MatchRelative().AtParent().AtName("remove_regex"),
																								),
																							},
																							MarkdownDescription: "Remove all",
																							Optional:            true,
																							Attributes:          map[string]schema.Attribute{},
																						},
																						"remove_regex": schema.StringAttribute{
																							Validators: []validator.String{
																								stringvalidator.ExactlyOneOf(
																									path.MatchRelative().AtParent().AtName("append"),
																									path.MatchRelative().AtParent().AtName("none"),
																									path.MatchRelative().AtParent().AtName("overwrite"),
																									path.MatchRelative().AtParent().AtName("remove_all"),
																								),
																							},
																							MarkdownDescription: "Remove regex",
																							Optional:            true,
																						},
																					},
																				},
																				"extended_community": schema.SingleNestedAttribute{
																					MarkdownDescription: "Extended community",
																					Optional:            true,
																					Attributes: map[string]schema.Attribute{
																						"append": schema.ListAttribute{
																							ElementType:         types.StringType,
																							MarkdownDescription: "Append",
																							Validators: []validator.List{
																								listvalidator.ExactlyOneOf(
																									path.MatchRelative().AtParent().AtName("none"),
																									path.MatchRelative().AtParent().AtName("overwrite"),
																									path.MatchRelative().AtParent().AtName("remove_all"),
																									path.MatchRelative().AtParent().AtName("remove_regex"),
																								),
																							},
																							Optional: true,
																						},
																						"none": schema.SingleNestedAttribute{
																							Validators: []validator.Object{
																								objectvalidator.ExactlyOneOf(
																									path.MatchRelative().AtParent().AtName("append"),
																									path.MatchRelative().AtParent().AtName("overwrite"),
																									path.MatchRelative().AtParent().AtName("remove_all"),
																									path.MatchRelative().AtParent().AtName("remove_regex"),
																								),
																							},
																							MarkdownDescription: "None",
																							Optional:            true,
																							Attributes:          map[string]schema.Attribute{},
																						},
																						"overwrite": schema.ListAttribute{
																							ElementType:         types.StringType,
																							MarkdownDescription: "Overwrite",
																							Validators: []validator.List{
																								listvalidator.ExactlyOneOf(
																									path.MatchRelative().AtParent().AtName("append"),
																									path.MatchRelative().AtParent().AtName("none"),
																									path.MatchRelative().AtParent().AtName("remove_all"),
																									path.MatchRelative().AtParent().AtName("remove_regex"),
																								),
																							},
																							Optional: true,
																						},
																						"remove_all": schema.SingleNestedAttribute{
																							Validators: []validator.Object{
																								objectvalidator.ExactlyOneOf(
																									path.MatchRelative().AtParent().AtName("append"),
																									path.MatchRelative().AtParent().AtName("none"),
																									path.MatchRelative().AtParent().AtName("overwrite"),
																									path.MatchRelative().AtParent().AtName("remove_regex"),
																								),
																							},
																							MarkdownDescription: "Remove all",
																							Optional:            true,
																							Attributes:          map[string]schema.Attribute{},
																						},
																						"remove_regex": schema.StringAttribute{
																							Validators: []validator.String{
																								stringvalidator.ExactlyOneOf(
																									path.MatchRelative().AtParent().AtName("append"),
																									path.MatchRelative().AtParent().AtName("none"),
																									path.MatchRelative().AtParent().AtName("overwrite"),
																									path.MatchRelative().AtParent().AtName("remove_all"),
																								),
																							},
																							MarkdownDescription: "Remove regex",
																							Optional:            true,
																						},
																					},
																				},
																				"local_preference": schema.Float64Attribute{
																					MarkdownDescription: "Local preference",
																					Optional:            true,
																				},
																				"med": schema.Float64Attribute{
																					MarkdownDescription: "Med",
																					Optional:            true,
																				},
																				"nexthop": schema.StringAttribute{
																					MarkdownDescription: "Nexthop",
																					Optional:            true,
																				},
																				"origin": schema.StringAttribute{
																					Validators: []validator.String{
																						stringvalidator.OneOf("igp", "egp", "incomplete"),
																					},
																					MarkdownDescription: "Origin",
																					Optional:            true,
																				},
																				"weight": schema.Float64Attribute{
																					MarkdownDescription: "Weight",
																					Optional:            true,
																				},
																			},
																		},
																	},
																},
																"deny": schema.SingleNestedAttribute{
																	Validators: []validator.Object{
																		objectvalidator.ExactlyOneOf(
																			path.MatchRelative().AtParent().AtName("allow"),
																		),
																	},
																	MarkdownDescription: "Deny",
																	Optional:            true,
																	Attributes:          map[string]schema.Attribute{},
																},
															},
														},
														"enable": schema.BoolAttribute{
															MarkdownDescription: "Enable",
															Optional:            true,
														},
														"match": schema.SingleNestedAttribute{
															MarkdownDescription: "Match",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"address_prefix": schema.ListNestedAttribute{
																	MarkdownDescription: "Address prefix",
																	Optional:            true,
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"exact": schema.BoolAttribute{
																				MarkdownDescription: "Exact",
																				Optional:            true,
																			},
																			"name": schema.StringAttribute{
																				MarkdownDescription: "Name",
																				Required:            true,
																			},
																		},
																	},
																},
																"afi": schema.StringAttribute{
																	Validators: []validator.String{
																		stringvalidator.OneOf("ip", "ipv6"),
																	},
																	MarkdownDescription: "Afi",
																	Optional:            true,
																},
																"as_path": schema.SingleNestedAttribute{
																	MarkdownDescription: "As path",
																	Optional:            true,
																	Attributes: map[string]schema.Attribute{
																		"regex": schema.StringAttribute{
																			MarkdownDescription: "Regex",
																			Optional:            true,
																		},
																	},
																},
																"community": schema.SingleNestedAttribute{
																	MarkdownDescription: "Community",
																	Optional:            true,
																	Attributes: map[string]schema.Attribute{
																		"regex": schema.StringAttribute{
																			MarkdownDescription: "Regex",
																			Optional:            true,
																		},
																	},
																},
																"extended_community": schema.SingleNestedAttribute{
																	MarkdownDescription: "Extended community",
																	Optional:            true,
																	Attributes: map[string]schema.Attribute{
																		"regex": schema.StringAttribute{
																			MarkdownDescription: "Regex",
																			Optional:            true,
																		},
																	},
																},
																"from_peer": schema.ListAttribute{
																	ElementType:         types.StringType,
																	MarkdownDescription: "From peer",
																	Optional:            true,
																},
																"med": schema.Float64Attribute{
																	MarkdownDescription: "Med",
																	Optional:            true,
																},
																"nexthop": schema.ListAttribute{
																	ElementType:         types.StringType,
																	MarkdownDescription: "Nexthop",
																	Optional:            true,
																},
																"route_table": schema.StringAttribute{
																	Validators: []validator.String{
																		stringvalidator.OneOf("unicast", "multicast", "both"),
																	},
																	MarkdownDescription: "Route table",
																	Optional:            true,
																},
																"safi": schema.StringAttribute{
																	Validators: []validator.String{
																		stringvalidator.OneOf("ip", "ipv6"),
																	},
																	MarkdownDescription: "Safi",
																	Optional:            true,
																},
															},
														},
														"name": schema.StringAttribute{
															MarkdownDescription: "Name",
															Required:            true,
														},
														"used_by": schema.ListAttribute{
															ElementType:         types.StringType,
															MarkdownDescription: "Used by",
															Optional:            true,
														},
													},
												},
											},
										},
									},
								},
							},
							"redist_rules": schema.ListNestedAttribute{
								MarkdownDescription: "Redist rules",
								Optional:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"address_family_identifier": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("ipv4", "ipv6"),
											},
											MarkdownDescription: "Address family identifier",
											Optional:            true,
										},
										"enable": schema.BoolAttribute{
											MarkdownDescription: "Enable",
											Optional:            true,
										},
										"metric": schema.Float64Attribute{
											MarkdownDescription: "Metric",
											Optional:            true,
										},
										"name": schema.StringAttribute{
											MarkdownDescription: "Name",
											Required:            true,
										},
										"route_table": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("unicast", "multicast", "both"),
											},
											MarkdownDescription: "Route table",
											Optional:            true,
										},
										"set_as_path_limit": schema.Float64Attribute{
											MarkdownDescription: "Set as path limit",
											Optional:            true,
										},
										"set_community": schema.ListAttribute{
											ElementType:         types.StringType,
											MarkdownDescription: "Set community",
											Optional:            true,
										},
										"set_extended_community": schema.ListAttribute{
											ElementType:         types.StringType,
											MarkdownDescription: "Set extended community",
											Optional:            true,
										},
										"set_local_preference": schema.Float64Attribute{
											MarkdownDescription: "Set local preference",
											Optional:            true,
										},
										"set_med": schema.Float64Attribute{
											MarkdownDescription: "Set med",
											Optional:            true,
										},
										"set_origin": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("igp", "egp", "incomplete"),
											},
											MarkdownDescription: "Set origin",
											Optional:            true,
										},
									},
								},
							},
							"redistribution_profile": schema.SingleNestedAttribute{
								MarkdownDescription: "Redistribution profile",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"ipv4": schema.SingleNestedAttribute{
										MarkdownDescription: "Ipv4",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"unicast": schema.StringAttribute{
												MarkdownDescription: "Unicast",
												Optional:            true,
											},
										},
									},
									"ipv6": schema.SingleNestedAttribute{
										MarkdownDescription: "Ipv6",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"unicast": schema.StringAttribute{
												MarkdownDescription: "Unicast",
												Optional:            true,
											},
										},
									},
								},
							},
							"reject_default_route": schema.BoolAttribute{
								MarkdownDescription: "Reject default route",
								Optional:            true,
							},
							"router_id": schema.StringAttribute{
								MarkdownDescription: "Router id",
								Optional:            true,
							},
						},
					},
					"ecmp": schema.SingleNestedAttribute{
						MarkdownDescription: "Ecmp",
						Optional:            true,
						Attributes: map[string]schema.Attribute{
							"algorithm": schema.SingleNestedAttribute{
								MarkdownDescription: "Algorithm",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"balanced_round_robin": schema.SingleNestedAttribute{
										Validators: []validator.Object{
											objectvalidator.ExactlyOneOf(
												path.MatchRelative().AtParent().AtName("ip_hash"),
												path.MatchRelative().AtParent().AtName("ip_modulo"),
												path.MatchRelative().AtParent().AtName("weighted_round_robin"),
											),
										},
										MarkdownDescription: "Balanced round robin",
										Optional:            true,
										Attributes:          map[string]schema.Attribute{},
									},
									"ip_hash": schema.SingleNestedAttribute{
										Validators: []validator.Object{
											objectvalidator.ExactlyOneOf(
												path.MatchRelative().AtParent().AtName("balanced_round_robin"),
												path.MatchRelative().AtParent().AtName("ip_modulo"),
												path.MatchRelative().AtParent().AtName("weighted_round_robin"),
											),
										},
										MarkdownDescription: "Ip hash",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"hash_seed": schema.Float64Attribute{
												MarkdownDescription: "Hash seed",
												Optional:            true,
											},
											"src_only": schema.BoolAttribute{
												MarkdownDescription: "Src only",
												Optional:            true,
											},
											"use_port": schema.BoolAttribute{
												MarkdownDescription: "Use port",
												Optional:            true,
											},
										},
									},
									"ip_modulo": schema.SingleNestedAttribute{
										Validators: []validator.Object{
											objectvalidator.ExactlyOneOf(
												path.MatchRelative().AtParent().AtName("balanced_round_robin"),
												path.MatchRelative().AtParent().AtName("ip_hash"),
												path.MatchRelative().AtParent().AtName("weighted_round_robin"),
											),
										},
										MarkdownDescription: "Ip modulo",
										Optional:            true,
										Attributes:          map[string]schema.Attribute{},
									},
									"weighted_round_robin": schema.SingleNestedAttribute{
										Validators: []validator.Object{
											objectvalidator.ExactlyOneOf(
												path.MatchRelative().AtParent().AtName("balanced_round_robin"),
												path.MatchRelative().AtParent().AtName("ip_hash"),
												path.MatchRelative().AtParent().AtName("ip_modulo"),
											),
										},
										MarkdownDescription: "Weighted round robin",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"interface": schema.ListNestedAttribute{
												MarkdownDescription: "Interface",
												Optional:            true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															MarkdownDescription: "Name",
															Required:            true,
														},
														"weight": schema.Float64Attribute{
															MarkdownDescription: "Weight",
															Optional:            true,
														},
													},
												},
											},
										},
									},
								},
							},
							"enable": schema.BoolAttribute{
								MarkdownDescription: "Enable",
								Optional:            true,
							},
							"max_path": schema.Float64Attribute{
								MarkdownDescription: "Max path",
								Optional:            true,
							},
							"strict_source_path": schema.BoolAttribute{
								MarkdownDescription: "Strict source path",
								Optional:            true,
							},
							"symmetric_return": schema.BoolAttribute{
								MarkdownDescription: "Symmetric return",
								Optional:            true,
							},
						},
					},
					"global_vrid": schema.Float64Attribute{
						MarkdownDescription: "Global vrid",
						Optional:            true,
					},
					"interface": schema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "Interface",
						Optional:            true,
					},
					"multicast": schema.SingleNestedAttribute{
						MarkdownDescription: "Multicast",
						Optional:            true,
						Attributes: map[string]schema.Attribute{
							"enable": schema.BoolAttribute{
								MarkdownDescription: "Enable",
								Optional:            true,
							},
							"enable_v6": schema.BoolAttribute{
								MarkdownDescription: "Enable v6",
								Optional:            true,
							},
							"igmp": schema.SingleNestedAttribute{
								MarkdownDescription: "Igmp",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"dynamic": schema.SingleNestedAttribute{
										MarkdownDescription: "Dynamic",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"interface": schema.ListNestedAttribute{
												MarkdownDescription: "Interface",
												Optional:            true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"group_filter": schema.StringAttribute{
															MarkdownDescription: "Group filter",
															Optional:            true,
														},
														"max_groups": schema.StringAttribute{
															MarkdownDescription: "Max groups",
															Optional:            true,
														},
														"max_sources": schema.StringAttribute{
															MarkdownDescription: "Max sources",
															Optional:            true,
														},
														"name": schema.StringAttribute{
															MarkdownDescription: "Name",
															Required:            true,
														},
														"query_profile": schema.StringAttribute{
															MarkdownDescription: "Query profile",
															Optional:            true,
														},
														"robustness": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.OneOf("1", "2", "3", "4", "5", "6", "7"),
															},
															MarkdownDescription: "Robustness",
															Optional:            true,
														},
														"router_alert_policing": schema.BoolAttribute{
															MarkdownDescription: "Router alert policing",
															Optional:            true,
														},
														"version": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.OneOf("2", "3"),
															},
															MarkdownDescription: "Version",
															Optional:            true,
														},
													},
												},
											},
										},
									},
									"enable": schema.BoolAttribute{
										MarkdownDescription: "Enable",
										Optional:            true,
									},
									"static": schema.ListNestedAttribute{
										MarkdownDescription: "Static",
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"group_address": schema.StringAttribute{
													MarkdownDescription: "Group address",
													Optional:            true,
												},
												"interface": schema.StringAttribute{
													MarkdownDescription: "Interface",
													Optional:            true,
												},
												"name": schema.StringAttribute{
													MarkdownDescription: "Name",
													Required:            true,
												},
												"source_address": schema.StringAttribute{
													MarkdownDescription: "Source address",
													Optional:            true,
												},
											},
										},
									},
								},
							},
							"interface_group": schema.ListNestedAttribute{
								MarkdownDescription: "Interface group",
								Optional:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"description": schema.StringAttribute{
											MarkdownDescription: "Description",
											Optional:            true,
										},
										"group_permission": schema.SingleNestedAttribute{
											MarkdownDescription: "Group permission",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"any_source_multicast": schema.ListNestedAttribute{
													MarkdownDescription: "Any source multicast",
													Optional:            true,
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"group_address": schema.StringAttribute{
																MarkdownDescription: "Group address",
																Optional:            true,
															},
															"included": schema.BoolAttribute{
																MarkdownDescription: "Included",
																Optional:            true,
															},
															"name": schema.StringAttribute{
																MarkdownDescription: "Name",
																Required:            true,
															},
														},
													},
												},
												"source_specific_multicast": schema.ListNestedAttribute{
													MarkdownDescription: "Source specific multicast",
													Optional:            true,
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"group_address": schema.StringAttribute{
																MarkdownDescription: "Group address",
																Optional:            true,
															},
															"included": schema.BoolAttribute{
																MarkdownDescription: "Included",
																Optional:            true,
															},
															"name": schema.StringAttribute{
																MarkdownDescription: "Name",
																Required:            true,
															},
															"source_address": schema.StringAttribute{
																MarkdownDescription: "Source address",
																Optional:            true,
															},
														},
													},
												},
											},
										},
										"igmp": schema.SingleNestedAttribute{
											MarkdownDescription: "Igmp",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"enable": schema.BoolAttribute{
													MarkdownDescription: "Enable",
													Optional:            true,
												},
												"immediate_leave": schema.BoolAttribute{
													MarkdownDescription: "Immediate leave",
													Optional:            true,
												},
												"last_member_query_interval": schema.Float64Attribute{
													MarkdownDescription: "Last member query interval",
													Optional:            true,
												},
												"max_groups": schema.StringAttribute{
													MarkdownDescription: "Max groups",
													Optional:            true,
												},
												"max_query_response_time": schema.Float64Attribute{
													MarkdownDescription: "Max query response time",
													Optional:            true,
												},
												"max_sources": schema.StringAttribute{
													MarkdownDescription: "Max sources",
													Optional:            true,
												},
												"mode": schema.StringAttribute{
													Validators: []validator.String{
														stringvalidator.OneOf("router", "host"),
													},
													MarkdownDescription: "Mode",
													Optional:            true,
												},
												"query_interval": schema.Float64Attribute{
													MarkdownDescription: "Query interval",
													Optional:            true,
												},
												"robustness": schema.StringAttribute{
													Validators: []validator.String{
														stringvalidator.OneOf("1", "2", "3", "4", "5", "6", "7"),
													},
													MarkdownDescription: "Robustness",
													Optional:            true,
												},
												"router_alert_policing": schema.BoolAttribute{
													MarkdownDescription: "Router alert policing",
													Optional:            true,
												},
												"version": schema.StringAttribute{
													Validators: []validator.String{
														stringvalidator.OneOf("1", "2", "3"),
													},
													MarkdownDescription: "Version",
													Optional:            true,
												},
											},
										},
										"interface": schema.ListAttribute{
											ElementType:         types.StringType,
											MarkdownDescription: "Interface",
											Optional:            true,
										},
										"name": schema.StringAttribute{
											MarkdownDescription: "Name",
											Required:            true,
										},
										"pim": schema.SingleNestedAttribute{
											MarkdownDescription: "Pim",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"allowed_neighbors": schema.ListNestedAttribute{
													MarkdownDescription: "Allowed neighbors",
													Optional:            true,
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																MarkdownDescription: "Name",
																Required:            true,
															},
														},
													},
												},
												"assert_interval": schema.Float64Attribute{
													MarkdownDescription: "Assert interval",
													Optional:            true,
												},
												"bsr_border": schema.BoolAttribute{
													MarkdownDescription: "Bsr border",
													Optional:            true,
												},
												"dr_priority": schema.Float64Attribute{
													MarkdownDescription: "Dr priority",
													Optional:            true,
												},
												"enable": schema.BoolAttribute{
													MarkdownDescription: "Enable",
													Optional:            true,
												},
												"hello_interval": schema.Float64Attribute{
													MarkdownDescription: "Hello interval",
													Optional:            true,
												},
												"join_prune_interval": schema.Float64Attribute{
													MarkdownDescription: "Join prune interval",
													Optional:            true,
												},
											},
										},
									},
								},
							},
							"mode": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("PIM-SM", "IGMP-Proxy"),
								},
								MarkdownDescription: "Mode",
								Optional:            true,
							},
							"msdp": schema.SingleNestedAttribute{
								MarkdownDescription: "Msdp",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"enable": schema.BoolAttribute{
										MarkdownDescription: "Enable",
										Optional:            true,
									},
									"global_authentication": schema.StringAttribute{
										MarkdownDescription: "Global authentication",
										Optional:            true,
									},
									"global_timer": schema.StringAttribute{
										MarkdownDescription: "Global timer",
										Optional:            true,
									},
									"originator_id": schema.SingleNestedAttribute{
										MarkdownDescription: "Originator id",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"interface": schema.StringAttribute{
												MarkdownDescription: "Interface",
												Optional:            true,
											},
											"ip": schema.StringAttribute{
												MarkdownDescription: "Ip",
												Optional:            true,
											},
										},
									},
									"peer": schema.ListNestedAttribute{
										MarkdownDescription: "Peer",
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"authentication": schema.StringAttribute{
													MarkdownDescription: "Authentication",
													Optional:            true,
												},
												"enable": schema.BoolAttribute{
													MarkdownDescription: "Enable",
													Optional:            true,
												},
												"inbound_sa_filter": schema.StringAttribute{
													MarkdownDescription: "Inbound sa filter",
													Optional:            true,
												},
												"local_address": schema.SingleNestedAttribute{
													MarkdownDescription: "Local address",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"interface": schema.StringAttribute{
															MarkdownDescription: "Interface",
															Optional:            true,
														},
														"ip": schema.StringAttribute{
															MarkdownDescription: "Ip",
															Optional:            true,
														},
													},
												},
												"max_sa": schema.Float64Attribute{
													MarkdownDescription: "Max sa",
													Optional:            true,
												},
												"name": schema.StringAttribute{
													MarkdownDescription: "Name",
													Required:            true,
												},
												"outbound_sa_filter": schema.StringAttribute{
													MarkdownDescription: "Outbound sa filter",
													Optional:            true,
												},
												"peer_address": schema.SingleNestedAttribute{
													MarkdownDescription: "Peer address",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"fqdn": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("ip"),
																),
															},
															MarkdownDescription: "Fqdn",
															Optional:            true,
														},
														"ip": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("fqdn"),
																),
															},
															MarkdownDescription: "Ip",
															Optional:            true,
														},
													},
												},
												"peer_as": schema.StringAttribute{
													MarkdownDescription: "Peer as",
													Optional:            true,
												},
											},
										},
									},
								},
							},
							"pim": schema.SingleNestedAttribute{
								MarkdownDescription: "Pim",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"enable": schema.BoolAttribute{
										MarkdownDescription: "Enable",
										Optional:            true,
									},
									"group_permission": schema.StringAttribute{
										MarkdownDescription: "Group permission",
										Optional:            true,
									},
									"if_timer_global": schema.StringAttribute{
										MarkdownDescription: "If timer global",
										Optional:            true,
									},
									"interface": schema.ListNestedAttribute{
										MarkdownDescription: "Interface",
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"description": schema.StringAttribute{
													MarkdownDescription: "Description",
													Optional:            true,
												},
												"dr_priority": schema.Float64Attribute{
													MarkdownDescription: "Dr priority",
													Optional:            true,
												},
												"if_timer": schema.StringAttribute{
													MarkdownDescription: "If timer",
													Optional:            true,
												},
												"name": schema.StringAttribute{
													MarkdownDescription: "Name",
													Required:            true,
												},
												"neighbor_filter": schema.StringAttribute{
													MarkdownDescription: "Neighbor filter",
													Optional:            true,
												},
												"send_bsm": schema.BoolAttribute{
													MarkdownDescription: "Send bsm",
													Optional:            true,
												},
											},
										},
									},
									"route_ageout_time": schema.Float64Attribute{
										MarkdownDescription: "Route ageout time",
										Optional:            true,
									},
									"rp": schema.SingleNestedAttribute{
										MarkdownDescription: "Rp",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"external_rp": schema.ListNestedAttribute{
												MarkdownDescription: "External rp",
												Optional:            true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"group_list": schema.StringAttribute{
															MarkdownDescription: "Group list",
															Optional:            true,
														},
														"name": schema.StringAttribute{
															MarkdownDescription: "Name",
															Optional:            true,
														},
														"override": schema.BoolAttribute{
															MarkdownDescription: "Override",
															Optional:            true,
														},
													},
												},
											},
											"local_rp": schema.SingleNestedAttribute{
												MarkdownDescription: "Local rp",
												Optional:            true,
												Attributes: map[string]schema.Attribute{
													"candidate_rp": schema.SingleNestedAttribute{
														Validators: []validator.Object{
															objectvalidator.ExactlyOneOf(
																path.MatchRelative().AtParent().AtName("static_rp"),
															),
														},
														MarkdownDescription: "Candidate rp",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"address": schema.StringAttribute{
																MarkdownDescription: "Address",
																Optional:            true,
															},
															"advertisement_interval": schema.Float64Attribute{
																MarkdownDescription: "Advertisement interval",
																Optional:            true,
															},
															"group_list": schema.StringAttribute{
																MarkdownDescription: "Group list",
																Optional:            true,
															},
															"interface": schema.StringAttribute{
																MarkdownDescription: "Interface",
																Optional:            true,
															},
															"priority": schema.Float64Attribute{
																MarkdownDescription: "Priority",
																Optional:            true,
															},
														},
													},
													"static_rp": schema.SingleNestedAttribute{
														Validators: []validator.Object{
															objectvalidator.ExactlyOneOf(
																path.MatchRelative().AtParent().AtName("candidate_rp"),
															),
														},
														MarkdownDescription: "Static rp",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"address": schema.StringAttribute{
																MarkdownDescription: "Address",
																Optional:            true,
															},
															"group_list": schema.StringAttribute{
																MarkdownDescription: "Group list",
																Optional:            true,
															},
															"interface": schema.StringAttribute{
																MarkdownDescription: "Interface",
																Optional:            true,
															},
															"override": schema.BoolAttribute{
																MarkdownDescription: "Override",
																Optional:            true,
															},
														},
													},
												},
											},
										},
									},
									"rpf_lookup_mode": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("mrib-then-urib", "mrib-only", "urib-only"),
										},
										MarkdownDescription: "Rpf lookup mode",
										Optional:            true,
									},
									"spt_threshold": schema.ListNestedAttribute{
										MarkdownDescription: "Spt threshold",
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													MarkdownDescription: "Name",
													Required:            true,
												},
												"threshold": schema.StringAttribute{
													MarkdownDescription: "Threshold",
													Optional:            true,
												},
											},
										},
									},
									"ssm_address_space": schema.SingleNestedAttribute{
										MarkdownDescription: "Ssm address space",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"group_list": schema.StringAttribute{
												MarkdownDescription: "Group list",
												Optional:            true,
											},
										},
									},
								},
							},
							"route_ageout_time": schema.Float64Attribute{
								MarkdownDescription: "Route ageout time",
								Optional:            true,
							},
							"rp": schema.SingleNestedAttribute{
								MarkdownDescription: "Rp",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"external_rp": schema.ListNestedAttribute{
										MarkdownDescription: "External rp",
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"group_addresses": schema.ListAttribute{
													ElementType:         types.StringType,
													MarkdownDescription: "Group addresses",
													Optional:            true,
												},
												"name": schema.StringAttribute{
													MarkdownDescription: "Name",
													Required:            true,
												},
												"override": schema.BoolAttribute{
													MarkdownDescription: "Override",
													Optional:            true,
												},
											},
										},
									},
									"local_rp": schema.SingleNestedAttribute{
										MarkdownDescription: "Local rp",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"candidate_rp": schema.SingleNestedAttribute{
												Validators: []validator.Object{
													objectvalidator.ExactlyOneOf(
														path.MatchRelative().AtParent().AtName("static_rp"),
													),
												},
												MarkdownDescription: "Candidate rp",
												Optional:            true,
												Attributes: map[string]schema.Attribute{
													"address": schema.StringAttribute{
														MarkdownDescription: "Address",
														Optional:            true,
													},
													"advertisement_interval": schema.Float64Attribute{
														MarkdownDescription: "Advertisement interval",
														Optional:            true,
													},
													"group_addresses": schema.ListAttribute{
														ElementType:         types.StringType,
														MarkdownDescription: "Group addresses",
														Optional:            true,
													},
													"interface": schema.StringAttribute{
														MarkdownDescription: "Interface",
														Optional:            true,
													},
													"priority": schema.Float64Attribute{
														MarkdownDescription: "Priority",
														Optional:            true,
													},
												},
											},
											"static_rp": schema.SingleNestedAttribute{
												Validators: []validator.Object{
													objectvalidator.ExactlyOneOf(
														path.MatchRelative().AtParent().AtName("candidate_rp"),
													),
												},
												MarkdownDescription: "Static rp",
												Optional:            true,
												Attributes: map[string]schema.Attribute{
													"address": schema.StringAttribute{
														MarkdownDescription: "Address",
														Optional:            true,
													},
													"group_addresses": schema.ListAttribute{
														ElementType:         types.StringType,
														MarkdownDescription: "Group addresses",
														Optional:            true,
													},
													"interface": schema.StringAttribute{
														MarkdownDescription: "Interface",
														Optional:            true,
													},
													"override": schema.BoolAttribute{
														MarkdownDescription: "Override",
														Optional:            true,
													},
												},
											},
										},
									},
								},
							},
							"spt_threshold": schema.ListNestedAttribute{
								MarkdownDescription: "Spt threshold",
								Optional:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											MarkdownDescription: "Name",
											Required:            true,
										},
										"threshold": schema.StringAttribute{
											MarkdownDescription: "Threshold",
											Optional:            true,
										},
									},
								},
							},
							"ssm_address_space": schema.ListNestedAttribute{
								MarkdownDescription: "Ssm address space",
								Optional:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"group_address": schema.StringAttribute{
											MarkdownDescription: "Group address",
											Optional:            true,
										},
										"included": schema.BoolAttribute{
											MarkdownDescription: "Included",
											Optional:            true,
										},
										"name": schema.StringAttribute{
											MarkdownDescription: "Name",
											Required:            true,
										},
									},
								},
							},
							"static_route": schema.ListNestedAttribute{
								MarkdownDescription: "Static route",
								Optional:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"destination": schema.StringAttribute{
											MarkdownDescription: "Destination",
											Optional:            true,
										},
										"interface": schema.StringAttribute{
											MarkdownDescription: "Interface",
											Optional:            true,
										},
										"name": schema.StringAttribute{
											MarkdownDescription: "Name",
											Required:            true,
										},
										"nexthop": schema.SingleNestedAttribute{
											MarkdownDescription: "Nexthop",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"ip_address": schema.StringAttribute{
													MarkdownDescription: "Ip address",
													Optional:            true,
												},
											},
										},
										"preference": schema.Float64Attribute{
											MarkdownDescription: "Preference",
											Optional:            true,
										},
									},
								},
							},
						},
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "Name",
						Required:            true,
					},
					"ospf": schema.SingleNestedAttribute{
						MarkdownDescription: "Ospf",
						Optional:            true,
						Attributes: map[string]schema.Attribute{
							"allow_redist_default_route": schema.BoolAttribute{
								MarkdownDescription: "Allow redist default route",
								Optional:            true,
							},
							"area": schema.ListNestedAttribute{
								MarkdownDescription: "Area",
								Optional:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"authentication": schema.StringAttribute{
											MarkdownDescription: "Authentication",
											Optional:            true,
										},
										"interface": schema.ListNestedAttribute{
											MarkdownDescription: "Interface",
											Optional:            true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"authentication": schema.StringAttribute{
														MarkdownDescription: "Authentication",
														Optional:            true,
													},
													"bfd": schema.SingleNestedAttribute{
														MarkdownDescription: "Bfd",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"profile": schema.StringAttribute{
																MarkdownDescription: "Profile",
																Optional:            true,
															},
														},
													},
													"enable": schema.BoolAttribute{
														MarkdownDescription: "Enable",
														Optional:            true,
													},
													"link_type": schema.SingleNestedAttribute{
														MarkdownDescription: "Link type",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"broadcast": schema.SingleNestedAttribute{
																Validators: []validator.Object{
																	objectvalidator.ExactlyOneOf(
																		path.MatchRelative().AtParent().AtName("p2mp"),
																		path.MatchRelative().AtParent().AtName("p2p"),
																	),
																},
																MarkdownDescription: "Broadcast",
																Optional:            true,
																Attributes:          map[string]schema.Attribute{},
															},
															"p2mp": schema.SingleNestedAttribute{
																Validators: []validator.Object{
																	objectvalidator.ExactlyOneOf(
																		path.MatchRelative().AtParent().AtName("broadcast"),
																		path.MatchRelative().AtParent().AtName("p2p"),
																	),
																},
																MarkdownDescription: "P2mp",
																Optional:            true,
																Attributes: map[string]schema.Attribute{
																	"neighbor": schema.ListNestedAttribute{
																		MarkdownDescription: "Neighbor",
																		Optional:            true,
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					MarkdownDescription: "Name",
																					Required:            true,
																				},
																				"priority": schema.Float64Attribute{
																					MarkdownDescription: "Priority",
																					Optional:            true,
																				},
																			},
																		},
																	},
																},
															},
															"p2p": schema.SingleNestedAttribute{
																Validators: []validator.Object{
																	objectvalidator.ExactlyOneOf(
																		path.MatchRelative().AtParent().AtName("broadcast"),
																		path.MatchRelative().AtParent().AtName("p2mp"),
																	),
																},
																MarkdownDescription: "P2p",
																Optional:            true,
																Attributes:          map[string]schema.Attribute{},
															},
														},
													},
													"metric": schema.Float64Attribute{
														MarkdownDescription: "Metric",
														Optional:            true,
													},
													"mtu_ignore": schema.BoolAttribute{
														MarkdownDescription: "Mtu ignore",
														Optional:            true,
													},
													"name": schema.StringAttribute{
														MarkdownDescription: "Name",
														Required:            true,
													},
													"passive": schema.BoolAttribute{
														MarkdownDescription: "Passive",
														Optional:            true,
													},
													"priority": schema.Float64Attribute{
														MarkdownDescription: "Priority",
														Optional:            true,
													},
													"timing": schema.StringAttribute{
														MarkdownDescription: "Timing",
														Optional:            true,
													},
													"vr_timing": schema.SingleNestedAttribute{
														MarkdownDescription: "Vr timing",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"dead_counts": schema.Float64Attribute{
																MarkdownDescription: "Dead counts",
																Optional:            true,
															},
															"gr_delay": schema.Float64Attribute{
																MarkdownDescription: "Gr delay",
																Optional:            true,
															},
															"hello_interval": schema.Float64Attribute{
																MarkdownDescription: "Hello interval",
																Optional:            true,
															},
															"retransmit_interval": schema.Float64Attribute{
																MarkdownDescription: "Retransmit interval",
																Optional:            true,
															},
															"transit_delay": schema.Float64Attribute{
																MarkdownDescription: "Transit delay",
																Optional:            true,
															},
														},
													},
												},
											},
										},
										"name": schema.StringAttribute{
											MarkdownDescription: "Name",
											Required:            true,
										},
										"range": schema.ListNestedAttribute{
											MarkdownDescription: "Range",
											Optional:            true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"advertise": schema.BoolAttribute{
														MarkdownDescription: "Advertise",
														Optional:            true,
													},
													"name": schema.StringAttribute{
														MarkdownDescription: "Name",
														Required:            true,
													},
													"substitute": schema.StringAttribute{
														MarkdownDescription: "Substitute",
														Optional:            true,
													},
												},
											},
										},
										"type": schema.SingleNestedAttribute{
											MarkdownDescription: "Type",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"normal": schema.SingleNestedAttribute{
													Validators: []validator.Object{
														objectvalidator.ExactlyOneOf(
															path.MatchRelative().AtParent().AtName("nssa"),
															path.MatchRelative().AtParent().AtName("stub"),
														),
													},
													MarkdownDescription: "Normal",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"abr": schema.SingleNestedAttribute{
															MarkdownDescription: "Abr",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"export_list": schema.StringAttribute{
																	MarkdownDescription: "Export list",
																	Optional:            true,
																},
																"import_list": schema.StringAttribute{
																	MarkdownDescription: "Import list",
																	Optional:            true,
																},
																"inbound_filter_list": schema.StringAttribute{
																	MarkdownDescription: "Inbound filter list",
																	Optional:            true,
																},
																"outbound_filter_list": schema.StringAttribute{
																	MarkdownDescription: "Outbound filter list",
																	Optional:            true,
																},
															},
														},
													},
												},
												"nssa": schema.SingleNestedAttribute{
													Validators: []validator.Object{
														objectvalidator.ExactlyOneOf(
															path.MatchRelative().AtParent().AtName("normal"),
															path.MatchRelative().AtParent().AtName("stub"),
														),
													},
													MarkdownDescription: "Nssa",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"abr": schema.SingleNestedAttribute{
															MarkdownDescription: "Abr",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"export_list": schema.StringAttribute{
																	MarkdownDescription: "Export list",
																	Optional:            true,
																},
																"import_list": schema.StringAttribute{
																	MarkdownDescription: "Import list",
																	Optional:            true,
																},
																"inbound_filter_list": schema.StringAttribute{
																	MarkdownDescription: "Inbound filter list",
																	Optional:            true,
																},
																"nssa_ext_range": schema.ListNestedAttribute{
																	MarkdownDescription: "Nssa ext range",
																	Optional:            true,
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"advertise": schema.BoolAttribute{
																				MarkdownDescription: "Advertise",
																				Optional:            true,
																			},
																			"name": schema.StringAttribute{
																				MarkdownDescription: "Name",
																				Required:            true,
																			},
																			"route_tag": schema.Float64Attribute{
																				MarkdownDescription: "Route tag",
																				Optional:            true,
																			},
																		},
																	},
																},
																"outbound_filter_list": schema.StringAttribute{
																	MarkdownDescription: "Outbound filter list",
																	Optional:            true,
																},
															},
														},
														"accept_summary": schema.BoolAttribute{
															MarkdownDescription: "Accept summary",
															Optional:            true,
														},
														"default_information_originate": schema.SingleNestedAttribute{
															MarkdownDescription: "Default information originate",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"metric": schema.Float64Attribute{
																	MarkdownDescription: "Metric",
																	Optional:            true,
																},
																"metric_type": schema.StringAttribute{
																	Validators: []validator.String{
																		stringvalidator.OneOf("type-1", "type-2"),
																	},
																	MarkdownDescription: "Metric type",
																	Optional:            true,
																},
															},
														},
														"default_route": schema.SingleNestedAttribute{
															MarkdownDescription: "Default route",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"advertise": schema.SingleNestedAttribute{
																	Validators: []validator.Object{
																		objectvalidator.ExactlyOneOf(
																			path.MatchRelative().AtParent().AtName("disable"),
																		),
																	},
																	MarkdownDescription: "Advertise",
																	Optional:            true,
																	Attributes: map[string]schema.Attribute{
																		"metric": schema.Float64Attribute{
																			MarkdownDescription: "Metric",
																			Optional:            true,
																		},
																		"type": schema.StringAttribute{
																			Validators: []validator.String{
																				stringvalidator.OneOf("ext-1", "ext-2"),
																			},
																			MarkdownDescription: "Type",
																			Optional:            true,
																		},
																	},
																},
																"disable": schema.SingleNestedAttribute{
																	Validators: []validator.Object{
																		objectvalidator.ExactlyOneOf(
																			path.MatchRelative().AtParent().AtName("advertise"),
																		),
																	},
																	MarkdownDescription: "Disable",
																	Optional:            true,
																	Attributes:          map[string]schema.Attribute{},
																},
															},
														},
														"no_summary": schema.BoolAttribute{
															MarkdownDescription: "No summary",
															Optional:            true,
														},
														"nssa_ext_range": schema.ListNestedAttribute{
															MarkdownDescription: "Nssa ext range",
															Optional:            true,
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"advertise": schema.SingleNestedAttribute{
																		MarkdownDescription: "Advertise",
																		Optional:            true,
																		Attributes:          map[string]schema.Attribute{},
																	},
																	"name": schema.StringAttribute{
																		MarkdownDescription: "Name",
																		Required:            true,
																	},
																	"suppress": schema.SingleNestedAttribute{
																		MarkdownDescription: "Suppress",
																		Optional:            true,
																		Attributes:          map[string]schema.Attribute{},
																	},
																},
															},
														},
													},
												},
												"stub": schema.SingleNestedAttribute{
													Validators: []validator.Object{
														objectvalidator.ExactlyOneOf(
															path.MatchRelative().AtParent().AtName("normal"),
															path.MatchRelative().AtParent().AtName("nssa"),
														),
													},
													MarkdownDescription: "Stub",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"abr": schema.SingleNestedAttribute{
															MarkdownDescription: "Abr",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"export_list": schema.StringAttribute{
																	MarkdownDescription: "Export list",
																	Optional:            true,
																},
																"import_list": schema.StringAttribute{
																	MarkdownDescription: "Import list",
																	Optional:            true,
																},
																"inbound_filter_list": schema.StringAttribute{
																	MarkdownDescription: "Inbound filter list",
																	Optional:            true,
																},
																"outbound_filter_list": schema.StringAttribute{
																	MarkdownDescription: "Outbound filter list",
																	Optional:            true,
																},
															},
														},
														"accept_summary": schema.BoolAttribute{
															MarkdownDescription: "Accept summary",
															Optional:            true,
														},
														"default_route": schema.SingleNestedAttribute{
															MarkdownDescription: "Default route",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"advertise": schema.SingleNestedAttribute{
																	Validators: []validator.Object{
																		objectvalidator.ExactlyOneOf(
																			path.MatchRelative().AtParent().AtName("disable"),
																		),
																	},
																	MarkdownDescription: "Advertise",
																	Optional:            true,
																	Attributes: map[string]schema.Attribute{
																		"metric": schema.Float64Attribute{
																			MarkdownDescription: "Metric",
																			Optional:            true,
																		},
																	},
																},
																"disable": schema.SingleNestedAttribute{
																	Validators: []validator.Object{
																		objectvalidator.ExactlyOneOf(
																			path.MatchRelative().AtParent().AtName("advertise"),
																		),
																	},
																	MarkdownDescription: "Disable",
																	Optional:            true,
																	Attributes:          map[string]schema.Attribute{},
																},
															},
														},
														"default_route_metric": schema.Float64Attribute{
															MarkdownDescription: "Default route metric",
															Optional:            true,
														},
														"no_summary": schema.BoolAttribute{
															MarkdownDescription: "No summary",
															Optional:            true,
														},
													},
												},
											},
										},
										"virtual_link": schema.ListNestedAttribute{
											MarkdownDescription: "Virtual link",
											Optional:            true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"authentication": schema.StringAttribute{
														MarkdownDescription: "Authentication",
														Optional:            true,
													},
													"bfd": schema.SingleNestedAttribute{
														MarkdownDescription: "Bfd",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"profile": schema.StringAttribute{
																MarkdownDescription: "Profile",
																Optional:            true,
															},
														},
													},
													"enable": schema.BoolAttribute{
														MarkdownDescription: "Enable",
														Optional:            true,
													},
													"instance_id": schema.Float64Attribute{
														MarkdownDescription: "Instance id",
														Optional:            true,
													},
													"interface_id": schema.Float64Attribute{
														MarkdownDescription: "Interface id",
														Optional:            true,
													},
													"name": schema.StringAttribute{
														MarkdownDescription: "Name",
														Required:            true,
													},
													"neighbor_id": schema.StringAttribute{
														MarkdownDescription: "Neighbor id",
														Optional:            true,
													},
													"passive": schema.BoolAttribute{
														MarkdownDescription: "Passive",
														Optional:            true,
													},
													"timing": schema.StringAttribute{
														MarkdownDescription: "Timing",
														Optional:            true,
													},
													"transit_area_id": schema.StringAttribute{
														MarkdownDescription: "Transit area id",
														Optional:            true,
													},
													"vr_timing": schema.SingleNestedAttribute{
														MarkdownDescription: "Vr timing",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"dead_counts": schema.Float64Attribute{
																MarkdownDescription: "Dead counts",
																Optional:            true,
															},
															"hello_interval": schema.Float64Attribute{
																MarkdownDescription: "Hello interval",
																Optional:            true,
															},
															"retransmit_interval": schema.Float64Attribute{
																MarkdownDescription: "Retransmit interval",
																Optional:            true,
															},
															"transit_delay": schema.Float64Attribute{
																MarkdownDescription: "Transit delay",
																Optional:            true,
															},
														},
													},
												},
											},
										},
										"vr_range": schema.ListNestedAttribute{
											MarkdownDescription: "Vr range",
											Optional:            true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"advertise": schema.SingleNestedAttribute{
														MarkdownDescription: "Advertise",
														Optional:            true,
														Attributes:          map[string]schema.Attribute{},
													},
													"name": schema.StringAttribute{
														MarkdownDescription: "Name",
														Required:            true,
													},
													"suppress": schema.SingleNestedAttribute{
														MarkdownDescription: "Suppress",
														Optional:            true,
														Attributes:          map[string]schema.Attribute{},
													},
												},
											},
										},
									},
								},
							},
							"auth_profile": schema.ListNestedAttribute{
								MarkdownDescription: "Auth profile",
								Optional:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"md5": schema.ListNestedAttribute{
											MarkdownDescription: "Md5",
											Optional:            true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														MarkdownDescription: "Key",
														Optional:            true,
													},
													"name": schema.Float64Attribute{
														MarkdownDescription: "Name",
														Required:            true,
													},
													"preferred": schema.BoolAttribute{
														MarkdownDescription: "Preferred",
														Optional:            true,
													},
												},
											},
										},
										"name": schema.StringAttribute{
											MarkdownDescription: "Name",
											Required:            true,
										},
										"password": schema.StringAttribute{
											MarkdownDescription: "Password",
											Optional:            true,
										},
									},
								},
							},
							"enable": schema.BoolAttribute{
								MarkdownDescription: "Enable",
								Optional:            true,
							},
							"export_rules": schema.ListNestedAttribute{
								MarkdownDescription: "Export rules",
								Optional:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"metric": schema.Float64Attribute{
											MarkdownDescription: "Metric",
											Optional:            true,
										},
										"name": schema.StringAttribute{
											MarkdownDescription: "Name",
											Required:            true,
										},
										"new_path_type": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("ext-1", "ext-2"),
											},
											MarkdownDescription: "New path type",
											Optional:            true,
										},
										"new_tag": schema.StringAttribute{
											MarkdownDescription: "New tag",
											Optional:            true,
										},
									},
								},
							},
							"flood_prevention": schema.SingleNestedAttribute{
								MarkdownDescription: "Flood prevention",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"hello": schema.SingleNestedAttribute{
										MarkdownDescription: "Hello",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"enable": schema.BoolAttribute{
												MarkdownDescription: "Enable",
												Optional:            true,
											},
											"max_packet": schema.Float64Attribute{
												MarkdownDescription: "Max packet",
												Optional:            true,
											},
										},
									},
									"lsa": schema.SingleNestedAttribute{
										MarkdownDescription: "Lsa",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"enable": schema.BoolAttribute{
												MarkdownDescription: "Enable",
												Optional:            true,
											},
											"max_packet": schema.Float64Attribute{
												MarkdownDescription: "Max packet",
												Optional:            true,
											},
										},
									},
								},
							},
							"global_bfd": schema.SingleNestedAttribute{
								MarkdownDescription: "Global bfd",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"profile": schema.StringAttribute{
										MarkdownDescription: "Profile",
										Optional:            true,
									},
								},
							},
							"global_if_timer": schema.StringAttribute{
								MarkdownDescription: "Global if timer",
								Optional:            true,
							},
							"graceful_restart": schema.SingleNestedAttribute{
								MarkdownDescription: "Graceful restart",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"enable": schema.BoolAttribute{
										MarkdownDescription: "Enable",
										Optional:            true,
									},
									"grace_period": schema.Float64Attribute{
										MarkdownDescription: "Grace period",
										Optional:            true,
									},
									"helper_enable": schema.BoolAttribute{
										MarkdownDescription: "Helper enable",
										Optional:            true,
									},
									"max_neighbor_restart_time": schema.Float64Attribute{
										MarkdownDescription: "Max neighbor restart time",
										Optional:            true,
									},
									"strict__l_s_a_checking": schema.BoolAttribute{
										MarkdownDescription: "Strict l s a checking",
										Optional:            true,
									},
								},
							},
							"redistribution_profile": schema.StringAttribute{
								MarkdownDescription: "Redistribution profile",
								Optional:            true,
							},
							"reject_default_route": schema.BoolAttribute{
								MarkdownDescription: "Reject default route",
								Optional:            true,
							},
							"rfc1583": schema.BoolAttribute{
								MarkdownDescription: "Rfc1583",
								Optional:            true,
							},
							"router_id": schema.StringAttribute{
								MarkdownDescription: "Router id",
								Optional:            true,
							},
							"spf_timer": schema.StringAttribute{
								MarkdownDescription: "Spf timer",
								Optional:            true,
							},
							"vr_timers": schema.SingleNestedAttribute{
								MarkdownDescription: "Vr timers",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"lsa_interval": schema.Float64Attribute{
										MarkdownDescription: "Lsa interval",
										Optional:            true,
									},
									"spf_calculation_delay": schema.Float64Attribute{
										MarkdownDescription: "Spf calculation delay",
										Optional:            true,
									},
								},
							},
						},
					},
					"ospfv3": schema.SingleNestedAttribute{
						MarkdownDescription: "Ospfv3",
						Optional:            true,
						Attributes: map[string]schema.Attribute{
							"allow_redist_default_route": schema.BoolAttribute{
								MarkdownDescription: "Allow redist default route",
								Optional:            true,
							},
							"area": schema.ListNestedAttribute{
								MarkdownDescription: "Area",
								Optional:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"authentication": schema.StringAttribute{
											MarkdownDescription: "Authentication",
											Optional:            true,
										},
										"interface": schema.ListNestedAttribute{
											MarkdownDescription: "Interface",
											Optional:            true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"authentication": schema.StringAttribute{
														MarkdownDescription: "Authentication",
														Optional:            true,
													},
													"bfd": schema.SingleNestedAttribute{
														MarkdownDescription: "Bfd",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"profile": schema.StringAttribute{
																MarkdownDescription: "Profile",
																Optional:            true,
															},
														},
													},
													"enable": schema.BoolAttribute{
														MarkdownDescription: "Enable",
														Optional:            true,
													},
													"instance_id": schema.Float64Attribute{
														MarkdownDescription: "Instance id",
														Optional:            true,
													},
													"link_type": schema.SingleNestedAttribute{
														MarkdownDescription: "Link type",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"broadcast": schema.SingleNestedAttribute{
																Validators: []validator.Object{
																	objectvalidator.ExactlyOneOf(
																		path.MatchRelative().AtParent().AtName("p2mp"),
																		path.MatchRelative().AtParent().AtName("p2p"),
																	),
																},
																MarkdownDescription: "Broadcast",
																Optional:            true,
																Attributes:          map[string]schema.Attribute{},
															},
															"p2mp": schema.SingleNestedAttribute{
																Validators: []validator.Object{
																	objectvalidator.ExactlyOneOf(
																		path.MatchRelative().AtParent().AtName("broadcast"),
																		path.MatchRelative().AtParent().AtName("p2p"),
																	),
																},
																MarkdownDescription: "P2mp",
																Optional:            true,
																Attributes: map[string]schema.Attribute{
																	"neighbor": schema.ListNestedAttribute{
																		MarkdownDescription: "Neighbor",
																		Optional:            true,
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					MarkdownDescription: "Name",
																					Required:            true,
																				},
																				"priority": schema.Float64Attribute{
																					MarkdownDescription: "Priority",
																					Optional:            true,
																				},
																			},
																		},
																	},
																},
															},
															"p2p": schema.SingleNestedAttribute{
																Validators: []validator.Object{
																	objectvalidator.ExactlyOneOf(
																		path.MatchRelative().AtParent().AtName("broadcast"),
																		path.MatchRelative().AtParent().AtName("p2mp"),
																	),
																},
																MarkdownDescription: "P2p",
																Optional:            true,
																Attributes:          map[string]schema.Attribute{},
															},
														},
													},
													"metric": schema.Float64Attribute{
														MarkdownDescription: "Metric",
														Optional:            true,
													},
													"mtu_ignore": schema.BoolAttribute{
														MarkdownDescription: "Mtu ignore",
														Optional:            true,
													},
													"name": schema.StringAttribute{
														MarkdownDescription: "Name",
														Required:            true,
													},
													"neighbor": schema.ListNestedAttribute{
														MarkdownDescription: "Neighbor",
														Optional:            true,
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	MarkdownDescription: "Name",
																	Required:            true,
																},
															},
														},
													},
													"passive": schema.BoolAttribute{
														MarkdownDescription: "Passive",
														Optional:            true,
													},
													"priority": schema.Float64Attribute{
														MarkdownDescription: "Priority",
														Optional:            true,
													},
													"timing": schema.StringAttribute{
														MarkdownDescription: "Timing",
														Optional:            true,
													},
													"vr_timing": schema.SingleNestedAttribute{
														MarkdownDescription: "Vr timing",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"dead_counts": schema.Float64Attribute{
																MarkdownDescription: "Dead counts",
																Optional:            true,
															},
															"gr_delay": schema.Float64Attribute{
																MarkdownDescription: "Gr delay",
																Optional:            true,
															},
															"hello_interval": schema.Float64Attribute{
																MarkdownDescription: "Hello interval",
																Optional:            true,
															},
															"retransmit_interval": schema.Float64Attribute{
																MarkdownDescription: "Retransmit interval",
																Optional:            true,
															},
															"transit_delay": schema.Float64Attribute{
																MarkdownDescription: "Transit delay",
																Optional:            true,
															},
														},
													},
												},
											},
										},
										"name": schema.StringAttribute{
											MarkdownDescription: "Name",
											Required:            true,
										},
										"range": schema.ListNestedAttribute{
											MarkdownDescription: "Range",
											Optional:            true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"advertise": schema.BoolAttribute{
														MarkdownDescription: "Advertise",
														Optional:            true,
													},
													"name": schema.StringAttribute{
														MarkdownDescription: "Name",
														Required:            true,
													},
												},
											},
										},
										"type": schema.SingleNestedAttribute{
											MarkdownDescription: "Type",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"normal": schema.SingleNestedAttribute{
													Validators: []validator.Object{
														objectvalidator.ExactlyOneOf(
															path.MatchRelative().AtParent().AtName("nssa"),
															path.MatchRelative().AtParent().AtName("stub"),
														),
													},
													MarkdownDescription: "Normal",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"abr": schema.SingleNestedAttribute{
															MarkdownDescription: "Abr",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"export_list": schema.StringAttribute{
																	MarkdownDescription: "Export list",
																	Optional:            true,
																},
																"import_list": schema.StringAttribute{
																	MarkdownDescription: "Import list",
																	Optional:            true,
																},
																"inbound_filter_list": schema.StringAttribute{
																	MarkdownDescription: "Inbound filter list",
																	Optional:            true,
																},
																"outbound_filter_list": schema.StringAttribute{
																	MarkdownDescription: "Outbound filter list",
																	Optional:            true,
																},
															},
														},
													},
												},
												"nssa": schema.SingleNestedAttribute{
													Validators: []validator.Object{
														objectvalidator.ExactlyOneOf(
															path.MatchRelative().AtParent().AtName("normal"),
															path.MatchRelative().AtParent().AtName("stub"),
														),
													},
													MarkdownDescription: "Nssa",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"abr": schema.SingleNestedAttribute{
															MarkdownDescription: "Abr",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"export_list": schema.StringAttribute{
																	MarkdownDescription: "Export list",
																	Optional:            true,
																},
																"import_list": schema.StringAttribute{
																	MarkdownDescription: "Import list",
																	Optional:            true,
																},
																"inbound_filter_list": schema.StringAttribute{
																	MarkdownDescription: "Inbound filter list",
																	Optional:            true,
																},
																"nssa_ext_range": schema.ListNestedAttribute{
																	MarkdownDescription: "Nssa ext range",
																	Optional:            true,
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"advertise": schema.SingleNestedAttribute{
																				MarkdownDescription: "Advertise",
																				Optional:            true,
																				Attributes:          map[string]schema.Attribute{},
																			},
																			"name": schema.StringAttribute{
																				MarkdownDescription: "Name",
																				Required:            true,
																			},
																			"route_tag": schema.Float64Attribute{
																				MarkdownDescription: "Route tag",
																				Optional:            true,
																			},
																			"suppress": schema.SingleNestedAttribute{
																				MarkdownDescription: "Suppress",
																				Optional:            true,
																				Attributes:          map[string]schema.Attribute{},
																			},
																		},
																	},
																},
																"outbound_filter_list": schema.StringAttribute{
																	MarkdownDescription: "Outbound filter list",
																	Optional:            true,
																},
															},
														},
														"accept_summary": schema.BoolAttribute{
															MarkdownDescription: "Accept summary",
															Optional:            true,
														},
														"default_information_originate": schema.SingleNestedAttribute{
															MarkdownDescription: "Default information originate",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"metric": schema.Float64Attribute{
																	MarkdownDescription: "Metric",
																	Optional:            true,
																},
																"metric_type": schema.StringAttribute{
																	Validators: []validator.String{
																		stringvalidator.OneOf("type-1", "type-2"),
																	},
																	MarkdownDescription: "Metric type",
																	Optional:            true,
																},
															},
														},
														"default_route": schema.SingleNestedAttribute{
															MarkdownDescription: "Default route",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"advertise": schema.SingleNestedAttribute{
																	Validators: []validator.Object{
																		objectvalidator.ExactlyOneOf(
																			path.MatchRelative().AtParent().AtName("disable"),
																		),
																	},
																	MarkdownDescription: "Advertise",
																	Optional:            true,
																	Attributes: map[string]schema.Attribute{
																		"metric": schema.Float64Attribute{
																			MarkdownDescription: "Metric",
																			Optional:            true,
																		},
																		"type": schema.StringAttribute{
																			Validators: []validator.String{
																				stringvalidator.OneOf("ext-1", "ext-2"),
																			},
																			MarkdownDescription: "Type",
																			Optional:            true,
																		},
																	},
																},
																"disable": schema.SingleNestedAttribute{
																	Validators: []validator.Object{
																		objectvalidator.ExactlyOneOf(
																			path.MatchRelative().AtParent().AtName("advertise"),
																		),
																	},
																	MarkdownDescription: "Disable",
																	Optional:            true,
																	Attributes:          map[string]schema.Attribute{},
																},
															},
														},
														"no_summary": schema.BoolAttribute{
															MarkdownDescription: "No summary",
															Optional:            true,
														},
														"nssa_ext_range": schema.ListNestedAttribute{
															MarkdownDescription: "Nssa ext range",
															Optional:            true,
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"advertise": schema.SingleNestedAttribute{
																		MarkdownDescription: "Advertise",
																		Optional:            true,
																		Attributes:          map[string]schema.Attribute{},
																	},
																	"name": schema.StringAttribute{
																		MarkdownDescription: "Name",
																		Required:            true,
																	},
																	"route_tag": schema.Float64Attribute{
																		MarkdownDescription: "Route tag",
																		Optional:            true,
																	},
																	"suppress": schema.SingleNestedAttribute{
																		MarkdownDescription: "Suppress",
																		Optional:            true,
																		Attributes:          map[string]schema.Attribute{},
																	},
																},
															},
														},
													},
												},
												"stub": schema.SingleNestedAttribute{
													Validators: []validator.Object{
														objectvalidator.ExactlyOneOf(
															path.MatchRelative().AtParent().AtName("normal"),
															path.MatchRelative().AtParent().AtName("nssa"),
														),
													},
													MarkdownDescription: "Stub",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"abr": schema.SingleNestedAttribute{
															MarkdownDescription: "Abr",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"export_list": schema.StringAttribute{
																	MarkdownDescription: "Export list",
																	Optional:            true,
																},
																"import_list": schema.StringAttribute{
																	MarkdownDescription: "Import list",
																	Optional:            true,
																},
																"inbound_filter_list": schema.StringAttribute{
																	MarkdownDescription: "Inbound filter list",
																	Optional:            true,
																},
																"outbound_filter_list": schema.StringAttribute{
																	MarkdownDescription: "Outbound filter list",
																	Optional:            true,
																},
															},
														},
														"accept_summary": schema.BoolAttribute{
															MarkdownDescription: "Accept summary",
															Optional:            true,
														},
														"default_route": schema.SingleNestedAttribute{
															MarkdownDescription: "Default route",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"advertise": schema.SingleNestedAttribute{
																	Validators: []validator.Object{
																		objectvalidator.ExactlyOneOf(
																			path.MatchRelative().AtParent().AtName("disable"),
																		),
																	},
																	MarkdownDescription: "Advertise",
																	Optional:            true,
																	Attributes: map[string]schema.Attribute{
																		"metric": schema.Float64Attribute{
																			MarkdownDescription: "Metric",
																			Optional:            true,
																		},
																	},
																},
																"disable": schema.SingleNestedAttribute{
																	Validators: []validator.Object{
																		objectvalidator.ExactlyOneOf(
																			path.MatchRelative().AtParent().AtName("advertise"),
																		),
																	},
																	MarkdownDescription: "Disable",
																	Optional:            true,
																	Attributes:          map[string]schema.Attribute{},
																},
															},
														},
														"default_route_metric": schema.Float64Attribute{
															MarkdownDescription: "Default route metric",
															Optional:            true,
														},
														"no_summary": schema.BoolAttribute{
															MarkdownDescription: "No summary",
															Optional:            true,
														},
													},
												},
											},
										},
										"virtual_link": schema.ListNestedAttribute{
											MarkdownDescription: "Virtual link",
											Optional:            true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"authentication": schema.StringAttribute{
														MarkdownDescription: "Authentication",
														Optional:            true,
													},
													"bfd": schema.SingleNestedAttribute{
														MarkdownDescription: "Bfd",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"profile": schema.StringAttribute{
																MarkdownDescription: "Profile",
																Optional:            true,
															},
														},
													},
													"enable": schema.BoolAttribute{
														MarkdownDescription: "Enable",
														Optional:            true,
													},
													"instance_id": schema.Float64Attribute{
														MarkdownDescription: "Instance id",
														Optional:            true,
													},
													"interface_id": schema.Float64Attribute{
														MarkdownDescription: "Interface id",
														Optional:            true,
													},
													"name": schema.StringAttribute{
														MarkdownDescription: "Name",
														Required:            true,
													},
													"neighbor_id": schema.StringAttribute{
														MarkdownDescription: "Neighbor id",
														Optional:            true,
													},
													"passive": schema.BoolAttribute{
														MarkdownDescription: "Passive",
														Optional:            true,
													},
													"timing": schema.StringAttribute{
														MarkdownDescription: "Timing",
														Optional:            true,
													},
													"transit_area_id": schema.StringAttribute{
														MarkdownDescription: "Transit area id",
														Optional:            true,
													},
													"vr_timing": schema.SingleNestedAttribute{
														MarkdownDescription: "Vr timing",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"dead_counts": schema.Float64Attribute{
																MarkdownDescription: "Dead counts",
																Optional:            true,
															},
															"hello_interval": schema.Float64Attribute{
																MarkdownDescription: "Hello interval",
																Optional:            true,
															},
															"retransmit_interval": schema.Float64Attribute{
																MarkdownDescription: "Retransmit interval",
																Optional:            true,
															},
															"transit_delay": schema.Float64Attribute{
																MarkdownDescription: "Transit delay",
																Optional:            true,
															},
														},
													},
												},
											},
										},
										"vr_range": schema.ListNestedAttribute{
											MarkdownDescription: "Vr range",
											Optional:            true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"advertise": schema.SingleNestedAttribute{
														MarkdownDescription: "Advertise",
														Optional:            true,
														Attributes:          map[string]schema.Attribute{},
													},
													"name": schema.StringAttribute{
														MarkdownDescription: "Name",
														Required:            true,
													},
													"suppress": schema.SingleNestedAttribute{
														MarkdownDescription: "Suppress",
														Optional:            true,
														Attributes:          map[string]schema.Attribute{},
													},
												},
											},
										},
									},
								},
							},
							"auth_profile": schema.ListNestedAttribute{
								MarkdownDescription: "Auth profile",
								Optional:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"ah": schema.SingleNestedAttribute{
											MarkdownDescription: "Ah",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"md5": schema.SingleNestedAttribute{
													Validators: []validator.Object{
														objectvalidator.ExactlyOneOf(
															path.MatchRelative().AtParent().AtName("sha1"),
															path.MatchRelative().AtParent().AtName("sha256"),
															path.MatchRelative().AtParent().AtName("sha384"),
															path.MatchRelative().AtParent().AtName("sha512"),
														),
													},
													MarkdownDescription: "Md5",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															MarkdownDescription: "Key",
															Optional:            true,
														},
													},
												},
												"sha1": schema.SingleNestedAttribute{
													Validators: []validator.Object{
														objectvalidator.ExactlyOneOf(
															path.MatchRelative().AtParent().AtName("md5"),
															path.MatchRelative().AtParent().AtName("sha256"),
															path.MatchRelative().AtParent().AtName("sha384"),
															path.MatchRelative().AtParent().AtName("sha512"),
														),
													},
													MarkdownDescription: "Sha1",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															MarkdownDescription: "Key",
															Optional:            true,
														},
													},
												},
												"sha256": schema.SingleNestedAttribute{
													Validators: []validator.Object{
														objectvalidator.ExactlyOneOf(
															path.MatchRelative().AtParent().AtName("md5"),
															path.MatchRelative().AtParent().AtName("sha1"),
															path.MatchRelative().AtParent().AtName("sha384"),
															path.MatchRelative().AtParent().AtName("sha512"),
														),
													},
													MarkdownDescription: "Sha256",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															MarkdownDescription: "Key",
															Optional:            true,
														},
													},
												},
												"sha384": schema.SingleNestedAttribute{
													Validators: []validator.Object{
														objectvalidator.ExactlyOneOf(
															path.MatchRelative().AtParent().AtName("md5"),
															path.MatchRelative().AtParent().AtName("sha1"),
															path.MatchRelative().AtParent().AtName("sha256"),
															path.MatchRelative().AtParent().AtName("sha512"),
														),
													},
													MarkdownDescription: "Sha384",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															MarkdownDescription: "Key",
															Optional:            true,
														},
													},
												},
												"sha512": schema.SingleNestedAttribute{
													Validators: []validator.Object{
														objectvalidator.ExactlyOneOf(
															path.MatchRelative().AtParent().AtName("md5"),
															path.MatchRelative().AtParent().AtName("sha1"),
															path.MatchRelative().AtParent().AtName("sha256"),
															path.MatchRelative().AtParent().AtName("sha384"),
														),
													},
													MarkdownDescription: "Sha512",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															MarkdownDescription: "Key",
															Optional:            true,
														},
													},
												},
											},
										},
										"esp": schema.SingleNestedAttribute{
											MarkdownDescription: "Esp",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"authentication": schema.SingleNestedAttribute{
													MarkdownDescription: "Authentication",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"md5": schema.SingleNestedAttribute{
															Validators: []validator.Object{
																objectvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("none"),
																	path.MatchRelative().AtParent().AtName("sha1"),
																	path.MatchRelative().AtParent().AtName("sha256"),
																	path.MatchRelative().AtParent().AtName("sha384"),
																	path.MatchRelative().AtParent().AtName("sha512"),
																),
															},
															MarkdownDescription: "Md5",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	MarkdownDescription: "Key",
																	Optional:            true,
																},
															},
														},
														"none": schema.SingleNestedAttribute{
															Validators: []validator.Object{
																objectvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("md5"),
																	path.MatchRelative().AtParent().AtName("sha1"),
																	path.MatchRelative().AtParent().AtName("sha256"),
																	path.MatchRelative().AtParent().AtName("sha384"),
																	path.MatchRelative().AtParent().AtName("sha512"),
																),
															},
															MarkdownDescription: "None",
															Optional:            true,
															Attributes:          map[string]schema.Attribute{},
														},
														"sha1": schema.SingleNestedAttribute{
															Validators: []validator.Object{
																objectvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("md5"),
																	path.MatchRelative().AtParent().AtName("none"),
																	path.MatchRelative().AtParent().AtName("sha256"),
																	path.MatchRelative().AtParent().AtName("sha384"),
																	path.MatchRelative().AtParent().AtName("sha512"),
																),
															},
															MarkdownDescription: "Sha1",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	MarkdownDescription: "Key",
																	Optional:            true,
																},
															},
														},
														"sha256": schema.SingleNestedAttribute{
															Validators: []validator.Object{
																objectvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("md5"),
																	path.MatchRelative().AtParent().AtName("none"),
																	path.MatchRelative().AtParent().AtName("sha1"),
																	path.MatchRelative().AtParent().AtName("sha384"),
																	path.MatchRelative().AtParent().AtName("sha512"),
																),
															},
															MarkdownDescription: "Sha256",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	MarkdownDescription: "Key",
																	Optional:            true,
																},
															},
														},
														"sha384": schema.SingleNestedAttribute{
															Validators: []validator.Object{
																objectvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("md5"),
																	path.MatchRelative().AtParent().AtName("none"),
																	path.MatchRelative().AtParent().AtName("sha1"),
																	path.MatchRelative().AtParent().AtName("sha256"),
																	path.MatchRelative().AtParent().AtName("sha512"),
																),
															},
															MarkdownDescription: "Sha384",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	MarkdownDescription: "Key",
																	Optional:            true,
																},
															},
														},
														"sha512": schema.SingleNestedAttribute{
															Validators: []validator.Object{
																objectvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("md5"),
																	path.MatchRelative().AtParent().AtName("none"),
																	path.MatchRelative().AtParent().AtName("sha1"),
																	path.MatchRelative().AtParent().AtName("sha256"),
																	path.MatchRelative().AtParent().AtName("sha384"),
																),
															},
															MarkdownDescription: "Sha512",
															Optional:            true,
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	MarkdownDescription: "Key",
																	Optional:            true,
																},
															},
														},
													},
												},
												"encryption": schema.SingleNestedAttribute{
													MarkdownDescription: "Encryption",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"algorithm": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.OneOf("3des", "aes-128-cbc", "aes-192-cbc", "aes-256-cbc", "null"),
															},
															MarkdownDescription: "Algorithm",
															Optional:            true,
														},
														"key": schema.StringAttribute{
															MarkdownDescription: "Key",
															Optional:            true,
														},
													},
												},
											},
										},
										"name": schema.StringAttribute{
											MarkdownDescription: "Name",
											Required:            true,
										},
										"spi": schema.StringAttribute{
											MarkdownDescription: "Spi",
											Optional:            true,
										},
									},
								},
							},
							"disable_transit_traffic": schema.BoolAttribute{
								MarkdownDescription: "Disable transit traffic",
								Optional:            true,
							},
							"enable": schema.BoolAttribute{
								MarkdownDescription: "Enable",
								Optional:            true,
							},
							"export_rules": schema.ListNestedAttribute{
								MarkdownDescription: "Export rules",
								Optional:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"metric": schema.Float64Attribute{
											MarkdownDescription: "Metric",
											Optional:            true,
										},
										"name": schema.StringAttribute{
											MarkdownDescription: "Name",
											Required:            true,
										},
										"new_path_type": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("ext-1", "ext-2"),
											},
											MarkdownDescription: "New path type",
											Optional:            true,
										},
										"new_tag": schema.StringAttribute{
											MarkdownDescription: "New tag",
											Optional:            true,
										},
									},
								},
							},
							"global_bfd": schema.SingleNestedAttribute{
								MarkdownDescription: "Global bfd",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"profile": schema.StringAttribute{
										MarkdownDescription: "Profile",
										Optional:            true,
									},
								},
							},
							"global_if_timer": schema.StringAttribute{
								MarkdownDescription: "Global if timer",
								Optional:            true,
							},
							"graceful_restart": schema.SingleNestedAttribute{
								MarkdownDescription: "Graceful restart",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"enable": schema.BoolAttribute{
										MarkdownDescription: "Enable",
										Optional:            true,
									},
									"grace_period": schema.Float64Attribute{
										MarkdownDescription: "Grace period",
										Optional:            true,
									},
									"helper_enable": schema.BoolAttribute{
										MarkdownDescription: "Helper enable",
										Optional:            true,
									},
									"max_neighbor_restart_time": schema.Float64Attribute{
										MarkdownDescription: "Max neighbor restart time",
										Optional:            true,
									},
									"strict__l_s_a_checking": schema.BoolAttribute{
										MarkdownDescription: "Strict l s a checking",
										Optional:            true,
									},
								},
							},
							"redistribution_profile": schema.StringAttribute{
								MarkdownDescription: "Redistribution profile",
								Optional:            true,
							},
							"reject_default_route": schema.BoolAttribute{
								MarkdownDescription: "Reject default route",
								Optional:            true,
							},
							"router_id": schema.StringAttribute{
								MarkdownDescription: "Router id",
								Optional:            true,
							},
							"spf_timer": schema.StringAttribute{
								MarkdownDescription: "Spf timer",
								Optional:            true,
							},
							"vr_timers": schema.SingleNestedAttribute{
								MarkdownDescription: "Vr timers",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"lsa_interval": schema.Float64Attribute{
										MarkdownDescription: "Lsa interval",
										Optional:            true,
									},
									"spf_calculation_delay": schema.Float64Attribute{
										MarkdownDescription: "Spf calculation delay",
										Optional:            true,
									},
								},
							},
						},
					},
					"rib_filter": schema.SingleNestedAttribute{
						MarkdownDescription: "Rib filter",
						Optional:            true,
						Attributes: map[string]schema.Attribute{
							"ipv4": schema.SingleNestedAttribute{
								MarkdownDescription: "Ipv4",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"bgp": schema.SingleNestedAttribute{
										MarkdownDescription: "Bgp",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"route_map": schema.StringAttribute{
												MarkdownDescription: "Route map",
												Optional:            true,
											},
										},
									},
									"ospf": schema.SingleNestedAttribute{
										MarkdownDescription: "Ospf",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"route_map": schema.StringAttribute{
												MarkdownDescription: "Route map",
												Optional:            true,
											},
										},
									},
									"rip": schema.SingleNestedAttribute{
										MarkdownDescription: "Rip",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"route_map": schema.StringAttribute{
												MarkdownDescription: "Route map",
												Optional:            true,
											},
										},
									},
									"static": schema.SingleNestedAttribute{
										MarkdownDescription: "Static",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"route_map": schema.StringAttribute{
												MarkdownDescription: "Route map",
												Optional:            true,
											},
										},
									},
								},
							},
							"ipv6": schema.SingleNestedAttribute{
								MarkdownDescription: "Ipv6",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"bgp": schema.SingleNestedAttribute{
										MarkdownDescription: "Bgp",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"route_map": schema.StringAttribute{
												MarkdownDescription: "Route map",
												Optional:            true,
											},
										},
									},
									"ospfv3": schema.SingleNestedAttribute{
										MarkdownDescription: "Ospfv3",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"route_map": schema.StringAttribute{
												MarkdownDescription: "Route map",
												Optional:            true,
											},
										},
									},
									"static": schema.SingleNestedAttribute{
										MarkdownDescription: "Static",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"route_map": schema.StringAttribute{
												MarkdownDescription: "Route map",
												Optional:            true,
											},
										},
									},
								},
							},
						},
					},
					"rip": schema.SingleNestedAttribute{
						MarkdownDescription: "Rip",
						Optional:            true,
						Attributes: map[string]schema.Attribute{
							"auth_profile": schema.StringAttribute{
								MarkdownDescription: "Auth profile",
								Optional:            true,
							},
							"default_information_originate": schema.BoolAttribute{
								MarkdownDescription: "Default information originate",
								Optional:            true,
							},
							"enable": schema.BoolAttribute{
								MarkdownDescription: "Enable",
								Optional:            true,
							},
							"global_bfd": schema.SingleNestedAttribute{
								MarkdownDescription: "Global bfd",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"profile": schema.StringAttribute{
										MarkdownDescription: "Profile",
										Optional:            true,
									},
								},
							},
							"global_inbound_distribute_list": schema.SingleNestedAttribute{
								MarkdownDescription: "Global inbound distribute list",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"access_list": schema.StringAttribute{
										MarkdownDescription: "Access list",
										Optional:            true,
									},
								},
							},
							"global_outbound_distribute_list": schema.SingleNestedAttribute{
								MarkdownDescription: "Global outbound distribute list",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"access_list": schema.StringAttribute{
										MarkdownDescription: "Access list",
										Optional:            true,
									},
								},
							},
							"global_timer": schema.StringAttribute{
								MarkdownDescription: "Global timer",
								Optional:            true,
							},
							"interface": schema.ListNestedAttribute{
								MarkdownDescription: "Interface",
								Optional:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"authentication": schema.StringAttribute{
											MarkdownDescription: "Authentication",
											Optional:            true,
										},
										"bfd": schema.SingleNestedAttribute{
											MarkdownDescription: "Bfd",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"profile": schema.StringAttribute{
													MarkdownDescription: "Profile",
													Optional:            true,
												},
											},
										},
										"enable": schema.BoolAttribute{
											MarkdownDescription: "Enable",
											Optional:            true,
										},
										"interface_inbound_distribute_list": schema.SingleNestedAttribute{
											MarkdownDescription: "Interface inbound distribute list",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"access_list": schema.StringAttribute{
													MarkdownDescription: "Access list",
													Optional:            true,
												},
												"metric": schema.Float64Attribute{
													MarkdownDescription: "Metric",
													Optional:            true,
												},
											},
										},
										"interface_outbound_distribute_list": schema.SingleNestedAttribute{
											MarkdownDescription: "Interface outbound distribute list",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"access_list": schema.StringAttribute{
													MarkdownDescription: "Access list",
													Optional:            true,
												},
												"metric": schema.Float64Attribute{
													MarkdownDescription: "Metric",
													Optional:            true,
												},
											},
										},
										"mode": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("active", "passive", "send-only"),
											},
											MarkdownDescription: "Mode",
											Optional:            true,
										},
										"name": schema.StringAttribute{
											MarkdownDescription: "Name",
											Required:            true,
										},
										"split_horizon": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("split-horizon", "no-split-horizon", "no-split-horizon-with-poison-reverse"),
											},
											MarkdownDescription: "Split horizon",
											Optional:            true,
										},
									},
								},
							},
							"redistribution_profile": schema.StringAttribute{
								MarkdownDescription: "Redistribution profile",
								Optional:            true,
							},
						},
					},
					"routing_table": schema.SingleNestedAttribute{
						MarkdownDescription: "Routing table",
						Optional:            true,
						Attributes: map[string]schema.Attribute{
							"ip": schema.SingleNestedAttribute{
								MarkdownDescription: "Ip",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"static_route": schema.ListNestedAttribute{
										MarkdownDescription: "Static route",
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"admin_dist": schema.Float64Attribute{
													MarkdownDescription: "Admin dist",
													Optional:            true,
												},
												"bfd": schema.SingleNestedAttribute{
													MarkdownDescription: "Bfd",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"profile": schema.StringAttribute{
															MarkdownDescription: "Profile",
															Optional:            true,
														},
													},
												},
												"destination": schema.StringAttribute{
													MarkdownDescription: "Destination",
													Optional:            true,
												},
												"interface": schema.StringAttribute{
													MarkdownDescription: "Interface",
													Optional:            true,
												},
												"metric": schema.Float64Attribute{
													MarkdownDescription: "Metric",
													Optional:            true,
												},
												"name": schema.StringAttribute{
													MarkdownDescription: "Name",
													Required:            true,
												},
												"nexthop": schema.SingleNestedAttribute{
													MarkdownDescription: "Nexthop",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"discard": schema.SingleNestedAttribute{
															Validators: []validator.Object{
																objectvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("fqdn"),
																	path.MatchRelative().AtParent().AtName("ipv6_address"),
																	path.MatchRelative().AtParent().AtName("next_lr"),
																	path.MatchRelative().AtParent().AtName("next_vr"),
																	path.MatchRelative().AtParent().AtName("receive"),
																	path.MatchRelative().AtParent().AtName("tunnel"),
																),
															},
															MarkdownDescription: "Discard",
															Optional:            true,
															Attributes:          map[string]schema.Attribute{},
														},
														"fqdn": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("discard"),
																	path.MatchRelative().AtParent().AtName("ipv6_address"),
																	path.MatchRelative().AtParent().AtName("next_lr"),
																	path.MatchRelative().AtParent().AtName("next_vr"),
																	path.MatchRelative().AtParent().AtName("receive"),
																	path.MatchRelative().AtParent().AtName("tunnel"),
																),
															},
															MarkdownDescription: "Fqdn",
															Optional:            true,
														},
														"ipv6_address": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("discard"),
																	path.MatchRelative().AtParent().AtName("fqdn"),
																	path.MatchRelative().AtParent().AtName("next_lr"),
																	path.MatchRelative().AtParent().AtName("next_vr"),
																	path.MatchRelative().AtParent().AtName("receive"),
																	path.MatchRelative().AtParent().AtName("tunnel"),
																),
															},
															MarkdownDescription: "Ipv6 address",
															Optional:            true,
														},
														"next_lr": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("discard"),
																	path.MatchRelative().AtParent().AtName("fqdn"),
																	path.MatchRelative().AtParent().AtName("ipv6_address"),
																	path.MatchRelative().AtParent().AtName("next_vr"),
																	path.MatchRelative().AtParent().AtName("receive"),
																	path.MatchRelative().AtParent().AtName("tunnel"),
																),
															},
															MarkdownDescription: "Next lr",
															Optional:            true,
														},
														"next_vr": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("discard"),
																	path.MatchRelative().AtParent().AtName("fqdn"),
																	path.MatchRelative().AtParent().AtName("ipv6_address"),
																	path.MatchRelative().AtParent().AtName("next_lr"),
																	path.MatchRelative().AtParent().AtName("receive"),
																	path.MatchRelative().AtParent().AtName("tunnel"),
																),
															},
															MarkdownDescription: "Next vr",
															Optional:            true,
														},
														"receive": schema.SingleNestedAttribute{
															Validators: []validator.Object{
																objectvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("discard"),
																	path.MatchRelative().AtParent().AtName("fqdn"),
																	path.MatchRelative().AtParent().AtName("ipv6_address"),
																	path.MatchRelative().AtParent().AtName("next_lr"),
																	path.MatchRelative().AtParent().AtName("next_vr"),
																	path.MatchRelative().AtParent().AtName("tunnel"),
																),
															},
															MarkdownDescription: "Receive",
															Optional:            true,
															Attributes:          map[string]schema.Attribute{},
														},
														"tunnel": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("discard"),
																	path.MatchRelative().AtParent().AtName("fqdn"),
																	path.MatchRelative().AtParent().AtName("ipv6_address"),
																	path.MatchRelative().AtParent().AtName("next_lr"),
																	path.MatchRelative().AtParent().AtName("next_vr"),
																	path.MatchRelative().AtParent().AtName("receive"),
																),
															},
															MarkdownDescription: "Tunnel",
															Optional:            true,
														},
													},
												},
												"path_monitor": schema.SingleNestedAttribute{
													MarkdownDescription: "Path monitor",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"enable": schema.BoolAttribute{
															MarkdownDescription: "Enable",
															Optional:            true,
														},
														"failure_condition": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.OneOf("any", "all"),
															},
															MarkdownDescription: "Failure condition",
															Optional:            true,
														},
														"hold_time": schema.Float64Attribute{
															MarkdownDescription: "Hold time",
															Optional:            true,
														},
														"monitor_destinations": schema.ListNestedAttribute{
															MarkdownDescription: "Monitor destinations",
															Optional:            true,
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"count": schema.Float64Attribute{
																		MarkdownDescription: "Count",
																		Optional:            true,
																	},
																	"destination": schema.StringAttribute{
																		MarkdownDescription: "Destination",
																		Optional:            true,
																	},
																	"destination_fqdn": schema.StringAttribute{
																		MarkdownDescription: "Destination fqdn",
																		Optional:            true,
																	},
																	"enable": schema.BoolAttribute{
																		MarkdownDescription: "Enable",
																		Optional:            true,
																	},
																	"interval": schema.Float64Attribute{
																		MarkdownDescription: "Interval",
																		Optional:            true,
																	},
																	"name": schema.StringAttribute{
																		MarkdownDescription: "Name",
																		Required:            true,
																	},
																	"source": schema.StringAttribute{
																		MarkdownDescription: "Source",
																		Optional:            true,
																	},
																},
															},
														},
													},
												},
												"route_table": schema.SingleNestedAttribute{
													MarkdownDescription: "Route table",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"both": schema.SingleNestedAttribute{
															Validators: []validator.Object{
																objectvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("multicast"),
																	path.MatchRelative().AtParent().AtName("no_install"),
																	path.MatchRelative().AtParent().AtName("unicast"),
																),
															},
															MarkdownDescription: "Both",
															Optional:            true,
															Attributes:          map[string]schema.Attribute{},
														},
														"multicast": schema.SingleNestedAttribute{
															Validators: []validator.Object{
																objectvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("both"),
																	path.MatchRelative().AtParent().AtName("no_install"),
																	path.MatchRelative().AtParent().AtName("unicast"),
																),
															},
															MarkdownDescription: "Multicast",
															Optional:            true,
															Attributes:          map[string]schema.Attribute{},
														},
														"no_install": schema.SingleNestedAttribute{
															Validators: []validator.Object{
																objectvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("both"),
																	path.MatchRelative().AtParent().AtName("multicast"),
																	path.MatchRelative().AtParent().AtName("unicast"),
																),
															},
															MarkdownDescription: "No install",
															Optional:            true,
															Attributes:          map[string]schema.Attribute{},
														},
														"unicast": schema.SingleNestedAttribute{
															Validators: []validator.Object{
																objectvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("both"),
																	path.MatchRelative().AtParent().AtName("multicast"),
																	path.MatchRelative().AtParent().AtName("no_install"),
																),
															},
															MarkdownDescription: "Unicast",
															Optional:            true,
															Attributes:          map[string]schema.Attribute{},
														},
													},
												},
											},
										},
									},
								},
							},
							"ipv6": schema.SingleNestedAttribute{
								MarkdownDescription: "Ipv6",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"static_route": schema.ListNestedAttribute{
										MarkdownDescription: "Static route",
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"admin_dist": schema.Float64Attribute{
													MarkdownDescription: "Admin dist",
													Optional:            true,
												},
												"bfd": schema.SingleNestedAttribute{
													MarkdownDescription: "Bfd",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"profile": schema.StringAttribute{
															MarkdownDescription: "Profile",
															Optional:            true,
														},
													},
												},
												"destination": schema.StringAttribute{
													MarkdownDescription: "Destination",
													Optional:            true,
												},
												"interface": schema.StringAttribute{
													MarkdownDescription: "Interface",
													Optional:            true,
												},
												"metric": schema.Float64Attribute{
													MarkdownDescription: "Metric",
													Optional:            true,
												},
												"name": schema.StringAttribute{
													MarkdownDescription: "Name",
													Required:            true,
												},
												"nexthop": schema.SingleNestedAttribute{
													MarkdownDescription: "Nexthop",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"discard": schema.SingleNestedAttribute{
															Validators: []validator.Object{
																objectvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("fqdn"),
																	path.MatchRelative().AtParent().AtName("ipv6_address"),
																	path.MatchRelative().AtParent().AtName("next_lr"),
																	path.MatchRelative().AtParent().AtName("next_vr"),
																	path.MatchRelative().AtParent().AtName("receive"),
																	path.MatchRelative().AtParent().AtName("tunnel"),
																),
															},
															MarkdownDescription: "Discard",
															Optional:            true,
															Attributes:          map[string]schema.Attribute{},
														},
														"fqdn": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("discard"),
																	path.MatchRelative().AtParent().AtName("ipv6_address"),
																	path.MatchRelative().AtParent().AtName("next_lr"),
																	path.MatchRelative().AtParent().AtName("next_vr"),
																	path.MatchRelative().AtParent().AtName("receive"),
																	path.MatchRelative().AtParent().AtName("tunnel"),
																),
															},
															MarkdownDescription: "Fqdn",
															Optional:            true,
														},
														"ipv6_address": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("discard"),
																	path.MatchRelative().AtParent().AtName("fqdn"),
																	path.MatchRelative().AtParent().AtName("next_lr"),
																	path.MatchRelative().AtParent().AtName("next_vr"),
																	path.MatchRelative().AtParent().AtName("receive"),
																	path.MatchRelative().AtParent().AtName("tunnel"),
																),
															},
															MarkdownDescription: "Ipv6 address",
															Optional:            true,
														},
														"next_lr": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("discard"),
																	path.MatchRelative().AtParent().AtName("fqdn"),
																	path.MatchRelative().AtParent().AtName("ipv6_address"),
																	path.MatchRelative().AtParent().AtName("next_vr"),
																	path.MatchRelative().AtParent().AtName("receive"),
																	path.MatchRelative().AtParent().AtName("tunnel"),
																),
															},
															MarkdownDescription: "Next lr",
															Optional:            true,
														},
														"next_vr": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("discard"),
																	path.MatchRelative().AtParent().AtName("fqdn"),
																	path.MatchRelative().AtParent().AtName("ipv6_address"),
																	path.MatchRelative().AtParent().AtName("next_lr"),
																	path.MatchRelative().AtParent().AtName("receive"),
																	path.MatchRelative().AtParent().AtName("tunnel"),
																),
															},
															MarkdownDescription: "Next vr",
															Optional:            true,
														},
														"receive": schema.SingleNestedAttribute{
															Validators: []validator.Object{
																objectvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("discard"),
																	path.MatchRelative().AtParent().AtName("fqdn"),
																	path.MatchRelative().AtParent().AtName("ipv6_address"),
																	path.MatchRelative().AtParent().AtName("next_lr"),
																	path.MatchRelative().AtParent().AtName("next_vr"),
																	path.MatchRelative().AtParent().AtName("tunnel"),
																),
															},
															MarkdownDescription: "Receive",
															Optional:            true,
															Attributes:          map[string]schema.Attribute{},
														},
														"tunnel": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("discard"),
																	path.MatchRelative().AtParent().AtName("fqdn"),
																	path.MatchRelative().AtParent().AtName("ipv6_address"),
																	path.MatchRelative().AtParent().AtName("next_lr"),
																	path.MatchRelative().AtParent().AtName("next_vr"),
																	path.MatchRelative().AtParent().AtName("receive"),
																),
															},
															MarkdownDescription: "Tunnel",
															Optional:            true,
														},
													},
												},
												"option": schema.SingleNestedAttribute{
													MarkdownDescription: "Option",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"passive": schema.SingleNestedAttribute{
															MarkdownDescription: "Passive",
															Optional:            true,
															Attributes:          map[string]schema.Attribute{},
														},
													},
												},
												"path_monitor": schema.SingleNestedAttribute{
													MarkdownDescription: "Path monitor",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"enable": schema.BoolAttribute{
															MarkdownDescription: "Enable",
															Optional:            true,
														},
														"failure_condition": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.OneOf("any", "all"),
															},
															MarkdownDescription: "Failure condition",
															Optional:            true,
														},
														"hold_time": schema.Float64Attribute{
															MarkdownDescription: "Hold time",
															Optional:            true,
														},
														"monitor_destinations": schema.ListNestedAttribute{
															MarkdownDescription: "Monitor destinations",
															Optional:            true,
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"count": schema.Float64Attribute{
																		MarkdownDescription: "Count",
																		Optional:            true,
																	},
																	"destination": schema.StringAttribute{
																		MarkdownDescription: "Destination",
																		Optional:            true,
																	},
																	"destination_fqdn": schema.StringAttribute{
																		MarkdownDescription: "Destination fqdn",
																		Optional:            true,
																	},
																	"enable": schema.BoolAttribute{
																		MarkdownDescription: "Enable",
																		Optional:            true,
																	},
																	"interval": schema.Float64Attribute{
																		MarkdownDescription: "Interval",
																		Optional:            true,
																	},
																	"name": schema.StringAttribute{
																		MarkdownDescription: "Name",
																		Required:            true,
																	},
																	"source": schema.StringAttribute{
																		MarkdownDescription: "Source",
																		Optional:            true,
																	},
																},
															},
														},
													},
												},
												"route_table": schema.SingleNestedAttribute{
													MarkdownDescription: "Route table",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"both": schema.SingleNestedAttribute{
															MarkdownDescription: "Both",
															Optional:            true,
															Attributes:          map[string]schema.Attribute{},
														},
														"multicast": schema.SingleNestedAttribute{
															MarkdownDescription: "Multicast",
															Optional:            true,
															Attributes:          map[string]schema.Attribute{},
														},
														"no_install": schema.SingleNestedAttribute{
															MarkdownDescription: "No install",
															Optional:            true,
															Attributes:          map[string]schema.Attribute{},
														},
														"unicast": schema.SingleNestedAttribute{
															MarkdownDescription: "Unicast",
															Optional:            true,
															Attributes:          map[string]schema.Attribute{},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
					"sdwan_type": schema.StringAttribute{
						MarkdownDescription: "Sdwan type",
						Optional:            true,
					},
					"vr_admin_dists": schema.SingleNestedAttribute{
						MarkdownDescription: "Vr admin dists",
						Optional:            true,
						Attributes: map[string]schema.Attribute{
							"ebgp": schema.Float64Attribute{
								MarkdownDescription: "Ebgp",
								Optional:            true,
							},
							"ibgp": schema.Float64Attribute{
								MarkdownDescription: "Ibgp",
								Optional:            true,
							},
							"ospf_ext": schema.Float64Attribute{
								MarkdownDescription: "Ospf ext",
								Optional:            true,
							},
							"ospf_int": schema.Float64Attribute{
								MarkdownDescription: "Ospf int",
								Optional:            true,
							},
							"ospfv3_ext": schema.Float64Attribute{
								MarkdownDescription: "Ospfv3 ext",
								Optional:            true,
							},
							"ospfv3_int": schema.Float64Attribute{
								MarkdownDescription: "Ospfv3 int",
								Optional:            true,
							},
							"rip": schema.Float64Attribute{
								MarkdownDescription: "Rip",
								Optional:            true,
							},
							"static": schema.Float64Attribute{
								MarkdownDescription: "Static",
								Optional:            true,
							},
							"static_ipv6": schema.Float64Attribute{
								MarkdownDescription: "Static ipv6",
								Optional:            true,
							},
						},
					},
					"zone_name": schema.StringAttribute{
						MarkdownDescription: "Zone name",
						Optional:            true,
					},
				},
			},
		},
	},
}

// LogicalRoutersDataSourceSchema defines the schema for LogicalRouters data source
var LogicalRoutersDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "LogicalRouter data source",
	Attributes: map[string]dsschema.Attribute{
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
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Name",
			Optional:            true,
			Computed:            true,
		},
		"routing_stack": dsschema.StringAttribute{
			MarkdownDescription: "Routing stack",
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
		"vrf": dsschema.ListNestedAttribute{
			MarkdownDescription: "Vrf",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"admin_dists": dsschema.SingleNestedAttribute{
						MarkdownDescription: "Admin dists",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"bgp_external": dsschema.Float64Attribute{
								MarkdownDescription: "Bgp external",
								Computed:            true,
							},
							"bgp_internal": dsschema.Float64Attribute{
								MarkdownDescription: "Bgp internal",
								Computed:            true,
							},
							"bgp_local": dsschema.Float64Attribute{
								MarkdownDescription: "Bgp local",
								Computed:            true,
							},
							"ospf_ext": dsschema.Float64Attribute{
								MarkdownDescription: "Ospf ext",
								Computed:            true,
							},
							"ospf_inter": dsschema.Float64Attribute{
								MarkdownDescription: "Ospf inter",
								Computed:            true,
							},
							"ospf_intra": dsschema.Float64Attribute{
								MarkdownDescription: "Ospf intra",
								Computed:            true,
							},
							"ospfv3_ext": dsschema.Float64Attribute{
								MarkdownDescription: "Ospfv3 ext",
								Computed:            true,
							},
							"ospfv3_inter": dsschema.Float64Attribute{
								MarkdownDescription: "Ospfv3 inter",
								Computed:            true,
							},
							"ospfv3_intra": dsschema.Float64Attribute{
								MarkdownDescription: "Ospfv3 intra",
								Computed:            true,
							},
							"rip": dsschema.Float64Attribute{
								MarkdownDescription: "Rip",
								Computed:            true,
							},
							"static": dsschema.Float64Attribute{
								MarkdownDescription: "Static",
								Computed:            true,
							},
							"static_ipv6": dsschema.Float64Attribute{
								MarkdownDescription: "Static ipv6",
								Computed:            true,
							},
						},
					},
					"bgp": dsschema.SingleNestedAttribute{
						MarkdownDescription: "Bgp",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"advertise_network": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Advertise network",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"ipv4": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Ipv4",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"network": dsschema.ListNestedAttribute{
												MarkdownDescription: "Network",
												Computed:            true,
												NestedObject: dsschema.NestedAttributeObject{
													Attributes: map[string]dsschema.Attribute{
														"backdoor": dsschema.BoolAttribute{
															MarkdownDescription: "Backdoor",
															Computed:            true,
														},
														"multicast": dsschema.BoolAttribute{
															MarkdownDescription: "Multicast",
															Computed:            true,
														},
														"name": dsschema.StringAttribute{
															MarkdownDescription: "Name",
															Computed:            true,
														},
														"unicast": dsschema.BoolAttribute{
															MarkdownDescription: "Unicast",
															Computed:            true,
														},
													},
												},
											},
										},
									},
									"ipv6": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Ipv6",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"network": dsschema.ListNestedAttribute{
												MarkdownDescription: "Network",
												Computed:            true,
												NestedObject: dsschema.NestedAttributeObject{
													Attributes: map[string]dsschema.Attribute{
														"name": dsschema.StringAttribute{
															MarkdownDescription: "Name",
															Computed:            true,
														},
														"unicast": dsschema.BoolAttribute{
															MarkdownDescription: "Unicast",
															Computed:            true,
														},
													},
												},
											},
										},
									},
								},
							},
							"aggregate": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Aggregate",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"aggregate_med": dsschema.BoolAttribute{
										MarkdownDescription: "Aggregate med",
										Computed:            true,
									},
								},
							},
							"aggregate_routes": dsschema.ListNestedAttribute{
								MarkdownDescription: "Aggregate routes",
								Computed:            true,
								NestedObject: dsschema.NestedAttributeObject{
									Attributes: map[string]dsschema.Attribute{
										"as_set": dsschema.BoolAttribute{
											MarkdownDescription: "As set",
											Computed:            true,
										},
										"description": dsschema.StringAttribute{
											MarkdownDescription: "Description",
											Computed:            true,
										},
										"enable": dsschema.BoolAttribute{
											MarkdownDescription: "Enable",
											Computed:            true,
										},
										"name": dsschema.StringAttribute{
											MarkdownDescription: "Name",
											Computed:            true,
										},
										"same_med": dsschema.BoolAttribute{
											MarkdownDescription: "Same med",
											Computed:            true,
										},
										"summary_only": dsschema.BoolAttribute{
											MarkdownDescription: "Summary only",
											Computed:            true,
										},
										"type": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Type",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"ipv4": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Ipv4",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"attribute_map": dsschema.StringAttribute{
															MarkdownDescription: "Attribute map",
															Computed:            true,
														},
														"summary_prefix": dsschema.StringAttribute{
															MarkdownDescription: "Summary prefix",
															Computed:            true,
														},
														"suppress_map": dsschema.StringAttribute{
															MarkdownDescription: "Suppress map",
															Computed:            true,
														},
													},
												},
												"ipv6": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Ipv6",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"attribute_map": dsschema.StringAttribute{
															MarkdownDescription: "Attribute map",
															Computed:            true,
														},
														"summary_prefix": dsschema.StringAttribute{
															MarkdownDescription: "Summary prefix",
															Computed:            true,
														},
														"suppress_map": dsschema.StringAttribute{
															MarkdownDescription: "Suppress map",
															Computed:            true,
														},
													},
												},
											},
										},
									},
								},
							},
							"allow_redist_default_route": dsschema.BoolAttribute{
								MarkdownDescription: "Allow redist default route",
								Computed:            true,
							},
							"always_advertise_network_route": dsschema.BoolAttribute{
								MarkdownDescription: "Always advertise network route",
								Computed:            true,
							},
							"as_format": dsschema.StringAttribute{
								MarkdownDescription: "As format",
								Computed:            true,
							},
							"confederation_member_as": dsschema.StringAttribute{
								MarkdownDescription: "Confederation member as",
								Computed:            true,
							},
							"default_local_preference": dsschema.Float64Attribute{
								MarkdownDescription: "Default local preference",
								Computed:            true,
							},
							"ecmp_multi_as": dsschema.BoolAttribute{
								MarkdownDescription: "Ecmp multi as",
								Computed:            true,
							},
							"enable": dsschema.BoolAttribute{
								MarkdownDescription: "Enable",
								Computed:            true,
							},
							"enforce_first_as": dsschema.BoolAttribute{
								MarkdownDescription: "Enforce first as",
								Computed:            true,
							},
							"fast_external_failover": dsschema.BoolAttribute{
								MarkdownDescription: "Fast external failover",
								Computed:            true,
							},
							"global_bfd": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Global bfd",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"profile": dsschema.StringAttribute{
										MarkdownDescription: "Profile",
										Computed:            true,
									},
								},
							},
							"graceful_restart": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Graceful restart",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"enable": dsschema.BoolAttribute{
										MarkdownDescription: "Enable",
										Computed:            true,
									},
									"local_restart_time": dsschema.Float64Attribute{
										MarkdownDescription: "Local restart time",
										Computed:            true,
									},
									"max_peer_restart_time": dsschema.Float64Attribute{
										MarkdownDescription: "Max peer restart time",
										Computed:            true,
									},
									"stale_route_time": dsschema.Float64Attribute{
										MarkdownDescription: "Stale route time",
										Computed:            true,
									},
								},
							},
							"graceful_shutdown": dsschema.BoolAttribute{
								MarkdownDescription: "Graceful shutdown",
								Computed:            true,
							},
							"install_route": dsschema.BoolAttribute{
								MarkdownDescription: "Install route",
								Computed:            true,
							},
							"local_as": dsschema.StringAttribute{
								MarkdownDescription: "Local as",
								Computed:            true,
							},
							"med": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Med",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"always_compare_med": dsschema.BoolAttribute{
										MarkdownDescription: "Always compare med",
										Computed:            true,
									},
									"deterministic_med_comparison": dsschema.BoolAttribute{
										MarkdownDescription: "Deterministic med comparison",
										Computed:            true,
									},
								},
							},
							"peer_group": dsschema.ListNestedAttribute{
								MarkdownDescription: "Peer group",
								Computed:            true,
								NestedObject: dsschema.NestedAttributeObject{
									Attributes: map[string]dsschema.Attribute{
										"address_family": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Address family",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"ipv4": dsschema.StringAttribute{
													MarkdownDescription: "Ipv4",
													Computed:            true,
												},
												"ipv6": dsschema.StringAttribute{
													MarkdownDescription: "Ipv6",
													Computed:            true,
												},
											},
										},
										"aggregated_confed_as_path": dsschema.BoolAttribute{
											MarkdownDescription: "Aggregated confed as path",
											Computed:            true,
										},
										"connection_options": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Connection options",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"authentication": dsschema.StringAttribute{
													MarkdownDescription: "Authentication",
													Computed:            true,
												},
												"dampening": dsschema.StringAttribute{
													MarkdownDescription: "Dampening",
													Computed:            true,
												},
												"multihop": dsschema.Float64Attribute{
													MarkdownDescription: "Multihop",
													Computed:            true,
												},
												"timers": dsschema.StringAttribute{
													MarkdownDescription: "Timers",
													Computed:            true,
												},
											},
										},
										"enable": dsschema.BoolAttribute{
											MarkdownDescription: "Enable",
											Computed:            true,
										},
										"filtering_profile": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Filtering profile",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"ipv4": dsschema.StringAttribute{
													MarkdownDescription: "Ipv4",
													Computed:            true,
												},
												"ipv6": dsschema.StringAttribute{
													MarkdownDescription: "Ipv6",
													Computed:            true,
												},
											},
										},
										"name": dsschema.StringAttribute{
											MarkdownDescription: "Name",
											Computed:            true,
										},
										"peer": dsschema.ListNestedAttribute{
											MarkdownDescription: "Peer",
											Computed:            true,
											NestedObject: dsschema.NestedAttributeObject{
												Attributes: map[string]dsschema.Attribute{
													"bfd": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Bfd",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"multihop": dsschema.SingleNestedAttribute{
																MarkdownDescription: "Multihop",
																Computed:            true,
																Attributes: map[string]dsschema.Attribute{
																	"min_received_ttl": dsschema.Float64Attribute{
																		MarkdownDescription: "Min received ttl",
																		Computed:            true,
																	},
																},
															},
															"profile": dsschema.StringAttribute{
																MarkdownDescription: "Profile",
																Computed:            true,
															},
														},
													},
													"connection_options": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Connection options",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"authentication": dsschema.StringAttribute{
																MarkdownDescription: "Authentication",
																Computed:            true,
															},
															"dampening": dsschema.StringAttribute{
																MarkdownDescription: "Dampening",
																Computed:            true,
															},
															"hold_time": dsschema.StringAttribute{
																MarkdownDescription: "Hold time",
																Computed:            true,
															},
															"idle_hold_time": dsschema.Float64Attribute{
																MarkdownDescription: "Idle hold time",
																Computed:            true,
															},
															"incoming_bgp_connection": dsschema.SingleNestedAttribute{
																MarkdownDescription: "Incoming bgp connection",
																Computed:            true,
																Attributes: map[string]dsschema.Attribute{
																	"allow": dsschema.BoolAttribute{
																		MarkdownDescription: "Allow",
																		Computed:            true,
																	},
																	"remote_port": dsschema.Float64Attribute{
																		MarkdownDescription: "Remote port",
																		Computed:            true,
																	},
																},
															},
															"keep_alive_interval": dsschema.StringAttribute{
																MarkdownDescription: "Keep alive interval",
																Computed:            true,
															},
															"max_prefixes": dsschema.StringAttribute{
																MarkdownDescription: "Max prefixes",
																Computed:            true,
															},
															"min_route_adv_interval": dsschema.Float64Attribute{
																MarkdownDescription: "Min route adv interval",
																Computed:            true,
															},
															"multihop": dsschema.StringAttribute{
																MarkdownDescription: "Multihop",
																Computed:            true,
															},
															"open_delay_time": dsschema.Float64Attribute{
																MarkdownDescription: "Open delay time",
																Computed:            true,
															},
															"outgoing_bgp_connection": dsschema.SingleNestedAttribute{
																MarkdownDescription: "Outgoing bgp connection",
																Computed:            true,
																Attributes: map[string]dsschema.Attribute{
																	"allow": dsschema.BoolAttribute{
																		MarkdownDescription: "Allow",
																		Computed:            true,
																	},
																	"local_port": dsschema.Float64Attribute{
																		MarkdownDescription: "Local port",
																		Computed:            true,
																	},
																},
															},
															"timers": dsschema.StringAttribute{
																MarkdownDescription: "Timers",
																Computed:            true,
															},
														},
													},
													"enable": dsschema.BoolAttribute{
														MarkdownDescription: "Enable",
														Computed:            true,
													},
													"enable_mp_bgp": dsschema.BoolAttribute{
														MarkdownDescription: "Enable mp bgp",
														Computed:            true,
													},
													"enable_sender_side_loop_detection": dsschema.BoolAttribute{
														MarkdownDescription: "Enable sender side loop detection",
														Computed:            true,
													},
													"inherit": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Inherit",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"no": dsschema.SingleNestedAttribute{
																MarkdownDescription: "No",
																Computed:            true,
																Attributes: map[string]dsschema.Attribute{
																	"address_family": dsschema.SingleNestedAttribute{
																		MarkdownDescription: "Address family",
																		Computed:            true,
																		Attributes: map[string]dsschema.Attribute{
																			"ipv4": dsschema.StringAttribute{
																				MarkdownDescription: "Ipv4",
																				Computed:            true,
																			},
																			"ipv6": dsschema.StringAttribute{
																				MarkdownDescription: "Ipv6",
																				Computed:            true,
																			},
																		},
																	},
																	"filtering_profile": dsschema.SingleNestedAttribute{
																		MarkdownDescription: "Filtering profile",
																		Computed:            true,
																		Attributes: map[string]dsschema.Attribute{
																			"ipv4": dsschema.StringAttribute{
																				MarkdownDescription: "Ipv4",
																				Computed:            true,
																			},
																			"ipv6": dsschema.StringAttribute{
																				MarkdownDescription: "Ipv6",
																				Computed:            true,
																			},
																		},
																	},
																},
															},
															"yes": dsschema.SingleNestedAttribute{
																MarkdownDescription: "Yes",
																Computed:            true,
																Attributes:          map[string]dsschema.Attribute{},
															},
														},
													},
													"local_address": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Local address",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"interface": dsschema.StringAttribute{
																MarkdownDescription: "Interface",
																Computed:            true,
															},
															"ip": dsschema.StringAttribute{
																MarkdownDescription: "Ip",
																Computed:            true,
															},
														},
													},
													"name": dsschema.StringAttribute{
														MarkdownDescription: "Name",
														Computed:            true,
													},
													"passive": dsschema.BoolAttribute{
														MarkdownDescription: "Passive",
														Computed:            true,
													},
													"peer_address": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Peer address",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"fqdn": dsschema.StringAttribute{
																MarkdownDescription: "Fqdn",
																Computed:            true,
															},
															"ip": dsschema.StringAttribute{
																MarkdownDescription: "Ip",
																Computed:            true,
															},
														},
													},
													"peer_as": dsschema.StringAttribute{
														MarkdownDescription: "Peer as",
														Computed:            true,
													},
													"peering_type": dsschema.StringAttribute{
														MarkdownDescription: "Peering type",
														Computed:            true,
													},
													"reflector_client": dsschema.StringAttribute{
														MarkdownDescription: "Reflector client",
														Computed:            true,
													},
													"subsequent_address_family_identifier": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Subsequent address family identifier",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"multicast": dsschema.BoolAttribute{
																MarkdownDescription: "Multicast",
																Computed:            true,
															},
															"unicast": dsschema.BoolAttribute{
																MarkdownDescription: "Unicast",
																Computed:            true,
															},
														},
													},
												},
											},
										},
										"soft_reset_with_stored_info": dsschema.BoolAttribute{
											MarkdownDescription: "Soft reset with stored info",
											Computed:            true,
										},
										"type": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Type",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"ebgp": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Ebgp",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"export_nexthop": dsschema.StringAttribute{
															MarkdownDescription: "Export nexthop",
															Computed:            true,
														},
														"import_nexthop": dsschema.StringAttribute{
															MarkdownDescription: "Import nexthop",
															Computed:            true,
														},
														"remove_private_as": dsschema.BoolAttribute{
															MarkdownDescription: "Remove private as",
															Computed:            true,
														},
													},
												},
												"ebgp_confed": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Ebgp confed",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"export_nexthop": dsschema.StringAttribute{
															MarkdownDescription: "Export nexthop",
															Computed:            true,
														},
													},
												},
												"ibgp": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Ibgp",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"export_nexthop": dsschema.StringAttribute{
															MarkdownDescription: "Export nexthop",
															Computed:            true,
														},
													},
												},
												"ibgp_confed": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Ibgp confed",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"export_nexthop": dsschema.StringAttribute{
															MarkdownDescription: "Export nexthop",
															Computed:            true,
														},
													},
												},
											},
										},
									},
								},
							},
							"policy": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Policy",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"aggregation": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Aggregation",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"address": dsschema.ListNestedAttribute{
												MarkdownDescription: "Address",
												Computed:            true,
												NestedObject: dsschema.NestedAttributeObject{
													Attributes: map[string]dsschema.Attribute{
														"advertise_filters": dsschema.ListNestedAttribute{
															MarkdownDescription: "Advertise filters",
															Computed:            true,
															NestedObject: dsschema.NestedAttributeObject{
																Attributes: map[string]dsschema.Attribute{
																	"enable": dsschema.BoolAttribute{
																		MarkdownDescription: "Enable",
																		Computed:            true,
																	},
																	"match": dsschema.SingleNestedAttribute{
																		MarkdownDescription: "Match",
																		Computed:            true,
																		Attributes: map[string]dsschema.Attribute{
																			"address_prefix": dsschema.ListNestedAttribute{
																				MarkdownDescription: "Address prefix",
																				Computed:            true,
																				NestedObject: dsschema.NestedAttributeObject{
																					Attributes: map[string]dsschema.Attribute{
																						"exact": dsschema.BoolAttribute{
																							MarkdownDescription: "Exact",
																							Computed:            true,
																						},
																						"name": dsschema.StringAttribute{
																							MarkdownDescription: "Name",
																							Computed:            true,
																						},
																					},
																				},
																			},
																			"afi": dsschema.StringAttribute{
																				MarkdownDescription: "Afi",
																				Computed:            true,
																			},
																			"as_path": dsschema.SingleNestedAttribute{
																				MarkdownDescription: "As path",
																				Computed:            true,
																				Attributes: map[string]dsschema.Attribute{
																					"regex": dsschema.StringAttribute{
																						MarkdownDescription: "Regex",
																						Computed:            true,
																					},
																				},
																			},
																			"community": dsschema.SingleNestedAttribute{
																				MarkdownDescription: "Community",
																				Computed:            true,
																				Attributes: map[string]dsschema.Attribute{
																					"regex": dsschema.StringAttribute{
																						MarkdownDescription: "Regex",
																						Computed:            true,
																					},
																				},
																			},
																			"extended_community": dsschema.SingleNestedAttribute{
																				MarkdownDescription: "Extended community",
																				Computed:            true,
																				Attributes: map[string]dsschema.Attribute{
																					"regex": dsschema.StringAttribute{
																						MarkdownDescription: "Regex",
																						Computed:            true,
																					},
																				},
																			},
																			"from_peer": dsschema.ListAttribute{
																				ElementType:         types.StringType,
																				MarkdownDescription: "From peer",
																				Computed:            true,
																			},
																			"med": dsschema.Float64Attribute{
																				MarkdownDescription: "Med",
																				Computed:            true,
																			},
																			"nexthop": dsschema.ListAttribute{
																				ElementType:         types.StringType,
																				MarkdownDescription: "Nexthop",
																				Computed:            true,
																			},
																			"route_table": dsschema.StringAttribute{
																				MarkdownDescription: "Route table",
																				Computed:            true,
																			},
																			"safi": dsschema.StringAttribute{
																				MarkdownDescription: "Safi",
																				Computed:            true,
																			},
																		},
																	},
																	"name": dsschema.StringAttribute{
																		MarkdownDescription: "Name",
																		Computed:            true,
																	},
																},
															},
														},
														"aggregate_route_attributes": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Aggregate route attributes",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"as_path": dsschema.SingleNestedAttribute{
																	MarkdownDescription: "As path",
																	Computed:            true,
																	Attributes: map[string]dsschema.Attribute{
																		"none": dsschema.SingleNestedAttribute{
																			MarkdownDescription: "None",
																			Computed:            true,
																			Attributes:          map[string]dsschema.Attribute{},
																		},
																		"prepend": dsschema.Float64Attribute{
																			MarkdownDescription: "Prepend",
																			Computed:            true,
																		},
																		"remove": dsschema.SingleNestedAttribute{
																			MarkdownDescription: "Remove",
																			Computed:            true,
																			Attributes:          map[string]dsschema.Attribute{},
																		},
																		"remove_and_prepend": dsschema.Float64Attribute{
																			MarkdownDescription: "Remove and prepend",
																			Computed:            true,
																		},
																	},
																},
																"as_path_limit": dsschema.Float64Attribute{
																	MarkdownDescription: "As path limit",
																	Computed:            true,
																},
																"community": dsschema.SingleNestedAttribute{
																	MarkdownDescription: "Community",
																	Computed:            true,
																	Attributes: map[string]dsschema.Attribute{
																		"append": dsschema.ListAttribute{
																			ElementType:         types.StringType,
																			MarkdownDescription: "Append",
																			Computed:            true,
																		},
																		"none": dsschema.SingleNestedAttribute{
																			MarkdownDescription: "None",
																			Computed:            true,
																			Attributes:          map[string]dsschema.Attribute{},
																		},
																		"overwrite": dsschema.ListAttribute{
																			ElementType:         types.StringType,
																			MarkdownDescription: "Overwrite",
																			Computed:            true,
																		},
																		"remove_all": dsschema.SingleNestedAttribute{
																			MarkdownDescription: "Remove all",
																			Computed:            true,
																			Attributes:          map[string]dsschema.Attribute{},
																		},
																		"remove_regex": dsschema.StringAttribute{
																			MarkdownDescription: "Remove regex",
																			Computed:            true,
																		},
																	},
																},
																"extended_community": dsschema.SingleNestedAttribute{
																	MarkdownDescription: "Extended community",
																	Computed:            true,
																	Attributes: map[string]dsschema.Attribute{
																		"append": dsschema.ListAttribute{
																			ElementType:         types.StringType,
																			MarkdownDescription: "Append",
																			Computed:            true,
																		},
																		"none": dsschema.SingleNestedAttribute{
																			MarkdownDescription: "None",
																			Computed:            true,
																			Attributes:          map[string]dsschema.Attribute{},
																		},
																		"overwrite": dsschema.ListAttribute{
																			ElementType:         types.StringType,
																			MarkdownDescription: "Overwrite",
																			Computed:            true,
																		},
																		"remove_all": dsschema.SingleNestedAttribute{
																			MarkdownDescription: "Remove all",
																			Computed:            true,
																			Attributes:          map[string]dsschema.Attribute{},
																		},
																		"remove_regex": dsschema.StringAttribute{
																			MarkdownDescription: "Remove regex",
																			Computed:            true,
																		},
																	},
																},
																"local_preference": dsschema.Float64Attribute{
																	MarkdownDescription: "Local preference",
																	Computed:            true,
																},
																"med": dsschema.Float64Attribute{
																	MarkdownDescription: "Med",
																	Computed:            true,
																},
																"nexthop": dsschema.StringAttribute{
																	MarkdownDescription: "Nexthop",
																	Computed:            true,
																},
																"origin": dsschema.StringAttribute{
																	MarkdownDescription: "Origin",
																	Computed:            true,
																},
																"weight": dsschema.Float64Attribute{
																	MarkdownDescription: "Weight",
																	Computed:            true,
																},
															},
														},
														"as_set": dsschema.BoolAttribute{
															MarkdownDescription: "As set",
															Computed:            true,
														},
														"enable": dsschema.BoolAttribute{
															MarkdownDescription: "Enable",
															Computed:            true,
														},
														"name": dsschema.StringAttribute{
															MarkdownDescription: "Name",
															Computed:            true,
														},
														"prefix": dsschema.StringAttribute{
															MarkdownDescription: "Prefix",
															Computed:            true,
														},
														"summary": dsschema.BoolAttribute{
															MarkdownDescription: "Summary",
															Computed:            true,
														},
														"suppress_filters": dsschema.ListNestedAttribute{
															MarkdownDescription: "Suppress filters",
															Computed:            true,
															NestedObject: dsschema.NestedAttributeObject{
																Attributes: map[string]dsschema.Attribute{
																	"enable": dsschema.BoolAttribute{
																		MarkdownDescription: "Enable",
																		Computed:            true,
																	},
																	"match": dsschema.SingleNestedAttribute{
																		MarkdownDescription: "Match",
																		Computed:            true,
																		Attributes: map[string]dsschema.Attribute{
																			"address_prefix": dsschema.ListNestedAttribute{
																				MarkdownDescription: "Address prefix",
																				Computed:            true,
																				NestedObject: dsschema.NestedAttributeObject{
																					Attributes: map[string]dsschema.Attribute{
																						"exact": dsschema.BoolAttribute{
																							MarkdownDescription: "Exact",
																							Computed:            true,
																						},
																						"name": dsschema.StringAttribute{
																							MarkdownDescription: "Name",
																							Computed:            true,
																						},
																					},
																				},
																			},
																			"afi": dsschema.StringAttribute{
																				MarkdownDescription: "Afi",
																				Computed:            true,
																			},
																			"as_path": dsschema.SingleNestedAttribute{
																				MarkdownDescription: "As path",
																				Computed:            true,
																				Attributes: map[string]dsschema.Attribute{
																					"regex": dsschema.StringAttribute{
																						MarkdownDescription: "Regex",
																						Computed:            true,
																					},
																				},
																			},
																			"community": dsschema.SingleNestedAttribute{
																				MarkdownDescription: "Community",
																				Computed:            true,
																				Attributes: map[string]dsschema.Attribute{
																					"regex": dsschema.StringAttribute{
																						MarkdownDescription: "Regex",
																						Computed:            true,
																					},
																				},
																			},
																			"extended_community": dsschema.SingleNestedAttribute{
																				MarkdownDescription: "Extended community",
																				Computed:            true,
																				Attributes: map[string]dsschema.Attribute{
																					"regex": dsschema.StringAttribute{
																						MarkdownDescription: "Regex",
																						Computed:            true,
																					},
																				},
																			},
																			"from_peer": dsschema.ListAttribute{
																				ElementType:         types.StringType,
																				MarkdownDescription: "From peer",
																				Computed:            true,
																			},
																			"med": dsschema.Float64Attribute{
																				MarkdownDescription: "Med",
																				Computed:            true,
																			},
																			"nexthop": dsschema.ListAttribute{
																				ElementType:         types.StringType,
																				MarkdownDescription: "Nexthop",
																				Computed:            true,
																			},
																			"route_table": dsschema.StringAttribute{
																				MarkdownDescription: "Route table",
																				Computed:            true,
																			},
																			"safi": dsschema.StringAttribute{
																				MarkdownDescription: "Safi",
																				Computed:            true,
																			},
																		},
																	},
																	"name": dsschema.StringAttribute{
																		MarkdownDescription: "Name",
																		Computed:            true,
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									"conditional_advertisement": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Conditional advertisement",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"policy": dsschema.ListNestedAttribute{
												MarkdownDescription: "Policy",
												Computed:            true,
												NestedObject: dsschema.NestedAttributeObject{
													Attributes: map[string]dsschema.Attribute{
														"advertise_filters": dsschema.ListNestedAttribute{
															MarkdownDescription: "Advertise filters",
															Computed:            true,
															NestedObject: dsschema.NestedAttributeObject{
																Attributes: map[string]dsschema.Attribute{
																	"enable": dsschema.BoolAttribute{
																		MarkdownDescription: "Enable",
																		Computed:            true,
																	},
																	"match": dsschema.SingleNestedAttribute{
																		MarkdownDescription: "Match",
																		Computed:            true,
																		Attributes: map[string]dsschema.Attribute{
																			"address_prefix": dsschema.ListNestedAttribute{
																				MarkdownDescription: "Address prefix",
																				Computed:            true,
																				NestedObject: dsschema.NestedAttributeObject{
																					Attributes: map[string]dsschema.Attribute{
																						"exact": dsschema.BoolAttribute{
																							MarkdownDescription: "Exact",
																							Computed:            true,
																						},
																						"name": dsschema.StringAttribute{
																							MarkdownDescription: "Name",
																							Computed:            true,
																						},
																					},
																				},
																			},
																			"afi": dsschema.StringAttribute{
																				MarkdownDescription: "Afi",
																				Computed:            true,
																			},
																			"as_path": dsschema.SingleNestedAttribute{
																				MarkdownDescription: "As path",
																				Computed:            true,
																				Attributes: map[string]dsschema.Attribute{
																					"regex": dsschema.StringAttribute{
																						MarkdownDescription: "Regex",
																						Computed:            true,
																					},
																				},
																			},
																			"community": dsschema.SingleNestedAttribute{
																				MarkdownDescription: "Community",
																				Computed:            true,
																				Attributes: map[string]dsschema.Attribute{
																					"regex": dsschema.StringAttribute{
																						MarkdownDescription: "Regex",
																						Computed:            true,
																					},
																				},
																			},
																			"extended_community": dsschema.SingleNestedAttribute{
																				MarkdownDescription: "Extended community",
																				Computed:            true,
																				Attributes: map[string]dsschema.Attribute{
																					"regex": dsschema.StringAttribute{
																						MarkdownDescription: "Regex",
																						Computed:            true,
																					},
																				},
																			},
																			"from_peer": dsschema.ListAttribute{
																				ElementType:         types.StringType,
																				MarkdownDescription: "From peer",
																				Computed:            true,
																			},
																			"med": dsschema.Float64Attribute{
																				MarkdownDescription: "Med",
																				Computed:            true,
																			},
																			"nexthop": dsschema.ListAttribute{
																				ElementType:         types.StringType,
																				MarkdownDescription: "Nexthop",
																				Computed:            true,
																			},
																			"route_table": dsschema.StringAttribute{
																				MarkdownDescription: "Route table",
																				Computed:            true,
																			},
																			"safi": dsschema.StringAttribute{
																				MarkdownDescription: "Safi",
																				Computed:            true,
																			},
																		},
																	},
																	"name": dsschema.StringAttribute{
																		MarkdownDescription: "Name",
																		Computed:            true,
																	},
																},
															},
														},
														"enable": dsschema.BoolAttribute{
															MarkdownDescription: "Enable",
															Computed:            true,
														},
														"name": dsschema.StringAttribute{
															MarkdownDescription: "Name",
															Computed:            true,
														},
														"non_exist_filters": dsschema.ListNestedAttribute{
															MarkdownDescription: "Non exist filters",
															Computed:            true,
															NestedObject: dsschema.NestedAttributeObject{
																Attributes: map[string]dsschema.Attribute{
																	"enable": dsschema.BoolAttribute{
																		MarkdownDescription: "Enable",
																		Computed:            true,
																	},
																	"match": dsschema.SingleNestedAttribute{
																		MarkdownDescription: "Match",
																		Computed:            true,
																		Attributes: map[string]dsschema.Attribute{
																			"address_prefix": dsschema.ListNestedAttribute{
																				MarkdownDescription: "Address prefix",
																				Computed:            true,
																				NestedObject: dsschema.NestedAttributeObject{
																					Attributes: map[string]dsschema.Attribute{
																						"exact": dsschema.BoolAttribute{
																							MarkdownDescription: "Exact",
																							Computed:            true,
																						},
																						"name": dsschema.StringAttribute{
																							MarkdownDescription: "Name",
																							Computed:            true,
																						},
																					},
																				},
																			},
																			"afi": dsschema.StringAttribute{
																				MarkdownDescription: "Afi",
																				Computed:            true,
																			},
																			"as_path": dsschema.SingleNestedAttribute{
																				MarkdownDescription: "As path",
																				Computed:            true,
																				Attributes: map[string]dsschema.Attribute{
																					"regex": dsschema.StringAttribute{
																						MarkdownDescription: "Regex",
																						Computed:            true,
																					},
																				},
																			},
																			"community": dsschema.SingleNestedAttribute{
																				MarkdownDescription: "Community",
																				Computed:            true,
																				Attributes: map[string]dsschema.Attribute{
																					"regex": dsschema.StringAttribute{
																						MarkdownDescription: "Regex",
																						Computed:            true,
																					},
																				},
																			},
																			"extended_community": dsschema.SingleNestedAttribute{
																				MarkdownDescription: "Extended community",
																				Computed:            true,
																				Attributes: map[string]dsschema.Attribute{
																					"regex": dsschema.StringAttribute{
																						MarkdownDescription: "Regex",
																						Computed:            true,
																					},
																				},
																			},
																			"from_peer": dsschema.ListAttribute{
																				ElementType:         types.StringType,
																				MarkdownDescription: "From peer",
																				Computed:            true,
																			},
																			"med": dsschema.Float64Attribute{
																				MarkdownDescription: "Med",
																				Computed:            true,
																			},
																			"nexthop": dsschema.ListAttribute{
																				ElementType:         types.StringType,
																				MarkdownDescription: "Nexthop",
																				Computed:            true,
																			},
																			"route_table": dsschema.StringAttribute{
																				MarkdownDescription: "Route table",
																				Computed:            true,
																			},
																			"safi": dsschema.StringAttribute{
																				MarkdownDescription: "Safi",
																				Computed:            true,
																			},
																		},
																	},
																	"name": dsschema.StringAttribute{
																		MarkdownDescription: "Name",
																		Computed:            true,
																	},
																},
															},
														},
														"used_by": dsschema.ListAttribute{
															ElementType:         types.StringType,
															MarkdownDescription: "Used by",
															Computed:            true,
														},
													},
												},
											},
										},
									},
									"export": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Export",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"rules": dsschema.ListNestedAttribute{
												MarkdownDescription: "Rules",
												Computed:            true,
												NestedObject: dsschema.NestedAttributeObject{
													Attributes: map[string]dsschema.Attribute{
														"action": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Action",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"allow": dsschema.SingleNestedAttribute{
																	MarkdownDescription: "Allow",
																	Computed:            true,
																	Attributes: map[string]dsschema.Attribute{
																		"update": dsschema.SingleNestedAttribute{
																			MarkdownDescription: "Update",
																			Computed:            true,
																			Attributes: map[string]dsschema.Attribute{
																				"as_path": dsschema.SingleNestedAttribute{
																					MarkdownDescription: "As path",
																					Computed:            true,
																					Attributes: map[string]dsschema.Attribute{
																						"none": dsschema.SingleNestedAttribute{
																							MarkdownDescription: "None",
																							Computed:            true,
																							Attributes:          map[string]dsschema.Attribute{},
																						},
																						"prepend": dsschema.Float64Attribute{
																							MarkdownDescription: "Prepend",
																							Computed:            true,
																						},
																						"remove": dsschema.SingleNestedAttribute{
																							MarkdownDescription: "Remove",
																							Computed:            true,
																							Attributes:          map[string]dsschema.Attribute{},
																						},
																						"remove_and_prepend": dsschema.Float64Attribute{
																							MarkdownDescription: "Remove and prepend",
																							Computed:            true,
																						},
																					},
																				},
																				"as_path_limit": dsschema.Float64Attribute{
																					MarkdownDescription: "As path limit",
																					Computed:            true,
																				},
																				"community": dsschema.SingleNestedAttribute{
																					MarkdownDescription: "Community",
																					Computed:            true,
																					Attributes: map[string]dsschema.Attribute{
																						"append": dsschema.ListAttribute{
																							ElementType:         types.StringType,
																							MarkdownDescription: "Append",
																							Computed:            true,
																						},
																						"none": dsschema.SingleNestedAttribute{
																							MarkdownDescription: "None",
																							Computed:            true,
																							Attributes:          map[string]dsschema.Attribute{},
																						},
																						"overwrite": dsschema.ListAttribute{
																							ElementType:         types.StringType,
																							MarkdownDescription: "Overwrite",
																							Computed:            true,
																						},
																						"remove_all": dsschema.SingleNestedAttribute{
																							MarkdownDescription: "Remove all",
																							Computed:            true,
																							Attributes:          map[string]dsschema.Attribute{},
																						},
																						"remove_regex": dsschema.StringAttribute{
																							MarkdownDescription: "Remove regex",
																							Computed:            true,
																						},
																					},
																				},
																				"extended_community": dsschema.SingleNestedAttribute{
																					MarkdownDescription: "Extended community",
																					Computed:            true,
																					Attributes: map[string]dsschema.Attribute{
																						"append": dsschema.ListAttribute{
																							ElementType:         types.StringType,
																							MarkdownDescription: "Append",
																							Computed:            true,
																						},
																						"none": dsschema.SingleNestedAttribute{
																							MarkdownDescription: "None",
																							Computed:            true,
																							Attributes:          map[string]dsschema.Attribute{},
																						},
																						"overwrite": dsschema.ListAttribute{
																							ElementType:         types.StringType,
																							MarkdownDescription: "Overwrite",
																							Computed:            true,
																						},
																						"remove_all": dsschema.SingleNestedAttribute{
																							MarkdownDescription: "Remove all",
																							Computed:            true,
																							Attributes:          map[string]dsschema.Attribute{},
																						},
																						"remove_regex": dsschema.StringAttribute{
																							MarkdownDescription: "Remove regex",
																							Computed:            true,
																						},
																					},
																				},
																				"local_preference": dsschema.Float64Attribute{
																					MarkdownDescription: "Local preference",
																					Computed:            true,
																				},
																				"med": dsschema.Float64Attribute{
																					MarkdownDescription: "Med",
																					Computed:            true,
																				},
																				"nexthop": dsschema.StringAttribute{
																					MarkdownDescription: "Nexthop",
																					Computed:            true,
																				},
																				"origin": dsschema.StringAttribute{
																					MarkdownDescription: "Origin",
																					Computed:            true,
																				},
																			},
																		},
																	},
																},
																"deny": dsschema.SingleNestedAttribute{
																	MarkdownDescription: "Deny",
																	Computed:            true,
																	Attributes:          map[string]dsschema.Attribute{},
																},
															},
														},
														"enable": dsschema.BoolAttribute{
															MarkdownDescription: "Enable",
															Computed:            true,
														},
														"match": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Match",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"address_prefix": dsschema.ListNestedAttribute{
																	MarkdownDescription: "Address prefix",
																	Computed:            true,
																	NestedObject: dsschema.NestedAttributeObject{
																		Attributes: map[string]dsschema.Attribute{
																			"exact": dsschema.BoolAttribute{
																				MarkdownDescription: "Exact",
																				Computed:            true,
																			},
																			"name": dsschema.StringAttribute{
																				MarkdownDescription: "Name",
																				Computed:            true,
																			},
																		},
																	},
																},
																"afi": dsschema.StringAttribute{
																	MarkdownDescription: "Afi",
																	Computed:            true,
																},
																"as_path": dsschema.SingleNestedAttribute{
																	MarkdownDescription: "As path",
																	Computed:            true,
																	Attributes: map[string]dsschema.Attribute{
																		"regex": dsschema.StringAttribute{
																			MarkdownDescription: "Regex",
																			Computed:            true,
																		},
																	},
																},
																"community": dsschema.SingleNestedAttribute{
																	MarkdownDescription: "Community",
																	Computed:            true,
																	Attributes: map[string]dsschema.Attribute{
																		"regex": dsschema.StringAttribute{
																			MarkdownDescription: "Regex",
																			Computed:            true,
																		},
																	},
																},
																"extended_community": dsschema.SingleNestedAttribute{
																	MarkdownDescription: "Extended community",
																	Computed:            true,
																	Attributes: map[string]dsschema.Attribute{
																		"regex": dsschema.StringAttribute{
																			MarkdownDescription: "Regex",
																			Computed:            true,
																		},
																	},
																},
																"from_peer": dsschema.ListAttribute{
																	ElementType:         types.StringType,
																	MarkdownDescription: "From peer",
																	Computed:            true,
																},
																"med": dsschema.Float64Attribute{
																	MarkdownDescription: "Med",
																	Computed:            true,
																},
																"nexthop": dsschema.ListAttribute{
																	ElementType:         types.StringType,
																	MarkdownDescription: "Nexthop",
																	Computed:            true,
																},
																"route_table": dsschema.StringAttribute{
																	MarkdownDescription: "Route table",
																	Computed:            true,
																},
																"safi": dsschema.StringAttribute{
																	MarkdownDescription: "Safi",
																	Computed:            true,
																},
															},
														},
														"name": dsschema.StringAttribute{
															MarkdownDescription: "Name",
															Computed:            true,
														},
														"used_by": dsschema.ListAttribute{
															ElementType:         types.StringType,
															MarkdownDescription: "Used by",
															Computed:            true,
														},
													},
												},
											},
										},
									},
									"import": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Import",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"rules": dsschema.ListNestedAttribute{
												MarkdownDescription: "Rules",
												Computed:            true,
												NestedObject: dsschema.NestedAttributeObject{
													Attributes: map[string]dsschema.Attribute{
														"action": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Action",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"allow": dsschema.SingleNestedAttribute{
																	MarkdownDescription: "Allow",
																	Computed:            true,
																	Attributes: map[string]dsschema.Attribute{
																		"dampening": dsschema.StringAttribute{
																			MarkdownDescription: "Dampening",
																			Computed:            true,
																		},
																		"update": dsschema.SingleNestedAttribute{
																			MarkdownDescription: "Update",
																			Computed:            true,
																			Attributes: map[string]dsschema.Attribute{
																				"as_path": dsschema.SingleNestedAttribute{
																					MarkdownDescription: "As path",
																					Computed:            true,
																					Attributes: map[string]dsschema.Attribute{
																						"none": dsschema.SingleNestedAttribute{
																							MarkdownDescription: "None",
																							Computed:            true,
																							Attributes:          map[string]dsschema.Attribute{},
																						},
																						"prepend": dsschema.Float64Attribute{
																							MarkdownDescription: "Prepend",
																							Computed:            true,
																						},
																						"remove": dsschema.SingleNestedAttribute{
																							MarkdownDescription: "Remove",
																							Computed:            true,
																							Attributes:          map[string]dsschema.Attribute{},
																						},
																						"remove_and_prepend": dsschema.Float64Attribute{
																							MarkdownDescription: "Remove and prepend",
																							Computed:            true,
																						},
																					},
																				},
																				"as_path_limit": dsschema.Float64Attribute{
																					MarkdownDescription: "As path limit",
																					Computed:            true,
																				},
																				"community": dsschema.SingleNestedAttribute{
																					MarkdownDescription: "Community",
																					Computed:            true,
																					Attributes: map[string]dsschema.Attribute{
																						"append": dsschema.ListAttribute{
																							ElementType:         types.StringType,
																							MarkdownDescription: "Append",
																							Computed:            true,
																						},
																						"none": dsschema.SingleNestedAttribute{
																							MarkdownDescription: "None",
																							Computed:            true,
																							Attributes:          map[string]dsschema.Attribute{},
																						},
																						"overwrite": dsschema.ListAttribute{
																							ElementType:         types.StringType,
																							MarkdownDescription: "Overwrite",
																							Computed:            true,
																						},
																						"remove_all": dsschema.SingleNestedAttribute{
																							MarkdownDescription: "Remove all",
																							Computed:            true,
																							Attributes:          map[string]dsschema.Attribute{},
																						},
																						"remove_regex": dsschema.StringAttribute{
																							MarkdownDescription: "Remove regex",
																							Computed:            true,
																						},
																					},
																				},
																				"extended_community": dsschema.SingleNestedAttribute{
																					MarkdownDescription: "Extended community",
																					Computed:            true,
																					Attributes: map[string]dsschema.Attribute{
																						"append": dsschema.ListAttribute{
																							ElementType:         types.StringType,
																							MarkdownDescription: "Append",
																							Computed:            true,
																						},
																						"none": dsschema.SingleNestedAttribute{
																							MarkdownDescription: "None",
																							Computed:            true,
																							Attributes:          map[string]dsschema.Attribute{},
																						},
																						"overwrite": dsschema.ListAttribute{
																							ElementType:         types.StringType,
																							MarkdownDescription: "Overwrite",
																							Computed:            true,
																						},
																						"remove_all": dsschema.SingleNestedAttribute{
																							MarkdownDescription: "Remove all",
																							Computed:            true,
																							Attributes:          map[string]dsschema.Attribute{},
																						},
																						"remove_regex": dsschema.StringAttribute{
																							MarkdownDescription: "Remove regex",
																							Computed:            true,
																						},
																					},
																				},
																				"local_preference": dsschema.Float64Attribute{
																					MarkdownDescription: "Local preference",
																					Computed:            true,
																				},
																				"med": dsschema.Float64Attribute{
																					MarkdownDescription: "Med",
																					Computed:            true,
																				},
																				"nexthop": dsschema.StringAttribute{
																					MarkdownDescription: "Nexthop",
																					Computed:            true,
																				},
																				"origin": dsschema.StringAttribute{
																					MarkdownDescription: "Origin",
																					Computed:            true,
																				},
																				"weight": dsschema.Float64Attribute{
																					MarkdownDescription: "Weight",
																					Computed:            true,
																				},
																			},
																		},
																	},
																},
																"deny": dsschema.SingleNestedAttribute{
																	MarkdownDescription: "Deny",
																	Computed:            true,
																	Attributes:          map[string]dsschema.Attribute{},
																},
															},
														},
														"enable": dsschema.BoolAttribute{
															MarkdownDescription: "Enable",
															Computed:            true,
														},
														"match": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Match",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"address_prefix": dsschema.ListNestedAttribute{
																	MarkdownDescription: "Address prefix",
																	Computed:            true,
																	NestedObject: dsschema.NestedAttributeObject{
																		Attributes: map[string]dsschema.Attribute{
																			"exact": dsschema.BoolAttribute{
																				MarkdownDescription: "Exact",
																				Computed:            true,
																			},
																			"name": dsschema.StringAttribute{
																				MarkdownDescription: "Name",
																				Computed:            true,
																			},
																		},
																	},
																},
																"afi": dsschema.StringAttribute{
																	MarkdownDescription: "Afi",
																	Computed:            true,
																},
																"as_path": dsschema.SingleNestedAttribute{
																	MarkdownDescription: "As path",
																	Computed:            true,
																	Attributes: map[string]dsschema.Attribute{
																		"regex": dsschema.StringAttribute{
																			MarkdownDescription: "Regex",
																			Computed:            true,
																		},
																	},
																},
																"community": dsschema.SingleNestedAttribute{
																	MarkdownDescription: "Community",
																	Computed:            true,
																	Attributes: map[string]dsschema.Attribute{
																		"regex": dsschema.StringAttribute{
																			MarkdownDescription: "Regex",
																			Computed:            true,
																		},
																	},
																},
																"extended_community": dsschema.SingleNestedAttribute{
																	MarkdownDescription: "Extended community",
																	Computed:            true,
																	Attributes: map[string]dsschema.Attribute{
																		"regex": dsschema.StringAttribute{
																			MarkdownDescription: "Regex",
																			Computed:            true,
																		},
																	},
																},
																"from_peer": dsschema.ListAttribute{
																	ElementType:         types.StringType,
																	MarkdownDescription: "From peer",
																	Computed:            true,
																},
																"med": dsschema.Float64Attribute{
																	MarkdownDescription: "Med",
																	Computed:            true,
																},
																"nexthop": dsschema.ListAttribute{
																	ElementType:         types.StringType,
																	MarkdownDescription: "Nexthop",
																	Computed:            true,
																},
																"route_table": dsschema.StringAttribute{
																	MarkdownDescription: "Route table",
																	Computed:            true,
																},
																"safi": dsschema.StringAttribute{
																	MarkdownDescription: "Safi",
																	Computed:            true,
																},
															},
														},
														"name": dsschema.StringAttribute{
															MarkdownDescription: "Name",
															Computed:            true,
														},
														"used_by": dsschema.ListAttribute{
															ElementType:         types.StringType,
															MarkdownDescription: "Used by",
															Computed:            true,
														},
													},
												},
											},
										},
									},
								},
							},
							"redist_rules": dsschema.ListNestedAttribute{
								MarkdownDescription: "Redist rules",
								Computed:            true,
								NestedObject: dsschema.NestedAttributeObject{
									Attributes: map[string]dsschema.Attribute{
										"address_family_identifier": dsschema.StringAttribute{
											MarkdownDescription: "Address family identifier",
											Computed:            true,
										},
										"enable": dsschema.BoolAttribute{
											MarkdownDescription: "Enable",
											Computed:            true,
										},
										"metric": dsschema.Float64Attribute{
											MarkdownDescription: "Metric",
											Computed:            true,
										},
										"name": dsschema.StringAttribute{
											MarkdownDescription: "Name",
											Computed:            true,
										},
										"route_table": dsschema.StringAttribute{
											MarkdownDescription: "Route table",
											Computed:            true,
										},
										"set_as_path_limit": dsschema.Float64Attribute{
											MarkdownDescription: "Set as path limit",
											Computed:            true,
										},
										"set_community": dsschema.ListAttribute{
											ElementType:         types.StringType,
											MarkdownDescription: "Set community",
											Computed:            true,
										},
										"set_extended_community": dsschema.ListAttribute{
											ElementType:         types.StringType,
											MarkdownDescription: "Set extended community",
											Computed:            true,
										},
										"set_local_preference": dsschema.Float64Attribute{
											MarkdownDescription: "Set local preference",
											Computed:            true,
										},
										"set_med": dsschema.Float64Attribute{
											MarkdownDescription: "Set med",
											Computed:            true,
										},
										"set_origin": dsschema.StringAttribute{
											MarkdownDescription: "Set origin",
											Computed:            true,
										},
									},
								},
							},
							"redistribution_profile": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Redistribution profile",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"ipv4": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Ipv4",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"unicast": dsschema.StringAttribute{
												MarkdownDescription: "Unicast",
												Computed:            true,
											},
										},
									},
									"ipv6": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Ipv6",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"unicast": dsschema.StringAttribute{
												MarkdownDescription: "Unicast",
												Computed:            true,
											},
										},
									},
								},
							},
							"reject_default_route": dsschema.BoolAttribute{
								MarkdownDescription: "Reject default route",
								Computed:            true,
							},
							"router_id": dsschema.StringAttribute{
								MarkdownDescription: "Router id",
								Computed:            true,
							},
						},
					},
					"ecmp": dsschema.SingleNestedAttribute{
						MarkdownDescription: "Ecmp",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"algorithm": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Algorithm",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"balanced_round_robin": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Balanced round robin",
										Computed:            true,
										Attributes:          map[string]dsschema.Attribute{},
									},
									"ip_hash": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Ip hash",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"hash_seed": dsschema.Float64Attribute{
												MarkdownDescription: "Hash seed",
												Computed:            true,
											},
											"src_only": dsschema.BoolAttribute{
												MarkdownDescription: "Src only",
												Computed:            true,
											},
											"use_port": dsschema.BoolAttribute{
												MarkdownDescription: "Use port",
												Computed:            true,
											},
										},
									},
									"ip_modulo": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Ip modulo",
										Computed:            true,
										Attributes:          map[string]dsschema.Attribute{},
									},
									"weighted_round_robin": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Weighted round robin",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"interface": dsschema.ListNestedAttribute{
												MarkdownDescription: "Interface",
												Computed:            true,
												NestedObject: dsschema.NestedAttributeObject{
													Attributes: map[string]dsschema.Attribute{
														"name": dsschema.StringAttribute{
															MarkdownDescription: "Name",
															Computed:            true,
														},
														"weight": dsschema.Float64Attribute{
															MarkdownDescription: "Weight",
															Computed:            true,
														},
													},
												},
											},
										},
									},
								},
							},
							"enable": dsschema.BoolAttribute{
								MarkdownDescription: "Enable",
								Computed:            true,
							},
							"max_path": dsschema.Float64Attribute{
								MarkdownDescription: "Max path",
								Computed:            true,
							},
							"strict_source_path": dsschema.BoolAttribute{
								MarkdownDescription: "Strict source path",
								Computed:            true,
							},
							"symmetric_return": dsschema.BoolAttribute{
								MarkdownDescription: "Symmetric return",
								Computed:            true,
							},
						},
					},
					"global_vrid": dsschema.Float64Attribute{
						MarkdownDescription: "Global vrid",
						Computed:            true,
					},
					"interface": dsschema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "Interface",
						Computed:            true,
					},
					"multicast": dsschema.SingleNestedAttribute{
						MarkdownDescription: "Multicast",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"enable": dsschema.BoolAttribute{
								MarkdownDescription: "Enable",
								Computed:            true,
							},
							"enable_v6": dsschema.BoolAttribute{
								MarkdownDescription: "Enable v6",
								Computed:            true,
							},
							"igmp": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Igmp",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"dynamic": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Dynamic",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"interface": dsschema.ListNestedAttribute{
												MarkdownDescription: "Interface",
												Computed:            true,
												NestedObject: dsschema.NestedAttributeObject{
													Attributes: map[string]dsschema.Attribute{
														"group_filter": dsschema.StringAttribute{
															MarkdownDescription: "Group filter",
															Computed:            true,
														},
														"max_groups": dsschema.StringAttribute{
															MarkdownDescription: "Max groups",
															Computed:            true,
														},
														"max_sources": dsschema.StringAttribute{
															MarkdownDescription: "Max sources",
															Computed:            true,
														},
														"name": dsschema.StringAttribute{
															MarkdownDescription: "Name",
															Computed:            true,
														},
														"query_profile": dsschema.StringAttribute{
															MarkdownDescription: "Query profile",
															Computed:            true,
														},
														"robustness": dsschema.StringAttribute{
															MarkdownDescription: "Robustness",
															Computed:            true,
														},
														"router_alert_policing": dsschema.BoolAttribute{
															MarkdownDescription: "Router alert policing",
															Computed:            true,
														},
														"version": dsschema.StringAttribute{
															MarkdownDescription: "Version",
															Computed:            true,
														},
													},
												},
											},
										},
									},
									"enable": dsschema.BoolAttribute{
										MarkdownDescription: "Enable",
										Computed:            true,
									},
									"static": dsschema.ListNestedAttribute{
										MarkdownDescription: "Static",
										Computed:            true,
										NestedObject: dsschema.NestedAttributeObject{
											Attributes: map[string]dsschema.Attribute{
												"group_address": dsschema.StringAttribute{
													MarkdownDescription: "Group address",
													Computed:            true,
												},
												"interface": dsschema.StringAttribute{
													MarkdownDescription: "Interface",
													Computed:            true,
												},
												"name": dsschema.StringAttribute{
													MarkdownDescription: "Name",
													Computed:            true,
												},
												"source_address": dsschema.StringAttribute{
													MarkdownDescription: "Source address",
													Computed:            true,
												},
											},
										},
									},
								},
							},
							"interface_group": dsschema.ListNestedAttribute{
								MarkdownDescription: "Interface group",
								Computed:            true,
								NestedObject: dsschema.NestedAttributeObject{
									Attributes: map[string]dsschema.Attribute{
										"description": dsschema.StringAttribute{
											MarkdownDescription: "Description",
											Computed:            true,
										},
										"group_permission": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Group permission",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"any_source_multicast": dsschema.ListNestedAttribute{
													MarkdownDescription: "Any source multicast",
													Computed:            true,
													NestedObject: dsschema.NestedAttributeObject{
														Attributes: map[string]dsschema.Attribute{
															"group_address": dsschema.StringAttribute{
																MarkdownDescription: "Group address",
																Computed:            true,
															},
															"included": dsschema.BoolAttribute{
																MarkdownDescription: "Included",
																Computed:            true,
															},
															"name": dsschema.StringAttribute{
																MarkdownDescription: "Name",
																Computed:            true,
															},
														},
													},
												},
												"source_specific_multicast": dsschema.ListNestedAttribute{
													MarkdownDescription: "Source specific multicast",
													Computed:            true,
													NestedObject: dsschema.NestedAttributeObject{
														Attributes: map[string]dsschema.Attribute{
															"group_address": dsschema.StringAttribute{
																MarkdownDescription: "Group address",
																Computed:            true,
															},
															"included": dsschema.BoolAttribute{
																MarkdownDescription: "Included",
																Computed:            true,
															},
															"name": dsschema.StringAttribute{
																MarkdownDescription: "Name",
																Computed:            true,
															},
															"source_address": dsschema.StringAttribute{
																MarkdownDescription: "Source address",
																Computed:            true,
															},
														},
													},
												},
											},
										},
										"igmp": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Igmp",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"enable": dsschema.BoolAttribute{
													MarkdownDescription: "Enable",
													Computed:            true,
												},
												"immediate_leave": dsschema.BoolAttribute{
													MarkdownDescription: "Immediate leave",
													Computed:            true,
												},
												"last_member_query_interval": dsschema.Float64Attribute{
													MarkdownDescription: "Last member query interval",
													Computed:            true,
												},
												"max_groups": dsschema.StringAttribute{
													MarkdownDescription: "Max groups",
													Computed:            true,
												},
												"max_query_response_time": dsschema.Float64Attribute{
													MarkdownDescription: "Max query response time",
													Computed:            true,
												},
												"max_sources": dsschema.StringAttribute{
													MarkdownDescription: "Max sources",
													Computed:            true,
												},
												"mode": dsschema.StringAttribute{
													MarkdownDescription: "Mode",
													Computed:            true,
												},
												"query_interval": dsschema.Float64Attribute{
													MarkdownDescription: "Query interval",
													Computed:            true,
												},
												"robustness": dsschema.StringAttribute{
													MarkdownDescription: "Robustness",
													Computed:            true,
												},
												"router_alert_policing": dsschema.BoolAttribute{
													MarkdownDescription: "Router alert policing",
													Computed:            true,
												},
												"version": dsschema.StringAttribute{
													MarkdownDescription: "Version",
													Computed:            true,
												},
											},
										},
										"interface": dsschema.ListAttribute{
											ElementType:         types.StringType,
											MarkdownDescription: "Interface",
											Computed:            true,
										},
										"name": dsschema.StringAttribute{
											MarkdownDescription: "Name",
											Computed:            true,
										},
										"pim": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Pim",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"allowed_neighbors": dsschema.ListNestedAttribute{
													MarkdownDescription: "Allowed neighbors",
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
												"assert_interval": dsschema.Float64Attribute{
													MarkdownDescription: "Assert interval",
													Computed:            true,
												},
												"bsr_border": dsschema.BoolAttribute{
													MarkdownDescription: "Bsr border",
													Computed:            true,
												},
												"dr_priority": dsschema.Float64Attribute{
													MarkdownDescription: "Dr priority",
													Computed:            true,
												},
												"enable": dsschema.BoolAttribute{
													MarkdownDescription: "Enable",
													Computed:            true,
												},
												"hello_interval": dsschema.Float64Attribute{
													MarkdownDescription: "Hello interval",
													Computed:            true,
												},
												"join_prune_interval": dsschema.Float64Attribute{
													MarkdownDescription: "Join prune interval",
													Computed:            true,
												},
											},
										},
									},
								},
							},
							"mode": dsschema.StringAttribute{
								MarkdownDescription: "Mode",
								Computed:            true,
							},
							"msdp": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Msdp",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"enable": dsschema.BoolAttribute{
										MarkdownDescription: "Enable",
										Computed:            true,
									},
									"global_authentication": dsschema.StringAttribute{
										MarkdownDescription: "Global authentication",
										Computed:            true,
									},
									"global_timer": dsschema.StringAttribute{
										MarkdownDescription: "Global timer",
										Computed:            true,
									},
									"originator_id": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Originator id",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"interface": dsschema.StringAttribute{
												MarkdownDescription: "Interface",
												Computed:            true,
											},
											"ip": dsschema.StringAttribute{
												MarkdownDescription: "Ip",
												Computed:            true,
											},
										},
									},
									"peer": dsschema.ListNestedAttribute{
										MarkdownDescription: "Peer",
										Computed:            true,
										NestedObject: dsschema.NestedAttributeObject{
											Attributes: map[string]dsschema.Attribute{
												"authentication": dsschema.StringAttribute{
													MarkdownDescription: "Authentication",
													Computed:            true,
												},
												"enable": dsschema.BoolAttribute{
													MarkdownDescription: "Enable",
													Computed:            true,
												},
												"inbound_sa_filter": dsschema.StringAttribute{
													MarkdownDescription: "Inbound sa filter",
													Computed:            true,
												},
												"local_address": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Local address",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"interface": dsschema.StringAttribute{
															MarkdownDescription: "Interface",
															Computed:            true,
														},
														"ip": dsschema.StringAttribute{
															MarkdownDescription: "Ip",
															Computed:            true,
														},
													},
												},
												"max_sa": dsschema.Float64Attribute{
													MarkdownDescription: "Max sa",
													Computed:            true,
												},
												"name": dsschema.StringAttribute{
													MarkdownDescription: "Name",
													Computed:            true,
												},
												"outbound_sa_filter": dsschema.StringAttribute{
													MarkdownDescription: "Outbound sa filter",
													Computed:            true,
												},
												"peer_address": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Peer address",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"fqdn": dsschema.StringAttribute{
															MarkdownDescription: "Fqdn",
															Computed:            true,
														},
														"ip": dsschema.StringAttribute{
															MarkdownDescription: "Ip",
															Computed:            true,
														},
													},
												},
												"peer_as": dsschema.StringAttribute{
													MarkdownDescription: "Peer as",
													Computed:            true,
												},
											},
										},
									},
								},
							},
							"pim": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Pim",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"enable": dsschema.BoolAttribute{
										MarkdownDescription: "Enable",
										Computed:            true,
									},
									"group_permission": dsschema.StringAttribute{
										MarkdownDescription: "Group permission",
										Computed:            true,
									},
									"if_timer_global": dsschema.StringAttribute{
										MarkdownDescription: "If timer global",
										Computed:            true,
									},
									"interface": dsschema.ListNestedAttribute{
										MarkdownDescription: "Interface",
										Computed:            true,
										NestedObject: dsschema.NestedAttributeObject{
											Attributes: map[string]dsschema.Attribute{
												"description": dsschema.StringAttribute{
													MarkdownDescription: "Description",
													Computed:            true,
												},
												"dr_priority": dsschema.Float64Attribute{
													MarkdownDescription: "Dr priority",
													Computed:            true,
												},
												"if_timer": dsschema.StringAttribute{
													MarkdownDescription: "If timer",
													Computed:            true,
												},
												"name": dsschema.StringAttribute{
													MarkdownDescription: "Name",
													Computed:            true,
												},
												"neighbor_filter": dsschema.StringAttribute{
													MarkdownDescription: "Neighbor filter",
													Computed:            true,
												},
												"send_bsm": dsschema.BoolAttribute{
													MarkdownDescription: "Send bsm",
													Computed:            true,
												},
											},
										},
									},
									"route_ageout_time": dsschema.Float64Attribute{
										MarkdownDescription: "Route ageout time",
										Computed:            true,
									},
									"rp": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Rp",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"external_rp": dsschema.ListNestedAttribute{
												MarkdownDescription: "External rp",
												Computed:            true,
												NestedObject: dsschema.NestedAttributeObject{
													Attributes: map[string]dsschema.Attribute{
														"group_list": dsschema.StringAttribute{
															MarkdownDescription: "Group list",
															Computed:            true,
														},
														"name": dsschema.StringAttribute{
															MarkdownDescription: "Name",
															Computed:            true,
														},
														"override": dsschema.BoolAttribute{
															MarkdownDescription: "Override",
															Computed:            true,
														},
													},
												},
											},
											"local_rp": dsschema.SingleNestedAttribute{
												MarkdownDescription: "Local rp",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"candidate_rp": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Candidate rp",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"address": dsschema.StringAttribute{
																MarkdownDescription: "Address",
																Computed:            true,
															},
															"advertisement_interval": dsschema.Float64Attribute{
																MarkdownDescription: "Advertisement interval",
																Computed:            true,
															},
															"group_list": dsschema.StringAttribute{
																MarkdownDescription: "Group list",
																Computed:            true,
															},
															"interface": dsschema.StringAttribute{
																MarkdownDescription: "Interface",
																Computed:            true,
															},
															"priority": dsschema.Float64Attribute{
																MarkdownDescription: "Priority",
																Computed:            true,
															},
														},
													},
													"static_rp": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Static rp",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"address": dsschema.StringAttribute{
																MarkdownDescription: "Address",
																Computed:            true,
															},
															"group_list": dsschema.StringAttribute{
																MarkdownDescription: "Group list",
																Computed:            true,
															},
															"interface": dsschema.StringAttribute{
																MarkdownDescription: "Interface",
																Computed:            true,
															},
															"override": dsschema.BoolAttribute{
																MarkdownDescription: "Override",
																Computed:            true,
															},
														},
													},
												},
											},
										},
									},
									"rpf_lookup_mode": dsschema.StringAttribute{
										MarkdownDescription: "Rpf lookup mode",
										Computed:            true,
									},
									"spt_threshold": dsschema.ListNestedAttribute{
										MarkdownDescription: "Spt threshold",
										Computed:            true,
										NestedObject: dsschema.NestedAttributeObject{
											Attributes: map[string]dsschema.Attribute{
												"name": dsschema.StringAttribute{
													MarkdownDescription: "Name",
													Computed:            true,
												},
												"threshold": dsschema.StringAttribute{
													MarkdownDescription: "Threshold",
													Computed:            true,
												},
											},
										},
									},
									"ssm_address_space": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Ssm address space",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"group_list": dsschema.StringAttribute{
												MarkdownDescription: "Group list",
												Computed:            true,
											},
										},
									},
								},
							},
							"route_ageout_time": dsschema.Float64Attribute{
								MarkdownDescription: "Route ageout time",
								Computed:            true,
							},
							"rp": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Rp",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"external_rp": dsschema.ListNestedAttribute{
										MarkdownDescription: "External rp",
										Computed:            true,
										NestedObject: dsschema.NestedAttributeObject{
											Attributes: map[string]dsschema.Attribute{
												"group_addresses": dsschema.ListAttribute{
													ElementType:         types.StringType,
													MarkdownDescription: "Group addresses",
													Computed:            true,
												},
												"name": dsschema.StringAttribute{
													MarkdownDescription: "Name",
													Computed:            true,
												},
												"override": dsschema.BoolAttribute{
													MarkdownDescription: "Override",
													Computed:            true,
												},
											},
										},
									},
									"local_rp": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Local rp",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"candidate_rp": dsschema.SingleNestedAttribute{
												MarkdownDescription: "Candidate rp",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"address": dsschema.StringAttribute{
														MarkdownDescription: "Address",
														Computed:            true,
													},
													"advertisement_interval": dsschema.Float64Attribute{
														MarkdownDescription: "Advertisement interval",
														Computed:            true,
													},
													"group_addresses": dsschema.ListAttribute{
														ElementType:         types.StringType,
														MarkdownDescription: "Group addresses",
														Computed:            true,
													},
													"interface": dsschema.StringAttribute{
														MarkdownDescription: "Interface",
														Computed:            true,
													},
													"priority": dsschema.Float64Attribute{
														MarkdownDescription: "Priority",
														Computed:            true,
													},
												},
											},
											"static_rp": dsschema.SingleNestedAttribute{
												MarkdownDescription: "Static rp",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"address": dsschema.StringAttribute{
														MarkdownDescription: "Address",
														Computed:            true,
													},
													"group_addresses": dsschema.ListAttribute{
														ElementType:         types.StringType,
														MarkdownDescription: "Group addresses",
														Computed:            true,
													},
													"interface": dsschema.StringAttribute{
														MarkdownDescription: "Interface",
														Computed:            true,
													},
													"override": dsschema.BoolAttribute{
														MarkdownDescription: "Override",
														Computed:            true,
													},
												},
											},
										},
									},
								},
							},
							"spt_threshold": dsschema.ListNestedAttribute{
								MarkdownDescription: "Spt threshold",
								Computed:            true,
								NestedObject: dsschema.NestedAttributeObject{
									Attributes: map[string]dsschema.Attribute{
										"name": dsschema.StringAttribute{
											MarkdownDescription: "Name",
											Computed:            true,
										},
										"threshold": dsschema.StringAttribute{
											MarkdownDescription: "Threshold",
											Computed:            true,
										},
									},
								},
							},
							"ssm_address_space": dsschema.ListNestedAttribute{
								MarkdownDescription: "Ssm address space",
								Computed:            true,
								NestedObject: dsschema.NestedAttributeObject{
									Attributes: map[string]dsschema.Attribute{
										"group_address": dsschema.StringAttribute{
											MarkdownDescription: "Group address",
											Computed:            true,
										},
										"included": dsschema.BoolAttribute{
											MarkdownDescription: "Included",
											Computed:            true,
										},
										"name": dsschema.StringAttribute{
											MarkdownDescription: "Name",
											Computed:            true,
										},
									},
								},
							},
							"static_route": dsschema.ListNestedAttribute{
								MarkdownDescription: "Static route",
								Computed:            true,
								NestedObject: dsschema.NestedAttributeObject{
									Attributes: map[string]dsschema.Attribute{
										"destination": dsschema.StringAttribute{
											MarkdownDescription: "Destination",
											Computed:            true,
										},
										"interface": dsschema.StringAttribute{
											MarkdownDescription: "Interface",
											Computed:            true,
										},
										"name": dsschema.StringAttribute{
											MarkdownDescription: "Name",
											Computed:            true,
										},
										"nexthop": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Nexthop",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"ip_address": dsschema.StringAttribute{
													MarkdownDescription: "Ip address",
													Computed:            true,
												},
											},
										},
										"preference": dsschema.Float64Attribute{
											MarkdownDescription: "Preference",
											Computed:            true,
										},
									},
								},
							},
						},
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "Name",
						Computed:            true,
					},
					"ospf": dsschema.SingleNestedAttribute{
						MarkdownDescription: "Ospf",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"allow_redist_default_route": dsschema.BoolAttribute{
								MarkdownDescription: "Allow redist default route",
								Computed:            true,
							},
							"area": dsschema.ListNestedAttribute{
								MarkdownDescription: "Area",
								Computed:            true,
								NestedObject: dsschema.NestedAttributeObject{
									Attributes: map[string]dsschema.Attribute{
										"authentication": dsschema.StringAttribute{
											MarkdownDescription: "Authentication",
											Computed:            true,
										},
										"interface": dsschema.ListNestedAttribute{
											MarkdownDescription: "Interface",
											Computed:            true,
											NestedObject: dsschema.NestedAttributeObject{
												Attributes: map[string]dsschema.Attribute{
													"authentication": dsschema.StringAttribute{
														MarkdownDescription: "Authentication",
														Computed:            true,
													},
													"bfd": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Bfd",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"profile": dsschema.StringAttribute{
																MarkdownDescription: "Profile",
																Computed:            true,
															},
														},
													},
													"enable": dsschema.BoolAttribute{
														MarkdownDescription: "Enable",
														Computed:            true,
													},
													"link_type": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Link type",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"broadcast": dsschema.SingleNestedAttribute{
																MarkdownDescription: "Broadcast",
																Computed:            true,
																Attributes:          map[string]dsschema.Attribute{},
															},
															"p2mp": dsschema.SingleNestedAttribute{
																MarkdownDescription: "P2mp",
																Computed:            true,
																Attributes: map[string]dsschema.Attribute{
																	"neighbor": dsschema.ListNestedAttribute{
																		MarkdownDescription: "Neighbor",
																		Computed:            true,
																		NestedObject: dsschema.NestedAttributeObject{
																			Attributes: map[string]dsschema.Attribute{
																				"name": dsschema.StringAttribute{
																					MarkdownDescription: "Name",
																					Computed:            true,
																				},
																				"priority": dsschema.Float64Attribute{
																					MarkdownDescription: "Priority",
																					Computed:            true,
																				},
																			},
																		},
																	},
																},
															},
															"p2p": dsschema.SingleNestedAttribute{
																MarkdownDescription: "P2p",
																Computed:            true,
																Attributes:          map[string]dsschema.Attribute{},
															},
														},
													},
													"metric": dsschema.Float64Attribute{
														MarkdownDescription: "Metric",
														Computed:            true,
													},
													"mtu_ignore": dsschema.BoolAttribute{
														MarkdownDescription: "Mtu ignore",
														Computed:            true,
													},
													"name": dsschema.StringAttribute{
														MarkdownDescription: "Name",
														Computed:            true,
													},
													"passive": dsschema.BoolAttribute{
														MarkdownDescription: "Passive",
														Computed:            true,
													},
													"priority": dsschema.Float64Attribute{
														MarkdownDescription: "Priority",
														Computed:            true,
													},
													"timing": dsschema.StringAttribute{
														MarkdownDescription: "Timing",
														Computed:            true,
													},
													"vr_timing": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Vr timing",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"dead_counts": dsschema.Float64Attribute{
																MarkdownDescription: "Dead counts",
																Computed:            true,
															},
															"gr_delay": dsschema.Float64Attribute{
																MarkdownDescription: "Gr delay",
																Computed:            true,
															},
															"hello_interval": dsschema.Float64Attribute{
																MarkdownDescription: "Hello interval",
																Computed:            true,
															},
															"retransmit_interval": dsschema.Float64Attribute{
																MarkdownDescription: "Retransmit interval",
																Computed:            true,
															},
															"transit_delay": dsschema.Float64Attribute{
																MarkdownDescription: "Transit delay",
																Computed:            true,
															},
														},
													},
												},
											},
										},
										"name": dsschema.StringAttribute{
											MarkdownDescription: "Name",
											Computed:            true,
										},
										"range": dsschema.ListNestedAttribute{
											MarkdownDescription: "Range",
											Computed:            true,
											NestedObject: dsschema.NestedAttributeObject{
												Attributes: map[string]dsschema.Attribute{
													"advertise": dsschema.BoolAttribute{
														MarkdownDescription: "Advertise",
														Computed:            true,
													},
													"name": dsschema.StringAttribute{
														MarkdownDescription: "Name",
														Computed:            true,
													},
													"substitute": dsschema.StringAttribute{
														MarkdownDescription: "Substitute",
														Computed:            true,
													},
												},
											},
										},
										"type": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Type",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"normal": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Normal",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"abr": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Abr",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"export_list": dsschema.StringAttribute{
																	MarkdownDescription: "Export list",
																	Computed:            true,
																},
																"import_list": dsschema.StringAttribute{
																	MarkdownDescription: "Import list",
																	Computed:            true,
																},
																"inbound_filter_list": dsschema.StringAttribute{
																	MarkdownDescription: "Inbound filter list",
																	Computed:            true,
																},
																"outbound_filter_list": dsschema.StringAttribute{
																	MarkdownDescription: "Outbound filter list",
																	Computed:            true,
																},
															},
														},
													},
												},
												"nssa": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Nssa",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"abr": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Abr",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"export_list": dsschema.StringAttribute{
																	MarkdownDescription: "Export list",
																	Computed:            true,
																},
																"import_list": dsschema.StringAttribute{
																	MarkdownDescription: "Import list",
																	Computed:            true,
																},
																"inbound_filter_list": dsschema.StringAttribute{
																	MarkdownDescription: "Inbound filter list",
																	Computed:            true,
																},
																"nssa_ext_range": dsschema.ListNestedAttribute{
																	MarkdownDescription: "Nssa ext range",
																	Computed:            true,
																	NestedObject: dsschema.NestedAttributeObject{
																		Attributes: map[string]dsschema.Attribute{
																			"advertise": dsschema.BoolAttribute{
																				MarkdownDescription: "Advertise",
																				Computed:            true,
																			},
																			"name": dsschema.StringAttribute{
																				MarkdownDescription: "Name",
																				Computed:            true,
																			},
																			"route_tag": dsschema.Float64Attribute{
																				MarkdownDescription: "Route tag",
																				Computed:            true,
																			},
																		},
																	},
																},
																"outbound_filter_list": dsschema.StringAttribute{
																	MarkdownDescription: "Outbound filter list",
																	Computed:            true,
																},
															},
														},
														"accept_summary": dsschema.BoolAttribute{
															MarkdownDescription: "Accept summary",
															Computed:            true,
														},
														"default_information_originate": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Default information originate",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"metric": dsschema.Float64Attribute{
																	MarkdownDescription: "Metric",
																	Computed:            true,
																},
																"metric_type": dsschema.StringAttribute{
																	MarkdownDescription: "Metric type",
																	Computed:            true,
																},
															},
														},
														"default_route": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Default route",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"advertise": dsschema.SingleNestedAttribute{
																	MarkdownDescription: "Advertise",
																	Computed:            true,
																	Attributes: map[string]dsschema.Attribute{
																		"metric": dsschema.Float64Attribute{
																			MarkdownDescription: "Metric",
																			Computed:            true,
																		},
																		"type": dsschema.StringAttribute{
																			MarkdownDescription: "Type",
																			Computed:            true,
																		},
																	},
																},
																"disable": dsschema.SingleNestedAttribute{
																	MarkdownDescription: "Disable",
																	Computed:            true,
																	Attributes:          map[string]dsschema.Attribute{},
																},
															},
														},
														"no_summary": dsschema.BoolAttribute{
															MarkdownDescription: "No summary",
															Computed:            true,
														},
														"nssa_ext_range": dsschema.ListNestedAttribute{
															MarkdownDescription: "Nssa ext range",
															Computed:            true,
															NestedObject: dsschema.NestedAttributeObject{
																Attributes: map[string]dsschema.Attribute{
																	"advertise": dsschema.SingleNestedAttribute{
																		MarkdownDescription: "Advertise",
																		Computed:            true,
																		Attributes:          map[string]dsschema.Attribute{},
																	},
																	"name": dsschema.StringAttribute{
																		MarkdownDescription: "Name",
																		Computed:            true,
																	},
																	"suppress": dsschema.SingleNestedAttribute{
																		MarkdownDescription: "Suppress",
																		Computed:            true,
																		Attributes:          map[string]dsschema.Attribute{},
																	},
																},
															},
														},
													},
												},
												"stub": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Stub",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"abr": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Abr",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"export_list": dsschema.StringAttribute{
																	MarkdownDescription: "Export list",
																	Computed:            true,
																},
																"import_list": dsschema.StringAttribute{
																	MarkdownDescription: "Import list",
																	Computed:            true,
																},
																"inbound_filter_list": dsschema.StringAttribute{
																	MarkdownDescription: "Inbound filter list",
																	Computed:            true,
																},
																"outbound_filter_list": dsschema.StringAttribute{
																	MarkdownDescription: "Outbound filter list",
																	Computed:            true,
																},
															},
														},
														"accept_summary": dsschema.BoolAttribute{
															MarkdownDescription: "Accept summary",
															Computed:            true,
														},
														"default_route": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Default route",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"advertise": dsschema.SingleNestedAttribute{
																	MarkdownDescription: "Advertise",
																	Computed:            true,
																	Attributes: map[string]dsschema.Attribute{
																		"metric": dsschema.Float64Attribute{
																			MarkdownDescription: "Metric",
																			Computed:            true,
																		},
																	},
																},
																"disable": dsschema.SingleNestedAttribute{
																	MarkdownDescription: "Disable",
																	Computed:            true,
																	Attributes:          map[string]dsschema.Attribute{},
																},
															},
														},
														"default_route_metric": dsschema.Float64Attribute{
															MarkdownDescription: "Default route metric",
															Computed:            true,
														},
														"no_summary": dsschema.BoolAttribute{
															MarkdownDescription: "No summary",
															Computed:            true,
														},
													},
												},
											},
										},
										"virtual_link": dsschema.ListNestedAttribute{
											MarkdownDescription: "Virtual link",
											Computed:            true,
											NestedObject: dsschema.NestedAttributeObject{
												Attributes: map[string]dsschema.Attribute{
													"authentication": dsschema.StringAttribute{
														MarkdownDescription: "Authentication",
														Computed:            true,
													},
													"bfd": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Bfd",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"profile": dsschema.StringAttribute{
																MarkdownDescription: "Profile",
																Computed:            true,
															},
														},
													},
													"enable": dsschema.BoolAttribute{
														MarkdownDescription: "Enable",
														Computed:            true,
													},
													"instance_id": dsschema.Float64Attribute{
														MarkdownDescription: "Instance id",
														Computed:            true,
													},
													"interface_id": dsschema.Float64Attribute{
														MarkdownDescription: "Interface id",
														Computed:            true,
													},
													"name": dsschema.StringAttribute{
														MarkdownDescription: "Name",
														Computed:            true,
													},
													"neighbor_id": dsschema.StringAttribute{
														MarkdownDescription: "Neighbor id",
														Computed:            true,
													},
													"passive": dsschema.BoolAttribute{
														MarkdownDescription: "Passive",
														Computed:            true,
													},
													"timing": dsschema.StringAttribute{
														MarkdownDescription: "Timing",
														Computed:            true,
													},
													"transit_area_id": dsschema.StringAttribute{
														MarkdownDescription: "Transit area id",
														Computed:            true,
													},
													"vr_timing": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Vr timing",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"dead_counts": dsschema.Float64Attribute{
																MarkdownDescription: "Dead counts",
																Computed:            true,
															},
															"hello_interval": dsschema.Float64Attribute{
																MarkdownDescription: "Hello interval",
																Computed:            true,
															},
															"retransmit_interval": dsschema.Float64Attribute{
																MarkdownDescription: "Retransmit interval",
																Computed:            true,
															},
															"transit_delay": dsschema.Float64Attribute{
																MarkdownDescription: "Transit delay",
																Computed:            true,
															},
														},
													},
												},
											},
										},
										"vr_range": dsschema.ListNestedAttribute{
											MarkdownDescription: "Vr range",
											Computed:            true,
											NestedObject: dsschema.NestedAttributeObject{
												Attributes: map[string]dsschema.Attribute{
													"advertise": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Advertise",
														Computed:            true,
														Attributes:          map[string]dsschema.Attribute{},
													},
													"name": dsschema.StringAttribute{
														MarkdownDescription: "Name",
														Computed:            true,
													},
													"suppress": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Suppress",
														Computed:            true,
														Attributes:          map[string]dsschema.Attribute{},
													},
												},
											},
										},
									},
								},
							},
							"auth_profile": dsschema.ListNestedAttribute{
								MarkdownDescription: "Auth profile",
								Computed:            true,
								NestedObject: dsschema.NestedAttributeObject{
									Attributes: map[string]dsschema.Attribute{
										"md5": dsschema.ListNestedAttribute{
											MarkdownDescription: "Md5",
											Computed:            true,
											NestedObject: dsschema.NestedAttributeObject{
												Attributes: map[string]dsschema.Attribute{
													"key": dsschema.StringAttribute{
														MarkdownDescription: "Key",
														Computed:            true,
													},
													"name": dsschema.Float64Attribute{
														MarkdownDescription: "Name",
														Computed:            true,
													},
													"preferred": dsschema.BoolAttribute{
														MarkdownDescription: "Preferred",
														Computed:            true,
													},
												},
											},
										},
										"name": dsschema.StringAttribute{
											MarkdownDescription: "Name",
											Computed:            true,
										},
										"password": dsschema.StringAttribute{
											MarkdownDescription: "Password",
											Computed:            true,
										},
									},
								},
							},
							"enable": dsschema.BoolAttribute{
								MarkdownDescription: "Enable",
								Computed:            true,
							},
							"export_rules": dsschema.ListNestedAttribute{
								MarkdownDescription: "Export rules",
								Computed:            true,
								NestedObject: dsschema.NestedAttributeObject{
									Attributes: map[string]dsschema.Attribute{
										"metric": dsschema.Float64Attribute{
											MarkdownDescription: "Metric",
											Computed:            true,
										},
										"name": dsschema.StringAttribute{
											MarkdownDescription: "Name",
											Computed:            true,
										},
										"new_path_type": dsschema.StringAttribute{
											MarkdownDescription: "New path type",
											Computed:            true,
										},
										"new_tag": dsschema.StringAttribute{
											MarkdownDescription: "New tag",
											Computed:            true,
										},
									},
								},
							},
							"flood_prevention": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Flood prevention",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"hello": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Hello",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"enable": dsschema.BoolAttribute{
												MarkdownDescription: "Enable",
												Computed:            true,
											},
											"max_packet": dsschema.Float64Attribute{
												MarkdownDescription: "Max packet",
												Computed:            true,
											},
										},
									},
									"lsa": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Lsa",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"enable": dsschema.BoolAttribute{
												MarkdownDescription: "Enable",
												Computed:            true,
											},
											"max_packet": dsschema.Float64Attribute{
												MarkdownDescription: "Max packet",
												Computed:            true,
											},
										},
									},
								},
							},
							"global_bfd": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Global bfd",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"profile": dsschema.StringAttribute{
										MarkdownDescription: "Profile",
										Computed:            true,
									},
								},
							},
							"global_if_timer": dsschema.StringAttribute{
								MarkdownDescription: "Global if timer",
								Computed:            true,
							},
							"graceful_restart": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Graceful restart",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"enable": dsschema.BoolAttribute{
										MarkdownDescription: "Enable",
										Computed:            true,
									},
									"grace_period": dsschema.Float64Attribute{
										MarkdownDescription: "Grace period",
										Computed:            true,
									},
									"helper_enable": dsschema.BoolAttribute{
										MarkdownDescription: "Helper enable",
										Computed:            true,
									},
									"max_neighbor_restart_time": dsschema.Float64Attribute{
										MarkdownDescription: "Max neighbor restart time",
										Computed:            true,
									},
									"strict__l_s_a_checking": dsschema.BoolAttribute{
										MarkdownDescription: "Strict l s a checking",
										Computed:            true,
									},
								},
							},
							"redistribution_profile": dsschema.StringAttribute{
								MarkdownDescription: "Redistribution profile",
								Computed:            true,
							},
							"reject_default_route": dsschema.BoolAttribute{
								MarkdownDescription: "Reject default route",
								Computed:            true,
							},
							"rfc1583": dsschema.BoolAttribute{
								MarkdownDescription: "Rfc1583",
								Computed:            true,
							},
							"router_id": dsschema.StringAttribute{
								MarkdownDescription: "Router id",
								Computed:            true,
							},
							"spf_timer": dsschema.StringAttribute{
								MarkdownDescription: "Spf timer",
								Computed:            true,
							},
							"vr_timers": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Vr timers",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"lsa_interval": dsschema.Float64Attribute{
										MarkdownDescription: "Lsa interval",
										Computed:            true,
									},
									"spf_calculation_delay": dsschema.Float64Attribute{
										MarkdownDescription: "Spf calculation delay",
										Computed:            true,
									},
								},
							},
						},
					},
					"ospfv3": dsschema.SingleNestedAttribute{
						MarkdownDescription: "Ospfv3",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"allow_redist_default_route": dsschema.BoolAttribute{
								MarkdownDescription: "Allow redist default route",
								Computed:            true,
							},
							"area": dsschema.ListNestedAttribute{
								MarkdownDescription: "Area",
								Computed:            true,
								NestedObject: dsschema.NestedAttributeObject{
									Attributes: map[string]dsschema.Attribute{
										"authentication": dsschema.StringAttribute{
											MarkdownDescription: "Authentication",
											Computed:            true,
										},
										"interface": dsschema.ListNestedAttribute{
											MarkdownDescription: "Interface",
											Computed:            true,
											NestedObject: dsschema.NestedAttributeObject{
												Attributes: map[string]dsschema.Attribute{
													"authentication": dsschema.StringAttribute{
														MarkdownDescription: "Authentication",
														Computed:            true,
													},
													"bfd": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Bfd",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"profile": dsschema.StringAttribute{
																MarkdownDescription: "Profile",
																Computed:            true,
															},
														},
													},
													"enable": dsschema.BoolAttribute{
														MarkdownDescription: "Enable",
														Computed:            true,
													},
													"instance_id": dsschema.Float64Attribute{
														MarkdownDescription: "Instance id",
														Computed:            true,
													},
													"link_type": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Link type",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"broadcast": dsschema.SingleNestedAttribute{
																MarkdownDescription: "Broadcast",
																Computed:            true,
																Attributes:          map[string]dsschema.Attribute{},
															},
															"p2mp": dsschema.SingleNestedAttribute{
																MarkdownDescription: "P2mp",
																Computed:            true,
																Attributes: map[string]dsschema.Attribute{
																	"neighbor": dsschema.ListNestedAttribute{
																		MarkdownDescription: "Neighbor",
																		Computed:            true,
																		NestedObject: dsschema.NestedAttributeObject{
																			Attributes: map[string]dsschema.Attribute{
																				"name": dsschema.StringAttribute{
																					MarkdownDescription: "Name",
																					Computed:            true,
																				},
																				"priority": dsschema.Float64Attribute{
																					MarkdownDescription: "Priority",
																					Computed:            true,
																				},
																			},
																		},
																	},
																},
															},
															"p2p": dsschema.SingleNestedAttribute{
																MarkdownDescription: "P2p",
																Computed:            true,
																Attributes:          map[string]dsschema.Attribute{},
															},
														},
													},
													"metric": dsschema.Float64Attribute{
														MarkdownDescription: "Metric",
														Computed:            true,
													},
													"mtu_ignore": dsschema.BoolAttribute{
														MarkdownDescription: "Mtu ignore",
														Computed:            true,
													},
													"name": dsschema.StringAttribute{
														MarkdownDescription: "Name",
														Computed:            true,
													},
													"neighbor": dsschema.ListNestedAttribute{
														MarkdownDescription: "Neighbor",
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
													"passive": dsschema.BoolAttribute{
														MarkdownDescription: "Passive",
														Computed:            true,
													},
													"priority": dsschema.Float64Attribute{
														MarkdownDescription: "Priority",
														Computed:            true,
													},
													"timing": dsschema.StringAttribute{
														MarkdownDescription: "Timing",
														Computed:            true,
													},
													"vr_timing": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Vr timing",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"dead_counts": dsschema.Float64Attribute{
																MarkdownDescription: "Dead counts",
																Computed:            true,
															},
															"gr_delay": dsschema.Float64Attribute{
																MarkdownDescription: "Gr delay",
																Computed:            true,
															},
															"hello_interval": dsschema.Float64Attribute{
																MarkdownDescription: "Hello interval",
																Computed:            true,
															},
															"retransmit_interval": dsschema.Float64Attribute{
																MarkdownDescription: "Retransmit interval",
																Computed:            true,
															},
															"transit_delay": dsschema.Float64Attribute{
																MarkdownDescription: "Transit delay",
																Computed:            true,
															},
														},
													},
												},
											},
										},
										"name": dsschema.StringAttribute{
											MarkdownDescription: "Name",
											Computed:            true,
										},
										"range": dsschema.ListNestedAttribute{
											MarkdownDescription: "Range",
											Computed:            true,
											NestedObject: dsschema.NestedAttributeObject{
												Attributes: map[string]dsschema.Attribute{
													"advertise": dsschema.BoolAttribute{
														MarkdownDescription: "Advertise",
														Computed:            true,
													},
													"name": dsschema.StringAttribute{
														MarkdownDescription: "Name",
														Computed:            true,
													},
												},
											},
										},
										"type": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Type",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"normal": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Normal",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"abr": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Abr",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"export_list": dsschema.StringAttribute{
																	MarkdownDescription: "Export list",
																	Computed:            true,
																},
																"import_list": dsschema.StringAttribute{
																	MarkdownDescription: "Import list",
																	Computed:            true,
																},
																"inbound_filter_list": dsschema.StringAttribute{
																	MarkdownDescription: "Inbound filter list",
																	Computed:            true,
																},
																"outbound_filter_list": dsschema.StringAttribute{
																	MarkdownDescription: "Outbound filter list",
																	Computed:            true,
																},
															},
														},
													},
												},
												"nssa": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Nssa",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"abr": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Abr",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"export_list": dsschema.StringAttribute{
																	MarkdownDescription: "Export list",
																	Computed:            true,
																},
																"import_list": dsschema.StringAttribute{
																	MarkdownDescription: "Import list",
																	Computed:            true,
																},
																"inbound_filter_list": dsschema.StringAttribute{
																	MarkdownDescription: "Inbound filter list",
																	Computed:            true,
																},
																"nssa_ext_range": dsschema.ListNestedAttribute{
																	MarkdownDescription: "Nssa ext range",
																	Computed:            true,
																	NestedObject: dsschema.NestedAttributeObject{
																		Attributes: map[string]dsschema.Attribute{
																			"advertise": dsschema.SingleNestedAttribute{
																				MarkdownDescription: "Advertise",
																				Computed:            true,
																				Attributes:          map[string]dsschema.Attribute{},
																			},
																			"name": dsschema.StringAttribute{
																				MarkdownDescription: "Name",
																				Computed:            true,
																			},
																			"route_tag": dsschema.Float64Attribute{
																				MarkdownDescription: "Route tag",
																				Computed:            true,
																			},
																			"suppress": dsschema.SingleNestedAttribute{
																				MarkdownDescription: "Suppress",
																				Computed:            true,
																				Attributes:          map[string]dsschema.Attribute{},
																			},
																		},
																	},
																},
																"outbound_filter_list": dsschema.StringAttribute{
																	MarkdownDescription: "Outbound filter list",
																	Computed:            true,
																},
															},
														},
														"accept_summary": dsschema.BoolAttribute{
															MarkdownDescription: "Accept summary",
															Computed:            true,
														},
														"default_information_originate": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Default information originate",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"metric": dsschema.Float64Attribute{
																	MarkdownDescription: "Metric",
																	Computed:            true,
																},
																"metric_type": dsschema.StringAttribute{
																	MarkdownDescription: "Metric type",
																	Computed:            true,
																},
															},
														},
														"default_route": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Default route",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"advertise": dsschema.SingleNestedAttribute{
																	MarkdownDescription: "Advertise",
																	Computed:            true,
																	Attributes: map[string]dsschema.Attribute{
																		"metric": dsschema.Float64Attribute{
																			MarkdownDescription: "Metric",
																			Computed:            true,
																		},
																		"type": dsschema.StringAttribute{
																			MarkdownDescription: "Type",
																			Computed:            true,
																		},
																	},
																},
																"disable": dsschema.SingleNestedAttribute{
																	MarkdownDescription: "Disable",
																	Computed:            true,
																	Attributes:          map[string]dsschema.Attribute{},
																},
															},
														},
														"no_summary": dsschema.BoolAttribute{
															MarkdownDescription: "No summary",
															Computed:            true,
														},
														"nssa_ext_range": dsschema.ListNestedAttribute{
															MarkdownDescription: "Nssa ext range",
															Computed:            true,
															NestedObject: dsschema.NestedAttributeObject{
																Attributes: map[string]dsschema.Attribute{
																	"advertise": dsschema.SingleNestedAttribute{
																		MarkdownDescription: "Advertise",
																		Computed:            true,
																		Attributes:          map[string]dsschema.Attribute{},
																	},
																	"name": dsschema.StringAttribute{
																		MarkdownDescription: "Name",
																		Computed:            true,
																	},
																	"route_tag": dsschema.Float64Attribute{
																		MarkdownDescription: "Route tag",
																		Computed:            true,
																	},
																	"suppress": dsschema.SingleNestedAttribute{
																		MarkdownDescription: "Suppress",
																		Computed:            true,
																		Attributes:          map[string]dsschema.Attribute{},
																	},
																},
															},
														},
													},
												},
												"stub": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Stub",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"abr": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Abr",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"export_list": dsschema.StringAttribute{
																	MarkdownDescription: "Export list",
																	Computed:            true,
																},
																"import_list": dsschema.StringAttribute{
																	MarkdownDescription: "Import list",
																	Computed:            true,
																},
																"inbound_filter_list": dsschema.StringAttribute{
																	MarkdownDescription: "Inbound filter list",
																	Computed:            true,
																},
																"outbound_filter_list": dsschema.StringAttribute{
																	MarkdownDescription: "Outbound filter list",
																	Computed:            true,
																},
															},
														},
														"accept_summary": dsschema.BoolAttribute{
															MarkdownDescription: "Accept summary",
															Computed:            true,
														},
														"default_route": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Default route",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"advertise": dsschema.SingleNestedAttribute{
																	MarkdownDescription: "Advertise",
																	Computed:            true,
																	Attributes: map[string]dsschema.Attribute{
																		"metric": dsschema.Float64Attribute{
																			MarkdownDescription: "Metric",
																			Computed:            true,
																		},
																	},
																},
																"disable": dsschema.SingleNestedAttribute{
																	MarkdownDescription: "Disable",
																	Computed:            true,
																	Attributes:          map[string]dsschema.Attribute{},
																},
															},
														},
														"default_route_metric": dsschema.Float64Attribute{
															MarkdownDescription: "Default route metric",
															Computed:            true,
														},
														"no_summary": dsschema.BoolAttribute{
															MarkdownDescription: "No summary",
															Computed:            true,
														},
													},
												},
											},
										},
										"virtual_link": dsschema.ListNestedAttribute{
											MarkdownDescription: "Virtual link",
											Computed:            true,
											NestedObject: dsschema.NestedAttributeObject{
												Attributes: map[string]dsschema.Attribute{
													"authentication": dsschema.StringAttribute{
														MarkdownDescription: "Authentication",
														Computed:            true,
													},
													"bfd": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Bfd",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"profile": dsschema.StringAttribute{
																MarkdownDescription: "Profile",
																Computed:            true,
															},
														},
													},
													"enable": dsschema.BoolAttribute{
														MarkdownDescription: "Enable",
														Computed:            true,
													},
													"instance_id": dsschema.Float64Attribute{
														MarkdownDescription: "Instance id",
														Computed:            true,
													},
													"interface_id": dsschema.Float64Attribute{
														MarkdownDescription: "Interface id",
														Computed:            true,
													},
													"name": dsschema.StringAttribute{
														MarkdownDescription: "Name",
														Computed:            true,
													},
													"neighbor_id": dsschema.StringAttribute{
														MarkdownDescription: "Neighbor id",
														Computed:            true,
													},
													"passive": dsschema.BoolAttribute{
														MarkdownDescription: "Passive",
														Computed:            true,
													},
													"timing": dsschema.StringAttribute{
														MarkdownDescription: "Timing",
														Computed:            true,
													},
													"transit_area_id": dsschema.StringAttribute{
														MarkdownDescription: "Transit area id",
														Computed:            true,
													},
													"vr_timing": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Vr timing",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"dead_counts": dsschema.Float64Attribute{
																MarkdownDescription: "Dead counts",
																Computed:            true,
															},
															"hello_interval": dsschema.Float64Attribute{
																MarkdownDescription: "Hello interval",
																Computed:            true,
															},
															"retransmit_interval": dsschema.Float64Attribute{
																MarkdownDescription: "Retransmit interval",
																Computed:            true,
															},
															"transit_delay": dsschema.Float64Attribute{
																MarkdownDescription: "Transit delay",
																Computed:            true,
															},
														},
													},
												},
											},
										},
										"vr_range": dsschema.ListNestedAttribute{
											MarkdownDescription: "Vr range",
											Computed:            true,
											NestedObject: dsschema.NestedAttributeObject{
												Attributes: map[string]dsschema.Attribute{
													"advertise": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Advertise",
														Computed:            true,
														Attributes:          map[string]dsschema.Attribute{},
													},
													"name": dsschema.StringAttribute{
														MarkdownDescription: "Name",
														Computed:            true,
													},
													"suppress": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Suppress",
														Computed:            true,
														Attributes:          map[string]dsschema.Attribute{},
													},
												},
											},
										},
									},
								},
							},
							"auth_profile": dsschema.ListNestedAttribute{
								MarkdownDescription: "Auth profile",
								Computed:            true,
								NestedObject: dsschema.NestedAttributeObject{
									Attributes: map[string]dsschema.Attribute{
										"ah": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Ah",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"md5": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Md5",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"key": dsschema.StringAttribute{
															MarkdownDescription: "Key",
															Computed:            true,
														},
													},
												},
												"sha1": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Sha1",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"key": dsschema.StringAttribute{
															MarkdownDescription: "Key",
															Computed:            true,
														},
													},
												},
												"sha256": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Sha256",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"key": dsschema.StringAttribute{
															MarkdownDescription: "Key",
															Computed:            true,
														},
													},
												},
												"sha384": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Sha384",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"key": dsschema.StringAttribute{
															MarkdownDescription: "Key",
															Computed:            true,
														},
													},
												},
												"sha512": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Sha512",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"key": dsschema.StringAttribute{
															MarkdownDescription: "Key",
															Computed:            true,
														},
													},
												},
											},
										},
										"esp": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Esp",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"authentication": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Authentication",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"md5": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Md5",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"key": dsschema.StringAttribute{
																	MarkdownDescription: "Key",
																	Computed:            true,
																},
															},
														},
														"none": dsschema.SingleNestedAttribute{
															MarkdownDescription: "None",
															Computed:            true,
															Attributes:          map[string]dsschema.Attribute{},
														},
														"sha1": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Sha1",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"key": dsschema.StringAttribute{
																	MarkdownDescription: "Key",
																	Computed:            true,
																},
															},
														},
														"sha256": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Sha256",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"key": dsschema.StringAttribute{
																	MarkdownDescription: "Key",
																	Computed:            true,
																},
															},
														},
														"sha384": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Sha384",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"key": dsschema.StringAttribute{
																	MarkdownDescription: "Key",
																	Computed:            true,
																},
															},
														},
														"sha512": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Sha512",
															Computed:            true,
															Attributes: map[string]dsschema.Attribute{
																"key": dsschema.StringAttribute{
																	MarkdownDescription: "Key",
																	Computed:            true,
																},
															},
														},
													},
												},
												"encryption": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Encryption",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"algorithm": dsschema.StringAttribute{
															MarkdownDescription: "Algorithm",
															Computed:            true,
														},
														"key": dsschema.StringAttribute{
															MarkdownDescription: "Key",
															Computed:            true,
														},
													},
												},
											},
										},
										"name": dsschema.StringAttribute{
											MarkdownDescription: "Name",
											Computed:            true,
										},
										"spi": dsschema.StringAttribute{
											MarkdownDescription: "Spi",
											Computed:            true,
										},
									},
								},
							},
							"disable_transit_traffic": dsschema.BoolAttribute{
								MarkdownDescription: "Disable transit traffic",
								Computed:            true,
							},
							"enable": dsschema.BoolAttribute{
								MarkdownDescription: "Enable",
								Computed:            true,
							},
							"export_rules": dsschema.ListNestedAttribute{
								MarkdownDescription: "Export rules",
								Computed:            true,
								NestedObject: dsschema.NestedAttributeObject{
									Attributes: map[string]dsschema.Attribute{
										"metric": dsschema.Float64Attribute{
											MarkdownDescription: "Metric",
											Computed:            true,
										},
										"name": dsschema.StringAttribute{
											MarkdownDescription: "Name",
											Computed:            true,
										},
										"new_path_type": dsschema.StringAttribute{
											MarkdownDescription: "New path type",
											Computed:            true,
										},
										"new_tag": dsschema.StringAttribute{
											MarkdownDescription: "New tag",
											Computed:            true,
										},
									},
								},
							},
							"global_bfd": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Global bfd",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"profile": dsschema.StringAttribute{
										MarkdownDescription: "Profile",
										Computed:            true,
									},
								},
							},
							"global_if_timer": dsschema.StringAttribute{
								MarkdownDescription: "Global if timer",
								Computed:            true,
							},
							"graceful_restart": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Graceful restart",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"enable": dsschema.BoolAttribute{
										MarkdownDescription: "Enable",
										Computed:            true,
									},
									"grace_period": dsschema.Float64Attribute{
										MarkdownDescription: "Grace period",
										Computed:            true,
									},
									"helper_enable": dsschema.BoolAttribute{
										MarkdownDescription: "Helper enable",
										Computed:            true,
									},
									"max_neighbor_restart_time": dsschema.Float64Attribute{
										MarkdownDescription: "Max neighbor restart time",
										Computed:            true,
									},
									"strict__l_s_a_checking": dsschema.BoolAttribute{
										MarkdownDescription: "Strict l s a checking",
										Computed:            true,
									},
								},
							},
							"redistribution_profile": dsschema.StringAttribute{
								MarkdownDescription: "Redistribution profile",
								Computed:            true,
							},
							"reject_default_route": dsschema.BoolAttribute{
								MarkdownDescription: "Reject default route",
								Computed:            true,
							},
							"router_id": dsschema.StringAttribute{
								MarkdownDescription: "Router id",
								Computed:            true,
							},
							"spf_timer": dsschema.StringAttribute{
								MarkdownDescription: "Spf timer",
								Computed:            true,
							},
							"vr_timers": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Vr timers",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"lsa_interval": dsschema.Float64Attribute{
										MarkdownDescription: "Lsa interval",
										Computed:            true,
									},
									"spf_calculation_delay": dsschema.Float64Attribute{
										MarkdownDescription: "Spf calculation delay",
										Computed:            true,
									},
								},
							},
						},
					},
					"rib_filter": dsschema.SingleNestedAttribute{
						MarkdownDescription: "Rib filter",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"ipv4": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Ipv4",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"bgp": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Bgp",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"route_map": dsschema.StringAttribute{
												MarkdownDescription: "Route map",
												Computed:            true,
											},
										},
									},
									"ospf": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Ospf",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"route_map": dsschema.StringAttribute{
												MarkdownDescription: "Route map",
												Computed:            true,
											},
										},
									},
									"rip": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Rip",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"route_map": dsschema.StringAttribute{
												MarkdownDescription: "Route map",
												Computed:            true,
											},
										},
									},
									"static": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Static",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"route_map": dsschema.StringAttribute{
												MarkdownDescription: "Route map",
												Computed:            true,
											},
										},
									},
								},
							},
							"ipv6": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Ipv6",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"bgp": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Bgp",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"route_map": dsschema.StringAttribute{
												MarkdownDescription: "Route map",
												Computed:            true,
											},
										},
									},
									"ospfv3": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Ospfv3",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"route_map": dsschema.StringAttribute{
												MarkdownDescription: "Route map",
												Computed:            true,
											},
										},
									},
									"static": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Static",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"route_map": dsschema.StringAttribute{
												MarkdownDescription: "Route map",
												Computed:            true,
											},
										},
									},
								},
							},
						},
					},
					"rip": dsschema.SingleNestedAttribute{
						MarkdownDescription: "Rip",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"auth_profile": dsschema.StringAttribute{
								MarkdownDescription: "Auth profile",
								Computed:            true,
							},
							"default_information_originate": dsschema.BoolAttribute{
								MarkdownDescription: "Default information originate",
								Computed:            true,
							},
							"enable": dsschema.BoolAttribute{
								MarkdownDescription: "Enable",
								Computed:            true,
							},
							"global_bfd": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Global bfd",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"profile": dsschema.StringAttribute{
										MarkdownDescription: "Profile",
										Computed:            true,
									},
								},
							},
							"global_inbound_distribute_list": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Global inbound distribute list",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"access_list": dsschema.StringAttribute{
										MarkdownDescription: "Access list",
										Computed:            true,
									},
								},
							},
							"global_outbound_distribute_list": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Global outbound distribute list",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"access_list": dsschema.StringAttribute{
										MarkdownDescription: "Access list",
										Computed:            true,
									},
								},
							},
							"global_timer": dsschema.StringAttribute{
								MarkdownDescription: "Global timer",
								Computed:            true,
							},
							"interface": dsschema.ListNestedAttribute{
								MarkdownDescription: "Interface",
								Computed:            true,
								NestedObject: dsschema.NestedAttributeObject{
									Attributes: map[string]dsschema.Attribute{
										"authentication": dsschema.StringAttribute{
											MarkdownDescription: "Authentication",
											Computed:            true,
										},
										"bfd": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Bfd",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"profile": dsschema.StringAttribute{
													MarkdownDescription: "Profile",
													Computed:            true,
												},
											},
										},
										"enable": dsschema.BoolAttribute{
											MarkdownDescription: "Enable",
											Computed:            true,
										},
										"interface_inbound_distribute_list": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Interface inbound distribute list",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"access_list": dsschema.StringAttribute{
													MarkdownDescription: "Access list",
													Computed:            true,
												},
												"metric": dsschema.Float64Attribute{
													MarkdownDescription: "Metric",
													Computed:            true,
												},
											},
										},
										"interface_outbound_distribute_list": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Interface outbound distribute list",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"access_list": dsschema.StringAttribute{
													MarkdownDescription: "Access list",
													Computed:            true,
												},
												"metric": dsschema.Float64Attribute{
													MarkdownDescription: "Metric",
													Computed:            true,
												},
											},
										},
										"mode": dsschema.StringAttribute{
											MarkdownDescription: "Mode",
											Computed:            true,
										},
										"name": dsschema.StringAttribute{
											MarkdownDescription: "Name",
											Computed:            true,
										},
										"split_horizon": dsschema.StringAttribute{
											MarkdownDescription: "Split horizon",
											Computed:            true,
										},
									},
								},
							},
							"redistribution_profile": dsschema.StringAttribute{
								MarkdownDescription: "Redistribution profile",
								Computed:            true,
							},
						},
					},
					"routing_table": dsschema.SingleNestedAttribute{
						MarkdownDescription: "Routing table",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"ip": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Ip",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"static_route": dsschema.ListNestedAttribute{
										MarkdownDescription: "Static route",
										Computed:            true,
										NestedObject: dsschema.NestedAttributeObject{
											Attributes: map[string]dsschema.Attribute{
												"admin_dist": dsschema.Float64Attribute{
													MarkdownDescription: "Admin dist",
													Computed:            true,
												},
												"bfd": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Bfd",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"profile": dsschema.StringAttribute{
															MarkdownDescription: "Profile",
															Computed:            true,
														},
													},
												},
												"destination": dsschema.StringAttribute{
													MarkdownDescription: "Destination",
													Computed:            true,
												},
												"interface": dsschema.StringAttribute{
													MarkdownDescription: "Interface",
													Computed:            true,
												},
												"metric": dsschema.Float64Attribute{
													MarkdownDescription: "Metric",
													Computed:            true,
												},
												"name": dsschema.StringAttribute{
													MarkdownDescription: "Name",
													Computed:            true,
												},
												"nexthop": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Nexthop",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"discard": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Discard",
															Computed:            true,
															Attributes:          map[string]dsschema.Attribute{},
														},
														"fqdn": dsschema.StringAttribute{
															MarkdownDescription: "Fqdn",
															Computed:            true,
														},
														"ipv6_address": dsschema.StringAttribute{
															MarkdownDescription: "Ipv6 address",
															Computed:            true,
														},
														"next_lr": dsschema.StringAttribute{
															MarkdownDescription: "Next lr",
															Computed:            true,
														},
														"next_vr": dsschema.StringAttribute{
															MarkdownDescription: "Next vr",
															Computed:            true,
														},
														"receive": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Receive",
															Computed:            true,
															Attributes:          map[string]dsschema.Attribute{},
														},
														"tunnel": dsschema.StringAttribute{
															MarkdownDescription: "Tunnel",
															Computed:            true,
														},
													},
												},
												"path_monitor": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Path monitor",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"enable": dsschema.BoolAttribute{
															MarkdownDescription: "Enable",
															Computed:            true,
														},
														"failure_condition": dsschema.StringAttribute{
															MarkdownDescription: "Failure condition",
															Computed:            true,
														},
														"hold_time": dsschema.Float64Attribute{
															MarkdownDescription: "Hold time",
															Computed:            true,
														},
														"monitor_destinations": dsschema.ListNestedAttribute{
															MarkdownDescription: "Monitor destinations",
															Computed:            true,
															NestedObject: dsschema.NestedAttributeObject{
																Attributes: map[string]dsschema.Attribute{
																	"count": dsschema.Float64Attribute{
																		MarkdownDescription: "Count",
																		Computed:            true,
																	},
																	"destination": dsschema.StringAttribute{
																		MarkdownDescription: "Destination",
																		Computed:            true,
																	},
																	"destination_fqdn": dsschema.StringAttribute{
																		MarkdownDescription: "Destination fqdn",
																		Computed:            true,
																	},
																	"enable": dsschema.BoolAttribute{
																		MarkdownDescription: "Enable",
																		Computed:            true,
																	},
																	"interval": dsschema.Float64Attribute{
																		MarkdownDescription: "Interval",
																		Computed:            true,
																	},
																	"name": dsschema.StringAttribute{
																		MarkdownDescription: "Name",
																		Computed:            true,
																	},
																	"source": dsschema.StringAttribute{
																		MarkdownDescription: "Source",
																		Computed:            true,
																	},
																},
															},
														},
													},
												},
												"route_table": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Route table",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"both": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Both",
															Computed:            true,
															Attributes:          map[string]dsschema.Attribute{},
														},
														"multicast": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Multicast",
															Computed:            true,
															Attributes:          map[string]dsschema.Attribute{},
														},
														"no_install": dsschema.SingleNestedAttribute{
															MarkdownDescription: "No install",
															Computed:            true,
															Attributes:          map[string]dsschema.Attribute{},
														},
														"unicast": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Unicast",
															Computed:            true,
															Attributes:          map[string]dsschema.Attribute{},
														},
													},
												},
											},
										},
									},
								},
							},
							"ipv6": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Ipv6",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"static_route": dsschema.ListNestedAttribute{
										MarkdownDescription: "Static route",
										Computed:            true,
										NestedObject: dsschema.NestedAttributeObject{
											Attributes: map[string]dsschema.Attribute{
												"admin_dist": dsschema.Float64Attribute{
													MarkdownDescription: "Admin dist",
													Computed:            true,
												},
												"bfd": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Bfd",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"profile": dsschema.StringAttribute{
															MarkdownDescription: "Profile",
															Computed:            true,
														},
													},
												},
												"destination": dsschema.StringAttribute{
													MarkdownDescription: "Destination",
													Computed:            true,
												},
												"interface": dsschema.StringAttribute{
													MarkdownDescription: "Interface",
													Computed:            true,
												},
												"metric": dsschema.Float64Attribute{
													MarkdownDescription: "Metric",
													Computed:            true,
												},
												"name": dsschema.StringAttribute{
													MarkdownDescription: "Name",
													Computed:            true,
												},
												"nexthop": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Nexthop",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"discard": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Discard",
															Computed:            true,
															Attributes:          map[string]dsschema.Attribute{},
														},
														"fqdn": dsschema.StringAttribute{
															MarkdownDescription: "Fqdn",
															Computed:            true,
														},
														"ipv6_address": dsschema.StringAttribute{
															MarkdownDescription: "Ipv6 address",
															Computed:            true,
														},
														"next_lr": dsschema.StringAttribute{
															MarkdownDescription: "Next lr",
															Computed:            true,
														},
														"next_vr": dsschema.StringAttribute{
															MarkdownDescription: "Next vr",
															Computed:            true,
														},
														"receive": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Receive",
															Computed:            true,
															Attributes:          map[string]dsschema.Attribute{},
														},
														"tunnel": dsschema.StringAttribute{
															MarkdownDescription: "Tunnel",
															Computed:            true,
														},
													},
												},
												"option": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Option",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"passive": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Passive",
															Computed:            true,
															Attributes:          map[string]dsschema.Attribute{},
														},
													},
												},
												"path_monitor": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Path monitor",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"enable": dsschema.BoolAttribute{
															MarkdownDescription: "Enable",
															Computed:            true,
														},
														"failure_condition": dsschema.StringAttribute{
															MarkdownDescription: "Failure condition",
															Computed:            true,
														},
														"hold_time": dsschema.Float64Attribute{
															MarkdownDescription: "Hold time",
															Computed:            true,
														},
														"monitor_destinations": dsschema.ListNestedAttribute{
															MarkdownDescription: "Monitor destinations",
															Computed:            true,
															NestedObject: dsschema.NestedAttributeObject{
																Attributes: map[string]dsschema.Attribute{
																	"count": dsschema.Float64Attribute{
																		MarkdownDescription: "Count",
																		Computed:            true,
																	},
																	"destination": dsschema.StringAttribute{
																		MarkdownDescription: "Destination",
																		Computed:            true,
																	},
																	"destination_fqdn": dsschema.StringAttribute{
																		MarkdownDescription: "Destination fqdn",
																		Computed:            true,
																	},
																	"enable": dsschema.BoolAttribute{
																		MarkdownDescription: "Enable",
																		Computed:            true,
																	},
																	"interval": dsschema.Float64Attribute{
																		MarkdownDescription: "Interval",
																		Computed:            true,
																	},
																	"name": dsschema.StringAttribute{
																		MarkdownDescription: "Name",
																		Computed:            true,
																	},
																	"source": dsschema.StringAttribute{
																		MarkdownDescription: "Source",
																		Computed:            true,
																	},
																},
															},
														},
													},
												},
												"route_table": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Route table",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"both": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Both",
															Computed:            true,
															Attributes:          map[string]dsschema.Attribute{},
														},
														"multicast": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Multicast",
															Computed:            true,
															Attributes:          map[string]dsschema.Attribute{},
														},
														"no_install": dsschema.SingleNestedAttribute{
															MarkdownDescription: "No install",
															Computed:            true,
															Attributes:          map[string]dsschema.Attribute{},
														},
														"unicast": dsschema.SingleNestedAttribute{
															MarkdownDescription: "Unicast",
															Computed:            true,
															Attributes:          map[string]dsschema.Attribute{},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
					"sdwan_type": dsschema.StringAttribute{
						MarkdownDescription: "Sdwan type",
						Computed:            true,
					},
					"vr_admin_dists": dsschema.SingleNestedAttribute{
						MarkdownDescription: "Vr admin dists",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"ebgp": dsschema.Float64Attribute{
								MarkdownDescription: "Ebgp",
								Computed:            true,
							},
							"ibgp": dsschema.Float64Attribute{
								MarkdownDescription: "Ibgp",
								Computed:            true,
							},
							"ospf_ext": dsschema.Float64Attribute{
								MarkdownDescription: "Ospf ext",
								Computed:            true,
							},
							"ospf_int": dsschema.Float64Attribute{
								MarkdownDescription: "Ospf int",
								Computed:            true,
							},
							"ospfv3_ext": dsschema.Float64Attribute{
								MarkdownDescription: "Ospfv3 ext",
								Computed:            true,
							},
							"ospfv3_int": dsschema.Float64Attribute{
								MarkdownDescription: "Ospfv3 int",
								Computed:            true,
							},
							"rip": dsschema.Float64Attribute{
								MarkdownDescription: "Rip",
								Computed:            true,
							},
							"static": dsschema.Float64Attribute{
								MarkdownDescription: "Static",
								Computed:            true,
							},
							"static_ipv6": dsschema.Float64Attribute{
								MarkdownDescription: "Static ipv6",
								Computed:            true,
							},
						},
					},
					"zone_name": dsschema.StringAttribute{
						MarkdownDescription: "Zone name",
						Computed:            true,
					},
				},
			},
		},
	},
}

// LogicalRoutersListModel represents the data model for a list data source.
type LogicalRoutersListModel struct {
	Tfid    types.String     `tfsdk:"tfid"`
	Data    []LogicalRouters `tfsdk:"data"`
	Limit   types.Int64      `tfsdk:"limit"`
	Offset  types.Int64      `tfsdk:"offset"`
	Name    types.String     `tfsdk:"name"`
	Total   types.Int64      `tfsdk:"total"`
	Folder  types.String     `tfsdk:"folder"`
	Snippet types.String     `tfsdk:"snippet"`
	Device  types.String     `tfsdk:"device"`
}

// LogicalRoutersListDataSourceSchema defines the schema for a list data source.
var LogicalRoutersListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: LogicalRoutersDataSourceSchema.Attributes,
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
