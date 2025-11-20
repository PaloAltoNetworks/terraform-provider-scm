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

// BgpRouteMapRedistributions represents the Terraform model for BgpRouteMapRedistributions
type BgpRouteMapRedistributions struct {
	Tfid            types.String          `tfsdk:"tfid"`
	Bgp             basetypes.ObjectValue `tfsdk:"bgp"`
	ConnectedStatic basetypes.ObjectValue `tfsdk:"connected_static"`
	Description     basetypes.StringValue `tfsdk:"description"`
	Device          basetypes.StringValue `tfsdk:"device"`
	Folder          basetypes.StringValue `tfsdk:"folder"`
	Id              basetypes.StringValue `tfsdk:"id"`
	Name            basetypes.StringValue `tfsdk:"name"`
	Ospf            basetypes.ObjectValue `tfsdk:"ospf"`
	Snippet         basetypes.StringValue `tfsdk:"snippet"`
}

// BgpRouteMapRedistributionsBgp represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsBgp struct {
	Ospf basetypes.ObjectValue `tfsdk:"ospf"`
	Rib  basetypes.ObjectValue `tfsdk:"rib"`
}

// BgpRouteMapRedistributionsBgpOspf represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsBgpOspf struct {
	RouteMap basetypes.ListValue `tfsdk:"route_map"`
}

// BgpRouteMapRedistributionsBgpOspfRouteMapInner represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsBgpOspfRouteMapInner struct {
	Action      basetypes.StringValue `tfsdk:"action"`
	Description basetypes.StringValue `tfsdk:"description"`
	Match       basetypes.ObjectValue `tfsdk:"match"`
	Name        basetypes.Int64Value  `tfsdk:"name"`
	Set         basetypes.ObjectValue `tfsdk:"set"`
}

// BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch struct {
	AsPathAccessList  basetypes.StringValue `tfsdk:"as_path_access_list"`
	ExtendedCommunity basetypes.StringValue `tfsdk:"extended_community"`
	Interface         basetypes.StringValue `tfsdk:"interface"`
	Ipv4              basetypes.ObjectValue `tfsdk:"ipv4"`
	LargeCommunity    basetypes.StringValue `tfsdk:"large_community"`
	LocalPreference   basetypes.Int64Value  `tfsdk:"local_preference"`
	Metric            basetypes.Int64Value  `tfsdk:"metric"`
	Origin            basetypes.StringValue `tfsdk:"origin"`
	Peer              basetypes.StringValue `tfsdk:"peer"`
	RegularCommunity  basetypes.StringValue `tfsdk:"regular_community"`
	Tag               basetypes.Int64Value  `tfsdk:"tag"`
}

// BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4 represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4 struct {
	Address     basetypes.ObjectValue `tfsdk:"address"`
	NextHop     basetypes.ObjectValue `tfsdk:"next_hop"`
	RouteSource basetypes.ObjectValue `tfsdk:"route_source"`
}

// BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address struct {
	AccessList basetypes.StringValue `tfsdk:"access_list"`
	PrefixList basetypes.StringValue `tfsdk:"prefix_list"`
}

// BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop struct {
	AccessList basetypes.StringValue `tfsdk:"access_list"`
	PrefixList basetypes.StringValue `tfsdk:"prefix_list"`
}

// BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource struct {
	AccessList basetypes.StringValue `tfsdk:"access_list"`
	PrefixList basetypes.StringValue `tfsdk:"prefix_list"`
}

// BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet struct {
	Metric     basetypes.ObjectValue `tfsdk:"metric"`
	MetricType basetypes.StringValue `tfsdk:"metric_type"`
	Tag        basetypes.Int64Value  `tfsdk:"tag"`
}

// BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric struct {
	Action basetypes.StringValue `tfsdk:"action"`
	Value  basetypes.Int64Value  `tfsdk:"value"`
}

// BgpRouteMapRedistributionsBgpRib represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsBgpRib struct {
	RouteMap basetypes.ListValue `tfsdk:"route_map"`
}

// BgpRouteMapRedistributionsBgpRibRouteMapInner represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsBgpRibRouteMapInner struct {
	Action      basetypes.StringValue `tfsdk:"action"`
	Description basetypes.StringValue `tfsdk:"description"`
	Match       basetypes.ObjectValue `tfsdk:"match"`
	Name        basetypes.Int64Value  `tfsdk:"name"`
	Set         basetypes.ObjectValue `tfsdk:"set"`
}

// BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch struct {
	AsPathAccessList  basetypes.StringValue `tfsdk:"as_path_access_list"`
	ExtendedCommunity basetypes.StringValue `tfsdk:"extended_community"`
	Interface         basetypes.StringValue `tfsdk:"interface"`
	Ipv4              basetypes.ObjectValue `tfsdk:"ipv4"`
	LargeCommunity    basetypes.StringValue `tfsdk:"large_community"`
	LocalPreference   basetypes.Int64Value  `tfsdk:"local_preference"`
	Metric            basetypes.Int64Value  `tfsdk:"metric"`
	Origin            basetypes.StringValue `tfsdk:"origin"`
	Peer              basetypes.StringValue `tfsdk:"peer"`
	RegularCommunity  basetypes.StringValue `tfsdk:"regular_community"`
	Tag               basetypes.Int64Value  `tfsdk:"tag"`
}

// BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4 represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4 struct {
	Address     basetypes.ObjectValue `tfsdk:"address"`
	NextHop     basetypes.ObjectValue `tfsdk:"next_hop"`
	RouteSource basetypes.ObjectValue `tfsdk:"route_source"`
}

// BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address struct {
	AccessList basetypes.StringValue `tfsdk:"access_list"`
	PrefixList basetypes.StringValue `tfsdk:"prefix_list"`
}

// BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop struct {
	AccessList basetypes.StringValue `tfsdk:"access_list"`
	PrefixList basetypes.StringValue `tfsdk:"prefix_list"`
}

// BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource struct {
	AccessList basetypes.StringValue `tfsdk:"access_list"`
	PrefixList basetypes.StringValue `tfsdk:"prefix_list"`
}

// BgpRouteMapRedistributionsBgpRibRouteMapInnerSet represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsBgpRibRouteMapInnerSet struct {
	SourceAddress basetypes.StringValue `tfsdk:"source_address"`
}

// BgpRouteMapRedistributionsConnectedStatic represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStatic struct {
	Bgp  basetypes.ObjectValue `tfsdk:"bgp"`
	Ospf basetypes.ObjectValue `tfsdk:"ospf"`
	Rib  basetypes.ObjectValue `tfsdk:"rib"`
}

// BgpRouteMapRedistributionsConnectedStaticBgp represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticBgp struct {
	RouteMap basetypes.ListValue `tfsdk:"route_map"`
}

// BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner struct {
	Action      basetypes.StringValue `tfsdk:"action"`
	Description basetypes.StringValue `tfsdk:"description"`
	Match       basetypes.ObjectValue `tfsdk:"match"`
	Name        basetypes.Int64Value  `tfsdk:"name"`
	Set         basetypes.ObjectValue `tfsdk:"set"`
}

// BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch struct {
	Interface basetypes.StringValue `tfsdk:"interface"`
	Ipv4      basetypes.ObjectValue `tfsdk:"ipv4"`
	Metric    basetypes.Int64Value  `tfsdk:"metric"`
}

// BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4 represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4 struct {
	Address basetypes.ObjectValue `tfsdk:"address"`
	NextHop basetypes.ObjectValue `tfsdk:"next_hop"`
}

// BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4Address represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4Address struct {
	AccessList basetypes.StringValue `tfsdk:"access_list"`
	PrefixList basetypes.StringValue `tfsdk:"prefix_list"`
}

// BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4NextHop represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4NextHop struct {
	AccessList basetypes.StringValue `tfsdk:"access_list"`
	PrefixList basetypes.StringValue `tfsdk:"prefix_list"`
}

// BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet struct {
	Aggregator       basetypes.ObjectValue `tfsdk:"aggregator"`
	AspathPrepend    basetypes.ListValue   `tfsdk:"aspath_prepend"`
	AtomicAggregate  basetypes.BoolValue   `tfsdk:"atomic_aggregate"`
	Ipv4             basetypes.ObjectValue `tfsdk:"ipv4"`
	LargeCommunity   basetypes.ListValue   `tfsdk:"large_community"`
	LocalPreference  basetypes.Int64Value  `tfsdk:"local_preference"`
	Metric           basetypes.ObjectValue `tfsdk:"metric"`
	Origin           basetypes.StringValue `tfsdk:"origin"`
	OriginatorId     basetypes.StringValue `tfsdk:"originator_id"`
	RegularCommunity basetypes.ListValue   `tfsdk:"regular_community"`
	Tag              basetypes.Int64Value  `tfsdk:"tag"`
	Weight           basetypes.Int64Value  `tfsdk:"weight"`
}

// BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator struct {
	As       basetypes.Int64Value  `tfsdk:"as"`
	RouterId basetypes.StringValue `tfsdk:"router_id"`
}

// BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4 represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4 struct {
	NextHop       basetypes.StringValue `tfsdk:"next_hop"`
	SourceAddress basetypes.StringValue `tfsdk:"source_address"`
}

// BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric struct {
	Action basetypes.StringValue `tfsdk:"action"`
	Value  basetypes.Int64Value  `tfsdk:"value"`
}

// BgpRouteMapRedistributionsConnectedStaticOspf represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticOspf struct {
	RouteMap basetypes.ListValue `tfsdk:"route_map"`
}

// BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner struct {
	Action      basetypes.StringValue `tfsdk:"action"`
	Description basetypes.StringValue `tfsdk:"description"`
	Match       basetypes.ObjectValue `tfsdk:"match"`
	Name        basetypes.Int64Value  `tfsdk:"name"`
	Set         basetypes.ObjectValue `tfsdk:"set"`
}

// BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch struct {
	Interface basetypes.StringValue `tfsdk:"interface"`
	Ipv4      basetypes.ObjectValue `tfsdk:"ipv4"`
	Metric    basetypes.Int64Value  `tfsdk:"metric"`
}

// BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4 represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4 struct {
	Address basetypes.ObjectValue `tfsdk:"address"`
	NextHop basetypes.ObjectValue `tfsdk:"next_hop"`
}

// BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4Address represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4Address struct {
	AccessList basetypes.StringValue `tfsdk:"access_list"`
	PrefixList basetypes.StringValue `tfsdk:"prefix_list"`
}

// BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4NextHop represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4NextHop struct {
	AccessList basetypes.StringValue `tfsdk:"access_list"`
	PrefixList basetypes.StringValue `tfsdk:"prefix_list"`
}

// BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet struct {
	Metric     basetypes.ObjectValue `tfsdk:"metric"`
	MetricType basetypes.StringValue `tfsdk:"metric_type"`
	Tag        basetypes.Int64Value  `tfsdk:"tag"`
}

// BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSetMetric represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSetMetric struct {
	Action basetypes.StringValue `tfsdk:"action"`
	Value  basetypes.Int64Value  `tfsdk:"value"`
}

// BgpRouteMapRedistributionsConnectedStaticRib represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticRib struct {
	RouteMap basetypes.ListValue `tfsdk:"route_map"`
}

// BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner struct {
	Action      basetypes.StringValue `tfsdk:"action"`
	Description basetypes.StringValue `tfsdk:"description"`
	Match       basetypes.ObjectValue `tfsdk:"match"`
	Name        basetypes.Int64Value  `tfsdk:"name"`
	Set         basetypes.ObjectValue `tfsdk:"set"`
}

// BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch struct {
	Interface basetypes.StringValue `tfsdk:"interface"`
	Ipv4      basetypes.ObjectValue `tfsdk:"ipv4"`
	Metric    basetypes.Int64Value  `tfsdk:"metric"`
}

// BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4 represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4 struct {
	Address basetypes.ObjectValue `tfsdk:"address"`
	NextHop basetypes.ObjectValue `tfsdk:"next_hop"`
}

// BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4Address represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4Address struct {
	AccessList basetypes.StringValue `tfsdk:"access_list"`
	PrefixList basetypes.StringValue `tfsdk:"prefix_list"`
}

// BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4NextHop represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4NextHop struct {
	AccessList basetypes.StringValue `tfsdk:"access_list"`
	PrefixList basetypes.StringValue `tfsdk:"prefix_list"`
}

// BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerSet represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerSet struct {
	SourceAddress basetypes.StringValue `tfsdk:"source_address"`
}

// BgpRouteMapRedistributionsOspf represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsOspf struct {
	Bgp basetypes.ObjectValue `tfsdk:"bgp"`
	Rib basetypes.ObjectValue `tfsdk:"rib"`
}

// BgpRouteMapRedistributionsOspfBgp represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsOspfBgp struct {
	RouteMap basetypes.ListValue `tfsdk:"route_map"`
}

// BgpRouteMapRedistributionsOspfBgpRouteMapInner represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsOspfBgpRouteMapInner struct {
	Action      basetypes.StringValue `tfsdk:"action"`
	Description basetypes.StringValue `tfsdk:"description"`
	Match       basetypes.ObjectValue `tfsdk:"match"`
	Name        basetypes.Int64Value  `tfsdk:"name"`
	Set         basetypes.ObjectValue `tfsdk:"set"`
}

// BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch struct {
	Address   basetypes.ObjectValue `tfsdk:"address"`
	Interface basetypes.StringValue `tfsdk:"interface"`
	Metric    basetypes.Int64Value  `tfsdk:"metric"`
	NextHop   basetypes.ObjectValue `tfsdk:"next_hop"`
	Tag       basetypes.Int64Value  `tfsdk:"tag"`
}

// BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress struct {
	AccessList basetypes.StringValue `tfsdk:"access_list"`
	PrefixList basetypes.StringValue `tfsdk:"prefix_list"`
}

// BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop struct {
	AccessList basetypes.StringValue `tfsdk:"access_list"`
	PrefixList basetypes.StringValue `tfsdk:"prefix_list"`
}

// BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet struct {
	Aggregator       basetypes.ObjectValue `tfsdk:"aggregator"`
	AspathPrepend    basetypes.ListValue   `tfsdk:"aspath_prepend"`
	AtomicAggregate  basetypes.BoolValue   `tfsdk:"atomic_aggregate"`
	Ipv4             basetypes.ObjectValue `tfsdk:"ipv4"`
	LargeCommunity   basetypes.ListValue   `tfsdk:"large_community"`
	LocalPreference  basetypes.Int64Value  `tfsdk:"local_preference"`
	Metric           basetypes.ObjectValue `tfsdk:"metric"`
	Origin           basetypes.StringValue `tfsdk:"origin"`
	OriginatorId     basetypes.StringValue `tfsdk:"originator_id"`
	RegularCommunity basetypes.ListValue   `tfsdk:"regular_community"`
	Tag              basetypes.Int64Value  `tfsdk:"tag"`
	Weight           basetypes.Int64Value  `tfsdk:"weight"`
}

// BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator struct {
	As       basetypes.Int64Value  `tfsdk:"as"`
	RouterId basetypes.StringValue `tfsdk:"router_id"`
}

// BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetIpv4 represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetIpv4 struct {
	NextHop       basetypes.StringValue `tfsdk:"next_hop"`
	SourceAddress basetypes.StringValue `tfsdk:"source_address"`
}

// BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetMetric represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetMetric struct {
	Action basetypes.StringValue `tfsdk:"action"`
	Value  basetypes.Int64Value  `tfsdk:"value"`
}

// BgpRouteMapRedistributionsOspfRib represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsOspfRib struct {
	RouteMap basetypes.ListValue `tfsdk:"route_map"`
}

// BgpRouteMapRedistributionsOspfRibRouteMapInner represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsOspfRibRouteMapInner struct {
	Action      basetypes.StringValue `tfsdk:"action"`
	Description basetypes.StringValue `tfsdk:"description"`
	Match       basetypes.ObjectValue `tfsdk:"match"`
	Name        basetypes.Int64Value  `tfsdk:"name"`
	Set         basetypes.ObjectValue `tfsdk:"set"`
}

// BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch struct {
	Address   basetypes.ObjectValue `tfsdk:"address"`
	Interface basetypes.StringValue `tfsdk:"interface"`
	Metric    basetypes.Int64Value  `tfsdk:"metric"`
	NextHop   basetypes.ObjectValue `tfsdk:"next_hop"`
	Tag       basetypes.Int64Value  `tfsdk:"tag"`
}

// BgpRouteMapRedistributionsOspfRibRouteMapInnerMatchAddress represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsOspfRibRouteMapInnerMatchAddress struct {
	AccessList basetypes.StringValue `tfsdk:"access_list"`
	PrefixList basetypes.StringValue `tfsdk:"prefix_list"`
}

// BgpRouteMapRedistributionsOspfRibRouteMapInnerMatchNextHop represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsOspfRibRouteMapInnerMatchNextHop struct {
	AccessList basetypes.StringValue `tfsdk:"access_list"`
	PrefixList basetypes.StringValue `tfsdk:"prefix_list"`
}

// BgpRouteMapRedistributionsOspfRibRouteMapInnerSet represents a nested structure within the BgpRouteMapRedistributions model
type BgpRouteMapRedistributionsOspfRibRouteMapInnerSet struct {
	SourceAddress basetypes.StringValue `tfsdk:"source_address"`
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributions model.
func (o BgpRouteMapRedistributions) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"bgp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ospf": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"route_map": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action":      basetypes.StringType{},
								"description": basetypes.StringType{},
								"match": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"as_path_access_list": basetypes.StringType{},
										"extended_community":  basetypes.StringType{},
										"interface":           basetypes.StringType{},
										"ipv4": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"address": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"access_list": basetypes.StringType{},
														"prefix_list": basetypes.StringType{},
													},
												},
												"next_hop": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"access_list": basetypes.StringType{},
														"prefix_list": basetypes.StringType{},
													},
												},
												"route_source": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"access_list": basetypes.StringType{},
														"prefix_list": basetypes.StringType{},
													},
												},
											},
										},
										"large_community":   basetypes.StringType{},
										"local_preference":  basetypes.Int64Type{},
										"metric":            basetypes.Int64Type{},
										"origin":            basetypes.StringType{},
										"peer":              basetypes.StringType{},
										"regular_community": basetypes.StringType{},
										"tag":               basetypes.Int64Type{},
									},
								},
								"name": basetypes.Int64Type{},
								"set": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"metric": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"action": basetypes.StringType{},
												"value":  basetypes.Int64Type{},
											},
										},
										"metric_type": basetypes.StringType{},
										"tag":         basetypes.Int64Type{},
									},
								},
							},
						}},
					},
				},
				"rib": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"route_map": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action":      basetypes.StringType{},
								"description": basetypes.StringType{},
								"match": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"as_path_access_list": basetypes.StringType{},
										"extended_community":  basetypes.StringType{},
										"interface":           basetypes.StringType{},
										"ipv4": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"address": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"access_list": basetypes.StringType{},
														"prefix_list": basetypes.StringType{},
													},
												},
												"next_hop": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"access_list": basetypes.StringType{},
														"prefix_list": basetypes.StringType{},
													},
												},
												"route_source": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"access_list": basetypes.StringType{},
														"prefix_list": basetypes.StringType{},
													},
												},
											},
										},
										"large_community":   basetypes.StringType{},
										"local_preference":  basetypes.Int64Type{},
										"metric":            basetypes.Int64Type{},
										"origin":            basetypes.StringType{},
										"peer":              basetypes.StringType{},
										"regular_community": basetypes.StringType{},
										"tag":               basetypes.Int64Type{},
									},
								},
								"name": basetypes.Int64Type{},
								"set": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"source_address": basetypes.StringType{},
									},
								},
							},
						}},
					},
				},
			},
		},
		"connected_static": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"bgp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"route_map": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action":      basetypes.StringType{},
								"description": basetypes.StringType{},
								"match": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"interface": basetypes.StringType{},
										"ipv4": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"address": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"access_list": basetypes.StringType{},
														"prefix_list": basetypes.StringType{},
													},
												},
												"next_hop": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"access_list": basetypes.StringType{},
														"prefix_list": basetypes.StringType{},
													},
												},
											},
										},
										"metric": basetypes.Int64Type{},
									},
								},
								"name": basetypes.Int64Type{},
								"set": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"aggregator": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"as":        basetypes.Int64Type{},
												"router_id": basetypes.StringType{},
											},
										},
										"aspath_prepend":   basetypes.ListType{ElemType: basetypes.Int64Type{}},
										"atomic_aggregate": basetypes.BoolType{},
										"ipv4": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"next_hop":       basetypes.StringType{},
												"source_address": basetypes.StringType{},
											},
										},
										"large_community":  basetypes.ListType{ElemType: basetypes.StringType{}},
										"local_preference": basetypes.Int64Type{},
										"metric": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"action": basetypes.StringType{},
												"value":  basetypes.Int64Type{},
											},
										},
										"origin":            basetypes.StringType{},
										"originator_id":     basetypes.StringType{},
										"regular_community": basetypes.ListType{ElemType: basetypes.StringType{}},
										"tag":               basetypes.Int64Type{},
										"weight":            basetypes.Int64Type{},
									},
								},
							},
						}},
					},
				},
				"ospf": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"route_map": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action":      basetypes.StringType{},
								"description": basetypes.StringType{},
								"match": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"interface": basetypes.StringType{},
										"ipv4": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"address": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"access_list": basetypes.StringType{},
														"prefix_list": basetypes.StringType{},
													},
												},
												"next_hop": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"access_list": basetypes.StringType{},
														"prefix_list": basetypes.StringType{},
													},
												},
											},
										},
										"metric": basetypes.Int64Type{},
									},
								},
								"name": basetypes.Int64Type{},
								"set": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"metric": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"action": basetypes.StringType{},
												"value":  basetypes.Int64Type{},
											},
										},
										"metric_type": basetypes.StringType{},
										"tag":         basetypes.Int64Type{},
									},
								},
							},
						}},
					},
				},
				"rib": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"route_map": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action":      basetypes.StringType{},
								"description": basetypes.StringType{},
								"match": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"interface": basetypes.StringType{},
										"ipv4": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"address": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"access_list": basetypes.StringType{},
														"prefix_list": basetypes.StringType{},
													},
												},
												"next_hop": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"access_list": basetypes.StringType{},
														"prefix_list": basetypes.StringType{},
													},
												},
											},
										},
										"metric": basetypes.Int64Type{},
									},
								},
								"name": basetypes.Int64Type{},
								"set": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"source_address": basetypes.StringType{},
									},
								},
							},
						}},
					},
				},
			},
		},
		"description": basetypes.StringType{},
		"device":      basetypes.StringType{},
		"folder":      basetypes.StringType{},
		"id":          basetypes.StringType{},
		"name":        basetypes.StringType{},
		"ospf": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"bgp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"route_map": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action":      basetypes.StringType{},
								"description": basetypes.StringType{},
								"match": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"address": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"access_list": basetypes.StringType{},
												"prefix_list": basetypes.StringType{},
											},
										},
										"interface": basetypes.StringType{},
										"metric":    basetypes.Int64Type{},
										"next_hop": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"access_list": basetypes.StringType{},
												"prefix_list": basetypes.StringType{},
											},
										},
										"tag": basetypes.Int64Type{},
									},
								},
								"name": basetypes.Int64Type{},
								"set": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"aggregator": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"as":        basetypes.Int64Type{},
												"router_id": basetypes.StringType{},
											},
										},
										"aspath_prepend":   basetypes.ListType{ElemType: basetypes.Int64Type{}},
										"atomic_aggregate": basetypes.BoolType{},
										"ipv4": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"next_hop":       basetypes.StringType{},
												"source_address": basetypes.StringType{},
											},
										},
										"large_community":  basetypes.ListType{ElemType: basetypes.StringType{}},
										"local_preference": basetypes.Int64Type{},
										"metric": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"action": basetypes.StringType{},
												"value":  basetypes.Int64Type{},
											},
										},
										"origin":            basetypes.StringType{},
										"originator_id":     basetypes.StringType{},
										"regular_community": basetypes.ListType{ElemType: basetypes.StringType{}},
										"tag":               basetypes.Int64Type{},
										"weight":            basetypes.Int64Type{},
									},
								},
							},
						}},
					},
				},
				"rib": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"route_map": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action":      basetypes.StringType{},
								"description": basetypes.StringType{},
								"match": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"address": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"access_list": basetypes.StringType{},
												"prefix_list": basetypes.StringType{},
											},
										},
										"interface": basetypes.StringType{},
										"metric":    basetypes.Int64Type{},
										"next_hop": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"access_list": basetypes.StringType{},
												"prefix_list": basetypes.StringType{},
											},
										},
										"tag": basetypes.Int64Type{},
									},
								},
								"name": basetypes.Int64Type{},
								"set": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"source_address": basetypes.StringType{},
									},
								},
							},
						}},
					},
				},
			},
		},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributions objects.
func (o BgpRouteMapRedistributions) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsBgp model.
func (o BgpRouteMapRedistributionsBgp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ospf": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"route_map": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action":      basetypes.StringType{},
						"description": basetypes.StringType{},
						"match": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"as_path_access_list": basetypes.StringType{},
								"extended_community":  basetypes.StringType{},
								"interface":           basetypes.StringType{},
								"ipv4": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"address": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"access_list": basetypes.StringType{},
												"prefix_list": basetypes.StringType{},
											},
										},
										"next_hop": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"access_list": basetypes.StringType{},
												"prefix_list": basetypes.StringType{},
											},
										},
										"route_source": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"access_list": basetypes.StringType{},
												"prefix_list": basetypes.StringType{},
											},
										},
									},
								},
								"large_community":   basetypes.StringType{},
								"local_preference":  basetypes.Int64Type{},
								"metric":            basetypes.Int64Type{},
								"origin":            basetypes.StringType{},
								"peer":              basetypes.StringType{},
								"regular_community": basetypes.StringType{},
								"tag":               basetypes.Int64Type{},
							},
						},
						"name": basetypes.Int64Type{},
						"set": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"metric": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"action": basetypes.StringType{},
										"value":  basetypes.Int64Type{},
									},
								},
								"metric_type": basetypes.StringType{},
								"tag":         basetypes.Int64Type{},
							},
						},
					},
				}},
			},
		},
		"rib": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"route_map": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action":      basetypes.StringType{},
						"description": basetypes.StringType{},
						"match": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"as_path_access_list": basetypes.StringType{},
								"extended_community":  basetypes.StringType{},
								"interface":           basetypes.StringType{},
								"ipv4": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"address": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"access_list": basetypes.StringType{},
												"prefix_list": basetypes.StringType{},
											},
										},
										"next_hop": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"access_list": basetypes.StringType{},
												"prefix_list": basetypes.StringType{},
											},
										},
										"route_source": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"access_list": basetypes.StringType{},
												"prefix_list": basetypes.StringType{},
											},
										},
									},
								},
								"large_community":   basetypes.StringType{},
								"local_preference":  basetypes.Int64Type{},
								"metric":            basetypes.Int64Type{},
								"origin":            basetypes.StringType{},
								"peer":              basetypes.StringType{},
								"regular_community": basetypes.StringType{},
								"tag":               basetypes.Int64Type{},
							},
						},
						"name": basetypes.Int64Type{},
						"set": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"source_address": basetypes.StringType{},
							},
						},
					},
				}},
			},
		},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsBgp objects.
func (o BgpRouteMapRedistributionsBgp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsBgpOspf model.
func (o BgpRouteMapRedistributionsBgpOspf) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"route_map": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action":      basetypes.StringType{},
				"description": basetypes.StringType{},
				"match": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"as_path_access_list": basetypes.StringType{},
						"extended_community":  basetypes.StringType{},
						"interface":           basetypes.StringType{},
						"ipv4": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"access_list": basetypes.StringType{},
										"prefix_list": basetypes.StringType{},
									},
								},
								"next_hop": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"access_list": basetypes.StringType{},
										"prefix_list": basetypes.StringType{},
									},
								},
								"route_source": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"access_list": basetypes.StringType{},
										"prefix_list": basetypes.StringType{},
									},
								},
							},
						},
						"large_community":   basetypes.StringType{},
						"local_preference":  basetypes.Int64Type{},
						"metric":            basetypes.Int64Type{},
						"origin":            basetypes.StringType{},
						"peer":              basetypes.StringType{},
						"regular_community": basetypes.StringType{},
						"tag":               basetypes.Int64Type{},
					},
				},
				"name": basetypes.Int64Type{},
				"set": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"metric": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action": basetypes.StringType{},
								"value":  basetypes.Int64Type{},
							},
						},
						"metric_type": basetypes.StringType{},
						"tag":         basetypes.Int64Type{},
					},
				},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsBgpOspf objects.
func (o BgpRouteMapRedistributionsBgpOspf) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsBgpOspfRouteMapInner model.
func (o BgpRouteMapRedistributionsBgpOspfRouteMapInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action":      basetypes.StringType{},
		"description": basetypes.StringType{},
		"match": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"as_path_access_list": basetypes.StringType{},
				"extended_community":  basetypes.StringType{},
				"interface":           basetypes.StringType{},
				"ipv4": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"access_list": basetypes.StringType{},
								"prefix_list": basetypes.StringType{},
							},
						},
						"next_hop": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"access_list": basetypes.StringType{},
								"prefix_list": basetypes.StringType{},
							},
						},
						"route_source": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"access_list": basetypes.StringType{},
								"prefix_list": basetypes.StringType{},
							},
						},
					},
				},
				"large_community":   basetypes.StringType{},
				"local_preference":  basetypes.Int64Type{},
				"metric":            basetypes.Int64Type{},
				"origin":            basetypes.StringType{},
				"peer":              basetypes.StringType{},
				"regular_community": basetypes.StringType{},
				"tag":               basetypes.Int64Type{},
			},
		},
		"name": basetypes.Int64Type{},
		"set": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"metric": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action": basetypes.StringType{},
						"value":  basetypes.Int64Type{},
					},
				},
				"metric_type": basetypes.StringType{},
				"tag":         basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsBgpOspfRouteMapInner objects.
func (o BgpRouteMapRedistributionsBgpOspfRouteMapInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch model.
func (o BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"as_path_access_list": basetypes.StringType{},
		"extended_community":  basetypes.StringType{},
		"interface":           basetypes.StringType{},
		"ipv4": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"access_list": basetypes.StringType{},
						"prefix_list": basetypes.StringType{},
					},
				},
				"next_hop": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"access_list": basetypes.StringType{},
						"prefix_list": basetypes.StringType{},
					},
				},
				"route_source": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"access_list": basetypes.StringType{},
						"prefix_list": basetypes.StringType{},
					},
				},
			},
		},
		"large_community":   basetypes.StringType{},
		"local_preference":  basetypes.Int64Type{},
		"metric":            basetypes.Int64Type{},
		"origin":            basetypes.StringType{},
		"peer":              basetypes.StringType{},
		"regular_community": basetypes.StringType{},
		"tag":               basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch objects.
func (o BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatch) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4 model.
func (o BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"access_list": basetypes.StringType{},
				"prefix_list": basetypes.StringType{},
			},
		},
		"next_hop": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"access_list": basetypes.StringType{},
				"prefix_list": basetypes.StringType{},
			},
		},
		"route_source": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"access_list": basetypes.StringType{},
				"prefix_list": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4 objects.
func (o BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address model.
func (o BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"access_list": basetypes.StringType{},
		"prefix_list": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address objects.
func (o BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4Address) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop model.
func (o BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"access_list": basetypes.StringType{},
		"prefix_list": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop objects.
func (o BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4NextHop) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource model.
func (o BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"access_list": basetypes.StringType{},
		"prefix_list": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource objects.
func (o BgpRouteMapRedistributionsBgpOspfRouteMapInnerMatchIpv4RouteSource) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet model.
func (o BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"metric": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action": basetypes.StringType{},
				"value":  basetypes.Int64Type{},
			},
		},
		"metric_type": basetypes.StringType{},
		"tag":         basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet objects.
func (o BgpRouteMapRedistributionsBgpOspfRouteMapInnerSet) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric model.
func (o BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action": basetypes.StringType{},
		"value":  basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric objects.
func (o BgpRouteMapRedistributionsBgpOspfRouteMapInnerSetMetric) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsBgpRib model.
func (o BgpRouteMapRedistributionsBgpRib) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"route_map": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action":      basetypes.StringType{},
				"description": basetypes.StringType{},
				"match": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"as_path_access_list": basetypes.StringType{},
						"extended_community":  basetypes.StringType{},
						"interface":           basetypes.StringType{},
						"ipv4": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"access_list": basetypes.StringType{},
										"prefix_list": basetypes.StringType{},
									},
								},
								"next_hop": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"access_list": basetypes.StringType{},
										"prefix_list": basetypes.StringType{},
									},
								},
								"route_source": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"access_list": basetypes.StringType{},
										"prefix_list": basetypes.StringType{},
									},
								},
							},
						},
						"large_community":   basetypes.StringType{},
						"local_preference":  basetypes.Int64Type{},
						"metric":            basetypes.Int64Type{},
						"origin":            basetypes.StringType{},
						"peer":              basetypes.StringType{},
						"regular_community": basetypes.StringType{},
						"tag":               basetypes.Int64Type{},
					},
				},
				"name": basetypes.Int64Type{},
				"set": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"source_address": basetypes.StringType{},
					},
				},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsBgpRib objects.
func (o BgpRouteMapRedistributionsBgpRib) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsBgpRibRouteMapInner model.
func (o BgpRouteMapRedistributionsBgpRibRouteMapInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action":      basetypes.StringType{},
		"description": basetypes.StringType{},
		"match": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"as_path_access_list": basetypes.StringType{},
				"extended_community":  basetypes.StringType{},
				"interface":           basetypes.StringType{},
				"ipv4": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"access_list": basetypes.StringType{},
								"prefix_list": basetypes.StringType{},
							},
						},
						"next_hop": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"access_list": basetypes.StringType{},
								"prefix_list": basetypes.StringType{},
							},
						},
						"route_source": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"access_list": basetypes.StringType{},
								"prefix_list": basetypes.StringType{},
							},
						},
					},
				},
				"large_community":   basetypes.StringType{},
				"local_preference":  basetypes.Int64Type{},
				"metric":            basetypes.Int64Type{},
				"origin":            basetypes.StringType{},
				"peer":              basetypes.StringType{},
				"regular_community": basetypes.StringType{},
				"tag":               basetypes.Int64Type{},
			},
		},
		"name": basetypes.Int64Type{},
		"set": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"source_address": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsBgpRibRouteMapInner objects.
func (o BgpRouteMapRedistributionsBgpRibRouteMapInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch model.
func (o BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"as_path_access_list": basetypes.StringType{},
		"extended_community":  basetypes.StringType{},
		"interface":           basetypes.StringType{},
		"ipv4": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"access_list": basetypes.StringType{},
						"prefix_list": basetypes.StringType{},
					},
				},
				"next_hop": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"access_list": basetypes.StringType{},
						"prefix_list": basetypes.StringType{},
					},
				},
				"route_source": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"access_list": basetypes.StringType{},
						"prefix_list": basetypes.StringType{},
					},
				},
			},
		},
		"large_community":   basetypes.StringType{},
		"local_preference":  basetypes.Int64Type{},
		"metric":            basetypes.Int64Type{},
		"origin":            basetypes.StringType{},
		"peer":              basetypes.StringType{},
		"regular_community": basetypes.StringType{},
		"tag":               basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch objects.
func (o BgpRouteMapRedistributionsBgpRibRouteMapInnerMatch) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4 model.
func (o BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"access_list": basetypes.StringType{},
				"prefix_list": basetypes.StringType{},
			},
		},
		"next_hop": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"access_list": basetypes.StringType{},
				"prefix_list": basetypes.StringType{},
			},
		},
		"route_source": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"access_list": basetypes.StringType{},
				"prefix_list": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4 objects.
func (o BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address model.
func (o BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"access_list": basetypes.StringType{},
		"prefix_list": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address objects.
func (o BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4Address) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop model.
func (o BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"access_list": basetypes.StringType{},
		"prefix_list": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop objects.
func (o BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4NextHop) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource model.
func (o BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"access_list": basetypes.StringType{},
		"prefix_list": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource objects.
func (o BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsBgpRibRouteMapInnerSet model.
func (o BgpRouteMapRedistributionsBgpRibRouteMapInnerSet) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"source_address": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsBgpRibRouteMapInnerSet objects.
func (o BgpRouteMapRedistributionsBgpRibRouteMapInnerSet) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStatic model.
func (o BgpRouteMapRedistributionsConnectedStatic) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"bgp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"route_map": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action":      basetypes.StringType{},
						"description": basetypes.StringType{},
						"match": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"interface": basetypes.StringType{},
								"ipv4": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"address": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"access_list": basetypes.StringType{},
												"prefix_list": basetypes.StringType{},
											},
										},
										"next_hop": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"access_list": basetypes.StringType{},
												"prefix_list": basetypes.StringType{},
											},
										},
									},
								},
								"metric": basetypes.Int64Type{},
							},
						},
						"name": basetypes.Int64Type{},
						"set": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"aggregator": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"as":        basetypes.Int64Type{},
										"router_id": basetypes.StringType{},
									},
								},
								"aspath_prepend":   basetypes.ListType{ElemType: basetypes.Int64Type{}},
								"atomic_aggregate": basetypes.BoolType{},
								"ipv4": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"next_hop":       basetypes.StringType{},
										"source_address": basetypes.StringType{},
									},
								},
								"large_community":  basetypes.ListType{ElemType: basetypes.StringType{}},
								"local_preference": basetypes.Int64Type{},
								"metric": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"action": basetypes.StringType{},
										"value":  basetypes.Int64Type{},
									},
								},
								"origin":            basetypes.StringType{},
								"originator_id":     basetypes.StringType{},
								"regular_community": basetypes.ListType{ElemType: basetypes.StringType{}},
								"tag":               basetypes.Int64Type{},
								"weight":            basetypes.Int64Type{},
							},
						},
					},
				}},
			},
		},
		"ospf": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"route_map": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action":      basetypes.StringType{},
						"description": basetypes.StringType{},
						"match": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"interface": basetypes.StringType{},
								"ipv4": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"address": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"access_list": basetypes.StringType{},
												"prefix_list": basetypes.StringType{},
											},
										},
										"next_hop": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"access_list": basetypes.StringType{},
												"prefix_list": basetypes.StringType{},
											},
										},
									},
								},
								"metric": basetypes.Int64Type{},
							},
						},
						"name": basetypes.Int64Type{},
						"set": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"metric": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"action": basetypes.StringType{},
										"value":  basetypes.Int64Type{},
									},
								},
								"metric_type": basetypes.StringType{},
								"tag":         basetypes.Int64Type{},
							},
						},
					},
				}},
			},
		},
		"rib": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"route_map": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action":      basetypes.StringType{},
						"description": basetypes.StringType{},
						"match": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"interface": basetypes.StringType{},
								"ipv4": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"address": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"access_list": basetypes.StringType{},
												"prefix_list": basetypes.StringType{},
											},
										},
										"next_hop": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"access_list": basetypes.StringType{},
												"prefix_list": basetypes.StringType{},
											},
										},
									},
								},
								"metric": basetypes.Int64Type{},
							},
						},
						"name": basetypes.Int64Type{},
						"set": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"source_address": basetypes.StringType{},
							},
						},
					},
				}},
			},
		},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStatic objects.
func (o BgpRouteMapRedistributionsConnectedStatic) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticBgp model.
func (o BgpRouteMapRedistributionsConnectedStaticBgp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"route_map": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action":      basetypes.StringType{},
				"description": basetypes.StringType{},
				"match": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"interface": basetypes.StringType{},
						"ipv4": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"access_list": basetypes.StringType{},
										"prefix_list": basetypes.StringType{},
									},
								},
								"next_hop": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"access_list": basetypes.StringType{},
										"prefix_list": basetypes.StringType{},
									},
								},
							},
						},
						"metric": basetypes.Int64Type{},
					},
				},
				"name": basetypes.Int64Type{},
				"set": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"aggregator": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"as":        basetypes.Int64Type{},
								"router_id": basetypes.StringType{},
							},
						},
						"aspath_prepend":   basetypes.ListType{ElemType: basetypes.Int64Type{}},
						"atomic_aggregate": basetypes.BoolType{},
						"ipv4": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"next_hop":       basetypes.StringType{},
								"source_address": basetypes.StringType{},
							},
						},
						"large_community":  basetypes.ListType{ElemType: basetypes.StringType{}},
						"local_preference": basetypes.Int64Type{},
						"metric": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action": basetypes.StringType{},
								"value":  basetypes.Int64Type{},
							},
						},
						"origin":            basetypes.StringType{},
						"originator_id":     basetypes.StringType{},
						"regular_community": basetypes.ListType{ElemType: basetypes.StringType{}},
						"tag":               basetypes.Int64Type{},
						"weight":            basetypes.Int64Type{},
					},
				},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticBgp objects.
func (o BgpRouteMapRedistributionsConnectedStaticBgp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner model.
func (o BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action":      basetypes.StringType{},
		"description": basetypes.StringType{},
		"match": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"interface": basetypes.StringType{},
				"ipv4": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"access_list": basetypes.StringType{},
								"prefix_list": basetypes.StringType{},
							},
						},
						"next_hop": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"access_list": basetypes.StringType{},
								"prefix_list": basetypes.StringType{},
							},
						},
					},
				},
				"metric": basetypes.Int64Type{},
			},
		},
		"name": basetypes.Int64Type{},
		"set": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"aggregator": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"as":        basetypes.Int64Type{},
						"router_id": basetypes.StringType{},
					},
				},
				"aspath_prepend":   basetypes.ListType{ElemType: basetypes.Int64Type{}},
				"atomic_aggregate": basetypes.BoolType{},
				"ipv4": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"next_hop":       basetypes.StringType{},
						"source_address": basetypes.StringType{},
					},
				},
				"large_community":  basetypes.ListType{ElemType: basetypes.StringType{}},
				"local_preference": basetypes.Int64Type{},
				"metric": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action": basetypes.StringType{},
						"value":  basetypes.Int64Type{},
					},
				},
				"origin":            basetypes.StringType{},
				"originator_id":     basetypes.StringType{},
				"regular_community": basetypes.ListType{ElemType: basetypes.StringType{}},
				"tag":               basetypes.Int64Type{},
				"weight":            basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner objects.
func (o BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch model.
func (o BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"interface": basetypes.StringType{},
		"ipv4": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"access_list": basetypes.StringType{},
						"prefix_list": basetypes.StringType{},
					},
				},
				"next_hop": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"access_list": basetypes.StringType{},
						"prefix_list": basetypes.StringType{},
					},
				},
			},
		},
		"metric": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch objects.
func (o BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatch) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4 model.
func (o BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"access_list": basetypes.StringType{},
				"prefix_list": basetypes.StringType{},
			},
		},
		"next_hop": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"access_list": basetypes.StringType{},
				"prefix_list": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4 objects.
func (o BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4Address model.
func (o BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4Address) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"access_list": basetypes.StringType{},
		"prefix_list": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4Address objects.
func (o BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4Address) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4NextHop model.
func (o BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4NextHop) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"access_list": basetypes.StringType{},
		"prefix_list": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4NextHop objects.
func (o BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerMatchIpv4NextHop) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet model.
func (o BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"aggregator": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"as":        basetypes.Int64Type{},
				"router_id": basetypes.StringType{},
			},
		},
		"aspath_prepend":   basetypes.ListType{ElemType: basetypes.Int64Type{}},
		"atomic_aggregate": basetypes.BoolType{},
		"ipv4": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"next_hop":       basetypes.StringType{},
				"source_address": basetypes.StringType{},
			},
		},
		"large_community":  basetypes.ListType{ElemType: basetypes.StringType{}},
		"local_preference": basetypes.Int64Type{},
		"metric": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action": basetypes.StringType{},
				"value":  basetypes.Int64Type{},
			},
		},
		"origin":            basetypes.StringType{},
		"originator_id":     basetypes.StringType{},
		"regular_community": basetypes.ListType{ElemType: basetypes.StringType{}},
		"tag":               basetypes.Int64Type{},
		"weight":            basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet objects.
