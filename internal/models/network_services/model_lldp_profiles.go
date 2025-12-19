package models

import (
	"regexp"

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

// LldpProfiles represents the Terraform model for LldpProfiles
type LldpProfiles struct {
	Tfid                   types.String          `tfsdk:"tfid"`
	Device                 basetypes.StringValue `tfsdk:"device"`
	Folder                 basetypes.StringValue `tfsdk:"folder"`
	Id                     basetypes.StringValue `tfsdk:"id"`
	Mode                   basetypes.StringValue `tfsdk:"mode"`
	Name                   basetypes.StringValue `tfsdk:"name"`
	OptionTlvs             basetypes.ObjectValue `tfsdk:"option_tlvs"`
	Snippet                basetypes.StringValue `tfsdk:"snippet"`
	SnmpSyslogNotification basetypes.BoolValue   `tfsdk:"snmp_syslog_notification"`
}

// LldpProfilesOptionTlvs represents a nested structure within the LldpProfiles model
type LldpProfilesOptionTlvs struct {
	ManagementAddress  basetypes.ObjectValue `tfsdk:"management_address"`
	PortDescription    basetypes.BoolValue   `tfsdk:"port_description"`
	SystemCapabilities basetypes.BoolValue   `tfsdk:"system_capabilities"`
	SystemDescription  basetypes.BoolValue   `tfsdk:"system_description"`
	SystemName         basetypes.BoolValue   `tfsdk:"system_name"`
}

// LldpProfilesOptionTlvsManagementAddress represents a nested structure within the LldpProfiles model
type LldpProfilesOptionTlvsManagementAddress struct {
	Enabled basetypes.BoolValue `tfsdk:"enabled"`
	Iplist  basetypes.ListValue `tfsdk:"iplist"`
}

// LldpProfilesOptionTlvsManagementAddressIplistInner represents a nested structure within the LldpProfiles model
type LldpProfilesOptionTlvsManagementAddressIplistInner struct {
	Interface basetypes.StringValue `tfsdk:"interface"`
	Ipv4      basetypes.StringValue `tfsdk:"ipv4"`
	Ipv6      basetypes.StringValue `tfsdk:"ipv6"`
	Name      basetypes.StringValue `tfsdk:"name"`
}

// AttrTypes defines the attribute types for the LldpProfiles model.
func (o LldpProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":   basetypes.StringType{},
		"device": basetypes.StringType{},
		"folder": basetypes.StringType{},
		"id":     basetypes.StringType{},
		"mode":   basetypes.StringType{},
		"name":   basetypes.StringType{},
		"option_tlvs": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"management_address": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enabled": basetypes.BoolType{},
						"iplist": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"interface": basetypes.StringType{},
								"ipv4":      basetypes.StringType{},
								"ipv6":      basetypes.StringType{},
								"name":      basetypes.StringType{},
							},
						}},
					},
				},
				"port_description":    basetypes.BoolType{},
				"system_capabilities": basetypes.BoolType{},
				"system_description":  basetypes.BoolType{},
				"system_name":         basetypes.BoolType{},
			},
		},
		"snippet":                  basetypes.StringType{},
		"snmp_syslog_notification": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of LldpProfiles objects.
func (o LldpProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LldpProfilesOptionTlvs model.
func (o LldpProfilesOptionTlvs) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"management_address": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enabled": basetypes.BoolType{},
				"iplist": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"interface": basetypes.StringType{},
						"ipv4":      basetypes.StringType{},
						"ipv6":      basetypes.StringType{},
						"name":      basetypes.StringType{},
					},
				}},
			},
		},
		"port_description":    basetypes.BoolType{},
		"system_capabilities": basetypes.BoolType{},
		"system_description":  basetypes.BoolType{},
		"system_name":         basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of LldpProfilesOptionTlvs objects.
func (o LldpProfilesOptionTlvs) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LldpProfilesOptionTlvsManagementAddress model.
func (o LldpProfilesOptionTlvsManagementAddress) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enabled": basetypes.BoolType{},
		"iplist": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"interface": basetypes.StringType{},
				"ipv4":      basetypes.StringType{},
				"ipv6":      basetypes.StringType{},
				"name":      basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of LldpProfilesOptionTlvsManagementAddress objects.
