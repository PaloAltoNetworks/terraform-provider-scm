package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
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

// BgpFilteringProfiles represents the Terraform model for BgpFilteringProfiles
type BgpFilteringProfiles struct {
	Tfid        types.String          `tfsdk:"tfid"`
	Description basetypes.StringValue `tfsdk:"description"`
	Device      basetypes.StringValue `tfsdk:"device"`
	Folder      basetypes.StringValue `tfsdk:"folder"`
	Id          basetypes.StringValue `tfsdk:"id"`
	Ipv4        basetypes.ObjectValue `tfsdk:"ipv4"`
	Name        basetypes.StringValue `tfsdk:"name"`
	Snippet     basetypes.StringValue `tfsdk:"snippet"`
}

// BgpFilteringProfilesIpv4 represents a nested structure within the BgpFilteringProfiles model
type BgpFilteringProfilesIpv4 struct {
	Ipv4 basetypes.ObjectValue `tfsdk:"ipv4"`
}

// BgpFilteringProfilesIpv4Ipv4 represents a nested structure within the BgpFilteringProfiles model
type BgpFilteringProfilesIpv4Ipv4 struct {
	Multicast basetypes.ObjectValue `tfsdk:"multicast"`
	Unicast   basetypes.ObjectValue `tfsdk:"unicast"`
}

// BgpFilteringProfilesIpv4Ipv4Multicast represents a nested structure within the BgpFilteringProfiles model
type BgpFilteringProfilesIpv4Ipv4Multicast struct {
	ConditionalAdvertisement basetypes.ObjectValue `tfsdk:"conditional_advertisement"`
	FilterList               basetypes.ObjectValue `tfsdk:"filter_list"`
	InboundNetworkFilters    basetypes.ObjectValue `tfsdk:"inbound_network_filters"`
	Inherit                  basetypes.BoolValue   `tfsdk:"inherit"`
	OutboundNetworkFilters   basetypes.ObjectValue `tfsdk:"outbound_network_filters"`
	RouteMaps                basetypes.ObjectValue `tfsdk:"route_maps"`
	UnsuppressMap            basetypes.StringValue `tfsdk:"unsuppress_map"`
}

// BgpFilterConditionalAdvertisement represents a nested structure within the BgpFilteringProfiles model
type BgpFilterConditionalAdvertisement struct {
	Exist    basetypes.ObjectValue `tfsdk:"exist"`
	NonExist basetypes.ObjectValue `tfsdk:"non_exist"`
}

// BgpFilterConditionalAdvertisementExist represents a nested structure within the BgpFilteringProfiles model
type BgpFilterConditionalAdvertisementExist struct {
	AdvertiseMap basetypes.StringValue `tfsdk:"advertise_map"`
	ExistMap     basetypes.StringValue `tfsdk:"exist_map"`
}

// BgpFilterConditionalAdvertisementNonExist represents a nested structure within the BgpFilteringProfiles model
type BgpFilterConditionalAdvertisementNonExist struct {
	AdvertiseMap basetypes.StringValue `tfsdk:"advertise_map"`
	NonExistMap  basetypes.StringValue `tfsdk:"non_exist_map"`
}

// BgpFilterFilterList represents a nested structure within the BgpFilteringProfiles model
type BgpFilterFilterList struct {
	Inbound  basetypes.StringValue `tfsdk:"inbound"`
	Outbound basetypes.StringValue `tfsdk:"outbound"`
}

// BgpFilterInboundNetworkFilters represents a nested structure within the BgpFilteringProfiles model
type BgpFilterInboundNetworkFilters struct {
	DistributeList basetypes.StringValue `tfsdk:"distribute_list"`
	PrefixList     basetypes.StringValue `tfsdk:"prefix_list"`
}

