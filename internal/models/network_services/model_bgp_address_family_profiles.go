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

// BgpAddressFamilyProfiles represents the Terraform model for BgpAddressFamilyProfiles
type BgpAddressFamilyProfiles struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Device  basetypes.StringValue `tfsdk:"device"`
	Folder  basetypes.StringValue `tfsdk:"folder"`
	Id      basetypes.StringValue `tfsdk:"id"`
	Ipv4    basetypes.ObjectValue `tfsdk:"ipv4"`
	Name    basetypes.StringValue `tfsdk:"name"`
	Snippet basetypes.StringValue `tfsdk:"snippet"`
}

// BgpAddressFamilyProfilesIpv4 represents a nested structure within the BgpAddressFamilyProfiles model
type BgpAddressFamilyProfilesIpv4 struct {
	Ipv4 basetypes.ObjectValue `tfsdk:"ipv4"`
}

// BgpAddressFamilyProfilesIpv4Ipv4 represents a nested structure within the BgpAddressFamilyProfiles model
type BgpAddressFamilyProfilesIpv4Ipv4 struct {
	Multicast basetypes.ObjectValue `tfsdk:"multicast"`
	Unicast   basetypes.ObjectValue `tfsdk:"unicast"`
}

// BgpAddressFamily represents a nested structure within the BgpAddressFamilyProfiles model
type BgpAddressFamily struct {
	AddPath                    basetypes.ObjectValue `tfsdk:"add_path"`
	AllowasIn                  basetypes.ObjectValue `tfsdk:"allowas_in"`
	AsOverride                 basetypes.BoolValue   `tfsdk:"as_override"`
	DefaultOriginate           basetypes.BoolValue   `tfsdk:"default_originate"`
	DefaultOriginateMap        basetypes.StringValue `tfsdk:"default_originate_map"`
	Enable                     basetypes.BoolValue   `tfsdk:"enable"`
	MaximumPrefix              basetypes.ObjectValue `tfsdk:"maximum_prefix"`
	NextHop                    basetypes.ObjectValue `tfsdk:"next_hop"`
	Orf                        basetypes.ObjectValue `tfsdk:"orf"`
	RemovePrivateAS            basetypes.ObjectValue `tfsdk:"remove_private_as"`
	RouteReflectorClient       basetypes.BoolValue   `tfsdk:"route_reflector_client"`
	SendCommunity              basetypes.ObjectValue `tfsdk:"send_community"`
	SoftReconfigWithStoredInfo basetypes.BoolValue   `tfsdk:"soft_reconfig_with_stored_info"`
}

// BgpAddressFamilyAddPath represents a nested structure within the BgpAddressFamilyProfiles model
type BgpAddressFamilyAddPath struct {
	TxAllPaths      basetypes.BoolValue `tfsdk:"tx_all_paths"`
	TxBestpathPerAS basetypes.BoolValue `tfsdk:"tx_bestpath_per_as"`
}

// BgpAddressFamilyAllowasIn represents a nested structure within the BgpAddressFamilyProfiles model
type BgpAddressFamilyAllowasIn struct {
	Occurrence basetypes.Int64Value  `tfsdk:"occurrence"`
	Origin     basetypes.ObjectValue `tfsdk:"origin"`
}

// BgpAddressFamilyMaximumPrefix represents a nested structure within the BgpAddressFamilyProfiles model
type BgpAddressFamilyMaximumPrefix struct {
	Action      basetypes.ObjectValue `tfsdk:"action"`
	NumPrefixes basetypes.Int64Value  `tfsdk:"num_prefixes"`
	Threshold   basetypes.Int64Value  `tfsdk:"threshold"`
}

// BgpAddressFamilyMaximumPrefixAction represents a nested structure within the BgpAddressFamilyProfiles model
type BgpAddressFamilyMaximumPrefixAction struct {
	Restart     basetypes.ObjectValue `tfsdk:"restart"`
	WarningOnly basetypes.ObjectValue `tfsdk:"warning_only"`
}

// BgpAddressFamilyMaximumPrefixActionRestart represents a nested structure within the BgpAddressFamilyProfiles model
type BgpAddressFamilyMaximumPrefixActionRestart struct {
	Interval basetypes.Int64Value `tfsdk:"interval"`
}

