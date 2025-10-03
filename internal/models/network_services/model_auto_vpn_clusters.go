package models

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
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

// AutoVpnClusters represents the Terraform model for AutoVpnClusters
type AutoVpnClusters struct {
	Tfid                   types.String          `tfsdk:"tfid"`
	Branches               basetypes.ListValue   `tfsdk:"branches"`
	EnableMeshBetweenHubs  basetypes.BoolValue   `tfsdk:"enable_mesh_between_hubs"`
	EnableMeshInterconnect basetypes.BoolValue   `tfsdk:"enable_mesh_interconnect"`
	EnableSdwan            basetypes.BoolValue   `tfsdk:"enable_sdwan"`
	Gateways               basetypes.ListValue   `tfsdk:"gateways"`
	Id                     basetypes.StringValue `tfsdk:"id"`
	Name                   basetypes.StringValue `tfsdk:"name"`
	Type                   basetypes.StringValue `tfsdk:"type"`
}

// AutoVpnClustersBranchesInner represents a nested structure within the AutoVpnClusters model
type AutoVpnClustersBranchesInner struct {
	BgpRedistributionProfile basetypes.StringValue `tfsdk:"bgp_redistribution_profile"`
	Interfaces               basetypes.ListValue   `tfsdk:"interfaces"`
	LogicalRouter            basetypes.StringValue `tfsdk:"logical_router"`
	Name                     basetypes.StringValue `tfsdk:"name"`
	PrivateInterfaces        basetypes.ListValue   `tfsdk:"private_interfaces"`
	Site                     basetypes.StringValue `tfsdk:"site"`
}

// AutoVpnClustersBranchesInnerInterfacesInner represents a nested structure within the AutoVpnClusters model
type AutoVpnClustersBranchesInnerInterfacesInner struct {
	DhcpIp            basetypes.StringValue `tfsdk:"dhcp_ip"`
	Name              basetypes.StringValue `tfsdk:"name"`
	SdwanLinkSettings basetypes.ObjectValue `tfsdk:"sdwan_link_settings"`
}

// AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings represents a nested structure within the AutoVpnClusters model
type AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings struct {
	SdwanGateway          basetypes.StringValue `tfsdk:"sdwan_gateway"`
	SdwanInterfaceProfile basetypes.StringValue `tfsdk:"sdwan_interface_profile"`
	UpstreamNat           basetypes.ObjectValue `tfsdk:"upstream_nat"`
}

// AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat represents a nested structure within the AutoVpnClusters model
type AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat struct {
	Enable   basetypes.BoolValue   `tfsdk:"enable"`
	StaticIp basetypes.ObjectValue `tfsdk:"static_ip"`
}

// AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp represents a nested structure within the AutoVpnClusters model
type AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp struct {
	Fqdn      basetypes.StringValue `tfsdk:"fqdn"`
	IpAddress basetypes.StringValue `tfsdk:"ip_address"`
}

// AutoVpnClustersBranchesInnerPrivateInterfacesInner represents a nested structure within the AutoVpnClusters model
type AutoVpnClustersBranchesInnerPrivateInterfacesInner struct {
	Name              basetypes.StringValue `tfsdk:"name"`
	SdwanLinkSettings basetypes.ObjectValue `tfsdk:"sdwan_link_settings"`
}

// AutoVpnClustersGatewaysInner represents a nested structure within the AutoVpnClusters model
type AutoVpnClustersGatewaysInner struct {
	AllowDiaVpnFailover      basetypes.BoolValue   `tfsdk:"allow_dia_vpn_failover"`
	BgpRedistributionProfile basetypes.StringValue `tfsdk:"bgp_redistribution_profile"`
	Interfaces               basetypes.ListValue   `tfsdk:"interfaces"`
	LogicalRouter            basetypes.StringValue `tfsdk:"logical_router"`
	Name                     basetypes.StringValue `tfsdk:"name"`
	Priority                 basetypes.StringValue `tfsdk:"priority"`
	PrivateInterfaces        basetypes.ListValue   `tfsdk:"private_interfaces"`
	Site                     basetypes.StringValue `tfsdk:"site"`
}

