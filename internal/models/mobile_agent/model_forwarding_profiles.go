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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// Package: mobile_agent
// This file contains models for the mobile_agent SDK package

// ForwardingProfiles represents the Terraform model for ForwardingProfiles
type ForwardingProfiles struct {
	Tfid             types.String          `tfsdk:"tfid"`
	DefinitionMethod basetypes.StringValue `tfsdk:"definition_method"`
	Description      basetypes.StringValue `tfsdk:"description"`
	Id               basetypes.StringValue `tfsdk:"id"`
	Name             basetypes.StringValue `tfsdk:"name"`
	Type             basetypes.ObjectValue `tfsdk:"type"`
	Folder           basetypes.StringValue `tfsdk:"folder"`
}

// ForwardingProfilesType represents a nested structure within the ForwardingProfiles model
type ForwardingProfilesType struct {
	GlobalProtectProxy basetypes.ObjectValue `tfsdk:"global_protect_proxy"`
	PacFile            basetypes.ObjectValue `tfsdk:"pac_file"`
	ZtnaAgent          basetypes.ObjectValue `tfsdk:"ztna_agent"`
}

// ForwardingProfileGlobalProtectProxyGlobalProtectProxy represents a nested structure within the ForwardingProfiles model
type ForwardingProfileGlobalProtectProxyGlobalProtectProxy struct {
	BlockRule       basetypes.ObjectValue `tfsdk:"block_rule"`
	ForwardingRules basetypes.ListValue   `tfsdk:"forwarding_rules"`
	PacUpload       basetypes.BoolValue   `tfsdk:"pac_upload"`
}

// BlockRuleBasic represents a nested structure within the ForwardingProfiles model
type BlockRuleBasic struct {
	AllowTcp basetypes.ObjectValue `tfsdk:"allow_tcp"`
	AllowUdp basetypes.ObjectValue `tfsdk:"allow_udp"`
	Enable   basetypes.BoolValue   `tfsdk:"enable"`
}

// BlockRuleBasicAllowTcp represents a nested structure within the ForwardingProfiles model
type BlockRuleBasicAllowTcp struct {
	EnableLocations basetypes.BoolValue `tfsdk:"enable_locations"`
	Locations       basetypes.ListValue `tfsdk:"locations"`
}

// BlockRuleBasicAllowUdp represents a nested structure within the ForwardingProfiles model
type BlockRuleBasicAllowUdp struct {
	Destinations       basetypes.StringValue `tfsdk:"destinations"`
	EnableDestinations basetypes.BoolValue   `tfsdk:"enable_destinations"`
	EnableLocations    basetypes.BoolValue   `tfsdk:"enable_locations"`
	Locations          basetypes.ListValue   `tfsdk:"locations"`
}

// ForwardingRuleBasic represents a nested structure within the ForwardingProfiles model
type ForwardingRuleBasic struct {
	Connectivity  basetypes.StringValue `tfsdk:"connectivity"`
	Destinations  basetypes.StringValue `tfsdk:"destinations"`
	Enabled       basetypes.BoolValue   `tfsdk:"enabled"`
	Name          basetypes.StringValue `tfsdk:"name"`
	UserLocations basetypes.StringValue `tfsdk:"user_locations"`
}

// ForwardingProfilePacFilePacFile represents a nested structure within the ForwardingProfiles model
type ForwardingProfilePacFilePacFile struct {
	BlockRule       basetypes.ObjectValue `tfsdk:"block_rule"`
	ForwardingRules basetypes.ListValue   `tfsdk:"forwarding_rules"`
	PacUpload       basetypes.BoolValue   `tfsdk:"pac_upload"`
}

// ForwardingProfileZtnaAgentZtnaAgent represents a nested structure within the ForwardingProfiles model
type ForwardingProfileZtnaAgentZtnaAgent struct {
	BlockRule       basetypes.ObjectValue `tfsdk:"block_rule"`
	ForwardingRules basetypes.ListValue   `tfsdk:"forwarding_rules"`
	PacUpload       basetypes.BoolValue   `tfsdk:"pac_upload"`
}

// BlockRuleZtna represents a nested structure within the ForwardingProfiles model
type BlockRuleZtna struct {
	AllowIcmpForTroubleshooting                       basetypes.BoolValue `tfsdk:"allow_icmp_for_troubleshooting"`
	BlockAllOtherUnmatchedOutboundConnections         basetypes.BoolValue `tfsdk:"block_all_other_unmatched_outbound_connections"`
	BlockInboundAccessWhenConnectedToTunnel           basetypes.BoolValue `tfsdk:"block_inbound_access_when_connected_to_tunnel"`
	BlockNonTcpNonUdpTrafficWhenConnectedToTunnel     basetypes.BoolValue `tfsdk:"block_non_tcp_non_udp_traffic_when_connected_to_tunnel"`
	BlockOutboundLanAccessWhenConnectedToTunnel       basetypes.BoolValue `tfsdk:"block_outbound_lan_access_when_connected_to_tunnel"`
	EnforcerFqdnDnsResolutionViaDnsServers            basetypes.BoolValue `tfsdk:"enforcer_fqdn_dns_resolution_via_dns_servers"`
	ResolveAllFqdnsUsingDnsServersAssignedByTheTunnel basetypes.BoolValue `tfsdk:"resolve_all_fqdns_using_dns_servers_assigned_by_the_tunnel"`
}

