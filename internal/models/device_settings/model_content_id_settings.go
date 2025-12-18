package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: device_settings
// This file contains models for the device_settings SDK package

// ContentIdSettings represents the Terraform model for ContentIdSettings
type ContentIdSettings struct {
	Tfid      types.String          `tfsdk:"tfid"`
	ContentId basetypes.ObjectValue `tfsdk:"content_id"`
	Device    basetypes.StringValue `tfsdk:"device"`
	Folder    basetypes.StringValue `tfsdk:"folder"`
	Id        basetypes.StringValue `tfsdk:"id"`
	Snippet   basetypes.StringValue `tfsdk:"snippet"`
}

// ContentIdSettingsContentId represents a nested structure within the ContentIdSettings model
type ContentIdSettingsContentId struct {
	AllowForwardDecryptedContent basetypes.BoolValue   `tfsdk:"allow_forward_decrypted_content"`
	AllowHttpRange               basetypes.BoolValue   `tfsdk:"allow_http_range"`
	Application                  basetypes.ObjectValue `tfsdk:"application"`
	ExtendedCaptureSegment       basetypes.Int64Value  `tfsdk:"extended_capture_segment"`
	StripXFwdFor                 basetypes.BoolValue   `tfsdk:"strip_x_fwd_for"`
	TcpBypassExceedQueue         basetypes.BoolValue   `tfsdk:"tcp_bypass_exceed_queue"`
	UdpBypassExceedQueue         basetypes.BoolValue   `tfsdk:"udp_bypass_exceed_queue"`
	XForwardedFor                basetypes.StringValue `tfsdk:"x_forwarded_for"`
}

// ContentIdSettingsContentIdApplication represents a nested structure within the ContentIdSettings model
type ContentIdSettingsContentIdApplication struct {
	BypassExceedQueue basetypes.BoolValue `tfsdk:"bypass_exceed_queue"`
}

// AttrTypes defines the attribute types for the ContentIdSettings model.
func (o ContentIdSettings) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"content_id": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"allow_forward_decrypted_content": basetypes.BoolType{},
				"allow_http_range":                basetypes.BoolType{},
				"application": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"bypass_exceed_queue": basetypes.BoolType{},
					},
				},
				"extended_capture_segment": basetypes.Int64Type{},
				"strip_x_fwd_for":          basetypes.BoolType{},
				"tcp_bypass_exceed_queue":  basetypes.BoolType{},
				"udp_bypass_exceed_queue":  basetypes.BoolType{},
				"x_forwarded_for":          basetypes.StringType{},
			},
		},
		"device":  basetypes.StringType{},
		"folder":  basetypes.StringType{},
		"id":      basetypes.StringType{},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ContentIdSettings objects.
func (o ContentIdSettings) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ContentIdSettingsContentId model.
func (o ContentIdSettingsContentId) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"allow_forward_decrypted_content": basetypes.BoolType{},
		"allow_http_range":                basetypes.BoolType{},
		"application": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"bypass_exceed_queue": basetypes.BoolType{},
			},
		},
		"extended_capture_segment": basetypes.Int64Type{},
		"strip_x_fwd_for":          basetypes.BoolType{},
		"tcp_bypass_exceed_queue":  basetypes.BoolType{},
		"udp_bypass_exceed_queue":  basetypes.BoolType{},
		"x_forwarded_for":          basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ContentIdSettingsContentId objects.
