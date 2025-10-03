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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: network_services
// This file contains models for the network_services SDK package

// QosProfiles represents the Terraform model for QosProfiles
type QosProfiles struct {
	Tfid               types.String          `tfsdk:"tfid"`
	AggregateBandwidth basetypes.ObjectValue `tfsdk:"aggregate_bandwidth"`
	ClassBandwidthType basetypes.ObjectValue `tfsdk:"class_bandwidth_type"`
	Device             basetypes.StringValue `tfsdk:"device"`
	Folder             basetypes.StringValue `tfsdk:"folder"`
	Id                 basetypes.StringValue `tfsdk:"id"`
	Name               basetypes.StringValue `tfsdk:"name"`
	Snippet            basetypes.StringValue `tfsdk:"snippet"`
}

// QosProfilesAggregateBandwidth represents a nested structure within the QosProfiles model
type QosProfilesAggregateBandwidth struct {
	EgressGuaranteed basetypes.Int64Value `tfsdk:"egress_guaranteed"`
	EgressMax        basetypes.Int64Value `tfsdk:"egress_max"`
}

// QosProfilesClassBandwidthType represents a nested structure within the QosProfiles model
type QosProfilesClassBandwidthType struct {
	Mbps       basetypes.ObjectValue `tfsdk:"mbps"`
	Percentage basetypes.ObjectValue `tfsdk:"percentage"`
}

// QosProfilesClassBandwidthTypeMbps represents a nested structure within the QosProfiles model
type QosProfilesClassBandwidthTypeMbps struct {
	Class basetypes.ListValue `tfsdk:"class"`
}

// QosProfilesClassBandwidthTypeMbpsClassInner represents a nested structure within the QosProfiles model
type QosProfilesClassBandwidthTypeMbpsClassInner struct {
	ClassBandwidth basetypes.ObjectValue `tfsdk:"class_bandwidth"`
	Name           basetypes.StringValue `tfsdk:"name"`
	Priority       basetypes.StringValue `tfsdk:"priority"`
}

// QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth represents a nested structure within the QosProfiles model
type QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth struct {
	EgressGuaranteed basetypes.Int64Value `tfsdk:"egress_guaranteed"`
	EgressMax        basetypes.Int64Value `tfsdk:"egress_max"`
}

// QosProfilesClassBandwidthTypePercentage represents a nested structure within the QosProfiles model
type QosProfilesClassBandwidthTypePercentage struct {
	Class basetypes.ListValue `tfsdk:"class"`
}

// QosProfilesClassBandwidthTypePercentageClassInner represents a nested structure within the QosProfiles model
type QosProfilesClassBandwidthTypePercentageClassInner struct {
	ClassBandwidth basetypes.ObjectValue `tfsdk:"class_bandwidth"`
	Name           basetypes.StringValue `tfsdk:"name"`
	Priority       basetypes.StringValue `tfsdk:"priority"`
}

// QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth represents a nested structure within the QosProfiles model
type QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth struct {
	EgressGuaranteed basetypes.Int64Value `tfsdk:"egress_guaranteed"`
	EgressMax        basetypes.Int64Value `tfsdk:"egress_max"`
}

// AttrTypes defines the attribute types for the QosProfiles model.
func (o QosProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid": basetypes.StringType{},
		"aggregate_bandwidth": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"egress_guaranteed": basetypes.Int64Type{},
				"egress_max":        basetypes.Int64Type{},
			},
		},
		"class_bandwidth_type": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"mbps": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"class": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"class_bandwidth": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"egress_guaranteed": basetypes.Int64Type{},
										"egress_max":        basetypes.Int64Type{},
									},
								},
								"name":     basetypes.StringType{},
								"priority": basetypes.StringType{},
							},
						}},
					},
				},
				"percentage": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"class": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"class_bandwidth": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"egress_guaranteed": basetypes.Int64Type{},
										"egress_max":        basetypes.Int64Type{},
									},
								},
								"name":     basetypes.StringType{},
								"priority": basetypes.StringType{},
							},
						}},
					},
				},
			},
		},
		"device":  basetypes.StringType{},
		"folder":  basetypes.StringType{},
		"id":      basetypes.StringType{},
		"name":    basetypes.StringType{},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of QosProfiles objects.
