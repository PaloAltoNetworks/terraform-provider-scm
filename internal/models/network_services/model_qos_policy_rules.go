package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// Package: network_services
// This file contains models for the network_services SDK package

// QosPolicyRules represents the Terraform model for QosPolicyRules
type QosPolicyRules struct {
	Tfid             types.String          `tfsdk:"tfid"`
	RelativePosition basetypes.StringValue `tfsdk:"relative_position"`
	TargetRule       basetypes.StringValue `tfsdk:"target_rule"`
	Action           basetypes.ObjectValue `tfsdk:"action"`
	Description      basetypes.StringValue `tfsdk:"description"`
	Device           basetypes.StringValue `tfsdk:"device"`
	DscpTos          basetypes.ObjectValue `tfsdk:"dscp_tos"`
	Folder           basetypes.StringValue `tfsdk:"folder"`
	Id               basetypes.StringValue `tfsdk:"id"`
	Name             basetypes.StringValue `tfsdk:"name"`
	Schedule         basetypes.StringValue `tfsdk:"schedule"`
	Snippet          basetypes.StringValue `tfsdk:"snippet"`
	Position         basetypes.StringValue `tfsdk:"position"`
}

// QosPolicyRulesAction represents a nested structure within the QosPolicyRules model
type QosPolicyRulesAction struct {
	Class basetypes.StringValue `tfsdk:"class"`
}

// QosPolicyRulesDscpTos represents a nested structure within the QosPolicyRules model
type QosPolicyRulesDscpTos struct {
	Codepoints basetypes.ListValue `tfsdk:"codepoints"`
}

// QosPolicyRulesDscpTosCodepointsInner represents a nested structure within the QosPolicyRules model
type QosPolicyRulesDscpTosCodepointsInner struct {
	Name basetypes.StringValue `tfsdk:"name"`
	Type basetypes.ObjectValue `tfsdk:"type"`
}

// QosPolicyRulesDscpTosCodepointsInnerType represents a nested structure within the QosPolicyRules model
type QosPolicyRulesDscpTosCodepointsInnerType struct {
	Af     basetypes.ObjectValue `tfsdk:"af"`
	Cs     basetypes.ObjectValue `tfsdk:"cs"`
	Custom basetypes.ObjectValue `tfsdk:"custom"`
	Ef     basetypes.ObjectValue `tfsdk:"ef"`
	Tos    basetypes.ObjectValue `tfsdk:"tos"`
}

// QosPolicyRulesDscpTosCodepointsInnerTypeAf represents a nested structure within the QosPolicyRules model
type QosPolicyRulesDscpTosCodepointsInnerTypeAf struct {
	Codepoint basetypes.StringValue `tfsdk:"codepoint"`
}

// QosPolicyRulesDscpTosCodepointsInnerTypeCustom represents a nested structure within the QosPolicyRules model
type QosPolicyRulesDscpTosCodepointsInnerTypeCustom struct {
	Codepoint basetypes.ObjectValue `tfsdk:"codepoint"`
}

// QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint represents a nested structure within the QosPolicyRules model
type QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint struct {
	BinaryValue   basetypes.StringValue `tfsdk:"binary_value"`
	CodepointName basetypes.StringValue `tfsdk:"codepoint_name"`
}

// AttrTypes defines the attribute types for the QosPolicyRules model.
func (o QosPolicyRules) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":              basetypes.StringType{},
		"relative_position": basetypes.StringType{},
		"target_rule":       basetypes.StringType{},
		"action": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"class": basetypes.StringType{},
			},
		},
		"description": basetypes.StringType{},
		"device":      basetypes.StringType{},
		"dscp_tos": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"codepoints": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name": basetypes.StringType{},
						"type": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"af": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"codepoint": basetypes.StringType{},
									},
								},
								"cs": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"codepoint": basetypes.StringType{},
									},
								},
								"custom": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"codepoint": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"binary_value":   basetypes.StringType{},
												"codepoint_name": basetypes.StringType{},
											},
										},
									},
								},
								"ef": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{},
								},
								"tos": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"codepoint": basetypes.StringType{},
									},
								},
							},
						},
					},
				}},
			},
		},
		"folder":   basetypes.StringType{},
		"id":       basetypes.StringType{},
		"name":     basetypes.StringType{},
		"schedule": basetypes.StringType{},
		"snippet":  basetypes.StringType{},
		"position": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of QosPolicyRules objects.