// AutoVpnClustersGatewaysInnerInterfacesInner represents a nested structure within the AutoVpnClusters model
type AutoVpnClustersGatewaysInnerInterfacesInner struct {
	DhcpIp            basetypes.StringValue `tfsdk:"dhcp_ip"`
	Name              basetypes.StringValue `tfsdk:"name"`
	SdwanLinkSettings basetypes.ObjectValue `tfsdk:"sdwan_link_settings"`
}

// AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings represents a nested structure within the AutoVpnClusters model
type AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings struct {
	SdwanGateway          basetypes.StringValue `tfsdk:"sdwan_gateway"`
	SdwanInterfaceProfile basetypes.StringValue `tfsdk:"sdwan_interface_profile"`
	UpstreamNat           basetypes.ObjectValue `tfsdk:"upstream_nat"`
}

// AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat represents a nested structure within the AutoVpnClusters model
type AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat struct {
	Enable   basetypes.BoolValue   `tfsdk:"enable"`
	StaticIp basetypes.ObjectValue `tfsdk:"static_ip"`
}

// AutoVpnClustersGatewaysInnerPrivateInterfacesInner represents a nested structure within the AutoVpnClusters model
type AutoVpnClustersGatewaysInnerPrivateInterfacesInner struct {
	Name              basetypes.StringValue `tfsdk:"name"`
	SdwanLinkSettings basetypes.ObjectValue `tfsdk:"sdwan_link_settings"`
}

// AttrTypes defines the attribute types for the AutoVpnClusters model.
func (o AutoVpnClusters) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"branches": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"bgp_redistribution_profile": basetypes.StringType{},
				"interfaces": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"dhcp_ip": basetypes.StringType{},
						"name":    basetypes.StringType{},
						"sdwan_link_settings": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"sdwan_gateway":           basetypes.StringType{},
								"sdwan_interface_profile": basetypes.StringType{},
								"upstream_nat": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"enable": basetypes.BoolType{},
										"static_ip": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"fqdn":       basetypes.StringType{},
												"ip_address": basetypes.StringType{},
											},
										},
									},
								},
							},
						},
					},
				}},
				"logical_router": basetypes.StringType{},
				"name":           basetypes.StringType{},
				"private_interfaces": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name": basetypes.StringType{},
						"sdwan_link_settings": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"sdwan_gateway":           basetypes.StringType{},
								"sdwan_interface_profile": basetypes.StringType{},
								"upstream_nat": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"enable": basetypes.BoolType{},
										"static_ip": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"fqdn":       basetypes.StringType{},
												"ip_address": basetypes.StringType{},
											},
										},
									},
								},
							},
						},
					},
				}},
				"site": basetypes.StringType{},
			},
		}},
		"enable_mesh_between_hubs": basetypes.BoolType{},
		"enable_mesh_interconnect": basetypes.BoolType{},
		"enable_sdwan":             basetypes.BoolType{},
		"gateways": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"allow_dia_vpn_failover":     basetypes.BoolType{},
				"bgp_redistribution_profile": basetypes.StringType{},
				"interfaces": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"dhcp_ip": basetypes.StringType{},
						"name":    basetypes.StringType{},
						"sdwan_link_settings": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"sdwan_gateway":           basetypes.StringType{},
								"sdwan_interface_profile": basetypes.StringType{},
								"upstream_nat": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"enable": basetypes.BoolType{},
										"static_ip": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"fqdn":       basetypes.StringType{},
												"ip_address": basetypes.StringType{},
											},
										},
									},
								},
							},
						},
					},
				}},
				"logical_router": basetypes.StringType{},
				"name":           basetypes.StringType{},
				"priority":       basetypes.StringType{},
				"private_interfaces": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name": basetypes.StringType{},
						"sdwan_link_settings": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"sdwan_gateway":           basetypes.StringType{},
								"sdwan_interface_profile": basetypes.StringType{},
								"upstream_nat": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"enable": basetypes.BoolType{},
										"static_ip": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"fqdn":       basetypes.StringType{},
												"ip_address": basetypes.StringType{},
											},
										},
									},
								},
							},
						},
					},
				}},
				"site": basetypes.StringType{},
			},
		}},
		"id":   basetypes.StringType{},
		"name": basetypes.StringType{},
		"type": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AutoVpnClusters objects.
