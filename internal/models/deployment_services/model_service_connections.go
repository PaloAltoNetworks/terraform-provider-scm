package models

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: deployment_services
// This file contains models for the deployment_services SDK package

// ServiceConnections represents the Terraform model for ServiceConnections
type ServiceConnections struct {
	Tfid                 types.String          `tfsdk:"tfid"`
	EncryptedValues      basetypes.MapValue    `tfsdk:"encrypted_values"`
	BackupSC             basetypes.StringValue `tfsdk:"backup_sc"`
	BgpPeer              basetypes.ObjectValue `tfsdk:"bgp_peer"`
	Id                   basetypes.StringValue `tfsdk:"id"`
	IpsecTunnel          basetypes.StringValue `tfsdk:"ipsec_tunnel"`
	Name                 basetypes.StringValue `tfsdk:"name"`
	NatPool              basetypes.StringValue `tfsdk:"nat_pool"`
	NoExportCommunity    basetypes.StringValue `tfsdk:"no_export_community"`
	OnboardingType       basetypes.StringValue `tfsdk:"onboarding_type"`
	Protocol             basetypes.ObjectValue `tfsdk:"protocol"`
	Qos                  basetypes.ObjectValue `tfsdk:"qos"`
	Region               basetypes.StringValue `tfsdk:"region"`
	SecondaryIpsecTunnel basetypes.StringValue `tfsdk:"secondary_ipsec_tunnel"`
	SourceNat            basetypes.BoolValue   `tfsdk:"source_nat"`
	Subnets              basetypes.ListValue   `tfsdk:"subnets"`
	Folder               basetypes.StringValue `tfsdk:"folder"`
}

// ServiceConnectionsBgpPeer represents a nested structure within the ServiceConnections model
type ServiceConnectionsBgpPeer struct {
	LocalIpAddress   basetypes.StringValue `tfsdk:"local_ip_address"`
	LocalIpv6Address basetypes.StringValue `tfsdk:"local_ipv6_address"`
	PeerIpAddress    basetypes.StringValue `tfsdk:"peer_ip_address"`
	PeerIpv6Address  basetypes.StringValue `tfsdk:"peer_ipv6_address"`
	Secret           basetypes.StringValue `tfsdk:"secret"`
}

// ServiceConnectionsProtocol represents a nested structure within the ServiceConnections model
type ServiceConnectionsProtocol struct {
	Bgp basetypes.ObjectValue `tfsdk:"bgp"`
}

// ServiceConnectionsProtocolBgp represents a nested structure within the ServiceConnections model
type ServiceConnectionsProtocolBgp struct {
	DoNotExportRoutes         basetypes.BoolValue   `tfsdk:"do_not_export_routes"`
	Enable                    basetypes.BoolValue   `tfsdk:"enable"`
	FastFailover              basetypes.BoolValue   `tfsdk:"fast_failover"`
	LocalIpAddress            basetypes.StringValue `tfsdk:"local_ip_address"`
	OriginateDefaultRoute     basetypes.BoolValue   `tfsdk:"originate_default_route"`
	PeerAs                    basetypes.StringValue `tfsdk:"peer_as"`
	PeerIpAddress             basetypes.StringValue `tfsdk:"peer_ip_address"`
	Secret                    basetypes.StringValue `tfsdk:"secret"`
	SummarizeMobileUserRoutes basetypes.BoolValue   `tfsdk:"summarize_mobile_user_routes"`
}

// ServiceConnectionsQos represents a nested structure within the ServiceConnections model
type ServiceConnectionsQos struct {
	Enable     basetypes.BoolValue   `tfsdk:"enable"`
	QosProfile basetypes.StringValue `tfsdk:"qos_profile"`
}

