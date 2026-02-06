package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// Package: device_settings
// This file contains models for the device_settings SDK package

// SessionTimeouts represents the Terraform model for SessionTimeouts
type SessionTimeouts struct {
	Tfid            types.String          `tfsdk:"tfid"`
	Device          basetypes.StringValue `tfsdk:"device"`
	Folder          basetypes.StringValue `tfsdk:"folder"`
	Id              basetypes.StringValue `tfsdk:"id"`
	SessionTimeouts basetypes.ObjectValue `tfsdk:"session_timeouts"`
	Snippet         basetypes.StringValue `tfsdk:"snippet"`
}

// SessionTimeoutsSessionTimeouts represents a nested structure within the SessionTimeouts model
type SessionTimeoutsSessionTimeouts struct {
	TimeoutCaptivePortal    basetypes.Int64Value `tfsdk:"timeout_captive_portal"`
	TimeoutDefault          basetypes.Int64Value `tfsdk:"timeout_default"`
	TimeoutDiscardDefault   basetypes.Int64Value `tfsdk:"timeout_discard_default"`
	TimeoutDiscardTcp       basetypes.Int64Value `tfsdk:"timeout_discard_tcp"`
	TimeoutDiscardUdp       basetypes.Int64Value `tfsdk:"timeout_discard_udp"`
	TimeoutIcmp             basetypes.Int64Value `tfsdk:"timeout_icmp"`
	TimeoutScan             basetypes.Int64Value `tfsdk:"timeout_scan"`
	TimeoutTcp              basetypes.Int64Value `tfsdk:"timeout_tcp"`
	TimeoutTcpHalfClosed    basetypes.Int64Value `tfsdk:"timeout_tcp_half_closed"`
	TimeoutTcpTimeWait      basetypes.Int64Value `tfsdk:"timeout_tcp_time_wait"`
	TimeoutTcpUnverifiedRst basetypes.Int64Value `tfsdk:"timeout_tcp_unverified_rst"`
	TimeoutTcphandshake     basetypes.Int64Value `tfsdk:"timeout_tcphandshake"`
	TimeoutTcpinit          basetypes.Int64Value `tfsdk:"timeout_tcpinit"`
	TimeoutUdp              basetypes.Int64Value `tfsdk:"timeout_udp"`
}

// AttrTypes defines the attribute types for the SessionTimeouts model.
func (o SessionTimeouts) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":   basetypes.StringType{},
		"device": basetypes.StringType{},
		"folder": basetypes.StringType{},
		"id":     basetypes.StringType{},
		"session_timeouts": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"timeout_captive_portal":     basetypes.Int64Type{},
				"timeout_default":            basetypes.Int64Type{},
				"timeout_discard_default":    basetypes.Int64Type{},
				"timeout_discard_tcp":        basetypes.Int64Type{},
				"timeout_discard_udp":        basetypes.Int64Type{},
				"timeout_icmp":               basetypes.Int64Type{},
				"timeout_scan":               basetypes.Int64Type{},
				"timeout_tcp":                basetypes.Int64Type{},
				"timeout_tcp_half_closed":    basetypes.Int64Type{},
				"timeout_tcp_time_wait":      basetypes.Int64Type{},
				"timeout_tcp_unverified_rst": basetypes.Int64Type{},
				"timeout_tcphandshake":       basetypes.Int64Type{},
				"timeout_tcpinit":            basetypes.Int64Type{},
				"timeout_udp":                basetypes.Int64Type{},
			},
		},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of SessionTimeouts objects.
func (o SessionTimeouts) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SessionTimeoutsSessionTimeouts model.
func (o SessionTimeoutsSessionTimeouts) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"timeout_captive_portal":     basetypes.Int64Type{},
		"timeout_default":            basetypes.Int64Type{},
		"timeout_discard_default":    basetypes.Int64Type{},
		"timeout_discard_tcp":        basetypes.Int64Type{},
		"timeout_discard_udp":        basetypes.Int64Type{},
		"timeout_icmp":               basetypes.Int64Type{},
		"timeout_scan":               basetypes.Int64Type{},
		"timeout_tcp":                basetypes.Int64Type{},
		"timeout_tcp_half_closed":    basetypes.Int64Type{},
		"timeout_tcp_time_wait":      basetypes.Int64Type{},
		"timeout_tcp_unverified_rst": basetypes.Int64Type{},
		"timeout_tcphandshake":       basetypes.Int64Type{},
		"timeout_tcpinit":            basetypes.Int64Type{},
		"timeout_udp":                basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of SessionTimeoutsSessionTimeouts objects.
