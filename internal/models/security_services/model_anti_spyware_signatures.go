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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: security_services
// This file contains models for the security_services SDK package

// AntiSpywareSignatures represents the Terraform model for AntiSpywareSignatures
type AntiSpywareSignatures struct {
	Tfid          types.String          `tfsdk:"tfid"`
	Bugtraq       basetypes.ListValue   `tfsdk:"bugtraq"`
	Comment       basetypes.StringValue `tfsdk:"comment"`
	Cve           basetypes.ListValue   `tfsdk:"cve"`
	DefaultAction basetypes.ObjectValue `tfsdk:"default_action"`
	Device        basetypes.StringValue `tfsdk:"device"`
	Direction     basetypes.StringValue `tfsdk:"direction"`
	Folder        basetypes.StringValue `tfsdk:"folder"`
	Id            basetypes.StringValue `tfsdk:"id"`
	Reference     basetypes.ListValue   `tfsdk:"reference"`
	Severity      basetypes.StringValue `tfsdk:"severity"`
	Signature     basetypes.ObjectValue `tfsdk:"signature"`
	Snippet       basetypes.StringValue `tfsdk:"snippet"`
	ThreatId      basetypes.Int64Value  `tfsdk:"threat_id"`
	Threatname    basetypes.StringValue `tfsdk:"threatname"`
	Vendor        basetypes.ListValue   `tfsdk:"vendor"`
}

// AntiSpywareSignaturesDefaultAction represents a nested structure within the AntiSpywareSignatures model
type AntiSpywareSignaturesDefaultAction struct {
	Alert       basetypes.ObjectValue `tfsdk:"alert"`
	Allow       basetypes.ObjectValue `tfsdk:"allow"`
	BlockIp     basetypes.ObjectValue `tfsdk:"block_ip"`
	Drop        basetypes.ObjectValue `tfsdk:"drop"`
	ResetBoth   basetypes.ObjectValue `tfsdk:"reset_both"`
	ResetClient basetypes.ObjectValue `tfsdk:"reset_client"`
	ResetServer basetypes.ObjectValue `tfsdk:"reset_server"`
}

// AntiSpywareSignaturesDefaultActionBlockIp represents a nested structure within the AntiSpywareSignatures model
type AntiSpywareSignaturesDefaultActionBlockIp struct {
	Duration basetypes.Int64Value  `tfsdk:"duration"`
	TrackBy  basetypes.StringValue `tfsdk:"track_by"`
}

// AntiSpywareSignaturesSignature represents a nested structure within the AntiSpywareSignatures model
type AntiSpywareSignaturesSignature struct {
	Combination basetypes.ObjectValue `tfsdk:"combination"`
	Standard    basetypes.ListValue   `tfsdk:"standard"`
}

// AttrTypes defines the attribute types for the AntiSpywareSignatures model.
func (o AntiSpywareSignatures) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":    basetypes.StringType{},
		"bugtraq": basetypes.ListType{ElemType: basetypes.StringType{}},
		"comment": basetypes.StringType{},
		"cve":     basetypes.ListType{ElemType: basetypes.StringType{}},
		"default_action": basetypes.ObjectType{
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
		"device":    basetypes.StringType{},
		"direction": basetypes.StringType{},
		"folder":    basetypes.StringType{},
		"id":        basetypes.StringType{},
		"reference": basetypes.ListType{ElemType: basetypes.StringType{}},
		"severity":  basetypes.StringType{},
		"signature": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"combination": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"and_condition": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name": basetypes.StringType{},
								"or_condition": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"name":      basetypes.StringType{},
										"threat_id": basetypes.StringType{},
									},
								}},
							},
						}},
						"order_free": basetypes.BoolType{},
						"time_attribute": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"interval":  basetypes.Int64Type{},
								"threshold": basetypes.Int64Type{},
								"track_by":  basetypes.StringType{},
							},
						},
					},
				},
				"standard": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"and_condition": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name": basetypes.StringType{},
								"or_condition": basetypes.ListType{ElemType: basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"name": basetypes.StringType{},
										"operator": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"equal_to": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"context": basetypes.StringType{},
														"negate":  basetypes.BoolType{},
														"qualifier": basetypes.ListType{ElemType: basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"name":  basetypes.StringType{},
																"value": basetypes.StringType{},
															},
														}},
														"value": basetypes.Int64Type{},
													},
												},
												"greater_than": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"context": basetypes.StringType{},
														"qualifier": basetypes.ListType{ElemType: basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"name":  basetypes.StringType{},
																"value": basetypes.StringType{},
															},
														}},
														"value": basetypes.Int64Type{},
													},
												},
												"less_than": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"context": basetypes.StringType{},
														"qualifier": basetypes.ListType{ElemType: basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"name":  basetypes.StringType{},
																"value": basetypes.StringType{},
															},
														}},
														"value": basetypes.Int64Type{},
													},
												},
												"pattern_match": basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"context": basetypes.StringType{},
														"negate":  basetypes.BoolType{},
														"pattern": basetypes.StringType{},
														"qualifier": basetypes.ListType{ElemType: basetypes.ObjectType{
															AttrTypes: map[string]attr.Type{
																"name":  basetypes.StringType{},
																"value": basetypes.StringType{},
															},
														}},
													},
												},
											},
										},
									},
								}},
							},
						}},
						"comment":    basetypes.StringType{},
						"name":       basetypes.StringType{},
						"order_free": basetypes.BoolType{},
						"scope":      basetypes.StringType{},
					},
				}},
			},
		},
		"snippet":    basetypes.StringType{},
		"threat_id":  basetypes.Int64Type{},
		"threatname": basetypes.StringType{},
		"vendor":     basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of AntiSpywareSignatures objects.
