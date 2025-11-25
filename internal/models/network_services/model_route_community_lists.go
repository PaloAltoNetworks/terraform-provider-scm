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

// RouteCommunityLists represents the Terraform model for RouteCommunityLists
type RouteCommunityLists struct {
	Tfid        types.String          `tfsdk:"tfid"`
	Description basetypes.StringValue `tfsdk:"description"`
	Device      basetypes.StringValue `tfsdk:"device"`
	Folder      basetypes.StringValue `tfsdk:"folder"`
	Id          basetypes.StringValue `tfsdk:"id"`
	Name        basetypes.StringValue `tfsdk:"name"`
	Snippet     basetypes.StringValue `tfsdk:"snippet"`
	Type        basetypes.ObjectValue `tfsdk:"type"`
}

// RouteCommunityListsType represents a nested structure within the RouteCommunityLists model
type RouteCommunityListsType struct {
	Extended basetypes.ObjectValue `tfsdk:"extended"`
	Large    basetypes.ObjectValue `tfsdk:"large"`
	Regular  basetypes.ObjectValue `tfsdk:"regular"`
}

// RouteCommunityListsTypeExtended represents a nested structure within the RouteCommunityLists model
type RouteCommunityListsTypeExtended struct {
	ExtendedEntry basetypes.ListValue `tfsdk:"extended_entry"`
}

// RouteCommunityListsTypeExtendedExtendedEntryInner represents a nested structure within the RouteCommunityLists model
type RouteCommunityListsTypeExtendedExtendedEntryInner struct {
	Action  basetypes.StringValue `tfsdk:"action"`
	LcRegex basetypes.ListValue   `tfsdk:"lc_regex"`
	Name    basetypes.Int64Value  `tfsdk:"name"`
}

// RouteCommunityListsTypeLarge represents a nested structure within the RouteCommunityLists model
type RouteCommunityListsTypeLarge struct {
	LargeEntry basetypes.ListValue `tfsdk:"large_entry"`
}

// RouteCommunityListsTypeLargeLargeEntryInner represents a nested structure within the RouteCommunityLists model
type RouteCommunityListsTypeLargeLargeEntryInner struct {
	Action  basetypes.StringValue `tfsdk:"action"`
	LcRegex basetypes.ListValue   `tfsdk:"lc_regex"`
	Name    basetypes.Int64Value  `tfsdk:"name"`
}

// RouteCommunityListsTypeRegular represents a nested structure within the RouteCommunityLists model
type RouteCommunityListsTypeRegular struct {
	RegularEntry basetypes.ListValue `tfsdk:"regular_entry"`
}

// RouteCommunityListsTypeRegularRegularEntryInner represents a nested structure within the RouteCommunityLists model
type RouteCommunityListsTypeRegularRegularEntryInner struct {
	Action    basetypes.StringValue `tfsdk:"action"`
	Community basetypes.ListValue   `tfsdk:"community"`
	Name      basetypes.Int64Value  `tfsdk:"name"`
}

// AttrTypes defines the attribute types for the RouteCommunityLists model.
func (o RouteCommunityLists) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":        basetypes.StringType{},
		"description": basetypes.StringType{},
		"device":      basetypes.StringType{},
		"folder":      basetypes.StringType{},
		"id":          basetypes.StringType{},
		"name":        basetypes.StringType{},
		"snippet":     basetypes.StringType{},
		"type": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"extended": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"extended_entry": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action":   basetypes.StringType{},
								"lc_regex": basetypes.ListType{ElemType: basetypes.StringType{}},
								"name":     basetypes.Int64Type{},
							},
						}},
					},
				},
				"large": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"large_entry": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action":   basetypes.StringType{},
								"lc_regex": basetypes.ListType{ElemType: basetypes.StringType{}},
								"name":     basetypes.Int64Type{},
							},
						}},
					},
				},
				"regular": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"regular_entry": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action":    basetypes.StringType{},
								"community": basetypes.ListType{ElemType: basetypes.StringType{}},
								"name":      basetypes.Int64Type{},
							},
						}},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of RouteCommunityLists objects.
