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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// Package: objects
// This file contains models for the objects SDK package

// Applications represents the Terraform model for Applications
type Applications struct {
	Tfid                   types.String          `tfsdk:"tfid"`
	AbleToTransferFile     basetypes.BoolValue   `tfsdk:"able_to_transfer_file"`
	AlgDisableCapability   basetypes.StringValue `tfsdk:"alg_disable_capability"`
	Category               basetypes.StringValue `tfsdk:"category"`
	ConsumeBigBandwidth    basetypes.BoolValue   `tfsdk:"consume_big_bandwidth"`
	DataIdent              basetypes.BoolValue   `tfsdk:"data_ident"`
	Default                basetypes.ObjectValue `tfsdk:"default"`
	Description            basetypes.StringValue `tfsdk:"description"`
	Device                 basetypes.StringValue `tfsdk:"device"`
	EvasiveBehavior        basetypes.BoolValue   `tfsdk:"evasive_behavior"`
	FileTypeIdent          basetypes.BoolValue   `tfsdk:"file_type_ident"`
	Folder                 basetypes.StringValue `tfsdk:"folder"`
	HasKnownVulnerability  basetypes.BoolValue   `tfsdk:"has_known_vulnerability"`
	Id                     basetypes.StringValue `tfsdk:"id"`
	Name                   basetypes.StringValue `tfsdk:"name"`
	NoAppidCaching         basetypes.BoolValue   `tfsdk:"no_appid_caching"`
	ParentApp              basetypes.StringValue `tfsdk:"parent_app"`
	PervasiveUse           basetypes.BoolValue   `tfsdk:"pervasive_use"`
	ProneToMisuse          basetypes.BoolValue   `tfsdk:"prone_to_misuse"`
	Risk                   basetypes.StringValue `tfsdk:"risk"`
	Signature              basetypes.ListValue   `tfsdk:"signature"`
	Snippet                basetypes.StringValue `tfsdk:"snippet"`
	Subcategory            basetypes.StringValue `tfsdk:"subcategory"`
	TcpHalfClosedTimeout   basetypes.Int64Value  `tfsdk:"tcp_half_closed_timeout"`
	TcpTimeWaitTimeout     basetypes.Int64Value  `tfsdk:"tcp_time_wait_timeout"`
	TcpTimeout             basetypes.Int64Value  `tfsdk:"tcp_timeout"`
	Technology             basetypes.StringValue `tfsdk:"technology"`
	Timeout                basetypes.Int64Value  `tfsdk:"timeout"`
	TunnelApplications     basetypes.BoolValue   `tfsdk:"tunnel_applications"`
	TunnelOtherApplication basetypes.BoolValue   `tfsdk:"tunnel_other_application"`
	UdpTimeout             basetypes.Int64Value  `tfsdk:"udp_timeout"`
	UsedByMalware          basetypes.BoolValue   `tfsdk:"used_by_malware"`
	VirusIdent             basetypes.BoolValue   `tfsdk:"virus_ident"`
}

// ApplicationsDefault represents a nested structure within the Applications model
type ApplicationsDefault struct {
	IdentByIcmp6Type  basetypes.ObjectValue `tfsdk:"ident_by_icmp6_type"`
	IdentByIcmpType   basetypes.ObjectValue `tfsdk:"ident_by_icmp_type"`
	IdentByIpProtocol basetypes.StringValue `tfsdk:"ident_by_ip_protocol"`
	Port              basetypes.ListValue   `tfsdk:"port"`
}

// ApplicationsDefaultIdentByIcmp6Type represents a nested structure within the Applications model
type ApplicationsDefaultIdentByIcmp6Type struct {
	Code basetypes.StringValue `tfsdk:"code"`
	Type basetypes.StringValue `tfsdk:"type"`
}

// ApplicationsSignatureInner represents a nested structure within the Applications model
type ApplicationsSignatureInner struct {
	AndCondition basetypes.ListValue   `tfsdk:"and_condition"`
	Comment      basetypes.StringValue `tfsdk:"comment"`
	Name         basetypes.StringValue `tfsdk:"name"`
	OrderFree    basetypes.BoolValue   `tfsdk:"order_free"`
	Scope        basetypes.StringValue `tfsdk:"scope"`
}

