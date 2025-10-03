package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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

// RouteAccessLists represents the Terraform model for RouteAccessLists
type RouteAccessLists struct {
	Tfid        types.String          `tfsdk:"tfid"`
	Description basetypes.StringValue `tfsdk:"description"`
	Device      basetypes.StringValue `tfsdk:"device"`
	Folder      basetypes.StringValue `tfsdk:"folder"`
	Id          basetypes.StringValue `tfsdk:"id"`
	Name        basetypes.StringValue `tfsdk:"name"`
	Snippet     basetypes.StringValue `tfsdk:"snippet"`
	Type        basetypes.ObjectValue `tfsdk:"type"`
}

// RouteAccessListsType represents a nested structure within the RouteAccessLists model
type RouteAccessListsType struct {
	Ipv4 basetypes.ObjectValue `tfsdk:"ipv4"`
}

// RouteAccessListsTypeIpv4 represents a nested structure within the RouteAccessLists model
type RouteAccessListsTypeIpv4 struct {
	Ipv4Entry basetypes.ListValue `tfsdk:"ipv4_entry"`
}

// RouteAccessListsTypeIpv4Ipv4EntryInner represents a nested structure within the RouteAccessLists model
type RouteAccessListsTypeIpv4Ipv4EntryInner struct {
	Action             basetypes.StringValue `tfsdk:"action"`
	DestinationAddress basetypes.ObjectValue `tfsdk:"destination_address"`
	Name               basetypes.Int64Value  `tfsdk:"name"`
	SourceAddress      basetypes.ObjectValue `tfsdk:"source_address"`
}

// RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress represents a nested structure within the RouteAccessLists model
type RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress struct {
	Address  basetypes.StringValue `tfsdk:"address"`
	Wildcard basetypes.StringValue `tfsdk:"wildcard"`
}

// RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress represents a nested structure within the RouteAccessLists model
type RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress struct {
	Address  basetypes.StringValue `tfsdk:"address"`
	Wildcard basetypes.StringValue `tfsdk:"wildcard"`
}

// AttrTypes defines the attribute types for the RouteAccessLists model.
func (o RouteAccessLists) AttrTypes() map[string]attr.Type {
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
								"destination_address": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"address":  basetypes.StringType{},
										"wildcard": basetypes.StringType{},
									},
								},
								"name": basetypes.Int64Type{},
								"source_address": basetypes.ObjectType{
									AttrTypes: map[string]attr.Type{
										"address":  basetypes.StringType{},
										"wildcard": basetypes.StringType{},
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

// AttrType returns the attribute type for a list of RouteAccessLists objects.
func (o RouteAccessLists) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RouteAccessListsType model.
func (o RouteAccessListsType) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ipv4": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ipv4_entry": basetypes.ListType{ElemType: basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"action": basetypes.StringType{},
						"destination_address": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address":  basetypes.StringType{},
								"wildcard": basetypes.StringType{},
							},
						},
						"name": basetypes.Int64Type{},
						"source_address": basetypes.ObjectType{
							AttrTypes: map[string]attr.Type{
								"address":  basetypes.StringType{},
								"wildcard": basetypes.StringType{},
							},
						},
					},
				}},
			},
		},
	}
}

// AttrType returns the attribute type for a list of RouteAccessListsType objects.
func (o RouteAccessListsType) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RouteAccessListsTypeIpv4 model.
func (o RouteAccessListsTypeIpv4) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"ipv4_entry": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action": basetypes.StringType{},
				"destination_address": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address":  basetypes.StringType{},
						"wildcard": basetypes.StringType{},
					},
				},
				"name": basetypes.Int64Type{},
				"source_address": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"address":  basetypes.StringType{},
						"wildcard": basetypes.StringType{},
					},
				},
			},
		}},
	}
}

// AttrType returns the attribute type for a list of RouteAccessListsTypeIpv4 objects.
func (o RouteAccessListsTypeIpv4) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RouteAccessListsTypeIpv4Ipv4EntryInner model.
func (o RouteAccessListsTypeIpv4Ipv4EntryInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action": basetypes.StringType{},
		"destination_address": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address":  basetypes.StringType{},
				"wildcard": basetypes.StringType{},
			},
		},
		"name": basetypes.Int64Type{},
		"source_address": basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"address":  basetypes.StringType{},
				"wildcard": basetypes.StringType{},
			},
		},
	}
}

