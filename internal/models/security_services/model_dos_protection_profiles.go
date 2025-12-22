package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: security_services
// This file contains models for the security_services SDK package

// DosProtectionProfiles represents the Terraform model for DosProtectionProfiles
type DosProtectionProfiles struct {
	Tfid        types.String          `tfsdk:"tfid"`
	Description basetypes.StringValue `tfsdk:"description"`
	Device      basetypes.StringValue `tfsdk:"device"`
	Flood       basetypes.ObjectValue `tfsdk:"flood"`
	Folder      basetypes.StringValue `tfsdk:"folder"`
	Id          basetypes.StringValue `tfsdk:"id"`
	Name        basetypes.StringValue `tfsdk:"name"`
	Resource    basetypes.ObjectValue `tfsdk:"resource"`
	Snippet     basetypes.StringValue `tfsdk:"snippet"`
	Type        basetypes.StringValue `tfsdk:"type"`
}

// DosProtectionProfilesFlood represents a nested structure within the DosProtectionProfiles model
type DosProtectionProfilesFlood struct {
	Icmp    basetypes.ObjectValue `tfsdk:"icmp"`
	Icmpv6  basetypes.ObjectValue `tfsdk:"icmpv6"`
	OtherIp basetypes.ObjectValue `tfsdk:"other_ip"`
	TcpSyn  basetypes.ObjectValue `tfsdk:"tcp_syn"`
	Udp     basetypes.ObjectValue `tfsdk:"udp"`
}

// DosProtectionProfilesFloodIcmp represents a nested structure within the DosProtectionProfiles model
type DosProtectionProfilesFloodIcmp struct {
	Enable basetypes.BoolValue   `tfsdk:"enable"`
	Red    basetypes.ObjectValue `tfsdk:"red"`
}

// DosProtectionProfilesFloodIcmpRed represents a nested structure within the DosProtectionProfiles model
type DosProtectionProfilesFloodIcmpRed struct {
	ActivateRate basetypes.Int64Value  `tfsdk:"activate_rate"`
	AlarmRate    basetypes.Int64Value  `tfsdk:"alarm_rate"`
	Block        basetypes.ObjectValue `tfsdk:"block"`
	MaximalRate  basetypes.Int64Value  `tfsdk:"maximal_rate"`
}

// DosProtectionProfilesFloodIcmpRedBlock represents a nested structure within the DosProtectionProfiles model
type DosProtectionProfilesFloodIcmpRedBlock struct {
	Duration basetypes.Int64Value `tfsdk:"duration"`
}

// DosProtectionProfilesFloodTcpSyn represents a nested structure within the DosProtectionProfiles model
type DosProtectionProfilesFloodTcpSyn struct {
	Enable     basetypes.BoolValue   `tfsdk:"enable"`
	Red        basetypes.ObjectValue `tfsdk:"red"`
	SynCookies basetypes.ObjectValue `tfsdk:"syn_cookies"`
}

// DosProtectionProfilesFloodTcpSynSynCookies represents a nested structure within the DosProtectionProfiles model
type DosProtectionProfilesFloodTcpSynSynCookies struct {
	ActivateRate basetypes.Int64Value  `tfsdk:"activate_rate"`
	AlarmRate    basetypes.Int64Value  `tfsdk:"alarm_rate"`
	Block        basetypes.ObjectValue `tfsdk:"block"`
	MaximalRate  basetypes.Int64Value  `tfsdk:"maximal_rate"`
}

// DosProtectionProfilesFloodTcpSynSynCookiesBlock represents a nested structure within the DosProtectionProfiles model
type DosProtectionProfilesFloodTcpSynSynCookiesBlock struct {
	Duration basetypes.Int64Value `tfsdk:"duration"`
}

// DosProtectionProfilesResource represents a nested structure within the DosProtectionProfiles model
type DosProtectionProfilesResource struct {
	Sessions basetypes.ObjectValue `tfsdk:"sessions"`
}

// DosProtectionProfilesResourceSessions represents a nested structure within the DosProtectionProfiles model
type DosProtectionProfilesResourceSessions struct {
	Enabled            basetypes.BoolValue  `tfsdk:"enabled"`
	MaxConcurrentLimit basetypes.Int64Value `tfsdk:"max_concurrent_limit"`
}

