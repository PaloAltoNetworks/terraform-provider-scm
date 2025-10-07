package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: network_services
// This file contains models for the network_services SDK package

// Layer3Subinterfaces represents the Terraform model for Layer3Subinterfaces
type Layer3Subinterfaces struct {
	Tfid                       types.String           `tfsdk:"tfid"`
	Arp                        basetypes.ListValue    `tfsdk:"arp"`
	Comment                    basetypes.StringValue  `tfsdk:"comment"`
	DdnsConfig                 basetypes.ObjectValue  `tfsdk:"ddns_config"`
	Device                     basetypes.StringValue  `tfsdk:"device"`
	DhcpClient                 basetypes.ObjectValue  `tfsdk:"dhcp_client"`
	Folder                     basetypes.StringValue  `tfsdk:"folder"`
	Id                         basetypes.StringValue  `tfsdk:"id"`
	InterfaceManagementProfile basetypes.StringValue  `tfsdk:"interface_management_profile"`
	Ip                         basetypes.ListValue    `tfsdk:"ip"`
	Mtu                        basetypes.Float64Value `tfsdk:"mtu"`
	Name                       basetypes.StringValue  `tfsdk:"name"`
	ParentInterface            basetypes.StringValue  `tfsdk:"parent_interface"`
	Snippet                    basetypes.StringValue  `tfsdk:"snippet"`
	Tag                        basetypes.Float64Value `tfsdk:"tag"`
}

// AttrTypes defines the attribute types for the Layer3Subinterfaces model.
func (o Layer3Subinterfaces) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"arp": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"hw_address": basetypes.StringType{},
				"name":       basetypes.StringType{},
			},
		}},
		"comment": basetypes.StringType{},
		"ddns_config": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ddns_cert_profile":    basetypes.StringType{},
				"ddns_enabled":         basetypes.BoolType{},
				"ddns_hostname":        basetypes.StringType{},
				"ddns_ip":              basetypes.StringType{},
				"ddns_update_interval": basetypes.Int64Type{},
				"ddns_vendor":          basetypes.StringType{},
				"ddns_vendor_config":   basetypes.StringType{},
			},
		},
		"device": basetypes.StringType{},
		"dhcp_client": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"create_default_route": basetypes.BoolType{},
				"default_route_metric": basetypes.Int64Type{},
				"enable":               basetypes.BoolType{},
				"send_hostname": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable":   basetypes.BoolType{},
						"hostname": basetypes.StringType{},
					},
				},
			},
		},
		"folder":                       basetypes.StringType{},
		"id":                           basetypes.StringType{},
		"interface_management_profile": basetypes.StringType{},
		"ip":                           basetypes.ListType{ElemType: basetypes.StringType{}},
		"mtu":                          basetypes.NumberType{},
		"name":                         basetypes.StringType{},
		"parent_interface":             basetypes.StringType{},
		"snippet":                      basetypes.StringType{},
		"tag":                          basetypes.NumberType{},
	}
}

// AttrType returns the attribute type for a list of Layer3Subinterfaces objects.
func (o Layer3Subinterfaces) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// Layer3SubinterfacesResourceSchema defines the schema for Layer3Subinterfaces resource
var Layer3SubinterfacesResourceSchema = schema.Schema{
	MarkdownDescription: "Layer3Subinterface resource",
	Attributes: map[string]schema.Attribute{
		"arp": schema.ListNestedAttribute{
			MarkdownDescription: "ARP configuration",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"hw_address": schema.StringAttribute{
						MarkdownDescription: "Hw address",
						Optional:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "IP address",
						Optional:            true,
					},
				},
			},
		},
		"comment": schema.StringAttribute{
			MarkdownDescription: "Description",
			Optional:            true,
		},
		"ddns_config": schema.SingleNestedAttribute{
			MarkdownDescription: "Ddns config",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"ddns_cert_profile": schema.StringAttribute{
					MarkdownDescription: "Ddns cert profile",
					Optional:            true,
				},
				"ddns_enabled": schema.BoolAttribute{
					MarkdownDescription: "Ddns enabled",
					Optional:            true,
				},
				"ddns_hostname": schema.StringAttribute{
					MarkdownDescription: "Ddns hostname",
					Optional:            true,
				},
				"ddns_ip": schema.StringAttribute{
					MarkdownDescription: "Ddns ip",
					Optional:            true,
				},
				"ddns_update_interval": schema.Int64Attribute{
					MarkdownDescription: "Ddns update interval",
					Optional:            true,
				},
				"ddns_vendor": schema.StringAttribute{
					MarkdownDescription: "Ddns vendor",
					Optional:            true,
				},
				"ddns_vendor_config": schema.StringAttribute{
					MarkdownDescription: "Ddns vendor config",
					Optional:            true,
				},
			},
		},
		"device": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
			},
			MarkdownDescription: "The device in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"dhcp_client": schema.SingleNestedAttribute{
			MarkdownDescription: "Dhcp client",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"create_default_route": schema.BoolAttribute{
					MarkdownDescription: "Create default route",
					Optional:            true,
				},
				"default_route_metric": schema.Int64Attribute{
					MarkdownDescription: "Default route metric",
					Optional:            true,
				},
				"enable": schema.BoolAttribute{
					MarkdownDescription: "Enable",
					Optional:            true,
				},
				"send_hostname": schema.SingleNestedAttribute{
					MarkdownDescription: "Send hostname",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"enable": schema.BoolAttribute{
							MarkdownDescription: "Enable",
							Optional:            true,
						},
						"hostname": schema.StringAttribute{
							MarkdownDescription: "Hostname",
							Optional:            true,
						},
					},
				},
			},
		},
		"folder": schema.StringAttribute{
			Validators: []validator.String{
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
		"interface_management_profile": schema.StringAttribute{
			MarkdownDescription: "Interface management profile",
			Optional:            true,
		},
		"ip": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Ip",
			Optional:            true,
		},
		"mtu": schema.Float64Attribute{
			Validators: []validator.Float64{
				float64validator.Between(576.000000, 9216.000000),
			},
			MarkdownDescription: "MTU",
			Optional:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "L3 sub-interface name",
			Required:            true,
		},
		"parent_interface": schema.StringAttribute{
			MarkdownDescription: "Parent interface",
			Optional:            true,
		},
		"snippet": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
			},
			MarkdownDescription: "The snippet in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"tag": schema.Float64Attribute{
			Validators: []validator.Float64{
				float64validator.Between(1.000000, 4096.000000),
			},
			MarkdownDescription: "VLAN tag",
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

