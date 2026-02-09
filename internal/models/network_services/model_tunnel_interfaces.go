package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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

// Package: network_services
// This file contains models for the network_services SDK package

// TunnelInterfaces represents the Terraform model for TunnelInterfaces
type TunnelInterfaces struct {
	Tfid                       types.String          `tfsdk:"tfid"`
	Comment                    basetypes.StringValue `tfsdk:"comment"`
	DefaultValue               basetypes.StringValue `tfsdk:"default_value"`
	Device                     basetypes.StringValue `tfsdk:"device"`
	Folder                     basetypes.StringValue `tfsdk:"folder"`
	Id                         basetypes.StringValue `tfsdk:"id"`
	InterfaceManagementProfile basetypes.StringValue `tfsdk:"interface_management_profile"`
	Ip                         basetypes.ListValue   `tfsdk:"ip"`
	Ipv6                       basetypes.ObjectValue `tfsdk:"ipv6"`
	Mtu                        basetypes.Int64Value  `tfsdk:"mtu"`
	Name                       basetypes.StringValue `tfsdk:"name"`
	Snippet                    basetypes.StringValue `tfsdk:"snippet"`
}

// TunnelInterfacesIpInner represents a nested structure within the TunnelInterfaces model
type TunnelInterfacesIpInner struct {
	Name basetypes.StringValue `tfsdk:"name"`
}

// TunnelInterfacesIpv6 represents a nested structure within the TunnelInterfaces model
type TunnelInterfacesIpv6 struct {
	Address     basetypes.ListValue   `tfsdk:"address"`
	Enabled     basetypes.BoolValue   `tfsdk:"enabled"`
	InterfaceId basetypes.StringValue `tfsdk:"interface_id"`
}

// TunnelInterfacesIpv6AddressInner represents a nested structure within the TunnelInterfaces model
type TunnelInterfacesIpv6AddressInner struct {
	Anycast           basetypes.ObjectValue `tfsdk:"anycast"`
	EnableOnInterface basetypes.BoolValue   `tfsdk:"enable_on_interface"`
	Name              basetypes.StringValue `tfsdk:"name"`
	Prefix            basetypes.ObjectValue `tfsdk:"prefix"`
}

// AttrTypes defines the attribute types for the TunnelInterfaces model.
func (o TunnelInterfaces) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":                         basetypes.StringType{},
		"comment":                      basetypes.StringType{},
		"default_value":                basetypes.StringType{},
		"device":                       basetypes.StringType{},
		"folder":                       basetypes.StringType{},
		"id":                           basetypes.StringType{},
		"interface_management_profile": basetypes.StringType{},
		"ip": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name": basetypes.StringType{},
			},
		}},
		"ipv6": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"anycast": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"enable_on_interface": basetypes.BoolType{},
						"name":                basetypes.StringType{},
						"prefix": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				}},
				"enabled":      basetypes.BoolType{},
				"interface_id": basetypes.StringType{},
			},
		},
		"mtu":     basetypes.Int64Type{},
		"name":    basetypes.StringType{},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of TunnelInterfaces objects.
func (o TunnelInterfaces) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the TunnelInterfacesIpInner model.
func (o TunnelInterfacesIpInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of TunnelInterfacesIpInner objects.
func (o TunnelInterfacesIpInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the TunnelInterfacesIpv6 model.
func (o TunnelInterfacesIpv6) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"anycast": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"enable_on_interface": basetypes.BoolType{},
				"name":                basetypes.StringType{},
				"prefix": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		}},
		"enabled":      basetypes.BoolType{},
		"interface_id": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of TunnelInterfacesIpv6 objects.