// ApplicationsSignatureInnerAndConditionInner represents a nested structure within the Applications model
type ApplicationsSignatureInnerAndConditionInner struct {
	Name        basetypes.StringValue `tfsdk:"name"`
	OrCondition basetypes.ListValue   `tfsdk:"or_condition"`
}

// ApplicationsSignatureInnerAndConditionInnerOrConditionInner represents a nested structure within the Applications model
type ApplicationsSignatureInnerAndConditionInnerOrConditionInner struct {
	Name     basetypes.StringValue `tfsdk:"name"`
	Operator basetypes.ObjectValue `tfsdk:"operator"`
}

// ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator represents a nested structure within the Applications model
type ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator struct {
	EqualTo      basetypes.ObjectValue `tfsdk:"equal_to"`
	GreaterThan  basetypes.ObjectValue `tfsdk:"greater_than"`
	LessThan     basetypes.ObjectValue `tfsdk:"less_than"`
	PatternMatch basetypes.ObjectValue `tfsdk:"pattern_match"`
}

// ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo represents a nested structure within the Applications model
type ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo struct {
	Context  basetypes.StringValue `tfsdk:"context"`
	Mask     basetypes.StringValue `tfsdk:"mask"`
	Position basetypes.StringValue `tfsdk:"position"`
	Value    basetypes.StringValue `tfsdk:"value"`
}

// ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan represents a nested structure within the Applications model
type ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan struct {
	Context   basetypes.StringValue `tfsdk:"context"`
	Qualifier basetypes.ListValue   `tfsdk:"qualifier"`
	Value     basetypes.Int64Value  `tfsdk:"value"`
}

// ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner represents a nested structure within the Applications model
type ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner struct {
	Name  basetypes.StringValue `tfsdk:"name"`
	Value basetypes.StringValue `tfsdk:"value"`
}

// ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch represents a nested structure within the Applications model
type ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch struct {
	Context   basetypes.StringValue `tfsdk:"context"`
	Pattern   basetypes.StringValue `tfsdk:"pattern"`
	Qualifier basetypes.ListValue   `tfsdk:"qualifier"`
}

// AttrTypes defines the attribute types for the Applications model.
func (o Applications) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":                   basetypes.StringType{},
		"able_to_transfer_file":  basetypes.BoolType{},
		"alg_disable_capability": basetypes.StringType{},
		"category":               basetypes.StringType{},
		"consume_big_bandwidth":  basetypes.BoolType{},
		"data_ident":             basetypes.BoolType{},
		"default": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ident_by_icmp6_type": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"code": basetypes.StringType{},
						"type": basetypes.StringType{},
					},
				},
				"ident_by_icmp_type": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"code": basetypes.StringType{},
						"type": basetypes.StringType{},
					},
				},
				"ident_by_ip_protocol": basetypes.StringType{},
				"port":                 basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		},
		"description":             basetypes.StringType{},
		"device":                  basetypes.StringType{},
		"evasive_behavior":        basetypes.BoolType{},
		"file_type_ident":         basetypes.BoolType{},
		"folder":                  basetypes.StringType{},
		"has_known_vulnerability": basetypes.BoolType{},
		"id":                      basetypes.StringType{},
		"name":                    basetypes.StringType{},
		"no_appid_caching":        basetypes.BoolType{},
		"parent_app":              basetypes.StringType{},
		"pervasive_use":           basetypes.BoolType{},
		"prone_to_misuse":         basetypes.BoolType{},
		"risk":                    basetypes.StringType{},
		"signature": basetypes.ListType{ElemType: basetypes.ObjectType{
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
												"context":  basetypes.StringType{},
												"mask":     basetypes.StringType{},
												"position": basetypes.StringType{},
												"value":    basetypes.StringType{},
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
		"snippet":                  basetypes.StringType{},
		"subcategory":              basetypes.StringType{},
		"tcp_half_closed_timeout":  basetypes.Int64Type{},
		"tcp_time_wait_timeout":    basetypes.Int64Type{},
		"tcp_timeout":              basetypes.Int64Type{},
		"technology":               basetypes.StringType{},
		"timeout":                  basetypes.Int64Type{},
		"tunnel_applications":      basetypes.BoolType{},
		"tunnel_other_application": basetypes.BoolType{},
		"udp_timeout":              basetypes.Int64Type{},
		"used_by_malware":          basetypes.BoolType{},
		"virus_ident":              basetypes.BoolType{},
	}
}