// ForwardingRuleZtna represents a nested structure within the ForwardingProfiles model
type ForwardingRuleZtna struct {
	Connectivity       basetypes.StringValue `tfsdk:"connectivity"`
	Destinations       basetypes.StringValue `tfsdk:"destinations"`
	Enabled            basetypes.BoolValue   `tfsdk:"enabled"`
	Name               basetypes.StringValue `tfsdk:"name"`
	SourceApplications basetypes.StringValue `tfsdk:"source_applications"`
	TrafficType        basetypes.StringValue `tfsdk:"traffic_type"`
	UserLocations      basetypes.StringValue `tfsdk:"user_locations"`
}

// AttrTypes defines the attribute types for the ForwardingProfiles model.
func (o ForwardingProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":              basetypes.StringType{},
		"definition_method": basetypes.StringType{},
		"description":       basetypes.StringType{},
		"id":                basetypes.StringType{},
		"name":              basetypes.StringType{},
		"type": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"global_protect_proxy": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"block_rule": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"allow_tcp": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"enable_locations": basetypes.BoolType{},
										"locations":        basetypes.ListType{ElemType: basetypes.StringType{}},
									},
								},
								"allow_udp": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"destinations":        basetypes.StringType{},
										"enable_destinations": basetypes.BoolType{},
										"enable_locations":    basetypes.BoolType{},
										"locations":           basetypes.ListType{ElemType: basetypes.StringType{}},
									},
								},
								"enable": basetypes.BoolType{},
							},
						},
						"forwarding_rules": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"connectivity":   basetypes.StringType{},
								"destinations":   basetypes.StringType{},
								"enabled":        basetypes.BoolType{},
								"name":           basetypes.StringType{},
								"user_locations": basetypes.StringType{},
							},
						}},
						"pac_upload": basetypes.BoolType{},
					},
				},
				"pac_file": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"block_rule": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"allow_tcp": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"enable_locations": basetypes.BoolType{},
										"locations":        basetypes.ListType{ElemType: basetypes.StringType{}},
									},
								},
								"allow_udp": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"destinations":        basetypes.StringType{},
										"enable_destinations": basetypes.BoolType{},
										"enable_locations":    basetypes.BoolType{},
										"locations":           basetypes.ListType{ElemType: basetypes.StringType{}},
									},
								},
								"enable": basetypes.BoolType{},
							},
						},
						"forwarding_rules": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"connectivity":   basetypes.StringType{},
								"destinations":   basetypes.StringType{},
								"enabled":        basetypes.BoolType{},
								"name":           basetypes.StringType{},
								"user_locations": basetypes.StringType{},
							},
						}},
						"pac_upload": basetypes.BoolType{},
					},
				},
				"ztna_agent": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"block_rule": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"allow_icmp_for_troubleshooting":                             basetypes.BoolType{},
								"block_all_other_unmatched_outbound_connections":             basetypes.BoolType{},
								"block_inbound_access_when_connected_to_tunnel":              basetypes.BoolType{},
								"block_non_tcp_non_udp_traffic_when_connected_to_tunnel":     basetypes.BoolType{},
								"block_outbound_lan_access_when_connected_to_tunnel":         basetypes.BoolType{},
								"enforcer_fqdn_dns_resolution_via_dns_servers":               basetypes.BoolType{},
								"resolve_all_fqdns_using_dns_servers_assigned_by_the_tunnel": basetypes.BoolType{},
							},
						},
						"forwarding_rules": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"connectivity":        basetypes.StringType{},
								"destinations":        basetypes.StringType{},
								"enabled":             basetypes.BoolType{},
								"name":                basetypes.StringType{},
								"source_applications": basetypes.StringType{},
								"traffic_type":        basetypes.StringType{},
								"user_locations":      basetypes.StringType{},
							},
						}},
						"pac_upload": basetypes.BoolType{},
					},
				},
			},
		},
		"folder": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ForwardingProfiles objects.
