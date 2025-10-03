package models

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: deployment_services
// This file contains models for the deployment_services SDK package

// SharedInfrastructureSettings represents the Terraform model for SharedInfrastructureSettings
type SharedInfrastructureSettings struct {
	Tfid                           types.String          `tfsdk:"tfid"`
	ApiKey                         basetypes.StringValue `tfsdk:"api_key"`
	CaptivePortalRedirectIpAddress basetypes.StringValue `tfsdk:"captive_portal_redirect_ip_address"`
	ConnectorApplicationBlocks     basetypes.ObjectValue `tfsdk:"connector_application_blocks"`
	ConnectorConnectorBlocks       basetypes.ObjectValue `tfsdk:"connector_connector_blocks"`
	EgressIpNotificationUrl        basetypes.StringValue `tfsdk:"egress_ip_notification_url"`
	Folder                         basetypes.StringValue `tfsdk:"folder"`
	InfraBgpAs                     basetypes.StringValue `tfsdk:"infra_bgp_as"`
	InfrastructureSubnet           basetypes.StringValue `tfsdk:"infrastructure_subnet"`
	InfrastructureSubnetIpv6       basetypes.StringValue `tfsdk:"infrastructure_subnet_ipv6"`
	Ipv6                           basetypes.BoolValue   `tfsdk:"ipv6"`
	LoopbackIps                    basetypes.ListValue   `tfsdk:"loopback_ips"`
	TunnelMonitorIpAddress         basetypes.StringValue `tfsdk:"tunnel_monitor_ip_address"`
}

// EditSharedInfrastructureSettingsConnectorApplicationBlocks represents a nested structure within the SharedInfrastructureSettings model
type EditSharedInfrastructureSettingsConnectorApplicationBlocks struct {
	Member basetypes.ListValue `tfsdk:"member"`
}

// EditSharedInfrastructureSettingsConnectorConnectorBlocks represents a nested structure within the SharedInfrastructureSettings model
type EditSharedInfrastructureSettingsConnectorConnectorBlocks struct {
	Member basetypes.ListValue `tfsdk:"member"`
}

// AttrTypes defines the attribute types for the SharedInfrastructureSettings model.
func (o SharedInfrastructureSettings) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":                               basetypes.StringType{},
		"api_key":                            basetypes.StringType{},
		"captive_portal_redirect_ip_address": basetypes.StringType{},
		"connector_application_blocks": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"member": basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		},
		"connector_connector_blocks": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"member": basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		},
		"egress_ip_notification_url": basetypes.StringType{},
		"folder":                     basetypes.StringType{},
		"infra_bgp_as":               basetypes.StringType{},
		"infrastructure_subnet":      basetypes.StringType{},
		"infrastructure_subnet_ipv6": basetypes.StringType{},
		"ipv6":                       basetypes.BoolType{},
		"loopback_ips":               basetypes.ListType{ElemType: basetypes.StringType{}},
		"tunnel_monitor_ip_address":  basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of SharedInfrastructureSettings objects.