// BgpFilter represents a nested structure within the BgpFilteringProfiles model
type BgpFilter struct {
	ConditionalAdvertisement basetypes.ObjectValue `tfsdk:"conditional_advertisement"`
	FilterList               basetypes.ObjectValue `tfsdk:"filter_list"`
	InboundNetworkFilters    basetypes.ObjectValue `tfsdk:"inbound_network_filters"`
	OutboundNetworkFilters   basetypes.ObjectValue `tfsdk:"outbound_network_filters"`
	RouteMaps                basetypes.ObjectValue `tfsdk:"route_maps"`
	UnsuppressMap            basetypes.StringValue `tfsdk:"unsuppress_map"`
}

// AttrTypes defines the attribute types for the BgpFilteringProfiles model.
func (o BgpFilteringProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":        basetypes.StringType{},
		"description": basetypes.StringType{},
		"device":      basetypes.StringType{},
		"folder":      basetypes.StringType{},
		"id":          basetypes.StringType{},
		"ipv4": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ipv4": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"multicast": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"conditional_advertisement": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"exist": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"advertise_map": basetypes.StringType{},
												"exist_map":     basetypes.StringType{},
											},
										},
										"non_exist": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"advertise_map": basetypes.StringType{},
												"non_exist_map": basetypes.StringType{},
											},
										},
									},
								},
								"filter_list": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"inbound":  basetypes.StringType{},
										"outbound": basetypes.StringType{},
									},
								},
								"inbound_network_filters": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"distribute_list": basetypes.StringType{},
										"prefix_list":     basetypes.StringType{},
									},
								},
								"inherit": basetypes.BoolType{},
								"outbound_network_filters": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"distribute_list": basetypes.StringType{},
										"prefix_list":     basetypes.StringType{},
									},
								},
								"route_maps": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"inbound":  basetypes.StringType{},
										"outbound": basetypes.StringType{},
									},
								},
								"unsuppress_map": basetypes.StringType{},
							},
						},
						"unicast": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"conditional_advertisement": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"exist": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"advertise_map": basetypes.StringType{},
												"exist_map":     basetypes.StringType{},
											},
										},
										"non_exist": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"advertise_map": basetypes.StringType{},
												"non_exist_map": basetypes.StringType{},
											},
										},
									},
								},
								"filter_list": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"inbound":  basetypes.StringType{},
										"outbound": basetypes.StringType{},
									},
								},
								"inbound_network_filters": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"distribute_list": basetypes.StringType{},
										"prefix_list":     basetypes.StringType{},
									},
								},
								"outbound_network_filters": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"distribute_list": basetypes.StringType{},
										"prefix_list":     basetypes.StringType{},
									},
								},
								"route_maps": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"inbound":  basetypes.StringType{},
										"outbound": basetypes.StringType{},
									},
								},
								"unsuppress_map": basetypes.StringType{},
							},
						},
					},
				},
			},
		},
		"name":    basetypes.StringType{},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpFilteringProfiles objects.
