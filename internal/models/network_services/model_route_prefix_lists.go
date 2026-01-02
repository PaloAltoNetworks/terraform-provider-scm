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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Package: network_services
// This file contains models for the network_services SDK package

// RoutePrefixLists represents the Terraform model for RoutePrefixLists
type RoutePrefixLists struct {
	Tfid        types.String          `tfsdk:"tfid"`
	Description basetypes.StringValue `tfsdk:"description"`
	Device      basetypes.StringValue `tfsdk:"device"`
	Folder      basetypes.StringValue `tfsdk:"folder"`
	Id          basetypes.StringValue `tfsdk:"id"`
	Name        basetypes.StringValue `tfsdk:"name"`
	Snippet     basetypes.StringValue `tfsdk:"snippet"`
	Type        basetypes.ObjectValue `tfsdk:"type"`
}

// RoutePrefixListsType represents a nested structure within the RoutePrefixLists model
type RoutePrefixListsType struct {
	Ipv4 basetypes.ObjectValue `tfsdk:"ipv4"`
}

// RoutePrefixListsTypeIpv4 represents a nested structure within the RoutePrefixLists model
type RoutePrefixListsTypeIpv4 struct {
	Ipv4Entry basetypes.ListValue `tfsdk:"ipv4_entry"`
}

// RoutePrefixListsTypeIpv4Ipv4EntryInner represents a nested structure within the RoutePrefixLists model
type RoutePrefixListsTypeIpv4Ipv4EntryInner struct {
	Action basetypes.StringValue `tfsdk:"action"`
	Name   basetypes.Int64Value  `tfsdk:"name"`
	Prefix basetypes.ObjectValue `tfsdk:"prefix"`
}

// RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix represents a nested structure within the RoutePrefixLists model
type RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix struct {
	Entry   basetypes.ObjectValue `tfsdk:"entry"`
	Network basetypes.StringValue `tfsdk:"network"`
}

// RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry represents a nested structure within the RoutePrefixLists model
type RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry struct {
	GreaterThanOrEqual basetypes.Int64Value  `tfsdk:"greater_than_or_equal"`
	LessThanOrEqual    basetypes.Int64Value  `tfsdk:"less_than_or_equal"`
	Network            basetypes.StringValue `tfsdk:"network"`
}

// AttrTypes defines the attribute types for the RoutePrefixLists model.
func (o RoutePrefixLists) AttrTypes() map[string]attr.Type {
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
				"ipv4": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ipv4_entry": basetypes.ListType{ElemType: basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action": basetypes.StringType{},
								"name":   basetypes.Int64Type{},
								"prefix": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"entry": basetypes.ObjectType{
											AttrTypes: map[string]attr.Type{
												"greater_than_or_equal": basetypes.Int64Type{},
												"less_than_or_equal":    basetypes.Int64Type{},
												"network":               basetypes.StringType{},
											},
										},
										"network": basetypes.StringType{},
									},
								},
							},
						}},
					},
				},
			},
		},
	}
}

// AttrType returns the attribute type for a list of RoutePrefixLists objects.
func (o RoutePrefixLists) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RoutePrefixListsType model.
func (o RoutePrefixListsType) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ipv4": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ipv4_entry": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action": basetypes.StringType{},
						"name":   basetypes.Int64Type{},
						"prefix": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"entry": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"greater_than_or_equal": basetypes.Int64Type{},
										"less_than_or_equal":    basetypes.Int64Type{},
										"network":               basetypes.StringType{},
									},
								},
								"network": basetypes.StringType{},
							},
						},
					},
				}},
			},
		},
	}
}

// AttrType returns the attribute type for a list of RoutePrefixListsType objects.
func (o RoutePrefixListsType) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RoutePrefixListsTypeIpv4 model.
func (o RoutePrefixListsTypeIpv4) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ipv4_entry": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action": basetypes.StringType{},
				"name":   basetypes.Int64Type{},
				"prefix": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"entry": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"greater_than_or_equal": basetypes.Int64Type{},
								"less_than_or_equal":    basetypes.Int64Type{},
								"network":               basetypes.StringType{},
							},
						},
						"network": basetypes.StringType{},
					},
				},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of RoutePrefixListsTypeIpv4 objects.
func (o RoutePrefixListsTypeIpv4) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RoutePrefixListsTypeIpv4Ipv4EntryInner model.
func (o RoutePrefixListsTypeIpv4Ipv4EntryInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action": basetypes.StringType{},
		"name":   basetypes.Int64Type{},
		"prefix": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"entry": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"greater_than_or_equal": basetypes.Int64Type{},
						"less_than_or_equal":    basetypes.Int64Type{},
						"network":               basetypes.StringType{},
					},
				},
				"network": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of RoutePrefixListsTypeIpv4Ipv4EntryInner objects.