func (o LldpProfilesOptionTlvsManagementAddress) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the LldpProfilesOptionTlvsManagementAddressIplistInner model.
func (o LldpProfilesOptionTlvsManagementAddressIplistInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"interface": basetypes.StringType{},
		"ipv4":      basetypes.StringType{},
		"ipv6":      basetypes.StringType{},
		"name":      basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of LldpProfilesOptionTlvsManagementAddressIplistInner objects.
func (o LldpProfilesOptionTlvsManagementAddressIplistInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// LldpProfilesResourceSchema defines the schema for LldpProfiles resource
var LldpProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "LldpProfile resource",
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
			MarkdownDescription: "The device in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
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
			MarkdownDescription: "The folder in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
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
		"mode": schema.StringAttribute{
			MarkdownDescription: "LLDP mode",
			Optional:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "LLDP profile name",
			Required:            true,
		},
		"option_tlvs": schema.SingleNestedAttribute{
			MarkdownDescription: "Option tlvs",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"management_address": schema.SingleNestedAttribute{
					MarkdownDescription: "Management address",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"enabled": schema.BoolAttribute{
							MarkdownDescription: "Management address enabled",
							Optional:            true,
						},
						"iplist": schema.ListNestedAttribute{
							MarkdownDescription: "Iplist",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"interface": schema.StringAttribute{
										MarkdownDescription: "Interface",
										Optional:            true,
									},
									"ipv4": schema.StringAttribute{
										MarkdownDescription: "IPv4 Address",
										Optional:            true,
									},
									"ipv6": schema.StringAttribute{
										MarkdownDescription: "IPv6 Address",
										Optional:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: "Name",
										Optional:            true,
									},
								},
							},
						},
					},
				},
				"port_description": schema.BoolAttribute{
					MarkdownDescription: "Option TLV Port Description",
					Optional:            true,
				},
				"system_capabilities": schema.BoolAttribute{
					MarkdownDescription: "Option TLV System Capabilities",
					Optional:            true,
				},
				"system_description": schema.BoolAttribute{
					MarkdownDescription: "Option TLV System Description",
					Optional:            true,
				},
				"system_name": schema.BoolAttribute{
					MarkdownDescription: "Option TLV System Name",
					Optional:            true,
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
			MarkdownDescription: "The snippet in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"snmp_syslog_notification": schema.BoolAttribute{
			MarkdownDescription: "SNMP syslog notification",
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

// LldpProfilesDataSourceSchema defines the schema for LldpProfiles data source
var LldpProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "LldpProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"mode": dsschema.StringAttribute{
			MarkdownDescription: "LLDP mode",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "LLDP profile name",
			Optional:            true,
			Computed:            true,
		},
		"option_tlvs": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Option tlvs",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"management_address": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Management address",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"enabled": dsschema.BoolAttribute{
							MarkdownDescription: "Management address enabled",
							Computed:            true,
						},
						"iplist": dsschema.ListNestedAttribute{
							MarkdownDescription: "Iplist",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"interface": dsschema.StringAttribute{
										MarkdownDescription: "Interface",
										Computed:            true,
									},
									"ipv4": dsschema.StringAttribute{
										MarkdownDescription: "IPv4 Address",
										Computed:            true,
									},
									"ipv6": dsschema.StringAttribute{
										MarkdownDescription: "IPv6 Address",
										Computed:            true,
									},
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Name",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"port_description": dsschema.BoolAttribute{
					MarkdownDescription: "Option TLV Port Description",
					Computed:            true,
				},
				"system_capabilities": dsschema.BoolAttribute{
					MarkdownDescription: "Option TLV System Capabilities",
					Computed:            true,
				},
				"system_description": dsschema.BoolAttribute{
					MarkdownDescription: "Option TLV System Description",
					Computed:            true,
				},
				"system_name": dsschema.BoolAttribute{
					MarkdownDescription: "Option TLV System Name",
					Computed:            true,
				},
			},
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"snmp_syslog_notification": dsschema.BoolAttribute{
			MarkdownDescription: "SNMP syslog notification",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// LldpProfilesListModel represents the data model for a list data source.
type LldpProfilesListModel struct {
	Tfid    types.String   `tfsdk:"tfid"`
	Data    []LldpProfiles `tfsdk:"data"`
	Limit   types.Int64    `tfsdk:"limit"`
	Offset  types.Int64    `tfsdk:"offset"`
	Name    types.String   `tfsdk:"name"`
	Total   types.Int64    `tfsdk:"total"`
	Folder  types.String   `tfsdk:"folder"`
	Snippet types.String   `tfsdk:"snippet"`
	Device  types.String   `tfsdk:"device"`
}

// LldpProfilesListDataSourceSchema defines the schema for a list data source.
var LldpProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: LldpProfilesDataSourceSchema.Attributes,
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
