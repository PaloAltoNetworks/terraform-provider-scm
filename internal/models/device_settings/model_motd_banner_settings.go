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

// MotdBannerSettings represents the Terraform model for MotdBannerSettings
type MotdBannerSettings struct {
	Tfid          types.String          `tfsdk:"tfid"`
	Device        basetypes.StringValue `tfsdk:"device"`
	Folder        basetypes.StringValue `tfsdk:"folder"`
	Id            basetypes.StringValue `tfsdk:"id"`
	MotdAndBanner basetypes.ObjectValue `tfsdk:"motd_and_banner"`
	Snippet       basetypes.StringValue `tfsdk:"snippet"`
}

// MotdBannerSettingsMotdAndBanner represents a nested structure within the MotdBannerSettings model
type MotdBannerSettingsMotdAndBanner struct {
	BannerFooter            basetypes.StringValue `tfsdk:"banner_footer"`
	BannerFooterColor       basetypes.StringValue `tfsdk:"banner_footer_color"`
	BannerFooterTextColor   basetypes.StringValue `tfsdk:"banner_footer_text_color"`
	BannerHeader            basetypes.StringValue `tfsdk:"banner_header"`
	BannerHeaderColor       basetypes.StringValue `tfsdk:"banner_header_color"`
	BannerHeaderFooterMatch basetypes.BoolValue   `tfsdk:"banner_header_footer_match"`
	BannerHeaderTextColor   basetypes.StringValue `tfsdk:"banner_header_text_color"`
	Message                 basetypes.StringValue `tfsdk:"message"`
	MotdColor               basetypes.StringValue `tfsdk:"motd_color"`
	MotdDoNotDisplayAgain   basetypes.BoolValue   `tfsdk:"motd_do_not_display_again"`
	MotdEnable              basetypes.BoolValue   `tfsdk:"motd_enable"`
	MotdTitle               basetypes.StringValue `tfsdk:"motd_title"`
	Severity                basetypes.StringValue `tfsdk:"severity"`
}

// AttrTypes defines the attribute types for the MotdBannerSettings model.
func (o MotdBannerSettings) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":   basetypes.StringType{},
		"device": basetypes.StringType{},
		"folder": basetypes.StringType{},
		"id":     basetypes.StringType{},
		"motd_and_banner": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"banner_footer":              basetypes.StringType{},
				"banner_footer_color":        basetypes.StringType{},
				"banner_footer_text_color":   basetypes.StringType{},
				"banner_header":              basetypes.StringType{},
				"banner_header_color":        basetypes.StringType{},
				"banner_header_footer_match": basetypes.BoolType{},
				"banner_header_text_color":   basetypes.StringType{},
				"message":                    basetypes.StringType{},
				"motd_color":                 basetypes.StringType{},
				"motd_do_not_display_again":  basetypes.BoolType{},
				"motd_enable":                basetypes.BoolType{},
				"motd_title":                 basetypes.StringType{},
				"severity":                   basetypes.StringType{},
			},
		},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of MotdBannerSettings objects.
