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

// Wildcards represents the Terraform model for Wildcards
type Wildcards struct {
	Tfid         types.String          `tfsdk:"tfid"`
	AppEnabled   basetypes.BoolValue   `tfsdk:"app_enabled"`
	Applications basetypes.ObjectValue `tfsdk:"applications"`
	CreatedTime  basetypes.StringValue `tfsdk:"created_time"`
	Description  basetypes.StringValue `tfsdk:"description"`
	EnablePolicy basetypes.BoolValue   `tfsdk:"enable_policy"`
	Fqdn         basetypes.StringValue `tfsdk:"fqdn"`
	Group        basetypes.StringValue `tfsdk:"group"`
	IcmpAllowed  basetypes.BoolValue   `tfsdk:"icmp_allowed"`
	Id           basetypes.StringValue `tfsdk:"id"`
	Name         basetypes.StringValue `tfsdk:"name"`
	Oid          basetypes.StringValue `tfsdk:"oid"`
	ProbePort    basetypes.StringValue `tfsdk:"probe_port"`
	ProbeType    basetypes.StringValue `tfsdk:"probe_type"`
	TcpPort      basetypes.StringValue `tfsdk:"tcp_port"`
	UdpPort      basetypes.StringValue `tfsdk:"udp_port"`
	UpdatedTime  basetypes.StringValue `tfsdk:"updated_time"`
	UseDcIp      basetypes.BoolValue   `tfsdk:"use_dc_ip"`
}

// AttrTypes defines the attribute types for the Wildcards model.
func (o Wildcards) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":        basetypes.StringType{},
		"app_enabled": basetypes.BoolType{},
		"applications": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"created_time":  basetypes.StringType{},
		"description":   basetypes.StringType{},
		"enable_policy": basetypes.BoolType{},
		"fqdn":          basetypes.StringType{},
		"group":         basetypes.StringType{},
		"icmp_allowed":  basetypes.BoolType{},
		"id":            basetypes.StringType{},
		"name":          basetypes.StringType{},
		"oid":           basetypes.StringType{},
		"probe_port":    basetypes.StringType{},
		"probe_type":    basetypes.StringType{},
		"tcp_port":      basetypes.StringType{},
		"udp_port":      basetypes.StringType{},
		"updated_time":  basetypes.StringType{},
		"use_dc_ip":     basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of Wildcards objects.
func (o Wildcards) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// WildcardsResourceSchema defines the schema for Wildcards resource
var WildcardsResourceSchema = schema.Schema{
	MarkdownDescription: "Wildcard resource",
	Attributes: map[string]schema.Attribute{
		"app_enabled": schema.BoolAttribute{
			MarkdownDescription: "Whether the wildcard is enabled.\n\nIf omitted, defaults to false.",
			Optional:            true,
			Computed:            true,
		},
		"applications": schema.SingleNestedAttribute{
			MarkdownDescription: "Discovered Wildcard App oid to name mapping.",
			Computed:            true,
			Attributes:          map[string]schema.Attribute{},
		},
		"created_time": schema.StringAttribute{
			MarkdownDescription: "Created time",
			Computed:            true,
		},
		"description": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[\\p{L}\\p{N}\\p{P}\\s,.:_-]*$"), "pattern must match "+"^[\\p{L}\\p{N}\\p{P}\\s,.:_-]*$"),
			},
			MarkdownDescription: "Description",
			Optional:            true,
			Computed:            true,
		},
		"enable_policy": schema.BoolAttribute{
			MarkdownDescription: "Whether policy is enabled for this wildcard.\n\nIf omitted, defaults to false.",
			Optional:            true,
			Computed:            true,
		},
		"fqdn": schema.StringAttribute{
			MarkdownDescription: "The wildcard to match.",
			Required:            true,
		},
		"group": schema.StringAttribute{
			MarkdownDescription: "A comma separated list of connector group IDs",
			Required:            true,
		},
		"icmp_allowed": schema.BoolAttribute{
			MarkdownDescription: "Whether ICMP is allowed for this wildcard.\n\nIf omitted, defaults to true.",
			Optional:            true,
			Computed:            true,
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "Id of the entry as returned by list/get operations.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(64),
				stringvalidator.LengthAtLeast(1),
				stringvalidator.RegexMatches(regexp.MustCompile("^[\\p{L}\\p{N}\\p{P}\\s,.:_-]*$"), "pattern must match "+"^[\\p{L}\\p{N}\\p{P}\\s,.:_-]*$"),
			},
			MarkdownDescription: "Name of the wildcard.\n\nIt can only be 64 characters long\nand contain unicode text, space, dash, or underscore, or period.",
			Required:            true,
		},
		"oid": schema.StringAttribute{
			MarkdownDescription: "Id of the entry.",
			Computed:            true,
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
			MarkdownDescription: "The probing type.\n\nThe value can be `tcp_ping`, `icmp_ping`, or omitted.",
			Optional:            true,
			Computed:            true,
		},
		"tcp_port": schema.StringAttribute{
			MarkdownDescription: "TCP port number(s).\n\nIt can be a single port number, multiple port numbers separated by comma,\nor a port range like 8000-9000.\n\nIf both tcp_port and udp_port are omitted, tcp_port defaults to 443.",
			Optional:            true,
			Computed:            true,
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"udp_port": schema.StringAttribute{
			MarkdownDescription: "UDP port number(s).\n\nIt can be a single port number, multiple port numbers separated by comma,\nor a port range like 8000-9000.\n\nIf both tcp_port and udp_port are omitted, tcp_port defaults to 443.",
			Optional:            true,
			Computed:            true,
		},
		"updated_time": schema.StringAttribute{
			MarkdownDescription: "Updated time",
			Computed:            true,
		},
		"use_dc_ip": schema.BoolAttribute{
			MarkdownDescription: "Whether to use datacenter IP for this wildcard.\n\nIf omitted, defaults to false.",
			Optional:            true,
			Computed:            true,
		},
	},
}