func (o RouteCommunityLists) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RouteCommunityListsType model.
func (o RouteCommunityListsType) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"extended": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"extended_entry": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action":   basetypes.StringType{},
						"lc_regex": basetypes.ListType{ElemType: basetypes.StringType{}},
						"name":     basetypes.Int64Type{},
					},
				}},
			},
		},
		"large": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"large_entry": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action":   basetypes.StringType{},
						"lc_regex": basetypes.ListType{ElemType: basetypes.StringType{}},
						"name":     basetypes.Int64Type{},
					},
				}},
			},
		},
		"regular": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"regular_entry": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action":    basetypes.StringType{},
						"community": basetypes.ListType{ElemType: basetypes.StringType{}},
						"name":      basetypes.Int64Type{},
					},
				}},
			},
		},
	}
}

// AttrType returns the attribute type for a list of RouteCommunityListsType objects.
func (o RouteCommunityListsType) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RouteCommunityListsTypeExtended model.
func (o RouteCommunityListsTypeExtended) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"extended_entry": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action":   basetypes.StringType{},
				"lc_regex": basetypes.ListType{ElemType: basetypes.StringType{}},
				"name":     basetypes.Int64Type{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of RouteCommunityListsTypeExtended objects.
func (o RouteCommunityListsTypeExtended) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RouteCommunityListsTypeExtendedExtendedEntryInner model.
func (o RouteCommunityListsTypeExtendedExtendedEntryInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action":   basetypes.StringType{},
		"lc_regex": basetypes.ListType{ElemType: basetypes.StringType{}},
		"name":     basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of RouteCommunityListsTypeExtendedExtendedEntryInner objects.
func (o RouteCommunityListsTypeExtendedExtendedEntryInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RouteCommunityListsTypeLarge model.
func (o RouteCommunityListsTypeLarge) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"large_entry": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action":   basetypes.StringType{},
				"lc_regex": basetypes.ListType{ElemType: basetypes.StringType{}},
				"name":     basetypes.Int64Type{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of RouteCommunityListsTypeLarge objects.
func (o RouteCommunityListsTypeLarge) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RouteCommunityListsTypeLargeLargeEntryInner model.
func (o RouteCommunityListsTypeLargeLargeEntryInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action":   basetypes.StringType{},
		"lc_regex": basetypes.ListType{ElemType: basetypes.StringType{}},
		"name":     basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of RouteCommunityListsTypeLargeLargeEntryInner objects.
func (o RouteCommunityListsTypeLargeLargeEntryInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RouteCommunityListsTypeRegular model.
func (o RouteCommunityListsTypeRegular) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"regular_entry": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action":    basetypes.StringType{},
				"community": basetypes.ListType{ElemType: basetypes.StringType{}},
				"name":      basetypes.Int64Type{},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of RouteCommunityListsTypeRegular objects.
func (o RouteCommunityListsTypeRegular) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RouteCommunityListsTypeRegularRegularEntryInner model.
func (o RouteCommunityListsTypeRegularRegularEntryInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action":    basetypes.StringType{},
		"community": basetypes.ListType{ElemType: basetypes.StringType{}},
		"name":      basetypes.Int64Type{},
	}
}

// AttrType returns the attribute type for a list of RouteCommunityListsTypeRegularRegularEntryInner objects.
func (o RouteCommunityListsTypeRegularRegularEntryInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// RouteCommunityListsResourceSchema defines the schema for RouteCommunityLists resource
var RouteCommunityListsResourceSchema = schema.Schema{
	MarkdownDescription: "RouteCommunityList resource",
	Attributes: map[string]schema.Attribute{
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
		"name": schema.StringAttribute{
			MarkdownDescription: "Route community list name",
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
		"type": schema.SingleNestedAttribute{
			MarkdownDescription: "Type",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"extended": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("large"),
							path.MatchRelative().AtParent().AtName("regular"),
						),
					},
					MarkdownDescription: "Extended\n\n> ℹ️ **Note:** You must specify exactly one of `extended`, `large`, and `regular`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"extended_entry": schema.ListNestedAttribute{
							MarkdownDescription: "Extended community lists",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"action": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("deny", "permit"),
										},
										MarkdownDescription: "Action",
										Optional:            true,
									},
									"lc_regex": schema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "Extended community regular expression",
										Validators: []validator.List{
											listvalidator.SizeAtMost(8),
										},
										Optional: true,
									},
									"name": schema.Int64Attribute{
										Validators: []validator.Int64{
											int64validator.Between(1, 65535),
										},
										MarkdownDescription: "Sequence number",
										Optional:            true,
									},
								},
							},
						},
					},
				},
				"large": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("extended"),
							path.MatchRelative().AtParent().AtName("regular"),
						),
					},
					MarkdownDescription: "Large\n\n> ℹ️ **Note:** You must specify exactly one of `extended`, `large`, and `regular`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"large_entry": schema.ListNestedAttribute{
							MarkdownDescription: "Large community lists",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"action": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("deny", "permit"),
										},
										MarkdownDescription: "Action",
										Optional:            true,
									},
									"lc_regex": schema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "Large community regular expression",
										Validators: []validator.List{
											listvalidator.SizeAtMost(8),
										},
										Optional: true,
									},
									"name": schema.Int64Attribute{
										Validators: []validator.Int64{
											int64validator.Between(1, 65535),
										},
										MarkdownDescription: "Sequence number",
										Optional:            true,
									},
								},
							},
						},
					},
				},
				"regular": schema.SingleNestedAttribute{
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(
							path.MatchRelative().AtParent().AtName("extended"),
							path.MatchRelative().AtParent().AtName("large"),
						),
					},
					MarkdownDescription: "Regular\n\n> ℹ️ **Note:** You must specify exactly one of `extended`, `large`, and `regular`.",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"regular_entry": schema.ListNestedAttribute{
							MarkdownDescription: "Regular community lists",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"action": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("deny", "permit"),
										},
										MarkdownDescription: "Action",
										Optional:            true,
									},
									"community": schema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "Communities",
										Optional:            true,
									},
									"name": schema.Int64Attribute{
										Validators: []validator.Int64{
											int64validator.Between(1, 65535),
										},
										MarkdownDescription: "Sequence number",
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
}