func (o BgpFilteringProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpFilteringProfilesIpv4 model.
func (o BgpFilteringProfilesIpv4) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ipv4": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"multicast": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"conditional_advertisement": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"exist": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"advertise_map": basetypes.StringType{},
										"exist_map":     basetypes.StringType{},
									},
								},
								"non_exist": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"advertise_map": basetypes.StringType{},
										"non_exist_map": basetypes.StringType{},
									},
								},
							},
						},
						"filter_list": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"inbound":  basetypes.StringType{},
								"outbound": basetypes.StringType{},
							},
						},
						"inbound_network_filters": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"distribute_list": basetypes.StringType{},
								"prefix_list":     basetypes.StringType{},
							},
						},
						"inherit": basetypes.BoolType{},
						"outbound_network_filters": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"distribute_list": basetypes.StringType{},
								"prefix_list":     basetypes.StringType{},
							},
						},
						"route_maps": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"inbound":  basetypes.StringType{},
								"outbound": basetypes.StringType{},
							},
						},
						"unsuppress_map": basetypes.StringType{},
					},
				},
				"unicast": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"conditional_advertisement": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"exist": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"advertise_map": basetypes.StringType{},
										"exist_map":     basetypes.StringType{},
									},
								},
								"non_exist": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"advertise_map": basetypes.StringType{},
										"non_exist_map": basetypes.StringType{},
									},
								},
							},
						},
						"filter_list": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"inbound":  basetypes.StringType{},
								"outbound": basetypes.StringType{},
							},
						},
						"inbound_network_filters": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"distribute_list": basetypes.StringType{},
								"prefix_list":     basetypes.StringType{},
							},
						},
						"outbound_network_filters": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"distribute_list": basetypes.StringType{},
								"prefix_list":     basetypes.StringType{},
							},
						},
						"route_maps": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"inbound":  basetypes.StringType{},
								"outbound": basetypes.StringType{},
							},
						},
						"unsuppress_map": basetypes.StringType{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of BgpFilteringProfilesIpv4 objects.
func (o BgpFilteringProfilesIpv4) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpFilteringProfilesIpv4Ipv4 model.
func (o BgpFilteringProfilesIpv4Ipv4) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"multicast": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"conditional_advertisement": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"exist": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"advertise_map": basetypes.StringType{},
								"exist_map":     basetypes.StringType{},
							},
						},
						"non_exist": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"advertise_map": basetypes.StringType{},
								"non_exist_map": basetypes.StringType{},
							},
						},
					},
				},
				"filter_list": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"inbound":  basetypes.StringType{},
						"outbound": basetypes.StringType{},
					},
				},
				"inbound_network_filters": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"distribute_list": basetypes.StringType{},
						"prefix_list":     basetypes.StringType{},
					},
				},
				"inherit": basetypes.BoolType{},
				"outbound_network_filters": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"distribute_list": basetypes.StringType{},
						"prefix_list":     basetypes.StringType{},
					},
				},
				"route_maps": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"inbound":  basetypes.StringType{},
						"outbound": basetypes.StringType{},
					},
				},
				"unsuppress_map": basetypes.StringType{},
			},
		},
		"unicast": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"conditional_advertisement": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"exist": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"advertise_map": basetypes.StringType{},
								"exist_map":     basetypes.StringType{},
							},
						},
						"non_exist": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"advertise_map": basetypes.StringType{},
								"non_exist_map": basetypes.StringType{},
							},
						},
					},
				},
				"filter_list": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"inbound":  basetypes.StringType{},
						"outbound": basetypes.StringType{},
					},
				},
				"inbound_network_filters": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"distribute_list": basetypes.StringType{},
						"prefix_list":     basetypes.StringType{},
					},
				},
				"outbound_network_filters": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"distribute_list": basetypes.StringType{},
						"prefix_list":     basetypes.StringType{},
					},
				},
				"route_maps": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"inbound":  basetypes.StringType{},
						"outbound": basetypes.StringType{},
					},
				},
				"unsuppress_map": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of BgpFilteringProfilesIpv4Ipv4 objects.
func (o BgpFilteringProfilesIpv4Ipv4) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpFilteringProfilesIpv4Ipv4Multicast model.
func (o BgpFilteringProfilesIpv4Ipv4Multicast) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"conditional_advertisement": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"exist": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"advertise_map": basetypes.StringType{},
						"exist_map":     basetypes.StringType{},
					},
				},
				"non_exist": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"advertise_map": basetypes.StringType{},
						"non_exist_map": basetypes.StringType{},
					},
				},
			},
		},
		"filter_list": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"inbound":  basetypes.StringType{},
				"outbound": basetypes.StringType{},
			},
		},
		"inbound_network_filters": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"distribute_list": basetypes.StringType{},
				"prefix_list":     basetypes.StringType{},
			},
		},
		"inherit": basetypes.BoolType{},
		"outbound_network_filters": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"distribute_list": basetypes.StringType{},
				"prefix_list":     basetypes.StringType{},
			},
		},
		"route_maps": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"inbound":  basetypes.StringType{},
				"outbound": basetypes.StringType{},
			},
		},
		"unsuppress_map": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpFilteringProfilesIpv4Ipv4Multicast objects.