// AttrType returns the attribute type for a list of RouteAccessListsTypeIpv4Ipv4EntryInner objects.
func (o RouteAccessListsTypeIpv4Ipv4EntryInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress model.
func (o RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address":  basetypes.StringType{},
		"wildcard": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress objects.
func (o RouteAccessListsTypeIpv4Ipv4EntryInnerDestinationAddress) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress model.
func (o RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"address":  basetypes.StringType{},
		"wildcard": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress objects.
func (o RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// RouteAccessListsResourceSchema defines the schema for RouteAccessLists resource
var RouteAccessListsResourceSchema = schema.Schema{
	MarkdownDescription: "Manages SCM RouteAccessLists objects",
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
			MarkdownDescription: "Route access list name",
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
		"type": schema.SingleNestedAttribute{
			MarkdownDescription: "Type",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"ipv4": schema.SingleNestedAttribute{
					MarkdownDescription: "Ipv4",
					Optional:            true,
					Attributes: map[string]schema.Attribute{
						"ipv4_entry": schema.ListNestedAttribute{
							MarkdownDescription: "IPv4 access lists",
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
									"destination_address": schema.SingleNestedAttribute{
										MarkdownDescription: "Destination address",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"address": schema.StringAttribute{
												Validators: []validator.String{
													stringvalidator.ExactlyOneOf(
														path.MatchRelative().AtParent().AtName("wildcard"),
														path.MatchRelative().AtParent().AtName("entry"),
													),
												},
												MarkdownDescription: "Destination IP address",
												Optional:            true,
											},
											"wildcard": schema.StringAttribute{
												Validators: []validator.String{
													stringvalidator.ExactlyOneOf(
														path.MatchRelative().AtParent().AtName("address"),
														path.MatchRelative().AtParent().AtName("entry"),
													),
												},
												MarkdownDescription: "Destination IP wildcard",
												Optional:            true,
											},
										},
									},
									"name": schema.Int64Attribute{
										Validators: []validator.Int64{
											int64validator.Between(1, 65535),
										},
										MarkdownDescription: "Sequence number",
										Optional:            true,
									},
									"source_address": schema.SingleNestedAttribute{
										MarkdownDescription: "Source address",
										Optional:            true,
										Attributes: map[string]schema.Attribute{
											"address": schema.StringAttribute{
												Validators: []validator.String{
													stringvalidator.ExactlyOneOf(
														path.MatchRelative().AtParent().AtName("wildcard"),
														path.MatchRelative().AtParent().AtName("entry"),
													),
												},
												MarkdownDescription: "Source IP address",
												Optional:            true,
											},
											"wildcard": schema.StringAttribute{
												Validators: []validator.String{
													stringvalidator.ExactlyOneOf(
														path.MatchRelative().AtParent().AtName("address"),
														path.MatchRelative().AtParent().AtName("entry"),
													),
												},
												MarkdownDescription: "Source IP wildcard",
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

// RouteAccessListsDataSourceSchema defines the schema for RouteAccessLists data source
var RouteAccessListsDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "RouteAccessLists data source",
	Attributes: map[string]dsschema.Attribute{
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
			MarkdownDescription: "UUID of the resource",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "Route access list name",
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
		"type": dsschema.SingleNestedAttribute{
			MarkdownDescription: "Type",
			Computed:            true,
			Attributes: map[string]dsschema.Attribute{
				"ipv4": dsschema.SingleNestedAttribute{
					MarkdownDescription: "Ipv4",
					Computed:            true,
					Attributes: map[string]dsschema.Attribute{
						"ipv4_entry": dsschema.ListNestedAttribute{
							MarkdownDescription: "IPv4 access lists",
							Computed:            true,
							NestedObject: dsschema.NestedAttributeObject{
								Attributes: map[string]dsschema.Attribute{
									"action": dsschema.StringAttribute{
										MarkdownDescription: "Action",
										Computed:            true,
									},
									"destination_address": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Destination address",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"address": dsschema.StringAttribute{
												MarkdownDescription: "Destination IP address",
												Computed:            true,
											},
											"wildcard": dsschema.StringAttribute{
												MarkdownDescription: "Destination IP wildcard",
												Computed:            true,
											},
										},
									},
									"name": dsschema.Int64Attribute{
										MarkdownDescription: "Sequence number",
										Computed:            true,
									},
									"source_address": dsschema.SingleNestedAttribute{
										MarkdownDescription: "Source address",
										Computed:            true,
										Attributes: map[string]dsschema.Attribute{
											"address": dsschema.StringAttribute{
												MarkdownDescription: "Source IP address",
												Computed:            true,
											},
											"wildcard": dsschema.StringAttribute{
												MarkdownDescription: "Source IP wildcard",
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

// RouteAccessListsListModel represents the data model for a list data source.
type RouteAccessListsListModel struct {
	Tfid    types.String       `tfsdk:"tfid"`
	Data    []RouteAccessLists `tfsdk:"data"`
	Limit   types.Int64        `tfsdk:"limit"`
	Offset  types.Int64        `tfsdk:"offset"`
	Name    types.String       `tfsdk:"name"`
	Total   types.Int64        `tfsdk:"total"`
	Folder  types.String       `tfsdk:"folder"`
	Snippet types.String       `tfsdk:"snippet"`
	Device  types.String       `tfsdk:"device"`
}

// RouteAccessListsListDataSourceSchema defines the schema for a list data source.
var RouteAccessListsListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: RouteAccessListsDataSourceSchema.Attributes,
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
