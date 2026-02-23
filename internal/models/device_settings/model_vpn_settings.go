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

// VpnSettings represents the Terraform model for VpnSettings
type VpnSettings struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Device  basetypes.StringValue `tfsdk:"device"`
	Folder  basetypes.StringValue `tfsdk:"folder"`
	Id      basetypes.StringValue `tfsdk:"id"`
	Snippet basetypes.StringValue `tfsdk:"snippet"`
	Vpn     basetypes.ObjectValue `tfsdk:"vpn"`
}

// VpnSettingsVpn represents a nested structure within the VpnSettings model
type VpnSettingsVpn struct {
	Ikev2 basetypes.ObjectValue `tfsdk:"ikev2"`
}

// VpnSettingsVpnIkev2 represents a nested structure within the VpnSettings model
type VpnSettingsVpnIkev2 struct {
	CertificateCacheSize basetypes.Int64Value `tfsdk:"certificate_cache_size"`
	CookieThreshold      basetypes.Int64Value `tfsdk:"cookie_threshold"`
	MaxHalfOpenedSa      basetypes.Int64Value `tfsdk:"max_half_opened_sa"`
}

// AttrTypes defines the attribute types for the VpnSettings model.
func (o VpnSettings) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":    basetypes.StringType{},
		"device":  basetypes.StringType{},
		"folder":  basetypes.StringType{},
		"id":      basetypes.StringType{},
		"snippet": basetypes.StringType{},
		"vpn": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ikev2": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"certificate_cache_size": basetypes.Int64Type{},
						"cookie_threshold":       basetypes.Int64Type{},
						"max_half_opened_sa":     basetypes.Int64Type{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of VpnSettings objects.
func (o VpnSettings) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the VpnSettingsVpn model.
func (o VpnSettingsVpn) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ikev2": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"certificate_cache_size": basetypes.Int64Type{},
				"cookie_threshold":       basetypes.Int64Type{},
				"max_half_opened_sa":     basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of VpnSettingsVpn objects.
func (o VpnSettingsVpn) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the VpnSettingsVpnIkev2 model.
func (o VpnSettingsVpnIkev2) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"certificate_cache_size": basetypes.Int64Type{},
		"cookie_threshold":       basetypes.Int64Type{},
		"max_half_opened_sa":     basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of VpnSettingsVpnIkev2 objects.
func (o VpnSettingsVpnIkev2) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// VpnSettingsResourceSchema defines the schema for VpnSettings resource
var VpnSettingsResourceSchema = schema.Schema{
	MarkdownDescription: "VpnSetting resource",
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
		"vpn": schema.SingleNestedAttribute{
			MarkdownDescription: "Vpn",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"ikev2": schema.SingleNestedAttribute{
					MarkdownDescription: "Ikev2",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"certificate_cache_size": schema.Int64Attribute{
							Validators: []validator.Int64{
								int64validator.Between(0, 4000),
							},
							MarkdownDescription: "Maximum cached certificates",
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(500),
						},
						"cookie_threshold": schema.Int64Attribute{
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
							MarkdownDescription: "Cookie activation threshold",
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(500),
						},
						"max_half_opened_sa": schema.Int64Attribute{
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
							MarkdownDescription: "Maximum half-opened SA",
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(65535),
						},
					},
				},
			},
		},
	},
}

// VpnSettingsDataSourceSchema defines the schema for VpnSettings data source
var VpnSettingsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "VpnSetting data source",
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
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"vpn": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Vpn",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"ikev2": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Ikev2",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"certificate_cache_size": dsschema.Int64Attribute{
							MarkdownDescription: "Maximum cached certificates",
							Computed:            true,
						},
						"cookie_threshold": dsschema.Int64Attribute{
							MarkdownDescription: "Cookie activation threshold",
							Computed:            true,
						},
						"max_half_opened_sa": dsschema.Int64Attribute{
							MarkdownDescription: "Maximum half-opened SA",
							Computed:            true,
						},
					},
				},
			},
		},
	},
}

// VpnSettingsListModel represents the data model for a list data source.
type VpnSettingsListModel struct {
	Tfid    types.String  `tfsdk:"tfid"`
	Data    []VpnSettings `tfsdk:"data"`
	Limit   types.Int64   `tfsdk:"limit"`
	Offset  types.Int64   `tfsdk:"offset"`
	Name    types.String  `tfsdk:"name"`
	Total   types.Int64   `tfsdk:"total"`
	Folder  types.String  `tfsdk:"folder"`
	Snippet types.String  `tfsdk:"snippet"`
	Device  types.String  `tfsdk:"device"`
}

// VpnSettingsListDataSourceSchema defines the schema for a list data source.
var VpnSettingsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: VpnSettingsDataSourceSchema.Attributes,
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