// AttrTypes defines the attribute types for the ServiceConnections model.
func (o ServiceConnections) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":             basetypes.StringType{},
		"encrypted_values": basetypes.MapType{ElemType: basetypes.StringType{}},
		"backup_sc":        basetypes.StringType{},
		"bgp_peer": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"local_ip_address":   basetypes.StringType{},
				"local_ipv6_address": basetypes.StringType{},
				"peer_ip_address":    basetypes.StringType{},
				"peer_ipv6_address":  basetypes.StringType{},
				"secret":             basetypes.StringType{},
			},
		},
		"id":                  basetypes.StringType{},
		"ipsec_tunnel":        basetypes.StringType{},
		"name":                basetypes.StringType{},
		"nat_pool":            basetypes.StringType{},
		"no_export_community": basetypes.StringType{},
		"onboarding_type":     basetypes.StringType{},
		"protocol": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"bgp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"do_not_export_routes":         basetypes.BoolType{},
						"enable":                       basetypes.BoolType{},
						"fast_failover":                basetypes.BoolType{},
						"local_ip_address":             basetypes.StringType{},
						"originate_default_route":      basetypes.BoolType{},
						"peer_as":                      basetypes.StringType{},
						"peer_ip_address":              basetypes.StringType{},
						"secret":                       basetypes.StringType{},
						"summarize_mobile_user_routes": basetypes.BoolType{},
					},
				},
			},
		},
		"qos": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable":      basetypes.BoolType{},
				"qos_profile": basetypes.StringType{},
			},
		},
		"region":                 basetypes.StringType{},
		"secondary_ipsec_tunnel": basetypes.StringType{},
		"source_nat":             basetypes.BoolType{},
		"subnets":                basetypes.ListType{ElemType: basetypes.StringType{}},
		"folder":                 basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ServiceConnections objects.
func (o ServiceConnections) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ServiceConnectionsBgpPeer model.
func (o ServiceConnectionsBgpPeer) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"local_ip_address":   basetypes.StringType{},
		"local_ipv6_address": basetypes.StringType{},
		"peer_ip_address":    basetypes.StringType{},
		"peer_ipv6_address":  basetypes.StringType{},
		"secret":             basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ServiceConnectionsBgpPeer objects.
func (o ServiceConnectionsBgpPeer) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ServiceConnectionsProtocol model.
func (o ServiceConnectionsProtocol) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"bgp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"do_not_export_routes":         basetypes.BoolType{},
				"enable":                       basetypes.BoolType{},
				"fast_failover":                basetypes.BoolType{},
				"local_ip_address":             basetypes.StringType{},
				"originate_default_route":      basetypes.BoolType{},
				"peer_as":                      basetypes.StringType{},
				"peer_ip_address":              basetypes.StringType{},
				"secret":                       basetypes.StringType{},
				"summarize_mobile_user_routes": basetypes.BoolType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ServiceConnectionsProtocol objects.
