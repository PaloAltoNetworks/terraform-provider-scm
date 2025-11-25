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
)

// Package: network_services
// This file contains models for the network_services SDK package

// LoopbackInterfaces represents the Terraform model for LoopbackInterfaces
type LoopbackInterfaces struct {
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

// LoopbackInterfacesIpInner represents a nested structure within the LoopbackInterfaces model
type LoopbackInterfacesIpInner struct {
	Name basetypes.StringValue `tfsdk:"name"`
}

// LoopbackInterfacesIpv6 represents a nested structure within the LoopbackInterfaces model
type LoopbackInterfacesIpv6 struct {
	Address basetypes.ListValue `tfsdk:"address"`
	Enabled basetypes.BoolValue `tfsdk:"enabled"`
}

// LoopbackInterfacesIpv6AddressInner represents a nested structure within the LoopbackInterfaces model
type LoopbackInterfacesIpv6AddressInner struct {
	EnableOnInterface basetypes.BoolValue   `tfsdk:"enable_on_interface"`
	InterfaceId       basetypes.StringValue `tfsdk:"interface_id"`
	Name              basetypes.StringValue `tfsdk:"name"`
}

// AttrTypes defines the attribute types for the LoopbackInterfaces model.
func (o LoopbackInterfaces) AttrTypes() map[string]attr.Type {
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
						"enable_on_interface": basetypes.BoolType{},
						"interface_id":        basetypes.StringType{},
						"name":                basetypes.StringType{},
					},
				}},
				"enabled": basetypes.BoolType{},
			},
		},
		"mtu":     basetypes.Int64Type{},
		"name":    basetypes.StringType{},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LoopbackInterfaces objects.
func (o LoopbackInterfaces) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LoopbackInterfacesIpInner model.
func (o LoopbackInterfacesIpInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LoopbackInterfacesIpInner objects.
func (o LoopbackInterfacesIpInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LoopbackInterfacesIpv6 model.
func (o LoopbackInterfacesIpv6) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable_on_interface": basetypes.BoolType{},
				"interface_id":        basetypes.StringType{},
				"name":                basetypes.StringType{},
			},
		}},
		"enabled": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of LoopbackInterfacesIpv6 objects.
func (o LoopbackInterfacesIpv6) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LoopbackInterfacesIpv6AddressInner model.
func (o LoopbackInterfacesIpv6AddressInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable_on_interface": basetypes.BoolType{},
		"interface_id":        basetypes.StringType{},
		"name":                basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LoopbackInterfacesIpv6AddressInner objects.
func (o LoopbackInterfacesIpv6AddressInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// LoopbackInterfacesResourceSchema defines the schema for LoopbackInterfaces resource
var LoopbackInterfacesResourceSchema = schema.Schema{
	MarkdownDescription: "LoopbackInterface resource",
	Attributes: map[string]schema.Attribute{
		"comment": schema.StringAttribute{
			MarkdownDescription: "Description",
			Optional:            true,
		},
		"default_value": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.RegexMatches(regexp.MustCompile("^loopback\\.([1-9][0-9]{0,3})$"), "pattern must match "+"^loopback\\.([1-9][0-9]{0,3})$"),
			},
			MarkdownDescription: "Default interface assignment",
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
		"ip": schema.ListNestedAttribute{
			MarkdownDescription: "Loopback IP Parent",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						MarkdownDescription: "Loopback IP address(es)",
						Required:            true,
					},
				},
			},
		},
		"ipv6": schema.SingleNestedAttribute{
			MarkdownDescription: "Loopback IPv6 Configuration",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"address": schema.ListNestedAttribute{
					MarkdownDescription: "IPv6 Address Parent",
					Optional:            true,
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"enable_on_interface": schema.BoolAttribute{
								MarkdownDescription: "Enable Address on Interface",
								Optional:            true,
								Computed:            true,
								Default:             booldefault.StaticBool(true),
							},
							"interface_id": schema.StringAttribute{
								MarkdownDescription: "Interface ID",
								Optional:            true,
								Computed:            true,
								Default:             stringdefault.StaticString("EUI-64"),
							},
							"name": schema.StringAttribute{
								MarkdownDescription: "IPv6 Address",
								Optional:            true,
							},
						},
					},
				},
				"enabled": schema.BoolAttribute{
					MarkdownDescription: "Enable IPv6",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
			},
		},
		"mtu": schema.Int64Attribute{
			Validators: []validator.Int64{
				int64validator.Between(576, 9216),
			},
			MarkdownDescription: "MTU",
			Optional:            true,
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.RegexMatches(regexp.MustCompile("^\\$[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^\\$[a-zA-Z\\d\\-_\\. ]+$"),
			},
			MarkdownDescription: "Loopback Interface name",
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

// LoopbackInterfacesDataSourceSchema defines the schema for LoopbackInterfaces data source
var LoopbackInterfacesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "LoopbackInterface data source",
	Attributes: map[string]dsschema.Attribute{
		"comment": dsschema.StringAttribute{
			MarkdownDescription: "Description",
			Computed:            true,
		},
		"default_value": dsschema.StringAttribute{
			MarkdownDescription: "Default interface assignment",
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
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"interface_management_profile": dsschema.StringAttribute{
			MarkdownDescription: "Interface management profile",
			Computed:            true,
		},
		"ip": dsschema.ListNestedAttribute{
			MarkdownDescription: "Loopback IP Parent",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"name": dsschema.StringAttribute{
						MarkdownDescription: "Loopback IP address(es)",
						Computed:            true,
					},
				},
			},
		},
		"ipv6": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Loopback IPv6 Configuration",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"address": dsschema.ListNestedAttribute{
					MarkdownDescription: "IPv6 Address Parent",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"enable_on_interface": dsschema.BoolAttribute{
								MarkdownDescription: "Enable Address on Interface",
								Computed:            true,
							},
							"interface_id": dsschema.StringAttribute{
								MarkdownDescription: "Interface ID",
								Computed:            true,
							},
							"name": dsschema.StringAttribute{
								MarkdownDescription: "IPv6 Address",
								Computed:            true,
							},
						},
					},
				},
				"enabled": dsschema.BoolAttribute{
					MarkdownDescription: "Enable IPv6",
					Computed:            true,
				},
			},
		},
		"mtu": dsschema.Int64Attribute{
			MarkdownDescription: "MTU",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Loopback Interface name",
			Optional:            true,
			Computed:            true,
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

// LoopbackInterfacesListModel represents the data model for a list data source.
type LoopbackInterfacesListModel struct {
	Tfid    types.String         `tfsdk:"tfid"`
	Data    []LoopbackInterfaces `tfsdk:"data"`
	Limit   types.Int64          `tfsdk:"limit"`
	Offset  types.Int64          `tfsdk:"offset"`
	Name    types.String         `tfsdk:"name"`
	Total   types.Int64          `tfsdk:"total"`
	Folder  types.String         `tfsdk:"folder"`
	Snippet types.String         `tfsdk:"snippet"`
	Device  types.String         `tfsdk:"device"`
}

// LoopbackInterfacesListDataSourceSchema defines the schema for a list data source.
var LoopbackInterfacesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: LoopbackInterfacesDataSourceSchema.Attributes,
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