func (o BgpFilteringProfilesIpv4Ipv4Multicast) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpFilterConditionalAdvertisement model.
func (o BgpFilterConditionalAdvertisement) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"exist": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"advertise_map": basetypes.StringType{},
				"exist_map":     basetypes.StringType{},
			},
		},
		"non_exist": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"advertise_map": basetypes.StringType{},
				"non_exist_map": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of BgpFilterConditionalAdvertisement objects.
func (o BgpFilterConditionalAdvertisement) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpFilterConditionalAdvertisementExist model.
func (o BgpFilterConditionalAdvertisementExist) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"advertise_map": basetypes.StringType{},
		"exist_map":     basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpFilterConditionalAdvertisementExist objects.
func (o BgpFilterConditionalAdvertisementExist) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpFilterConditionalAdvertisementNonExist model.
func (o BgpFilterConditionalAdvertisementNonExist) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"advertise_map": basetypes.StringType{},
		"non_exist_map": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpFilterConditionalAdvertisementNonExist objects.
func (o BgpFilterConditionalAdvertisementNonExist) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpFilterFilterList model.
func (o BgpFilterFilterList) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"inbound":  basetypes.StringType{},
		"outbound": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpFilterFilterList objects.
func (o BgpFilterFilterList) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpFilterInboundNetworkFilters model.
func (o BgpFilterInboundNetworkFilters) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"distribute_list": basetypes.StringType{},
		"prefix_list":     basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpFilterInboundNetworkFilters objects.
func (o BgpFilterInboundNetworkFilters) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpFilter model.
func (o BgpFilter) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"conditional_advertisement": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"exist": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"advertise_map": basetypes.StringType{},
						"exist_map":     basetypes.StringType{},
					},
				},
				"non_exist": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"advertise_map": basetypes.StringType{},
						"non_exist_map": basetypes.StringType{},
					},
				},
			},
		},
		"filter_list": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"inbound":  basetypes.StringType{},
				"outbound": basetypes.StringType{},
			},
		},
		"inbound_network_filters": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"distribute_list": basetypes.StringType{},
				"prefix_list":     basetypes.StringType{},
			},
		},
		"outbound_network_filters": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"distribute_list": basetypes.StringType{},
				"prefix_list":     basetypes.StringType{},
			},
		},
		"route_maps": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"inbound":  basetypes.StringType{},
				"outbound": basetypes.StringType{},
			},
		},
		"unsuppress_map": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpFilter objects.