// RouteCommunityListsDataSourceSchema defines the schema for RouteCommunityLists data source
var RouteCommunityListsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "RouteCommunityList data source",
	Attributes: map[string]dsschema.Attribute{
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Description",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Route community list name",
			Optional:            true,
			Computed:            true,
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined\n\n> ℹ️ **Note:** You must specify exactly one of `device`, `folder`, and `snippet`.",
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
		"type": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Type",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"extended": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Extended\n\n> ℹ️ **Note:** You must specify exactly one of `extended`, `large`, and `regular`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"extended_entry": dsschema.ListNestedAttribute{
							MarkdownDescription: "Extended community lists",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"action": dsschema.StringAttribute{
										MarkdownDescription: "Action",
										Computed:            true,
									},
									"lc_regex": dsschema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "Extended community regular expression",
										Computed:            true,
									},
									"name": dsschema.Int64Attribute{
										MarkdownDescription: "Sequence number",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"large": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Large\n\n> ℹ️ **Note:** You must specify exactly one of `extended`, `large`, and `regular`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"large_entry": dsschema.ListNestedAttribute{
							MarkdownDescription: "Large community lists",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"action": dsschema.StringAttribute{
										MarkdownDescription: "Action",
										Computed:            true,
									},
									"lc_regex": dsschema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "Large community regular expression",
										Computed:            true,
									},
									"name": dsschema.Int64Attribute{
										MarkdownDescription: "Sequence number",
										Computed:            true,
									},
								},
							},
						},
					},
				},
				"regular": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Regular\n\n> ℹ️ **Note:** You must specify exactly one of `extended`, `large`, and `regular`.",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"regular_entry": dsschema.ListNestedAttribute{
							MarkdownDescription: "Regular community lists",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"action": dsschema.StringAttribute{
										MarkdownDescription: "Action",
										Computed:            true,
									},
									"community": dsschema.ListAttribute{
										ElementType:         types.StringType,
										MarkdownDescription: "Communities",
										Computed:            true,
									},
									"name": dsschema.Int64Attribute{
										MarkdownDescription: "Sequence number",
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
}

// RouteCommunityListsListModel represents the data model for a list data source.
type RouteCommunityListsListModel struct {
	Tfid    types.String          `tfsdk:"tfid"`
	Data    []RouteCommunityLists `tfsdk:"data"`
	Limit   types.Int64           `tfsdk:"limit"`
	Offset  types.Int64           `tfsdk:"offset"`
	Name    types.String          `tfsdk:"name"`
	Total   types.Int64           `tfsdk:"total"`
	Folder  types.String          `tfsdk:"folder"`
	Snippet types.String          `tfsdk:"snippet"`
	Device  types.String          `tfsdk:"device"`
}

// RouteCommunityListsListDataSourceSchema defines the schema for a list data source.
var RouteCommunityListsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: RouteCommunityListsDataSourceSchema.Attributes,
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
