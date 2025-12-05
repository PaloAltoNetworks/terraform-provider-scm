package models

import (
	"regexp"

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

// Package: device_settings
// This file contains models for the device_settings SDK package

// GeneralSettings represents the Terraform model for GeneralSettings
type GeneralSettings struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Device  basetypes.StringValue `tfsdk:"device"`
	Folder  basetypes.StringValue `tfsdk:"folder"`
	General basetypes.ObjectValue `tfsdk:"general"`
	Id      basetypes.StringValue `tfsdk:"id"`
	Snippet basetypes.StringValue `tfsdk:"snippet"`
}

// GeneralSettingsGeneral represents a nested structure within the GeneralSettings model
type GeneralSettingsGeneral struct {
	AckLoginBanner       basetypes.BoolValue   `tfsdk:"ack_login_banner"`
	Domain               basetypes.StringValue `tfsdk:"domain"`
	GeoLocation          basetypes.ObjectValue `tfsdk:"geo_location"`
	Locale               basetypes.StringValue `tfsdk:"locale"`
	LoginBanner          basetypes.StringValue `tfsdk:"login_banner"`
	Setting              basetypes.ObjectValue `tfsdk:"setting"`
	SslTlsServiceProfile basetypes.StringValue `tfsdk:"ssl_tls_service_profile"`
	Timezone             basetypes.StringValue `tfsdk:"timezone"`
}

// GeneralSettingsGeneralGeoLocation represents a nested structure within the GeneralSettings model
type GeneralSettingsGeneralGeoLocation struct {
	Latitude  basetypes.Float64Value `tfsdk:"latitude"`
	Longitude basetypes.Float64Value `tfsdk:"longitude"`
}

// GeneralSettingsGeneralSetting represents a nested structure within the GeneralSettings model
type GeneralSettingsGeneralSetting struct {
	AutoMacDetect      basetypes.BoolValue   `tfsdk:"auto_mac_detect"`
	FailOpen           basetypes.BoolValue   `tfsdk:"fail_open"`
	Management         basetypes.ObjectValue `tfsdk:"management"`
	TunnelAcceleration basetypes.BoolValue   `tfsdk:"tunnel_acceleration"`
}

// GeneralSettingsGeneralSettingManagement represents a nested structure within the GeneralSettings model
type GeneralSettingsGeneralSettingManagement struct {
	AutoAcquireCommitLock            basetypes.BoolValue `tfsdk:"auto_acquire_commit_lock"`
	EnableCertificateExpirationCheck basetypes.BoolValue `tfsdk:"enable_certificate_expiration_check"`
}

// AttrTypes defines the attribute types for the GeneralSettings model.
func (o GeneralSettings) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":   basetypes.StringType{},
		"device": basetypes.StringType{},
		"folder": basetypes.StringType{},
		"general": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ack_login_banner": basetypes.BoolType{},
				"domain":           basetypes.StringType{},
				"geo_location": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"latitude":  basetypes.Float64Type{},
						"longitude": basetypes.Float64Type{},
					},
				},
				"locale":       basetypes.StringType{},
				"login_banner": basetypes.StringType{},
				"setting": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"auto_mac_detect": basetypes.BoolType{},
						"fail_open":       basetypes.BoolType{},
						"management": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"auto_acquire_commit_lock":            basetypes.BoolType{},
								"enable_certificate_expiration_check": basetypes.BoolType{},
							},
						},
						"tunnel_acceleration": basetypes.BoolType{},
					},
				},
				"ssl_tls_service_profile": basetypes.StringType{},
				"timezone":                basetypes.StringType{},
			},
		},
		"id":      basetypes.StringType{},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of GeneralSettings objects.
func (o GeneralSettings) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the GeneralSettingsGeneral model.
func (o GeneralSettingsGeneral) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ack_login_banner": basetypes.BoolType{},
		"domain":           basetypes.StringType{},
		"geo_location": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"latitude":  basetypes.Float64Type{},
				"longitude": basetypes.Float64Type{},
			},
		},
		"locale":       basetypes.StringType{},
		"login_banner": basetypes.StringType{},
		"setting": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"auto_mac_detect": basetypes.BoolType{},
				"fail_open":       basetypes.BoolType{},
				"management": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"auto_acquire_commit_lock":            basetypes.BoolType{},
						"enable_certificate_expiration_check": basetypes.BoolType{},
					},
				},
				"tunnel_acceleration": basetypes.BoolType{},
			},
		},
		"ssl_tls_service_profile": basetypes.StringType{},
		"timezone":                basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of GeneralSettingsGeneral objects.
func (o GeneralSettingsGeneral) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the GeneralSettingsGeneralGeoLocation model.
func (o GeneralSettingsGeneralGeoLocation) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"latitude":  basetypes.Float64Type{},
		"longitude": basetypes.Float64Type{},
	}
}

// AttrType returns the attribute type for a list of GeneralSettingsGeneralGeoLocation objects.
func (o GeneralSettingsGeneralGeoLocation) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the GeneralSettingsGeneralSetting model.
func (o GeneralSettingsGeneralSetting) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"auto_mac_detect": basetypes.BoolType{},
		"fail_open":       basetypes.BoolType{},
		"management": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"auto_acquire_commit_lock":            basetypes.BoolType{},
				"enable_certificate_expiration_check": basetypes.BoolType{},
			},
		},
		"tunnel_acceleration": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of GeneralSettingsGeneralSetting objects.
func (o GeneralSettingsGeneralSetting) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the GeneralSettingsGeneralSettingManagement model.
func (o GeneralSettingsGeneralSettingManagement) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"auto_acquire_commit_lock":            basetypes.BoolType{},
		"enable_certificate_expiration_check": basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of GeneralSettingsGeneralSettingManagement objects.