func (o TunnelInterfacesIpv6) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the TunnelInterfacesIpv6AddressInner model.
func (o TunnelInterfacesIpv6AddressInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"anycast": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"enable_on_interface": basetypes.BoolType{},
		"name":                basetypes.StringType{},
		"prefix": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of TunnelInterfacesIpv6AddressInner objects.
func (o TunnelInterfacesIpv6AddressInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// TunnelInterfacesResourceSchema defines the schema for TunnelInterfaces resource
var TunnelInterfacesResourceSchema = schema.Schema{
	MarkdownDescription: "TunnelInterface resource",
	Attributes: map[string]schema.Attribute{
		"comment": schema.StringAttribute{
			MarkdownDescription: "Description for tunnel interface",
			Optional:            true,
		},
		"default_value": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.RegexMatches(regexp.MustCompile("^tunnel\\.([1-9][0-9]{0,3})$"), "pattern must match "+"^tunnel\\.([1-9][0-9]{0,3})$"),
			},
			MarkdownDescription: "Default interface assignment for tunnel interface",
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
				utils.FolderValidator(),
			},
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "UUID of the resource for tunnel interface",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"interface_management_profile": schema.StringAttribute{
			MarkdownDescription: "Interface management profile for tunnel interface",
			Optional:            true,
		},
		"ip": schema.ListNestedAttribute{
			MarkdownDescription: "Tunnel Interface IP Parent",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						MarkdownDescription: "Tunnel Interface IP address(es)",
						Required:            true,
					},
				},
			},
		},
		"ipv6": schema.SingleNestedAttribute{
			MarkdownDescription: "Tunnel Interface IPv6 Configuration",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"address": schema.ListNestedAttribute{
					MarkdownDescription: "IPv6 Address Parent for tunnel interface",
					Optional:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"anycast": schema.SingleNestedAttribute{
								MarkdownDescription: "Anycast for tunnel interface",
								Optional:            true,
								Attributes:          map[string]schema.Attribute{},
							},
							"enable_on_interface": schema.BoolAttribute{
								MarkdownDescription: "Enable Address on Interface for tunnel interface",
								Optional:            true,
								Computed:            true,
								Default:             booldefault.StaticBool(true),
							},
							"name": schema.StringAttribute{
								MarkdownDescription: "IPv6 Address for tunnel interface",
								Optional:            true,
							},
							"prefix": schema.SingleNestedAttribute{
								MarkdownDescription: "Use interface ID as host portion for tunnel interface",
								Optional:            true,
								Attributes:          map[string]schema.Attribute{},
							},
						},
					},
				},
				"enabled": schema.BoolAttribute{
					MarkdownDescription: "Enable IPv6 for tunnel interface",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"interface_id": schema.StringAttribute{
					MarkdownDescription: "Interface ID for tunnel interface",
					Optional:            true,
					Computed:            true,
					Default:             stringdefault.StaticString("EUI-64"),
				},
			},
		},
		"mtu": schema.Int64Attribute{
			Validators: []validator.Int64{
				int64validator.Between(576, 9216),
			},
			MarkdownDescription: "MTU for tunnel interface",
			Optional:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "L3 sub-interface name for tunnel interface",
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

// TunnelInterfacesDataSourceSchema defines the schema for TunnelInterfaces data source
var TunnelInterfacesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "TunnelInterface data source",
	Attributes: map[string]dsschema.Attribute{
		"comment": dsschema.StringAttribute{
			MarkdownDescription: "Description for tunnel interface",
			Computed:            true,
		},
		"default_value": dsschema.StringAttribute{
			MarkdownDescription: "Default interface assignment for tunnel interface",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "UUID of the resource for tunnel interface",
			Required:            true,
		},
		"interface_management_profile": dsschema.StringAttribute{
			MarkdownDescription: "Interface management profile for tunnel interface",
			Computed:            true,
		},
		"ip": dsschema.ListNestedAttribute{
			MarkdownDescription: "Tunnel Interface IP Parent",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"name": dsschema.StringAttribute{
						MarkdownDescription: "Tunnel Interface IP address(es)",
						Computed:            true,
					},
				},
			},
		},
		"ipv6": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Tunnel Interface IPv6 Configuration",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"address": dsschema.ListNestedAttribute{
					MarkdownDescription: "IPv6 Address Parent for tunnel interface",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"anycast": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Anycast for tunnel interface",
								Computed:            true,
								Attributes:          map[string]dsschema.Attribute{},
							},
							"enable_on_interface": dsschema.BoolAttribute{
								MarkdownDescription: "Enable Address on Interface for tunnel interface",
								Computed:            true,
							},
							"name": dsschema.StringAttribute{
								MarkdownDescription: "IPv6 Address for tunnel interface",
								Computed:            true,
							},
							"prefix": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Use interface ID as host portion for tunnel interface",
								Computed:            true,
								Attributes:          map[string]dsschema.Attribute{},
							},
						},
					},
				},
				"enabled": dsschema.BoolAttribute{
					MarkdownDescription: "Enable IPv6 for tunnel interface",
					Computed:            true,
				},
				"interface_id": dsschema.StringAttribute{
					MarkdownDescription: "Interface ID for tunnel interface",
					Computed:            true,
				},
			},
		},
		"mtu": dsschema.Int64Attribute{
			MarkdownDescription: "MTU for tunnel interface",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "L3 sub-interface name for tunnel interface",
			Optional:            true,
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// TunnelInterfacesListModel represents the data model for a list data source.
type TunnelInterfacesListModel struct {
	Tfid    types.String       `tfsdk:"tfid"`
	Data    []TunnelInterfaces `tfsdk:"data"`
	Limit   types.Int64        `tfsdk:"limit"`
	Offset  types.Int64        `tfsdk:"offset"`
	Name    types.String       `tfsdk:"name"`
	Total   types.Int64        `tfsdk:"total"`
	Folder  types.String       `tfsdk:"folder"`
	Snippet types.String       `tfsdk:"snippet"`
	Device  types.String       `tfsdk:"device"`
}

// TunnelInterfacesListDataSourceSchema defines the schema for a list data source.
var TunnelInterfacesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: TunnelInterfacesDataSourceSchema.Attributes,
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
