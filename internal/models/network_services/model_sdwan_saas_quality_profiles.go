package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
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

// SdwanSaasQualityProfiles represents the Terraform model for SdwanSaasQualityProfiles
type SdwanSaasQualityProfiles struct {
	Tfid        types.String          `tfsdk:"tfid"`
	Device      basetypes.StringValue `tfsdk:"device"`
	Folder      basetypes.StringValue `tfsdk:"folder"`
	Id          basetypes.StringValue `tfsdk:"id"`
	MonitorMode basetypes.ObjectValue `tfsdk:"monitor_mode"`
	Name        basetypes.StringValue `tfsdk:"name"`
	Snippet     basetypes.StringValue `tfsdk:"snippet"`
}

// SdwanSaasQualityProfilesMonitorMode represents a nested structure within the SdwanSaasQualityProfiles model
type SdwanSaasQualityProfilesMonitorMode struct {
	Adaptive  basetypes.ObjectValue `tfsdk:"adaptive"`
	HttpHttps basetypes.ObjectValue `tfsdk:"http_https"`
	StaticIp  basetypes.ObjectValue `tfsdk:"static_ip"`
}

// SdwanSaasQualityProfilesMonitorModeHttpHttps represents a nested structure within the SdwanSaasQualityProfiles model
type SdwanSaasQualityProfilesMonitorModeHttpHttps struct {
	MonitoredUrl  basetypes.StringValue `tfsdk:"monitored_url"`
	ProbeInterval basetypes.Int64Value  `tfsdk:"probe_interval"`
}

// SdwanSaasQualityProfilesMonitorModeStaticIp represents a nested structure within the SdwanSaasQualityProfiles model
type SdwanSaasQualityProfilesMonitorModeStaticIp struct {
	Fqdn      basetypes.ObjectValue `tfsdk:"fqdn"`
	IpAddress basetypes.ListValue   `tfsdk:"ip_address"`
}

// SdwanSaasQualityProfilesMonitorModeStaticIpFqdn represents a nested structure within the SdwanSaasQualityProfiles model
type SdwanSaasQualityProfilesMonitorModeStaticIpFqdn struct {
	FqdnName      basetypes.StringValue `tfsdk:"fqdn_name"`
	ProbeInterval basetypes.Int64Value  `tfsdk:"probe_interval"`
}

// SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner represents a nested structure within the SdwanSaasQualityProfiles model
type SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner struct {
	Name          basetypes.StringValue `tfsdk:"name"`
	ProbeInterval basetypes.Int64Value  `tfsdk:"probe_interval"`
}

// AttrTypes defines the attribute types for the SdwanSaasQualityProfiles model.
func (o SdwanSaasQualityProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":   basetypes.StringType{},
		"device": basetypes.StringType{},
		"folder": basetypes.StringType{},
		"id":     basetypes.StringType{},
		"monitor_mode": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"adaptive": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"http_https": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"monitored_url":  basetypes.StringType{},
						"probe_interval": basetypes.Int64Type{},
					},
				},
				"static_ip": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"fqdn": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"fqdn_name":      basetypes.StringType{},
								"probe_interval": basetypes.Int64Type{},
							},
						},
						"ip_address": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":           basetypes.StringType{},
								"probe_interval": basetypes.Int64Type{},
							},
						}},
					},
				},
			},
		},
		"name":    basetypes.StringType{},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of SdwanSaasQualityProfiles objects.
func (o SdwanSaasQualityProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SdwanSaasQualityProfilesMonitorMode model.
func (o SdwanSaasQualityProfilesMonitorMode) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"adaptive": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"http_https": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"monitored_url":  basetypes.StringType{},
				"probe_interval": basetypes.Int64Type{},
			},
		},
		"static_ip": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"fqdn": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"fqdn_name":      basetypes.StringType{},
						"probe_interval": basetypes.Int64Type{},
					},
				},
				"ip_address": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":           basetypes.StringType{},
						"probe_interval": basetypes.Int64Type{},
					},
				}},
			},
		},
	}
}

// AttrType returns the attribute type for a list of SdwanSaasQualityProfilesMonitorMode objects.
func (o SdwanSaasQualityProfilesMonitorMode) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SdwanSaasQualityProfilesMonitorModeHttpHttps model.
func (o SdwanSaasQualityProfilesMonitorModeHttpHttps) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"monitored_url":  basetypes.StringType{},
		"probe_interval": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of SdwanSaasQualityProfilesMonitorModeHttpHttps objects.
