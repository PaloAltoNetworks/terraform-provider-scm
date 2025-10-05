package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
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

// Package: security_services
// This file contains models for the security_services SDK package

// AntiSpywareProfiles represents the Terraform model for AntiSpywareProfiles
type AntiSpywareProfiles struct {
	Tfid                     types.String          `tfsdk:"tfid"`
	CloudInlineAnalysis      basetypes.BoolValue   `tfsdk:"cloud_inline_analysis"`
	Description              basetypes.StringValue `tfsdk:"description"`
	Device                   basetypes.StringValue `tfsdk:"device"`
	Folder                   basetypes.StringValue `tfsdk:"folder"`
	Id                       basetypes.StringValue `tfsdk:"id"`
	InlineExceptionEdlUrl    basetypes.ListValue   `tfsdk:"inline_exception_edl_url"`
	InlineExceptionIpAddress basetypes.ListValue   `tfsdk:"inline_exception_ip_address"`
	MicaEngineSpywareEnabled basetypes.ListValue   `tfsdk:"mica_engine_spyware_enabled"`
	Name                     basetypes.StringValue `tfsdk:"name"`
	Rules                    basetypes.ListValue   `tfsdk:"rules"`
	Snippet                  basetypes.StringValue `tfsdk:"snippet"`
	ThreatException          basetypes.ListValue   `tfsdk:"threat_exception"`
}

// AntiSpywareProfilesMicaEngineSpywareEnabledInner represents a nested structure within the AntiSpywareProfiles model
type AntiSpywareProfilesMicaEngineSpywareEnabledInner struct {
	InlinePolicyAction basetypes.StringValue `tfsdk:"inline_policy_action"`
	Name               basetypes.StringValue `tfsdk:"name"`
}

// AntiSpywareProfilesRulesInner represents a nested structure within the AntiSpywareProfiles model
type AntiSpywareProfilesRulesInner struct {
	Action        basetypes.ObjectValue `tfsdk:"action"`
	Category      basetypes.StringValue `tfsdk:"category"`
	Name          basetypes.StringValue `tfsdk:"name"`
	PacketCapture basetypes.StringValue `tfsdk:"packet_capture"`
	Severity      basetypes.ListValue   `tfsdk:"severity"`
	ThreatName    basetypes.StringValue `tfsdk:"threat_name"`
}

// AntiSpywareProfilesRulesInnerAction represents a nested structure within the AntiSpywareProfiles model
type AntiSpywareProfilesRulesInnerAction struct {
	Alert       basetypes.ObjectValue `tfsdk:"alert"`
	Allow       basetypes.ObjectValue `tfsdk:"allow"`
	BlockIp     basetypes.ObjectValue `tfsdk:"block_ip"`
	Drop        basetypes.ObjectValue `tfsdk:"drop"`
	ResetBoth   basetypes.ObjectValue `tfsdk:"reset_both"`
	ResetClient basetypes.ObjectValue `tfsdk:"reset_client"`
	ResetServer basetypes.ObjectValue `tfsdk:"reset_server"`
}

// AntiSpywareProfilesRulesInnerActionBlockIp represents a nested structure within the AntiSpywareProfiles model
type AntiSpywareProfilesRulesInnerActionBlockIp struct {
	Duration basetypes.Int64Value  `tfsdk:"duration"`
	TrackBy  basetypes.StringValue `tfsdk:"track_by"`
}

// AntiSpywareProfilesThreatExceptionInner represents a nested structure within the AntiSpywareProfiles model
type AntiSpywareProfilesThreatExceptionInner struct {
	Action        basetypes.ObjectValue `tfsdk:"action"`
	ExemptIp      basetypes.ListValue   `tfsdk:"exempt_ip"`
	Name          basetypes.StringValue `tfsdk:"name"`
	Notes         basetypes.StringValue `tfsdk:"notes"`
	PacketCapture basetypes.StringValue `tfsdk:"packet_capture"`
}