func (o QosProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the QosProfilesAggregateBandwidth model.
func (o QosProfilesAggregateBandwidth) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"egress_guaranteed": basetypes.Int64Type{},
		"egress_max":        basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of QosProfilesAggregateBandwidth objects.
func (o QosProfilesAggregateBandwidth) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the QosProfilesClassBandwidthType model.
func (o QosProfilesClassBandwidthType) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"mbps": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"class": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"class_bandwidth": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"egress_guaranteed": basetypes.Int64Type{},
								"egress_max":        basetypes.Int64Type{},
							},
						},
						"name":     basetypes.StringType{},
						"priority": basetypes.StringType{},
					},
				}},
			},
		},
		"percentage": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"class": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"class_bandwidth": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"egress_guaranteed": basetypes.Int64Type{},
								"egress_max":        basetypes.Int64Type{},
							},
						},
						"name":     basetypes.StringType{},
						"priority": basetypes.StringType{},
					},
				}},
			},
		},
	}
}

// AttrType returns the attribute type for a list of QosProfilesClassBandwidthType objects.
func (o QosProfilesClassBandwidthType) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the QosProfilesClassBandwidthTypeMbps model.
func (o QosProfilesClassBandwidthTypeMbps) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"class": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"class_bandwidth": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"egress_guaranteed": basetypes.Int64Type{},
						"egress_max":        basetypes.Int64Type{},
					},
				},
				"name":     basetypes.StringType{},
				"priority": basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of QosProfilesClassBandwidthTypeMbps objects.
func (o QosProfilesClassBandwidthTypeMbps) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the QosProfilesClassBandwidthTypeMbpsClassInner model.
func (o QosProfilesClassBandwidthTypeMbpsClassInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"class_bandwidth": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"egress_guaranteed": basetypes.Int64Type{},
				"egress_max":        basetypes.Int64Type{},
			},
		},
		"name":     basetypes.StringType{},
		"priority": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of QosProfilesClassBandwidthTypeMbpsClassInner objects.
func (o QosProfilesClassBandwidthTypeMbpsClassInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth model.
func (o QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"egress_guaranteed": basetypes.Int64Type{},
		"egress_max":        basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth objects.
func (o QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the QosProfilesClassBandwidthTypePercentage model.
func (o QosProfilesClassBandwidthTypePercentage) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"class": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"class_bandwidth": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"egress_guaranteed": basetypes.Int64Type{},
						"egress_max":        basetypes.Int64Type{},
					},
				},
				"name":     basetypes.StringType{},
				"priority": basetypes.StringType{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of QosProfilesClassBandwidthTypePercentage objects.
func (o QosProfilesClassBandwidthTypePercentage) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the QosProfilesClassBandwidthTypePercentageClassInner model.
func (o QosProfilesClassBandwidthTypePercentageClassInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"class_bandwidth": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"egress_guaranteed": basetypes.Int64Type{},
				"egress_max":        basetypes.Int64Type{},
			},
		},
		"name":     basetypes.StringType{},
		"priority": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of QosProfilesClassBandwidthTypePercentageClassInner objects.
func (o QosProfilesClassBandwidthTypePercentageClassInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth model.
func (o QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"egress_guaranteed": basetypes.Int64Type{},
		"egress_max":        basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth objects.
func (o QosProfilesClassBandwidthTypePercentageClassInnerClassBandwidth) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// QosProfilesResourceSchema defines the schema for QosProfiles resource
var QosProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "QosProfile resource",
	Attributes: map[string]schema.Attribute{
		"aggregate_bandwidth": schema.SingleNestedAttribute{
			MarkdownDescription: "Aggregate bandwidth",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"egress_guaranteed": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(0, 16000),
					},
					MarkdownDescription: "guaranteed sending bandwidth in mbps",
					Optional:            true,
				},
				"egress_max": schema.Int64Attribute{
					Validators: []validator.Int64{
						int64validator.Between(0, 60000),
					},
					MarkdownDescription: "max sending bandwidth in mbps",
					Optional:            true,
				},
			},
		},
		"class_bandwidth_type": schema.SingleNestedAttribute{
			MarkdownDescription: "Class bandwidth type",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"mbps": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("percentage"),
						),
					},
					MarkdownDescription: "Mbps",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"class": schema.ListNestedAttribute{
							MarkdownDescription: "QoS setting for traffic classes",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"class_bandwidth": schema.SingleNestedAttribute{
										MarkdownDescription: "Class bandwidth",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"egress_guaranteed": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(0, 60000),
												},
												MarkdownDescription: "guaranteed sending bandwidth in mbps",
												Optional:            true,
											},
											"egress_max": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(0, 60000),
												},
												MarkdownDescription: "max sending bandwidth in mbps",
												Optional:            true,
											},
										},
									},
									"name": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.LengthAtMost(31),
										},
										MarkdownDescription: "Traffic class",
										Optional:            true,
									},
									"priority": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("real-time", "high", "medium", "low"),
										},
										MarkdownDescription: "traffic class priority",
										Optional:            true,
										Computed:            true,
										Default:             stringdefault.StaticString("medium"),
									},
								},
							},
						},
					},
				},
				"percentage": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("mbps"),
						),
					},
					MarkdownDescription: "Percentage",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"class": schema.ListNestedAttribute{
							MarkdownDescription: "QoS setting for traffic classes",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"class_bandwidth": schema.SingleNestedAttribute{
										MarkdownDescription: "Class bandwidth",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"egress_guaranteed": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(0, 100),
												},
												MarkdownDescription: "guaranteed sending bandwidth in percentage",
												Optional:            true,
											},
											"egress_max": schema.Int64Attribute{
												Validators: []validator.Int64{
													int64validator.Between(0, 100),
												},
												MarkdownDescription: "max sending bandwidth in percentage",
												Optional:            true,
											},
										},
									},
									"name": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.LengthAtMost(31),
										},
										MarkdownDescription: "Traffic class",
										Optional:            true,
									},
									"priority": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("real-time", "high", "medium", "low"),
										},
										MarkdownDescription: "traffic class priority",
										Optional:            true,
										Computed:            true,
										Default:             stringdefault.StaticString("medium"),
									},
								},
							},
						},
					},
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
		"name": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.LengthAtMost(31),
			},
			MarkdownDescription: "Alphanumeric string begin with letter: [0-9a-zA-Z._-]",
			Required:            true,
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
	},
}