func (o SharedInfrastructureSettings) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the EditSharedInfrastructureSettingsConnectorApplicationBlocks model.
func (o EditSharedInfrastructureSettingsConnectorApplicationBlocks) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"member": basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of EditSharedInfrastructureSettingsConnectorApplicationBlocks objects.
func (o EditSharedInfrastructureSettingsConnectorApplicationBlocks) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the EditSharedInfrastructureSettingsConnectorConnectorBlocks model.
func (o EditSharedInfrastructureSettingsConnectorConnectorBlocks) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"member": basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of EditSharedInfrastructureSettingsConnectorConnectorBlocks objects.
func (o EditSharedInfrastructureSettingsConnectorConnectorBlocks) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// SharedInfrastructureSettingsResourceSchema defines the schema for SharedInfrastructureSettings resource
var SharedInfrastructureSettingsResourceSchema = schema.Schema{
	MarkdownDescription: "SharedInfrastructureSetting resource",
	Attributes: map[string]schema.Attribute{
		"api_key": schema.StringAttribute{
			MarkdownDescription: "Api key",
			Optional:            true,
		},
		"captive_portal_redirect_ip_address": schema.StringAttribute{
			MarkdownDescription: "Captive portal redirect ip address",
			Optional:            true,
		},
		"connector_application_blocks": schema.SingleNestedAttribute{
			MarkdownDescription: "Connector application blocks",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"member": schema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Member",
					Optional:            true,
				},
			},
		},
		"connector_connector_blocks": schema.SingleNestedAttribute{
			MarkdownDescription: "Connector connector blocks",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"member": schema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Member",
					Optional:            true,
				},
			},
		},
		"egress_ip_notification_url": schema.StringAttribute{
			MarkdownDescription: "Egress ip notification url",
			Optional:            true,
		},
		"folder": schema.StringAttribute{
			MarkdownDescription: "The folder containing the shared infrastructure settings",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("Shared"),
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"infra_bgp_as": schema.StringAttribute{
			MarkdownDescription: "Infra bgp as",
			Optional:            true,
		},
		"infrastructure_subnet": schema.StringAttribute{
			MarkdownDescription: "Infrastructure subnet",
			Optional:            true,
		},
		"infrastructure_subnet_ipv6": schema.StringAttribute{
			MarkdownDescription: "Infrastructure subnet ipv6",
			Optional:            true,
		},
		"ipv6": schema.BoolAttribute{
			MarkdownDescription: "Ipv6",
			Optional:            true,
		},
		"loopback_ips": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Loopback ips",
			Optional:            true,
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"tunnel_monitor_ip_address": schema.StringAttribute{
			MarkdownDescription: "Tunnel monitor ip address",
			Optional:            true,
		},
	},
}

// SharedInfrastructureSettingsDataSourceSchema defines the schema for SharedInfrastructureSettings data source
var SharedInfrastructureSettingsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "SharedInfrastructureSetting data source",
	Attributes: map[string]dsschema.Attribute{
		"api_key": dsschema.StringAttribute{
			MarkdownDescription: "Api key",
			Computed:            true,
		},
		"captive_portal_redirect_ip_address": dsschema.StringAttribute{
			MarkdownDescription: "Captive portal redirect ip address",
			Computed:            true,
		},
		"connector_application_blocks": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Connector application blocks",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"member": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Member",
					Computed:            true,
				},
			},
		},
		"connector_connector_blocks": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Connector connector blocks",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"member": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Member",
					Computed:            true,
				},
			},
		},
		"egress_ip_notification_url": dsschema.StringAttribute{
			MarkdownDescription: "Egress ip notification url",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder containing the shared infrastructure settings",
			Computed:            true,
		},
		"infra_bgp_as": dsschema.StringAttribute{
			MarkdownDescription: "Infra bgp as",
			Computed:            true,
		},
		"infrastructure_subnet": dsschema.StringAttribute{
			MarkdownDescription: "Infrastructure subnet",
			Computed:            true,
		},
		"infrastructure_subnet_ipv6": dsschema.StringAttribute{
			MarkdownDescription: "Infrastructure subnet ipv6",
			Computed:            true,
		},
		"ipv6": dsschema.BoolAttribute{
			MarkdownDescription: "Ipv6",
			Computed:            true,
		},
		"loopback_ips": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Loopback ips",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"tunnel_monitor_ip_address": dsschema.StringAttribute{
			MarkdownDescription: "Tunnel monitor ip address",
			Computed:            true,
		},
	},
}

// SharedInfrastructureSettingsListModel represents the data model for a list data source.
type SharedInfrastructureSettingsListModel struct {
	Tfid    types.String                   `tfsdk:"tfid"`
	Data    []SharedInfrastructureSettings `tfsdk:"data"`
	Limit   types.Int64                    `tfsdk:"limit"`
	Offset  types.Int64                    `tfsdk:"offset"`
	Name    types.String                   `tfsdk:"name"`
	Total   types.Int64                    `tfsdk:"total"`
	Folder  types.String                   `tfsdk:"folder"`
	Snippet types.String                   `tfsdk:"snippet"`
	Device  types.String                   `tfsdk:"device"`
}

// SharedInfrastructureSettingsListDataSourceSchema defines the schema for a list data source.
var SharedInfrastructureSettingsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: SharedInfrastructureSettingsDataSourceSchema.Attributes,
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