func (o BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSet) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator model.
func (o BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"as":        basetypes.Int64Type{},
		"router_id": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator objects.
func (o BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetAggregator) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4 model.
func (o BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"next_hop":       basetypes.StringType{},
		"source_address": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4 objects.
func (o BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric model.
func (o BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action": basetypes.StringType{},
		"value":  basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric objects.
func (o BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticOspf model.
func (o BgpRouteMapRedistributionsConnectedStaticOspf) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"route_map": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action":      basetypes.StringType{},
				"description": basetypes.StringType{},
				"match": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"interface": basetypes.StringType{},
						"ipv4": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"access_list": basetypes.StringType{},
										"prefix_list": basetypes.StringType{},
									},
								},
								"next_hop": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"access_list": basetypes.StringType{},
										"prefix_list": basetypes.StringType{},
									},
								},
							},
						},
						"metric": basetypes.Int64Type{},
					},
				},
				"name": basetypes.Int64Type{},
				"set": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"metric": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action": basetypes.StringType{},
								"value":  basetypes.Int64Type{},
							},
						},
						"metric_type": basetypes.StringType{},
						"tag":         basetypes.Int64Type{},
					},
				},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticOspf objects.
func (o BgpRouteMapRedistributionsConnectedStaticOspf) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner model.
func (o BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action":      basetypes.StringType{},
		"description": basetypes.StringType{},
		"match": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"interface": basetypes.StringType{},
				"ipv4": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"access_list": basetypes.StringType{},
								"prefix_list": basetypes.StringType{},
							},
						},
						"next_hop": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"access_list": basetypes.StringType{},
								"prefix_list": basetypes.StringType{},
							},
						},
					},
				},
				"metric": basetypes.Int64Type{},
			},
		},
		"name": basetypes.Int64Type{},
		"set": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"metric": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action": basetypes.StringType{},
						"value":  basetypes.Int64Type{},
					},
				},
				"metric_type": basetypes.StringType{},
				"tag":         basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner objects.
func (o BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch model.
func (o BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"interface": basetypes.StringType{},
		"ipv4": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"access_list": basetypes.StringType{},
						"prefix_list": basetypes.StringType{},
					},
				},
				"next_hop": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"access_list": basetypes.StringType{},
						"prefix_list": basetypes.StringType{},
					},
				},
			},
		},
		"metric": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch objects.
func (o BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatch) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4 model.
func (o BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"access_list": basetypes.StringType{},
				"prefix_list": basetypes.StringType{},
			},
		},
		"next_hop": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"access_list": basetypes.StringType{},
				"prefix_list": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4 objects.
func (o BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4Address model.
func (o BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4Address) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"access_list": basetypes.StringType{},
		"prefix_list": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4Address objects.
func (o BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4Address) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4NextHop model.
func (o BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4NextHop) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"access_list": basetypes.StringType{},
		"prefix_list": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4NextHop objects.
func (o BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerMatchIpv4NextHop) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet model.
func (o BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"metric": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action": basetypes.StringType{},
				"value":  basetypes.Int64Type{},
			},
		},
		"metric_type": basetypes.StringType{},
		"tag":         basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet objects.
func (o BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSet) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSetMetric model.
func (o BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSetMetric) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action": basetypes.StringType{},
		"value":  basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSetMetric objects.
func (o BgpRouteMapRedistributionsConnectedStaticOspfRouteMapInnerSetMetric) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticRib model.
func (o BgpRouteMapRedistributionsConnectedStaticRib) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"route_map": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action":      basetypes.StringType{},
				"description": basetypes.StringType{},
				"match": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"interface": basetypes.StringType{},
						"ipv4": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"access_list": basetypes.StringType{},
										"prefix_list": basetypes.StringType{},
									},
								},
								"next_hop": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"access_list": basetypes.StringType{},
										"prefix_list": basetypes.StringType{},
									},
								},
							},
						},
						"metric": basetypes.Int64Type{},
					},
				},
				"name": basetypes.Int64Type{},
				"set": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"source_address": basetypes.StringType{},
					},
				},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticRib objects.
func (o BgpRouteMapRedistributionsConnectedStaticRib) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner model.
func (o BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action":      basetypes.StringType{},
		"description": basetypes.StringType{},
		"match": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"interface": basetypes.StringType{},
				"ipv4": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"access_list": basetypes.StringType{},
								"prefix_list": basetypes.StringType{},
							},
						},
						"next_hop": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"access_list": basetypes.StringType{},
								"prefix_list": basetypes.StringType{},
							},
						},
					},
				},
				"metric": basetypes.Int64Type{},
			},
		},
		"name": basetypes.Int64Type{},
		"set": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"source_address": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner objects.
func (o BgpRouteMapRedistributionsConnectedStaticRibRouteMapInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch model.
func (o BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"interface": basetypes.StringType{},
		"ipv4": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"access_list": basetypes.StringType{},
						"prefix_list": basetypes.StringType{},
					},
				},
				"next_hop": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"access_list": basetypes.StringType{},
						"prefix_list": basetypes.StringType{},
					},
				},
			},
		},
		"metric": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch objects.
func (o BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatch) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4 model.
func (o BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"access_list": basetypes.StringType{},
				"prefix_list": basetypes.StringType{},
			},
		},
		"next_hop": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"access_list": basetypes.StringType{},
				"prefix_list": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4 objects.
func (o BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4Address model.
func (o BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4Address) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"access_list": basetypes.StringType{},
		"prefix_list": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4Address objects.
func (o BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4Address) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4NextHop model.
func (o BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4NextHop) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"access_list": basetypes.StringType{},
		"prefix_list": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4NextHop objects.
func (o BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerMatchIpv4NextHop) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerSet model.
func (o BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerSet) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"source_address": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerSet objects.
func (o BgpRouteMapRedistributionsConnectedStaticRibRouteMapInnerSet) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsOspf model.
func (o BgpRouteMapRedistributionsOspf) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"bgp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"route_map": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action":      basetypes.StringType{},
						"description": basetypes.StringType{},
						"match": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"access_list": basetypes.StringType{},
										"prefix_list": basetypes.StringType{},
									},
								},
								"interface": basetypes.StringType{},
								"metric":    basetypes.Int64Type{},
								"next_hop": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"access_list": basetypes.StringType{},
										"prefix_list": basetypes.StringType{},
									},
								},
								"tag": basetypes.Int64Type{},
							},
						},
						"name": basetypes.Int64Type{},
						"set": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"aggregator": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"as":        basetypes.Int64Type{},
										"router_id": basetypes.StringType{},
									},
								},
								"aspath_prepend":   basetypes.ListType{ElemType: basetypes.Int64Type{}},
								"atomic_aggregate": basetypes.BoolType{},
								"ipv4": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"next_hop":       basetypes.StringType{},
										"source_address": basetypes.StringType{},
									},
								},
								"large_community":  basetypes.ListType{ElemType: basetypes.StringType{}},
								"local_preference": basetypes.Int64Type{},
								"metric": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"action": basetypes.StringType{},
										"value":  basetypes.Int64Type{},
									},
								},
								"origin":            basetypes.StringType{},
								"originator_id":     basetypes.StringType{},
								"regular_community": basetypes.ListType{ElemType: basetypes.StringType{}},
								"tag":               basetypes.Int64Type{},
								"weight":            basetypes.Int64Type{},
							},
						},
					},
				}},
			},
		},
		"rib": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"route_map": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action":      basetypes.StringType{},
						"description": basetypes.StringType{},
						"match": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"access_list": basetypes.StringType{},
										"prefix_list": basetypes.StringType{},
									},
								},
								"interface": basetypes.StringType{},
								"metric":    basetypes.Int64Type{},
								"next_hop": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"access_list": basetypes.StringType{},
										"prefix_list": basetypes.StringType{},
									},
								},
								"tag": basetypes.Int64Type{},
							},
						},
						"name": basetypes.Int64Type{},
						"set": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"source_address": basetypes.StringType{},
							},
						},
					},
				}},
			},
		},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsOspf objects.
func (o BgpRouteMapRedistributionsOspf) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsOspfBgp model.
func (o BgpRouteMapRedistributionsOspfBgp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"route_map": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action":      basetypes.StringType{},
				"description": basetypes.StringType{},
				"match": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"access_list": basetypes.StringType{},
								"prefix_list": basetypes.StringType{},
							},
						},
						"interface": basetypes.StringType{},
						"metric":    basetypes.Int64Type{},
						"next_hop": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"access_list": basetypes.StringType{},
								"prefix_list": basetypes.StringType{},
							},
						},
						"tag": basetypes.Int64Type{},
					},
				},
				"name": basetypes.Int64Type{},
				"set": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"aggregator": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"as":        basetypes.Int64Type{},
								"router_id": basetypes.StringType{},
							},
						},
						"aspath_prepend":   basetypes.ListType{ElemType: basetypes.Int64Type{}},
						"atomic_aggregate": basetypes.BoolType{},
						"ipv4": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"next_hop":       basetypes.StringType{},
								"source_address": basetypes.StringType{},
							},
						},
						"large_community":  basetypes.ListType{ElemType: basetypes.StringType{}},
						"local_preference": basetypes.Int64Type{},
						"metric": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action": basetypes.StringType{},
								"value":  basetypes.Int64Type{},
							},
						},
						"origin":            basetypes.StringType{},
						"originator_id":     basetypes.StringType{},
						"regular_community": basetypes.ListType{ElemType: basetypes.StringType{}},
						"tag":               basetypes.Int64Type{},
						"weight":            basetypes.Int64Type{},
					},
				},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsOspfBgp objects.
func (o BgpRouteMapRedistributionsOspfBgp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsOspfBgpRouteMapInner model.
func (o BgpRouteMapRedistributionsOspfBgpRouteMapInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action":      basetypes.StringType{},
		"description": basetypes.StringType{},
		"match": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"access_list": basetypes.StringType{},
						"prefix_list": basetypes.StringType{},
					},
				},
				"interface": basetypes.StringType{},
				"metric":    basetypes.Int64Type{},
				"next_hop": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"access_list": basetypes.StringType{},
						"prefix_list": basetypes.StringType{},
					},
				},
				"tag": basetypes.Int64Type{},
			},
		},
		"name": basetypes.Int64Type{},
		"set": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"aggregator": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"as":        basetypes.Int64Type{},
						"router_id": basetypes.StringType{},
					},
				},
				"aspath_prepend":   basetypes.ListType{ElemType: basetypes.Int64Type{}},
				"atomic_aggregate": basetypes.BoolType{},
				"ipv4": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"next_hop":       basetypes.StringType{},
						"source_address": basetypes.StringType{},
					},
				},
				"large_community":  basetypes.ListType{ElemType: basetypes.StringType{}},
				"local_preference": basetypes.Int64Type{},
				"metric": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action": basetypes.StringType{},
						"value":  basetypes.Int64Type{},
					},
				},
				"origin":            basetypes.StringType{},
				"originator_id":     basetypes.StringType{},
				"regular_community": basetypes.ListType{ElemType: basetypes.StringType{}},
				"tag":               basetypes.Int64Type{},
				"weight":            basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsOspfBgpRouteMapInner objects.
func (o BgpRouteMapRedistributionsOspfBgpRouteMapInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch model.
func (o BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"access_list": basetypes.StringType{},
				"prefix_list": basetypes.StringType{},
			},
		},
		"interface": basetypes.StringType{},
		"metric":    basetypes.Int64Type{},
		"next_hop": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"access_list": basetypes.StringType{},
				"prefix_list": basetypes.StringType{},
			},
		},
		"tag": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch objects.
func (o BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatch) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress model.
func (o BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"access_list": basetypes.StringType{},
		"prefix_list": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress objects.
func (o BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchAddress) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop model.
func (o BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"access_list": basetypes.StringType{},
		"prefix_list": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop objects.
func (o BgpRouteMapRedistributionsOspfBgpRouteMapInnerMatchNextHop) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet model.
func (o BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"aggregator": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"as":        basetypes.Int64Type{},
				"router_id": basetypes.StringType{},
			},
		},
		"aspath_prepend":   basetypes.ListType{ElemType: basetypes.Int64Type{}},
		"atomic_aggregate": basetypes.BoolType{},
		"ipv4": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"next_hop":       basetypes.StringType{},
				"source_address": basetypes.StringType{},
			},
		},
		"large_community":  basetypes.ListType{ElemType: basetypes.StringType{}},
		"local_preference": basetypes.Int64Type{},
		"metric": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action": basetypes.StringType{},
				"value":  basetypes.Int64Type{},
			},
		},
		"origin":            basetypes.StringType{},
		"originator_id":     basetypes.StringType{},
		"regular_community": basetypes.ListType{ElemType: basetypes.StringType{}},
		"tag":               basetypes.Int64Type{},
		"weight":            basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet objects.
func (o BgpRouteMapRedistributionsOspfBgpRouteMapInnerSet) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator model.
func (o BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"as":        basetypes.Int64Type{},
		"router_id": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator objects.
func (o BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetAggregator) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetIpv4 model.
func (o BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetIpv4) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"next_hop":       basetypes.StringType{},
		"source_address": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetIpv4 objects.
func (o BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetIpv4) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetMetric model.
func (o BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetMetric) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action": basetypes.StringType{},
		"value":  basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetMetric objects.
func (o BgpRouteMapRedistributionsOspfBgpRouteMapInnerSetMetric) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsOspfRib model.
func (o BgpRouteMapRedistributionsOspfRib) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"route_map": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action":      basetypes.StringType{},
				"description": basetypes.StringType{},
				"match": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"access_list": basetypes.StringType{},
								"prefix_list": basetypes.StringType{},
							},
						},
						"interface": basetypes.StringType{},
						"metric":    basetypes.Int64Type{},
						"next_hop": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"access_list": basetypes.StringType{},
								"prefix_list": basetypes.StringType{},
							},
						},
						"tag": basetypes.Int64Type{},
					},
				},
				"name": basetypes.Int64Type{},
				"set": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"source_address": basetypes.StringType{},
					},
				},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsOspfRib objects.
func (o BgpRouteMapRedistributionsOspfRib) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsOspfRibRouteMapInner model.
func (o BgpRouteMapRedistributionsOspfRibRouteMapInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action":      basetypes.StringType{},
		"description": basetypes.StringType{},
		"match": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"access_list": basetypes.StringType{},
						"prefix_list": basetypes.StringType{},
					},
				},
				"interface": basetypes.StringType{},
				"metric":    basetypes.Int64Type{},
				"next_hop": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"access_list": basetypes.StringType{},
						"prefix_list": basetypes.StringType{},
					},
				},
				"tag": basetypes.Int64Type{},
			},
		},
		"name": basetypes.Int64Type{},
		"set": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"source_address": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsOspfRibRouteMapInner objects.
func (o BgpRouteMapRedistributionsOspfRibRouteMapInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch model.
func (o BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"access_list": basetypes.StringType{},
				"prefix_list": basetypes.StringType{},
			},
		},
		"interface": basetypes.StringType{},
		"metric":    basetypes.Int64Type{},
		"next_hop": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"access_list": basetypes.StringType{},
				"prefix_list": basetypes.StringType{},
			},
		},
		"tag": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch objects.
