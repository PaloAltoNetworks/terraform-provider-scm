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

// Package: device_settings
// This file contains models for the device_settings SDK package

// TcpSettings represents the Terraform model for TcpSettings
type TcpSettings struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Device  basetypes.StringValue `tfsdk:"device"`
	Folder  basetypes.StringValue `tfsdk:"folder"`
	Id      basetypes.StringValue `tfsdk:"id"`
	Snippet basetypes.StringValue `tfsdk:"snippet"`
	Tcp     basetypes.ObjectValue `tfsdk:"tcp"`
}

// TcpSettingsTcp represents a nested structure within the TcpSettings model
type TcpSettingsTcp struct {
	AllowChallengeAck    basetypes.BoolValue   `tfsdk:"allow_challenge_ack"`
	AsymmetricPath       basetypes.StringValue `tfsdk:"asymmetric_path"`
	BypassExceedOoQueue  basetypes.BoolValue   `tfsdk:"bypass_exceed_oo_queue"`
	CheckTimestampOption basetypes.BoolValue   `tfsdk:"check_timestamp_option"`
	DropZeroFlag         basetypes.BoolValue   `tfsdk:"drop_zero_flag"`
	SiptcpCleartextProxy basetypes.StringValue `tfsdk:"siptcp_cleartext_proxy"`
	StripMptcpOption     basetypes.BoolValue   `tfsdk:"strip_mptcp_option"`
	TcpRetransmitScan    basetypes.BoolValue   `tfsdk:"tcp_retransmit_scan"`
	UrgentData           basetypes.StringValue `tfsdk:"urgent_data"`
}

// AttrTypes defines the attribute types for the TcpSettings model.
func (o TcpSettings) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":    basetypes.StringType{},
		"device":  basetypes.StringType{},
		"folder":  basetypes.StringType{},
		"id":      basetypes.StringType{},
		"snippet": basetypes.StringType{},
		"tcp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"allow_challenge_ack":    basetypes.BoolType{},
				"asymmetric_path":        basetypes.StringType{},
				"bypass_exceed_oo_queue": basetypes.BoolType{},
				"check_timestamp_option": basetypes.BoolType{},
				"drop_zero_flag":         basetypes.BoolType{},
				"siptcp_cleartext_proxy": basetypes.StringType{},
				"strip_mptcp_option":     basetypes.BoolType{},
				"tcp_retransmit_scan":    basetypes.BoolType{},
				"urgent_data":            basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of TcpSettings objects.
func (o TcpSettings) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the TcpSettingsTcp model.
func (o TcpSettingsTcp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"allow_challenge_ack":    basetypes.BoolType{},
		"asymmetric_path":        basetypes.StringType{},
		"bypass_exceed_oo_queue": basetypes.BoolType{},
		"check_timestamp_option": basetypes.BoolType{},
		"drop_zero_flag":         basetypes.BoolType{},
		"siptcp_cleartext_proxy": basetypes.StringType{},
		"strip_mptcp_option":     basetypes.BoolType{},
		"tcp_retransmit_scan":    basetypes.BoolType{},
		"urgent_data":            basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of TcpSettingsTcp objects.
func (o TcpSettingsTcp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// TcpSettingsResourceSchema defines the schema for TcpSettings resource
var TcpSettingsResourceSchema = schema.Schema{
	MarkdownDescription: "TcpSetting resource",
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
		"tcp": schema.SingleNestedAttribute{
			MarkdownDescription: "Tcp",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"allow_challenge_ack": schema.BoolAttribute{
					MarkdownDescription: "Allow arbitrary ACK in response to SYN?",
					Optional:            true,
				},
				"asymmetric_path": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("drop", "bypass"),
					},
					MarkdownDescription: "Asymmetric path action",
					Optional:            true,
				},
				"bypass_exceed_oo_queue": schema.BoolAttribute{
					MarkdownDescription: "Forward segments exceeding TCP out-of-order queue?",
					Optional:            true,
				},
				"check_timestamp_option": schema.BoolAttribute{
					MarkdownDescription: "Drop segments with null timestamp option?",
					Optional:            true,
				},
				"drop_zero_flag": schema.BoolAttribute{
					MarkdownDescription: "Drop segments without flag?",
					Optional:            true,
				},
				"siptcp_cleartext_proxy": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("0", "2", "3"),
					},
					MarkdownDescription: "SIP TCP cleartext action (`'0'` = Always Off, `'1'` = Always Enabled, `'2'` = Automatically enable proxy when needed)",
					Optional:            true,
				},
				"strip_mptcp_option": schema.BoolAttribute{
					MarkdownDescription: "Strip MPTCP option?",
					Optional:            true,
				},
				"tcp_retransmit_scan": schema.BoolAttribute{
					MarkdownDescription: "TCP retransmit scan?",
					Optional:            true,
				},
				"urgent_data": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("clear", "oobinline"),
					},
					MarkdownDescription: "Urgent data flag action",
					Optional:            true,
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
	},
}

// TcpSettingsDataSourceSchema defines the schema for TcpSettings data source
var TcpSettingsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "TcpSetting data source",
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
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"tcp": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Tcp",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"allow_challenge_ack": dsschema.BoolAttribute{
					MarkdownDescription: "Allow arbitrary ACK in response to SYN?",
					Computed:            true,
				},
				"asymmetric_path": dsschema.StringAttribute{
					MarkdownDescription: "Asymmetric path action",
					Computed:            true,
				},
				"bypass_exceed_oo_queue": dsschema.BoolAttribute{
					MarkdownDescription: "Forward segments exceeding TCP out-of-order queue?",
					Computed:            true,
				},
				"check_timestamp_option": dsschema.BoolAttribute{
					MarkdownDescription: "Drop segments with null timestamp option?",
					Computed:            true,
				},
				"drop_zero_flag": dsschema.BoolAttribute{
					MarkdownDescription: "Drop segments without flag?",
					Computed:            true,
				},
				"siptcp_cleartext_proxy": dsschema.StringAttribute{
					MarkdownDescription: "SIP TCP cleartext action (`'0'` = Always Off, `'1'` = Always Enabled, `'2'` = Automatically enable proxy when needed)",
					Computed:            true,
				},
				"strip_mptcp_option": dsschema.BoolAttribute{
					MarkdownDescription: "Strip MPTCP option?",
					Computed:            true,
				},
				"tcp_retransmit_scan": dsschema.BoolAttribute{
					MarkdownDescription: "TCP retransmit scan?",
					Computed:            true,
				},
				"urgent_data": dsschema.StringAttribute{
					MarkdownDescription: "Urgent data flag action",
					Computed:            true,
				},
			},
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// TcpSettingsListModel represents the data model for a list data source.
type TcpSettingsListModel struct {
	Tfid    types.String  `tfsdk:"tfid"`
	Data    []TcpSettings `tfsdk:"data"`
	Limit   types.Int64   `tfsdk:"limit"`
	Offset  types.Int64   `tfsdk:"offset"`
	Name    types.String  `tfsdk:"name"`
	Total   types.Int64   `tfsdk:"total"`
	Folder  types.String  `tfsdk:"folder"`
	Snippet types.String  `tfsdk:"snippet"`
	Device  types.String  `tfsdk:"device"`
}

// TcpSettingsListDataSourceSchema defines the schema for a list data source.
var TcpSettingsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: TcpSettingsDataSourceSchema.Attributes,
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