// AttrTypes defines the attribute types for the DosProtectionProfiles model.
func (o DosProtectionProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":        basetypes.StringType{},
		"description": basetypes.StringType{},
		"device":      basetypes.StringType{},
		"flood": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"icmp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.BoolType{},
						"red": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"activate_rate": basetypes.Int64Type{},
								"alarm_rate":    basetypes.Int64Type{},
								"block": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"duration": basetypes.Int64Type{},
									},
								},
								"maximal_rate": basetypes.Int64Type{},
							},
						},
					},
				},
				"icmpv6": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.BoolType{},
						"red": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"activate_rate": basetypes.Int64Type{},
								"alarm_rate":    basetypes.Int64Type{},
								"block": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"duration": basetypes.Int64Type{},
									},
								},
								"maximal_rate": basetypes.Int64Type{},
							},
						},
					},
				},
				"other_ip": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.BoolType{},
						"red": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"activate_rate": basetypes.Int64Type{},
								"alarm_rate":    basetypes.Int64Type{},
								"block": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"duration": basetypes.Int64Type{},
									},
								},
								"maximal_rate": basetypes.Int64Type{},
							},
						},
					},
				},
				"tcp_syn": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.BoolType{},
						"red": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"activate_rate": basetypes.Int64Type{},
								"alarm_rate":    basetypes.Int64Type{},
								"block": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"duration": basetypes.Int64Type{},
									},
								},
								"maximal_rate": basetypes.Int64Type{},
							},
						},
						"syn_cookies": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"activate_rate": basetypes.Int64Type{},
								"alarm_rate":    basetypes.Int64Type{},
								"block": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"duration": basetypes.Int64Type{},
									},
								},
								"maximal_rate": basetypes.Int64Type{},
							},
						},
					},
				},
				"udp": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enable": basetypes.BoolType{},
						"red": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"activate_rate": basetypes.Int64Type{},
								"alarm_rate":    basetypes.Int64Type{},
								"block": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"duration": basetypes.Int64Type{},
									},
								},
								"maximal_rate": basetypes.Int64Type{},
							},
						},
					},
				},
			},
		},
		"folder": basetypes.StringType{},
		"id":     basetypes.StringType{},
		"name":   basetypes.StringType{},
		"resource": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"sessions": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"enabled":              basetypes.BoolType{},
						"max_concurrent_limit": basetypes.Int64Type{},
					},
				},
			},
		},
		"snippet": basetypes.StringType{},
		"type":    basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of DosProtectionProfiles objects.
func (o DosProtectionProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DosProtectionProfilesFlood model.
func (o DosProtectionProfilesFlood) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"icmp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable": basetypes.BoolType{},
				"red": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"activate_rate": basetypes.Int64Type{},
						"alarm_rate":    basetypes.Int64Type{},
						"block": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"duration": basetypes.Int64Type{},
							},
						},
						"maximal_rate": basetypes.Int64Type{},
					},
				},
			},
		},
		"icmpv6": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable": basetypes.BoolType{},
				"red": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"activate_rate": basetypes.Int64Type{},
						"alarm_rate":    basetypes.Int64Type{},
						"block": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"duration": basetypes.Int64Type{},
							},
						},
						"maximal_rate": basetypes.Int64Type{},
					},
				},
			},
		},
		"other_ip": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable": basetypes.BoolType{},
				"red": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"activate_rate": basetypes.Int64Type{},
						"alarm_rate":    basetypes.Int64Type{},
						"block": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"duration": basetypes.Int64Type{},
							},
						},
						"maximal_rate": basetypes.Int64Type{},
					},
				},
			},
		},
		"tcp_syn": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable": basetypes.BoolType{},
				"red": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"activate_rate": basetypes.Int64Type{},
						"alarm_rate":    basetypes.Int64Type{},
						"block": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"duration": basetypes.Int64Type{},
							},
						},
						"maximal_rate": basetypes.Int64Type{},
					},
				},
				"syn_cookies": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"activate_rate": basetypes.Int64Type{},
						"alarm_rate":    basetypes.Int64Type{},
						"block": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"duration": basetypes.Int64Type{},
							},
						},
						"maximal_rate": basetypes.Int64Type{},
					},
				},
			},
		},
		"udp": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enable": basetypes.BoolType{},
				"red": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"activate_rate": basetypes.Int64Type{},
						"alarm_rate":    basetypes.Int64Type{},
						"block": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"duration": basetypes.Int64Type{},
							},
						},
						"maximal_rate": basetypes.Int64Type{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of DosProtectionProfilesFlood objects.
func (o DosProtectionProfilesFlood) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DosProtectionProfilesFloodIcmp model.
func (o DosProtectionProfilesFloodIcmp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable": basetypes.BoolType{},
		"red": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"activate_rate": basetypes.Int64Type{},
				"alarm_rate":    basetypes.Int64Type{},
				"block": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"duration": basetypes.Int64Type{},
					},
				},
				"maximal_rate": basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of DosProtectionProfilesFloodIcmp objects.