func (o BgpRouteMapRedistributionsOspfRibRouteMapInnerMatch) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsOspfRibRouteMapInnerMatchAddress model.
func (o BgpRouteMapRedistributionsOspfRibRouteMapInnerMatchAddress) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"access_list": basetypes.StringType{},
		"prefix_list": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsOspfRibRouteMapInnerMatchAddress objects.
func (o BgpRouteMapRedistributionsOspfRibRouteMapInnerMatchAddress) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsOspfRibRouteMapInnerMatchNextHop model.
func (o BgpRouteMapRedistributionsOspfRibRouteMapInnerMatchNextHop) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"access_list": basetypes.StringType{},
		"prefix_list": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsOspfRibRouteMapInnerMatchNextHop objects.
func (o BgpRouteMapRedistributionsOspfRibRouteMapInnerMatchNextHop) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapRedistributionsOspfRibRouteMapInnerSet model.
func (o BgpRouteMapRedistributionsOspfRibRouteMapInnerSet) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"source_address": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapRedistributionsOspfRibRouteMapInnerSet objects.
func (o BgpRouteMapRedistributionsOspfRibRouteMapInnerSet) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// BgpRouteMapRedistributionsResourceSchema defines the schema for BgpRouteMapRedistributions resource
var BgpRouteMapRedistributionsResourceSchema = schema.Schema{
	MarkdownDescription: "BgpRouteMapRedistribution resource",
	Attributes: map[string]schema.Attribute{
		"bgp": schema.SingleNestedAttribute{
			Validators: []validator.Object{
				objectvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("connected_static"),
					path.MatchRelative().AtParent().AtName("ospf"),
				),
			},
			MarkdownDescription: "Bgp\n\n> ℹ️ **Note:** You must specify exactly one of `bgp`, `connected_static`, and `ospf`.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"ospf": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("rib"),
						),
					},
					MarkdownDescription: "Ospf\n\n> ℹ️ **Note:** You must specify exactly one of `ospf` and `rib`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"route_map": schema.ListNestedAttribute{
							MarkdownDescription: "BGP Root OSPF Route maps",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"action": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("permit", "deny"),
										},
										MarkdownDescription: "BGP Root OSPF Route maps Action",
										Optional:            true,
									},
									"description": schema.StringAttribute{
										MarkdownDescription: "BGP Root OSPF Route maps Description",
										Optional:            true,
									},
									"match": schema.SingleNestedAttribute{
										MarkdownDescription: "Match",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"as_path_access_list": schema.StringAttribute{
												MarkdownDescription: "BGP Root OSPF Route maps match AS path access list",
												Optional:            true,
											},
											"extended_community": schema.StringAttribute{
												MarkdownDescription: "EBGP Root OSPF Route maps match xtended community",
												Optional:            true,
											},
											"interface": schema.StringAttribute{
												MarkdownDescription: "BGP Root OSPF Route maps match Interface",
												Optional:            true,
											},
											"ipv4": schema.SingleNestedAttribute{
												MarkdownDescription: "BGP Root OSPF Route maps match bgp-route-map-redistributions ipv4 object",
												Optional:            true,
												Attributes: map[string]schema.Attribute{
													"address": schema.SingleNestedAttribute{
														MarkdownDescription: "BGP Root OSPF Route maps match bgp-route-map-redistributions ipv4 object address",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"access_list": schema.StringAttribute{
																MarkdownDescription: "BGP Root OSPF Route maps match ipv4 Access list",
																Optional:            true,
															},
															"prefix_list": schema.StringAttribute{
																MarkdownDescription: "BGP Root OSPF Route maps match ipv4 Prefix list",
																Optional:            true,
															},
														},
													},
													"next_hop": schema.SingleNestedAttribute{
														MarkdownDescription: "BGP Root OSPF Route maps match  bgp-route-map-redistributions ipv4 object next_hop",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"access_list": schema.StringAttribute{
																MarkdownDescription: "BGP Root OSPF Route maps ipv4 next_vr hop Access list",
																Optional:            true,
															},
															"prefix_list": schema.StringAttribute{
																MarkdownDescription: "BGP Root OSPF Route maps ipv4 next hop Prefix list",
																Optional:            true,
															},
														},
													},
													"route_source": schema.SingleNestedAttribute{
														MarkdownDescription: "BGP Root OSPF Route maps ipv4 bgp-route-map-redistributions ipv4 object route_source",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"access_list": schema.StringAttribute{
																MarkdownDescription: "BGP Root OSPF Route maps ipv4 route source Access list",
																Optional:            true,
															},
															"prefix_list": schema.StringAttribute{
																MarkdownDescription: "BGP Root OSPF Route maps ipv4 route source Prefix list",
																Optional:            true,
															},
														},
													},
												},
											},
											"large_community": schema.StringAttribute{
												MarkdownDescription: "BGP Root OSPF Route maps match Large community",
												Optional:            true,
											},
											"local_preference": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(1, 4294967295),
												},
												MarkdownDescription: "BGP Root OSPF Route maps match Local preference",
												Optional:            true,
											},
											"metric": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(0, 4294967295),
												},
												MarkdownDescription: "BGP Root OSPF Route maps match Metric",
												Optional:            true,
											},
											"origin": schema.StringAttribute{
												MarkdownDescription: "BGP Root OSPF Route maps match Origin",
												Optional:            true,
											},
											"peer": schema.StringAttribute{
												Validators: []validator.String{
													stringvalidator.OneOf("local", "none"),
												},
												MarkdownDescription: "BGP Root OSPF Route maps match Peer",
												Optional:            true,
											},
											"regular_community": schema.StringAttribute{
												MarkdownDescription: "BGP Root OSPF Route maps match Regular community",
												Optional:            true,
											},
											"tag": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(1, 4294967295),
												},
												MarkdownDescription: "BGP Root OSPF Route maps match Tag",
												Optional:            true,
											},
										},
									},
									"name": schema.Int64Attribute{
										Validators: []validator.Int64{
											int64validator.Between(1, 65535),
										},
										MarkdownDescription: "BGP Root OSPF Route maps Sequence number",
										Optional:            true,
									},
									"set": schema.SingleNestedAttribute{
										MarkdownDescription: "BGP Root OSPF Set",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"metric": schema.SingleNestedAttribute{
												MarkdownDescription: "Metric",
												Optional:            true,
												Attributes: map[string]schema.Attribute{
													"action": schema.StringAttribute{
														Validators: []validator.String{
															stringvalidator.OneOf("set", "add", "subtract"),
														},
														MarkdownDescription: "BGP Root OSPF Route maps set Metric action",
														Optional:            true,
													},
													"value": schema.Int64Attribute{
														Validators: []validator.Int64{
															int64validator.Between(0, 4294967295),
														},
														MarkdownDescription: "BGP Root OSPF Route maps set Metric value",
														Optional:            true,
													},
												},
											},
											"metric_type": schema.StringAttribute{
												Validators: []validator.String{
													stringvalidator.OneOf("type-1", "type-2"),
												},
												MarkdownDescription: "BGP Root OSPF Route maps set Metric type",
												Optional:            true,
											},
											"tag": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(1, 4294967295),
												},
												MarkdownDescription: "BGP Root OSPF Route maps set Tag",
												Optional:            true,
											},
										},
									},
								},
							},
						},
					},
				},
				"rib": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("ospf"),
						),
					},
					MarkdownDescription: "BGP Root RIB\n\n> ℹ️ **Note:** You must specify exactly one of `ospf` and `rib`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"route_map": schema.ListNestedAttribute{
							MarkdownDescription: "BGP Root RIB Route maps",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"action": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("permit", "deny"),
										},
										MarkdownDescription: "BGP Root RIB Route maps Action",
										Optional:            true,
									},
									"description": schema.StringAttribute{
										MarkdownDescription: "BGP Root RIB Route maps Description",
										Optional:            true,
									},
									"match": schema.SingleNestedAttribute{
										MarkdownDescription: "match attribute for BG Rib route map",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"as_path_access_list": schema.StringAttribute{
												MarkdownDescription: "BGP Root RIB Route maps match AS path access list",
												Optional:            true,
											},
											"extended_community": schema.StringAttribute{
												MarkdownDescription: "BGP Root RIB Route maps match Extended community",
												Optional:            true,
											},
											"interface": schema.StringAttribute{
												MarkdownDescription: "BGP Root RIB Route maps match Interface",
												Optional:            true,
											},
											"ipv4": schema.SingleNestedAttribute{
												MarkdownDescription: "BGP Route Map Redistributions Root BGP rib Route Map IPv4",
												Optional:            true,
												Attributes: map[string]schema.Attribute{
													"address": schema.SingleNestedAttribute{
														MarkdownDescription: "bgp-route-map-redistributions ipv4 rib object address",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"access_list": schema.StringAttribute{
																MarkdownDescription: "BGP Root RIB Route maps match ipv Access list",
																Optional:            true,
															},
															"prefix_list": schema.StringAttribute{
																MarkdownDescription: "BGP Root RIB Route maps match ipv Prefix list",
																Optional:            true,
															},
														},
													},
													"next_hop": schema.SingleNestedAttribute{
														MarkdownDescription: "bgp-route-map-redistributions ipv4 rib object next_hop",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"access_list": schema.StringAttribute{
																MarkdownDescription: "BGP Root RIB Route maps match ipv next hop Access list",
																Optional:            true,
															},
															"prefix_list": schema.StringAttribute{
																MarkdownDescription: "BGP Root RIB Route maps match ipv next hop Prefix list",
																Optional:            true,
															},
														},
													},
													"route_source": schema.SingleNestedAttribute{
														MarkdownDescription: "Route source",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"access_list": schema.StringAttribute{
																MarkdownDescription: "BGP Root RIB Route maps match ipv route source Access list",
																Optional:            true,
															},
															"prefix_list": schema.StringAttribute{
																MarkdownDescription: "BGP Root RIB Route maps match ipv route source Prefix list",
																Optional:            true,
															},
														},
													},
												},
											},
											"large_community": schema.StringAttribute{
												MarkdownDescription: "BGP Root RIB Route maps match Large community",
												Optional:            true,
											},
											"local_preference": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(1, 4294967295),
												},
												MarkdownDescription: "BGP Root RIB Route maps match Local preference",
												Optional:            true,
											},
											"metric": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(0, 4294967295),
												},
												MarkdownDescription: "BGP Root RIB Route maps match Metric",
												Optional:            true,
											},
											"origin": schema.StringAttribute{
												MarkdownDescription: "BGP Root RIB Route maps match Origin",
												Optional:            true,
											},
											"peer": schema.StringAttribute{
												Validators: []validator.String{
													stringvalidator.OneOf("local", "none"),
												},
												MarkdownDescription: "BGP Root RIB Route maps match Peer",
												Optional:            true,
											},
											"regular_community": schema.StringAttribute{
												MarkdownDescription: "BGP Root RIB Route maps match Regular community",
												Optional:            true,
											},
											"tag": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(1, 4294967295),
												},
												MarkdownDescription: "BGP Root RIB Route maps match Tag",
												Optional:            true,
											},
										},
									},
									"name": schema.Int64Attribute{
										Validators: []validator.Int64{
											int64validator.Between(1, 65535),
										},
										MarkdownDescription: "BGP Root RIB Route maps Sequence number",
										Optional:            true,
									},
									"set": schema.SingleNestedAttribute{
										MarkdownDescription: "Set attributes for BGP route map",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"source_address": schema.StringAttribute{
												MarkdownDescription: "BGP Root RIB Route maps set Source address",
												Optional:            true,
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
		"connected_static": schema.SingleNestedAttribute{
			Validators: []validator.Object{
				objectvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("bgp"),
					path.MatchRelative().AtParent().AtName("ospf"),
				),
			},
			MarkdownDescription: "Connected static\n\n> ℹ️ **Note:** You must specify exactly one of `bgp`, `connected_static`, and `ospf`.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"bgp": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("ospf"),
							path.MatchRelative().AtParent().AtName("rib"),
						),
					},
					MarkdownDescription: "Connected Static Root BGP\n\n> ℹ️ **Note:** You must specify exactly one of `bgp`, `ospf`, and `rib`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"route_map": schema.ListNestedAttribute{
							MarkdownDescription: "Connected Static BGP Route maps",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"action": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("permit", "deny"),
										},
										MarkdownDescription: "Connected Static BGP Route maps Action",
										Optional:            true,
									},
									"description": schema.StringAttribute{
										MarkdownDescription: "Connected Static BGP Route maps Description",
										Optional:            true,
									},
									"match": schema.SingleNestedAttribute{
										MarkdownDescription: "Match",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"interface": schema.StringAttribute{
												MarkdownDescription: "Connected Static BGP Route maps match Interface",
												Optional:            true,
											},
											"ipv4": schema.SingleNestedAttribute{
												MarkdownDescription: "bgp-route-map-redistributions connected-static ipv4",
												Optional:            true,
												Attributes: map[string]schema.Attribute{
													"address": schema.SingleNestedAttribute{
														MarkdownDescription: "Address",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"access_list": schema.StringAttribute{
																MarkdownDescription: "Connected Static BGP Route maps match ip4 Access list",
																Optional:            true,
															},
															"prefix_list": schema.StringAttribute{
																MarkdownDescription: "Connected Static BGP Route maps match ip4  Prefix list",
																Optional:            true,
															},
														},
													},
													"next_hop": schema.SingleNestedAttribute{
														MarkdownDescription: "Next hop",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"access_list": schema.StringAttribute{
																MarkdownDescription: "Connected Static BGP Route maps match ip4 next hop Access list",
																Optional:            true,
															},
															"prefix_list": schema.StringAttribute{
																MarkdownDescription: "Connected Static BGP Route maps match ip4 next hop Prefix list",
																Optional:            true,
															},
														},
													},
												},
											},
											"metric": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(0, 4294967295),
												},
												MarkdownDescription: "Connected Static BGP Route maps match Metric",
												Optional:            true,
											},
										},
									},
									"name": schema.Int64Attribute{
										Validators: []validator.Int64{
											int64validator.Between(1, 65535),
										},
										MarkdownDescription: "Connected Static BGP Route maps Sequence number",
										Optional:            true,
									},
									"set": schema.SingleNestedAttribute{
										MarkdownDescription: "Set",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"aggregator": schema.SingleNestedAttribute{
												MarkdownDescription: "bgp-route-map-redistributions connected_static aggregator",
												Optional:            true,
												Attributes: map[string]schema.Attribute{
													"as": schema.Int64Attribute{
														Validators: []validator.Int64{
															int64validator.Between(1, 4294967295),
														},
														MarkdownDescription: "Connected Static BGP Route maps set Aggregator AS",
														Optional:            true,
													},
													"router_id": schema.StringAttribute{
														MarkdownDescription: "Connected Static BGP Route maps set Router ID",
														Optional:            true,
													},
												},
											},
											"aspath_prepend": schema.ListAttribute{
												ElementType:         types.Int64Type,
												MarkdownDescription: "Connected Static BGP Route maps set AS numbers",
												Optional:            true,
											},
											"atomic_aggregate": schema.BoolAttribute{
												MarkdownDescription: "Connected Static BGP Route maps set Enable BGP atomic aggregate?",
												Optional:            true,
											},
											"ipv4": schema.SingleNestedAttribute{
												MarkdownDescription: "Ipv4",
												Optional:            true,
												Attributes: map[string]schema.Attribute{
													"next_hop": schema.StringAttribute{
														MarkdownDescription: "Connected Static BGP Route maps set Next ipv4 hop",
														Optional:            true,
													},
													"source_address": schema.StringAttribute{
														MarkdownDescription: "Connected Static BGP Route maps set ipv4 Source address",
														Optional:            true,
													},
												},
											},
											"large_community": schema.ListAttribute{
												ElementType:         types.StringType,
												MarkdownDescription: "Connected Static  BGP Route maps set Large communities",
												Optional:            true,
											},
											"local_preference": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(0, 4294967295),
												},
												MarkdownDescription: "Connected Static BGP Route maps set Local preference",
												Optional:            true,
											},
											"metric": schema.SingleNestedAttribute{
												MarkdownDescription: "Metric",
												Optional:            true,
												Attributes: map[string]schema.Attribute{
													"action": schema.StringAttribute{
														Validators: []validator.String{
															stringvalidator.OneOf("set", "add", "substract"),
														},
														MarkdownDescription: "Connected Static BGP Route maps set Metric action",
														Optional:            true,
													},
													"value": schema.Int64Attribute{
														Validators: []validator.Int64{
															int64validator.Between(0, 4294967295),
														},
														MarkdownDescription: "Connected Static BGP Route maps set Metric value",
														Optional:            true,
													},
												},
											},
											"origin": schema.StringAttribute{
												Validators: []validator.String{
													stringvalidator.OneOf("none", "egp", "igp", "incomplete"),
												},
												MarkdownDescription: "Connected Static BGP Route maps set Origin",
												Optional:            true,
											},
											"originator_id": schema.StringAttribute{
												MarkdownDescription: "Connected Static BGP Route maps set Originator ID",
												Optional:            true,
											},
											"regular_community": schema.ListAttribute{
												ElementType:         types.StringType,
												MarkdownDescription: "Connected Static  BGP Route maps set Regular communities",
												Optional:            true,
											},
											"tag": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(1, 4294967295),
												},
												MarkdownDescription: "Connected Static BGP Route maps set Tag",
												Optional:            true,
											},
											"weight": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(0, 4294967295),
												},
												MarkdownDescription: "Connected Static BGP Route maps set Weight",
												Optional:            true,
											},
										},
									},
								},
							},
						},
					},
				},
				"ospf": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("bgp"),
							path.MatchRelative().AtParent().AtName("rib"),
						),
					},
					MarkdownDescription: "Ospf\n\n> ℹ️ **Note:** You must specify exactly one of `bgp`, `ospf`, and `rib`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"route_map": schema.ListNestedAttribute{
							MarkdownDescription: "Connected Static  BGP OSPF Route maps",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"action": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("permit", "deny"),
										},
										MarkdownDescription: "Connected Static BGP OSPF Route map Action",
										Optional:            true,
									},
									"description": schema.StringAttribute{
										MarkdownDescription: "Connected Static BGP OSPF Route map Description",
										Optional:            true,
									},
									"match": schema.SingleNestedAttribute{
										MarkdownDescription: "Match",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"interface": schema.StringAttribute{
												MarkdownDescription: "Connected Static BGP OSPF Route map Interface",
												Optional:            true,
											},
											"ipv4": schema.SingleNestedAttribute{
												MarkdownDescription: "bgp-route-map-redistributions connected-static match ipv4",
												Optional:            true,
												Attributes: map[string]schema.Attribute{
													"address": schema.SingleNestedAttribute{
														MarkdownDescription: "Connected Static Root OSPF Address",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"access_list": schema.StringAttribute{
																MarkdownDescription: "Connected Static BGP OSPF Route map ipv4 Access list",
																Optional:            true,
															},
															"prefix_list": schema.StringAttribute{
																MarkdownDescription: "Connected Static BGP OSPF Route map ipv4 Prefix list",
																Optional:            true,
															},
														},
													},
													"next_hop": schema.SingleNestedAttribute{
														MarkdownDescription: "Next hop",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"access_list": schema.StringAttribute{
																MarkdownDescription: "Connected Static BGP OSPF Route map ipv4 next hop Access list",
																Optional:            true,
															},
															"prefix_list": schema.StringAttribute{
																MarkdownDescription: "Connected Static BGP OSPF Route map ipv4 next hop Prefix list",
																Optional:            true,
															},
														},
													},
												},
											},
											"metric": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(0, 4294967295),
												},
												MarkdownDescription: "Connected Static BGP OSPF Route map Metric",
												Optional:            true,
											},
										},
									},
									"name": schema.Int64Attribute{
										Validators: []validator.Int64{
											int64validator.Between(1, 65535),
										},
										MarkdownDescription: "Connected Static BGP OSPF Route map Sequence number",
										Optional:            true,
									},
									"set": schema.SingleNestedAttribute{
										MarkdownDescription: "Connected Static Root OSPF Set",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"metric": schema.SingleNestedAttribute{
												MarkdownDescription: "Metric",
												Optional:            true,
												Attributes: map[string]schema.Attribute{
													"action": schema.StringAttribute{
														Validators: []validator.String{
															stringvalidator.OneOf("set", "add", "substract"),
														},
														MarkdownDescription: "Connected Static BGP OSPF Route map set Metric action",
														Optional:            true,
													},
													"value": schema.Int64Attribute{
														Validators: []validator.Int64{
															int64validator.Between(0, 4294967295),
														},
														MarkdownDescription: "Connected Static BGP OSPF Route map set Metric value",
														Optional:            true,
													},
												},
											},
											"metric_type": schema.StringAttribute{
												Validators: []validator.String{
													stringvalidator.OneOf("type-1", "type-2"),
												},
												MarkdownDescription: "Connected Static BGP OSPF Route map set Metric type",
												Optional:            true,
											},
											"tag": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(1, 4294967295),
												},
												MarkdownDescription: "Connected Static BGP OSPF Route map set Tag",
												Optional:            true,
											},
										},
									},
								},
							},
						},
					},
				},
				"rib": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("bgp"),
							path.MatchRelative().AtParent().AtName("ospf"),
						),
					},
					MarkdownDescription: "Rib\n\n> ℹ️ **Note:** You must specify exactly one of `bgp`, `ospf`, and `rib`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"route_map": schema.ListNestedAttribute{
							MarkdownDescription: "Connected Static BGP Rib Route maps",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"action": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("permit", "deny"),
										},
										MarkdownDescription: "Connected Static BGP Rib Route maps Action",
										Optional:            true,
									},
									"description": schema.StringAttribute{
										MarkdownDescription: "Connected Static BGP Rib Route maps Description",
										Optional:            true,
									},
									"match": schema.SingleNestedAttribute{
										MarkdownDescription: "Match",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"interface": schema.StringAttribute{
												MarkdownDescription: "Connected Static BGP Rib Route maps Interface",
												Optional:            true,
											},
											"ipv4": schema.SingleNestedAttribute{
												MarkdownDescription: "Ipv4",
												Optional:            true,
												Attributes: map[string]schema.Attribute{
													"address": schema.SingleNestedAttribute{
														MarkdownDescription: "Connected Static BGP Rib Route maps ipv4 address",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"access_list": schema.StringAttribute{
																MarkdownDescription: "Connected Static BGP Rib Route maps ipv4 Access list",
																Optional:            true,
															},
															"prefix_list": schema.StringAttribute{
																MarkdownDescription: "Connected Static BGP Rib Route maps ipv4 Prefix list",
																Optional:            true,
															},
														},
													},
													"next_hop": schema.SingleNestedAttribute{
														MarkdownDescription: "Next hop",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"access_list": schema.StringAttribute{
																MarkdownDescription: "Connected Static BGP Rib Route maps ipv4 nect hop Access list",
																Optional:            true,
															},
															"prefix_list": schema.StringAttribute{
																MarkdownDescription: "Connected Static BGP Rib Route maps ipv4 next hop Prefix list",
																Optional:            true,
															},
														},
													},
												},
											},
											"metric": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(0, 4294967295),
												},
												MarkdownDescription: "Connected Static BGP Rib Route maps Metric",
												Optional:            true,
											},
										},
									},
									"name": schema.Int64Attribute{
										Validators: []validator.Int64{
											int64validator.Between(1, 65535),
										},
										MarkdownDescription: "Connected Static BGP Rib Route maps Sequence number",
										Optional:            true,
									},
									"set": schema.SingleNestedAttribute{
										MarkdownDescription: "Connected Static Root RIB set",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"source_address": schema.StringAttribute{
												MarkdownDescription: "Connected Static BGP Rib Route Map Distribution Source address",
												Optional:            true,
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
		"description": schema.StringAttribute{
			MarkdownDescription: "BGP Route Map Redistributions Description",
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
			},
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "BGP Route Map Redistributions UUID of the resource",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "BGP Route Map Redistributions Name",
			Required:            true,
		},
		"ospf": schema.SingleNestedAttribute{
			Validators: []validator.Object{
				objectvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("bgp"),
					path.MatchRelative().AtParent().AtName("connected_static"),
				),
			},
			MarkdownDescription: "Ospf\n\n> ℹ️ **Note:** You must specify exactly one of `bgp`, `connected_static`, and `ospf`.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"bgp": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("rib"),
						),
					},
					MarkdownDescription: "OSPF Root BGP\n\n> ℹ️ **Note:** You must specify exactly one of `bgp` and `rib`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"route_map": schema.ListNestedAttribute{
							MarkdownDescription: "OSPF BGP Route maps",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"action": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("permit", "deny"),
										},
										MarkdownDescription: "OSPF BGP Route maps Action",
										Optional:            true,
									},
									"description": schema.StringAttribute{
										MarkdownDescription: "OSPF BGP Route maps Description",
										Optional:            true,
									},
									"match": schema.SingleNestedAttribute{
										MarkdownDescription: "Match",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"address": schema.SingleNestedAttribute{
												MarkdownDescription: "bgp-route-map-redistributions ospf address",
												Optional:            true,
												Attributes: map[string]schema.Attribute{
													"access_list": schema.StringAttribute{
														MarkdownDescription: "OSPF BGP Route maps match Access list",
														Optional:            true,
													},
													"prefix_list": schema.StringAttribute{
														MarkdownDescription: "OSPF BGP Route maps match Prefix list",
														Optional:            true,
													},
												},
											},
											"interface": schema.StringAttribute{
												MarkdownDescription: "OSPF BGP Route maps Interface",
												Optional:            true,
											},
											"metric": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(0, 4294967295),
												},
												MarkdownDescription: "OSPF BGP Route maps Metric",
												Optional:            true,
											},
											"next_hop": schema.SingleNestedAttribute{
												MarkdownDescription: "bgp-route-map-redistributions ospf next_hop",
												Optional:            true,
												Attributes: map[string]schema.Attribute{
													"access_list": schema.StringAttribute{
														MarkdownDescription: "OSPF BGP Route maps next_hop Access list",
														Optional:            true,
													},
													"prefix_list": schema.StringAttribute{
														MarkdownDescription: "OSPF BGP Route maps next_hop Prefix list",
														Optional:            true,
													},
												},
											},
											"tag": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(1, 4294967295),
												},
												MarkdownDescription: "OSPF BGP Route maps Tag",
												Optional:            true,
											},
										},
									},
									"name": schema.Int64Attribute{
										Validators: []validator.Int64{
											int64validator.Between(1, 65535),
										},
										MarkdownDescription: "OSPF BGP Route maps Sequence number",
										Optional:            true,
									},
									"set": schema.SingleNestedAttribute{
										MarkdownDescription: "OSPF Root Set",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"aggregator": schema.SingleNestedAttribute{
												MarkdownDescription: "bgp-route-map-redistributions set aggregator",
												Optional:            true,
												Attributes: map[string]schema.Attribute{
													"as": schema.Int64Attribute{
														Validators: []validator.Int64{
															int64validator.Between(1, 4294967295),
														},
														MarkdownDescription: "OSPF BGP Route maps set Aggregator AS",
														Optional:            true,
													},
													"router_id": schema.StringAttribute{
														MarkdownDescription: "OSPF BGP Route maps set Router ID",
														Optional:            true,
													},
												},
											},
											"aspath_prepend": schema.ListAttribute{
												ElementType:         types.Int64Type,
												MarkdownDescription: "OSPF BGP Route maps set AS numbers",
												Optional:            true,
											},
											"atomic_aggregate": schema.BoolAttribute{
												MarkdownDescription: "OSPF BGP Route maps set Enable BGP atomic aggregate?",
												Optional:            true,
											},
											"ipv4": schema.SingleNestedAttribute{
												MarkdownDescription: "Ipv4",
												Optional:            true,
												Attributes: map[string]schema.Attribute{
													"next_hop": schema.StringAttribute{
														MarkdownDescription: "OSPF BGP Route maps set ipv4 Next hop",
														Optional:            true,
													},
													"source_address": schema.StringAttribute{
														MarkdownDescription: "OSPF BGP Route maps set ipv4 Source address",
														Optional:            true,
													},
												},
											},
											"large_community": schema.ListAttribute{
												ElementType:         types.StringType,
												MarkdownDescription: "OSPF BGP Route maps set Large communities",
												Optional:            true,
											},
											"local_preference": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(0, 4294967295),
												},
												MarkdownDescription: "OSPF BGP Route maps set Local preference",
												Optional:            true,
											},
											"metric": schema.SingleNestedAttribute{
												MarkdownDescription: "Metric",
												Optional:            true,
												Attributes: map[string]schema.Attribute{
													"action": schema.StringAttribute{
														Validators: []validator.String{
															stringvalidator.OneOf("set", "add", "substract"),
														},
														MarkdownDescription: "OSPF BGP Route maps set Metric action",
														Optional:            true,
													},
													"value": schema.Int64Attribute{
														Validators: []validator.Int64{
															int64validator.Between(0, 4294967295),
														},
														MarkdownDescription: "OSPF BGP Route maps set Metric value",
														Optional:            true,
													},
												},
											},
											"origin": schema.StringAttribute{
												Validators: []validator.String{
													stringvalidator.OneOf("none", "egp", "igp", "incomplete"),
												},
												MarkdownDescription: "OSPF BGP Route maps set Origin",
												Optional:            true,
											},
											"originator_id": schema.StringAttribute{
												MarkdownDescription: "OSPF BGP Route maps set Originator ID",
												Optional:            true,
											},
											"regular_community": schema.ListAttribute{
												ElementType:         types.StringType,
												MarkdownDescription: "OSPF BGP Route maps set Regular communities",
												Optional:            true,
											},
											"tag": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(1, 4294967295),
												},
												MarkdownDescription: "OSPF BGP Route maps set Tag",
												Optional:            true,
											},
											"weight": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(0, 4294967295),
												},
												MarkdownDescription: "OSPF BGP Route maps set Weight",
												Optional:            true,
											},
										},
									},
								},
							},
						},
					},
				},
				"rib": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("bgp"),
						),
					},
					MarkdownDescription: "Rib\n\n> ℹ️ **Note:** You must specify exactly one of `bgp` and `rib`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"route_map": schema.ListNestedAttribute{
							MarkdownDescription: "OSPF RIB Route maps set Route maps",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"action": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("permit", "deny"),
										},
										MarkdownDescription: "OSPF RIB Route maps Action",
										Optional:            true,
									},
									"description": schema.StringAttribute{
										MarkdownDescription: "OSPF RIB Route maps Description",
										Optional:            true,
									},
									"match": schema.SingleNestedAttribute{
										MarkdownDescription: "Match",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"address": schema.SingleNestedAttribute{
												MarkdownDescription: "OSPF RIB Route maps address",
												Optional:            true,
												Attributes: map[string]schema.Attribute{
													"access_list": schema.StringAttribute{
														MarkdownDescription: "OSPF RIB Route maps address Access list",
														Optional:            true,
													},
													"prefix_list": schema.StringAttribute{
														MarkdownDescription: "OSPF RIB Route maps address Prefix list",
														Optional:            true,
													},
												},
											},
											"interface": schema.StringAttribute{
												MarkdownDescription: "OSPF RIB Route maps Interface",
												Optional:            true,
											},
											"metric": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(0, 4294967295),
												},
												MarkdownDescription: "OSPF RIB Route maps Metric",
												Optional:            true,
											},
											"next_hop": schema.SingleNestedAttribute{
												MarkdownDescription: "OSPF RIB Route maps next_hop",
												Optional:            true,
												Attributes: map[string]schema.Attribute{
													"access_list": schema.StringAttribute{
														MarkdownDescription: "OSPF RIB Route maps next_hop Access list",
														Optional:            true,
													},
													"prefix_list": schema.StringAttribute{
														MarkdownDescription: "OSPF RIB Route maps next_hop Prefix list",
														Optional:            true,
													},
												},
											},
											"tag": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(1, 4294967295),
												},
												MarkdownDescription: "OSPF RIB Route maps tag",
												Optional:            true,
											},
										},
									},
									"name": schema.Int64Attribute{
										Validators: []validator.Int64{
											int64validator.Between(1, 65535),
										},
										MarkdownDescription: "OSPF RIB Route mapsSequence number",
										Optional:            true,
									},
									"set": schema.SingleNestedAttribute{
										MarkdownDescription: "OSPF RIB Route maps set",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"source_address": schema.StringAttribute{
												MarkdownDescription: "OSPF RIB Route maps set Source address",
												Optional:            true,
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

// BgpRouteMapRedistributionsDataSourceSchema defines the schema for BgpRouteMapRedistributions data source
var BgpRouteMapRedistributionsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "BgpRouteMapRedistribution data source",
	Attributes: map[string]dsschema.Attribute{
		"bgp": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Bgp\n\n> ℹ️ **Note:** You must specify exactly one of `bgp`, `connected_static`, and `ospf`.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"ospf": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Ospf\n\n> ℹ️ **Note:** You must specify exactly one of `ospf` and `rib`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"route_map": dsschema.ListNestedAttribute{
							MarkdownDescription: "BGP Root OSPF Route maps",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"action": dsschema.StringAttribute{
										MarkdownDescription: "BGP Root OSPF Route maps Action",
										Computed:            true,
									},
									"description": dsschema.StringAttribute{
										MarkdownDescription: "BGP Root OSPF Route maps Description",
										Computed:            true,
									},
									"match": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Match",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"as_path_access_list": dsschema.StringAttribute{
												MarkdownDescription: "BGP Root OSPF Route maps match AS path access list",
												Computed:            true,
											},
											"extended_community": dsschema.StringAttribute{
												MarkdownDescription: "EBGP Root OSPF Route maps match xtended community",
												Computed:            true,
											},
											"interface": dsschema.StringAttribute{
												MarkdownDescription: "BGP Root OSPF Route maps match Interface",
												Computed:            true,
											},
											"ipv4": dsschema.SingleNestedAttribute{
												MarkdownDescription: "BGP Root OSPF Route maps match bgp-route-map-redistributions ipv4 object",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"address": dsschema.SingleNestedAttribute{
														MarkdownDescription: "BGP Root OSPF Route maps match bgp-route-map-redistributions ipv4 object address",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"access_list": dsschema.StringAttribute{
																MarkdownDescription: "BGP Root OSPF Route maps match ipv4 Access list",
																Computed:            true,
															},
															"prefix_list": dsschema.StringAttribute{
																MarkdownDescription: "BGP Root OSPF Route maps match ipv4 Prefix list",
																Computed:            true,
															},
														},
													},
													"next_hop": dsschema.SingleNestedAttribute{
														MarkdownDescription: "BGP Root OSPF Route maps match  bgp-route-map-redistributions ipv4 object next_hop",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"access_list": dsschema.StringAttribute{
																MarkdownDescription: "BGP Root OSPF Route maps ipv4 next_vr hop Access list",
																Computed:            true,
															},
															"prefix_list": dsschema.StringAttribute{
																MarkdownDescription: "BGP Root OSPF Route maps ipv4 next hop Prefix list",
																Computed:            true,
															},
														},
													},
													"route_source": dsschema.SingleNestedAttribute{
														MarkdownDescription: "BGP Root OSPF Route maps ipv4 bgp-route-map-redistributions ipv4 object route_source",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"access_list": dsschema.StringAttribute{
																MarkdownDescription: "BGP Root OSPF Route maps ipv4 route source Access list",
																Computed:            true,
															},
															"prefix_list": dsschema.StringAttribute{
																MarkdownDescription: "BGP Root OSPF Route maps ipv4 route source Prefix list",
																Computed:            true,
															},
														},
													},
												},
											},
											"large_community": dsschema.StringAttribute{
												MarkdownDescription: "BGP Root OSPF Route maps match Large community",
												Computed:            true,
											},
											"local_preference": dsschema.Int64Attribute{
												MarkdownDescription: "BGP Root OSPF Route maps match Local preference",
												Computed:            true,
											},
											"metric": dsschema.Int64Attribute{
												MarkdownDescription: "BGP Root OSPF Route maps match Metric",
												Computed:            true,
											},
											"origin": dsschema.StringAttribute{
												MarkdownDescription: "BGP Root OSPF Route maps match Origin",
												Computed:            true,
											},
											"peer": dsschema.StringAttribute{
												MarkdownDescription: "BGP Root OSPF Route maps match Peer",
												Computed:            true,
											},
											"regular_community": dsschema.StringAttribute{
												MarkdownDescription: "BGP Root OSPF Route maps match Regular community",
												Computed:            true,
											},
											"tag": dsschema.Int64Attribute{
												MarkdownDescription: "BGP Root OSPF Route maps match Tag",
												Computed:            true,
											},
										},
									},
									"name": dsschema.Int64Attribute{
										MarkdownDescription: "BGP Root OSPF Route maps Sequence number",
										Computed:            true,
									},
									"set": dsschema.SingleNestedAttribute{
										MarkdownDescription: "BGP Root OSPF Set",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"metric": dsschema.SingleNestedAttribute{
												MarkdownDescription: "Metric",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"action": dsschema.StringAttribute{
														MarkdownDescription: "BGP Root OSPF Route maps set Metric action",
														Computed:            true,
													},
													"value": dsschema.Int64Attribute{
														MarkdownDescription: "BGP Root OSPF Route maps set Metric value",
														Computed:            true,
													},
												},
											},
											"metric_type": dsschema.StringAttribute{
												MarkdownDescription: "BGP Root OSPF Route maps set Metric type",
												Computed:            true,
											},
											"tag": dsschema.Int64Attribute{
												MarkdownDescription: "BGP Root OSPF Route maps set Tag",
												Computed:            true,
											},
										},
									},
								},
							},
						},
					},
				},
				"rib": dsschema.SingleNestedAttribute{
					MarkdownDescription: "BGP Root RIB\n\n> ℹ️ **Note:** You must specify exactly one of `ospf` and `rib`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"route_map": dsschema.ListNestedAttribute{
							MarkdownDescription: "BGP Root RIB Route maps",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"action": dsschema.StringAttribute{
										MarkdownDescription: "BGP Root RIB Route maps Action",
										Computed:            true,
									},
									"description": dsschema.StringAttribute{
										MarkdownDescription: "BGP Root RIB Route maps Description",
										Computed:            true,
									},
									"match": dsschema.SingleNestedAttribute{
										MarkdownDescription: "match attribute for BG Rib route map",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"as_path_access_list": dsschema.StringAttribute{
												MarkdownDescription: "BGP Root RIB Route maps match AS path access list",
												Computed:            true,
											},
											"extended_community": dsschema.StringAttribute{
												MarkdownDescription: "BGP Root RIB Route maps match Extended community",
												Computed:            true,
											},
											"interface": dsschema.StringAttribute{
												MarkdownDescription: "BGP Root RIB Route maps match Interface",
												Computed:            true,
											},
											"ipv4": dsschema.SingleNestedAttribute{
												MarkdownDescription: "BGP Route Map Redistributions Root BGP rib Route Map IPv4",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"address": dsschema.SingleNestedAttribute{
														MarkdownDescription: "bgp-route-map-redistributions ipv4 rib object address",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"access_list": dsschema.StringAttribute{
																MarkdownDescription: "BGP Root RIB Route maps match ipv Access list",
																Computed:            true,
															},
															"prefix_list": dsschema.StringAttribute{
																MarkdownDescription: "BGP Root RIB Route maps match ipv Prefix list",
																Computed:            true,
															},
														},
													},
													"next_hop": dsschema.SingleNestedAttribute{
														MarkdownDescription: "bgp-route-map-redistributions ipv4 rib object next_hop",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"access_list": dsschema.StringAttribute{
																MarkdownDescription: "BGP Root RIB Route maps match ipv next hop Access list",
																Computed:            true,
															},
															"prefix_list": dsschema.StringAttribute{
																MarkdownDescription: "BGP Root RIB Route maps match ipv next hop Prefix list",
																Computed:            true,
															},
														},
													},
													"route_source": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Route source",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"access_list": dsschema.StringAttribute{
																MarkdownDescription: "BGP Root RIB Route maps match ipv route source Access list",
																Computed:            true,
															},
															"prefix_list": dsschema.StringAttribute{
																MarkdownDescription: "BGP Root RIB Route maps match ipv route source Prefix list",
																Computed:            true,
															},
														},
													},
												},
											},
											"large_community": dsschema.StringAttribute{
												MarkdownDescription: "BGP Root RIB Route maps match Large community",
												Computed:            true,
											},
											"local_preference": dsschema.Int64Attribute{
												MarkdownDescription: "BGP Root RIB Route maps match Local preference",
												Computed:            true,
											},
											"metric": dsschema.Int64Attribute{
												MarkdownDescription: "BGP Root RIB Route maps match Metric",
												Computed:            true,
											},
											"origin": dsschema.StringAttribute{
												MarkdownDescription: "BGP Root RIB Route maps match Origin",
												Computed:            true,
											},
											"peer": dsschema.StringAttribute{
												MarkdownDescription: "BGP Root RIB Route maps match Peer",
												Computed:            true,
											},
											"regular_community": dsschema.StringAttribute{
												MarkdownDescription: "BGP Root RIB Route maps match Regular community",
												Computed:            true,
											},
											"tag": dsschema.Int64Attribute{
												MarkdownDescription: "BGP Root RIB Route maps match Tag",
												Computed:            true,
											},
										},
									},
									"name": dsschema.Int64Attribute{
										MarkdownDescription: "BGP Root RIB Route maps Sequence number",
										Computed:            true,
									},
									"set": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Set attributes for BGP route map",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"source_address": dsschema.StringAttribute{
												MarkdownDescription: "BGP Root RIB Route maps set Source address",
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
		},
		"connected_static": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Connected static\n\n> ℹ️ **Note:** You must specify exactly one of `bgp`, `connected_static`, and `ospf`.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"bgp": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Connected Static Root BGP\n\n> ℹ️ **Note:** You must specify exactly one of `bgp`, `ospf`, and `rib`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"route_map": dsschema.ListNestedAttribute{
							MarkdownDescription: "Connected Static BGP Route maps",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"action": dsschema.StringAttribute{
										MarkdownDescription: "Connected Static BGP Route maps Action",
										Computed:            true,
									},
									"description": dsschema.StringAttribute{
										MarkdownDescription: "Connected Static BGP Route maps Description",
										Computed:            true,
									},
									"match": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Match",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"interface": dsschema.StringAttribute{
												MarkdownDescription: "Connected Static BGP Route maps match Interface",
												Computed:            true,
											},
											"ipv4": dsschema.SingleNestedAttribute{
												MarkdownDescription: "bgp-route-map-redistributions connected-static ipv4",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"address": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Address",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"access_list": dsschema.StringAttribute{
																MarkdownDescription: "Connected Static BGP Route maps match ip4 Access list",
																Computed:            true,
															},
															"prefix_list": dsschema.StringAttribute{
																MarkdownDescription: "Connected Static BGP Route maps match ip4  Prefix list",
																Computed:            true,
															},
														},
													},
													"next_hop": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Next hop",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"access_list": dsschema.StringAttribute{
																MarkdownDescription: "Connected Static BGP Route maps match ip4 next hop Access list",
																Computed:            true,
															},
															"prefix_list": dsschema.StringAttribute{
																MarkdownDescription: "Connected Static BGP Route maps match ip4 next hop Prefix list",
																Computed:            true,
															},
														},
													},
												},
											},
											"metric": dsschema.Int64Attribute{
												MarkdownDescription: "Connected Static BGP Route maps match Metric",
												Computed:            true,
											},
										},
									},
									"name": dsschema.Int64Attribute{
										MarkdownDescription: "Connected Static BGP Route maps Sequence number",
										Computed:            true,
									},
									"set": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Set",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"aggregator": dsschema.SingleNestedAttribute{
												MarkdownDescription: "bgp-route-map-redistributions connected_static aggregator",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"as": dsschema.Int64Attribute{
														MarkdownDescription: "Connected Static BGP Route maps set Aggregator AS",
														Computed:            true,
													},
													"router_id": dsschema.StringAttribute{
														MarkdownDescription: "Connected Static BGP Route maps set Router ID",
														Computed:            true,
													},
												},
											},
											"aspath_prepend": dsschema.ListAttribute{
												ElementType:         types.Int64Type,
												MarkdownDescription: "Connected Static BGP Route maps set AS numbers",
												Computed:            true,
											},
											"atomic_aggregate": dsschema.BoolAttribute{
												MarkdownDescription: "Connected Static BGP Route maps set Enable BGP atomic aggregate?",
												Computed:            true,
											},
											"ipv4": dsschema.SingleNestedAttribute{
												MarkdownDescription: "Ipv4",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"next_hop": dsschema.StringAttribute{
														MarkdownDescription: "Connected Static BGP Route maps set Next ipv4 hop",
														Computed:            true,
													},
													"source_address": dsschema.StringAttribute{
														MarkdownDescription: "Connected Static BGP Route maps set ipv4 Source address",
														Computed:            true,
													},
												},
											},
											"large_community": dsschema.ListAttribute{
												ElementType:         types.StringType,
												MarkdownDescription: "Connected Static  BGP Route maps set Large communities",
												Computed:            true,
											},
											"local_preference": dsschema.Int64Attribute{
												MarkdownDescription: "Connected Static BGP Route maps set Local preference",
												Computed:            true,
											},
											"metric": dsschema.SingleNestedAttribute{
												MarkdownDescription: "Metric",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"action": dsschema.StringAttribute{
														MarkdownDescription: "Connected Static BGP Route maps set Metric action",
														Computed:            true,
													},
													"value": dsschema.Int64Attribute{
														MarkdownDescription: "Connected Static BGP Route maps set Metric value",
														Computed:            true,
													},
												},
											},
											"origin": dsschema.StringAttribute{
												MarkdownDescription: "Connected Static BGP Route maps set Origin",
												Computed:            true,
											},
											"originator_id": dsschema.StringAttribute{
												MarkdownDescription: "Connected Static BGP Route maps set Originator ID",
												Computed:            true,
											},
											"regular_community": dsschema.ListAttribute{
												ElementType:         types.StringType,
												MarkdownDescription: "Connected Static  BGP Route maps set Regular communities",
												Computed:            true,
											},
											"tag": dsschema.Int64Attribute{
												MarkdownDescription: "Connected Static BGP Route maps set Tag",
												Computed:            true,
											},
											"weight": dsschema.Int64Attribute{
												MarkdownDescription: "Connected Static BGP Route maps set Weight",
												Computed:            true,
											},
										},
									},
								},
							},
						},
					},
				},
				"ospf": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Ospf\n\n> ℹ️ **Note:** You must specify exactly one of `bgp`, `ospf`, and `rib`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"route_map": dsschema.ListNestedAttribute{
							MarkdownDescription: "Connected Static  BGP OSPF Route maps",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"action": dsschema.StringAttribute{
										MarkdownDescription: "Connected Static BGP OSPF Route map Action",
										Computed:            true,
									},
									"description": dsschema.StringAttribute{
										MarkdownDescription: "Connected Static BGP OSPF Route map Description",
										Computed:            true,
									},
									"match": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Match",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"interface": dsschema.StringAttribute{
												MarkdownDescription: "Connected Static BGP OSPF Route map Interface",
												Computed:            true,
											},
											"ipv4": dsschema.SingleNestedAttribute{
												MarkdownDescription: "bgp-route-map-redistributions connected-static match ipv4",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"address": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Connected Static Root OSPF Address",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"access_list": dsschema.StringAttribute{
																MarkdownDescription: "Connected Static BGP OSPF Route map ipv4 Access list",
																Computed:            true,
															},
															"prefix_list": dsschema.StringAttribute{
																MarkdownDescription: "Connected Static BGP OSPF Route map ipv4 Prefix list",
																Computed:            true,
															},
														},
													},
													"next_hop": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Next hop",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"access_list": dsschema.StringAttribute{
																MarkdownDescription: "Connected Static BGP OSPF Route map ipv4 next hop Access list",
																Computed:            true,
															},
															"prefix_list": dsschema.StringAttribute{
																MarkdownDescription: "Connected Static BGP OSPF Route map ipv4 next hop Prefix list",
																Computed:            true,
															},
														},
													},
												},
											},
											"metric": dsschema.Int64Attribute{
												MarkdownDescription: "Connected Static BGP OSPF Route map Metric",
												Computed:            true,
											},
										},
									},
									"name": dsschema.Int64Attribute{
										MarkdownDescription: "Connected Static BGP OSPF Route map Sequence number",
										Computed:            true,
									},
									"set": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Connected Static Root OSPF Set",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"metric": dsschema.SingleNestedAttribute{
												MarkdownDescription: "Metric",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"action": dsschema.StringAttribute{
														MarkdownDescription: "Connected Static BGP OSPF Route map set Metric action",
														Computed:            true,
													},
													"value": dsschema.Int64Attribute{
														MarkdownDescription: "Connected Static BGP OSPF Route map set Metric value",
														Computed:            true,
													},
												},
											},
											"metric_type": dsschema.StringAttribute{
												MarkdownDescription: "Connected Static BGP OSPF Route map set Metric type",
												Computed:            true,
											},
											"tag": dsschema.Int64Attribute{
												MarkdownDescription: "Connected Static BGP OSPF Route map set Tag",
												Computed:            true,
											},
										},
									},
								},
							},
						},
					},
				},
				"rib": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Rib\n\n> ℹ️ **Note:** You must specify exactly one of `bgp`, `ospf`, and `rib`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"route_map": dsschema.ListNestedAttribute{
							MarkdownDescription: "Connected Static BGP Rib Route maps",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"action": dsschema.StringAttribute{
										MarkdownDescription: "Connected Static BGP Rib Route maps Action",
										Computed:            true,
									},
									"description": dsschema.StringAttribute{
										MarkdownDescription: "Connected Static BGP Rib Route maps Description",
										Computed:            true,
									},
									"match": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Match",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"interface": dsschema.StringAttribute{
												MarkdownDescription: "Connected Static BGP Rib Route maps Interface",
												Computed:            true,
											},
											"ipv4": dsschema.SingleNestedAttribute{
												MarkdownDescription: "Ipv4",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"address": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Connected Static BGP Rib Route maps ipv4 address",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"access_list": dsschema.StringAttribute{
																MarkdownDescription: "Connected Static BGP Rib Route maps ipv4 Access list",
																Computed:            true,
															},
															"prefix_list": dsschema.StringAttribute{
																MarkdownDescription: "Connected Static BGP Rib Route maps ipv4 Prefix list",
																Computed:            true,
															},
														},
													},
													"next_hop": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Next hop",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"access_list": dsschema.StringAttribute{
																MarkdownDescription: "Connected Static BGP Rib Route maps ipv4 nect hop Access list",
																Computed:            true,
															},
															"prefix_list": dsschema.StringAttribute{
																MarkdownDescription: "Connected Static BGP Rib Route maps ipv4 next hop Prefix list",
																Computed:            true,
															},
														},
													},
												},
											},
											"metric": dsschema.Int64Attribute{
												MarkdownDescription: "Connected Static BGP Rib Route maps Metric",
												Computed:            true,
											},
										},
									},
									"name": dsschema.Int64Attribute{
										MarkdownDescription: "Connected Static BGP Rib Route maps Sequence number",
										Computed:            true,
									},
									"set": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Connected Static Root RIB set",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"source_address": dsschema.StringAttribute{
												MarkdownDescription: "Connected Static BGP Rib Route Map Distribution Source address",
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
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "BGP Route Map Redistributions Description",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "BGP Route Map Redistributions UUID of the resource",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "BGP Route Map Redistributions Name",
			Optional:            true,
			Computed:            true,
		},
		"ospf": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Ospf\n\n> ℹ️ **Note:** You must specify exactly one of `bgp`, `connected_static`, and `ospf`.",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"bgp": dsschema.SingleNestedAttribute{
					MarkdownDescription: "OSPF Root BGP\n\n> ℹ️ **Note:** You must specify exactly one of `bgp` and `rib`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"route_map": dsschema.ListNestedAttribute{
							MarkdownDescription: "OSPF BGP Route maps",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"action": dsschema.StringAttribute{
										MarkdownDescription: "OSPF BGP Route maps Action",
										Computed:            true,
									},
									"description": dsschema.StringAttribute{
										MarkdownDescription: "OSPF BGP Route maps Description",
										Computed:            true,
									},
									"match": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Match",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"address": dsschema.SingleNestedAttribute{
												MarkdownDescription: "bgp-route-map-redistributions ospf address",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"access_list": dsschema.StringAttribute{
														MarkdownDescription: "OSPF BGP Route maps match Access list",
														Computed:            true,
													},
													"prefix_list": dsschema.StringAttribute{
														MarkdownDescription: "OSPF BGP Route maps match Prefix list",
														Computed:            true,
													},
												},
											},
											"interface": dsschema.StringAttribute{
												MarkdownDescription: "OSPF BGP Route maps Interface",
												Computed:            true,
											},
											"metric": dsschema.Int64Attribute{
												MarkdownDescription: "OSPF BGP Route maps Metric",
												Computed:            true,
											},
											"next_hop": dsschema.SingleNestedAttribute{
												MarkdownDescription: "bgp-route-map-redistributions ospf next_hop",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"access_list": dsschema.StringAttribute{
														MarkdownDescription: "OSPF BGP Route maps next_hop Access list",
														Computed:            true,
													},
													"prefix_list": dsschema.StringAttribute{
														MarkdownDescription: "OSPF BGP Route maps next_hop Prefix list",
														Computed:            true,
													},
												},
											},
											"tag": dsschema.Int64Attribute{
												MarkdownDescription: "OSPF BGP Route maps Tag",
												Computed:            true,
											},
										},
									},
									"name": dsschema.Int64Attribute{
										MarkdownDescription: "OSPF BGP Route maps Sequence number",
										Computed:            true,
									},
									"set": dsschema.SingleNestedAttribute{
										MarkdownDescription: "OSPF Root Set",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"aggregator": dsschema.SingleNestedAttribute{
												MarkdownDescription: "bgp-route-map-redistributions set aggregator",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"as": dsschema.Int64Attribute{
														MarkdownDescription: "OSPF BGP Route maps set Aggregator AS",
														Computed:            true,
													},
													"router_id": dsschema.StringAttribute{
														MarkdownDescription: "OSPF BGP Route maps set Router ID",
														Computed:            true,
													},
												},
											},
											"aspath_prepend": dsschema.ListAttribute{
												ElementType:         types.Int64Type,
												MarkdownDescription: "OSPF BGP Route maps set AS numbers",
												Computed:            true,
											},
											"atomic_aggregate": dsschema.BoolAttribute{
												MarkdownDescription: "OSPF BGP Route maps set Enable BGP atomic aggregate?",
												Computed:            true,
											},
											"ipv4": dsschema.SingleNestedAttribute{
												MarkdownDescription: "Ipv4",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"next_hop": dsschema.StringAttribute{
														MarkdownDescription: "OSPF BGP Route maps set ipv4 Next hop",
														Computed:            true,
													},
													"source_address": dsschema.StringAttribute{
														MarkdownDescription: "OSPF BGP Route maps set ipv4 Source address",
														Computed:            true,
													},
												},
											},
											"large_community": dsschema.ListAttribute{
												ElementType:         types.StringType,
												MarkdownDescription: "OSPF BGP Route maps set Large communities",
												Computed:            true,
											},
											"local_preference": dsschema.Int64Attribute{
												MarkdownDescription: "OSPF BGP Route maps set Local preference",
												Computed:            true,
											},
											"metric": dsschema.SingleNestedAttribute{
												MarkdownDescription: "Metric",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"action": dsschema.StringAttribute{
														MarkdownDescription: "OSPF BGP Route maps set Metric action",
														Computed:            true,
													},
													"value": dsschema.Int64Attribute{
														MarkdownDescription: "OSPF BGP Route maps set Metric value",
														Computed:            true,
													},
												},
											},
											"origin": dsschema.StringAttribute{
												MarkdownDescription: "OSPF BGP Route maps set Origin",
												Computed:            true,
											},
											"originator_id": dsschema.StringAttribute{
												MarkdownDescription: "OSPF BGP Route maps set Originator ID",
												Computed:            true,
											},
											"regular_community": dsschema.ListAttribute{
												ElementType:         types.StringType,
												MarkdownDescription: "OSPF BGP Route maps set Regular communities",
												Computed:            true,
											},
											"tag": dsschema.Int64Attribute{
												MarkdownDescription: "OSPF BGP Route maps set Tag",
												Computed:            true,
											},
											"weight": dsschema.Int64Attribute{
												MarkdownDescription: "OSPF BGP Route maps set Weight",
												Computed:            true,
											},
										},
									},
								},
							},
						},
					},
				},
				"rib": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Rib\n\n> ℹ️ **Note:** You must specify exactly one of `bgp` and `rib`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"route_map": dsschema.ListNestedAttribute{
							MarkdownDescription: "OSPF RIB Route maps set Route maps",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"action": dsschema.StringAttribute{
										MarkdownDescription: "OSPF RIB Route maps Action",
										Computed:            true,
									},
									"description": dsschema.StringAttribute{
										MarkdownDescription: "OSPF RIB Route maps Description",
										Computed:            true,
									},
									"match": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Match",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"address": dsschema.SingleNestedAttribute{
												MarkdownDescription: "OSPF RIB Route maps address",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"access_list": dsschema.StringAttribute{
														MarkdownDescription: "OSPF RIB Route maps address Access list",
														Computed:            true,
													},
													"prefix_list": dsschema.StringAttribute{
														MarkdownDescription: "OSPF RIB Route maps address Prefix list",
														Computed:            true,
													},
												},
											},
											"interface": dsschema.StringAttribute{
												MarkdownDescription: "OSPF RIB Route maps Interface",
												Computed:            true,
											},
											"metric": dsschema.Int64Attribute{
												MarkdownDescription: "OSPF RIB Route maps Metric",
												Computed:            true,
											},
											"next_hop": dsschema.SingleNestedAttribute{
												MarkdownDescription: "OSPF RIB Route maps next_hop",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"access_list": dsschema.StringAttribute{
														MarkdownDescription: "OSPF RIB Route maps next_hop Access list",
														Computed:            true,
													},
													"prefix_list": dsschema.StringAttribute{
														MarkdownDescription: "OSPF RIB Route maps next_hop Prefix list",
														Computed:            true,
													},
												},
											},
											"tag": dsschema.Int64Attribute{
												MarkdownDescription: "OSPF RIB Route maps tag",
												Computed:            true,
											},
										},
									},
									"name": dsschema.Int64Attribute{
										MarkdownDescription: "OSPF RIB Route mapsSequence number",
										Computed:            true,
									},
									"set": dsschema.SingleNestedAttribute{
										MarkdownDescription: "OSPF RIB Route maps set",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"source_address": dsschema.StringAttribute{
												MarkdownDescription: "OSPF RIB Route maps set Source address",
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
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// BgpRouteMapRedistributionsListModel represents the data model for a list data source.
type BgpRouteMapRedistributionsListModel struct {
	Tfid    types.String                 `tfsdk:"tfid"`
	Data    []BgpRouteMapRedistributions `tfsdk:"data"`
	Limit   types.Int64                  `tfsdk:"limit"`
	Offset  types.Int64                  `tfsdk:"offset"`
	Name    types.String                 `tfsdk:"name"`
	Total   types.Int64                  `tfsdk:"total"`
	Folder  types.String                 `tfsdk:"folder"`
	Snippet types.String                 `tfsdk:"snippet"`
	Device  types.String                 `tfsdk:"device"`
}

// BgpRouteMapRedistributionsListDataSourceSchema defines the schema for a list data source.
var BgpRouteMapRedistributionsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: BgpRouteMapRedistributionsDataSourceSchema.Attributes,
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