// AttrType returns the attribute type for a list of Applications objects.
func (o Applications) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ApplicationsDefault model.
func (o ApplicationsDefault) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ident_by_icmp6_type": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"code": basetypes.StringType{},
				"type": basetypes.StringType{},
			},
		},
		"ident_by_icmp_type": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"code": basetypes.StringType{},
				"type": basetypes.StringType{},
			},
		},
		"ident_by_ip_protocol": basetypes.StringType{},
		"port":                 basetypes.ListType{ElemType: basetypes.StringType{}},
	}
}

// AttrType returns the attribute type for a list of ApplicationsDefault objects.
func (o ApplicationsDefault) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ApplicationsDefaultIdentByIcmp6Type model.
func (o ApplicationsDefaultIdentByIcmp6Type) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"code": basetypes.StringType{},
		"type": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ApplicationsDefaultIdentByIcmp6Type objects.
func (o ApplicationsDefaultIdentByIcmp6Type) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ApplicationsSignatureInner model.
func (o ApplicationsSignatureInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
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
										"context":  basetypes.StringType{},
										"mask":     basetypes.StringType{},
										"position": basetypes.StringType{},
										"value":    basetypes.StringType{},
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
	}
}

// AttrType returns the attribute type for a list of ApplicationsSignatureInner objects.
func (o ApplicationsSignatureInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ApplicationsSignatureInnerAndConditionInner model.
func (o ApplicationsSignatureInnerAndConditionInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": basetypes.StringType{},
		"or_condition": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name": basetypes.StringType{},
				"operator": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"equal_to": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"context":  basetypes.StringType{},
								"mask":     basetypes.StringType{},
								"position": basetypes.StringType{},
								"value":    basetypes.StringType{},
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
	}
}

// AttrType returns the attribute type for a list of ApplicationsSignatureInnerAndConditionInner objects.
func (o ApplicationsSignatureInnerAndConditionInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ApplicationsSignatureInnerAndConditionInnerOrConditionInner model.
func (o ApplicationsSignatureInnerAndConditionInnerOrConditionInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": basetypes.StringType{},
		"operator": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"equal_to": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"context":  basetypes.StringType{},
						"mask":     basetypes.StringType{},
						"position": basetypes.StringType{},
						"value":    basetypes.StringType{},
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
	}
}

// AttrType returns the attribute type for a list of ApplicationsSignatureInnerAndConditionInnerOrConditionInner objects.
func (o ApplicationsSignatureInnerAndConditionInnerOrConditionInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator model.
func (o ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"equal_to": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"context":  basetypes.StringType{},
				"mask":     basetypes.StringType{},
				"position": basetypes.StringType{},
				"value":    basetypes.StringType{},
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
				"pattern": basetypes.StringType{},
				"qualifier": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"name":  basetypes.StringType{},
						"value": basetypes.StringType{},
					},
				}},
			},
		},
	}
}

// AttrType returns the attribute type for a list of ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator objects.
func (o ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperator) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo model.
func (o ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"context":  basetypes.StringType{},
		"mask":     basetypes.StringType{},
		"position": basetypes.StringType{},
		"value":    basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo objects.
func (o ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorEqualTo) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan model.
func (o ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"context": basetypes.StringType{},
		"qualifier": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":  basetypes.StringType{},
				"value": basetypes.StringType{},
			},
		}},
		"value": basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan objects.