func (o DosProtectionProfilesFloodIcmp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DosProtectionProfilesFloodIcmpRed model.
func (o DosProtectionProfilesFloodIcmpRed) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"activate_rate": basetypes.Int64Type{},
		"alarm_rate":    basetypes.Int64Type{},
		"block": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"duration": basetypes.Int64Type{},
			},
		},
		"maximal_rate": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of DosProtectionProfilesFloodIcmpRed objects.
func (o DosProtectionProfilesFloodIcmpRed) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DosProtectionProfilesFloodIcmpRedBlock model.
func (o DosProtectionProfilesFloodIcmpRedBlock) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"duration": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of DosProtectionProfilesFloodIcmpRedBlock objects.
func (o DosProtectionProfilesFloodIcmpRedBlock) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DosProtectionProfilesFloodTcpSyn model.
func (o DosProtectionProfilesFloodTcpSyn) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enable": basetypes.BoolType{},
		"red": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"activate_rate": basetypes.Int64Type{},
				"alarm_rate":    basetypes.Int64Type{},
				"block": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"duration": basetypes.Int64Type{},
					},
				},
				"maximal_rate": basetypes.Int64Type{},
			},
		},
		"syn_cookies": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"activate_rate": basetypes.Int64Type{},
				"alarm_rate":    basetypes.Int64Type{},
				"block": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"duration": basetypes.Int64Type{},
					},
				},
				"maximal_rate": basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of DosProtectionProfilesFloodTcpSyn objects.
func (o DosProtectionProfilesFloodTcpSyn) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DosProtectionProfilesFloodTcpSynSynCookies model.
func (o DosProtectionProfilesFloodTcpSynSynCookies) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"activate_rate": basetypes.Int64Type{},
		"alarm_rate":    basetypes.Int64Type{},
		"block": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"duration": basetypes.Int64Type{},
			},
		},
		"maximal_rate": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of DosProtectionProfilesFloodTcpSynSynCookies objects.