func (o MotdBannerSettings) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the MotdBannerSettingsMotdAndBanner model.
func (o MotdBannerSettingsMotdAndBanner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"banner_footer":              basetypes.StringType{},
		"banner_footer_color":        basetypes.StringType{},
		"banner_footer_text_color":   basetypes.StringType{},
		"banner_header":              basetypes.StringType{},
		"banner_header_color":        basetypes.StringType{},
		"banner_header_footer_match": basetypes.BoolType{},
		"banner_header_text_color":   basetypes.StringType{},
		"message":                    basetypes.StringType{},
		"motd_color":                 basetypes.StringType{},
		"motd_do_not_display_again":  basetypes.BoolType{},
		"motd_enable":                basetypes.BoolType{},
		"motd_title":                 basetypes.StringType{},
		"severity":                   basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of MotdBannerSettingsMotdAndBanner objects.
func (o MotdBannerSettingsMotdAndBanner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// MotdBannerSettingsResourceSchema defines the schema for MotdBannerSettings resource
var MotdBannerSettingsResourceSchema = schema.Schema{
	MarkdownDescription: "MotdBannerSetting resource",
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
		"motd_and_banner": schema.SingleNestedAttribute{
			MarkdownDescription: "Motd and banner",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"banner_footer": schema.StringAttribute{
					MarkdownDescription: "Banner footer",
					Optional:            true,
				},
				"banner_footer_color": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("color1", "color2", "color3", "color4", "color5", "color6", "color7", "color8", "color9", "color10", "color11", "color12", "color13", "color14", "color15", "color16", "color17"),
					},
					MarkdownDescription: "The following list details the supported values and their colors.\n\n- `color1` = Red\n- `color2` = Green\n- `color3` = Blue\n- `color4` = Yellow\n- `color5` = Copper\n- `color6` = Orange\n- `color7` = Purple\n- `color8` = Gray\n- `color9` = Light Green\n- `color10` = Cyan\n- `color11` = Light Gray\n- `color12` = Blue Gray\n- `color13` = Lime\n- `color14` = Black\n- `color15` = Gold\n- `color16` = Brown\n- `color17` = Olive\n",
					Optional:            true,
				},
				"banner_footer_text_color": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("color1", "color2", "color3", "color4", "color5", "color6", "color7", "color8", "color9", "color10", "color11", "color12", "color13", "color14", "color15", "color16", "color17"),
					},
					MarkdownDescription: "The following list details the supported values and their colors.\n\n- `color1` = Red\n- `color2` = Green\n- `color3` = Blue\n- `color4` = Yellow\n- `color5` = Copper\n- `color6` = Orange\n- `color7` = Purple\n- `color8` = Gray\n- `color9` = Light Green\n- `color10` = Cyan\n- `color11` = Light Gray\n- `color12` = Blue Gray\n- `color13` = Lime\n- `color14` = Black\n- `color15` = Gold\n- `color16` = Brown\n- `color17` = Olive\n",
					Optional:            true,
				},
				"banner_header": schema.StringAttribute{
					MarkdownDescription: "Banner header",
					Optional:            true,
				},
				"banner_header_color": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("color1", "color2", "color3", "color4", "color5", "color6", "color7", "color8", "color9", "color10", "color11", "color12", "color13", "color14", "color15", "color16", "color17"),
					},
					MarkdownDescription: "The following list details the supported values and their colors.\n\n- `color1` = Red\n- `color2` = Green\n- `color3` = Blue\n- `color4` = Yellow\n- `color5` = Copper\n- `color6` = Orange\n- `color7` = Purple\n- `color8` = Gray\n- `color9` = Light Green\n- `color10` = Cyan\n- `color11` = Light Gray\n- `color12` = Blue Gray\n- `color13` = Lime\n- `color14` = Black\n- `color15` = Gold\n- `color16` = Brown\n- `color17` = Olive\n",
					Optional:            true,
				},
				"banner_header_footer_match": schema.BoolAttribute{
					MarkdownDescription: "Banner header footer match",
					Optional:            true,
				},
				"banner_header_text_color": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("color1", "color2", "color3", "color4", "color5", "color6", "color7", "color8", "color9", "color10", "color11", "color12", "color13", "color14", "color15", "color16", "color17"),
					},
					MarkdownDescription: "The following list details the supported values and their colors.\n\n- `color1` = Red\n- `color2` = Green\n- `color3` = Blue\n- `color4` = Yellow\n- `color5` = Copper\n- `color6` = Orange\n- `color7` = Purple\n- `color8` = Gray\n- `color9` = Light Green\n- `color10` = Cyan\n- `color11` = Light Gray\n- `color12` = Blue Gray\n- `color13` = Lime\n- `color14` = Black\n- `color15` = Gold\n- `color16` = Brown\n- `color17` = Olive\n",
					Optional:            true,
				},
				"message": schema.StringAttribute{
					MarkdownDescription: "Message",
					Optional:            true,
				},
				"motd_color": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("color1", "color2", "color3", "color4", "color5", "color6", "color7", "color8", "color9", "color10", "color11", "color12", "color13", "color14", "color15", "color16", "color17"),
					},
					MarkdownDescription: "The following list details the supported values and their colors.\n\n- `color1` = Red\n- `color2` = Green\n- `color3` = Blue\n- `color4` = Yellow\n- `color5` = Copper\n- `color6` = Orange\n- `color7` = Purple\n- `color8` = Gray\n- `color9` = Light Green\n- `color10` = Cyan\n- `color11` = Light Gray\n- `color12` = Blue Gray\n- `color13` = Lime\n- `color14` = Black\n- `color15` = Gold\n- `color16` = Brown\n- `color17` = Olive\n",
					Optional:            true,
				},
				"motd_do_not_display_again": schema.BoolAttribute{
					MarkdownDescription: "Motd do not display again",
					Optional:            true,
				},
				"motd_enable": schema.BoolAttribute{
					MarkdownDescription: "Motd enable",
					Optional:            true,
				},
				"motd_title": schema.StringAttribute{
					MarkdownDescription: "Motd title",
					Optional:            true,
				},
				"severity": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("warning", "question", "error", "info"),
					},
					MarkdownDescription: "Severity",
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