// BgpAddressFamilyNextHop represents a nested structure within the BgpAddressFamilyProfiles model
type BgpAddressFamilyNextHop struct {
	Self      basetypes.ObjectValue `tfsdk:"self"`
	SelfForce basetypes.ObjectValue `tfsdk:"self_force"`
}

// BgpAddressFamilyOrf represents a nested structure within the BgpAddressFamilyProfiles model
type BgpAddressFamilyOrf struct {
	OrfPrefixList basetypes.StringValue `tfsdk:"orf_prefix_list"`
}

// BgpAddressFamilyRemovePrivateAS represents a nested structure within the BgpAddressFamilyProfiles model
type BgpAddressFamilyRemovePrivateAS struct {
	All       basetypes.ObjectValue `tfsdk:"all"`
	ReplaceAS basetypes.ObjectValue `tfsdk:"replace_as"`
}

// BgpAddressFamilySendCommunity represents a nested structure within the BgpAddressFamilyProfiles model
type BgpAddressFamilySendCommunity struct {
	All      basetypes.ObjectValue `tfsdk:"all"`
	Both     basetypes.ObjectValue `tfsdk:"both"`
	Extended basetypes.ObjectValue `tfsdk:"extended"`
	Large    basetypes.ObjectValue `tfsdk:"large"`
	Standard basetypes.ObjectValue `tfsdk:"standard"`
}

// AttrTypes defines the attribute types for the BgpAddressFamilyProfiles model.
func (o BgpAddressFamilyProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":   basetypes.StringType{},
		"device": basetypes.StringType{},
		"folder": basetypes.StringType{},
		"id":     basetypes.StringType{},
		"ipv4": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ipv4": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"multicast": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"add_path": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"tx_all_paths":       basetypes.BoolType{},
										"tx_bestpath_per_as": basetypes.BoolType{},
									},
								},
								"allowas_in": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"occurrence": basetypes.Int64Type{},
										"origin": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								},
								"as_override":           basetypes.BoolType{},
								"default_originate":     basetypes.BoolType{},
								"default_originate_map": basetypes.StringType{},
								"enable":                basetypes.BoolType{},
								"maximum_prefix": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"action": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"restart": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"interval": basetypes.Int64Type{},
													},
												},
												"warning_only": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
											},
										},
										"num_prefixes": basetypes.Int64Type{},
										"threshold":    basetypes.Int64Type{},
									},
								},
								"next_hop": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"self": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"self_force": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								},
								"orf": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"orf_prefix_list": basetypes.StringType{},
									},
								},
								"remove_private_as": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"all": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"replace_as": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								},
								"route_reflector_client": basetypes.BoolType{},
								"send_community": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"all": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"both": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"extended": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"large": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"standard": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								},
								"soft_reconfig_with_stored_info": basetypes.BoolType{},
							},
						},
						"unicast": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"add_path": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"tx_all_paths":       basetypes.BoolType{},
										"tx_bestpath_per_as": basetypes.BoolType{},
									},
								},
								"allowas_in": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"occurrence": basetypes.Int64Type{},
										"origin": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								},
								"as_override":           basetypes.BoolType{},
								"default_originate":     basetypes.BoolType{},
								"default_originate_map": basetypes.StringType{},
								"enable":                basetypes.BoolType{},
								"maximum_prefix": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"action": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"restart": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"interval": basetypes.Int64Type{},
													},
												},
												"warning_only": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{},
												},
											},
										},
										"num_prefixes": basetypes.Int64Type{},
										"threshold":    basetypes.Int64Type{},
									},
								},
								"next_hop": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"self": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"self_force": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								},
								"orf": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"orf_prefix_list": basetypes.StringType{},
									},
								},
								"remove_private_as": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"all": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"replace_as": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								},
								"route_reflector_client": basetypes.BoolType{},
								"send_community": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"all": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"both": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"extended": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"large": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
										"standard": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								},
								"soft_reconfig_with_stored_info": basetypes.BoolType{},
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