func (o ServiceConnectionsProtocol) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ServiceConnectionsProtocolBgp model.
func (o ServiceConnectionsProtocolBgp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"do_not_export_routes":         basetypes.BoolType{},
		"enable":                       basetypes.BoolType{},
		"fast_failover":                basetypes.BoolType{},
		"local_ip_address":             basetypes.StringType{},
		"originate_default_route":      basetypes.BoolType{},
		"peer_as":                      basetypes.StringType{},
		"peer_ip_address":              basetypes.StringType{},
		"secret":                       basetypes.StringType{},
		"summarize_mobile_user_routes": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of ServiceConnectionsProtocolBgp objects.
func (o ServiceConnectionsProtocolBgp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ServiceConnectionsQos model.
func (o ServiceConnectionsQos) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable":      basetypes.BoolType{},
		"qos_profile": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ServiceConnectionsQos objects.
func (o ServiceConnectionsQos) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// ServiceConnectionsResourceSchema defines the schema for ServiceConnections resource
var ServiceConnectionsResourceSchema = schema.Schema{
	MarkdownDescription: "ServiceConnection resource",
	Attributes: map[string]schema.Attribute{
		"backup_sc": schema.StringAttribute{
			MarkdownDescription: "Backup s c",
			Optional:            true,
		},
		"bgp_peer": schema.SingleNestedAttribute{
			MarkdownDescription: "Bgp peer",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"local_ip_address": schema.StringAttribute{
					MarkdownDescription: "Local ip address",
					Optional:            true,
				},
				"local_ipv6_address": schema.StringAttribute{
					MarkdownDescription: "Local ipv6 address",
					Optional:            true,
				},
				"peer_ip_address": schema.StringAttribute{
					MarkdownDescription: "Peer ip address",
					Optional:            true,
				},
				"peer_ipv6_address": schema.StringAttribute{
					MarkdownDescription: "Peer ipv6 address",
					Optional:            true,
				},
				"secret": schema.StringAttribute{
					MarkdownDescription: "Secret",
					Optional:            true,
					Sensitive:           true,
				},
			},
		},
		"encrypted_values": schema.MapAttribute{
			ElementType:         basetypes.StringType{},
			MarkdownDescription: "Map of sensitive values returned from the API.",
			Computed:            true,
			Sensitive:           true,
		},
		"folder": schema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the service connection",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"ipsec_tunnel": schema.StringAttribute{
			MarkdownDescription: "Ipsec tunnel",
			Required:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the service connection",
			Required:            true,
		},
		"nat_pool": schema.StringAttribute{
			MarkdownDescription: "Nat pool",
			Optional:            true,
		},
		"no_export_community": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("Disabled", "Enabled-In", "Enabled-Out", "Enabled-Both"),
			},
			MarkdownDescription: "No export community",
			Optional:            true,
		},
		"onboarding_type": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("classic"),
			},
			MarkdownDescription: "Onboarding type",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("classic"),
		},
		"protocol": schema.SingleNestedAttribute{
			MarkdownDescription: "Protocol",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"bgp": schema.SingleNestedAttribute{
					MarkdownDescription: "Bgp",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"do_not_export_routes": schema.BoolAttribute{
							MarkdownDescription: "Do not export routes",
							Optional:            true,
						},
						"enable": schema.BoolAttribute{
							MarkdownDescription: "Enable",
							Optional:            true,
						},
						"fast_failover": schema.BoolAttribute{
							MarkdownDescription: "Fast failover",
							Optional:            true,
						},
						"local_ip_address": schema.StringAttribute{
							MarkdownDescription: "Local ip address",
							Optional:            true,
						},
						"originate_default_route": schema.BoolAttribute{
							MarkdownDescription: "Originate default route",
							Optional:            true,
						},
						"peer_as": schema.StringAttribute{
							MarkdownDescription: "Peer as",
							Required:            true,
						},
						"peer_ip_address": schema.StringAttribute{
							MarkdownDescription: "Peer ip address",
							Optional:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Secret",
							Optional:            true,
							Sensitive:           true,
						},
						"summarize_mobile_user_routes": schema.BoolAttribute{
							MarkdownDescription: "Summarize mobile user routes",
							Optional:            true,
						},
					},
				},
			},
		},
		"qos": schema.SingleNestedAttribute{
			MarkdownDescription: "Qos",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Enable",
					Optional:            true,
				},
				"qos_profile": schema.StringAttribute{
					MarkdownDescription: "Qos profile",
					Optional:            true,
				},
			},
		},
		"region": schema.StringAttribute{
			MarkdownDescription: "Region",
			Required:            true,
		},
		"secondary_ipsec_tunnel": schema.StringAttribute{
			MarkdownDescription: "Secondary ipsec tunnel",
			Optional:            true,
		},
		"source_nat": schema.BoolAttribute{
			MarkdownDescription: "Source nat",
			Optional:            true,
		},
		"subnets": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Subnets",
			Optional:            true,
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