func (o ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThan) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner model.
func (o ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":  basetypes.StringType{},
		"value": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner objects.
func (o ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorGreaterThanQualifierInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch model.
func (o ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"context": basetypes.StringType{},
		"pattern": basetypes.StringType{},
		"qualifier": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":  basetypes.StringType{},
				"value": basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch objects.
func (o ApplicationsSignatureInnerAndConditionInnerOrConditionInnerOperatorPatternMatch) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// ApplicationsResourceSchema defines the schema for Applications resource
var ApplicationsResourceSchema = schema.Schema{
	MarkdownDescription: "Application resource",
	Attributes: map[string]schema.Attribute{
		"able_to_transfer_file": schema.BoolAttribute{
			MarkdownDescription: "Able to transfer file",
			Optional:            true,
		},
		"alg_disable_capability": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(127),
			},
			MarkdownDescription: "Alg disable capability",
			Optional:            true,
		},
		"category": schema.StringAttribute{
			MarkdownDescription: "Category",
			Required:            true,
		},
		"consume_big_bandwidth": schema.BoolAttribute{
			MarkdownDescription: "Consume big bandwidth",
			Optional:            true,
		},
		"data_ident": schema.BoolAttribute{
			MarkdownDescription: "Data ident",
			Optional:            true,
		},
		"default": schema.SingleNestedAttribute{
			MarkdownDescription: "Default",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"ident_by_icmp6_type": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("ident_by_icmp_type"),
							path.MatchRelative().AtParent().AtName("ident_by_ip_protocol"),
							path.MatchRelative().AtParent().AtName("port"),
						),
					},
					MarkdownDescription: "Ident by icmp6 type\n\n> ℹ️ **Note:** You must specify exactly one of `ident_by_icmp6_type`, `ident_by_icmp_type`, `ident_by_ip_protocol`, and `port`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"code": schema.StringAttribute{
							MarkdownDescription: "Code",
							Optional:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Type",
							Required:            true,
						},
					},
				},
				"ident_by_icmp_type": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("ident_by_icmp6_type"),
							path.MatchRelative().AtParent().AtName("ident_by_ip_protocol"),
							path.MatchRelative().AtParent().AtName("port"),
						),
					},
					MarkdownDescription: "Ident by icmp type\n\n> ℹ️ **Note:** You must specify exactly one of `ident_by_icmp6_type`, `ident_by_icmp_type`, `ident_by_ip_protocol`, and `port`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"code": schema.StringAttribute{
							MarkdownDescription: "Code",
							Optional:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Type",
							Required:            true,
						},
					},
				},
				"ident_by_ip_protocol": schema.StringAttribute{
					Validators: []validator.String{
						stringvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("ident_by_icmp6_type"),
							path.MatchRelative().AtParent().AtName("ident_by_icmp_type"),
							path.MatchRelative().AtParent().AtName("port"),
						),
					},
					MarkdownDescription: "Ident by ip protocol\n\n> ℹ️ **Note:** You must specify exactly one of `ident_by_icmp6_type`, `ident_by_icmp_type`, `ident_by_ip_protocol`, and `port`.",
					Optional:            true,
				},
				"port": schema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Port\n\n> ℹ️ **Note:** You must specify exactly one of `ident_by_icmp6_type`, `ident_by_icmp_type`, `ident_by_ip_protocol`, and `port`.",
					Validators: []validator.List{
						listvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("ident_by_icmp6_type"),
							path.MatchRelative().AtParent().AtName("ident_by_icmp_type"),
							path.MatchRelative().AtParent().AtName("ident_by_ip_protocol"),
						),
						listvalidator.ValueStringsAre(stringvalidator.LengthAtMost(63)),
					},
					Optional: true,
				},
			},
		},
		"description": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(8192),
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
		"evasive_behavior": schema.BoolAttribute{
			MarkdownDescription: "Evasive behavior",
			Optional:            true,
		},
		"file_type_ident": schema.BoolAttribute{
			MarkdownDescription: "File type ident",
			Optional:            true,
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
		"has_known_vulnerability": schema.BoolAttribute{
			MarkdownDescription: "Has known vulnerability",
			Optional:            true,
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the application",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(31),
			},
			MarkdownDescription: "The name of the application",
			Required:            true,
		},
		"no_appid_caching": schema.BoolAttribute{
			MarkdownDescription: "No appid caching",
			Optional:            true,
		},
		"parent_app": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(127),
			},
			MarkdownDescription: "Parent app",
			Optional:            true,
		},
		"pervasive_use": schema.BoolAttribute{
			MarkdownDescription: "Pervasive use",
			Optional:            true,
		},
		"prone_to_misuse": schema.BoolAttribute{
			MarkdownDescription: "Prone to misuse",
			Optional:            true,
		},
		"risk": schema.StringAttribute{
			MarkdownDescription: "Risk",
			Required:            true,
		},
		"signature": schema.ListNestedAttribute{
			MarkdownDescription: "Signature",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"and_condition": schema.ListNestedAttribute{
						MarkdownDescription: "And condition",
						Optional:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.LengthAtMost(31),
									},
									MarkdownDescription: "Alphanumeric string [ 0-9a-zA-Z._-]",
									Required:            true,
								},
								"or_condition": schema.ListNestedAttribute{
									MarkdownDescription: "Or condition",
									Optional:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Validators: []validator.String{
													stringvalidator.LengthAtMost(31),
												},
												MarkdownDescription: "Alphanumeric string [ 0-9a-zA-Z._-]",
												Required:            true,
											},
											"operator": schema.SingleNestedAttribute{
												MarkdownDescription: "Operator",
												Required:            true,
												Attributes: map[string]schema.Attribute{
													"equal_to": schema.SingleNestedAttribute{
														Validators: []validator.Object{
															objectvalidator.ExactlyOneOf(
																path.MatchRelative().AtParent().AtName("greater_than"),
																path.MatchRelative().AtParent().AtName("less_than"),
																path.MatchRelative().AtParent().AtName("pattern_match"),
															),
														},
														MarkdownDescription: "Equal to\n\n> ℹ️ **Note:** You must specify exactly one of `equal_to`, `greater_than`, `less_than`, and `pattern_match`.",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																MarkdownDescription: "Context",
																Required:            true,
															},
															"mask": schema.StringAttribute{
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(10),
																	stringvalidator.RegexMatches(regexp.MustCompile("^[0][xX][0-9A-Fa-f]{8}$"), "pattern must match "+"^[0][xX][0-9A-Fa-f]{8}$"),
																},
																MarkdownDescription: "4-byte hex value",
																Optional:            true,
															},
															"position": schema.StringAttribute{
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(127),
																},
																MarkdownDescription: "Position",
																Optional:            true,
															},
															"value": schema.StringAttribute{
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(10),
																},
																MarkdownDescription: "Value",
																Required:            true,
															},
														},
													},
													"greater_than": schema.SingleNestedAttribute{
														Validators: []validator.Object{
															objectvalidator.ExactlyOneOf(
																path.MatchRelative().AtParent().AtName("equal_to"),
																path.MatchRelative().AtParent().AtName("less_than"),
																path.MatchRelative().AtParent().AtName("pattern_match"),
															),
														},
														MarkdownDescription: "Greater than\n\n> ℹ️ **Note:** You must specify exactly one of `equal_to`, `greater_than`, `less_than`, and `pattern_match`.",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(127),
																},
																MarkdownDescription: "Context",
																Required:            true,
															},
															"qualifier": schema.ListNestedAttribute{
																MarkdownDescription: "Qualifier",
																Optional:            true,
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Validators: []validator.String{
																				stringvalidator.LengthAtMost(31),
																			},
																			MarkdownDescription: "Alphanumeric string [ 0-9a-zA-Z._-]",
																			Required:            true,
																		},
																		"value": schema.StringAttribute{
																			MarkdownDescription: "Value",
																			Required:            true,
																		},
																	},
																},
															},
															"value": schema.Int64Attribute{
																Validators: []validator.Int64{
																	int64validator.Between(0, 4294967295),
																},
																MarkdownDescription: "Value",
																Required:            true,
															},
														},
													},
													"less_than": schema.SingleNestedAttribute{
														Validators: []validator.Object{
															objectvalidator.ExactlyOneOf(
																path.MatchRelative().AtParent().AtName("equal_to"),
																path.MatchRelative().AtParent().AtName("greater_than"),
																path.MatchRelative().AtParent().AtName("pattern_match"),
															),
														},
														MarkdownDescription: "Less than\n\n> ℹ️ **Note:** You must specify exactly one of `equal_to`, `greater_than`, `less_than`, and `pattern_match`.",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(127),
																},
																MarkdownDescription: "Context",
																Required:            true,
															},
															"qualifier": schema.ListNestedAttribute{
																MarkdownDescription: "Qualifier",
																Optional:            true,
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Validators: []validator.String{
																				stringvalidator.LengthAtMost(31),
																			},
																			MarkdownDescription: "Alphanumeric string [ 0-9a-zA-Z._-]",
																			Required:            true,
																		},
																		"value": schema.StringAttribute{
																			MarkdownDescription: "Value",
																			Required:            true,
																		},
																	},
																},
															},
															"value": schema.Int64Attribute{
																Validators: []validator.Int64{
																	int64validator.Between(0, 4294967295),
																},
																MarkdownDescription: "Value",
																Required:            true,
															},
														},
													},
													"pattern_match": schema.SingleNestedAttribute{
														Validators: []validator.Object{
															objectvalidator.ExactlyOneOf(
																path.MatchRelative().AtParent().AtName("equal_to"),
																path.MatchRelative().AtParent().AtName("greater_than"),
																path.MatchRelative().AtParent().AtName("less_than"),
															),
														},
														MarkdownDescription: "Pattern match\n\n> ℹ️ **Note:** You must specify exactly one of `equal_to`, `greater_than`, `less_than`, and `pattern_match`.",
														Optional:            true,
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(127),
																},
																MarkdownDescription: "Context",
																Required:            true,
															},
															"pattern": schema.StringAttribute{
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(127),
																},
																MarkdownDescription: "Pattern",
																Required:            true,
															},
															"qualifier": schema.ListNestedAttribute{
																MarkdownDescription: "Qualifier",
																Optional:            true,
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Validators: []validator.String{
																				stringvalidator.LengthAtMost(31),
																			},
																			MarkdownDescription: "Alphanumeric string [ 0-9a-zA-Z._-]",
																			Required:            true,
																		},
																		"value": schema.StringAttribute{
																			MarkdownDescription: "Value",
																			Required:            true,
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
						Validators: []validator.String{
							stringvalidator.LengthAtMost(31),
						},
						MarkdownDescription: "Alphanumeric string [ 0-9a-zA-Z._-]",
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
						Computed:            true,
						Default:             stringdefault.StaticString("protocol-data-unit"),
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
		"subcategory": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(63),
			},
			MarkdownDescription: "Subcategory",
			Optional:            true,
		},
		"tcp_half_closed_timeout": schema.Int64Attribute{
			Validators: []validator.Int64{
				int64validator.Between(1, 604800),
			},
			MarkdownDescription: "timeout for half-close session in seconds",
			Optional:            true,
		},
		"tcp_time_wait_timeout": schema.Int64Attribute{
			Validators: []validator.Int64{
				int64validator.Between(1, 600),
			},
			MarkdownDescription: "timeout for session in time_wait state in seconds",
			Optional:            true,
		},
		"tcp_timeout": schema.Int64Attribute{
			Validators: []validator.Int64{
				int64validator.Between(0, 604800),
			},
			MarkdownDescription: "timeout in seconds",
			Optional:            true,
		},
		"technology": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(63),
			},
			MarkdownDescription: "Technology",
			Optional:            true,
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"timeout": schema.Int64Attribute{
			Validators: []validator.Int64{
				int64validator.Between(0, 604800),
			},
			MarkdownDescription: "timeout in seconds",
			Optional:            true,
		},
		"tunnel_applications": schema.BoolAttribute{
			MarkdownDescription: "Tunnel applications",
			Optional:            true,
		},
		"tunnel_other_application": schema.BoolAttribute{
			MarkdownDescription: "Tunnel other application",
			Optional:            true,
		},
		"udp_timeout": schema.Int64Attribute{
			Validators: []validator.Int64{
				int64validator.Between(0, 604800),
			},
			MarkdownDescription: "timeout in seconds",
			Optional:            true,
		},
		"used_by_malware": schema.BoolAttribute{
			MarkdownDescription: "Used by malware",
			Optional:            true,
		},
		"virus_ident": schema.BoolAttribute{
			MarkdownDescription: "Virus ident",
			Optional:            true,
		},
	},
}