func (o AntiSpywareSignatures) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AntiSpywareSignaturesDefaultAction model.
func (o AntiSpywareSignaturesDefaultAction) AttrTypes() map[string]attr.Type {
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

// AttrType returns the attribute type for a list of AntiSpywareSignaturesDefaultAction objects.
func (o AntiSpywareSignaturesDefaultAction) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AntiSpywareSignaturesDefaultActionBlockIp model.
func (o AntiSpywareSignaturesDefaultActionBlockIp) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"duration": basetypes.Int64Type{},
		"track_by": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of AntiSpywareSignaturesDefaultActionBlockIp objects.
func (o AntiSpywareSignaturesDefaultActionBlockIp) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the AntiSpywareSignaturesSignature model.
func (o AntiSpywareSignaturesSignature) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"combination": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"and_condition": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name": basetypes.StringType{},
						"or_condition": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name":      basetypes.StringType{},
								"threat_id": basetypes.StringType{},
							},
						}},
					},
				}},
				"order_free": basetypes.BoolType{},
				"time_attribute": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"interval":  basetypes.Int64Type{},
						"threshold": basetypes.Int64Type{},
						"track_by":  basetypes.StringType{},
					},
				},
			},
		},
		"standard": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"and_condition": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name": basetypes.StringType{},
						"or_condition": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"name": basetypes.StringType{},
								"operator": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"equal_to": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"context": basetypes.StringType{},
												"negate":  basetypes.BoolType{},
												"qualifier": basetypes.ListType{ElemType: basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"name":  basetypes.StringType{},
														"value": basetypes.StringType{},
													},
												}},
												"value": basetypes.Int64Type{},
											},
										},
										"greater_than": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"context": basetypes.StringType{},
												"qualifier": basetypes.ListType{ElemType: basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"name":  basetypes.StringType{},
														"value": basetypes.StringType{},
													},
												}},
												"value": basetypes.Int64Type{},
											},
										},
										"less_than": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"context": basetypes.StringType{},
												"qualifier": basetypes.ListType{ElemType: basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"name":  basetypes.StringType{},
														"value": basetypes.StringType{},
													},
												}},
												"value": basetypes.Int64Type{},
											},
										},
										"pattern_match": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"context": basetypes.StringType{},
												"negate":  basetypes.BoolType{},
												"pattern": basetypes.StringType{},
												"qualifier": basetypes.ListType{ElemType: basetypes.ObjectType{
													AttrTypes: map[string]attr.Type{
														"name":  basetypes.StringType{},
														"value": basetypes.StringType{},
													},
												}},
											},
										},
									},
								},
							},
						}},
					},
				}},
				"comment":    basetypes.StringType{},
				"name":       basetypes.StringType{},
				"order_free": basetypes.BoolType{},
				"scope":      basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of AntiSpywareSignaturesSignature objects.