// ServiceConnectionsDataSourceSchema defines the schema for ServiceConnections data source
var ServiceConnectionsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "ServiceConnection data source",
	Attributes: map[string]dsschema.Attribute{
		"backup_sc": dsschema.StringAttribute{
			MarkdownDescription: "Backup s c",
			Computed:            true,
		},
		"bgp_peer": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Bgp peer",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"local_ip_address": dsschema.StringAttribute{
					MarkdownDescription: "Local ip address",
					Computed:            true,
				},
				"local_ipv6_address": dsschema.StringAttribute{
					MarkdownDescription: "Local ipv6 address",
					Computed:            true,
				},
				"peer_ip_address": dsschema.StringAttribute{
					MarkdownDescription: "Peer ip address",
					Computed:            true,
				},
				"peer_ipv6_address": dsschema.StringAttribute{
					MarkdownDescription: "Peer ipv6 address",
					Computed:            true,
				},
				"secret": dsschema.StringAttribute{
					MarkdownDescription: "Secret",
					Computed:            true,
					Sensitive:           true,
				},
			},
		},
		"encrypted_values": dsschema.MapAttribute{
			ElementType:         basetypes.StringType{},
			MarkdownDescription: "Map of sensitive values returned from the API.",
			Computed:            true,
			Sensitive:           true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n",
			Optional:            true,
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the service connection",
			Required:            true,
		},
		"ipsec_tunnel": dsschema.StringAttribute{
			MarkdownDescription: "Ipsec tunnel",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the service connection",
			Optional:            true,
			Computed:            true,
		},
		"nat_pool": dsschema.StringAttribute{
			MarkdownDescription: "Nat pool",
			Computed:            true,
		},
		"no_export_community": dsschema.StringAttribute{
			MarkdownDescription: "No export community",
			Computed:            true,
		},
		"onboarding_type": dsschema.StringAttribute{
			MarkdownDescription: "Onboarding type",
			Computed:            true,
		},
		"protocol": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Protocol",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"bgp": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Bgp",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"do_not_export_routes": dsschema.BoolAttribute{
							MarkdownDescription: "Do not export routes",
							Computed:            true,
						},
						"enable": dsschema.BoolAttribute{
							MarkdownDescription: "Enable",
							Computed:            true,
						},
						"fast_failover": dsschema.BoolAttribute{
							MarkdownDescription: "Fast failover",
							Computed:            true,
						},
						"local_ip_address": dsschema.StringAttribute{
							MarkdownDescription: "Local ip address",
							Computed:            true,
						},
						"originate_default_route": dsschema.BoolAttribute{
							MarkdownDescription: "Originate default route",
							Computed:            true,
						},
						"peer_as": dsschema.StringAttribute{
							MarkdownDescription: "Peer as",
							Computed:            true,
						},
						"peer_ip_address": dsschema.StringAttribute{
							MarkdownDescription: "Peer ip address",
							Computed:            true,
						},
						"secret": dsschema.StringAttribute{
							MarkdownDescription: "Secret",
							Computed:            true,
							Sensitive:           true,
						},
						"summarize_mobile_user_routes": dsschema.BoolAttribute{
							MarkdownDescription: "Summarize mobile user routes",
							Computed:            true,
						},
					},
				},
			},
		},
		"qos": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Qos",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"enable": dsschema.BoolAttribute{
					MarkdownDescription: "Enable",
					Computed:            true,
				},
				"qos_profile": dsschema.StringAttribute{
					MarkdownDescription: "Qos profile",
					Computed:            true,
				},
			},
		},
		"region": dsschema.StringAttribute{
			MarkdownDescription: "Region",
			Computed:            true,
		},
		"secondary_ipsec_tunnel": dsschema.StringAttribute{
			MarkdownDescription: "Secondary ipsec tunnel",
			Computed:            true,
		},
		"source_nat": dsschema.BoolAttribute{
			MarkdownDescription: "Source nat",
			Computed:            true,
		},
		"subnets": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Subnets",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// ServiceConnectionsListModel represents the data model for a list data source.
type ServiceConnectionsListModel struct {
	Tfid    types.String         `tfsdk:"tfid"`
	Data    []ServiceConnections `tfsdk:"data"`
	Limit   types.Int64          `tfsdk:"limit"`
	Offset  types.Int64          `tfsdk:"offset"`
	Name    types.String         `tfsdk:"name"`
	Total   types.Int64          `tfsdk:"total"`
	Folder  types.String         `tfsdk:"folder"`
	Snippet types.String         `tfsdk:"snippet"`
	Device  types.String         `tfsdk:"device"`
}

// ServiceConnectionsListDataSourceSchema defines the schema for a list data source.
var ServiceConnectionsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: ServiceConnectionsDataSourceSchema.Attributes,
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