func (o ForwardingProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ForwardingProfilesType model.
func (o ForwardingProfilesType) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"global_protect_proxy": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"block_rule": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"allow_tcp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable_locations": basetypes.BoolType{},
								"locations":        basetypes.ListType{ElemType: basetypes.StringType{}},
							},
						},
						"allow_udp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"destinations":        basetypes.StringType{},
								"enable_destinations": basetypes.BoolType{},
								"enable_locations":    basetypes.BoolType{},
								"locations":           basetypes.ListType{ElemType: basetypes.StringType{}},
							},
						},
						"enable": basetypes.BoolType{},
					},
				},
				"forwarding_rules": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"connectivity":   basetypes.StringType{},
						"destinations":   basetypes.StringType{},
						"enabled":        basetypes.BoolType{},
						"name":           basetypes.StringType{},
						"user_locations": basetypes.StringType{},
					},
				}},
				"pac_upload": basetypes.BoolType{},
			},
		},
		"pac_file": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"block_rule": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"allow_tcp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"enable_locations": basetypes.BoolType{},
								"locations":        basetypes.ListType{ElemType: basetypes.StringType{}},
							},
						},
						"allow_udp": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"destinations":        basetypes.StringType{},
								"enable_destinations": basetypes.BoolType{},
								"enable_locations":    basetypes.BoolType{},
								"locations":           basetypes.ListType{ElemType: basetypes.StringType{}},
							},
						},
						"enable": basetypes.BoolType{},
					},
				},
				"forwarding_rules": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"connectivity":   basetypes.StringType{},
						"destinations":   basetypes.StringType{},
						"enabled":        basetypes.BoolType{},
						"name":           basetypes.StringType{},
						"user_locations": basetypes.StringType{},
					},
				}},
				"pac_upload": basetypes.BoolType{},
			},
		},
		"ztna_agent": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"block_rule": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"allow_icmp_for_troubleshooting":                             basetypes.BoolType{},
						"block_all_other_unmatched_outbound_connections":             basetypes.BoolType{},
						"block_inbound_access_when_connected_to_tunnel":              basetypes.BoolType{},
						"block_non_tcp_non_udp_traffic_when_connected_to_tunnel":     basetypes.BoolType{},
						"block_outbound_lan_access_when_connected_to_tunnel":         basetypes.BoolType{},
						"enforcer_fqdn_dns_resolution_via_dns_servers":               basetypes.BoolType{},
						"resolve_all_fqdns_using_dns_servers_assigned_by_the_tunnel": basetypes.BoolType{},
					},
				},
				"forwarding_rules": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"connectivity":        basetypes.StringType{},
						"destinations":        basetypes.StringType{},
						"enabled":             basetypes.BoolType{},
						"name":                basetypes.StringType{},
						"source_applications": basetypes.StringType{},
						"traffic_type":        basetypes.StringType{},
						"user_locations":      basetypes.StringType{},
					},
				}},
				"pac_upload": basetypes.BoolType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ForwardingProfilesType objects.
func (o ForwardingProfilesType) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ForwardingProfileGlobalProtectProxyGlobalProtectProxy model.
func (o ForwardingProfileGlobalProtectProxyGlobalProtectProxy) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"block_rule": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"allow_tcp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable_locations": basetypes.BoolType{},
						"locations":        basetypes.ListType{ElemType: basetypes.StringType{}},
					},
				},
				"allow_udp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"destinations":        basetypes.StringType{},
						"enable_destinations": basetypes.BoolType{},
						"enable_locations":    basetypes.BoolType{},
						"locations":           basetypes.ListType{ElemType: basetypes.StringType{}},
					},
				},
				"enable": basetypes.BoolType{},
			},
		},
		"forwarding_rules": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"connectivity":   basetypes.StringType{},
				"destinations":   basetypes.StringType{},
				"enabled":        basetypes.BoolType{},
				"name":           basetypes.StringType{},
				"user_locations": basetypes.StringType{},
			},
		}},
		"pac_upload": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of ForwardingProfileGlobalProtectProxyGlobalProtectProxy objects.
func (o ForwardingProfileGlobalProtectProxyGlobalProtectProxy) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BlockRuleBasic model.
func (o BlockRuleBasic) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"allow_tcp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable_locations": basetypes.BoolType{},
				"locations":        basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		},
		"allow_udp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"destinations":        basetypes.StringType{},
				"enable_destinations": basetypes.BoolType{},
				"enable_locations":    basetypes.BoolType{},
				"locations":           basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		},
		"enable": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of BlockRuleBasic objects.
func (o BlockRuleBasic) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BlockRuleBasicAllowTcp model.
func (o BlockRuleBasicAllowTcp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable_locations": basetypes.BoolType{},
		"locations":        basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of BlockRuleBasicAllowTcp objects.
func (o BlockRuleBasicAllowTcp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BlockRuleBasicAllowUdp model.
func (o BlockRuleBasicAllowUdp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"destinations":        basetypes.StringType{},
		"enable_destinations": basetypes.BoolType{},
		"enable_locations":    basetypes.BoolType{},
		"locations":           basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of BlockRuleBasicAllowUdp objects.
func (o BlockRuleBasicAllowUdp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ForwardingRuleBasic model.
func (o ForwardingRuleBasic) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"connectivity":   basetypes.StringType{},
		"destinations":   basetypes.StringType{},
		"enabled":        basetypes.BoolType{},
		"name":           basetypes.StringType{},
		"user_locations": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ForwardingRuleBasic objects.
func (o ForwardingRuleBasic) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ForwardingProfilePacFilePacFile model.
func (o ForwardingProfilePacFilePacFile) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"block_rule": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"allow_tcp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable_locations": basetypes.BoolType{},
						"locations":        basetypes.ListType{ElemType: basetypes.StringType{}},
					},
				},
				"allow_udp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"destinations":        basetypes.StringType{},
						"enable_destinations": basetypes.BoolType{},
						"enable_locations":    basetypes.BoolType{},
						"locations":           basetypes.ListType{ElemType: basetypes.StringType{}},
					},
				},
				"enable": basetypes.BoolType{},
			},
		},
		"forwarding_rules": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"connectivity":   basetypes.StringType{},
				"destinations":   basetypes.StringType{},
				"enabled":        basetypes.BoolType{},
				"name":           basetypes.StringType{},
				"user_locations": basetypes.StringType{},
			},
		}},
		"pac_upload": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of ForwardingProfilePacFilePacFile objects.
