package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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

// BgpRouteMaps represents the Terraform model for BgpRouteMaps
type BgpRouteMaps struct {
	Tfid        types.String          `tfsdk:"tfid"`
	Description basetypes.StringValue `tfsdk:"description"`
	Device      basetypes.StringValue `tfsdk:"device"`
	Folder      basetypes.StringValue `tfsdk:"folder"`
	Id          basetypes.StringValue `tfsdk:"id"`
	Name        basetypes.StringValue `tfsdk:"name"`
	RouteMap    basetypes.ListValue   `tfsdk:"route_map"`
	Snippet     basetypes.StringValue `tfsdk:"snippet"`
}

// BgpRouteMapsRouteMapInner represents a nested structure within the BgpRouteMaps model
type BgpRouteMapsRouteMapInner struct {
	Action      basetypes.StringValue `tfsdk:"action"`
	Description basetypes.StringValue `tfsdk:"description"`
	Match       basetypes.ObjectValue `tfsdk:"match"`
	Name        basetypes.Int64Value  `tfsdk:"name"`
	Set         basetypes.ObjectValue `tfsdk:"set"`
}

// BgpRouteMapsRouteMapInnerMatch represents a nested structure within the BgpRouteMaps model
type BgpRouteMapsRouteMapInnerMatch struct {
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

// BgpRouteMapsRouteMapInnerMatchIpv4 represents a nested structure within the BgpRouteMaps model
type BgpRouteMapsRouteMapInnerMatchIpv4 struct {
	Address     basetypes.ObjectValue `tfsdk:"address"`
	NextHop     basetypes.ObjectValue `tfsdk:"next_hop"`
	RouteSource basetypes.ObjectValue `tfsdk:"route_source"`
}

// BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource represents a nested structure within the BgpRouteMaps model
type BgpRouteMapRedistributionsBgpRibRouteMapInnerMatchIpv4RouteSource struct {
	AccessList basetypes.StringValue `tfsdk:"access_list"`
	PrefixList basetypes.StringValue `tfsdk:"prefix_list"`
}

// BgpRouteMapsRouteMapInnerSet represents a nested structure within the BgpRouteMaps model
type BgpRouteMapsRouteMapInnerSet struct {
	Aggregator                basetypes.ObjectValue `tfsdk:"aggregator"`
	AspathExclude             basetypes.ListValue   `tfsdk:"aspath_exclude"`
	AspathPrepend             basetypes.ListValue   `tfsdk:"aspath_prepend"`
	AtomicAggregate           basetypes.BoolValue   `tfsdk:"atomic_aggregate"`
	Ipv4                      basetypes.ObjectValue `tfsdk:"ipv4"`
	LargeCommunity            basetypes.ListValue   `tfsdk:"large_community"`
	LocalPreference           basetypes.Int64Value  `tfsdk:"local_preference"`
	Metric                    basetypes.ObjectValue `tfsdk:"metric"`
	Origin                    basetypes.StringValue `tfsdk:"origin"`
	OriginatorId              basetypes.StringValue `tfsdk:"originator_id"`
	OverwriteLargeCommunity   basetypes.BoolValue   `tfsdk:"overwrite_large_community"`
	OverwriteRegularCommunity basetypes.BoolValue   `tfsdk:"overwrite_regular_community"`
	RegularCommunity          basetypes.ListValue   `tfsdk:"regular_community"`
	RemoveLargeCommunity      basetypes.StringValue `tfsdk:"remove_large_community"`
	RemoveRegularCommunity    basetypes.StringValue `tfsdk:"remove_regular_community"`
	Tag                       basetypes.Int64Value  `tfsdk:"tag"`
	Weight                    basetypes.Int64Value  `tfsdk:"weight"`
}

// BgpRouteMapsRouteMapInnerSetAggregator represents a nested structure within the BgpRouteMaps model
type BgpRouteMapsRouteMapInnerSetAggregator struct {
	As       basetypes.Int64Value  `tfsdk:"as"`
	RouterId basetypes.StringValue `tfsdk:"router_id"`
}

// BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4 represents a nested structure within the BgpRouteMaps model
type BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetIpv4 struct {
	NextHop       basetypes.StringValue `tfsdk:"next_hop"`
	SourceAddress basetypes.StringValue `tfsdk:"source_address"`
}

// BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric represents a nested structure within the BgpRouteMaps model
type BgpRouteMapRedistributionsConnectedStaticBgpRouteMapInnerSetMetric struct {
	Action basetypes.StringValue `tfsdk:"action"`
	Value  basetypes.Int64Value  `tfsdk:"value"`
}

// AttrTypes defines the attribute types for the BgpRouteMaps model.
func (o BgpRouteMaps) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":        basetypes.StringType{},
		"description": basetypes.StringType{},
		"device":      basetypes.StringType{},
		"folder":      basetypes.StringType{},
		"id":          basetypes.StringType{},
		"name":        basetypes.StringType{},
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
						"aggregator": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"as":        basetypes.Int64Type{},
								"router_id": basetypes.StringType{},
							},
						},
						"aspath_exclude":   basetypes.ListType{ElemType: basetypes.Int64Type{}},
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
						"origin":                      basetypes.StringType{},
						"originator_id":               basetypes.StringType{},
						"overwrite_large_community":   basetypes.BoolType{},
						"overwrite_regular_community": basetypes.BoolType{},
						"regular_community":           basetypes.ListType{ElemType: basetypes.StringType{}},
						"remove_large_community":      basetypes.StringType{},
						"remove_regular_community":    basetypes.StringType{},
						"tag":                         basetypes.Int64Type{},
						"weight":                      basetypes.Int64Type{},
					},
				},
			},
		}},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMaps objects.