func (o AntiSpywareSignaturesSignature) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AntiSpywareSignaturesResourceSchema defines the schema for AntiSpywareSignatures resource
var AntiSpywareSignaturesResourceSchema = schema.Schema{
	MarkdownDescription: "AntiSpywareSignature resource",
	Attributes: map[string]schema.Attribute{
		"bugtraq": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Bugtraq",
			Optional:            true,
		},
		"comment": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(256),
			},
			MarkdownDescription: "Comment",
			Optional:            true,
		},
		"cve": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Cve",
			Optional:            true,
		},
		"default_action": schema.SingleNestedAttribute{
			MarkdownDescription: "anti spyware signature default action",
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
					MarkdownDescription: "anti spyware signature block ip",
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
		"direction": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("client2server", "server2client", "both"),
			},
			MarkdownDescription: "Direction",
			Optional:            true,
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
			MarkdownDescription: "UUID of the resource",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"reference": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Reference",
			Optional:            true,
		},
		"severity": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.OneOf("critical", "low", "high", "medium", "informational"),
			},
			MarkdownDescription: "Severity",
			Optional:            true,
		},
		"signature": schema.SingleNestedAttribute{
			MarkdownDescription: "anti spyware signature",
			Optional:            true,
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"combination": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("standard"),
						),
					},
					MarkdownDescription: "Combination",
					Optional:            true,
					Computed:            true,
					Attributes: map[string]schema.Attribute{
						"and_condition": schema.ListNestedAttribute{
							MarkdownDescription: "And condition",
							Optional:            true,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Name",
										Optional:            true,
									},
									"or_condition": schema.ListNestedAttribute{
										MarkdownDescription: "Or condition",
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													MarkdownDescription: "Name",
													Optional:            true,
												},
												"threat_id": schema.StringAttribute{
													MarkdownDescription: "Threat id",
													Optional:            true,
												},
											},
										},
									},
								},
							},
						},
						"order_free": schema.BoolAttribute{
							MarkdownDescription: "Order free",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"time_attribute": schema.SingleNestedAttribute{
							MarkdownDescription: "Time attribute",
							Optional:            true,
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"interval": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(1, 3600),
									},
									MarkdownDescription: "Interval",
									Optional:            true,
									Computed:            true,
								},
								"threshold": schema.Int64Attribute{
									Validators: []validator.Int64{
										int64validator.Between(1, 255),
									},
									MarkdownDescription: "Threshold",
									Optional:            true,
									Computed:            true,
								},
								"track_by": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("source-and-destination", "source", "destination"),
									},
									MarkdownDescription: "Track by",
									Optional:            true,
									Computed:            true,
								},
							},
						},
					},
				},
				"standard": schema.ListNestedAttribute{
					Validators: []validator.List{
						listvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("combination"),
						),
					},
					MarkdownDescription: "Standard",
					Optional:            true,
					Computed:            true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"and_condition": schema.ListNestedAttribute{
								MarkdownDescription: "And condition",
								Optional:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											MarkdownDescription: "Name",
											Optional:            true,
										},
										"or_condition": schema.ListNestedAttribute{
											MarkdownDescription: "Or condition",
											Optional:            true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														MarkdownDescription: "Name",
														Optional:            true,
													},
													"operator": schema.SingleNestedAttribute{
														MarkdownDescription: "Operator",
														Optional:            true,
														Computed:            true,
														Attributes: map[string]schema.Attribute{
															"equal_to": schema.SingleNestedAttribute{
																MarkdownDescription: "Equal to",
																Optional:            true,
																Computed:            true,
																Attributes: map[string]schema.Attribute{
																	"context": schema.StringAttribute{
																		MarkdownDescription: "Context",
																		Optional:            true,
																		Computed:            true,
																	},
																	"negate": schema.BoolAttribute{
																		MarkdownDescription: "Negate",
																		Optional:            true,
																		Computed:            true,
																		Default:             booldefault.StaticBool(false),
																	},
																	"qualifier": schema.ListNestedAttribute{
																		MarkdownDescription: "Qualifier",
																		Optional:            true,
																		Computed:            true,
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					MarkdownDescription: "Name",
																					Optional:            true,
																				},
																				"value": schema.StringAttribute{
																					MarkdownDescription: "Value",
																					Optional:            true,
																				},
																			},
																		},
																	},
																	"value": schema.Int64Attribute{
																		Validators: []validator.Int64{
																			int64validator.Between(0, 4294967295),
																		},
																		MarkdownDescription: "Value",
																		Optional:            true,
																		Computed:            true,
																	},
																},
															},
															"greater_than": schema.SingleNestedAttribute{
																MarkdownDescription: "Greater than",
																Optional:            true,
																Computed:            true,
																Attributes: map[string]schema.Attribute{
																	"context": schema.StringAttribute{
																		MarkdownDescription: "Context",
																		Optional:            true,
																		Computed:            true,
																	},
																	"qualifier": schema.ListNestedAttribute{
																		MarkdownDescription: "Qualifier",
																		Optional:            true,
																		Computed:            true,
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					MarkdownDescription: "Name",
																					Optional:            true,
																				},
																				"value": schema.StringAttribute{
																					MarkdownDescription: "Value",
																					Optional:            true,
																				},
																			},
																		},
																	},
																	"value": schema.Int64Attribute{
																		Validators: []validator.Int64{
																			int64validator.Between(0, 4294967295),
																		},
																		MarkdownDescription: "Value",
																		Optional:            true,
																		Computed:            true,
																	},
																},
															},
															"less_than": schema.SingleNestedAttribute{
																MarkdownDescription: "Less than",
																Optional:            true,
																Computed:            true,
																Attributes: map[string]schema.Attribute{
																	"context": schema.StringAttribute{
																		MarkdownDescription: "Context",
																		Optional:            true,
																		Computed:            true,
																	},
																	"qualifier": schema.ListNestedAttribute{
																		MarkdownDescription: "Qualifier",
																		Optional:            true,
																		Computed:            true,
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					MarkdownDescription: "Name",
																					Optional:            true,
																				},
																				"value": schema.StringAttribute{
																					MarkdownDescription: "Value",
																					Optional:            true,
																				},
																			},
																		},
																	},
																	"value": schema.Int64Attribute{
																		Validators: []validator.Int64{
																			int64validator.Between(0, 4294967295),
																		},
																		MarkdownDescription: "Value",
																		Optional:            true,
																		Computed:            true,
																	},
																},
															},
															"pattern_match": schema.SingleNestedAttribute{
																MarkdownDescription: "Pattern match",
																Optional:            true,
																Computed:            true,
																Attributes: map[string]schema.Attribute{
																	"context": schema.StringAttribute{
																		MarkdownDescription: "Context",
																		Optional:            true,
																		Computed:            true,
																	},
																	"negate": schema.BoolAttribute{
																		MarkdownDescription: "Negate",
																		Optional:            true,
																		Computed:            true,
																		Default:             booldefault.StaticBool(false),
																	},
																	"pattern": schema.StringAttribute{
																		MarkdownDescription: "Pattern",
																		Optional:            true,
																		Computed:            true,
																	},
																	"qualifier": schema.ListNestedAttribute{
																		MarkdownDescription: "Qualifier",
																		Optional:            true,
																		Computed:            true,
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					MarkdownDescription: "Name",
																					Optional:            true,
																				},
																				"value": schema.StringAttribute{
																					MarkdownDescription: "Value",
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
										},
									},
								},
							},
							"comment": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.LengthAtMost(256),
								},
								MarkdownDescription: "Comment",
								Optional:            true,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: "Name",
								Required:            true,
							},
							"order_free": schema.BoolAttribute{
								MarkdownDescription: "Order free",
								Optional:            true,
								Computed:            true,
								Default:             booldefault.StaticBool(false),
							},
							"scope": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("protocol-data-unit", "session"),
								},
								MarkdownDescription: "Scope",
								Optional:            true,
							},
						},
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
		"threat_id": schema.Int64Attribute{
			Validators: []validator.Int64{
				int64validator.Between(15000, 70000000),
			},
			MarkdownDescription: "threat id range <15000-18000> and <6900001-7000000>",
			Required:            true,
		},
		"threatname": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(1024),
			},
			MarkdownDescription: "Threatname",
			Required:            true,
		},
		"vendor": schema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Vendor",
			Optional:            true,
		},
	},
}