func (o ForwardingProfilePacFilePacFile) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ForwardingProfileZtnaAgentZtnaAgent model.
func (o ForwardingProfileZtnaAgentZtnaAgent) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"block_rule": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"allow_icmp_for_troubleshooting":                             basetypes.BoolType{},
				"block_all_other_unmatched_outbound_connections":             basetypes.BoolType{},
				"block_inbound_access_when_connected_to_tunnel":              basetypes.BoolType{},
				"block_non_tcp_non_udp_traffic_when_connected_to_tunnel":     basetypes.BoolType{},
				"block_outbound_lan_access_when_connected_to_tunnel":         basetypes.BoolType{},
				"enforcer_fqdn_dns_resolution_via_dns_servers":               basetypes.BoolType{},
				"resolve_all_fqdns_using_dns_servers_assigned_by_the_tunnel": basetypes.BoolType{},
			},
		},
		"forwarding_rules": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"connectivity":        basetypes.StringType{},
				"destinations":        basetypes.StringType{},
				"enabled":             basetypes.BoolType{},
				"name":                basetypes.StringType{},
				"source_applications": basetypes.StringType{},
				"traffic_type":        basetypes.StringType{},
				"user_locations":      basetypes.StringType{},
			},
		}},
		"pac_upload": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of ForwardingProfileZtnaAgentZtnaAgent objects.
func (o ForwardingProfileZtnaAgentZtnaAgent) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the BlockRuleZtna model.
func (o BlockRuleZtna) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"allow_icmp_for_troubleshooting":                             basetypes.BoolType{},
		"block_all_other_unmatched_outbound_connections":             basetypes.BoolType{},
		"block_inbound_access_when_connected_to_tunnel":              basetypes.BoolType{},
		"block_non_tcp_non_udp_traffic_when_connected_to_tunnel":     basetypes.BoolType{},
		"block_outbound_lan_access_when_connected_to_tunnel":         basetypes.BoolType{},
		"enforcer_fqdn_dns_resolution_via_dns_servers":               basetypes.BoolType{},
		"resolve_all_fqdns_using_dns_servers_assigned_by_the_tunnel": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of BlockRuleZtna objects.
func (o BlockRuleZtna) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ForwardingRuleZtna model.
func (o ForwardingRuleZtna) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"connectivity":        basetypes.StringType{},
		"destinations":        basetypes.StringType{},
		"enabled":             basetypes.BoolType{},
		"name":                basetypes.StringType{},
		"source_applications": basetypes.StringType{},
		"traffic_type":        basetypes.StringType{},
		"user_locations":      basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ForwardingRuleZtna objects.