func (o BgpRouteMaps) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapsRouteMapInner model.
func (o BgpRouteMapsRouteMapInner) AttrTypes() map[string]attr.Type {
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
				"aggregator": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"as":        basetypes.Int64Type{},
						"router_id": basetypes.StringType{},
					},
				},
				"aspath_exclude":   basetypes.ListType{ElemType: basetypes.Int64Type{}},
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
				"origin":                      basetypes.StringType{},
				"originator_id":               basetypes.StringType{},
				"overwrite_large_community":   basetypes.BoolType{},
				"overwrite_regular_community": basetypes.BoolType{},
				"regular_community":           basetypes.ListType{ElemType: basetypes.StringType{}},
				"remove_large_community":      basetypes.StringType{},
				"remove_regular_community":    basetypes.StringType{},
				"tag":                         basetypes.Int64Type{},
				"weight":                      basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapsRouteMapInner objects.
func (o BgpRouteMapsRouteMapInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapsRouteMapInnerMatch model.
func (o BgpRouteMapsRouteMapInnerMatch) AttrTypes() map[string]attr.Type {
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

// AttrType returns the attribute type for a list of BgpRouteMapsRouteMapInnerMatch objects.
func (o BgpRouteMapsRouteMapInnerMatch) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapsRouteMapInnerMatchIpv4 model.
func (o BgpRouteMapsRouteMapInnerMatchIpv4) AttrTypes() map[string]attr.Type {
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

// AttrType returns the attribute type for a list of BgpRouteMapsRouteMapInnerMatchIpv4 objects.
func (o BgpRouteMapsRouteMapInnerMatchIpv4) AttrType() attr.Type {
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

// AttrTypes defines the attribute types for the BgpRouteMapsRouteMapInnerSet model.
func (o BgpRouteMapsRouteMapInnerSet) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"aggregator": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"as":        basetypes.Int64Type{},
				"router_id": basetypes.StringType{},
			},
		},
		"aspath_exclude":   basetypes.ListType{ElemType: basetypes.Int64Type{}},
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
		"origin":                      basetypes.StringType{},
		"originator_id":               basetypes.StringType{},
		"overwrite_large_community":   basetypes.BoolType{},
		"overwrite_regular_community": basetypes.BoolType{},
		"regular_community":           basetypes.ListType{ElemType: basetypes.StringType{}},
		"remove_large_community":      basetypes.StringType{},
		"remove_regular_community":    basetypes.StringType{},
		"tag":                         basetypes.Int64Type{},
		"weight":                      basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapsRouteMapInnerSet objects.
func (o BgpRouteMapsRouteMapInnerSet) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpRouteMapsRouteMapInnerSetAggregator model.
func (o BgpRouteMapsRouteMapInnerSetAggregator) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"as":        basetypes.Int64Type{},
		"router_id": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpRouteMapsRouteMapInnerSetAggregator objects.
func (o BgpRouteMapsRouteMapInnerSetAggregator) AttrType() attr.Type {
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

// BgpRouteMapsResourceSchema defines the schema for BgpRouteMaps resource
var BgpRouteMapsResourceSchema = schema.Schema{
	MarkdownDescription: "BgpRouteMap resource",
	Attributes: map[string]schema.Attribute{
		"description": schema.StringAttribute{
			MarkdownDescription: "Description",
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
		"name": schema.StringAttribute{
			MarkdownDescription: "Name",
			Required:            true,
		},
		"route_map": schema.ListNestedAttribute{
			MarkdownDescription: "Route map",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"action": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("permit", "deny"),
						},
						MarkdownDescription: "Action",
						Optional:            true,
					},
					"description": schema.StringAttribute{
						MarkdownDescription: "Description",
						Optional:            true,
					},
					"match": schema.SingleNestedAttribute{
						MarkdownDescription: "Match",
						Optional:            true,
						Attributes: map[string]schema.Attribute{
							"as_path_access_list": schema.StringAttribute{
								MarkdownDescription: "AS path access list",
								Optional:            true,
							},
							"extended_community": schema.StringAttribute{
								MarkdownDescription: "Extended community",
								Optional:            true,
							},
							"interface": schema.StringAttribute{
								MarkdownDescription: "Interface",
								Optional:            true,
							},
							"ipv4": schema.SingleNestedAttribute{
								MarkdownDescription: "bgp-route-maps ipv4 object",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"address": schema.SingleNestedAttribute{
										MarkdownDescription: "Address",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"access_list": schema.StringAttribute{
												MarkdownDescription: "Access list",
												Optional:            true,
											},
											"prefix_list": schema.StringAttribute{
												MarkdownDescription: "Prefix list",
												Optional:            true,
											},
										},
									},
									"next_hop": schema.SingleNestedAttribute{
										MarkdownDescription: "Next hop",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"access_list": schema.StringAttribute{
												MarkdownDescription: "Access list",
												Optional:            true,
											},
											"prefix_list": schema.StringAttribute{
												MarkdownDescription: "Prefix list",
												Optional:            true,
											},
										},
									},
									"route_source": schema.SingleNestedAttribute{
										MarkdownDescription: "Route source",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"access_list": schema.StringAttribute{
												MarkdownDescription: "Access list",
												Optional:            true,
											},
											"prefix_list": schema.StringAttribute{
												MarkdownDescription: "Prefix list",
												Optional:            true,
											},
										},
									},
								},
							},
							"large_community": schema.StringAttribute{
								MarkdownDescription: "Large community",
								Optional:            true,
							},
							"local_preference": schema.Int64Attribute{
								Validators: []validator.Int64{
									int64validator.Between(0, 4294967295),
								},
								MarkdownDescription: "Local preference",
								Optional:            true,
							},
							"metric": schema.Int64Attribute{
								Validators: []validator.Int64{
									int64validator.Between(0, 4294967295),
								},
								MarkdownDescription: "Metric",
								Optional:            true,
							},
							"origin": schema.StringAttribute{
								MarkdownDescription: "Origin",
								Optional:            true,
							},
							"peer": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("local", "none"),
								},
								MarkdownDescription: "Peer",
								Optional:            true,
							},
							"regular_community": schema.StringAttribute{
								MarkdownDescription: "Regular community",
								Optional:            true,
							},
							"tag": schema.Int64Attribute{
								Validators: []validator.Int64{
									int64validator.Between(1, 4294967295),
								},
								MarkdownDescription: "Tag",
								Optional:            true,
							},
						},
					},
					"name": schema.Int64Attribute{
						Validators: []validator.Int64{
							int64validator.Between(1, 65535),
						},
						MarkdownDescription: "Sequence number",
						Optional:            true,
					},
					"set": schema.SingleNestedAttribute{
						MarkdownDescription: "Set",
						Optional:            true,
						Attributes: map[string]schema.Attribute{
							"aggregator": schema.SingleNestedAttribute{
								MarkdownDescription: "bgp-route-maps aggregator",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"as": schema.Int64Attribute{
										Validators: []validator.Int64{
											int64validator.Between(1, 4294967295),
										},
										MarkdownDescription: "Aggregator AS",
										Optional:            true,
									},
									"router_id": schema.StringAttribute{
										MarkdownDescription: "Router ID",
										Optional:            true,
									},
								},
							},
							"aspath_exclude": schema.ListAttribute{
								ElementType:         types.Int64Type,
								MarkdownDescription: "Aspath exclude",
								Optional:            true,
							},
							"aspath_prepend": schema.ListAttribute{
								ElementType:         types.Int64Type,
								MarkdownDescription: "Aspath prepend",
								Optional:            true,
							},
							"atomic_aggregate": schema.BoolAttribute{
								MarkdownDescription: "Enable BGP atomic aggregate?",
								Optional:            true,
							},
							"ipv4": schema.SingleNestedAttribute{
								MarkdownDescription: "Ipv4",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"next_hop": schema.StringAttribute{
										MarkdownDescription: "Next hop",
										Optional:            true,
									},
									"source_address": schema.StringAttribute{
										MarkdownDescription: "Source address",
										Optional:            true,
									},
								},
							},
							"large_community": schema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "Large community",
								Optional:            true,
							},
							"local_preference": schema.Int64Attribute{
								Validators: []validator.Int64{
									int64validator.Between(0, 4294967295),
								},
								MarkdownDescription: "Local preference",
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
										MarkdownDescription: "Metric action",
										Optional:            true,
									},
									"value": schema.Int64Attribute{
										Validators: []validator.Int64{
											int64validator.Between(0, 4294967295),
										},
										MarkdownDescription: "Metric value",
										Optional:            true,
									},
								},
							},
							"origin": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("none", "egp", "igp", "incomplete"),
								},
								MarkdownDescription: "Origin",
								Optional:            true,
							},
							"originator_id": schema.StringAttribute{
								MarkdownDescription: "Originator ID",
								Optional:            true,
							},
							"overwrite_large_community": schema.BoolAttribute{
								MarkdownDescription: "Overwrite large community?",
								Optional:            true,
							},
							"overwrite_regular_community": schema.BoolAttribute{
								MarkdownDescription: "Overwrite regular community?",
								Optional:            true,
							},
							"regular_community": schema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "Regular community",
								Optional:            true,
							},
							"remove_large_community": schema.StringAttribute{
								MarkdownDescription: "Remove large community name",
								Optional:            true,
							},
							"remove_regular_community": schema.StringAttribute{
								MarkdownDescription: "Remove regular community name",
								Optional:            true,
							},
							"tag": schema.Int64Attribute{
								Validators: []validator.Int64{
									int64validator.Between(0, 4294967295),
								},
								MarkdownDescription: "Tag",
								Optional:            true,
							},
							"weight": schema.Int64Attribute{
								Validators: []validator.Int64{
									int64validator.Between(0, 4294967295),
								},
								MarkdownDescription: "Weight",
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

// BgpRouteMapsDataSourceSchema defines the schema for BgpRouteMaps data source
var BgpRouteMapsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "BgpRouteMap data source",
	Attributes: map[string]dsschema.Attribute{
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Description",
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
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Name",
			Optional:            true,
			Computed:            true,
		},
		"route_map": dsschema.ListNestedAttribute{
			MarkdownDescription: "Route map",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"action": dsschema.StringAttribute{
						MarkdownDescription: "Action",
						Computed:            true,
					},
					"description": dsschema.StringAttribute{
						MarkdownDescription: "Description",
						Computed:            true,
					},
					"match": dsschema.SingleNestedAttribute{
						MarkdownDescription: "Match",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"as_path_access_list": dsschema.StringAttribute{
								MarkdownDescription: "AS path access list",
								Computed:            true,
							},
							"extended_community": dsschema.StringAttribute{
								MarkdownDescription: "Extended community",
								Computed:            true,
							},
							"interface": dsschema.StringAttribute{
								MarkdownDescription: "Interface",
								Computed:            true,
							},
							"ipv4": dsschema.SingleNestedAttribute{
								MarkdownDescription: "bgp-route-maps ipv4 object",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"address": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Address",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"access_list": dsschema.StringAttribute{
												MarkdownDescription: "Access list",
												Computed:            true,
											},
											"prefix_list": dsschema.StringAttribute{
												MarkdownDescription: "Prefix list",
												Computed:            true,
											},
										},
									},
									"next_hop": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Next hop",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"access_list": dsschema.StringAttribute{
												MarkdownDescription: "Access list",
												Computed:            true,
											},
											"prefix_list": dsschema.StringAttribute{
												MarkdownDescription: "Prefix list",
												Computed:            true,
											},
										},
									},
									"route_source": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Route source",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"access_list": dsschema.StringAttribute{
												MarkdownDescription: "Access list",
												Computed:            true,
											},
											"prefix_list": dsschema.StringAttribute{
												MarkdownDescription: "Prefix list",
												Computed:            true,
											},
										},
									},
								},
							},
							"large_community": dsschema.StringAttribute{
								MarkdownDescription: "Large community",
								Computed:            true,
							},
							"local_preference": dsschema.Int64Attribute{
								MarkdownDescription: "Local preference",
								Computed:            true,
							},
							"metric": dsschema.Int64Attribute{
								MarkdownDescription: "Metric",
								Computed:            true,
							},
							"origin": dsschema.StringAttribute{
								MarkdownDescription: "Origin",
								Computed:            true,
							},
							"peer": dsschema.StringAttribute{
								MarkdownDescription: "Peer",
								Computed:            true,
							},
							"regular_community": dsschema.StringAttribute{
								MarkdownDescription: "Regular community",
								Computed:            true,
							},
							"tag": dsschema.Int64Attribute{
								MarkdownDescription: "Tag",
								Computed:            true,
							},
						},
					},
					"name": dsschema.Int64Attribute{
						MarkdownDescription: "Sequence number",
						Computed:            true,
					},
					"set": dsschema.SingleNestedAttribute{
						MarkdownDescription: "Set",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"aggregator": dsschema.SingleNestedAttribute{
								MarkdownDescription: "bgp-route-maps aggregator",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"as": dsschema.Int64Attribute{
										MarkdownDescription: "Aggregator AS",
										Computed:            true,
									},
									"router_id": dsschema.StringAttribute{
										MarkdownDescription: "Router ID",
										Computed:            true,
									},
								},
							},
							"aspath_exclude": dsschema.ListAttribute{
								ElementType:         types.Int64Type,
								MarkdownDescription: "Aspath exclude",
								Computed:            true,
							},
							"aspath_prepend": dsschema.ListAttribute{
								ElementType:         types.Int64Type,
								MarkdownDescription: "Aspath prepend",
								Computed:            true,
							},
							"atomic_aggregate": dsschema.BoolAttribute{
								MarkdownDescription: "Enable BGP atomic aggregate?",
								Computed:            true,
							},
							"ipv4": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Ipv4",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"next_hop": dsschema.StringAttribute{
										MarkdownDescription: "Next hop",
										Computed:            true,
									},
									"source_address": dsschema.StringAttribute{
										MarkdownDescription: "Source address",
										Computed:            true,
									},
								},
							},
							"large_community": dsschema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "Large community",
								Computed:            true,
							},
							"local_preference": dsschema.Int64Attribute{
								MarkdownDescription: "Local preference",
								Computed:            true,
							},
							"metric": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Metric",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"action": dsschema.StringAttribute{
										MarkdownDescription: "Metric action",
										Computed:            true,
									},
									"value": dsschema.Int64Attribute{
										MarkdownDescription: "Metric value",
										Computed:            true,
									},
								},
							},
							"origin": dsschema.StringAttribute{
								MarkdownDescription: "Origin",
								Computed:            true,
							},
							"originator_id": dsschema.StringAttribute{
								MarkdownDescription: "Originator ID",
								Computed:            true,
							},
							"overwrite_large_community": dsschema.BoolAttribute{
								MarkdownDescription: "Overwrite large community?",
								Computed:            true,
							},
							"overwrite_regular_community": dsschema.BoolAttribute{
								MarkdownDescription: "Overwrite regular community?",
								Computed:            true,
							},
							"regular_community": dsschema.ListAttribute{
								ElementType:         types.StringType,
								MarkdownDescription: "Regular community",
								Computed:            true,
							},
							"remove_large_community": dsschema.StringAttribute{
								MarkdownDescription: "Remove large community name",
								Computed:            true,
							},
							"remove_regular_community": dsschema.StringAttribute{
								MarkdownDescription: "Remove regular community name",
								Computed:            true,
							},
							"tag": dsschema.Int64Attribute{
								MarkdownDescription: "Tag",
								Computed:            true,
							},
							"weight": dsschema.Int64Attribute{
								MarkdownDescription: "Weight",
								Computed:            true,
							},
						},
					},
				},
			},
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

// BgpRouteMapsListModel represents the data model for a list data source.
type BgpRouteMapsListModel struct {
	Tfid    types.String   `tfsdk:"tfid"`
	Data    []BgpRouteMaps `tfsdk:"data"`
	Limit   types.Int64    `tfsdk:"limit"`
	Offset  types.Int64    `tfsdk:"offset"`
	Name    types.String   `tfsdk:"name"`
	Total   types.Int64    `tfsdk:"total"`
	Folder  types.String   `tfsdk:"folder"`
	Snippet types.String   `tfsdk:"snippet"`
	Device  types.String   `tfsdk:"device"`
}

// BgpRouteMapsListDataSourceSchema defines the schema for a list data source.
var BgpRouteMapsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: BgpRouteMapsDataSourceSchema.Attributes,
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