// AntiSpywareSignaturesDataSourceSchema defines the schema for AntiSpywareSignatures data source
var AntiSpywareSignaturesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "AntiSpywareSignature data source",
	Attributes: map[string]dsschema.Attribute{
		"bugtraq": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Bugtraq",
			Computed:            true,
		},
		"comment": dsschema.StringAttribute{
			MarkdownDescription: "Comment",
			Computed:            true,
		},
		"cve": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Cve",
			Computed:            true,
		},
		"default_action": dsschema.SingleNestedAttribute{
			MarkdownDescription: "anti spyware signature default action",
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
					MarkdownDescription: "anti spyware signature block ip",
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
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Computed:            true,
		},
		"direction": dsschema.StringAttribute{
			MarkdownDescription: "Direction",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"reference": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Reference",
			Computed:            true,
		},
		"severity": dsschema.StringAttribute{
			MarkdownDescription: "Severity",
			Computed:            true,
		},
		"signature": dsschema.SingleNestedAttribute{
			MarkdownDescription: "anti spyware signature",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"combination": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Combination",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"and_condition": dsschema.ListNestedAttribute{
							MarkdownDescription: "And condition",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Name",
										Computed:            true,
									},
									"or_condition": dsschema.ListNestedAttribute{
										MarkdownDescription: "Or condition",
										Computed:            true,
										NestedObject: dsschema.NestedAttributeObject{
											Attributes: map[string]dsschema.Attribute{
												"name": dsschema.StringAttribute{
													MarkdownDescription: "Name",
													Computed:            true,
												},
												"threat_id": dsschema.StringAttribute{
													MarkdownDescription: "Threat id",
													Computed:            true,
												},
											},
										},
									},
								},
							},
						},
						"order_free": dsschema.BoolAttribute{
							MarkdownDescription: "Order free",
							Computed:            true,
						},
						"time_attribute": dsschema.SingleNestedAttribute{
							MarkdownDescription: "Time attribute",
							Computed:            true,
							Attributes: map[string]dsschema.Attribute{
								"interval": dsschema.Int64Attribute{
									MarkdownDescription: "Interval",
									Computed:            true,
								},
								"threshold": dsschema.Int64Attribute{
									MarkdownDescription: "Threshold",
									Computed:            true,
								},
								"track_by": dsschema.StringAttribute{
									MarkdownDescription: "Track by",
									Computed:            true,
								},
							},
						},
					},
				},
				"standard": dsschema.ListNestedAttribute{
					MarkdownDescription: "Standard",
					Computed:            true,
					NestedObject: dsschema.NestedAttributeObject{
						Attributes: map[string]dsschema.Attribute{
							"and_condition": dsschema.ListNestedAttribute{
								MarkdownDescription: "And condition",
								Computed:            true,
								NestedObject: dsschema.NestedAttributeObject{
									Attributes: map[string]dsschema.Attribute{
										"name": dsschema.StringAttribute{
											MarkdownDescription: "Name",
											Computed:            true,
										},
										"or_condition": dsschema.ListNestedAttribute{
											MarkdownDescription: "Or condition",
											Computed:            true,
											NestedObject: dsschema.NestedAttributeObject{
												Attributes: map[string]dsschema.Attribute{
													"name": dsschema.StringAttribute{
														MarkdownDescription: "Name",
														Computed:            true,
													},
													"operator": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Operator",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"equal_to": dsschema.SingleNestedAttribute{
																MarkdownDescription: "Equal to",
																Computed:            true,
																Attributes: map[string]dsschema.Attribute{
																	"context": dsschema.StringAttribute{
																		MarkdownDescription: "Context",
																		Computed:            true,
																	},
																	"negate": dsschema.BoolAttribute{
																		MarkdownDescription: "Negate",
																		Computed:            true,
																	},
																	"qualifier": dsschema.ListNestedAttribute{
																		MarkdownDescription: "Qualifier",
																		Computed:            true,
																		NestedObject: dsschema.NestedAttributeObject{
																			Attributes: map[string]dsschema.Attribute{
																				"name": dsschema.StringAttribute{
																					MarkdownDescription: "Name",
																					Computed:            true,
																				},
																				"value": dsschema.StringAttribute{
																					MarkdownDescription: "Value",
																					Computed:            true,
																				},
																			},
																		},
																	},
																	"value": dsschema.Int64Attribute{
																		MarkdownDescription: "Value",
																		Computed:            true,
																	},
																},
															},
															"greater_than": dsschema.SingleNestedAttribute{
																MarkdownDescription: "Greater than",
																Computed:            true,
																Attributes: map[string]dsschema.Attribute{
																	"context": dsschema.StringAttribute{
																		MarkdownDescription: "Context",
																		Computed:            true,
																	},
																	"qualifier": dsschema.ListNestedAttribute{
																		MarkdownDescription: "Qualifier",
																		Computed:            true,
																		NestedObject: dsschema.NestedAttributeObject{
																			Attributes: map[string]dsschema.Attribute{
																				"name": dsschema.StringAttribute{
																					MarkdownDescription: "Name",
																					Computed:            true,
																				},
																				"value": dsschema.StringAttribute{
																					MarkdownDescription: "Value",
																					Computed:            true,
																				},
																			},
																		},
																	},
																	"value": dsschema.Int64Attribute{
																		MarkdownDescription: "Value",
																		Computed:            true,
																	},
																},
															},
															"less_than": dsschema.SingleNestedAttribute{
																MarkdownDescription: "Less than",
																Computed:            true,
																Attributes: map[string]dsschema.Attribute{
																	"context": dsschema.StringAttribute{
																		MarkdownDescription: "Context",
																		Computed:            true,
																	},
																	"qualifier": dsschema.ListNestedAttribute{
																		MarkdownDescription: "Qualifier",
																		Computed:            true,
																		NestedObject: dsschema.NestedAttributeObject{
																			Attributes: map[string]dsschema.Attribute{
																				"name": dsschema.StringAttribute{
																					MarkdownDescription: "Name",
																					Computed:            true,
																				},
																				"value": dsschema.StringAttribute{
																					MarkdownDescription: "Value",
																					Computed:            true,
																				},
																			},
																		},
																	},
																	"value": dsschema.Int64Attribute{
																		MarkdownDescription: "Value",
																		Computed:            true,
																	},
																},
															},
															"pattern_match": dsschema.SingleNestedAttribute{
																MarkdownDescription: "Pattern match",
																Computed:            true,
																Attributes: map[string]dsschema.Attribute{
																	"context": dsschema.StringAttribute{
																		MarkdownDescription: "Context",
																		Computed:            true,
																	},
																	"negate": dsschema.BoolAttribute{
																		MarkdownDescription: "Negate",
																		Computed:            true,
																	},
																	"pattern": dsschema.StringAttribute{
																		MarkdownDescription: "Pattern",
																		Computed:            true,
																	},
																	"qualifier": dsschema.ListNestedAttribute{
																		MarkdownDescription: "Qualifier",
																		Computed:            true,
																		NestedObject: dsschema.NestedAttributeObject{
																			Attributes: map[string]dsschema.Attribute{
																				"name": dsschema.StringAttribute{
																					MarkdownDescription: "Name",
																					Computed:            true,
																				},
																				"value": dsschema.StringAttribute{
																					MarkdownDescription: "Value",
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
										},
									},
								},
							},
							"comment": dsschema.StringAttribute{
								MarkdownDescription: "Comment",
								Computed:            true,
							},
							"name": dsschema.StringAttribute{
								MarkdownDescription: "Name",
								Computed:            true,
							},
							"order_free": dsschema.BoolAttribute{
								MarkdownDescription: "Order free",
								Computed:            true,
							},
							"scope": dsschema.StringAttribute{
								MarkdownDescription: "Scope",
								Computed:            true,
							},
						},
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
		"threat_id": dsschema.Int64Attribute{
			MarkdownDescription: "threat id range <15000-18000> and <6900001-7000000>",
			Computed:            true,
		},
		"threatname": dsschema.StringAttribute{
			MarkdownDescription: "Threatname",
			Computed:            true,
		},
		"vendor": dsschema.ListAttribute{
			ElementType:         types.StringType,
			MarkdownDescription: "Vendor",
			Computed:            true,
		},
	},
}

// AntiSpywareSignaturesListModel represents the data model for a list data source.
type AntiSpywareSignaturesListModel struct {
	Tfid    types.String            `tfsdk:"tfid"`
	Data    []AntiSpywareSignatures `tfsdk:"data"`
	Limit   types.Int64             `tfsdk:"limit"`
	Offset  types.Int64             `tfsdk:"offset"`
	Name    types.String            `tfsdk:"name"`
	Total   types.Int64             `tfsdk:"total"`
	Folder  types.String            `tfsdk:"folder"`
	Snippet types.String            `tfsdk:"snippet"`
	Device  types.String            `tfsdk:"device"`
}

// AntiSpywareSignaturesListDataSourceSchema defines the schema for a list data source.
var AntiSpywareSignaturesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: AntiSpywareSignaturesDataSourceSchema.Attributes,
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