// AttrType returns the attribute type for a list of BgpAddressFamilyProfiles objects.
func (o BgpAddressFamilyProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpAddressFamilyProfilesIpv4 model.
func (o BgpAddressFamilyProfilesIpv4) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ipv4": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"multicast": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"add_path": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"tx_all_paths":       basetypes.BoolType{},
								"tx_bestpath_per_as": basetypes.BoolType{},
							},
						},
						"allowas_in": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"occurrence": basetypes.Int64Type{},
								"origin": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
						"as_override":           basetypes.BoolType{},
						"default_originate":     basetypes.BoolType{},
						"default_originate_map": basetypes.StringType{},
						"enable":                basetypes.BoolType{},
						"maximum_prefix": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"restart": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"interval": basetypes.Int64Type{},
											},
										},
										"warning_only": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								},
								"num_prefixes": basetypes.Int64Type{},
								"threshold":    basetypes.Int64Type{},
							},
						},
						"next_hop": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"self": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"self_force": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
						"orf": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"orf_prefix_list": basetypes.StringType{},
							},
						},
						"remove_private_as": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"all": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"replace_as": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
						"route_reflector_client": basetypes.BoolType{},
						"send_community": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"all": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"both": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"extended": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"large": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"standard": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
						"soft_reconfig_with_stored_info": basetypes.BoolType{},
					},
				},
				"unicast": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"add_path": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"tx_all_paths":       basetypes.BoolType{},
								"tx_bestpath_per_as": basetypes.BoolType{},
							},
						},
						"allowas_in": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"occurrence": basetypes.Int64Type{},
								"origin": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
						"as_override":           basetypes.BoolType{},
						"default_originate":     basetypes.BoolType{},
						"default_originate_map": basetypes.StringType{},
						"enable":                basetypes.BoolType{},
						"maximum_prefix": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"restart": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"interval": basetypes.Int64Type{},
											},
										},
										"warning_only": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{},
										},
									},
								},
								"num_prefixes": basetypes.Int64Type{},
								"threshold":    basetypes.Int64Type{},
							},
						},
						"next_hop": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"self": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"self_force": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
						"orf": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"orf_prefix_list": basetypes.StringType{},
							},
						},
						"remove_private_as": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"all": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"replace_as": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
						"route_reflector_client": basetypes.BoolType{},
						"send_community": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"all": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"both": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"extended": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"large": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"standard": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
						"soft_reconfig_with_stored_info": basetypes.BoolType{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of BgpAddressFamilyProfilesIpv4 objects.
func (o BgpAddressFamilyProfilesIpv4) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpAddressFamilyProfilesIpv4Ipv4 model.
func (o BgpAddressFamilyProfilesIpv4Ipv4) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"multicast": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"add_path": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"tx_all_paths":       basetypes.BoolType{},
						"tx_bestpath_per_as": basetypes.BoolType{},
					},
				},
				"allowas_in": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"occurrence": basetypes.Int64Type{},
						"origin": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
				"as_override":           basetypes.BoolType{},
				"default_originate":     basetypes.BoolType{},
				"default_originate_map": basetypes.StringType{},
				"enable":                basetypes.BoolType{},
				"maximum_prefix": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"restart": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"interval": basetypes.Int64Type{},
									},
								},
								"warning_only": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
						"num_prefixes": basetypes.Int64Type{},
						"threshold":    basetypes.Int64Type{},
					},
				},
				"next_hop": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"self": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"self_force": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
				"orf": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"orf_prefix_list": basetypes.StringType{},
					},
				},
				"remove_private_as": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"all": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"replace_as": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
				"route_reflector_client": basetypes.BoolType{},
				"send_community": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"all": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"both": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"extended": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"large": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"standard": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
				"soft_reconfig_with_stored_info": basetypes.BoolType{},
			},
		},
		"unicast": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"add_path": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"tx_all_paths":       basetypes.BoolType{},
						"tx_bestpath_per_as": basetypes.BoolType{},
					},
				},
				"allowas_in": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"occurrence": basetypes.Int64Type{},
						"origin": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
				"as_override":           basetypes.BoolType{},
				"default_originate":     basetypes.BoolType{},
				"default_originate_map": basetypes.StringType{},
				"enable":                basetypes.BoolType{},
				"maximum_prefix": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"restart": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"interval": basetypes.Int64Type{},
									},
								},
								"warning_only": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
							},
						},
						"num_prefixes": basetypes.Int64Type{},
						"threshold":    basetypes.Int64Type{},
					},
				},
				"next_hop": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"self": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"self_force": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
				"orf": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"orf_prefix_list": basetypes.StringType{},
					},
				},
				"remove_private_as": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"all": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"replace_as": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
				"route_reflector_client": basetypes.BoolType{},
				"send_community": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"all": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"both": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"extended": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"large": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"standard": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
				"soft_reconfig_with_stored_info": basetypes.BoolType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of BgpAddressFamilyProfilesIpv4Ipv4 objects.