func (o DosProtectionProfilesFloodTcpSynSynCookies) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DosProtectionProfilesFloodTcpSynSynCookiesBlock model.
func (o DosProtectionProfilesFloodTcpSynSynCookiesBlock) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"duration": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of DosProtectionProfilesFloodTcpSynSynCookiesBlock objects.
func (o DosProtectionProfilesFloodTcpSynSynCookiesBlock) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DosProtectionProfilesResource model.
func (o DosProtectionProfilesResource) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"sessions": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"enabled":              basetypes.BoolType{},
				"max_concurrent_limit": basetypes.Int64Type{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of DosProtectionProfilesResource objects.
func (o DosProtectionProfilesResource) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the DosProtectionProfilesResourceSessions model.
func (o DosProtectionProfilesResourceSessions) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"enabled":              basetypes.BoolType{},
		"max_concurrent_limit": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of DosProtectionProfilesResourceSessions objects.
func (o DosProtectionProfilesResourceSessions) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// DosProtectionProfilesResourceSchema defines the schema for DosProtectionProfiles resource
var DosProtectionProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "DosProtectionProfile resource",
	Attributes: map[string]schema.Attribute{
		"description": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(255),
			},
			MarkdownDescription: "Description",
			Optional:            true,
		},
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
		"flood": schema.SingleNestedAttribute{
			MarkdownDescription: "Flood",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"icmp": schema.SingleNestedAttribute{
					MarkdownDescription: "Icmp",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"enable": schema.BoolAttribute{
							MarkdownDescription: "Enable",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"red": schema.SingleNestedAttribute{
							MarkdownDescription: "Red",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"activate_rate": schema.Int64Attribute{
									MarkdownDescription: "Activate rate",
									Optional:            true,
								},
								"alarm_rate": schema.Int64Attribute{
									MarkdownDescription: "Alarm rate",
									Optional:            true,
								},
								"block": schema.SingleNestedAttribute{
									MarkdownDescription: "Block",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"duration": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.Between(1, 21600),
											},
											MarkdownDescription: "Duration",
											Optional:            true,
											Computed:            true,
											Default:             int64default.StaticInt64(300),
										},
									},
								},
								"maximal_rate": schema.Int64Attribute{
									MarkdownDescription: "Maximal rate",
									Optional:            true,
								},
							},
						},
					},
				},
				"icmpv6": schema.SingleNestedAttribute{
					MarkdownDescription: "Icmpv6",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"enable": schema.BoolAttribute{
							MarkdownDescription: "Enable",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"red": schema.SingleNestedAttribute{
							MarkdownDescription: "Red",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"activate_rate": schema.Int64Attribute{
									MarkdownDescription: "Activate rate",
									Optional:            true,
								},
								"alarm_rate": schema.Int64Attribute{
									MarkdownDescription: "Alarm rate",
									Optional:            true,
								},
								"block": schema.SingleNestedAttribute{
									MarkdownDescription: "Block",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"duration": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.Between(1, 21600),
											},
											MarkdownDescription: "Duration",
											Optional:            true,
											Computed:            true,
											Default:             int64default.StaticInt64(300),
										},
									},
								},
								"maximal_rate": schema.Int64Attribute{
									MarkdownDescription: "Maximal rate",
									Optional:            true,
								},
							},
						},
					},
				},
				"other_ip": schema.SingleNestedAttribute{
					MarkdownDescription: "Other ip",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"enable": schema.BoolAttribute{
							MarkdownDescription: "Enable",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"red": schema.SingleNestedAttribute{
							MarkdownDescription: "Red",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"activate_rate": schema.Int64Attribute{
									MarkdownDescription: "Activate rate",
									Optional:            true,
								},
								"alarm_rate": schema.Int64Attribute{
									MarkdownDescription: "Alarm rate",
									Optional:            true,
								},
								"block": schema.SingleNestedAttribute{
									MarkdownDescription: "Block",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"duration": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.Between(1, 21600),
											},
											MarkdownDescription: "Duration",
											Optional:            true,
											Computed:            true,
											Default:             int64default.StaticInt64(300),
										},
									},
								},
								"maximal_rate": schema.Int64Attribute{
									MarkdownDescription: "Maximal rate",
									Optional:            true,
								},
							},
						},
					},
				},
				"tcp_syn": schema.SingleNestedAttribute{
					MarkdownDescription: "Tcp syn",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"enable": schema.BoolAttribute{
							MarkdownDescription: "Enable",
							Optional:            true,
						},
						"red": schema.SingleNestedAttribute{
							MarkdownDescription: "Red",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"activate_rate": schema.Int64Attribute{
									MarkdownDescription: "Activate rate",
									Optional:            true,
								},
								"alarm_rate": schema.Int64Attribute{
									MarkdownDescription: "Alarm rate",
									Optional:            true,
								},
								"block": schema.SingleNestedAttribute{
									MarkdownDescription: "Block",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"duration": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.Between(1, 21600),
											},
											MarkdownDescription: "Duration",
											Optional:            true,
											Computed:            true,
											Default:             int64default.StaticInt64(300),
										},
									},
								},
								"maximal_rate": schema.Int64Attribute{
									MarkdownDescription: "Maximal rate",
									Optional:            true,
								},
							},
						},
						"syn_cookies": schema.SingleNestedAttribute{
							MarkdownDescription: "Syn cookies",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"activate_rate": schema.Int64Attribute{
									MarkdownDescription: "Activate rate",
									Optional:            true,
								},
								"alarm_rate": schema.Int64Attribute{
									MarkdownDescription: "Alarm rate",
									Optional:            true,
								},
								"block": schema.SingleNestedAttribute{
									MarkdownDescription: "Block",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"duration": schema.Int64Attribute{
											MarkdownDescription: "Duration",
											Optional:            true,
										},
									},
								},
								"maximal_rate": schema.Int64Attribute{
									MarkdownDescription: "Maximal rate",
									Optional:            true,
								},
							},
						},
					},
				},
				"udp": schema.SingleNestedAttribute{
					MarkdownDescription: "Udp",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"enable": schema.BoolAttribute{
							MarkdownDescription: "Enable",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"red": schema.SingleNestedAttribute{
							MarkdownDescription: "Red",
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"activate_rate": schema.Int64Attribute{
									MarkdownDescription: "Activate rate",
									Optional:            true,
								},
								"alarm_rate": schema.Int64Attribute{
									MarkdownDescription: "Alarm rate",
									Optional:            true,
								},
								"block": schema.SingleNestedAttribute{
									MarkdownDescription: "Block",
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"duration": schema.Int64Attribute{
											Validators: []validator.Int64{
												int64validator.Between(1, 21600),
											},
											MarkdownDescription: "Duration",
											Optional:            true,
											Computed:            true,
											Default:             int64default.StaticInt64(300),
										},
									},
								},
								"maximal_rate": schema.Int64Attribute{
									MarkdownDescription: "Maximal rate",
									Optional:            true,
								},
							},
						},
					},
				},
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
			MarkdownDescription: "The UUID of the DNS security profile",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(31),
			},
			MarkdownDescription: "Profile name",
			Required:            true,
		},
		"resource": schema.SingleNestedAttribute{
			MarkdownDescription: "Resource",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"sessions": schema.SingleNestedAttribute{
					MarkdownDescription: "Sessions",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"enabled": schema.BoolAttribute{
							MarkdownDescription: "Enabled",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"max_concurrent_limit": schema.Int64Attribute{
							MarkdownDescription: "Max concurrent limit",
							Optional:            true,
						},
					},
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
		"type": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("aggregate", "classified"),
			},
			MarkdownDescription: "Type",
			Required:            true,
		},
	},
}