// MotdBannerSettingsDataSourceSchema defines the schema for MotdBannerSettings data source
var MotdBannerSettingsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "MotdBannerSetting data source",
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
		"motd_and_banner": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Motd and banner",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"banner_footer": dsschema.StringAttribute{
					MarkdownDescription: "Banner footer",
					Computed:            true,
				},
				"banner_footer_color": dsschema.StringAttribute{
					MarkdownDescription: "The following list details the supported values and their colors.\n\n- `color1` = Red\n- `color2` = Green\n- `color3` = Blue\n- `color4` = Yellow\n- `color5` = Copper\n- `color6` = Orange\n- `color7` = Purple\n- `color8` = Gray\n- `color9` = Light Green\n- `color10` = Cyan\n- `color11` = Light Gray\n- `color12` = Blue Gray\n- `color13` = Lime\n- `color14` = Black\n- `color15` = Gold\n- `color16` = Brown\n- `color17` = Olive\n",
					Computed:            true,
				},
				"banner_footer_text_color": dsschema.StringAttribute{
					MarkdownDescription: "The following list details the supported values and their colors.\n\n- `color1` = Red\n- `color2` = Green\n- `color3` = Blue\n- `color4` = Yellow\n- `color5` = Copper\n- `color6` = Orange\n- `color7` = Purple\n- `color8` = Gray\n- `color9` = Light Green\n- `color10` = Cyan\n- `color11` = Light Gray\n- `color12` = Blue Gray\n- `color13` = Lime\n- `color14` = Black\n- `color15` = Gold\n- `color16` = Brown\n- `color17` = Olive\n",
					Computed:            true,
				},
				"banner_header": dsschema.StringAttribute{
					MarkdownDescription: "Banner header",
					Computed:            true,
				},
				"banner_header_color": dsschema.StringAttribute{
					MarkdownDescription: "The following list details the supported values and their colors.\n\n- `color1` = Red\n- `color2` = Green\n- `color3` = Blue\n- `color4` = Yellow\n- `color5` = Copper\n- `color6` = Orange\n- `color7` = Purple\n- `color8` = Gray\n- `color9` = Light Green\n- `color10` = Cyan\n- `color11` = Light Gray\n- `color12` = Blue Gray\n- `color13` = Lime\n- `color14` = Black\n- `color15` = Gold\n- `color16` = Brown\n- `color17` = Olive\n",
					Computed:            true,
				},
				"banner_header_footer_match": dsschema.BoolAttribute{
					MarkdownDescription: "Banner header footer match",
					Computed:            true,
				},
				"banner_header_text_color": dsschema.StringAttribute{
					MarkdownDescription: "The following list details the supported values and their colors.\n\n- `color1` = Red\n- `color2` = Green\n- `color3` = Blue\n- `color4` = Yellow\n- `color5` = Copper\n- `color6` = Orange\n- `color7` = Purple\n- `color8` = Gray\n- `color9` = Light Green\n- `color10` = Cyan\n- `color11` = Light Gray\n- `color12` = Blue Gray\n- `color13` = Lime\n- `color14` = Black\n- `color15` = Gold\n- `color16` = Brown\n- `color17` = Olive\n",
					Computed:            true,
				},
				"message": dsschema.StringAttribute{
					MarkdownDescription: "Message",
					Computed:            true,
				},
				"motd_color": dsschema.StringAttribute{
					MarkdownDescription: "The following list details the supported values and their colors.\n\n- `color1` = Red\n- `color2` = Green\n- `color3` = Blue\n- `color4` = Yellow\n- `color5` = Copper\n- `color6` = Orange\n- `color7` = Purple\n- `color8` = Gray\n- `color9` = Light Green\n- `color10` = Cyan\n- `color11` = Light Gray\n- `color12` = Blue Gray\n- `color13` = Lime\n- `color14` = Black\n- `color15` = Gold\n- `color16` = Brown\n- `color17` = Olive\n",
					Computed:            true,
				},
				"motd_do_not_display_again": dsschema.BoolAttribute{
					MarkdownDescription: "Motd do not display again",
					Computed:            true,
				},
				"motd_enable": dsschema.BoolAttribute{
					MarkdownDescription: "Motd enable",
					Computed:            true,
				},
				"motd_title": dsschema.StringAttribute{
					MarkdownDescription: "Motd title",
					Computed:            true,
				},
				"severity": dsschema.StringAttribute{
					MarkdownDescription: "Severity",
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

// MotdBannerSettingsListModel represents the data model for a list data source.
type MotdBannerSettingsListModel struct {
	Tfid    types.String         `tfsdk:"tfid"`
	Data    []MotdBannerSettings `tfsdk:"data"`
	Limit   types.Int64          `tfsdk:"limit"`
	Offset  types.Int64          `tfsdk:"offset"`
	Name    types.String         `tfsdk:"name"`
	Total   types.Int64          `tfsdk:"total"`
	Folder  types.String         `tfsdk:"folder"`
	Snippet types.String         `tfsdk:"snippet"`
	Device  types.String         `tfsdk:"device"`
}

// MotdBannerSettingsListDataSourceSchema defines the schema for a list data source.
var MotdBannerSettingsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: MotdBannerSettingsDataSourceSchema.Attributes,
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