func (o BgpAddressFamilyProfilesIpv4Ipv4) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpAddressFamily model.
func (o BgpAddressFamily) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"add_path": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"tx_all_paths":       basetypes.BoolType{},
				"tx_bestpath_per_as": basetypes.BoolType{},
			},
		},
		"allowas_in": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"occurrence": basetypes.Int64Type{},
				"origin": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
		"as_override":           basetypes.BoolType{},
		"default_originate":     basetypes.BoolType{},
		"default_originate_map": basetypes.StringType{},
		"enable":                basetypes.BoolType{},
		"maximum_prefix": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"restart": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"interval": basetypes.Int64Type{},
							},
						},
						"warning_only": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
				"num_prefixes": basetypes.Int64Type{},
				"threshold":    basetypes.Int64Type{},
			},
		},
		"next_hop": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"self": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"self_force": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
		"orf": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"orf_prefix_list": basetypes.StringType{},
			},
		},
		"remove_private_as": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"all": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"replace_as": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
		"route_reflector_client": basetypes.BoolType{},
		"send_community": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"all": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"both": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"extended": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"large": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"standard": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
		"soft_reconfig_with_stored_info": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of BgpAddressFamily objects.
func (o BgpAddressFamily) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpAddressFamilyAddPath model.
func (o BgpAddressFamilyAddPath) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tx_all_paths":       basetypes.BoolType{},
		"tx_bestpath_per_as": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of BgpAddressFamilyAddPath objects.
func (o BgpAddressFamilyAddPath) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpAddressFamilyAllowasIn model.
func (o BgpAddressFamilyAllowasIn) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"occurrence": basetypes.Int64Type{},
		"origin": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of BgpAddressFamilyAllowasIn objects.
func (o BgpAddressFamilyAllowasIn) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpAddressFamilyMaximumPrefix model.
func (o BgpAddressFamilyMaximumPrefix) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"restart": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"interval": basetypes.Int64Type{},
					},
				},
				"warning_only": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
		"num_prefixes": basetypes.Int64Type{},
		"threshold":    basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of BgpAddressFamilyMaximumPrefix objects.
func (o BgpAddressFamilyMaximumPrefix) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpAddressFamilyMaximumPrefixAction model.
func (o BgpAddressFamilyMaximumPrefixAction) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"restart": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"interval": basetypes.Int64Type{},
			},
		},
		"warning_only": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of BgpAddressFamilyMaximumPrefixAction objects.