func (o ForwardingRuleZtna) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// ForwardingProfilesResourceSchema defines the schema for ForwardingProfiles resource
var ForwardingProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "ForwardingProfile resource",
	Attributes: map[string]schema.Attribute{
		"definition_method": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("rules", "pac-file"),
			},
			MarkdownDescription: "Enable forwarding rule for forwarding profile",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("rules"),
		},
		"description": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(1023),
			},
			MarkdownDescription: "Forwarding profile description",
			Optional:            true,
		},
		"folder": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("Mobile Users"),
				utils.FolderValidator(),
			},
			MarkdownDescription: "The folder in which the resource is defined\n",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the forwarding profile",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[0-9a-zA-Z._ -]+$"), "pattern must match "+"^[0-9a-zA-Z._ -]+$"),
			},
			MarkdownDescription: "forwarding profile name as an alphanumeric string [ 0-9a-zA-Z._ -]",
			Required:            true,
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"type": schema.SingleNestedAttribute{
			MarkdownDescription: "Forwarding profile type configuration (PAC file, GlobalProtect proxy, or ZTNA agent)",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"global_protect_proxy": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("pac_file"),
							path.MatchRelative().AtParent().AtName("ztna_agent"),
						),
					},
					MarkdownDescription: "Global Protect proxy-based forwarding configuration\n\n> ℹ️ **Note:** You must specify exactly one of `global_protect_proxy`, `pac_file`, and `ztna_agent`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"block_rule": schema.SingleNestedAttribute{
							MarkdownDescription: "Basic block rule configuration for PAC file and GlobalProtect proxy profiles",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"allow_tcp": schema.SingleNestedAttribute{
									MarkdownDescription: "TCP traffic allowlist configuration",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"enable_locations": schema.BoolAttribute{
											MarkdownDescription: "Enable locations for allow-tcp",
											Optional:            true,
										},
										"locations": schema.ListAttribute{
											ElementType:         types.StringType,
											MarkdownDescription: "List of user locations allowed for TCP traffic",
											Validators: []validator.List{
												listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(64)),
											},
											Optional: true,
										},
									},
								},
								"allow_udp": schema.SingleNestedAttribute{
									MarkdownDescription: "UDP traffic allowlist configuration with location and destination support",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"destinations": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.LengthAtMost(64),
											},
											MarkdownDescription: "Destination addresses or networks allowed for UDP traffic",
											Optional:            true,
										},
										"enable_destinations": schema.BoolAttribute{
											MarkdownDescription: "Enable destinations for allow-udp",
											Optional:            true,
										},
										"enable_locations": schema.BoolAttribute{
											MarkdownDescription: "Enable locations for allow-udp",
											Optional:            true,
										},
										"locations": schema.ListAttribute{
											ElementType:         types.StringType,
											MarkdownDescription: "List of user locations allowed for UDP traffic",
											Validators: []validator.List{
												listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(64)),
											},
											Optional: true,
										},
									},
								},
								"enable": schema.BoolAttribute{
									MarkdownDescription: "Enable block rule",
									Optional:            true,
								},
							},
						},
						"forwarding_rules": schema.ListNestedAttribute{
							MarkdownDescription: "List of GlobalProtect proxy-based forwarding rules",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"connectivity": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.LengthAtMost(64),
										},
										MarkdownDescription: "Connectivity method for this forwarding rule (e.g. direct)",
										Optional:            true,
										Computed:            true,
										Default:             stringdefault.StaticString("direct"),
									},
									"destinations": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.LengthAtMost(64),
										},
										MarkdownDescription: "Destination scope this forwarding rule applies to",
										Optional:            true,
										Computed:            true,
										Default:             stringdefault.StaticString("Any"),
									},
									"enabled": schema.BoolAttribute{
										MarkdownDescription: "Enable a basic forwarding rule",
										Optional:            true,
										Computed:            true,
										Default:             booldefault.StaticBool(true),
									},
									"name": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.LengthAtMost(64),
											stringvalidator.RegexMatches(regexp.MustCompile("^[0-9a-zA-Z._ -]+$"), "pattern must match "+"^[0-9a-zA-Z._ -]+$"),
										},
										MarkdownDescription: "Basic forwarding rule name as an alphanumeric string [ 0-9a-zA-Z._ -]",
										Required:            true,
									},
									"user_locations": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.LengthAtMost(64),
										},
										MarkdownDescription: "User location scope this rule applies to",
										Optional:            true,
										Computed:            true,
										Default:             stringdefault.StaticString("Any"),
									},
								},
							},
						},
						"pac_upload": schema.BoolAttribute{
							MarkdownDescription: "User uploaded PAC file for Global Protect proxy-based forwarding configuration",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
					},
				},
				"pac_file": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("global_protect_proxy"),
							path.MatchRelative().AtParent().AtName("ztna_agent"),
						),
					},
					MarkdownDescription: "PAC file based forwarding configuration\n\n> ℹ️ **Note:** You must specify exactly one of `global_protect_proxy`, `pac_file`, and `ztna_agent`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"block_rule": schema.SingleNestedAttribute{
							MarkdownDescription: "Basic block rule configuration for PAC file and GlobalProtect proxy profiles",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"allow_tcp": schema.SingleNestedAttribute{
									MarkdownDescription: "TCP traffic allowlist configuration",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"enable_locations": schema.BoolAttribute{
											MarkdownDescription: "Enable locations for allow-tcp",
											Optional:            true,
										},
										"locations": schema.ListAttribute{
											ElementType:         types.StringType,
											MarkdownDescription: "List of user locations allowed for TCP traffic",
											Validators: []validator.List{
												listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(64)),
											},
											Optional: true,
										},
									},
								},
								"allow_udp": schema.SingleNestedAttribute{
									MarkdownDescription: "UDP traffic allowlist configuration with location and destination support",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"destinations": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.LengthAtMost(64),
											},
											MarkdownDescription: "Destination addresses or networks allowed for UDP traffic",
											Optional:            true,
										},
										"enable_destinations": schema.BoolAttribute{
											MarkdownDescription: "Enable destinations for allow-udp",
											Optional:            true,
										},
										"enable_locations": schema.BoolAttribute{
											MarkdownDescription: "Enable locations for allow-udp",
											Optional:            true,
										},
										"locations": schema.ListAttribute{
											ElementType:         types.StringType,
											MarkdownDescription: "List of user locations allowed for UDP traffic",
											Validators: []validator.List{
												listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(64)),
											},
											Optional: true,
										},
									},
								},
								"enable": schema.BoolAttribute{
									MarkdownDescription: "Enable block rule",
									Optional:            true,
								},
							},
						},
						"forwarding_rules": schema.ListNestedAttribute{
							MarkdownDescription: "List of PAC file-based forwarding rules",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"connectivity": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.LengthAtMost(64),
										},
										MarkdownDescription: "Connectivity method for this forwarding rule (e.g. direct)",
										Optional:            true,
										Computed:            true,
										Default:             stringdefault.StaticString("direct"),
									},
									"destinations": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.LengthAtMost(64),
										},
										MarkdownDescription: "Destination scope this forwarding rule applies to",
										Optional:            true,
										Computed:            true,
										Default:             stringdefault.StaticString("Any"),
									},
									"enabled": schema.BoolAttribute{
										MarkdownDescription: "Enable a basic forwarding rule",
										Optional:            true,
										Computed:            true,
										Default:             booldefault.StaticBool(true),
									},
									"name": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.LengthAtMost(64),
											stringvalidator.RegexMatches(regexp.MustCompile("^[0-9a-zA-Z._ -]+$"), "pattern must match "+"^[0-9a-zA-Z._ -]+$"),
										},
										MarkdownDescription: "Basic forwarding rule name as an alphanumeric string [ 0-9a-zA-Z._ -]",
										Required:            true,
									},
									"user_locations": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.LengthAtMost(64),
										},
										MarkdownDescription: "User location scope this rule applies to",
										Optional:            true,
										Computed:            true,
										Default:             stringdefault.StaticString("Any"),
									},
								},
							},
						},
						"pac_upload": schema.BoolAttribute{
							MarkdownDescription: "User upload PAC file for PAC file based forwarding configuration",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
					},
				},
				"ztna_agent": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("global_protect_proxy"),
							path.MatchRelative().AtParent().AtName("pac_file"),
						),
					},
					MarkdownDescription: "ZTNA agent-based forwarding configuration\n\n> ℹ️ **Note:** You must specify exactly one of `global_protect_proxy`, `pac_file`, and `ztna_agent`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"block_rule": schema.SingleNestedAttribute{
							MarkdownDescription: "ZTNA block rule configuration",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"allow_icmp_for_troubleshooting": schema.BoolAttribute{
									MarkdownDescription: "Allow ICMP for troubleshooting",
									Optional:            true,
									Computed:            true,
									Default:             booldefault.StaticBool(false),
								},
								"block_all_other_unmatched_outbound_connections": schema.BoolAttribute{
									MarkdownDescription: "Block all other unmatched outbound connections",
									Optional:            true,
									Computed:            true,
									Default:             booldefault.StaticBool(false),
								},
								"block_inbound_access_when_connected_to_tunnel": schema.BoolAttribute{
									MarkdownDescription: "Block inbound access when connected to tunnel",
									Optional:            true,
									Computed:            true,
									Default:             booldefault.StaticBool(false),
								},
								"block_non_tcp_non_udp_traffic_when_connected_to_tunnel": schema.BoolAttribute{
									MarkdownDescription: "Block Non-TCP Non UDP based traffic when connected to tunnel",
									Optional:            true,
									Computed:            true,
									Default:             booldefault.StaticBool(false),
								},
								"block_outbound_lan_access_when_connected_to_tunnel": schema.BoolAttribute{
									MarkdownDescription: "Block outbound LAN access when connected to tunnel",
									Optional:            true,
									Computed:            true,
									Default:             booldefault.StaticBool(false),
								},
								"enforcer_fqdn_dns_resolution_via_dns_servers": schema.BoolAttribute{
									MarkdownDescription: "Enforce FQDN DNS resolution via tunnel DNS servers",
									Optional:            true,
									Computed:            true,
									Default:             booldefault.StaticBool(true),
								},
								"resolve_all_fqdns_using_dns_servers_assigned_by_the_tunnel": schema.BoolAttribute{
									MarkdownDescription: "Resolve All FQDNs using DNS servers assigned by the tunnel (Windows Only)",
									Optional:            true,
									Computed:            true,
									Default:             booldefault.StaticBool(true),
								},
							},
						},
						"forwarding_rules": schema.ListNestedAttribute{
							MarkdownDescription: "List of ZTNA agent-based forwarding rules",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"connectivity": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.LengthAtMost(64),
										},
										MarkdownDescription: "Connectivity method for this ZTNA forwarding rule (e.g. direct)",
										Optional:            true,
										Computed:            true,
										Default:             stringdefault.StaticString("direct"),
									},
									"destinations": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.LengthAtMost(64),
										},
										MarkdownDescription: "Destination scope this ZTNA forwarding rule applies to",
										Optional:            true,
										Computed:            true,
										Default:             stringdefault.StaticString("Any"),
									},
									"enabled": schema.BoolAttribute{
										MarkdownDescription: "Enable a forwarding rule ztna",
										Optional:            true,
										Computed:            true,
										Default:             booldefault.StaticBool(true),
									},
									"name": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.LengthAtMost(64),
											stringvalidator.RegexMatches(regexp.MustCompile("^[0-9a-zA-Z._ -]+$"), "pattern must match "+"^[0-9a-zA-Z._ -]+$"),
										},
										MarkdownDescription: "Forwarding rule ZTNA name as an alphanumeric string [ 0-9a-zA-Z._ -]",
										Required:            true,
									},
									"source_applications": schema.StringAttribute{
										MarkdownDescription: "Source applications this ZTNA rule applies to",
										Optional:            true,
										Computed:            true,
										Default:             stringdefault.StaticString("Any"),
									},
									"traffic_type": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("dns", "dns-and-network-traffic", "network-traffic"),
										},
										MarkdownDescription: "Type of traffic this ZTNA rule applies to (dns, network, or both)",
										Optional:            true,
										Computed:            true,
										Default:             stringdefault.StaticString("dns"),
									},
									"user_locations": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.LengthAtMost(64),
										},
										MarkdownDescription: "User location scope this ZTNA rule applies to",
										Optional:            true,
										Computed:            true,
										Default:             stringdefault.StaticString("Any"),
									},
								},
							},
						},
						"pac_upload": schema.BoolAttribute{
							MarkdownDescription: "User uploaded PAC file for a ZTNA agent-based forwarding configuration",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
					},
				},
			},
		},
	},
}