// AntiSpywareProfilesThreatExceptionInnerAction represents a nested structure within the AntiSpywareProfiles model
type AntiSpywareProfilesThreatExceptionInnerAction struct {
	Alert       basetypes.ObjectValue `tfsdk:"alert"`
	Allow       basetypes.ObjectValue `tfsdk:"allow"`
	BlockIp     basetypes.ObjectValue `tfsdk:"block_ip"`
	Default     basetypes.ObjectValue `tfsdk:"default"`
	Drop        basetypes.ObjectValue `tfsdk:"drop"`
	ResetBoth   basetypes.ObjectValue `tfsdk:"reset_both"`
	ResetClient basetypes.ObjectValue `tfsdk:"reset_client"`
	ResetServer basetypes.ObjectValue `tfsdk:"reset_server"`
}

// AntiSpywareProfilesThreatExceptionInnerActionBlockIp represents a nested structure within the AntiSpywareProfiles model
type AntiSpywareProfilesThreatExceptionInnerActionBlockIp struct {
	Duration basetypes.Int64Value  `tfsdk:"duration"`
	TrackBy  basetypes.StringValue `tfsdk:"track_by"`
}

// AntiSpywareProfilesThreatExceptionInnerExemptIpInner represents a nested structure within the AntiSpywareProfiles model
type AntiSpywareProfilesThreatExceptionInnerExemptIpInner struct {
	Name basetypes.StringValue `tfsdk:"name"`
}

// AttrTypes defines the attribute types for the AntiSpywareProfiles model.
func (o AntiSpywareProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":                        basetypes.StringType{},
		"cloud_inline_analysis":       basetypes.BoolType{},
		"description":                 basetypes.StringType{},
		"device":                      basetypes.StringType{},
		"folder":                      basetypes.StringType{},
		"id":                          basetypes.StringType{},
		"inline_exception_edl_url":    basetypes.ListType{ElemType: basetypes.StringType{}},
		"inline_exception_ip_address": basetypes.ListType{ElemType: basetypes.StringType{}},
		"mica_engine_spyware_enabled": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"inline_policy_action": basetypes.StringType{},
				"name":                 basetypes.StringType{},
			},
		}},
		"name": basetypes.StringType{},
		"rules": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"alert": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"allow": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"block_ip": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"duration": basetypes.Int64Type{},
								"track_by": basetypes.StringType{},
							},
						},
						"drop": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"reset_both": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"reset_client": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"reset_server": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
				"category":       basetypes.StringType{},
				"name":           basetypes.StringType{},
				"packet_capture": basetypes.StringType{},
				"severity":       basetypes.ListType{ElemType: basetypes.StringType{}},
				"threat_name":    basetypes.StringType{},
			},
		}},
		"snippet": basetypes.StringType{},
		"threat_exception": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"alert": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"allow": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"block_ip": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"duration": basetypes.Int64Type{},
								"track_by": basetypes.StringType{},
							},
						},
						"default": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"drop": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"reset_both": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"reset_client": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"reset_server": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
					},
				},
				"exempt_ip": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name": basetypes.StringType{},
					},
				}},
				"name":           basetypes.StringType{},
				"notes":          basetypes.StringType{},
				"packet_capture": basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of AntiSpywareProfiles objects.