func (o ContentIdSettingsContentId) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ContentIdSettingsContentIdApplication model.
func (o ContentIdSettingsContentIdApplication) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"bypass_exceed_queue": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of ContentIdSettingsContentIdApplication objects.
func (o ContentIdSettingsContentIdApplication) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// ContentIdSettingsResourceSchema defines the schema for ContentIdSettings resource
var ContentIdSettingsResourceSchema = schema.Schema{
	MarkdownDescription: "ContentIdSetting resource",
	Attributes: map[string]schema.Attribute{
		"content_id": schema.SingleNestedAttribute{
			MarkdownDescription: "Content id",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"allow_forward_decrypted_content": schema.BoolAttribute{
					MarkdownDescription: "Allow forward decrypted content",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"allow_http_range": schema.BoolAttribute{
					MarkdownDescription: "Allow http range",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"application": schema.SingleNestedAttribute{
					MarkdownDescription: "Application",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"bypass_exceed_queue": schema.BoolAttribute{
							MarkdownDescription: "Bypass exceed queue",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
					},
				},
				"extended_capture_segment": schema.Int64Attribute{
					MarkdownDescription: "Extended capture segment",
					Optional:            true,
					Computed:            true,
					Default:             int64default.StaticInt64(5),
				},
				"strip_x_fwd_for": schema.BoolAttribute{
					MarkdownDescription: "Strip x fwd for",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"tcp_bypass_exceed_queue": schema.BoolAttribute{
					MarkdownDescription: "Tcp bypass exceed queue",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"udp_bypass_exceed_queue": schema.BoolAttribute{
					MarkdownDescription: "Udp bypass exceed queue",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
				"x_forwarded_for": schema.StringAttribute{
					MarkdownDescription: "X forwarded for",
					Optional:            true,
					Computed:            true,
					Default:             stringdefault.StaticString("0"),
				},
			},
		},
		"device": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("folder"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
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
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
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
		"snippet": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("device"),
					path.MatchRelative().AtParent().AtName("folder"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
			},
			MarkdownDescription: "The snippet in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
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

// ContentIdSettingsDataSourceSchema defines the schema for ContentIdSettings data source
var ContentIdSettingsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "ContentIdSetting data source",
	Attributes: map[string]dsschema.Attribute{
		"content_id": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Content id",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"allow_forward_decrypted_content": dsschema.BoolAttribute{
					MarkdownDescription: "Allow forward decrypted content",
					Computed:            true,
				},
				"allow_http_range": dsschema.BoolAttribute{
					MarkdownDescription: "Allow http range",
					Computed:            true,
				},
				"application": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Application",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"bypass_exceed_queue": dsschema.BoolAttribute{
							MarkdownDescription: "Bypass exceed queue",
							Computed:            true,
						},
					},
				},
				"extended_capture_segment": dsschema.Int64Attribute{
					MarkdownDescription: "Extended capture segment",
					Computed:            true,
				},
				"strip_x_fwd_for": dsschema.BoolAttribute{
					MarkdownDescription: "Strip x fwd for",
					Computed:            true,
				},
				"tcp_bypass_exceed_queue": dsschema.BoolAttribute{
					MarkdownDescription: "Tcp bypass exceed queue",
					Computed:            true,
				},
				"udp_bypass_exceed_queue": dsschema.BoolAttribute{
					MarkdownDescription: "Udp bypass exceed queue",
					Computed:            true,
				},
				"x_forwarded_for": dsschema.StringAttribute{
					MarkdownDescription: "X forwarded for",
					Computed:            true,
				},
			},
		},
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
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// ContentIdSettingsListModel represents the data model for a list data source.
type ContentIdSettingsListModel struct {
	Tfid    types.String        `tfsdk:"tfid"`
	Data    []ContentIdSettings `tfsdk:"data"`
	Limit   types.Int64         `tfsdk:"limit"`
	Offset  types.Int64         `tfsdk:"offset"`
	Name    types.String        `tfsdk:"name"`
	Total   types.Int64         `tfsdk:"total"`
	Folder  types.String        `tfsdk:"folder"`
	Snippet types.String        `tfsdk:"snippet"`
	Device  types.String        `tfsdk:"device"`
}

// ContentIdSettingsListDataSourceSchema defines the schema for a list data source.
var ContentIdSettingsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: ContentIdSettingsDataSourceSchema.Attributes,
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