func (o QosPolicyRules) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the QosPolicyRulesAction model.
func (o QosPolicyRulesAction) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"class": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of QosPolicyRulesAction objects.
func (o QosPolicyRulesAction) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the QosPolicyRulesDscpTos model.
func (o QosPolicyRulesDscpTos) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"codepoints": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name": basetypes.StringType{},
				"type": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"af": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"codepoint": basetypes.StringType{},
							},
						},
						"cs": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"codepoint": basetypes.StringType{},
							},
						},
						"custom": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"codepoint": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"binary_value":   basetypes.StringType{},
										"codepoint_name": basetypes.StringType{},
									},
								},
							},
						},
						"ef": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{},
						},
						"tos": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"codepoint": basetypes.StringType{},
							},
						},
					},
				},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of QosPolicyRulesDscpTos objects.
func (o QosPolicyRulesDscpTos) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the QosPolicyRulesDscpTosCodepointsInner model.
func (o QosPolicyRulesDscpTosCodepointsInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": basetypes.StringType{},
		"type": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"af": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"codepoint": basetypes.StringType{},
					},
				},
				"cs": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"codepoint": basetypes.StringType{},
					},
				},
				"custom": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"codepoint": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"binary_value":   basetypes.StringType{},
								"codepoint_name": basetypes.StringType{},
							},
						},
					},
				},
				"ef": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{},
				},
				"tos": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"codepoint": basetypes.StringType{},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of QosPolicyRulesDscpTosCodepointsInner objects.
func (o QosPolicyRulesDscpTosCodepointsInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the QosPolicyRulesDscpTosCodepointsInnerType model.
func (o QosPolicyRulesDscpTosCodepointsInnerType) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"af": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"codepoint": basetypes.StringType{},
			},
		},
		"cs": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"codepoint": basetypes.StringType{},
			},
		},
		"custom": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"codepoint": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"binary_value":   basetypes.StringType{},
						"codepoint_name": basetypes.StringType{},
					},
				},
			},
		},
		"ef": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{},
		},
		"tos": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"codepoint": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of QosPolicyRulesDscpTosCodepointsInnerType objects.
func (o QosPolicyRulesDscpTosCodepointsInnerType) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the QosPolicyRulesDscpTosCodepointsInnerTypeAf model.
func (o QosPolicyRulesDscpTosCodepointsInnerTypeAf) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"codepoint": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of QosPolicyRulesDscpTosCodepointsInnerTypeAf objects.
func (o QosPolicyRulesDscpTosCodepointsInnerTypeAf) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the QosPolicyRulesDscpTosCodepointsInnerTypeCustom model.
func (o QosPolicyRulesDscpTosCodepointsInnerTypeCustom) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"codepoint": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"binary_value":   basetypes.StringType{},
				"codepoint_name": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of QosPolicyRulesDscpTosCodepointsInnerTypeCustom objects.