func (o RoutePrefixListsTypeIpv4Ipv4EntryInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix model.
func (o RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"entry": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"greater_than_or_equal": basetypes.Int64Type{},
				"less_than_or_equal":    basetypes.Int64Type{},
				"network":               basetypes.StringType{},
			},
		},
		"network": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix objects.
func (o RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry model.
func (o RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"greater_than_or_equal": basetypes.Int64Type{},
		"less_than_or_equal":    basetypes.Int64Type{},
		"network":               basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry objects.
func (o RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// RoutePrefixListsResourceSchema defines the schema for RoutePrefixLists resource
var RoutePrefixListsResourceSchema = schema.Schema{
	MarkdownDescription: "RoutePrefixList resource",
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
			MarkdownDescription: "Filter prefix list name",
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
			MarkdownDescription: "Address Family Type",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"ipv4": schema.SingleNestedAttribute{
					MarkdownDescription: "Ipv4",
					Required:            true,
					Attributes: map[string]schema.Attribute{
						"ipv4_entry": schema.ListNestedAttribute{
							MarkdownDescription: "IPv4 prefix lists",
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
									"name": schema.Int64Attribute{
										Validators: []validator.Int64{
											int64validator.Between(1, 65535),
										},
										MarkdownDescription: "Sequence number",
										Optional:            true,
									},
									"prefix": schema.SingleNestedAttribute{
										MarkdownDescription: "Prefix",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"entry": schema.SingleNestedAttribute{
												Validators: []validator.Object{
													objectvalidator.ConflictsWith(
														path.MatchRelative().AtParent().AtName("network"),
													),
												},
												MarkdownDescription: "Entry\n\n> ℹ️ **Note:** You must specify exactly one of `entry` and `network`.",
												Optional:            true,
												Attributes: map[string]schema.Attribute{
													"greater_than_or_equal": schema.Int64Attribute{
														Validators: []validator.Int64{
															int64validator.Between(0, 32),
														},
														MarkdownDescription: "Greater than or equal to",
														Optional:            true,
													},
													"less_than_or_equal": schema.Int64Attribute{
														Validators: []validator.Int64{
															int64validator.Between(0, 32),
														},
														MarkdownDescription: "Less than or equal to",
														Optional:            true,
													},
													"network": schema.StringAttribute{
														MarkdownDescription: "Network",
														Optional:            true,
													},
												},
											},
											"network": schema.StringAttribute{
												Validators: []validator.String{
													stringvalidator.ConflictsWith(
														path.MatchRelative().AtParent().AtName("entry"),
													),
													stringvalidator.OneOf("any"),
												},
												MarkdownDescription: "Network\n\n> ℹ️ **Note:** You must specify exactly one of `entry` and `network`.",
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
}

// RoutePrefixListsDataSourceSchema defines the schema for RoutePrefixLists data source
var RoutePrefixListsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "RoutePrefixList data source",
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
			MarkdownDescription: "Filter prefix list name",
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
		"type": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Address Family Type",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"ipv4": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Ipv4",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"ipv4_entry": dsschema.ListNestedAttribute{
							MarkdownDescription: "IPv4 prefix lists",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"action": dsschema.StringAttribute{
										MarkdownDescription: "Action",
										Computed:            true,
									},
									"name": dsschema.Int64Attribute{
										MarkdownDescription: "Sequence number",
										Computed:            true,
									},
									"prefix": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Prefix",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"entry": dsschema.SingleNestedAttribute{
												MarkdownDescription: "Entry\n\n> ℹ️ **Note:** You must specify exactly one of `entry` and `network`.",
												Computed:            true,
												Attributes: map[string]dsschema.Attribute{
													"greater_than_or_equal": dsschema.Int64Attribute{
														MarkdownDescription: "Greater than or equal to",
														Computed:            true,
													},
													"less_than_or_equal": dsschema.Int64Attribute{
														MarkdownDescription: "Less than or equal to",
														Computed:            true,
													},
													"network": dsschema.StringAttribute{
														MarkdownDescription: "Network",
														Computed:            true,
													},
												},
											},
											"network": dsschema.StringAttribute{
												MarkdownDescription: "Network\n\n> ℹ️ **Note:** You must specify exactly one of `entry` and `network`.",
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
}

// RoutePrefixListsListModel represents the data model for a list data source.
type RoutePrefixListsListModel struct {
	Tfid    types.String       `tfsdk:"tfid"`
	Data    []RoutePrefixLists `tfsdk:"data"`
	Limit   types.Int64        `tfsdk:"limit"`
	Offset  types.Int64        `tfsdk:"offset"`
	Name    types.String       `tfsdk:"name"`
	Total   types.Int64        `tfsdk:"total"`
	Folder  types.String       `tfsdk:"folder"`
	Snippet types.String       `tfsdk:"snippet"`
	Device  types.String       `tfsdk:"device"`
}

// RoutePrefixListsListDataSourceSchema defines the schema for a list data source.
var RoutePrefixListsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: RoutePrefixListsDataSourceSchema.Attributes,
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