func (o AntiSpywareProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AntiSpywareProfilesMicaEngineSpywareEnabledInner model.
func (o AntiSpywareProfilesMicaEngineSpywareEnabledInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"inline_policy_action": basetypes.StringType{},
		"name":                 basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AntiSpywareProfilesMicaEngineSpywareEnabledInner objects.
func (o AntiSpywareProfilesMicaEngineSpywareEnabledInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AntiSpywareProfilesRulesInner model.
func (o AntiSpywareProfilesRulesInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"alert": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"allow": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"block_ip": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"duration": basetypes.Int64Type{},
						"track_by": basetypes.StringType{},
					},
				},
				"drop": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"reset_both": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"reset_client": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"reset_server": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
		"category":       basetypes.StringType{},
		"name":           basetypes.StringType{},
		"packet_capture": basetypes.StringType{},
		"severity":       basetypes.ListType{ElemType: basetypes.StringType{}},
		"threat_name":    basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AntiSpywareProfilesRulesInner objects.
func (o AntiSpywareProfilesRulesInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AntiSpywareProfilesRulesInnerAction model.
func (o AntiSpywareProfilesRulesInnerAction) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"alert": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"allow": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"block_ip": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"duration": basetypes.Int64Type{},
				"track_by": basetypes.StringType{},
			},
		},
		"drop": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"reset_both": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"reset_client": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"reset_server": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of AntiSpywareProfilesRulesInnerAction objects.
func (o AntiSpywareProfilesRulesInnerAction) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AntiSpywareProfilesRulesInnerActionBlockIp model.
func (o AntiSpywareProfilesRulesInnerActionBlockIp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"duration": basetypes.Int64Type{},
		"track_by": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AntiSpywareProfilesRulesInnerActionBlockIp objects.
func (o AntiSpywareProfilesRulesInnerActionBlockIp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AntiSpywareProfilesThreatExceptionInner model.
func (o AntiSpywareProfilesThreatExceptionInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"alert": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"allow": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"block_ip": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"duration": basetypes.Int64Type{},
						"track_by": basetypes.StringType{},
					},
				},
				"default": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"drop": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"reset_both": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"reset_client": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"reset_server": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
			},
		},
		"exempt_ip": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name": basetypes.StringType{},
			},
		}},
		"name":           basetypes.StringType{},
		"notes":          basetypes.StringType{},
		"packet_capture": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AntiSpywareProfilesThreatExceptionInner objects.
func (o AntiSpywareProfilesThreatExceptionInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AntiSpywareProfilesThreatExceptionInnerAction model.
func (o AntiSpywareProfilesThreatExceptionInnerAction) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"alert": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"allow": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"block_ip": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"duration": basetypes.Int64Type{},
				"track_by": basetypes.StringType{},
			},
		},
		"default": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"drop": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"reset_both": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"reset_client": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"reset_server": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
	}
}