func (o SessionTimeoutsSessionTimeouts) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// SessionTimeoutsResourceSchema defines the schema for SessionTimeouts resource
var SessionTimeoutsResourceSchema = schema.Schema{
	MarkdownDescription: "SessionTimeout resource",
	Attributes: map[string]schema.Attribute{
		"device": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("folder"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
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
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
				utils.FolderValidator(),
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
		"session_timeouts": schema.SingleNestedAttribute{
			MarkdownDescription: "Session timeouts",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"timeout_captive_portal": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(1, 15999999),
					},
					MarkdownDescription: "Captive Portal (seconds)",
					Optional:            true,
					Computed:            true,
					Default:             int64default.StaticInt64(30),
				},
				"timeout_default": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(1, 15999999),
					},
					MarkdownDescription: "Default timeout (seconds)",
					Optional:            true,
					Computed:            true,
					Default:             int64default.StaticInt64(30),
				},
				"timeout_discard_default": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(1, 15999999),
					},
					MarkdownDescription: "Discard default (seconds)",
					Optional:            true,
					Computed:            true,
					Default:             int64default.StaticInt64(60),
				},
				"timeout_discard_tcp": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(1, 15999999),
					},
					MarkdownDescription: "Discard TCP (seconds)",
					Optional:            true,
					Computed:            true,
					Default:             int64default.StaticInt64(90),
				},
				"timeout_discard_udp": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(1, 15999999),
					},
					MarkdownDescription: "Discard UDP (seconds)",
					Optional:            true,
					Computed:            true,
					Default:             int64default.StaticInt64(60),
				},
				"timeout_icmp": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(1, 15999999),
					},
					MarkdownDescription: "ICMP (seconds)",
					Optional:            true,
					Computed:            true,
					Default:             int64default.StaticInt64(6),
				},
				"timeout_scan": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(5, 30),
					},
					MarkdownDescription: "Scan (seconds)",
					Optional:            true,
					Computed:            true,
					Default:             int64default.StaticInt64(10),
				},
				"timeout_tcp": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(1, 15999999),
					},
					MarkdownDescription: "TCP (seconds)",
					Optional:            true,
					Computed:            true,
					Default:             int64default.StaticInt64(3600),
				},
				"timeout_tcp_half_closed": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(1, 604800),
					},
					MarkdownDescription: "TCP Half Closed (seconds)",
					Optional:            true,
					Computed:            true,
					Default:             int64default.StaticInt64(120),
				},
				"timeout_tcp_time_wait": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(1, 600),
					},
					MarkdownDescription: "TCP Time Wait (seconds)",
					Optional:            true,
					Computed:            true,
					Default:             int64default.StaticInt64(15),
				},
				"timeout_tcp_unverified_rst": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(1, 600),
					},
					MarkdownDescription: "Unverified RST (seconds)",
					Optional:            true,
					Computed:            true,
					Default:             int64default.StaticInt64(30),
				},
				"timeout_tcphandshake": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(1, 60),
					},
					MarkdownDescription: "TCP handshake (seconds)",
					Optional:            true,
					Computed:            true,
					Default:             int64default.StaticInt64(10),
				},
				"timeout_tcpinit": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(1, 60),
					},
					MarkdownDescription: "TCP init (seconds)",
					Optional:            true,
					Computed:            true,
					Default:             int64default.StaticInt64(5),
				},
				"timeout_udp": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(1, 15999999),
					},
					MarkdownDescription: "UDP (seconds)",
					Optional:            true,
					Computed:            true,
					Default:             int64default.StaticInt64(30),
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
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
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

// SessionTimeoutsDataSourceSchema defines the schema for SessionTimeouts data source
var SessionTimeoutsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "SessionTimeout data source",
	Attributes: map[string]dsschema.Attribute{
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
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"session_timeouts": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Session timeouts",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"timeout_captive_portal": dsschema.Int64Attribute{
					MarkdownDescription: "Captive Portal (seconds)",
					Computed:            true,
				},
				"timeout_default": dsschema.Int64Attribute{
					MarkdownDescription: "Default timeout (seconds)",
					Computed:            true,
				},
				"timeout_discard_default": dsschema.Int64Attribute{
					MarkdownDescription: "Discard default (seconds)",
					Computed:            true,
				},
				"timeout_discard_tcp": dsschema.Int64Attribute{
					MarkdownDescription: "Discard TCP (seconds)",
					Computed:            true,
				},
				"timeout_discard_udp": dsschema.Int64Attribute{
					MarkdownDescription: "Discard UDP (seconds)",
					Computed:            true,
				},
				"timeout_icmp": dsschema.Int64Attribute{
					MarkdownDescription: "ICMP (seconds)",
					Computed:            true,
				},
				"timeout_scan": dsschema.Int64Attribute{
					MarkdownDescription: "Scan (seconds)",
					Computed:            true,
				},
				"timeout_tcp": dsschema.Int64Attribute{
					MarkdownDescription: "TCP (seconds)",
					Computed:            true,
				},
				"timeout_tcp_half_closed": dsschema.Int64Attribute{
					MarkdownDescription: "TCP Half Closed (seconds)",
					Computed:            true,
				},
				"timeout_tcp_time_wait": dsschema.Int64Attribute{
					MarkdownDescription: "TCP Time Wait (seconds)",
					Computed:            true,
				},
				"timeout_tcp_unverified_rst": dsschema.Int64Attribute{
					MarkdownDescription: "Unverified RST (seconds)",
					Computed:            true,
				},
				"timeout_tcphandshake": dsschema.Int64Attribute{
					MarkdownDescription: "TCP handshake (seconds)",
					Computed:            true,
				},
				"timeout_tcpinit": dsschema.Int64Attribute{
					MarkdownDescription: "TCP init (seconds)",
					Computed:            true,
				},
				"timeout_udp": dsschema.Int64Attribute{
					MarkdownDescription: "UDP (seconds)",
					Computed:            true,
				},
			},
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

// SessionTimeoutsListModel represents the data model for a list data source.
type SessionTimeoutsListModel struct {
	Tfid    types.String      `tfsdk:"tfid"`
	Data    []SessionTimeouts `tfsdk:"data"`
	Limit   types.Int64       `tfsdk:"limit"`
	Offset  types.Int64       `tfsdk:"offset"`
	Name    types.String      `tfsdk:"name"`
	Total   types.Int64       `tfsdk:"total"`
	Folder  types.String      `tfsdk:"folder"`
	Snippet types.String      `tfsdk:"snippet"`
	Device  types.String      `tfsdk:"device"`
}

// SessionTimeoutsListDataSourceSchema defines the schema for a list data source.
var SessionTimeoutsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: SessionTimeoutsDataSourceSchema.Attributes,
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