func (o GeneralSettingsGeneralSettingManagement) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// GeneralSettingsResourceSchema defines the schema for GeneralSettings resource
var GeneralSettingsResourceSchema = schema.Schema{
	MarkdownDescription: "GeneralSetting resource",
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
		"general": schema.SingleNestedAttribute{
			MarkdownDescription: "General",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"ack_login_banner": schema.BoolAttribute{
					MarkdownDescription: "Force admins to acknowledge login banner",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
				"domain": schema.StringAttribute{
					MarkdownDescription: "DNS domain",
					Optional:            true,
					Computed:            true,
				},
				"geo_location": schema.SingleNestedAttribute{
					MarkdownDescription: "Geographic coordinates",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"latitude": schema.Float64Attribute{
							MarkdownDescription: "Latitude",
							Required:            true,
						},
						"longitude": schema.Float64Attribute{
							MarkdownDescription: "Longitude",
							Required:            true,
						},
					},
				},
				"locale": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.OneOf("en", "es", "ja", "fr", "zh_CN", "zh_TW"),
					},
					MarkdownDescription: "Locale",
					Optional:            true,
					Computed:            true,
					Default:             stringdefault.StaticString("en"),
				},
				"login_banner": schema.StringAttribute{
					MarkdownDescription: "Logon banner",
					Optional:            true,
					Computed:            true,
				},
				"setting": schema.SingleNestedAttribute{
					MarkdownDescription: "Setting",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"auto_mac_detect": schema.BoolAttribute{
							MarkdownDescription: "Use hypervisor assigned MAC addresses",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"fail_open": schema.BoolAttribute{
							MarkdownDescription: "Fail open",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"management": schema.SingleNestedAttribute{
							MarkdownDescription: "Management",
							Optional:            true,
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"auto_acquire_commit_lock": schema.BoolAttribute{
									MarkdownDescription: "Automatically acquire commit lock",
									Optional:            true,
									Computed:            true,
									Default:             booldefault.StaticBool(false),
								},
								"enable_certificate_expiration_check": schema.BoolAttribute{
									MarkdownDescription: "Certificate expiration check",
									Optional:            true,
									Computed:            true,
									Default:             booldefault.StaticBool(false),
								},
							},
						},
						"tunnel_acceleration": schema.BoolAttribute{
							MarkdownDescription: "Tunnel acceleration",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(true),
						},
					},
				},
				"ssl_tls_service_profile": schema.StringAttribute{
					MarkdownDescription: "SSL/TLS service profile",
					Optional:            true,
					Computed:            true,
				},
				"timezone": schema.StringAttribute{
					MarkdownDescription: "Timezone",
					Optional:            true,
					Computed:            true,
				},
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

// GeneralSettingsDataSourceSchema defines the schema for GeneralSettings data source
var GeneralSettingsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "GeneralSetting data source",
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
		"general": dsschema.SingleNestedAttribute{
			MarkdownDescription: "General",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"ack_login_banner": dsschema.BoolAttribute{
					MarkdownDescription: "Force admins to acknowledge login banner",
					Computed:            true,
				},
				"domain": dsschema.StringAttribute{
					MarkdownDescription: "DNS domain",
					Computed:            true,
				},
				"geo_location": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Geographic coordinates",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"latitude": dsschema.Float64Attribute{
							MarkdownDescription: "Latitude",
							Computed:            true,
						},
						"longitude": dsschema.Float64Attribute{
							MarkdownDescription: "Longitude",
							Computed:            true,
						},
					},
				},
				"locale": dsschema.StringAttribute{
					MarkdownDescription: "Locale",
					Computed:            true,
				},
				"login_banner": dsschema.StringAttribute{
					MarkdownDescription: "Logon banner",
					Computed:            true,
				},
				"setting": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Setting",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"auto_mac_detect": dsschema.BoolAttribute{
							MarkdownDescription: "Use hypervisor assigned MAC addresses",
							Computed:            true,
						},
						"fail_open": dsschema.BoolAttribute{
							MarkdownDescription: "Fail open",
							Computed:            true,
						},
						"management": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Management",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"auto_acquire_commit_lock": dsschema.BoolAttribute{
									MarkdownDescription: "Automatically acquire commit lock",
									Computed:            true,
								},
								"enable_certificate_expiration_check": dsschema.BoolAttribute{
									MarkdownDescription: "Certificate expiration check",
									Computed:            true,
								},
							},
						},
						"tunnel_acceleration": dsschema.BoolAttribute{
							MarkdownDescription: "Tunnel acceleration",
							Computed:            true,
						},
					},
				},
				"ssl_tls_service_profile": dsschema.StringAttribute{
					MarkdownDescription: "SSL/TLS service profile",
					Computed:            true,
				},
				"timezone": dsschema.StringAttribute{
					MarkdownDescription: "Timezone",
					Computed:            true,
				},
			},
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

// GeneralSettingsListModel represents the data model for a list data source.
type GeneralSettingsListModel struct {
	Tfid    types.String      `tfsdk:"tfid"`
	Data    []GeneralSettings `tfsdk:"data"`
	Limit   types.Int64       `tfsdk:"limit"`
	Offset  types.Int64       `tfsdk:"offset"`
	Name    types.String      `tfsdk:"name"`
	Total   types.Int64       `tfsdk:"total"`
	Folder  types.String      `tfsdk:"folder"`
	Snippet types.String      `tfsdk:"snippet"`
	Device  types.String      `tfsdk:"device"`
}

// GeneralSettingsListDataSourceSchema defines the schema for a list data source.
var GeneralSettingsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: GeneralSettingsDataSourceSchema.Attributes,
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