func (o QosPolicyRulesDscpTosCodepointsInnerTypeCustom) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint model.
func (o QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"binary_value":   basetypes.StringType{},
		"codepoint_name": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint objects.
func (o QosPolicyRulesDscpTosCodepointsInnerTypeCustomCodepoint) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// QosPolicyRulesResourceSchema defines the schema for QosPolicyRules resource
var QosPolicyRulesResourceSchema = schema.Schema{
	MarkdownDescription: "QosPolicyRule resource",
	Attributes: map[string]schema.Attribute{
		"action": schema.SingleNestedAttribute{
			MarkdownDescription: "Action",
			Required:            true,
			Attributes: map[string]schema.Attribute{
				"class": schema.StringAttribute{
					MarkdownDescription: "Class",
					Optional:            true,
				},
			},
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
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
			},
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"dscp_tos": schema.SingleNestedAttribute{
			MarkdownDescription: "Dscp tos",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"codepoints": schema.ListNestedAttribute{
					MarkdownDescription: "Codepoints",
					Optional:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								MarkdownDescription: "Name",
								Optional:            true,
							},
							"type": schema.SingleNestedAttribute{
								MarkdownDescription: "Type",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"af": schema.SingleNestedAttribute{
										Validators: []validator.Object{
											objectvalidator.ConflictsWith(
												path.MatchRelative().AtParent().AtName("cs"),
												path.MatchRelative().AtParent().AtName("custom"),
												path.MatchRelative().AtParent().AtName("ef"),
												path.MatchRelative().AtParent().AtName("tos"),
											),
										},
										MarkdownDescription: "Af\n\n> ℹ️ **Note:** You must specify exactly one of `af`, `cs`, `custom`, `ef`, and `tos`.",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"codepoint": schema.StringAttribute{
												MarkdownDescription: "Codepoint",
												Optional:            true,
											},
										},
									},
									"cs": schema.SingleNestedAttribute{
										Validators: []validator.Object{
											objectvalidator.ConflictsWith(
												path.MatchRelative().AtParent().AtName("af"),
												path.MatchRelative().AtParent().AtName("custom"),
												path.MatchRelative().AtParent().AtName("ef"),
												path.MatchRelative().AtParent().AtName("tos"),
											),
										},
										MarkdownDescription: "Cs\n\n> ℹ️ **Note:** You must specify exactly one of `af`, `cs`, `custom`, `ef`, and `tos`.",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"codepoint": schema.StringAttribute{
												MarkdownDescription: "Codepoint",
												Optional:            true,
											},
										},
									},
									"custom": schema.SingleNestedAttribute{
										Validators: []validator.Object{
											objectvalidator.ConflictsWith(
												path.MatchRelative().AtParent().AtName("af"),
												path.MatchRelative().AtParent().AtName("cs"),
												path.MatchRelative().AtParent().AtName("ef"),
												path.MatchRelative().AtParent().AtName("tos"),
											),
										},
										MarkdownDescription: "Custom\n\n> ℹ️ **Note:** You must specify exactly one of `af`, `cs`, `custom`, `ef`, and `tos`.",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"codepoint": schema.SingleNestedAttribute{
												MarkdownDescription: "Codepoint",
												Optional:            true,
												Attributes: map[string]schema.Attribute{
													"binary_value": schema.StringAttribute{
														MarkdownDescription: "Binary value",
														Optional:            true,
													},
													"codepoint_name": schema.StringAttribute{
														MarkdownDescription: "Codepoint name",
														Optional:            true,
													},
												},
											},
										},
									},
									"ef": schema.SingleNestedAttribute{
										Validators: []validator.Object{
											objectvalidator.ConflictsWith(
												path.MatchRelative().AtParent().AtName("af"),
												path.MatchRelative().AtParent().AtName("cs"),
												path.MatchRelative().AtParent().AtName("custom"),
												path.MatchRelative().AtParent().AtName("tos"),
											),
										},
										MarkdownDescription: "Ef\n\n> ℹ️ **Note:** You must specify exactly one of `af`, `cs`, `custom`, `ef`, and `tos`.",
										Optional:            true,
										Attributes:          map[string]schema.Attribute{},
									},
									"tos": schema.SingleNestedAttribute{
										Validators: []validator.Object{
											objectvalidator.ConflictsWith(
												path.MatchRelative().AtParent().AtName("af"),
												path.MatchRelative().AtParent().AtName("cs"),
												path.MatchRelative().AtParent().AtName("custom"),
												path.MatchRelative().AtParent().AtName("ef"),
											),
										},
										MarkdownDescription: "Tos\n\n> ℹ️ **Note:** You must specify exactly one of `af`, `cs`, `custom`, `ef`, and `tos`.",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"codepoint": schema.StringAttribute{
												MarkdownDescription: "Codepoint",
												Optional:            true,
											},
										},
									},
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
		"name": schema.StringAttribute{
			MarkdownDescription: "Name",
			Required:            true,
		},
		"position": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("pre", "post"),
			},
			MarkdownDescription: "The relative position of the rule",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("pre"),
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"relative_position": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("before", "after", "top", "bottom"),
			},
			MarkdownDescription: "Relative positioning rule. String must be one of these: `\"before\"`, `\"after\"`, `\"top\"`, `\"bottom\"`. If not specified, rule is created at the bottom of the ruleset.",
			Optional:            true,
		},
		"schedule": schema.StringAttribute{
			MarkdownDescription: "Schedule",
			Optional:            true,
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
		"target_rule": schema.StringAttribute{
			MarkdownDescription: "The name or UUID of the rule to position this rule relative to. Required when `relative_position` is `\"before\"` or `\"after\"`.",
			Optional:            true,
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

// QosPolicyRulesDataSourceSchema defines the schema for QosPolicyRules data source
var QosPolicyRulesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "QosPolicyRule data source",
	Attributes: map[string]dsschema.Attribute{
		"action": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Action",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"class": dsschema.StringAttribute{
					MarkdownDescription: "Class",
					Computed:            true,
				},
			},
		},
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Description",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"dscp_tos": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Dscp tos",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"codepoints": dsschema.ListNestedAttribute{
					MarkdownDescription: "Codepoints",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"name": dsschema.StringAttribute{
								MarkdownDescription: "Name",
								Computed:            true,
							},
							"type": dsschema.SingleNestedAttribute{
								MarkdownDescription: "Type",
								Computed:            true,
								Attributes: map[string]dsschema.Attribute{
									"af": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Af\n\n> ℹ️ **Note:** You must specify exactly one of `af`, `cs`, `custom`, `ef`, and `tos`.",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"codepoint": dsschema.StringAttribute{
												MarkdownDescription: "Codepoint",
												Computed:            true,
											},
										},
									},
									"cs": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Cs\n\n> ℹ️ **Note:** You must specify exactly one of `af`, `cs`, `custom`, `ef`, and `tos`.",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"codepoint": dsschema.StringAttribute{
												MarkdownDescription: "Codepoint",
												Computed:            true,
											},
										},
									},
									"custom": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Custom\n\n> ℹ️ **Note:** You must specify exactly one of `af`, `cs`, `custom`, `ef`, and `tos`.",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"codepoint": dsschema.SingleNestedAttribute{
												MarkdownDescription: "Codepoint",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"binary_value": dsschema.StringAttribute{
														MarkdownDescription: "Binary value",
														Computed:            true,
													},
													"codepoint_name": dsschema.StringAttribute{
														MarkdownDescription: "Codepoint name",
														Computed:            true,
													},
												},
											},
										},
									},
									"ef": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Ef\n\n> ℹ️ **Note:** You must specify exactly one of `af`, `cs`, `custom`, `ef`, and `tos`.",
										Computed:            true,
										Attributes:          map[string]dsschema.Attribute{},
									},
									"tos": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Tos\n\n> ℹ️ **Note:** You must specify exactly one of `af`, `cs`, `custom`, `ef`, and `tos`.",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"codepoint": dsschema.StringAttribute{
												MarkdownDescription: "Codepoint",
												Computed:            true,
											},
										},
									},
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
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Name",
			Optional:            true,
			Computed:            true,
		},
		"position": dsschema.StringAttribute{
			MarkdownDescription: "The relative position of the rule",
			Computed:            true,
		},
		"relative_position": dsschema.StringAttribute{
			MarkdownDescription: "Relative positioning rule. String must be one of these: `\"before\"`, `\"after\"`, `\"top\"`, `\"bottom\"`. If not specified, rule is created at the bottom of the ruleset.",
			Computed:            true,
		},
		"schedule": dsschema.StringAttribute{
			MarkdownDescription: "Schedule",
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"target_rule": dsschema.StringAttribute{
			MarkdownDescription: "The name or UUID of the rule to position this rule relative to. Required when `relative_position` is `\"before\"` or `\"after\"`.",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// QosPolicyRulesListModel represents the data model for a list data source.
type QosPolicyRulesListModel struct {
	Tfid     types.String          `tfsdk:"tfid"`
	Data     []QosPolicyRules      `tfsdk:"data"`
	Limit    types.Int64           `tfsdk:"limit"`
	Offset   types.Int64           `tfsdk:"offset"`
	Name     types.String          `tfsdk:"name"`
	Total    types.Int64           `tfsdk:"total"`
	Folder   types.String          `tfsdk:"folder"`
	Snippet  types.String          `tfsdk:"snippet"`
	Device   types.String          `tfsdk:"device"`
	Position basetypes.StringValue `tfsdk:"position"`
}

// QosPolicyRulesListDataSourceSchema defines the schema for a list data source.
var QosPolicyRulesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: QosPolicyRulesDataSourceSchema.Attributes,
			},
		},
		"limit":   dsschema.Int64Attribute{Description: "The max number of items to return. Default: 200.", Optional: true},
		"offset":  dsschema.Int64Attribute{Description: "The offset of the first item to return.", Optional: true},
		"name":    dsschema.StringAttribute{Description: "The name of the item.", Optional: true},
		"total":   dsschema.Int64Attribute{Description: "The total number of items.", Computed: true},
		"folder":  dsschema.StringAttribute{Description: "The folder of the item. Default: Shared.", Optional: true},
		"snippet": dsschema.StringAttribute{Description: "The snippet of the item.", Optional: true},
		"device":  dsschema.StringAttribute{Description: "The device of the item.", Optional: true},
		"position": dsschema.StringAttribute{
			Description: "The relative position of the rule",
			Required:    true,
		},
	},
}