func (o BgpAddressFamilyMaximumPrefixAction) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpAddressFamilyMaximumPrefixActionRestart model.
func (o BgpAddressFamilyMaximumPrefixActionRestart) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"interval": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of BgpAddressFamilyMaximumPrefixActionRestart objects.
func (o BgpAddressFamilyMaximumPrefixActionRestart) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpAddressFamilyNextHop model.
func (o BgpAddressFamilyNextHop) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"self": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"self_force": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of BgpAddressFamilyNextHop objects.
func (o BgpAddressFamilyNextHop) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpAddressFamilyOrf model.
func (o BgpAddressFamilyOrf) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"orf_prefix_list": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of BgpAddressFamilyOrf objects.
func (o BgpAddressFamilyOrf) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpAddressFamilyRemovePrivateAS model.
func (o BgpAddressFamilyRemovePrivateAS) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"all": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"replace_as": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of BgpAddressFamilyRemovePrivateAS objects.
func (o BgpAddressFamilyRemovePrivateAS) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BgpAddressFamilySendCommunity model.
func (o BgpAddressFamilySendCommunity) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"all": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"both": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"extended": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"large": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"standard": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of BgpAddressFamilySendCommunity objects.
func (o BgpAddressFamilySendCommunity) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// BgpAddressFamilyProfilesResourceSchema defines the schema for BgpAddressFamilyProfiles resource
var BgpAddressFamilyProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "BgpAddressFamilyProfile resource",
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
								"add_path": schema.SingleNestedAttribute{
									MarkdownDescription: "Add path",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"tx_all_paths": schema.BoolAttribute{
											MarkdownDescription: "Advertise all paths to peer?",
											Optional:            true,
										},
										"tx_bestpath_per_as": schema.BoolAttribute{
											MarkdownDescription: "Tx bestpath per a s",
											Optional:            true,
										},
									},
								},
								"allowas_in": schema.SingleNestedAttribute{
									MarkdownDescription: "Allowas in",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"occurrence": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("origin"),
												),
												int64validator.Between(1, 10),
											},
											MarkdownDescription: "Number of times the firewalls own AS can be in an AS_PATH",
											Optional:            true,
										},
										"origin": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("occurrence"),
												),
											},
											MarkdownDescription: "Origin",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
									},
								},
								"as_override": schema.BoolAttribute{
									MarkdownDescription: "Override ASNs in outbound updates if AS-Path equals Remote-AS?",
									Optional:            true,
								},
								"default_originate": schema.BoolAttribute{
									MarkdownDescription: "Originate default route?",
									Optional:            true,
								},
								"default_originate_map": schema.StringAttribute{
									MarkdownDescription: "Default originate route map",
									Optional:            true,
								},
								"enable": schema.BoolAttribute{
									MarkdownDescription: "Enable?",
									Optional:            true,
								},
								"maximum_prefix": schema.SingleNestedAttribute{
									MarkdownDescription: "Maximum prefix",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"action": schema.SingleNestedAttribute{
											MarkdownDescription: "Action",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"restart": schema.SingleNestedAttribute{
													Validators: []validator.Object{
														objectvalidator.ConflictsWith(
															path.MatchRelative().AtParent().AtName("warning_only"),
														),
													},
													MarkdownDescription: "Restart",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"interval": schema.Int64Attribute{
															Validators: []validator.Int64{
																int64validator.Between(1, 65535),
															},
															MarkdownDescription: "Restart interval",
															Optional:            true,
														},
													},
												},
												"warning_only": schema.SingleNestedAttribute{
													Validators: []validator.Object{
														objectvalidator.ConflictsWith(
															path.MatchRelative().AtParent().AtName("restart"),
														),
													},
													MarkdownDescription: "Warning only",
													Optional:            true,
													Attributes:          map[string]schema.Attribute{},
												},
											},
										},
										"num_prefixes": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.Between(1, 4294967295),
											},
											MarkdownDescription: "Maximum number of prefixes",
											Optional:            true,
										},
										"threshold": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.Between(1, 100),
											},
											MarkdownDescription: "Threshold percentage of the maximum number of prefixes",
											Optional:            true,
										},
									},
								},
								"next_hop": schema.SingleNestedAttribute{
									MarkdownDescription: "Next hop",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"self": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("self_force"),
												),
											},
											MarkdownDescription: "Self",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
										"self_force": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("self"),
												),
											},
											MarkdownDescription: "Self force",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
									},
								},
								"orf": schema.SingleNestedAttribute{
									MarkdownDescription: "Orf",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"orf_prefix_list": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("none", "both", "receive", "send"),
											},
											MarkdownDescription: "ORF prefix list",
											Optional:            true,
										},
									},
								},
								"remove_private_as": schema.SingleNestedAttribute{
									MarkdownDescription: "Remove private a s",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"all": schema.SingleNestedAttribute{
											MarkdownDescription: "All",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
										"replace_as": schema.SingleNestedAttribute{
											MarkdownDescription: "Replace a s",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
									},
								},
								"route_reflector_client": schema.BoolAttribute{
									MarkdownDescription: "Route reflector client?",
									Optional:            true,
								},
								"send_community": schema.SingleNestedAttribute{
									MarkdownDescription: "Send community",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"all": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("both"),
													path.MatchRelative().AtParent().AtName("extended"),
													path.MatchRelative().AtParent().AtName("large"),
													path.MatchRelative().AtParent().AtName("standard"),
												),
											},
											MarkdownDescription: "All",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
										"both": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("all"),
													path.MatchRelative().AtParent().AtName("extended"),
													path.MatchRelative().AtParent().AtName("large"),
													path.MatchRelative().AtParent().AtName("standard"),
												),
											},
											MarkdownDescription: "Both",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
										"extended": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("all"),
													path.MatchRelative().AtParent().AtName("both"),
													path.MatchRelative().AtParent().AtName("large"),
													path.MatchRelative().AtParent().AtName("standard"),
												),
											},
											MarkdownDescription: "Extended",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
										"large": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("all"),
													path.MatchRelative().AtParent().AtName("both"),
													path.MatchRelative().AtParent().AtName("extended"),
													path.MatchRelative().AtParent().AtName("standard"),
												),
											},
											MarkdownDescription: "Large",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
										"standard": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("all"),
													path.MatchRelative().AtParent().AtName("both"),
													path.MatchRelative().AtParent().AtName("extended"),
													path.MatchRelative().AtParent().AtName("large"),
												),
											},
											MarkdownDescription: "Standard",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
									},
								},
								"soft_reconfig_with_stored_info": schema.BoolAttribute{
									MarkdownDescription: "Soft reconfiguration of peer with stored routes?",
									Optional:            true,
								},
							},
						},
						"unicast": schema.SingleNestedAttribute{
							MarkdownDescription: "Unicast",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"add_path": schema.SingleNestedAttribute{
									MarkdownDescription: "Add path",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"tx_all_paths": schema.BoolAttribute{
											MarkdownDescription: "Advertise all paths to peer?",
											Optional:            true,
										},
										"tx_bestpath_per_as": schema.BoolAttribute{
											MarkdownDescription: "Tx bestpath per a s",
											Optional:            true,
										},
									},
								},
								"allowas_in": schema.SingleNestedAttribute{
									MarkdownDescription: "Allowas in",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"occurrence": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("origin"),
												),
												int64validator.Between(1, 10),
											},
											MarkdownDescription: "Number of times the firewalls own AS can be in an AS_PATH",
											Optional:            true,
										},
										"origin": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("occurrence"),
												),
											},
											MarkdownDescription: "Origin",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
									},
								},
								"as_override": schema.BoolAttribute{
									MarkdownDescription: "Override ASNs in outbound updates if AS-Path equals Remote-AS?",
									Optional:            true,
								},
								"default_originate": schema.BoolAttribute{
									MarkdownDescription: "Originate default route?",
									Optional:            true,
								},
								"default_originate_map": schema.StringAttribute{
									MarkdownDescription: "Default originate route map",
									Optional:            true,
								},
								"enable": schema.BoolAttribute{
									MarkdownDescription: "Enable?",
									Optional:            true,
								},
								"maximum_prefix": schema.SingleNestedAttribute{
									MarkdownDescription: "Maximum prefix",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"action": schema.SingleNestedAttribute{
											MarkdownDescription: "Action",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"restart": schema.SingleNestedAttribute{
													Validators: []validator.Object{
														objectvalidator.ConflictsWith(
															path.MatchRelative().AtParent().AtName("warning_only"),
														),
													},
													MarkdownDescription: "Restart",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"interval": schema.Int64Attribute{
															Validators: []validator.Int64{
																int64validator.Between(1, 65535),
															},
															MarkdownDescription: "Restart interval",
															Optional:            true,
														},
													},
												},
												"warning_only": schema.SingleNestedAttribute{
													Validators: []validator.Object{
														objectvalidator.ConflictsWith(
															path.MatchRelative().AtParent().AtName("restart"),
														),
													},
													MarkdownDescription: "Warning only",
													Optional:            true,
													Attributes:          map[string]schema.Attribute{},
												},
											},
										},
										"num_prefixes": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.Between(1, 4294967295),
											},
											MarkdownDescription: "Maximum number of prefixes",
											Optional:            true,
										},
										"threshold": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.Between(1, 100),
											},
											MarkdownDescription: "Threshold percentage of the maximum number of prefixes",
											Optional:            true,
										},
									},
								},
								"next_hop": schema.SingleNestedAttribute{
									MarkdownDescription: "Next hop",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"self": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("self_force"),
												),
											},
											MarkdownDescription: "Self",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
										"self_force": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("self"),
												),
											},
											MarkdownDescription: "Self force",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
									},
								},
								"orf": schema.SingleNestedAttribute{
									MarkdownDescription: "Orf",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"orf_prefix_list": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("none", "both", "receive", "send"),
											},
											MarkdownDescription: "ORF prefix list",
											Optional:            true,
										},
									},
								},
								"remove_private_as": schema.SingleNestedAttribute{
									MarkdownDescription: "Remove private a s",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"all": schema.SingleNestedAttribute{
											MarkdownDescription: "All",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
										"replace_as": schema.SingleNestedAttribute{
											MarkdownDescription: "Replace a s",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
									},
								},
								"route_reflector_client": schema.BoolAttribute{
									MarkdownDescription: "Route reflector client?",
									Optional:            true,
								},
								"send_community": schema.SingleNestedAttribute{
									MarkdownDescription: "Send community",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"all": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("both"),
													path.MatchRelative().AtParent().AtName("extended"),
													path.MatchRelative().AtParent().AtName("large"),
													path.MatchRelative().AtParent().AtName("standard"),
												),
											},
											MarkdownDescription: "All",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
										"both": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("all"),
													path.MatchRelative().AtParent().AtName("extended"),
													path.MatchRelative().AtParent().AtName("large"),
													path.MatchRelative().AtParent().AtName("standard"),
												),
											},
											MarkdownDescription: "Both",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
										"extended": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("all"),
													path.MatchRelative().AtParent().AtName("both"),
													path.MatchRelative().AtParent().AtName("large"),
													path.MatchRelative().AtParent().AtName("standard"),
												),
											},
											MarkdownDescription: "Extended",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
										"large": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("all"),
													path.MatchRelative().AtParent().AtName("both"),
													path.MatchRelative().AtParent().AtName("extended"),
													path.MatchRelative().AtParent().AtName("standard"),
												),
											},
											MarkdownDescription: "Large",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
										"standard": schema.SingleNestedAttribute{
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(
													path.MatchRelative().AtParent().AtName("all"),
													path.MatchRelative().AtParent().AtName("both"),
													path.MatchRelative().AtParent().AtName("extended"),
													path.MatchRelative().AtParent().AtName("large"),
												),
											},
											MarkdownDescription: "Standard",
											Optional:            true,
											Attributes:          map[string]schema.Attribute{},
										},
									},
								},
								"soft_reconfig_with_stored_info": schema.BoolAttribute{
									MarkdownDescription: "Soft reconfiguration of peer with stored routes?",
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

// BgpAddressFamilyProfilesDataSourceSchema defines the schema for BgpAddressFamilyProfiles data source
var BgpAddressFamilyProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "BgpAddressFamilyProfile data source",
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
								"add_path": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Add path",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"tx_all_paths": dsschema.BoolAttribute{
											MarkdownDescription: "Advertise all paths to peer?",
											Computed:            true,
										},
										"tx_bestpath_per_as": dsschema.BoolAttribute{
											MarkdownDescription: "Tx bestpath per a s",
											Computed:            true,
										},
									},
								},
								"allowas_in": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Allowas in",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"occurrence": dsschema.Int64Attribute{
											MarkdownDescription: "Number of times the firewalls own AS can be in an AS_PATH",
											Computed:            true,
										},
										"origin": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Origin",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
									},
								},
								"as_override": dsschema.BoolAttribute{
									MarkdownDescription: "Override ASNs in outbound updates if AS-Path equals Remote-AS?",
									Computed:            true,
								},
								"default_originate": dsschema.BoolAttribute{
									MarkdownDescription: "Originate default route?",
									Computed:            true,
								},
								"default_originate_map": dsschema.StringAttribute{
									MarkdownDescription: "Default originate route map",
									Computed:            true,
								},
								"enable": dsschema.BoolAttribute{
									MarkdownDescription: "Enable?",
									Computed:            true,
								},
								"maximum_prefix": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Maximum prefix",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"action": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Action",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"restart": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Restart",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"interval": dsschema.Int64Attribute{
															MarkdownDescription: "Restart interval",
															Computed:            true,
														},
													},
												},
												"warning_only": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Warning only",
													Computed:            true,
													Attributes:          map[string]dsschema.Attribute{},
												},
											},
										},
										"num_prefixes": dsschema.Int64Attribute{
											MarkdownDescription: "Maximum number of prefixes",
											Computed:            true,
										},
										"threshold": dsschema.Int64Attribute{
											MarkdownDescription: "Threshold percentage of the maximum number of prefixes",
											Computed:            true,
										},
									},
								},
								"next_hop": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Next hop",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"self": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Self",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
										"self_force": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Self force",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
									},
								},
								"orf": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Orf",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"orf_prefix_list": dsschema.StringAttribute{
											MarkdownDescription: "ORF prefix list",
											Computed:            true,
										},
									},
								},
								"remove_private_as": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Remove private a s",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"all": dsschema.SingleNestedAttribute{
											MarkdownDescription: "All",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
										"replace_as": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Replace a s",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
									},
								},
								"route_reflector_client": dsschema.BoolAttribute{
									MarkdownDescription: "Route reflector client?",
									Computed:            true,
								},
								"send_community": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Send community",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"all": dsschema.SingleNestedAttribute{
											MarkdownDescription: "All",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
										"both": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Both",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
										"extended": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Extended",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
										"large": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Large",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
										"standard": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Standard",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
									},
								},
								"soft_reconfig_with_stored_info": dsschema.BoolAttribute{
									MarkdownDescription: "Soft reconfiguration of peer with stored routes?",
									Computed:            true,
								},
							},
						},
						"unicast": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Unicast",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"add_path": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Add path",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"tx_all_paths": dsschema.BoolAttribute{
											MarkdownDescription: "Advertise all paths to peer?",
											Computed:            true,
										},
										"tx_bestpath_per_as": dsschema.BoolAttribute{
											MarkdownDescription: "Tx bestpath per a s",
											Computed:            true,
										},
									},
								},
								"allowas_in": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Allowas in",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"occurrence": dsschema.Int64Attribute{
											MarkdownDescription: "Number of times the firewalls own AS can be in an AS_PATH",
											Computed:            true,
										},
										"origin": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Origin",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
									},
								},
								"as_override": dsschema.BoolAttribute{
									MarkdownDescription: "Override ASNs in outbound updates if AS-Path equals Remote-AS?",
									Computed:            true,
								},
								"default_originate": dsschema.BoolAttribute{
									MarkdownDescription: "Originate default route?",
									Computed:            true,
								},
								"default_originate_map": dsschema.StringAttribute{
									MarkdownDescription: "Default originate route map",
									Computed:            true,
								},
								"enable": dsschema.BoolAttribute{
									MarkdownDescription: "Enable?",
									Computed:            true,
								},
								"maximum_prefix": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Maximum prefix",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"action": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Action",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"restart": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Restart",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"interval": dsschema.Int64Attribute{
															MarkdownDescription: "Restart interval",
															Computed:            true,
														},
													},
												},
												"warning_only": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Warning only",
													Computed:            true,
													Attributes:          map[string]dsschema.Attribute{},
												},
											},
										},
										"num_prefixes": dsschema.Int64Attribute{
											MarkdownDescription: "Maximum number of prefixes",
											Computed:            true,
										},
										"threshold": dsschema.Int64Attribute{
											MarkdownDescription: "Threshold percentage of the maximum number of prefixes",
											Computed:            true,
										},
									},
								},
								"next_hop": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Next hop",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"self": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Self",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
										"self_force": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Self force",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
									},
								},
								"orf": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Orf",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"orf_prefix_list": dsschema.StringAttribute{
											MarkdownDescription: "ORF prefix list",
											Computed:            true,
										},
									},
								},
								"remove_private_as": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Remove private a s",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"all": dsschema.SingleNestedAttribute{
											MarkdownDescription: "All",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
										"replace_as": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Replace a s",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
									},
								},
								"route_reflector_client": dsschema.BoolAttribute{
									MarkdownDescription: "Route reflector client?",
									Computed:            true,
								},
								"send_community": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Send community",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"all": dsschema.SingleNestedAttribute{
											MarkdownDescription: "All",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
										"both": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Both",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
										"extended": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Extended",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
										"large": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Large",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
										"standard": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Standard",
											Computed:            true,
											Attributes:          map[string]dsschema.Attribute{},
										},
									},
								},
								"soft_reconfig_with_stored_info": dsschema.BoolAttribute{
									MarkdownDescription: "Soft reconfiguration of peer with stored routes?",
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

// BgpAddressFamilyProfilesListModel represents the data model for a list data source.
type BgpAddressFamilyProfilesListModel struct {
	Tfid    types.String               `tfsdk:"tfid"`
	Data    []BgpAddressFamilyProfiles `tfsdk:"data"`
	Limit   types.Int64                `tfsdk:"limit"`
	Offset  types.Int64                `tfsdk:"offset"`
	Name    types.String               `tfsdk:"name"`
	Total   types.Int64                `tfsdk:"total"`
	Folder  types.String               `tfsdk:"folder"`
	Snippet types.String               `tfsdk:"snippet"`
	Device  types.String               `tfsdk:"device"`
}

// BgpAddressFamilyProfilesListDataSourceSchema defines the schema for a list data source.
var BgpAddressFamilyProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: BgpAddressFamilyProfilesDataSourceSchema.Attributes,
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