// ApplicationsDataSourceSchema defines the schema for Applications data source
var ApplicationsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Application data source",
	Attributes: map[string]dsschema.Attribute{
		"able_to_transfer_file": dsschema.BoolAttribute{
			MarkdownDescription: "Able to transfer file",
			Computed:            true,
		},
		"alg_disable_capability": dsschema.StringAttribute{
			MarkdownDescription: "Alg disable capability",
			Computed:            true,
		},
		"category": dsschema.StringAttribute{
			MarkdownDescription: "Category",
			Computed:            true,
		},
		"consume_big_bandwidth": dsschema.BoolAttribute{
			MarkdownDescription: "Consume big bandwidth",
			Computed:            true,
		},
		"data_ident": dsschema.BoolAttribute{
			MarkdownDescription: "Data ident",
			Computed:            true,
		},
		"default": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Default",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"ident_by_icmp6_type": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Ident by icmp6 type\n\n> ℹ️ **Note:** You must specify exactly one of `ident_by_icmp6_type`, `ident_by_icmp_type`, `ident_by_ip_protocol`, and `port`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"code": dsschema.StringAttribute{
							MarkdownDescription: "Code",
							Computed:            true,
						},
						"type": dsschema.StringAttribute{
							MarkdownDescription: "Type",
							Computed:            true,
						},
					},
				},
				"ident_by_icmp_type": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Ident by icmp type\n\n> ℹ️ **Note:** You must specify exactly one of `ident_by_icmp6_type`, `ident_by_icmp_type`, `ident_by_ip_protocol`, and `port`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"code": dsschema.StringAttribute{
							MarkdownDescription: "Code",
							Computed:            true,
						},
						"type": dsschema.StringAttribute{
							MarkdownDescription: "Type",
							Computed:            true,
						},
					},
				},
				"ident_by_ip_protocol": dsschema.StringAttribute{
					MarkdownDescription: "Ident by ip protocol\n\n> ℹ️ **Note:** You must specify exactly one of `ident_by_icmp6_type`, `ident_by_icmp_type`, `ident_by_ip_protocol`, and `port`.",
					Computed:            true,
				},
				"port": dsschema.ListAttribute{
					ElementType:         types.StringType,
					MarkdownDescription: "Port\n\n> ℹ️ **Note:** You must specify exactly one of `ident_by_icmp6_type`, `ident_by_icmp_type`, `ident_by_ip_protocol`, and `port`.",
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
		"evasive_behavior": dsschema.BoolAttribute{
			MarkdownDescription: "Evasive behavior",
			Computed:            true,
		},
		"file_type_ident": dsschema.BoolAttribute{
			MarkdownDescription: "File type ident",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"has_known_vulnerability": dsschema.BoolAttribute{
			MarkdownDescription: "Has known vulnerability",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the application",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the application",
			Optional:            true,
			Computed:            true,
		},
		"no_appid_caching": dsschema.BoolAttribute{
			MarkdownDescription: "No appid caching",
			Computed:            true,
		},
		"parent_app": dsschema.StringAttribute{
			MarkdownDescription: "Parent app",
			Computed:            true,
		},
		"pervasive_use": dsschema.BoolAttribute{
			MarkdownDescription: "Pervasive use",
			Computed:            true,
		},
		"prone_to_misuse": dsschema.BoolAttribute{
			MarkdownDescription: "Prone to misuse",
			Computed:            true,
		},
		"risk": dsschema.StringAttribute{
			MarkdownDescription: "Risk",
			Computed:            true,
		},
		"signature": dsschema.ListNestedAttribute{
			MarkdownDescription: "Signature",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"and_condition": dsschema.ListNestedAttribute{
						MarkdownDescription: "And condition",
						Computed:            true,
						NestedObject: dsschema.NestedAttributeObject{
							Attributes: map[string]dsschema.Attribute{
								"name": dsschema.StringAttribute{
									MarkdownDescription: "Alphanumeric string [ 0-9a-zA-Z._-]",
									Computed:            true,
								},
								"or_condition": dsschema.ListNestedAttribute{
									MarkdownDescription: "Or condition",
									Computed:            true,
									NestedObject: dsschema.NestedAttributeObject{
										Attributes: map[string]dsschema.Attribute{
											"name": dsschema.StringAttribute{
												MarkdownDescription: "Alphanumeric string [ 0-9a-zA-Z._-]",
												Computed:            true,
											},
											"operator": dsschema.SingleNestedAttribute{
												MarkdownDescription: "Operator",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"equal_to": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Equal to\n\n> ℹ️ **Note:** You must specify exactly one of `equal_to`, `greater_than`, `less_than`, and `pattern_match`.",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"context": dsschema.StringAttribute{
																MarkdownDescription: "Context",
																Computed:            true,
															},
															"mask": dsschema.StringAttribute{
																MarkdownDescription: "4-byte hex value",
																Computed:            true,
															},
															"position": dsschema.StringAttribute{
																MarkdownDescription: "Position",
																Computed:            true,
															},
															"value": dsschema.StringAttribute{
																MarkdownDescription: "Value",
																Computed:            true,
															},
														},
													},
													"greater_than": dsschema.SingleNestedAttribute{
														MarkdownDescription: "Greater than\n\n> ℹ️ **Note:** You must specify exactly one of `equal_to`, `greater_than`, `less_than`, and `pattern_match`.",
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
																			MarkdownDescription: "Alphanumeric string [ 0-9a-zA-Z._-]",
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
														MarkdownDescription: "Less than\n\n> ℹ️ **Note:** You must specify exactly one of `equal_to`, `greater_than`, `less_than`, and `pattern_match`.",
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
																			MarkdownDescription: "Alphanumeric string [ 0-9a-zA-Z._-]",
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
														MarkdownDescription: "Pattern match\n\n> ℹ️ **Note:** You must specify exactly one of `equal_to`, `greater_than`, `less_than`, and `pattern_match`.",
														Computed:            true,
														Attributes: map[string]dsschema.Attribute{
															"context": dsschema.StringAttribute{
																MarkdownDescription: "Context",
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
																			MarkdownDescription: "Alphanumeric string [ 0-9a-zA-Z._-]",
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
						MarkdownDescription: "Alphanumeric string [ 0-9a-zA-Z._-]",
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
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Optional:            true,
			Computed:            true,
		},
		"subcategory": dsschema.StringAttribute{
			MarkdownDescription: "Subcategory",
			Computed:            true,
		},
		"tcp_half_closed_timeout": dsschema.Int64Attribute{
			MarkdownDescription: "timeout for half-close session in seconds",
			Computed:            true,
		},
		"tcp_time_wait_timeout": dsschema.Int64Attribute{
			MarkdownDescription: "timeout for session in time_wait state in seconds",
			Computed:            true,
		},
		"tcp_timeout": dsschema.Int64Attribute{
			MarkdownDescription: "timeout in seconds",
			Computed:            true,
		},
		"technology": dsschema.StringAttribute{
			MarkdownDescription: "Technology",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"timeout": dsschema.Int64Attribute{
			MarkdownDescription: "timeout in seconds",
			Computed:            true,
		},
		"tunnel_applications": dsschema.BoolAttribute{
			MarkdownDescription: "Tunnel applications",
			Computed:            true,
		},
		"tunnel_other_application": dsschema.BoolAttribute{
			MarkdownDescription: "Tunnel other application",
			Computed:            true,
		},
		"udp_timeout": dsschema.Int64Attribute{
			MarkdownDescription: "timeout in seconds",
			Computed:            true,
		},
		"used_by_malware": dsschema.BoolAttribute{
			MarkdownDescription: "Used by malware",
			Computed:            true,
		},
		"virus_ident": dsschema.BoolAttribute{
			MarkdownDescription: "Virus ident",
			Computed:            true,
		},
	},
}

// ApplicationsListModel represents the data model for a list data source.
type ApplicationsListModel struct {
	Tfid    types.String   `tfsdk:"tfid"`
	Data    []Applications `tfsdk:"data"`
	Limit   types.Int64    `tfsdk:"limit"`
	Offset  types.Int64    `tfsdk:"offset"`
	Name    types.String   `tfsdk:"name"`
	Total   types.Int64    `tfsdk:"total"`
	Folder  types.String   `tfsdk:"folder"`
	Snippet types.String   `tfsdk:"snippet"`
	Device  types.String   `tfsdk:"device"`
}

// ApplicationsListDataSourceSchema defines the schema for a list data source.
var ApplicationsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: ApplicationsDataSourceSchema.Attributes,
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