// AttrType returns the attribute type for a list of AntiSpywareProfilesThreatExceptionInnerAction objects.
func (o AntiSpywareProfilesThreatExceptionInnerAction) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AntiSpywareProfilesThreatExceptionInnerActionBlockIp model.
func (o AntiSpywareProfilesThreatExceptionInnerActionBlockIp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"duration": basetypes.Int64Type{},
		"track_by": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AntiSpywareProfilesThreatExceptionInnerActionBlockIp objects.
func (o AntiSpywareProfilesThreatExceptionInnerActionBlockIp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AntiSpywareProfilesThreatExceptionInnerExemptIpInner model.
func (o AntiSpywareProfilesThreatExceptionInnerExemptIpInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AntiSpywareProfilesThreatExceptionInnerExemptIpInner objects.
func (o AntiSpywareProfilesThreatExceptionInnerExemptIpInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AntiSpywareProfilesResourceSchema defines the schema for AntiSpywareProfiles resource
var AntiSpywareProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "AntiSpywareProfile resource",
	Attributes: map[string]schema.Attribute{
		"cloud_inline_analysis": schema.BoolAttribute{
			MarkdownDescription: "Cloud inline analysis",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
		},
		"description": schema.StringAttribute{
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
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
			},
			MarkdownDescription: "The device in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"folder": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("snippet"),
					path.MatchRelative().AtParent().AtName("device"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
			},
			MarkdownDescription: "The folder in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the anti-spyware profile",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"inline_exception_edl_url": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Inline exception edl url",
			Optional:            true,
		},
		"inline_exception_ip_address": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Inline exception ip address",
			Optional:            true,
		},
		"mica_engine_spyware_enabled": schema.ListNestedAttribute{
			MarkdownDescription: "Mica engine spyware enabled",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"inline_policy_action": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("alert", "allow", "drop", "reset-both", "reset-client", "reset-server"),
						},
						MarkdownDescription: "Inline policy action",
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("alert"),
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "Name",
						Optional:            true,
					},
				},
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the anti-spyware profile",
			Required:            true,
		},
		"rules": schema.ListNestedAttribute{
			MarkdownDescription: "Rules",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"action": schema.SingleNestedAttribute{
						MarkdownDescription: "anti spyware profiles rules default action",
						Optional:            true,
						Attributes: map[string]schema.Attribute{
							"alert": schema.SingleNestedAttribute{
								Validators: []validator.Object{
									objectvalidator.ExactlyOneOf(
										path.MatchRelative().AtParent().AtName("allow"),
										path.MatchRelative().AtParent().AtName("drop"),
										path.MatchRelative().AtParent().AtName("reset_client"),
										path.MatchRelative().AtParent().AtName("reset_server"),
										path.MatchRelative().AtParent().AtName("reset_both"),
										path.MatchRelative().AtParent().AtName("block_ip"),
									),
								},
								MarkdownDescription: "Alert",
								Optional:            true,
								Attributes:          map[string]schema.Attribute{},
							},
							"allow": schema.SingleNestedAttribute{
								Validators: []validator.Object{
									objectvalidator.ExactlyOneOf(
										path.MatchRelative().AtParent().AtName("alert"),
										path.MatchRelative().AtParent().AtName("drop"),
										path.MatchRelative().AtParent().AtName("reset_client"),
										path.MatchRelative().AtParent().AtName("reset_server"),
										path.MatchRelative().AtParent().AtName("reset_both"),
										path.MatchRelative().AtParent().AtName("block_ip"),
									),
								},
								MarkdownDescription: "Allow",
								Optional:            true,
								Attributes:          map[string]schema.Attribute{},
							},
							"block_ip": schema.SingleNestedAttribute{
								Validators: []validator.Object{
									objectvalidator.ExactlyOneOf(
										path.MatchRelative().AtParent().AtName("allow"),
										path.MatchRelative().AtParent().AtName("alert"),
										path.MatchRelative().AtParent().AtName("drop"),
										path.MatchRelative().AtParent().AtName("reset_client"),
										path.MatchRelative().AtParent().AtName("reset_server"),
										path.MatchRelative().AtParent().AtName("reset_both"),
									),
								},
								MarkdownDescription: "anti spyware profiles rules action block ip",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"duration": schema.Int64Attribute{
										Validators: []validator.Int64{
											int64validator.Between(1, 3600),
										},
										MarkdownDescription: "Duration",
										Optional:            true,
									},
									"track_by": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("source-and-destination", "source"),
										},
										MarkdownDescription: "Track by",
										Optional:            true,
									},
								},
							},
							"drop": schema.SingleNestedAttribute{
								Validators: []validator.Object{
									objectvalidator.ExactlyOneOf(
										path.MatchRelative().AtParent().AtName("allow"),
										path.MatchRelative().AtParent().AtName("alert"),
										path.MatchRelative().AtParent().AtName("reset_client"),
										path.MatchRelative().AtParent().AtName("reset_server"),
										path.MatchRelative().AtParent().AtName("reset_both"),
										path.MatchRelative().AtParent().AtName("block_ip"),
									),
								},
								MarkdownDescription: "Drop",
								Optional:            true,
								Attributes:          map[string]schema.Attribute{},
							},
							"reset_both": schema.SingleNestedAttribute{
								Validators: []validator.Object{
									objectvalidator.ExactlyOneOf(
										path.MatchRelative().AtParent().AtName("allow"),
										path.MatchRelative().AtParent().AtName("alert"),
										path.MatchRelative().AtParent().AtName("drop"),
										path.MatchRelative().AtParent().AtName("reset_client"),
										path.MatchRelative().AtParent().AtName("reset_server"),
										path.MatchRelative().AtParent().AtName("block_ip"),
									),
								},
								MarkdownDescription: "Reset both",
								Optional:            true,
								Attributes:          map[string]schema.Attribute{},
							},
							"reset_client": schema.SingleNestedAttribute{
								Validators: []validator.Object{
									objectvalidator.ExactlyOneOf(
										path.MatchRelative().AtParent().AtName("allow"),
										path.MatchRelative().AtParent().AtName("alert"),
										path.MatchRelative().AtParent().AtName("drop"),
										path.MatchRelative().AtParent().AtName("reset_server"),
										path.MatchRelative().AtParent().AtName("reset_both"),
										path.MatchRelative().AtParent().AtName("block_ip"),
									),
								},
								MarkdownDescription: "Reset client",
								Optional:            true,
								Attributes:          map[string]schema.Attribute{},
							},
							"reset_server": schema.SingleNestedAttribute{
								Validators: []validator.Object{
									objectvalidator.ExactlyOneOf(
										path.MatchRelative().AtParent().AtName("allow"),
										path.MatchRelative().AtParent().AtName("alert"),
										path.MatchRelative().AtParent().AtName("drop"),
										path.MatchRelative().AtParent().AtName("reset_client"),
										path.MatchRelative().AtParent().AtName("reset_both"),
										path.MatchRelative().AtParent().AtName("block_ip"),
									),
								},
								MarkdownDescription: "Reset server",
								Optional:            true,
								Attributes:          map[string]schema.Attribute{},
							},
						},
					},
					"category": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("dns-proxy", "backdoor", "data-theft", "autogen", "spyware", "dns-security", "downloader", "dns-phishing", "phishing-kit", "cryptominer", "hacktool", "dns-benign", "dns-wildfire", "botnet", "dns-grayware", "inline-cloud-c2", "keylogger", "p2p-communication", "domain-edl", "webshell", "command-and-control", "dns-ddns", "net-worm", "any", "tls-fingerprint", "dns-new-domain", "dns", "fraud", "dns-c2", "adware", "post-exploitation", "dns-malware", "browser-hijack", "dns-parked"),
						},
						MarkdownDescription: "Category",
						Optional:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "Name",
						Optional:            true,
					},
					"packet_capture": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("disable", "single-packet", "extended-capture"),
						},
						MarkdownDescription: "Packet capture",
						Optional:            true,
					},
					"severity": schema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "Severity",
						Optional:            true,
					},
					"threat_name": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(4),
						},
						MarkdownDescription: "Threat name",
						Optional:            true,
					},
				},
			},
		},
		"snippet": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("folder"),
					path.MatchRelative().AtParent().AtName("device"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d-_\\. ]+$"),
			},
			MarkdownDescription: "The snippet in which the resource is defined",
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
		"threat_exception": schema.ListNestedAttribute{
			MarkdownDescription: "Threat exception",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"action": schema.SingleNestedAttribute{
						MarkdownDescription: "anti spyware profiles threat exception default action",
						Optional:            true,
						Attributes: map[string]schema.Attribute{
							"alert": schema.SingleNestedAttribute{
								Validators: []validator.Object{
									objectvalidator.ExactlyOneOf(
										path.MatchRelative().AtParent().AtName("default"),
										path.MatchRelative().AtParent().AtName("allow"),
										path.MatchRelative().AtParent().AtName("drop"),
										path.MatchRelative().AtParent().AtName("reset_client"),
										path.MatchRelative().AtParent().AtName("reset_server"),
										path.MatchRelative().AtParent().AtName("reset_both"),
										path.MatchRelative().AtParent().AtName("block_ip"),
									),
								},
								MarkdownDescription: "Alert",
								Optional:            true,
								Attributes:          map[string]schema.Attribute{},
							},
							"allow": schema.SingleNestedAttribute{
								Validators: []validator.Object{
									objectvalidator.ExactlyOneOf(
										path.MatchRelative().AtParent().AtName("default"),
										path.MatchRelative().AtParent().AtName("alert"),
										path.MatchRelative().AtParent().AtName("drop"),
										path.MatchRelative().AtParent().AtName("reset_client"),
										path.MatchRelative().AtParent().AtName("reset_server"),
										path.MatchRelative().AtParent().AtName("reset_both"),
										path.MatchRelative().AtParent().AtName("block_ip"),
									),
								},
								MarkdownDescription: "Allow",
								Optional:            true,
								Attributes:          map[string]schema.Attribute{},
							},
							"block_ip": schema.SingleNestedAttribute{
								Validators: []validator.Object{
									objectvalidator.ExactlyOneOf(
										path.MatchRelative().AtParent().AtName("default"),
										path.MatchRelative().AtParent().AtName("allow"),
										path.MatchRelative().AtParent().AtName("alert"),
										path.MatchRelative().AtParent().AtName("drop"),
										path.MatchRelative().AtParent().AtName("reset_client"),
										path.MatchRelative().AtParent().AtName("reset_server"),
										path.MatchRelative().AtParent().AtName("reset_both"),
									),
								},
								MarkdownDescription: "anti spyware profiles threat exception action block ip",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"duration": schema.Int64Attribute{
										Validators: []validator.Int64{
											int64validator.Between(1, 3600),
										},
										MarkdownDescription: "Duration",
										Optional:            true,
									},
									"track_by": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("source-and-destination", "source"),
										},
										MarkdownDescription: "Track by",
										Optional:            true,
									},
								},
							},
							"default": schema.SingleNestedAttribute{
								Validators: []validator.Object{
									objectvalidator.ExactlyOneOf(
										path.MatchRelative().AtParent().AtName("allow"),
										path.MatchRelative().AtParent().AtName("alert"),
										path.MatchRelative().AtParent().AtName("drop"),
										path.MatchRelative().AtParent().AtName("reset_client"),
										path.MatchRelative().AtParent().AtName("reset_server"),
										path.MatchRelative().AtParent().AtName("reset_both"),
										path.MatchRelative().AtParent().AtName("block_ip"),
									),
								},
								MarkdownDescription: "Default",
								Optional:            true,
								Attributes:          map[string]schema.Attribute{},
							},
							"drop": schema.SingleNestedAttribute{
								Validators: []validator.Object{
									objectvalidator.ExactlyOneOf(
										path.MatchRelative().AtParent().AtName("default"),
										path.MatchRelative().AtParent().AtName("allow"),
										path.MatchRelative().AtParent().AtName("alert"),
										path.MatchRelative().AtParent().AtName("reset_client"),
										path.MatchRelative().AtParent().AtName("reset_server"),
										path.MatchRelative().AtParent().AtName("reset_both"),
										path.MatchRelative().AtParent().AtName("block_ip"),
									),
								},
								MarkdownDescription: "Drop",
								Optional:            true,
								Attributes:          map[string]schema.Attribute{},
							},
							"reset_both": schema.SingleNestedAttribute{
								Validators: []validator.Object{
									objectvalidator.ExactlyOneOf(
										path.MatchRelative().AtParent().AtName("default"),
										path.MatchRelative().AtParent().AtName("allow"),
										path.MatchRelative().AtParent().AtName("alert"),
										path.MatchRelative().AtParent().AtName("drop"),
										path.MatchRelative().AtParent().AtName("reset_client"),
										path.MatchRelative().AtParent().AtName("reset_server"),
										path.MatchRelative().AtParent().AtName("block_ip"),
									),
								},
								MarkdownDescription: "Reset both",
								Optional:            true,
								Attributes:          map[string]schema.Attribute{},
							},
							"reset_client": schema.SingleNestedAttribute{
								Validators: []validator.Object{
									objectvalidator.ExactlyOneOf(
										path.MatchRelative().AtParent().AtName("default"),
										path.MatchRelative().AtParent().AtName("allow"),
										path.MatchRelative().AtParent().AtName("alert"),
										path.MatchRelative().AtParent().AtName("drop"),
										path.MatchRelative().AtParent().AtName("reset_server"),
										path.MatchRelative().AtParent().AtName("reset_both"),
										path.MatchRelative().AtParent().AtName("block_ip"),
									),
								},
								MarkdownDescription: "Reset client",
								Optional:            true,
								Attributes:          map[string]schema.Attribute{},
							},
							"reset_server": schema.SingleNestedAttribute{
								Validators: []validator.Object{
									objectvalidator.ExactlyOneOf(
										path.MatchRelative().AtParent().AtName("default"),
										path.MatchRelative().AtParent().AtName("allow"),
										path.MatchRelative().AtParent().AtName("alert"),
										path.MatchRelative().AtParent().AtName("drop"),
										path.MatchRelative().AtParent().AtName("reset_client"),
										path.MatchRelative().AtParent().AtName("reset_both"),
										path.MatchRelative().AtParent().AtName("block_ip"),
									),
								},
								MarkdownDescription: "Reset server",
								Optional:            true,
								Attributes:          map[string]schema.Attribute{},
							},
						},
					},
					"exempt_ip": schema.ListNestedAttribute{
						MarkdownDescription: "Exempt ip",
						Optional:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									MarkdownDescription: "Name",
									Required:            true,
								},
							},
						},
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "Name",
						Optional:            true,
					},
					"notes": schema.StringAttribute{
						MarkdownDescription: "Notes",
						Optional:            true,
					},
					"packet_capture": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("disable", "single-packet", "extended-capture"),
						},
						MarkdownDescription: "Packet capture",
						Optional:            true,
					},
				},
			},
		},
	},
}