func (o BgpFilter) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// BgpFilteringProfilesResourceSchema defines the schema for BgpFilteringProfiles resource
var BgpFilteringProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "BgpFilteringProfile resource",
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
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
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
					path.MatchRelative().AtParent().AtName("snippet"),
					path.MatchRelative().AtParent().AtName("device"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
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
		"ipv4": schema.SingleNestedAttribute{
			MarkdownDescription: "Ipv4",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"ipv4": schema.SingleNestedAttribute{
					MarkdownDescription: "Ipv4",
					Required:            true,
					Attributes: map[string]schema.Attribute{
						"multicast": schema.SingleNestedAttribute{
							MarkdownDescription: "Multicast",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"conditional_advertisement": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("inherit"),
											path.MatchRelative().AtParent().AtName("unsuppress_map"),
											path.MatchRelative().AtParent().AtName("filter_list"),
											path.MatchRelative().AtParent().AtName("inbound_network_filters"),
											path.MatchRelative().AtParent().AtName("outbound_network_filters"),
											path.MatchRelative().AtParent().AtName("route_maps"),
										),
									},
									MarkdownDescription: "Conditional advertisement",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"exist": schema.SingleNestedAttribute{
											MarkdownDescription: "Exist",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"advertise_map": schema.StringAttribute{
													MarkdownDescription: "Advertise map",
													Optional:            true,
												},
												"exist_map": schema.StringAttribute{
													MarkdownDescription: "Exist map",
													Optional:            true,
												},
											},
										},
										"non_exist": schema.SingleNestedAttribute{
											MarkdownDescription: "Non exist",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"advertise_map": schema.StringAttribute{
													MarkdownDescription: "Advertise map",
													Optional:            true,
												},
												"non_exist_map": schema.StringAttribute{
													MarkdownDescription: "Non exist map",
													Optional:            true,
												},
											},
										},
									},
								},
								"filter_list": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("inherit"),
											path.MatchRelative().AtParent().AtName("unsuppress_map"),
											path.MatchRelative().AtParent().AtName("conditional_advertisement"),
											path.MatchRelative().AtParent().AtName("inbound_network_filters"),
											path.MatchRelative().AtParent().AtName("outbound_network_filters"),
											path.MatchRelative().AtParent().AtName("route_maps"),
										),
									},
									MarkdownDescription: "Filter list",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"inbound": schema.StringAttribute{
											MarkdownDescription: "Inbound",
											Optional:            true,
										},
										"outbound": schema.StringAttribute{
											MarkdownDescription: "Outbound",
											Optional:            true,
										},
									},
								},
								"inbound_network_filters": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("inherit"),
											path.MatchRelative().AtParent().AtName("unsuppress_map"),
											path.MatchRelative().AtParent().AtName("conditional_advertisement"),
											path.MatchRelative().AtParent().AtName("filter_list"),
											path.MatchRelative().AtParent().AtName("outbound_network_filters"),
											path.MatchRelative().AtParent().AtName("route_maps"),
										),
									},
									MarkdownDescription: "Inbound network filters",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"distribute_list": schema.StringAttribute{
											MarkdownDescription: "Distribute list",
											Optional:            true,
										},
										"prefix_list": schema.StringAttribute{
											MarkdownDescription: "Prefix list",
											Optional:            true,
										},
									},
								},
								"inherit": schema.BoolAttribute{
									Validators: []validator.Bool{
										boolvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("unsuppress_map"),
											path.MatchRelative().AtParent().AtName("conditional_advertisement"),
											path.MatchRelative().AtParent().AtName("filter_list"),
											path.MatchRelative().AtParent().AtName("inbound_network_filters"),
											path.MatchRelative().AtParent().AtName("outbound_network_filters"),
											path.MatchRelative().AtParent().AtName("route_maps"),
										),
									},
									MarkdownDescription: "Inherit from unicast",
									Optional:            true,
								},
								"outbound_network_filters": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("inherit"),
											path.MatchRelative().AtParent().AtName("unsuppress_map"),
											path.MatchRelative().AtParent().AtName("conditional_advertisement"),
											path.MatchRelative().AtParent().AtName("filter_list"),
											path.MatchRelative().AtParent().AtName("inbound_network_filters"),
											path.MatchRelative().AtParent().AtName("route_maps"),
										),
									},
									MarkdownDescription: "Outbound network filters",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"distribute_list": schema.StringAttribute{
											MarkdownDescription: "Distribute list",
											Optional:            true,
										},
										"prefix_list": schema.StringAttribute{
											MarkdownDescription: "Prefix list",
											Optional:            true,
										},
									},
								},
								"route_maps": schema.SingleNestedAttribute{
									Validators: []validator.Object{
										objectvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("inherit"),
											path.MatchRelative().AtParent().AtName("unsuppress_map"),
											path.MatchRelative().AtParent().AtName("conditional_advertisement"),
											path.MatchRelative().AtParent().AtName("filter_list"),
											path.MatchRelative().AtParent().AtName("inbound_network_filters"),
											path.MatchRelative().AtParent().AtName("outbound_network_filters"),
										),
									},
									MarkdownDescription: "Route maps",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"inbound": schema.StringAttribute{
											MarkdownDescription: "Inbound",
											Optional:            true,
										},
										"outbound": schema.StringAttribute{
											MarkdownDescription: "Outbound",
											Optional:            true,
										},
									},
								},
								"unsuppress_map": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.ExactlyOneOf(
											path.MatchRelative().AtParent().AtName("inherit"),
											path.MatchRelative().AtParent().AtName("conditional_advertisement"),
											path.MatchRelative().AtParent().AtName("filter_list"),
											path.MatchRelative().AtParent().AtName("inbound_network_filters"),
											path.MatchRelative().AtParent().AtName("outbound_network_filters"),
											path.MatchRelative().AtParent().AtName("route_maps"),
										),
									},
									MarkdownDescription: "Unsuppress map",
									Optional:            true,
								},
							},
						},
						"unicast": schema.SingleNestedAttribute{
							MarkdownDescription: "Unicast",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"conditional_advertisement": schema.SingleNestedAttribute{
									MarkdownDescription: "Conditional advertisement",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"exist": schema.SingleNestedAttribute{
											MarkdownDescription: "Exist",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"advertise_map": schema.StringAttribute{
													MarkdownDescription: "Advertise map",
													Optional:            true,
												},
												"exist_map": schema.StringAttribute{
													MarkdownDescription: "Exist map",
													Optional:            true,
												},
											},
										},
										"non_exist": schema.SingleNestedAttribute{
											MarkdownDescription: "Non exist",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"advertise_map": schema.StringAttribute{
													MarkdownDescription: "Advertise map",
													Optional:            true,
												},
												"non_exist_map": schema.StringAttribute{
													MarkdownDescription: "Non exist map",
													Optional:            true,
												},
											},
										},
									},
								},
								"filter_list": schema.SingleNestedAttribute{
									MarkdownDescription: "Filter list",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"inbound": schema.StringAttribute{
											MarkdownDescription: "Inbound",
											Optional:            true,
										},
										"outbound": schema.StringAttribute{
											MarkdownDescription: "Outbound",
											Optional:            true,
										},
									},
								},
								"inbound_network_filters": schema.SingleNestedAttribute{
									MarkdownDescription: "Inbound network filters",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"distribute_list": schema.StringAttribute{
											MarkdownDescription: "Distribute list",
											Optional:            true,
										},
										"prefix_list": schema.StringAttribute{
											MarkdownDescription: "Prefix list",
											Optional:            true,
										},
									},
								},
								"outbound_network_filters": schema.SingleNestedAttribute{
									MarkdownDescription: "Outbound network filters",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"distribute_list": schema.StringAttribute{
											MarkdownDescription: "Distribute list",
											Optional:            true,
										},
										"prefix_list": schema.StringAttribute{
											MarkdownDescription: "Prefix list",
											Optional:            true,
										},
									},
								},
								"route_maps": schema.SingleNestedAttribute{
									MarkdownDescription: "Route maps",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"inbound": schema.StringAttribute{
											MarkdownDescription: "Inbound",
											Optional:            true,
										},
										"outbound": schema.StringAttribute{
											MarkdownDescription: "Outbound",
											Optional:            true,
										},
									},
								},
								"unsuppress_map": schema.StringAttribute{
									MarkdownDescription: "Unsuppress map",
									Optional:            true,
								},
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
		"snippet": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("folder"),
					path.MatchRelative().AtParent().AtName("device"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
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

// BgpFilteringProfilesDataSourceSchema defines the schema for BgpFilteringProfiles data source
var BgpFilteringProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "BgpFilteringProfile data source",
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
		"ipv4": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Ipv4",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"ipv4": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Ipv4",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"multicast": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Multicast",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"conditional_advertisement": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Conditional advertisement",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"exist": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Exist",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"advertise_map": dsschema.StringAttribute{
													MarkdownDescription: "Advertise map",
													Computed:            true,
												},
												"exist_map": dsschema.StringAttribute{
													MarkdownDescription: "Exist map",
													Computed:            true,
												},
											},
										},
										"non_exist": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Non exist",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"advertise_map": dsschema.StringAttribute{
													MarkdownDescription: "Advertise map",
													Computed:            true,
												},
												"non_exist_map": dsschema.StringAttribute{
													MarkdownDescription: "Non exist map",
													Computed:            true,
												},
											},
										},
									},
								},
								"filter_list": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Filter list",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"inbound": dsschema.StringAttribute{
											MarkdownDescription: "Inbound",
											Computed:            true,
										},
										"outbound": dsschema.StringAttribute{
											MarkdownDescription: "Outbound",
											Computed:            true,
										},
									},
								},
								"inbound_network_filters": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Inbound network filters",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"distribute_list": dsschema.StringAttribute{
											MarkdownDescription: "Distribute list",
											Computed:            true,
										},
										"prefix_list": dsschema.StringAttribute{
											MarkdownDescription: "Prefix list",
											Computed:            true,
										},
									},
								},
								"inherit": dsschema.BoolAttribute{
									MarkdownDescription: "Inherit from unicast",
									Computed:            true,
								},
								"outbound_network_filters": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Outbound network filters",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"distribute_list": dsschema.StringAttribute{
											MarkdownDescription: "Distribute list",
											Computed:            true,
										},
										"prefix_list": dsschema.StringAttribute{
											MarkdownDescription: "Prefix list",
											Computed:            true,
										},
									},
								},
								"route_maps": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Route maps",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"inbound": dsschema.StringAttribute{
											MarkdownDescription: "Inbound",
											Computed:            true,
										},
										"outbound": dsschema.StringAttribute{
											MarkdownDescription: "Outbound",
											Computed:            true,
										},
									},
								},
								"unsuppress_map": dsschema.StringAttribute{
									MarkdownDescription: "Unsuppress map",
									Computed:            true,
								},
							},
						},
						"unicast": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Unicast",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"conditional_advertisement": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Conditional advertisement",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"exist": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Exist",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"advertise_map": dsschema.StringAttribute{
													MarkdownDescription: "Advertise map",
													Computed:            true,
												},
												"exist_map": dsschema.StringAttribute{
													MarkdownDescription: "Exist map",
													Computed:            true,
												},
											},
										},
										"non_exist": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Non exist",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"advertise_map": dsschema.StringAttribute{
													MarkdownDescription: "Advertise map",
													Computed:            true,
												},
												"non_exist_map": dsschema.StringAttribute{
													MarkdownDescription: "Non exist map",
													Computed:            true,
												},
											},
										},
									},
								},
								"filter_list": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Filter list",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"inbound": dsschema.StringAttribute{
											MarkdownDescription: "Inbound",
											Computed:            true,
										},
										"outbound": dsschema.StringAttribute{
											MarkdownDescription: "Outbound",
											Computed:            true,
										},
									},
								},
								"inbound_network_filters": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Inbound network filters",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"distribute_list": dsschema.StringAttribute{
											MarkdownDescription: "Distribute list",
											Computed:            true,
										},
										"prefix_list": dsschema.StringAttribute{
											MarkdownDescription: "Prefix list",
											Computed:            true,
										},
									},
								},
								"outbound_network_filters": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Outbound network filters",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"distribute_list": dsschema.StringAttribute{
											MarkdownDescription: "Distribute list",
											Computed:            true,
										},
										"prefix_list": dsschema.StringAttribute{
											MarkdownDescription: "Prefix list",
											Computed:            true,
										},
									},
								},
								"route_maps": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Route maps",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"inbound": dsschema.StringAttribute{
											MarkdownDescription: "Inbound",
											Computed:            true,
										},
										"outbound": dsschema.StringAttribute{
											MarkdownDescription: "Outbound",
											Computed:            true,
										},
									},
								},
								"unsuppress_map": dsschema.StringAttribute{
									MarkdownDescription: "Unsuppress map",
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Name",
			Optional:            true,
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
	},
}

// BgpFilteringProfilesListModel represents the data model for a list data source.
type BgpFilteringProfilesListModel struct {
	Tfid    types.String           `tfsdk:"tfid"`
	Data    []BgpFilteringProfiles `tfsdk:"data"`
	Limit   types.Int64            `tfsdk:"limit"`
	Offset  types.Int64            `tfsdk:"offset"`
	Name    types.String           `tfsdk:"name"`
	Total   types.Int64            `tfsdk:"total"`
	Folder  types.String           `tfsdk:"folder"`
	Snippet types.String           `tfsdk:"snippet"`
	Device  types.String           `tfsdk:"device"`
}

// BgpFilteringProfilesListDataSourceSchema defines the schema for a list data source.
var BgpFilteringProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: BgpFilteringProfilesDataSourceSchema.Attributes,
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