func (o AutoVpnClusters) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AutoVpnClustersBranchesInner model.
func (o AutoVpnClustersBranchesInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"bgp_redistribution_profile": basetypes.StringType{},
		"interfaces": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"dhcp_ip": basetypes.StringType{},
				"name":    basetypes.StringType{},
				"sdwan_link_settings": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"sdwan_gateway":           basetypes.StringType{},
						"sdwan_interface_profile": basetypes.StringType{},
						"upstream_nat": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable": basetypes.BoolType{},
								"static_ip": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"fqdn":       basetypes.StringType{},
										"ip_address": basetypes.StringType{},
									},
								},
							},
						},
					},
				},
			},
		}},
		"logical_router": basetypes.StringType{},
		"name":           basetypes.StringType{},
		"private_interfaces": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name": basetypes.StringType{},
				"sdwan_link_settings": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"sdwan_gateway":           basetypes.StringType{},
						"sdwan_interface_profile": basetypes.StringType{},
						"upstream_nat": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable": basetypes.BoolType{},
								"static_ip": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"fqdn":       basetypes.StringType{},
										"ip_address": basetypes.StringType{},
									},
								},
							},
						},
					},
				},
			},
		}},
		"site": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AutoVpnClustersBranchesInner objects.
func (o AutoVpnClustersBranchesInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AutoVpnClustersBranchesInnerInterfacesInner model.
func (o AutoVpnClustersBranchesInnerInterfacesInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"dhcp_ip": basetypes.StringType{},
		"name":    basetypes.StringType{},
		"sdwan_link_settings": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"sdwan_gateway":           basetypes.StringType{},
				"sdwan_interface_profile": basetypes.StringType{},
				"upstream_nat": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.BoolType{},
						"static_ip": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"fqdn":       basetypes.StringType{},
								"ip_address": basetypes.StringType{},
							},
						},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of AutoVpnClustersBranchesInnerInterfacesInner objects.
func (o AutoVpnClustersBranchesInnerInterfacesInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings model.
func (o AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"sdwan_gateway":           basetypes.StringType{},
		"sdwan_interface_profile": basetypes.StringType{},
		"upstream_nat": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable": basetypes.BoolType{},
				"static_ip": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"fqdn":       basetypes.StringType{},
						"ip_address": basetypes.StringType{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings objects.
func (o AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettings) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat model.
func (o AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable": basetypes.BoolType{},
		"static_ip": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"fqdn":       basetypes.StringType{},
				"ip_address": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat objects.
func (o AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNat) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp model.
func (o AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"fqdn":       basetypes.StringType{},
		"ip_address": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp objects.
func (o AutoVpnClustersBranchesInnerInterfacesInnerSdwanLinkSettingsUpstreamNatStaticIp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AutoVpnClustersBranchesInnerPrivateInterfacesInner model.
func (o AutoVpnClustersBranchesInnerPrivateInterfacesInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": basetypes.StringType{},
		"sdwan_link_settings": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"sdwan_gateway":           basetypes.StringType{},
				"sdwan_interface_profile": basetypes.StringType{},
				"upstream_nat": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.BoolType{},
						"static_ip": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"fqdn":       basetypes.StringType{},
								"ip_address": basetypes.StringType{},
							},
						},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of AutoVpnClustersBranchesInnerPrivateInterfacesInner objects.
func (o AutoVpnClustersBranchesInnerPrivateInterfacesInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AutoVpnClustersGatewaysInner model.
func (o AutoVpnClustersGatewaysInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"allow_dia_vpn_failover":     basetypes.BoolType{},
		"bgp_redistribution_profile": basetypes.StringType{},
		"interfaces": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"dhcp_ip": basetypes.StringType{},
				"name":    basetypes.StringType{},
				"sdwan_link_settings": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"sdwan_gateway":           basetypes.StringType{},
						"sdwan_interface_profile": basetypes.StringType{},
						"upstream_nat": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable": basetypes.BoolType{},
								"static_ip": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"fqdn":       basetypes.StringType{},
										"ip_address": basetypes.StringType{},
									},
								},
							},
						},
					},
				},
			},
		}},
		"logical_router": basetypes.StringType{},
		"name":           basetypes.StringType{},
		"priority":       basetypes.StringType{},
		"private_interfaces": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name": basetypes.StringType{},
				"sdwan_link_settings": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"sdwan_gateway":           basetypes.StringType{},
						"sdwan_interface_profile": basetypes.StringType{},
						"upstream_nat": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable": basetypes.BoolType{},
								"static_ip": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"fqdn":       basetypes.StringType{},
										"ip_address": basetypes.StringType{},
									},
								},
							},
						},
					},
				},
			},
		}},
		"site": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AutoVpnClustersGatewaysInner objects.
func (o AutoVpnClustersGatewaysInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AutoVpnClustersGatewaysInnerInterfacesInner model.
func (o AutoVpnClustersGatewaysInnerInterfacesInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"dhcp_ip": basetypes.StringType{},
		"name":    basetypes.StringType{},
		"sdwan_link_settings": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"sdwan_gateway":           basetypes.StringType{},
				"sdwan_interface_profile": basetypes.StringType{},
				"upstream_nat": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.BoolType{},
						"static_ip": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"fqdn":       basetypes.StringType{},
								"ip_address": basetypes.StringType{},
							},
						},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of AutoVpnClustersGatewaysInnerInterfacesInner objects.
func (o AutoVpnClustersGatewaysInnerInterfacesInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings model.
func (o AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"sdwan_gateway":           basetypes.StringType{},
		"sdwan_interface_profile": basetypes.StringType{},
		"upstream_nat": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable": basetypes.BoolType{},
				"static_ip": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"fqdn":       basetypes.StringType{},
						"ip_address": basetypes.StringType{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings objects.
func (o AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettings) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat model.
func (o AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable": basetypes.BoolType{},
		"static_ip": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"fqdn":       basetypes.StringType{},
				"ip_address": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat objects.
func (o AutoVpnClustersGatewaysInnerInterfacesInnerSdwanLinkSettingsUpstreamNat) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AutoVpnClustersGatewaysInnerPrivateInterfacesInner model.
func (o AutoVpnClustersGatewaysInnerPrivateInterfacesInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": basetypes.StringType{},
		"sdwan_link_settings": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"sdwan_gateway":           basetypes.StringType{},
				"sdwan_interface_profile": basetypes.StringType{},
				"upstream_nat": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.BoolType{},
						"static_ip": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"fqdn":       basetypes.StringType{},
								"ip_address": basetypes.StringType{},
							},
						},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of AutoVpnClustersGatewaysInnerPrivateInterfacesInner objects.
func (o AutoVpnClustersGatewaysInnerPrivateInterfacesInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AutoVpnClustersResourceSchema defines the schema for AutoVpnClusters resource
var AutoVpnClustersResourceSchema = schema.Schema{
	MarkdownDescription: "AutoVpnCluster resource",
	Attributes: map[string]schema.Attribute{
		"branches": schema.ListNestedAttribute{
			MarkdownDescription: "Branches",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"bgp_redistribution_profile": schema.StringAttribute{
						MarkdownDescription: "BGP redistribution profile",
						Optional:            true,
					},
					"interfaces": schema.ListNestedAttribute{
						Validators: []validator.List{
							listvalidator.SizeAtMost(4),
						},
						MarkdownDescription: "Interfaces",
						Optional:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"dhcp_ip": schema.StringAttribute{
									MarkdownDescription: "DHCP IP",
									Optional:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: "Ethernet interface",
									Optional:            true,
								},
								"sdwan_link_settings": schema.SingleNestedAttribute{
									MarkdownDescription: "Sdwan link settings",
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{
										"sdwan_gateway": schema.StringAttribute{
											MarkdownDescription: "Next hop gateway",
											Optional:            true,
											Computed:            true,
										},
										"sdwan_interface_profile": schema.StringAttribute{
											MarkdownDescription: "SD-WAN interface profile",
											Optional:            true,
											Computed:            true,
										},
										"upstream_nat": schema.SingleNestedAttribute{
											MarkdownDescription: "Upstream nat",
											Optional:            true,
											Computed:            true,
											Attributes: map[string]schema.Attribute{
												"enable": schema.BoolAttribute{
													MarkdownDescription: "Upstream NAT?",
													Optional:            true,
													Computed:            true,
													Default:             booldefault.StaticBool(false),
												},
												"static_ip": schema.SingleNestedAttribute{
													MarkdownDescription: "Static ip",
													Optional:            true,
													Computed:            true,
													Attributes: map[string]schema.Attribute{
														"fqdn": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("ip_address"),
																),
															},
															MarkdownDescription: "FQDN",
															Optional:            true,
															Computed:            true,
														},
														"ip_address": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("fqdn"),
																),
															},
															MarkdownDescription: "IP address",
															Optional:            true,
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
					"logical_router": schema.StringAttribute{
						MarkdownDescription: "Router",
						Optional:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "Branch firewall serial number",
						Optional:            true,
					},
					"private_interfaces": schema.ListNestedAttribute{
						Validators: []validator.List{
							listvalidator.SizeAtMost(4),
						},
						MarkdownDescription: "Private interfaces",
						Optional:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									MarkdownDescription: "Ethernet interface",
									Optional:            true,
								},
								"sdwan_link_settings": schema.SingleNestedAttribute{
									MarkdownDescription: "Sdwan link settings",
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{
										"sdwan_gateway": schema.StringAttribute{
											MarkdownDescription: "Next hop gateway",
											Optional:            true,
											Computed:            true,
										},
										"sdwan_interface_profile": schema.StringAttribute{
											MarkdownDescription: "SD-WAN interface profile",
											Optional:            true,
											Computed:            true,
										},
										"upstream_nat": schema.SingleNestedAttribute{
											MarkdownDescription: "Upstream nat",
											Optional:            true,
											Computed:            true,
											Attributes: map[string]schema.Attribute{
												"enable": schema.BoolAttribute{
													MarkdownDescription: "Upstream NAT?",
													Optional:            true,
													Computed:            true,
													Default:             booldefault.StaticBool(false),
												},
												"static_ip": schema.SingleNestedAttribute{
													MarkdownDescription: "Static ip",
													Optional:            true,
													Computed:            true,
													Attributes: map[string]schema.Attribute{
														"fqdn": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("ip_address"),
																),
															},
															MarkdownDescription: "FQDN",
															Optional:            true,
															Computed:            true,
														},
														"ip_address": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("fqdn"),
																),
															},
															MarkdownDescription: "IP address",
															Optional:            true,
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
					"site": schema.StringAttribute{
						MarkdownDescription: "Site name",
						Optional:            true,
					},
				},
			},
		},
		"enable_mesh_between_hubs": schema.BoolAttribute{
			MarkdownDescription: "Enable mesh between hubs?",
			Optional:            true,
		},
		"enable_mesh_interconnect": schema.BoolAttribute{
			MarkdownDescription: "Enable mesh interconnect?",
			Optional:            true,
		},
		"enable_sdwan": schema.BoolAttribute{
			MarkdownDescription: "Enable SD-WAN?",
			Optional:            true,
		},
		"gateways": schema.ListNestedAttribute{
			MarkdownDescription: "Hubs",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"allow_dia_vpn_failover": schema.BoolAttribute{
						MarkdownDescription: "Allow DIA to VPN failover on branch device for the hub?",
						Optional:            true,
					},
					"bgp_redistribution_profile": schema.StringAttribute{
						MarkdownDescription: "BGP redistribution file",
						Optional:            true,
					},
					"interfaces": schema.ListNestedAttribute{
						MarkdownDescription: "Interfaces",
						Optional:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"dhcp_ip": schema.StringAttribute{
									MarkdownDescription: "DHCP IP",
									Optional:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: "Ethernet interface",
									Optional:            true,
								},
								"sdwan_link_settings": schema.SingleNestedAttribute{
									MarkdownDescription: "Sdwan link settings",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"sdwan_gateway": schema.StringAttribute{
											MarkdownDescription: "Next hop gateway",
											Optional:            true,
										},
										"sdwan_interface_profile": schema.StringAttribute{
											MarkdownDescription: "SD-WAN interface profile",
											Optional:            true,
										},
										"upstream_nat": schema.SingleNestedAttribute{
											MarkdownDescription: "Upstream nat",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"enable": schema.BoolAttribute{
													MarkdownDescription: "Upstream NAT?",
													Optional:            true,
												},
												"static_ip": schema.SingleNestedAttribute{
													MarkdownDescription: "Static ip",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"fqdn": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("ip_address"),
																),
															},
															MarkdownDescription: "FQDN",
															Optional:            true,
														},
														"ip_address": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("fqdn"),
																),
															},
															MarkdownDescription: "IP address",
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
					"logical_router": schema.StringAttribute{
						MarkdownDescription: "Router",
						Optional:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "Hub firewall serial number",
						Optional:            true,
					},
					"priority": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("1", "2", "3", "4", "5", "6", "7", "8"),
						},
						MarkdownDescription: "Priority",
						Optional:            true,
					},
					"private_interfaces": schema.ListNestedAttribute{
						MarkdownDescription: "Private interfaces",
						Optional:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									MarkdownDescription: "Ethernet interface",
									Optional:            true,
								},
								"sdwan_link_settings": schema.SingleNestedAttribute{
									MarkdownDescription: "Sdwan link settings",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"sdwan_gateway": schema.StringAttribute{
											MarkdownDescription: "Next hop gateway",
											Optional:            true,
										},
										"sdwan_interface_profile": schema.StringAttribute{
											MarkdownDescription: "SD-WAN interface profile",
											Optional:            true,
										},
										"upstream_nat": schema.SingleNestedAttribute{
											MarkdownDescription: "Upstream nat",
											Optional:            true,
											Attributes: map[string]schema.Attribute{
												"enable": schema.BoolAttribute{
													MarkdownDescription: "Upstream NAT?",
													Optional:            true,
												},
												"static_ip": schema.SingleNestedAttribute{
													MarkdownDescription: "Static ip",
													Optional:            true,
													Attributes: map[string]schema.Attribute{
														"fqdn": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("ip_address"),
																),
															},
															MarkdownDescription: "FQDN",
															Optional:            true,
														},
														"ip_address": schema.StringAttribute{
															Validators: []validator.String{
																stringvalidator.ExactlyOneOf(
																	path.MatchRelative().AtParent().AtName("fqdn"),
																),
															},
															MarkdownDescription: "IP address",
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
					"site": schema.StringAttribute{
						MarkdownDescription: "Site name",
						Optional:            true,
					},
				},
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
			MarkdownDescription: "VPN cluster name",
			Optional:            true,
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"type": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("hub-spoke", "mesh"),
			},
			MarkdownDescription: "VPN cluster type",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("hub-spoke"),
		},
	},
}

// AutoVpnClustersDataSourceSchema defines the schema for AutoVpnClusters data source
var AutoVpnClustersDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "AutoVpnCluster data source",
	Attributes: map[string]dsschema.Attribute{
		"branches": dsschema.ListNestedAttribute{
			MarkdownDescription: "Branches",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"bgp_redistribution_profile": dsschema.StringAttribute{
						MarkdownDescription: "BGP redistribution profile",
						Computed:            true,
					},
					"interfaces": dsschema.ListNestedAttribute{
						MarkdownDescription: "Interfaces",
						Computed:            true,
						NestedObject: dsschema.NestedAttributeObject{
							Attributes: map[string]dsschema.Attribute{
								"dhcp_ip": dsschema.StringAttribute{
									MarkdownDescription: "DHCP IP",
									Computed:            true,
								},
								"name": dsschema.StringAttribute{
									MarkdownDescription: "Ethernet interface",
									Computed:            true,
								},
								"sdwan_link_settings": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Sdwan link settings",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"sdwan_gateway": dsschema.StringAttribute{
											MarkdownDescription: "Next hop gateway",
											Computed:            true,
										},
										"sdwan_interface_profile": dsschema.StringAttribute{
											MarkdownDescription: "SD-WAN interface profile",
											Computed:            true,
										},
										"upstream_nat": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Upstream nat",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"enable": dsschema.BoolAttribute{
													MarkdownDescription: "Upstream NAT?",
													Computed:            true,
												},
												"static_ip": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Static ip",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"fqdn": dsschema.StringAttribute{
															MarkdownDescription: "FQDN",
															Computed:            true,
														},
														"ip_address": dsschema.StringAttribute{
															MarkdownDescription: "IP address",
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
					"logical_router": dsschema.StringAttribute{
						MarkdownDescription: "Router",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "Branch firewall serial number",
						Computed:            true,
					},
					"private_interfaces": dsschema.ListNestedAttribute{
						MarkdownDescription: "Private interfaces",
						Computed:            true,
						NestedObject: dsschema.NestedAttributeObject{
							Attributes: map[string]dsschema.Attribute{
								"name": dsschema.StringAttribute{
									MarkdownDescription: "Ethernet interface",
									Computed:            true,
								},
								"sdwan_link_settings": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Sdwan link settings",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"sdwan_gateway": dsschema.StringAttribute{
											MarkdownDescription: "Next hop gateway",
											Computed:            true,
										},
										"sdwan_interface_profile": dsschema.StringAttribute{
											MarkdownDescription: "SD-WAN interface profile",
											Computed:            true,
										},
										"upstream_nat": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Upstream nat",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"enable": dsschema.BoolAttribute{
													MarkdownDescription: "Upstream NAT?",
													Computed:            true,
												},
												"static_ip": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Static ip",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"fqdn": dsschema.StringAttribute{
															MarkdownDescription: "FQDN",
															Computed:            true,
														},
														"ip_address": dsschema.StringAttribute{
															MarkdownDescription: "IP address",
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
					"site": dsschema.StringAttribute{
						MarkdownDescription: "Site name",
						Computed:            true,
					},
				},
			},
		},
		"enable_mesh_between_hubs": dsschema.BoolAttribute{
			MarkdownDescription: "Enable mesh between hubs?",
			Computed:            true,
		},
		"enable_mesh_interconnect": dsschema.BoolAttribute{
			MarkdownDescription: "Enable mesh interconnect?",
			Computed:            true,
		},
		"enable_sdwan": dsschema.BoolAttribute{
			MarkdownDescription: "Enable SD-WAN?",
			Computed:            true,
		},
		"gateways": dsschema.ListNestedAttribute{
			MarkdownDescription: "Hubs",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"allow_dia_vpn_failover": dsschema.BoolAttribute{
						MarkdownDescription: "Allow DIA to VPN failover on branch device for the hub?",
						Computed:            true,
					},
					"bgp_redistribution_profile": dsschema.StringAttribute{
						MarkdownDescription: "BGP redistribution file",
						Computed:            true,
					},
					"interfaces": dsschema.ListNestedAttribute{
						MarkdownDescription: "Interfaces",
						Computed:            true,
						NestedObject: dsschema.NestedAttributeObject{
							Attributes: map[string]dsschema.Attribute{
								"dhcp_ip": dsschema.StringAttribute{
									MarkdownDescription: "DHCP IP",
									Computed:            true,
								},
								"name": dsschema.StringAttribute{
									MarkdownDescription: "Ethernet interface",
									Computed:            true,
								},
								"sdwan_link_settings": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Sdwan link settings",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"sdwan_gateway": dsschema.StringAttribute{
											MarkdownDescription: "Next hop gateway",
											Computed:            true,
										},
										"sdwan_interface_profile": dsschema.StringAttribute{
											MarkdownDescription: "SD-WAN interface profile",
											Computed:            true,
										},
										"upstream_nat": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Upstream nat",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"enable": dsschema.BoolAttribute{
													MarkdownDescription: "Upstream NAT?",
													Computed:            true,
												},
												"static_ip": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Static ip",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"fqdn": dsschema.StringAttribute{
															MarkdownDescription: "FQDN",
															Computed:            true,
														},
														"ip_address": dsschema.StringAttribute{
															MarkdownDescription: "IP address",
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
					"logical_router": dsschema.StringAttribute{
						MarkdownDescription: "Router",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "Hub firewall serial number",
						Computed:            true,
					},
					"priority": dsschema.StringAttribute{
						MarkdownDescription: "Priority",
						Computed:            true,
					},
					"private_interfaces": dsschema.ListNestedAttribute{
						MarkdownDescription: "Private interfaces",
						Computed:            true,
						NestedObject: dsschema.NestedAttributeObject{
							Attributes: map[string]dsschema.Attribute{
								"name": dsschema.StringAttribute{
									MarkdownDescription: "Ethernet interface",
									Computed:            true,
								},
								"sdwan_link_settings": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Sdwan link settings",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"sdwan_gateway": dsschema.StringAttribute{
											MarkdownDescription: "Next hop gateway",
											Computed:            true,
										},
										"sdwan_interface_profile": dsschema.StringAttribute{
											MarkdownDescription: "SD-WAN interface profile",
											Computed:            true,
										},
										"upstream_nat": dsschema.SingleNestedAttribute{
											MarkdownDescription: "Upstream nat",
											Computed:            true,
											Attributes: map[string]dsschema.Attribute{
												"enable": dsschema.BoolAttribute{
													MarkdownDescription: "Upstream NAT?",
													Computed:            true,
												},
												"static_ip": dsschema.SingleNestedAttribute{
													MarkdownDescription: "Static ip",
													Computed:            true,
													Attributes: map[string]dsschema.Attribute{
														"fqdn": dsschema.StringAttribute{
															MarkdownDescription: "FQDN",
															Computed:            true,
														},
														"ip_address": dsschema.StringAttribute{
															MarkdownDescription: "IP address",
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
					"site": dsschema.StringAttribute{
						MarkdownDescription: "Site name",
						Computed:            true,
					},
				},
			},
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "VPN cluster name",
			Optional:            true,
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"type": dsschema.StringAttribute{
			MarkdownDescription: "VPN cluster type",
			Computed:            true,
		},
	},
}

// AutoVpnClustersListModel represents the data model for a list data source.
type AutoVpnClustersListModel struct {
	Tfid    types.String      `tfsdk:"tfid"`
	Data    []AutoVpnClusters `tfsdk:"data"`
	Limit   types.Int64       `tfsdk:"limit"`
	Offset  types.Int64       `tfsdk:"offset"`
	Name    types.String      `tfsdk:"name"`
	Total   types.Int64       `tfsdk:"total"`
	Folder  types.String      `tfsdk:"folder"`
	Snippet types.String      `tfsdk:"snippet"`
	Device  types.String      `tfsdk:"device"`
}

// AutoVpnClustersListDataSourceSchema defines the schema for a list data source.
var AutoVpnClustersListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: AutoVpnClustersDataSourceSchema.Attributes,
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