// AntiSpywareProfilesDataSourceSchema defines the schema for AntiSpywareProfiles data source
var AntiSpywareProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "AntiSpywareProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"cloud_inline_analysis": dsschema.BoolAttribute{
			MarkdownDescription: "Cloud inline analysis",
			Computed:            true,
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Description",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the anti-spyware profile",
			Required:            true,
		},
		"inline_exception_edl_url": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Inline exception edl url",
			Computed:            true,
		},
		"inline_exception_ip_address": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Inline exception ip address",
			Computed:            true,
		},
		"mica_engine_spyware_enabled": dsschema.ListNestedAttribute{
			MarkdownDescription: "Mica engine spyware enabled",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"inline_policy_action": dsschema.StringAttribute{
						MarkdownDescription: "Inline policy action",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "Name",
						Computed:            true,
					},
				},
			},
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the anti-spyware profile",
			Optional:            true,
			Computed:            true,
		},
		"rules": dsschema.ListNestedAttribute{
			MarkdownDescription: "Rules",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"action": dsschema.SingleNestedAttribute{
						MarkdownDescription: "anti spyware profiles rules default action",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"alert": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Alert",
								Computed:            true,
								Attributes:          map[string]dsschema.Attribute{},
							},
							"allow": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Allow",
								Computed:            true,
								Attributes:          map[string]dsschema.Attribute{},
							},
							"block_ip": dsschema.SingleNestedAttribute{
								MarkdownDescription: "anti spyware profiles rules action block ip",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"duration": dsschema.Int64Attribute{
										MarkdownDescription: "Duration",
										Computed:            true,
									},
									"track_by": dsschema.StringAttribute{
										MarkdownDescription: "Track by",
										Computed:            true,
									},
								},
							},
							"drop": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Drop",
								Computed:            true,
								Attributes:          map[string]dsschema.Attribute{},
							},
							"reset_both": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Reset both",
								Computed:            true,
								Attributes:          map[string]dsschema.Attribute{},
							},
							"reset_client": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Reset client",
								Computed:            true,
								Attributes:          map[string]dsschema.Attribute{},
							},
							"reset_server": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Reset server",
								Computed:            true,
								Attributes:          map[string]dsschema.Attribute{},
							},
						},
					},
					"category": dsschema.StringAttribute{
						MarkdownDescription: "Category",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "Name",
						Computed:            true,
					},
					"packet_capture": dsschema.StringAttribute{
						MarkdownDescription: "Packet capture",
						Computed:            true,
					},
					"severity": dsschema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "Severity",
						Computed:            true,
					},
					"threat_name": dsschema.StringAttribute{
						MarkdownDescription: "Threat name",
						Computed:            true,
					},
				},
			},
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"threat_exception": dsschema.ListNestedAttribute{
			MarkdownDescription: "Threat exception",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"action": dsschema.SingleNestedAttribute{
						MarkdownDescription: "anti spyware profiles threat exception default action",
						Computed:            true,
						Attributes: map[string]dsschema.Attribute{
							"alert": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Alert",
								Computed:            true,
								Attributes:          map[string]dsschema.Attribute{},
							},
							"allow": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Allow",
								Computed:            true,
								Attributes:          map[string]dsschema.Attribute{},
							},
							"block_ip": dsschema.SingleNestedAttribute{
								MarkdownDescription: "anti spyware profiles threat exception action block ip",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"duration": dsschema.Int64Attribute{
										MarkdownDescription: "Duration",
										Computed:            true,
									},
									"track_by": dsschema.StringAttribute{
										MarkdownDescription: "Track by",
										Computed:            true,
									},
								},
							},
							"default": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Default",
								Computed:            true,
								Attributes:          map[string]dsschema.Attribute{},
							},
							"drop": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Drop",
								Computed:            true,
								Attributes:          map[string]dsschema.Attribute{},
							},
							"reset_both": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Reset both",
								Computed:            true,
								Attributes:          map[string]dsschema.Attribute{},
							},
							"reset_client": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Reset client",
								Computed:            true,
								Attributes:          map[string]dsschema.Attribute{},
							},
							"reset_server": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Reset server",
								Computed:            true,
								Attributes:          map[string]dsschema.Attribute{},
							},
						},
					},
					"exempt_ip": dsschema.ListNestedAttribute{
						MarkdownDescription: "Exempt ip",
						Computed:            true,
						NestedObject: dsschema.NestedAttributeObject{
							Attributes: map[string]dsschema.Attribute{
								"name": dsschema.StringAttribute{
									MarkdownDescription: "Name",
									Computed:            true,
								},
							},
						},
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "Name",
						Computed:            true,
					},
					"notes": dsschema.StringAttribute{
						MarkdownDescription: "Notes",
						Computed:            true,
					},
					"packet_capture": dsschema.StringAttribute{
						MarkdownDescription: "Packet capture",
						Computed:            true,
					},
				},
			},
		},
	},
}

// AntiSpywareProfilesListModel represents the data model for a list data source.
type AntiSpywareProfilesListModel struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Data    []AntiSpywareProfiles `tfsdk:"data"`
	Limit   types.Int64           `tfsdk:"limit"`
	Offset  types.Int64           `tfsdk:"offset"`
	Name    types.String          `tfsdk:"name"`
	Total   types.Int64           `tfsdk:"total"`
	Folder  types.String          `tfsdk:"folder"`
	Snippet types.String          `tfsdk:"snippet"`
	Device  types.String          `tfsdk:"device"`
}

// AntiSpywareProfilesListDataSourceSchema defines the schema for a list data source.
var AntiSpywareProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: AntiSpywareProfilesDataSourceSchema.Attributes,
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