// ForwardingProfilesDataSourceSchema defines the schema for ForwardingProfiles data source
var ForwardingProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "ForwardingProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"definition_method": dsschema.StringAttribute{
			MarkdownDescription: "Enable forwarding rule for forwarding profile",
			Computed:            true,
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Forwarding profile description",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n",
			Optional:            true,
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the forwarding profile",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "forwarding profile name as an alphanumeric string [ 0-9a-zA-Z._ -]",
			Optional:            true,
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"type": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Forwarding profile type configuration (PAC file, GlobalProtect proxy, or ZTNA agent)",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"global_protect_proxy": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Global Protect proxy-based forwarding configuration\n\n> ℹ️ **Note:** You must specify exactly one of `global_protect_proxy`, `pac_file`, and `ztna_agent`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"block_rule": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Basic block rule configuration for PAC file and GlobalProtect proxy profiles",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"allow_tcp": dsschema.SingleNestedAttribute{
									MarkdownDescription: "TCP traffic allowlist configuration",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"enable_locations": dsschema.BoolAttribute{
											MarkdownDescription: "Enable locations for allow-tcp",
											Computed:            true,
										},
										"locations": dsschema.ListAttribute{
											ElementType:         types.StringType,
											MarkdownDescription: "List of user locations allowed for TCP traffic",
											Computed:            true,
										},
									},
								},
								"allow_udp": dsschema.SingleNestedAttribute{
									MarkdownDescription: "UDP traffic allowlist configuration with location and destination support",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"destinations": dsschema.StringAttribute{
											MarkdownDescription: "Destination addresses or networks allowed for UDP traffic",
											Computed:            true,
										},
										"enable_destinations": dsschema.BoolAttribute{
											MarkdownDescription: "Enable destinations for allow-udp",
											Computed:            true,
										},
										"enable_locations": dsschema.BoolAttribute{
											MarkdownDescription: "Enable locations for allow-udp",
											Computed:            true,
										},
										"locations": dsschema.ListAttribute{
											ElementType:         types.StringType,
											MarkdownDescription: "List of user locations allowed for UDP traffic",
											Computed:            true,
										},
									},
								},
								"enable": dsschema.BoolAttribute{
									MarkdownDescription: "Enable block rule",
									Computed:            true,
								},
							},
						},
						"forwarding_rules": dsschema.ListNestedAttribute{
							MarkdownDescription: "List of GlobalProtect proxy-based forwarding rules",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"connectivity": dsschema.StringAttribute{
										MarkdownDescription: "Connectivity method for this forwarding rule (e.g. direct)",
										Computed:            true,
									},
									"destinations": dsschema.StringAttribute{
										MarkdownDescription: "Destination scope this forwarding rule applies to",
										Computed:            true,
									},
									"enabled": dsschema.BoolAttribute{
										MarkdownDescription: "Enable a basic forwarding rule",
										Computed:            true,
									},
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Basic forwarding rule name as an alphanumeric string [ 0-9a-zA-Z._ -]",
										Computed:            true,
									},
									"user_locations": dsschema.StringAttribute{
										MarkdownDescription: "User location scope this rule applies to",
										Computed:            true,
									},
								},
							},
						},
						"pac_upload": dsschema.BoolAttribute{
							MarkdownDescription: "User uploaded PAC file for Global Protect proxy-based forwarding configuration",
							Computed:            true,
						},
					},
				},
				"pac_file": dsschema.SingleNestedAttribute{
					MarkdownDescription: "PAC file based forwarding configuration\n\n> ℹ️ **Note:** You must specify exactly one of `global_protect_proxy`, `pac_file`, and `ztna_agent`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"block_rule": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Basic block rule configuration for PAC file and GlobalProtect proxy profiles",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"allow_tcp": dsschema.SingleNestedAttribute{
									MarkdownDescription: "TCP traffic allowlist configuration",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"enable_locations": dsschema.BoolAttribute{
											MarkdownDescription: "Enable locations for allow-tcp",
											Computed:            true,
										},
										"locations": dsschema.ListAttribute{
											ElementType:         types.StringType,
											MarkdownDescription: "List of user locations allowed for TCP traffic",
											Computed:            true,
										},
									},
								},
								"allow_udp": dsschema.SingleNestedAttribute{
									MarkdownDescription: "UDP traffic allowlist configuration with location and destination support",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"destinations": dsschema.StringAttribute{
											MarkdownDescription: "Destination addresses or networks allowed for UDP traffic",
											Computed:            true,
										},
										"enable_destinations": dsschema.BoolAttribute{
											MarkdownDescription: "Enable destinations for allow-udp",
											Computed:            true,
										},
										"enable_locations": dsschema.BoolAttribute{
											MarkdownDescription: "Enable locations for allow-udp",
											Computed:            true,
										},
										"locations": dsschema.ListAttribute{
											ElementType:         types.StringType,
											MarkdownDescription: "List of user locations allowed for UDP traffic",
											Computed:            true,
										},
									},
								},
								"enable": dsschema.BoolAttribute{
									MarkdownDescription: "Enable block rule",
									Computed:            true,
								},
							},
						},
						"forwarding_rules": dsschema.ListNestedAttribute{
							MarkdownDescription: "List of PAC file-based forwarding rules",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"connectivity": dsschema.StringAttribute{
										MarkdownDescription: "Connectivity method for this forwarding rule (e.g. direct)",
										Computed:            true,
									},
									"destinations": dsschema.StringAttribute{
										MarkdownDescription: "Destination scope this forwarding rule applies to",
										Computed:            true,
									},
									"enabled": dsschema.BoolAttribute{
										MarkdownDescription: "Enable a basic forwarding rule",
										Computed:            true,
									},
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Basic forwarding rule name as an alphanumeric string [ 0-9a-zA-Z._ -]",
										Computed:            true,
									},
									"user_locations": dsschema.StringAttribute{
										MarkdownDescription: "User location scope this rule applies to",
										Computed:            true,
									},
								},
							},
						},
						"pac_upload": dsschema.BoolAttribute{
							MarkdownDescription: "User upload PAC file for PAC file based forwarding configuration",
							Computed:            true,
						},
					},
				},
				"ztna_agent": dsschema.SingleNestedAttribute{
					MarkdownDescription: "ZTNA agent-based forwarding configuration\n\n> ℹ️ **Note:** You must specify exactly one of `global_protect_proxy`, `pac_file`, and `ztna_agent`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"block_rule": dsschema.SingleNestedAttribute{
							MarkdownDescription: "ZTNA block rule configuration",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"allow_icmp_for_troubleshooting": dsschema.BoolAttribute{
									MarkdownDescription: "Allow ICMP for troubleshooting",
									Computed:            true,
								},
								"block_all_other_unmatched_outbound_connections": dsschema.BoolAttribute{
									MarkdownDescription: "Block all other unmatched outbound connections",
									Computed:            true,
								},
								"block_inbound_access_when_connected_to_tunnel": dsschema.BoolAttribute{
									MarkdownDescription: "Block inbound access when connected to tunnel",
									Computed:            true,
								},
								"block_non_tcp_non_udp_traffic_when_connected_to_tunnel": dsschema.BoolAttribute{
									MarkdownDescription: "Block Non-TCP Non UDP based traffic when connected to tunnel",
									Computed:            true,
								},
								"block_outbound_lan_access_when_connected_to_tunnel": dsschema.BoolAttribute{
									MarkdownDescription: "Block outbound LAN access when connected to tunnel",
									Computed:            true,
								},
								"enforcer_fqdn_dns_resolution_via_dns_servers": dsschema.BoolAttribute{
									MarkdownDescription: "Enforce FQDN DNS resolution via tunnel DNS servers",
									Computed:            true,
								},
								"resolve_all_fqdns_using_dns_servers_assigned_by_the_tunnel": dsschema.BoolAttribute{
									MarkdownDescription: "Resolve All FQDNs using DNS servers assigned by the tunnel (Windows Only)",
									Computed:            true,
								},
							},
						},
						"forwarding_rules": dsschema.ListNestedAttribute{
							MarkdownDescription: "List of ZTNA agent-based forwarding rules",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"connectivity": dsschema.StringAttribute{
										MarkdownDescription: "Connectivity method for this ZTNA forwarding rule (e.g. direct)",
										Computed:            true,
									},
									"destinations": dsschema.StringAttribute{
										MarkdownDescription: "Destination scope this ZTNA forwarding rule applies to",
										Computed:            true,
									},
									"enabled": dsschema.BoolAttribute{
										MarkdownDescription: "Enable a forwarding rule ztna",
										Computed:            true,
									},
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Forwarding rule ZTNA name as an alphanumeric string [ 0-9a-zA-Z._ -]",
										Computed:            true,
									},
									"source_applications": dsschema.StringAttribute{
										MarkdownDescription: "Source applications this ZTNA rule applies to",
										Computed:            true,
									},
									"traffic_type": dsschema.StringAttribute{
										MarkdownDescription: "Type of traffic this ZTNA rule applies to (dns, network, or both)",
										Computed:            true,
									},
									"user_locations": dsschema.StringAttribute{
										MarkdownDescription: "User location scope this ZTNA rule applies to",
										Computed:            true,
									},
								},
							},
						},
						"pac_upload": dsschema.BoolAttribute{
							MarkdownDescription: "User uploaded PAC file for a ZTNA agent-based forwarding configuration",
							Computed:            true,
						},
					},
				},
			},
		},
	},
}

// ForwardingProfilesListModel represents the data model for a list data source.
type ForwardingProfilesListModel struct {
	Tfid    types.String         `tfsdk:"tfid"`
	Data    []ForwardingProfiles `tfsdk:"data"`
	Limit   types.Int64          `tfsdk:"limit"`
	Offset  types.Int64          `tfsdk:"offset"`
	Name    types.String         `tfsdk:"name"`
	Total   types.Int64          `tfsdk:"total"`
	Folder  types.String         `tfsdk:"folder"`
	Snippet types.String         `tfsdk:"snippet"`
	Device  types.String         `tfsdk:"device"`
}

// ForwardingProfilesListDataSourceSchema defines the schema for a list data source.
var ForwardingProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: ForwardingProfilesDataSourceSchema.Attributes,
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