// DosProtectionProfilesDataSourceSchema defines the schema for DosProtectionProfiles data source
var DosProtectionProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "DosProtectionProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Description",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"flood": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Flood",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"icmp": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Icmp",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"enable": dsschema.BoolAttribute{
							MarkdownDescription: "Enable",
							Computed:            true,
						},
						"red": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Red",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"activate_rate": dsschema.Int64Attribute{
									MarkdownDescription: "Activate rate",
									Computed:            true,
								},
								"alarm_rate": dsschema.Int64Attribute{
									MarkdownDescription: "Alarm rate",
									Computed:            true,
								},
								"block": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Block",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"duration": dsschema.Int64Attribute{
											MarkdownDescription: "Duration",
											Computed:            true,
										},
									},
								},
								"maximal_rate": dsschema.Int64Attribute{
									MarkdownDescription: "Maximal rate",
									Computed:            true,
								},
							},
						},
					},
				},
				"icmpv6": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Icmpv6",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"enable": dsschema.BoolAttribute{
							MarkdownDescription: "Enable",
							Computed:            true,
						},
						"red": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Red",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"activate_rate": dsschema.Int64Attribute{
									MarkdownDescription: "Activate rate",
									Computed:            true,
								},
								"alarm_rate": dsschema.Int64Attribute{
									MarkdownDescription: "Alarm rate",
									Computed:            true,
								},
								"block": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Block",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"duration": dsschema.Int64Attribute{
											MarkdownDescription: "Duration",
											Computed:            true,
										},
									},
								},
								"maximal_rate": dsschema.Int64Attribute{
									MarkdownDescription: "Maximal rate",
									Computed:            true,
								},
							},
						},
					},
				},
				"other_ip": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Other ip",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"enable": dsschema.BoolAttribute{
							MarkdownDescription: "Enable",
							Computed:            true,
						},
						"red": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Red",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"activate_rate": dsschema.Int64Attribute{
									MarkdownDescription: "Activate rate",
									Computed:            true,
								},
								"alarm_rate": dsschema.Int64Attribute{
									MarkdownDescription: "Alarm rate",
									Computed:            true,
								},
								"block": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Block",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"duration": dsschema.Int64Attribute{
											MarkdownDescription: "Duration",
											Computed:            true,
										},
									},
								},
								"maximal_rate": dsschema.Int64Attribute{
									MarkdownDescription: "Maximal rate",
									Computed:            true,
								},
							},
						},
					},
				},
				"tcp_syn": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Tcp syn",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"enable": dsschema.BoolAttribute{
							MarkdownDescription: "Enable",
							Computed:            true,
						},
						"red": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Red",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"activate_rate": dsschema.Int64Attribute{
									MarkdownDescription: "Activate rate",
									Computed:            true,
								},
								"alarm_rate": dsschema.Int64Attribute{
									MarkdownDescription: "Alarm rate",
									Computed:            true,
								},
								"block": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Block",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"duration": dsschema.Int64Attribute{
											MarkdownDescription: "Duration",
											Computed:            true,
										},
									},
								},
								"maximal_rate": dsschema.Int64Attribute{
									MarkdownDescription: "Maximal rate",
									Computed:            true,
								},
							},
						},
						"syn_cookies": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Syn cookies",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"activate_rate": dsschema.Int64Attribute{
									MarkdownDescription: "Activate rate",
									Computed:            true,
								},
								"alarm_rate": dsschema.Int64Attribute{
									MarkdownDescription: "Alarm rate",
									Computed:            true,
								},
								"block": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Block",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"duration": dsschema.Int64Attribute{
											MarkdownDescription: "Duration",
											Computed:            true,
										},
									},
								},
								"maximal_rate": dsschema.Int64Attribute{
									MarkdownDescription: "Maximal rate",
									Computed:            true,
								},
							},
						},
					},
				},
				"udp": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Udp",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"enable": dsschema.BoolAttribute{
							MarkdownDescription: "Enable",
							Computed:            true,
						},
						"red": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Red",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"activate_rate": dsschema.Int64Attribute{
									MarkdownDescription: "Activate rate",
									Computed:            true,
								},
								"alarm_rate": dsschema.Int64Attribute{
									MarkdownDescription: "Alarm rate",
									Computed:            true,
								},
								"block": dsschema.SingleNestedAttribute{
									MarkdownDescription: "Block",
									Computed:            true,
									Attributes: map[string]dsschema.Attribute{
										"duration": dsschema.Int64Attribute{
											MarkdownDescription: "Duration",
											Computed:            true,
										},
									},
								},
								"maximal_rate": dsschema.Int64Attribute{
									MarkdownDescription: "Maximal rate",
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the DNS security profile",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Profile name",
			Optional:            true,
			Computed:            true,
		},
		"resource": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Resource",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"sessions": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Sessions",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"enabled": dsschema.BoolAttribute{
							MarkdownDescription: "Enabled",
							Computed:            true,
						},
						"max_concurrent_limit": dsschema.Int64Attribute{
							MarkdownDescription: "Max concurrent limit",
							Computed:            true,
						},
					},
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
		"type": dsschema.StringAttribute{
			MarkdownDescription: "Type",
			Computed:            true,
		},
	},
}

// DosProtectionProfilesListModel represents the data model for a list data source.
type DosProtectionProfilesListModel struct {
	Tfid    types.String            `tfsdk:"tfid"`
	Data    []DosProtectionProfiles `tfsdk:"data"`
	Limit   types.Int64             `tfsdk:"limit"`
	Offset  types.Int64             `tfsdk:"offset"`
	Name    types.String            `tfsdk:"name"`
	Total   types.Int64             `tfsdk:"total"`
	Folder  types.String            `tfsdk:"folder"`
	Snippet types.String            `tfsdk:"snippet"`
	Device  types.String            `tfsdk:"device"`
}

// DosProtectionProfilesListDataSourceSchema defines the schema for a list data source.
var DosProtectionProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: DosProtectionProfilesDataSourceSchema.Attributes,
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
