package models

import (
	"regexp"

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

// Package: ztna_connector_all
// This file contains models for the ztna_connector_all SDK package

// Applications represents the Terraform model for Applications
type Applications struct {
	Tfid        types.String          `tfsdk:"tfid"`
	AnycastIp   basetypes.StringValue `tfsdk:"anycast_ip"`
	AppEnabled  basetypes.BoolValue   `tfsdk:"app_enabled"`
	CreatedTime basetypes.StringValue `tfsdk:"created_time"`
	Description basetypes.StringValue `tfsdk:"description"`
	Group       basetypes.StringValue `tfsdk:"group"`
	IcmpAllowed basetypes.BoolValue   `tfsdk:"icmp_allowed"`
	Name        basetypes.StringValue `tfsdk:"name"`
	Oid         basetypes.StringValue `tfsdk:"oid"`
	Spec        basetypes.ListValue   `tfsdk:"spec"`
	UpdatedTime basetypes.StringValue `tfsdk:"updated_time"`
	UseDcIp     basetypes.BoolValue   `tfsdk:"use_dc_ip"`
}

// ApplicationsSpecInner represents a nested structure within the Applications model
type ApplicationsSpecInner struct {
	Fqdn      basetypes.StringValue `tfsdk:"fqdn"`
	ProbePort basetypes.StringValue `tfsdk:"probe_port"`
	ProbeType basetypes.StringValue `tfsdk:"probe_type"`
	TcpPort   basetypes.StringValue `tfsdk:"tcp_port"`
	UdpPort   basetypes.StringValue `tfsdk:"udp_port"`
}

// AttrTypes defines the attribute types for the Applications model.
func (o Applications) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":         basetypes.StringType{},
		"anycast_ip":   basetypes.StringType{},
		"app_enabled":  basetypes.BoolType{},
		"created_time": basetypes.StringType{},
		"description":  basetypes.StringType{},
		"group":        basetypes.StringType{},
		"icmp_allowed": basetypes.BoolType{},
		"name":         basetypes.StringType{},
		"oid":          basetypes.StringType{},
		"spec": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"fqdn":       basetypes.StringType{},
				"probe_port": basetypes.StringType{},
				"probe_type": basetypes.StringType{},
				"tcp_port":   basetypes.StringType{},
				"udp_port":   basetypes.StringType{},
			},
		}},
		"updated_time": basetypes.StringType{},
		"use_dc_ip":    basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of Applications objects.
func (o Applications) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ApplicationsSpecInner model.
func (o ApplicationsSpecInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"fqdn":       basetypes.StringType{},
		"probe_port": basetypes.StringType{},
		"probe_type": basetypes.StringType{},
		"tcp_port":   basetypes.StringType{},
		"udp_port":   basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ApplicationsSpecInner objects.
func (o ApplicationsSpecInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// ApplicationsResourceSchema defines the schema for Applications resource
var ApplicationsResourceSchema = schema.Schema{
	MarkdownDescription: "Application resource",
	Attributes: map[string]schema.Attribute{
		"anycast_ip": schema.StringAttribute{
			MarkdownDescription: "Anycast IP address assigned to this FQDN rule.",
			Computed:            true,
		},
		"app_enabled": schema.BoolAttribute{
			MarkdownDescription: "Whether the FQDN rule is enabled.\n\nIf omitted, defaults to false.",
			Optional:            true,
			Computed:            true,
		},
		"created_time": schema.StringAttribute{
			MarkdownDescription: "Created time",
			Computed:            true,
		},
		"description": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(1023),
			},
			MarkdownDescription: "The description of the resource",
			Optional:            true,
			Computed:            true,
		},
		"group": schema.StringAttribute{
			MarkdownDescription: "A comma separated list of connector group IDs",
			Required:            true,
		},
		"icmp_allowed": schema.BoolAttribute{
			MarkdownDescription: "Whether ICMP is allowed for this FQDN rule.\n\nIf omitted, defaults to true.",
			Optional:            true,
			Computed:            true,
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(63),
				stringvalidator.LengthAtLeast(1),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
			},
			MarkdownDescription: "Name of the FQDN rule.",
			Required:            true,
		},
		"oid": schema.StringAttribute{
			MarkdownDescription: "The UUID of the resource",
			Computed:            true,
		},
		"spec": schema.ListNestedAttribute{
			MarkdownDescription: "Spec",
			Required:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"fqdn": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile("^([a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\\.)+[a-zA-Z]{2,63}$"), "pattern must match "+"^([a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\\.)+[a-zA-Z]{2,63}$"),
						},
						MarkdownDescription: "FQDN of the rule.",
						Required:            true,
					},
					"probe_port": schema.StringAttribute{
						MarkdownDescription: "The probing port if the `probe_type` is `tcp_ping`.",
						Optional:            true,
						Computed:            true,
					},
					"probe_type": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("tcp_ping", "icmp_ping"),
						},
						MarkdownDescription: "The probing type.\n\nThe value can be `tcp_ping`, `icmp_ping`, or omitted. Possible values are `tcp_ping` and `icmp_ping`.",
						Optional:            true,
						Computed:            true,
					},
					"tcp_port": schema.StringAttribute{
						MarkdownDescription: "TCP port number(s).\n\nIt can be a single port number, multiple port numbers separated by comma,\nor a port range like 8000-9000.\n\nIf both tcp_port and udp_port are omitted, tcp_port defaults to 443.",
						Optional:            true,
						Computed:            true,
					},
					"udp_port": schema.StringAttribute{
						MarkdownDescription: "UDP port number(s).\n\nIt can be a single port number, multiple port numbers separated by comma,\nor a port range like 8000-9000.\n\nIf both tcp_port and udp_port are omitted, tcp_port defaults to 443.",
						Optional:            true,
						Computed:            true,
					},
				},
			},
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"updated_time": schema.StringAttribute{
			MarkdownDescription: "Updated time",
			Computed:            true,
		},
		"use_dc_ip": schema.BoolAttribute{
			MarkdownDescription: "Whether to use datacenter IP for this FQDN rule.\n\nIf omitted, defaults to false.",
			Optional:            true,
			Computed:            true,
		},
	},
}