// WildcardsDataSourceSchema defines the schema for Wildcards data source
var WildcardsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Wildcard data source",
	Attributes: map[string]dsschema.Attribute{
		"app_enabled": dsschema.BoolAttribute{
			MarkdownDescription: "Whether the wildcard is enabled.\n\nIf omitted, defaults to false.",
			Computed:            true,
		},
		"applications": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Discovered Wildcard App oid to name mapping.",
			Computed:            true,
			Attributes:          map[string]dsschema.Attribute{},
		},
		"created_time": dsschema.StringAttribute{
			MarkdownDescription: "Created time",
			Computed:            true,
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Description",
			Computed:            true,
		},
		"enable_policy": dsschema.BoolAttribute{
			MarkdownDescription: "Whether policy is enabled for this wildcard.\n\nIf omitted, defaults to false.",
			Computed:            true,
		},
		"fqdn": dsschema.StringAttribute{
			MarkdownDescription: "The wildcard to match.",
			Computed:            true,
		},
		"group": dsschema.StringAttribute{
			MarkdownDescription: "A comma separated list of connector group IDs",
			Computed:            true,
		},
		"icmp_allowed": dsschema.BoolAttribute{
			MarkdownDescription: "Whether ICMP is allowed for this wildcard.\n\nIf omitted, defaults to true.",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "Id of the entry as returned by list/get operations.",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Name of the wildcard.\n\nIt can only be 64 characters long\nand contain unicode text, space, dash, or underscore, or period.",
			Optional:            true,
			Computed:            true,
		},
		"oid": dsschema.StringAttribute{
			MarkdownDescription: "Id of the entry.",
			Computed:            true,
		},
		"probe_port": dsschema.StringAttribute{
			MarkdownDescription: "The probing port if the `probe_type` is `tcp_ping`.",
			Computed:            true,
		},
		"probe_type": dsschema.StringAttribute{
			MarkdownDescription: "The probing type.\n\nThe value can be `tcp_ping`, `icmp_ping`, or omitted.",
			Computed:            true,
		},
		"tcp_port": dsschema.StringAttribute{
			MarkdownDescription: "TCP port number(s).\n\nIt can be a single port number, multiple port numbers separated by comma,\nor a port range like 8000-9000.\n\nIf both tcp_port and udp_port are omitted, tcp_port defaults to 443.",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"udp_port": dsschema.StringAttribute{
			MarkdownDescription: "UDP port number(s).\n\nIt can be a single port number, multiple port numbers separated by comma,\nor a port range like 8000-9000.\n\nIf both tcp_port and udp_port are omitted, tcp_port defaults to 443.",
			Computed:            true,
		},
		"updated_time": dsschema.StringAttribute{
			MarkdownDescription: "Updated time",
			Computed:            true,
		},
		"use_dc_ip": dsschema.BoolAttribute{
			MarkdownDescription: "Whether to use datacenter IP for this wildcard.\n\nIf omitted, defaults to false.",
			Computed:            true,
		},
	},
}

// WildcardsListModel represents the data model for a list data source.
type WildcardsListModel struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Data    []Wildcards           `tfsdk:"data"`
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

// WildcardsListDataSourceSchema defines the schema for a list data source.
var WildcardsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: WildcardsDataSourceSchema.Attributes,
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