// Layer3SubinterfacesDataSourceSchema defines the schema for Layer3Subinterfaces data source
var Layer3SubinterfacesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Layer3Subinterface data source",
	Attributes: map[string]dsschema.Attribute{
		"arp": dsschema.ListNestedAttribute{
			MarkdownDescription: "ARP configuration",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"hw_address": dsschema.StringAttribute{
						MarkdownDescription: "Hw address",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "IP address",
						Computed:            true,
					},
				},
			},
		},
		"comment": dsschema.StringAttribute{
			MarkdownDescription: "Description",
			Computed:            true,
		},
		"ddns_config": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Ddns config",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"ddns_cert_profile": dsschema.StringAttribute{
					MarkdownDescription: "Ddns cert profile",
					Computed:            true,
				},
				"ddns_enabled": dsschema.BoolAttribute{
					MarkdownDescription: "Ddns enabled",
					Computed:            true,
				},
				"ddns_hostname": dsschema.StringAttribute{
					MarkdownDescription: "Ddns hostname",
					Computed:            true,
				},
				"ddns_ip": dsschema.StringAttribute{
					MarkdownDescription: "Ddns ip",
					Computed:            true,
				},
				"ddns_update_interval": dsschema.Int64Attribute{
					MarkdownDescription: "Ddns update interval",
					Computed:            true,
				},
				"ddns_vendor": dsschema.StringAttribute{
					MarkdownDescription: "Ddns vendor",
					Computed:            true,
				},
				"ddns_vendor_config": dsschema.StringAttribute{
					MarkdownDescription: "Ddns vendor config",
					Computed:            true,
				},
			},
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"dhcp_client": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Dhcp client",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"create_default_route": dsschema.BoolAttribute{
					MarkdownDescription: "Create default route",
					Computed:            true,
				},
				"default_route_metric": dsschema.Int64Attribute{
					MarkdownDescription: "Default route metric",
					Computed:            true,
				},
				"enable": dsschema.BoolAttribute{
					MarkdownDescription: "Enable",
					Computed:            true,
				},
				"send_hostname": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Send hostname",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"enable": dsschema.BoolAttribute{
							MarkdownDescription: "Enable",
							Computed:            true,
						},
						"hostname": dsschema.StringAttribute{
							MarkdownDescription: "Hostname",
							Computed:            true,
						},
					},
				},
			},
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"interface_management_profile": dsschema.StringAttribute{
			MarkdownDescription: "Interface management profile",
			Computed:            true,
		},
		"ip": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Ip",
			Computed:            true,
		},
		"mtu": dsschema.Float64Attribute{
			MarkdownDescription: "MTU",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "L3 sub-interface name",
			Optional:            true,
			Computed:            true,
		},
		"parent_interface": dsschema.StringAttribute{
			MarkdownDescription: "Parent interface",
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"tag": dsschema.Float64Attribute{
			MarkdownDescription: "VLAN tag",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// Layer3SubinterfacesListModel represents the data model for a list data source.
type Layer3SubinterfacesListModel struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Data    []Layer3Subinterfaces `tfsdk:"data"`
	Limit   types.Int64           `tfsdk:"limit"`
	Offset  types.Int64           `tfsdk:"offset"`
	Name    types.String          `tfsdk:"name"`
	Total   types.Int64           `tfsdk:"total"`
	Folder  types.String          `tfsdk:"folder"`
	Snippet types.String          `tfsdk:"snippet"`
	Device  types.String          `tfsdk:"device"`
}

// Layer3SubinterfacesListDataSourceSchema defines the schema for a list data source.
var Layer3SubinterfacesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: Layer3SubinterfacesDataSourceSchema.Attributes,
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