// QosProfilesDataSourceSchema defines the schema for QosProfiles data source
var QosProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "QosProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"aggregate_bandwidth": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Aggregate bandwidth",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"egress_guaranteed": dsschema.Int64Attribute{
					MarkdownDescription: "guaranteed sending bandwidth in mbps",
					Computed:            true,
				},
				"egress_max": dsschema.Int64Attribute{
					MarkdownDescription: "max sending bandwidth in mbps",
					Computed:            true,
				},
			},
		},
		"class_bandwidth_type": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Class bandwidth type",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"mbps": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Mbps",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"class": dsschema.ListNestedAttribute{
							MarkdownDescription: "QoS setting for traffic classes",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"class_bandwidth": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Class bandwidth",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"egress_guaranteed": dsschema.Int64Attribute{
												MarkdownDescription: "guaranteed sending bandwidth in mbps",
												Computed:            true,
											},
											"egress_max": dsschema.Int64Attribute{
												MarkdownDescription: "max sending bandwidth in mbps",
												Computed:            true,
											},
										},
									},
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Traffic class",
										Computed:            true,
									},
									"priority": dsschema.StringAttribute{
										MarkdownDescription: "traffic class priority",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"percentage": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Percentage",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"class": dsschema.ListNestedAttribute{
							MarkdownDescription: "QoS setting for traffic classes",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"class_bandwidth": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Class bandwidth",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"egress_guaranteed": dsschema.Int64Attribute{
												MarkdownDescription: "guaranteed sending bandwidth in percentage",
												Computed:            true,
											},
											"egress_max": dsschema.Int64Attribute{
												MarkdownDescription: "max sending bandwidth in percentage",
												Computed:            true,
											},
										},
									},
									"name": dsschema.StringAttribute{
										MarkdownDescription: "Traffic class",
										Computed:            true,
									},
									"priority": dsschema.StringAttribute{
										MarkdownDescription: "traffic class priority",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
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
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Alphanumeric string begin with letter: [0-9a-zA-Z._-]",
			Optional:            true,
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// QosProfilesListModel represents the data model for a list data source.
type QosProfilesListModel struct {
	Tfid    types.String  `tfsdk:"tfid"`
	Data    []QosProfiles `tfsdk:"data"`
	Limit   types.Int64   `tfsdk:"limit"`
	Offset  types.Int64   `tfsdk:"offset"`
	Name    types.String  `tfsdk:"name"`
	Total   types.Int64   `tfsdk:"total"`
	Folder  types.String  `tfsdk:"folder"`
	Snippet types.String  `tfsdk:"snippet"`
	Device  types.String  `tfsdk:"device"`
}

// QosProfilesListDataSourceSchema defines the schema for a list data source.
var QosProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: QosProfilesDataSourceSchema.Attributes,
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