// ApplicationsDataSourceSchema defines the schema for Applications data source
var ApplicationsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Application data source",
	Attributes: map[string]dsschema.Attribute{
		"anycast_ip": dsschema.StringAttribute{
			MarkdownDescription: "Anycast IP address assigned to this FQDN rule.",
			Computed:            true,
		},
		"app_enabled": dsschema.BoolAttribute{
			MarkdownDescription: "Whether the FQDN rule is enabled.\n\nIf omitted, defaults to false.",
			Computed:            true,
		},
		"created_time": dsschema.StringAttribute{
			MarkdownDescription: "Created time",
			Computed:            true,
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "The description of the resource",
			Computed:            true,
		},
		"group": dsschema.StringAttribute{
			MarkdownDescription: "A comma separated list of connector group IDs",
			Computed:            true,
		},
		"icmp_allowed": dsschema.BoolAttribute{
			MarkdownDescription: "Whether ICMP is allowed for this FQDN rule.\n\nIf omitted, defaults to true.",
			Computed:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Name of the FQDN rule.",
			Optional:            true,
			Computed:            true,
		},
		"oid": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the resource",
			Computed:            true,
		},
		"spec": dsschema.ListNestedAttribute{
			MarkdownDescription: "Spec",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"fqdn": dsschema.StringAttribute{
						MarkdownDescription: "FQDN of the rule.",
						Computed:            true,
					},
					"probe_port": dsschema.StringAttribute{
						MarkdownDescription: "The probing port if the `probe_type` is `tcp_ping`.",
						Computed:            true,
					},
					"probe_type": dsschema.StringAttribute{
						MarkdownDescription: "The probing type.\n\nThe value can be `tcp_ping`, `icmp_ping`, or omitted. Possible values are `tcp_ping` and `icmp_ping`.",
						Computed:            true,
					},
					"tcp_port": dsschema.StringAttribute{
						MarkdownDescription: "TCP port number(s).\n\nIt can be a single port number, multiple port numbers separated by comma,\nor a port range like 8000-9000.\n\nIf both tcp_port and udp_port are omitted, tcp_port defaults to 443.",
						Computed:            true,
					},
					"udp_port": dsschema.StringAttribute{
						MarkdownDescription: "UDP port number(s).\n\nIt can be a single port number, multiple port numbers separated by comma,\nor a port range like 8000-9000.\n\nIf both tcp_port and udp_port are omitted, tcp_port defaults to 443.",
						Computed:            true,
					},
				},
			},
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"updated_time": dsschema.StringAttribute{
			MarkdownDescription: "Updated time",
			Computed:            true,
		},
		"use_dc_ip": dsschema.BoolAttribute{
			MarkdownDescription: "Whether to use datacenter IP for this FQDN rule.\n\nIf omitted, defaults to false.",
			Computed:            true,
		},
	},
}

// ApplicationsListModel represents the data model for a list data source.
type ApplicationsListModel struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Data    []Applications        `tfsdk:"data"`
	Limit   types.Int64           `tfsdk:"limit"`
	Offset  types.Int64           `tfsdk:"offset"`
	Name    types.String          `tfsdk:"name"`
	Total   types.Int64           `tfsdk:"total"`
	Folder  types.String          `tfsdk:"folder"`
	Snippet types.String          `tfsdk:"snippet"`
	Device  types.String          `tfsdk:"device"`
	Sort    basetypes.StringValue `tfsdk:"sort"`
	Search  basetypes.StringValue `tfsdk:"search"`
	Filters basetypes.StringValue `tfsdk:"filters"`
}

// ApplicationsListDataSourceSchema defines the schema for a list data source.
var ApplicationsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: ApplicationsDataSourceSchema.Attributes,
			},
		},
		"limit":   dsschema.Int64Attribute{Description: "The max number of items to return. Default: 200.", Optional: true},
		"offset":  dsschema.Int64Attribute{Description: "The offset of the first item to return.", Optional: true},
		"name":    dsschema.StringAttribute{Description: "The name of the item.", Optional: true},
		"total":   dsschema.Int64Attribute{Description: "The total number of items.", Computed: true},
		"folder":  dsschema.StringAttribute{Description: "The folder of the item. Default: Shared.", Optional: true},
		"snippet": dsschema.StringAttribute{Description: "The snippet of the item.", Optional: true},
		"device":  dsschema.StringAttribute{Description: "The device of the item.", Optional: true},
		"sort": dsschema.StringAttribute{
			Description: "List of fields from item response to sort by.",
			Optional:    true,
		},
		"search": dsschema.StringAttribute{
			Description: "String to filter list results by. Is searched over multiple fields in each object. Multiple searches can be specified.",
			Optional:    true,
		},
		"filters": dsschema.StringAttribute{
			Description: "String to filter list results by searching one specified field of an object. Multiple filters can be specified.",
			Optional:    true,
		},
	},
}