func (o SdwanSaasQualityProfilesMonitorModeHttpHttps) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SdwanSaasQualityProfilesMonitorModeStaticIp model.
func (o SdwanSaasQualityProfilesMonitorModeStaticIp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"fqdn": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"fqdn_name":      basetypes.StringType{},
				"probe_interval": basetypes.Int64Type{},
			},
		},
		"ip_address": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":           basetypes.StringType{},
				"probe_interval": basetypes.Int64Type{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of SdwanSaasQualityProfilesMonitorModeStaticIp objects.
func (o SdwanSaasQualityProfilesMonitorModeStaticIp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SdwanSaasQualityProfilesMonitorModeStaticIpFqdn model.
func (o SdwanSaasQualityProfilesMonitorModeStaticIpFqdn) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"fqdn_name":      basetypes.StringType{},
		"probe_interval": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of SdwanSaasQualityProfilesMonitorModeStaticIpFqdn objects.
func (o SdwanSaasQualityProfilesMonitorModeStaticIpFqdn) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner model.
func (o SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":           basetypes.StringType{},
		"probe_interval": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner objects.
func (o SdwanSaasQualityProfilesMonitorModeStaticIpIpAddressInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// SdwanSaasQualityProfilesResourceSchema defines the schema for SdwanSaasQualityProfiles resource
var SdwanSaasQualityProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "SdwanSaasQualityProfile resource",
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
		"monitor_mode": schema.SingleNestedAttribute{
			MarkdownDescription: "Monitor mode",
			Required:            true,
			Attributes: map[string]schema.Attribute{
				"adaptive": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("http_https"),
							path.MatchRelative().AtParent().AtName("static_ip"),
						),
					},
					MarkdownDescription: "Adaptive\n\n> ℹ️ **Note:** You must specify exactly one of `adaptive`, `http_https`, and `static_ip`.",
					Optional:            true,
					Attributes:          map[string]schema.Attribute{},
				},
				"http_https": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("adaptive"),
							path.MatchRelative().AtParent().AtName("static_ip"),
						),
					},
					MarkdownDescription: "Http https\n\n> ℹ️ **Note:** You must specify exactly one of `adaptive`, `http_https`, and `static_ip`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"monitored_url": schema.StringAttribute{
							MarkdownDescription: "Monitored URL",
							Required:            true,
						},
						"probe_interval": schema.Int64Attribute{
							Validators: []validator.Int64{
								int64validator.Between(1, 60),
							},
							MarkdownDescription: "Probe interval (seconds)",
							Required:            true,
						},
					},
				},
				"static_ip": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("adaptive"),
							path.MatchRelative().AtParent().AtName("http_https"),
						),
					},
					MarkdownDescription: "Static ip\n\n> ℹ️ **Note:** You must specify exactly one of `adaptive`, `http_https`, and `static_ip`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"fqdn": schema.SingleNestedAttribute{
							Validators: []validator.Object{
								objectvalidator.ConflictsWith(
									path.MatchRelative().AtParent().AtName("ip_address"),
								),
							},
							MarkdownDescription: "Fqdn\n\n> ℹ️ **Note:** You must specify exactly one of `fqdn` and `ip_address`.",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"fqdn_name": schema.StringAttribute{
									MarkdownDescription: "FQDN",
									Required:            true,
								},
								"probe_interval": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(1, 60),
									},
									MarkdownDescription: "Probe interval (seconds)",
									Required:            true,
								},
							},
						},
						"ip_address": schema.ListNestedAttribute{
							Validators: []validator.List{
								listvalidator.ConflictsWith(
									path.MatchRelative().AtParent().AtName("fqdn"),
								),
							},
							MarkdownDescription: "List of IP addresses\n\n> ℹ️ **Note:** You must specify exactly one of `fqdn` and `ip_address`.",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "IP address",
										Required:            true,
									},
									"probe_interval": schema.Int64Attribute{
										Validators: []validator.Int64{
											int64validator.Between(1, 60),
										},
										MarkdownDescription: "Probe interval (seconds)",
										Required:            true,
									},
								},
							},
						},
					},
				},
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "Profile name",
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

// SdwanSaasQualityProfilesDataSourceSchema defines the schema for SdwanSaasQualityProfiles data source
var SdwanSaasQualityProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "SdwanSaasQualityProfile data source",
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
		"monitor_mode": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Monitor mode",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"adaptive": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Adaptive\n\n> ℹ️ **Note:** You must specify exactly one of `adaptive`, `http_https`, and `static_ip`.",
					Computed:            true,
					Attributes:          map[string]dsschema.Attribute{},
				},
				"http_https": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Http https\n\n> ℹ️ **Note:** You must specify exactly one of `adaptive`, `http_https`, and `static_ip`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"monitored_url": dsschema.StringAttribute{
							MarkdownDescription: "Monitored URL",
							Computed:            true,
						},
						"probe_interval": dsschema.Int64Attribute{
							MarkdownDescription: "Probe interval (seconds)",
							Computed:            true,
						},
					},
				},
				"static_ip": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Static ip\n\n> ℹ️ **Note:** You must specify exactly one of `adaptive`, `http_https`, and `static_ip`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"fqdn": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Fqdn\n\n> ℹ️ **Note:** You must specify exactly one of `fqdn` and `ip_address`.",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"fqdn_name": dsschema.StringAttribute{
									MarkdownDescription: "FQDN",
									Computed:            true,
								},
								"probe_interval": dsschema.Int64Attribute{
									MarkdownDescription: "Probe interval (seconds)",
									Computed:            true,
								},
							},
						},
						"ip_address": dsschema.ListNestedAttribute{
							MarkdownDescription: "List of IP addresses\n\n> ℹ️ **Note:** You must specify exactly one of `fqdn` and `ip_address`.",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "IP address",
										Computed:            true,
									},
									"probe_interval": dsschema.Int64Attribute{
										MarkdownDescription: "Probe interval (seconds)",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Profile name",
			Optional:            true,
			Computed:            true,
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

// SdwanSaasQualityProfilesListModel represents the data model for a list data source.
type SdwanSaasQualityProfilesListModel struct {
	Tfid    types.String               `tfsdk:"tfid"`
	Data    []SdwanSaasQualityProfiles `tfsdk:"data"`
	Limit   types.Int64                `tfsdk:"limit"`
	Offset  types.Int64                `tfsdk:"offset"`
	Name    types.String               `tfsdk:"name"`
	Total   types.Int64                `tfsdk:"total"`
	Folder  types.String               `tfsdk:"folder"`
	Snippet types.String               `tfsdk:"snippet"`
	Device  types.String               `tfsdk:"device"`
}

// SdwanSaasQualityProfilesListDataSourceSchema defines the schema for a list data source.
var SdwanSaasQualityProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: SdwanSaasQualityProfilesDataSourceSchema.Attributes,
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
